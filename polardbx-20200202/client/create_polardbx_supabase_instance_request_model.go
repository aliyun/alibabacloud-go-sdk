// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolardbxSupabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreatePolardbxSupabaseInstanceRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreatePolardbxSupabaseInstanceRequest
	GetClientToken() *string
	SetDashboardPassword(v string) *CreatePolardbxSupabaseInstanceRequest
	GetDashboardPassword() *string
	SetDbInstanceDescription(v string) *CreatePolardbxSupabaseInstanceRequest
	GetDbInstanceDescription() *string
	SetDbPassword(v string) *CreatePolardbxSupabaseInstanceRequest
	GetDbPassword() *string
	SetPayType(v string) *CreatePolardbxSupabaseInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreatePolardbxSupabaseInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *CreatePolardbxSupabaseInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePolardbxSupabaseInstanceRequest
	GetResourceGroupId() *string
	SetTenantMode(v bool) *CreatePolardbxSupabaseInstanceRequest
	GetTenantMode() *bool
	SetUsedTime(v int32) *CreatePolardbxSupabaseInstanceRequest
	GetUsedTime() *int32
	SetVSwitchId(v string) *CreatePolardbxSupabaseInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreatePolardbxSupabaseInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreatePolardbxSupabaseInstanceRequest
	GetZoneId() *string
}

type CreatePolardbxSupabaseInstanceRequest struct {
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dashboard password.
	//
	// This parameter is required.
	//
	// example:
	//
	// dTyMQem0o3HOAFh!
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Supabase实例
	DbInstanceDescription *string `json:"DbInstanceDescription,omitempty" xml:"DbInstanceDescription,omitempty"`
	// The database password.
	//
	// This parameter is required.
	//
	// example:
	//
	// dTyMQem0o3HOAFh!
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// The billing type. Valid values:
	//
	// - PREPAY: subscription.
	//
	// - POSTPAY: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle. Valid values:
	//
	// - Year
	//
	// - Month
	//
	// - Hour
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable multi-tenant mode.
	//
	// example:
	//
	// false
	TenantMode *bool `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ndoug37dtwoedlmru0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreatePolardbxSupabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolardbxSupabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetDbInstanceDescription() *string {
	return s.DbInstanceDescription
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetDbPassword() *string {
	return s.DbPassword
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetTenantMode() *bool {
	return s.TenantMode
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreatePolardbxSupabaseInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetAutoRenew(v bool) *CreatePolardbxSupabaseInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetClientToken(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetDashboardPassword(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.DashboardPassword = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetDbInstanceDescription(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.DbInstanceDescription = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetDbPassword(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.DbPassword = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetPayType(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetPeriod(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetRegionId(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetResourceGroupId(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetTenantMode(v bool) *CreatePolardbxSupabaseInstanceRequest {
	s.TenantMode = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetUsedTime(v int32) *CreatePolardbxSupabaseInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetVSwitchId(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetVpcId(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) SetZoneId(v string) *CreatePolardbxSupabaseInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
