// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *UpdateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type UpdateConfigDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// cdc-8e45ff4e06a3a8****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDERTHG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateConfigDeliveryChannelResponseBody) SetRequestId(v string) *UpdateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
