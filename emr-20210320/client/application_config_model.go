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
	// 应用名称。从EMR控制台集群创建页面可查看到指定发行版的应用名称列表。
	//
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 应用配置文件名。
	//
	// This parameter is required.
	//
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// 配置项键。
	//
	// This parameter is required.
	//
	// example:
	//
	// dfs.namenode.checkpoint.period
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// 配置项值。
	//
	// example:
	//
	// 3600s
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	// 配置范围。取值范围：
	//
	// - CLUSTER：集群级别。
	//
	// - NODE_GROUP：节点组级别。
	//
	// 默认值：CLUSTER。
	//
	// example:
	//
	// NODE_GROUP
	ConfigScope *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	// 节点组ID。ConfigScope取值NODE_GROUP时，该参数生效。NodeGroupId参数优先级高于NodeGroupName。
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// 节点组名称。ConfigScope取值NODE_GROUP时，且参数NodeGroupId为空时生效，该参数生效。
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
