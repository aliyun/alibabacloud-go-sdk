// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTaskWaitingQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyTaskWaitingQueueResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribeComfyTaskWaitingQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeComfyTaskWaitingQueueResponseBody
	GetRequestId() *string
	SetTaskWaitingQueue(v *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) *DescribeComfyTaskWaitingQueueResponseBody
	GetTaskWaitingQueue() *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue
}

type DescribeComfyTaskWaitingQueueResponseBody struct {
	// The status code. A value of 0 indicates success.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The waiting queue information.
	TaskWaitingQueue *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue `json:"TaskWaitingQueue,omitempty" xml:"TaskWaitingQueue,omitempty" type:"Struct"`
}

func (s DescribeComfyTaskWaitingQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTaskWaitingQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) GetTaskWaitingQueue() *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue {
	return s.TaskWaitingQueue
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) SetCode(v int64) *DescribeComfyTaskWaitingQueueResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) SetMessage(v string) *DescribeComfyTaskWaitingQueueResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) SetRequestId(v string) *DescribeComfyTaskWaitingQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) SetTaskWaitingQueue(v *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) *DescribeComfyTaskWaitingQueueResponseBody {
	s.TaskWaitingQueue = v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponseBody) Validate() error {
	if s.TaskWaitingQueue != nil {
		if err := s.TaskWaitingQueue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue struct {
	// The number of waiting tasks.
	//
	// example:
	//
	// 20
	WaitingCount *int64 `json:"WaitingCount,omitempty" xml:"WaitingCount,omitempty"`
}

func (s DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) GoString() string {
	return s.String()
}

func (s *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) GetWaitingCount() *int64 {
	return s.WaitingCount
}

func (s *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) SetWaitingCount(v int64) *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue {
	s.WaitingCount = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponseBodyTaskWaitingQueue) Validate() error {
	return dara.Validate(s)
}
