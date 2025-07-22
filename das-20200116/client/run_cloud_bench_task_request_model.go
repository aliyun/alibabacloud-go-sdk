// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCloudBenchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *RunCloudBenchTaskRequest
	GetTaskId() *string
}

type RunCloudBenchTaskRequest struct {
	// The stress testing task ID. You can call the [DescribeCloudBenchTasks](https://help.aliyun.com/document_detail/230670.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5cec704-0518-430f-8263-76f4dcds****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RunCloudBenchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCloudBenchTaskRequest) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCloudBenchTaskRequest) SetTaskId(v string) *RunCloudBenchTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RunCloudBenchTaskRequest) Validate() error {
	return dara.Validate(s)
}
