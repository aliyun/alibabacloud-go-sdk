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
	// The error code. A value of OK indicates that the request was successful.
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
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	// The node code.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cron expression for automatic scheduling. For more information, refer to the Linux cron expression syntax.
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
	// The catalog for database SQL nodes. This parameter takes effect only for data source types that require a catalog, such as Presto.
	//
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// The data source ID for database SQL nodes.
	//
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The schema for database SQL nodes. This parameter takes effect only for data source types that require a schema, such as Oracle.
	//
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// The node ID in the node directory tree.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Indicates whether the node has a development environment node.
	HasDevNode *bool `json:"HasDevNode,omitempty" xml:"HasDevNode,omitempty"`
	// The node name.
	//
	// example:
	//
	// 测试任务1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the node needs to be published.
	NeedPublish *bool `json:"NeedPublish,omitempty" xml:"NeedPublish,omitempty"`
	// The node description.
	//
	// example:
	//
	// xx测试
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The source of the node, indicating the organization or application that created the node.
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
	// 测试任务1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	// The user ID of the current operator.
	//
	// example:
	//
	// 30231123
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// The name of the node owner.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The user ID of the node owner.
	//
	// example:
	//
	// 30231123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The list of custom node parameters.
	ParamList []*GetBatchTaskInfoByVersionResponseBodyTaskInfoParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// Indicates whether the node scheduling is paused.
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The scheduling priority of the node. Valid values: 1 to 9. A larger value indicates a lower priority.
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
	// Indicates whether the node has been published.
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test xx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Indicates whether the node can be rerun.
	Rerunable *bool `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	// The scheduling period. Valid values:
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
	// - MINUTELY.
	//
	// example:
	//
	// DAILY
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The scheduling type. Valid values:
	//
	// - 1: periodic node.
	//
	// - 3: manual node.
	//
	// example:
	//
	// 3
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The Spark client information.
	SparkClientInfo *GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// The publish status. Valid values:
	//
	// - 0: draft.
	//
	// - 1: submitted.
	//
	// - 100: in development.
	//
	// example:
	//
	// 测试任务1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The node type. For more information, see the API operation for creating a batch task.
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
	// The end time in the format of HH:mm.
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
	// The interval unit. Valid values: MINUTE, HOUR.
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
	// The start time in the format of HH:mm.
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

type GetBatchTaskInfoByVersionResponseBodyTaskInfoSparkClientInfo struct {
	// The Spark client version.
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
	// The dependency strategy. Valid values: ALL, FIRST, LAST, NEAR.
	//
	// example:
	//
	// LAST
	DependStrategy *string `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	// The fields of the dependent logical table.
	FieldList []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// The type of the upstream dependency node. Valid values:
	//
	// - PHYSICAL: physical node.
	//
	// - LOGICAL: logical table dependency.
	//
	// example:
	//
	// PHYSICAL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The period difference. A value of 0 indicates a same-period dependency. A positive number indicates a dependency on the previous N periods.
	//
	// example:
	//
	// 1
	PeriodDiff *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	// Indicates whether the upstream node is enabled.
	SourceNodeEnabled *bool `json:"SourceNodeEnabled,omitempty" xml:"SourceNodeEnabled,omitempty"`
	// The upstream node ID.
	//
	// example:
	//
	// n_2001
	SourceNodeId *string `json:"SourceNodeId,omitempty" xml:"SourceNodeId,omitempty"`
	// The upstream node name.
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
	// The username of the upstream node owner.
	//
	// example:
	//
	// 张三
	SourceNodeUserName *string `json:"SourceNodeUserName,omitempty" xml:"SourceNodeUserName,omitempty"`
	// The input table name.
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
	// The period offset. This parameter is required when PeriodType is set to LAST_N_PERIOD.
	//
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
	// The dependency period type. Valid values:
	//
	// - CURRENT_PERIOD
	//
	// - LAST_PERIOD
	//
	// - LAST_N_PERIOD
	//
	// - LAST_24_HOUR.
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
