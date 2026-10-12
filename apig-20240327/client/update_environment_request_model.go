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
	// The alias of the environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test environment
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The description of the environment. You can specify information such as the purpose of the environment and the owner.
	//
	// example:
	//
	// This is the test environment for the xx project of xx
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
