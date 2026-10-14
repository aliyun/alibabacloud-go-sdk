// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateARMServerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateARMServerInstancesRequest
	GetAmount() *int32
	SetAutoRenew(v bool) *CreateARMServerInstancesRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreateARMServerInstancesRequest
	GetAutoUseCoupon() *bool
	SetCidr(v string) *CreateARMServerInstancesRequest
	GetCidr() *string
	SetEnsRegionId(v string) *CreateARMServerInstancesRequest
	GetEnsRegionId() *string
	SetEnvironmentVar(v string) *CreateARMServerInstancesRequest
	GetEnvironmentVar() *string
	SetFrequency(v int32) *CreateARMServerInstancesRequest
	GetFrequency() *int32
	SetImageId(v string) *CreateARMServerInstancesRequest
	GetImageId() *string
	SetInstanceBillingCycle(v string) *CreateARMServerInstancesRequest
	GetInstanceBillingCycle() *string
	SetInstanceType(v string) *CreateARMServerInstancesRequest
	GetInstanceType() *string
	SetKeyPairName(v string) *CreateARMServerInstancesRequest
	GetKeyPairName() *string
	SetNameSpace(v string) *CreateARMServerInstancesRequest
	GetNameSpace() *string
	SetPayType(v string) *CreateARMServerInstancesRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateARMServerInstancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateARMServerInstancesRequest
	GetPeriodUnit() *string
	SetResolution(v string) *CreateARMServerInstancesRequest
	GetResolution() *string
	SetServerName(v string) *CreateARMServerInstancesRequest
	GetServerName() *string
	SetServerType(v string) *CreateARMServerInstancesRequest
	GetServerType() *string
	SetTag(v []*CreateARMServerInstancesRequestTag) *CreateARMServerInstancesRequest
	GetTag() []*CreateARMServerInstancesRequestTag
}

type CreateARMServerInstancesRequest struct {
	// The number of instances to create. Minimum value: **1**. Maximum value: **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The IPv4 CIDR block used by the instance, in CIDR notation.
	//
	// example:
	//
	// 10.0.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ENS edge zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-guiyang-12
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// One or more environment variables to set during EAIS instance initialization.
	//
	// example:
	//
	// [object Object]
	EnvironmentVar *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	// The refresh rate, in Hz. Valid values: 30 and 60.
	//
	// example:
	//
	// 30
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourImage ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance billing cycle. Valid values:
	//
	// - Hour
	//
	// - Day
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The instance type of the AIC container instance. Valid values:
	//
	// - aic.cf52r.c1.np
	//
	// - aic.cf52r.c2.np
	//
	// - aic.cf53r.c2.np
	//
	// - aic.cf52r.c4.np
	//
	// - aic.cf53r.c3.np
	//
	// - aic.cf52r.c3.np
	//
	// - aic.cf53r.c1.np
	//
	// - aic.cf53r.c5.np
	//
	// - aic.cf53r.c6
	//
	// - aic.cf53r.c4.np
	//
	// - aic.cf53r.c6.np
	//
	// - aic.cf53r.c7.np
	//
	// - aic.cf52m1r.c5.np
	//
	// - aic.cf53r.c8.np
	//
	// - aic.cf53r.c7
	//
	// - aic.cf52m1r.c2.np
	//
	// - aic.cf52m1r.c1.np
	//
	// - aic.cf52m1r.c3.np
	//
	// - aic.cf52m1r.c4.np
	//
	// - aic.cf52m1r.c6
	//
	// - ens.a6c2
	//
	// This parameter is required.
	//
	// example:
	//
	// aic.cf53r.c6.np
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The key pair name.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The namespace.
	//
	// example:
	//
	// pre
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The billing type. Valid values: **PrePaid*	- (subscription).
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The duration for which you want to purchase the resource.
	//
	// - If you do not specify PeriodUnit, the resource is purchased on a monthly basis by default. Currently, only Days and Month are supported.
	//
	// - If PeriodUnit is set to Day, Period can only be set to 3.
	//
	// - If PeriodUnit is set to Month, Period can be set to 1 to 9 or 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The resolution. Valid values:
	//
	// - 1920*864
	//
	// - 1080*1920
	//
	// - 1920*1080
	//
	// - 720*1280
	//
	// - 2400*1080
	//
	// - 1080*2400
	//
	// - 1280*720
	//
	// - 864*1920
	//
	// This parameter is required.
	//
	// example:
	//
	// 720*1280
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The service name.
	//
	// example:
	//
	// AIC-Server
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// The ARM board specification. Valid values:
	//
	// - cas.cf53r
	//
	// - cas.cf52r
	//
	// - cas.cf52m1r
	//
	// - cas.tg52g2
	//
	// - ens.afq-c2m3i.medium
	//
	// This parameter is required.
	//
	// example:
	//
	// cas.cf53r
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The list of tags bound to the instance.
	Tag []*CreateARMServerInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateARMServerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateARMServerInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateARMServerInstancesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateARMServerInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateARMServerInstancesRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateARMServerInstancesRequest) GetCidr() *string {
	return s.Cidr
}

func (s *CreateARMServerInstancesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateARMServerInstancesRequest) GetEnvironmentVar() *string {
	return s.EnvironmentVar
}

func (s *CreateARMServerInstancesRequest) GetFrequency() *int32 {
	return s.Frequency
}

func (s *CreateARMServerInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateARMServerInstancesRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateARMServerInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateARMServerInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateARMServerInstancesRequest) GetNameSpace() *string {
	return s.NameSpace
}

func (s *CreateARMServerInstancesRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateARMServerInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateARMServerInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateARMServerInstancesRequest) GetResolution() *string {
	return s.Resolution
}

func (s *CreateARMServerInstancesRequest) GetServerName() *string {
	return s.ServerName
}

func (s *CreateARMServerInstancesRequest) GetServerType() *string {
	return s.ServerType
}

func (s *CreateARMServerInstancesRequest) GetTag() []*CreateARMServerInstancesRequestTag {
	return s.Tag
}

func (s *CreateARMServerInstancesRequest) SetAmount(v int32) *CreateARMServerInstancesRequest {
	s.Amount = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetAutoRenew(v bool) *CreateARMServerInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetAutoUseCoupon(v bool) *CreateARMServerInstancesRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetCidr(v string) *CreateARMServerInstancesRequest {
	s.Cidr = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetEnsRegionId(v string) *CreateARMServerInstancesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetEnvironmentVar(v string) *CreateARMServerInstancesRequest {
	s.EnvironmentVar = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetFrequency(v int32) *CreateARMServerInstancesRequest {
	s.Frequency = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetImageId(v string) *CreateARMServerInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetInstanceBillingCycle(v string) *CreateARMServerInstancesRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetInstanceType(v string) *CreateARMServerInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetKeyPairName(v string) *CreateARMServerInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetNameSpace(v string) *CreateARMServerInstancesRequest {
	s.NameSpace = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetPayType(v string) *CreateARMServerInstancesRequest {
	s.PayType = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetPeriod(v int32) *CreateARMServerInstancesRequest {
	s.Period = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetPeriodUnit(v string) *CreateARMServerInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetResolution(v string) *CreateARMServerInstancesRequest {
	s.Resolution = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetServerName(v string) *CreateARMServerInstancesRequest {
	s.ServerName = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetServerType(v string) *CreateARMServerInstancesRequest {
	s.ServerType = &v
	return s
}

func (s *CreateARMServerInstancesRequest) SetTag(v []*CreateARMServerInstancesRequestTag) *CreateARMServerInstancesRequest {
	s.Tag = v
	return s
}

func (s *CreateARMServerInstancesRequest) Validate() error {
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

type CreateARMServerInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// backend
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateARMServerInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateARMServerInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *CreateARMServerInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateARMServerInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateARMServerInstancesRequestTag) SetKey(v string) *CreateARMServerInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *CreateARMServerInstancesRequestTag) SetValue(v string) *CreateARMServerInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *CreateARMServerInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
