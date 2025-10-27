// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBResourceGroupRequest
	GetClientToken() *string
	SetClusterMode(v string) *CreateDBResourceGroupRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *CreateDBResourceGroupRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *CreateDBResourceGroupRequest
	GetDBClusterId() *string
	SetEngine(v string) *CreateDBResourceGroupRequest
	GetEngine() *string
	SetEngineParams(v map[string]interface{}) *CreateDBResourceGroupRequest
	GetEngineParams() map[string]interface{}
	SetGroupName(v string) *CreateDBResourceGroupRequest
	GetGroupName() *string
	SetGroupType(v string) *CreateDBResourceGroupRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *CreateDBResourceGroupRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *CreateDBResourceGroupRequest
	GetMaxComputeResource() *string
	SetMinClusterCount(v int32) *CreateDBResourceGroupRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *CreateDBResourceGroupRequest
	GetMinComputeResource() *string
	SetNodeNum(v int32) *CreateDBResourceGroupRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *CreateDBResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDBResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBResourceGroupRequest
	GetResourceOwnerId() *int64
}

type CreateDBResourceGroupRequest struct {
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
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
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

func (s CreateDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBResourceGroupRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateDBResourceGroupRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *CreateDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBResourceGroupRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBResourceGroupRequest) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *CreateDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDBResourceGroupRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *CreateDBResourceGroupRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *CreateDBResourceGroupRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *CreateDBResourceGroupRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *CreateDBResourceGroupRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *CreateDBResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBResourceGroupRequest) SetClientToken(v string) *CreateDBResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClusterMode(v string) *CreateDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetDBClusterId(v string) *CreateDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEngine(v string) *CreateDBResourceGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEngineParams(v map[string]interface{}) *CreateDBResourceGroupRequest {
	s.EngineParams = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupName(v string) *CreateDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupType(v string) *CreateDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetNodeNum(v int32) *CreateDBResourceGroupRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetOwnerAccount(v string) *CreateDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetOwnerId(v int64) *CreateDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceOwnerAccount(v string) *CreateDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceOwnerId(v int64) *CreateDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
