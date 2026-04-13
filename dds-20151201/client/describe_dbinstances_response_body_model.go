// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstances(v *DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody
	GetDBInstances() *DescribeDBInstancesResponseBodyDBInstances
	SetPageNumber(v int32) *DescribeDBInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeDBInstancesResponseBody struct {
	DBInstances *DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0E4FE33F-5510-5758-8FA7-A6672CDE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances in the query results.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) GetDBInstances() *DescribeDBInstancesResponseBodyDBInstances {
	return s.DBInstances
}

func (s *DescribeDBInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBInstancesResponseBody) SetDBInstances(v *DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int32) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) Validate() error {
	if s.DBInstances != nil {
		if err := s.DBInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	DBInstance []*DescribeDBInstancesResponseBodyDBInstancesDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDBInstance() []*DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBInstance(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstance) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) Validate() error {
	if s.DBInstance != nil {
		for _, item := range s.DBInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstance struct {
	BackupRetentionPolicy *int32                                                          `json:"BackupRetentionPolicy,omitempty" xml:"BackupRetentionPolicy,omitempty"`
	CapacityUnit          *string                                                         `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	ChargeType            *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreationTime          *string                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceClass       *string                                                         `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string                                                         `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string                                                         `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *string                                                         `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage     *int32                                                          `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceType        *string                                                         `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DestroyTime           *string                                                         `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	Engine                *string                                                         `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                                         `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime            *string                                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HiddenZoneId          *string                                                         `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	KindCode              *string                                                         `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	LastDowngradeTime     *string                                                         `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	LockMode              *string                                                         `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MongosList            *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Struct"`
	NetworkType           *string                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId              *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReleaseTime           *string                                                         `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	ReplicationFactor     *string                                                         `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceGroupId       *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryZoneId       *string                                                         `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	ShardList             *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList  `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Struct"`
	StorageType           *string                                                         `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                  *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcAuthMode           *string                                                         `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	ZoneId                *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetBackupRetentionPolicy() *int32 {
	return s.BackupRetentionPolicy
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetKindCode() *string {
	return s.KindCode
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetLastDowngradeTime() *string {
	return s.LastDowngradeTime
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetMongosList() *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList {
	return s.MongosList
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetShardList() *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList {
	return s.ShardList
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetTags() *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags {
	return s.Tags
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetBackupRetentionPolicy(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.BackupRetentionPolicy = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetCapacityUnit(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetChargeType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetCreationTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceStorage(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDestroyTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetHiddenZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.HiddenZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetKindCode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetLastDowngradeTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.LastDowngradeTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetMongosList(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.MongosList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetNetworkType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetReleaseTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetReplicationFactor(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetSecondaryZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetShardList(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ShardList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetStorageType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetTags(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetVpcAuthMode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) Validate() error {
	if s.MongosList != nil {
		if err := s.MongosList.Validate(); err != nil {
			return err
		}
	}
	if s.ShardList != nil {
		if err := s.ShardList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList struct {
	MongosAttribute []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute `json:"MongosAttribute,omitempty" xml:"MongosAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) GetMongosAttribute() []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	return s.MongosAttribute
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) SetMongosAttribute(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList {
	s.MongosAttribute = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) Validate() error {
	if s.MongosAttribute != nil {
		for _, item := range s.MongosAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute struct {
	NodeClass       *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList struct {
	ShardAttribute []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute `json:"ShardAttribute,omitempty" xml:"ShardAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) GetShardAttribute() []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	return s.ShardAttribute
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) SetShardAttribute(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList {
	s.ShardAttribute = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) Validate() error {
	if s.ShardAttribute != nil {
		for _, item := range s.ShardAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute struct {
	NodeClass        *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription  *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId           *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeStorage      *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	ReadonlyReplicas *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeStorage(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetReadonlyReplicas(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ReadonlyReplicas = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) GetTag() []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag {
	return s.Tag
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) Validate() error {
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

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
