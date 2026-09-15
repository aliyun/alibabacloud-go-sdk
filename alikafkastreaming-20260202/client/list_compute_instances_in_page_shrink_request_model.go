// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesInPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListComputeInstancesInPageShrinkRequest
	GetCurrentPage() *int32
	SetInstanceIdsShrink(v string) *ListComputeInstancesInPageShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageSize(v int32) *ListComputeInstancesInPageShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListComputeInstancesInPageShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListComputeInstancesInPageShrinkRequest
	GetResourceGroupId() *string
}

type ListComputeInstancesInPageShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListComputeInstancesInPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListComputeInstancesInPageShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListComputeInstancesInPageShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeInstancesInPageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesInPageShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListComputeInstancesInPageShrinkRequest) SetCurrentPage(v int32) *ListComputeInstancesInPageShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetInstanceIdsShrink(v string) *ListComputeInstancesInPageShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetPageSize(v int32) *ListComputeInstancesInPageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetRegionId(v string) *ListComputeInstancesInPageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetResourceGroupId(v string) *ListComputeInstancesInPageShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
