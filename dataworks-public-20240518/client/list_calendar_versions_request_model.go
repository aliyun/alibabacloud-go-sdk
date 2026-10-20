// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListCalendarVersionsRequest
	GetId() *int64
	SetPageNumber(v int32) *ListCalendarVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCalendarVersionsRequest
	GetPageSize() *int32
}

type ListCalendarVersionsRequest struct {
	// The ID of the scheduling calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCalendarVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListCalendarVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalendarVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalendarVersionsRequest) SetId(v int64) *ListCalendarVersionsRequest {
	s.Id = &v
	return s
}

func (s *ListCalendarVersionsRequest) SetPageNumber(v int32) *ListCalendarVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCalendarVersionsRequest) SetPageSize(v int32) *ListCalendarVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCalendarVersionsRequest) Validate() error {
	return dara.Validate(s)
}
