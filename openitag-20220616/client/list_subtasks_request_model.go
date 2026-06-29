// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSubtasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSubtasksRequest
	GetPageSize() *int32
}

type ListSubtasksRequest struct {
	// The page number of the subtask List. The starting value is 1, and the default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of subtasks displayed per page in a paged query. Default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubtasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubtasksRequest) GoString() string {
	return s.String()
}

func (s *ListSubtasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSubtasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubtasksRequest) SetPageNumber(v int32) *ListSubtasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubtasksRequest) SetPageSize(v int32) *ListSubtasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubtasksRequest) Validate() error {
	return dara.Validate(s)
}
