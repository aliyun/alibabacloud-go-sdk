// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetData() *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetSuccess() *bool
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetData() *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
