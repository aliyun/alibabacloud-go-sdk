// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBInstanceShrinkRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceShrinkRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceShrinkRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceShrinkRequest
	GetColdDataEnabled() *bool
	SetDBInstanceClass(v string) *ModifyDBInstanceShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBNodesShrink(v string) *ModifyDBInstanceShrinkRequest
	GetDBNodesShrink() *string
	SetDirection(v string) *ModifyDBInstanceShrinkRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceShrinkRequest
	GetEffectiveTime() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceShrinkRequest
	GetIoAccelerationEnabled() *string
	SetOwnerAccount(v string) *ModifyDBInstanceShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceShrinkRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBInstanceShrinkRequest
	GetParameterGroupId() *string
	SetParametersShrink(v string) *ModifyDBInstanceShrinkRequest
	GetParametersShrink() *string
	SetPromotionCode(v string) *ModifyDBInstanceShrinkRequest
	GetPromotionCode() *string
	SetResourceGroupId(v string) *ModifyDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyDBInstanceShrinkRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceShrinkRequest
	GetTargetMinorVersion() *string
}

type ModifyDBInstanceShrinkRequest struct {
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- **true*	- (default): Yes.
	//
	// 	- **false**: No.
	//
	// > After a coupon is used, if a downgrade is required, the amount deducted by the coupon will not be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Specifies whether to enable the [I/O performance burst feature for Premium ESSD cloud disks](https://help.aliyun.com/document_detail/2340501.html). Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The instance edition. Valid values:
	//
	// - **Basic**: Basic Edition
	//
	// - **HighAvailability**: High-availability Edition
	//
	// - **cluster**: Cluster Edition
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// <props="china">Specifies whether to enable the [data archiving feature for general cloud disks](https://help.aliyun.com/document_detail/2701832.html). Valid values:
	//
	// - <props="china">**true**: Enabled.
	//
	// - <props="china">**false**: Disabled.
	//
	// <props="intl">Reserved parameter.
	//
	// example:
	//
	// true
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// pg.n4.2c.1m
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The [target storage capacity](https://help.aliyun.com/document_detail/26312.html). Unit: GB. You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available storage capacity range for the target instance type.
	//
	// > 	- You must specify at least one of this parameter and the **DBInstanceClass*	- parameter.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to view the current storage capacity of the instance.
	//
	// example:
	//
	// 500
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **general_essd**: Premium ESSD (recommended)
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The node information.
	DBNodesShrink *string `json:"DBNodes,omitempty" xml:"DBNodes,omitempty"`
	// The type of specification change for the instance. Valid values:
	//
	// - **Up*	- (default): Upgrades a subscription instance, or upgrades or downgrades a pay-as-you-go instance.
	//
	// - **Down**: Downgrades a subscription instance.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The effective period for the new configurations. Valid values:
	//
	// > **Changing some configurations may affect the instance**. Read the [Impact section in the feature documentation](https://help.aliyun.com/document_detail/96061.html) carefully before you configure this parameter, and perform the operation during off-peak hours.
	//
	// 	- **Immediate*	- (default): Takes effect immediately.
	//
	// 	- **MaintainTime**: Takes effect within the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **ScheduleTime**: Specifies a time for the switchover. The time must be at least 12 hours later than the current time, and the actual switch time follows the formula: EffectiveTime = ScheduleTime + SwitchTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE) feature for Premium ESSD cloud disks](https://help.aliyun.com/document_detail/2527067.html). Valid values:
	//
	// -  **1**: Enabled.
	//
	// -  **0**: Disabled.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	// example:
	//
	// rpg-dp****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The parameters and their values. All parameter values are of the string type. You can call the DescribeParameterTemplates operation to query parameter names and values.
	//
	// > If you specify the **ParameterGroupId*	- parameter and a parameter is modified by both the ParameterGroupId and Parameters parameters, the modification in the Parameters parameter takes precedence.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// aliwood-1688-mobile-promotion
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies the execution time for modifying the parameter. This parameter is required when EffectiveTime is set to ScheduleTime. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// > The time must be later than the current time (the time when the API call is executed).
	//
	// example:
	//
	// 2019-10-17T18:50:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The [minor engine version](https://help.aliyun.com/document_detail/126002.html) of the PostgreSQL instance. When you change the instance type and an error indicates that the minor engine version is not supported, you must specify this parameter **to upgrade the minor engine version during the instance type change**.
	//
	// Format: `rds_postgres_<Major version number>00_<Minor version number>`. Example for version 12 on 20200830: `rds_postgres_1200_20200830`.
	//
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
}

func (s ModifyDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceShrinkRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceShrinkRequest) GetDBNodesShrink() *string {
	return s.DBNodesShrink
}

func (s *ModifyDBInstanceShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceShrinkRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBInstanceShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *ModifyDBInstanceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceShrinkRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceShrinkRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceShrinkRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetCategory(v string) *ModifyDBInstanceShrinkRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceShrinkRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDBInstanceClass(v string) *ModifyDBInstanceShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDBInstanceId(v string) *ModifyDBInstanceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDBNodesShrink(v string) *ModifyDBInstanceShrinkRequest {
	s.DBNodesShrink = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetDirection(v string) *ModifyDBInstanceShrinkRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetEffectiveTime(v string) *ModifyDBInstanceShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetOwnerAccount(v string) *ModifyDBInstanceShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetOwnerId(v int64) *ModifyDBInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetParameterGroupId(v string) *ModifyDBInstanceShrinkRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetParametersShrink(v string) *ModifyDBInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetPromotionCode(v string) *ModifyDBInstanceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetResourceGroupId(v string) *ModifyDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetSwitchTime(v string) *ModifyDBInstanceShrinkRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceShrinkRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
