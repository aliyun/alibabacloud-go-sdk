// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceGroupDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpServiceGroupDetail
	GetAccountId() *string
	SetAlias(v string) *PdpServiceGroupDetail
	GetAlias() *string
	SetDescription(v string) *PdpServiceGroupDetail
	GetDescription() *string
	SetEdasId(v string) *PdpServiceGroupDetail
	GetEdasId() *string
	SetEnterpriseId(v int64) *PdpServiceGroupDetail
	GetEnterpriseId() *int64
	SetEnv(v string) *PdpServiceGroupDetail
	GetEnv() *string
	SetEnvType(v string) *PdpServiceGroupDetail
	GetEnvType() *string
	SetGmtCreate(v string) *PdpServiceGroupDetail
	GetGmtCreate() *string
	SetGmtModified(v string) *PdpServiceGroupDetail
	GetGmtModified() *string
	SetGroupType(v string) *PdpServiceGroupDetail
	GetGroupType() *string
	SetId(v int64) *PdpServiceGroupDetail
	GetId() *int64
	SetLogProject(v string) *PdpServiceGroupDetail
	GetLogProject() *string
	SetLogStore(v string) *PdpServiceGroupDetail
	GetLogStore() *string
	SetName(v string) *PdpServiceGroupDetail
	GetName() *string
	SetOrgType(v string) *PdpServiceGroupDetail
	GetOrgType() *string
	SetPbcId(v int64) *PdpServiceGroupDetail
	GetPbcId() *int64
	SetPipelineId(v string) *PdpServiceGroupDetail
	GetPipelineId() *string
	SetProductId(v int64) *PdpServiceGroupDetail
	GetProductId() *int64
	SetProductName(v string) *PdpServiceGroupDetail
	GetProductName() *string
	SetRegion(v string) *PdpServiceGroupDetail
	GetRegion() *string
	SetRequestId(v string) *PdpServiceGroupDetail
	GetRequestId() *string
	SetServiceId(v int64) *PdpServiceGroupDetail
	GetServiceId() *int64
	SetType(v string) *PdpServiceGroupDetail
	GetType() *string
}

type PdpServiceGroupDetail struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 员工管理
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 979ae3bf-805b-4561-8019-43b034179441
	EdasId *string `json:"edasId,omitempty" xml:"edasId,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// DEV
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// DEFAULT
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// k8s-log-custom-asi-region-cn-hangzhou-1883014953109615
	LogProject *string `json:"logProject,omitempty" xml:"logProject,omitempty"`
	// example:
	//
	// asi-addon-stdout
	LogStore *string `json:"logStore,omitempty" xml:"logStore,omitempty"`
	// example:
	//
	// employee
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// INNER
	OrgType *string `json:"orgType,omitempty" xml:"orgType,omitempty"`
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 979ae3bf-805b-4561-801
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// yunmall
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// SDK
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpServiceGroupDetail) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceGroupDetail) GoString() string {
	return s.String()
}

func (s *PdpServiceGroupDetail) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpServiceGroupDetail) GetAlias() *string {
	return s.Alias
}

func (s *PdpServiceGroupDetail) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceGroupDetail) GetEdasId() *string {
	return s.EdasId
}

func (s *PdpServiceGroupDetail) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpServiceGroupDetail) GetEnv() *string {
	return s.Env
}

func (s *PdpServiceGroupDetail) GetEnvType() *string {
	return s.EnvType
}

func (s *PdpServiceGroupDetail) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpServiceGroupDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PdpServiceGroupDetail) GetGroupType() *string {
	return s.GroupType
}

func (s *PdpServiceGroupDetail) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceGroupDetail) GetLogProject() *string {
	return s.LogProject
}

func (s *PdpServiceGroupDetail) GetLogStore() *string {
	return s.LogStore
}

func (s *PdpServiceGroupDetail) GetName() *string {
	return s.Name
}

func (s *PdpServiceGroupDetail) GetOrgType() *string {
	return s.OrgType
}

func (s *PdpServiceGroupDetail) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpServiceGroupDetail) GetPipelineId() *string {
	return s.PipelineId
}

func (s *PdpServiceGroupDetail) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpServiceGroupDetail) GetProductName() *string {
	return s.ProductName
}

func (s *PdpServiceGroupDetail) GetRegion() *string {
	return s.Region
}

func (s *PdpServiceGroupDetail) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpServiceGroupDetail) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpServiceGroupDetail) GetType() *string {
	return s.Type
}

func (s *PdpServiceGroupDetail) SetAccountId(v string) *PdpServiceGroupDetail {
	s.AccountId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetAlias(v string) *PdpServiceGroupDetail {
	s.Alias = &v
	return s
}

func (s *PdpServiceGroupDetail) SetDescription(v string) *PdpServiceGroupDetail {
	s.Description = &v
	return s
}

func (s *PdpServiceGroupDetail) SetEdasId(v string) *PdpServiceGroupDetail {
	s.EdasId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetEnterpriseId(v int64) *PdpServiceGroupDetail {
	s.EnterpriseId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetEnv(v string) *PdpServiceGroupDetail {
	s.Env = &v
	return s
}

func (s *PdpServiceGroupDetail) SetEnvType(v string) *PdpServiceGroupDetail {
	s.EnvType = &v
	return s
}

func (s *PdpServiceGroupDetail) SetGmtCreate(v string) *PdpServiceGroupDetail {
	s.GmtCreate = &v
	return s
}

func (s *PdpServiceGroupDetail) SetGmtModified(v string) *PdpServiceGroupDetail {
	s.GmtModified = &v
	return s
}

func (s *PdpServiceGroupDetail) SetGroupType(v string) *PdpServiceGroupDetail {
	s.GroupType = &v
	return s
}

func (s *PdpServiceGroupDetail) SetId(v int64) *PdpServiceGroupDetail {
	s.Id = &v
	return s
}

func (s *PdpServiceGroupDetail) SetLogProject(v string) *PdpServiceGroupDetail {
	s.LogProject = &v
	return s
}

func (s *PdpServiceGroupDetail) SetLogStore(v string) *PdpServiceGroupDetail {
	s.LogStore = &v
	return s
}

func (s *PdpServiceGroupDetail) SetName(v string) *PdpServiceGroupDetail {
	s.Name = &v
	return s
}

func (s *PdpServiceGroupDetail) SetOrgType(v string) *PdpServiceGroupDetail {
	s.OrgType = &v
	return s
}

func (s *PdpServiceGroupDetail) SetPbcId(v int64) *PdpServiceGroupDetail {
	s.PbcId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetPipelineId(v string) *PdpServiceGroupDetail {
	s.PipelineId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetProductId(v int64) *PdpServiceGroupDetail {
	s.ProductId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetProductName(v string) *PdpServiceGroupDetail {
	s.ProductName = &v
	return s
}

func (s *PdpServiceGroupDetail) SetRegion(v string) *PdpServiceGroupDetail {
	s.Region = &v
	return s
}

func (s *PdpServiceGroupDetail) SetRequestId(v string) *PdpServiceGroupDetail {
	s.RequestId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetServiceId(v int64) *PdpServiceGroupDetail {
	s.ServiceId = &v
	return s
}

func (s *PdpServiceGroupDetail) SetType(v string) *PdpServiceGroupDetail {
	s.Type = &v
	return s
}

func (s *PdpServiceGroupDetail) Validate() error {
	return dara.Validate(s)
}
