// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAttachInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsIds(v []*string) *ListKVCacheStoreAttachInfoRequest
	GetKvcsIds() []*string
	SetMaxResults(v int64) *ListKVCacheStoreAttachInfoRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListKVCacheStoreAttachInfoRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListKVCacheStoreAttachInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKVCacheStoreAttachInfoRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKVCacheStoreAttachInfoRequest
	GetRegionId() *string
}

type ListKVCacheStoreAttachInfoRequest struct {
	// The list of KVCacheStore KvcsId values to query. A maximum of 100 values can be specified.
	//
	// This parameter is required.
	KvcsIds []*string `json:"KvcsIds,omitempty" xml:"KvcsIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return in a single request. Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request. If you set NextToken, the PageSize and PageNumber request parameters become ineffective, and the TotalCount value in the response is invalid.
	//
	// example:
	//
	// your-client-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for a paged query. Used together with PageSize. If the value exceeds the total number of pages, the last page of data is returned.
	//
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKVCacheStoreAttachInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAttachInfoRequest) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAttachInfoRequest) GetKvcsIds() []*string {
	return s.KvcsIds
}

func (s *ListKVCacheStoreAttachInfoRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListKVCacheStoreAttachInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKVCacheStoreAttachInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKVCacheStoreAttachInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKVCacheStoreAttachInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoreAttachInfoRequest) SetKvcsIds(v []*string) *ListKVCacheStoreAttachInfoRequest {
	s.KvcsIds = v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) SetMaxResults(v int64) *ListKVCacheStoreAttachInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) SetNextToken(v string) *ListKVCacheStoreAttachInfoRequest {
	s.NextToken = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) SetPageNumber(v int32) *ListKVCacheStoreAttachInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) SetPageSize(v int32) *ListKVCacheStoreAttachInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) SetRegionId(v string) *ListKVCacheStoreAttachInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) Validate() error {
	return dara.Validate(s)
}
