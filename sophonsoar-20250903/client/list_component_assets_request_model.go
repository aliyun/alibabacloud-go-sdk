// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssetUuid(v string) *ListComponentAssetsRequest
	GetComponentAssetUuid() *string
	SetComponentName(v string) *ListComponentAssetsRequest
	GetComponentName() *string
	SetLang(v string) *ListComponentAssetsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListComponentAssetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentAssetsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListComponentAssetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentAssetsRequest
	GetPageSize() *int32
	SetRoleFor(v int64) *ListComponentAssetsRequest
	GetRoleFor() *int64
}

type ListComponentAssetsRequest struct {
	// Asset UUID.
	//
	// example:
	//
	// 1C5F11E9-****-51F0-****-43BB312A0557
	ComponentAssetUuid *string `json:"ComponentAssetUuid,omitempty" xml:"ComponentAssetUuid,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// SLS
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language type for requests and responses. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Maximum number of results to return. Range: 0~100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next query. Value: Not required for the first query or if there is no next query. For subsequent queries, use the NextToken value returned from the previous API call.
	//
	// example:
	//
	// kt0BLbenY2XCyRfsmoEcVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Page number for paginated queries, with a default value of 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page for paginated queries. Range: 1~100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource directory member account ID.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListComponentAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentAssetsRequest) GetComponentAssetUuid() *string {
	return s.ComponentAssetUuid
}

func (s *ListComponentAssetsRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListComponentAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListComponentAssetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentAssetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentAssetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentAssetsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListComponentAssetsRequest) SetComponentAssetUuid(v string) *ListComponentAssetsRequest {
	s.ComponentAssetUuid = &v
	return s
}

func (s *ListComponentAssetsRequest) SetComponentName(v string) *ListComponentAssetsRequest {
	s.ComponentName = &v
	return s
}

func (s *ListComponentAssetsRequest) SetLang(v string) *ListComponentAssetsRequest {
	s.Lang = &v
	return s
}

func (s *ListComponentAssetsRequest) SetMaxResults(v int32) *ListComponentAssetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComponentAssetsRequest) SetNextToken(v string) *ListComponentAssetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListComponentAssetsRequest) SetPageNumber(v int32) *ListComponentAssetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentAssetsRequest) SetPageSize(v int32) *ListComponentAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentAssetsRequest) SetRoleFor(v int64) *ListComponentAssetsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListComponentAssetsRequest) Validate() error {
	return dara.Validate(s)
}
