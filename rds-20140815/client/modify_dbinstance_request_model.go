// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBInstanceRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceRequest
	GetColdDataEnabled() *bool
	SetDBInstanceClass(v string) *ModifyDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDBNodes(v []*ModifyDBInstanceRequestDBNodes) *ModifyDBInstanceRequest
	GetDBNodes() []*ModifyDBInstanceRequestDBNodes
	SetDirection(v string) *ModifyDBInstanceRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceRequest
	GetEffectiveTime() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceRequest
	GetIoAccelerationEnabled() *string
	SetOwnerAccount(v string) *ModifyDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBInstanceRequest
	GetParameterGroupId() *string
	SetParameters(v map[string]*string) *ModifyDBInstanceRequest
	GetParameters() map[string]*string
	SetPromotionCode(v string) *ModifyDBInstanceRequest
	GetPromotionCode() *string
	SetResourceGroupId(v string) *ModifyDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyDBInstanceRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceRequest
	GetTargetMinorVersion() *string
}

type ModifyDBInstanceRequest struct {
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- **true*	- (default): Automatically uses coupons.
	//
	// 	- **false**: Does not automatically use coupons.
	//
	// > After a coupon is used, the amount deducted by the coupon is not refunded if you downgrade the instance specifications.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Specifies whether to enable the [I/O burst feature for premium performance disks](https://help.aliyun.com/document_detail/2340501.html). Valid values:
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
	// <props="china">Specifies whether to enable the [cold data archiving feature](https://help.aliyun.com/document_detail/2701832.html) for general-purpose cloud disks. Valid values:
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
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The [target storage capacity](https://help.aliyun.com/document_detail/26312.html), in GB. You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available storage capacity range for the target instance type.
	//
	// > 	- You must specify at least one of this parameter and the **DBInstanceClass*	- parameter.
	//
	// > 	- You can call [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) to query the current storage capacity of the instance.
	//
	// example:
	//
	// 500
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **general_essd**: premium performance disk (recommended)
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
	DBNodes []*ModifyDBInstanceRequestDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The type of specification change. Valid values:
	//
	// - **Up*	- (default): Upgrades a subscription instance or upgrades/downgrades a pay-as-you-go instance.
	//
	// - **Down**: Downgrades a subscription instance.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The time when the new configurations take effect. Valid values:
	//
	// > **Changing some configurations may affect the instance**. Read the impact section in the [feature documentation](https://help.aliyun.com/document_detail/96061.html) before you configure this parameter. Perform the operation during off-peak hours.
	//
	// 	- **Immediate*	- (default): The new configurations take effect immediately.
	//
	// 	- **MaintainTime**: The new configurations take effect during the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **ScheduleTime**: The new configurations take effect at a specified time. The specified time must be at least 12 hours later than the current time. The actual switchover time follows the formula: EffectiveTime = ScheduleTime + SwitchTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE) feature](https://help.aliyun.com/document_detail/2527067.html) for premium performance disks. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
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
	// The parameters and their values. All parameter values are of the STRING type. You can call DescribeParameterTemplates to query parameter names and values.
	//
	// > If you specify the **ParameterGroupId*	- parameter and both the ParameterGroupId and Parameters parameters modify the same parameter, the modification specified by the Parameters parameter takes precedence.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// The scheduled time for executing the parameter modification. The EffectiveTime parameter must be set to ScheduleTime. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > The specified time must be later than the current time (the time when the call is made).
	//
	// example:
	//
	// 2019-10-17T18:50:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The [minor engine version](https://help.aliyun.com/document_detail/126002.html) of the PostgreSQL instance. If the specification change fails because the current minor engine version is not supported, specify the minor engine version to **upgrade the minor engine version during the specification change**.
	//
	// Format: `rds_postgres_<major version>00_<minor version>`. Example for version 12 with minor version 20200830: `rds_postgres_1200_20200830`.
	//
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
}

func (s ModifyDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceRequest) GetDBNodes() []*ModifyDBInstanceRequestDBNodes {
	return s.DBNodes
}

func (s *ModifyDBInstanceRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBInstanceRequest) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyDBInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetCategory(v string) *ModifyDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceClass(v string) *ModifyDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceId(v string) *ModifyDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBNodes(v []*ModifyDBInstanceRequestDBNodes) *ModifyDBInstanceRequest {
	s.DBNodes = v
	return s
}

func (s *ModifyDBInstanceRequest) SetDirection(v string) *ModifyDBInstanceRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetEffectiveTime(v string) *ModifyDBInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetOwnerAccount(v string) *ModifyDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetOwnerId(v int64) *ModifyDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetParameterGroupId(v string) *ModifyDBInstanceRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetParameters(v map[string]*string) *ModifyDBInstanceRequest {
	s.Parameters = v
	return s
}

func (s *ModifyDBInstanceRequest) SetPromotionCode(v string) *ModifyDBInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceGroupId(v string) *ModifyDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetSwitchTime(v string) *ModifyDBInstanceRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceRequest) Validate() error {
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBInstanceRequestDBNodes struct {
	// The unique identifier of the node, which is used to specify a node.
	//
	// > This parameter is valid only for Cluster Edition instances.
	//
	// example:
	//
	// 28542293
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **Master**: primary node.
	//
	// 	- **Slave**: secondary node.
	//
	// > For Cluster Edition instances, you can leave this parameter empty and specify the NodeId parameter to identify the node.
	//
	// example:
	//
	// Master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The vSwitch ID of the instance.
	//
	// example:
	//
	// vsw-bp1g7uym6ia6yroes6dkm
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shanghai-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyDBInstanceRequestDBNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceRequestDBNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceRequestDBNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBInstanceRequestDBNodes) GetRole() *string {
	return s.Role
}

func (s *ModifyDBInstanceRequestDBNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceRequestDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceRequestDBNodes) SetNodeId(v string) *ModifyDBInstanceRequestDBNodes {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetRole(v string) *ModifyDBInstanceRequestDBNodes {
	s.Role = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetVSwitchId(v string) *ModifyDBInstanceRequestDBNodes {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetZoneId(v string) *ModifyDBInstanceRequestDBNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) Validate() error {
	return dara.Validate(s)
}
