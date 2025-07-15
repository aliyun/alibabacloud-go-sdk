// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartLivePullToPushResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RestartLivePullToPushResponseBody
	GetTaskId() *string
}

type RestartLivePullToPushResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// a05e6b15-15af-405b-a4a2-01522450****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The new task ID.
	//
	// example:
	//
	// fb0d4ac7-c7e3-4978-9743-0bf2f6e8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *RestartLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartLivePullToPushResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartLivePullToPushResponseBody) SetRequestId(v string) *RestartLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartLivePullToPushResponseBody) SetTaskId(v string) *RestartLivePullToPushResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}
