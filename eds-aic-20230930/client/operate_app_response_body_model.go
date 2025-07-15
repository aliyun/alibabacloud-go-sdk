// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *OperateAppResponseBody
	GetTaskId() *string
}

type OperateAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-imr0fufqgac2z****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAppResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *OperateAppResponseBody) SetRequestId(v string) *OperateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppResponseBody) SetTaskId(v string) *OperateAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *OperateAppResponseBody) Validate() error {
	return dara.Validate(s)
}
