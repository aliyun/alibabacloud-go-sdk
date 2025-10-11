// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *GetDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type GetDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
