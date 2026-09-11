// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ModifyConsumerGroupPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody
	GetSuccess() *string
}

type ModifyConsumerGroupPasswordResponseBody struct {
	// The error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A06B5CFF-9576-4BC1-BE62-A3D43E1F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyConsumerGroupPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetRequestId(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
