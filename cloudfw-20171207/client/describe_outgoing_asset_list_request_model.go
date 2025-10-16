// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingAssetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsRegion(v string) *DescribeOutgoingAssetListRequest
	GetAssetsRegion() *string
	SetCurrentPage(v string) *DescribeOutgoingAssetListRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeOutgoingAssetListRequest
	GetEndTime() *string
	SetGroupName(v string) *DescribeOutgoingAssetListRequest
	GetGroupName() *string
	SetIPType(v string) *DescribeOutgoingAssetListRequest
	GetIPType() *string
	SetLang(v string) *DescribeOutgoingAssetListRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeOutgoingAssetListRequest
	GetNatGatewayId() *string
	SetNatGatewayName(v string) *DescribeOutgoingAssetListRequest
	GetNatGatewayName() *string
	SetOrder(v string) *DescribeOutgoingAssetListRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingAssetListRequest
	GetPageSize() *string
	SetPrivateIP(v string) *DescribeOutgoingAssetListRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeOutgoingAssetListRequest
	GetPublicIP() *string
	SetResourceType(v string) *DescribeOutgoingAssetListRequest
	GetResourceType() *string
	SetSecurityRisk(v string) *DescribeOutgoingAssetListRequest
	GetSecurityRisk() *string
	SetSort(v string) *DescribeOutgoingAssetListRequest
	GetSort() *string
	SetStartTime(v string) *DescribeOutgoingAssetListRequest
	GetStartTime() *string
}

type DescribeOutgoingAssetListRequest struct {
	// example:
	//
	// cn-beijing
	AssetsRegion *string `json:"AssetsRegion,omitempty" xml:"AssetsRegion,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1736438400
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// subscribe
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// NatPrivate
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// ngw-bp123456g******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
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
	// 10.200.33.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// example:
	//
	// 47.116.70.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// NatEIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// risk
	SecurityRisk *string `json:"SecurityRisk,omitempty" xml:"SecurityRisk,omitempty"`
	// example:
	//
	// InBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1743647114
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingAssetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingAssetListRequest) GetAssetsRegion() *string {
	return s.AssetsRegion
}

func (s *DescribeOutgoingAssetListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingAssetListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingAssetListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOutgoingAssetListRequest) GetIPType() *string {
	return s.IPType
}

func (s *DescribeOutgoingAssetListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingAssetListRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingAssetListRequest) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeOutgoingAssetListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingAssetListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingAssetListRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingAssetListRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingAssetListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOutgoingAssetListRequest) GetSecurityRisk() *string {
	return s.SecurityRisk
}

func (s *DescribeOutgoingAssetListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingAssetListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingAssetListRequest) SetAssetsRegion(v string) *DescribeOutgoingAssetListRequest {
	s.AssetsRegion = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetCurrentPage(v string) *DescribeOutgoingAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetEndTime(v string) *DescribeOutgoingAssetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetGroupName(v string) *DescribeOutgoingAssetListRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetIPType(v string) *DescribeOutgoingAssetListRequest {
	s.IPType = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetLang(v string) *DescribeOutgoingAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetNatGatewayId(v string) *DescribeOutgoingAssetListRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetNatGatewayName(v string) *DescribeOutgoingAssetListRequest {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetOrder(v string) *DescribeOutgoingAssetListRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetPageSize(v string) *DescribeOutgoingAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetPrivateIP(v string) *DescribeOutgoingAssetListRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetPublicIP(v string) *DescribeOutgoingAssetListRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetResourceType(v string) *DescribeOutgoingAssetListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetSecurityRisk(v string) *DescribeOutgoingAssetListRequest {
	s.SecurityRisk = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetSort(v string) *DescribeOutgoingAssetListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) SetStartTime(v string) *DescribeOutgoingAssetListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingAssetListRequest) Validate() error {
	return dara.Validate(s)
}
