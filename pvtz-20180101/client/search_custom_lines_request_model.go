// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCustomLinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestampEnd(v int64) *SearchCustomLinesRequest
	GetCreateTimestampEnd() *int64
	SetCreateTimestampStart(v int64) *SearchCustomLinesRequest
	GetCreateTimestampStart() *int64
	SetCreator(v []*string) *SearchCustomLinesRequest
	GetCreator() []*string
	SetIpv4(v string) *SearchCustomLinesRequest
	GetIpv4() *string
	SetLang(v string) *SearchCustomLinesRequest
	GetLang() *string
	SetName(v string) *SearchCustomLinesRequest
	GetName() *string
	SetPageNumber(v int32) *SearchCustomLinesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCustomLinesRequest
	GetPageSize() *int32
	SetUpdateTimestampEnd(v int64) *SearchCustomLinesRequest
	GetUpdateTimestampEnd() *int64
	SetUpdateTimestampStart(v int64) *SearchCustomLinesRequest
	GetUpdateTimestampStart() *int64
}

type SearchCustomLinesRequest struct {
	// The end of the time range during which the custom lines are created to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672136518234
	CreateTimestampEnd *int64 `json:"CreateTimestampEnd,omitempty" xml:"CreateTimestampEnd,omitempty"`
	// The beginning of the time range during which the custom lines are created to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672136518123
	CreateTimestampStart *int64 `json:"CreateTimestampStart,omitempty" xml:"CreateTimestampStart,omitempty"`
	// The IDs of the creators for the custom lines.
	Creator []*string `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Repeated"`
	// The IPv4 address.
	//
	// example:
	//
	// 1.1.1.1
	Ipv4 *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the custom line.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The end of the time range during which the custom lines are updated to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672136518000
	UpdateTimestampEnd *int64 `json:"UpdateTimestampEnd,omitempty" xml:"UpdateTimestampEnd,omitempty"`
	// The beginning of the time range during which the custom lines are updated to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672136515000
	UpdateTimestampStart *int64 `json:"UpdateTimestampStart,omitempty" xml:"UpdateTimestampStart,omitempty"`
}

func (s SearchCustomLinesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesRequest) GetCreateTimestampEnd() *int64 {
	return s.CreateTimestampEnd
}

func (s *SearchCustomLinesRequest) GetCreateTimestampStart() *int64 {
	return s.CreateTimestampStart
}

func (s *SearchCustomLinesRequest) GetCreator() []*string {
	return s.Creator
}

func (s *SearchCustomLinesRequest) GetIpv4() *string {
	return s.Ipv4
}

func (s *SearchCustomLinesRequest) GetLang() *string {
	return s.Lang
}

func (s *SearchCustomLinesRequest) GetName() *string {
	return s.Name
}

func (s *SearchCustomLinesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCustomLinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCustomLinesRequest) GetUpdateTimestampEnd() *int64 {
	return s.UpdateTimestampEnd
}

func (s *SearchCustomLinesRequest) GetUpdateTimestampStart() *int64 {
	return s.UpdateTimestampStart
}

func (s *SearchCustomLinesRequest) SetCreateTimestampEnd(v int64) *SearchCustomLinesRequest {
	s.CreateTimestampEnd = &v
	return s
}

func (s *SearchCustomLinesRequest) SetCreateTimestampStart(v int64) *SearchCustomLinesRequest {
	s.CreateTimestampStart = &v
	return s
}

func (s *SearchCustomLinesRequest) SetCreator(v []*string) *SearchCustomLinesRequest {
	s.Creator = v
	return s
}

func (s *SearchCustomLinesRequest) SetIpv4(v string) *SearchCustomLinesRequest {
	s.Ipv4 = &v
	return s
}

func (s *SearchCustomLinesRequest) SetLang(v string) *SearchCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *SearchCustomLinesRequest) SetName(v string) *SearchCustomLinesRequest {
	s.Name = &v
	return s
}

func (s *SearchCustomLinesRequest) SetPageNumber(v int32) *SearchCustomLinesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCustomLinesRequest) SetPageSize(v int32) *SearchCustomLinesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCustomLinesRequest) SetUpdateTimestampEnd(v int64) *SearchCustomLinesRequest {
	s.UpdateTimestampEnd = &v
	return s
}

func (s *SearchCustomLinesRequest) SetUpdateTimestampStart(v int64) *SearchCustomLinesRequest {
	s.UpdateTimestampStart = &v
	return s
}

func (s *SearchCustomLinesRequest) Validate() error {
	return dara.Validate(s)
}
