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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskInfo  *GetBatchTaskInfoByVersionResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression       *string                                                            `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	CustomScheduleConfig *GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
	// example:
	//
	// dag_102121211
	DagId *string `json:"DagId,omitempty" xml:"DagId,omitempty"`
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
	// 12113111
	FileId     *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	HasDevNode *bool  `json:"HasDevNode,omitempty" xml:"HasDevNode,omitempty"`
	// example:
	//
	// 测试任务1
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedPublish *bool   `json:"NeedPublish,omitempty" xml:"NeedPublish,omitempty"`
	// example:
	//
	// xx测试
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// example:
	//
	// openapi
	NodeFrom *string `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
	// example:
	//
	// n_1011_21232132322
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 测试任务1
	NodeName           *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeOutputNameList []*string `json:"NodeOutputNameList,omitempty" xml:"NodeOutputNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NodeStatus *int32 `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// example:
	//
	// 30231123
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30231123
	OwnerUserId *string                                                   `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	ParamList   []*GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	Paused      *bool                                                     `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Published *bool  `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// test xx
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Rerunable *bool   `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// example:
	//
	// 3
	ScheduleType    *int32                                                        `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	SparkClientInfo *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// example:
	//
	// 测试任务1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 21
	TaskType     *int32                                                       `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetHasDevNode() *bool {
	return s.HasDevNode
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetOwnerUserId() *string {
	return s.OwnerUserId
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetRemark() *string {
	return s.Remark
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) GetRerunable() *bool {
	return s.Rerunable
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetFileId(v int64) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetHasDevNode(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.HasDevNode = &v
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOwnerName(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OwnerName = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetOwnerUserId(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.OwnerUserId = &v
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

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetRemark(v string) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Remark = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponseBodyTaskInfo) SetRerunable(v bool) *GetBatchTaskInfoByVersionResponseBodyTaskInfo {
	s.Rerunable = &v
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

type GetBatchTaskInfoByVersionResponseBodyTaskInfoCustomScheduleConfig struct {
	// example:
	//
	// 20:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
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

type GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

type GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo struct {
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
	DependPeriod *GetBatchTaskInfoByVersionResponseBodyTaskInfoUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// example:
	//
	// LAST
	DependStrategy *string   `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	FieldList      []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// example:
	//
	// PHYSICAL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 1
	PeriodDiff        *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	SourceNodeEnabled *bool  `json:"SourceNodeEnabled,omitempty" xml:"SourceNodeEnabled,omitempty"`
	// example:
	//
	// n_2001
	SourceNodeId *string `json:"SourceNodeId,omitempty" xml:"SourceNodeId,omitempty"`
	// example:
	//
	// t_input1
	SourceNodeName *string `json:"SourceNodeName,omitempty" xml:"SourceNodeName,omitempty"`
	// example:
	//
	// t_input1
	SourceNodeOutputName *string `json:"SourceNodeOutputName,omitempty" xml:"SourceNodeOutputName,omitempty"`
	// example:
	//
	// 张三
	SourceNodeUserName *string `json:"SourceNodeUserName,omitempty" xml:"SourceNodeUserName,omitempty"`
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
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
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
