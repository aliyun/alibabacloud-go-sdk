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
	// The page number to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IP version of the asset. Valid values:
	//
	// - **4*	- (default): IPv4
	//
	// - **6**: IPv6
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// Filters for assets discovered within a specific time window. Valid values:
	//
	// - **discovered in 1 hour**: The asset was added within the last hour.
	//
	// - **discovered in 1 day**: The asset was added within the last day.
	//
	// - **discovered in 7 days**: The asset was added within the last 7 days.
	//
	// example:
	//
	// discovered in 1 hour
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// Specifies whether to query information about outbound traffic.
	//
	// example:
	//
	// true
	OutStatistic *string `json:"OutStatistic,omitempty" xml:"OutStatistic,omitempty"`
	// The number of assets to return per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of your Cloud Firewall instance.
	//
	// > For more information about the regions that Cloud Firewall supports, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The asset type. Valid values:
	//
	// - **BastionHostEgressIP**: The egress IP address of a Bastionhost instance.
	//
	// - **BastionHostIngressIP**: The ingress IP address of a Bastionhost instance.
	//
	// - **EcsEIP**: The Elastic IP Address (EIP) of an ECS instance.
	//
	// - **EcsPublicIP**: The public IP address of an ECS instance.
	//
	// - **EIP**: An Elastic IP Address (EIP).
	//
	// - **EniEIP**: The EIP of an elastic network interface (ENI).
	//
	// - **NatEIP**: The EIP of a NAT Gateway instance.
	//
	// - **SlbEIP**: The EIP of a Server Load Balancer (SLB) or Classic Load Balancer (CLB) instance.
	//
	// - **SlbPublicIP**: The public IP address of a Server Load Balancer (SLB) or Classic Load Balancer (CLB) instance.
	//
	// - **NatPublicIP**: The public IP address of a NAT Gateway instance.
	//
	// - **HAVIP**: A High-availability Virtual IP (HAVIP).
	//
	// - **NlbEIP**: The EIP of a Network Load Balancer (NLB) instance.
	//
	// - **ApiGatewayEIP**: The public IP address of an API Gateway instance.
	//
	// - **AlbEIP**: The EIP of an Application Load Balancer (ALB) instance.
	//
	// - **AiGatewayEIP**: The public IP address of an AI Gateway instance.
	//
	// - **GaEIP**: The EIP of a Global Accelerator (GA) instance.
	//
	// - **SwasEIP**: The public IP address of a Simple Application Server instance.
	//
	// - **EcdEIP**: The public IP address of a Wuying instance.
	//
	// - **BastionHostIP**: The IP address of a Bastionhost instance.
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
	// The status of the data leak detection feature.
	//
	// example:
	//
	// open
	SensitiveStatus *string `json:"SensitiveStatus,omitempty" xml:"SensitiveStatus,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// - **pass**: The security group policy is enforced.
	//
	// - **block**: The security group policy is not enforced.
	//
	// - **unsupport**: The asset does not support security group policies.
	//
	// > If you do not specify this parameter, assets are queried regardless of the security group policy status.
	//
	// example:
	//
	// pass
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The protection status of the asset. Valid values:
	//
	// - **open**: Protection is enabled.
	//
	// - **opening**: Protection is being enabled.
	//
	// - **closed**: Protection is disabled.
	//
	// - **closing**: Protection is being disabled.
	//
	// > If you do not specify this parameter, assets are queried regardless of their protection status.
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// eip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type of the user. Valid values:
	//
	// - **buy*	- (default): A user with a paid subscription.
	//
	// - **free**: A user on the free tier.
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
