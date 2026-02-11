// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceInput
	GetDescription() *string
	SetName(v string) *CreateWorkspaceInput
	GetName() *string
	SetResourceGroupId(v string) *CreateWorkspaceInput
	GetResourceGroupId() *string
}

type CreateWorkspaceInput struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s CreateWorkspaceInput) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceInput) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceInput) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceInput) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceInput) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWorkspaceInput) SetDescription(v string) *CreateWorkspaceInput {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceInput) SetName(v string) *CreateWorkspaceInput {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceInput) SetResourceGroupId(v string) *CreateWorkspaceInput {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWorkspaceInput) Validate() error {
	return dara.Validate(s)
}
