// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// example:
	//
	// true
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientInfoShrink     *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" xml:"InternetMaxBandwidth,omitempty"`
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	StorageSize   *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s CreateRenderingInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceShrinkRequest) GoString() string {
	return s.String()
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
