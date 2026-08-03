// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptGenerateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAICoachScriptGenerateTaskResponseBodyData) *GetAICoachScriptGenerateTaskResponseBody
	GetData() *GetAICoachScriptGenerateTaskResponseBodyData
	SetErrorCode(v string) *GetAICoachScriptGenerateTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAICoachScriptGenerateTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAICoachScriptGenerateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAICoachScriptGenerateTaskResponseBody
	GetSuccess() *bool
}

type GetAICoachScriptGenerateTaskResponseBody struct {
	Data         *GetAICoachScriptGenerateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode    *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAICoachScriptGenerateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptGenerateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptGenerateTaskResponseBody) GetData() *GetAICoachScriptGenerateTaskResponseBodyData {
	return s.Data
}

func (s *GetAICoachScriptGenerateTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAICoachScriptGenerateTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAICoachScriptGenerateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachScriptGenerateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAICoachScriptGenerateTaskResponseBody) SetData(v *GetAICoachScriptGenerateTaskResponseBodyData) *GetAICoachScriptGenerateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBody) SetErrorCode(v string) *GetAICoachScriptGenerateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBody) SetErrorMessage(v string) *GetAICoachScriptGenerateTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBody) SetRequestId(v string) *GetAICoachScriptGenerateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBody) SetSuccess(v bool) *GetAICoachScriptGenerateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachScriptGenerateTaskResponseBodyData struct {
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TotalTokens    *int64  `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetAICoachScriptGenerateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptGenerateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) SetScriptRecordId(v string) *GetAICoachScriptGenerateTaskResponseBodyData {
	s.ScriptRecordId = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) SetStatus(v string) *GetAICoachScriptGenerateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) SetTaskId(v string) *GetAICoachScriptGenerateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) SetTotalTokens(v int64) *GetAICoachScriptGenerateTaskResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
