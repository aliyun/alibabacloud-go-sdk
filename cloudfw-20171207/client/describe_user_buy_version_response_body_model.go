// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnectorQuota(v int64) *DescribeUserBuyVersionResponseBody
	GetAckClusterConnectorQuota() *int64
	SetAliUid(v int64) *DescribeUserBuyVersionResponseBody
	GetAliUid() *int64
	SetDefaultBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetDefaultBandwidth() *int64
	SetExpire(v int64) *DescribeUserBuyVersionResponseBody
	GetExpire() *int64
	SetExtensionBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetExtensionBandwidth() *int64
	SetGeneralInstance(v int64) *DescribeUserBuyVersionResponseBody
	GetGeneralInstance() *int64
	SetInstanceId(v string) *DescribeUserBuyVersionResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *DescribeUserBuyVersionResponseBody
	GetInstanceStatus() *string
	SetInternetBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetInternetBandwidth() *int64
	SetIpNumber(v int64) *DescribeUserBuyVersionResponseBody
	GetIpNumber() *int64
	SetLogStatus(v bool) *DescribeUserBuyVersionResponseBody
	GetLogStatus() *bool
	SetLogStorage(v int64) *DescribeUserBuyVersionResponseBody
	GetLogStorage() *int64
	SetMajorVersion(v int64) *DescribeUserBuyVersionResponseBody
	GetMajorVersion() *int64
	SetMaxOverflow(v int64) *DescribeUserBuyVersionResponseBody
	GetMaxOverflow() *int64
	SetNatBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetNatBandwidth() *int64
	SetPrivateDnsConnectorQuota(v int64) *DescribeUserBuyVersionResponseBody
	GetPrivateDnsConnectorQuota() *int64
	SetRequestId(v string) *DescribeUserBuyVersionResponseBody
	GetRequestId() *string
	SetSdl(v int64) *DescribeUserBuyVersionResponseBody
	GetSdl() *int64
	SetStartTime(v int64) *DescribeUserBuyVersionResponseBody
	GetStartTime() *int64
	SetTemporaryBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetTemporaryBandwidth() *int64
	SetThreatIntelligence(v int64) *DescribeUserBuyVersionResponseBody
	GetThreatIntelligence() *int64
	SetUserStatus(v bool) *DescribeUserBuyVersionResponseBody
	GetUserStatus() *bool
	SetVersion(v int32) *DescribeUserBuyVersionResponseBody
	GetVersion() *int32
	SetVpcBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetVpcBandwidth() *int64
	SetVpcNumber(v int64) *DescribeUserBuyVersionResponseBody
	GetVpcNumber() *int64
}

type DescribeUserBuyVersionResponseBody struct {
	// The quota for ACK cluster connectors.
	//
	// example:
	//
	// 5
	AckClusterConnectorQuota *int64 `json:"AckClusterConnectorQuota,omitempty" xml:"AckClusterConnectorQuota,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 119898001566xxxx
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The default bandwidth of the edition.
	//
	// example:
	//
	// 1000
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" xml:"DefaultBandwidth,omitempty"`
	// The expiration time of the Cloud Firewall instance.
	//
	// > The value is a UNIX timestamp in milliseconds.
	//
	// > This parameter does not apply to pay-as-you-go editions.
	//
	// example:
	//
	// 1726934400000
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The extended bandwidth.
	//
	// example:
	//
	// 1000
	ExtensionBandwidth *int64 `json:"ExtensionBandwidth,omitempty" xml:"ExtensionBandwidth,omitempty"`
	// The number of general-purpose instances.
	//
	// example:
	//
	// 10
	GeneralInstance *int64 `json:"GeneralInstance,omitempty" xml:"GeneralInstance,omitempty"`
	// The ID of the Cloud Firewall instance.
	//
	// > This parameter does not apply to trial editions.
	//
	// example:
	//
	// vipcloudfw-cn-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operational status of the Cloud Firewall instance. Valid values:
	//
	// - **normal**: The instance is running as expected.
	//
	// - **init**: The instance is being initialized.
	//
	// - **deleting**: The instance is being deleted.
	//
	// - **abnormal**: The instance is in an abnormal state.
	//
	// - **free**: No valid instance is available.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The purchased traffic processing capability for the Internet firewall.
	//
	// example:
	//
	// 3000
	InternetBandwidth *int64 `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty"`
	// The number of public IP addresses that can be protected.
	//
	// > This parameter applies only to subscription instances.
	//
	// example:
	//
	// 63
	IpNumber *int64 `json:"IpNumber,omitempty" xml:"IpNumber,omitempty"`
	// Indicates whether log delivery is enabled. Valid values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// true
	LogStatus *bool `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// The purchased log storage capacity.
	//
	// > This parameter applies only to subscription instances.
	//
	// example:
	//
	// 3000
	LogStorage *int64 `json:"LogStorage,omitempty" xml:"LogStorage,omitempty"`
	// The major version.
	//
	// example:
	//
	// 1
	MajorVersion *int64 `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// Indicates whether elastic billing for excess traffic is enabled. Valid values:
	//
	// - **1000000**: Enabled
	//
	// - **0**: Disabled
	//
	// > This parameter applies only to subscription instances.
	//
	// example:
	//
	// 0
	MaxOverflow *int64 `json:"MaxOverflow,omitempty" xml:"MaxOverflow,omitempty"`
	// The purchased traffic processing capability for the NAT firewall.
	//
	// example:
	//
	// 3000
	NatBandwidth *int64 `json:"NatBandwidth,omitempty" xml:"NatBandwidth,omitempty"`
	// The quota for private DNS connectors.
	//
	// example:
	//
	// 5
	PrivateDnsConnectorQuota *int64 `json:"PrivateDnsConnectorQuota,omitempty" xml:"PrivateDnsConnectorQuota,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F71B03EE-xxxxx-91D79CC6AA1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether data leakage prevention is enabled.
	//
	// example:
	//
	// 1
	Sdl *int64 `json:"Sdl,omitempty" xml:"Sdl,omitempty"`
	// The time when the Cloud Firewall instance was enabled.
	//
	// > The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1692504764000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The temporary bandwidth.
	//
	// example:
	//
	// 1000
	TemporaryBandwidth *int64 `json:"TemporaryBandwidth,omitempty" xml:"TemporaryBandwidth,omitempty"`
	// Indicates whether threat intelligence is enabled.
	//
	// example:
	//
	// 1
	ThreatIntelligence *int64 `json:"ThreatIntelligence,omitempty" xml:"ThreatIntelligence,omitempty"`
	// The status of the Cloud Firewall instance. Valid values:
	//
	// - **true**: The instance is valid.
	//
	// - **false**: The instance is invalid.
	//
	// example:
	//
	// true
	UserStatus *bool `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	// The edition of the Cloud Firewall instance. Valid values:
	//
	// - **2**: Premium Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **4**: Ultimate Edition
	//
	// - **10**: Pay-as-you-go
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The purchased traffic processing capability for the VPC firewall.
	//
	// example:
	//
	// 3000
	VpcBandwidth *int64 `json:"VpcBandwidth,omitempty" xml:"VpcBandwidth,omitempty"`
	// The number of purchased VPC firewalls.
	//
	// > This parameter applies only to subscription instances.
	//
	// example:
	//
	// 21
	VpcNumber *int64 `json:"VpcNumber,omitempty" xml:"VpcNumber,omitempty"`
}

func (s DescribeUserBuyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyVersionResponseBody) GetAckClusterConnectorQuota() *int64 {
	return s.AckClusterConnectorQuota
}

func (s *DescribeUserBuyVersionResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeUserBuyVersionResponseBody) GetDefaultBandwidth() *int64 {
	return s.DefaultBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *DescribeUserBuyVersionResponseBody) GetExtensionBandwidth() *int64 {
	return s.ExtensionBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetGeneralInstance() *int64 {
	return s.GeneralInstance
}

func (s *DescribeUserBuyVersionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserBuyVersionResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeUserBuyVersionResponseBody) GetInternetBandwidth() *int64 {
	return s.InternetBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetIpNumber() *int64 {
	return s.IpNumber
}

func (s *DescribeUserBuyVersionResponseBody) GetLogStatus() *bool {
	return s.LogStatus
}

func (s *DescribeUserBuyVersionResponseBody) GetLogStorage() *int64 {
	return s.LogStorage
}

func (s *DescribeUserBuyVersionResponseBody) GetMajorVersion() *int64 {
	return s.MajorVersion
}

func (s *DescribeUserBuyVersionResponseBody) GetMaxOverflow() *int64 {
	return s.MaxOverflow
}

func (s *DescribeUserBuyVersionResponseBody) GetNatBandwidth() *int64 {
	return s.NatBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetPrivateDnsConnectorQuota() *int64 {
	return s.PrivateDnsConnectorQuota
}

func (s *DescribeUserBuyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBuyVersionResponseBody) GetSdl() *int64 {
	return s.Sdl
}

func (s *DescribeUserBuyVersionResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUserBuyVersionResponseBody) GetTemporaryBandwidth() *int64 {
	return s.TemporaryBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetThreatIntelligence() *int64 {
	return s.ThreatIntelligence
}

func (s *DescribeUserBuyVersionResponseBody) GetUserStatus() *bool {
	return s.UserStatus
}

func (s *DescribeUserBuyVersionResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeUserBuyVersionResponseBody) GetVpcBandwidth() *int64 {
	return s.VpcBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetVpcNumber() *int64 {
	return s.VpcNumber
}

func (s *DescribeUserBuyVersionResponseBody) SetAckClusterConnectorQuota(v int64) *DescribeUserBuyVersionResponseBody {
	s.AckClusterConnectorQuota = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetAliUid(v int64) *DescribeUserBuyVersionResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetDefaultBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.DefaultBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetExpire(v int64) *DescribeUserBuyVersionResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetExtensionBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.ExtensionBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetGeneralInstance(v int64) *DescribeUserBuyVersionResponseBody {
	s.GeneralInstance = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetInstanceId(v string) *DescribeUserBuyVersionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetInstanceStatus(v string) *DescribeUserBuyVersionResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetInternetBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.InternetBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetIpNumber(v int64) *DescribeUserBuyVersionResponseBody {
	s.IpNumber = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetLogStatus(v bool) *DescribeUserBuyVersionResponseBody {
	s.LogStatus = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetLogStorage(v int64) *DescribeUserBuyVersionResponseBody {
	s.LogStorage = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetMajorVersion(v int64) *DescribeUserBuyVersionResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetMaxOverflow(v int64) *DescribeUserBuyVersionResponseBody {
	s.MaxOverflow = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetNatBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.NatBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetPrivateDnsConnectorQuota(v int64) *DescribeUserBuyVersionResponseBody {
	s.PrivateDnsConnectorQuota = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetRequestId(v string) *DescribeUserBuyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetSdl(v int64) *DescribeUserBuyVersionResponseBody {
	s.Sdl = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetStartTime(v int64) *DescribeUserBuyVersionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetTemporaryBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.TemporaryBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetThreatIntelligence(v int64) *DescribeUserBuyVersionResponseBody {
	s.ThreatIntelligence = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetUserStatus(v bool) *DescribeUserBuyVersionResponseBody {
	s.UserStatus = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetVersion(v int32) *DescribeUserBuyVersionResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetVpcBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.VpcBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetVpcNumber(v int64) *DescribeUserBuyVersionResponseBody {
	s.VpcNumber = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
