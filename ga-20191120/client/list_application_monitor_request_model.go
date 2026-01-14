// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListApplicationMonitorRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationMonitorRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListApplicationMonitorRequest
	GetRegionId() *string
	SetSearchValue(v string) *ListApplicationMonitorRequest
	GetSearchValue() *string
}

type ListApplicationMonitorRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// １
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that is used to search for origin probing tasks. You can enter a URL, an IP address, a GA instance ID, or a listener ID to perform a fuzzy match.
	//
	// example:
	//
	// aliyun
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s ListApplicationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationMonitorRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationMonitorRequest) GetSearchValue() *string {
	return s.SearchValue
}

func (s *ListApplicationMonitorRequest) SetPageNumber(v int32) *ListApplicationMonitorRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetPageSize(v int32) *ListApplicationMonitorRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetRegionId(v string) *ListApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetSearchValue(v string) *ListApplicationMonitorRequest {
	s.SearchValue = &v
	return s
}

func (s *ListApplicationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
