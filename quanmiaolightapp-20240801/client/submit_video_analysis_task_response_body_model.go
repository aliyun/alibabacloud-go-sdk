// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitVideoAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *SubmitVideoAnalysisTaskResponseBodyData) *SubmitVideoAnalysisTaskResponseBody
	GetData() *SubmitVideoAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitVideoAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitVideoAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitVideoAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitVideoAnalysisTaskResponseBody
	GetSuccess() *bool
}

type SubmitVideoAnalysisTaskResponseBody struct {
	// example:
	//
	// xx
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitVideoAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetData() *SubmitVideoAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetCode(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetData(v *SubmitVideoAnalysisTaskResponseBodyData) *SubmitVideoAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitVideoAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetMessage(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetRequestId(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitVideoAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVideoAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitVideoAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
