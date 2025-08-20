// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateTaskResponseBodyData) *CreateTaskResponseBody
	GetData() *CreateTaskResponseBodyData
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateTaskResponseBody
	GetSuccess() *string
}

type CreateTaskResponseBody struct {
	Data *CreateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 9F1DB065-AE0D-5EE3-B1AF-48632CB0831C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetData() *CreateTaskResponseBodyData {
	return s.Data
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateTaskResponseBody) SetData(v *CreateTaskResponseBodyData) *CreateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v string) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTaskResponseBodyData struct {
	// example:
	//
	// 20240905-********-93E9-5D45-B4EF-045743A34071
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskResponseBodyData) SetTaskId(v string) *CreateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
