// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitTopIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVisitTopIpResponseBody
	GetRequestId() *string
	SetTopIp(v []*DescribeVisitTopIpResponseBodyTopIp) *DescribeVisitTopIpResponseBody
	GetTopIp() []*DescribeVisitTopIpResponseBodyTopIp
}

type DescribeVisitTopIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D2B8DAE-A761-58CB-A68D-74989E4831DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 IP addresses from which requests are sent.
	TopIp []*DescribeVisitTopIpResponseBodyTopIp `json:"TopIp,omitempty" xml:"TopIp,omitempty" type:"Repeated"`
}

func (s DescribeVisitTopIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVisitTopIpResponseBody) GetTopIp() []*DescribeVisitTopIpResponseBodyTopIp {
	return s.TopIp
}

func (s *DescribeVisitTopIpResponseBody) SetRequestId(v string) *DescribeVisitTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitTopIpResponseBody) SetTopIp(v []*DescribeVisitTopIpResponseBodyTopIp) *DescribeVisitTopIpResponseBody {
	s.TopIp = v
	return s
}

func (s *DescribeVisitTopIpResponseBody) Validate() error {
	if s.TopIp != nil {
		for _, item := range s.TopIp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVisitTopIpResponseBodyTopIp struct {
	// The ordinal number of the area to which the IP address belongs.
	//
	// example:
	//
	// 310000
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The total number of requests that are sent from the IP address.
	//
	// example:
	//
	// 2622
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 1.1.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ISP.
	//
	// example:
	//
	// AAA
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeVisitTopIpResponseBodyTopIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitTopIpResponseBodyTopIp) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBodyTopIp) GetArea() *string {
	return s.Area
}

func (s *DescribeVisitTopIpResponseBodyTopIp) GetCount() *int64 {
	return s.Count
}

func (s *DescribeVisitTopIpResponseBodyTopIp) GetIp() *string {
	return s.Ip
}

func (s *DescribeVisitTopIpResponseBodyTopIp) GetIsp() *string {
	return s.Isp
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetArea(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Area = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetCount(v int64) *DescribeVisitTopIpResponseBodyTopIp {
	s.Count = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Ip = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIsp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Isp = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) Validate() error {
	return dara.Validate(s)
}
