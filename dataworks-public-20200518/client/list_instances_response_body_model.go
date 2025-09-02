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
	// The ID of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID of the node.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// Invalid.Tenant.ProjectNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// The project does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The error message that is returned for the instance.
	//
	// This parameter is deprecated. You can call the [GetInstanceLog](https://help.aliyun.com/document_detail/173983.html) operation to query the error information related to the node.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
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
	return dara.Validate(s)
}

type ListInstancesResponseBodyData struct {
	// The name of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the name of the node.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The time when the node was scheduled to run.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the instance. Valid values: 1, 3, 5, 7, and 8.
	//
	// A greater value indicates a higher priority. Default value: 1.
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
	return dara.Validate(s)
}

type ListInstancesResponseBodyDataInstances struct {
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
	// 123123
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1590416703313
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The time when the node stopped running.
	//
	// example:
	//
	// 1590416703313
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 1590416703313
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// You cannot specify the sorting method for the instances to be returned by this operation. By default, the instances are sorted in descending order of the time when the instances were created.
	//
	// example:
	//
	// 1590336000000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The ID of the workflow to which the node belongs.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The number of times the node can be rerun. The value of this parameter can be empty or an integer that is greater than or equal to 0.
	//
	// 	- If the value of this parameter is empty, the number of times that the node can be rerun is not specified.
	//
	// 	- If the value of this parameter is 0, the node cannot be rerun.
	//
	// 	- If the value of this parameter is a positive integer such as n, the node can be rerun n times. For example, if the value of this parameter is 1, the node can be rerun once. If the value of this parameter is 2, the node can be rerun twice.
	//
	// example:
	//
	// odps_first
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The interval at which the node is rerun after the node fails to run. Unit: milliseconds.
	//
	// example:
	//
	// 1590416703313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID of the node.
	//
	// example:
	//
	// Test
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// 1590422400000
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The time when the instance started to wait for resources.
	//
	// example:
	//
	// 33845
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The data timestamp of the instance. In most cases, the value is one day before the time when the instance was run.
	//
	// example:
	//
	// DAILY
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The operation that you want to perform.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- NOT_RUN: The node is not run.
	//
	// 	- WAIT_TIME: The node is waiting for the scheduling time to arrive.
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
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The name of the account that is used to run the instance. For example, if an account named Test was used to run the instance to backfill data, the value of this parameter is Test.
	//
	// example:
	//
	// error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the Alibaba Cloud account used by the workspace administrator. You can log on to the Alibaba Cloud Management Console and view the ID on the Security Settings page of the Account Center console.
	//
	// example:
	//
	// 1590416703313
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The number of the page to return. Minimum value:1. Maximum value: 100.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the workflow. You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the name of the workflow.
	//
	// example:
	//
	// 1590416703313
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV. The value PROD indicates the production environment. The value DEV indicates the development environment.
	//
	// example:
	//
	// 33115
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// kzh
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The table and partition filter expression in Data Quality that are associated with the node.
	//
	// example:
	//
	// bizdate=$bizdate tbods=$tbods
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the type of the node.
	//
	// example:
	//
	// 123456
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL(0): The node is an auto triggered node. The scheduling system regularly runs the node.
	//
	// 	- MANUAL(1): The node is a manually triggered node. The scheduling system does not regularly run the node.
	//
	// 	- PAUSE(2): The node is a frozen node. The scheduling system regularly runs the node but sets the status of the node to failed when the scheduling system starts to run the node.
	//
	// 	- SKIP(3): The node is a dry-run node. The scheduling system regularly runs the node but sets the status of the node to succeeded when the scheduling system starts to run the node.
	//
	// 	- SKIP_UNCHOOSE(4): The node is an unselected node in a temporary workflow. This type of node exists only in temporary workflows. The scheduling system sets the status of the node to succeeded when the scheduling system starts to run the node.
	//
	// 	- SKIP_CYCLE(5): The node is a node that is scheduled by week or month and is waiting for the scheduling time to arrive. The scheduling system regularly runs the node but sets the status of the node to succeeded when the scheduling system starts to run the node.
	//
	// 	- CONDITION_UNCHOOSE(6): The node is not selected by its ancestor branch node and is run as a dry-run node.
	//
	//     REALTIME_DEPRECATED(7): The node has instances that are generated in real time but deprecated. The scheduling system sets the status of the node to succeeded.
	//
	// example:
	//
	// 60000
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- NOT_RUN: The node is not run.
	//
	// 	- WAIT_TIME: The node is waiting for the scheduling time to arrive.
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
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The data timestamp of the instances that you want to query. Specify the timestamp in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workspace. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID of the workspace.
	//
	// example:
	//
	// 0
	TaskRerunTime *int32 `json:"TaskRerunTime,omitempty" xml:"TaskRerunTime,omitempty"`
	// The information about the instances.
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
