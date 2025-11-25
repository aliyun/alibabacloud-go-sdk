// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcFirewallZoneResponseBody
	GetRequestId() *string
	SetZoneList(v [][]*DescribeVpcFirewallZoneResponseBodyZoneList) *DescribeVpcFirewallZoneResponseBody
	GetZoneList() [][]*DescribeVpcFirewallZoneResponseBodyZoneList
}

type DescribeVpcFirewallZoneResponseBody struct {
	// example:
	//
	// 337A4DBA-8A01-5E9C-99CA-84293E13****
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneList  [][]*DescribeVpcFirewallZoneResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallZoneResponseBody) GetZoneList() [][]*DescribeVpcFirewallZoneResponseBodyZoneList {
	return s.ZoneList
}

func (s *DescribeVpcFirewallZoneResponseBody) SetRequestId(v string) *DescribeVpcFirewallZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallZoneResponseBody) SetZoneList(v [][]*DescribeVpcFirewallZoneResponseBodyZoneList) *DescribeVpcFirewallZoneResponseBody {
	s.ZoneList = v
	return s
}

func (s *DescribeVpcFirewallZoneResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallZoneResponseBodyZoneList struct {
	// example:
	//
	// cn-hangzhou-c
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
}

func (s DescribeVpcFirewallZoneResponseBodyZoneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallZoneResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallZoneResponseBodyZoneList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcFirewallZoneResponseBodyZoneList) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeVpcFirewallZoneResponseBodyZoneList) SetZoneId(v string) *DescribeVpcFirewallZoneResponseBodyZoneList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcFirewallZoneResponseBodyZoneList) SetLocalName(v string) *DescribeVpcFirewallZoneResponseBodyZoneList {
	s.LocalName = &v
	return s
}

func (s *DescribeVpcFirewallZoneResponseBodyZoneList) Validate() error {
	return dara.Validate(s)
}
