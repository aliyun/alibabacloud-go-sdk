// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v *CreateRenderingInstanceRequestAttributes) *CreateRenderingInstanceRequest
	GetAttributes() *CreateRenderingInstanceRequestAttributes
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
	// > Unless you have specific requirements, keep the default values. For customers with special requirements, fill in the relevant parameters after communication and confirmation.
	//
	// Attribute information.
	Attributes *CreateRenderingInstanceRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Struct"`
	// > This value is valid only when `InstanceChargeType` is `PrePaid` (subscription).
	//
	// Enable or disable auto-renewal. Values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Client information.
	ClientInfo *CreateRenderingInstanceRequestClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	// > This value is valid only when `InstanceChargeType` is `PostPaid` (pay-as-you-go).
	//
	// Billing type. Valid values:
	//
	// - Hour: Hourly.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The billing method for the instance. Valid values:
	//
	// - PrePaid (default): Subscription.
	//
	// - PostPaid: Pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// > Only one billing method is allowed. If a billing method already exists, the new value is invalid by default, and the existing one prevails. Note that this field is required when a user makes the first purchase.
	//
	// Network billing type. Value:
	//
	// - 95BandwidthByMonth: Monthly 95th percentile bandwidth.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// Maximum bandwidth, in Mbps. Default is 10.
	//
	// example:
	//
	// 10
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" xml:"InternetMaxBandwidth,omitempty"`
	// > This value is valid only when `InstanceChargeType` is `PrePaid` (subscription).
	//
	// The subscription period. Valid values (Note: 12 is converted to one year; other values are in months):
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
	// Cloud application service instance specifications.
	//
	// - For crs.cp.\\	- series specifications, choose between subscription or pay-as-you-go billing.
	//
	// - For crs.vm.\\	- series specifications, choose between subscription or pay-as-you-go billing.
	//
	// - For other series, only subscription billing is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	// The cloud storage capacity used by the cloud application service instance (Note: not local storage).
	//
	// example:
	//
	// 20
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s CreateRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceRequest) GetAttributes() *CreateRenderingInstanceRequestAttributes {
	return s.Attributes
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

func (s *CreateRenderingInstanceRequest) SetAttributes(v *CreateRenderingInstanceRequestAttributes) *CreateRenderingInstanceRequest {
	s.Attributes = v
	return s
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
	if s.Attributes != nil {
		if err := s.Attributes.Validate(); err != nil {
			return err
		}
	}
	if s.ClientInfo != nil {
		if err := s.ClientInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRenderingInstanceRequestAttributes struct {
	// Streaming service.
	//
	// example:
	//
	// ON
	EdgeMediaService *string `json:"EdgeMediaService,omitempty" xml:"EdgeMediaService,omitempty"`
	// Inbound access configuration.
	//
	// example:
	//
	// ON
	InAccess *string `json:"InAccess,omitempty" xml:"InAccess,omitempty"`
	// Outbound access configuration.
	//
	// example:
	//
	// ON
	OutAccess *string `json:"OutAccess,omitempty" xml:"OutAccess,omitempty"`
	// Resource domain.
	//
	// example:
	//
	// Public
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateRenderingInstanceRequestAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceRequestAttributes) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceRequestAttributes) GetEdgeMediaService() *string {
	return s.EdgeMediaService
}

func (s *CreateRenderingInstanceRequestAttributes) GetInAccess() *string {
	return s.InAccess
}

func (s *CreateRenderingInstanceRequestAttributes) GetOutAccess() *string {
	return s.OutAccess
}

func (s *CreateRenderingInstanceRequestAttributes) GetZone() *string {
	return s.Zone
}

func (s *CreateRenderingInstanceRequestAttributes) SetEdgeMediaService(v string) *CreateRenderingInstanceRequestAttributes {
	s.EdgeMediaService = &v
	return s
}

func (s *CreateRenderingInstanceRequestAttributes) SetInAccess(v string) *CreateRenderingInstanceRequestAttributes {
	s.InAccess = &v
	return s
}

func (s *CreateRenderingInstanceRequestAttributes) SetOutAccess(v string) *CreateRenderingInstanceRequestAttributes {
	s.OutAccess = &v
	return s
}

func (s *CreateRenderingInstanceRequestAttributes) SetZone(v string) *CreateRenderingInstanceRequestAttributes {
	s.Zone = &v
	return s
}

func (s *CreateRenderingInstanceRequestAttributes) Validate() error {
	return dara.Validate(s)
}

type CreateRenderingInstanceRequestClientInfo struct {
	// Client IP address.
	//
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
