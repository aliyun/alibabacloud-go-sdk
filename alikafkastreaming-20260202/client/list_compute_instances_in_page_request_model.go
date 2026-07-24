// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesInPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListComputeInstancesInPageRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListComputeInstancesInPageRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ListComputeInstancesInPageRequest
	GetInstanceIds() []*string
	SetOrderId(v string) *ListComputeInstancesInPageRequest
	GetOrderId() *string
	SetPageSize(v int32) *ListComputeInstancesInPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListComputeInstancesInPageRequest
	GetRegionId() *string
}

type ListComputeInstancesInPageRequest struct {
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OrderId     *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListComputeInstancesInPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageRequest) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListComputeInstancesInPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesInPageRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListComputeInstancesInPageRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ListComputeInstancesInPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeInstancesInPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesInPageRequest) SetCurrentPage(v int32) *ListComputeInstancesInPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetInstanceId(v string) *ListComputeInstancesInPageRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetInstanceIds(v []*string) *ListComputeInstancesInPageRequest {
	s.InstanceIds = v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetOrderId(v string) *ListComputeInstancesInPageRequest {
	s.OrderId = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetPageSize(v int32) *ListComputeInstancesInPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetRegionId(v string) *ListComputeInstancesInPageRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) Validate() error {
	return dara.Validate(s)
}
