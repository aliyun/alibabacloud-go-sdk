// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLogstashRequest
	GetDescription() *string
	SetNetworkConfig(v *CreateLogstashRequestNetworkConfig) *CreateLogstashRequest
	GetNetworkConfig() *CreateLogstashRequestNetworkConfig
	SetNodeAmount(v int32) *CreateLogstashRequest
	GetNodeAmount() *int32
	SetNodeSpec(v *CreateLogstashRequestNodeSpec) *CreateLogstashRequest
	GetNodeSpec() *CreateLogstashRequestNodeSpec
	SetPaymentInfo(v *CreateLogstashRequestPaymentInfo) *CreateLogstashRequest
	GetPaymentInfo() *CreateLogstashRequestPaymentInfo
	SetPaymentType(v string) *CreateLogstashRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *CreateLogstashRequest
	GetResourceGroupId() *string
	SetVersion(v string) *CreateLogstashRequest
	GetVersion() *string
	SetClientToken(v string) *CreateLogstashRequest
	GetClientToken() *string
}

type CreateLogstashRequest struct {
	// example:
	//
	// ls-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	NetworkConfig *CreateLogstashRequestNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// This parameter is required.
	NodeSpec    *CreateLogstashRequestNodeSpec    `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentInfo *CreateLogstashRequestPaymentInfo `json:"paymentInfo,omitempty" xml:"paymentInfo,omitempty" type:"Struct"`
	// example:
	//
	// prepaid
	PaymentType     *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6.7_with_X-Pack
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashRequest) GoString() string {
	return s.String()
}

func (s *CreateLogstashRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLogstashRequest) GetNetworkConfig() *CreateLogstashRequestNetworkConfig {
	return s.NetworkConfig
}

func (s *CreateLogstashRequest) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *CreateLogstashRequest) GetNodeSpec() *CreateLogstashRequestNodeSpec {
	return s.NodeSpec
}

func (s *CreateLogstashRequest) GetPaymentInfo() *CreateLogstashRequestPaymentInfo {
	return s.PaymentInfo
}

func (s *CreateLogstashRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateLogstashRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLogstashRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateLogstashRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLogstashRequest) SetDescription(v string) *CreateLogstashRequest {
	s.Description = &v
	return s
}

func (s *CreateLogstashRequest) SetNetworkConfig(v *CreateLogstashRequestNetworkConfig) *CreateLogstashRequest {
	s.NetworkConfig = v
	return s
}

func (s *CreateLogstashRequest) SetNodeAmount(v int32) *CreateLogstashRequest {
	s.NodeAmount = &v
	return s
}

func (s *CreateLogstashRequest) SetNodeSpec(v *CreateLogstashRequestNodeSpec) *CreateLogstashRequest {
	s.NodeSpec = v
	return s
}

func (s *CreateLogstashRequest) SetPaymentInfo(v *CreateLogstashRequestPaymentInfo) *CreateLogstashRequest {
	s.PaymentInfo = v
	return s
}

func (s *CreateLogstashRequest) SetPaymentType(v string) *CreateLogstashRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateLogstashRequest) SetResourceGroupId(v string) *CreateLogstashRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLogstashRequest) SetVersion(v string) *CreateLogstashRequest {
	s.Version = &v
	return s
}

func (s *CreateLogstashRequest) SetClientToken(v string) *CreateLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLogstashRequest) Validate() error {
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	if s.PaymentInfo != nil {
		if err := s.PaymentInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLogstashRequestNetworkConfig struct {
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1k4ec6s7sjdbudw****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s CreateLogstashRequestNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashRequestNetworkConfig) GoString() string {
	return s.String()
}

func (s *CreateLogstashRequestNetworkConfig) GetType() *string {
	return s.Type
}

func (s *CreateLogstashRequestNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLogstashRequestNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *CreateLogstashRequestNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateLogstashRequestNetworkConfig) SetType(v string) *CreateLogstashRequestNetworkConfig {
	s.Type = &v
	return s
}

func (s *CreateLogstashRequestNetworkConfig) SetVpcId(v string) *CreateLogstashRequestNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *CreateLogstashRequestNetworkConfig) SetVsArea(v string) *CreateLogstashRequestNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *CreateLogstashRequestNetworkConfig) SetVswitchId(v string) *CreateLogstashRequestNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *CreateLogstashRequestNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type CreateLogstashRequestNodeSpec struct {
	// example:
	//
	// 50
	Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// logstash.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateLogstashRequestNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashRequestNodeSpec) GoString() string {
	return s.String()
}

func (s *CreateLogstashRequestNodeSpec) GetDisk() *int64 {
	return s.Disk
}

func (s *CreateLogstashRequestNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateLogstashRequestNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *CreateLogstashRequestNodeSpec) SetDisk(v int64) *CreateLogstashRequestNodeSpec {
	s.Disk = &v
	return s
}

func (s *CreateLogstashRequestNodeSpec) SetDiskType(v string) *CreateLogstashRequestNodeSpec {
	s.DiskType = &v
	return s
}

func (s *CreateLogstashRequestNodeSpec) SetSpec(v string) *CreateLogstashRequestNodeSpec {
	s.Spec = &v
	return s
}

func (s *CreateLogstashRequestNodeSpec) Validate() error {
	return dara.Validate(s)
}

type CreateLogstashRequestPaymentInfo struct {
	// example:
	//
	// 3
	AutoRenewDuration *int64 `json:"autoRenewDuration,omitempty" xml:"autoRenewDuration,omitempty"`
	// example:
	//
	// 1
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// false
	IsAutoRenew *bool `json:"isAutoRenew,omitempty" xml:"isAutoRenew,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s CreateLogstashRequestPaymentInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashRequestPaymentInfo) GoString() string {
	return s.String()
}

func (s *CreateLogstashRequestPaymentInfo) GetAutoRenewDuration() *int64 {
	return s.AutoRenewDuration
}

func (s *CreateLogstashRequestPaymentInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateLogstashRequestPaymentInfo) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *CreateLogstashRequestPaymentInfo) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateLogstashRequestPaymentInfo) SetAutoRenewDuration(v int64) *CreateLogstashRequestPaymentInfo {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateLogstashRequestPaymentInfo) SetDuration(v int64) *CreateLogstashRequestPaymentInfo {
	s.Duration = &v
	return s
}

func (s *CreateLogstashRequestPaymentInfo) SetIsAutoRenew(v bool) *CreateLogstashRequestPaymentInfo {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateLogstashRequestPaymentInfo) SetPricingCycle(v string) *CreateLogstashRequestPaymentInfo {
	s.PricingCycle = &v
	return s
}

func (s *CreateLogstashRequestPaymentInfo) Validate() error {
	return dara.Validate(s)
}
