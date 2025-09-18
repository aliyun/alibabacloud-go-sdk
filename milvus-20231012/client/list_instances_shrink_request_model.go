// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstancesShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *ListInstancesShrinkRequest
	GetClusterName() *string
	SetPageNumber(v int32) *ListInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListInstancesShrinkRequest
	GetTagShrink() *string
}

type ListInstancesShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-123xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-123xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstancesShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListInstancesShrinkRequest) SetClusterId(v string) *ListInstancesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetClusterName(v string) *ListInstancesShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetRegionId(v string) *ListInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagShrink(v string) *ListInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
