// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGetDBListFromAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateGetDBListFromAgentTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateGetDBListFromAgentTaskResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateGetDBListFromAgentTaskResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateGetDBListFromAgentTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGetDBListFromAgentTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *CreateGetDBListFromAgentTaskResponseBody
	GetTaskId() *int64
}

type CreateGetDBListFromAgentTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This backupPlan can\\"t support this action
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB4DFD5E-3618-498D-BE35-4DBEA0072122
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGetDBListFromAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGetDBListFromAgentTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetErrCode(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetErrMessage(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetHttpStatusCode(v int32) *CreateGetDBListFromAgentTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetRequestId(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetSuccess(v bool) *CreateGetDBListFromAgentTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetTaskId(v int64) *CreateGetDBListFromAgentTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
