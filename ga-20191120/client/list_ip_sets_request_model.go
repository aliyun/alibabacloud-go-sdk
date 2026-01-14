// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListIpSetsRequest
	GetAcceleratorId() *string
	SetPageNumber(v int32) *ListIpSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIpSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListIpSetsRequest
	GetRegionId() *string
}

type ListIpSetsRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmqy****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpSetsRequest) GoString() string {
	return s.String()
}

func (s *ListIpSetsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListIpSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIpSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIpSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpSetsRequest) SetAcceleratorId(v string) *ListIpSetsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListIpSetsRequest) SetPageNumber(v int32) *ListIpSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIpSetsRequest) SetPageSize(v int32) *ListIpSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIpSetsRequest) SetRegionId(v string) *ListIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpSetsRequest) Validate() error {
	return dara.Validate(s)
}
