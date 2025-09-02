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
	// The instances.
	Instances []*ListInstanceHistoryResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
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
	return dara.Validate(s)
}

type ListInstanceHistoryResponseBodyInstances struct {
	// The time when the instance started to be run. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
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
	// The time when the instance was generated.
	//
	// example:
	//
	// 1590416703313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the node started to be run. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
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
	// Indicates whether the instance is associated with a monitoring rule in Data Quality. Valid values:
	//
	// 	- 0: The instance is associated with a monitoring rule in Data Quality.
	//
	// 	- 1: The instance is not associated with a monitoring rule in Data Quality.
	//
	// example:
	//
	// 1
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The error message. This parameter is deprecated. You can call the GetInstanceLog operation to query the error information related to the node.
	//
	// example:
	//
	// error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the running of the node was complete. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1590416703313
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The historical record number of the instance.
	//
	// example:
	//
	// 1
	InstanceHistoryId *int64 `json:"InstanceHistoryId,omitempty" xml:"InstanceHistoryId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the node was last modified.
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
	// The status of the node that generates the instance. Valid values:
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
	// 	- FAILURE: The node fails to be run.
	//
	// 	- SUCCESS: The node is successfully run.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduling type of the node. Valid values:
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
