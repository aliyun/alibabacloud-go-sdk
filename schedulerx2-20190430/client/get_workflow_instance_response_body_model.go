// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkflowInstanceResponseBody
	GetCode() *int32
	SetData(v *GetWorkflowInstanceResponseBodyData) *GetWorkflowInstanceResponseBody
	GetData() *GetWorkflowInstanceResponseBodyData
	SetMessage(v string) *GetWorkflowInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkflowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkflowInstanceResponseBody
	GetSuccess() *bool
}

type GetWorkflowInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the workflow instance.
	Data *GetWorkflowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// workflowId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkflowInstanceResponseBody) GetData() *GetWorkflowInstanceResponseBodyData {
	return s.Data
}

func (s *GetWorkflowInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowInstanceResponseBody) SetCode(v int32) *GetWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetData(v *GetWorkflowInstanceResponseBodyData) *GetWorkflowInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetMessage(v string) *GetWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRequestId(v string) *GetWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetSuccess(v bool) *GetWorkflowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyData struct {
	// The directed acyclic graph (DAG) of the workflow instance, including nodes and dependencies.
	WfInstanceDag *GetWorkflowInstanceResponseBodyDataWfInstanceDag `json:"WfInstanceDag,omitempty" xml:"WfInstanceDag,omitempty" type:"Struct"`
	// The details of the workflow instance, including the state of the workflow instance, the time when the workflow instance started to run, the time when the workflow instance stopped running, the state of each job instance, and the dependencies between job instances.
	WfInstanceInfo *GetWorkflowInstanceResponseBodyDataWfInstanceInfo `json:"WfInstanceInfo,omitempty" xml:"WfInstanceInfo,omitempty" type:"Struct"`
}

func (s GetWorkflowInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyData) GetWfInstanceDag() *GetWorkflowInstanceResponseBodyDataWfInstanceDag {
	return s.WfInstanceDag
}

func (s *GetWorkflowInstanceResponseBodyData) GetWfInstanceInfo() *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	return s.WfInstanceInfo
}

func (s *GetWorkflowInstanceResponseBodyData) SetWfInstanceDag(v *GetWorkflowInstanceResponseBodyDataWfInstanceDag) *GetWorkflowInstanceResponseBodyData {
	s.WfInstanceDag = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyData) SetWfInstanceInfo(v *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) *GetWorkflowInstanceResponseBodyData {
	s.WfInstanceInfo = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDag struct {
	// The dependencies between job instances.
	Edges []*GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The job instances.
	Nodes []*GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDag) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDag) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) GetEdges() []*GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges {
	return s.Edges
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) GetNodes() []*GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	return s.Nodes
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) SetEdges(v []*GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) *GetWorkflowInstanceResponseBodyDataWfInstanceDag {
	s.Edges = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) SetNodes(v []*GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) *GetWorkflowInstanceResponseBodyDataWfInstanceDag {
	s.Nodes = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges struct {
	// The upstream job instance of the current job instance. A value of 0 indicates that the upstream job instance is the root node.
	//
	// example:
	//
	// 24058798
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The downstream job instance of the current job instance.
	//
	// example:
	//
	// 24058796
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) GetSource() *int64 {
	return s.Source
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) GetTarget() *int64 {
	return s.Target
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) SetSource(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges {
	s.Source = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) SetTarget(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges {
	s.Target = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes struct {
	// The number of retries when the job failed.
	//
	// example:
	//
	// 0
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// The data timestamp of the job.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the job stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1472
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the job instance.
	//
	// example:
	//
	// 24058796
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The job name.
	//
	// example:
	//
	// TestJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The execution result of the job.
	//
	// example:
	//
	// code=200
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the job was scheduled.
	//
	// example:
	//
	// 2023-01-03 18:00:03
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the job started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:03
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the job instance. Valid values: 1: The job instance is waiting for execution. 3: The job instance is running. 4: The job instance is run. 5: The job instance failed to run. 9: The job instance is rejected. Enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus.
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The worker on which the job instance run.
	//
	// example:
	//
	// 10.163.0.101:34027
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetAttempt() *int32 {
	return s.Attempt
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetDataTime() *string {
	return s.DataTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetEndTime() *string {
	return s.EndTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetJobId() *int64 {
	return s.JobId
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetJobInstanceId() *int64 {
	return s.JobInstanceId
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetJobName() *string {
	return s.JobName
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetResult() *string {
	return s.Result
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetStartTime() *string {
	return s.StartTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GetWorkAddr() *string {
	return s.WorkAddr
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetAttempt(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Attempt = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetDataTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.DataTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetEndTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.EndTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobId(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobInstanceId(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobInstanceId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobName(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobName = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetResult(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Result = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetScheduleTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.ScheduleTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetStartTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.StartTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetStatus(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetWorkAddr(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.WorkAddr = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyDataWfInstanceInfo struct {
	// The data timestamp of the workflow instance.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the workflow instance stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the workflow instance was scheduled to run.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the workflow instance started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the workflow instance. Valid values:
	//
	// 	- 1: pending
	//
	// 	- 2: preparing
	//
	// 	- 3: running
	//
	// 	- 4: successful
	//
	// 	- 5: failed
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GetDataTime() *string {
	return s.DataTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetDataTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.DataTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetEndTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.EndTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetScheduleTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.ScheduleTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetStartTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.StartTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetStatus(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) Validate() error {
	return dara.Validate(s)
}
