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
	// The number of instances to create. Valid values: **1*	- to **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal for the subscription. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Cidr          *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-guiyang-12
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Set one or more environment variables during EAIS instance initialization.
	//
	// example:
	//
	// [object Object]
	EnvironmentVar *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	// The refresh rate. Unit: Hz. Valid values: 30 and 60.
	//
	// example:
	//
	// 30
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourImage ID
	ImageId              *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The specification of the Android in Container (AIC) instance. Examples:
	//
	// 	- aic.cf52r.c1.np
	//
	// 	- aic.cf52r.c2.np
	//
	// 	- aic.cf53r.c2.np
	//
	// 	- aic.cf52r.c4.np
	//
	// 	- aic.cf53r.c3.np
	//
	// 	- aic.cf52r.c3.np
	//
	// 	- aic.cf53r.c1.np
	//
	// 	- aic.cf53r.c5.np
	//
	// 	- aic.cf53r.c6
	//
	// 	- aic.cf53r.c4.np
	//
	// 	- aic.cf53r.c6.np
	//
	// 	- aic.cf53r.c7.np
	//
	// 	- aic.cf52m1r.c5.np
	//
	// 	- aic.cf53r.c8.np
	//
	// 	- aic.cf53r.c7
	//
	// 	- aic.cf52m1r.c2.np
	//
	// 	- aic.cf52m1r.c1.np
	//
	// 	- aic.cf52m1r.c3.np
	//
	// 	- aic.cf52m1r.c4.np
	//
	// 	- aic.cf52m1r.c6
	//
	// 	- ens.a6c2
	//
	// This parameter is required.
	//
	// example:
	//
	// aic.cf53r.c6.np
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the key pair.
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
	// The billing method. Set the value to **PrePaid**. PrePaid specifies the subscription billing method.
	//
	// >  Only PrePaid is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If you leave PeriodUnit empty, the instance is purchased on a monthly basis. Valid values: Day and Month.
	//
	// 	- If you set PeriodUnit to Day, you can set Period only to 3.
	//
	// 	- If you set PeriodUnit to Month, you can set Period to a value within the range of [1,9], or set the value to 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// 	- If you leave PeriodUnit empty, the instance is purchased on a monthly basis. Valid values: Day and Month.
	//
	// 	- If you set PeriodUnit to Day, you can set Period only to 3.
	//
	// 	- If you set PeriodUnit to Month, you can set Period to a value within the range of [1,9], or set the value to 12.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The resolution. Examples:
	//
	// 	- 1920\\*864
	//
	// 	- 1080\\*1920
	//
	// 	- 1920\\*1080
	//
	// 	- 720\\*1280
	//
	// 	- 2400\\*1080
	//
	// 	- 1080\\*2400
	//
	// 	- 1280\\*720
	//
	// 	- 864\\*1920
	//
	// This parameter is required.
	//
	// example:
	//
	// 720*1280
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// AIC-Server
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// The specification of the ARM server. Examples:
	//
	// 	- cas.cf53r
	//
	// 	- cas.cf52r
	//
	// 	- cas.cf52m1r
	//
	// 	- cas.tg52g2
	//
	// 	- ens.afq-c2m3i.medium
	//
	// This parameter is required.
	//
	// example:
	//
	// cas.cf53r
	ServerType *string                               `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	Tag        []*CreateARMServerInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateARMServerInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
