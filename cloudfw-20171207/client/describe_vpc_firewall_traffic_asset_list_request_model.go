// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficAssetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetCurrentPage() *string
	SetDomain(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetDomain() *string
	SetEndTime(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetEndTime() *string
	SetIP(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetIP() *string
	SetIsAITraffic(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetIsAITraffic() *string
	SetLang(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetLang() *string
	SetOrder(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetPageSize() *string
	SetSort(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetSort() *string
	SetStartTime(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetStartTime() *string
	SetVpcId(v string) *DescribeVpcFirewallTrafficAssetListRequest
	GetVpcId() *string
}

type DescribeVpcFirewallTrafficAssetListRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// www.****.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 47.92.x.x
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// vpc-m5ewlqkuf7orclr1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallTrafficAssetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetIP() *string {
	return s.IP
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetIsAITraffic() *string {
	return s.IsAITraffic
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetCurrentPage(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetDomain(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetEndTime(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetIP(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.IP = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetIsAITraffic(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.IsAITraffic = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetLang(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetOrder(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetPageSize(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetSort(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetStartTime(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) SetVpcId(v string) *DescribeVpcFirewallTrafficAssetListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListRequest) Validate() error {
	return dara.Validate(s)
}
