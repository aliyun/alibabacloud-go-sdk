// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeSourceItem interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CodeSourceItem
	GetAccessibility() *string
	SetCloneType(v int32) *CodeSourceItem
	GetCloneType() *int32
	SetCodeBranch(v string) *CodeSourceItem
	GetCodeBranch() *string
	SetCodeCommit(v string) *CodeSourceItem
	GetCodeCommit() *string
	SetCodeRepo(v string) *CodeSourceItem
	GetCodeRepo() *string
	SetCodeRepoAccessToken(v string) *CodeSourceItem
	GetCodeRepoAccessToken() *string
	SetCodeRepoUserName(v string) *CodeSourceItem
	GetCodeRepoUserName() *string
	SetCodeSourceId(v string) *CodeSourceItem
	GetCodeSourceId() *string
	SetDescription(v string) *CodeSourceItem
	GetDescription() *string
	SetDisplayName(v string) *CodeSourceItem
	GetDisplayName() *string
	SetGmtCreateTime(v string) *CodeSourceItem
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *CodeSourceItem
	GetGmtModifyTime() *string
	SetMountPath(v string) *CodeSourceItem
	GetMountPath() *string
	SetUserId(v string) *CodeSourceItem
	GetUserId() *string
	SetWorkspaceId(v string) *CodeSourceItem
	GetWorkspaceId() *string
}

type CodeSourceItem struct {
	// The visibility of the code source. Valid values:
	//
	// 	- PRIVATE: Visible only to you and the administrator in the workspace.
	//
	// 	- PUBLIC: Visible to all users in the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CloneType     *int32  `json:"CloneType,omitempty" xml:"CloneType,omitempty"`
	// The code branch.
	//
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// The code commit ID
	//
	// example:
	//
	// 44da10**********
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The address of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/****
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The token used to access the code repository.
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username of the code repository.
	//
	// example:
	//
	// user
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The code source ID.
	//
	// example:
	//
	// code-202**********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The code source description.
	//
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The code source name.
	//
	// example:
	//
	// MyCodeSourceName1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The local mount path of the code.
	//
	// example:
	//
	// /root/code/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 1157290171663117
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CodeSourceItem) String() string {
	return dara.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
}

func (s *CodeSourceItem) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CodeSourceItem) GetCloneType() *int32 {
	return s.CloneType
}

func (s *CodeSourceItem) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *CodeSourceItem) GetCodeCommit() *string {
	return s.CodeCommit
}

func (s *CodeSourceItem) GetCodeRepo() *string {
	return s.CodeRepo
}

func (s *CodeSourceItem) GetCodeRepoAccessToken() *string {
	return s.CodeRepoAccessToken
}

func (s *CodeSourceItem) GetCodeRepoUserName() *string {
	return s.CodeRepoUserName
}

func (s *CodeSourceItem) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *CodeSourceItem) GetDescription() *string {
	return s.Description
}

func (s *CodeSourceItem) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CodeSourceItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CodeSourceItem) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *CodeSourceItem) GetMountPath() *string {
	return s.MountPath
}

func (s *CodeSourceItem) GetUserId() *string {
	return s.UserId
}

func (s *CodeSourceItem) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CodeSourceItem) SetAccessibility(v string) *CodeSourceItem {
	s.Accessibility = &v
	return s
}

func (s *CodeSourceItem) SetCloneType(v int32) *CodeSourceItem {
	s.CloneType = &v
	return s
}

func (s *CodeSourceItem) SetCodeBranch(v string) *CodeSourceItem {
	s.CodeBranch = &v
	return s
}

func (s *CodeSourceItem) SetCodeCommit(v string) *CodeSourceItem {
	s.CodeCommit = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepo(v string) *CodeSourceItem {
	s.CodeRepo = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoAccessToken(v string) *CodeSourceItem {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoUserName(v string) *CodeSourceItem {
	s.CodeRepoUserName = &v
	return s
}

func (s *CodeSourceItem) SetCodeSourceId(v string) *CodeSourceItem {
	s.CodeSourceId = &v
	return s
}

func (s *CodeSourceItem) SetDescription(v string) *CodeSourceItem {
	s.Description = &v
	return s
}

func (s *CodeSourceItem) SetDisplayName(v string) *CodeSourceItem {
	s.DisplayName = &v
	return s
}

func (s *CodeSourceItem) SetGmtCreateTime(v string) *CodeSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *CodeSourceItem) SetGmtModifyTime(v string) *CodeSourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *CodeSourceItem) SetMountPath(v string) *CodeSourceItem {
	s.MountPath = &v
	return s
}

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

func (s *CodeSourceItem) SetWorkspaceId(v string) *CodeSourceItem {
	s.WorkspaceId = &v
	return s
}

func (s *CodeSourceItem) Validate() error {
	return dara.Validate(s)
}
