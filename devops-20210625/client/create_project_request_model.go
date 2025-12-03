// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomCode(v string) *CreateProjectRequest
	GetCustomCode() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
	SetScope(v string) *CreateProjectRequest
	GetScope() *string
	SetTemplateIdentifier(v string) *CreateProjectRequest
	GetTemplateIdentifier() *string
}

type CreateProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ABCD
	CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 704eaxxxx5efede61xxx5
	TemplateIdentifier *string `json:"templateIdentifier,omitempty" xml:"templateIdentifier,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetCustomCode() *string {
	return s.CustomCode
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateProjectRequest) GetTemplateIdentifier() *string {
	return s.TemplateIdentifier
}

func (s *CreateProjectRequest) SetCustomCode(v string) *CreateProjectRequest {
	s.CustomCode = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetScope(v string) *CreateProjectRequest {
	s.Scope = &v
	return s
}

func (s *CreateProjectRequest) SetTemplateIdentifier(v string) *CreateProjectRequest {
	s.TemplateIdentifier = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
