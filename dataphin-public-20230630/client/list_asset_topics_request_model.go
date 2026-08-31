// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAssetTopicsRequestListQuery) *ListAssetTopicsRequest
	GetListQuery() *ListAssetTopicsRequestListQuery
	SetOpTenantId(v int64) *ListAssetTopicsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListAssetTopicsRequest
	GetOpUserId() *string
}

type ListAssetTopicsRequest struct {
	// The query parameters.
	//
	// This parameter is required.
	ListQuery *ListAssetTopicsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListAssetTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsRequest) GetListQuery() *ListAssetTopicsRequestListQuery {
	return s.ListQuery
}

func (s *ListAssetTopicsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAssetTopicsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListAssetTopicsRequest) SetListQuery(v *ListAssetTopicsRequestListQuery) *ListAssetTopicsRequest {
	s.ListQuery = v
	return s
}

func (s *ListAssetTopicsRequest) SetOpTenantId(v int64) *ListAssetTopicsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAssetTopicsRequest) SetOpUserId(v string) *ListAssetTopicsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListAssetTopicsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssetTopicsRequestListQuery struct {
	// The asset type. Valid values: TABLE, INDEX, API, DASHBOARD.
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The keyword for the topic name. Maximum length: 256 characters.
	//
	// example:
	//
	// Core Metrics
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 9770420
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Default value: 50. Valid values: 1 to 200.
	//
	// example:
	//
	// 7428337
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAssetTopicsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsRequestListQuery) GetAssetType() *string {
	return s.AssetType
}

func (s *ListAssetTopicsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAssetTopicsRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAssetTopicsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAssetTopicsRequestListQuery) SetAssetType(v string) *ListAssetTopicsRequestListQuery {
	s.AssetType = &v
	return s
}

func (s *ListAssetTopicsRequestListQuery) SetKeyword(v string) *ListAssetTopicsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListAssetTopicsRequestListQuery) SetPage(v int32) *ListAssetTopicsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAssetTopicsRequestListQuery) SetPageSize(v int32) *ListAssetTopicsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAssetTopicsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
