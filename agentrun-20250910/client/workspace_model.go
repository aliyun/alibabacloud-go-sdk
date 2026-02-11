// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspace interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Workspace
	GetCreatedAt() *string
	SetDescription(v string) *Workspace
	GetDescription() *string
	SetIsDefault(v bool) *Workspace
	GetIsDefault() *bool
	SetName(v string) *Workspace
	GetName() *string
	SetResourceGroupId(v string) *Workspace
	GetResourceGroupId() *string
	SetUpdatedAt(v string) *Workspace
	GetUpdatedAt() *string
	SetWorkspaceArn(v string) *Workspace
	GetWorkspaceArn() *string
	SetWorkspaceId(v string) *Workspace
	GetWorkspaceId() *string
}

type Workspace struct {
	CreatedAt       *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	IsDefault       *bool   `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	WorkspaceArn    *string `json:"workspaceArn,omitempty" xml:"workspaceArn,omitempty"`
	WorkspaceId     *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s Workspace) String() string {
	return dara.Prettify(s)
}

func (s Workspace) GoString() string {
	return s.String()
}

func (s *Workspace) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Workspace) GetDescription() *string {
	return s.Description
}

func (s *Workspace) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *Workspace) GetName() *string {
	return s.Name
}

func (s *Workspace) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Workspace) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Workspace) GetWorkspaceArn() *string {
	return s.WorkspaceArn
}

func (s *Workspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Workspace) SetCreatedAt(v string) *Workspace {
	s.CreatedAt = &v
	return s
}

func (s *Workspace) SetDescription(v string) *Workspace {
	s.Description = &v
	return s
}

func (s *Workspace) SetIsDefault(v bool) *Workspace {
	s.IsDefault = &v
	return s
}

func (s *Workspace) SetName(v string) *Workspace {
	s.Name = &v
	return s
}

func (s *Workspace) SetResourceGroupId(v string) *Workspace {
	s.ResourceGroupId = &v
	return s
}

func (s *Workspace) SetUpdatedAt(v string) *Workspace {
	s.UpdatedAt = &v
	return s
}

func (s *Workspace) SetWorkspaceArn(v string) *Workspace {
	s.WorkspaceArn = &v
	return s
}

func (s *Workspace) SetWorkspaceId(v string) *Workspace {
	s.WorkspaceId = &v
	return s
}

func (s *Workspace) Validate() error {
	return dara.Validate(s)
}
