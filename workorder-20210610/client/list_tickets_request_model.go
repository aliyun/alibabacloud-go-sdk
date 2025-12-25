// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListTicketsRequest
	GetEndDate() *int64
	SetKeyword(v string) *ListTicketsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListTicketsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTicketsRequest
	GetPageSize() *int32
	SetStartDate(v int64) *ListTicketsRequest
	GetStartDate() *int64
	SetStatusList(v []*string) *ListTicketsRequest
	GetStatusList() []*string
	SetTicketId(v string) *ListTicketsRequest
	GetTicketId() *string
	SetTicketIdList(v []*string) *ListTicketsRequest
	GetTicketIdList() []*string
	SetUid(v string) *ListTicketsRequest
	GetUid() *string
}

type ListTicketsRequest struct {
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
	TicketIdList []*string `json:"TicketIdList,omitempty" xml:"TicketIdList,omitempty" type:"Repeated"`
	// UID
	//
	// example:
	//
	// 1902070573958003
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListTicketsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTicketsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListTicketsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListTicketsRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketsRequest) GetTicketIdList() []*string {
	return s.TicketIdList
}

func (s *ListTicketsRequest) GetUid() *string {
	return s.Uid
}

func (s *ListTicketsRequest) SetEndDate(v int64) *ListTicketsRequest {
	s.EndDate = &v
	return s
}

func (s *ListTicketsRequest) SetKeyword(v string) *ListTicketsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTicketsRequest) SetPageNumber(v int32) *ListTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int32) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetStartDate(v int64) *ListTicketsRequest {
	s.StartDate = &v
	return s
}

func (s *ListTicketsRequest) SetStatusList(v []*string) *ListTicketsRequest {
	s.StatusList = v
	return s
}

func (s *ListTicketsRequest) SetTicketId(v string) *ListTicketsRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketsRequest) SetTicketIdList(v []*string) *ListTicketsRequest {
	s.TicketIdList = v
	return s
}

func (s *ListTicketsRequest) SetUid(v string) *ListTicketsRequest {
	s.Uid = &v
	return s
}

func (s *ListTicketsRequest) Validate() error {
	return dara.Validate(s)
}
