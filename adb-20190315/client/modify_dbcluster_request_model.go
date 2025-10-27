// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v string) *ModifyDBClusterRequest
	GetComputeResource() *string
	SetDBClusterCategory(v string) *ModifyDBClusterRequest
	GetDBClusterCategory() *string
	SetDBClusterId(v string) *ModifyDBClusterRequest
	GetDBClusterId() *string
	SetDBNodeClass(v string) *ModifyDBClusterRequest
	GetDBNodeClass() *string
	SetDBNodeGroupCount(v string) *ModifyDBClusterRequest
	GetDBNodeGroupCount() *string
	SetDBNodeStorage(v string) *ModifyDBClusterRequest
	GetDBNodeStorage() *string
	SetDiskPerformanceLevel(v string) *ModifyDBClusterRequest
	GetDiskPerformanceLevel() *string
	SetElasticIOResource(v int32) *ModifyDBClusterRequest
	GetElasticIOResource() *int32
	SetElasticIOResourceSize(v string) *ModifyDBClusterRequest
	GetElasticIOResourceSize() *string
	SetExecutorCount(v string) *ModifyDBClusterRequest
	GetExecutorCount() *string
	SetMode(v string) *ModifyDBClusterRequest
	GetMode() *string
	SetModifyType(v string) *ModifyDBClusterRequest
	GetModifyType() *string
	SetOwnerAccount(v string) *ModifyDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterRequest
	GetResourceOwnerId() *int64
	SetStorageResource(v string) *ModifyDBClusterRequest
	GetStorageResource() *string
}

type ModifyDBClusterRequest struct {
	// The computing resources of the cluster. You can call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/190632.html) operation to query the computing resources that are available within a region.
	//
	// > This parameter must be specified when Mode is set to Flexible.
	//
	// example:
	//
	// 32Core128GBNEW
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Cluster**: reserved mode for Cluster Edition.
	//
	// 	- **MixedStorage**: elastic mode for Cluster Edition.
	//
	// > If you set DBClusterCategory to Cluster, you must set Mode to Reserver. If you set DBClusterCategory to MixedStorage, you must set Mode to Flexible. Otherwise, you fail to change the specifications of the cluster.
	//
	// example:
	//
	// MixedStorage
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node specifications of the cluster. Valid values:
	//
	// 	- **C8**
	//
	// 	- **C32**
	//
	// > This parameter must be specified when Mode is set to Reserver.
	//
	// example:
	//
	// C32
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of node groups. Valid values: 1 to 200.
	//
	// > This parameter must be specified when Mode is set to Reserver.
	//
	// example:
	//
	// 2
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity per node. Unit: GB.
	//
	// 	- Valid values when DBClusterClass is set to C8: 100 to 2000.
	//
	// 	- Valid values when DBClusterClass is set to C32: 100 to 8000.
	//
	// >
	//
	// 	- This parameter must be specified when Mode is set to Reserver.
	//
	// 	- The storage capacity less than 1,000 GB increases in 100 GB increments. The storage capacity greater than 1,000 GB increases in 1,000 GB increments.
	//
	// example:
	//
	// 200
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The Enterprise SSD (ESSD) performance level of the cluster. Valid values:
	//
	// 	- PL0
	//
	// 	- PL1
	//
	// 	- PL2
	//
	// 	- PL3
	//
	// example:
	//
	// PL1
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The number of EIUs. The number of EIUs that you can purchase varies based on the single-node EIU specifications.
	//
	// 	- If the single-node EIU specifications are 8 cores and 64 GB, you can purchase up to 32 EIUs.
	//
	// 	- If the single-node EIU specifications are 12 cores and 96 GB, you can purchase up to 16 EIUs.
	//
	// example:
	//
	// 2
	ElasticIOResource *int32 `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// The single-node specifications of an elastic I/O unit (EIU). Valid values:
	//
	// 	- **8Core64GB**: If you set the parameter to **8Core64GB**, the specifications of an EIU are 24 cores and 192 GB memory.
	//
	// 	- **12Core96GB**: If you set the parameter to **12Core96GB**, the specifications of an EIU are 36 cores and 288 GB memory.
	//
	// >  This parameter takes effect only when your cluster meets the following requirements:
	//
	// 	- The cluster is in elastic mode.
	//
	// 	- If the cluster resides in the China (Guangzhou), China (Shenzhen), China (Hangzhou), China (Shanghai), China (Qingdao), China (Beijing), or China (Zhangjiakou) region, the cluster has 16 cores and 64 GB memory or higher specifications.
	//
	// 	- If the cluster resides in another region, the cluster has 32 cores and 128 GB memory or higher specifications.
	//
	// example:
	//
	// 8Core64GB
	ElasticIOResourceSize *string `json:"ElasticIOResourceSize,omitempty" xml:"ElasticIOResourceSize,omitempty"`
	// N/A
	//
	// example:
	//
	// None
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// 	- **Reserver**: the reserved mode.
	//
	// 	- **Flexible**: the elastic mode.
	//
	// example:
	//
	// Flexible
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The change type. Valid values:
	//
	// 	- **Upgrade**
	//
	// 	- **Downgrade**
	//
	// example:
	//
	// Upgrade
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// N/A
	//
	// example:
	//
	// None
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *ModifyDBClusterRequest) GetDBClusterCategory() *string {
	return s.DBClusterCategory
}

func (s *ModifyDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *ModifyDBClusterRequest) GetDBNodeGroupCount() *string {
	return s.DBNodeGroupCount
}

func (s *ModifyDBClusterRequest) GetDBNodeStorage() *string {
	return s.DBNodeStorage
}

func (s *ModifyDBClusterRequest) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *ModifyDBClusterRequest) GetElasticIOResource() *int32 {
	return s.ElasticIOResource
}

func (s *ModifyDBClusterRequest) GetElasticIOResourceSize() *string {
	return s.ElasticIOResourceSize
}

func (s *ModifyDBClusterRequest) GetExecutorCount() *string {
	return s.ExecutorCount
}

func (s *ModifyDBClusterRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyDBClusterRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterRequest) GetStorageResource() *string {
	return s.StorageResource
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterCategory(v string) *ModifyDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeClass(v string) *ModifyDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeGroupCount(v string) *ModifyDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeStorage(v string) *ModifyDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDiskPerformanceLevel(v string) *ModifyDBClusterRequest {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResource(v int32) *ModifyDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResourceSize(v string) *ModifyDBClusterRequest {
	s.ElasticIOResourceSize = &v
	return s
}

func (s *ModifyDBClusterRequest) SetExecutorCount(v string) *ModifyDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetMode(v string) *ModifyDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetModifyType(v string) *ModifyDBClusterRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *ModifyDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
