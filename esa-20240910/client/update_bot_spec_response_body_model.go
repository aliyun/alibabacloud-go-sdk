// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBotSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateBotSpecResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *UpdateBotSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateBotSpecResponseBody
	GetRequestId() *string
}

type UpdateBotSpecResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// esa-bot-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 31223****11
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBotSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBotSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBotSpecResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBotSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateBotSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBotSpecResponseBody) SetInstanceId(v string) *UpdateBotSpecResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateBotSpecResponseBody) SetOrderId(v string) *UpdateBotSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateBotSpecResponseBody) SetRequestId(v string) *UpdateBotSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBotSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
