// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeAssetListRequest
	GetCurrentPage() *string
	SetIpVersion(v string) *DescribeAssetListRequest
	GetIpVersion() *string
	SetLang(v string) *DescribeAssetListRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeAssetListRequest
	GetMemberUid() *int64
	SetNewResourceTag(v string) *DescribeAssetListRequest
	GetNewResourceTag() *string
	SetOutStatistic(v string) *DescribeAssetListRequest
	GetOutStatistic() *string
	SetPageSize(v string) *DescribeAssetListRequest
	GetPageSize() *string
	SetRegionNo(v string) *DescribeAssetListRequest
	GetRegionNo() *string
	SetResourceType(v string) *DescribeAssetListRequest
	GetResourceType() *string
	SetSearchItem(v string) *DescribeAssetListRequest
	GetSearchItem() *string
	SetSensitiveStatus(v string) *DescribeAssetListRequest
	GetSensitiveStatus() *string
	SetSgStatus(v string) *DescribeAssetListRequest
	GetSgStatus() *string
	SetStatus(v string) *DescribeAssetListRequest
	GetStatus() *string
	SetType(v string) *DescribeAssetListRequest
	GetType() *string
	SetUserType(v string) *DescribeAssetListRequest
	GetUserType() *string
}

type DescribeAssetListRequest struct {
	// The page number of the current page in a paginated query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IP version of the assets protected by Cloud Firewall. Valid values:
	//
	// - **4*	- (default): IPv4.
	//
	// - **6**: IPv6.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language type of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the Cloud Firewall member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the asset was discovered. Valid values:
	//
	// - **discovered in 1 hour**: The asset was discovered within 1 hour.
	//
	// - **discovered in 1 day**: The asset was discovered within 1 day.
	//
	// - **discovered in 7 days**: The asset was discovered within 7 days.
	//
	// example:
	//
	// discovered in 1 hour
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// Specifies whether to query outbound traffic information.
	//
	// example:
	//
	// true
	OutStatistic *string `json:"OutStatistic,omitempty" xml:"OutStatistic,omitempty"`
	// The number of Cloud Firewall-protected assets to display on each page in a paginated query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Cloud Firewall.
	//
	// > For more information about regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The asset type. Valid values:
	//
	// - **BastionHostEgressIP**: Bastion host egress IP.
	//
	// - **BastionHostIngressIP**: Bastion host ingress IP.
	//
	// - **EcsEIP**: ECS EIP.
	//
	// - **EcsPublicIP**: ECS public IP.
	//
	// - **EIP**: Elastic IP address.
	//
	// - **EniEIP**: Elastic network interface EIP.
	//
	// - **NatEIP**: NAT EIP.
	//
	// - **SlbEIP**: SLB EIP (CLB EIP).
	//
	// - **SlbPublicIP**: SLB public IP (CLB public IP).
	//
	// - **NatPublicIP**: NAT public IP.
	//
	// - **HAVIP**: High-availability virtual IP.
	//
	// - **NlbEIP**: NLB EIP.
	//
	// - **ApiGatewayEIP**: API Gateway public IP.
	//
	// - **AlbEIP**: ALB EIP.
	//
	// - **AiGatewayEIP**: AI Gateway public IP.
	//
	// - **GaEIP**: GA EIP.
	//
	// - **SwasEIP**: Simple Application Server public IP.
	//
	// - **EcdEIP**: Elastic Desktop Service public IP.
	//
	// - **BastionHostIP**: Bastion host IP.
	//
	// example:
	//
	// EIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The IP address or instance ID of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The status of data leakage detection.
	//
	// example:
	//
	// open
	SensitiveStatus *string `json:"SensitiveStatus,omitempty" xml:"SensitiveStatus,omitempty"`
	// The security group policy status. Valid values:
	//
	// - **pass**: Delivered.
	//
	// - **block**: Not delivered.
	//
	// - **unsupport**: Not supported.
	//
	// > If this parameter is not set, all security group policy statuses are queried.
	//
	// example:
	//
	// pass
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The Cloud Firewall status. Valid values:
	//
	// - **open**: Protected.
	//
	// - **opening**: Protection enabling.
	//
	// - **closed**: Not protected.
	//
	// - **closing**: Protection disabling.
	//
	// > If this parameter is not set, all firewall statuses are queried.
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// eip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user type. Valid values:
	//
	// - **buy*	- (default): Paid user.
	//
	// - **free**: Free user.
	//
	// example:
	//
	// buy
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeAssetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAssetListRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeAssetListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAssetListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeAssetListRequest) GetNewResourceTag() *string {
	return s.NewResourceTag
}

func (s *DescribeAssetListRequest) GetOutStatistic() *string {
	return s.OutStatistic
}

func (s *DescribeAssetListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAssetListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAssetListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeAssetListRequest) GetSearchItem() *string {
	return s.SearchItem
}

func (s *DescribeAssetListRequest) GetSensitiveStatus() *string {
	return s.SensitiveStatus
}

func (s *DescribeAssetListRequest) GetSgStatus() *string {
	return s.SgStatus
}

func (s *DescribeAssetListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAssetListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAssetListRequest) GetUserType() *string {
	return s.UserType
}

func (s *DescribeAssetListRequest) SetCurrentPage(v string) *DescribeAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAssetListRequest) SetIpVersion(v string) *DescribeAssetListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListRequest) SetLang(v string) *DescribeAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetListRequest) SetMemberUid(v int64) *DescribeAssetListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListRequest) SetNewResourceTag(v string) *DescribeAssetListRequest {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListRequest) SetOutStatistic(v string) *DescribeAssetListRequest {
	s.OutStatistic = &v
	return s
}

func (s *DescribeAssetListRequest) SetPageSize(v string) *DescribeAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAssetListRequest) SetRegionNo(v string) *DescribeAssetListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAssetListRequest) SetResourceType(v string) *DescribeAssetListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListRequest) SetSearchItem(v string) *DescribeAssetListRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribeAssetListRequest) SetSensitiveStatus(v string) *DescribeAssetListRequest {
	s.SensitiveStatus = &v
	return s
}

func (s *DescribeAssetListRequest) SetSgStatus(v string) *DescribeAssetListRequest {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListRequest) SetStatus(v string) *DescribeAssetListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAssetListRequest) SetType(v string) *DescribeAssetListRequest {
	s.Type = &v
	return s
}

func (s *DescribeAssetListRequest) SetUserType(v string) *DescribeAssetListRequest {
	s.UserType = &v
	return s
}

func (s *DescribeAssetListRequest) Validate() error {
	return dara.Validate(s)
}
