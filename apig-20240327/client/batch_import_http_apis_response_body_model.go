// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchImportHttpApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchImportHttpApisResponseBody
	GetCode() *string
	SetData(v *BatchImportHttpApisResponseBodyData) *BatchImportHttpApisResponseBody
	GetData() *BatchImportHttpApisResponseBodyData
	SetMessage(v string) *BatchImportHttpApisResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchImportHttpApisResponseBody
	GetRequestId() *string
}

type BatchImportHttpApisResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchImportHttpApisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchImportHttpApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchImportHttpApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchImportHttpApisResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchImportHttpApisResponseBody) GetData() *BatchImportHttpApisResponseBodyData {
	return s.Data
}

func (s *BatchImportHttpApisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchImportHttpApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchImportHttpApisResponseBody) SetCode(v string) *BatchImportHttpApisResponseBody {
	s.Code = &v
	return s
}

func (s *BatchImportHttpApisResponseBody) SetData(v *BatchImportHttpApisResponseBodyData) *BatchImportHttpApisResponseBody {
	s.Data = v
	return s
}

func (s *BatchImportHttpApisResponseBody) SetMessage(v string) *BatchImportHttpApisResponseBody {
	s.Message = &v
	return s
}

func (s *BatchImportHttpApisResponseBody) SetRequestId(v string) *BatchImportHttpApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchImportHttpApisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchImportHttpApisResponseBodyData struct {
	// example:
	//
	// async-task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchImportHttpApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchImportHttpApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchImportHttpApisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchImportHttpApisResponseBodyData) SetTaskId(v string) *BatchImportHttpApisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BatchImportHttpApisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
