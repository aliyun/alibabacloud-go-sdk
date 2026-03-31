// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeleteAggregateConfigDeliveryChannelRequest
	GetAggregatorId() *string
	SetDeliveryChannelId(v string) *DeleteAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type DeleteAggregateConfigDeliveryChannelRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-23c6626622af0041****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the ID of a delivery channel, see [ListAggregateConfigDeliveryChannels](https://help.aliyun.com/document_detail/429842.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdc-38c3013b46c9002c****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteAggregateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigDeliveryChannelRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregateConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *DeleteAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
