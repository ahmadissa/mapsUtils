package Utils

import (
	"strconv"
	"strings"
	"time"
)

func GetString(obj map[string]interface{}, key string) (string, bool) {
	switch obj[key].(type) {
	case string:
		return obj[key].(string), true
	}
	return "", false
}

func GetTime(obj map[string]interface{}, key string) (time.Time, bool) {
	switch obj[key].(type) {
	case time.Time:
		return obj[key].(time.Time), true
	}
	return time.Time{}, false
}
func GetBool(obj map[string]interface{}, key string) (bool, bool) {
	switch obj[key].(type) {
	case bool:
		return obj[key].(bool), true
	}
	return false, false
}

//StrToInt
func StrToInt(v string) (int, error) {
	v = strings.TrimSpace(v)
	res, e := strconv.Atoi(v)
	if e != nil {
		return 0, e
	}
	return res, nil
}

//StrToInt64
func StrToInt64(v string) (int64, error) {
	v = strings.TrimSpace(v)
	res, e := strconv.ParseInt(v, 10, 64)
	if e != nil {
		return 0, e
	}
	return res, nil
}

func GetInt(obj map[string]interface{}, key string) (int, bool) {
	switch obj[key].(type) {
	case int:
		return obj[key].(int), true
	}
	return 0, false
}
func GetFloat64(obj map[string]interface{}, key string) (float64, bool) {
	switch obj[key].(type) {
	case float64:
		return obj[key].(float64), true
	}
	return 0, false
}
func GetInt64(obj map[string]interface{}, key string) (int64, bool) {
	switch obj[key].(type) {
	case int64:
		return obj[key].(int64), true
	}
	return 0, false
}
