// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateSubscriptionInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateSubscriptionInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateSubscriptionInstanceResponseBody
	GetRequestId() *string
	SetSubscriptionInstanceId(v string) *CreateSubscriptionInstanceResponseBody
	GetSubscriptionInstanceId() *string
	SetSuccess(v string) *CreateSubscriptionInstanceResponseBody
	GetSuccess() *string
}

type CreateSubscriptionInstanceResponseBody struct {
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
	// 4E9564C5-F99D-4176-A6BA-2D7F2DC8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the change tracking instance.
	//
	// example:
	//
	// dtsfen11q2g23x****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSubscriptionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateSubscriptionInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateSubscriptionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubscriptionInstanceResponseBody) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *CreateSubscriptionInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrCode(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrMessage(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetRequestId(v string) *CreateSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSubscriptionInstanceId(v string) *CreateSubscriptionInstanceResponseBody {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSuccess(v string) *CreateSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
