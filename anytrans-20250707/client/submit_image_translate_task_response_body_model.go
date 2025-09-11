// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitImageTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitImageTranslateTaskResponseBodyData) *SubmitImageTranslateTaskResponseBody
	GetData() *SubmitImageTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitImageTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitImageTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImageTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitImageTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitImageTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitImageTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42FF90E5-5D40-5797-AAF6-8A4D837CCCD5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitImageTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitImageTranslateTaskResponseBody) GetData() *SubmitImageTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitImageTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitImageTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitImageTranslateTaskResponseBody) SetCode(v string) *SubmitImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetData(v *SubmitImageTranslateTaskResponseBodyData) *SubmitImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitImageTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetMessage(v string) *SubmitImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetRequestId(v string) *SubmitImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetSuccess(v bool) *SubmitImageTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitImageTranslateTaskResponseBodyData struct {
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2746f4be-cff2-465e-a2c6-12bff30ce0f9
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitImageTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitImageTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitImageTranslateTaskResponseBodyData) SetStatus(v string) *SubmitImageTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitImageTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
