// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifySubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifySubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ModifySubscriptionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ModifySubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySubscriptionResponseBody
	GetSuccess() *string
}

type ModifySubscriptionResponseBody struct {
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 068FA72F-4800-4A54-90BB-94806068****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifySubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifySubscriptionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifySubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySubscriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySubscriptionResponseBody) SetErrCode(v string) *ModifySubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetErrMessage(v string) *ModifySubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetHttpStatusCode(v string) *ModifySubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetRequestId(v string) *ModifySubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetSuccess(v string) *ModifySubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
