// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodesShrink(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyNodesShrink() *string
	SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceShrinkRequest
	GetEffectiveSpecificTime() *string
	SetEffectiveTime(v string) *ModifyDBProxyInstanceShrinkRequest
	GetEffectiveTime() *string
	SetMigrateAZShrink(v string) *ModifyDBProxyInstanceShrinkRequest
	GetMigrateAZShrink() *string
	SetOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBProxyInstanceShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetVSwitchIds(v string) *ModifyDBProxyInstanceShrinkRequest
	GetVSwitchIds() *string
}

type ModifyDBProxyInstanceShrinkRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The number of database proxies. If you set this parameter to 0, the database proxy feature is disabled for the instance. Valid values: **1*	- to **16**.
	//
	// >  The capability of the database proxy feature to process requests increases with the number of database proxies that are enabled. You can monitor the load on the instance and specify an appropriate number of database proxies based on the load monitoring data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// The database proxy type. Valid values:
	//
	// 	- **common**: general-purpose database proxy
	//
	// 	- **exclusive*	- (default): dedicated database proxy
	//
	// This parameter is required.
	//
	// example:
	//
	// DedicatedProxy
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// List of proxy nodes.
	//
	// > This parameter must be passed when the current proxy instance is deployed in multiple availability zones.
	DBProxyNodesShrink *string `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty"`
	// The point in time that you want to specify. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If the **EffectiveTime*	- parameter is set to **SpecificTime**, you must specify this parameter.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	EffectiveSpecificTime *string `json:"EffectiveSpecificTime,omitempty" xml:"EffectiveSpecificTime,omitempty"`
	// The effective time. Valid values:
	//
	// 	- **Immediate**: The effective time is immediate.
	//
	// 	- **MaintainTime**: The effective time is within the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecificTime**: The effective time is a specified point in time.
	//
	// Default value: **MaintainTime**.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The list of available zones for migration agents.
	//
	// > Currently, only RDS MySQL cloud disk version agent instance migration is supported.
	MigrateAZShrink *string `json:"MigrateAZ,omitempty" xml:"MigrateAZ,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch in the destination zone. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/610431.html) operation to query existing vSwitches.
	//
	// >  Only database proxies for ApsaraDB RDS for MySQL instances that use cloud disks can be migrated to different zones.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
}

func (s ModifyDBProxyInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyNodesShrink() *string {
	return s.DBProxyNodesShrink
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetEffectiveSpecificTime() *string {
	return s.EffectiveSpecificTime
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetMigrateAZShrink() *string {
	return s.MigrateAZShrink
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBInstanceId(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyEngineType(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyNodesShrink(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyNodesShrink = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.EffectiveSpecificTime = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetEffectiveTime(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetMigrateAZShrink(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.MigrateAZShrink = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetRegionId(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetVSwitchIds(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
