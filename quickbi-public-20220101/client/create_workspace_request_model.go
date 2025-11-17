// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPublish(v bool) *CreateWorkspaceRequest
	GetAllowPublish() *bool
	SetAllowShare(v bool) *CreateWorkspaceRequest
	GetAllowShare() *bool
	SetAllowViewAll(v bool) *CreateWorkspaceRequest
	GetAllowViewAll() *bool
	SetDefaultShareToAll(v bool) *CreateWorkspaceRequest
	GetDefaultShareToAll() *bool
	SetOnlyAdminCreateDatasource(v bool) *CreateWorkspaceRequest
	GetOnlyAdminCreateDatasource() *bool
	SetUseComment(v bool) *CreateWorkspaceRequest
	GetUseComment() *bool
	SetWorkspaceDescription(v string) *CreateWorkspaceRequest
	GetWorkspaceDescription() *string
	SetWorkspaceName(v string) *CreateWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateWorkspaceRequest struct {
	// example:
	//
	// true
	AllowPublish *bool `json:"AllowPublish,omitempty" xml:"AllowPublish,omitempty"`
	// example:
	//
	// true
	AllowShare *bool `json:"AllowShare,omitempty" xml:"AllowShare,omitempty"`
	// example:
	//
	// true
	AllowViewAll *bool `json:"AllowViewAll,omitempty" xml:"AllowViewAll,omitempty"`
	// example:
	//
	// false
	DefaultShareToAll *bool `json:"DefaultShareToAll,omitempty" xml:"DefaultShareToAll,omitempty"`
	// example:
	//
	// false
	OnlyAdminCreateDatasource *bool `json:"OnlyAdminCreateDatasource,omitempty" xml:"OnlyAdminCreateDatasource,omitempty"`
	// example:
	//
	// true
	UseComment           *bool   `json:"UseComment,omitempty" xml:"UseComment,omitempty"`
	WorkspaceDescription *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	// This parameter is required.
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetAllowPublish() *bool {
	return s.AllowPublish
}

func (s *CreateWorkspaceRequest) GetAllowShare() *bool {
	return s.AllowShare
}

func (s *CreateWorkspaceRequest) GetAllowViewAll() *bool {
	return s.AllowViewAll
}

func (s *CreateWorkspaceRequest) GetDefaultShareToAll() *bool {
	return s.DefaultShareToAll
}

func (s *CreateWorkspaceRequest) GetOnlyAdminCreateDatasource() *bool {
	return s.OnlyAdminCreateDatasource
}

func (s *CreateWorkspaceRequest) GetUseComment() *bool {
	return s.UseComment
}

func (s *CreateWorkspaceRequest) GetWorkspaceDescription() *string {
	return s.WorkspaceDescription
}

func (s *CreateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceRequest) SetAllowPublish(v bool) *CreateWorkspaceRequest {
	s.AllowPublish = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAllowShare(v bool) *CreateWorkspaceRequest {
	s.AllowShare = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAllowViewAll(v bool) *CreateWorkspaceRequest {
	s.AllowViewAll = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDefaultShareToAll(v bool) *CreateWorkspaceRequest {
	s.DefaultShareToAll = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOnlyAdminCreateDatasource(v bool) *CreateWorkspaceRequest {
	s.OnlyAdminCreateDatasource = &v
	return s
}

func (s *CreateWorkspaceRequest) SetUseComment(v bool) *CreateWorkspaceRequest {
	s.UseComment = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceDescription(v string) *CreateWorkspaceRequest {
	s.WorkspaceDescription = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
