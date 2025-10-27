// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifySyncJobRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifySyncJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySyncJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySyncJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySyncJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySyncJobRequest
	GetResourceOwnerId() *int64
	SetSourceDBCluster(v []*ModifySyncJobRequestSourceDBCluster) *ModifySyncJobRequest
	GetSourceDBCluster() []*ModifySyncJobRequestSourceDBCluster
	SetSyncPlatform(v string) *ModifySyncJobRequest
	GetSyncPlatform() *string
}

type ModifySyncJobRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp12bh6z59nh8497f
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source instances or clusters.
	//
	// if can be null:
	// true
	SourceDBCluster []*ModifySyncJobRequestSourceDBCluster `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty" type:"Repeated"`
	// The synchronization platform.
	//
	// example:
	//
	// ADB-CDC
	SyncPlatform *string `json:"SyncPlatform,omitempty" xml:"SyncPlatform,omitempty"`
}

func (s ModifySyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySyncJobRequest) GoString() string {
	return s.String()
}

func (s *ModifySyncJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySyncJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySyncJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySyncJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySyncJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySyncJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySyncJobRequest) GetSourceDBCluster() []*ModifySyncJobRequestSourceDBCluster {
	return s.SourceDBCluster
}

func (s *ModifySyncJobRequest) GetSyncPlatform() *string {
	return s.SyncPlatform
}

func (s *ModifySyncJobRequest) SetDBClusterId(v string) *ModifySyncJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySyncJobRequest) SetOwnerAccount(v string) *ModifySyncJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySyncJobRequest) SetOwnerId(v int64) *ModifySyncJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySyncJobRequest) SetRegionId(v string) *ModifySyncJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySyncJobRequest) SetResourceOwnerAccount(v string) *ModifySyncJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySyncJobRequest) SetResourceOwnerId(v int64) *ModifySyncJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySyncJobRequest) SetSourceDBCluster(v []*ModifySyncJobRequestSourceDBCluster) *ModifySyncJobRequest {
	s.SourceDBCluster = v
	return s
}

func (s *ModifySyncJobRequest) SetSyncPlatform(v string) *ModifySyncJobRequest {
	s.SyncPlatform = &v
	return s
}

func (s *ModifySyncJobRequest) Validate() error {
	if s.SourceDBCluster != nil {
		for _, item := range s.SourceDBCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySyncJobRequestSourceDBCluster struct {
	// The ID of the source instance or cluster. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2zepqn129i9s3l2z3,rm-2zea4dj583129ksp6
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The ID of the synchronization job.
	//
	// example:
	//
	// dts-xxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The operation type.
	//
	// Valid values:
	//
	// 	- Create
	//
	// 	- Modify
	//
	// This parameter is required.
	//
	// example:
	//
	// Create
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source database type.
	//
	// Valid values:
	//
	// 	- rds: ApsaraDB RDS.
	//
	// 	- sls: Simple Log Service.
	//
	// 	- polardb: PolarDB.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifySyncJobRequestSourceDBCluster) String() string {
	return dara.Prettify(s)
}

func (s ModifySyncJobRequestSourceDBCluster) GoString() string {
	return s.String()
}

func (s *ModifySyncJobRequestSourceDBCluster) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *ModifySyncJobRequestSourceDBCluster) GetJobId() *string {
	return s.JobId
}

func (s *ModifySyncJobRequestSourceDBCluster) GetOperateType() *string {
	return s.OperateType
}

func (s *ModifySyncJobRequestSourceDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySyncJobRequestSourceDBCluster) GetType() *string {
	return s.Type
}

func (s *ModifySyncJobRequestSourceDBCluster) SetClusterIds(v string) *ModifySyncJobRequestSourceDBCluster {
	s.ClusterIds = &v
	return s
}

func (s *ModifySyncJobRequestSourceDBCluster) SetJobId(v string) *ModifySyncJobRequestSourceDBCluster {
	s.JobId = &v
	return s
}

func (s *ModifySyncJobRequestSourceDBCluster) SetOperateType(v string) *ModifySyncJobRequestSourceDBCluster {
	s.OperateType = &v
	return s
}

func (s *ModifySyncJobRequestSourceDBCluster) SetRegionId(v string) *ModifySyncJobRequestSourceDBCluster {
	s.RegionId = &v
	return s
}

func (s *ModifySyncJobRequestSourceDBCluster) SetType(v string) *ModifySyncJobRequestSourceDBCluster {
	s.Type = &v
	return s
}

func (s *ModifySyncJobRequestSourceDBCluster) Validate() error {
	return dara.Validate(s)
}
