// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateResourceResponseBody
	GetTaskId() *string
}

type UpdateResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the asynchronous task. If the operation is asynchronous, this field is returned. In this case, the HTTP status code 202 is returned.
	//
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetTaskId(v string) *UpdateResourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
