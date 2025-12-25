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
	// The deadline for ticket creation. This parameter is used in conjunction with StartDate to query tickets submitted within the specified start and end time ranges.
	//
	// example:
	//
	// 1623396736000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ticket keyword, which is used for fuzzy search to match the content of the Description field when a ticket is created.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Paging query page number parameters
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries queried by page parameter
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time when the ticket was created. This parameter is used with EndDate to query tickets that are created within the specified start and end time ranges.
	//
	// example:
	//
	// 1623396736000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Multiple ticket statuses
	//
	// example:
	//
	// 6
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// Work Order Number
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// Multiple job numbers
	TicketIdListShrink *string `json:"TicketIdList,omitempty" xml:"TicketIdList,omitempty"`
	// UID
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
