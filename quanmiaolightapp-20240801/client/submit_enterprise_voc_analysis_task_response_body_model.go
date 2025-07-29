// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseVocAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *SubmitEnterpriseVocAnalysisTaskResponseBodyData) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetData() *SubmitEnterpriseVocAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitEnterpriseVocAnalysisTaskResponseBody
	GetSuccess() *bool
}

type SubmitEnterpriseVocAnalysisTaskResponseBody struct {
	// example:
	//
	// NoPermission
	Code *string                                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitEnterpriseVocAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 403
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetData() *SubmitEnterpriseVocAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetCode(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetData(v *SubmitEnterpriseVocAnalysisTaskResponseBodyData) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetMessage(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetRequestId(v string) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitEnterpriseVocAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitEnterpriseVocAnalysisTaskResponseBodyData struct {
	// example:
	//
	// a0cc71ec-fe07-47e5-bf12-6e1c46081c98
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitEnterpriseVocAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
