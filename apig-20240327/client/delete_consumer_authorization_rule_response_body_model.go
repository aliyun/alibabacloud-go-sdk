// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConsumerAuthorizationRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteConsumerAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
}

type DeleteConsumerAuthorizationRuleResponseBody struct {
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
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C3B9A12-3868-5EB9-8BEA-F99E03DD125C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) SetCode(v string) *DeleteConsumerAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) SetMessage(v string) *DeleteConsumerAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *DeleteConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
