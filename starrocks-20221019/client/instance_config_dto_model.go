// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceConfigDto interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *InstanceConfigDto
	GetConfigKey() *string
	SetConfigType(v string) *InstanceConfigDto
	GetConfigType() *string
	SetConfigValue(v string) *InstanceConfigDto
	GetConfigValue() *string
	SetNodeGroupId(v string) *InstanceConfigDto
	GetNodeGroupId() *string
}

type InstanceConfigDto struct {
	// example:
	//
	// storage_root_path
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// example:
	//
	// BE
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// example:
	//
	// value1
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
	// example:
	//
	// ng-e24924dxxxxx
	NodeGroupId *string `json:"nodeGroupId,omitempty" xml:"nodeGroupId,omitempty"`
}

func (s InstanceConfigDto) String() string {
	return dara.Prettify(s)
}

func (s InstanceConfigDto) GoString() string {
	return s.String()
}

func (s *InstanceConfigDto) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *InstanceConfigDto) GetConfigType() *string {
	return s.ConfigType
}

func (s *InstanceConfigDto) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *InstanceConfigDto) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *InstanceConfigDto) SetConfigKey(v string) *InstanceConfigDto {
	s.ConfigKey = &v
	return s
}

func (s *InstanceConfigDto) SetConfigType(v string) *InstanceConfigDto {
	s.ConfigType = &v
	return s
}

func (s *InstanceConfigDto) SetConfigValue(v string) *InstanceConfigDto {
	s.ConfigValue = &v
	return s
}

func (s *InstanceConfigDto) SetNodeGroupId(v string) *InstanceConfigDto {
	s.NodeGroupId = &v
	return s
}

func (s *InstanceConfigDto) Validate() error {
	return dara.Validate(s)
}
