// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddVPCIPs(v string) *CreateDBInstanceRequest
	GetAddVPCIPs() *string
	SetCacheSize(v int32) *CreateDBInstanceRequest
	GetCacheSize() *int32
	SetChargeType(v string) *CreateDBInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetClusterNodeCount(v int32) *CreateDBInstanceRequest
	GetClusterNodeCount() *int32
	SetClusterNodeType(v string) *CreateDBInstanceRequest
	GetClusterNodeType() *string
	SetConnectionString(v string) *CreateDBInstanceRequest
	GetConnectionString() *string
	SetDBInstanceClass(v string) *CreateDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDeployScheme(v string) *CreateDBInstanceRequest
	GetDeployScheme() *string
	SetEngine(v string) *CreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetMultiZone(v []*CreateDBInstanceRequestMultiZone) *CreateDBInstanceRequest
	GetMultiZone() []*CreateDBInstanceRequestMultiZone
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetScaleMax(v float64) *CreateDBInstanceRequest
	GetScaleMax() *float64
	SetScaleMin(v float64) *CreateDBInstanceRequest
	GetScaleMin() *float64
	SetSecurityIPList(v string) *CreateDBInstanceRequest
	GetSecurityIPList() *string
	SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest
	GetTag() []*CreateDBInstanceRequestTag
	SetUsedTime(v int32) *CreateDBInstanceRequest
	GetUsedTime() *int32
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
	// Specifies whether to add the virtual private cloud (VPC) CIDR block to the IP address whitelist. Valid values:
	//
	// 	- 1: yes.
	//
	// 	- 0: no.
	//
	// example:
	//
	// 1
	AddVPCIPs *string `json:"AddVPCIPs,omitempty" xml:"AddVPCIPs,omitempty"`
	// The reserved cache size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200GB
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AB
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterNodeCount *int32  `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	ClusterNodeType  *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance type. You can call the [DescribeAllDBInstanceClass](https://help.aliyun.com/document_detail/2853363.html) operation to query instance types.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb.xlarge
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance description.
	//
	// example:
	//
	// The instance is created for testing.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The deployment method of the instance.
	//
	// example:
	//
	// single_az
	DeployScheme *string `json:"DeployScheme,omitempty" xml:"DeployScheme,omitempty"`
	// The database engine of the instance. Default value: **selectdb**.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Default value: **3.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The configurations of multi-zone deployment.
	//
	// >
	//
	// 	- This parameter takes effect and is required only when DeployScheme is set to multi_az.
	//
	// if can be null:
	// false
	MultiZone []*CreateDBInstanceRequestMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// The unit of the subscription duration of the cluster. Valid values:
	//
	// 	- **Year**: subscription on a yearly basis.
	//
	// 	- **Month**: subscription on a monthly basis.
	//
	// >  This parameter takes effect and is required only when **ChargeType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzt2zaluvuvqa_fake
	ResourceGroupId *string  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleMax        *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin        *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The IP addresses in the whitelist of the instance. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 192.168.1.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The instance tags.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If Period is set to Year, valid values of UsedTime are 1, 2, 3, 4, and 5.
	//
	// 	- If Period is set to Month, valid values of UsedTime are 1 to 12.
	//
	// >  This parameter takes effect and is required only when **ChargeType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAddVPCIPs() *string {
	return s.AddVPCIPs
}

func (s *CreateDBInstanceRequest) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *CreateDBInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetClusterNodeCount() *int32 {
	return s.ClusterNodeCount
}

func (s *CreateDBInstanceRequest) GetClusterNodeType() *string {
	return s.ClusterNodeType
}

func (s *CreateDBInstanceRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceRequest) GetDeployScheme() *string {
	return s.DeployScheme
}

func (s *CreateDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceRequest) GetMultiZone() []*CreateDBInstanceRequestMultiZone {
	return s.MultiZone
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceRequest) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *CreateDBInstanceRequest) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *CreateDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceRequest) GetTag() []*CreateDBInstanceRequestTag {
	return s.Tag
}

func (s *CreateDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetAddVPCIPs(v string) *CreateDBInstanceRequest {
	s.AddVPCIPs = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCacheSize(v int32) *CreateDBInstanceRequest {
	s.CacheSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetChargeType(v string) *CreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClusterNodeCount(v int32) *CreateDBInstanceRequest {
	s.ClusterNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClusterNodeType(v string) *CreateDBInstanceRequest {
	s.ClusterNodeType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetConnectionString(v string) *CreateDBInstanceRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDeployScheme(v string) *CreateDBInstanceRequest {
	s.DeployScheme = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMultiZone(v []*CreateDBInstanceRequestMultiZone) *CreateDBInstanceRequest {
	s.MultiZone = v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetScaleMax(v float64) *CreateDBInstanceRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceRequest) SetScaleMin(v float64) *CreateDBInstanceRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestMultiZone struct {
	// The vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequestMultiZone) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestMultiZone) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestMultiZone) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateDBInstanceRequestMultiZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequestMultiZone) SetVSwitchIds(v []*string) *CreateDBInstanceRequestMultiZone {
	s.VSwitchIds = v
	return s
}

func (s *CreateDBInstanceRequestMultiZone) SetZoneId(v string) *CreateDBInstanceRequestMultiZone {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequestMultiZone) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceRequestTag) SetKey(v string) *CreateDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTag) SetValue(v string) *CreateDBInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
