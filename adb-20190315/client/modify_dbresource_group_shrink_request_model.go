// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBResourceGroupShrinkRequest
	GetClientToken() *string
	SetClusterMode(v string) *ModifyDBResourceGroupShrinkRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *ModifyDBResourceGroupShrinkRequest
	GetDBClusterId() *string
	SetEngineParamsShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetEngineParamsShrink() *string
	SetGroupName(v string) *ModifyDBResourceGroupShrinkRequest
	GetGroupName() *string
	SetGroupType(v string) *ModifyDBResourceGroupShrinkRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetMaxComputeResource() *string
	SetMinClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetMinComputeResource() *string
	SetNodeNum(v int32) *ModifyDBResourceGroupShrinkRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *ModifyDBResourceGroupShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest
	GetOwnerId() *int64
	SetPoolUserListShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetPoolUserListShrink() *string
	SetResourceOwnerAccount(v string) *ModifyDBResourceGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyDBResourceGroupShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-t7241****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The working mode of the resource group. Valid values:
	//
	// 	- **Disable*	- (default)
	//
	// 	- **AutoScale**
	//
	// example:
	//
	// AutoScale
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The resource specifications of a single compute cluster. Unit: ACU.
	//
	// example:
	//
	// 16ACU
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Spark application configuration parameters that can be applied to all Spark jobs executed in the resource group. If you want to configure parameters for a specific Spark job, you can specify values for the parameters in the code editor when you submit the job.
	//
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParamsShrink *string `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution mode. Valid values:
	//
	// 	- **interactive*	- (default)
	//
	// 	- **batch**
	//
	// 	- **job**
	//
	// >  For more information, see [Query execution modes](https://help.aliyun.com/document_detail/189502.html).
	//
	// example:
	//
	// batch
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The maximum number of compute clusters that are allowed in the resource group. Maximum value: 10.
	//
	// example:
	//
	// 4
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources, which refers to the amount of resources that are not allocated in the cluster. Unit: ACU.
	//
	// 	- When GroupType is set to interactive, set this parameter to a value in increments of 16ACU.
	//
	// 	- When GroupType is set to job, set this parameter to a value in increments of 8ACU.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// The minimum number of compute clusters that are required in the resource group. Minimum value: 1.
	//
	// example:
	//
	// 1
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: AnalyticDB compute unit (ACU).
	//
	// 	- When GroupType is set to interactive, set this parameter to 16ACU.
	//
	// 	- When GroupType is set to job, set this parameter to 0ACU.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The number of nodes. Default value: 0.
	//
	// 	- Each node is configured with the resources of 16 cores and 64 GB memory.
	//
	// 	- Make sure that the amount of resources of the nodes (Number of nodes × 16 cores and 64 GB memory) is less than or equal to the amount of unused resources of the cluster.
	//
	// example:
	//
	// 1
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The database accounts with which to associate the resource group. They can be standard accounts or privileged accounts.
	PoolUserListShrink   *string `json:"PoolUserList,omitempty" xml:"PoolUserList,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBResourceGroupShrinkRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *ModifyDBResourceGroupShrinkRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetEngineParamsShrink() *string {
	return s.EngineParamsShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ModifyDBResourceGroupShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetPoolUserListShrink() *string {
	return s.PoolUserListShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClientToken(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterMode(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetDBClusterId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetEngineParamsShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.EngineParamsShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupName(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupType(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetNodeNum(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetOwnerAccount(v string) *ModifyDBResourceGroupShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetPoolUserListShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.PoolUserListShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
