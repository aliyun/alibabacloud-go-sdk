// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListInstancesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstancesRequest
	GetRegionId() *string
	SetTenantIsolationType(v string) *ListInstancesRequest
	GetTenantIsolationType() *string
}

type ListInstancesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of HSMs that is classified based on resource isolation. Valid values:
	//
	// - vsm: Virtual security modules (VSMs).
	//
	// - hostedHsm: Dedicated HSMs.
	//
	// example:
	//
	// vsm
	TenantIsolationType *string `json:"TenantIsolationType,omitempty" xml:"TenantIsolationType,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesRequest) GetTenantIsolationType() *string {
	return s.TenantIsolationType
}

func (s *ListInstancesRequest) SetCurrentPage(v int32) *ListInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetTenantIsolationType(v string) *ListInstancesRequest {
	s.TenantIsolationType = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
