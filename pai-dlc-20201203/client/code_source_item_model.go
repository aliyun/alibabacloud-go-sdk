// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeSourceItem interface {
	dara.Model
	String() string
	GoString() string
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
	SetUserId(v string) *CodeSourceItem
	GetUserId() *string
}

type CodeSourceItem struct {
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// example:
	//
	// 44da1*******
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// example:
	//
	// https://code.aliyun.com/pai-dlc/examples.git
	CodeRepo            *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// example:
	//
	// user
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// example:
	//
	// code-20210111103721-85qz*****
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MyCodeSourceName1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// 115**********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CodeSourceItem) String() string {
	return dara.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
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

func (s *CodeSourceItem) GetUserId() *string {
	return s.UserId
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

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

func (s *CodeSourceItem) Validate() error {
	return dara.Validate(s)
}
