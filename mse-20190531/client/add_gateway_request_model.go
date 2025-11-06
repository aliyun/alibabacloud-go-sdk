// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *AddGatewayRequest
	GetChargeType() *string
	SetClbNetworkType(v string) *AddGatewayRequest
	GetClbNetworkType() *string
	SetEnableHardwareAcceleration(v bool) *AddGatewayRequest
	GetEnableHardwareAcceleration() *bool
	SetEnableSls(v bool) *AddGatewayRequest
	GetEnableSls() *bool
	SetEnableXtrace(v bool) *AddGatewayRequest
	GetEnableXtrace() *bool
	SetEnterpriseSecurityGroup(v bool) *AddGatewayRequest
	GetEnterpriseSecurityGroup() *bool
	SetInternetSlbSpec(v string) *AddGatewayRequest
	GetInternetSlbSpec() *string
	SetManagedEntryNetworkType(v string) *AddGatewayRequest
	GetManagedEntryNetworkType() *string
	SetMserVersion(v string) *AddGatewayRequest
	GetMserVersion() *string
	SetName(v string) *AddGatewayRequest
	GetName() *string
	SetNlbNetworkType(v string) *AddGatewayRequest
	GetNlbNetworkType() *string
	SetRegion(v string) *AddGatewayRequest
	GetRegion() *string
	SetReplica(v int32) *AddGatewayRequest
	GetReplica() *int32
	SetRequestPars(v string) *AddGatewayRequest
	GetRequestPars() *string
	SetResourceGroupId(v string) *AddGatewayRequest
	GetResourceGroupId() *string
	SetSlbSpec(v string) *AddGatewayRequest
	GetSlbSpec() *string
	SetSpec(v string) *AddGatewayRequest
	GetSpec() *string
	SetTag(v []*AddGatewayRequestTag) *AddGatewayRequest
	GetTag() []*AddGatewayRequestTag
	SetVSwitchId(v string) *AddGatewayRequest
	GetVSwitchId() *string
	SetVSwitchId2(v string) *AddGatewayRequest
	GetVSwitchId2() *string
	SetVpc(v string) *AddGatewayRequest
	GetVpc() *string
	SetXtraceRatio(v string) *AddGatewayRequest
	GetXtraceRatio() *string
	SetZoneInfo(v []*AddGatewayRequestZoneInfo) *AddGatewayRequest
	GetZoneInfo() []*AddGatewayRequestZoneInfo
}

type AddGatewayRequest struct {
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
	Tag []*AddGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	ZoneInfo []*AddGatewayRequestZoneInfo `json:"ZoneInfo,omitempty" xml:"ZoneInfo,omitempty" type:"Repeated"`
}

func (s AddGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *AddGatewayRequest) GetClbNetworkType() *string {
	return s.ClbNetworkType
}

func (s *AddGatewayRequest) GetEnableHardwareAcceleration() *bool {
	return s.EnableHardwareAcceleration
}

func (s *AddGatewayRequest) GetEnableSls() *bool {
	return s.EnableSls
}

func (s *AddGatewayRequest) GetEnableXtrace() *bool {
	return s.EnableXtrace
}

func (s *AddGatewayRequest) GetEnterpriseSecurityGroup() *bool {
	return s.EnterpriseSecurityGroup
}

func (s *AddGatewayRequest) GetInternetSlbSpec() *string {
	return s.InternetSlbSpec
}

func (s *AddGatewayRequest) GetManagedEntryNetworkType() *string {
	return s.ManagedEntryNetworkType
}

func (s *AddGatewayRequest) GetMserVersion() *string {
	return s.MserVersion
}

func (s *AddGatewayRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayRequest) GetNlbNetworkType() *string {
	return s.NlbNetworkType
}

func (s *AddGatewayRequest) GetRegion() *string {
	return s.Region
}

func (s *AddGatewayRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *AddGatewayRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *AddGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddGatewayRequest) GetSlbSpec() *string {
	return s.SlbSpec
}

func (s *AddGatewayRequest) GetSpec() *string {
	return s.Spec
}

func (s *AddGatewayRequest) GetTag() []*AddGatewayRequestTag {
	return s.Tag
}

func (s *AddGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddGatewayRequest) GetVSwitchId2() *string {
	return s.VSwitchId2
}

func (s *AddGatewayRequest) GetVpc() *string {
	return s.Vpc
}

func (s *AddGatewayRequest) GetXtraceRatio() *string {
	return s.XtraceRatio
}

func (s *AddGatewayRequest) GetZoneInfo() []*AddGatewayRequestZoneInfo {
	return s.ZoneInfo
}

func (s *AddGatewayRequest) SetAcceptLanguage(v string) *AddGatewayRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayRequest) SetChargeType(v string) *AddGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *AddGatewayRequest) SetClbNetworkType(v string) *AddGatewayRequest {
	s.ClbNetworkType = &v
	return s
}

func (s *AddGatewayRequest) SetEnableHardwareAcceleration(v bool) *AddGatewayRequest {
	s.EnableHardwareAcceleration = &v
	return s
}

func (s *AddGatewayRequest) SetEnableSls(v bool) *AddGatewayRequest {
	s.EnableSls = &v
	return s
}

func (s *AddGatewayRequest) SetEnableXtrace(v bool) *AddGatewayRequest {
	s.EnableXtrace = &v
	return s
}

func (s *AddGatewayRequest) SetEnterpriseSecurityGroup(v bool) *AddGatewayRequest {
	s.EnterpriseSecurityGroup = &v
	return s
}

func (s *AddGatewayRequest) SetInternetSlbSpec(v string) *AddGatewayRequest {
	s.InternetSlbSpec = &v
	return s
}

func (s *AddGatewayRequest) SetManagedEntryNetworkType(v string) *AddGatewayRequest {
	s.ManagedEntryNetworkType = &v
	return s
}

func (s *AddGatewayRequest) SetMserVersion(v string) *AddGatewayRequest {
	s.MserVersion = &v
	return s
}

func (s *AddGatewayRequest) SetName(v string) *AddGatewayRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayRequest) SetNlbNetworkType(v string) *AddGatewayRequest {
	s.NlbNetworkType = &v
	return s
}

func (s *AddGatewayRequest) SetRegion(v string) *AddGatewayRequest {
	s.Region = &v
	return s
}

func (s *AddGatewayRequest) SetReplica(v int32) *AddGatewayRequest {
	s.Replica = &v
	return s
}

func (s *AddGatewayRequest) SetRequestPars(v string) *AddGatewayRequest {
	s.RequestPars = &v
	return s
}

func (s *AddGatewayRequest) SetResourceGroupId(v string) *AddGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddGatewayRequest) SetSlbSpec(v string) *AddGatewayRequest {
	s.SlbSpec = &v
	return s
}

func (s *AddGatewayRequest) SetSpec(v string) *AddGatewayRequest {
	s.Spec = &v
	return s
}

func (s *AddGatewayRequest) SetTag(v []*AddGatewayRequestTag) *AddGatewayRequest {
	s.Tag = v
	return s
}

func (s *AddGatewayRequest) SetVSwitchId(v string) *AddGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddGatewayRequest) SetVSwitchId2(v string) *AddGatewayRequest {
	s.VSwitchId2 = &v
	return s
}

func (s *AddGatewayRequest) SetVpc(v string) *AddGatewayRequest {
	s.Vpc = &v
	return s
}

func (s *AddGatewayRequest) SetXtraceRatio(v string) *AddGatewayRequest {
	s.XtraceRatio = &v
	return s
}

func (s *AddGatewayRequest) SetZoneInfo(v []*AddGatewayRequestZoneInfo) *AddGatewayRequest {
	s.ZoneInfo = v
	return s
}

func (s *AddGatewayRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneInfo != nil {
		for _, item := range s.ZoneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayRequestTag struct {
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

func (s AddGatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRequestTag) GoString() string {
	return s.String()
}

func (s *AddGatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddGatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddGatewayRequestTag) SetKey(v string) *AddGatewayRequestTag {
	s.Key = &v
	return s
}

func (s *AddGatewayRequestTag) SetValue(v string) *AddGatewayRequestTag {
	s.Value = &v
	return s
}

func (s *AddGatewayRequestTag) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRequestZoneInfo struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddGatewayRequestZoneInfo) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRequestZoneInfo) GoString() string {
	return s.String()
}

func (s *AddGatewayRequestZoneInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddGatewayRequestZoneInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddGatewayRequestZoneInfo) SetVSwitchId(v string) *AddGatewayRequestZoneInfo {
	s.VSwitchId = &v
	return s
}

func (s *AddGatewayRequestZoneInfo) SetZoneId(v string) *AddGatewayRequestZoneInfo {
	s.ZoneId = &v
	return s
}

func (s *AddGatewayRequestZoneInfo) Validate() error {
	return dara.Validate(s)
}
