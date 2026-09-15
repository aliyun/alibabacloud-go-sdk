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
	SetInstanceIds(v []*string) *ListComputeInstancesInPageRequest
	GetInstanceIds() []*string
	SetPageSize(v int32) *ListComputeInstancesInPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListComputeInstancesInPageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListComputeInstancesInPageRequest
	GetResourceGroupId() *string
}

type ListComputeInstancesInPageRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s ListComputeInstancesInPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageRequest) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListComputeInstancesInPageRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListComputeInstancesInPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeInstancesInPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesInPageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListComputeInstancesInPageRequest) SetCurrentPage(v int32) *ListComputeInstancesInPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) SetInstanceIds(v []*string) *ListComputeInstancesInPageRequest {
	s.InstanceIds = v
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

func (s *ListComputeInstancesInPageRequest) SetResourceGroupId(v string) *ListComputeInstancesInPageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListComputeInstancesInPageRequest) Validate() error {
	return dara.Validate(s)
}
