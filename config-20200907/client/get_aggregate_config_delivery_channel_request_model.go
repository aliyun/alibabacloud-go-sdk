// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigDeliveryChannelRequest
	GetAggregatorId() *string
	SetDeliveryChannelId(v string) *GetAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type GetAggregateConfigDeliveryChannelRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the ID of a delivery channel, see [ListAggregateConfigDeliveryChannels](https://help.aliyun.com/document_detail/429842.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdc-d9106457e0d900b1****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *GetAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
