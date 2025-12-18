// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *CreateAggregateConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *CreateAggregateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type CreateAggregateConfigDeliveryChannelResponseBody struct {
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

func (s CreateAggregateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *CreateAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
