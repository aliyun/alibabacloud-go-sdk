// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNodesResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *DeleteNodesResponseBody
	GetTaskId() *string
}

type DeleteNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteNodesResponseBody) SetRequestId(v string) *DeleteNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodesResponseBody) SetSuccess(v bool) *DeleteNodesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNodesResponseBody) SetTaskId(v string) *DeleteNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
