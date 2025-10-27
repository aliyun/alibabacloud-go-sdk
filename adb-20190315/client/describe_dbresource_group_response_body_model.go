// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBResourceGroupResponseBody
	GetDBClusterId() *string
	SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody
	GetGroupsInfo() []*DescribeDBResourceGroupResponseBodyGroupsInfo
	SetRequestId(v string) *DescribeDBResourceGroupResponseBody
	GetRequestId() *string
}

type DescribeDBResourceGroupResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried resource group.
	GroupsInfo []*DescribeDBResourceGroupResponseBodyGroupsInfo `json:"GroupsInfo,omitempty" xml:"GroupsInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBResourceGroupResponseBody) GetGroupsInfo() []*DescribeDBResourceGroupResponseBodyGroupsInfo {
	return s.GroupsInfo
}

func (s *DescribeDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBResourceGroupResponseBody) SetDBClusterId(v string) *DescribeDBResourceGroupResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody {
	s.GroupsInfo = v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetRequestId(v string) *DescribeDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) Validate() error {
	if s.GroupsInfo != nil {
		for _, item := range s.GroupsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyGroupsInfo struct {
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
	// The endpoint of the resource group.
	//
	// >  This parameter is returned only when the value of Engine is SparkWarehouse.
	//
	// example:
	//
	// amv-bp1nw64y******.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2022-10-09 16:57:35.241
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The minimum amount of elastic computing resources. Unit: ACU.
	//
	// example:
	//
	// 16ACU
	ElasticMinComputeResource *string `json:"ElasticMinComputeResource,omitempty" xml:"ElasticMinComputeResource,omitempty"`
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
	// example:
	//
	// USER_DEFAULT
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
	// The database accounts that are associated with the resource group.
	GroupUserList []*string `json:"GroupUserList,omitempty" xml:"GroupUserList,omitempty" type:"Repeated"`
	// The database accounts that are associated with the resource group. Multiple database accounts are separated by commas (,).
	//
	// example:
	//
	// testb,testc
	GroupUsers *string `json:"GroupUsers,omitempty" xml:"GroupUsers,omitempty"`
	// The maximum number of compute clusters that are allowed in the resource group. Maximum value: 10.
	//
	// example:
	//
	// 4
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources, which refers to the amount of resources that are not allocated in the cluster. Unit: ACU.
	//
	// 	- If the value of GroupType is **interactive**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 16ACU.
	//
	// 	- If the value of GroupType is **job**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 8ACU.
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
	// 	- If the value of GroupType is **interactive**, 16ACU is returned.
	//
	// 	- If the value of GroupType is **job**, 0ACU is returned.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The number of nodes. Each node provides 16 cores and 64 GB memory.
	//
	// example:
	//
	// 2
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The port number of the resource group.
	//
	// >  This parameter is returned only when the value of Engine is SparkWarehouse.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The number of compute clusters running in the resource group.
	//
	// example:
	//
	// 2
	RunningClusterCount *int32 `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Running**
	//
	// 	- **Scaling**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource group was updated.
	//
	// example:
	//
	// 2022-11-09 16:57:35.241
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetElasticMinComputeResource() *string {
	return s.ElasticMinComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupUserList() []*string {
	return s.GroupUserList
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupUsers() *string {
	return s.GroupUsers
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetRunningClusterCount() *int32 {
	return s.RunningClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterMode(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterMode = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterSizeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterSizeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetConnectionString(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetCreateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetElasticMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ElasticMinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEngine(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Engine = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEngineParams(v map[string]interface{}) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.EngineParams = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupUserList(v []*string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupUserList = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupUsers(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupUsers = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetNodeNum(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.NodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetPort(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRunningClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.RunningClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetStatus(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Status = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetUpdateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) Validate() error {
	return dara.Validate(s)
}
