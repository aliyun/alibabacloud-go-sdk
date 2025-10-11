// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *CreateDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *CreateDeliveryChannelResponseBody
	GetRequestId() *string
}

type CreateDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-0bzhsqpnk***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 42A89312-0616-591E-B614-07BC87D3D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *CreateDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateDeliveryChannelResponseBody) SetRequestId(v string) *CreateDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
