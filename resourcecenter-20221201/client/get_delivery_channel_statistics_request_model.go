// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *GetDeliveryChannelStatisticsRequest
	GetDeliveryChannelId() *string
}

type GetDeliveryChannelStatisticsRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetDeliveryChannelStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetDeliveryChannelStatisticsRequest) SetDeliveryChannelId(v string) *GetDeliveryChannelStatisticsRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
