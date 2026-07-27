// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGwConsumerOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditToken(v string) *CreateGwConsumerOrderResponseBody
	GetCreditToken() *string
	SetExpireTime(v string) *CreateGwConsumerOrderResponseBody
	GetExpireTime() *string
	SetGatewayId(v string) *CreateGwConsumerOrderResponseBody
	GetGatewayId() *string
	SetOrderId(v string) *CreateGwConsumerOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateGwConsumerOrderResponseBody
	GetRequestId() *string
}

type CreateGwConsumerOrderResponseBody struct {
	// The redemption code used for subsequent activation.
	//
	// example:
	//
	// "9"
	CreditToken *string `json:"CreditToken,omitempty" xml:"CreditToken,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the AI gateway instance.
	//
	// example:
	//
	// pg-2ze24rr575j5b18cg
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The order ID returned after the order is placed.
	//
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGwConsumerOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGwConsumerOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGwConsumerOrderResponseBody) GetCreditToken() *string {
	return s.CreditToken
}

func (s *CreateGwConsumerOrderResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateGwConsumerOrderResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGwConsumerOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateGwConsumerOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGwConsumerOrderResponseBody) SetCreditToken(v string) *CreateGwConsumerOrderResponseBody {
	s.CreditToken = &v
	return s
}

func (s *CreateGwConsumerOrderResponseBody) SetExpireTime(v string) *CreateGwConsumerOrderResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateGwConsumerOrderResponseBody) SetGatewayId(v string) *CreateGwConsumerOrderResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGwConsumerOrderResponseBody) SetOrderId(v string) *CreateGwConsumerOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateGwConsumerOrderResponseBody) SetRequestId(v string) *CreateGwConsumerOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGwConsumerOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
