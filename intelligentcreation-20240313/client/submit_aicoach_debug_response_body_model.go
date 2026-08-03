// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICoachDebugResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitAICoachDebugResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SubmitAICoachDebugResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SubmitAICoachDebugResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAICoachDebugResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *SubmitAICoachDebugResponseBody
	GetTaskId() *string
}

type SubmitAICoachDebugResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitAICoachDebugResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitAICoachDebugResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitAICoachDebugResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAICoachDebugResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAICoachDebugResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitAICoachDebugResponseBody) SetErrorCode(v string) *SubmitAICoachDebugResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAICoachDebugResponseBody) SetErrorMessage(v string) *SubmitAICoachDebugResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitAICoachDebugResponseBody) SetRequestId(v string) *SubmitAICoachDebugResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAICoachDebugResponseBody) SetSuccess(v bool) *SubmitAICoachDebugResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAICoachDebugResponseBody) SetTaskId(v string) *SubmitAICoachDebugResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitAICoachDebugResponseBody) Validate() error {
	return dara.Validate(s)
}
