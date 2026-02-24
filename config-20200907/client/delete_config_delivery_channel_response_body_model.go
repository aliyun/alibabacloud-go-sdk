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
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
