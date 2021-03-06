// Copyright 2019 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"time"

	"github.com/minio/kes"
	"gopkg.in/yaml.v2"
)

type serverConfig struct {
	Addr string       `yaml:"address"`
	Root kes.Identity `yaml:"root"`

	TLS struct {
		KeyPath  string `yaml:"key"`
		CertPath string `yaml:"cert"`
		Proxy    struct {
			Identities []kes.Identity `yaml:"identities"`
			Header     struct {
				ClientCert string `yaml:"cert"`
			} `yaml:"header"`
		} `yaml:"proxy"`
	} `yaml:"tls"`

	Policies map[string]struct {
		Paths      []string       `yaml:"paths"`
		Identities []kes.Identity `yaml:"identities"`
	} `yaml:"policy"`

	Cache struct {
		Expiry struct {
			Any    time.Duration `yaml:"any"`
			Unused time.Duration `yaml:"unused"`
		} `yaml:"expiry"`
	} `yaml:"cache"`

	Log struct {
		Error string `yaml:"error"`
		Audit string `yaml:"audit"`
	} `yaml:"log"`

	Keys struct {
		Fs struct {
			Path string `yaml:"path"`
		} `yaml:"fs"`

		Vault struct {
			Endpoint  string `yaml:"endpoint"`
			Namespace string `yaml:"namespace"`

			Prefix string `yaml:"prefix"`

			AppRole struct {
				ID     string        `yaml:"id"`
				Secret string        `yaml:"secret"`
				Retry  time.Duration `yaml:"retry"`
			} `yaml:"approle"`

			TLS struct {
				KeyPath  string `yaml:"key"`
				CertPath string `yaml:"cert"`
				CAPath   string `yaml:"ca"`
			} `yaml:"tls"`

			Status struct {
				Ping time.Duration `yaml:"ping"`
			} `yaml:"status"`
		} `yaml:"vault"`

		Aws struct {
			SecretsManager struct {
				Endpoint string `yaml:"endpoint"`
				Region   string `yaml:"region"`
				KmsKey   string ` yaml:"kmskey"`

				Login struct {
					AccessKey    string `yaml:"accesskey"`
					SecretKey    string `yaml:"secretkey"`
					SessionToken string `yaml:"token"`
				} `yaml:"credentials"`
			} `yaml:"secretsmanager"`
		} `yaml:"aws"`
	} `yaml:"keys"`
}

func loadServerConfig(path string) (config serverConfig, err error) {
	if path == "" {
		return config, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		file.Close()
		return config, err
	}
	return config, file.Close()
}
