// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyInstanceRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodes(v []*ModifyDBProxyInstanceRequestDBProxyNodes) *ModifyDBProxyInstanceRequest
	GetDBProxyNodes() []*ModifyDBProxyInstanceRequestDBProxyNodes
	SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceRequest
	GetEffectiveSpecificTime() *string
	SetEffectiveTime(v string) *ModifyDBProxyInstanceRequest
	GetEffectiveTime() *string
	SetMigrateAZ(v []*ModifyDBProxyInstanceRequestMigrateAZ) *ModifyDBProxyInstanceRequest
	GetMigrateAZ() []*ModifyDBProxyInstanceRequestMigrateAZ
	SetOwnerId(v int64) *ModifyDBProxyInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBProxyInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyInstanceRequest
	GetResourceOwnerId() *int64
	SetVSwitchIds(v string) *ModifyDBProxyInstanceRequest
	GetVSwitchIds() *string
}

type ModifyDBProxyInstanceRequest struct {
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The number of proxy instances. If this parameter is set to 0, the proxy service of this type is disabled for the instance. Valid values: **1*	- to **16**.
	//
	// > More proxy instances can handle more requests. You can check the load of proxy instances based on monitoring data and then specify an appropriate number of proxy instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// The type of the database proxy instance. Valid values:
	//
	// - **common**: general-purpose database proxy
	//
	// - **exclusive**: dedicated database proxy (default)
	//
	// This parameter is required.
	//
	// example:
	//
	// exclusive
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// The list of proxy nodes.
	//
	// > This parameter is required when the current proxy instance uses multi-active zone deployment.
	DBProxyNodes []*ModifyDBProxyInstanceRequestDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
	// The specified time for the modification to take effect. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > This parameter is required when **EffectiveTime*	- is set to **SpecificTime**.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	EffectiveSpecificTime *string `json:"EffectiveSpecificTime,omitempty" xml:"EffectiveSpecificTime,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate**: The modification takes effect immediately.
	//
	// 	- **MaintainTime**: The modification takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecificTime**: The modification takes effect at a specified time.
	//
	// Default value: **MaintainTime**.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The list of active zones for proxy migration.
	//
	// > Currently, only ApsaraDB RDS for MySQL proxy instances with cloud disks support active zone migration.
	MigrateAZ []*ModifyDBProxyInstanceRequestMigrateAZ `json:"MigrateAZ,omitempty" xml:"MigrateAZ,omitempty" type:"Repeated"`
	OwnerId   *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call DescribeRegions to obtain the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
}

func (s ModifyDBProxyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyNodes() []*ModifyDBProxyInstanceRequestDBProxyNodes {
	return s.DBProxyNodes
}

func (s *ModifyDBProxyInstanceRequest) GetEffectiveSpecificTime() *string {
	return s.EffectiveSpecificTime
}

func (s *ModifyDBProxyInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBProxyInstanceRequest) GetMigrateAZ() []*ModifyDBProxyInstanceRequestMigrateAZ {
	return s.MigrateAZ
}

func (s *ModifyDBProxyInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyInstanceRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ModifyDBProxyInstanceRequest) SetDBInstanceId(v string) *ModifyDBProxyInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyEngineType(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyNodes(v []*ModifyDBProxyInstanceRequestDBProxyNodes) *ModifyDBProxyInstanceRequest {
	s.DBProxyNodes = v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceRequest {
	s.EffectiveSpecificTime = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetEffectiveTime(v string) *ModifyDBProxyInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetMigrateAZ(v []*ModifyDBProxyInstanceRequestMigrateAZ) *ModifyDBProxyInstanceRequest {
	s.MigrateAZ = v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetOwnerId(v int64) *ModifyDBProxyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetRegionId(v string) *ModifyDBProxyInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetResourceOwnerId(v int64) *ModifyDBProxyInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetVSwitchIds(v string) *ModifyDBProxyInstanceRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) Validate() error {
	if s.DBProxyNodes != nil {
		for _, item := range s.DBProxyNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MigrateAZ != nil {
		for _, item := range s.MigrateAZ {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBProxyInstanceRequestDBProxyNodes struct {
	// The number of CPU cores for the node. Valid values: **1*	- to **16**.
	//
	// > This parameter is required when **DBProxyNodes*	- is specified.
	//
	// example:
	//
	// 1
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The number of proxy nodes in the zone. Valid values: **1*	- to **2**.
	//
	// > This parameter is required when **DBProxyNodes*	- is specified.
	//
	// example:
	//
	// 2
	NodeCounts *string `json:"nodeCounts,omitempty" xml:"nodeCounts,omitempty"`
	// The zone ID of the node.
	//
	// > This parameter is required when **DBProxyNodes*	- is specified.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ModifyDBProxyInstanceRequestDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequestDBProxyNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetNodeCounts() *string {
	return s.NodeCounts
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetCpuCores(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetNodeCounts(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.NodeCounts = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetZoneId(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) Validate() error {
	return dara.Validate(s)
}

type ModifyDBProxyInstanceRequestMigrateAZ struct {
	// The proxy endpoint ID. You can call DescribeDBProxyEndpoint to obtain the proxy endpoint ID.
	//
	// > This parameter is required when **MigrateAZ*	- is specified.
	//
	// example:
	//
	// yhw429********
	DbProxyEndpointId *string `json:"dbProxyEndpointId,omitempty" xml:"dbProxyEndpointId,omitempty"`
	// The ID of the destination vSwitch for the proxy instance migration.
	//
	// > This parameter is required when **MigrateAZ*	- is specified.
	//
	// example:
	//
	// vsw-sw0qq49d1m****
	DestVSwitchId *string `json:"destVSwitchId,omitempty" xml:"destVSwitchId,omitempty"`
	// The ID of the destination VPC for the proxy instance migration.
	//
	// > This parameter is required when **MigrateAZ*	- is specified.
	//
	// example:
	//
	// vpc-2vcicu73rdylp****
	DestVpcId *string `json:"destVpcId,omitempty" xml:"destVpcId,omitempty"`
}

func (s ModifyDBProxyInstanceRequestMigrateAZ) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequestMigrateAZ) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDbProxyEndpointId() *string {
	return s.DbProxyEndpointId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDestVSwitchId() *string {
	return s.DestVSwitchId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDestVpcId() *string {
	return s.DestVpcId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDbProxyEndpointId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DbProxyEndpointId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDestVSwitchId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DestVSwitchId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDestVpcId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DestVpcId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) Validate() error {
	return dara.Validate(s)
}
