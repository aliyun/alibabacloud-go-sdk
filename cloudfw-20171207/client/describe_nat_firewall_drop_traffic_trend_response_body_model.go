// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallDropTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeNatFirewallDropTrafficTrendResponseBodyDataList) *DescribeNatFirewallDropTrafficTrendResponseBody
	GetDataList() []*DescribeNatFirewallDropTrafficTrendResponseBodyDataList
	SetDropSessionMax(v int64) *DescribeNatFirewallDropTrafficTrendResponseBody
	GetDropSessionMax() *int64
	SetDropSessionMaxTime(v string) *DescribeNatFirewallDropTrafficTrendResponseBody
	GetDropSessionMaxTime() *string
	SetRequestId(v string) *DescribeNatFirewallDropTrafficTrendResponseBody
	GetRequestId() *string
}

type DescribeNatFirewallDropTrafficTrendResponseBody struct {
	DataList []*DescribeNatFirewallDropTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 62436
	DropSessionMax *int64 `json:"DropSessionMax,omitempty" xml:"DropSessionMax,omitempty"`
	// example:
	//
	// 1525662720
	DropSessionMaxTime *string `json:"DropSessionMaxTime,omitempty" xml:"DropSessionMaxTime,omitempty"`
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatFirewallDropTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallDropTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) GetDataList() []*DescribeNatFirewallDropTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) GetDropSessionMax() *int64 {
	return s.DropSessionMax
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) GetDropSessionMaxTime() *string {
	return s.DropSessionMaxTime
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) SetDataList(v []*DescribeNatFirewallDropTrafficTrendResponseBodyDataList) *DescribeNatFirewallDropTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) SetDropSessionMax(v int64) *DescribeNatFirewallDropTrafficTrendResponseBody {
	s.DropSessionMax = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) SetDropSessionMaxTime(v string) *DescribeNatFirewallDropTrafficTrendResponseBody {
	s.DropSessionMaxTime = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) SetRequestId(v string) *DescribeNatFirewallDropTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBody) Validate() error {
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

type DescribeNatFirewallDropTrafficTrendResponseBodyDataList struct {
	// example:
	//
	// 10
	DropSession *int64 `json:"DropSession,omitempty" xml:"DropSession,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 153188
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeNatFirewallDropTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallDropTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) GetDropSession() *int64 {
	return s.DropSession
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) GetTotalSession() *int64 {
	return s.TotalSession
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) SetDropSession(v int64) *DescribeNatFirewallDropTrafficTrendResponseBodyDataList {
	s.DropSession = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) SetTime(v int64) *DescribeNatFirewallDropTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) SetTotalSession(v int64) *DescribeNatFirewallDropTrafficTrendResponseBodyDataList {
	s.TotalSession = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
