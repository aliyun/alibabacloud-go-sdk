// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListLoadBalancerRegionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLoadBalancerRegionsResponseBody
	GetPageSize() *int32
	SetRegions(v []*ListLoadBalancerRegionsResponseBodyRegions) *ListLoadBalancerRegionsResponseBody
	GetRegions() []*ListLoadBalancerRegionsResponseBodyRegions
	SetRequestId(v string) *ListLoadBalancerRegionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLoadBalancerRegionsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListLoadBalancerRegionsResponseBody
	GetTotalPage() *int32
}

type ListLoadBalancerRegionsResponseBody struct {
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page
	//
	// example:
	//
	// 1024
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of region information
	Regions []*ListLoadBalancerRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 81A5E222-24BF-17EF-9E80-A68D9B8F363D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListLoadBalancerRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLoadBalancerRegionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLoadBalancerRegionsResponseBody) GetRegions() []*ListLoadBalancerRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListLoadBalancerRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLoadBalancerRegionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLoadBalancerRegionsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListLoadBalancerRegionsResponseBody) SetPageNumber(v int32) *ListLoadBalancerRegionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetPageSize(v int32) *ListLoadBalancerRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetRegions(v []*ListLoadBalancerRegionsResponseBodyRegions) *ListLoadBalancerRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetRequestId(v string) *ListLoadBalancerRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetTotalCount(v int32) *ListLoadBalancerRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetTotalPage(v int32) *ListLoadBalancerRegionsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLoadBalancerRegionsResponseBodyRegions struct {
	// Primary region Chinese full name
	//
	// example:
	//
	// 东南亚
	RegionCnName *string `json:"RegionCnName,omitempty" xml:"RegionCnName,omitempty"`
	// Primary region code
	//
	// example:
	//
	// SEAS
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// Primary region English full name
	//
	// example:
	//
	// South East Asia
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	// List of secondary region information
	SubRegions []*ListLoadBalancerRegionsResponseBodyRegionsSubRegions `json:"SubRegions,omitempty" xml:"SubRegions,omitempty" type:"Repeated"`
}

func (s ListLoadBalancerRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) GetRegionCnName() *string {
	return s.RegionCnName
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) GetRegionEnName() *string {
	return s.RegionEnName
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) GetSubRegions() []*ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	return s.SubRegions
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionCnName(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionCnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionCode(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionEnName(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetSubRegions(v []*ListLoadBalancerRegionsResponseBodyRegionsSubRegions) *ListLoadBalancerRegionsResponseBodyRegions {
	s.SubRegions = v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) Validate() error {
	if s.SubRegions != nil {
		for _, item := range s.SubRegions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLoadBalancerRegionsResponseBodyRegionsSubRegions struct {
	// Secondary region Chinese full name
	//
	// example:
	//
	// 印度尼西亚
	SubRegionCnName *string `json:"SubRegionCnName,omitempty" xml:"SubRegionCnName,omitempty"`
	// Secondary region code
	//
	// example:
	//
	// ID
	SubRegionCode *string `json:"SubRegionCode,omitempty" xml:"SubRegionCode,omitempty"`
	// Secondary region English full name
	//
	// example:
	//
	// Indonesia
	SubRegionEnName *string `json:"SubRegionEnName,omitempty" xml:"SubRegionEnName,omitempty"`
}

func (s ListLoadBalancerRegionsResponseBodyRegionsSubRegions) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBodyRegionsSubRegions) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) GetSubRegionCnName() *string {
	return s.SubRegionCnName
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) GetSubRegionCode() *string {
	return s.SubRegionCode
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) GetSubRegionEnName() *string {
	return s.SubRegionEnName
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionCnName(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionCnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionCode(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionEnName(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionEnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) Validate() error {
	return dara.Validate(s)
}
