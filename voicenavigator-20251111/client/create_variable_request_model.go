// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVariableRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateVariableRequest
	GetDisplayName() *string
	SetInstanceId(v string) *CreateVariableRequest
	GetInstanceId() *string
	SetName(v string) *CreateVariableRequest
	GetName() *string
}

type CreateVariableRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 5daac920-d6c1-429f-a95f-2a798f5255b5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// userType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVariableRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateVariableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVariableRequest) GetName() *string {
	return s.Name
}

func (s *CreateVariableRequest) SetDescription(v string) *CreateVariableRequest {
	s.Description = &v
	return s
}

func (s *CreateVariableRequest) SetDisplayName(v string) *CreateVariableRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateVariableRequest) SetInstanceId(v string) *CreateVariableRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVariableRequest) SetName(v string) *CreateVariableRequest {
	s.Name = &v
	return s
}

func (s *CreateVariableRequest) Validate() error {
	return dara.Validate(s)
}
