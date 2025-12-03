// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVariableGroupRequest
	GetDescription() *string
	SetName(v string) *UpdateVariableGroupRequest
	GetName() *string
	SetVariables(v string) *UpdateVariableGroupRequest
	GetVariables() *string
}

type UpdateVariableGroupRequest struct {
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

func (s UpdateVariableGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVariableGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateVariableGroupRequest) GetVariables() *string {
	return s.Variables
}

func (s *UpdateVariableGroupRequest) SetDescription(v string) *UpdateVariableGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateVariableGroupRequest) SetName(v string) *UpdateVariableGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateVariableGroupRequest) SetVariables(v string) *UpdateVariableGroupRequest {
	s.Variables = &v
	return s
}

func (s *UpdateVariableGroupRequest) Validate() error {
	return dara.Validate(s)
}
