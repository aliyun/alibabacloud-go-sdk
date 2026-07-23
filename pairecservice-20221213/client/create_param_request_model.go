// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *CreateParamRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CreateParamRequest
	GetInstanceId() *string
	SetName(v string) *CreateParamRequest
	GetName() *string
	SetSceneId(v string) *CreateParamRequest
	GetSceneId() *string
	SetType(v string) *CreateParamRequest
	GetType() *string
	SetValue(v string) *CreateParamRequest
	GetValue() *string
}

type CreateParamRequest struct {
	// The environment to which the parameter belongs. Valid values:
	//
	// - Daily: daily environment.
	//
	// - Pre: staging environment.
	//
	// - Prod: production environment.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. You can call the ListInstances operation to obtain the instance ID.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scene ID. You can call the ListScenes operation to obtain the scene ID.
	//
	// example:
	//
	// 4
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The parameter type. Valid values:
	//
	// - Normal: common parameter.
	//
	// - Encryption: encrypted parameter.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParamRequest) GoString() string {
	return s.String()
}

func (s *CreateParamRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CreateParamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateParamRequest) GetName() *string {
	return s.Name
}

func (s *CreateParamRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateParamRequest) GetType() *string {
	return s.Type
}

func (s *CreateParamRequest) GetValue() *string {
	return s.Value
}

func (s *CreateParamRequest) SetEnvironment(v string) *CreateParamRequest {
	s.Environment = &v
	return s
}

func (s *CreateParamRequest) SetInstanceId(v string) *CreateParamRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateParamRequest) SetName(v string) *CreateParamRequest {
	s.Name = &v
	return s
}

func (s *CreateParamRequest) SetSceneId(v string) *CreateParamRequest {
	s.SceneId = &v
	return s
}

func (s *CreateParamRequest) SetType(v string) *CreateParamRequest {
	s.Type = &v
	return s
}

func (s *CreateParamRequest) SetValue(v string) *CreateParamRequest {
	s.Value = &v
	return s
}

func (s *CreateParamRequest) Validate() error {
	return dara.Validate(s)
}
