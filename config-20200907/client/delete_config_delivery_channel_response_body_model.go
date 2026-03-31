// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *DeleteConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *DeleteConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type DeleteConfigDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// cdc-38c32e87cadb002c****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35F1DA37-ECB5-54E9-AC22-0D9111A665AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *DeleteConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *DeleteConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *DeleteConfigDeliveryChannelResponseBody) SetRequestId(v string) *DeleteConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
