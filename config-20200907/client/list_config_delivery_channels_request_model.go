// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigDeliveryChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelIds(v string) *ListConfigDeliveryChannelsRequest
	GetDeliveryChannelIds() *string
}

type ListConfigDeliveryChannelsRequest struct {
	// The ID of the delivery channel. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// cdc-d9106457e0d900b1****
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s ListConfigDeliveryChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsRequest) GetDeliveryChannelIds() *string {
	return s.DeliveryChannelIds
}

func (s *ListConfigDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *ListConfigDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

func (s *ListConfigDeliveryChannelsRequest) Validate() error {
	return dara.Validate(s)
}
