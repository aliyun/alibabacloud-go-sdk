// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySlsDispatchConfigResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifySlsDispatchConfigResponseBody
	GetTaskId() *string
}

type ModifySlsDispatchConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CD646BA-490F-5584-9272-B6FFE3BB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID. Modifying log configurations is an asynchronous task. This field indicates the unique identifier of the task. You can use this ID to query the status of the task.
	//
	// example:
	//
	// d8c995ec-40a5-4382-a3f3-57713096****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifySlsDispatchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySlsDispatchConfigResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifySlsDispatchConfigResponseBody) SetRequestId(v string) *ModifySlsDispatchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySlsDispatchConfigResponseBody) SetTaskId(v string) *ModifySlsDispatchConfigResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySlsDispatchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
