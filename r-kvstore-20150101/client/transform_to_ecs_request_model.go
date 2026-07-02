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
	SetIsAcrossZone(v bool) *TransformToEcsRequest
	GetIsAcrossZone() *bool
	SetIzNo(v string) *TransformToEcsRequest
	GetIzNo() *string
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
	SetSecondaryIzNo(v string) *TransformToEcsRequest
	GetSecondaryIzNo() *string
	SetShardCount(v int64) *TransformToEcsRequest
	GetShardCount() *int64
	SetVSwitchId(v string) *TransformToEcsRequest
	GetVSwitchId() *string
}

type TransformToEcsRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal cycle. Unit: month. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required if you set **AutoRenew*	- to **true**.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The billing method of the target instance. Valid values:
	//
	// - **PostPaid**: pay-as-you-go
	//
	// - **PrePaid**: subscription. If you set this parameter to PrePaid, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run to check the request. The check items include the required parameters, request format, service limits, and available resources. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request and creates an instance after the request passes the check.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when to switch the database after data migration. Valid values:
	//
	// - **Immediately**: The database is immediately switched after the migration is complete.
	//
	// - **MaintainTime**: The database is switched within the maintenance window.
	//
	// > Default value: **Immediately**.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The Redis-compatible version of the instance. Valid values: **5.0**, **6.0**, and **7.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type of the target cloud-native instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/26350.html).
	//
	// > If you want to convert a cluster instance, you must specify the corresponding cloud-native cluster instance type that includes .with.proxy in its name and specify the ShardCount parameter.
	//
	// >
	//
	// > - For a cluster instance, you must provide the corresponding cloud-native cluster specification that includes `.proxy`. You must also specify the number of shards by using the `ShardCount` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tair.rdb.1g
	//
	// tair.rdb.with.proxy.1g
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the classic instance that you want to convert.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to deploy the instance across availability zones. This feature is supported only for cluster instances.
	IsAcrossZone *bool `json:"IsAcrossZone,omitempty" xml:"IsAcrossZone,omitempty"`
	// The ID of the availability zone.
	IzNo         *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: month. Valid values: **1**, **2**, **3**, **4**, **5**, 6, 7, 8, 9, 12, 24, and 36.
	//
	// > This parameter is available and required only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 1
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the secondary availability zone.
	SecondaryIzNo *string `json:"SecondaryIzNo,omitempty" xml:"SecondaryIzNo,omitempty"`
	// The number of data shards in the cloud-native cluster instance.
	//
	// example:
	//
	// 2
	ShardCount *int64 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *TransformToEcsRequest) GetIsAcrossZone() *bool {
	return s.IsAcrossZone
}

func (s *TransformToEcsRequest) GetIzNo() *string {
	return s.IzNo
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

func (s *TransformToEcsRequest) GetSecondaryIzNo() *string {
	return s.SecondaryIzNo
}

func (s *TransformToEcsRequest) GetShardCount() *int64 {
	return s.ShardCount
}

func (s *TransformToEcsRequest) GetVSwitchId() *string {
	return s.VSwitchId
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

func (s *TransformToEcsRequest) SetIsAcrossZone(v bool) *TransformToEcsRequest {
	s.IsAcrossZone = &v
	return s
}

func (s *TransformToEcsRequest) SetIzNo(v string) *TransformToEcsRequest {
	s.IzNo = &v
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

func (s *TransformToEcsRequest) SetSecondaryIzNo(v string) *TransformToEcsRequest {
	s.SecondaryIzNo = &v
	return s
}

func (s *TransformToEcsRequest) SetShardCount(v int64) *TransformToEcsRequest {
	s.ShardCount = &v
	return s
}

func (s *TransformToEcsRequest) SetVSwitchId(v string) *TransformToEcsRequest {
	s.VSwitchId = &v
	return s
}

func (s *TransformToEcsRequest) Validate() error {
	return dara.Validate(s)
}
