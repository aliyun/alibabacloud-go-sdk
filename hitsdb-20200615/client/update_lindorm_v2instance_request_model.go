// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityStorageSize(v int32) *UpdateLindormV2InstanceRequest
	GetCapacityStorageSize() *int32
	SetCloudStorageSize(v int32) *UpdateLindormV2InstanceRequest
	GetCloudStorageSize() *int32
	SetCloudStorageType(v string) *UpdateLindormV2InstanceRequest
	GetCloudStorageType() *string
	SetEnableCapacityStorage(v bool) *UpdateLindormV2InstanceRequest
	GetEnableCapacityStorage() *bool
	SetEngineList(v []*UpdateLindormV2InstanceRequestEngineList) *UpdateLindormV2InstanceRequest
	GetEngineList() []*UpdateLindormV2InstanceRequestEngineList
	SetInstanceId(v string) *UpdateLindormV2InstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateLindormV2InstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLindormV2InstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpdateLindormV2InstanceRequest
	GetSecurityToken() *string
}

type UpdateLindormV2InstanceRequest struct {
	// The capacity of the storage-optimized storage.
	//
	// example:
	//
	// 10000
	CapacityStorageSize *int32 `json:"CapacityStorageSize,omitempty" xml:"CapacityStorageSize,omitempty"`
	// The cloud storage capacity. Unit: GB.
	//
	// example:
	//
	// 480
	CloudStorageSize *int32 `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	// The cloud storage class.
	//
	// - **PerformanceStorage**: performance cloud storage.
	//
	// - **StandardStorage**: standard cloud storage.
	//
	// example:
	//
	// PerformanceStorage
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	// Specifies whether to enable storage-optimized storage.
	//
	// example:
	//
	// false
	EnableCapacityStorage *bool `json:"EnableCapacityStorage,omitempty" xml:"EnableCapacityStorage,omitempty"`
	// A list of engine types.
	//
	// This parameter is required.
	EngineList []*UpdateLindormV2InstanceRequestEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance is located. To query the latest region list, call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceRequest) GetCapacityStorageSize() *int32 {
	return s.CapacityStorageSize
}

func (s *UpdateLindormV2InstanceRequest) GetCloudStorageSize() *int32 {
	return s.CloudStorageSize
}

func (s *UpdateLindormV2InstanceRequest) GetCloudStorageType() *string {
	return s.CloudStorageType
}

func (s *UpdateLindormV2InstanceRequest) GetEnableCapacityStorage() *bool {
	return s.EnableCapacityStorage
}

func (s *UpdateLindormV2InstanceRequest) GetEngineList() []*UpdateLindormV2InstanceRequestEngineList {
	return s.EngineList
}

func (s *UpdateLindormV2InstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLindormV2InstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLindormV2InstanceRequest) SetCapacityStorageSize(v int32) *UpdateLindormV2InstanceRequest {
	s.CapacityStorageSize = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetCloudStorageSize(v int32) *UpdateLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetCloudStorageType(v string) *UpdateLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetEnableCapacityStorage(v bool) *UpdateLindormV2InstanceRequest {
	s.EnableCapacityStorage = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetEngineList(v []*UpdateLindormV2InstanceRequestEngineList) *UpdateLindormV2InstanceRequest {
	s.EngineList = v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetInstanceId(v string) *UpdateLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetOwnerAccount(v string) *UpdateLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetOwnerId(v int64) *UpdateLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetRegionId(v string) *UpdateLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *UpdateLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetResourceOwnerId(v int64) *UpdateLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) SetSecurityToken(v string) *UpdateLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLindormV2InstanceRequest) Validate() error {
	if s.EngineList != nil {
		for _, item := range s.EngineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLindormV2InstanceRequestEngineList struct {
	// The engine type. Valid values:
	//
	// - TABLE: LindormTable.
	//
	// - TSDB: LindormTSDB.
	//
	// - LSEARCH: search engine.
	//
	// - LTS: LTS engine.
	//
	// - LVECTOR: vector engine.
	//
	// - LCOLUMN: column store.
	//
	// - LAI: AI engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// A list of engine node groups.
	NodeGroupList []*UpdateLindormV2InstanceRequestEngineListNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
}

func (s UpdateLindormV2InstanceRequestEngineList) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceRequestEngineList) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceRequestEngineList) GetEngineType() *string {
	return s.EngineType
}

func (s *UpdateLindormV2InstanceRequestEngineList) GetNodeGroupList() []*UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	return s.NodeGroupList
}

func (s *UpdateLindormV2InstanceRequestEngineList) SetEngineType(v string) *UpdateLindormV2InstanceRequestEngineList {
	s.EngineType = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineList) SetNodeGroupList(v []*UpdateLindormV2InstanceRequestEngineListNodeGroupList) *UpdateLindormV2InstanceRequestEngineList {
	s.NodeGroupList = v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineList) Validate() error {
	if s.NodeGroupList != nil {
		for _, item := range s.NodeGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLindormV2InstanceRequestEngineListNodeGroupList struct {
	// The ID of the node group.
	//
	// example:
	//
	// ix90Yes
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The disk size of a single node. This parameter is not required.
	//
	// example:
	//
	// 0
	NodeDiskSize *int32 `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	// The disk type of the node. This parameter is not required. **This parameter is available only for specific scenarios and is accessible to users on a whitelist.**
	//
	// example:
	//
	// cloud_essd
	NodeDiskType *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	// The node specifications of the engine.
	//
	// - **lindorm.c.2xlarge**: 8 cores and 16 GB of memory.
	//
	// - **lindorm.g.2xlarge**: 8 cores and 32 GB of memory.
	//
	// - **lindorm.c.4xlarge**: 16 cores and 32 GB of memory.
	//
	// - **lindorm.g.4xlarge**: 16 cores and 64 GB of memory.
	//
	// - **lindorm.c.8xlarge**: 32 cores and 64 GB of memory.
	//
	// - **lindorm.g.8xlarge**: 32 cores and 128 GB of memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// lindorm.g.2xlarge
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	// The name of the node group. **This parameter is required.*	- You can obtain the name by calling the GetLindormV2Instance operation.
	//
	// example:
	//
	// groupName
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s UpdateLindormV2InstanceRequestEngineListNodeGroupList) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceRequestEngineListNodeGroupList) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskSize() *int32 {
	return s.NodeDiskSize
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskType() *string {
	return s.NodeDiskType
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetGroupId(v string) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.GroupId = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeCount(v int32) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskSize(v int32) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskType(v string) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeSpec(v string) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) SetResourceGroupName(v string) *UpdateLindormV2InstanceRequestEngineListNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

func (s *UpdateLindormV2InstanceRequestEngineListNodeGroupList) Validate() error {
	return dara.Validate(s)
}
