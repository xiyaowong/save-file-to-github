// Package main provides ...
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

var (
	Owner string
	Repo  string
	Token string
)

func InitConfig() {
	exeDir, err := GetExeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg, err := ini.ShadowLoad(filepath.Join(exeDir, "config.ini"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sec, _ := cfg.GetSection(ini.DEFAULT_SECTION)
	Owner = sec.Key("owner").String()
	Repo = sec.Key("repo").String()
	Token = sec.Key("token").String()
	if Owner == "" || Repo == "" || Token == "" {
		fmt.Println("请配置")
		os.Exit(1)
	}
}
