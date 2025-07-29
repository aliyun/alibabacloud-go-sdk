// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTagMiningAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitTagMiningAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *SubmitTagMiningAnalysisTaskResponseBodyData) *SubmitTagMiningAnalysisTaskResponseBody
	GetData() *SubmitTagMiningAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitTagMiningAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTagMiningAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitTagMiningAnalysisTaskResponseBody
	GetSuccess() *bool
}

type SubmitTagMiningAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitTagMiningAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetData() *SubmitTagMiningAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetCode(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetData(v *SubmitTagMiningAnalysisTaskResponseBodyData) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetMessage(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetRequestId(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitTagMiningAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTagMiningAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitTagMiningAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
