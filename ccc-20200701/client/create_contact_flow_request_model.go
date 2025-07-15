// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateContactFlowRequest
	GetDefinition() *string
	SetDescription(v string) *CreateContactFlowRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateContactFlowRequest
	GetInstanceId() *string
	SetName(v string) *CreateContactFlowRequest
	GetName() *string
	SetType(v string) *CreateContactFlowRequest
	GetType() *string
}

type CreateContactFlowRequest struct {
	// This parameter is required.
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAIN_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateContactFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateContactFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateContactFlowRequest) GetName() *string {
	return s.Name
}

func (s *CreateContactFlowRequest) GetType() *string {
	return s.Type
}

func (s *CreateContactFlowRequest) SetDefinition(v string) *CreateContactFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateContactFlowRequest) SetDescription(v string) *CreateContactFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateContactFlowRequest) SetInstanceId(v string) *CreateContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateContactFlowRequest) SetName(v string) *CreateContactFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateContactFlowRequest) SetType(v string) *CreateContactFlowRequest {
	s.Type = &v
	return s
}

func (s *CreateContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
