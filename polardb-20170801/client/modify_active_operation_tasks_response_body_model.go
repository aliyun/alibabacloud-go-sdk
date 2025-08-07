// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyActiveOperationTasksResponseBody
	GetRequestId() *string
	SetTaskIds(v string) *ModifyActiveOperationTasksResponseBody
	GetTaskIds() *string
}

type ModifyActiveOperationTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 42CD2EF5-D77E-5AD4-961B-159330D98286
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task IDs.
	//
	// example:
	//
	// 11111,22222
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s ModifyActiveOperationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationTasksResponseBody) GetTaskIds() *string {
	return s.TaskIds
}

func (s *ModifyActiveOperationTasksResponseBody) SetRequestId(v string) *ModifyActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) SetTaskIds(v string) *ModifyActiveOperationTasksResponseBody {
	s.TaskIds = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
