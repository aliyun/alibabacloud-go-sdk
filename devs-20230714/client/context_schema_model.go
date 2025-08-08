// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextSchema interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ContextSchema
	GetDescription() *string
	SetHint(v string) *ContextSchema
	GetHint() *string
	SetName(v string) *ContextSchema
	GetName() *string
	SetRequired(v bool) *ContextSchema
	GetRequired() *bool
	SetType(v string) *ContextSchema
	GetType() *string
}

type ContextSchema struct {
	// example:
	//
	// [git](https://git-scm.com/) address for [git clone](https://git-scm.com/docs/git-clone).
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// git@gitlab.alibaba-inc.com:serverless/lambda.git
	Hint *string `json:"hint,omitempty" xml:"hint,omitempty"`
	// example:
	//
	// gitRepoUrl
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ContextSchema) String() string {
	return dara.Prettify(s)
}

func (s ContextSchema) GoString() string {
	return s.String()
}

func (s *ContextSchema) GetDescription() *string {
	return s.Description
}

func (s *ContextSchema) GetHint() *string {
	return s.Hint
}

func (s *ContextSchema) GetName() *string {
	return s.Name
}

func (s *ContextSchema) GetRequired() *bool {
	return s.Required
}

func (s *ContextSchema) GetType() *string {
	return s.Type
}

func (s *ContextSchema) SetDescription(v string) *ContextSchema {
	s.Description = &v
	return s
}

func (s *ContextSchema) SetHint(v string) *ContextSchema {
	s.Hint = &v
	return s
}

func (s *ContextSchema) SetName(v string) *ContextSchema {
	s.Name = &v
	return s
}

func (s *ContextSchema) SetRequired(v bool) *ContextSchema {
	s.Required = &v
	return s
}

func (s *ContextSchema) SetType(v string) *ContextSchema {
	s.Type = &v
	return s
}

func (s *ContextSchema) Validate() error {
	return dara.Validate(s)
}
