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
	// Specifies whether reports in the workspace can be made public. Default value: true.
	//
	// example:
	//
	// true
	AllowPublish *bool `json:"AllowPublish,omitempty" xml:"AllowPublish,omitempty"`
	// Specifies whether reports in the workspace can be shared. Default value: true.
	//
	// example:
	//
	// true
	AllowShare *bool `json:"AllowShare,omitempty" xml:"AllowShare,omitempty"`
	// Specifies whether the workspace is in collaboration mode. Default value: true.
	//
	// example:
	//
	// true
	AllowViewAll *bool `json:"AllowViewAll,omitempty" xml:"AllowViewAll,omitempty"`
	// Specifies whether to grant read permissions on the works in the workspace to all workspace members by default. Default value: false.
	//
	// example:
	//
	// false
	DefaultShareToAll *bool `json:"DefaultShareToAll,omitempty" xml:"DefaultShareToAll,omitempty"`
	// Specifies whether only administrators can create data sources in the workspace. Default value: false.
	//
	// example:
	//
	// false
	OnlyAdminCreateDatasource *bool `json:"OnlyAdminCreateDatasource,omitempty" xml:"OnlyAdminCreateDatasource,omitempty"`
	// Specifies whether to use table remarks when you create a dataset in the workspace. Default value: true.
	//
	// example:
	//
	// true
	UseComment *bool `json:"UseComment,omitempty" xml:"UseComment,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// test
	WorkspaceDescription *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
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
