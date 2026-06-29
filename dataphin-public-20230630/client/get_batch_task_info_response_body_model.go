// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTaskInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBatchTaskInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTaskInfoResponseBody
	GetSuccess() *bool
	SetTaskInfo(v *GetBatchTaskInfoResponseBodyTaskInfo) *GetBatchTaskInfoResponseBody
	GetTaskInfo() *GetBatchTaskInfoResponseBodyTaskInfo
}

type GetBatchTaskInfoResponseBody struct {
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
	TaskInfo *GetBatchTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s GetBatchTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTaskInfoResponseBody) GetTaskInfo() *GetBatchTaskInfoResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *GetBatchTaskInfoResponseBody) SetCode(v string) *GetBatchTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTaskInfoResponseBody) SetHttpStatusCode(v int32) *GetBatchTaskInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTaskInfoResponseBody) SetMessage(v string) *GetBatchTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTaskInfoResponseBody) SetRequestId(v string) *GetBatchTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBody) SetSuccess(v bool) *GetBatchTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTaskInfoResponseBody) SetTaskInfo(v *GetBatchTaskInfoResponseBodyTaskInfo) *GetBatchTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

func (s *GetBatchTaskInfoResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTaskInfoResponseBodyTaskInfo struct {
	// The node code.
	//
	// example:
	//
	// show tables;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cron expression for automatic scheduling. Refer to the Linux cron expression format.
	//
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The custom scheduling interval configuration.
	CustomScheduleConfig *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig `json:"CustomScheduleConfig,omitempty" xml:"CustomScheduleConfig,omitempty" type:"Struct"`
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
	// The user ID of the development owner.
	//
	// example:
	//
	// 30231123
	DevelopOwnerId *string `json:"DevelopOwnerId,omitempty" xml:"DevelopOwnerId,omitempty"`
	// The name of the development owner.
	//
	// example:
	//
	// 张三
	DevelopOwnerName *string `json:"DevelopOwnerName,omitempty" xml:"DevelopOwnerName,omitempty"`
	// The node ID in the directory tree.
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
	// The node source, indicating the organization or application that created the node.
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
	// The user ID of the O&M owner.
	//
	// example:
	//
	// 30231123
	OpsOwnerId *string `json:"OpsOwnerId,omitempty" xml:"OpsOwnerId,omitempty"`
	// The name of the O&M owner.
	//
	// example:
	//
	// 张三
	OpsOwnerName *string `json:"OpsOwnerName,omitempty" xml:"OpsOwnerName,omitempty"`
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
	ParamList []*GetBatchTaskInfoResponseBodyTaskInfoParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// Indicates whether the node is paused for scheduling.
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
	// Indicates whether the node is published.
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
	SparkClientInfo *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo `json:"SparkClientInfo,omitempty" xml:"SparkClientInfo,omitempty" type:"Struct"`
	// The submit status. Valid values:
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
	// The node type. For more information, refer to the create offline compute node operation.
	//
	// example:
	//
	// 21
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The upstream dependencies.
	UpStreamList []*GetBatchTaskInfoResponseBodyTaskInfoUpStreamList `json:"UpStreamList,omitempty" xml:"UpStreamList,omitempty" type:"Repeated"`
}

func (s GetBatchTaskInfoResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetCustomScheduleConfig() *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	return s.CustomScheduleConfig
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDagId() *string {
	return s.DagId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDevelopOwnerId() *string {
	return s.DevelopOwnerId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetDevelopOwnerName() *string {
	return s.DevelopOwnerName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetHasDevNode() *bool {
	return s.HasDevNode
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNeedPublish() *bool {
	return s.NeedPublish
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeFrom() *string {
	return s.NodeFrom
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeOutputNameList() []*string {
	return s.NodeOutputNameList
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetOpsOwnerId() *string {
	return s.OpsOwnerId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetOpsOwnerName() *string {
	return s.OpsOwnerName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetParamList() []*GetBatchTaskInfoResponseBodyTaskInfoParamList {
	return s.ParamList
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetPaused() *bool {
	return s.Paused
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetPublished() *bool {
	return s.Published
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetRemark() *string {
	return s.Remark
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetRerunable() *bool {
	return s.Rerunable
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetSparkClientInfo() *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo {
	return s.SparkClientInfo
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) GetUpStreamList() []*GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	return s.UpStreamList
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetCode(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Code = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetCronExpression(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.CronExpression = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetCustomScheduleConfig(v *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.CustomScheduleConfig = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDagId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DagId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDataSourceCatalog(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DataSourceCatalog = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDataSourceId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDataSourceSchema(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DataSourceSchema = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDevelopOwnerId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DevelopOwnerId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetDevelopOwnerName(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.DevelopOwnerName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetFileId(v int64) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetHasDevNode(v bool) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.HasDevNode = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetName(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Name = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNeedPublish(v bool) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NeedPublish = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeDescription(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeDescription = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeFrom(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeFrom = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeName(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeOutputNameList(v []*string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeOutputNameList = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetNodeStatus(v int32) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.NodeStatus = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetOperatorUserId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.OperatorUserId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetOpsOwnerId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.OpsOwnerId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetOpsOwnerName(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.OpsOwnerName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetOwnerName(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.OwnerName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetOwnerUserId(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.OwnerUserId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetParamList(v []*GetBatchTaskInfoResponseBodyTaskInfoParamList) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.ParamList = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetPaused(v bool) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Paused = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetPriority(v int32) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Priority = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetProjectId(v int64) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetPublished(v bool) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Published = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetRemark(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Remark = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetRerunable(v bool) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Rerunable = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetSchedulePeriod(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.SchedulePeriod = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetScheduleType(v int32) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetSparkClientInfo(v *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.SparkClientInfo = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetStatus(v string) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetTaskType(v int32) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.TaskType = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) SetUpStreamList(v []*GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) *GetBatchTaskInfoResponseBodyTaskInfo {
	s.UpStreamList = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfo) Validate() error {
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

type GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig struct {
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
	// The interval unit. Valid values:
	//
	// - MINUTE
	//
	// - HOUR.
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

func (s GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) SetEndTime(v string) *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	s.EndTime = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) SetInterval(v int32) *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	s.Interval = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) SetIntervalUnit(v string) *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	s.IntervalUnit = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) SetSchedulePeriod(v string) *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) SetStartTime(v string) *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig {
	s.StartTime = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoCustomScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoResponseBodyTaskInfoParamList struct {
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

func (s GetBatchTaskInfoResponseBodyTaskInfoParamList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfoParamList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoParamList) GetKey() *string {
	return s.Key
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoParamList) GetValue() *string {
	return s.Value
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoParamList) SetKey(v string) *GetBatchTaskInfoResponseBodyTaskInfoParamList {
	s.Key = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoParamList) SetValue(v string) *GetBatchTaskInfoResponseBodyTaskInfoParamList {
	s.Value = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoParamList) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo struct {
	// The Spark client version name.
	//
	// example:
	//
	// abc
	SparkClientVersion *string `json:"SparkClientVersion,omitempty" xml:"SparkClientVersion,omitempty"`
}

func (s GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) GetSparkClientVersion() *string {
	return s.SparkClientVersion
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) SetSparkClientVersion(v string) *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo {
	s.SparkClientVersion = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoSparkClientInfo) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskInfoResponseBodyTaskInfoUpStreamList struct {
	// The dependency period.
	DependPeriod *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod `json:"DependPeriod,omitempty" xml:"DependPeriod,omitempty" type:"Struct"`
	// The dependency strategy. Valid values: ALL, FIRST, LAST, NEAR.
	//
	// example:
	//
	// LAST
	DependStrategy *string `json:"DependStrategy,omitempty" xml:"DependStrategy,omitempty"`
	// The dependent logical table fields.
	FieldList []*string `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
	// The upstream dependency node type. Valid values:
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
	// The name of the input table.
	//
	// example:
	//
	// t_input1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
}

func (s GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetDependPeriod() *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod {
	return s.DependPeriod
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetDependStrategy() *string {
	return s.DependStrategy
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetFieldList() []*string {
	return s.FieldList
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetPeriodDiff() *int32 {
	return s.PeriodDiff
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceNodeEnabled() *bool {
	return s.SourceNodeEnabled
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceNodeId() *string {
	return s.SourceNodeId
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceNodeName() *string {
	return s.SourceNodeName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceNodeOutputName() *string {
	return s.SourceNodeOutputName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceNodeUserName() *string {
	return s.SourceNodeUserName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetDependPeriod(v *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.DependPeriod = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetDependStrategy(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.DependStrategy = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetFieldList(v []*string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.FieldList = v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetNodeType(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.NodeType = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetPeriodDiff(v int32) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.PeriodDiff = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceNodeEnabled(v bool) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceNodeEnabled = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceNodeId(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceNodeId = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceNodeName(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceNodeName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceNodeOutputName(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceNodeOutputName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceNodeUserName(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceNodeUserName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) SetSourceTableName(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList {
	s.SourceTableName = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamList) Validate() error {
	if s.DependPeriod != nil {
		if err := s.DependPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod struct {
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

func (s GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) GetPeriodOffset() *int32 {
	return s.PeriodOffset
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) SetPeriodOffset(v int32) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) SetPeriodType(v string) *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod {
	s.PeriodType = &v
	return s
}

func (s *GetBatchTaskInfoResponseBodyTaskInfoUpStreamListDependPeriod) Validate() error {
	return dara.Validate(s)
}
