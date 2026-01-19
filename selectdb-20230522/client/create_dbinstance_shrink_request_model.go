// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddVPCIPs(v string) *CreateDBInstanceShrinkRequest
	GetAddVPCIPs() *string
	SetCacheSize(v int32) *CreateDBInstanceShrinkRequest
	GetCacheSize() *int32
	SetChargeType(v string) *CreateDBInstanceShrinkRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetClusterNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetClusterNodeCount() *int32
	SetClusterNodeType(v string) *CreateDBInstanceShrinkRequest
	GetClusterNodeType() *string
	SetConfigPatternType(v string) *CreateDBInstanceShrinkRequest
	GetConfigPatternType() *string
	SetConnectionString(v string) *CreateDBInstanceShrinkRequest
	GetConnectionString() *string
	SetDBInstanceClass(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceDescription() *string
	SetDeployScheme(v string) *CreateDBInstanceShrinkRequest
	GetDeployScheme() *string
	SetEngine(v string) *CreateDBInstanceShrinkRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceShrinkRequest
	GetEngineVersion() *string
	SetMultiZoneShrink(v string) *CreateDBInstanceShrinkRequest
	GetMultiZoneShrink() *string
	SetPeriod(v string) *CreateDBInstanceShrinkRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetScaleMax(v float64) *CreateDBInstanceShrinkRequest
	GetScaleMax() *float64
	SetScaleMin(v float64) *CreateDBInstanceShrinkRequest
	GetScaleMin() *float64
	SetSecurityIPList(v string) *CreateDBInstanceShrinkRequest
	GetSecurityIPList() *string
	SetTagShrink(v string) *CreateDBInstanceShrinkRequest
	GetTagShrink() *string
	SetUsedTime(v int32) *CreateDBInstanceShrinkRequest
	GetUsedTime() *int32
	SetVSwitchId(v string) *CreateDBInstanceShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBInstanceShrinkRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateDBInstanceShrinkRequest
	GetZoneId() *string
}

type CreateDBInstanceShrinkRequest struct {
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
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterNodeCount  *int32  `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	ClusterNodeType   *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	ConfigPatternType *string `json:"ConfigPatternType,omitempty" xml:"ConfigPatternType,omitempty"`
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
	MultiZoneShrink *string `json:"MultiZone,omitempty" xml:"MultiZone,omitempty"`
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
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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

func (s CreateDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) GetAddVPCIPs() *string {
	return s.AddVPCIPs
}

func (s *CreateDBInstanceShrinkRequest) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *CreateDBInstanceShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetClusterNodeCount() *int32 {
	return s.ClusterNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetClusterNodeType() *string {
	return s.ClusterNodeType
}

func (s *CreateDBInstanceShrinkRequest) GetConfigPatternType() *string {
	return s.ConfigPatternType
}

func (s *CreateDBInstanceShrinkRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceShrinkRequest) GetDeployScheme() *string {
	return s.DeployScheme
}

func (s *CreateDBInstanceShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBInstanceShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceShrinkRequest) GetMultiZoneShrink() *string {
	return s.MultiZoneShrink
}

func (s *CreateDBInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceShrinkRequest) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *CreateDBInstanceShrinkRequest) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *CreateDBInstanceShrinkRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateDBInstanceShrinkRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceShrinkRequest) SetAddVPCIPs(v string) *CreateDBInstanceShrinkRequest {
	s.AddVPCIPs = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCacheSize(v int32) *CreateDBInstanceShrinkRequest {
	s.CacheSize = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetChargeType(v string) *CreateDBInstanceShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClusterNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.ClusterNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClusterNodeType(v string) *CreateDBInstanceShrinkRequest {
	s.ClusterNodeType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetConfigPatternType(v string) *CreateDBInstanceShrinkRequest {
	s.ConfigPatternType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetConnectionString(v string) *CreateDBInstanceShrinkRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceClass(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDeployScheme(v string) *CreateDBInstanceShrinkRequest {
	s.DeployScheme = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngine(v string) *CreateDBInstanceShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetMultiZoneShrink(v string) *CreateDBInstanceShrinkRequest {
	s.MultiZoneShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPeriod(v string) *CreateDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceOwnerId(v int64) *CreateDBInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetScaleMax(v float64) *CreateDBInstanceShrinkRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetScaleMin(v float64) *CreateDBInstanceShrinkRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSecurityIPList(v string) *CreateDBInstanceShrinkRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTagShrink(v string) *CreateDBInstanceShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUsedTime(v int32) *CreateDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVpcId(v string) *CreateDBInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
