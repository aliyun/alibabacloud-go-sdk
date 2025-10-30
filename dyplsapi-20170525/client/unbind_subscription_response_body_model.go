// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeId(v string) *UnbindSubscriptionResponseBody
	GetChargeId() *string
	SetCode(v string) *UnbindSubscriptionResponseBody
	GetCode() *string
	SetMessage(v string) *UnbindSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindSubscriptionResponseBody
	GetRequestId() *string
}

type UnbindSubscriptionResponseBody struct {
	// A deprecated parameter.
	//
	// example:
	//
	// true
	ChargeId *string `json:"ChargeId,omitempty" xml:"ChargeId,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A9901233D36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponseBody) GetChargeId() *string {
	return s.ChargeId
}

func (s *UnbindSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSubscriptionResponseBody) SetChargeId(v string) *UnbindSubscriptionResponseBody {
	s.ChargeId = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetCode(v string) *UnbindSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetMessage(v string) *UnbindSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetRequestId(v string) *UnbindSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
