// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLivePullToPushResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopLivePullToPushResponseBody
	GetTaskId() *string
}

type StopLivePullToPushResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// a05e6b15-15af-405b-a4a2-0152245d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The new task ID.
	//
	// example:
	//
	// fb0d4ac7-c7e3-4978-9743-0bf2f6e8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *StopLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLivePullToPushResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopLivePullToPushResponseBody) SetRequestId(v string) *StopLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLivePullToPushResponseBody) SetTaskId(v string) *StopLivePullToPushResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}
