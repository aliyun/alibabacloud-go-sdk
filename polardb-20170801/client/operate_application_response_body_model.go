// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateApplicationResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *OperateApplicationResponseBody
	GetTaskId() *int32
}

type OperateApplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The asynchronous task ID. For Ray restart, the parent task ID of restart_ray_cluster is returned. For DisableHeadScheduling and EnableHeadScheduling, the modify_ray_head_scheduling task ID is returned. A success response only indicates that the request is accepted. Historical applications or MCP paths may not return a task ID, and the default output is 0, which does not represent a valid task.
	//
	// example:
	//
	// 123456
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *OperateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateApplicationResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *OperateApplicationResponseBody) SetRequestId(v string) *OperateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateApplicationResponseBody) SetTaskId(v int32) *OperateApplicationResponseBody {
	s.TaskId = &v
	return s
}

func (s *OperateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
