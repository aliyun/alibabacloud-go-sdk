// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CreateDBInstanceShrinkRequest
	GetBackupSetId() *string
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceDescription() *string
	SetDeploySchema(v string) *CreateDBInstanceShrinkRequest
	GetDeploySchema() *string
	SetEngine(v string) *CreateDBInstanceShrinkRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceShrinkRequest
	GetEngineVersion() *string
	SetMultiZoneShrink(v string) *CreateDBInstanceShrinkRequest
	GetMultiZoneShrink() *string
	SetNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetNodeCount() *int32
	SetNodeScaleMax(v int32) *CreateDBInstanceShrinkRequest
	GetNodeScaleMax() *int32
	SetNodeScaleMin(v int32) *CreateDBInstanceShrinkRequest
	GetNodeScaleMin() *int32
	SetRegionId(v string) *CreateDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetScaleMax(v string) *CreateDBInstanceShrinkRequest
	GetScaleMax() *string
	SetScaleMin(v string) *CreateDBInstanceShrinkRequest
	GetScaleMin() *string
	SetSourceDBInstanceId(v string) *CreateDBInstanceShrinkRequest
	GetSourceDBInstanceId() *string
	SetStorageQuota(v int64) *CreateDBInstanceShrinkRequest
	GetStorageQuota() *int64
	SetStorageType(v string) *CreateDBInstanceShrinkRequest
	GetStorageType() *string
	SetVpcId(v string) *CreateDBInstanceShrinkRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateDBInstanceShrinkRequest
	GetVswitchId() *string
	SetZoneId(v string) *CreateDBInstanceShrinkRequest
	GetZoneId() *string
}

type CreateDBInstanceShrinkRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// 1
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token. Make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Used for test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The deployment status of the cluster.
	//
	// example:
	//
	// multi_az
	DeploySchema *string `json:"DeploySchema,omitempty" xml:"DeploySchema,omitempty"`
	// The engine type.
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 23.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The configurations of multi-zone deployment.
	MultiZoneShrink *string `json:"MultiZone,omitempty" xml:"MultiZone,omitempty"`
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// 4
	NodeScaleMax *int32 `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
	// example:
	//
	// 32
	NodeScaleMin *int32 `json:"NodeScaleMin,omitempty" xml:"NodeScaleMin,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum capacity for auto scaling.
	//
	// example:
	//
	// 32
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for auto scaling.
	//
	// example:
	//
	// 8
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-2ze1*********
	SourceDBInstanceId *string `json:"SourceDBInstanceId,omitempty" xml:"SourceDBInstanceId,omitempty"`
	// example:
	//
	// 100
	StorageQuota *int64 `json:"StorageQuota,omitempty" xml:"StorageQuota,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6xmupdn7v6ui9f****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf632qye9oqt4x4sr****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceShrinkRequest) GetDeploySchema() *string {
	return s.DeploySchema
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

func (s *CreateDBInstanceShrinkRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetNodeScaleMax() *int32 {
	return s.NodeScaleMax
}

func (s *CreateDBInstanceShrinkRequest) GetNodeScaleMin() *int32 {
	return s.NodeScaleMin
}

func (s *CreateDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateDBInstanceShrinkRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateDBInstanceShrinkRequest) GetSourceDBInstanceId() *string {
	return s.SourceDBInstanceId
}

func (s *CreateDBInstanceShrinkRequest) GetStorageQuota() *int64 {
	return s.StorageQuota
}

func (s *CreateDBInstanceShrinkRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceShrinkRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceShrinkRequest) SetBackupSetId(v string) *CreateDBInstanceShrinkRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDeploySchema(v string) *CreateDBInstanceShrinkRequest {
	s.DeploySchema = &v
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

func (s *CreateDBInstanceShrinkRequest) SetNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetNodeScaleMax(v int32) *CreateDBInstanceShrinkRequest {
	s.NodeScaleMax = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetNodeScaleMin(v int32) *CreateDBInstanceShrinkRequest {
	s.NodeScaleMin = &v
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

func (s *CreateDBInstanceShrinkRequest) SetScaleMax(v string) *CreateDBInstanceShrinkRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetScaleMin(v string) *CreateDBInstanceShrinkRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSourceDBInstanceId(v string) *CreateDBInstanceShrinkRequest {
	s.SourceDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageQuota(v int64) *CreateDBInstanceShrinkRequest {
	s.StorageQuota = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageType(v string) *CreateDBInstanceShrinkRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVpcId(v string) *CreateDBInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVswitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
