// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudStorageSize(v int64) *ModifyLindormV2InstanceRequest
	GetCloudStorageSize() *int64
	SetCloudStorageType(v string) *ModifyLindormV2InstanceRequest
	GetCloudStorageType() *string
	SetEngineType(v string) *ModifyLindormV2InstanceRequest
	GetEngineType() *string
	SetInstanceId(v string) *ModifyLindormV2InstanceRequest
	GetInstanceId() *string
	SetNodeGroupList(v []*ModifyLindormV2InstanceRequestNodeGroupList) *ModifyLindormV2InstanceRequest
	GetNodeGroupList() []*ModifyLindormV2InstanceRequestNodeGroupList
	SetOwnerAccount(v string) *ModifyLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLindormV2InstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLindormV2InstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyLindormV2InstanceRequest
	GetSecurityToken() *string
	SetUpgradeType(v string) *ModifyLindormV2InstanceRequest
	GetUpgradeType() *string
}

type ModifyLindormV2InstanceRequest struct {
	CloudStorageSize *int64  `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	EngineType       *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is required.
	InstanceId    *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeGroupList []*ModifyLindormV2InstanceRequestNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
	OwnerAccount  *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s ModifyLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceRequest) GetCloudStorageSize() *int64 {
	return s.CloudStorageSize
}

func (s *ModifyLindormV2InstanceRequest) GetCloudStorageType() *string {
	return s.CloudStorageType
}

func (s *ModifyLindormV2InstanceRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ModifyLindormV2InstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyLindormV2InstanceRequest) GetNodeGroupList() []*ModifyLindormV2InstanceRequestNodeGroupList {
	return s.NodeGroupList
}

func (s *ModifyLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLindormV2InstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyLindormV2InstanceRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *ModifyLindormV2InstanceRequest) SetCloudStorageSize(v int64) *ModifyLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetCloudStorageType(v string) *ModifyLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetEngineType(v string) *ModifyLindormV2InstanceRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetInstanceId(v string) *ModifyLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetNodeGroupList(v []*ModifyLindormV2InstanceRequestNodeGroupList) *ModifyLindormV2InstanceRequest {
	s.NodeGroupList = v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetOwnerAccount(v string) *ModifyLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetOwnerId(v int64) *ModifyLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetRegionId(v string) *ModifyLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *ModifyLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetResourceOwnerId(v int64) *ModifyLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetSecurityToken(v string) *ModifyLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetUpgradeType(v string) *ModifyLindormV2InstanceRequest {
	s.UpgradeType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyLindormV2InstanceRequestNodeGroupList struct {
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NodeCount         *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeDiskSize      *int64  `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType      *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	NodeSpec          *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ModifyLindormV2InstanceRequestNodeGroupList) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceRequestNodeGroupList) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetNodeCount() *string {
	return s.NodeCount
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetNodeDiskSize() *int64 {
	return s.NodeDiskSize
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetNodeDiskType() *string {
	return s.NodeDiskType
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetGroupId(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.GroupId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeCount(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeDiskSize(v int64) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeDiskType(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeSpec(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetResourceGroupName(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) Validate() error {
	return dara.Validate(s)
}
