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
	// example:
	//
	// 5
	AckClusterConnectorQuota *int64 `json:"AckClusterConnectorQuota,omitempty" xml:"AckClusterConnectorQuota,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase Cloud Firewall.
	//
	// example:
	//
	// 119898001566xxxx
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 1000
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" xml:"DefaultBandwidth,omitempty"`
	// The time when Cloud Firewall expires.
	//
	// >  The value is a timestamp in milliseconds.
	//
	// >  If you use Cloud Firewall that uses the pay-as-you-go billing method, ignore this parameter.
	//
	// example:
	//
	// 1726934400000
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// 1000
	ExtensionBandwidth *int64 `json:"ExtensionBandwidth,omitempty" xml:"ExtensionBandwidth,omitempty"`
	// example:
	//
	// 10
	GeneralInstance *int64 `json:"GeneralInstance,omitempty" xml:"GeneralInstance,omitempty"`
	// The instance ID of Cloud Firewall.
	//
	// >  If you use a trial of Cloud Firewall, ignore this parameter.
	//
	// example:
	//
	// vipcloudfw-cn-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of Cloud Firewall. Valid values:
	//
	// 	- **normal**: Cloud Firewall is running as expected.
	//
	// 	- **init**: Cloud Firewall is being initialized.
	//
	// 	- **deleting**: Cloud Firewall is being deleted.
	//
	// 	- **abnormal**: An exception occurs in Cloud Firewall.
	//
	// 	- **free**: Cloud Firewall is invalid.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The peak Internet traffic that can be protected.
	//
	// example:
	//
	// 3000
	InternetBandwidth *int64 `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty"`
	// The number of public IP addresses that can be protected.
	//
	// >  This parameter takes effect only for Cloud Firewall that uses the subscription billing method.
	//
	// example:
	//
	// 63
	IpNumber *int64 `json:"IpNumber,omitempty" xml:"IpNumber,omitempty"`
	// Indicates whether log delivery is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	LogStatus *bool `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// The log storage capacity.
	//
	// >  This parameter takes effect only for Cloud Firewall that uses the subscription billing method.
	//
	// example:
	//
	// 3000
	LogStorage *int64 `json:"LogStorage,omitempty" xml:"LogStorage,omitempty"`
	// example:
	//
	// 1
	MajorVersion *int64 `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The status of the burstable protected traffic feature. Valid values:
	//
	// 	- **1000000**: enabled.
	//
	// 	- **0**: disabled.
	//
	// >  This parameter takes effect only for Cloud Firewall that uses the subscription billing method.
	//
	// example:
	//
	// 0
	MaxOverflow *int64 `json:"MaxOverflow,omitempty" xml:"MaxOverflow,omitempty"`
	// The peak traffic of NAT private network that can be protected.
	//
	// example:
	//
	// 3000
	NatBandwidth *int64 `json:"NatBandwidth,omitempty" xml:"NatBandwidth,omitempty"`
	// example:
	//
	// 5
	PrivateDnsConnectorQuota *int64 `json:"PrivateDnsConnectorQuota,omitempty" xml:"PrivateDnsConnectorQuota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F71B03EE-xxxxx-91D79CC6AA1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Sdl *int64 `json:"Sdl,omitempty" xml:"Sdl,omitempty"`
	// The time when Cloud Firewall was activated.
	//
	// >  The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1692504764000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1000
	TemporaryBandwidth *int64 `json:"TemporaryBandwidth,omitempty" xml:"TemporaryBandwidth,omitempty"`
	// example:
	//
	// 1
	ThreatIntelligence *int64 `json:"ThreatIntelligence,omitempty" xml:"ThreatIntelligence,omitempty"`
	// Indicates whether Cloud Firewall is valid. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UserStatus *bool `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	// The edition of Cloud Firewall. Valid values:
	//
	// 	- **2**: Premium Edition.
	//
	// 	- **3**: Enterprise Edition.
	//
	// 	- **4**: Ultimate Edition.
	//
	// 	- **10**: Cloud Firewall that uses the pay-as-you-go billing method.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The peak cross-VPC traffic that can be protected.
	//
	// example:
	//
	// 3000
	VpcBandwidth *int64 `json:"VpcBandwidth,omitempty" xml:"VpcBandwidth,omitempty"`
	// The number of virtual private clouds (VPCs) that can be protected.
	//
	// >  This parameter takes effect only for Cloud Firewall that uses the subscription billing method.
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
