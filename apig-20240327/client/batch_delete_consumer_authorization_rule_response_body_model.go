// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody
	GetCode() *string
	SetMessage(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
}

type BatchDeleteConsumerAuthorizationRuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 464F9EA0-1052-51BD-8187-D292AA2D8D24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchDeleteConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) SetCode(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) SetMessage(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *BatchDeleteConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
