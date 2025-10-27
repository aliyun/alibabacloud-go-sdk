// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAvailableDBClusterListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSyncAvailableDBClusterListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSyncAvailableDBClusterListRequest
	GetOwnerId() *int64
	SetQueryType(v string) *DescribeSyncAvailableDBClusterListRequest
	GetQueryType() *string
	SetRegionId(v string) *DescribeSyncAvailableDBClusterListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSyncAvailableDBClusterListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSyncAvailableDBClusterListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSyncAvailableDBClusterListRequest
	GetResourceOwnerId() *int64
	SetSourceDBCluster(v []*DescribeSyncAvailableDBClusterListRequestSourceDBCluster) *DescribeSyncAvailableDBClusterListRequest
	GetSourceDBCluster() []*DescribeSyncAvailableDBClusterListRequestSourceDBCluster
	SetSyncPlatform(v string) *DescribeSyncAvailableDBClusterListRequest
	GetSyncPlatform() *string
}

type DescribeSyncAvailableDBClusterListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The query type.
	//
	// Valid values:
	//
	// 	- Target
	//
	// 	- Source
	//
	// This parameter is required.
	//
	// example:
	//
	// Source
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
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
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source instances or clusters.
	SourceDBCluster []*DescribeSyncAvailableDBClusterListRequestSourceDBCluster `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty" type:"Repeated"`
	// The synchronization platform.
	//
	// example:
	//
	// ADB-CDC
	SyncPlatform *string `json:"SyncPlatform,omitempty" xml:"SyncPlatform,omitempty"`
}

func (s DescribeSyncAvailableDBClusterListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAvailableDBClusterListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetSourceDBCluster() []*DescribeSyncAvailableDBClusterListRequestSourceDBCluster {
	return s.SourceDBCluster
}

func (s *DescribeSyncAvailableDBClusterListRequest) GetSyncPlatform() *string {
	return s.SyncPlatform
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetOwnerAccount(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetOwnerId(v int64) *DescribeSyncAvailableDBClusterListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetQueryType(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetRegionId(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetResourceGroupId(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetResourceOwnerAccount(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetResourceOwnerId(v int64) *DescribeSyncAvailableDBClusterListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetSourceDBCluster(v []*DescribeSyncAvailableDBClusterListRequestSourceDBCluster) *DescribeSyncAvailableDBClusterListRequest {
	s.SourceDBCluster = v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) SetSyncPlatform(v string) *DescribeSyncAvailableDBClusterListRequest {
	s.SyncPlatform = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequest) Validate() error {
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

type DescribeSyncAvailableDBClusterListRequestSourceDBCluster struct {
	// The ID of the source instance or cluster. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rm-bp1l3yh04y7us147n
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The database type of the source instance or cluster.
	//
	// Valid values:
	//
	// 	- rds: ApsaraDB RDS.
	//
	// 	- sls: Simple Log Service.
	//
	// 	- polardb: PolarDB.
	//
	// example:
	//
	// rds
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSyncAvailableDBClusterListRequestSourceDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAvailableDBClusterListRequestSourceDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) GetType() *string {
	return s.Type
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) SetClusterIds(v string) *DescribeSyncAvailableDBClusterListRequestSourceDBCluster {
	s.ClusterIds = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) SetRegionId(v string) *DescribeSyncAvailableDBClusterListRequestSourceDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) SetType(v string) *DescribeSyncAvailableDBClusterListRequestSourceDBCluster {
	s.Type = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListRequestSourceDBCluster) Validate() error {
	return dara.Validate(s)
}
