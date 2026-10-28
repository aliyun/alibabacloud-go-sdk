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
	SetVscId(v string) *ListKVCacheStoreAttachInfoRequest
	GetVscId() *string
}

type ListKVCacheStoreAttachInfoRequest struct {
	// The list of KvcsIds of the KVCacheStores to query. You can specify up to 100 KvcsIds.
	//
	// This parameter is required.
	KvcsIds []*string `json:"KvcsIds,omitempty" xml:"KvcsIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return in a single call. Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the last call to this operation. You do not need to specify this parameter for the first call. If NextToken is specified, the PageSize and PageNumber request parameters become invalid, and the TotalCount value in the response is invalid.
	//
	// example:
	//
	// your-client-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for a paging query. This parameter is used together with PageSize. If the specified page number exceeds the total number of pages, the last page is returned.
	//
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paging query.
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
	// The ID of the compute-side VSC. If you specify this parameter, the mount records are filtered by the specified VSC.
	//
	// example:
	//
	// vsc-dhag651dsda
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
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

func (s *ListKVCacheStoreAttachInfoRequest) GetVscId() *string {
	return s.VscId
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

func (s *ListKVCacheStoreAttachInfoRequest) SetVscId(v string) *ListKVCacheStoreAttachInfoRequest {
	s.VscId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoRequest) Validate() error {
	return dara.Validate(s)
}
