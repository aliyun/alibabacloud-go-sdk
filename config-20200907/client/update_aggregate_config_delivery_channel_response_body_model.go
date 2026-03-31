// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type UpdateAggregateConfigDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// cdc-8e45ff4e06a3a8****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDERTHG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
