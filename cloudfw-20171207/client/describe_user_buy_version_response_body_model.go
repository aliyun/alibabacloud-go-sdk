// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *DescribeUserBuyVersionResponseBody
	GetAliUid() *int64
	SetExpire(v int64) *DescribeUserBuyVersionResponseBody
	GetExpire() *int64
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
	SetMaxOverflow(v int64) *DescribeUserBuyVersionResponseBody
	GetMaxOverflow() *int64
	SetNatBandwidth(v int64) *DescribeUserBuyVersionResponseBody
	GetNatBandwidth() *int64
	SetRequestId(v string) *DescribeUserBuyVersionResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *DescribeUserBuyVersionResponseBody
	GetStartTime() *int64
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
	// The ID of the Alibaba Cloud account that is used to purchase Cloud Firewall.
	//
	// example:
	//
	// 119898001566xxxx
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// F71B03EE-xxxxx-91D79CC6AA1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when Cloud Firewall was activated.
	//
	// >  The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1692504764000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeUserBuyVersionResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeUserBuyVersionResponseBody) GetExpire() *int64 {
	return s.Expire
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

func (s *DescribeUserBuyVersionResponseBody) GetMaxOverflow() *int64 {
	return s.MaxOverflow
}

func (s *DescribeUserBuyVersionResponseBody) GetNatBandwidth() *int64 {
	return s.NatBandwidth
}

func (s *DescribeUserBuyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBuyVersionResponseBody) GetStartTime() *int64 {
	return s.StartTime
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

func (s *DescribeUserBuyVersionResponseBody) SetAliUid(v int64) *DescribeUserBuyVersionResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetExpire(v int64) *DescribeUserBuyVersionResponseBody {
	s.Expire = &v
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

func (s *DescribeUserBuyVersionResponseBody) SetMaxOverflow(v int64) *DescribeUserBuyVersionResponseBody {
	s.MaxOverflow = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetNatBandwidth(v int64) *DescribeUserBuyVersionResponseBody {
	s.NatBandwidth = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetRequestId(v string) *DescribeUserBuyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBuyVersionResponseBody) SetStartTime(v int64) *DescribeUserBuyVersionResponseBody {
	s.StartTime = &v
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
