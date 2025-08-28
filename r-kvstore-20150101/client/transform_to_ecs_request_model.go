// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToEcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *TransformToEcsRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v int64) *TransformToEcsRequest
	GetAutoRenewPeriod() *int64
	SetChargeType(v string) *TransformToEcsRequest
	GetChargeType() *string
	SetDryRun(v bool) *TransformToEcsRequest
	GetDryRun() *bool
	SetEffectiveTime(v string) *TransformToEcsRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *TransformToEcsRequest
	GetEngineVersion() *string
	SetInstanceClass(v string) *TransformToEcsRequest
	GetInstanceClass() *string
	SetInstanceId(v string) *TransformToEcsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *TransformToEcsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformToEcsRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *TransformToEcsRequest
	GetPeriod() *int64
	SetResourceOwnerAccount(v string) *TransformToEcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformToEcsRequest
	GetResourceOwnerId() *int64
	SetShardCount(v int64) *TransformToEcsRequest
	GetShardCount() *int64
}

type TransformToEcsRequest struct {
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tair.rdb.1g
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2
	ShardCount *int64 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
}

func (s TransformToEcsRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformToEcsRequest) GoString() string {
	return s.String()
}

func (s *TransformToEcsRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformToEcsRequest) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *TransformToEcsRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransformToEcsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *TransformToEcsRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *TransformToEcsRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *TransformToEcsRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *TransformToEcsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformToEcsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformToEcsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformToEcsRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *TransformToEcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformToEcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformToEcsRequest) GetShardCount() *int64 {
	return s.ShardCount
}

func (s *TransformToEcsRequest) SetAutoRenew(v string) *TransformToEcsRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformToEcsRequest) SetAutoRenewPeriod(v int64) *TransformToEcsRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *TransformToEcsRequest) SetChargeType(v string) *TransformToEcsRequest {
	s.ChargeType = &v
	return s
}

func (s *TransformToEcsRequest) SetDryRun(v bool) *TransformToEcsRequest {
	s.DryRun = &v
	return s
}

func (s *TransformToEcsRequest) SetEffectiveTime(v string) *TransformToEcsRequest {
	s.EffectiveTime = &v
	return s
}

func (s *TransformToEcsRequest) SetEngineVersion(v string) *TransformToEcsRequest {
	s.EngineVersion = &v
	return s
}

func (s *TransformToEcsRequest) SetInstanceClass(v string) *TransformToEcsRequest {
	s.InstanceClass = &v
	return s
}

func (s *TransformToEcsRequest) SetInstanceId(v string) *TransformToEcsRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformToEcsRequest) SetOwnerAccount(v string) *TransformToEcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformToEcsRequest) SetOwnerId(v int64) *TransformToEcsRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformToEcsRequest) SetPeriod(v int64) *TransformToEcsRequest {
	s.Period = &v
	return s
}

func (s *TransformToEcsRequest) SetResourceOwnerAccount(v string) *TransformToEcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformToEcsRequest) SetResourceOwnerId(v int64) *TransformToEcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformToEcsRequest) SetShardCount(v int64) *TransformToEcsRequest {
	s.ShardCount = &v
	return s
}

func (s *TransformToEcsRequest) Validate() error {
	return dara.Validate(s)
}
