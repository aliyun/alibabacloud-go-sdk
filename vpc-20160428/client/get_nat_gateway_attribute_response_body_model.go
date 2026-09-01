// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v *GetNatGatewayAttributeResponseBodyAccessMode) *GetNatGatewayAttributeResponseBody
	GetAccessMode() *GetNatGatewayAttributeResponseBodyAccessMode
	SetAvailabilityMode(v string) *GetNatGatewayAttributeResponseBody
	GetAvailabilityMode() *string
	SetBillingConfig(v *GetNatGatewayAttributeResponseBodyBillingConfig) *GetNatGatewayAttributeResponseBody
	GetBillingConfig() *GetNatGatewayAttributeResponseBodyBillingConfig
	SetBusinessStatus(v string) *GetNatGatewayAttributeResponseBody
	GetBusinessStatus() *string
	SetCreationTime(v string) *GetNatGatewayAttributeResponseBody
	GetCreationTime() *string
	SetDeletionProtectionInfo(v *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) *GetNatGatewayAttributeResponseBody
	GetDeletionProtectionInfo() *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo
	SetDescription(v string) *GetNatGatewayAttributeResponseBody
	GetDescription() *string
	SetEcsMetricEnabled(v bool) *GetNatGatewayAttributeResponseBody
	GetEcsMetricEnabled() *bool
	SetEnableSessionLog(v bool) *GetNatGatewayAttributeResponseBody
	GetEnableSessionLog() *bool
	SetExpiredTime(v string) *GetNatGatewayAttributeResponseBody
	GetExpiredTime() *string
	SetForwardTable(v *GetNatGatewayAttributeResponseBodyForwardTable) *GetNatGatewayAttributeResponseBody
	GetForwardTable() *GetNatGatewayAttributeResponseBodyForwardTable
	SetFullNatTable(v *GetNatGatewayAttributeResponseBodyFullNatTable) *GetNatGatewayAttributeResponseBody
	GetFullNatTable() *GetNatGatewayAttributeResponseBodyFullNatTable
	SetIpList(v []*GetNatGatewayAttributeResponseBodyIpList) *GetNatGatewayAttributeResponseBody
	GetIpList() []*GetNatGatewayAttributeResponseBodyIpList
	SetLogDelivery(v *GetNatGatewayAttributeResponseBodyLogDelivery) *GetNatGatewayAttributeResponseBody
	GetLogDelivery() *GetNatGatewayAttributeResponseBodyLogDelivery
	SetName(v string) *GetNatGatewayAttributeResponseBody
	GetName() *string
	SetNatGatewayId(v string) *GetNatGatewayAttributeResponseBody
	GetNatGatewayId() *string
	SetNatType(v string) *GetNatGatewayAttributeResponseBody
	GetNatType() *string
	SetNetworkType(v string) *GetNatGatewayAttributeResponseBody
	GetNetworkType() *string
	SetPrivateInfo(v *GetNatGatewayAttributeResponseBodyPrivateInfo) *GetNatGatewayAttributeResponseBody
	GetPrivateInfo() *GetNatGatewayAttributeResponseBodyPrivateInfo
	SetPrivateLinkEnabled(v bool) *GetNatGatewayAttributeResponseBody
	GetPrivateLinkEnabled() *bool
	SetPrivateLinkMode(v string) *GetNatGatewayAttributeResponseBody
	GetPrivateLinkMode() *string
	SetRegionId(v string) *GetNatGatewayAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetNatGatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetNatGatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetSnatTable(v *GetNatGatewayAttributeResponseBodySnatTable) *GetNatGatewayAttributeResponseBody
	GetSnatTable() *GetNatGatewayAttributeResponseBodySnatTable
	SetStatus(v string) *GetNatGatewayAttributeResponseBody
	GetStatus() *string
	SetVpcId(v string) *GetNatGatewayAttributeResponseBody
	GetVpcId() *string
}

type GetNatGatewayAttributeResponseBody struct {
	// The access mode for reverse access to the VPC NAT gateway.
	AccessMode       *GetNatGatewayAttributeResponseBodyAccessMode `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" type:"Struct"`
	AvailabilityMode *string                                       `json:"AvailabilityMode,omitempty" xml:"AvailabilityMode,omitempty"`
	// The billing configuration information.
	BillingConfig *GetNatGatewayAttributeResponseBodyBillingConfig `json:"BillingConfig,omitempty" xml:"BillingConfig,omitempty" type:"Struct"`
	// The business status of the NAT gateway. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: Locked due to overdue payment.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the NAT gateway was created. The time is displayed in the format of YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2021-12-08T12:20:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The deletion protection information.
	DeletionProtectionInfo *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo `json:"DeletionProtectionInfo,omitempty" xml:"DeletionProtectionInfo,omitempty" type:"Struct"`
	// The description of the NAT gateway instance.
	//
	// example:
	//
	// NAT
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the gateway traffic monitoring feature is enabled. Valid values:
	//
	// - **true**: The gateway traffic monitoring feature is enabled.
	//
	// - **false**: The gateway traffic monitoring feature is disabled.
	//
	// example:
	//
	// true
	EcsMetricEnabled *bool `json:"EcsMetricEnabled,omitempty" xml:"EcsMetricEnabled,omitempty"`
	// Indicates whether session logging is enabled. Valid values:
	//
	// - **true**: Session logging is enabled.
	//
	// - **false**: Session logging is disabled.
	//
	// example:
	//
	// true
	EnableSessionLog *bool `json:"EnableSessionLog,omitempty" xml:"EnableSessionLog,omitempty"`
	// The expiration time of the NAT gateway instance.
	//
	// example:
	//
	// 2021-12-26T12:20:20Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The DNAT table information.
	ForwardTable *GetNatGatewayAttributeResponseBodyForwardTable `json:"ForwardTable,omitempty" xml:"ForwardTable,omitempty" type:"Struct"`
	// The FULLNAT table information.
	FullNatTable *GetNatGatewayAttributeResponseBodyFullNatTable `json:"FullNatTable,omitempty" xml:"FullNatTable,omitempty" type:"Struct"`
	// The list of elastic IP addresses (EIPs) associated with the Internet NAT gateway.
	IpList []*GetNatGatewayAttributeResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The session log configuration information.
	LogDelivery *GetNatGatewayAttributeResponseBodyLogDelivery `json:"LogDelivery,omitempty" xml:"LogDelivery,omitempty" type:"Struct"`
	// The name of the NAT gateway instance.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway instance.
	//
	// example:
	//
	// ngw-bp1047e2d4z7kf2ki****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The type of the Internet NAT gateway. The value is **Enhanced**, which indicates an enhanced NAT gateway.
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The type of the NAT gateway. Valid values:
	//
	// - **internet**: Internet NAT gateway.
	//
	// - **intranet**: VPC NAT gateway.
	//
	// example:
	//
	// internet
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The private network information of the NAT gateway instance.
	PrivateInfo *GetNatGatewayAttributeResponseBodyPrivateInfo `json:"PrivateInfo,omitempty" xml:"PrivateInfo,omitempty" type:"Struct"`
	// Indicates whether private connectivity is supported. Valid values:
	//
	// - **true**: Private connectivity is supported.
	//
	// - **false**: Private connectivity is not supported.
	//
	// example:
	//
	// true
	PrivateLinkEnabled *bool `json:"PrivateLinkEnabled,omitempty" xml:"PrivateLinkEnabled,omitempty"`
	// The conversion mode of the private connectivity service. Valid values:
	//
	// - **FullNat**: FULLNAT mode.
	//
	// - **Geneve**: Geneve mode.
	//
	// example:
	//
	// FullNat
	PrivateLinkMode *string `json:"PrivateLinkMode,omitempty" xml:"PrivateLinkMode,omitempty"`
	// The region ID of the NAT gateway instance.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The SNAT table information.
	SnatTable *GetNatGatewayAttributeResponseBodySnatTable `json:"SnatTable,omitempty" xml:"SnatTable,omitempty" type:"Struct"`
	// The status of the NAT gateway. Valid values:
	//
	// - **Creating**: The NAT gateway is being created. Creating a NAT gateway is an asynchronous operation. The NAT gateway remains in the **Creating*	- state until the operation is complete.
	//
	// - **Available**: The NAT gateway is available. This is a stable state after the NAT gateway is created.
	//
	// - **Modifying**: The NAT gateway is being modified. Modifying a NAT gateway is an asynchronous operation. The NAT gateway remains in the **Modifying*	- state until the operation is complete.
	//
	// - **Deleting**: The NAT gateway is being deleted. Deleting a NAT gateway is an asynchronous operation. The NAT gateway remains in the **Deleting*	- state until the operation is complete.
	//
	// - **Converting**: The NAT gateway is being converted. Converting a standard NAT gateway to an enhanced NAT gateway is an asynchronous operation. The NAT gateway remains in the **Converting*	- state until the operation is complete.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC to which the NAT gateway instance belongs.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetNatGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBody) GetAccessMode() *GetNatGatewayAttributeResponseBodyAccessMode {
	return s.AccessMode
}

func (s *GetNatGatewayAttributeResponseBody) GetAvailabilityMode() *string {
	return s.AvailabilityMode
}

func (s *GetNatGatewayAttributeResponseBody) GetBillingConfig() *GetNatGatewayAttributeResponseBodyBillingConfig {
	return s.BillingConfig
}

func (s *GetNatGatewayAttributeResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *GetNatGatewayAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetNatGatewayAttributeResponseBody) GetDeletionProtectionInfo() *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo {
	return s.DeletionProtectionInfo
}

func (s *GetNatGatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetNatGatewayAttributeResponseBody) GetEcsMetricEnabled() *bool {
	return s.EcsMetricEnabled
}

func (s *GetNatGatewayAttributeResponseBody) GetEnableSessionLog() *bool {
	return s.EnableSessionLog
}

func (s *GetNatGatewayAttributeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetNatGatewayAttributeResponseBody) GetForwardTable() *GetNatGatewayAttributeResponseBodyForwardTable {
	return s.ForwardTable
}

func (s *GetNatGatewayAttributeResponseBody) GetFullNatTable() *GetNatGatewayAttributeResponseBodyFullNatTable {
	return s.FullNatTable
}

func (s *GetNatGatewayAttributeResponseBody) GetIpList() []*GetNatGatewayAttributeResponseBodyIpList {
	return s.IpList
}

func (s *GetNatGatewayAttributeResponseBody) GetLogDelivery() *GetNatGatewayAttributeResponseBodyLogDelivery {
	return s.LogDelivery
}

func (s *GetNatGatewayAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *GetNatGatewayAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatGatewayAttributeResponseBody) GetNatType() *string {
	return s.NatType
}

func (s *GetNatGatewayAttributeResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetNatGatewayAttributeResponseBody) GetPrivateInfo() *GetNatGatewayAttributeResponseBodyPrivateInfo {
	return s.PrivateInfo
}

func (s *GetNatGatewayAttributeResponseBody) GetPrivateLinkEnabled() *bool {
	return s.PrivateLinkEnabled
}

func (s *GetNatGatewayAttributeResponseBody) GetPrivateLinkMode() *string {
	return s.PrivateLinkMode
}

func (s *GetNatGatewayAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNatGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNatGatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetNatGatewayAttributeResponseBody) GetSnatTable() *GetNatGatewayAttributeResponseBodySnatTable {
	return s.SnatTable
}

func (s *GetNatGatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetNatGatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNatGatewayAttributeResponseBody) SetAccessMode(v *GetNatGatewayAttributeResponseBodyAccessMode) *GetNatGatewayAttributeResponseBody {
	s.AccessMode = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetAvailabilityMode(v string) *GetNatGatewayAttributeResponseBody {
	s.AvailabilityMode = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetBillingConfig(v *GetNatGatewayAttributeResponseBodyBillingConfig) *GetNatGatewayAttributeResponseBody {
	s.BillingConfig = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetBusinessStatus(v string) *GetNatGatewayAttributeResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetCreationTime(v string) *GetNatGatewayAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetDeletionProtectionInfo(v *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) *GetNatGatewayAttributeResponseBody {
	s.DeletionProtectionInfo = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetDescription(v string) *GetNatGatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetEcsMetricEnabled(v bool) *GetNatGatewayAttributeResponseBody {
	s.EcsMetricEnabled = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetEnableSessionLog(v bool) *GetNatGatewayAttributeResponseBody {
	s.EnableSessionLog = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetExpiredTime(v string) *GetNatGatewayAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetForwardTable(v *GetNatGatewayAttributeResponseBodyForwardTable) *GetNatGatewayAttributeResponseBody {
	s.ForwardTable = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetFullNatTable(v *GetNatGatewayAttributeResponseBodyFullNatTable) *GetNatGatewayAttributeResponseBody {
	s.FullNatTable = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetIpList(v []*GetNatGatewayAttributeResponseBodyIpList) *GetNatGatewayAttributeResponseBody {
	s.IpList = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetLogDelivery(v *GetNatGatewayAttributeResponseBodyLogDelivery) *GetNatGatewayAttributeResponseBody {
	s.LogDelivery = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetName(v string) *GetNatGatewayAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetNatGatewayId(v string) *GetNatGatewayAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetNatType(v string) *GetNatGatewayAttributeResponseBody {
	s.NatType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetNetworkType(v string) *GetNatGatewayAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetPrivateInfo(v *GetNatGatewayAttributeResponseBodyPrivateInfo) *GetNatGatewayAttributeResponseBody {
	s.PrivateInfo = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetPrivateLinkEnabled(v bool) *GetNatGatewayAttributeResponseBody {
	s.PrivateLinkEnabled = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetPrivateLinkMode(v string) *GetNatGatewayAttributeResponseBody {
	s.PrivateLinkMode = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetRegionId(v string) *GetNatGatewayAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetRequestId(v string) *GetNatGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetResourceGroupId(v string) *GetNatGatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetSnatTable(v *GetNatGatewayAttributeResponseBodySnatTable) *GetNatGatewayAttributeResponseBody {
	s.SnatTable = v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetStatus(v string) *GetNatGatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) SetVpcId(v string) *GetNatGatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBody) Validate() error {
	if s.AccessMode != nil {
		if err := s.AccessMode.Validate(); err != nil {
			return err
		}
	}
	if s.BillingConfig != nil {
		if err := s.BillingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeletionProtectionInfo != nil {
		if err := s.DeletionProtectionInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ForwardTable != nil {
		if err := s.ForwardTable.Validate(); err != nil {
			return err
		}
	}
	if s.FullNatTable != nil {
		if err := s.FullNatTable.Validate(); err != nil {
			return err
		}
	}
	if s.IpList != nil {
		for _, item := range s.IpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogDelivery != nil {
		if err := s.LogDelivery.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateInfo != nil {
		if err := s.PrivateInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SnatTable != nil {
		if err := s.SnatTable.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNatGatewayAttributeResponseBodyAccessMode struct {
	// The access mode value. Valid values:
	//
	// - **route**: Route mode.
	//
	// - **tunnel**: Tunnel mode.
	//
	// example:
	//
	// route
	ModeValue *string `json:"ModeValue,omitempty" xml:"ModeValue,omitempty"`
	// The tunnel mode type. Valid values:
	//
	// - **geneve**: Geneve type.
	//
	// example:
	//
	// geneve
	TunnelType *string `json:"TunnelType,omitempty" xml:"TunnelType,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyAccessMode) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyAccessMode) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyAccessMode) GetModeValue() *string {
	return s.ModeValue
}

func (s *GetNatGatewayAttributeResponseBodyAccessMode) GetTunnelType() *string {
	return s.TunnelType
}

func (s *GetNatGatewayAttributeResponseBodyAccessMode) SetModeValue(v string) *GetNatGatewayAttributeResponseBodyAccessMode {
	s.ModeValue = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyAccessMode) SetTunnelType(v string) *GetNatGatewayAttributeResponseBodyAccessMode {
	s.TunnelType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyAccessMode) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyBillingConfig struct {
	// Indicates whether automatic payment is enabled. When the value of **InstanceChargeType*	- is **PrePaid**, the following values are returned:
	//
	// - **false**: Automatic payment is disabled. After an order is generated, go to the Order Center to complete the payment.
	//
	// - **true**: Automatic payment is enabled. The order is automatically paid.
	//
	// When the value of **InstanceChargeType*	- is **PostPaid**, an empty value is returned.
	//
	// example:
	//
	// false
	AutoPay *string `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// <props="china">The billing method of the NAT gateway instance. Valid values:
	//
	// <props="china">- **PostPaid**: pay-as-you-go.
	//
	// <props="china">- **PrePaid**: subscription.
	//
	// <props="intl">The billing method of the NAT gateway instance. Valid values: **PostPaid*	- (pay-as-you-go).
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The billing type of the NAT gateway instance. Valid values:
	//
	// - **PayBySpec**: Billed by defined specifications.
	//
	// - **PayByLcu**: Billed by usage.
	//
	// example:
	//
	// PayBySpec
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The specification of the Internet NAT gateway instance. When **InternetChargeType*	- is **PayBySpec**, the following values are returned:
	//
	// - **Small**: small.
	//
	// - **Middle**: medium.
	//
	// - **Large**: large.
	//
	// When **InternetChargeType*	- is **PayByLcu**, an empty value is returned.
	//
	// example:
	//
	// Small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyBillingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyBillingConfig) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) GetAutoPay() *string {
	return s.AutoPay
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) GetSpec() *string {
	return s.Spec
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) SetAutoPay(v string) *GetNatGatewayAttributeResponseBodyBillingConfig {
	s.AutoPay = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) SetInstanceChargeType(v string) *GetNatGatewayAttributeResponseBodyBillingConfig {
	s.InstanceChargeType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) SetInternetChargeType(v string) *GetNatGatewayAttributeResponseBodyBillingConfig {
	s.InternetChargeType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) SetSpec(v string) *GetNatGatewayAttributeResponseBodyBillingConfig {
	s.Spec = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyBillingConfig) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyDeletionProtectionInfo struct {
	// Indicates whether deletion protection is enabled.
	//
	// - **true**: Deletion protection is enabled.
	//
	// - **false**: Deletion protection is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) SetEnabled(v bool) *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo {
	s.Enabled = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyDeletionProtectionInfo) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyForwardTable struct {
	// The number of DNAT entries.
	//
	// example:
	//
	// 1
	ForwardEntryCount *int32 `json:"ForwardEntryCount,omitempty" xml:"ForwardEntryCount,omitempty"`
	// The ID of the DNAT table.
	//
	// example:
	//
	// ftb-uf6gj3mhsg94qsqst****
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyForwardTable) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyForwardTable) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyForwardTable) GetForwardEntryCount() *int32 {
	return s.ForwardEntryCount
}

func (s *GetNatGatewayAttributeResponseBodyForwardTable) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *GetNatGatewayAttributeResponseBodyForwardTable) SetForwardEntryCount(v int32) *GetNatGatewayAttributeResponseBodyForwardTable {
	s.ForwardEntryCount = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyForwardTable) SetForwardTableId(v string) *GetNatGatewayAttributeResponseBodyForwardTable {
	s.ForwardTableId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyForwardTable) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyFullNatTable struct {
	// The number of FULLNAT entries.
	//
	// example:
	//
	// 1
	FullNatEntryCount *int64 `json:"FullNatEntryCount,omitempty" xml:"FullNatEntryCount,omitempty"`
	// The ID of the FULLNAT table.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyFullNatTable) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyFullNatTable) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyFullNatTable) GetFullNatEntryCount() *int64 {
	return s.FullNatEntryCount
}

func (s *GetNatGatewayAttributeResponseBodyFullNatTable) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *GetNatGatewayAttributeResponseBodyFullNatTable) SetFullNatEntryCount(v int64) *GetNatGatewayAttributeResponseBodyFullNatTable {
	s.FullNatEntryCount = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyFullNatTable) SetFullNatTableId(v string) *GetNatGatewayAttributeResponseBodyFullNatTable {
	s.FullNatTableId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyFullNatTable) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyIpList struct {
	// The instance ID of the EIP.
	//
	// example:
	//
	// eip-bp13e9i2qst4g6jzi****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The IP address of the EIP.
	//
	// example:
	//
	// 116.33.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The association status of the EIP bound to the Internet NAT gateway instance.
	//
	// - **idle**: The EIP is not associated with a SNAT entry or DNAT entry.
	//
	// - **UsedBySnatTable**: The EIP is associated with a SNAT entry.
	//
	// - **UsedByForwardTable**: The EIP is associated with a DNAT entry.
	//
	// example:
	//
	// idle
	UsingStatus *string `json:"UsingStatus,omitempty" xml:"UsingStatus,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyIpList) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyIpList) GetAllocationId() *string {
	return s.AllocationId
}

func (s *GetNatGatewayAttributeResponseBodyIpList) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetNatGatewayAttributeResponseBodyIpList) GetUsingStatus() *string {
	return s.UsingStatus
}

func (s *GetNatGatewayAttributeResponseBodyIpList) SetAllocationId(v string) *GetNatGatewayAttributeResponseBodyIpList {
	s.AllocationId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyIpList) SetIpAddress(v string) *GetNatGatewayAttributeResponseBodyIpList {
	s.IpAddress = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyIpList) SetUsingStatus(v string) *GetNatGatewayAttributeResponseBodyIpList {
	s.UsingStatus = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyIpList) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyLogDelivery struct {
	// The error message for session log writing failure.
	//
	// example:
	//
	// LogStoreNotExist: logstore session_log_test does not exist
	DeliverLogsErrorMessage *string `json:"DeliverLogsErrorMessage,omitempty" xml:"DeliverLogsErrorMessage,omitempty"`
	// The status of session log writing. Valid values:
	//
	// - **Succsess**: Succeeded.
	//
	// - **Failure**: Failed.
	//
	// example:
	//
	// Failure
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// The destination type for session log delivery. Valid values:
	//
	// **sls**: Alibaba Cloud Simple Log Service (SLS).
	//
	// example:
	//
	// sls
	LogDeliveryType *string `json:"LogDeliveryType,omitempty" xml:"LogDeliveryType,omitempty"`
	// The destination address to which session logs are written.
	//
	// example:
	//
	// acs:log:cn-hangzhou:0000:project/nat_session_log_project/logstore/session_log_test
	LogDestination *string `json:"LogDestination,omitempty" xml:"LogDestination,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyLogDelivery) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyLogDelivery) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) GetDeliverLogsErrorMessage() *string {
	return s.DeliverLogsErrorMessage
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) GetLogDeliveryType() *string {
	return s.LogDeliveryType
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) GetLogDestination() *string {
	return s.LogDestination
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) SetDeliverLogsErrorMessage(v string) *GetNatGatewayAttributeResponseBodyLogDelivery {
	s.DeliverLogsErrorMessage = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) SetDeliveryStatus(v string) *GetNatGatewayAttributeResponseBodyLogDelivery {
	s.DeliveryStatus = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) SetLogDeliveryType(v string) *GetNatGatewayAttributeResponseBodyLogDelivery {
	s.LogDeliveryType = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) SetLogDestination(v string) *GetNatGatewayAttributeResponseBodyLogDelivery {
	s.LogDestination = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyLogDelivery) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodyPrivateInfo struct {
	// The instance ID of the elastic network interfaces (ENIs) network interface controller (NIC).
	//
	// example:
	//
	// eni-bp1cmgtoaka8vfyg****
	EniInstanceId *string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	// The zone to which the NAT gateway instance belongs.
	//
	// example:
	//
	// cn-qingdao-b
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The maximum bandwidth value. Unit: Mbit/s.
	//
	// example:
	//
	// 5120
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the vSwitch to which the NAT gateway instance belongs.
	//
	// example:
	//
	// vsw-bp1s2laxhdf9ayjbo***
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodyPrivateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodyPrivateInfo) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) GetEniInstanceId() *string {
	return s.EniInstanceId
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) GetIzNo() *string {
	return s.IzNo
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) GetMaxBandwidth() *int32 {
	return s.MaxBandwidth
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) SetEniInstanceId(v string) *GetNatGatewayAttributeResponseBodyPrivateInfo {
	s.EniInstanceId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) SetIzNo(v string) *GetNatGatewayAttributeResponseBodyPrivateInfo {
	s.IzNo = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) SetMaxBandwidth(v int32) *GetNatGatewayAttributeResponseBodyPrivateInfo {
	s.MaxBandwidth = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) SetPrivateIpAddress(v string) *GetNatGatewayAttributeResponseBodyPrivateInfo {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) SetVswitchId(v string) *GetNatGatewayAttributeResponseBodyPrivateInfo {
	s.VswitchId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodyPrivateInfo) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayAttributeResponseBodySnatTable struct {
	// The number of SNAT entries.
	//
	// example:
	//
	// 1
	SnatEntryCount *int32 `json:"SnatEntryCount,omitempty" xml:"SnatEntryCount,omitempty"`
	// The ID of the SNAT table.
	//
	// example:
	//
	// stb-SnatTableIds****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
}

func (s GetNatGatewayAttributeResponseBodySnatTable) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponseBodySnatTable) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponseBodySnatTable) GetSnatEntryCount() *int32 {
	return s.SnatEntryCount
}

func (s *GetNatGatewayAttributeResponseBodySnatTable) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *GetNatGatewayAttributeResponseBodySnatTable) SetSnatEntryCount(v int32) *GetNatGatewayAttributeResponseBodySnatTable {
	s.SnatEntryCount = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodySnatTable) SetSnatTableId(v string) *GetNatGatewayAttributeResponseBodySnatTable {
	s.SnatTableId = &v
	return s
}

func (s *GetNatGatewayAttributeResponseBodySnatTable) Validate() error {
	return dara.Validate(s)
}
