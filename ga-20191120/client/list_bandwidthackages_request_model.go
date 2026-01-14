// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListBandwidthackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBandwidthackagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListBandwidthackagesRequest
	GetRegionId() *string
}

type ListBandwidthackagesRequest struct {
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

func (s ListBandwidthackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBandwidthackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBandwidthackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBandwidthackagesRequest) SetPageNumber(v int32) *ListBandwidthackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetPageSize(v int32) *ListBandwidthackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetRegionId(v string) *ListBandwidthackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthackagesRequest) Validate() error {
	return dara.Validate(s)
}
