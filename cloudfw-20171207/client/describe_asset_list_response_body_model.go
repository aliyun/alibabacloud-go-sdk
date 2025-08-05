// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeAssetListResponseBodyAssets) *DescribeAssetListResponseBody
	GetAssets() []*DescribeAssetListResponseBodyAssets
	SetRequestId(v string) *DescribeAssetListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAssetListResponseBody
	GetTotalCount() *int32
}

type DescribeAssetListResponseBody struct {
	// The assets that are protected by Cloud Firewall.
	Assets []*DescribeAssetListResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets that are protected by Cloud Firewall.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBody) GetAssets() []*DescribeAssetListResponseBodyAssets {
	return s.Assets
}

func (s *DescribeAssetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAssetListResponseBody) SetAssets(v []*DescribeAssetListResponseBodyAssets) *DescribeAssetListResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeAssetListResponseBody) SetRequestId(v string) *DescribeAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetListResponseBody) SetTotalCount(v int32) *DescribeAssetListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAssetListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetListResponseBodyAssets struct {
	// The UID of the Alibaba Cloud account.
	//
	// >  The value of this parameter indicates the management account to which the member is added.
	//
	// example:
	//
	// 158039427902****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the cloud resource with which the asset is associated.
	//
	// example:
	//
	// i-8vbdrjrxzt78****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// instance01
	BindInstanceName *string `json:"BindInstanceName,omitempty" xml:"BindInstanceName,omitempty"`
	// The timestamp when the asset is added to Cloud Firewall.
	//
	// example:
	//
	// 2023-02-28 10:29:58
	CreateTimeStamp *string `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The internal IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// 	- **4**: IPv4
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Outbound traffic in the last 7 days.
	//
	// example:
	//
	// 0
	Last7DayOutTrafficBytes *int64 `json:"Last7DayOutTrafficBytes,omitempty" xml:"Last7DayOutTrafficBytes,omitempty"`
	// The UID of the member.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance name of the asset that is protected by Cloud Firewall.
	//
	// example:
	//
	// instance01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The remarks of the asset. Valid values:
	//
	// 	- **REGION_NOT_SUPPORT**: The region is not supported.
	//
	// 	- **NETWORK_NOT_SUPPORT**: The network is not supported.
	//
	// example:
	//
	// REGION_NOT_SUPPORT
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The status of the firewall. Valid values:
	//
	// 	- **open**: enabled.
	//
	// 	- **opening**: being enabled.
	//
	// 	- **closed**: disabled.
	//
	// 	- **closing**: being disabled.
	//
	// example:
	//
	// open
	ProtectStatus *string `json:"ProtectStatus,omitempty" xml:"ProtectStatus,omitempty"`
	// The ID of the region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// Indicates whether the firewall is supported in the region in which the asset resides. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-8vbdrjrxzt78****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
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
	// 	- **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	//
	// 	- **SlbPublicIP**: the public IP address of an SLB instance
	//
	// 	- **NatPublicIP**: the public IP address of a NAT gateway
	//
	// 	- **HAVIP**: the high-availability virtual IP address (HAVIP)
	//
	// example:
	//
	// EIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The risk level of the asset. Valid values:
	//
	// 	- **low**: low
	//
	// 	- **middle**: medium
	//
	// 	- **hight**: high
	//
	// >  The value of this parameter is returned only when the UserType parameter is set to free.
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Data leakage detection enabled status.
	//
	// example:
	//
	// open
	SensitiveDataStatus *string `json:"SensitiveDataStatus,omitempty" xml:"SensitiveDataStatus,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// 	- **pass**: applied
	//
	// 	- **block**: not applied
	//
	// 	- **unsupport**: unsupported
	//
	// example:
	//
	// block
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The time when the status of the security group was last checked. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615082937
	SgStatusTime *int64 `json:"SgStatusTime,omitempty" xml:"SgStatusTime,omitempty"`
	// Indicates whether traffic redirection is supported for the asset. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// eip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetListResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetListResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBodyAssets) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeAssetListResponseBodyAssets) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *DescribeAssetListResponseBodyAssets) GetBindInstanceName() *string {
	return s.BindInstanceName
}

func (s *DescribeAssetListResponseBodyAssets) GetCreateTimeStamp() *string {
	return s.CreateTimeStamp
}

func (s *DescribeAssetListResponseBodyAssets) GetInternetAddress() *string {
	return s.InternetAddress
}

func (s *DescribeAssetListResponseBodyAssets) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *DescribeAssetListResponseBodyAssets) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeAssetListResponseBodyAssets) GetLast7DayOutTrafficBytes() *int64 {
	return s.Last7DayOutTrafficBytes
}

func (s *DescribeAssetListResponseBodyAssets) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeAssetListResponseBodyAssets) GetName() *string {
	return s.Name
}

func (s *DescribeAssetListResponseBodyAssets) GetNewResourceTag() *string {
	return s.NewResourceTag
}

func (s *DescribeAssetListResponseBodyAssets) GetNote() *string {
	return s.Note
}

func (s *DescribeAssetListResponseBodyAssets) GetProtectStatus() *string {
	return s.ProtectStatus
}

func (s *DescribeAssetListResponseBodyAssets) GetRegionID() *string {
	return s.RegionID
}

func (s *DescribeAssetListResponseBodyAssets) GetRegionStatus() *string {
	return s.RegionStatus
}

func (s *DescribeAssetListResponseBodyAssets) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeAssetListResponseBodyAssets) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeAssetListResponseBodyAssets) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAssetListResponseBodyAssets) GetSensitiveDataStatus() *string {
	return s.SensitiveDataStatus
}

func (s *DescribeAssetListResponseBodyAssets) GetSgStatus() *string {
	return s.SgStatus
}

func (s *DescribeAssetListResponseBodyAssets) GetSgStatusTime() *int64 {
	return s.SgStatusTime
}

func (s *DescribeAssetListResponseBodyAssets) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *DescribeAssetListResponseBodyAssets) GetType() *string {
	return s.Type
}

func (s *DescribeAssetListResponseBodyAssets) SetAliUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.AliUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceName(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceName = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetCreateTimeStamp(v string) *DescribeAssetListResponseBodyAssets {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetInternetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.InternetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIntranetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.IntranetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIpVersion(v int32) *DescribeAssetListResponseBodyAssets {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetLast7DayOutTrafficBytes(v int64) *DescribeAssetListResponseBodyAssets {
	s.Last7DayOutTrafficBytes = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetMemberUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetName(v string) *DescribeAssetListResponseBodyAssets {
	s.Name = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNewResourceTag(v string) *DescribeAssetListResponseBodyAssets {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNote(v string) *DescribeAssetListResponseBodyAssets {
	s.Note = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetProtectStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.ProtectStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionID(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionID = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceType(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRiskLevel(v string) *DescribeAssetListResponseBodyAssets {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSensitiveDataStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SensitiveDataStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatusTime(v int64) *DescribeAssetListResponseBodyAssets {
	s.SgStatusTime = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSyncStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SyncStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetType(v string) *DescribeAssetListResponseBodyAssets {
	s.Type = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
