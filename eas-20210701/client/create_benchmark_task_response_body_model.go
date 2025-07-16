// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateBenchmarkTaskResponseBody
	GetMessage() *string
	SetRegion(v string) *CreateBenchmarkTaskResponseBody
	GetRegion() *string
	SetRequestId(v string) *CreateBenchmarkTaskResponseBody
	GetRequestId() *string
	SetTaskName(v string) *CreateBenchmarkTaskResponseBody
	GetTaskName() *string
}

type CreateBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark  task [foo] is Creating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region where the stress testing task is performed.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-1076
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBenchmarkTaskResponseBody) GetRegion() *string {
	return s.Region
}

func (s *CreateBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBenchmarkTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateBenchmarkTaskResponseBody) SetMessage(v string) *CreateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRegion(v string) *CreateBenchmarkTaskResponseBody {
	s.Region = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRequestId(v string) *CreateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetTaskName(v string) *CreateBenchmarkTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
