// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteResourceResponseBody
	GetTaskId() *string
}

type DeleteResourceResponseBody struct {
	// The request ID.
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

func (s DeleteResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetTaskId(v string) *DeleteResourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
