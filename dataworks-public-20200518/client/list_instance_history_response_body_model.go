// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstanceHistoryResponseBodyInstances) *ListInstanceHistoryResponseBody
	GetInstances() []*ListInstanceHistoryResponseBodyInstances
	SetRequestId(v string) *ListInstanceHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceHistoryResponseBody
	GetSuccess() *bool
}

type ListInstanceHistoryResponseBody struct {
	// The list of instances.
	Instances []*ListInstanceHistoryResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryResponseBody) GetInstances() []*ListInstanceHistoryResponseBodyInstances {
	return s.Instances
}

func (s *ListInstanceHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceHistoryResponseBody) SetInstances(v []*ListInstanceHistoryResponseBodyInstances) *ListInstanceHistoryResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstanceHistoryResponseBody) SetRequestId(v string) *ListInstanceHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceHistoryResponseBody) SetSuccess(v bool) *ListInstanceHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceHistoryResponseBody) Validate() error {
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

type ListInstanceHistoryResponseBodyInstances struct {
	// The time when the instance started running, in timestamp format.
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
	// The business date on which the scheduled node was run. This value is typically one day before the run time of the node.
	//
	// The value is a 13-digit number, such as `1590336000000`.
	//
	// example:
	//
	// 1590336000000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The time when the instance was created.
	//
	// The value is a 13-digit number, such as `1590416703313`.
	//
	// example:
	//
	// 1590416703313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The scheduled run time of the node, in timestamp format.
	//
	// example:
	//
	// 1590422400000
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 33845
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The Data Quality Check (DQC) type. Valid values:
	//
	// - 0: associated with DQC.
	//
	// - 1: not associated with DQC.
	//
	// example:
	//
	// 1
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// **[Deprecated]*	- The error message returned when the instance failed to run. This field is deprecated. You can call the GetInstanceLog operation to obtain the error information of the node.
	//
	// example:
	//
	// error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the scheduled node finished running, in timestamp format.
	//
	// example:
	//
	// 1590416703313
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The history archive ID of the instance.
	//
	// example:
	//
	// 1
	InstanceHistoryId *int64 `json:"InstanceHistoryId,omitempty" xml:"InstanceHistoryId,omitempty"`
	// The ID of the instance.
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
	// The ID of the node.
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
	// The status of the node. Valid values:
	//
	// - NOT_RUN: The node is not run.
	//
	// - WAIT_TIME: The node is waiting for the scheduled time (DueTime or CycTime) to arrive.
	//
	// - WAIT_RESOURCE: The node is waiting for resources.
	//
	// - RUNNING: The node is running.
	//
	// - CHECKING: The node is sent to Data Quality for data verification.
	//
	// - CHECKING_CONDITION: The node is undergoing branch condition verification.
	//
	// - FAILURE: The node failed to run.
	//
	// - SUCCESS: The node ran successfully.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduling type of the node instance. Valid values:
	//
	// - NORMAL(0): A normal scheduling node. The node is scheduled on a daily basis.
	//
	// - MANUAL(1): A manual node. The node is not scheduled on a daily basis.
	//
	// - PAUSE(2): A frozen node. The node is scheduled on a daily basis, but is set to failed when scheduling starts.
	//
	// - SKIP(3): A dry-run node. The node is scheduled on a daily basis, but is set to successful when scheduling starts.
	//
	// - SKIP_UNCHOOSE(4): A node that is not selected in a temporary workflow. This type of node exists only in temporary workflows and is set to successful when scheduling starts.
	//
	// - SKIP_CYCLE(5): A weekly or monthly node that has not reached its run cycle. The node is scheduled on a daily basis, but is set to successful when scheduling starts.
	//
	// - CONDITION_UNCHOOSE(6): A downstream node that is not selected by an upstream branch (IF) node. The node is directly set to dry-run.
	//
	// - REALTIME_DEPRECATED(7): An expired periodic instance generated in real time. This type of node is directly set to successful.
	//
	// example:
	//
	// NORMAL(0)
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListInstanceHistoryResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryResponseBodyInstances) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListInstanceHistoryResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetCycTime() *int64 {
	return s.CycTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetDagId() *int64 {
	return s.DagId
}

func (s *ListInstanceHistoryResponseBodyInstances) GetDagType() *string {
	return s.DagType
}

func (s *ListInstanceHistoryResponseBodyInstances) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstanceHistoryResponseBodyInstances) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetInstanceHistoryId() *int64 {
	return s.InstanceHistoryId
}

func (s *ListInstanceHistoryResponseBodyInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListInstanceHistoryResponseBodyInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListInstanceHistoryResponseBodyInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListInstanceHistoryResponseBodyInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInstanceHistoryResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceHistoryResponseBodyInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListInstanceHistoryResponseBodyInstances) SetBeginRunningTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.BeginRunningTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetBeginWaitResTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.BeginWaitResTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetBeginWaitTimeTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetBizdate(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.Bizdate = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetCreateTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetCycTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.CycTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetDagId(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.DagId = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetDagType(v string) *ListInstanceHistoryResponseBodyInstances {
	s.DagType = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetErrorMessage(v string) *ListInstanceHistoryResponseBodyInstances {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetFinishTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.FinishTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetInstanceHistoryId(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.InstanceHistoryId = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetInstanceId(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetModifyTime(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetNodeId(v int64) *ListInstanceHistoryResponseBodyInstances {
	s.NodeId = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetNodeName(v string) *ListInstanceHistoryResponseBodyInstances {
	s.NodeName = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetStatus(v string) *ListInstanceHistoryResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) SetTaskType(v string) *ListInstanceHistoryResponseBodyInstances {
	s.TaskType = &v
	return s
}

func (s *ListInstanceHistoryResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
