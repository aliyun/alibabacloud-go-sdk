// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoByVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTaskInfoByVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBatchTaskInfoByVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTaskInfoByVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTaskInfoByVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTaskInfoByVersionResponseBody
	GetSuccess() *bool
	SetTaskInfo(v *GetBatchTaskInfoByVersionResponseBodyTaskInfo) *GetBatchTaskInfoByVersionResponseBody
	GetTaskInfo() *GetBatchTaskInfoByVersionResponseBodyTaskInfo
}

type GetBatchTaskInfoByVersionResponseBody struct {
	// The error code. OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message of the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The node details.
	TaskInfo *GetBatchTaskInfoByVersionResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s GetBatchTaskInfoByVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTaskInfoByVersionResponseBody) GetTaskInfo() *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetCode(v string) *GetBatchTaskInfoByVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetHttpStatusCode(v int32) *GetBatchTaskInfoByVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetMessage(v string) *GetBatchTaskInfoByVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetRequestId(v string) *GetBatchTaskInfoByVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetSuccess(v bool) *GetBatchTaskInfoByVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) SetTaskInfo(v *GetBatchTaskInfoByVersionResponseBodyTaskInfo) *GetBatchTaskInfoByVersionResponseBody {
	s.TaskInfo = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfo struct {
	// The code of the node.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the offline code template.
	//
	// example:
	//
	// st_123456
	CodeTemplateId *string `json:"CodeTemplateId,omitempty" xml:"CodeTemplateId,omitempty"`
	// The version of the offline code template.
	//
	// example:
	//
	// 1
	CodeTemplateVersion *int32 `json:"CodeTemplateVersion,omitempty" xml:"CodeTemplateVersion,omitempty"`
	// The current CPU resource configuration of the node. If queried by version, this is the configuration of the corresponding version.
	//
	// example:
	//
	// 0.5
	CpuResource *string `json:"CpuResource,omitempty" xml:"CpuResource,omitempty"`
	// The cron expression for automatic scheduling. You can refer to Linux cron expressions.
	//
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The custom scheduling interval configuration.
	CustomScheduleConfig *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
	// The ID of the DAG to which the node belongs.
	//
	// example:
	//
	// dag_102121211
	DagId *string `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The catalog of the data source. This parameter applies only to database SQL nodes and data source types that require a catalog, such as Presto.
	//
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// The ID of the data source for the database SQL node.
	//
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The schema of the data source. This parameter applies only to database SQL nodes and data source types that require a schema, such as Oracle.
	//
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// The Python connection information.
	DatabricksPythonConnectInfo *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo `json:"DatabricksPythonConnectInfo,omitempty" xml:"DatabricksPythonConnectInfo,omitempty" type:"Struct"`
	// The IDs of the developers.
	DevelopOwnerIdList []*string `json:"DevelopOwnerIdList,omitempty" xml:"DevelopOwnerIdList,omitempty" type:"Repeated"`
	// The names of the developers.
	DevelopOwnerNameList []*string `json:"DevelopOwnerNameList,omitempty" xml:"DevelopOwnerNameList,omitempty" type:"Repeated"`
	// The retry configuration for failed executions. This is the current configuration of the node. If queried by version, this is the configuration of the corresponding version.
	ExecuteRerunConfig *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig `json:"ExecuteRerunConfig,omitempty" xml:"ExecuteRerunConfig,omitempty" type:"Struct"`
	// The execution timeout configuration. This is the current configuration of the node. If queried by version, this is the configuration of the corresponding version.
	ExecuteTimeOutConfig *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig `json:"ExecuteTimeOutConfig,omitempty" xml:"ExecuteTimeOutConfig,omitempty" type:"Struct"`
	// The node ID of the task directory tree.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Indicates whether a development environment node exists.
	HasDevNode *bool `json:"HasDevNode,omitempty" xml:"HasDevNode,omitempty"`
	// The current CPU resource configuration of the node. If queried by version, this is the configuration of the corresponding version.
	//
	// example:
	//
	// 1024
	MemoryResource *string `json:"MemoryResource,omitempty" xml:"MemoryResource,omitempty"`
	// The task name.
	//
	// example:
	//
	// Test task 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether publishing is required.
	NeedPublish *bool `json:"NeedPublish,omitempty" xml:"NeedPublish,omitempty"`
	// The task description.
	//
	// example:
	//
	// xx test
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The source of the node, which indicates the organization or application that created the node.
	//
	// example:
	//
	// openapi
	NodeFrom *string `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_1011_21232132322
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// example:
	//
	// Test task 1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The list of node output names.
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// The node status. Valid values: 1: Normal. 2: Paused. 3: Dry run.
	//
	// example:
	//
	// 1
	NodeStatus *int32 `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The ID of the user who performs the current operation.
	//
	// example:
	//
	// 30231123
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// The list of O&M owner IDs.
	OpsOwnerIdList []*string `json:"OpsOwnerIdList,omitempty" xml:"OpsOwnerIdList,omitempty" type:"Repeated"`
	// The list of O&M owner names.
	OpsOwnerNameList []*string `json:"OpsOwnerNameList,omitempty" xml:"OpsOwnerNameList,omitempty" type:"Repeated"`
	// The name of the node owner.
	//
	// example:
	//
	// Zhang San
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The user ID of the node owner.
	//
	// example:
	//
	// 30231123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The PAI configuration.
	PaiConfig *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig `json:"PaiConfig,omitempty" xml:"PaiConfig,omitempty" type:"Struct"`
	// The list of custom parameters for the node.
	ParamList []*GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// Specifies whether node scheduling is paused.
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The task scheduling priority. Valid values: 1 to 9. A larger value indicates a lower priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether the task is published.
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// The list of third-party packages required by the Python task.
	PythonModuleList []*GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The Python version for the Python task, such as Python 3.7.
	//
	// example:
	//
	// Python 3.7
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test xx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Specifies whether the node can be rerun.
	Rerunable *bool `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	// Specifies whether to enable retry on timeout. This reflects the current configuration of the node. When queried by version, it reflects the configuration of the corresponding version.
	//
	// example:
	//
	// false
	RetryOnTimeoutEnable *bool `json:"RetryOnTimeoutEnable,omitempty" xml:"RetryOnTimeoutEnable,omitempty"`
	// The schedule period. Valid values:
	//
	// - YEARLY: Yearly.
	//
	// - MONTHLY: Monthly.
	//
	// - WEEKLY: Weekly.
	//
	// - DAILY: Daily.
	//
	// - HOURLY: Hourly.
	//
	// - MINUTELY: Minutely.
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The node type. Valid values:
	//
	// - 1: Periodic node.
	//
	// - 3: Manual node.
	//
	// example:
	//
	// 3
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The Spark client.
	SparkClientInfo *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// The current CPU resource configuration of the node. If queried by version, this is the configuration of the corresponding version.
	//
	// example:
	//
	// false
	SparkDynamicResourceSwitch *bool `json:"SparkDynamicResourceSwitch,omitempty" xml:"SparkDynamicResourceSwitch,omitempty"`
	// The effective scope of the Spark resource configuration. Valid values:
	//
	// - ONLY_PROD: Production environment only.
	//
	// - PROD_AND_DEV: Production and development environments.
	//
	// This reflects the current configuration of the node. When queried by version, it reflects the configuration of the corresponding version.
	//
	// example:
	//
	// ONLY_PROD
	SparkResourceEffectScope *string `json:"SparkResourceEffectScope,omitempty" xml:"SparkResourceEffectScope,omitempty"`
	// The Spark SQL resource configuration. The key is the configuration item name and the value is a string. The configuration is returned as a JSON string, such as {"spark.executor.memory":"2g"}. This reflects the current configuration of the node. When queried by version, it reflects the configuration of the corresponding version.
	//
	// example:
	//
	// {"spark.executor.memory":"2g"}
	SparkSqlResourceConfigMap *string `json:"SparkSqlResourceConfigMap,omitempty" xml:"SparkSqlResourceConfigMap,omitempty"`
	// The MaxCompute Spark version.
	//
	// example:
	//
	// Spark 3.1.1
	SparkVersion *string `json:"SparkVersion,omitempty" xml:"SparkVersion,omitempty"`
	// The publish status. Valid values:
	//
	// - 0: Draft.
	//
	// - 1: Submitted.
	//
	// - 100: In development.
	//
	// example:
	//
	// Test task 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type. For more information, see the API operation for creating an offline computing node.
	//
	// example:
	//
	// 21
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The upstream dependencies.
	UpStreamList []*GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCodeTemplateId() *string {
	return s.CodeTemplateId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCodeTemplateVersion() *int32 {
	return s.CodeTemplateVersion
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCpuResource() *string {
	return s.CpuResource
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetCustomScheduleConfig() *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	return s.CustomScheduleConfig
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDagId() *string {
	return s.DagId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDatabricksPythonConnectInfo() *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo {
	return s.DatabricksPythonConnectInfo
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDevelopOwnerIdList() []*string {
	return s.DevelopOwnerIdList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetDevelopOwnerNameList() []*string {
	return s.DevelopOwnerNameList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetExecuteRerunConfig() *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig {
	return s.ExecuteRerunConfig
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetExecuteTimeOutConfig() *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig {
	return s.ExecuteTimeOutConfig
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetHasDevNode() *bool {
	return s.HasDevNode
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetMemoryResource() *string {
	return s.MemoryResource
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNeedPublish() *bool {
	return s.NeedPublish
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeFrom() *string {
	return s.NodeFrom
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeOutputNameList() []*string {
	return s.NodeOutputNameList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOpsOwnerIdList() []*string {
	return s.OpsOwnerIdList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOpsOwnerNameList() []*string {
	return s.OpsOwnerNameList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPaiConfig() *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig {
	return s.PaiConfig
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetParamList() []*GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList {
	return s.ParamList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPaused() *bool {
	return s.Paused
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPublished() *bool {
	return s.Published
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPythonModuleList() []*GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList {
	return s.PythonModuleList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetRemark() *string {
	return s.Remark
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetRerunable() *bool {
	return s.Rerunable
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetRetryOnTimeoutEnable() *bool {
	return s.RetryOnTimeoutEnable
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSparkClientInfo() *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo {
	return s.SparkClientInfo
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSparkDynamicResourceSwitch() *bool {
	return s.SparkDynamicResourceSwitch
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSparkResourceEffectScope() *string {
	return s.SparkResourceEffectScope
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSparkSqlResourceConfigMap() *string {
	return s.SparkSqlResourceConfigMap
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetUpStreamList() []*GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	return s.UpStreamList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCode(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Code = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCodeTemplateId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.CodeTemplateId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCodeTemplateVersion(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.CodeTemplateVersion = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCpuResource(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.CpuResource = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCronExpression(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.CronExpression = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetCustomScheduleConfig(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.CustomScheduleConfig = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDagId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DagId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDataSourceCatalog(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DataSourceCatalog = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDataSourceId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDataSourceSchema(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DataSourceSchema = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDatabricksPythonConnectInfo(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DatabricksPythonConnectInfo = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDevelopOwnerIdList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DevelopOwnerIdList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetDevelopOwnerNameList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.DevelopOwnerNameList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetExecuteRerunConfig(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.ExecuteRerunConfig = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetExecuteTimeOutConfig(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.ExecuteTimeOutConfig = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetFileId(v int64) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetHasDevNode(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.HasDevNode = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetMemoryResource(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.MemoryResource = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Name = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNeedPublish(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NeedPublish = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeDescription(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeDescription = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeFrom(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeFrom = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeOutputNameList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeOutputNameList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetNodeStatus(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.NodeStatus = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOperatorUserId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OperatorUserId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOpsOwnerIdList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OpsOwnerIdList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOpsOwnerNameList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OpsOwnerNameList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOwnerName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OwnerName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOwnerUserId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OwnerUserId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPaiConfig(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.PaiConfig = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetParamList(v []*GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.ParamList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPaused(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Paused = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPriority(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Priority = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetProjectId(v int64) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPublished(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Published = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPythonModuleList(v []*GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.PythonModuleList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetPythonVersion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.PythonVersion = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetRemark(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Remark = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetRerunable(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Rerunable = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetRetryOnTimeoutEnable(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.RetryOnTimeoutEnable = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSchedulePeriod(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SchedulePeriod = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetScheduleType(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSparkClientInfo(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SparkClientInfo = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSparkDynamicResourceSwitch(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SparkDynamicResourceSwitch = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSparkResourceEffectScope(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SparkResourceEffectScope = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSparkSqlResourceConfigMap(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SparkSqlResourceConfigMap = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetSparkVersion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.SparkVersion = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetStatus(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetTaskType(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.TaskType = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetUpStreamList(v []*GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.UpStreamList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) Validate() error {
	if s.CustomScheduleConfig != nil {
		if err := s.CustomScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DatabricksPythonConnectInfo != nil {
		if err := s.DatabricksPythonConnectInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ExecuteRerunConfig != nil {
		if err := s.ExecuteRerunConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ExecuteTimeOutConfig != nil {
		if err := s.ExecuteTimeOutConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PaiConfig != nil {
		if err := s.PaiConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PythonModuleList != nil {
		for _, item := range s.PythonModuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SparkClientInfo != nil {
		if err := s.SparkClientInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UpStreamList != nil {
		for _, item := range s.UpStreamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig struct {
	// The end time in the HH:mm format.
	//
	// example:
	//
	// 20:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The custom interval.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The unit of the interval. Valid values: MINUTE and HOUR.
	//
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// The scheduling period.
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The start time in the HH:mm format.
	//
	// example:
	//
	// 08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) SetEndTime(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	s.EndTime = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) SetInterval(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	s.Interval = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) SetIntervalUnit(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	s.IntervalUnit = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) SetSchedulePeriod(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) SetStartTime(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig {
	s.StartTime = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo struct {
	// The HTTP path in the development environment.
	//
	// example:
	//
	// /sql/1.0/warehouses/dev
	DevHttpPath *string `json:"DevHttpPath,omitempty" xml:"DevHttpPath,omitempty"`
	// Specifies whether to enable Databricks Python Connect.
	//
	// example:
	//
	// true
	EnableDatabricksPythonConnect *bool `json:"EnableDatabricksPythonConnect,omitempty" xml:"EnableDatabricksPythonConnect,omitempty"`
	// The HTTP path in the production environment.
	//
	// example:
	//
	// /sql/1.0/warehouses/prod
	ProdHttpPath *string `json:"ProdHttpPath,omitempty" xml:"ProdHttpPath,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) GetDevHttpPath() *string {
	return s.DevHttpPath
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) GetEnableDatabricksPythonConnect() *bool {
	return s.EnableDatabricksPythonConnect
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) GetProdHttpPath() *string {
	return s.ProdHttpPath
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) SetDevHttpPath(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo {
	s.DevHttpPath = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) SetEnableDatabricksPythonConnect(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo {
	s.EnableDatabricksPythonConnect = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) SetProdHttpPath(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo {
	s.ProdHttpPath = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoDatabricksPythonConnectInfo) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig struct {
	// The custom retry interval for failed executions.
	//
	// example:
	//
	// 60
	CustomRerunInterval *string `json:"CustomRerunInterval,omitempty" xml:"CustomRerunInterval,omitempty"`
	// The custom number of retries for failed executions.
	//
	// example:
	//
	// 3
	CustomRerunTimes *string `json:"CustomRerunTimes,omitempty" xml:"CustomRerunTimes,omitempty"`
	// Specifies whether to follow the system retry configuration.
	//
	// example:
	//
	// false
	FollowSystem *bool `json:"FollowSystem,omitempty" xml:"FollowSystem,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) GetCustomRerunInterval() *string {
	return s.CustomRerunInterval
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) GetCustomRerunTimes() *string {
	return s.CustomRerunTimes
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) SetCustomRerunInterval(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig {
	s.CustomRerunInterval = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) SetCustomRerunTimes(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig {
	s.CustomRerunTimes = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) SetFollowSystem(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig {
	s.FollowSystem = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteRerunConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig struct {
	// The custom timeout period.
	//
	// example:
	//
	// 1800
	CustomValues *string `json:"CustomValues,omitempty" xml:"CustomValues,omitempty"`
	// Specifies whether to follow the system retry configuration.
	//
	// example:
	//
	// false
	FollowSystem *bool `json:"FollowSystem,omitempty" xml:"FollowSystem,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) GetCustomValues() *string {
	return s.CustomValues
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) SetCustomValues(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig {
	s.CustomValues = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) SetFollowSystem(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig {
	s.FollowSystem = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoExecuteTimeOutConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig struct {
	// The PAI experiment ID.
	//
	// example:
	//
	// 789012
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The PAI region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The PAI workspace ID.
	//
	// example:
	//
	// 123456
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" xml:"WorkSpaceId,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) GetRegion() *string {
	return s.Region
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) GetWorkSpaceId() *string {
	return s.WorkSpaceId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) SetExperimentId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig {
	s.ExperimentId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) SetRegion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig {
	s.Region = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) SetWorkSpaceId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig {
	s.WorkSpaceId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPaiConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList struct {
	// The parameter name.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) GetKey() *string {
	return s.Key
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) GetValue() *string {
	return s.Value
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) SetKey(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList {
	s.Key = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) SetValue(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList {
	s.Value = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList struct {
	// The ID of the third-party package.
	//
	// example:
	//
	// 1001
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// The name of the third-party package.
	//
	// example:
	//
	// pandas
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Python version required by the third-party package.
	//
	// example:
	//
	// Python 3.7
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
	// The version of the third-party package.
	//
	// example:
	//
	// 1.3.5
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) GetVersion() *string {
	return s.Version
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) SetModuleId(v int64) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList {
	s.ModuleId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) SetName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList {
	s.Name = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) SetPythonVersion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList {
	s.PythonVersion = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) SetVersion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList {
	s.Version = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoPythonModuleList) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo struct {
	// The name of the Spark client.
	//
	// example:
	//
	// abc
	SparkClientVersion *string `json:"SparkClientVersion,omitempty" xml:"SparkClientVersion,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) GetSparkClientVersion() *string {
	return s.SparkClientVersion
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) SetSparkClientVersion(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo {
	s.SparkClientVersion = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList struct {
	// The dependency period.
	DependPeriod *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// The dependency strategy. Valid values:
	//
	// - ALL: All.
	//
	// - FIRST: First.
	//
	// - LAST: Last.
	//
	// - NEAR: Nearest.
	//
	// example:
	//
	// LAST
	DependStrategy *string `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	// The dependent logical table fields.
	FieldList []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// The upstream dependency node type. Valid values:
	//
	// - PHYSICAL: Physical node.
	//
	// - LOGICAL: Logical table dependency.
	//
	// example:
	//
	// PHYSICAL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The period difference. A value of 0 indicates a same-period dependency. A positive integer indicates a dependency on the previous N periods.
	//
	// example:
	//
	// 1
	PeriodDiff *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	// Specifies whether the upstream node is enabled.
	SourceNodeEnabled *bool `json:"SourceNodeEnabled,omitempty" xml:"SourceNodeEnabled,omitempty"`
	// The ID of the upstream node.
	//
	// example:
	//
	// n_2001
	SourceNodeId *string `json:"SourceNodeId,omitempty" xml:"SourceNodeId,omitempty"`
	// The name of the upstream node.
	//
	// example:
	//
	// t_input1
	SourceNodeName *string `json:"SourceNodeName,omitempty" xml:"SourceNodeName,omitempty"`
	// The output name of the upstream node.
	//
	// example:
	//
	// t_input1
	SourceNodeOutputName *string `json:"SourceNodeOutputName,omitempty" xml:"SourceNodeOutputName,omitempty"`
	// The upstream username.
	//
	// example:
	//
	// Zhang San
	SourceNodeUserName *string `json:"SourceNodeUserName,omitempty" xml:"SourceNodeUserName,omitempty"`
	// The name of the input table.
	//
	// example:
	//
	// t_input1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetDependPeriod() *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod {
	return s.DependPeriod
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetDependStrategy() *string {
	return s.DependStrategy
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetFieldList() []*string {
	return s.FieldList
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetPeriodDiff() *int32 {
	return s.PeriodDiff
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceNodeEnabled() *bool {
	return s.SourceNodeEnabled
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceNodeId() *string {
	return s.SourceNodeId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceNodeName() *string {
	return s.SourceNodeName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceNodeOutputName() *string {
	return s.SourceNodeOutputName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceNodeUserName() *string {
	return s.SourceNodeUserName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetDependPeriod(v *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.DependPeriod = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetDependStrategy(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.DependStrategy = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetFieldList(v []*string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.FieldList = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetNodeType(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.NodeType = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetPeriodDiff(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.PeriodDiff = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceNodeEnabled(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceNodeEnabled = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceNodeId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceNodeId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceNodeName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceNodeName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceNodeOutputName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceNodeOutputName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceNodeUserName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceNodeUserName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) SetSourceTableName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList {
	s.SourceTableName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamList) Validate() error {
	if s.DependPeriod != nil {
		if err := s.DependPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod struct {
	// The period offset. This parameter is required when the PeriodType parameter is set to LAST_N_PERIOD.
	//
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
	// The dependency period type. Valid values:
	//
	// - CURRENT_PERIOD: Current period.
	//
	// - LAST_PERIOD: Last period.
	//
	// - LAST_N_PERIOD: Last N periods.
	//
	// - LAST_24_HOUR: Last 24 hours.
	//
	// example:
	//
	// CURRENT_PERIOD
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) GetPeriodOffset() *int32 {
	return s.PeriodOffset
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) SetPeriodOffset(v int32) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) SetPeriodType(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod {
	s.PeriodType = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod) Validate() error {
	return dara.Validate(s)
}
