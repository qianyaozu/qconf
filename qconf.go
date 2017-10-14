package qconf

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type Config struct {
	kv map[string]interface{}
}

func LoadConfiguration(path string) (*Config, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := new(Config)
	kv := make(map[string]interface{})
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		index := strings.Index(line, "=")
		if index <= 0 {
			return nil, errors.New(" configuration format requires an equals between the key and value")
		}
		key := strings.ToLower(strings.TrimSpace(line[0:index]))
		value := strings.TrimSpace(line[index+1:])
		value = strings.Trim(value, "\"'") //clear quotes
		kv[key] = value
	}
	config.kv = kv
	return config, scanner.Err()
}

func (conf *Config) getString(key string) string {
	value := conf.get(key)
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func (conf *Config) getInteger(key string) (int64, error) {
	value := conf.get(key)
	if value == nil {
		return 0, errors.New("it is not a Integer")
	}
	b, err := strconv.ParseInt(fmt.Sprint(value), 0, 64)

	if err != nil {
		return 0, err
	}
	return b, nil
}

func (conf *Config) getBoolean(key string) (bool, error) {
	value := conf.get(key)
	if value == nil {
		return false, errors.New("it is not a boolean")
	}
	b, ok := value.(bool)
	if !ok {
		return false, errors.New("it is not a bool")
	}
	return b, nil
}

func (conf *Config) get(key string) interface{} {
	if v, ok := conf.kv[strings.ToLower(key)]; ok {
		return v
	} else {
		fmt.Println("not ok")
		return nil
	}

}
func (conf *Config) Keys() []string {
	list := make([]string, 0)
	for k, _ := range conf.kv {
		list = append(list, k)
	}
	return list
}
