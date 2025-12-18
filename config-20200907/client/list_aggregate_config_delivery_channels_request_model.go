// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigDeliveryChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateConfigDeliveryChannelsRequest
	GetAggregatorId() *string
	SetDeliveryChannelIds(v string) *ListAggregateConfigDeliveryChannelsRequest
	GetDeliveryChannelIds() *string
}

type ListAggregateConfigDeliveryChannelsRequest struct {
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
	// The IDs of the delivery channels. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// cdc-d9106457e0d900b1****
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigDeliveryChannelsRequest) GetDeliveryChannelIds() *string {
	return s.DeliveryChannelIds
}

func (s *ListAggregateConfigDeliveryChannelsRequest) SetAggregatorId(v string) *ListAggregateConfigDeliveryChannelsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *ListAggregateConfigDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsRequest) Validate() error {
	return dara.Validate(s)
}
