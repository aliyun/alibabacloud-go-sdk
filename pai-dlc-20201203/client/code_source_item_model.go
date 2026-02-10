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
	// The branch of the code repository. If you configure this parameter when you call the CreateJob API operation, the branch is overwritten.
	//
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// The commit ID. If you configure this parameter when you call the CreateJob API operation, the commit is overwritten.
	//
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487bb72ea63
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// The URL of the code repository.
	//
	// example:
	//
	// https://code.aliyun.com/pai-dlc/examples.git
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// The access token used to access the code repository.
	//
	// example:
	//
	// xxxx
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// The username used to access the private code repository.
	//
	// example:
	//
	// user
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// The ID of the code source.
	//
	// example:
	//
	// code-20210111103721-********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The description of the code source.
	//
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the code source.
	//
	// example:
	//
	// MyCodeSourceName1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the code source was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the code source was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The UID of the Alibaba Cloud user who creates the code source.
	//
	// example:
	//
	// 115729017166****
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
