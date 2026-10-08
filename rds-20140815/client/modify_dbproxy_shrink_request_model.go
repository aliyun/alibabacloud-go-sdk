// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigDBProxyService(v string) *ModifyDBProxyShrinkRequest
	GetConfigDBProxyService() *string
	SetDBInstanceId(v string) *ModifyDBProxyShrinkRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodesShrink(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyNodesShrink() *string
	SetInstanceNetworkType(v string) *ModifyDBProxyShrinkRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *ModifyDBProxyShrinkRequest
	GetOwnerId() *int64
	SetPersistentConnectionStatus(v string) *ModifyDBProxyShrinkRequest
	GetPersistentConnectionStatus() *string
	SetRegionId(v string) *ModifyDBProxyShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDBProxyShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyShrinkRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBProxyShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBProxyShrinkRequest
	GetVSwitchId() *string
}

type ModifyDBProxyShrinkRequest struct {
	// Specifies whether to enable, disable, or modify the database proxy. Valid values:
	//
	// 	- **Startup**: Enables the database proxy.
	//
	// 	- **Shutdown**: Disables the database proxy.
	//
	// 	- **Modify**: Modifies the database proxy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Startup
	ConfigDBProxyService *string `json:"ConfigDBProxyService,omitempty" xml:"ConfigDBProxyService,omitempty"`
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The number of proxy instances. Valid values: **1*	- to **16**. Default value: **1**.
	//
	// > More proxy instances can handle more requests. You can check the monitoring data to understand the load on proxy instances and then set an appropriate number of proxy instances.
	//
	// example:
	//
	// 1
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// The type of the database proxy instance. Valid values:
	//
	// - **common**: general-purpose database proxy
	//
	// - **exclusive**: dedicated database proxy (default)
	//
	// example:
	//
	// exclusive
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// The list of proxy nodes.
	DBProxyNodesShrink *string `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty"`
	// The network type of the instance. Only Virtual Private Cloud (VPC) is supported. Set the value to **VPC**.
	//
	// > This parameter is required when you enable the database proxy.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable persistent connections. Valid values:
	//
	// - **Enabled**: enables persistent connections.
	//
	// - **Disabled**: disables persistent connections.
	//
	// > - Only RDS MySQL supports this parameter.
	//
	// > - To modify the persistent connection status, set **ConfigDBProxyService*	- to **Modify**.
	//
	// example:
	//
	// Disabled
	PersistentConnectionStatus *string `json:"PersistentConnectionStatus,omitempty" xml:"PersistentConnectionStatus,omitempty"`
	// The region ID. You can call DescribeRegions to obtain the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The VPC ID of the instance. You can call DescribeDBInstanceAttribute to obtain the VPC ID.
	//
	// > This parameter is required when you enable the database proxy.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance. You can call DescribeDBInstanceAttribute to obtain the vSwitch ID.
	//
	// > This parameter is required when you enable the database proxy.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBProxyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyShrinkRequest) GetConfigDBProxyService() *string {
	return s.ConfigDBProxyService
}

func (s *ModifyDBProxyShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyNodesShrink() *string {
	return s.DBProxyNodesShrink
}

func (s *ModifyDBProxyShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBProxyShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyShrinkRequest) GetPersistentConnectionStatus() *string {
	return s.PersistentConnectionStatus
}

func (s *ModifyDBProxyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBProxyShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBProxyShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBProxyShrinkRequest) SetConfigDBProxyService(v string) *ModifyDBProxyShrinkRequest {
	s.ConfigDBProxyService = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBInstanceId(v string) *ModifyDBProxyShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyEngineType(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyNodesShrink(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyNodesShrink = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetInstanceNetworkType(v string) *ModifyDBProxyShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetOwnerId(v int64) *ModifyDBProxyShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetPersistentConnectionStatus(v string) *ModifyDBProxyShrinkRequest {
	s.PersistentConnectionStatus = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetRegionId(v string) *ModifyDBProxyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceGroupId(v string) *ModifyDBProxyShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBProxyShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetVPCId(v string) *ModifyDBProxyShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetVSwitchId(v string) *ModifyDBProxyShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
