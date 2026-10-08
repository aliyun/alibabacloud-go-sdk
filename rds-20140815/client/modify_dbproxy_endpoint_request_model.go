// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCausalConsistReadTimeout(v string) *ModifyDBProxyEndpointRequest
	GetCausalConsistReadTimeout() *string
	SetConfigDBProxyFeatures(v string) *ModifyDBProxyEndpointRequest
	GetConfigDBProxyFeatures() *string
	SetDBInstanceId(v string) *ModifyDBProxyEndpointRequest
	GetDBInstanceId() *string
	SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyEndpointRequest
	GetDBProxyEngineType() *string
	SetDbEndpointAliases(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointAliases() *string
	SetDbEndpointCostThresholdForDuckdb(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointCostThresholdForDuckdb() *string
	SetDbEndpointMinSlaveCount(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointMinSlaveCount() *string
	SetDbEndpointOperator(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointOperator() *string
	SetDbEndpointReadWriteMode(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointReadWriteMode() *string
	SetDbEndpointType(v string) *ModifyDBProxyEndpointRequest
	GetDbEndpointType() *string
	SetEffectiveSpecificTime(v string) *ModifyDBProxyEndpointRequest
	GetEffectiveSpecificTime() *string
	SetEffectiveTime(v string) *ModifyDBProxyEndpointRequest
	GetEffectiveTime() *string
	SetOwnerId(v int64) *ModifyDBProxyEndpointRequest
	GetOwnerId() *int64
	SetReadOnlyInstanceDistributionType(v string) *ModifyDBProxyEndpointRequest
	GetReadOnlyInstanceDistributionType() *string
	SetReadOnlyInstanceMaxDelayTime(v string) *ModifyDBProxyEndpointRequest
	GetReadOnlyInstanceMaxDelayTime() *string
	SetReadOnlyInstanceWeight(v string) *ModifyDBProxyEndpointRequest
	GetReadOnlyInstanceWeight() *string
	SetRegionId(v string) *ModifyDBProxyEndpointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyEndpointRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ModifyDBProxyEndpointRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyDBProxyEndpointRequest
	GetVpcId() *string
}

type ModifyDBProxyEndpointRequest struct {
	// The timeout period for read consistency. Unit: milliseconds. Default value: **10**. Valid values: **0 to 60000**.
	//
	// example:
	//
	// 10
	CausalConsistReadTimeout *string `json:"CausalConsistReadTimeout,omitempty" xml:"CausalConsistReadTimeout,omitempty"`
	// The proxy features that you want to enable for the proxy endpoint. Separate multiple features with semicolons (;). Format: `Feature 1:Status;Feature 2:Status;...`. Do not add a semicolon (;) at the end.
	//
	// Valid values for features:
	//
	// 	- **ReadWriteSpliting**: Read/write splitting.
	//
	// 	- **ConnectionPersist**: Connection pool.
	//
	// 	- **TransactionReadSqlRouteOptimizeStatus**: Transaction splitting.
	//
	// 	- **AZProximityAccess**: Nearest access.
	//
	// 	- **CausalConsistRead**: Read consistency.
	//
	// 	- **HtapFilter**: HTAP automatic request distribution among row store and column store nodes.
	//
	// Valid values for status:
	//
	// 	- **1**: Enabled.
	//
	// 	- **0**: Disabled.
	//
	// > - ApsaraDB RDS for PostgreSQL supports only **ReadWriteSpliting**.
	//
	// > - The nearest access feature is supported only by the dedicated database proxy for MySQL.
	//
	// example:
	//
	// ReadWriteSpliting:1;ConnectionPersist:0
	ConfigDBProxyFeatures *string `json:"ConfigDBProxyFeatures,omitempty" xml:"ConfigDBProxyFeatures,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp145737x5bi6****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the proxy endpoint. You can call DescribeDBProxyEndpoint to query the ID.
	//
	// > - MySQL: This parameter is required when **DbEndpointOperator*	- is set to **Delete*	- or **Modify**.
	//
	// > - PostgreSQL: This parameter is required when **DbEndpointOperator*	- is set to **Delete**, **Modify**, or **Create**.
	//
	// example:
	//
	// gos787jog2wk0y****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The description of the proxy endpoint.
	//
	// example:
	//
	// test-proxy
	DbEndpointAliases                *string `json:"DbEndpointAliases,omitempty" xml:"DbEndpointAliases,omitempty"`
	DbEndpointCostThresholdForDuckdb *string `json:"DbEndpointCostThresholdForDuckdb,omitempty" xml:"DbEndpointCostThresholdForDuckdb,omitempty"`
	// The minimum number of reserved instances.
	//
	// example:
	//
	// 2
	DbEndpointMinSlaveCount *string `json:"DbEndpointMinSlaveCount,omitempty" xml:"DbEndpointMinSlaveCount,omitempty"`
	// The type of operation. Valid values:
	//
	// 	- **Modify**: The default value. Modifies the proxy endpoint.
	//
	// 	- **Create**: Creates a proxy endpoint.
	//
	// 	- **Delete**: Deletes a proxy endpoint.
	//
	// example:
	//
	// Modify
	DbEndpointOperator *string `json:"DbEndpointOperator,omitempty" xml:"DbEndpointOperator,omitempty"`
	// The read/write mode. Valid values:
	//
	// 	- **ReadWrite**: Connects to the primary instance and can accept write requests.
	//
	// 	- **ReadOnly**: The default value. Does not connect to the primary instance and cannot accept write requests.
	//
	// > 	- This parameter is required when **DbEndpointOperator*	- is set to **Create**.
	//
	// > 	- For ApsaraDB RDS for MySQL instances, if you change this parameter from **ReadWrite*	- to **ReadOnly**, the transaction splitting feature is disabled.
	//
	// example:
	//
	// ReadWrite
	DbEndpointReadWriteMode *string `json:"DbEndpointReadWriteMode,omitempty" xml:"DbEndpointReadWriteMode,omitempty"`
	// The type of the proxy endpoint. This is a reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// RWSplit
	DbEndpointType *string `json:"DbEndpointType,omitempty" xml:"DbEndpointType,omitempty"`
	// The specified time at which the change takes effect. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > This parameter is required when **EffectiveTime*	- is set to **SpecificTime**.
	//
	// example:
	//
	// 2023-05-06T07:08:09Z
	EffectiveSpecificTime *string `json:"EffectiveSpecificTime,omitempty" xml:"EffectiveSpecificTime,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate**: The change takes effect immediately.
	//
	// 	- **MaintainTime**: The change takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecificTime**: The change takes effect at a specified time.
	//
	// Default value: **MaintainTime**.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mode used to allocate read weights. Valid values:
	//
	// 	- **Standard**: The default value. Read weights are automatically allocated based on instance specifications.
	//
	// 	- **Custom**: Custom read weights.
	//
	// > This parameter is required only when read/write splitting is enabled. For more information about read weight allocation, see [Read weight allocation](https://help.aliyun.com/document_detail/96076.html) for MySQL and [Enable and configure the database proxy service](https://help.aliyun.com/document_detail/418272.html) for PostgreSQL.
	//
	// example:
	//
	// Standard
	ReadOnlyInstanceDistributionType *string `json:"ReadOnlyInstanceDistributionType,omitempty" xml:"ReadOnlyInstanceDistributionType,omitempty"`
	// The maximum latency threshold for read-only instances in read/write splitting. If the latency of a read-only instance exceeds this value, read traffic is not routed to the instance. Unit: seconds. If you do not specify this parameter, the current value is retained. Valid values: **0*	- to **3600**.
	//
	// >- This parameter is required only when read/write splitting is enabled.
	//
	// >- Default value: **30*	- seconds when the read/write mode is set to read/write (read/write splitting), and **-1*	- (disabled) when the read/write mode is set to read-only.
	//
	// example:
	//
	// 30
	ReadOnlyInstanceMaxDelayTime *string `json:"ReadOnlyInstanceMaxDelayTime,omitempty" xml:"ReadOnlyInstanceMaxDelayTime,omitempty"`
	// The custom read weights to allocate to the primary instance and read-only instances. The value must be in increments of 100. Maximum value: 10000. Format:
	//
	// - Regular instance: `{"PrimaryInstanceID":"Weight","ReadOnlyInstanceID":"Weight"...}`
	//
	//     Example: `{"rm-uf6wjk5****":"500","rr-tfhfgk5xxx":"200"...}`
	//
	// - ApsaraDB RDS for MySQL cluster instance: `{"ReadOnlyInstanceID":"Weight","DBClusterNode":{"PrimaryNodeID":"Weight","SecondaryNodeID":"Weight","SecondaryNodeID":"Weight"...}}`
	//
	//     Example: `{"rr-tfhfgk5****":"200","DBClusterNode":{"rn-2z****":"0","rn-2z****":"400","rn-2z****":"400"...}}`
	//
	//     > **DBClusterNode*	- is a request parameter specific to cluster instances. It contains the **NodeID*	- and **Weight*	- of the primary and secondary nodes.
	//
	// example:
	//
	// {"rm-uf6wjk5****":"500","rr-tfhfgk5xxx":"200"...}
	ReadOnlyInstanceWeight *string `json:"ReadOnlyInstanceWeight,omitempty" xml:"ReadOnlyInstanceWeight,omitempty"`
	// The region ID. You can call DescribeRegions to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitch ID that corresponds to the zone of the proxy endpoint. Default value: the vSwitch ID of the default endpoint of the proxy instance. You can call DescribeVSwitches to query available vSwitches.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID that corresponds to the zone of the proxy endpoint. Default value: the VPC ID of the default endpoint of the proxy instance. You can call DescribeDBInstanceAttribute to query the default VPC of the instance.
	//
	// example:
	//
	// vpc-2zeusejj******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyDBProxyEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointRequest) GetCausalConsistReadTimeout() *string {
	return s.CausalConsistReadTimeout
}

func (s *ModifyDBProxyEndpointRequest) GetConfigDBProxyFeatures() *string {
	return s.ConfigDBProxyFeatures
}

func (s *ModifyDBProxyEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyEndpointRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *ModifyDBProxyEndpointRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointAliases() *string {
	return s.DbEndpointAliases
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointCostThresholdForDuckdb() *string {
	return s.DbEndpointCostThresholdForDuckdb
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointMinSlaveCount() *string {
	return s.DbEndpointMinSlaveCount
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointOperator() *string {
	return s.DbEndpointOperator
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointReadWriteMode() *string {
	return s.DbEndpointReadWriteMode
}

func (s *ModifyDBProxyEndpointRequest) GetDbEndpointType() *string {
	return s.DbEndpointType
}

func (s *ModifyDBProxyEndpointRequest) GetEffectiveSpecificTime() *string {
	return s.EffectiveSpecificTime
}

func (s *ModifyDBProxyEndpointRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBProxyEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyEndpointRequest) GetReadOnlyInstanceDistributionType() *string {
	return s.ReadOnlyInstanceDistributionType
}

func (s *ModifyDBProxyEndpointRequest) GetReadOnlyInstanceMaxDelayTime() *string {
	return s.ReadOnlyInstanceMaxDelayTime
}

func (s *ModifyDBProxyEndpointRequest) GetReadOnlyInstanceWeight() *string {
	return s.ReadOnlyInstanceWeight
}

func (s *ModifyDBProxyEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyEndpointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBProxyEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyDBProxyEndpointRequest) SetCausalConsistReadTimeout(v string) *ModifyDBProxyEndpointRequest {
	s.CausalConsistReadTimeout = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetConfigDBProxyFeatures(v string) *ModifyDBProxyEndpointRequest {
	s.ConfigDBProxyFeatures = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDBInstanceId(v string) *ModifyDBProxyEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDBProxyEngineType(v string) *ModifyDBProxyEndpointRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointAliases(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointAliases = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointCostThresholdForDuckdb(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointCostThresholdForDuckdb = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointMinSlaveCount(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointMinSlaveCount = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointOperator(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointOperator = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointReadWriteMode(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointReadWriteMode = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetDbEndpointType(v string) *ModifyDBProxyEndpointRequest {
	s.DbEndpointType = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetEffectiveSpecificTime(v string) *ModifyDBProxyEndpointRequest {
	s.EffectiveSpecificTime = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetEffectiveTime(v string) *ModifyDBProxyEndpointRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetOwnerId(v int64) *ModifyDBProxyEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetReadOnlyInstanceDistributionType(v string) *ModifyDBProxyEndpointRequest {
	s.ReadOnlyInstanceDistributionType = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetReadOnlyInstanceMaxDelayTime(v string) *ModifyDBProxyEndpointRequest {
	s.ReadOnlyInstanceMaxDelayTime = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetReadOnlyInstanceWeight(v string) *ModifyDBProxyEndpointRequest {
	s.ReadOnlyInstanceWeight = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetRegionId(v string) *ModifyDBProxyEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBProxyEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetVSwitchId(v string) *ModifyDBProxyEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) SetVpcId(v string) *ModifyDBProxyEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDBProxyEndpointRequest) Validate() error {
	return dara.Validate(s)
}
