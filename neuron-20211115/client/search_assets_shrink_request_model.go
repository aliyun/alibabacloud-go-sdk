// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetIndustrysShrink(v string) *SearchAssetsShrinkRequest
	GetAssetIndustrysShrink() *string
	SetAssetName(v string) *SearchAssetsShrinkRequest
	GetAssetName() *string
	SetAssetTypesShrink(v string) *SearchAssetsShrinkRequest
	GetAssetTypesShrink() *string
	SetCompanyId(v int64) *SearchAssetsShrinkRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *SearchAssetsShrinkRequest
	GetMarketId() *int64
	SetOrderBy(v string) *SearchAssetsShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchAssetsShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *SearchAssetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchAssetsShrinkRequest
	GetPageSize() *int32
}

type SearchAssetsShrinkRequest struct {
	AssetIndustrysShrink *string `json:"assetIndustrys,omitempty" xml:"assetIndustrys,omitempty"`
	AssetName            *string `json:"assetName,omitempty" xml:"assetName,omitempty"`
	AssetTypesShrink     *string `json:"assetTypes,omitempty" xml:"assetTypes,omitempty"`
	CompanyId            *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId             *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy              *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection       *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber           *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize             *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s SearchAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAssetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchAssetsShrinkRequest) GetAssetIndustrysShrink() *string {
	return s.AssetIndustrysShrink
}

func (s *SearchAssetsShrinkRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *SearchAssetsShrinkRequest) GetAssetTypesShrink() *string {
	return s.AssetTypesShrink
}

func (s *SearchAssetsShrinkRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *SearchAssetsShrinkRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *SearchAssetsShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchAssetsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchAssetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAssetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAssetsShrinkRequest) SetAssetIndustrysShrink(v string) *SearchAssetsShrinkRequest {
	s.AssetIndustrysShrink = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetAssetName(v string) *SearchAssetsShrinkRequest {
	s.AssetName = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetAssetTypesShrink(v string) *SearchAssetsShrinkRequest {
	s.AssetTypesShrink = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetCompanyId(v int64) *SearchAssetsShrinkRequest {
	s.CompanyId = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetMarketId(v int64) *SearchAssetsShrinkRequest {
	s.MarketId = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetOrderBy(v string) *SearchAssetsShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetOrderDirection(v string) *SearchAssetsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetPageNumber(v int32) *SearchAssetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchAssetsShrinkRequest) SetPageSize(v int32) *SearchAssetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAssetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
