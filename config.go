package chisel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	Version string
	Auth    string
	Server  string
	Remotes []*Remote
}

const pre = "chisel-"

func DecodeConfig(s string) (*Config, error) {
	if !strings.HasPrefix(s, pre) {
		return nil, fmt.Errorf("Invalid config")
	}
	s = strings.TrimPrefix(s, pre)
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("Invalid base64 config")
	}
	c := &Config{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON config")
	}
	return c, nil
}

func EncodeConfig(c *Config) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return pre + base64.StdEncoding.EncodeToString(b), nil
}