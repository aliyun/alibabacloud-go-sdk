// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v string) *MigrateDBClusterRequest
	GetComputeResource() *string
	SetDBClusterId(v string) *MigrateDBClusterRequest
	GetDBClusterId() *string
	SetDryRun(v bool) *MigrateDBClusterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *MigrateDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateDBClusterRequest
	GetOwnerId() *int64
	SetProductForm(v string) *MigrateDBClusterRequest
	GetProductForm() *string
	SetProductVersion(v string) *MigrateDBClusterRequest
	GetProductVersion() *string
	SetReservedNodeCount(v int32) *MigrateDBClusterRequest
	GetReservedNodeCount() *int32
	SetReservedNodeSize(v string) *MigrateDBClusterRequest
	GetReservedNodeSize() *string
	SetResourceOwnerAccount(v string) *MigrateDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateDBClusterRequest
	GetResourceOwnerId() *int64
	SetSecondaryVSwitchId(v string) *MigrateDBClusterRequest
	GetSecondaryVSwitchId() *string
	SetSecondaryZoneId(v string) *MigrateDBClusterRequest
	GetSecondaryZoneId() *string
	SetShardNumber(v string) *MigrateDBClusterRequest
	GetShardNumber() *string
	SetStorageResource(v string) *MigrateDBClusterRequest
	GetStorageResource() *string
	SetStorageResourceSize(v string) *MigrateDBClusterRequest
	GetStorageResourceSize() *string
}

type MigrateDBClusterRequest struct {
	// The amount of reserved computing resources.
	//
	// Valid values: 0ACU to 4096ACU. Step size: 16. Each AnalyticDB compute unit (ACU) is approximately equal to 1 core and 4 GB memory.
	//
	// example:
	//
	// 32ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// LegacyForm
	ProductForm          *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	ProductVersion       *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	ReservedNodeCount    *int32  `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	ReservedNodeSize     *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecondaryVSwitchId   *string `json:"SecondaryVSwitchId,omitempty" xml:"SecondaryVSwitchId,omitempty"`
	SecondaryZoneId      *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The number of shards that you want to change during the data migration.
	//
	// example:
	//
	// 128
	ShardNumber *string `json:"ShardNumber,omitempty" xml:"ShardNumber,omitempty"`
	// The amount of reserved storage resources. Valid values: 0ACU to 2064ACU. The value must be in increments of the number of ACUs specified by the StorageResourceSize parameter × 3 (default value: 24ACU). Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The node specifications of reserved storage resources. Valid values: 8ACU, 12ACU, and 16ACU.
	//
	// example:
	//
	// 8ACU
	StorageResourceSize *string `json:"StorageResourceSize,omitempty" xml:"StorageResourceSize,omitempty"`
}

func (s MigrateDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterRequest) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *MigrateDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *MigrateDBClusterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *MigrateDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateDBClusterRequest) GetProductForm() *string {
	return s.ProductForm
}

func (s *MigrateDBClusterRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *MigrateDBClusterRequest) GetReservedNodeCount() *int32 {
	return s.ReservedNodeCount
}

func (s *MigrateDBClusterRequest) GetReservedNodeSize() *string {
	return s.ReservedNodeSize
}

func (s *MigrateDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateDBClusterRequest) GetSecondaryVSwitchId() *string {
	return s.SecondaryVSwitchId
}

func (s *MigrateDBClusterRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *MigrateDBClusterRequest) GetShardNumber() *string {
	return s.ShardNumber
}

func (s *MigrateDBClusterRequest) GetStorageResource() *string {
	return s.StorageResource
}

func (s *MigrateDBClusterRequest) GetStorageResourceSize() *string {
	return s.StorageResourceSize
}

func (s *MigrateDBClusterRequest) SetComputeResource(v string) *MigrateDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *MigrateDBClusterRequest) SetDBClusterId(v string) *MigrateDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetDryRun(v bool) *MigrateDBClusterRequest {
	s.DryRun = &v
	return s
}

func (s *MigrateDBClusterRequest) SetOwnerAccount(v string) *MigrateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateDBClusterRequest) SetOwnerId(v int64) *MigrateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetProductForm(v string) *MigrateDBClusterRequest {
	s.ProductForm = &v
	return s
}

func (s *MigrateDBClusterRequest) SetProductVersion(v string) *MigrateDBClusterRequest {
	s.ProductVersion = &v
	return s
}

func (s *MigrateDBClusterRequest) SetReservedNodeCount(v int32) *MigrateDBClusterRequest {
	s.ReservedNodeCount = &v
	return s
}

func (s *MigrateDBClusterRequest) SetReservedNodeSize(v string) *MigrateDBClusterRequest {
	s.ReservedNodeSize = &v
	return s
}

func (s *MigrateDBClusterRequest) SetResourceOwnerAccount(v string) *MigrateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateDBClusterRequest) SetResourceOwnerId(v int64) *MigrateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetSecondaryVSwitchId(v string) *MigrateDBClusterRequest {
	s.SecondaryVSwitchId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetSecondaryZoneId(v string) *MigrateDBClusterRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetShardNumber(v string) *MigrateDBClusterRequest {
	s.ShardNumber = &v
	return s
}

func (s *MigrateDBClusterRequest) SetStorageResource(v string) *MigrateDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *MigrateDBClusterRequest) SetStorageResourceSize(v string) *MigrateDBClusterRequest {
	s.StorageResourceSize = &v
	return s
}

func (s *MigrateDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
