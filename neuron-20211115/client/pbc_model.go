// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbc interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *Pbc
	GetAlias() *string
	SetAssetType(v string) *Pbc
	GetAssetType() *string
	SetCompany(v string) *Pbc
	GetCompany() *string
	SetCompanyId(v int64) *Pbc
	GetCompanyId() *int64
	SetDescription(v string) *Pbc
	GetDescription() *string
	SetDeveloperId(v string) *Pbc
	GetDeveloperId() *string
	SetDeveloperName(v string) *Pbc
	GetDeveloperName() *string
	SetForkCount(v int32) *Pbc
	GetForkCount() *int32
	SetId(v int64) *Pbc
	GetId() *int64
	SetIndustry(v string) *Pbc
	GetIndustry() *string
	SetInvorkCount(v int64) *Pbc
	GetInvorkCount() *int64
	SetIsWatched(v bool) *Pbc
	GetIsWatched() *bool
	SetLatestVersionId(v int64) *Pbc
	GetLatestVersionId() *int64
	SetMarketId(v int64) *Pbc
	GetMarketId() *int64
	SetName(v string) *Pbc
	GetName() *string
	SetRequestId(v string) *Pbc
	GetRequestId() *string
	SetType(v string) *Pbc
	GetType() *string
	SetWatchCount(v int64) *Pbc
	GetWatchCount() *int64
}

type Pbc struct {
	// example:
	//
	// 基础商品
	Alias     *string `json:"alias,omitempty" xml:"alias,omitempty"`
	AssetType *string `json:"assetType,omitempty" xml:"assetType,omitempty"`
	// example:
	//
	// 企业服务
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// example:
	//
	// 12
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 基础商品PBC
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	DeveloperId   *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	DeveloperName *string `json:"developerName,omitempty" xml:"developerName,omitempty"`
	// example:
	//
	// 1
	ForkCount *int32 `json:"forkCount,omitempty" xml:"forkCount,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// common
	Industry        *string `json:"industry,omitempty" xml:"industry,omitempty"`
	InvorkCount     *int64  `json:"invorkCount,omitempty" xml:"invorkCount,omitempty"`
	IsWatched       *bool   `json:"isWatched,omitempty" xml:"isWatched,omitempty"`
	LatestVersionId *int64  `json:"latestVersionId,omitempty" xml:"latestVersionId,omitempty"`
	MarketId        *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	// example:
	//
	// product
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// foundation
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	WatchCount *int64  `json:"watchCount,omitempty" xml:"watchCount,omitempty"`
}

func (s Pbc) String() string {
	return dara.Prettify(s)
}

func (s Pbc) GoString() string {
	return s.String()
}

func (s *Pbc) GetAlias() *string {
	return s.Alias
}

func (s *Pbc) GetAssetType() *string {
	return s.AssetType
}

func (s *Pbc) GetCompany() *string {
	return s.Company
}

func (s *Pbc) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *Pbc) GetDescription() *string {
	return s.Description
}

func (s *Pbc) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *Pbc) GetDeveloperName() *string {
	return s.DeveloperName
}

func (s *Pbc) GetForkCount() *int32 {
	return s.ForkCount
}

func (s *Pbc) GetId() *int64 {
	return s.Id
}

func (s *Pbc) GetIndustry() *string {
	return s.Industry
}

func (s *Pbc) GetInvorkCount() *int64 {
	return s.InvorkCount
}

func (s *Pbc) GetIsWatched() *bool {
	return s.IsWatched
}

func (s *Pbc) GetLatestVersionId() *int64 {
	return s.LatestVersionId
}

func (s *Pbc) GetMarketId() *int64 {
	return s.MarketId
}

func (s *Pbc) GetName() *string {
	return s.Name
}

func (s *Pbc) GetRequestId() *string {
	return s.RequestId
}

func (s *Pbc) GetType() *string {
	return s.Type
}

func (s *Pbc) GetWatchCount() *int64 {
	return s.WatchCount
}

func (s *Pbc) SetAlias(v string) *Pbc {
	s.Alias = &v
	return s
}

func (s *Pbc) SetAssetType(v string) *Pbc {
	s.AssetType = &v
	return s
}

func (s *Pbc) SetCompany(v string) *Pbc {
	s.Company = &v
	return s
}

func (s *Pbc) SetCompanyId(v int64) *Pbc {
	s.CompanyId = &v
	return s
}

func (s *Pbc) SetDescription(v string) *Pbc {
	s.Description = &v
	return s
}

func (s *Pbc) SetDeveloperId(v string) *Pbc {
	s.DeveloperId = &v
	return s
}

func (s *Pbc) SetDeveloperName(v string) *Pbc {
	s.DeveloperName = &v
	return s
}

func (s *Pbc) SetForkCount(v int32) *Pbc {
	s.ForkCount = &v
	return s
}

func (s *Pbc) SetId(v int64) *Pbc {
	s.Id = &v
	return s
}

func (s *Pbc) SetIndustry(v string) *Pbc {
	s.Industry = &v
	return s
}

func (s *Pbc) SetInvorkCount(v int64) *Pbc {
	s.InvorkCount = &v
	return s
}

func (s *Pbc) SetIsWatched(v bool) *Pbc {
	s.IsWatched = &v
	return s
}

func (s *Pbc) SetLatestVersionId(v int64) *Pbc {
	s.LatestVersionId = &v
	return s
}

func (s *Pbc) SetMarketId(v int64) *Pbc {
	s.MarketId = &v
	return s
}

func (s *Pbc) SetName(v string) *Pbc {
	s.Name = &v
	return s
}

func (s *Pbc) SetRequestId(v string) *Pbc {
	s.RequestId = &v
	return s
}

func (s *Pbc) SetType(v string) *Pbc {
	s.Type = &v
	return s
}

func (s *Pbc) SetWatchCount(v int64) *Pbc {
	s.WatchCount = &v
	return s
}

func (s *Pbc) Validate() error {
	return dara.Validate(s)
}
