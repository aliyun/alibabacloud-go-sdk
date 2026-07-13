// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomLineResponseBody
	GetCode() *string
	SetDomainName(v string) *DescribeCustomLineResponseBody
	GetDomainName() *string
	SetId(v int64) *DescribeCustomLineResponseBody
	GetId() *int64
	SetIpSegmentList(v []*DescribeCustomLineResponseBodyIpSegmentList) *DescribeCustomLineResponseBody
	GetIpSegmentList() []*DescribeCustomLineResponseBodyIpSegmentList
	SetName(v string) *DescribeCustomLineResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeCustomLineResponseBody
	GetRequestId() *string
}

type DescribeCustomLineResponseBody struct {
	// The code of the custom line.
	//
	// example:
	//
	// hra0yc-*********
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the custom line.
	//
	// example:
	//
	// 5*******
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of IP address segments. Use a hyphen (-) to separate the start and end IP addresses. Each line represents one segment. You can specify from 1 to 50 segments. For a single IP address, use the format IP1-IP1. IP address segments cannot overlap.
	IpSegmentList []*DescribeCustomLineResponseBodyIpSegmentList `json:"IpSegmentList,omitempty" xml:"IpSegmentList,omitempty" type:"Repeated"`
	// The name of the custom line.
	//
	// example:
	//
	// 测试线路
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomLineResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCustomLineResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomLineResponseBody) GetIpSegmentList() []*DescribeCustomLineResponseBodyIpSegmentList {
	return s.IpSegmentList
}

func (s *DescribeCustomLineResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomLineResponseBody) SetCode(v string) *DescribeCustomLineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetDomainName(v string) *DescribeCustomLineResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetId(v int64) *DescribeCustomLineResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetIpSegmentList(v []*DescribeCustomLineResponseBodyIpSegmentList) *DescribeCustomLineResponseBody {
	s.IpSegmentList = v
	return s
}

func (s *DescribeCustomLineResponseBody) SetName(v string) *DescribeCustomLineResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetRequestId(v string) *DescribeCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLineResponseBody) Validate() error {
	if s.IpSegmentList != nil {
		for _, item := range s.IpSegmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomLineResponseBodyIpSegmentList struct {
	// The end IP address of the segment.
	//
	// example:
	//
	// 1.1.XX.XX
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The start IP address of the segment.
	//
	// example:
	//
	// 1.2.XX.XX
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s DescribeCustomLineResponseBodyIpSegmentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineResponseBodyIpSegmentList) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) GetEndIp() *string {
	return s.EndIp
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) GetStartIp() *string {
	return s.StartIp
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) SetEndIp(v string) *DescribeCustomLineResponseBodyIpSegmentList {
	s.EndIp = &v
	return s
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) SetStartIp(v string) *DescribeCustomLineResponseBodyIpSegmentList {
	s.StartIp = &v
	return s
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) Validate() error {
	return dara.Validate(s)
}
