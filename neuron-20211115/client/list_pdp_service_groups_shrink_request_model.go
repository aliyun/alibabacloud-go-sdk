// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServiceGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListPdpServiceGroupsShrinkRequest
	GetAlias() *string
	SetEnterpriseId(v int64) *ListPdpServiceGroupsShrinkRequest
	GetEnterpriseId() *int64
	SetEnv(v string) *ListPdpServiceGroupsShrinkRequest
	GetEnv() *string
	SetEnvType(v string) *ListPdpServiceGroupsShrinkRequest
	GetEnvType() *string
	SetGroupType(v string) *ListPdpServiceGroupsShrinkRequest
	GetGroupType() *string
	SetIdsShrink(v string) *ListPdpServiceGroupsShrinkRequest
	GetIdsShrink() *string
	SetName(v string) *ListPdpServiceGroupsShrinkRequest
	GetName() *string
	SetOrderBy(v string) *ListPdpServiceGroupsShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpServiceGroupsShrinkRequest
	GetOrderDirection() *string
	SetOrgType(v string) *ListPdpServiceGroupsShrinkRequest
	GetOrgType() *string
	SetPageNumber(v int32) *ListPdpServiceGroupsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpServiceGroupsShrinkRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListPdpServiceGroupsShrinkRequest
	GetPbcId() *int64
	SetProductId(v int64) *ListPdpServiceGroupsShrinkRequest
	GetProductId() *int64
	SetRegion(v string) *ListPdpServiceGroupsShrinkRequest
	GetRegion() *string
	SetRepoId(v string) *ListPdpServiceGroupsShrinkRequest
	GetRepoId() *string
	SetServiceId(v int64) *ListPdpServiceGroupsShrinkRequest
	GetServiceId() *int64
}

type ListPdpServiceGroupsShrinkRequest struct {
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// This parameter is required.
	//
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
	// SYSTEM
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	IdsShrink *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// gmt_create
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// INNER
	OrgType *string `json:"orgType,omitempty" xml:"orgType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// pbcId
	//
	// example:
	//
	// 64
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 22
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 12312
	RepoId *string `json:"repoId,omitempty" xml:"repoId,omitempty"`
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s ListPdpServiceGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServiceGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPdpServiceGroupsShrinkRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListPdpServiceGroupsShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListPdpServiceGroupsShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListPdpServiceGroupsShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListPdpServiceGroupsShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ListPdpServiceGroupsShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListPdpServiceGroupsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListPdpServiceGroupsShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpServiceGroupsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpServiceGroupsShrinkRequest) GetOrgType() *string {
	return s.OrgType
}

func (s *ListPdpServiceGroupsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpServiceGroupsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpServiceGroupsShrinkRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListPdpServiceGroupsShrinkRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListPdpServiceGroupsShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListPdpServiceGroupsShrinkRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListPdpServiceGroupsShrinkRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListPdpServiceGroupsShrinkRequest) SetAlias(v string) *ListPdpServiceGroupsShrinkRequest {
	s.Alias = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetEnterpriseId(v int64) *ListPdpServiceGroupsShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetEnv(v string) *ListPdpServiceGroupsShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetEnvType(v string) *ListPdpServiceGroupsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetGroupType(v string) *ListPdpServiceGroupsShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetIdsShrink(v string) *ListPdpServiceGroupsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetName(v string) *ListPdpServiceGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetOrderBy(v string) *ListPdpServiceGroupsShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetOrderDirection(v string) *ListPdpServiceGroupsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetOrgType(v string) *ListPdpServiceGroupsShrinkRequest {
	s.OrgType = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetPageNumber(v int32) *ListPdpServiceGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetPageSize(v int32) *ListPdpServiceGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetPbcId(v int64) *ListPdpServiceGroupsShrinkRequest {
	s.PbcId = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetProductId(v int64) *ListPdpServiceGroupsShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetRegion(v string) *ListPdpServiceGroupsShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetRepoId(v string) *ListPdpServiceGroupsShrinkRequest {
	s.RepoId = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) SetServiceId(v int64) *ListPdpServiceGroupsShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *ListPdpServiceGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
