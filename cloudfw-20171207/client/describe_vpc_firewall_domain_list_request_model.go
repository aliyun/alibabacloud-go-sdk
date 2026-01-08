// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeVpcFirewallDomainListRequest
	GetCurrentPage() *string
	SetDomain(v string) *DescribeVpcFirewallDomainListRequest
	GetDomain() *string
	SetEndTime(v string) *DescribeVpcFirewallDomainListRequest
	GetEndTime() *string
	SetIsAITraffic(v string) *DescribeVpcFirewallDomainListRequest
	GetIsAITraffic() *string
	SetLang(v string) *DescribeVpcFirewallDomainListRequest
	GetLang() *string
	SetOrder(v string) *DescribeVpcFirewallDomainListRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeVpcFirewallDomainListRequest
	GetPageSize() *string
	SetSort(v string) *DescribeVpcFirewallDomainListRequest
	GetSort() *string
	SetSrcIP(v string) *DescribeVpcFirewallDomainListRequest
	GetSrcIP() *string
	SetSrcVpcId(v string) *DescribeVpcFirewallDomainListRequest
	GetSrcVpcId() *string
	SetStartTime(v string) *DescribeVpcFirewallDomainListRequest
	GetStartTime() *string
}

type DescribeVpcFirewallDomainListRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// true
	IsAITraffic *string `json:"IsAITraffic,omitempty" xml:"IsAITraffic,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// SessionCount
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// desc
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

func (s DescribeVpcFirewallDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallDomainListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVpcFirewallDomainListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallDomainListRequest) GetIsAITraffic() *string {
	return s.IsAITraffic
}

func (s *DescribeVpcFirewallDomainListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallDomainListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallDomainListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallDomainListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallDomainListRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeVpcFirewallDomainListRequest) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeVpcFirewallDomainListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallDomainListRequest) SetCurrentPage(v string) *DescribeVpcFirewallDomainListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetDomain(v string) *DescribeVpcFirewallDomainListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetEndTime(v string) *DescribeVpcFirewallDomainListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetIsAITraffic(v string) *DescribeVpcFirewallDomainListRequest {
	s.IsAITraffic = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetLang(v string) *DescribeVpcFirewallDomainListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetOrder(v string) *DescribeVpcFirewallDomainListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetPageSize(v string) *DescribeVpcFirewallDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetSort(v string) *DescribeVpcFirewallDomainListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetSrcIP(v string) *DescribeVpcFirewallDomainListRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetSrcVpcId(v string) *DescribeVpcFirewallDomainListRequest {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) SetStartTime(v string) *DescribeVpcFirewallDomainListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainListRequest) Validate() error {
	return dara.Validate(s)
}
