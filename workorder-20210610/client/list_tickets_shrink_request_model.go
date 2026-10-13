// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListTicketsShrinkRequest
	GetEndDate() *int64
	SetKeyword(v string) *ListTicketsShrinkRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListTicketsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTicketsShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *ListTicketsShrinkRequest
	GetStartDate() *int64
	SetStatusList(v []*string) *ListTicketsShrinkRequest
	GetStatusList() []*string
	SetTicketId(v string) *ListTicketsShrinkRequest
	GetTicketId() *string
	SetTicketIdListShrink(v string) *ListTicketsShrinkRequest
	GetTicketIdListShrink() *string
	SetUid(v string) *ListTicketsShrinkRequest
	GetUid() *string
}

type ListTicketsShrinkRequest struct {
	// The end time when the ticket was created. Use this parameter with StartDate to query tickets created within a specific time range.
	//
	// example:
	//
	// 1623396736000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword for a fuzzy search. The keyword is used to match the content of the Description field of the ticket.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number for a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time when the ticket was created. Use this parameter with EndDate to query tickets created within a specific time range.
	//
	// example:
	//
	// 1623396736000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// A list of ticket statuses.
	//
	// example:
	//
	// 6
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The ticket ID.
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// A list of ticket IDs.
	TicketIdListShrink *string `json:"TicketIdList,omitempty" xml:"TicketIdList,omitempty"`
	// The UID.
	//
	// example:
	//
	// 1902070573958003
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListTicketsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListTicketsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTicketsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListTicketsShrinkRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListTicketsShrinkRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketsShrinkRequest) GetTicketIdListShrink() *string {
	return s.TicketIdListShrink
}

func (s *ListTicketsShrinkRequest) GetUid() *string {
	return s.Uid
}

func (s *ListTicketsShrinkRequest) SetEndDate(v int64) *ListTicketsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetKeyword(v string) *ListTicketsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageNumber(v int32) *ListTicketsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageSize(v int32) *ListTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStartDate(v int64) *ListTicketsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStatusList(v []*string) *ListTicketsShrinkRequest {
	s.StatusList = v
	return s
}

func (s *ListTicketsShrinkRequest) SetTicketId(v string) *ListTicketsShrinkRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetTicketIdListShrink(v string) *ListTicketsShrinkRequest {
	s.TicketIdListShrink = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetUid(v string) *ListTicketsShrinkRequest {
	s.Uid = &v
	return s
}

func (s *ListTicketsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
