// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest
	GetCreateParameters() []*UpdateFunctionInstanceRequestCreateParameters
	SetDescription(v string) *UpdateFunctionInstanceRequest
	GetDescription() *string
}

type UpdateFunctionInstanceRequest struct {
	// The creation parameters.
	CreateParameters []*UpdateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequest) GetCreateParameters() []*UpdateFunctionInstanceRequestCreateParameters {
	return s.CreateParameters
}

func (s *UpdateFunctionInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFunctionInstanceRequest) SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetDescription(v string) *UpdateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) Validate() error {
	if s.CreateParameters != nil {
		for _, item := range s.CreateParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFunctionInstanceRequestCreateParameters struct {
	// The parameter name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestCreateParameters) GetName() *string {
	return s.Name
}

func (s *UpdateFunctionInstanceRequestCreateParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetName(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetValue(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) Validate() error {
	return dara.Validate(s)
}
