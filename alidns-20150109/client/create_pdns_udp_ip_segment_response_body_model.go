// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsUdpIpSegmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePdnsUdpIpSegmentResponseBody
	GetRequestId() *string
	SetValidMessage(v string) *CreatePdnsUdpIpSegmentResponseBody
	GetValidMessage() *string
}

type CreatePdnsUdpIpSegmentResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ValidMessage *string `json:"ValidMessage,omitempty" xml:"ValidMessage,omitempty"`
}

func (s CreatePdnsUdpIpSegmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsUdpIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePdnsUdpIpSegmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePdnsUdpIpSegmentResponseBody) GetValidMessage() *string {
	return s.ValidMessage
}

func (s *CreatePdnsUdpIpSegmentResponseBody) SetRequestId(v string) *CreatePdnsUdpIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentResponseBody) SetValidMessage(v string) *CreatePdnsUdpIpSegmentResponseBody {
	s.ValidMessage = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentResponseBody) Validate() error {
	return dara.Validate(s)
}
