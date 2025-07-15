// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTopicSelectionPerspectiveAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetData() *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetSuccess() *bool
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetData() *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskName(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
