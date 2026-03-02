// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetIndustrys(v []*string) *SearchAssetsRequest
	GetAssetIndustrys() []*string
	SetAssetName(v string) *SearchAssetsRequest
	GetAssetName() *string
	SetAssetTypes(v []*string) *SearchAssetsRequest
	GetAssetTypes() []*string
	SetCompanyId(v int64) *SearchAssetsRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *SearchAssetsRequest
	GetMarketId() *int64
	SetOrderBy(v string) *SearchAssetsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchAssetsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *SearchAssetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchAssetsRequest
	GetPageSize() *int32
}

type SearchAssetsRequest struct {
	AssetIndustrys []*string `json:"assetIndustrys,omitempty" xml:"assetIndustrys,omitempty" type:"Repeated"`
	AssetName      *string   `json:"assetName,omitempty" xml:"assetName,omitempty"`
	AssetTypes     []*string `json:"assetTypes,omitempty" xml:"assetTypes,omitempty" type:"Repeated"`
	CompanyId      *int64    `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId       *int64    `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string   `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string   `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s SearchAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAssetsRequest) GoString() string {
	return s.String()
}

func (s *SearchAssetsRequest) GetAssetIndustrys() []*string {
	return s.AssetIndustrys
}

func (s *SearchAssetsRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *SearchAssetsRequest) GetAssetTypes() []*string {
	return s.AssetTypes
}

func (s *SearchAssetsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *SearchAssetsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *SearchAssetsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchAssetsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchAssetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAssetsRequest) SetAssetIndustrys(v []*string) *SearchAssetsRequest {
	s.AssetIndustrys = v
	return s
}

func (s *SearchAssetsRequest) SetAssetName(v string) *SearchAssetsRequest {
	s.AssetName = &v
	return s
}

func (s *SearchAssetsRequest) SetAssetTypes(v []*string) *SearchAssetsRequest {
	s.AssetTypes = v
	return s
}

func (s *SearchAssetsRequest) SetCompanyId(v int64) *SearchAssetsRequest {
	s.CompanyId = &v
	return s
}

func (s *SearchAssetsRequest) SetMarketId(v int64) *SearchAssetsRequest {
	s.MarketId = &v
	return s
}

func (s *SearchAssetsRequest) SetOrderBy(v string) *SearchAssetsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchAssetsRequest) SetOrderDirection(v string) *SearchAssetsRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchAssetsRequest) SetPageNumber(v int32) *SearchAssetsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchAssetsRequest) SetPageSize(v int32) *SearchAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAssetsRequest) Validate() error {
	return dara.Validate(s)
}
