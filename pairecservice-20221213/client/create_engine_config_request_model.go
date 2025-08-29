// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *CreateEngineConfigRequest
	GetConfigValue() *string
	SetDescription(v string) *CreateEngineConfigRequest
	GetDescription() *string
	SetEnvironment(v string) *CreateEngineConfigRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CreateEngineConfigRequest
	GetInstanceId() *string
	SetName(v string) *CreateEngineConfigRequest
	GetName() *string
}

type CreateEngineConfigRequest struct {
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
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *CreateEngineConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEngineConfigRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CreateEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEngineConfigRequest) GetName() *string {
	return s.Name
}

func (s *CreateEngineConfigRequest) SetConfigValue(v string) *CreateEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CreateEngineConfigRequest) SetDescription(v string) *CreateEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateEngineConfigRequest) SetEnvironment(v string) *CreateEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *CreateEngineConfigRequest) SetInstanceId(v string) *CreateEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEngineConfigRequest) SetName(v string) *CreateEngineConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
