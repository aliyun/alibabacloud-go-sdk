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
	SetInstanceId(v string) *ListComputeInstancesInPageShrinkRequest
	GetInstanceId() *string
	SetInstanceIdsShrink(v string) *ListComputeInstancesInPageShrinkRequest
	GetInstanceIdsShrink() *string
	SetOrderId(v string) *ListComputeInstancesInPageShrinkRequest
	GetOrderId() *string
	SetPageSize(v int32) *ListComputeInstancesInPageShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListComputeInstancesInPageShrinkRequest
	GetRegionId() *string
}

type ListComputeInstancesInPageShrinkRequest struct {
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListComputeInstancesInPageShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesInPageShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListComputeInstancesInPageShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ListComputeInstancesInPageShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeInstancesInPageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesInPageShrinkRequest) SetCurrentPage(v int32) *ListComputeInstancesInPageShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetInstanceId(v string) *ListComputeInstancesInPageShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetInstanceIdsShrink(v string) *ListComputeInstancesInPageShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListComputeInstancesInPageShrinkRequest) SetOrderId(v string) *ListComputeInstancesInPageShrinkRequest {
	s.OrderId = &v
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

func (s *ListComputeInstancesInPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
