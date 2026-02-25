// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterSetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateParameterSetAttributeRequest
	GetDescription() *string
	SetName(v string) *UpdateParameterSetAttributeRequest
	GetName() *string
	SetParameters(v []*UpdateParameterSetAttributeRequestParameters) *UpdateParameterSetAttributeRequest
	GetParameters() []*UpdateParameterSetAttributeRequestParameters
}

type UpdateParameterSetAttributeRequest struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name       *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	Parameters []*UpdateParameterSetAttributeRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
}

func (s UpdateParameterSetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateParameterSetAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateParameterSetAttributeRequest) GetParameters() []*UpdateParameterSetAttributeRequestParameters {
	return s.Parameters
}

func (s *UpdateParameterSetAttributeRequest) SetDescription(v string) *UpdateParameterSetAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterSetAttributeRequest) SetName(v string) *UpdateParameterSetAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateParameterSetAttributeRequest) SetParameters(v []*UpdateParameterSetAttributeRequestParameters) *UpdateParameterSetAttributeRequest {
	s.Parameters = v
	return s
}

func (s *UpdateParameterSetAttributeRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateParameterSetAttributeRequestParameters struct {
	// example:
	//
	// t
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-bp1mjm9exduos1bipw9x6
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateParameterSetAttributeRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterSetAttributeRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeRequestParameters) GetName() *string {
	return s.Name
}

func (s *UpdateParameterSetAttributeRequestParameters) GetStatus() *string {
	return s.Status
}

func (s *UpdateParameterSetAttributeRequestParameters) GetType() *string {
	return s.Type
}

func (s *UpdateParameterSetAttributeRequestParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateParameterSetAttributeRequestParameters) SetName(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Name = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) SetStatus(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Status = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) SetType(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Type = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) SetValue(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Value = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) Validate() error {
	return dara.Validate(s)
}
