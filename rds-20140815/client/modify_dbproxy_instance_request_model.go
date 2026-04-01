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
	DBProxyNodes []*ModifyDBProxyInstanceRequestDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
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
	MigrateAZ []*ModifyDBProxyInstanceRequestMigrateAZ `json:"MigrateAZ,omitempty" xml:"MigrateAZ,omitempty" type:"Repeated"`
	OwnerId   *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The number of cpu cores for the node, valid values: **1*	- to **16**.
	//
	// >This parameter is required when selecting **DBProxyNodes**.
	//
	// example:
	//
	// 1
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The number of proxy nodes in the availability zone, valid values: **1*	- to **16**.
	//
	// >This parameter is required when selecting **DBProxyNodes**.
	//
	// example:
	//
	// 2
	NodeCounts *string `json:"nodeCounts,omitempty" xml:"nodeCounts,omitempty"`
	// The id of the availability zone where the node is located.
	//
	// >This parameter is required when selecting **DBProxyNodes**.
	//
	// example:
	//
	// cn-hagnzhou-c
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
	// The proxy connection address ID. You can obtain it through the DescribeDBProxyEndpoint interface.
	//
	// > This parameter is required when MigrateAZ is selected.
	//
	// example:
	//
	// yhw429********
	DbProxyEndpointId *string `json:"dbProxyEndpointId,omitempty" xml:"dbProxyEndpointId,omitempty"`
	// The target VSwitchId corresponding to the proxy instance migration.
	//
	// > This parameter is required when MigrateAZ is selected.
	//
	// example:
	//
	// vsw-sw0qq49d1m****
	DestVSwitchId *string `json:"destVSwitchId,omitempty" xml:"destVSwitchId,omitempty"`
	// The target vpc id corresponding to the proxy instance migration.
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
