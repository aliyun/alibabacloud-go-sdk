// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServiceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListPdpServiceGroupsRequest
	GetAlias() *string
	SetEnterpriseId(v int64) *ListPdpServiceGroupsRequest
	GetEnterpriseId() *int64
	SetEnv(v string) *ListPdpServiceGroupsRequest
	GetEnv() *string
	SetEnvType(v string) *ListPdpServiceGroupsRequest
	GetEnvType() *string
	SetGroupType(v string) *ListPdpServiceGroupsRequest
	GetGroupType() *string
	SetIds(v []*int64) *ListPdpServiceGroupsRequest
	GetIds() []*int64
	SetName(v string) *ListPdpServiceGroupsRequest
	GetName() *string
	SetOrderBy(v string) *ListPdpServiceGroupsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpServiceGroupsRequest
	GetOrderDirection() *string
	SetOrgType(v string) *ListPdpServiceGroupsRequest
	GetOrgType() *string
	SetPageNumber(v int32) *ListPdpServiceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpServiceGroupsRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListPdpServiceGroupsRequest
	GetPbcId() *int64
	SetProductId(v int64) *ListPdpServiceGroupsRequest
	GetProductId() *int64
	SetRegion(v string) *ListPdpServiceGroupsRequest
	GetRegion() *string
	SetRepoId(v string) *ListPdpServiceGroupsRequest
	GetRepoId() *string
	SetServiceId(v int64) *ListPdpServiceGroupsRequest
	GetServiceId() *int64
}

type ListPdpServiceGroupsRequest struct {
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
	GroupType *string  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Ids       []*int64 `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
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

func (s ListPdpServiceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPdpServiceGroupsRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListPdpServiceGroupsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListPdpServiceGroupsRequest) GetEnv() *string {
	return s.Env
}

func (s *ListPdpServiceGroupsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListPdpServiceGroupsRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ListPdpServiceGroupsRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListPdpServiceGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListPdpServiceGroupsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpServiceGroupsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpServiceGroupsRequest) GetOrgType() *string {
	return s.OrgType
}

func (s *ListPdpServiceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpServiceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpServiceGroupsRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListPdpServiceGroupsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListPdpServiceGroupsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListPdpServiceGroupsRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListPdpServiceGroupsRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListPdpServiceGroupsRequest) SetAlias(v string) *ListPdpServiceGroupsRequest {
	s.Alias = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetEnterpriseId(v int64) *ListPdpServiceGroupsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetEnv(v string) *ListPdpServiceGroupsRequest {
	s.Env = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetEnvType(v string) *ListPdpServiceGroupsRequest {
	s.EnvType = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetGroupType(v string) *ListPdpServiceGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetIds(v []*int64) *ListPdpServiceGroupsRequest {
	s.Ids = v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetName(v string) *ListPdpServiceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetOrderBy(v string) *ListPdpServiceGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetOrderDirection(v string) *ListPdpServiceGroupsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetOrgType(v string) *ListPdpServiceGroupsRequest {
	s.OrgType = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetPageNumber(v int32) *ListPdpServiceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetPageSize(v int32) *ListPdpServiceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetPbcId(v int64) *ListPdpServiceGroupsRequest {
	s.PbcId = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetProductId(v int64) *ListPdpServiceGroupsRequest {
	s.ProductId = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetRegion(v string) *ListPdpServiceGroupsRequest {
	s.Region = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetRepoId(v string) *ListPdpServiceGroupsRequest {
	s.RepoId = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) SetServiceId(v int64) *ListPdpServiceGroupsRequest {
	s.ServiceId = &v
	return s
}

func (s *ListPdpServiceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
