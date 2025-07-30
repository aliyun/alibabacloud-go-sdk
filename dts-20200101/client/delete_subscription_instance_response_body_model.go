// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteSubscriptionInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteSubscriptionInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteSubscriptionInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSubscriptionInstanceResponseBody
	GetSuccess() *string
}

type DeleteSubscriptionInstanceResponseBody struct {
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
	// C306C198-7807-409D-930A-D6CE6C32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSubscriptionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteSubscriptionInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteSubscriptionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubscriptionInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrCode(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrMessage(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetRequestId(v string) *DeleteSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetSuccess(v string) *DeleteSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
