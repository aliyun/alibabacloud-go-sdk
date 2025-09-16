// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *DescribeSupportedZonesResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeSupportedZonesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSupportedZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSupportedZonesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeSupportedZonesResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int32) *DescribeSupportedZonesResponseBody
	GetTotalPage() *int32
	SetZoneIds(v []*string) *DescribeSupportedZonesResponseBody
	GetZoneIds() []*string
}

type DescribeSupportedZonesResponseBody struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23A9C718-DDAB-1696-B025-18FBC830F7C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	ZoneIds   []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeSupportedZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeSupportedZonesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSupportedZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportedZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSupportedZonesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSupportedZonesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSupportedZonesResponseBody) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *DescribeSupportedZonesResponseBody) SetPageIndex(v int32) *DescribeSupportedZonesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetPageSize(v int32) *DescribeSupportedZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetRequestId(v string) *DescribeSupportedZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetSuccess(v bool) *DescribeSupportedZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetTotalCount(v int64) *DescribeSupportedZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetTotalPage(v int32) *DescribeSupportedZonesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetZoneIds(v []*string) *DescribeSupportedZonesResponseBody {
	s.ZoneIds = v
	return s
}

func (s *DescribeSupportedZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
