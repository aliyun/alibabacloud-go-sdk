// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomSourceTopicAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitCustomSourceTopicAnalysisResponseBody
	GetCode() *string
	SetData(v *SubmitCustomSourceTopicAnalysisResponseBodyData) *SubmitCustomSourceTopicAnalysisResponseBody
	GetData() *SubmitCustomSourceTopicAnalysisResponseBodyData
	SetHttpStatusCode(v int32) *SubmitCustomSourceTopicAnalysisResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitCustomSourceTopicAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCustomSourceTopicAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCustomSourceTopicAnalysisResponseBody
	GetSuccess() *bool
}

type SubmitCustomSourceTopicAnalysisResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitCustomSourceTopicAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitCustomSourceTopicAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetData() *SubmitCustomSourceTopicAnalysisResponseBodyData {
	return s.Data
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetCode(v string) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetData(v *SubmitCustomSourceTopicAnalysisResponseBodyData) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetHttpStatusCode(v int32) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetMessage(v string) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetRequestId(v string) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) SetSuccess(v bool) *SubmitCustomSourceTopicAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitCustomSourceTopicAnalysisResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitCustomSourceTopicAnalysisResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *SubmitCustomSourceTopicAnalysisResponseBodyData) SetTaskId(v string) *SubmitCustomSourceTopicAnalysisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBodyData) SetTaskName(v string) *SubmitCustomSourceTopicAnalysisResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
