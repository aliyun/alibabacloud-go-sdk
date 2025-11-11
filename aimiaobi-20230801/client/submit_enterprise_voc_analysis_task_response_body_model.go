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
	// NoData
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitEnterpriseVocAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitEnterpriseVocAnalysisTaskResponseBodyData struct {
	// example:
	//
	// xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
