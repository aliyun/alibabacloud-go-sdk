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
	// The information about assets protected by Cloud Firewall.
	Assets []*DescribeAssetListResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of assets protected by Cloud Firewall.
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
	if s.Assets != nil {
		for _, item := range s.Assets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAssetListResponseBodyAssets struct {
	// The UID of the Alibaba Cloud account.
	//
	// > The primary account of the Cloud Firewall member account.
	//
	// example:
	//
	// 158039427902****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the bound asset instance.
	//
	// example:
	//
	// i-8vbdrjrxzt78****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The name of the bound asset instance.
	//
	// example:
	//
	// instance01
	BindInstanceName *string `json:"BindInstanceName,omitempty" xml:"BindInstanceName,omitempty"`
	// The time when Cloud Firewall discovered the asset. Time format: YYYY-MM-DD HH:mm:ss.
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
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The IP address version of the asset protected by Cloud Firewall.
	//
	// Valid values:
	//
	// - **4**: Indicates an IPv4 address.
	//
	// - **6**: Indicates an IPv6 address.
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The outbound traffic in the last 7 days.
	//
	// example:
	//
	// 0
	Last7DayOutTrafficBytes *int64 `json:"Last7DayOutTrafficBytes,omitempty" xml:"Last7DayOutTrafficBytes,omitempty"`
	// The UID of the Cloud Firewall member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance name of the asset protected by Cloud Firewall.
	//
	// example:
	//
	// instance01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The remarks of the asset. Valid values:
	//
	// - **REGION_NOT_SUPPORT**: Region not supported.
	//
	// - **NETWORK_NOT_SUPPORT**: Network not supported.
	//
	// example:
	//
	// REGION_NOT_SUPPORT
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The firewall status. Valid values:
	//
	// - **open**: Protected.
	//
	// - **opening**: Protection enabling.
	//
	// - **closed**: Not protected.
	//
	// - **closing**: Protection disabling.
	//
	// example:
	//
	// open
	ProtectStatus *string `json:"ProtectStatus,omitempty" xml:"ProtectStatus,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// Indicates whether the region of the asset supports enabling Cloud Firewall protection. Valid values:
	//
	// - **enable**: Supported.
	//
	// - **disable**: Not supported.
	//
	// example:
	//
	// enable
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The asset instance ID.
	//
	// example:
	//
	// i-8vbdrjrxzt78****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
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
	// The risk level of the asset. Valid values:
	//
	// - **low**: Low risk.
	//
	// - **middle**: Medium risk.
	//
	// - **hight**: High risk.
	//
	// > This parameter is returned only when the value of UserType is free.
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of data leakage detection.
	//
	// example:
	//
	// open
	SensitiveDataStatus *string `json:"SensitiveDataStatus,omitempty" xml:"SensitiveDataStatus,omitempty"`
	// The security group policy. Valid values:
	//
	// - **pass**: Delivered.
	//
	// - **block**: Not delivered.
	//
	// - **unsupport**: Not supported.
	//
	// example:
	//
	// block
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The last security group status detection time, in timestamp format. Unit: seconds.
	//
	// example:
	//
	// 1615082937
	SgStatusTime *int64 `json:"SgStatusTime,omitempty" xml:"SgStatusTime,omitempty"`
	// The traffic diversion support status of the asset. Valid values:
	//
	// - **enable**: Traffic diversion supported.
	//
	// - **disable**: Traffic diversion not supported.
	//
	// example:
	//
	// enable
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// Deprecated
	//
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
