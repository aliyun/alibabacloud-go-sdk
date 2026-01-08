// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainRelationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetCurrentPage() *string
	SetDomainList(v []*string) *DescribeVpcFirewallDomainRelationListRequest
	GetDomainList() []*string
	SetDstIP(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetDstIP() *string
	SetDstVpcId(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetDstVpcId() *string
	SetEndTime(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetEndTime() *string
	SetLang(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetLang() *string
	SetOrder(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetPageSize() *string
	SetSort(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetSort() *string
	SetSrcIP(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetSrcIP() *string
	SetSrcVpcId(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetSrcVpcId() *string
	SetStartTime(v string) *DescribeVpcFirewallDomainRelationListRequest
	GetStartTime() *string
}

type DescribeVpcFirewallDomainRelationListRequest struct {
	// example:
	//
	// 1
	CurrentPage *string   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DomainList  []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// 34.136.111.XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// example:
	//
	// vpc-bp10w5nb30r4jzfyc****
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// TotalBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 47.92.x.x
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// vpc-t4nlt09olhpazpoeg****
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVpcFirewallDomainRelationListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainRelationListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetDstVpcId() *string {
	return s.DstVpcId
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeVpcFirewallDomainRelationListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetCurrentPage(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetDomainList(v []*string) *DescribeVpcFirewallDomainRelationListRequest {
	s.DomainList = v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetDstIP(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetDstVpcId(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.DstVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetEndTime(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetLang(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetOrder(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetPageSize(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetSort(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetSrcIP(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetSrcVpcId(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) SetStartTime(v string) *DescribeVpcFirewallDomainRelationListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListRequest) Validate() error {
	return dara.Validate(s)
}
