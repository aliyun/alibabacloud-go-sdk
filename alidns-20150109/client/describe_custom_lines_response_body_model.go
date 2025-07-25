// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLines(v []*DescribeCustomLinesResponseBodyCustomLines) *DescribeCustomLinesResponseBody
	GetCustomLines() []*DescribeCustomLinesResponseBodyCustomLines
	SetPageNumber(v int32) *DescribeCustomLinesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomLinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCustomLinesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeCustomLinesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeCustomLinesResponseBody
	GetTotalPages() *int32
}

type DescribeCustomLinesResponseBody struct {
	// The custom lines.
	CustomLines []*DescribeCustomLinesResponseBodyCustomLines `json:"CustomLines,omitempty" xml:"CustomLines,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of custom lines.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeCustomLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBody) GetCustomLines() []*DescribeCustomLinesResponseBodyCustomLines {
	return s.CustomLines
}

func (s *DescribeCustomLinesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomLinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomLinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomLinesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeCustomLinesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeCustomLinesResponseBody) SetCustomLines(v []*DescribeCustomLinesResponseBodyCustomLines) *DescribeCustomLinesResponseBody {
	s.CustomLines = v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetPageNumber(v int32) *DescribeCustomLinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetPageSize(v int32) *DescribeCustomLinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetRequestId(v string) *DescribeCustomLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetTotalItems(v int32) *DescribeCustomLinesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetTotalPages(v int32) *DescribeCustomLinesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomLinesResponseBodyCustomLines struct {
	// The code of the custom line.
	//
	// example:
	//
	// hra0yc-597
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique ID of the custom line.
	//
	// example:
	//
	// 597
	Id            *int64                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	IpSegmentList []*DescribeCustomLinesResponseBodyCustomLinesIpSegmentList `json:"IpSegmentList,omitempty" xml:"IpSegmentList,omitempty" type:"Repeated"`
	// The name of the custom line.
	//
	// example:
	//
	// hra0yd-597
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCustomLinesResponseBodyCustomLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLines) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLines) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomLinesResponseBodyCustomLines) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomLinesResponseBodyCustomLines) GetIpSegmentList() []*DescribeCustomLinesResponseBodyCustomLinesIpSegmentList {
	return s.IpSegmentList
}

func (s *DescribeCustomLinesResponseBodyCustomLines) GetName() *string {
	return s.Name
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetCode(v string) *DescribeCustomLinesResponseBodyCustomLines {
	s.Code = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetId(v int64) *DescribeCustomLinesResponseBodyCustomLines {
	s.Id = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetIpSegmentList(v []*DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) *DescribeCustomLinesResponseBodyCustomLines {
	s.IpSegmentList = v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetName(v string) *DescribeCustomLinesResponseBodyCustomLines {
	s.Name = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomLinesResponseBodyCustomLinesIpSegmentList struct {
	EndIp   *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) GetEndIp() *string {
	return s.EndIp
}

func (s *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) GetStartIp() *string {
	return s.StartIp
}

func (s *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) SetEndIp(v string) *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList {
	s.EndIp = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) SetStartIp(v string) *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList {
	s.StartIp = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesIpSegmentList) Validate() error {
	return dara.Validate(s)
}
