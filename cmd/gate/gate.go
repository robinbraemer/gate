/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gate

import (
	"fmt"
	"github.com/spf13/viper"
	"go.minekube.com/gate/pkg/config"
	"go.minekube.com/gate/pkg/proxy"
	"go.uber.org/zap"
)

func Run() (err error) {
	var file config.File
	if err := viper.Unmarshal(&file); err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	if err := initLogger(file.Debug); err != nil {
		return fmt.Errorf("error initializing global logger: %w", err)
	}

	cfg, err := config.NewValid(&file)
	if err != nil {
		return fmt.Errorf("error validating config: %w", err)
	}

	return proxy.NewProxy(cfg).Run()
}

func initLogger(debug bool) (err error) {
	var l *zap.Logger
	if debug {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}
	zap.ReplaceGlobals(l)
	return nil
}