// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *UpdateEngineConfigRequest
	GetConfigValue() *string
	SetDescription(v string) *UpdateEngineConfigRequest
	GetDescription() *string
	SetEnvironment(v string) *UpdateEngineConfigRequest
	GetEnvironment() *string
	SetInstanceId(v string) *UpdateEngineConfigRequest
	GetInstanceId() *string
	SetName(v string) *UpdateEngineConfigRequest
	GetName() *string
	SetType(v string) *UpdateEngineConfigRequest
	GetType() *string
}

type UpdateEngineConfigRequest struct {
	// The content of the engine configuration.
	//
	// example:
	//
	// {
	//
	// 	"ListenConf": {
	//
	// 		"HttpAddr": "",
	//
	// 		"HttpPort": 8000
	//
	// 	}
	//
	// }
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description.
	//
	// example:
	//
	// update config
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The runtime environment.
	//
	// Valid values:
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
	// The instance ID. For information about how to obtain the instance ID, see [ListInstances](https://help.aliyun.com/document_detail/2411819.html).
	//
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The engine configuration name.
	//
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The engine configuration type.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateEngineConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEngineConfigRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *UpdateEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEngineConfigRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEngineConfigRequest) GetType() *string {
	return s.Type
}

func (s *UpdateEngineConfigRequest) SetConfigValue(v string) *UpdateEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetDescription(v string) *UpdateEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetEnvironment(v string) *UpdateEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetInstanceId(v string) *UpdateEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetName(v string) *UpdateEngineConfigRequest {
	s.Name = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetType(v string) *UpdateEngineConfigRequest {
	s.Type = &v
	return s
}

func (s *UpdateEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
