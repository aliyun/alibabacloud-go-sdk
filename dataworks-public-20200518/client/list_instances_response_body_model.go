// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() *ListInstancesResponseBodyData
	SetErrorCode(v string) *ListInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInstancesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
}

type ListInstancesResponseBody struct {
	// The list of instances.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ProjectNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The project does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetData() *ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// The instance information.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 66
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetInstances() []*ListInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *ListInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int32) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int32) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int32) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataInstances struct {
	// The baseline ID.
	//
	// example:
	//
	// 123123
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The time when the instance started running.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The time when the instance started waiting for resources.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The time when the instance started waiting for scheduling.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The data timestamp of the scheduled node. This is typically the day before the node runs.
	//
	// The value is a 13-digit number, such as `1590336000000`.
	//
	// example:
	//
	// 1590336000000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The business process ID.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The connection string.
	//
	// example:
	//
	// odps_source
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The time when the instance was created.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user who triggered the instance to run. For example, if user Test triggered a data backfill instance, the CreateUser is Test.
	//
	// example:
	//
	// Test
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The scheduled runtime of the node.
	//
	// The value is a 13-digit number, such as `1590422400000`.
	//
	// example:
	//
	// 1590422400000
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 33845
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The type of the workflow. Valid values:
	//
	// - DAILY(0): daily scheduling workflow.
	//
	// - MANUAL(1): manual task workflow.
	//
	// - SMOKE_TEST(2): smoke testing workflow.
	//
	// - SUPPLY_DATA(3): data backfill workflow.
	//
	// - MANUAL_FLOW(4): manually triggered dataflow PAI workflow (such as running a workflow in the IDE).
	//
	// - BUSINESS_PROCESS_DAG(5): manual business process workflow.
	//
	// example:
	//
	// DAILY
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The DQC partitioning rule string.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// The DQC type. Valid values:
	//
	// - 0: associated with DQC.
	//
	// - 1: not associated with DQC.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// **[Deprecated]*	- The error message of the instance run. You can call [GetInstanceLog](https://help.aliyun.com/document_detail/173983.html) to obtain the error information of the executed task.
	//
	// example:
	//
	// error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the scheduled node finished running.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the scheduled node was last modified.
	//
	// The value is a 13-digit number, such as `1590416703313`.
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
	// The node name.
	//
	// example:
	//
	// kzh
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The parameter information.
	//
	// example:
	//
	// bizdate=$bizdate tbods=$tbods
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The priority of the instance. Valid values: 1, 3, 5, 7, and 8.
	//
	// A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the associated business process.
	//
	// example:
	//
	// 123456
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The interval at which the node is rescheduled after a failure. Unit: milliseconds.
	//
	// example:
	//
	// 60000
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Indicates whether the instance task can be rerun.
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The status of the node. Valid values:
	//
	// - NOT_RUN(1): The node is not run.
	//
	// - WAIT_TIME(2): The node is waiting for the scheduled time to arrive.
	//
	// - WAIT_RESOURCE(3): The node has been sent to the execution engine and is waiting for resources to be scheduled.
	//
	// - RUNNING(4): The node is running.
	//
	// - CHECKING(7): The node has finished running and has been sent to Data Quality for data verification.
	//
	// - CHECKING_CONDITION(8): The node has finished running and is undergoing branch condition verification.
	//
	// - WAIT_TRIGGER(9): The node is waiting to be triggered. A trigger-based node enters this state after the waiting time elapses.
	//
	// - FAILURE(5): The node failed to run.
	//
	// - SUCCESS(6): The node ran successfully.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of remaining reruns for the instance. The value can be empty or an integer greater than or equal to 0.
	//
	// - Empty: The node corresponding to this instance does not have automatic rerun configured.
	//
	// - 0: The instance cannot be rerun.
	//
	// - An integer greater than 0 (n): The instance can be rerun n times. For example, if the value is 1, the remaining rerun count is 1. If the value is 2, the remaining rerun count is 2, and so on. The initial value is the automatic rerun count defined for the corresponding node plus 1.
	//
	// example:
	//
	// 0
	TaskRerunTime *int32 `json:"TaskRerunTime,omitempty" xml:"TaskRerunTime,omitempty"`
	// The scheduling type of the task instance. Valid values:
	//
	// - NORMAL(0): The node is a normal scheduled node that is triggered by daily scheduling.
	//
	// - MANUAL(1): The node is a manual node that is not triggered by daily scheduling.
	//
	// - PAUSE(2): The node is a frozen node that is triggered by daily scheduling but is set to failed when scheduling starts.
	//
	// - SKIP(3): The node is a dry-run node that is triggered by daily scheduling but is set to successful when scheduling starts.
	//
	// - SKIP_UNCHOOSE(4): The node is an unselected node in a temporary workflow. It exists only in temporary workflows and is set to successful when scheduling starts.
	//
	// - SKIP_CYCLE(5): The node is a weekly or monthly node whose scheduling cycle has not arrived. It is triggered by daily scheduling but is set to successful when scheduling starts.
	//
	// - CONDITION_UNCHOOSE(6): The upstream instance contains a branch (IF) node, but this downstream node is not selected by the branch node and is set to a dry-run node.
	//
	// - REALTIME_DEPRECATED(7): The node is an expired periodic instance generated in real time. This type of node is set to successful.
	//
	// example:
	//
	// NORMAL(0)
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListInstancesResponseBodyDataInstances) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *ListInstancesResponseBodyDataInstances) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *ListInstancesResponseBodyDataInstances) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *ListInstancesResponseBodyDataInstances) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListInstancesResponseBodyDataInstances) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListInstancesResponseBodyDataInstances) GetConnection() *string {
	return s.Connection
}

func (s *ListInstancesResponseBodyDataInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyDataInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListInstancesResponseBodyDataInstances) GetCycTime() *int64 {
	return s.CycTime
}

func (s *ListInstancesResponseBodyDataInstances) GetDagId() *int64 {
	return s.DagId
}

func (s *ListInstancesResponseBodyDataInstances) GetDagType() *string {
	return s.DagType
}

func (s *ListInstancesResponseBodyDataInstances) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *ListInstancesResponseBodyDataInstances) GetDqcType() *int32 {
	return s.DqcType
}

func (s *ListInstancesResponseBodyDataInstances) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstancesResponseBodyDataInstances) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListInstancesResponseBodyDataInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyDataInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListInstancesResponseBodyDataInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListInstancesResponseBodyDataInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInstancesResponseBodyDataInstances) GetParamValues() *string {
	return s.ParamValues
}

func (s *ListInstancesResponseBodyDataInstances) GetPriority() *int32 {
	return s.Priority
}

func (s *ListInstancesResponseBodyDataInstances) GetRelatedFlowId() *int64 {
	return s.RelatedFlowId
}

func (s *ListInstancesResponseBodyDataInstances) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *ListInstancesResponseBodyDataInstances) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *ListInstancesResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyDataInstances) GetTaskRerunTime() *int32 {
	return s.TaskRerunTime
}

func (s *ListInstancesResponseBodyDataInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListInstancesResponseBodyDataInstances) SetBaselineId(v int64) *ListInstancesResponseBodyDataInstances {
	s.BaselineId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetBeginRunningTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.BeginRunningTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetBeginWaitResTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.BeginWaitResTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetBeginWaitTimeTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetBizdate(v int64) *ListInstancesResponseBodyDataInstances {
	s.Bizdate = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetBusinessId(v int64) *ListInstancesResponseBodyDataInstances {
	s.BusinessId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetConnection(v string) *ListInstancesResponseBodyDataInstances {
	s.Connection = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetCreateTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetCreateUser(v string) *ListInstancesResponseBodyDataInstances {
	s.CreateUser = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetCycTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.CycTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetDagId(v int64) *ListInstancesResponseBodyDataInstances {
	s.DagId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetDagType(v string) *ListInstancesResponseBodyDataInstances {
	s.DagType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetDqcDescription(v string) *ListInstancesResponseBodyDataInstances {
	s.DqcDescription = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetDqcType(v int32) *ListInstancesResponseBodyDataInstances {
	s.DqcType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetErrorMessage(v string) *ListInstancesResponseBodyDataInstances {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetFinishTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.FinishTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v int64) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetModifyTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetNodeId(v int64) *ListInstancesResponseBodyDataInstances {
	s.NodeId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetNodeName(v string) *ListInstancesResponseBodyDataInstances {
	s.NodeName = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetParamValues(v string) *ListInstancesResponseBodyDataInstances {
	s.ParamValues = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPriority(v int32) *ListInstancesResponseBodyDataInstances {
	s.Priority = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetRelatedFlowId(v int64) *ListInstancesResponseBodyDataInstances {
	s.RelatedFlowId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetRepeatInterval(v int64) *ListInstancesResponseBodyDataInstances {
	s.RepeatInterval = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetRepeatability(v bool) *ListInstancesResponseBodyDataInstances {
	s.Repeatability = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStatus(v string) *ListInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTaskRerunTime(v int32) *ListInstancesResponseBodyDataInstances {
	s.TaskRerunTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTaskType(v string) *ListInstancesResponseBodyDataInstances {
	s.TaskType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}
