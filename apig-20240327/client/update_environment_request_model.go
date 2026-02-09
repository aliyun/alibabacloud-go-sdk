// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *UpdateEnvironmentRequest
	GetAlias() *string
	SetDescription(v string) *UpdateEnvironmentRequest
	GetDescription() *string
}

type UpdateEnvironmentRequest struct {
	// Schema of Response
	//
	// This parameter is required.
	//
	// example:
	//
	// The request ID, which is used to trace the API call link.
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The status code returned.
	//
	// example:
	//
	// The response message returned.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) GetAlias() *string {
	return s.Alias
}

func (s *UpdateEnvironmentRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEnvironmentRequest) SetAlias(v string) *UpdateEnvironmentRequest {
	s.Alias = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetDescription(v string) *UpdateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *UpdateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
