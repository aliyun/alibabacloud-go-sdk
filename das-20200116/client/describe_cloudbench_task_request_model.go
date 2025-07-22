// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeCloudbenchTaskRequest
	GetTaskId() *string
}

type DescribeCloudbenchTaskRequest struct {
	// The ID of the stress testing task. You can call the [DescribeCloudBenchTasks](https://help.aliyun.com/document_detail/230670.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5cec704-0518-430f-8263-76f4dcds****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCloudbenchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudbenchTaskRequest) SetTaskId(v string) *DescribeCloudbenchTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskRequest) Validate() error {
	return dara.Validate(s)
}
