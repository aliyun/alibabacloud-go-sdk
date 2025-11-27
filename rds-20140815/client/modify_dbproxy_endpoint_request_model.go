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
	// The consistency read timeout period. Unit: milliseconds. Default value: **10*	- Unit: milliseconds. Valid values: **0 to 60000**
	//
	// example:
	//
	// 10
	CausalConsistReadTimeout *string `json:"CausalConsistReadTimeout,omitempty" xml:"CausalConsistReadTimeout,omitempty"`
	// The capabilities that you want to enable for the proxy endpoint. If you specify more than one capability, separate the capabilities with semicolons (;). Format: `Capability 1:Status;Capability 2:Status;...`. Do not add a semicolon (;) at the end of the value.
	//
	// Valid capability values:
	//
	// 	- **ReadWriteSpliting**: read/write splitting
	//
	// 	- **ConnectionPersist**: connection pooling
	//
	// 	- **TransactionReadSqlRouteOptimizeStatus**: transaction splitting
	//
	// 	- **AZProximityAccess**: nearest access
	//
	// 	- **CausalConsistRead**: read consistency
	//
	// Valid status values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// >
	//
	// 	- If the instance runs PostgreSQL, you can enable only read/write splitting, which is specified by **ReadWriteSpliting**.
	//
	// 	- Nearest access is supported only by dedicated database proxies for RDS instances that run MySQL.
	//
	// example:
	//
	// ReadWriteSpliting:1;ConnectionPersist:0
	ConfigDBProxyFeatures *string `json:"ConfigDBProxyFeatures,omitempty" xml:"ConfigDBProxyFeatures,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp145737x5bi6****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the proxy endpoint. You can call the DescribeDBProxyEndpoint operation to query the proxy endpoint ID.
	//
	// > 	- If the instance runs MySQL and you set **DbEndpointOperator*	- to **Delete*	- or **Modify**, you must specify DBProxyEndpointId.
	//
	// > 	- If the instance runs PostgreSQL and you set **DbEndpointOperator*	- to **Delete**, **Modify**, or **Create**, you must specify DBProxyEndpointId.
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
	// The description of the proxy terminal.
	//
	// example:
	//
	// test-proxy
	DbEndpointAliases *string `json:"DbEndpointAliases,omitempty" xml:"DbEndpointAliases,omitempty"`
	// The minimum number of reserved instances.
	//
	// example:
	//
	// 2
	DbEndpointMinSlaveCount *string `json:"DbEndpointMinSlaveCount,omitempty" xml:"DbEndpointMinSlaveCount,omitempty"`
	// The type of operation that you want to perform. Valid values:
	//
	// 	- **Modify**: Modify a proxy terminal. This is the default value.
	//
	// 	- **Create**: Create a proxy terminal.
	//
	// 	- **Delete**: Delete a proxy terminal.
	//
	// example:
	//
	// Modify
	DbEndpointOperator *string `json:"DbEndpointOperator,omitempty" xml:"DbEndpointOperator,omitempty"`
	// The read and write attributes of the proxy terminal. Valid values:
	//
	// 	- **ReadWrite**: The proxy terminal connects to the primary instance and can receive both read and write requests.
	//
	// 	- **ReadOnly**: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
	//
	// > 	- If you set **DbEndpointOperator*	- to **Create**, you must also specify DbEndpointReadWriteMode.
	//
	// > 	- If the instance runs MySQL and you change the value of this parameter from **ReadWrite*	- to **ReadOnly**, the transaction splitting feature is disabled.
	//
	// example:
	//
	// ReadWrite
	DbEndpointReadWriteMode *string `json:"DbEndpointReadWriteMode,omitempty" xml:"DbEndpointReadWriteMode,omitempty"`
	// The type of the proxy terminal. This is a reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// RWSplit
	DbEndpointType *string `json:"DbEndpointType,omitempty" xml:"DbEndpointType,omitempty"`
	// The point in time that you want to specify. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If **EffectiveTime*	- is set to **SpecificTime**, you must specify this parameter.
	//
	// example:
	//
	// 2023-05-06T07:08:09Z
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
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy that is used to allocate read weights. Valid values:
	//
	// 	- **Standard*	- (default): The system automatically assigns read weights to the primary and read-only instances based on the specifications of these instances.
	//
	// 	- **Custom**: You must manually allocate read weights to the primary and read-only instances.
	//
	// >  You must specify this parameter when read/write splitting is enabled. For more information about the permission allocation policy, see [Modify the latency threshold and read weights of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/96076.html) and [Enable and configure the database proxy feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418272.html).
	//
	// example:
	//
	// Standard
	ReadOnlyInstanceDistributionType *string `json:"ReadOnlyInstanceDistributionType,omitempty" xml:"ReadOnlyInstanceDistributionType,omitempty"`
	// The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, the system no longer forwards read requests to the read-only instance. If you do not specify this parameter, the original value of this parameter is retained. Valid values: **0*	- to **3600**.
	//
	// >
	//
	// 	- You must specify this parameter only when read/write splitting is enabled.
	//
	// 	- If the database proxy endpoint has the read and write attributes, the default value of this parameter is **30*	- and read/write splitting is supported. If the database proxy endpoint has the read-only attribute, the default value of this parameter is **-1*	- and read/write splitting is not supported. Unit: seconds.
	//
	// example:
	//
	// 30
	ReadOnlyInstanceMaxDelayTime *string `json:"ReadOnlyInstanceMaxDelayTime,omitempty" xml:"ReadOnlyInstanceMaxDelayTime,omitempty"`
	// The read weights of the instance and its read-only instances. A read weight must be a multiple of 100 and cannot exceed 10000. Formats:
	//
	// 	- Standard instance: `{"ID of the primary instance":"Weight","ID of the read-only instance":"Weight"...}`
	//
	//     Example: `{"rm-uf6wjk5****":"500","rr-tfhfgk5xxx":"200"...}`
	//
	// 	- Instance on RDS Cluster Edition: `{"ID of the read-only instance":"Weight","DBClusterNode":{"ID of the primary node":"Weight","ID of the secondary node":"Weight","ID of the secondary node":"Weight"...}}`
	//
	//     Example: `{"rr-tfhfgk5****":"200","DBClusterNode":{"rn-2z****":"0","rn-2z****":"400","rn-2z****":"400"...}}`
	//
	//     > **DBClusterNode*	- is required if the instance runs RDS Cluster Edition. The DBClusterNode parameter includes information about **IDs*	- and **weights*	- of the primary and secondary nodes..
	//
	// example:
	//
	// {"rm-uf6wjk5xxxx":"500","rr-tfhfgk5xxx":"200"...}
	ReadOnlyInstanceWeight *string `json:"ReadOnlyInstanceWeight,omitempty" xml:"ReadOnlyInstanceWeight,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch in the zone in which the proxy endpoint is specified. The default value is the ID of the vSwitch that corresponds to the default terminal of the database proxy. You can call the DescribeVSwitches operation to query existing vSwitches.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the zone in which the proxy endpoint is specified. The default value is the VPC ID that corresponds to the default terminal of the database proxy. You can call the DescribeDBInstanceAttribute operation to query the default VPC of an instance.
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
