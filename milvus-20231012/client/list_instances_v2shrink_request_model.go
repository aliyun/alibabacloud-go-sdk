// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListInstancesV2ShrinkRequest
	GetRegionId() *string
	SetInstanceId(v string) *ListInstancesV2ShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListInstancesV2ShrinkRequest
	GetInstanceName() *string
	SetMaxResults(v int32) *ListInstancesV2ShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesV2ShrinkRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListInstancesV2ShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesV2ShrinkRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesV2ShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListInstancesV2ShrinkRequest
	GetTagShrink() *string
}

type ListInstancesV2ShrinkRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// milvus-test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// rg-123xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	TagShrink       *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListInstancesV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesV2ShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesV2ShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesV2ShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesV2ShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesV2ShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesV2ShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesV2ShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesV2ShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListInstancesV2ShrinkRequest) SetRegionId(v string) *ListInstancesV2ShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetInstanceId(v string) *ListInstancesV2ShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetInstanceName(v string) *ListInstancesV2ShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetMaxResults(v int32) *ListInstancesV2ShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetNextToken(v string) *ListInstancesV2ShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetPageNumber(v int32) *ListInstancesV2ShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetPageSize(v int32) *ListInstancesV2ShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetResourceGroupId(v string) *ListInstancesV2ShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) SetTagShrink(v string) *ListInstancesV2ShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListInstancesV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
