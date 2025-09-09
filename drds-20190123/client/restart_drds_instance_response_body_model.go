// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDrdsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDrdsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartDrdsInstanceResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *RestartDrdsInstanceResponseBody
	GetTaskId() *int64
}

type RestartDrdsInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the PolarDB-X instance.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDrdsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDrdsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartDrdsInstanceResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RestartDrdsInstanceResponseBody) SetRequestId(v string) *RestartDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetSuccess(v bool) *RestartDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetTaskId(v int64) *RestartDrdsInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
