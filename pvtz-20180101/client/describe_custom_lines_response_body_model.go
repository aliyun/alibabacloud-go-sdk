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
	if s.CustomLines != nil {
		if err := s.CustomLines.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.CustomLine != nil {
		for _, item := range s.CustomLine {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomLinesResponseBodyCustomLinesCustomLine struct {
	CreateTime      *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                                     `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Creator         *string                                                    `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorSubType  *string                                                    `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	CreatorType     *string                                                    `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	DnsCategory     *string                                                    `json:"DnsCategory,omitempty" xml:"DnsCategory,omitempty"`
	Ipv4s           *DescribeCustomLinesResponseBodyCustomLinesCustomLineIpv4s `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Struct"`
	LineId          *string                                                    `json:"LineId,omitempty" xml:"LineId,omitempty"`
	Name            *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime      *string                                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                                     `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
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
	if s.Ipv4s != nil {
		if err := s.Ipv4s.Validate(); err != nil {
			return err
		}
	}
	return nil
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
