// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayShrinkRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *AddGatewayShrinkRequest
	GetChargeType() *string
	SetClbNetworkType(v string) *AddGatewayShrinkRequest
	GetClbNetworkType() *string
	SetEnableHardwareAcceleration(v bool) *AddGatewayShrinkRequest
	GetEnableHardwareAcceleration() *bool
	SetEnableSls(v bool) *AddGatewayShrinkRequest
	GetEnableSls() *bool
	SetEnableXtrace(v bool) *AddGatewayShrinkRequest
	GetEnableXtrace() *bool
	SetEnterpriseSecurityGroup(v bool) *AddGatewayShrinkRequest
	GetEnterpriseSecurityGroup() *bool
	SetInternetSlbSpec(v string) *AddGatewayShrinkRequest
	GetInternetSlbSpec() *string
	SetManagedEntryNetworkType(v string) *AddGatewayShrinkRequest
	GetManagedEntryNetworkType() *string
	SetMserVersion(v string) *AddGatewayShrinkRequest
	GetMserVersion() *string
	SetName(v string) *AddGatewayShrinkRequest
	GetName() *string
	SetNlbNetworkType(v string) *AddGatewayShrinkRequest
	GetNlbNetworkType() *string
	SetRegion(v string) *AddGatewayShrinkRequest
	GetRegion() *string
	SetReplica(v int32) *AddGatewayShrinkRequest
	GetReplica() *int32
	SetRequestPars(v string) *AddGatewayShrinkRequest
	GetRequestPars() *string
	SetResourceGroupId(v string) *AddGatewayShrinkRequest
	GetResourceGroupId() *string
	SetSlbSpec(v string) *AddGatewayShrinkRequest
	GetSlbSpec() *string
	SetSpec(v string) *AddGatewayShrinkRequest
	GetSpec() *string
	SetTag(v []*AddGatewayShrinkRequestTag) *AddGatewayShrinkRequest
	GetTag() []*AddGatewayShrinkRequestTag
	SetVSwitchId(v string) *AddGatewayShrinkRequest
	GetVSwitchId() *string
	SetVSwitchId2(v string) *AddGatewayShrinkRequest
	GetVSwitchId2() *string
	SetVpc(v string) *AddGatewayShrinkRequest
	GetVpc() *string
	SetXtraceRatio(v string) *AddGatewayShrinkRequest
	GetXtraceRatio() *string
	SetZoneInfoShrink(v string) *AddGatewayShrinkRequest
	GetZoneInfoShrink() *string
}

type AddGatewayShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The billing method you specify when you purchase an ordinary instance.
	//
	// Valid values:
	//
	// 	- PREPAY: subscription
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The network type of the purchased Classic Load Balancer (CLB) instance that is billed based on LCUs.
	//
	// 	- pubnet: Internet
	//
	// 	- privatenet: private network
	//
	// 	- privatepubnet: Internet and private network
	//
	// example:
	//
	// pubnet
	ClbNetworkType *string `json:"ClbNetworkType,omitempty" xml:"ClbNetworkType,omitempty"`
	// Specifies whether to activate Tracing Analysis.
	//
	// example:
	//
	// false
	EnableHardwareAcceleration *bool `json:"EnableHardwareAcceleration,omitempty" xml:"EnableHardwareAcceleration,omitempty"`
	// The tag of the gateway.
	//
	// example:
	//
	// false
	EnableSls *bool `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	// The sampling rate of Tracing Analysis. Valid values: [1,100].
	//
	// example:
	//
	// false
	EnableXtrace *bool `json:"EnableXtrace,omitempty" xml:"EnableXtrace,omitempty"`
	// Specifies whether to enable hardware acceleration.
	//
	// example:
	//
	// false
	EnterpriseSecurityGroup *bool `json:"EnterpriseSecurityGroup,omitempty" xml:"EnterpriseSecurityGroup,omitempty"`
	// Deprecated
	//
	// The specifications of the Internet-facing Server Load Balancer (SLB) instance. Valid values:
	//
	// 	- slb.s1.small
	//
	// 	- slb.s2.smal
	//
	// 	- slb.s2.medium
	//
	// 	- slb.s3.small
	//
	// 	- slb.s3.medium
	//
	// 	- slb.s3.large
	//
	// example:
	//
	// slb.s2.small
	InternetSlbSpec *string `json:"InternetSlbSpec,omitempty" xml:"InternetSlbSpec,omitempty"`
	// example:
	//
	// pubnet
	ManagedEntryNetworkType *string `json:"ManagedEntryNetworkType,omitempty" xml:"ManagedEntryNetworkType,omitempty"`
	// The MSE instance type. Valid values:
	//
	// 	- mse_pro: ordinary instance
	//
	// 	- mse_serverless: serverless instance
	//
	// example:
	//
	// mse_pro
	MserVersion *string `json:"MserVersion,omitempty" xml:"MserVersion,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// test-ceshi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the Network Load Balancer (NLB) instance you specify when you purchase a serverless instance.
	//
	// 	- pubnet: Internet
	//
	// 	- privatenet: private network
	//
	// 	- privatepubnet: Internet and private network
	//
	// example:
	//
	// pubnet
	NlbNetworkType *string `json:"NlbNetworkType,omitempty" xml:"NlbNetworkType,omitempty"`
	// The specifications of the internal-facing Server Load Balancer (SLB) instance. Valid values:
	//
	// 	- slb.s1.small
	//
	// 	- slb.s2.small
	//
	// 	- slb.s2.medium
	//
	// 	- slb.s3.small
	//
	// 	- slb.s3.medium
	//
	// 	- slb.s3.large
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of nodes you specify when you purchase an ordinary instance.
	//
	// example:
	//
	// 2
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
	// The extended field.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// rg-acfm34x43l*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Deprecated
	//
	// The specifications of the internal-facing Server Load Balancer (SLB) instance. Valid values:
	//
	// 	- slb.s1.small
	//
	// 	- slb.s2.small
	//
	// 	- slb.s2.medium
	//
	// 	- slb.s3.small
	//
	// 	- slb.s3.medium
	//
	// 	- slb.s3.large
	//
	// example:
	//
	// slb.s2.small
	SlbSpec *string `json:"SlbSpec,omitempty" xml:"SlbSpec,omitempty"`
	// The node specifications you specify when you purchase an ordinary instance. Valid values:
	//
	// 	- MSE_GTW_16_32_200_c(16C32G)
	//
	// 	- MSE_GTW_2_4_200_c(2C4G)
	//
	// 	- MSE_GTW_4_8_200_c(4C8G)
	//
	// 	- MSE_GTW_8_16_200_c(8C16G)
	//
	// example:
	//
	// MSE_GTW_2_4_200_c
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tag object.
	Tag []*AddGatewayShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the primary vSwitch.
	//
	// example:
	//
	// vsw-bp1q8th57frl5khj2li43
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Specifies whether to use an advanced security group.
	//
	// example:
	//
	// vsw-wz9bu6o5vsvitt5mrxo6s
	VSwitchId2 *string `json:"VSwitchId2,omitempty" xml:"VSwitchId2,omitempty"`
	// The ID of the primary vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp15mncnrtm83uauxd1xb
	Vpc *string `json:"Vpc,omitempty" xml:"Vpc,omitempty"`
	// Specifies whether to activate Log Service.
	//
	// example:
	//
	// 10
	XtraceRatio *string `json:"XtraceRatio,omitempty" xml:"XtraceRatio,omitempty"`
	// The details of the zone.
	ZoneInfoShrink *string `json:"ZoneInfo,omitempty" xml:"ZoneInfo,omitempty"`
}

func (s AddGatewayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *AddGatewayShrinkRequest) GetClbNetworkType() *string {
	return s.ClbNetworkType
}

func (s *AddGatewayShrinkRequest) GetEnableHardwareAcceleration() *bool {
	return s.EnableHardwareAcceleration
}

func (s *AddGatewayShrinkRequest) GetEnableSls() *bool {
	return s.EnableSls
}

func (s *AddGatewayShrinkRequest) GetEnableXtrace() *bool {
	return s.EnableXtrace
}

func (s *AddGatewayShrinkRequest) GetEnterpriseSecurityGroup() *bool {
	return s.EnterpriseSecurityGroup
}

func (s *AddGatewayShrinkRequest) GetInternetSlbSpec() *string {
	return s.InternetSlbSpec
}

func (s *AddGatewayShrinkRequest) GetManagedEntryNetworkType() *string {
	return s.ManagedEntryNetworkType
}

func (s *AddGatewayShrinkRequest) GetMserVersion() *string {
	return s.MserVersion
}

func (s *AddGatewayShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayShrinkRequest) GetNlbNetworkType() *string {
	return s.NlbNetworkType
}

func (s *AddGatewayShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *AddGatewayShrinkRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *AddGatewayShrinkRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *AddGatewayShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddGatewayShrinkRequest) GetSlbSpec() *string {
	return s.SlbSpec
}

func (s *AddGatewayShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *AddGatewayShrinkRequest) GetTag() []*AddGatewayShrinkRequestTag {
	return s.Tag
}

func (s *AddGatewayShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddGatewayShrinkRequest) GetVSwitchId2() *string {
	return s.VSwitchId2
}

func (s *AddGatewayShrinkRequest) GetVpc() *string {
	return s.Vpc
}

func (s *AddGatewayShrinkRequest) GetXtraceRatio() *string {
	return s.XtraceRatio
}

func (s *AddGatewayShrinkRequest) GetZoneInfoShrink() *string {
	return s.ZoneInfoShrink
}

func (s *AddGatewayShrinkRequest) SetAcceptLanguage(v string) *AddGatewayShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetChargeType(v string) *AddGatewayShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetClbNetworkType(v string) *AddGatewayShrinkRequest {
	s.ClbNetworkType = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetEnableHardwareAcceleration(v bool) *AddGatewayShrinkRequest {
	s.EnableHardwareAcceleration = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetEnableSls(v bool) *AddGatewayShrinkRequest {
	s.EnableSls = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetEnableXtrace(v bool) *AddGatewayShrinkRequest {
	s.EnableXtrace = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetEnterpriseSecurityGroup(v bool) *AddGatewayShrinkRequest {
	s.EnterpriseSecurityGroup = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetInternetSlbSpec(v string) *AddGatewayShrinkRequest {
	s.InternetSlbSpec = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetManagedEntryNetworkType(v string) *AddGatewayShrinkRequest {
	s.ManagedEntryNetworkType = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetMserVersion(v string) *AddGatewayShrinkRequest {
	s.MserVersion = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetName(v string) *AddGatewayShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetNlbNetworkType(v string) *AddGatewayShrinkRequest {
	s.NlbNetworkType = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetRegion(v string) *AddGatewayShrinkRequest {
	s.Region = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetReplica(v int32) *AddGatewayShrinkRequest {
	s.Replica = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetRequestPars(v string) *AddGatewayShrinkRequest {
	s.RequestPars = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetResourceGroupId(v string) *AddGatewayShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetSlbSpec(v string) *AddGatewayShrinkRequest {
	s.SlbSpec = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetSpec(v string) *AddGatewayShrinkRequest {
	s.Spec = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetTag(v []*AddGatewayShrinkRequestTag) *AddGatewayShrinkRequest {
	s.Tag = v
	return s
}

func (s *AddGatewayShrinkRequest) SetVSwitchId(v string) *AddGatewayShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetVSwitchId2(v string) *AddGatewayShrinkRequest {
	s.VSwitchId2 = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetVpc(v string) *AddGatewayShrinkRequest {
	s.Vpc = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetXtraceRatio(v string) *AddGatewayShrinkRequest {
	s.XtraceRatio = &v
	return s
}

func (s *AddGatewayShrinkRequest) SetZoneInfoShrink(v string) *AddGatewayShrinkRequest {
	s.ZoneInfoShrink = &v
	return s
}

func (s *AddGatewayShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type AddGatewayShrinkRequestTag struct {
	// The value of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddGatewayShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *AddGatewayShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddGatewayShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddGatewayShrinkRequestTag) SetKey(v string) *AddGatewayShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *AddGatewayShrinkRequestTag) SetValue(v string) *AddGatewayShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *AddGatewayShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
