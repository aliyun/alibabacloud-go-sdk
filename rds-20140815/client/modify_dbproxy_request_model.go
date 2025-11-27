// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigDBProxyService(v string) *ModifyDBProxyRequest
	GetConfigDBProxyService() *string
	SetDBInstanceId(v string) *ModifyDBProxyRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodes(v []*ModifyDBProxyRequestDBProxyNodes) *ModifyDBProxyRequest
	GetDBProxyNodes() []*ModifyDBProxyRequestDBProxyNodes
	SetInstanceNetworkType(v string) *ModifyDBProxyRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *ModifyDBProxyRequest
	GetOwnerId() *int64
	SetPersistentConnectionStatus(v string) *ModifyDBProxyRequest
	GetPersistentConnectionStatus() *string
	SetRegionId(v string) *ModifyDBProxyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDBProxyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBProxyRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBProxyRequest
	GetVSwitchId() *string
}

type ModifyDBProxyRequest struct {
	// Specifies whether to enable or disable the database proxy feature. Valid values:
	//
	// 	- **Startup**: enables the feature.
	//
	// 	- **Shutdown**: disables the feature.
	//
	// 	- **Modify**: modifies the configuration of the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Startup
	ConfigDBProxyService *string `json:"ConfigDBProxyService,omitempty" xml:"ConfigDBProxyService,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The number of proxy instances that are enabled. Valid values: **1*	- to **16**. Default value: **1**.
	//
	// >  The capability of the database proxy to process requests increases with the number of proxy instances that are enabled. You can monitor the load on the instance and specify an appropriate number of proxy instances based on the load monitoring data.
	//
	// example:
	//
	// 1
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// The database proxy type. Valid values:
	//
	// 	- **common**: general-purpose database proxy
	//
	// 	- **exclusive*	- (default): dedicated database proxy
	//
	// example:
	//
	// common
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// The proxy nodes.
	DBProxyNodes []*ModifyDBProxyRequestDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
	// The network type of the instance. Only the VPC network type is supported. Set the value to **VPC**.
	//
	// >  If you enable the database proxy feature for the instance, you must specify this parameter.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable persistent connections. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// >
	//
	// 	- This parameter is available only for instances that run MySQL.
	//
	// 	- If you want to modify persistent connections, you must set the **ConfigDBProxyService*	- parameter to **Modify**.
	//
	// example:
	//
	// Enabled
	PersistentConnectionStatus *string `json:"PersistentConnectionStatus,omitempty" xml:"PersistentConnectionStatus,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the instance belongs. You can call the DescribeDBInstanceAttribute operation to query the ID.
	//
	// >  If you enable the database proxy feature for the instance, you must specify this parameter.
	//
	// example:
	//
	// vpc-xxxxxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch to which the instance belongs. You can call the DescribeDBInstanceAttribute operation to query the ID.
	//
	// >  If you enable the database proxy feature for the instance, you must specify this parameter.
	//
	// example:
	//
	// vsw-xxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyRequest) GetConfigDBProxyService() *string {
	return s.ConfigDBProxyService
}

func (s *ModifyDBProxyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyRequest) GetDBProxyNodes() []*ModifyDBProxyRequestDBProxyNodes {
	return s.DBProxyNodes
}

func (s *ModifyDBProxyRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyRequest) GetPersistentConnectionStatus() *string {
	return s.PersistentConnectionStatus
}

func (s *ModifyDBProxyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBProxyRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBProxyRequest) SetConfigDBProxyService(v string) *ModifyDBProxyRequest {
	s.ConfigDBProxyService = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBInstanceId(v string) *ModifyDBProxyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyEngineType(v string) *ModifyDBProxyRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyNodes(v []*ModifyDBProxyRequestDBProxyNodes) *ModifyDBProxyRequest {
	s.DBProxyNodes = v
	return s
}

func (s *ModifyDBProxyRequest) SetInstanceNetworkType(v string) *ModifyDBProxyRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetOwnerId(v int64) *ModifyDBProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetPersistentConnectionStatus(v string) *ModifyDBProxyRequest {
	s.PersistentConnectionStatus = &v
	return s
}

func (s *ModifyDBProxyRequest) SetRegionId(v string) *ModifyDBProxyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceGroupId(v string) *ModifyDBProxyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceOwnerId(v int64) *ModifyDBProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetVPCId(v string) *ModifyDBProxyRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetVSwitchId(v string) *ModifyDBProxyRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBProxyRequest) Validate() error {
	if s.DBProxyNodes != nil {
		for _, item := range s.DBProxyNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBProxyRequestDBProxyNodes struct {
	// The number of CPU cores of the node. Valid values: **1*	- to **16**.
	//
	// >  This parameter is required when you configure the **DBProxyNodes*	- parameter.
	//
	// example:
	//
	// 1
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The number of proxy nodes in the zone. Valid values: **1*	- and **2**.
	//
	// >  This parameter is required when you configure the **DBProxyNodes*	- parameter.
	//
	// example:
	//
	// 2
	NodeCounts *string `json:"nodeCounts,omitempty" xml:"nodeCounts,omitempty"`
	// The ID of the zone in which the node resides.
	//
	// >  This parameter is required when you configure the **DBProxyNodes*	- parameter.
	//
	// example:
	//
	// cn-hagnzhou-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ModifyDBProxyRequestDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyRequestDBProxyNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetNodeCounts() *string {
	return s.NodeCounts
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetCpuCores(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetNodeCounts(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.NodeCounts = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetZoneId(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) Validate() error {
	return dara.Validate(s)
}
