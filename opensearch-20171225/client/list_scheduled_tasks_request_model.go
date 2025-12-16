// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListScheduledTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScheduledTasksRequest
	GetPageSize() *int32
	SetType(v string) *ListScheduledTasksRequest
	GetType() *string
}

type ListScheduledTasksRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scheduled task type. Valid values:
	//
	// 	- wipe: data cleaning.
	//
	// 	- fork: reindexing.
	//
	// 	- check-status: application status check.
	//
	// 	- index: reindexing.
	//
	// example:
	//
	// wipe
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScheduledTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduledTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListScheduledTasksRequest) SetPageNumber(v int32) *ListScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int32) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetType(v string) *ListScheduledTasksRequest {
	s.Type = &v
	return s
}

func (s *ListScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
