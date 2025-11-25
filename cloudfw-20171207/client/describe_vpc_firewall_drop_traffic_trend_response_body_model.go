// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDropTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) *DescribeVpcFirewallDropTrafficTrendResponseBody
	GetDataList() []*DescribeVpcFirewallDropTrafficTrendResponseBodyDataList
	SetDropSessionMax(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBody
	GetDropSessionMax() *int64
	SetRequestId(v string) *DescribeVpcFirewallDropTrafficTrendResponseBody
	GetRequestId() *string
}

type DescribeVpcFirewallDropTrafficTrendResponseBody struct {
	DataList []*DescribeVpcFirewallDropTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	DropSessionMax *int64 `json:"DropSessionMax,omitempty" xml:"DropSessionMax,omitempty"`
	// example:
	//
	// C87C1797-02E6-5EEB-A943-4416207D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcFirewallDropTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDropTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) GetDataList() []*DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) GetDropSessionMax() *int64 {
	return s.DropSessionMax
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) SetDataList(v []*DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) *DescribeVpcFirewallDropTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) SetDropSessionMax(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBody {
	s.DropSessionMax = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) SetRequestId(v string) *DescribeVpcFirewallDropTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDropTrafficTrendResponseBodyDataList struct {
	// example:
	//
	// 12
	AclDrop *int64 `json:"AclDrop,omitempty" xml:"AclDrop,omitempty"`
	// example:
	//
	// 2018-08-25 12:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 10
	DropSession *int64 `json:"DropSession,omitempty" xml:"DropSession,omitempty"`
	// example:
	//
	// 5
	IpsDrop *int64 `json:"IpsDrop,omitempty" xml:"IpsDrop,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 153188
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetAclDrop() *int64 {
	return s.AclDrop
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetDataTime() *string {
	return s.DataTime
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetDropSession() *int64 {
	return s.DropSession
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetIpsDrop() *int64 {
	return s.IpsDrop
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) GetTotalSession() *int64 {
	return s.TotalSession
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetAclDrop(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.AclDrop = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetDataTime(v string) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.DataTime = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetDropSession(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.DropSession = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetIpsDrop(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.IpsDrop = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetTime(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) SetTotalSession(v int64) *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList {
	s.TotalSession = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
