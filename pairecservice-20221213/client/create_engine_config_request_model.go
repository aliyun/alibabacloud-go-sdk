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
	SetType(v string) *CreateEngineConfigRequest
	GetType() *string
}

type CreateEngineConfigRequest struct {
	// The content of the engine configuration.
	//
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description.
	//
	// example:
	//
	// this is a test config
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The runtime environment. Valid values:
	//
	// - Daily: daily environment.
	//
	// - Pre: staging environment.
	//
	// - Prod: production environment.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. You can obtain the ID from the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation.
	//
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the engine configuration.
	//
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the engine configuration.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *CreateEngineConfigRequest) GetType() *string {
	return s.Type
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

func (s *CreateEngineConfigRequest) SetType(v string) *CreateEngineConfigRequest {
	s.Type = &v
	return s
}

func (s *CreateEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
