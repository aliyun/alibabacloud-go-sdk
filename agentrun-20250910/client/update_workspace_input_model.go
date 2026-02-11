// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWorkspaceInput
	GetDescription() *string
}

type UpdateWorkspaceInput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateWorkspaceInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceInput) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkspaceInput) SetDescription(v string) *UpdateWorkspaceInput {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceInput) Validate() error {
	return dara.Validate(s)
}
