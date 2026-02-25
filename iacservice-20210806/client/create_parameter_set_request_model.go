// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateParameterSetRequest
	GetClientToken() *string
	SetDescription(v string) *CreateParameterSetRequest
	GetDescription() *string
	SetName(v string) *CreateParameterSetRequest
	GetName() *string
	SetParameters(v []*CreateParameterSetRequestParameters) *CreateParameterSetRequest
	GetParameters() []*CreateParameterSetRequestParameters
}

type CreateParameterSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name       *string                                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters []*CreateParameterSetRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
}

func (s CreateParameterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateParameterSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterSetRequest) GetName() *string {
	return s.Name
}

func (s *CreateParameterSetRequest) GetParameters() []*CreateParameterSetRequestParameters {
	return s.Parameters
}

func (s *CreateParameterSetRequest) SetClientToken(v string) *CreateParameterSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterSetRequest) SetDescription(v string) *CreateParameterSetRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterSetRequest) SetName(v string) *CreateParameterSetRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterSetRequest) SetParameters(v []*CreateParameterSetRequestParameters) *CreateParameterSetRequest {
	s.Parameters = v
	return s
}

func (s *CreateParameterSetRequest) Validate() error {
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

type CreateParameterSetRequestParameters struct {
	// example:
	//
	// test1121
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateParameterSetRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterSetRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateParameterSetRequestParameters) GetName() *string {
	return s.Name
}

func (s *CreateParameterSetRequestParameters) GetStatus() *string {
	return s.Status
}

func (s *CreateParameterSetRequestParameters) GetType() *string {
	return s.Type
}

func (s *CreateParameterSetRequestParameters) GetValue() *string {
	return s.Value
}

func (s *CreateParameterSetRequestParameters) SetName(v string) *CreateParameterSetRequestParameters {
	s.Name = &v
	return s
}

func (s *CreateParameterSetRequestParameters) SetStatus(v string) *CreateParameterSetRequestParameters {
	s.Status = &v
	return s
}

func (s *CreateParameterSetRequestParameters) SetType(v string) *CreateParameterSetRequestParameters {
	s.Type = &v
	return s
}

func (s *CreateParameterSetRequestParameters) SetValue(v string) *CreateParameterSetRequestParameters {
	s.Value = &v
	return s
}

func (s *CreateParameterSetRequestParameters) Validate() error {
	return dara.Validate(s)
}
