// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *DeleteDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type DeleteDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
