// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePdnsUdpIpSegmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidatePdnsUdpIpSegmentResponseBody
	GetRequestId() *string
}

type ValidatePdnsUdpIpSegmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidatePdnsUdpIpSegmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidatePdnsUdpIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *ValidatePdnsUdpIpSegmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidatePdnsUdpIpSegmentResponseBody) SetRequestId(v string) *ValidatePdnsUdpIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidatePdnsUdpIpSegmentResponseBody) Validate() error {
	return dara.Validate(s)
}
