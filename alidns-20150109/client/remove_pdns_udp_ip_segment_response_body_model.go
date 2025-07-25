// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsUdpIpSegmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePdnsUdpIpSegmentResponseBody
	GetRequestId() *string
}

type RemovePdnsUdpIpSegmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePdnsUdpIpSegmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsUdpIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePdnsUdpIpSegmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePdnsUdpIpSegmentResponseBody) SetRequestId(v string) *RemovePdnsUdpIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePdnsUdpIpSegmentResponseBody) Validate() error {
	return dara.Validate(s)
}
