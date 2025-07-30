// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v string) *CreateDBClusterRequest
	GetCacheSize() *string
	SetChargeType(v string) *CreateDBClusterRequest
	GetChargeType() *string
	SetClusterNodeCount(v int32) *CreateDBClusterRequest
	GetClusterNodeCount() *int32
	SetClusterNodeType(v string) *CreateDBClusterRequest
	GetClusterNodeType() *string
	SetDBClusterClass(v string) *CreateDBClusterRequest
	GetDBClusterClass() *string
	SetDBClusterDescription(v string) *CreateDBClusterRequest
	GetDBClusterDescription() *string
	SetDBInstanceId(v string) *CreateDBClusterRequest
	GetDBInstanceId() *string
	SetEngine(v string) *CreateDBClusterRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBClusterRequest
	GetEngineVersion() *string
	SetPeriod(v string) *CreateDBClusterRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CreateDBClusterRequest
	GetResourceOwnerId() *int64
	SetScaleMax(v float64) *CreateDBClusterRequest
	GetScaleMax() *float64
	SetScaleMin(v float64) *CreateDBClusterRequest
	GetScaleMin() *float64
	SetUsedTime(v string) *CreateDBClusterRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateDBClusterRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBClusterRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateDBClusterRequest
	GetZoneId() *string
}

type CreateDBClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 200
	CacheSize *string `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterNodeCount *int32  `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	ClusterNodeType  *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb.2xlarge
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleMax        *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin        *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) GetCacheSize() *string {
	return s.CacheSize
}

func (s *CreateDBClusterRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDBClusterRequest) GetClusterNodeCount() *int32 {
	return s.ClusterNodeCount
}

func (s *CreateDBClusterRequest) GetClusterNodeType() *string {
	return s.ClusterNodeType
}

func (s *CreateDBClusterRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *CreateDBClusterRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *CreateDBClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBClusterRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBClusterRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterRequest) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *CreateDBClusterRequest) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *CreateDBClusterRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBClusterRequest) SetCacheSize(v string) *CreateDBClusterRequest {
	s.CacheSize = &v
	return s
}

func (s *CreateDBClusterRequest) SetChargeType(v string) *CreateDBClusterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBClusterRequest) SetClusterNodeCount(v int32) *CreateDBClusterRequest {
	s.ClusterNodeCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetClusterNodeType(v string) *CreateDBClusterRequest {
	s.ClusterNodeType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterClass(v string) *CreateDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBInstanceId(v string) *CreateDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBClusterRequest) SetEngine(v string) *CreateDBClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBClusterRequest) SetEngineVersion(v string) *CreateDBClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleMax(v float64) *CreateDBClusterRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleMin(v float64) *CreateDBClusterRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVpcId(v string) *CreateDBClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
