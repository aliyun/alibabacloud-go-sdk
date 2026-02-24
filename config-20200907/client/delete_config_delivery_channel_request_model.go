// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *DeleteConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type DeleteConfigDeliveryChannelRequest struct {
	// This parameter is required.
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
