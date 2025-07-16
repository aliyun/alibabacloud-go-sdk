// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTaskLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskLogsRequest
	GetPageSize() *int32
}

type ListTaskLogsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTaskLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskLogsRequest) SetPageNumber(v int32) *ListTaskLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskLogsRequest) SetPageSize(v int32) *ListTaskLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskLogsRequest) Validate() error {
	return dara.Validate(s)
}
