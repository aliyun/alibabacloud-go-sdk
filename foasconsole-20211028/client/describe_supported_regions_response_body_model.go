// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *DescribeSupportedRegionsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeSupportedRegionsResponseBody
	GetPageSize() *int32
	SetRegions(v []*DescribeSupportedRegionsResponseBodyRegions) *DescribeSupportedRegionsResponseBody
	GetRegions() []*DescribeSupportedRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeSupportedRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSupportedRegionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeSupportedRegionsResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int32) *DescribeSupportedRegionsResponseBody
	GetTotalPage() *int32
}

type DescribeSupportedRegionsResponseBody struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Regions  []*DescribeSupportedRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// B21DC47E-8928-199A-9F32-36D45E4693B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeSupportedRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeSupportedRegionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSupportedRegionsResponseBody) GetRegions() []*DescribeSupportedRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeSupportedRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportedRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSupportedRegionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSupportedRegionsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSupportedRegionsResponseBody) SetPageIndex(v int32) *DescribeSupportedRegionsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetPageSize(v int32) *DescribeSupportedRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetRegions(v []*DescribeSupportedRegionsResponseBodyRegions) *DescribeSupportedRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetRequestId(v string) *DescribeSupportedRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetSuccess(v bool) *DescribeSupportedRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetTotalCount(v int64) *DescribeSupportedRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetTotalPage(v int32) *DescribeSupportedRegionsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSupportedRegionsResponseBodyRegions struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Extra       *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 华北2 (北京)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSupportedRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetDescription() *string {
	return s.Description
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetExtra() *string {
	return s.Extra
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetRegion() *string {
	return s.Region
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetDescription(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Description = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetExtra(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Extra = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegion(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Region = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegionName(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
