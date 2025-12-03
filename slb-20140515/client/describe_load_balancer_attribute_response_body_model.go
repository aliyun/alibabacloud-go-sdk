// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddress() *string
	SetAddressIPVersion(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddressIPVersion() *string
	SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddressType() *string
	SetAutoReleaseTime(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetAutoReleaseTime() *int64
	SetBackendServers(v *DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody
	GetBackendServers() *DescribeLoadBalancerAttributeResponseBodyBackendServers
	SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody
	GetBandwidth() *int32
	SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetCreateTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetCreateTimeStamp() *int64
	SetDeleteProtection(v string) *DescribeLoadBalancerAttributeResponseBody
	GetDeleteProtection() *string
	SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetEndTime() *string
	SetEndTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetEndTimeStamp() *int64
	SetInstanceChargeType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetInternetChargeType() *string
	SetListenerPorts(v *DescribeLoadBalancerAttributeResponseBodyListenerPorts) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPorts() *DescribeLoadBalancerAttributeResponseBodyListenerPorts
	SetListenerPortsAndProtocal(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPortsAndProtocal() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal
	SetListenerPortsAndProtocol(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPortsAndProtocol() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol
	SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerSpec() *string
	SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetMasterZoneId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetMasterZoneId() *string
	SetModificationProtectionReason(v string) *DescribeLoadBalancerAttributeResponseBody
	GetModificationProtectionReason() *string
	SetModificationProtectionStatus(v string) *DescribeLoadBalancerAttributeResponseBody
	GetModificationProtectionStatus() *string
	SetNetworkType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetNetworkType() *string
	SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRegionId() *string
	SetRegionIdAlias(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRegionIdAlias() *string
	SetRenewalCycUnit(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRenewalCycUnit() *string
	SetRenewalDuration(v int32) *DescribeLoadBalancerAttributeResponseBody
	GetRenewalDuration() *int32
	SetRenewalStatus(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRenewalStatus() *string
	SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetResourceGroupId() *string
	SetSlaveZoneId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetSlaveZoneId() *string
	SetTags(v *DescribeLoadBalancerAttributeResponseBodyTags) *DescribeLoadBalancerAttributeResponseBody
	GetTags() *DescribeLoadBalancerAttributeResponseBodyTags
	SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetVpcId() *string
}

type DescribeLoadBalancerAttributeResponseBody struct {
	// The service IP address of the CLB instance.
	//
	// example:
	//
	// 42.XX.XX.6
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The version of the IP address. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The address type of the CLB instance.
	//
	// example:
	//
	// internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The timestamp generated when the CLB instance is released.
	//
	// example:
	//
	// 1513947075000
	AutoReleaseTime *int64 `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The backend servers of the CLB instance.
	BackendServers *DescribeLoadBalancerAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The maximum bandwidth of the Internet-facing CLB instance that is billed on a pay-by-bandwidth basis.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The time when the CLB instance was created. The time is in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2017-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp generated when the CA certificate is uploaded.
	//
	// example:
	//
	// 1504147745000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// Indicates whether deletion protection is enabled for the CLB instance.
	//
	// Valid values: **on*	- and **off**.
	//
	// example:
	//
	// off
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// The time when the CLB instance expires.
	//
	// example:
	//
	// 2022-09-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The timestamp that indicates the expiration time of the CLB instance.
	//
	// example:
	//
	// 32493801600000
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// The metering method of the CLB instance. Valid values:
	//
	// 	- **PayBySpec*	- (default)
	//
	// 	- **PayByCLCU**
	//
	// > This parameter is available only on the China site and takes effect only when **PayType*	- is set to **PayOnDemand**.
	//
	// example:
	//
	// PayBySpec
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Internet-facing CLB instance. Valid values:
	//
	// 	- **paybytraffic**
	//
	// 	- **paybybandwidth**
	//
	// example:
	//
	// paybytraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The frontend port used by the CLB instance.
	ListenerPorts *DescribeLoadBalancerAttributeResponseBodyListenerPorts `json:"ListenerPorts,omitempty" xml:"ListenerPorts,omitempty" type:"Struct"`
	// The ports or protocols of the listeners.
	ListenerPortsAndProtocal *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal `json:"ListenerPortsAndProtocal,omitempty" xml:"ListenerPortsAndProtocal,omitempty" type:"Struct"`
	// The ports or protocols of the listeners.
	ListenerPortsAndProtocol *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol `json:"ListenerPortsAndProtocol,omitempty" xml:"ListenerPortsAndProtocol,omitempty" type:"Struct"`
	// The CLB instance ID.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the CLB instance.
	//
	// example:
	//
	// lb-instance-test
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specification of the CLB instance.
	//
	// >  Pay-as-you-go CLB instances are not subject to specifications. By default, **slb.lcu.elastic*	- is returned.
	//
	// example:
	//
	// slb.s1.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **inactive**: The CLB instance is disabled. CLB instances in the inactive state do not forward traffic.
	//
	// 	- **active**: The CLB instance is running as expected. Newly created CLB instances are in the **active*	- state by default.
	//
	// 	- **locked**: The CLB instance is locked. CLB instances may be locked due to overdue payments or other reasons.
	//
	// example:
	//
	// active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The ID of the primary zone to which the CLB instance belongs.
	//
	// example:
	//
	// cn-hangzhou-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The reason why the configuration read-only mode is enabled. The value is 1 to 80 characters in length. It starts with a letter and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  This parameter is valid only when **ModificationProtectionStatus*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// Managed instance
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	// Indicates whether the configuration read-only mode is enabled. Valid values:
	//
	// 	- **NonProtection**: The configuration read-only mode is disabled. After you disable the configuration read-only mode, the value of **ModificationProtectionReason*	- is cleared.
	//
	// 	- **ConsoleProtection**: The configuration read-only mode is enabled.
	//
	// >  If this parameter is set to **ConsoleProtection**, you cannot modify instance configurations in the CLB console. However, you can modify instance configurations by calling API operations.
	//
	// example:
	//
	// ConsoleProtection
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	// The network type of the CLB instance.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- Only **PayOnDemand*	- may be returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the CLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The alias of the region to which the CLB instance belongs.
	//
	// example:
	//
	// hangzhou
	RegionIdAlias *string `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	// The auto-renewal cycle. Valid values: **Year*	- and **Month**. Default value: Month.
	//
	// >  This parameter is valid only if you create a subscription CLB instance on the Alibaba Cloud China site. In this case, **PayType*	- must be set to **PrePay*	- and **RenewalStatus*	- must be set to **AutoRenewal**.
	//
	// example:
	//
	// Month
	RenewalCycUnit *string `json:"RenewalCycUnit,omitempty" xml:"RenewalCycUnit,omitempty"`
	// The auto-renewal duration. This parameter is valid only if **RenewalStatus*	- is set to **AutoRenewal**.
	//
	// 	- Valid values when **PeriodUnit*	- is set to **Year**: **1**~**5**.
	//
	// 	- Valid values when **PeriodUnit*	- is set to **Month**: **1**~ **9**.
	//
	// > This parameter is valid only when you create a subscription CLB instance on the Alibaba Cloud China site. In this case, the **PayType*	- parameter must be set to **PrePay**.
	//
	// example:
	//
	// 1
	RenewalDuration *int32 `json:"RenewalDuration,omitempty" xml:"RenewalDuration,omitempty"`
	// Indicates whether auto-renewal is enabled. Valid values:
	//
	// 	- **AutoRenewal**: Auto-renewal is enabled.
	//
	// 	- **Normal**: Auto-renewal is disabled. You must manually renew the CLB instance.
	//
	// 	- **NotRenewal**: The CLB instance will not be renewed upon expiration. If this value is returned, the system does not send notifications until three days before the expiration date.
	//
	//     **
	//
	//     **Note*	- This parameter is valid only when you create a subscription CLB instance on the Alibaba Cloud China site. In this case, **PayType*	- must be set to **PrePay**.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the secondary zone to which the CLB instance belongs.
	//
	// example:
	//
	// cn-hangzhou-d
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the internal-facing CLB instance belongs.
	//
	// example:
	//
	// vsw-255ecrwq5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) where the internal-facing CLB instance is deployed.
	//
	// example:
	//
	// vpc-25dvzy9f8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAutoReleaseTime() *int64 {
	return s.AutoReleaseTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBackendServers() *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	return s.BackendServers
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetDeleteProtection() *string {
	return s.DeleteProtection
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPorts() *DescribeLoadBalancerAttributeResponseBodyListenerPorts {
	return s.ListenerPorts
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPortsAndProtocal() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal {
	return s.ListenerPortsAndProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPortsAndProtocol() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol {
	return s.ListenerPortsAndProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetModificationProtectionReason() *string {
	return s.ModificationProtectionReason
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetModificationProtectionStatus() *string {
	return s.ModificationProtectionStatus
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRegionIdAlias() *string {
	return s.RegionIdAlias
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRenewalCycUnit() *string {
	return s.RenewalCycUnit
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRenewalDuration() *int32 {
	return s.RenewalDuration
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetTags() *DescribeLoadBalancerAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressIPVersion(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAutoReleaseTime(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBackendServers(v *DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetDeleteProtection(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.DeleteProtection = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetInstanceChargeType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetInternetChargeType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPorts(v *DescribeLoadBalancerAttributeResponseBodyListenerPorts) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPorts = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocal(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocal = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocol(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocol = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetMasterZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetModificationProtectionReason(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ModificationProtectionReason = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetModificationProtectionStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetNetworkType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionIdAlias(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalCycUnit(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalCycUnit = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalDuration(v int32) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalDuration = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetSlaveZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetTags(v *DescribeLoadBalancerAttributeResponseBodyTags) *DescribeLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVpcId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPorts != nil {
		if err := s.ListenerPorts.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPortsAndProtocal != nil {
		if err := s.ListenerPortsAndProtocal.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPortsAndProtocol != nil {
		if err := s.ListenerPortsAndProtocol.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyBackendServers struct {
	BackendServer []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetBackendServer() []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetBackendServer(v []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) Validate() error {
	if s.BackendServer != nil {
		for _, item := range s.BackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer struct {
	// The description of the backend server.
	//
	// > This parameter is not returned if Description is not set.
	//
	// example:
	//
	// backend server description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The backend server ID.
	//
	// example:
	//
	// i-2zej4lxhjoq1icu*****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The ID of the elastic network interface (ENI) or elastic container instance.
	//
	// example:
	//
	// 192.XX.XX.11
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server.
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 90
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetServerIp() *string {
	return s.ServerIp
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetServerIp(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.ServerIp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetType(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPorts struct {
	ListenerPort []*int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) GetListenerPort() []*int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) SetListenerPort(v []*int32) *DescribeLoadBalancerAttributeResponseBodyListenerPorts {
	s.ListenerPort = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal struct {
	ListenerPortAndProtocal []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal `json:"ListenerPortAndProtocal,omitempty" xml:"ListenerPortAndProtocal,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) GetListenerPortAndProtocal() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	return s.ListenerPortAndProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) SetListenerPortAndProtocal(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal {
	s.ListenerPortAndProtocal = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) Validate() error {
	if s.ListenerPortAndProtocal != nil {
		for _, item := range s.ListenerPortAndProtocal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal struct {
	// The frontend port that is used by the CLB instance.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the CLB instance.
	//
	// example:
	//
	// http
	ListenerProtocal *string `json:"ListenerProtocal,omitempty" xml:"ListenerProtocal,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GetListenerProtocal() *string {
	return s.ListenerProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerProtocal(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerProtocal = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol struct {
	ListenerPortAndProtocol []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol `json:"ListenerPortAndProtocol,omitempty" xml:"ListenerPortAndProtocol,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) GetListenerPortAndProtocol() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	return s.ListenerPortAndProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) SetListenerPortAndProtocol(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol {
	s.ListenerPortAndProtocol = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) Validate() error {
	if s.ListenerPortAndProtocol != nil {
		for _, item := range s.ListenerPortAndProtocol {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol struct {
	// Indicates whether the listener is enabled.
	//
	// example:
	//
	// Listener Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination listening port to which requests are forwarded. The port must be open and use HTTPS.
	//
	// example:
	//
	// 443
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether the listener is enabled.
	//
	// example:
	//
	// on
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The frontend port that is used by the CLB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the CLB instance.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetDescription(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetForwardPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerForward(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerProtocol(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyTags struct {
	Tag []*DescribeLoadBalancerAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyTags) GetTag() []*DescribeLoadBalancerAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancerAttributeResponseBodyTags) SetTag(v []*DescribeLoadBalancerAttributeResponseBodyTagsTag) *DescribeLoadBalancerAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyTagsTag struct {
	// The tag key. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string.
	//
	// The tag key can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeLoadBalancerAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeLoadBalancerAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
