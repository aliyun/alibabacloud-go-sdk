// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomHotTopicBroadcastJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetCode() *string
	SetData(v *SubmitCustomHotTopicBroadcastJobResponseBodyData) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetData() *SubmitCustomHotTopicBroadcastJobResponseBodyData
	SetHttpStatusCode(v int32) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCustomHotTopicBroadcastJobResponseBody
	GetSuccess() *bool
}

type SubmitCustomHotTopicBroadcastJobResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitCustomHotTopicBroadcastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitCustomHotTopicBroadcastJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetData() *SubmitCustomHotTopicBroadcastJobResponseBodyData {
	return s.Data
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetCode(v string) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetData(v *SubmitCustomHotTopicBroadcastJobResponseBodyData) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetHttpStatusCode(v int32) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetMessage(v string) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetRequestId(v string) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) SetSuccess(v bool) *SubmitCustomHotTopicBroadcastJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitCustomHotTopicBroadcastJobResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBodyData) SetTaskId(v string) *SubmitCustomHotTopicBroadcastJobResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
