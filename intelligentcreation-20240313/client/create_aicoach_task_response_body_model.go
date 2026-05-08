// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateAICoachTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAICoachTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAICoachTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAICoachTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateAICoachTaskResponseBody
	GetTaskId() *string
}

type CreateAICoachTaskResponseBody struct {
	// example:
	//
	// Deduct.DeductTaskAlreadySuccess
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 821882330423951360
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateAICoachTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAICoachTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAICoachTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAICoachTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAICoachTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAICoachTaskResponseBody) SetErrorCode(v string) *CreateAICoachTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetErrorMessage(v string) *CreateAICoachTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetRequestId(v string) *CreateAICoachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetSuccess(v bool) *CreateAICoachTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetTaskId(v string) *CreateAICoachTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
