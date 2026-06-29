// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListCatalogAssetsQuery(v *ListCatalogAssetsRequestListCatalogAssetsQuery) *ListCatalogAssetsRequest
	GetListCatalogAssetsQuery() *ListCatalogAssetsRequestListCatalogAssetsQuery
	SetOpTenantId(v int64) *ListCatalogAssetsRequest
	GetOpTenantId() *int64
}

type ListCatalogAssetsRequest struct {
	// The query parameters.
	//
	// This parameter is required.
	ListCatalogAssetsQuery *ListCatalogAssetsRequestListCatalogAssetsQuery `json:"ListCatalogAssetsQuery,omitempty" xml:"ListCatalogAssetsQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListCatalogAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsRequest) GetListCatalogAssetsQuery() *ListCatalogAssetsRequestListCatalogAssetsQuery {
	return s.ListCatalogAssetsQuery
}

func (s *ListCatalogAssetsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListCatalogAssetsRequest) SetListCatalogAssetsQuery(v *ListCatalogAssetsRequestListCatalogAssetsQuery) *ListCatalogAssetsRequest {
	s.ListCatalogAssetsQuery = v
	return s
}

func (s *ListCatalogAssetsRequest) SetOpTenantId(v int64) *ListCatalogAssetsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListCatalogAssetsRequest) Validate() error {
	if s.ListCatalogAssetsQuery != nil {
		if err := s.ListCatalogAssetsQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCatalogAssetsRequestListCatalogAssetsQuery struct {
	// The asset type. Default value: TABLE. Valid values:
	//
	// - TABLE: table, including views and materialized views.
	//
	// - INDEX: technical metric.
	//
	// - BIZ_INDEX: business metric.
	//
	// - API: API.
	//
	// - PAGE: dashboard.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The search keyword. Used when queryMode is set to ASSET_SEARCH. Supports keyword matching against the asset full name, asset name, asset display name, and asset description. If this parameter is not specified, all assets are queried.
	//
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The asset name. Used when queryMode is set to EXACT_MATCH. If this parameter is not specified, all assets are queried.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query type. Determines whether to use name for exact matching or keyword for fuzzy search. Default value: EXACT_MATCH. Valid values:
	//
	// - EXACT_MATCH: exact match.
	//
	// - ASSET_SEARCH: fuzzy search.
	//
	// example:
	//
	// EXACT_MATCH
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
}

func (s ListCatalogAssetsRequestListCatalogAssetsQuery) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsRequestListCatalogAssetsQuery) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetAssetType() *string {
	return s.AssetType
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetName() *string {
	return s.Name
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) GetQueryMode() *string {
	return s.QueryMode
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetAssetType(v string) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.AssetType = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetKeyword(v string) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.Keyword = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetName(v string) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.Name = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetPageNum(v int32) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.PageNum = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetPageSize(v int32) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.PageSize = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) SetQueryMode(v string) *ListCatalogAssetsRequestListCatalogAssetsQuery {
	s.QueryMode = &v
	return s
}

func (s *ListCatalogAssetsRequestListCatalogAssetsQuery) Validate() error {
	return dara.Validate(s)
}
