// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiAccountDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *CreateMultiAccountDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *CreateMultiAccountDeliveryChannelResponseBody
	GetRequestId() *string
}

type CreateMultiAccountDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 994BFEFE-4BB5-5A27-8917-4583DEEF2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *CreateMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
