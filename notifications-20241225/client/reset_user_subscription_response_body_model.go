// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetUserSubscriptionResponseBody
	GetCode() *string
	SetMessage(v string) *ResetUserSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetUserSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetUserSubscriptionResponseBody
	GetSuccess() *bool
}

type ResetUserSubscriptionResponseBody struct {
	// The error code returned when the call fails. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned when the call fails.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73FD6AE8-898F-5D09-9763-69B8A875488A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetUserSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetUserSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetUserSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetUserSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetUserSubscriptionResponseBody) SetCode(v string) *ResetUserSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *ResetUserSubscriptionResponseBody) SetMessage(v string) *ResetUserSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *ResetUserSubscriptionResponseBody) SetRequestId(v string) *ResetUserSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserSubscriptionResponseBody) SetSuccess(v bool) *ResetUserSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ResetUserSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
