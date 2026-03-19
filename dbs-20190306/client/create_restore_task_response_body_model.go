// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateRestoreTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateRestoreTaskResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateRestoreTaskResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateRestoreTaskResponseBody
	GetRequestId() *string
	SetRestoreTaskId(v string) *CreateRestoreTaskResponseBody
	GetRestoreTaskId() *string
	SetSuccess(v bool) *CreateRestoreTaskResponseBody
	GetSuccess() *bool
}

type CreateRestoreTaskResponseBody struct {
	// error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// request ID.
	//
	// example:
	//
	// 60AF7C5D-EF4D-4D48-8FD5-C0823FDF28AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// restore job ID.
	//
	// example:
	//
	// s102h*****
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRestoreTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateRestoreTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateRestoreTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRestoreTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRestoreTaskResponseBody) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *CreateRestoreTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRestoreTaskResponseBody) SetErrCode(v string) *CreateRestoreTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetErrMessage(v string) *CreateRestoreTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetHttpStatusCode(v int32) *CreateRestoreTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetRequestId(v string) *CreateRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetRestoreTaskId(v string) *CreateRestoreTaskResponseBody {
	s.RestoreTaskId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetSuccess(v bool) *CreateRestoreTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
