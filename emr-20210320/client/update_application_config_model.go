// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfigDescription(v string) *UpdateApplicationConfig
	GetConfigDescription() *string
	SetConfigFileName(v string) *UpdateApplicationConfig
	GetConfigFileName() *string
	SetConfigItemKey(v string) *UpdateApplicationConfig
	GetConfigItemKey() *string
	SetConfigItemValue(v string) *UpdateApplicationConfig
	GetConfigItemValue() *string
}

type UpdateApplicationConfig struct {
	// A description of the modification.
	//
	// example:
	//
	// dfs.namenode.checkpoint.period
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	// The name of the application configuration file.
	//
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// The configuration item key.
	//
	// example:
	//
	// dfs.namenode.checkpoint.period
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// The configuration item value.
	//
	// example:
	//
	// 3600s
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
}

func (s UpdateApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationConfig) GetConfigDescription() *string {
	return s.ConfigDescription
}

func (s *UpdateApplicationConfig) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *UpdateApplicationConfig) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *UpdateApplicationConfig) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *UpdateApplicationConfig) SetConfigDescription(v string) *UpdateApplicationConfig {
	s.ConfigDescription = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigFileName(v string) *UpdateApplicationConfig {
	s.ConfigFileName = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigItemKey(v string) *UpdateApplicationConfig {
	s.ConfigItemKey = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigItemValue(v string) *UpdateApplicationConfig {
	s.ConfigItemValue = &v
	return s
}

func (s *UpdateApplicationConfig) Validate() error {
	return dara.Validate(s)
}
