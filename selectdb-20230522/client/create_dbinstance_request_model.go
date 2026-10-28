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
	SetConfigPatternType(v string) *CreateDBInstanceRequest
	GetConfigPatternType() *string
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
	SetFEClassCode(v string) *CreateDBInstanceRequest
	GetFEClassCode() *string
	SetFENodeCount(v string) *CreateDBInstanceRequest
	GetFENodeCount() *string
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
	SetTdeEnabled(v string) *CreateDBInstanceRequest
	GetTdeEnabled() *string
	SetTdeMode(v string) *CreateDBInstanceRequest
	GetTdeMode() *string
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
	// Specifies whether to add the VPC CIDR block to the whitelist. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
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
	// 200
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go
	//
	// - **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of cluster nodes. This parameter takes effect only in Serverless mode.
	//
	// example:
	//
	// 1
	ClusterNodeCount *int32 `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	// The cluster node type. This parameter takes effect only in Serverless mode.
	//
	// example:
	//
	// base
	ClusterNodeType *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	// The configuration template applied to the instance.
	//
	// example:
	//
	// normal
	ConfigPatternType *string `json:"ConfigPatternType,omitempty" xml:"ConfigPatternType,omitempty"`
	// The database endpoint.
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
	// selectdb.g.2xlarge
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// New instance test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The deployment mode of the instance.
	//
	// example:
	//
	// single_az
	DeployScheme *string `json:"DeployScheme,omitempty" xml:"DeployScheme,omitempty"`
	// The database engine type. Default value: **selectdb**.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Default value: **4.1**. Valid values:
	//
	// - 4.0
	//
	// - 4.1
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.1
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The node specifications of a single node in the frontend (FE) cluster. This parameter takes effect only when you create a pay-as-you-go or subscription instance. Default value: selectdb_fe.xlarge16.
	//
	// example:
	//
	// selectdb_fe.xlarge16
	FEClassCode *string `json:"FEClassCode,omitempty" xml:"FEClassCode,omitempty"`
	// The number of nodes in the FE cluster. This parameter takes effect only when you create a pay-as-you-go or subscription instance. Default value: 3. Valid values: 1 to 11. The value must be an odd number.
	//
	// example:
	//
	// 3
	FENodeCount *string `json:"FENodeCount,omitempty" xml:"FENodeCount,omitempty"`
	// The multi-zone configuration.
	//
	// >
	//
	// > - This parameter takes effect and is required when DeployScheme is set to multi_az.
	//
	// if can be null:
	// false
	MultiZone []*CreateDBInstanceRequestMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// Specifies the billing cycle of the subscription cluster. Valid values:
	//
	// - **Year**: yearly.
	//
	// - **Month**: monthly.
	//
	// > This parameter takes effect and is required only when **ChargeType*	- is set to **Prepaid**.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzt2zaluvuvqa_fake
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum value of the automatic scaling range for the RCU (RDS Capacity Unit) of the instance.
	//
	// example:
	//
	// 32
	ScaleMax *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum value of the automatic scaling range for the RCU (RDS Capacity Unit) of the instance.
	//
	// example:
	//
	// 4
	ScaleMin *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The IP address whitelist of the instance. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 192.168.1.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The tag list information of the instance.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to enable TDE on the instance.
	//
	// example:
	//
	// True
	TdeEnabled *string `json:"TdeEnabled,omitempty" xml:"TdeEnabled,omitempty"`
	// The mode for enabling Transparent Data Encryption (TDE) on the instance.
	//
	// example:
	//
	// MANAGED
	TdeMode *string `json:"TdeMode,omitempty" xml:"TdeMode,omitempty"`
	// The subscription duration of the instance. The unit is specified by the Period parameter. Year indicates years, and Month indicates months. Valid values:
	//
	// - If Period is set to Year, valid values are 1, 2, 3, 4, and 5.
	//
	// - If Period is set to Month, valid values are 1 to 12.
	//
	// > This parameter takes effect and is required only when **ChargeType*	- is set to **Prepaid**.
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
	// The VPC ID.
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

func (s *CreateDBInstanceRequest) GetConfigPatternType() *string {
	return s.ConfigPatternType
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

func (s *CreateDBInstanceRequest) GetFEClassCode() *string {
	return s.FEClassCode
}

func (s *CreateDBInstanceRequest) GetFENodeCount() *string {
	return s.FENodeCount
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

func (s *CreateDBInstanceRequest) GetTdeEnabled() *string {
	return s.TdeEnabled
}

func (s *CreateDBInstanceRequest) GetTdeMode() *string {
	return s.TdeMode
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

func (s *CreateDBInstanceRequest) SetConfigPatternType(v string) *CreateDBInstanceRequest {
	s.ConfigPatternType = &v
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

func (s *CreateDBInstanceRequest) SetFEClassCode(v string) *CreateDBInstanceRequest {
	s.FEClassCode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetFENodeCount(v string) *CreateDBInstanceRequest {
	s.FENodeCount = &v
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

func (s *CreateDBInstanceRequest) SetTdeEnabled(v string) *CreateDBInstanceRequest {
	s.TdeEnabled = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTdeMode(v string) *CreateDBInstanceRequest {
	s.TdeMode = &v
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
	if s.MultiZone != nil {
		for _, item := range s.MultiZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBInstanceRequestMultiZone struct {
	// The list of vSwitch IDs.
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
