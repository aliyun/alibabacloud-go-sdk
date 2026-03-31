// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *DeleteAggregateConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *DeleteAggregateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type DeleteAggregateConfigDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// cdc-38c3013b46c9002c****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FCC2F05C-F672-5665-8102-0020DF66B9B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteAggregateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *DeleteAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *DeleteAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
