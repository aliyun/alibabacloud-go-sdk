// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelId(v string) *CreateConfigDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetRequestId(v string) *CreateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type CreateConfigDeliveryChannelResponseBody struct {
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

func (s CreateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *CreateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateConfigDeliveryChannelResponseBody) SetRequestId(v string) *CreateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
