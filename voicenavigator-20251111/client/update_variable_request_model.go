// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVariableRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateVariableRequest
	GetDisplayName() *string
	SetInstanceId(v string) *UpdateVariableRequest
	GetInstanceId() *string
	SetVariableId(v string) *UpdateVariableRequest
	GetVariableId() *string
}

type UpdateVariableRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 026ca0f4-483b-4252-ae1d-1f15f056f8b9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	VariableId *string `json:"VariableId,omitempty" xml:"VariableId,omitempty"`
}

func (s UpdateVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableRequest) GoString() string {
	return s.String()
}

func (s *UpdateVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVariableRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateVariableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateVariableRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *UpdateVariableRequest) SetDescription(v string) *UpdateVariableRequest {
	s.Description = &v
	return s
}

func (s *UpdateVariableRequest) SetDisplayName(v string) *UpdateVariableRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateVariableRequest) SetInstanceId(v string) *UpdateVariableRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateVariableRequest) SetVariableId(v string) *UpdateVariableRequest {
	s.VariableId = &v
	return s
}

func (s *UpdateVariableRequest) Validate() error {
	return dara.Validate(s)
}
