// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConsumerAuthorizationRuleResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateConsumerAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
}

type UpdateConsumerAuthorizationRuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854CD0D4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) SetCode(v string) *UpdateConsumerAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) SetMessage(v string) *UpdateConsumerAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
