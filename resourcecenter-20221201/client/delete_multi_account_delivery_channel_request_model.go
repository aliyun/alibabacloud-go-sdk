// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiAccountDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *DeleteMultiAccountDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type DeleteMultiAccountDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-0bzhsqpnkxxx
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
