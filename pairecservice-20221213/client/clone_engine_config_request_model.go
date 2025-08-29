// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *CloneEngineConfigRequest
	GetConfigValue() *string
	SetDescription(v string) *CloneEngineConfigRequest
	GetDescription() *string
	SetEnvironment(v string) *CloneEngineConfigRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CloneEngineConfigRequest
	GetInstanceId() *string
}

type CloneEngineConfigRequest struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *CloneEngineConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *CloneEngineConfigRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CloneEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneEngineConfigRequest) SetConfigValue(v string) *CloneEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CloneEngineConfigRequest) SetDescription(v string) *CloneEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *CloneEngineConfigRequest) SetEnvironment(v string) *CloneEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *CloneEngineConfigRequest) SetInstanceId(v string) *CloneEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
