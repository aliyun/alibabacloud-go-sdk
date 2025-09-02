// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetErrorCode(v string) *GetInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// The details of the instance.
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the instance fails to be scheduled.
	//
	// example:
	//
	// test
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetErrorCode(v string) *GetInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetErrorMessage(v string) *GetInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyData struct {
	// The baseline ID.
	//
	// example:
	//
	// 123123
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1590416703313
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The time when the instance started to wait for resources.
	//
	// example:
	//
	// 1590416703313
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The time when the instance started to wait to be scheduled.
	//
	// example:
	//
	// 1590416703313
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The data timestamp of the instance. In most cases, the value is one day before the time when the instance was run.
	//
	// example:
	//
	// 1590336000000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The connection string.
	//
	// example:
	//
	// odps_first
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The time when the instance was generated.
	//
	// example:
	//
	// 1590416703313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the instance.
	//
	// example:
	//
	// 111
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The scheduling time of the instance.
	//
	// example:
	//
	// 1590422400000
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 338450167
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The type of the workflow. Valid values:
	//
	// 	- DAILY: The workflow is used to run auto triggered nodes.
	//
	// 	- MANUAL: The workflow is used to run manually triggered nodes.
	//
	// 	- SMOKE_TEST: The workflow is used to perform smoke testing.
	//
	// 	- SUPPLY_DATA: The workflow is used to backfill data.
	//
	// example:
	//
	// DAILY
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The table and partition filter expression in Data Quality that are associated with the node.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// Indicates whether the instance is associated with a monitoring rule in Data Quality. Valid values:
	//
	// 	- 0: The instance is associated with a monitoring rule in Data Quality.
	//
	// 	- 1: The instance is not associated with a monitoring rule in Data Quality.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The time when the running of the instance was complete.
	//
	// example:
	//
	// 1590416703313
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 11713307578
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1590416703313
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 33115
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// kzh
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The owner of the instance.
	//
	// example:
	//
	// 111
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameters related to the node.
	//
	// example:
	//
	// bizdate=$bizdate tbods=$tbods tbdw=$tbdw tbpmic=$tbpmic tbpidx=$tbpidx tbptcif=$tbptcif
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The sequence number of the cycle. This parameter indicates the sequence number of the cycle of the instance on the current day.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The priority of the instance. Valid values: 1, 3, 5, 7, and 8. A greater value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workflow to which the node belongs.
	//
	// example:
	//
	// 123123
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The interval at which the node is rerun after the node fails to run. Unit: milliseconds.
	//
	// example:
	//
	// 60000
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Indicates whether the node that generates the instance can be rerun.
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The status of the node that generates the instance. Valid values:
	//
	// 	- NOT_RUN: The node is not run.
	//
	// 	- WAIT_TIME: The node is waiting for its scheduling time to arrive.
	//
	// 	- WAIT_RESOURCE: The node is waiting for resources.
	//
	// 	- RUNNING: The node is running.
	//
	// 	- CHECKING: Data quality is being checked for the node.
	//
	// 	- CHECKING_CONDITION: Branch conditions are being checked for the node.
	//
	// 	- FAILURE: The node fails to run.
	//
	// 	- SUCCESS: The node is successfully run.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of times the node can be rerun. The value of this parameter can be empty or an integer that is greater than or equal to 0.
	//
	// 	- If the value of this parameter is empty, the number of times that the node can be rerun is not specified.
	//
	// 	- If the value of this parameter is 0, the node cannot be rerun.
	//
	// 	- If the value of this parameter is a positive integer such as n, the node can still be rerun n times. For example, if the value of this parameter is 1, the node can still be rerun once. If the value of this parameter is 2, the node can still be rerun twice.
	//
	// example:
	//
	// 0
	TaskRerunTime *int32 `json:"TaskRerunTime,omitempty" xml:"TaskRerunTime,omitempty"`
	// The scheduling type of the node that generates the instance. Valid values:
	//
	// 	- NORMAL(0): The node is an auto triggered node. The scheduling system regularly runs the node.
	//
	// 	- MANUAL(1): The node is a manually triggered node. The scheduling system does not regularly run the node.
	//
	// 	- PAUSE(2): The node is a frozen node. The scheduling system regularly runs the node but sets the status of the node to failed when the scheduling system starts to run the node.
	//
	// 	- SKIP(3): The node is a dry-run node. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- SKIP_UNCHOOSE(4): The node is an unselected node in a temporary workflow. This type of node exists only in temporary workflows. The scheduling system sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- SKIP_CYCLE(5): The node is a node that is scheduled by the week or month and is waiting for the scheduling time to arrive. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- CONDITION_UNCHOOSE(6): The node is not selected by its ancestor branch node and is run as a dry-run node.
	//
	// 	- REALTIME_DEPRECATED(7): The node has instances that are generated in real time but deprecated. The scheduling system sets the status of the node to successful.
	//
	// example:
	//
	// NORMAL(0)
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetInstanceResponseBodyData) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *GetInstanceResponseBodyData) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *GetInstanceResponseBodyData) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *GetInstanceResponseBodyData) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetInstanceResponseBodyData) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetInstanceResponseBodyData) GetConnection() *string {
	return s.Connection
}

func (s *GetInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetInstanceResponseBodyData) GetCycTime() *int64 {
	return s.CycTime
}

func (s *GetInstanceResponseBodyData) GetDagId() *int64 {
	return s.DagId
}

func (s *GetInstanceResponseBodyData) GetDagType() *string {
	return s.DagType
}

func (s *GetInstanceResponseBodyData) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *GetInstanceResponseBodyData) GetDqcType() *int32 {
	return s.DqcType
}

func (s *GetInstanceResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetInstanceResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetInstanceResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetInstanceResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetInstanceResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetInstanceResponseBodyData) GetParamValues() *string {
	return s.ParamValues
}

func (s *GetInstanceResponseBodyData) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *GetInstanceResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetInstanceResponseBodyData) GetRelatedFlowId() *int64 {
	return s.RelatedFlowId
}

func (s *GetInstanceResponseBodyData) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *GetInstanceResponseBodyData) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *GetInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyData) GetTaskRerunTime() *int32 {
	return s.TaskRerunTime
}

func (s *GetInstanceResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetInstanceResponseBodyData) SetBaselineId(v int64) *GetInstanceResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBeginRunningTime(v int64) *GetInstanceResponseBodyData {
	s.BeginRunningTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBeginWaitResTime(v int64) *GetInstanceResponseBodyData {
	s.BeginWaitResTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBeginWaitTimeTime(v int64) *GetInstanceResponseBodyData {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBizdate(v int64) *GetInstanceResponseBodyData {
	s.Bizdate = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBusinessId(v int64) *GetInstanceResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetConnection(v string) *GetInstanceResponseBodyData {
	s.Connection = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreateTime(v int64) *GetInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreateUser(v string) *GetInstanceResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCycTime(v int64) *GetInstanceResponseBodyData {
	s.CycTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDagId(v int64) *GetInstanceResponseBodyData {
	s.DagId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDagType(v string) *GetInstanceResponseBodyData {
	s.DagType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDqcDescription(v string) *GetInstanceResponseBodyData {
	s.DqcDescription = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDqcType(v int32) *GetInstanceResponseBodyData {
	s.DqcType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetFinishTime(v int64) *GetInstanceResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v int64) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetModifyTime(v int64) *GetInstanceResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetNodeId(v int64) *GetInstanceResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetNodeName(v string) *GetInstanceResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetOwner(v string) *GetInstanceResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetParamValues(v string) *GetInstanceResponseBodyData {
	s.ParamValues = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetPeriodNumber(v int32) *GetInstanceResponseBodyData {
	s.PeriodNumber = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetPriority(v int32) *GetInstanceResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRelatedFlowId(v int64) *GetInstanceResponseBodyData {
	s.RelatedFlowId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRepeatInterval(v int64) *GetInstanceResponseBodyData {
	s.RepeatInterval = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRepeatability(v bool) *GetInstanceResponseBodyData {
	s.Repeatability = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTaskRerunTime(v int32) *GetInstanceResponseBodyData {
	s.TaskRerunTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTaskType(v string) *GetInstanceResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
