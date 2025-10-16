// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVpcListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeAccessInstanceVpcListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAccessInstanceVpcListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessInstanceVpcListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAccessInstanceVpcListResponseBody
	GetTotalCount() *int64
	SetVpcList(v []*DescribeAccessInstanceVpcListResponseBodyVpcList) *DescribeAccessInstanceVpcListResponseBody
	GetVpcList() []*DescribeAccessInstanceVpcListResponseBodyVpcList
}

type DescribeAccessInstanceVpcListResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 450D47F5-956E-543E-8502-2F71C8C54E72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcList    []*DescribeAccessInstanceVpcListResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s DescribeAccessInstanceVpcListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVpcListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVpcListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAccessInstanceVpcListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessInstanceVpcListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessInstanceVpcListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAccessInstanceVpcListResponseBody) GetVpcList() []*DescribeAccessInstanceVpcListResponseBodyVpcList {
	return s.VpcList
}

func (s *DescribeAccessInstanceVpcListResponseBody) SetPageNo(v int32) *DescribeAccessInstanceVpcListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBody) SetPageSize(v int32) *DescribeAccessInstanceVpcListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBody) SetRequestId(v string) *DescribeAccessInstanceVpcListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBody) SetTotalCount(v int64) *DescribeAccessInstanceVpcListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBody) SetVpcList(v []*DescribeAccessInstanceVpcListResponseBodyVpcList) *DescribeAccessInstanceVpcListResponseBody {
	s.VpcList = v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBody) Validate() error {
	if s.VpcList != nil {
		for _, item := range s.VpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessInstanceVpcListResponseBodyVpcList struct {
	// example:
	//
	// false
	FirewallVpc *bool `json:"FirewallVpc,omitempty" xml:"FirewallVpc,omitempty"`
	// example:
	//
	// vpc-2ze1t4irqj0fljlbb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// Cloud_Firewall_VPC
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeAccessInstanceVpcListResponseBodyVpcList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVpcListResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) GetFirewallVpc() *bool {
	return s.FirewallVpc
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) SetFirewallVpc(v bool) *DescribeAccessInstanceVpcListResponseBodyVpcList {
	s.FirewallVpc = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) SetVpcId(v string) *DescribeAccessInstanceVpcListResponseBodyVpcList {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) SetVpcName(v string) *DescribeAccessInstanceVpcListResponseBodyVpcList {
	s.VpcName = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponseBodyVpcList) Validate() error {
	return dara.Validate(s)
}
