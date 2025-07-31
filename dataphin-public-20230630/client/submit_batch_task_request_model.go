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
	SetSubmitCommand(v *SubmitBatchTaskRequestSubmitCommand) *SubmitBatchTaskRequest
	GetSubmitCommand() *SubmitBatchTaskRequestSubmitCommand
}

type SubmitBatchTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *SubmitBatchTaskRequest) GetSubmitCommand() *SubmitBatchTaskRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitBatchTaskRequest) SetOpTenantId(v int64) *SubmitBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitBatchTaskRequest) SetSubmitCommand(v *SubmitBatchTaskRequestSubmitCommand) *SubmitBatchTaskRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression       *string                                                  `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	CustomScheduleConfig *SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
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
	ParamList  []*SubmitBatchTaskRequestSubmitCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
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
	SparkClientInfo *SubmitBatchTaskRequestSubmitCommandSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	UpStreamList    []*SubmitBatchTaskRequestSubmitCommandUpStreamList  `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
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

func (s *SubmitBatchTaskRequestSubmitCommand) GetComment() *string {
	return s.Comment
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

func (s *SubmitBatchTaskRequestSubmitCommand) GetFileId() *int64 {
	return s.FileId
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

func (s *SubmitBatchTaskRequestSubmitCommand) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetSparkClientInfo() *SubmitBatchTaskRequestSubmitCommandSparkClientInfo {
	return s.SparkClientInfo
}

func (s *SubmitBatchTaskRequestSubmitCommand) GetUpStreamList() []*SubmitBatchTaskRequestSubmitCommandUpStreamList {
	return s.UpStreamList
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetCode(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Code = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetComment(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.Comment = &v
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

func (s *SubmitBatchTaskRequestSubmitCommand) SetFileId(v int64) *SubmitBatchTaskRequestSubmitCommand {
	s.FileId = &v
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

func (s *SubmitBatchTaskRequestSubmitCommand) SetSchedulePeriod(v string) *SubmitBatchTaskRequestSubmitCommand {
	s.SchedulePeriod = &v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetSparkClientInfo(v *SubmitBatchTaskRequestSubmitCommandSparkClientInfo) *SubmitBatchTaskRequestSubmitCommand {
	s.SparkClientInfo = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) SetUpStreamList(v []*SubmitBatchTaskRequestSubmitCommandUpStreamList) *SubmitBatchTaskRequestSubmitCommand {
	s.UpStreamList = v
	return s
}

func (s *SubmitBatchTaskRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandCustomScheduleConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 10:00
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

type SubmitBatchTaskRequestSubmitCommandParamList struct {
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

type SubmitBatchTaskRequestSubmitCommandSparkClientInfo struct {
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
	DependPeriod *SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// example:
	//
	// ALL
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
	return dara.Validate(s)
}

type SubmitBatchTaskRequestSubmitCommandUpStreamListDependPeriod struct {
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
