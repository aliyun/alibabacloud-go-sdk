// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageToVideoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *SubmitImageToVideoTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImageToVideoTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitImageToVideoTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *SubmitImageToVideoTaskResponseBody
	GetTaskId() *string
}

type SubmitImageToVideoTaskResponseBody struct {
	// example:
	//
	// job added successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 53AED51A-74CE-57CE-B1BF-2703F314EEC8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitImageToVideoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageToVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageToVideoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageToVideoTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitImageToVideoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitImageToVideoTaskResponseBody) SetMessage(v string) *SubmitImageToVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetRequestId(v string) *SubmitImageToVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetSuccess(v bool) *SubmitImageToVideoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetTaskId(v string) *SubmitImageToVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
