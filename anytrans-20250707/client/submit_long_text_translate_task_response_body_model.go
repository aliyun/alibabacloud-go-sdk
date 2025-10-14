// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLongTextTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitLongTextTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitLongTextTranslateTaskResponseBodyData) *SubmitLongTextTranslateTaskResponseBody
	GetData() *SubmitLongTextTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitLongTextTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitLongTextTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitLongTextTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitLongTextTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitLongTextTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitLongTextTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 400392FF-2F47-549A-A7FF-60FA4121D19E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitLongTextTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetData() *SubmitLongTextTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetCode(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetData(v *SubmitLongTextTranslateTaskResponseBodyData) *SubmitLongTextTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetMessage(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetRequestId(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetSuccess(v bool) *SubmitLongTextTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitLongTextTranslateTaskResponseBodyData struct {
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2746f4be-cff2-465e-a2c6-12bff30ce0f9
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitLongTextTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) SetStatus(v string) *SubmitLongTextTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitLongTextTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
