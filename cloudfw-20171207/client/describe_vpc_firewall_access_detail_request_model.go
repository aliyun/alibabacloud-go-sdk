// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAccessDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetIP(v string) *DescribeVpcFirewallAccessDetailRequest
	GetAssetIP() *string
	SetCurrentPage(v string) *DescribeVpcFirewallAccessDetailRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeVpcFirewallAccessDetailRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeVpcFirewallAccessDetailRequest
	GetEndTime() *string
	SetIPProtocol(v string) *DescribeVpcFirewallAccessDetailRequest
	GetIPProtocol() *string
	SetLang(v string) *DescribeVpcFirewallAccessDetailRequest
	GetLang() *string
	SetOrder(v string) *DescribeVpcFirewallAccessDetailRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPageSize() *string
	SetPeerAssetIP(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPeerAssetIP() *string
	SetPeerAssetInstanceId(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPeerAssetInstanceId() *string
	SetPeerAssetInstanceName(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPeerAssetInstanceName() *string
	SetPeerVpcId(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPeerVpcId() *string
	SetPort(v string) *DescribeVpcFirewallAccessDetailRequest
	GetPort() *string
	SetRiskLevel(v string) *DescribeVpcFirewallAccessDetailRequest
	GetRiskLevel() *string
	SetSort(v string) *DescribeVpcFirewallAccessDetailRequest
	GetSort() *string
	SetStartTime(v string) *DescribeVpcFirewallAccessDetailRequest
	GetStartTime() *string
	SetVpcId(v string) *DescribeVpcFirewallAccessDetailRequest
	GetVpcId() *string
}

type DescribeVpcFirewallAccessDetailRequest struct {
	// example:
	//
	// 192.0.XX.XX
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1729042555
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// tcp
	IPProtocol *string `json:"IPProtocol,omitempty" xml:"IPProtocol,omitempty"`
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
	// 10.125.1.XX
	PeerAssetIP *string `json:"PeerAssetIP,omitempty" xml:"PeerAssetIP,omitempty"`
	// example:
	//
	// i-123451
	PeerAssetInstanceId *string `json:"PeerAssetInstanceId,omitempty" xml:"PeerAssetInstanceId,omitempty"`
	// example:
	//
	// ecs22
	PeerAssetInstanceName *string `json:"PeerAssetInstanceName,omitempty" xml:"PeerAssetInstanceName,omitempty"`
	// example:
	//
	// vpc-90rq0anm6t8vbwbo****
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 3
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// InBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1655778046
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze4xj5kmb5udb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallAccessDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAccessDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetAssetIP() *string {
	return s.AssetIP
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetIPProtocol() *string {
	return s.IPProtocol
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPeerAssetIP() *string {
	return s.PeerAssetIP
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPeerAssetInstanceId() *string {
	return s.PeerAssetInstanceId
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPeerAssetInstanceName() *string {
	return s.PeerAssetInstanceName
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallAccessDetailRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetAssetIP(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetCurrentPage(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetDirection(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.Direction = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetEndTime(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetIPProtocol(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.IPProtocol = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetLang(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetOrder(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPageSize(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPeerAssetIP(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.PeerAssetIP = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPeerAssetInstanceId(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.PeerAssetInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPeerAssetInstanceName(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.PeerAssetInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPeerVpcId(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetPort(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.Port = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetRiskLevel(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetSort(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetStartTime(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) SetVpcId(v string) *DescribeVpcFirewallAccessDetailRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailRequest) Validate() error {
	return dara.Validate(s)
}
