// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGenerateAICoachScriptTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateGenerateAICoachScriptTaskResponseBodyData) *CreateGenerateAICoachScriptTaskResponseBody
	GetData() *CreateGenerateAICoachScriptTaskResponseBodyData
	SetErrorCode(v string) *CreateGenerateAICoachScriptTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateGenerateAICoachScriptTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateGenerateAICoachScriptTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGenerateAICoachScriptTaskResponseBody
	GetSuccess() *bool
}

type CreateGenerateAICoachScriptTaskResponseBody struct {
	Data         *CreateGenerateAICoachScriptTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode    *string                                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateGenerateAICoachScriptTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGenerateAICoachScriptTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) GetData() *CreateGenerateAICoachScriptTaskResponseBodyData {
	return s.Data
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) SetData(v *CreateGenerateAICoachScriptTaskResponseBodyData) *CreateGenerateAICoachScriptTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) SetErrorCode(v string) *CreateGenerateAICoachScriptTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) SetErrorMessage(v string) *CreateGenerateAICoachScriptTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) SetRequestId(v string) *CreateGenerateAICoachScriptTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) SetSuccess(v bool) *CreateGenerateAICoachScriptTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGenerateAICoachScriptTaskResponseBodyData struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateGenerateAICoachScriptTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGenerateAICoachScriptTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGenerateAICoachScriptTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateGenerateAICoachScriptTaskResponseBodyData) SetTaskId(v string) *CreateGenerateAICoachScriptTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
