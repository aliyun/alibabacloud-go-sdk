// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomPrompt interface {
	dara.Model
	String() string
	GoString() string
	SetRoleDefinition(v string) *CustomPrompt
	GetRoleDefinition() *string
}

type CustomPrompt struct {
	RoleDefinition *string `json:"RoleDefinition,omitempty" xml:"RoleDefinition,omitempty"`
}

func (s CustomPrompt) String() string {
	return dara.Prettify(s)
}

func (s CustomPrompt) GoString() string {
	return s.String()
}

func (s *CustomPrompt) GetRoleDefinition() *string {
	return s.RoleDefinition
}

func (s *CustomPrompt) SetRoleDefinition(v string) *CustomPrompt {
	s.RoleDefinition = &v
	return s
}

func (s *CustomPrompt) Validate() error {
	return dara.Validate(s)
}
