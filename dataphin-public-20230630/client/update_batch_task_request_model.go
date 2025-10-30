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
	SetUpdateCommand(v *UpdateBatchTaskRequestUpdateCommand) *UpdateBatchTaskRequest
	GetUpdateCommand() *UpdateBatchTaskRequestUpdateCommand
}

type UpdateBatchTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *UpdateBatchTaskRequest) GetUpdateCommand() *UpdateBatchTaskRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBatchTaskRequest) SetOpTenantId(v int64) *UpdateBatchTaskRequest {
	s.OpTenantId = &v
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
	// This parameter is required.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression       *string                                                  `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	CustomScheduleConfig *UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xx测试
	NodeDescription    *string   `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NodeStatus *int32                                          `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	ParamList  []*UpdateBatchTaskRequestUpdateCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId        *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// example:
	//
	// DAILY
	SchedulePeriod  *string                                             `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	SparkClientInfo *UpdateBatchTaskRequestUpdateCommandSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 21
	TaskType     *int32                                             `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpStreamList []*UpdateBatchTaskRequestUpdateCommandUpStreamList `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
}

func (s UpdateBatchTaskRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetCode() *string {
	return s.Code
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

func (s *UpdateBatchTaskRequestUpdateCommand) GetEngine() *string {
	return s.Engine
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetFileId() *int64 {
	return s.FileId
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

func (s *UpdateBatchTaskRequestUpdateCommand) GetParamList() []*UpdateBatchTaskRequestUpdateCommandParamList {
	return s.ParamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetSparkClientInfo() *UpdateBatchTaskRequestUpdateCommandSparkClientInfo {
	return s.SparkClientInfo
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetTaskType() *int32 {
	return s.TaskType
}

func (s *UpdateBatchTaskRequestUpdateCommand) GetUpStreamList() []*UpdateBatchTaskRequestUpdateCommandUpStreamList {
	return s.UpStreamList
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetCode(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.Code = &v
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

func (s *UpdateBatchTaskRequestUpdateCommand) SetEngine(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.Engine = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetFileId(v int64) *UpdateBatchTaskRequestUpdateCommand {
	s.FileId = &v
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

func (s *UpdateBatchTaskRequestUpdateCommand) SetParamList(v []*UpdateBatchTaskRequestUpdateCommandParamList) *UpdateBatchTaskRequestUpdateCommand {
	s.ParamList = v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetPriority(v int32) *UpdateBatchTaskRequestUpdateCommand {
	s.Priority = &v
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

func (s *UpdateBatchTaskRequestUpdateCommand) SetSchedulePeriod(v string) *UpdateBatchTaskRequestUpdateCommand {
	s.SchedulePeriod = &v
	return s
}

func (s *UpdateBatchTaskRequestUpdateCommand) SetSparkClientInfo(v *UpdateBatchTaskRequestUpdateCommandSparkClientInfo) *UpdateBatchTaskRequestUpdateCommand {
	s.SparkClientInfo = v
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

func (s *UpdateBatchTaskRequestUpdateCommand) Validate() error {
	if s.CustomScheduleConfig != nil {
		if err := s.CustomScheduleConfig.Validate(); err != nil {
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

type UpdateBatchTaskRequestUpdateCommandCustomScheduleConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 20:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
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

type UpdateBatchTaskRequestUpdateCommandParamList struct {
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	DependPeriod *UpdateBatchTaskRequestUpdateCommandUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// example:
	//
	// LAST
	DependStrategy *string   `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	FieldList      []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// example:
	//
	// PHYSICAL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PeriodDiff        *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	SourceNodeEnabled *bool  `json:"SourceNodeEnabled,omitempty" xml:"SourceNodeEnabled,omitempty"`
	// example:
	//
	// n_2001
	SourceNodeId *string `json:"SourceNodeId,omitempty" xml:"SourceNodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_input1
	SourceNodeOutputName *string `json:"SourceNodeOutputName,omitempty" xml:"SourceNodeOutputName,omitempty"`
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
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
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
