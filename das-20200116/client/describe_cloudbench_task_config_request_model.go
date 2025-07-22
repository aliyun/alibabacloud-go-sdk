// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeCloudbenchTaskConfigRequest
	GetTaskId() *string
}

type DescribeCloudbenchTaskConfigRequest struct {
	// The task ID. You can call the [DescribeCloudBenchTasks](https://help.aliyun.com/document_detail/230670.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5cec704-0518-430f-8263-76f4dcds****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCloudbenchTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudbenchTaskConfigRequest) SetTaskId(v string) *DescribeCloudbenchTaskConfigRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
