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
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
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
