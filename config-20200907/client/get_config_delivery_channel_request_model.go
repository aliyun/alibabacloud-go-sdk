// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *GetConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
}

type GetConfigDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the ID of a delivery channel, see [DescribeDeliveryChannels](https://help.aliyun.com/document_detail/429841.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdc-d9106457e0d900b1****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
