// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBResourceGroupRequest
	GetClientToken() *string
	SetClusterMode(v string) *ModifyDBResourceGroupRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *ModifyDBResourceGroupRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *ModifyDBResourceGroupRequest
	GetDBClusterId() *string
	SetEngineParams(v map[string]interface{}) *ModifyDBResourceGroupRequest
	GetEngineParams() map[string]interface{}
	SetGroupName(v string) *ModifyDBResourceGroupRequest
	GetGroupName() *string
	SetGroupType(v string) *ModifyDBResourceGroupRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *ModifyDBResourceGroupRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *ModifyDBResourceGroupRequest
	GetMaxComputeResource() *string
	SetMinClusterCount(v int32) *ModifyDBResourceGroupRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *ModifyDBResourceGroupRequest
	GetMinComputeResource() *string
	SetNodeNum(v int32) *ModifyDBResourceGroupRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *ModifyDBResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBResourceGroupRequest
	GetOwnerId() *int64
	SetPoolUserList(v []*string) *ModifyDBResourceGroupRequest
	GetPoolUserList() []*string
	SetResourceOwnerAccount(v string) *ModifyDBResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBResourceGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyDBResourceGroupRequest struct {
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
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
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
	PoolUserList         []*string `json:"PoolUserList,omitempty" xml:"PoolUserList,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBResourceGroupRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *ModifyDBResourceGroupRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *ModifyDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBResourceGroupRequest) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *ModifyDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ModifyDBResourceGroupRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *ModifyDBResourceGroupRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *ModifyDBResourceGroupRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *ModifyDBResourceGroupRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *ModifyDBResourceGroupRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ModifyDBResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBResourceGroupRequest) GetPoolUserList() []*string {
	return s.PoolUserList
}

func (s *ModifyDBResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBResourceGroupRequest) SetClientToken(v string) *ModifyDBResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetClusterMode(v string) *ModifyDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetDBClusterId(v string) *ModifyDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetEngineParams(v map[string]interface{}) *ModifyDBResourceGroupRequest {
	s.EngineParams = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupName(v string) *ModifyDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupType(v string) *ModifyDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetNodeNum(v int32) *ModifyDBResourceGroupRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetOwnerId(v int64) *ModifyDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetPoolUserList(v []*string) *ModifyDBResourceGroupRequest {
	s.PoolUserList = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
