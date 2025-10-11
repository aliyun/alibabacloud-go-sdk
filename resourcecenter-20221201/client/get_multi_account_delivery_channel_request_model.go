// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type GetMultiAccountDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
