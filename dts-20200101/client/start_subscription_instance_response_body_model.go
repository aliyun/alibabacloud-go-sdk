// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSubscriptionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StartSubscriptionInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartSubscriptionInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *StartSubscriptionInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartSubscriptionInstanceResponseBody
	GetSuccess() *string
	SetTaskId(v string) *StartSubscriptionInstanceResponseBody
	GetTaskId() *string
}

type StartSubscriptionInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B083F9AB-BE9B-4716-8AD3-CFA04391****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID. This parameter will be removed in the future.
	//
	// example:
	//
	// 11****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartSubscriptionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartSubscriptionInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartSubscriptionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSubscriptionInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartSubscriptionInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartSubscriptionInstanceResponseBody) SetErrCode(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrMessage(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetRequestId(v string) *StartSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetSuccess(v string) *StartSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetTaskId(v string) *StartSubscriptionInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
