// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *ClusterConfig
	GetConfigValue() *string
	SetEnableOverwrite(v bool) *ClusterConfig
	GetEnableOverwrite() *bool
	SetModuleName(v string) *ClusterConfig
	GetModuleName() *string
}

type ClusterConfig struct {
	// The configuration value.
	//
	// example:
	//
	// {"spark.driver.memory":"1g"}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// Specifies whether to overwrite the advanced settings of nodes in DataStudio. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableOverwrite *bool `json:"EnableOverwrite,omitempty" xml:"EnableOverwrite,omitempty"`
	// The module in which the cluster is configured. Valid values:
	//
	// 	- ide: DataStudio.
	//
	// 	- da: DataAnalysis.
	//
	// 	- scheduler.auto: Operation Center - auto triggered instances.
	//
	// 	- scheduler.backfill: Operation Center - data backfill instances.
	//
	// 	- scheduler.test: Operation Center - test instances.
	//
	// 	- scheduler.manual: Operation Center - manually triggered instances.
	//
	// example:
	//
	// ide
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s ClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s ClusterConfig) GoString() string {
	return s.String()
}

func (s *ClusterConfig) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ClusterConfig) GetEnableOverwrite() *bool {
	return s.EnableOverwrite
}

func (s *ClusterConfig) GetModuleName() *string {
	return s.ModuleName
}

func (s *ClusterConfig) SetConfigValue(v string) *ClusterConfig {
	s.ConfigValue = &v
	return s
}

func (s *ClusterConfig) SetEnableOverwrite(v bool) *ClusterConfig {
	s.EnableOverwrite = &v
	return s
}

func (s *ClusterConfig) SetModuleName(v string) *ClusterConfig {
	s.ModuleName = &v
	return s
}

func (s *ClusterConfig) Validate() error {
	return dara.Validate(s)
}
