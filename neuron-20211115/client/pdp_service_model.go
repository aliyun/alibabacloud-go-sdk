// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpService interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpService
	GetAccountId() *string
	SetAlias(v string) *PdpService
	GetAlias() *string
	SetCodeBranch(v string) *PdpService
	GetCodeBranch() *string
	SetDescription(v string) *PdpService
	GetDescription() *string
	SetEnterpriseId(v int64) *PdpService
	GetEnterpriseId() *int64
	SetExtraInfo(v string) *PdpService
	GetExtraInfo() *string
	SetGitRepo(v string) *PdpService
	GetGitRepo() *string
	SetGmtCreate(v string) *PdpService
	GetGmtCreate() *string
	SetGmtModified(v string) *PdpService
	GetGmtModified() *string
	SetId(v int64) *PdpService
	GetId() *int64
	SetJumpUrl(v string) *PdpService
	GetJumpUrl() *string
	SetName(v string) *PdpService
	GetName() *string
	SetOrgType(v string) *PdpService
	GetOrgType() *string
	SetPbcId(v int64) *PdpService
	GetPbcId() *int64
	SetRequestId(v string) *PdpService
	GetRequestId() *string
	SetStatus(v string) *PdpService
	GetStatus() *string
	SetType(v string) *PdpService
	GetType() *string
}

type PdpService struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
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
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// SDK
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// git@gitlab.alibaba-inc.com:neuron/manager-developer.git
	GitRepo *string `json:"gitRepo,omitempty" xml:"gitRepo,omitempty"`
	// example:
	//
	// 2022-2-22 11:11:2
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-2-22 11:11:2
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// https://cd.aone.alibaba-inc.com/ec/pipelines/44156?spm=a2o8d.aone_cd_assets_pages_pipeline_index.0.0.68b81c05WLmX79&name=%E3%80%90%E6%B5%8B%E8%AF%95%E3%80%91neuron-developer%E6%9D%AD%E5%B7%9E
	JumpUrl *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// employee
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// INNER
	OrgType *string `json:"orgType,omitempty" xml:"orgType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SDK
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpService) String() string {
	return dara.Prettify(s)
}

func (s PdpService) GoString() string {
	return s.String()
}

func (s *PdpService) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpService) GetAlias() *string {
	return s.Alias
}

func (s *PdpService) GetCodeBranch() *string {
	return s.CodeBranch
}

func (s *PdpService) GetDescription() *string {
	return s.Description
}

func (s *PdpService) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpService) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *PdpService) GetGitRepo() *string {
	return s.GitRepo
}

func (s *PdpService) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpService) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PdpService) GetId() *int64 {
	return s.Id
}

func (s *PdpService) GetJumpUrl() *string {
	return s.JumpUrl
}

func (s *PdpService) GetName() *string {
	return s.Name
}

func (s *PdpService) GetOrgType() *string {
	return s.OrgType
}

func (s *PdpService) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpService) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpService) GetStatus() *string {
	return s.Status
}

func (s *PdpService) GetType() *string {
	return s.Type
}

func (s *PdpService) SetAccountId(v string) *PdpService {
	s.AccountId = &v
	return s
}

func (s *PdpService) SetAlias(v string) *PdpService {
	s.Alias = &v
	return s
}

func (s *PdpService) SetCodeBranch(v string) *PdpService {
	s.CodeBranch = &v
	return s
}

func (s *PdpService) SetDescription(v string) *PdpService {
	s.Description = &v
	return s
}

func (s *PdpService) SetEnterpriseId(v int64) *PdpService {
	s.EnterpriseId = &v
	return s
}

func (s *PdpService) SetExtraInfo(v string) *PdpService {
	s.ExtraInfo = &v
	return s
}

func (s *PdpService) SetGitRepo(v string) *PdpService {
	s.GitRepo = &v
	return s
}

func (s *PdpService) SetGmtCreate(v string) *PdpService {
	s.GmtCreate = &v
	return s
}

func (s *PdpService) SetGmtModified(v string) *PdpService {
	s.GmtModified = &v
	return s
}

func (s *PdpService) SetId(v int64) *PdpService {
	s.Id = &v
	return s
}

func (s *PdpService) SetJumpUrl(v string) *PdpService {
	s.JumpUrl = &v
	return s
}

func (s *PdpService) SetName(v string) *PdpService {
	s.Name = &v
	return s
}

func (s *PdpService) SetOrgType(v string) *PdpService {
	s.OrgType = &v
	return s
}

func (s *PdpService) SetPbcId(v int64) *PdpService {
	s.PbcId = &v
	return s
}

func (s *PdpService) SetRequestId(v string) *PdpService {
	s.RequestId = &v
	return s
}

func (s *PdpService) SetStatus(v string) *PdpService {
	s.Status = &v
	return s
}

func (s *PdpService) SetType(v string) *PdpService {
	s.Type = &v
	return s
}

func (s *PdpService) Validate() error {
	return dara.Validate(s)
}
