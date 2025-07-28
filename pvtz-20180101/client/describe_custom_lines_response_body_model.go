// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLines(v *DescribeCustomLinesResponseBodyCustomLines) *DescribeCustomLinesResponseBody
	GetCustomLines() *DescribeCustomLinesResponseBodyCustomLines
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
	CustomLines *DescribeCustomLinesResponseBodyCustomLines `json:"CustomLines,omitempty" xml:"CustomLines,omitempty" type:"Struct"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 5
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeCustomLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBody) GetCustomLines() *DescribeCustomLinesResponseBodyCustomLines {
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

func (s *DescribeCustomLinesResponseBody) SetCustomLines(v *DescribeCustomLinesResponseBodyCustomLines) *DescribeCustomLinesResponseBody {
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
	CustomLine []*DescribeCustomLinesResponseBodyCustomLinesCustomLine `json:"CustomLine,omitempty" xml:"CustomLine,omitempty" type:"Repeated"`
}

func (s DescribeCustomLinesResponseBodyCustomLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLines) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLines) GetCustomLine() []*DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	return s.CustomLine
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetCustomLine(v []*DescribeCustomLinesResponseBodyCustomLinesCustomLine) *DescribeCustomLinesResponseBodyCustomLines {
	s.CustomLine = v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomLinesResponseBodyCustomLinesCustomLine struct {
	// The time when the custom line was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-25T08:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the custom line was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1671174074000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the custom line.
	//
	// example:
	//
	// 21312421
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The type of the creator for the custom line. Valid values:
	//
	// 	- CUSTOM: Alibaba Cloud account
	//
	// 	- SUB: RAM user
	//
	// 	- STS: assumed role that obtains the Security Token Service (STS) token of a RAM role
	//
	// 	- OTHER: other roles
	//
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// The role of the creator for the custom line. Valid values:
	//
	// 	- USER: user
	//
	// 	- SYSTEM: system
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	DnsCategory *string `json:"DnsCategory,omitempty" xml:"DnsCategory,omitempty"`
	// The IPv4 CIDR blocks.
	Ipv4s *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Struct"`
	// The unique ID of the custom line.
	//
	// example:
	//
	// 160002
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The name of the custom line.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the custom line was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-08-24T16:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the custom line was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1681879029000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeCustomLinesResponseBodyCustomLinesCustomLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLinesCustomLine) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetCreator() *string {
	return s.Creator
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetDnsCategory() *string {
	return s.DnsCategory
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetIpv4s() *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s {
	return s.Ipv4s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetLineId() *string {
	return s.LineId
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetName() *string {
	return s.Name
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetCreateTime(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreateTime = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetCreateTimestamp(v int64) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetCreator(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.Creator = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetCreatorSubType(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetCreatorType(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreatorType = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetDnsCategory(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.DnsCategory = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetIpv4s(v *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.Ipv4s = v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetLineId(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.LineId = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetName(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.Name = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetUpdateTime(v string) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) SetUpdateTimestamp(v int64) *DescribeCustomLinesResponseBodyCustomLinesCustomLine {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLine) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s struct {
	Ipv4 []*string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty" type:"Repeated"`
}

func (s DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) GetIpv4() []*string {
	return s.Ipv4
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) SetIpv4(v []*string) *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s {
	s.Ipv4 = v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s) Validate() error {
	return dara.Validate(s)
}
