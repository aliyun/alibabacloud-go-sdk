// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ApplicationConfig
	GetApplicationName() *string
	SetConfigFileName(v string) *ApplicationConfig
	GetConfigFileName() *string
	SetConfigItemKey(v string) *ApplicationConfig
	GetConfigItemKey() *string
	SetConfigItemValue(v string) *ApplicationConfig
	GetConfigItemValue() *string
	SetConfigScope(v string) *ApplicationConfig
	GetConfigScope() *string
	SetNodeGroupId(v string) *ApplicationConfig
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *ApplicationConfig
	GetNodeGroupName() *string
}

type ApplicationConfig struct {
	// The name of the application. You can view the application names of each EMR version on the cluster creation page in the EMR console.
	//
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the configuration file.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// The key of the configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// dfs.namenode.checkpoint.period
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 3600s
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	// The level at which you want to apply the configurations. Valid values:
	//
	// 	- CLUSTER
	//
	// 	- NODE_GROUP
	//
	// Default value: CLUSTER.
	//
	// example:
	//
	// NODE_GROUP
	ConfigScope *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	// The node group ID. This parameter takes effect only when the ConfigScope parameter is set to NODE_GROUP. The NodeGroupId parameter has a higher priority than the NodeGroupName parameter.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The name of the node group. This parameter takes effect only when the ConfigScope parameter is set to NODE_GROUP and the NodeGroupId parameter is not configured.
	//
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
}

func (s ApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s ApplicationConfig) GoString() string {
	return s.String()
}

func (s *ApplicationConfig) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ApplicationConfig) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ApplicationConfig) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ApplicationConfig) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ApplicationConfig) GetConfigScope() *string {
	return s.ConfigScope
}

func (s *ApplicationConfig) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ApplicationConfig) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ApplicationConfig) SetApplicationName(v string) *ApplicationConfig {
	s.ApplicationName = &v
	return s
}

func (s *ApplicationConfig) SetConfigFileName(v string) *ApplicationConfig {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfig) SetConfigItemKey(v string) *ApplicationConfig {
	s.ConfigItemKey = &v
	return s
}

func (s *ApplicationConfig) SetConfigItemValue(v string) *ApplicationConfig {
	s.ConfigItemValue = &v
	return s
}

func (s *ApplicationConfig) SetConfigScope(v string) *ApplicationConfig {
	s.ConfigScope = &v
	return s
}

func (s *ApplicationConfig) SetNodeGroupId(v string) *ApplicationConfig {
	s.NodeGroupId = &v
	return s
}

func (s *ApplicationConfig) SetNodeGroupName(v string) *ApplicationConfig {
	s.NodeGroupName = &v
	return s
}

func (s *ApplicationConfig) Validate() error {
	return dara.Validate(s)
}
