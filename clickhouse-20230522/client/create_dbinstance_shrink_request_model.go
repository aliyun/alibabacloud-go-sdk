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
	SetCategory(v string) *CreateDBInstanceShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceDescription() *string
	SetDBTimeZone(v string) *CreateDBInstanceShrinkRequest
	GetDBTimeZone() *string
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
	SetTags(v []*CreateDBInstanceShrinkRequestTags) *CreateDBInstanceShrinkRequest
	GetTags() []*CreateDBInstanceShrinkRequestTags
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
	// The edition of the instance. Valid value:
	//
	// - `enterprise`: Enterprise Edition
	//
	// example:
	//
	// enterprise
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A client-provided token to ensure request idempotence. It must be unique across requests, contain only ASCII characters, and not exceed 64 characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Cluster test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The time zone of the database, which must be an IANA time zone identifier.
	//
	// example:
	//
	// Asia/Shanghai
	DBTimeZone *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
	// The deployment mode of the instance.
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
	// The multi-zone configuration.
	MultiZoneShrink *string `json:"MultiZone,omitempty" xml:"MultiZone,omitempty"`
	// The number of nodes. Valid values: 2 to 16. This parameter is required when you configure an elastic scaling range by using the `NodeScaleMin` and `NodeScaleMax` parameters.
	//
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The maximum number of nodes for serverless elastic scaling. Valid values: 4 to 32. The value must be greater than the `NodeScaleMin` parameter.
	//
	// example:
	//
	// 4
	NodeScaleMax *int32 `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
	// The minimum number of nodes for serverless elastic scaling. Valid values: 4 to 32.
	//
	// example:
	//
	// 32
	NodeScaleMin *int32 `json:"NodeScaleMin,omitempty" xml:"NodeScaleMin,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is deprecated. Use the `NodeCount`, `NodeScaleMin`, and `NodeScaleMax` parameters to configure elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// This parameter is deprecated. Use the `NodeCount`, `NodeScaleMin`, and `NodeScaleMax` parameters to configure elastic scaling.
	//
	// example:
	//
	// 8
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The ID of the source instance. This parameter is required when restoring from a backup.
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
	// The tags to add to the instance.
	Tags []*CreateDBInstanceShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID.
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

func (s *CreateDBInstanceShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceShrinkRequest) GetDBTimeZone() *string {
	return s.DBTimeZone
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

func (s *CreateDBInstanceShrinkRequest) GetTags() []*CreateDBInstanceShrinkRequestTags {
	return s.Tags
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

func (s *CreateDBInstanceShrinkRequest) SetCategory(v string) *CreateDBInstanceShrinkRequest {
	s.Category = &v
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

func (s *CreateDBInstanceShrinkRequest) SetDBTimeZone(v string) *CreateDBInstanceShrinkRequest {
	s.DBTimeZone = &v
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

func (s *CreateDBInstanceShrinkRequest) SetTags(v []*CreateDBInstanceShrinkRequestTags) *CreateDBInstanceShrinkRequest {
	s.Tags = v
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBInstanceShrinkRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// user_123
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// example string
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceShrinkRequestTags) SetKey(v string) *CreateDBInstanceShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceShrinkRequestTags) SetValue(v string) *CreateDBInstanceShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
