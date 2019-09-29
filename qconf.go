package qconf

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	i := 0
	for scanner.Scan() {
		var line string
		if i == 0 {
			b := scanner.Bytes()
			if b[0] == 0xef && b[1] == 0xbb && b[2] == 0xbf {
				b = b[3:]
			}
			line = string(b) // scanner.Text()
		} else {
			line = scanner.Text()
		}
		i++
		line = strings.TrimSpace(line)

		if line == "" || strings.Index(line, "#") == 0 {
			continue
		}
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

func (conf *Config) GetString(key string) string {
	value := conf.Get(key)
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func (conf *Config) GetInteger(key string) (int64, error) {
	value := conf.Get(key)
	if value == nil {
		return 0, errors.New("it is not a Integer")
	}
	b, err := strconv.ParseInt(fmt.Sprint(value), 0, 64)

	if err != nil {
		return 0, err
	}
	return b, nil
}

func (conf *Config) GetBoolean(key string) (bool, error) {
	value := conf.Get(key)
	if value == nil {
		return false, errors.New("there is no key named:" + key)
	}
	b, ok := value.(bool)
	if !ok {
		return false, errors.New("it is not a bool")
	}
	return b, nil
}

func (conf *Config) Get(key string) interface{} {
	if v, ok := conf.kv[strings.ToLower(key)]; ok {
		return v
	} else {
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

func (conf *Config)  Save(){

}
