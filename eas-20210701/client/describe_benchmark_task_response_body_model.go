// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAgent(v int64) *DescribeBenchmarkTaskResponseBody
	GetAvailableAgent() *int64
	SetCallerUid(v string) *DescribeBenchmarkTaskResponseBody
	GetCallerUid() *string
	SetDesiredAgent(v int64) *DescribeBenchmarkTaskResponseBody
	GetDesiredAgent() *int64
	SetEndpoint(v string) *DescribeBenchmarkTaskResponseBody
	GetEndpoint() *string
	SetMessage(v string) *DescribeBenchmarkTaskResponseBody
	GetMessage() *string
	SetParentUid(v string) *DescribeBenchmarkTaskResponseBody
	GetParentUid() *string
	SetReason(v string) *DescribeBenchmarkTaskResponseBody
	GetReason() *string
	SetRequestId(v string) *DescribeBenchmarkTaskResponseBody
	GetRequestId() *string
	SetServiceName(v string) *DescribeBenchmarkTaskResponseBody
	GetServiceName() *string
	SetStatus(v string) *DescribeBenchmarkTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *DescribeBenchmarkTaskResponseBody
	GetTaskId() *string
	SetTaskName(v string) *DescribeBenchmarkTaskResponseBody
	GetTaskName() *string
	SetToken(v string) *DescribeBenchmarkTaskResponseBody
	GetToken() *string
}

type DescribeBenchmarkTaskResponseBody struct {
	// The number of instances that you can test.
	//
	// example:
	//
	// 4
	AvailableAgent *int64 `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	// The ID of the operation caller.
	//
	// example:
	//
	// 1640133467****
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The number of instances that you want to test.
	//
	// example:
	//
	// 4
	DesiredAgent *int64 `json:"DesiredAgent,omitempty" xml:"DesiredAgent,omitempty"`
	// The endpoint of the service gateway.
	//
	// example:
	//
	// 192342311234.pai-eas.cn-chengdu.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Benchmar task is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the Alibaba Cloud account that is used to call the operation.
	//
	// example:
	//
	// 1029728669****
	ParentUid *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The event or reason that causes the current state of the stress testing task.
	//
	// example:
	//
	// RUNNING
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the service that you want to test.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stress testing task.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Error
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the stress testing task.
	//
	// example:
	//
	// eas-b-gv4y86u****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-ae70
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The token for authentication when a stress testing task is created.
	//
	// example:
	//
	// 6062787a-9301****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponseBody) GetAvailableAgent() *int64 {
	return s.AvailableAgent
}

func (s *DescribeBenchmarkTaskResponseBody) GetCallerUid() *string {
	return s.CallerUid
}

func (s *DescribeBenchmarkTaskResponseBody) GetDesiredAgent() *int64 {
	return s.DesiredAgent
}

func (s *DescribeBenchmarkTaskResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBenchmarkTaskResponseBody) GetParentUid() *string {
	return s.ParentUid
}

func (s *DescribeBenchmarkTaskResponseBody) GetReason() *string {
	return s.Reason
}

func (s *DescribeBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBenchmarkTaskResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeBenchmarkTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeBenchmarkTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeBenchmarkTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeBenchmarkTaskResponseBody) GetToken() *string {
	return s.Token
}

func (s *DescribeBenchmarkTaskResponseBody) SetAvailableAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.AvailableAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetCallerUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.CallerUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetDesiredAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.DesiredAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetEndpoint(v string) *DescribeBenchmarkTaskResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetMessage(v string) *DescribeBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetParentUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.ParentUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetReason(v string) *DescribeBenchmarkTaskResponseBody {
	s.Reason = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetServiceName(v string) *DescribeBenchmarkTaskResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetStatus(v string) *DescribeBenchmarkTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskId(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskName(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetToken(v string) *DescribeBenchmarkTaskResponseBody {
	s.Token = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
