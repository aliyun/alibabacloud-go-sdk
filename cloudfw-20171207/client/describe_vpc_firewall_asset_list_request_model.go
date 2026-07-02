// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeVpcFirewallAssetListRequest
	GetAppName() *string
	SetAssetIP(v string) *DescribeVpcFirewallAssetListRequest
	GetAssetIP() *string
	SetCurrentPage(v string) *DescribeVpcFirewallAssetListRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeVpcFirewallAssetListRequest
	GetDirection() *string
	SetEcsInstanceId(v string) *DescribeVpcFirewallAssetListRequest
	GetEcsInstanceId() *string
	SetEcsInstanceName(v string) *DescribeVpcFirewallAssetListRequest
	GetEcsInstanceName() *string
	SetEndTime(v string) *DescribeVpcFirewallAssetListRequest
	GetEndTime() *string
	SetIPProtocol(v string) *DescribeVpcFirewallAssetListRequest
	GetIPProtocol() *string
	SetIsAITraffic(v string) *DescribeVpcFirewallAssetListRequest
	GetIsAITraffic() *string
	SetLang(v string) *DescribeVpcFirewallAssetListRequest
	GetLang() *string
	SetOrder(v string) *DescribeVpcFirewallAssetListRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeVpcFirewallAssetListRequest
	GetPageSize() *string
	SetPeerVpcId(v string) *DescribeVpcFirewallAssetListRequest
	GetPeerVpcId() *string
	SetPort(v string) *DescribeVpcFirewallAssetListRequest
	GetPort() *string
	SetRiskLevel(v string) *DescribeVpcFirewallAssetListRequest
	GetRiskLevel() *string
	SetSort(v string) *DescribeVpcFirewallAssetListRequest
	GetSort() *string
	SetStartTime(v string) *DescribeVpcFirewallAssetListRequest
	GetStartTime() *string
	SetVpcId(v string) *DescribeVpcFirewallAssetListRequest
	GetVpcId() *string
}

type DescribeVpcFirewallAssetListRequest struct {
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The traffic direction.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The instance ID of the ECS instance.
	//
	// example:
	//
	// i-hp3ez3rs9bxwt034****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the asset instance.
	//
	// example:
	//
	// test-ecs
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The end time. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1756952150
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// tcp
	IPProtocol *string `json:"IPProtocol,omitempty" xml:"IPProtocol,omitempty"`
	// Specifies whether to collect statistics only on traffic that accesses AI services.
	//
	// example:
	//
	// true
	IsAITraffic *string `json:"IsAITraffic,omitempty" xml:"IsAITraffic,omitempty"`
	// The language type of the request and response messages.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance ID of the peer VPC instance.
	//
	// example:
	//
	// vpc-90rq0anm6t8vbwbo****
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The port number.
	//
	// example:
	//
	// 5234
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The risk level.
	//
	// example:
	//
	// 3
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sort method.
	//
	// example:
	//
	// SessionCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The start time. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1534408189
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance ID of the VPC-connected instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9ulqcvly23w31zkh8sm****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallAssetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVpcFirewallAssetListRequest) GetAssetIP() *string {
	return s.AssetIP
}

func (s *DescribeVpcFirewallAssetListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallAssetListRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeVpcFirewallAssetListRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeVpcFirewallAssetListRequest) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *DescribeVpcFirewallAssetListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallAssetListRequest) GetIPProtocol() *string {
	return s.IPProtocol
}

func (s *DescribeVpcFirewallAssetListRequest) GetIsAITraffic() *string {
	return s.IsAITraffic
}

func (s *DescribeVpcFirewallAssetListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallAssetListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallAssetListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallAssetListRequest) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *DescribeVpcFirewallAssetListRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeVpcFirewallAssetListRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeVpcFirewallAssetListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallAssetListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallAssetListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallAssetListRequest) SetAppName(v string) *DescribeVpcFirewallAssetListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetAssetIP(v string) *DescribeVpcFirewallAssetListRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetCurrentPage(v string) *DescribeVpcFirewallAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetDirection(v string) *DescribeVpcFirewallAssetListRequest {
	s.Direction = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetEcsInstanceId(v string) *DescribeVpcFirewallAssetListRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetEcsInstanceName(v string) *DescribeVpcFirewallAssetListRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetEndTime(v string) *DescribeVpcFirewallAssetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetIPProtocol(v string) *DescribeVpcFirewallAssetListRequest {
	s.IPProtocol = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetIsAITraffic(v string) *DescribeVpcFirewallAssetListRequest {
	s.IsAITraffic = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetLang(v string) *DescribeVpcFirewallAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetOrder(v string) *DescribeVpcFirewallAssetListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetPageSize(v string) *DescribeVpcFirewallAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetPeerVpcId(v string) *DescribeVpcFirewallAssetListRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetPort(v string) *DescribeVpcFirewallAssetListRequest {
	s.Port = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetRiskLevel(v string) *DescribeVpcFirewallAssetListRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetSort(v string) *DescribeVpcFirewallAssetListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetStartTime(v string) *DescribeVpcFirewallAssetListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) SetVpcId(v string) *DescribeVpcFirewallAssetListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallAssetListRequest) Validate() error {
	return dara.Validate(s)
}
