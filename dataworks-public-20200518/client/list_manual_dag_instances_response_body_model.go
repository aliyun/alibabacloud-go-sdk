// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualDagInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListManualDagInstancesResponseBodyInstances) *ListManualDagInstancesResponseBody
	GetInstances() []*ListManualDagInstancesResponseBodyInstances
	SetRequestId(v string) *ListManualDagInstancesResponseBody
	GetRequestId() *string
}

type ListManualDagInstancesResponseBody struct {
	// The instances in the manually triggered workflow.
	Instances []*ListManualDagInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListManualDagInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManualDagInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListManualDagInstancesResponseBody) GetInstances() []*ListManualDagInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListManualDagInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManualDagInstancesResponseBody) SetInstances(v []*ListManualDagInstancesResponseBodyInstances) *ListManualDagInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListManualDagInstancesResponseBody) SetRequestId(v string) *ListManualDagInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManualDagInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListManualDagInstancesResponseBodyInstances struct {
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
	// The name of the account that is used to run the instance. For example, if you use an account named Test to run the instance, the value of this parameter is Test.
	//
	// example:
	//
	// Test
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The time when the instance was scheduled to run.
	//
	// example:
	//
	// 1605178414676
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The ID of the DAG for the instance in the manually triggered workflow.
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
	// The instance ID.
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
	// xxx=yyy
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
	// SUCCESS
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
	// 	- SKIP_CYCLE(5): The node is a node that is scheduled by week or month and is waiting for the scheduling time to arrive. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// 	- CONDITION_UNCHOOSE(6): The node is not selected by its ancestor branch node and is run as a dry-run node.
	//
	// 	- REALTIME_DEPRECATED(7): The node has instances that are generated in real time but deprecated. The scheduling system sets the status of the node to successful.
	//
	// example:
	//
	// MANUAL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListManualDagInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListManualDagInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListManualDagInstancesResponseBodyInstances) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListManualDagInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListManualDagInstancesResponseBodyInstances) GetCycTime() *int64 {
	return s.CycTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetDagId() *int64 {
	return s.DagId
}

func (s *ListManualDagInstancesResponseBodyInstances) GetDagType() *string {
	return s.DagType
}

func (s *ListManualDagInstancesResponseBodyInstances) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListManualDagInstancesResponseBodyInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListManualDagInstancesResponseBodyInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListManualDagInstancesResponseBodyInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListManualDagInstancesResponseBodyInstances) GetParamValues() *string {
	return s.ParamValues
}

func (s *ListManualDagInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListManualDagInstancesResponseBodyInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListManualDagInstancesResponseBodyInstances) SetBeginRunningTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.BeginRunningTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetBeginWaitResTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.BeginWaitResTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetBeginWaitTimeTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetBizDate(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.BizDate = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetCreateTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetCreateUser(v string) *ListManualDagInstancesResponseBodyInstances {
	s.CreateUser = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetCycTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.CycTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetDagId(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.DagId = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetDagType(v string) *ListManualDagInstancesResponseBodyInstances {
	s.DagType = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetFinishTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.FinishTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetInstanceId(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetModifyTime(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetNodeId(v int64) *ListManualDagInstancesResponseBodyInstances {
	s.NodeId = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetNodeName(v string) *ListManualDagInstancesResponseBodyInstances {
	s.NodeName = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetParamValues(v string) *ListManualDagInstancesResponseBodyInstances {
	s.ParamValues = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetStatus(v string) *ListManualDagInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) SetTaskType(v string) *ListManualDagInstancesResponseBodyInstances {
	s.TaskType = &v
	return s
}

func (s *ListManualDagInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
