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
	// This parameter is required.
	ListCatalogAssetsQuery *ListCatalogAssetsRequestListCatalogAssetsQuery `json:"ListCatalogAssetsQuery,omitempty" xml:"ListCatalogAssetsQuery,omitempty" type:"Struct"`
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
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
