// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeTaskDispatchStatusRequest
	GetTaskId() *string
}

type DescribeTaskDispatchStatusRequest struct {
	// The task ID, which is the unique identifier of the log configuration modification task. Obtain this value from the TaskId response parameter of the ModifySlsDispatchConfig operation.
	//
	// example:
	//
	// 65db4ce2418b44b3be7c3xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskDispatchStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskDispatchStatusRequest) SetTaskId(v string) *DescribeTaskDispatchStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
