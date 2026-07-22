// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallDropTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeFirewallDropTrendResponseBodyDataList) *DescribeFirewallDropTrendResponseBody
	GetDataList() []*DescribeFirewallDropTrendResponseBodyDataList
	SetMaxDropSession(v int64) *DescribeFirewallDropTrendResponseBody
	GetMaxDropSession() *int64
	SetMaxDropTime(v int64) *DescribeFirewallDropTrendResponseBody
	GetMaxDropTime() *int64
	SetRequestId(v string) *DescribeFirewallDropTrendResponseBody
	GetRequestId() *string
}

type DescribeFirewallDropTrendResponseBody struct {
	// The returned data list.
	DataList []*DescribeFirewallDropTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The maximum number of total blocked sessions.
	//
	// example:
	//
	// 300
	MaxDropSession *int64 `json:"MaxDropSession,omitempty" xml:"MaxDropSession,omitempty"`
	// The time when the maximum number of total blocked sessions occurred. The value is a UNIX timestamp in seconds, such as 1672502400.
	//
	// example:
	//
	// 1656837360
	MaxDropTime *int64 `json:"MaxDropTime,omitempty" xml:"MaxDropTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75E60025-43C5-5635-B7B7-272C5246****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFirewallDropTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropTrendResponseBody) GetDataList() []*DescribeFirewallDropTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeFirewallDropTrendResponseBody) GetMaxDropSession() *int64 {
	return s.MaxDropSession
}

func (s *DescribeFirewallDropTrendResponseBody) GetMaxDropTime() *int64 {
	return s.MaxDropTime
}

func (s *DescribeFirewallDropTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallDropTrendResponseBody) SetDataList(v []*DescribeFirewallDropTrendResponseBodyDataList) *DescribeFirewallDropTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeFirewallDropTrendResponseBody) SetMaxDropSession(v int64) *DescribeFirewallDropTrendResponseBody {
	s.MaxDropSession = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBody) SetMaxDropTime(v int64) *DescribeFirewallDropTrendResponseBody {
	s.MaxDropTime = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBody) SetRequestId(v string) *DescribeFirewallDropTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBody) Validate() error {
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

type DescribeFirewallDropTrendResponseBodyDataList struct {
	// The number of sessions blocked by the Internet firewall.
	//
	// example:
	//
	// 100
	InternetDropSession *int64 `json:"InternetDropSession,omitempty" xml:"InternetDropSession,omitempty"`
	// The number of sessions blocked by the NAT firewall.
	//
	// example:
	//
	// 100
	NatDropSession *int64 `json:"NatDropSession,omitempty" xml:"NatDropSession,omitempty"`
	// The time when the traffic occurred. The value is a UNIX timestamp in seconds.
	//
	// If the data at this point in time has not been processed, the values of other fields are -1.
	//
	// example:
	//
	// 1758474000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total number of sessions blocked by the firewall.
	//
	// example:
	//
	// 300
	TotalDropSession *int64 `json:"TotalDropSession,omitempty" xml:"TotalDropSession,omitempty"`
	// The number of sessions blocked by the VPC firewall.
	//
	// example:
	//
	// 100
	VpcDropSession *int64 `json:"VpcDropSession,omitempty" xml:"VpcDropSession,omitempty"`
}

func (s DescribeFirewallDropTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) GetInternetDropSession() *int64 {
	return s.InternetDropSession
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) GetNatDropSession() *int64 {
	return s.NatDropSession
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) GetTotalDropSession() *int64 {
	return s.TotalDropSession
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) GetVpcDropSession() *int64 {
	return s.VpcDropSession
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) SetInternetDropSession(v int64) *DescribeFirewallDropTrendResponseBodyDataList {
	s.InternetDropSession = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) SetNatDropSession(v int64) *DescribeFirewallDropTrendResponseBodyDataList {
	s.NatDropSession = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) SetTime(v int64) *DescribeFirewallDropTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) SetTotalDropSession(v int64) *DescribeFirewallDropTrendResponseBodyDataList {
	s.TotalDropSession = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) SetVpcDropSession(v int64) *DescribeFirewallDropTrendResponseBodyDataList {
	s.VpcDropSession = &v
	return s
}

func (s *DescribeFirewallDropTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
