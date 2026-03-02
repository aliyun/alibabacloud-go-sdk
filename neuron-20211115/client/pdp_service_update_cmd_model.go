// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpServiceUpdateCmd
	GetAlias() *string
	SetCodeBranch(v string) *PdpServiceUpdateCmd
	GetCodeBranch() *string
	SetDescription(v string) *PdpServiceUpdateCmd
	GetDescription() *string
	SetExtraInfo(v string) *PdpServiceUpdateCmd
	GetExtraInfo() *string
	SetGitRepo(v string) *PdpServiceUpdateCmd
	GetGitRepo() *string
	SetId(v int64) *PdpServiceUpdateCmd
	GetId() *int64
	SetJumpUrl(v string) *PdpServiceUpdateCmd
	GetJumpUrl() *string
}

type PdpServiceUpdateCmd struct {
	// example:
	//
	// 员工管理
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// dev
	CodeBranch *string `json:"codeBranch,omitempty" xml:"codeBranch,omitempty"`
	// example:
	//
	// 员工管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// SDK
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// git@gitlab.alibaba-inc.com:neuron/manager-developer.git
	GitRepo *string `json:"gitRepo,omitempty" xml:"gitRepo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// https://aone.alibaba-inc.com/appcenter/app/onlinetask/SUBMIT_RESOURCE_APPLY?appId=195041
	JumpUrl *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
}

func (s PdpServiceUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceUpdateCmd) GoString() string {
	return s.String()
}

func (s *PdpServiceUpdateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpServiceUpdateCmd) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *PdpServiceUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceUpdateCmd) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *PdpServiceUpdateCmd) GetGitRepo() *string {
	return s.GitRepo
}

func (s *PdpServiceUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceUpdateCmd) GetJumpUrl() *string {
	return s.JumpUrl
}

func (s *PdpServiceUpdateCmd) SetAlias(v string) *PdpServiceUpdateCmd {
	s.Alias = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetCodeBranch(v string) *PdpServiceUpdateCmd {
	s.CodeBranch = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetDescription(v string) *PdpServiceUpdateCmd {
	s.Description = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetExtraInfo(v string) *PdpServiceUpdateCmd {
	s.ExtraInfo = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetGitRepo(v string) *PdpServiceUpdateCmd {
	s.GitRepo = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetId(v int64) *PdpServiceUpdateCmd {
	s.Id = &v
	return s
}

func (s *PdpServiceUpdateCmd) SetJumpUrl(v string) *PdpServiceUpdateCmd {
	s.JumpUrl = &v
	return s
}

func (s *PdpServiceUpdateCmd) Validate() error {
	return dara.Validate(s)
}
