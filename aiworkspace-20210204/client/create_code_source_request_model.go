// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateCodeSourceRequest
	GetAccessibility() *string
	SetCloneType(v int32) *CreateCodeSourceRequest
	GetCloneType() *int32
	SetCodeBranch(v string) *CreateCodeSourceRequest
	GetCodeBranch() *string
	SetCodeCommit(v string) *CreateCodeSourceRequest
	GetCodeCommit() *string
	SetCodeRepo(v string) *CreateCodeSourceRequest
	GetCodeRepo() *string
	SetCodeRepoAccessToken(v string) *CreateCodeSourceRequest
	GetCodeRepoAccessToken() *string
	SetCodeRepoUserName(v string) *CreateCodeSourceRequest
	GetCodeRepoUserName() *string
	SetDescription(v string) *CreateCodeSourceRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateCodeSourceRequest
	GetDisplayName() *string
	SetMountPath(v string) *CreateCodeSourceRequest
	GetMountPath() *string
	SetWorkspaceId(v string) *CreateCodeSourceRequest
	GetWorkspaceId() *string
}

type CreateCodeSourceRequest struct {
	// The visibility of the code configuration. Valid values:
	//
	// - PUBLIC: The configuration is visible to everyone in the workspace.
	//
	// - PRIVATE: The configuration is visible only to you and workspace administrators.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CloneType     *int32  `json:"CloneType,omitempty" xml:"CloneType,omitempty"`
	// The code branch.
	//
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// The commit ID of the code. \\`CodeCommit\\` takes precedence over \\`CodeBranch\\`. If you specify \\`CodeCommit\\`, \\`CodeBranch\\` is ignored.
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The URL of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/******
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The access token for the code repository.
	//
	// example:
	//
	// ***
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username for the code repository.
	//
	// example:
	//
	// use***
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The description of the code configuration, which helps distinguish it from other configurations.
	//
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the code configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The local mount path for the code. The default is `/root/code/`.
	//
	// example:
	//
	// /root/code/code-source-1
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The workspace ID. For more information, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCodeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateCodeSourceRequest) GetCloneType() *int32 {
	return s.CloneType
}

func (s *CreateCodeSourceRequest) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *CreateCodeSourceRequest) GetCodeCommit() *string {
	return s.CodeCommit
}

func (s *CreateCodeSourceRequest) GetCodeRepo() *string {
	return s.CodeRepo
}

func (s *CreateCodeSourceRequest) GetCodeRepoAccessToken() *string {
	return s.CodeRepoAccessToken
}

func (s *CreateCodeSourceRequest) GetCodeRepoUserName() *string {
	return s.CodeRepoUserName
}

func (s *CreateCodeSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCodeSourceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCodeSourceRequest) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateCodeSourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCodeSourceRequest) SetAccessibility(v string) *CreateCodeSourceRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCloneType(v int32) *CreateCodeSourceRequest {
	s.CloneType = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeBranch(v string) *CreateCodeSourceRequest {
	s.CodeBranch = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeCommit(v string) *CreateCodeSourceRequest {
	s.CodeCommit = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepo(v string) *CreateCodeSourceRequest {
	s.CodeRepo = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoAccessToken(v string) *CreateCodeSourceRequest {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoUserName(v string) *CreateCodeSourceRequest {
	s.CodeRepoUserName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDescription(v string) *CreateCodeSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDisplayName(v string) *CreateCodeSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetMountPath(v string) *CreateCodeSourceRequest {
	s.MountPath = &v
	return s
}

func (s *CreateCodeSourceRequest) SetWorkspaceId(v string) *CreateCodeSourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCodeSourceRequest) Validate() error {
	return dara.Validate(s)
}
