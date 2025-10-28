// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeOrderTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRetryStatus(v bool) *RetryChangeOrderTaskRequest
	GetRetryStatus() *bool
	SetTaskId(v string) *RetryChangeOrderTaskRequest
	GetTaskId() *string
}

type RetryChangeOrderTaskRequest struct {
	// The retry status.
	//
	// example:
	//
	// true
	RetryStatus *bool `json:"RetryStatus,omitempty" xml:"RetryStatus,omitempty"`
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 823-bhjf-23u4-eiuf*
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RetryChangeOrderTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeOrderTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryChangeOrderTaskRequest) GetRetryStatus() *bool {
	return s.RetryStatus
}

func (s *RetryChangeOrderTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryChangeOrderTaskRequest) SetRetryStatus(v bool) *RetryChangeOrderTaskRequest {
	s.RetryStatus = &v
	return s
}

func (s *RetryChangeOrderTaskRequest) SetTaskId(v string) *RetryChangeOrderTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RetryChangeOrderTaskRequest) Validate() error {
	return dara.Validate(s)
}
