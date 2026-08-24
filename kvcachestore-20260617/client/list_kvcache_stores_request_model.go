// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsIds(v string) *ListKVCacheStoresRequest
	GetKvcsIds() *string
	SetMaxResults(v int32) *ListKVCacheStoresRequest
	GetMaxResults() *int32
	SetName(v string) *ListKVCacheStoresRequest
	GetName() *string
	SetNextToken(v string) *ListKVCacheStoresRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListKVCacheStoresRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKVCacheStoresRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKVCacheStoresRequest
	GetRegionId() *string
	SetStatus(v string) *ListKVCacheStoresRequest
	GetStatus() *string
	SetZoneId(v string) *ListKVCacheStoresRequest
	GetZoneId() *string
}

type ListKVCacheStoresRequest struct {
	// The list of KvcsId values. Separate multiple IDs with commas. A maximum of 100 IDs are supported.
	//
	// example:
	//
	// kvcs-87djda131
	KvcsIds *string `json:"KvcsIds,omitempty" xml:"KvcsIds,omitempty"`
	// The maximum number of entries per page for cursor-based pagination. Default value: 10. Maximum value: 100. This parameter is used together with NextToken.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The instance name filter. Prefix matching is used.
	//
	// example:
	//
	// obj-detect
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, use the NextToken value returned in the previous response. This parameter is mutually exclusive with PageNumber.
	//
	// example:
	//
	// a24c3a9cc8e6da77b10cffc4c93c7922e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1. This parameter takes precedence over NextToken if both are specified.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100. This parameter is used together with PageNumber.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status filter. Valid values: Creating, Available, InUse, Stopping, Stopped, and Deleting.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The zone ID, such as cn-hangzhou-a.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListKVCacheStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoresRequest) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoresRequest) GetKvcsIds() *string {
	return s.KvcsIds
}

func (s *ListKVCacheStoresRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKVCacheStoresRequest) GetName() *string {
	return s.Name
}

func (s *ListKVCacheStoresRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKVCacheStoresRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKVCacheStoresRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKVCacheStoresRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoresRequest) GetStatus() *string {
	return s.Status
}

func (s *ListKVCacheStoresRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListKVCacheStoresRequest) SetKvcsIds(v string) *ListKVCacheStoresRequest {
	s.KvcsIds = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetMaxResults(v int32) *ListKVCacheStoresRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetName(v string) *ListKVCacheStoresRequest {
	s.Name = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetNextToken(v string) *ListKVCacheStoresRequest {
	s.NextToken = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetPageNumber(v int32) *ListKVCacheStoresRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetPageSize(v int32) *ListKVCacheStoresRequest {
	s.PageSize = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetRegionId(v string) *ListKVCacheStoresRequest {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetStatus(v string) *ListKVCacheStoresRequest {
	s.Status = &v
	return s
}

func (s *ListKVCacheStoresRequest) SetZoneId(v string) *ListKVCacheStoresRequest {
	s.ZoneId = &v
	return s
}

func (s *ListKVCacheStoresRequest) Validate() error {
	return dara.Validate(s)
}
