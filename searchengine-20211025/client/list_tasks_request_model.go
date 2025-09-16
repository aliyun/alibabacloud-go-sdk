// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *ListTasksRequest
	GetEnd() *int64
	SetStart(v int64) *ListTasksRequest
	GetStart() *int64
}

type ListTasksRequest struct {
	// The timestamp that indicates the end of the time range to query.
	//
	// example:
	//
	// 1690423741577
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// The timestamp that indicates the beginning of the time range to query.
	//
	// example:
	//
	// 1687238865434
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetEnd() *int64 {
	return s.End
}

func (s *ListTasksRequest) GetStart() *int64 {
	return s.Start
}

func (s *ListTasksRequest) SetEnd(v int64) *ListTasksRequest {
	s.End = &v
	return s
}

func (s *ListTasksRequest) SetStart(v int64) *ListTasksRequest {
	s.Start = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
