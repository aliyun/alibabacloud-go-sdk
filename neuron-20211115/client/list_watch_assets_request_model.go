// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWatchAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListWatchAssetsRequest
	GetAccountId() *string
	SetAssetType(v string) *ListWatchAssetsRequest
	GetAssetType() *string
	SetCompanyId(v int64) *ListWatchAssetsRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListWatchAssetsRequest
	GetMarketId() *int64
	SetOrderBy(v string) *ListWatchAssetsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListWatchAssetsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListWatchAssetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWatchAssetsRequest
	GetPageSize() *int32
	SetUpshelfAssetId(v int64) *ListWatchAssetsRequest
	GetUpshelfAssetId() *int64
}

type ListWatchAssetsRequest struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AssetType      *string `json:"assetType,omitempty" xml:"assetType,omitempty"`
	CompanyId      *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UpshelfAssetId *int64  `json:"upshelfAssetId,omitempty" xml:"upshelfAssetId,omitempty"`
}

func (s ListWatchAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWatchAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListWatchAssetsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListWatchAssetsRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *ListWatchAssetsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListWatchAssetsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListWatchAssetsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWatchAssetsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListWatchAssetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWatchAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWatchAssetsRequest) GetUpshelfAssetId() *int64 {
	return s.UpshelfAssetId
}

func (s *ListWatchAssetsRequest) SetAccountId(v string) *ListWatchAssetsRequest {
	s.AccountId = &v
	return s
}

func (s *ListWatchAssetsRequest) SetAssetType(v string) *ListWatchAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *ListWatchAssetsRequest) SetCompanyId(v int64) *ListWatchAssetsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListWatchAssetsRequest) SetMarketId(v int64) *ListWatchAssetsRequest {
	s.MarketId = &v
	return s
}

func (s *ListWatchAssetsRequest) SetOrderBy(v string) *ListWatchAssetsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWatchAssetsRequest) SetOrderDirection(v string) *ListWatchAssetsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListWatchAssetsRequest) SetPageNumber(v int32) *ListWatchAssetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWatchAssetsRequest) SetPageSize(v int32) *ListWatchAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWatchAssetsRequest) SetUpshelfAssetId(v int64) *ListWatchAssetsRequest {
	s.UpshelfAssetId = &v
	return s
}

func (s *ListWatchAssetsRequest) Validate() error {
	return dara.Validate(s)
}
