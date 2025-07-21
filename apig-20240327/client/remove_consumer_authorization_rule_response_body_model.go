// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveConsumerAuthorizationRuleResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveConsumerAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
}

type RemoveConsumerAuthorizationRuleResponseBody struct {
	// The status code returned.
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
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) SetCode(v string) *RemoveConsumerAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) SetMessage(v string) *RemoveConsumerAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *RemoveConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
