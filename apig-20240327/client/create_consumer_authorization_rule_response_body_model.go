// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerAuthorizationRuleResponseBody
	GetCode() *string
	SetData(v *CreateConsumerAuthorizationRuleResponseBodyData) *CreateConsumerAuthorizationRuleResponseBody
	GetData() *CreateConsumerAuthorizationRuleResponseBodyData
	SetMessage(v string) *CreateConsumerAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
}

type CreateConsumerAuthorizationRuleResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateConsumerAuthorizationRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 6CC83C32-3B5A-57EE-9AFE-D0D51822C7BA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerAuthorizationRuleResponseBody) GetData() *CreateConsumerAuthorizationRuleResponseBodyData {
	return s.Data
}

func (s *CreateConsumerAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerAuthorizationRuleResponseBody) SetCode(v string) *CreateConsumerAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponseBody) SetData(v *CreateConsumerAuthorizationRuleResponseBodyData) *CreateConsumerAuthorizationRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponseBody) SetMessage(v string) *CreateConsumerAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *CreateConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateConsumerAuthorizationRuleResponseBodyData struct {
	// example:
	//
	// car-d06p196m1hkg9ukum5pg
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
}

func (s CreateConsumerAuthorizationRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRuleResponseBodyData) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *CreateConsumerAuthorizationRuleResponseBodyData) SetConsumerAuthorizationRuleId(v string) *CreateConsumerAuthorizationRuleResponseBodyData {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
