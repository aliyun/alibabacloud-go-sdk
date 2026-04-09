// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCustomLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLines(v *SearchCustomLinesResponseBodyCustomLines) *SearchCustomLinesResponseBody
	GetCustomLines() *SearchCustomLinesResponseBodyCustomLines
	SetPageNumber(v int32) *SearchCustomLinesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCustomLinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCustomLinesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCustomLinesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCustomLinesResponseBody
	GetTotalPages() *int32
}

type SearchCustomLinesResponseBody struct {
	CustomLines *SearchCustomLinesResponseBodyCustomLines `json:"CustomLines,omitempty" xml:"CustomLines,omitempty" type:"Struct"`
	// The page number. Default value: 1.
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
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
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

func (s SearchCustomLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesResponseBody) GetCustomLines() *SearchCustomLinesResponseBodyCustomLines {
	return s.CustomLines
}

func (s *SearchCustomLinesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCustomLinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCustomLinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCustomLinesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCustomLinesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCustomLinesResponseBody) SetCustomLines(v *SearchCustomLinesResponseBodyCustomLines) *SearchCustomLinesResponseBody {
	s.CustomLines = v
	return s
}

func (s *SearchCustomLinesResponseBody) SetPageNumber(v int32) *SearchCustomLinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCustomLinesResponseBody) SetPageSize(v int32) *SearchCustomLinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCustomLinesResponseBody) SetRequestId(v string) *SearchCustomLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCustomLinesResponseBody) SetTotalItems(v int32) *SearchCustomLinesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCustomLinesResponseBody) SetTotalPages(v int32) *SearchCustomLinesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCustomLinesResponseBody) Validate() error {
	if s.CustomLines != nil {
		if err := s.CustomLines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCustomLinesResponseBodyCustomLines struct {
	CustomLine []*SearchCustomLinesResponseBodyCustomLinesCustomLine `json:"CustomLine,omitempty" xml:"CustomLine,omitempty" type:"Repeated"`
}

func (s SearchCustomLinesResponseBodyCustomLines) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesResponseBodyCustomLines) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesResponseBodyCustomLines) GetCustomLine() []*SearchCustomLinesResponseBodyCustomLinesCustomLine {
	return s.CustomLine
}

func (s *SearchCustomLinesResponseBodyCustomLines) SetCustomLine(v []*SearchCustomLinesResponseBodyCustomLinesCustomLine) *SearchCustomLinesResponseBodyCustomLines {
	s.CustomLine = v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLines) Validate() error {
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

type SearchCustomLinesResponseBodyCustomLinesCustomLine struct {
	CreateTime      *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                                   `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Creator         *string                                                  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorSubType  *string                                                  `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	CreatorType     *string                                                  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	DnsCategory     *string                                                  `json:"DnsCategory,omitempty" xml:"DnsCategory,omitempty"`
	Ipv4s           *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Struct"`
	LineId          *string                                                  `json:"LineId,omitempty" xml:"LineId,omitempty"`
	Name            *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime      *string                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                                   `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchCustomLinesResponseBodyCustomLinesCustomLine) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesResponseBodyCustomLinesCustomLine) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetCreator() *string {
	return s.Creator
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetCreatorType() *string {
	return s.CreatorType
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetDnsCategory() *string {
	return s.DnsCategory
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetIpv4s() *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s {
	return s.Ipv4s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetLineId() *string {
	return s.LineId
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetName() *string {
	return s.Name
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetCreateTime(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreateTime = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetCreateTimestamp(v int64) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetCreator(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.Creator = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetCreatorSubType(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreatorSubType = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetCreatorType(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.CreatorType = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetDnsCategory(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.DnsCategory = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetIpv4s(v *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.Ipv4s = v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetLineId(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.LineId = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetName(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.Name = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetUpdateTime(v string) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.UpdateTime = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) SetUpdateTimestamp(v int64) *SearchCustomLinesResponseBodyCustomLinesCustomLine {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLine) Validate() error {
	if s.Ipv4s != nil {
		if err := s.Ipv4s.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s struct {
	Ipv4 []*string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty" type:"Repeated"`
}

func (s SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) GetIpv4() []*string {
	return s.Ipv4
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) SetIpv4(v []*string) *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s {
	s.Ipv4 = v
	return s
}

func (s *SearchCustomLinesResponseBodyCustomLinesCustomLineIpv4s) Validate() error {
	return dara.Validate(s)
}
