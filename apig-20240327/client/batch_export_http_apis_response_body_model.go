// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportHttpApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchExportHttpApisResponseBody
	GetCode() *string
	SetData(v *BatchExportHttpApisResponseBodyData) *BatchExportHttpApisResponseBody
	GetData() *BatchExportHttpApisResponseBodyData
	SetMessage(v string) *BatchExportHttpApisResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchExportHttpApisResponseBody
	GetRequestId() *string
}

type BatchExportHttpApisResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchExportHttpApisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchExportHttpApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchExportHttpApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExportHttpApisResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchExportHttpApisResponseBody) GetData() *BatchExportHttpApisResponseBodyData {
	return s.Data
}

func (s *BatchExportHttpApisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchExportHttpApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchExportHttpApisResponseBody) SetCode(v string) *BatchExportHttpApisResponseBody {
	s.Code = &v
	return s
}

func (s *BatchExportHttpApisResponseBody) SetData(v *BatchExportHttpApisResponseBodyData) *BatchExportHttpApisResponseBody {
	s.Data = v
	return s
}

func (s *BatchExportHttpApisResponseBody) SetMessage(v string) *BatchExportHttpApisResponseBody {
	s.Message = &v
	return s
}

func (s *BatchExportHttpApisResponseBody) SetRequestId(v string) *BatchExportHttpApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchExportHttpApisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchExportHttpApisResponseBodyData struct {
	// example:
	//
	// async-task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchExportHttpApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchExportHttpApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchExportHttpApisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchExportHttpApisResponseBodyData) SetTaskId(v string) *BatchExportHttpApisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BatchExportHttpApisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
