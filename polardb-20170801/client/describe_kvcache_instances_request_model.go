// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKVCacheInstanceDescription(v string) *DescribeKVCacheInstancesRequest
	GetKVCacheInstanceDescription() *string
	SetKVCacheInstanceIds(v string) *DescribeKVCacheInstancesRequest
	GetKVCacheInstanceIds() *string
	SetPageNumber(v int32) *DescribeKVCacheInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKVCacheInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKVCacheInstancesRequest
	GetRegionId() *string
	SetRelativeDBClusterId(v string) *DescribeKVCacheInstancesRequest
	GetRelativeDBClusterId() *string
}

type DescribeKVCacheInstancesRequest struct {
	// The description of the KVCache instance.
	//
	// example:
	//
	// Test cache instance
	KVCacheInstanceDescription *string `json:"KVCacheInstanceDescription,omitempty" xml:"KVCacheInstanceDescription,omitempty"`
	// The KVCache instance IDs. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// pkv-123456,pkv-789012
	KVCacheInstanceIds *string `json:"KVCacheInstanceIds,omitempty" xml:"KVCacheInstanceIds,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the associated PolarDB instance.
	//
	// example:
	//
	// pc-123456
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
}

func (s DescribeKVCacheInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstancesRequest) GetKVCacheInstanceDescription() *string {
	return s.KVCacheInstanceDescription
}

func (s *DescribeKVCacheInstancesRequest) GetKVCacheInstanceIds() *string {
	return s.KVCacheInstanceIds
}

func (s *DescribeKVCacheInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKVCacheInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKVCacheInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKVCacheInstancesRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeKVCacheInstancesRequest) SetKVCacheInstanceDescription(v string) *DescribeKVCacheInstancesRequest {
	s.KVCacheInstanceDescription = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) SetKVCacheInstanceIds(v string) *DescribeKVCacheInstancesRequest {
	s.KVCacheInstanceIds = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) SetPageNumber(v int32) *DescribeKVCacheInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) SetPageSize(v int32) *DescribeKVCacheInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) SetRegionId(v string) *DescribeKVCacheInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) SetRelativeDBClusterId(v string) *DescribeKVCacheInstancesRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeKVCacheInstancesRequest) Validate() error {
	return dara.Validate(s)
}
