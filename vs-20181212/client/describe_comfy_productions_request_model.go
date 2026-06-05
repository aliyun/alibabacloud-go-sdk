// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeComfyProductionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyProductionsRequest
	GetPageSize() *int32
	SetTaskId(v string) *DescribeComfyProductionsRequest
	GetTaskId() *string
}

type DescribeComfyProductionsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6c8234f4-d1e1-4cea-b08b-7926fbdea144
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeComfyProductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyProductionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyProductionsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeComfyProductionsRequest) SetPageNumber(v int32) *DescribeComfyProductionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyProductionsRequest) SetPageSize(v int32) *DescribeComfyProductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyProductionsRequest) SetTaskId(v string) *DescribeComfyProductionsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeComfyProductionsRequest) Validate() error {
	return dara.Validate(s)
}
