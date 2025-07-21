// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerAuthorizationRulesResponseBody
	GetCode() *string
	SetData(v *CreateConsumerAuthorizationRulesResponseBodyData) *CreateConsumerAuthorizationRulesResponseBody
	GetData() *CreateConsumerAuthorizationRulesResponseBodyData
	SetMessage(v string) *CreateConsumerAuthorizationRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerAuthorizationRulesResponseBody
	GetRequestId() *string
}

type CreateConsumerAuthorizationRulesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response parameters.
	Data *CreateConsumerAuthorizationRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerAuthorizationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerAuthorizationRulesResponseBody) GetData() *CreateConsumerAuthorizationRulesResponseBodyData {
	return s.Data
}

func (s *CreateConsumerAuthorizationRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConsumerAuthorizationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerAuthorizationRulesResponseBody) SetCode(v string) *CreateConsumerAuthorizationRulesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponseBody) SetData(v *CreateConsumerAuthorizationRulesResponseBodyData) *CreateConsumerAuthorizationRulesResponseBody {
	s.Data = v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponseBody) SetMessage(v string) *CreateConsumerAuthorizationRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponseBody) SetRequestId(v string) *CreateConsumerAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateConsumerAuthorizationRulesResponseBodyData struct {
	// The authentication rule IDs.
	ConsumerAuthorizationRuleIds []*string `json:"consumerAuthorizationRuleIds,omitempty" xml:"consumerAuthorizationRuleIds,omitempty" type:"Repeated"`
}

func (s CreateConsumerAuthorizationRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesResponseBodyData) GetConsumerAuthorizationRuleIds() []*string {
	return s.ConsumerAuthorizationRuleIds
}

func (s *CreateConsumerAuthorizationRulesResponseBodyData) SetConsumerAuthorizationRuleIds(v []*string) *CreateConsumerAuthorizationRulesResponseBodyData {
	s.ConsumerAuthorizationRuleIds = v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
