// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitBatchTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitBatchTaskRequest
	GetOpUserId() *string
	SetSubmitCommand(v *SubmitBatchTaskRequestSubmitCommand) *SubmitBatchTaskRequest
	GetSubmitCommand() *SubmitBatchTaskRequestSubmitCommand
}

type SubmitBatchTaskRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The edit request.
	//
	// This parameter is required.
	SubmitCommand *SubmitBatchTaskRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitBatchTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitBatchTaskRequest) GetSubmitCommand() *SubmitBatchTaskRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitBatchTaskRequest) SetOpTenantId(v int64) *SubmitBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitBatchTaskRequest) SetOpUserId(v string) *SubmitBatchTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitBatchTaskRequest) SetSubmitCommand(v *SubmitBatchTaskRequestSubmitCommand) *SubmitBatchTaskRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitBatchTaskRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchTaskRequestSubmitCommand struct {
	// The task code.
	//
	// This parameter is required.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The code template version number.
	//
	// example:
	//
	// 1
	CodeTemplateVersion *int32 `json:"CodeTemplateVersion,omitempty" xml:"CodeTemplateVersion,omitempty"`
	// The comment for the submit operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The conditional scheduling configuration. If this parameter is not specified, the current configuration is retained. If Enable is set to false, conditional scheduling is disabled and the original conditions are cleared. This parameter is supported only for periodic tasks.
	ConditionScheduleConfig *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig `json:"ConditionScheduleConfig,omitempty" xml:"ConditionScheduleConfig,omitempty" type:"Struct"`
	// The amount of CPU resources. If this parameter is not specified, the current configuration of the task is retained.
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
	CustomScheduleConfig *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
	// The task execution engine. For Python tasks, valid values:
	//
	// - PYTHON2_7
	//
	// - PYTHON3_7
	//
	// - PYTHON3_11
	//
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The failure retry configuration. If this parameter is not specified, the current configuration of the task is retained. If this parameter is specified, the entire configuration object is replaced.
	ExecuteRerunConfig *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig `json:"ExecuteRerunConfig,omitempty" xml:"ExecuteRerunConfig,omitempty" type:"Struct"`
	// The execution timeout configuration. If this parameter is not specified, the current configuration of the task is retained. If this parameter is specified, the entire configuration object is replaced.
	ExecuteTimeOutConfig *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig `json:"ExecuteTimeOutConfig,omitempty" xml:"ExecuteTimeOutConfig,omitempty" type:"Struct"`
	// The ID of the node in the task directory tree.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The amount of CPU resources. If this parameter is not specified, the current configuration of the task is retained.
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
	// The task description.
	//
	// example:
	//
	// xx test
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The list of node output names.
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// The node status. Valid values:
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
	// The ID of the offline code template.
	//
	// example:
	//
	// template_123456
	OfflineCodeTemplateId *string `json:"OfflineCodeTemplateId,omitempty" xml:"OfflineCodeTemplateId,omitempty"`
	// The list of offline code template parameters.
	OfflineCodeTemplateParams []*SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams `json:"OfflineCodeTemplateParams,omitempty" xml:"OfflineCodeTemplateParams,omitempty" type:"Repeated"`
	// The list of custom parameters.
	ParamList []*SubmitBatchTaskRequestSubmitCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The task scheduling priority. Valid values: 1 to 9. A larger value indicates a lower priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the project to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of Python third-party packages on which the node depends.
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The list of referenced resources. Each element is a resource name. If this parameter is not specified, the resource references of the original node are used. If a list is specified, the existing references are completely replaced. An empty list clears all references.
	ResourceReferenceList []*string `json:"ResourceReferenceList,omitempty" xml:"ResourceReferenceList,omitempty" type:"Repeated"`
	// Specifies whether to enable retry on timeout. If this parameter is not specified, the current configuration of the node is retained. If this parameter is explicitly set to false, the feature is disabled.
	//
	// example:
	//
	// false
	RetryOnTimeoutEnable *bool `json:"RetryOnTimeoutEnable,omitempty" xml:"RetryOnTimeoutEnable,omitempty"`
	// The scheduling period. Valid values:
	//
	// - YEARLY: yearly
	//
	// - MONTHLY: monthly
	//
	// - WEEKLY: weekly
	//
	// - DAILY: daily
	//
	// - HOURLY: hourly
	//
	// - MINUTELY: minutely
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The basic scheduling template. If this parameter is not specified, the current template is retained. An empty object unbinds the template. This parameter is supported only for periodic nodes.
	ScheduleTemplate *SubmitBatchTaskRequestSubmitCommandScheduleTemplate `json:"ScheduleTemplate,omitempty" xml:"ScheduleTemplate,omitempty" type:"Struct"`
	// The Spark client.
	SparkClientInfo *SubmitBatchTaskRequestSubmitCommandSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// Specifies whether to enable dynamic resources for Spark. If this parameter is not specified, the current configuration of the node is retained. If this parameter is explicitly set to false, the feature is disabled.
	//
	// example:
	//
	// false
	SparkDynamicResourceSwitch *bool `json:"SparkDynamicResourceSwitch,omitempty" xml:"SparkDynamicResourceSwitch,omitempty"`
	// The effective scope of the Spark resource configuration. ONLY_PROD indicates the production environment only. PROD_AND_DEV indicates both the production and development environments. If this parameter is not specified, the current configuration of the node is retained.
	//
	// example:
	//
	// ONLY_PROD
	SparkResourceEffectScope *string `json:"SparkResourceEffectScope,omitempty" xml:"SparkResourceEffectScope,omitempty"`
	// The Spark SQL resource configurations. The key is the configuration item name, and the value is a string. If this parameter is not specified, the current configurations of the node are retained. An empty map clears all configurations.
	SparkSqlResourceConfigMap map[string]*string `json:"SparkSqlResourceConfigMap,omitempty" xml:"SparkSqlResourceConfigMap,omitempty"`
	// The upstream dependencies.
	UpStreamList []*SubmitBatchTaskRequestSubmitCommandUpStreamList `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
}

func (s SubmitBatchTaskRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetCode() *string {
	return s.Code
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetCodeTemplateVersion() *int32 {
	return s.CodeTemplateVersion
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetComment() *string {
	return s.Comment
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetConditionScheduleConfig() *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig {
	return s.ConditionScheduleConfig
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetCpuResource() *string {
	return s.CpuResource
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetCronExpression() *string {
	return s.CronExpression
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetCustomScheduleConfig() *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	return s.CustomScheduleConfig
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetEngine() *string {
	return s.Engine
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetExecuteRerunConfig() *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig {
	return s.ExecuteRerunConfig
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetExecuteTimeOutConfig() *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig {
	return s.ExecuteTimeOutConfig
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetMemoryResource() *string {
	return s.MemoryResource
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetName() *string {
	return s.Name
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetNodeOutputNameList() []*string {
	return s.NodeOutputNameList
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetOfflineCodeTemplateId() *string {
	return s.OfflineCodeTemplateId
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetOfflineCodeTemplateParams() []*SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams {
	return s.OfflineCodeTemplateParams
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetParamList() []*SubmitBatchTaskRequestSubmitCommandParamList {
	return s.ParamList
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetResourceReferenceList() []*string {
	return s.ResourceReferenceList
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetRetryOnTimeoutEnable() *bool {
	return s.RetryOnTimeoutEnable
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetScheduleTemplate() *SubmitBatchTaskRequestSubmitCommandScheduleTemplate {
	return s.ScheduleTemplate
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSparkClientInfo() *SubmitBatchTaskRequestSubmitCommandSparkClientInfo {
	return s.SparkClientInfo
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSparkDynamicResourceSwitch() *bool {
	return s.SparkDynamicResourceSwitch
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSparkResourceEffectScope() *string {
	return s.SparkResourceEffectScope
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSparkSqlResourceConfigMap() map[string]*string {
	return s.SparkSqlResourceConfigMap
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetUpStreamList() []*SubmitBatchTaskRequestSubmitCommandUpStreamList {
	return s.UpStreamList
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCode(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Code = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCodeTemplateVersion(v int32) *SubmitBatchTaskRequestSubmitCommand {
	s.CodeTemplateVersion = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetComment(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Comment = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetConditionScheduleConfig(v *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) *SubmitBatchTaskRequestSubmitCommand {
	s.ConditionScheduleConfig = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCpuResource(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.CpuResource = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCronExpression(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.CronExpression = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCustomScheduleConfig(v *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) *SubmitBatchTaskRequestSubmitCommand {
	s.CustomScheduleConfig = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetEngine(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Engine = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetExecuteRerunConfig(v *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) *SubmitBatchTaskRequestSubmitCommand {
	s.ExecuteRerunConfig = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetExecuteTimeOutConfig(v *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) *SubmitBatchTaskRequestSubmitCommand {
	s.ExecuteTimeOutConfig = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetFileId(v int64) *SubmitBatchTaskRequestSubmitCommand {
	s.FileId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetMemoryResource(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.MemoryResource = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetName(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Name = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetNodeDescription(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.NodeDescription = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetNodeOutputNameList(v []*string) *SubmitBatchTaskRequestSubmitCommand {
	s.NodeOutputNameList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetNodeStatus(v int32) *SubmitBatchTaskRequestSubmitCommand {
	s.NodeStatus = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetOfflineCodeTemplateId(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.OfflineCodeTemplateId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetOfflineCodeTemplateParams(v []*SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) *SubmitBatchTaskRequestSubmitCommand {
	s.OfflineCodeTemplateParams = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetParamList(v []*SubmitBatchTaskRequestSubmitCommandParamList) *SubmitBatchTaskRequestSubmitCommand {
	s.ParamList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetPriority(v int32) *SubmitBatchTaskRequestSubmitCommand {
	s.Priority = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetProjectId(v int64) *SubmitBatchTaskRequestSubmitCommand {
	s.ProjectId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetPythonModuleList(v []*string) *SubmitBatchTaskRequestSubmitCommand {
	s.PythonModuleList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetResourceReferenceList(v []*string) *SubmitBatchTaskRequestSubmitCommand {
	s.ResourceReferenceList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetRetryOnTimeoutEnable(v bool) *SubmitBatchTaskRequestSubmitCommand {
	s.RetryOnTimeoutEnable = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSchedulePeriod(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.SchedulePeriod = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetScheduleTemplate(v *SubmitBatchTaskRequestSubmitCommandScheduleTemplate) *SubmitBatchTaskRequestSubmitCommand {
	s.ScheduleTemplate = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSparkClientInfo(v *SubmitBatchTaskRequestSubmitCommandSparkClientInfo) *SubmitBatchTaskRequestSubmitCommand {
	s.SparkClientInfo = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSparkDynamicResourceSwitch(v bool) *SubmitBatchTaskRequestSubmitCommand {
	s.SparkDynamicResourceSwitch = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSparkResourceEffectScope(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.SparkResourceEffectScope = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSparkSqlResourceConfigMap(v map[string]*string) *SubmitBatchTaskRequestSubmitCommand {
	s.SparkSqlResourceConfigMap = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetUpStreamList(v []*SubmitBatchTaskRequestSubmitCommandUpStreamList) *SubmitBatchTaskRequestSubmitCommand {
	s.UpStreamList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) Validate() error {
	if s.ConditionScheduleConfig != nil {
		if err := s.ConditionScheduleConfig.Validate(); err != nil {
			return err
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
	if s.OfflineCodeTemplateParams != nil {
		for _, item := range s.OfflineCodeTemplateParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.ScheduleTemplate != nil {
		if err := s.ScheduleTemplate.Validate(); err != nil {
			return err
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

type SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig struct {
	// Specifies whether to enable conditional scheduling. This parameter is required when the ConditionScheduleConfig object is passed and is verified by the backend. If this parameter is set to false, conditional scheduling is disabled and the previous configuration is cleared.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The list of custom conditions. This parameter is mutually exclusive with TemplateId. An empty list clears the previous conditions.
	ParamList []*SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The ID of the conditional scheduling template. This parameter is mutually exclusive with ParamList.
	//
	// example:
	//
	// 1002
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) GetEnable() *bool {
	return s.Enable
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) GetParamList() []*SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	return s.ParamList
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) SetEnable(v bool) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig {
	s.Enable = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) SetParamList(v []*SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig {
	s.ParamList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) SetTemplateId(v int64) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig {
	s.TemplateId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfig) Validate() error {
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList struct {
	// The condition name.
	//
	// example:
	//
	// workday
	ConditionName *string `json:"ConditionName,omitempty" xml:"ConditionName,omitempty"`
	// The cron expression used after a condition is matched.
	//
	// example:
	//
	// 0 0 5 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to enable the condition.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Specifies whether to follow the primary scheduling parameters.
	//
	// example:
	//
	// true
	FollowScheduleParam *bool `json:"FollowScheduleParam,omitempty" xml:"FollowScheduleParam,omitempty"`
	// The scheduling status used after a condition is matched. Valid values:
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
	// The JSON character string of the conditional expression tree. Internal key names use lower camel case.
	//
	// example:
	//
	// {}
	ScheduleConditionJson *string `json:"ScheduleConditionJson,omitempty" xml:"ScheduleConditionJson,omitempty"`
	// The scheduling time used after a condition is matched.
	//
	// example:
	//
	// 05:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetConditionName() *string {
	return s.ConditionName
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetCronExpression() *string {
	return s.CronExpression
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetEnable() *bool {
	return s.Enable
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetFollowScheduleParam() *bool {
	return s.FollowScheduleParam
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetScheduleConditionJson() *string {
	return s.ScheduleConditionJson
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetConditionName(v string) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.ConditionName = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetCronExpression(v string) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.CronExpression = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetEnable(v bool) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.Enable = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetFollowScheduleParam(v bool) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.FollowScheduleParam = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetNodeStatus(v int32) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.NodeStatus = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetScheduleConditionJson(v string) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.ScheduleConditionJson = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) SetScheduleTime(v string) *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList {
	s.ScheduleTime = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandConditionScheduleConfigParamList) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig struct {
	// The end time in the HH:mm format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00
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
	// - MINUTE: Minutes.
	//
	// - HOUR: Hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// The scheduling period. Valid values:
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
	// 10:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) SetEndTime(v string) *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	s.EndTime = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) SetInterval(v int32) *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	s.Interval = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) SetIntervalUnit(v string) *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	s.IntervalUnit = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) SetSchedulePeriod(v string) *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) SetStartTime(v string) *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig {
	s.StartTime = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig struct {
	// The custom retry interval for failures.
	//
	// example:
	//
	// 60
	CustomRerunInterval *string `json:"CustomRerunInterval,omitempty" xml:"CustomRerunInterval,omitempty"`
	// The custom number of retries for failures.
	//
	// example:
	//
	// 3
	CustomRerunTimes *string `json:"CustomRerunTimes,omitempty" xml:"CustomRerunTimes,omitempty"`
	// Specifies whether to follow the system configuration.
	//
	// example:
	//
	// false
	FollowSystem *bool `json:"FollowSystem,omitempty" xml:"FollowSystem,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) GetCustomRerunInterval() *string {
	return s.CustomRerunInterval
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) GetCustomRerunTimes() *string {
	return s.CustomRerunTimes
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) SetCustomRerunInterval(v string) *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig {
	s.CustomRerunInterval = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) SetCustomRerunTimes(v string) *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig {
	s.CustomRerunTimes = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) SetFollowSystem(v bool) *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig {
	s.FollowSystem = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteRerunConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig struct {
	// The custom timeout period.
	//
	// example:
	//
	// 1800
	CustomValues *string `json:"CustomValues,omitempty" xml:"CustomValues,omitempty"`
	// Specifies whether to follow the system configuration.
	//
	// example:
	//
	// false
	FollowSystem *bool `json:"FollowSystem,omitempty" xml:"FollowSystem,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) GetCustomValues() *string {
	return s.CustomValues
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) GetFollowSystem() *bool {
	return s.FollowSystem
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) SetCustomValues(v string) *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig {
	s.CustomValues = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) SetFollowSystem(v bool) *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig {
	s.FollowSystem = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandExecuteTimeOutConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams struct {
	// The parameter description.
	//
	// example:
	//
	// Parameter description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether encryption is enabled.
	EncryptEnabled *bool `json:"EncryptEnabled,omitempty" xml:"EncryptEnabled,omitempty"`
	// The parameter key.
	//
	// example:
	//
	// paramKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// paramValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) GetDescription() *string {
	return s.Description
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) GetEncryptEnabled() *bool {
	return s.EncryptEnabled
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) GetKey() *string {
	return s.Key
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) GetValue() *string {
	return s.Value
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) SetDescription(v string) *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams {
	s.Description = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) SetEncryptEnabled(v bool) *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams {
	s.EncryptEnabled = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) SetKey(v string) *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams {
	s.Key = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) SetValue(v string) *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams {
	s.Value = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandOfflineCodeTemplateParams) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandParamList struct {
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
	// key
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandParamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandParamList) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandParamList) GetKey() *string {
	return s.Key
}

func (s *SubmitBatchTaskRequestSubmitCommandParamList) GetValue() *string {
	return s.Value
}

func (s *SubmitBatchTaskRequestSubmitCommandParamList) SetKey(v string) *SubmitBatchTaskRequestSubmitCommandParamList {
	s.Key = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandParamList) SetValue(v string) *SubmitBatchTaskRequestSubmitCommandParamList {
	s.Value = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandParamList) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandScheduleTemplate struct {
	// The ID of the basic scheduling template. The value must be a positive number.
	//
	// example:
	//
	// 1001
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandScheduleTemplate) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandScheduleTemplate) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandScheduleTemplate) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *SubmitBatchTaskRequestSubmitCommandScheduleTemplate) SetTemplateId(v int64) *SubmitBatchTaskRequestSubmitCommandScheduleTemplate {
	s.TemplateId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandScheduleTemplate) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandSparkClientInfo struct {
	// The version name of the Spark client.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	SparkClientVersion *string `json:"SparkClientVersion,omitempty" xml:"SparkClientVersion,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandSparkClientInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandSparkClientInfo) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandSparkClientInfo) GetSparkClientVersion() *string {
	return s.SparkClientVersion
}

func (s *SubmitBatchTaskRequestSubmitCommandSparkClientInfo) SetSparkClientVersion(v string) *SubmitBatchTaskRequestSubmitCommandSparkClientInfo {
	s.SparkClientVersion = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandSparkClientInfo) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandUpStreamList struct {
	// The dependency period.
	DependPeriod *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// The dependency strategy. Valid values:
	//
	// - ALL: all
	//
	// - FIRST: first
	//
	// - LAST: last
	//
	// - NEAR: nearest
	//
	// example:
	//
	// ALL
	DependStrategy *string `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	// The dependent logical table fields.
	FieldList []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// The upstream dependency node type. Valid values:
	//
	// - PHYSICAL: physical node
	//
	// - LOGICAL: logical table dependency
	//
	// example:
	//
	// PHYSICAL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The period difference. A value of 0 indicates a dependency on the same period. A positive value indicates a dependency on the previous N periods.
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

func (s SubmitBatchTaskRequestSubmitCommandUpStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandUpStreamList) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetDependPeriod() *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod {
	return s.DependPeriod
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetDependStrategy() *string {
	return s.DependStrategy
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetFieldList() []*string {
	return s.FieldList
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetNodeType() *string {
	return s.NodeType
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetPeriodDiff() *int32 {
	return s.PeriodDiff
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetSourceNodeEnabled() *bool {
	return s.SourceNodeEnabled
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetSourceNodeId() *string {
	return s.SourceNodeId
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetSourceNodeOutputName() *string {
	return s.SourceNodeOutputName
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetDependPeriod(v *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.DependPeriod = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetDependStrategy(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.DependStrategy = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetFieldList(v []*string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.FieldList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetNodeType(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.NodeType = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetPeriodDiff(v int32) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.PeriodDiff = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetSourceNodeEnabled(v bool) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.SourceNodeEnabled = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetSourceNodeId(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.SourceNodeId = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetSourceNodeOutputName(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.SourceNodeOutputName = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) SetSourceTableName(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamList {
	s.SourceTableName = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamList) Validate() error {
	if s.DependPeriod != nil {
		if err := s.DependPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod struct {
	// The period offset. This parameter is required when dependencyPeriodType is set to LAST_N_PERIOD.
	//
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
	// The dependency period type. Valid values:
	//
	// - CURRENT_PERIOD: current period
	//
	// - LAST_PERIOD: last period
	//
	// - LAST_N_PERIOD: last N days
	//
	// - LAST_24_HOUR: last 24 hours
	//
	// This parameter is required.
	//
	// example:
	//
	// CURRENT_PERIOD
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) GetPeriodOffset() *int32 {
	return s.PeriodOffset
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) SetPeriodOffset(v int32) *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) SetPeriodType(v string) *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod {
	s.PeriodType = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod) Validate() error {
	return dara.Validate(s)
}
