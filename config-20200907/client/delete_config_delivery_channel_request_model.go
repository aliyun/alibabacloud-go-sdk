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
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the ID of a delivery channel, see [DescribeDeliveryChannels](https://help.aliyun.com/document_detail/429841.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdc-38c32e87cadb002c****
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
