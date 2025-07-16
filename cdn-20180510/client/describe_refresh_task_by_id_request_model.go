// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTaskByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeRefreshTaskByIdRequest
	GetTaskId() *string
}

type DescribeRefreshTaskByIdRequest struct {
	// The ID of the task that you want to query.
	//
	// You can call the [RefreshObjectCaches](https://help.aliyun.com/document_detail/91164.html) operation to query task IDs. Then, you can use the task IDs to query task status.
	//
	// You can specify up to 10 task IDs. Separate task IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTaskByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTaskByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTaskByIdRequest) SetTaskId(v string) *DescribeRefreshTaskByIdRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTaskByIdRequest) Validate() error {
	return dara.Validate(s)
}
