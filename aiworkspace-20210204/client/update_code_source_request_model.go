// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeBranch(v string) *UpdateCodeSourceRequest
	GetCodeBranch() *string
	SetCodeCommit(v string) *UpdateCodeSourceRequest
	GetCodeCommit() *string
	SetCodeRepo(v string) *UpdateCodeSourceRequest
	GetCodeRepo() *string
	SetCodeRepoAccessToken(v string) *UpdateCodeSourceRequest
	GetCodeRepoAccessToken() *string
	SetCodeRepoUserName(v string) *UpdateCodeSourceRequest
	GetCodeRepoUserName() *string
	SetDescription(v string) *UpdateCodeSourceRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateCodeSourceRequest
	GetDisplayName() *string
	SetMountPath(v string) *UpdateCodeSourceRequest
	GetMountPath() *string
}

type UpdateCodeSourceRequest struct {
	// The name of the code branch.
	//
	// example:
	//
	// dev
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// The code commit ID.
	//
	// example:
	//
	// 3a6*****
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The address of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/******
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The access token corresponding to the username.
	//
	// example:
	//
	// ***
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username used to access the code repository.
	//
	// example:
	//
	// demo-user
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The description of the code build.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the code build.
	//
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The default mount path.
	//
	// example:
	//
	// /root/code/code-source-1
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s UpdateCodeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateCodeSourceRequest) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *UpdateCodeSourceRequest) GetCodeCommit() *string {
	return s.CodeCommit
}

func (s *UpdateCodeSourceRequest) GetCodeRepo() *string {
	return s.CodeRepo
}

func (s *UpdateCodeSourceRequest) GetCodeRepoAccessToken() *string {
	return s.CodeRepoAccessToken
}

func (s *UpdateCodeSourceRequest) GetCodeRepoUserName() *string {
	return s.CodeRepoUserName
}

func (s *UpdateCodeSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCodeSourceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateCodeSourceRequest) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateCodeSourceRequest) SetCodeBranch(v string) *UpdateCodeSourceRequest {
	s.CodeBranch = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetCodeCommit(v string) *UpdateCodeSourceRequest {
	s.CodeCommit = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetCodeRepo(v string) *UpdateCodeSourceRequest {
	s.CodeRepo = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetCodeRepoAccessToken(v string) *UpdateCodeSourceRequest {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetCodeRepoUserName(v string) *UpdateCodeSourceRequest {
	s.CodeRepoUserName = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetDescription(v string) *UpdateCodeSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetDisplayName(v string) *UpdateCodeSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateCodeSourceRequest) SetMountPath(v string) *UpdateCodeSourceRequest {
	s.MountPath = &v
	return s
}

func (s *UpdateCodeSourceRequest) Validate() error {
	return dara.Validate(s)
}
