// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchQueryTaskStatusResponseBody
	GetCode() *string
	SetData(v *BatchQueryTaskStatusResponseBodyData) *BatchQueryTaskStatusResponseBody
	GetData() *BatchQueryTaskStatusResponseBodyData
	SetHttpStatusCode(v int32) *BatchQueryTaskStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchQueryTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchQueryTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchQueryTaskStatusResponseBody
	GetSuccess() *bool
}

type BatchQueryTaskStatusResponseBody struct {
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchQueryTaskStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-GHIJ-KLMNOPQRST
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryTaskStatusResponseBody) GetData() *BatchQueryTaskStatusResponseBodyData {
	return s.Data
}

func (s *BatchQueryTaskStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchQueryTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchQueryTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchQueryTaskStatusResponseBody) SetCode(v string) *BatchQueryTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) SetData(v *BatchQueryTaskStatusResponseBodyData) *BatchQueryTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) SetHttpStatusCode(v int32) *BatchQueryTaskStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) SetMessage(v string) *BatchQueryTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) SetRequestId(v string) *BatchQueryTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) SetSuccess(v bool) *BatchQueryTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchQueryTaskStatusResponseBodyData struct {
	TaskStatusList []*BatchQueryTaskStatusResponseBodyDataTaskStatusList `json:"taskStatusList,omitempty" xml:"taskStatusList,omitempty" type:"Repeated"`
}

func (s BatchQueryTaskStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusResponseBodyData) GetTaskStatusList() []*BatchQueryTaskStatusResponseBodyDataTaskStatusList {
	return s.TaskStatusList
}

func (s *BatchQueryTaskStatusResponseBodyData) SetTaskStatusList(v []*BatchQueryTaskStatusResponseBodyDataTaskStatusList) *BatchQueryTaskStatusResponseBodyData {
	s.TaskStatusList = v
	return s
}

func (s *BatchQueryTaskStatusResponseBodyData) Validate() error {
	if s.TaskStatusList != nil {
		for _, item := range s.TaskStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryTaskStatusResponseBodyDataTaskStatusList struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchQueryTaskStatusResponseBodyDataTaskStatusList) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusResponseBodyDataTaskStatusList) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) GetStatus() *string {
	return s.Status
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) SetErrorMessage(v string) *BatchQueryTaskStatusResponseBodyDataTaskStatusList {
	s.ErrorMessage = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) SetStatus(v string) *BatchQueryTaskStatusResponseBodyDataTaskStatusList {
	s.Status = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) SetTaskId(v string) *BatchQueryTaskStatusResponseBodyDataTaskStatusList {
	s.TaskId = &v
	return s
}

func (s *BatchQueryTaskStatusResponseBodyDataTaskStatusList) Validate() error {
	return dara.Validate(s)
}
