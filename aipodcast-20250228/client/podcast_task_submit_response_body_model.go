// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskSubmitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PodcastTaskSubmitResponseBody
	GetCode() *string
	SetData(v *PodcastTaskSubmitResponseBodyData) *PodcastTaskSubmitResponseBody
	GetData() *PodcastTaskSubmitResponseBodyData
	SetHttpStatusCode(v string) *PodcastTaskSubmitResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PodcastTaskSubmitResponseBody
	GetMessage() *string
	SetRequestId(v string) *PodcastTaskSubmitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PodcastTaskSubmitResponseBody
	GetSuccess() *bool
}

type PodcastTaskSubmitResponseBody struct {
	// example:
	//
	// "success"
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *PodcastTaskSubmitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// "success"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9CE5B91A-6E6B-55FB-A1AF-037DF01C84B3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PodcastTaskSubmitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponseBody) GetCode() *string {
	return s.Code
}

func (s *PodcastTaskSubmitResponseBody) GetData() *PodcastTaskSubmitResponseBodyData {
	return s.Data
}

func (s *PodcastTaskSubmitResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PodcastTaskSubmitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PodcastTaskSubmitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PodcastTaskSubmitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PodcastTaskSubmitResponseBody) SetCode(v string) *PodcastTaskSubmitResponseBody {
	s.Code = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetData(v *PodcastTaskSubmitResponseBodyData) *PodcastTaskSubmitResponseBody {
	s.Data = v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetHttpStatusCode(v string) *PodcastTaskSubmitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetMessage(v string) *PodcastTaskSubmitResponseBody {
	s.Message = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetRequestId(v string) *PodcastTaskSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetSuccess(v bool) *PodcastTaskSubmitResponseBody {
	s.Success = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) Validate() error {
	return dara.Validate(s)
}

type PodcastTaskSubmitResponseBodyData struct {
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCEEDED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s PodcastTaskSubmitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskSubmitResponseBodyData) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PodcastTaskSubmitResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *PodcastTaskSubmitResponseBodyData) SetTaskId(v string) *PodcastTaskSubmitResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskSubmitResponseBodyData) SetTaskStatus(v string) *PodcastTaskSubmitResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *PodcastTaskSubmitResponseBodyData) Validate() error {
	return dara.Validate(s)
}
