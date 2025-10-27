// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBResourceGroupShrinkRequest
	GetClientToken() *string
	SetClusterMode(v string) *CreateDBResourceGroupShrinkRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *CreateDBResourceGroupShrinkRequest
	GetDBClusterId() *string
	SetEngine(v string) *CreateDBResourceGroupShrinkRequest
	GetEngine() *string
	SetEngineParamsShrink(v string) *CreateDBResourceGroupShrinkRequest
	GetEngineParamsShrink() *string
	SetGroupName(v string) *CreateDBResourceGroupShrinkRequest
	GetGroupName() *string
	SetGroupType(v string) *CreateDBResourceGroupShrinkRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *CreateDBResourceGroupShrinkRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetMaxComputeResource() *string
	SetMinClusterCount(v int32) *CreateDBResourceGroupShrinkRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetMinComputeResource() *string
	SetNodeNum(v int32) *CreateDBResourceGroupShrinkRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *CreateDBResourceGroupShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBResourceGroupShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDBResourceGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBResourceGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateDBResourceGroupShrinkRequest struct {
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
	// The engine of the resource group. Valid values:
	//
	// 	- **AnalyticDB*	- (default)
	//
	// 	- **SparkWarehouse**
	//
	// example:
	//
	// AnalyticDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The Spark application configuration parameters that can be applied to all Spark jobs executed in the resource group. If you want to configure parameters for a specific Spark job, you can specify values for the parameters in the code editor when you submit the job.
	//
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParamsShrink *string `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The name of the resource group.
	//
	// 	- The name can be up to 255 characters in length.
	//
	// 	- The name must start with an uppercase letter or a digit.
	//
	// 	- The name can contain uppercase letters, digits, hyphens (-), and underscores (_).
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
	// example:
	//
	// interactive
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The maximum number of compute clusters that are allowed in the resource group. Maximum value: 10.
	//
	// example:
	//
	// 2
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources, which refers to the amount of resources that are not allocated in the cluster. Unit: ACU.
	//
	// 	- When GroupType is set to interactive, set this parameter to a value in increments of 16ACU.
	//
	// 	- When GroupType is set to job, set this parameter to a value in increments of 8ACU.
	//
	// example:
	//
	// 32ACU
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
	// 2
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBResourceGroupShrinkRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateDBResourceGroupShrinkRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBResourceGroupShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBResourceGroupShrinkRequest) GetEngineParamsShrink() *string {
	return s.EngineParamsShrink
}

func (s *CreateDBResourceGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDBResourceGroupShrinkRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *CreateDBResourceGroupShrinkRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *CreateDBResourceGroupShrinkRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *CreateDBResourceGroupShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBResourceGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBResourceGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBResourceGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBResourceGroupShrinkRequest) SetClientToken(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterMode(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetDBClusterId(v string) *CreateDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEngine(v string) *CreateDBResourceGroupShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEngineParamsShrink(v string) *CreateDBResourceGroupShrinkRequest {
	s.EngineParamsShrink = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupName(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupType(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetNodeNum(v int32) *CreateDBResourceGroupShrinkRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetOwnerAccount(v string) *CreateDBResourceGroupShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetOwnerId(v int64) *CreateDBResourceGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetResourceOwnerAccount(v string) *CreateDBResourceGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetResourceOwnerId(v int64) *CreateDBResourceGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
