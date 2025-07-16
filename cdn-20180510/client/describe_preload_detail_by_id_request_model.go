// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribePreloadDetailByIdRequest
	GetTaskId() *string
}

type DescribePreloadDetailByIdRequest struct {
	// Queries the details of a preload task by task ID. You can query one task ID at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15423123921
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreloadDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribePreloadDetailByIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreloadDetailByIdRequest) SetTaskId(v string) *DescribePreloadDetailByIdRequest {
	s.TaskId = &v
	return s
}

func (s *DescribePreloadDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
