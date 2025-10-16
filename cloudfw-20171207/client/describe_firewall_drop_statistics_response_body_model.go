// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallDropStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody
	GetAclDropCnt() *int64
	SetIpsDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody
	GetIpsDropCnt() *int64
	SetRequestId(v string) *DescribeFirewallDropStatisticsResponseBody
	GetRequestId() *string
	SetTotalDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody
	GetTotalDropCnt() *int64
	SetVulnDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody
	GetVulnDropCnt() *int64
}

type DescribeFirewallDropStatisticsResponseBody struct {
	// example:
	//
	// 20
	AclDropCnt *int64 `json:"AclDropCnt,omitempty" xml:"AclDropCnt,omitempty"`
	// example:
	//
	// 20
	IpsDropCnt *int64 `json:"IpsDropCnt,omitempty" xml:"IpsDropCnt,omitempty"`
	// example:
	//
	// BEA1D173-D5DB-582E-9637-438D5CE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	TotalDropCnt *int64 `json:"TotalDropCnt,omitempty" xml:"TotalDropCnt,omitempty"`
	// example:
	//
	// 10
	VulnDropCnt *int64 `json:"VulnDropCnt,omitempty" xml:"VulnDropCnt,omitempty"`
}

func (s DescribeFirewallDropStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropStatisticsResponseBody) GetAclDropCnt() *int64 {
	return s.AclDropCnt
}

func (s *DescribeFirewallDropStatisticsResponseBody) GetIpsDropCnt() *int64 {
	return s.IpsDropCnt
}

func (s *DescribeFirewallDropStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallDropStatisticsResponseBody) GetTotalDropCnt() *int64 {
	return s.TotalDropCnt
}

func (s *DescribeFirewallDropStatisticsResponseBody) GetVulnDropCnt() *int64 {
	return s.VulnDropCnt
}

func (s *DescribeFirewallDropStatisticsResponseBody) SetAclDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody {
	s.AclDropCnt = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponseBody) SetIpsDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody {
	s.IpsDropCnt = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponseBody) SetRequestId(v string) *DescribeFirewallDropStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponseBody) SetTotalDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody {
	s.TotalDropCnt = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponseBody) SetVulnDropCnt(v int64) *DescribeFirewallDropStatisticsResponseBody {
	s.VulnDropCnt = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
