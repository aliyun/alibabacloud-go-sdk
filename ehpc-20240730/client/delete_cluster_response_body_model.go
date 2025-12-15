// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteClusterResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *DeleteClusterResponseBody
	GetTaskId() *string
}

type DeleteClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F1AB6D8D-E185-4D94-859C-7CE7B8B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// F1AB6D8D-E185-4D94-859C-7CE7B8B7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
