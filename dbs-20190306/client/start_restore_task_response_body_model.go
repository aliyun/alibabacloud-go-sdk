// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StartRestoreTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartRestoreTaskResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StartRestoreTaskResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StartRestoreTaskResponseBody
	GetRequestId() *string
	SetRestoreTaskId(v string) *StartRestoreTaskResponseBody
	GetRestoreTaskId() *string
	SetSuccess(v bool) *StartRestoreTaskResponseBody
	GetSuccess() *bool
}

type StartRestoreTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60AF7C5D-EF4D-4D48-8FD5-C0823FDF28AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the restore task.
	//
	// example:
	//
	// s102h7rfXXXX
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartRestoreTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartRestoreTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartRestoreTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartRestoreTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRestoreTaskResponseBody) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *StartRestoreTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartRestoreTaskResponseBody) SetErrCode(v string) *StartRestoreTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetErrMessage(v string) *StartRestoreTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetHttpStatusCode(v int32) *StartRestoreTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetRequestId(v string) *StartRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetRestoreTaskId(v string) *StartRestoreTaskResponseBody {
	s.RestoreTaskId = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetSuccess(v bool) *StartRestoreTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartRestoreTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
