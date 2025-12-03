// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVariableGroupRequest
	GetDescription() *string
	SetName(v string) *CreateVariableGroupRequest
	GetName() *string
	SetVariables(v string) *CreateVariableGroupRequest
	GetVariables() *string
}

type CreateVariableGroupRequest struct {
	// example:
	//
	// 变量组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 变量组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"isEncrypted":true,"name":"name1","value":"vaue1"}]
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateVariableGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVariableGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateVariableGroupRequest) GetVariables() *string {
	return s.Variables
}

func (s *CreateVariableGroupRequest) SetDescription(v string) *CreateVariableGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateVariableGroupRequest) SetName(v string) *CreateVariableGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateVariableGroupRequest) SetVariables(v string) *CreateVariableGroupRequest {
	s.Variables = &v
	return s
}

func (s *CreateVariableGroupRequest) Validate() error {
	return dara.Validate(s)
}
