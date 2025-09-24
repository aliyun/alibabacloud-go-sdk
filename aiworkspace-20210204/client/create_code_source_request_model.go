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
	// The visibility of the code build. Valid values:
	//
	// 	- PUBLIC: The code build is visible to all members in the workspace.
	//
	// 	- PRIVATE: The code build is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The code branch.
	//
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The URL of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/******
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The token used to access the code repository.
	//
	// example:
	//
	// ***
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username of the code repository.
	//
	// example:
	//
	// use***
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The description of the code build, which helps you distinguish between code builds.
	//
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the code build.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The local mount path of the code. By default, the code is mounted to the /root/code/ path.
	//
	// example:
	//
	// /root/code/code-source-1
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
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
