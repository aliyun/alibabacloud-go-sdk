// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateBatchTaskRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateBatchTaskRequestUpdateCommand) *UpdateBatchTaskRequest
	GetUpdateCommand() *UpdateBatchTaskRequestUpdateCommand
}

type UpdateBatchTaskRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update request.
	//
	// This parameter is required.
	UpdateCommand *UpdateBatchTaskRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBatchTaskRequest) GetUpdateCommand() *UpdateBatchTaskRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBatchTaskRequest) SetOpTenantId(v int64) *UpdateBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTaskRequest) SetOpUserId(v string) *UpdateBatchTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateBatchTaskRequest) SetUpdateCommand(v *UpdateBatchTaskRequestUpdateCommand) *UpdateBatchTaskRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBatchTaskRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBatchTaskRequestUpdateCommand struct {
	// The base schedule template ID. This parameter applies when the configuration method is set to schedule template. If this parameter is not specified and no other scheduling fields are specified, the original binding is retained.
	//
	// example:
	//
	// 7305621095333696
	BaseScheduleTemplateId *int64 `json:"BaseScheduleTemplateId,omitempty" xml:"BaseScheduleTemplateId,omitempty"`
	// The task code.
	//
	// This parameter is required.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Specifies whether to enable conditional scheduling.
	//
	// example:
	//
	// true
	ConditionScheduleEnable *bool `json:"ConditionScheduleEnable,omitempty" xml:"ConditionScheduleEnable,omitempty"`
	// The conditional scheduling configuration details. This parameter fully overwrites the existing configuration. If this parameter is not specified, the original configuration remains unchanged. This parameter is supported only for periodic tasks.
	ConditionScheduleParamList []*UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList `json:"ConditionScheduleParamList,omitempty" xml:"ConditionScheduleParamList,omitempty" type:"Repeated"`
	// The conditional scheduling template ID.
	//
	// example:
	//
	// 7305621095333697
	ConditionScheduleTemplateId *int64 `json:"ConditionScheduleTemplateId,omitempty" xml:"ConditionScheduleTemplateId,omitempty"`
	// The cross-node parameters, which are the output parameters declared by the current node for downstream nodes to reference. This parameter fully overwrites the existing parameters. If this parameter is not specified, the original parameters remain unchanged.
	ContextParamList []*UpdateBatchTaskRequestUpdateCommandContextParamList `json:"ContextParamList,omitempty" xml:"ContextParamList,omitempty" type:"Repeated"`
	// The amount of CPU resources.
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
	CustomScheduleConfig *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
	// The data source catalog for database SQL tasks. This parameter takes effect only for data source types that require a catalog, such as Presto.
	//
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// The data source ID for database SQL tasks.
	//
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source schema for database SQL tasks. This parameter takes effect only for data source types that require a schema, such as Oracle.
	//
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// The HTTP path in the development environment for Databricks tasks.
	//
	// example:
	//
	// /sql/protocolv1/o/xxx
	DevHttpPath *string `json:"DevHttpPath,omitempty" xml:"DevHttpPath,omitempty"`
	// The schedule resource group ID for development tasks.
	//
	// example:
	//
	// rg-def456
	DevResourceGroupId *string `json:"DevResourceGroupId,omitempty" xml:"DevResourceGroupId,omitempty"`
	// The list of lead developer IDs.
	DevelopOwnerIdList []*string `json:"DevelopOwnerIdList,omitempty" xml:"DevelopOwnerIdList,omitempty" type:"Repeated"`
	// The task execution engine, such as Python tasks. Valid values:
	//
	// - PYTHON2_7: Python 2.7.
	//
	// - PYTHON3_7: Python 3.7.
	//
	// - PYTHON3_11: Python 3.11.
	//
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The failure retry configuration. If not specified, the current configuration of the node is retained. If specified, the entire configuration object is replaced.
	ExecuteRerunConfig *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig `json:"ExecuteRerunConfig,omitempty" xml:"ExecuteRerunConfig,omitempty" type:"Struct"`
	// The execution timeout configuration. If not specified, the current configuration of the node is retained. If specified, the entire configuration object is replaced.
	ExecuteTimeOutConfig *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig `json:"ExecuteTimeOutConfig,omitempty" xml:"ExecuteTimeOutConfig,omitempty" type:"Struct"`
	// The ID of the node in the task directory tree.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The amount of memory resources.
	//
	// example:
	//
	// 1024
	MemoryResource *string `json:"MemoryResource,omitempty" xml:"MemoryResource,omitempty"`
	// The name of the offline task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the node.
	//
	// example:
	//
	// xx test
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The list of node output names.
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// The status of the node. Valid values:
	//
	// - 1: Normal.
	//
	// - 2: Paused.
	//
	// - 3: Dry run.
	//
	// example:
	//
	// 1
	NodeStatus *int32 `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The list of O&M owner IDs. If not specified, the O&M owners are not modified.
	OpsOwnerIdList []*string `json:"OpsOwnerIdList,omitempty" xml:"OpsOwnerIdList,omitempty" type:"Repeated"`
	// The list of custom parameters.
	ParamList []*UpdateBatchTaskRequestUpdateCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The scheduling priority of the node. Valid values: 1 to 9. A larger value indicates a lower priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The HTTP path in the production environment. This parameter is used for Databricks nodes.
	//
	// example:
	//
	// /sql/protocolv1/o/yyy
	ProdHttpPath *string `json:"ProdHttpPath,omitempty" xml:"ProdHttpPath,omitempty"`
	// The ID of the project to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The third-party Python packages that the node depends on.
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The ID of the scheduling resource group for production nodes.
	//
	// example:
	//
	// rg-abc123
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of referenced resources, specified by resource name.
	ResourceReferenceList []*string `json:"ResourceReferenceList,omitempty" xml:"ResourceReferenceList,omitempty" type:"Repeated"`
	// Specifies whether to retry the node upon timeout.
	//
	// example:
	//
	// false
	RetryOnTimeoutEnable *bool `json:"RetryOnTimeoutEnable,omitempty" xml:"RetryOnTimeoutEnable,omitempty"`
	// The scheduling cycle. Valid values:
	//
	// - YEARLY
	//
	// - MONTHLY
	//
	// - WEEKLY
	//
	// - DAILY
	//
	// - HOURLY
	//
	// - MINUTELY
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The Spark client.
	SparkClientInfo *UpdateBatchTaskRequestUpdateCommandSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// The switch for Spark dynamic resources.
	//
	// example:
	//
	// false
	SparkDynamicResourceSwitch *bool `json:"SparkDynamicResourceSwitch,omitempty" xml:"SparkDynamicResourceSwitch,omitempty"`
	// The effective scope of the Spark resource configuration.
	//
	// example:
	//
	// ONLY_PROD
	SparkResourceEffectScope *string `json:"SparkResourceEffectScope,omitempty" xml:"SparkResourceEffectScope,omitempty"`
	// The resource configuration for the Spark node. The value is a JSON string, such as {"spark.executor.memory":"2g"}.
	//
	// example:
	//
	// {"spark.executor.memory":"2g"}
	SparkSqlResourceConfigMap *string `json:"SparkSqlResourceConfigMap,omitempty" xml:"SparkSqlResourceConfigMap,omitempty"`
	// The list of node tag names. This parameter overwrites all existing tags. If not specified, the original tags remain unchanged.
	TaskTagList []*string `json:"TaskTagList,omitempty" xml:"TaskTagList,omitempty" type:"Repeated"`
	// The type of the node. Valid values:
	//
	// - 1: Hive_SQL.
	//
	// - 5: MaxCompute_SQL.
	//
	// - 10: Shell.
	//
	// - 21: Python.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The upstream dependencies.
	UpStreamList []*UpdateBatchTaskRequestUpdateCommandUpStreamList `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
	// The end date of the validity period. Format: yyyy-MM-dd. This parameter must be specified together with ValidStartDate.
	//
	// example:
	//
	// 9999-12-31
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// The start date of the validity period. Format: yyyy-MM-dd. This parameter must be specified together with ValidEndDate.
	//
	// example:
	//
	// 2026-01-01
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetBaseScheduleTemplateId() *int64 {
	return s.BaseScheduleTemplateId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetCode() *string {
	return s.Code
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetConditionScheduleEnable() *bool {
	return s.ConditionScheduleEnable
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetConditionScheduleParamList() []*UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	return s.ConditionScheduleParamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetConditionScheduleTemplateId() *int64 {
	return s.ConditionScheduleTemplateId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetContextParamList() []*UpdateBatchTaskRequestUpdateCommandContextParamList {
	return s.ContextParamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetCpuResource() *string {
	return s.CpuResource
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetCustomScheduleConfig() *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	return s.CustomScheduleConfig
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDevHttpPath() *string {
	return s.DevHttpPath
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDevResourceGroupId() *string {
	return s.DevResourceGroupId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetDevelopOwnerIdList() []*string {
	return s.DevelopOwnerIdList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetEngine() *string {
	return s.Engine
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetExecuteRerunConfig() *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig {
	return s.ExecuteRerunConfig
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetExecuteTimeOutConfig() *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig {
	return s.ExecuteTimeOutConfig
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetMemoryResource() *string {
	return s.MemoryResource
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetNodeOutputNameList() []*string {
	return s.NodeOutputNameList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetOpsOwnerIdList() []*string {
	return s.OpsOwnerIdList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetParamList() []*UpdateBatchTaskRequestUpdateCommandParamList {
	return s.ParamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetProdHttpPath() *string {
	return s.ProdHttpPath
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetResourceReferenceList() []*string {
	return s.ResourceReferenceList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetRetryOnTimeoutEnable() *bool {
	return s.RetryOnTimeoutEnable
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSparkClientInfo() *UpdateBatchTaskRequestUpdateCommandSparkClientInfo {
	return s.SparkClientInfo
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSparkDynamicResourceSwitch() *bool {
	return s.SparkDynamicResourceSwitch
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSparkResourceEffectScope() *string {
	return s.SparkResourceEffectScope
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSparkSqlResourceConfigMap() *string {
	return s.SparkSqlResourceConfigMap
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetTaskTagList() []*string {
	return s.TaskTagList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetTaskType() *int32 {
	return s.TaskType
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetUpStreamList() []*UpdateBatchTaskRequestUpdateCommandUpStreamList {
	return s.UpStreamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetBaseScheduleTemplateId(v int64) *UpdateBatchTaskRequestUpdateCommand {
	s.BaseScheduleTemplateId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetCode(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.Code = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetConditionScheduleEnable(v bool) *UpdateBatchTaskRequestUpdateCommand {
	s.ConditionScheduleEnable = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetConditionScheduleParamList(v []*UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) *UpdateBatchTaskRequestUpdateCommand {
	s.ConditionScheduleParamList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetConditionScheduleTemplateId(v int64) *UpdateBatchTaskRequestUpdateCommand {
	s.ConditionScheduleTemplateId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetContextParamList(v []*UpdateBatchTaskRequestUpdateCommandContextParamList) *UpdateBatchTaskRequestUpdateCommand {
	s.ContextParamList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetCpuResource(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.CpuResource = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetCronExpression(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.CronExpression = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetCustomScheduleConfig(v *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) *UpdateBatchTaskRequestUpdateCommand {
	s.CustomScheduleConfig = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDataSourceCatalog(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.DataSourceCatalog = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDataSourceId(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.DataSourceId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDataSourceSchema(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.DataSourceSchema = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDevHttpPath(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.DevHttpPath = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDevResourceGroupId(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.DevResourceGroupId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetDevelopOwnerIdList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.DevelopOwnerIdList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetEngine(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.Engine = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetExecuteRerunConfig(v *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) *UpdateBatchTaskRequestUpdateCommand {
	s.ExecuteRerunConfig = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetExecuteTimeOutConfig(v *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) *UpdateBatchTaskRequestUpdateCommand {
	s.ExecuteTimeOutConfig = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetFileId(v int64) *UpdateBatchTaskRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetMemoryResource(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.MemoryResource = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetName(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetNodeDescription(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.NodeDescription = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetNodeOutputNameList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.NodeOutputNameList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetNodeStatus(v int32) *UpdateBatchTaskRequestUpdateCommand {
	s.NodeStatus = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetOpsOwnerIdList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.OpsOwnerIdList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetParamList(v []*UpdateBatchTaskRequestUpdateCommandParamList) *UpdateBatchTaskRequestUpdateCommand {
	s.ParamList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetPriority(v int32) *UpdateBatchTaskRequestUpdateCommand {
	s.Priority = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetProdHttpPath(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.ProdHttpPath = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetProjectId(v int64) *UpdateBatchTaskRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetPythonModuleList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.PythonModuleList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetResourceGroupId(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetResourceReferenceList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.ResourceReferenceList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetRetryOnTimeoutEnable(v bool) *UpdateBatchTaskRequestUpdateCommand {
	s.RetryOnTimeoutEnable = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSchedulePeriod(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.SchedulePeriod = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSparkClientInfo(v *UpdateBatchTaskRequestUpdateCommandSparkClientInfo) *UpdateBatchTaskRequestUpdateCommand {
	s.SparkClientInfo = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSparkDynamicResourceSwitch(v bool) *UpdateBatchTaskRequestUpdateCommand {
	s.SparkDynamicResourceSwitch = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSparkResourceEffectScope(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.SparkResourceEffectScope = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSparkSqlResourceConfigMap(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.SparkSqlResourceConfigMap = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetTaskTagList(v []*string) *UpdateBatchTaskRequestUpdateCommand {
	s.TaskTagList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetTaskType(v int32) *UpdateBatchTaskRequestUpdateCommand {
	s.TaskType = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetUpStreamList(v []*UpdateBatchTaskRequestUpdateCommandUpStreamList) *UpdateBatchTaskRequestUpdateCommand {
	s.UpStreamList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetValidEndDate(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.ValidEndDate = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetValidStartDate(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.ValidStartDate = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) Validate() error {
	if s.ConditionScheduleParamList != nil {
		for _, item := range s.ConditionScheduleParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ContextParamList != nil {
		for _, item := range s.ContextParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomScheduleConfig != nil {
		if err := s.CustomScheduleConfig.Validate(); err != nil {
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
	if s.ParamList != nil {
		for _, item := range s.ParamList {
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

type UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList struct {
	// The condition name.
	//
	// example:
	//
	// condition1
	ConditionName *string `json:"ConditionName,omitempty" xml:"ConditionName,omitempty"`
	// The cron expression used after the condition is met.
	//
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to enable the condition.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Specifies whether to follow the main scheduling parameters.
	//
	// example:
	//
	// true
	FollowScheduleParam *bool `json:"FollowScheduleParam,omitempty" xml:"FollowScheduleParam,omitempty"`
	// The node scheduling status after the condition is met. Valid values:
	//
	// - 1: Normal.
	//
	// - 2: Paused.
	//
	// - 3: Dry run.
	//
	// example:
	//
	// 1
	NodeStatus *int32 `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The JSON string of the conditional expression tree. The keys use lower camel case characters and must be parsed by the invoker.
	//
	// example:
	//
	// {"type":"EXPRESSION","operator":"or"}
	ScheduleConditionJson *string `json:"ScheduleConditionJson,omitempty" xml:"ScheduleConditionJson,omitempty"`
	// The schedule time used after the condition is met.
	//
	// example:
	//
	// 01:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetConditionName() *string {
	return s.ConditionName
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetFollowScheduleParam() *bool {
	return s.FollowScheduleParam
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetScheduleConditionJson() *string {
	return s.ScheduleConditionJson
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetConditionName(v string) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.ConditionName = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetCronExpression(v string) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.CronExpression = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetEnable(v bool) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.Enable = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetFollowScheduleParam(v bool) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.FollowScheduleParam = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetNodeStatus(v int32) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.NodeStatus = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetScheduleConditionJson(v string) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.ScheduleConditionJson = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) SetScheduleTime(v string) *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandConditionScheduleParamList) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandContextParamList struct {
	// The default value.
	//
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Test parameter
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// param1
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandContextParamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandContextParamList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) GetDesc() *string {
	return s.Desc
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) GetParamKey() *string {
	return s.ParamKey
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) SetDefaultValue(v string) *UpdateBatchTaskRequestUpdateCommandContextParamList {
	s.DefaultValue = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) SetDesc(v string) *UpdateBatchTaskRequestUpdateCommandContextParamList {
	s.Desc = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) SetParamKey(v string) *UpdateBatchTaskRequestUpdateCommandContextParamList {
	s.ParamKey = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandContextParamList) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig struct {
	// The end time in the HH:mm format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The custom interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The interval unit. Valid values:
	//
	// - MINUTE: Minute.
	//
	// - HOUR: Hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// The scheduling cycle. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The start time in the HH:mm format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) SetEndTime(v string) *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	s.EndTime = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) SetInterval(v int32) *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	s.Interval = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) SetIntervalUnit(v string) *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	s.IntervalUnit = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) SetSchedulePeriod(v string) *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) SetStartTime(v string) *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig {
	s.StartTime = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig struct {
	// The custom retry interval.
	//
	// example:
	//
	// 60
	CustomRerunInterval *string `json:"CustomRerunInterval,omitempty" xml:"CustomRerunInterval,omitempty"`
	// The custom retry count.
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

func (s UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) GetCustomRerunInterval() *string {
	return s.CustomRerunInterval
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) GetCustomRerunTimes() *string {
	return s.CustomRerunTimes
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) SetCustomRerunInterval(v string) *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig {
	s.CustomRerunInterval = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) SetCustomRerunTimes(v string) *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig {
	s.CustomRerunTimes = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) SetFollowSystem(v bool) *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig {
	s.FollowSystem = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteRerunConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig struct {
	// The custom timeout period.
	//
	// example:
	//
	// 1800
	CustomValues *string `json:"CustomValues,omitempty" xml:"CustomValues,omitempty"`
	// Specifies whether to follow the system timeout configuration.
	//
	// example:
	//
	// false
	FollowSystem *bool `json:"FollowSystem,omitempty" xml:"FollowSystem,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) GetCustomValues() *string {
	return s.CustomValues
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) SetCustomValues(v string) *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig {
	s.CustomValues = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) SetFollowSystem(v bool) *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig {
	s.FollowSystem = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandExecuteTimeOutConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandParamList struct {
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandParamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandParamList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandParamList) GetKey() *string {
	return s.Key
}

func (s *UpdateBatchTaskRequestUpdateCommandParamList) GetValue() *string {
	return s.Value
}

func (s *UpdateBatchTaskRequestUpdateCommandParamList) SetKey(v string) *UpdateBatchTaskRequestUpdateCommandParamList {
	s.Key = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandParamList) SetValue(v string) *UpdateBatchTaskRequestUpdateCommandParamList {
	s.Value = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandParamList) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandSparkClientInfo struct {
	// The version name of the Spark client.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	SparkClientVersion *string `json:"SparkClientVersion,omitempty" xml:"SparkClientVersion,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandSparkClientInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandSparkClientInfo) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandSparkClientInfo) GetSparkClientVersion() *string {
	return s.SparkClientVersion
}

func (s *UpdateBatchTaskRequestUpdateCommandSparkClientInfo) SetSparkClientVersion(v string) *UpdateBatchTaskRequestUpdateCommandSparkClientInfo {
	s.SparkClientVersion = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandSparkClientInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskRequestUpdateCommandUpStreamList struct {
	// The dependency period.
	DependPeriod *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
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
	// The period difference. A value of 0 indicates a same-period dependency. A positive number indicates a dependency on the previous N periods.
	//
	// This parameter is required.
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
	// The output name of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_input1
	SourceNodeOutputName *string `json:"SourceNodeOutputName,omitempty" xml:"SourceNodeOutputName,omitempty"`
	// The name of the input table.
	//
	// example:
	//
	// t_input1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandUpStreamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandUpStreamList) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetDependPeriod() *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod {
	return s.DependPeriod
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetDependStrategy() *string {
	return s.DependStrategy
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetFieldList() []*string {
	return s.FieldList
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetNodeType() *string {
	return s.NodeType
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetPeriodDiff() *int32 {
	return s.PeriodDiff
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetSourceNodeEnabled() *bool {
	return s.SourceNodeEnabled
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetSourceNodeId() *string {
	return s.SourceNodeId
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetSourceNodeOutputName() *string {
	return s.SourceNodeOutputName
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetDependPeriod(v *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.DependPeriod = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetDependStrategy(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.DependStrategy = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetFieldList(v []*string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.FieldList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetNodeType(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.NodeType = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetPeriodDiff(v int32) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.PeriodDiff = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetSourceNodeEnabled(v bool) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.SourceNodeEnabled = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetSourceNodeId(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.SourceNodeId = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetSourceNodeOutputName(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.SourceNodeOutputName = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) SetSourceTableName(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamList {
	s.SourceTableName = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamList) Validate() error {
	if s.DependPeriod != nil {
		if err := s.DependPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod struct {
	// The period offset. This parameter is required when dependencyPeriodType is set to LAST_N_PERIOD.
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
	// - LAST_N_PERIOD: Last N days.
	//
	// - LAST_24_HOUR: Last 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// CURRENT_PERIOD
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) GetPeriodOffset() *int32 {
	return s.PeriodOffset
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) SetPeriodOffset(v int32) *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) SetPeriodType(v string) *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod {
	s.PeriodType = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod) Validate() error {
	return dara.Validate(s)
}
