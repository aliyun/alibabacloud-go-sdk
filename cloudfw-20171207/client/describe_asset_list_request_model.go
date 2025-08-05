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
	// The page number. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall. Valid values:
	//
	// 	- **4**: IPv4 (default)
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is added to Cloud Firewall.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the asset was added. Valid values:
	//
	// 	- **discovered in 1 hour**: within one hour.
	//
	// 	- **discovered in 1 day**: within one day.
	//
	// 	- **discovered in 7 days**: within seven days.
	//
	// example:
	//
	// discovered in 1 hour
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// Whether to query external traffic information.
	//
	// example:
	//
	// true
	OutStatistic *string `json:"OutStatistic,omitempty" xml:"OutStatistic,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// > For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **BastionHostEgressIP**: the egress IP address of a bastion host
	//
	// 	- **BastionHostIngressIP**: the ingress IP address of a bastion host
	//
	// 	- **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	//
	// 	- **EcsPublicIP**: the public IP address of an ECS instance
	//
	// 	- **EIP**: the EIP
	//
	// 	- **EniEIP**: the EIP of an elastic network interface (ENI)
	//
	// 	- **NatEIP**: the EIP of a NAT gateway
	//
	// 	- **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance or a Classic Load Balancer (CLB) instance
	//
	// 	- **SlbPublicIP**: the public IP address of an SLB instance or a CLB instance
	//
	// 	- **NatPublicIP**: the public IP address of a NAT gateway
	//
	// 	- **HAVIP**: the high-availability virtual IP address (HAVIP)
	//
	// example:
	//
	// EIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The instance ID or IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// Data leakage detection activation status.
	//
	// example:
	//
	// open
	SensitiveStatus *string `json:"SensitiveStatus,omitempty" xml:"SensitiveStatus,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// 	- **pass**: delivered
	//
	// 	- **block**: undelivered
	//
	// 	- **unsupport**: unsupported
	//
	// > If you do not specify this parameter, the assets on which security group policies in all states take effect are queried.
	//
	// example:
	//
	// pass
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The status of the firewall. Valid values:
	//
	// 	- **open**: The firewall is enabled.
	//
	// 	- **opening**: The firewall is being enabled.
	//
	// 	- **closed**: The firewall is disabled.
	//
	// 	- **closing**: The firewall is being disabled.
	//
	// > If you do not specify this parameter, the assets that are configured for firewalls in all states are queried.
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
	// The edition of Cloud Firewall. Valid values:
	//
	// 	- **buy**: a paid edition (default)
	//
	// 	- **free**: Free Edition
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
