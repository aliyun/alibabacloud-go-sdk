// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListRoutineRoutesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRoutineRoutesRequest
	GetPageSize() *int32
	SetRoutineName(v string) *ListRoutineRoutesRequest
	GetRoutineName() *string
}

type ListRoutineRoutesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s ListRoutineRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineRoutesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRoutineRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRoutineRoutesRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineRoutesRequest) SetPageNumber(v int32) *ListRoutineRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineRoutesRequest) SetPageSize(v int32) *ListRoutineRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutineRoutesRequest) SetRoutineName(v string) *ListRoutineRoutesRequest {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineRoutesRequest) Validate() error {
	return dara.Validate(s)
}
