// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateApmRequest
	GetDescription() *string
	SetOutputES(v string) *UpdateApmRequest
	GetOutputES() *string
	SetOutputESPassword(v string) *UpdateApmRequest
	GetOutputESPassword() *string
	SetOutputESUserName(v string) *UpdateApmRequest
	GetOutputESUserName() *string
	SetToken(v string) *UpdateApmRequest
	GetToken() *string
}

type UpdateApmRequest struct {
	// example:
	//
	// APMtest
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// es-cn-i7m2fsfhc001x****
	OutputES *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	// example:
	//
	// ESPassword****
	OutputESPassword *string `json:"outputESPassword,omitempty" xml:"outputESPassword,omitempty"`
	// example:
	//
	// elastic
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	// example:
	//
	// AMPPassword****
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s UpdateApmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApmRequest) GoString() string {
	return s.String()
}

func (s *UpdateApmRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApmRequest) GetOutputES() *string {
	return s.OutputES
}

func (s *UpdateApmRequest) GetOutputESPassword() *string {
	return s.OutputESPassword
}

func (s *UpdateApmRequest) GetOutputESUserName() *string {
	return s.OutputESUserName
}

func (s *UpdateApmRequest) GetToken() *string {
	return s.Token
}

func (s *UpdateApmRequest) SetDescription(v string) *UpdateApmRequest {
	s.Description = &v
	return s
}

func (s *UpdateApmRequest) SetOutputES(v string) *UpdateApmRequest {
	s.OutputES = &v
	return s
}

func (s *UpdateApmRequest) SetOutputESPassword(v string) *UpdateApmRequest {
	s.OutputESPassword = &v
	return s
}

func (s *UpdateApmRequest) SetOutputESUserName(v string) *UpdateApmRequest {
	s.OutputESUserName = &v
	return s
}

func (s *UpdateApmRequest) SetToken(v string) *UpdateApmRequest {
	s.Token = &v
	return s
}

func (s *UpdateApmRequest) Validate() error {
	return dara.Validate(s)
}
