// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListScheduledTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListScheduledTasksRequest
	GetPageSize() *int64
	SetScheduledId(v string) *ListScheduledTasksRequest
	GetScheduledId() *string
}

type ListScheduledTasksRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on each page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the scheduled inspection configuration.
	//
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
}

func (s ListScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListScheduledTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ListScheduledTasksRequest) SetPageNumber(v int64) *ListScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int64) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetScheduledId(v string) *ListScheduledTasksRequest {
	s.ScheduledId = &v
	return s
}

func (s *ListScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
