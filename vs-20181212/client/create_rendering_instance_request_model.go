// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateRenderingInstanceRequest
	GetAutoRenew() *bool
	SetClientInfo(v *CreateRenderingInstanceRequestClientInfo) *CreateRenderingInstanceRequest
	GetClientInfo() *CreateRenderingInstanceRequestClientInfo
	SetInstanceBillingCycle(v string) *CreateRenderingInstanceRequest
	GetInstanceBillingCycle() *string
	SetInstanceChargeType(v string) *CreateRenderingInstanceRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *CreateRenderingInstanceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidth(v int64) *CreateRenderingInstanceRequest
	GetInternetMaxBandwidth() *int64
	SetPeriod(v string) *CreateRenderingInstanceRequest
	GetPeriod() *string
	SetRenderingSpec(v string) *CreateRenderingInstanceRequest
	GetRenderingSpec() *string
	SetStorageSize(v string) *CreateRenderingInstanceRequest
	GetStorageSize() *string
}

type CreateRenderingInstanceRequest struct {
	// example:
	//
	// true
	AutoRenew            *bool                                     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientInfo           *CreateRenderingInstanceRequestClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	InstanceBillingCycle *string                                   `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
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

func (s CreateRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRenderingInstanceRequest) GetClientInfo() *CreateRenderingInstanceRequestClientInfo {
	return s.ClientInfo
}

func (s *CreateRenderingInstanceRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateRenderingInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRenderingInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateRenderingInstanceRequest) GetInternetMaxBandwidth() *int64 {
	return s.InternetMaxBandwidth
}

func (s *CreateRenderingInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateRenderingInstanceRequest) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *CreateRenderingInstanceRequest) GetStorageSize() *string {
	return s.StorageSize
}

func (s *CreateRenderingInstanceRequest) SetAutoRenew(v bool) *CreateRenderingInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetClientInfo(v *CreateRenderingInstanceRequestClientInfo) *CreateRenderingInstanceRequest {
	s.ClientInfo = v
	return s
}

func (s *CreateRenderingInstanceRequest) SetInstanceBillingCycle(v string) *CreateRenderingInstanceRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetInstanceChargeType(v string) *CreateRenderingInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetInternetChargeType(v string) *CreateRenderingInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetInternetMaxBandwidth(v int64) *CreateRenderingInstanceRequest {
	s.InternetMaxBandwidth = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetPeriod(v string) *CreateRenderingInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetRenderingSpec(v string) *CreateRenderingInstanceRequest {
	s.RenderingSpec = &v
	return s
}

func (s *CreateRenderingInstanceRequest) SetStorageSize(v string) *CreateRenderingInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRenderingInstanceRequestClientInfo struct {
	// example:
	//
	// 172.21.128.110
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
}

func (s CreateRenderingInstanceRequestClientInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceRequestClientInfo) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceRequestClientInfo) GetClientIp() *string {
	return s.ClientIp
}

func (s *CreateRenderingInstanceRequestClientInfo) SetClientIp(v string) *CreateRenderingInstanceRequestClientInfo {
	s.ClientIp = &v
	return s
}

func (s *CreateRenderingInstanceRequestClientInfo) Validate() error {
	return dara.Validate(s)
}
