// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdpReflectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUdpReflectResponseBody
	GetRequestId() *string
	SetUdpSports(v []*string) *DescribeUdpReflectResponseBody
	GetUdpSports() []*string
}

type DescribeUdpReflectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F97A8766-FB4D-411A-9CD5-2CFF701B592F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the source ports of the UDP traffic that are filtered out by the filtering policies for UDP reflection attacks.
	UdpSports []*string `json:"UdpSports,omitempty" xml:"UdpSports,omitempty" type:"Repeated"`
}

func (s DescribeUdpReflectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdpReflectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUdpReflectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUdpReflectResponseBody) GetUdpSports() []*string {
	return s.UdpSports
}

func (s *DescribeUdpReflectResponseBody) SetRequestId(v string) *DescribeUdpReflectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUdpReflectResponseBody) SetUdpSports(v []*string) *DescribeUdpReflectResponseBody {
	s.UdpSports = v
	return s
}

func (s *DescribeUdpReflectResponseBody) Validate() error {
	return dara.Validate(s)
}
