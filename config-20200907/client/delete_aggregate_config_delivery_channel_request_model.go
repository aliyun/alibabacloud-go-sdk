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
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
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
