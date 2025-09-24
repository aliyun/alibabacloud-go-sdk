// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetCodeSourceResponseBody
	GetAccessibility() *string
	SetCodeBranch(v string) *GetCodeSourceResponseBody
	GetCodeBranch() *string
	SetCodeCommit(v string) *GetCodeSourceResponseBody
	GetCodeCommit() *string
	SetCodeRepo(v string) *GetCodeSourceResponseBody
	GetCodeRepo() *string
	SetCodeRepoAccessToken(v string) *GetCodeSourceResponseBody
	GetCodeRepoAccessToken() *string
	SetCodeRepoUserName(v string) *GetCodeSourceResponseBody
	GetCodeRepoUserName() *string
	SetCodeSourceId(v string) *GetCodeSourceResponseBody
	GetCodeSourceId() *string
	SetDescription(v string) *GetCodeSourceResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetCodeSourceResponseBody
	GetDisplayName() *string
	SetGmtCreateTime(v string) *GetCodeSourceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *GetCodeSourceResponseBody
	GetGmtModifyTime() *string
	SetMountPath(v string) *GetCodeSourceResponseBody
	GetMountPath() *string
	SetRequestId(v string) *GetCodeSourceResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetCodeSourceResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetCodeSourceResponseBody
	GetWorkspaceId() *string
}

type GetCodeSourceResponseBody struct {
	// The visibility of the code source. Valid values:
	//
	// 	- PRIVATE: Visible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: Visible to all members in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The code repository branch.
	//
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// The code commit ID.
	//
	// example:
	//
	// 44da10***********
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The address of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/pai-dlc/examples.git
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The token used to access the code repository.
	//
	// example:
	//
	// xxxx
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username of the code repository.
	//
	// example:
	//
	// user1
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The ID of the code source.
	//
	// example:
	//
	// code-202**********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The description of the code source.
	//
	// example:
	//
	// This is my data source 1.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the code source.
	//
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the code source was created, in the ISO8601 format.
	//
	// example:
	//
	// 2021-01-12T23:36:01.123Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the code source was modified, in the ISO8601 format.
	//
	// example:
	//
	// 2021-01-12T23:36:01.123Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The local mount path of the code.
	//
	// example:
	//
	// /root/code
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 1722********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCodeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetCodeSourceResponseBody) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *GetCodeSourceResponseBody) GetCodeCommit() *string {
	return s.CodeCommit
}

func (s *GetCodeSourceResponseBody) GetCodeRepo() *string {
	return s.CodeRepo
}

func (s *GetCodeSourceResponseBody) GetCodeRepoAccessToken() *string {
	return s.CodeRepoAccessToken
}

func (s *GetCodeSourceResponseBody) GetCodeRepoUserName() *string {
	return s.CodeRepoUserName
}

func (s *GetCodeSourceResponseBody) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *GetCodeSourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetCodeSourceResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCodeSourceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetCodeSourceResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *GetCodeSourceResponseBody) GetMountPath() *string {
	return s.MountPath
}

func (s *GetCodeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCodeSourceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetCodeSourceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCodeSourceResponseBody) SetAccessibility(v string) *GetCodeSourceResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeBranch(v string) *GetCodeSourceResponseBody {
	s.CodeBranch = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeCommit(v string) *GetCodeSourceResponseBody {
	s.CodeCommit = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepo(v string) *GetCodeSourceResponseBody {
	s.CodeRepo = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoAccessToken(v string) *GetCodeSourceResponseBody {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoUserName(v string) *GetCodeSourceResponseBody {
	s.CodeRepoUserName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeSourceId(v string) *GetCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDescription(v string) *GetCodeSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDisplayName(v string) *GetCodeSourceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtCreateTime(v string) *GetCodeSourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtModifyTime(v string) *GetCodeSourceResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetMountPath(v string) *GetCodeSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetRequestId(v string) *GetCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetUserId(v string) *GetCodeSourceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetWorkspaceId(v string) *GetCodeSourceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetCodeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
