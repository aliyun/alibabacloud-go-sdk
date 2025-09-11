// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHtmlTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitHtmlTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitHtmlTranslateTaskResponseBodyData) *SubmitHtmlTranslateTaskResponseBody
	GetData() *SubmitHtmlTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitHtmlTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitHtmlTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitHtmlTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitHtmlTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitHtmlTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitHtmlTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 72E4FDA1-5474-5DC1-8DFF-968BEEA65C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitHtmlTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetData() *SubmitHtmlTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHtmlTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetCode(v string) *SubmitHtmlTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetData(v *SubmitHtmlTranslateTaskResponseBodyData) *SubmitHtmlTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitHtmlTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetMessage(v string) *SubmitHtmlTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetRequestId(v string) *SubmitHtmlTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) SetSuccess(v bool) *SubmitHtmlTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitHtmlTranslateTaskResponseBodyData struct {
	// example:
	//
	// in_process
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2746f4be-cff2-465e-a2c6-12bff30ce0f9
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitHtmlTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitHtmlTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitHtmlTranslateTaskResponseBodyData) SetStatus(v string) *SubmitHtmlTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitHtmlTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
