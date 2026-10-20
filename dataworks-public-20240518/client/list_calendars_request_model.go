// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListCalendarsRequest
	GetName() *string
	SetOwner(v string) *ListCalendarsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListCalendarsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCalendarsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCalendarsRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListCalendarsRequest
	GetSortBy() *string
}

type ListCalendarsRequest struct {
	// The name of the scheduling calendar. Partial match is supported.
	//
	// example:
	//
	// calendar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the owner. Exact match is supported.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID. Queries the calendars applicable to this workspace, including those applicable to all workspaces. If this parameter is not specified, the calendars of the current tenant are queried.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sorting rule. The format is "Field Direction", separated by a space.
	//
	// - Field: ModifyTime, Name, or Id.
	//
	// - Direction: Asc for ascending order, or Desc for descending order. If this parameter is not specified, Asc is used.
	//
	// Default value: ModifyTime Desc.
	//
	// example:
	//
	// ModifyTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListCalendarsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarsRequest) GetName() *string {
	return s.Name
}

func (s *ListCalendarsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListCalendarsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalendarsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalendarsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCalendarsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCalendarsRequest) SetName(v string) *ListCalendarsRequest {
	s.Name = &v
	return s
}

func (s *ListCalendarsRequest) SetOwner(v string) *ListCalendarsRequest {
	s.Owner = &v
	return s
}

func (s *ListCalendarsRequest) SetPageNumber(v int32) *ListCalendarsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCalendarsRequest) SetPageSize(v int32) *ListCalendarsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCalendarsRequest) SetProjectId(v int64) *ListCalendarsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCalendarsRequest) SetSortBy(v string) *ListCalendarsRequest {
	s.SortBy = &v
	return s
}

func (s *ListCalendarsRequest) Validate() error {
	return dara.Validate(s)
}
