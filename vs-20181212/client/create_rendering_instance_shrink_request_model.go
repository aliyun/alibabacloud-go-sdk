// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributesShrink(v string) *CreateRenderingInstanceShrinkRequest
	GetAttributesShrink() *string
	SetAutoRenew(v bool) *CreateRenderingInstanceShrinkRequest
	GetAutoRenew() *bool
	SetClientInfoShrink(v string) *CreateRenderingInstanceShrinkRequest
	GetClientInfoShrink() *string
	SetInstanceBillingCycle(v string) *CreateRenderingInstanceShrinkRequest
	GetInstanceBillingCycle() *string
	SetInstanceChargeType(v string) *CreateRenderingInstanceShrinkRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *CreateRenderingInstanceShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidth(v int64) *CreateRenderingInstanceShrinkRequest
	GetInternetMaxBandwidth() *int64
	SetPeriod(v string) *CreateRenderingInstanceShrinkRequest
	GetPeriod() *string
	SetRenderingSpec(v string) *CreateRenderingInstanceShrinkRequest
	GetRenderingSpec() *string
	SetStorageSize(v string) *CreateRenderingInstanceShrinkRequest
	GetStorageSize() *string
}

type CreateRenderingInstanceShrinkRequest struct {
	// > Keep the default value unless you have special requirements. If you have special requirements, contact technical support to confirm and then specify the relevant parameters.
	//
	// The attribute information.
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// > This parameter takes effect only when `InstanceChargeType` is set to `PrePaid` (subscription).
	//
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The client information.
	ClientInfoShrink *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	// > This parameter takes effect only when `InstanceChargeType` is set to `PostPaid` (pay-as-you-go).
	//
	// The billing cycle. Valid values:
	//
	// - Hour: hour.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PrePaid (default): subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// > Only one billing method is supported. If a billing method already exists, the new value is invalid by default, and the existing billing method prevails. Note that this parameter is required when you purchase the instance for the first time.
	//
	// The network billing type. Valid values:
	//
	// - 95BandwidthByMonth: monthly 95th percentile bandwidth.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum bandwidth. Unit: Mbit/s. Default value: 10.
	//
	// example:
	//
	// 10
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" xml:"InternetMaxBandwidth,omitempty"`
	// > This parameter takes effect only when `InstanceChargeType` is set to `PrePaid` (subscription).
	//
	// The subscription duration. Valid values (Note: If you select 12, the duration is calculated as one year. For other values, the unit is months):
	//
	// - 1 (default)
	//
	// - 2
	//
	// - 3
	//
	// - 4
	//
	// - 5
	//
	// - 6
	//
	// - 7
	//
	// - 8
	//
	// - 9
	//
	// - 12
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The specification of the cloud application service instance.
	//
	// - If the specification belongs to the crs.cp.	- series, the billing method can be subscription or pay-as-you-go.
	//
	// - If the specification belongs to the crs.vm.	- series, the billing method can be subscription or pay-as-you-go.
	//
	// - For other series, only the subscription billing method is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	// The cloud storage capacity used by the cloud application service instance. Note: This is not local storage.
	//
	// example:
	//
	// 20
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s CreateRenderingInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceShrinkRequest) GetAttributesShrink() *string {
	return s.AttributesShrink
}

func (s *CreateRenderingInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRenderingInstanceShrinkRequest) GetClientInfoShrink() *string {
	return s.ClientInfoShrink
}

func (s *CreateRenderingInstanceShrinkRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateRenderingInstanceShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRenderingInstanceShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateRenderingInstanceShrinkRequest) GetInternetMaxBandwidth() *int64 {
	return s.InternetMaxBandwidth
}

func (s *CreateRenderingInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateRenderingInstanceShrinkRequest) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *CreateRenderingInstanceShrinkRequest) GetStorageSize() *string {
	return s.StorageSize
}

func (s *CreateRenderingInstanceShrinkRequest) SetAttributesShrink(v string) *CreateRenderingInstanceShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetAutoRenew(v bool) *CreateRenderingInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetClientInfoShrink(v string) *CreateRenderingInstanceShrinkRequest {
	s.ClientInfoShrink = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetInstanceBillingCycle(v string) *CreateRenderingInstanceShrinkRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetInstanceChargeType(v string) *CreateRenderingInstanceShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetInternetChargeType(v string) *CreateRenderingInstanceShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetInternetMaxBandwidth(v int64) *CreateRenderingInstanceShrinkRequest {
	s.InternetMaxBandwidth = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetPeriod(v string) *CreateRenderingInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetRenderingSpec(v string) *CreateRenderingInstanceShrinkRequest {
	s.RenderingSpec = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) SetStorageSize(v string) *CreateRenderingInstanceShrinkRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateRenderingInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
