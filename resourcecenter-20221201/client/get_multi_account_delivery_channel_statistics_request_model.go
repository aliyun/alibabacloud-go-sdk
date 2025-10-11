// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelStatisticsRequest
	GetDeliveryChannelId() *string
}

type GetMultiAccountDeliveryChannelStatisticsRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetMultiAccountDeliveryChannelStatisticsRequest) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelStatisticsRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
