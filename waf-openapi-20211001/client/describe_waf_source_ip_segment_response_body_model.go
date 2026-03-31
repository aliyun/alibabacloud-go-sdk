// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafSourceIpSegmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWafSourceIpSegmentResponseBody
	GetRequestId() *string
	SetWafSourceIp(v *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) *DescribeWafSourceIpSegmentResponseBody
	GetWafSourceIp() *DescribeWafSourceIpSegmentResponseBodyWafSourceIp
}

type DescribeWafSourceIpSegmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9087ADDC-9047-4D02-82A7-33021B58083C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The back-to-origin CIDR blocks that are used by the protection cluster.
	WafSourceIp *DescribeWafSourceIpSegmentResponseBodyWafSourceIp `json:"WafSourceIp,omitempty" xml:"WafSourceIp,omitempty" type:"Struct"`
}

func (s DescribeWafSourceIpSegmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWafSourceIpSegmentResponseBody) GetWafSourceIp() *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	return s.WafSourceIp
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetRequestId(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetWafSourceIp(v *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) *DescribeWafSourceIpSegmentResponseBody {
	s.WafSourceIp = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) Validate() error {
	if s.WafSourceIp != nil {
		if err := s.WafSourceIp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWafSourceIpSegmentResponseBodyWafSourceIp struct {
	// An array of back-to-origin IPv4 CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// An array of back-to-origin IPv6 CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) GetIPv4() []*string {
	return s.IPv4
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) GetIPv6() []*string {
	return s.IPv6
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv4(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv4 = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv6(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv6 = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) Validate() error {
	return dara.Validate(s)
}
