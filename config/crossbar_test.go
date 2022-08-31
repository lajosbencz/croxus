package config_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/lajosbencz/croxus/config"
)

func TestCrossbar(t *testing.T) {
	jsonCfg, err := os.ReadFile("../crossbar/config.json")
	if err != nil {
		panic(err)
	}
	cfg := &config.CrossbarConfig{}
	json.Unmarshal(jsonCfg, cfg)
	fmt.Printf("%#v\n", cfg)
	if cfg.Version != 2 {
		t.Error("version should equal 2")
	}
}
