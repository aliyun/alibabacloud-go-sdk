// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManualDagInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*GetManualDagInstancesResponseBodyInstances) *GetManualDagInstancesResponseBody
	GetInstances() []*GetManualDagInstancesResponseBodyInstances
	SetRequestId(v string) *GetManualDagInstancesResponseBody
	GetRequestId() *string
}

type GetManualDagInstancesResponseBody struct {
	// The instances in the manually triggered workflow.
	Instances []*GetManualDagInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetManualDagInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManualDagInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetManualDagInstancesResponseBody) GetInstances() []*GetManualDagInstancesResponseBodyInstances {
	return s.Instances
}

func (s *GetManualDagInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManualDagInstancesResponseBody) SetInstances(v []*GetManualDagInstancesResponseBodyInstances) *GetManualDagInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *GetManualDagInstancesResponseBody) SetRequestId(v string) *GetManualDagInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManualDagInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetManualDagInstancesResponseBodyInstances struct {
	// The time when the instance started to run.
	//
	// example:
	//
	// 1605178414676
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The time when the instance started to wait for resources.
	//
	// example:
	//
	// 1605178414676
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The time when the instance started to wait to be scheduled.
	//
	// example:
	//
	// 1605178414676
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The data timestamp of the instance. In most cases, the value is one day before the time when the instance was run.
	//
	// example:
	//
	// 1605178414676
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The time when the instance was generated.
	//
	// example:
	//
	// 1605178414676
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user who performed the operation.
	//
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The time when the instance was scheduled to run.
	//
	// example:
	//
	// 1605178414676
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The ID of the DAG for the manually triggered workflow.
	//
	// example:
	//
	// 350850491
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The type of the manually triggered workflow.
	//
	// example:
	//
	// 5
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The time when the instance finished running.
	//
	// example:
	//
	// 1605178414676
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the instance in the manually triggered workflow.
	//
	// example:
	//
	// 11726873619
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1605178414676
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the node in the manually triggered workflow.
	//
	// example:
	//
	// 37851
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test2
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The parameters related to the instance.
	//
	// example:
	//
	// xxx
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- NOT_RUN: The instance is not run.
	//
	// 	- WAIT_TIME: The instance is waiting for its scheduling time to arrive.
	//
	// 	- WAIT_RESOURCE: The instance is waiting for resources.
	//
	// 	- RUNNING: The instance is running.
	//
	// 	- CHECKING: Data quality is being checked for the instance.
	//
	// 	- CHECKING_CONDITION: Branch conditions are being checked for the instance.
	//
	// 	- FAILURE: The instance fails to be run.
	//
	// 	- SUCCESS: The instance is successfully run.
	//
	// example:
	//
	// WAIT_TIME
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduling type of the node that generates the instance. Valid values:
	//
	// 	- NORMAL(0): The node is an auto triggered node. The scheduling system regularly runs the node.
	//
	// 	- MANUAL(1): The node is a manually triggered node. The scheduling system does not regularly run the node.
	//
	// 	- PAUSE(2): The node is a paused node. The scheduling system regularly runs the node but sets the status of the node to failed when the scheduling system starts to run the node.
	//
	// 	- SKIP(3): The node is a dry-run node. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- SKIP_UNCHOOSE(4): The node is an unselected node in a temporary workflow. This type of node exists only in temporary workflows. The scheduling system sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- SKIP_CYCLE(5): The node is a node that is scheduled by week or month and is waiting for the scheduling time to arrive. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
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

func (s GetManualDagInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s GetManualDagInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetManualDagInstancesResponseBodyInstances) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetBizDate() *int64 {
	return s.BizDate
}

func (s *GetManualDagInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetManualDagInstancesResponseBodyInstances) GetCycTime() *int64 {
	return s.CycTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetDagId() *int64 {
	return s.DagId
}

func (s *GetManualDagInstancesResponseBodyInstances) GetDagType() *string {
	return s.DagType
}

func (s *GetManualDagInstancesResponseBodyInstances) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetManualDagInstancesResponseBodyInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetManualDagInstancesResponseBodyInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetManualDagInstancesResponseBodyInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *GetManualDagInstancesResponseBodyInstances) GetParamValues() *string {
	return s.ParamValues
}

func (s *GetManualDagInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *GetManualDagInstancesResponseBodyInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *GetManualDagInstancesResponseBodyInstances) SetBeginRunningTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.BeginRunningTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetBeginWaitResTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.BeginWaitResTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetBeginWaitTimeTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetBizDate(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.BizDate = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetCreateTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetCreateUser(v string) *GetManualDagInstancesResponseBodyInstances {
	s.CreateUser = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetCycTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.CycTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetDagId(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.DagId = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetDagType(v string) *GetManualDagInstancesResponseBodyInstances {
	s.DagType = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetFinishTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.FinishTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetInstanceId(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetModifyTime(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetNodeId(v int64) *GetManualDagInstancesResponseBodyInstances {
	s.NodeId = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetNodeName(v string) *GetManualDagInstancesResponseBodyInstances {
	s.NodeName = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetParamValues(v string) *GetManualDagInstancesResponseBodyInstances {
	s.ParamValues = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetStatus(v string) *GetManualDagInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) SetTaskType(v string) *GetManualDagInstancesResponseBodyInstances {
	s.TaskType = &v
	return s
}

func (s *GetManualDagInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
