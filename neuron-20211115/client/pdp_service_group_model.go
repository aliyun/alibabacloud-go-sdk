// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpServiceGroup
	GetAccountId() *string
	SetAlias(v string) *PdpServiceGroup
	GetAlias() *string
	SetDescription(v string) *PdpServiceGroup
	GetDescription() *string
	SetEdasId(v string) *PdpServiceGroup
	GetEdasId() *string
	SetEnterpriseId(v int64) *PdpServiceGroup
	GetEnterpriseId() *int64
	SetEnv(v string) *PdpServiceGroup
	GetEnv() *string
	SetEnvType(v string) *PdpServiceGroup
	GetEnvType() *string
	SetGmtCreate(v string) *PdpServiceGroup
	GetGmtCreate() *string
	SetGmtModified(v string) *PdpServiceGroup
	GetGmtModified() *string
	SetGroupType(v string) *PdpServiceGroup
	GetGroupType() *string
	SetId(v int64) *PdpServiceGroup
	GetId() *int64
	SetName(v string) *PdpServiceGroup
	GetName() *string
	SetOrgType(v string) *PdpServiceGroup
	GetOrgType() *string
	SetPbcId(v int64) *PdpServiceGroup
	GetPbcId() *int64
	SetPipelineId(v string) *PdpServiceGroup
	GetPipelineId() *string
	SetProductId(v int64) *PdpServiceGroup
	GetProductId() *int64
	SetProductName(v string) *PdpServiceGroup
	GetProductName() *string
	SetRegion(v string) *PdpServiceGroup
	GetRegion() *string
	SetRequestId(v string) *PdpServiceGroup
	GetRequestId() *string
	SetServiceId(v int64) *PdpServiceGroup
	GetServiceId() *int64
	SetType(v string) *PdpServiceGroup
	GetType() *string
}

type PdpServiceGroup struct {
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

func (s PdpServiceGroup) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceGroup) GoString() string {
	return s.String()
}

func (s *PdpServiceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpServiceGroup) GetAlias() *string {
	return s.Alias
}

func (s *PdpServiceGroup) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceGroup) GetEdasId() *string {
	return s.EdasId
}

func (s *PdpServiceGroup) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpServiceGroup) GetEnv() *string {
	return s.Env
}

func (s *PdpServiceGroup) GetEnvType() *string {
	return s.EnvType
}

func (s *PdpServiceGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpServiceGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PdpServiceGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *PdpServiceGroup) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceGroup) GetName() *string {
	return s.Name
}

func (s *PdpServiceGroup) GetOrgType() *string {
	return s.OrgType
}

func (s *PdpServiceGroup) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpServiceGroup) GetPipelineId() *string {
	return s.PipelineId
}

func (s *PdpServiceGroup) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpServiceGroup) GetProductName() *string {
	return s.ProductName
}

func (s *PdpServiceGroup) GetRegion() *string {
	return s.Region
}

func (s *PdpServiceGroup) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpServiceGroup) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpServiceGroup) GetType() *string {
	return s.Type
}

func (s *PdpServiceGroup) SetAccountId(v string) *PdpServiceGroup {
	s.AccountId = &v
	return s
}

func (s *PdpServiceGroup) SetAlias(v string) *PdpServiceGroup {
	s.Alias = &v
	return s
}

func (s *PdpServiceGroup) SetDescription(v string) *PdpServiceGroup {
	s.Description = &v
	return s
}

func (s *PdpServiceGroup) SetEdasId(v string) *PdpServiceGroup {
	s.EdasId = &v
	return s
}

func (s *PdpServiceGroup) SetEnterpriseId(v int64) *PdpServiceGroup {
	s.EnterpriseId = &v
	return s
}

func (s *PdpServiceGroup) SetEnv(v string) *PdpServiceGroup {
	s.Env = &v
	return s
}

func (s *PdpServiceGroup) SetEnvType(v string) *PdpServiceGroup {
	s.EnvType = &v
	return s
}

func (s *PdpServiceGroup) SetGmtCreate(v string) *PdpServiceGroup {
	s.GmtCreate = &v
	return s
}

func (s *PdpServiceGroup) SetGmtModified(v string) *PdpServiceGroup {
	s.GmtModified = &v
	return s
}

func (s *PdpServiceGroup) SetGroupType(v string) *PdpServiceGroup {
	s.GroupType = &v
	return s
}

func (s *PdpServiceGroup) SetId(v int64) *PdpServiceGroup {
	s.Id = &v
	return s
}

func (s *PdpServiceGroup) SetName(v string) *PdpServiceGroup {
	s.Name = &v
	return s
}

func (s *PdpServiceGroup) SetOrgType(v string) *PdpServiceGroup {
	s.OrgType = &v
	return s
}

func (s *PdpServiceGroup) SetPbcId(v int64) *PdpServiceGroup {
	s.PbcId = &v
	return s
}

func (s *PdpServiceGroup) SetPipelineId(v string) *PdpServiceGroup {
	s.PipelineId = &v
	return s
}

func (s *PdpServiceGroup) SetProductId(v int64) *PdpServiceGroup {
	s.ProductId = &v
	return s
}

func (s *PdpServiceGroup) SetProductName(v string) *PdpServiceGroup {
	s.ProductName = &v
	return s
}

func (s *PdpServiceGroup) SetRegion(v string) *PdpServiceGroup {
	s.Region = &v
	return s
}

func (s *PdpServiceGroup) SetRequestId(v string) *PdpServiceGroup {
	s.RequestId = &v
	return s
}

func (s *PdpServiceGroup) SetServiceId(v int64) *PdpServiceGroup {
	s.ServiceId = &v
	return s
}

func (s *PdpServiceGroup) SetType(v string) *PdpServiceGroup {
	s.Type = &v
	return s
}

func (s *PdpServiceGroup) Validate() error {
	return dara.Validate(s)
}
