// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CreateDBInstanceRequest
	GetBackupSetId() *string
	SetCategory(v string) *CreateDBInstanceRequest
	GetCategory() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBTimeZone(v string) *CreateDBInstanceRequest
	GetDBTimeZone() *string
	SetDeploySchema(v string) *CreateDBInstanceRequest
	GetDeploySchema() *string
	SetEngine(v string) *CreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetMultiZone(v []*CreateDBInstanceRequestMultiZone) *CreateDBInstanceRequest
	GetMultiZone() []*CreateDBInstanceRequestMultiZone
	SetNodeCount(v int32) *CreateDBInstanceRequest
	GetNodeCount() *int32
	SetNodeScaleMax(v int32) *CreateDBInstanceRequest
	GetNodeScaleMax() *int32
	SetNodeScaleMin(v int32) *CreateDBInstanceRequest
	GetNodeScaleMin() *int32
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetScaleMax(v string) *CreateDBInstanceRequest
	GetScaleMax() *string
	SetScaleMin(v string) *CreateDBInstanceRequest
	GetScaleMin() *string
	SetSourceDBInstanceId(v string) *CreateDBInstanceRequest
	GetSourceDBInstanceId() *string
	SetStorageQuota(v int64) *CreateDBInstanceRequest
	GetStorageQuota() *int64
	SetStorageType(v string) *CreateDBInstanceRequest
	GetStorageType() *string
	SetVpcId(v string) *CreateDBInstanceRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateDBInstanceRequest
	GetVswitchId() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// 1
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
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
	DBTimeZone            *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
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
	MultiZone []*CreateDBInstanceRequestMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
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

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceRequest) GetDBTimeZone() *string {
	return s.DBTimeZone
}

func (s *CreateDBInstanceRequest) GetDeploySchema() *string {
	return s.DeploySchema
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

func (s *CreateDBInstanceRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateDBInstanceRequest) GetNodeScaleMax() *int32 {
	return s.NodeScaleMax
}

func (s *CreateDBInstanceRequest) GetNodeScaleMin() *int32 {
	return s.NodeScaleMin
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateDBInstanceRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateDBInstanceRequest) GetSourceDBInstanceId() *string {
	return s.SourceDBInstanceId
}

func (s *CreateDBInstanceRequest) GetStorageQuota() *int64 {
	return s.StorageQuota
}

func (s *CreateDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetBackupSetId(v string) *CreateDBInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCategory(v string) *CreateDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBTimeZone(v string) *CreateDBInstanceRequest {
	s.DBTimeZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDeploySchema(v string) *CreateDBInstanceRequest {
	s.DeploySchema = &v
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

func (s *CreateDBInstanceRequest) SetNodeCount(v int32) *CreateDBInstanceRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNodeScaleMax(v int32) *CreateDBInstanceRequest {
	s.NodeScaleMax = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNodeScaleMin(v int32) *CreateDBInstanceRequest {
	s.NodeScaleMin = &v
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

func (s *CreateDBInstanceRequest) SetScaleMax(v string) *CreateDBInstanceRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceRequest) SetScaleMin(v string) *CreateDBInstanceRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSourceDBInstanceId(v string) *CreateDBInstanceRequest {
	s.SourceDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageQuota(v int64) *CreateDBInstanceRequest {
	s.StorageQuota = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVswitchId(v string) *CreateDBInstanceRequest {
	s.VswitchId = &v
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
	return nil
}

type CreateDBInstanceRequestMultiZone struct {
	// The vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
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
