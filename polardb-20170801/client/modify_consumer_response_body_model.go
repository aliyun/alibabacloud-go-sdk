// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ModifyConsumerResponseBody
	GetConsumerId() *string
	SetRequestId(v string) *ModifyConsumerResponseBody
	GetRequestId() *string
}

type ModifyConsumerResponseBody struct {
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ModifyConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumerResponseBody) SetConsumerId(v string) *ModifyConsumerResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *ModifyConsumerResponseBody) SetRequestId(v string) *ModifyConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
