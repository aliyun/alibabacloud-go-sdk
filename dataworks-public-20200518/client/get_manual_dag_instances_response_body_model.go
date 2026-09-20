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
	// The list of internal instances of the manual workflow.
	Instances []*GetManualDagInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The unique ID of the request.
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

type GetManualDagInstancesResponseBodyInstances struct {
	// The time when the instance node started to run.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The time when the instance node started to wait for resources.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The time when the instance node started to wait for scheduling.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The business date. This is typically the day before the node runs.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The time when the instance node was created.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The operator.
	//
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The scheduled time of the instance node.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	CycTime *int64 `json:"CycTime,omitempty" xml:"CycTime,omitempty"`
	// The DAG ID of the manual workflow instance.
	//
	// example:
	//
	// 350850491
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The type of the manual workflow.
	//
	// example:
	//
	// 5
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The time when the instance node finished running.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The internal instance ID.
	//
	// example:
	//
	// 11726873619
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The most recent modification time of the instance node.
	//
	// The value is a 13-digit number, for example, `1605178414676`.
	//
	// example:
	//
	// 1605178414676
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The internal node ID of the workflow.
	//
	// example:
	//
	// 37851
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// example:
	//
	// test2
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The parameter information of the instance.
	//
	// example:
	//
	// xxx
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The status of the instance node. Valid values:
	//
	// - NOT_RUN: The instance is not run.
	//
	// - WAIT_TIME: The instance is waiting for the scheduled dueTime or cycleTime.
	//
	// - WAIT_RESOURCE: The instance is waiting for resources.
	//
	// - RUNNING: The instance is running.
	//
	// - CHECKING: The instance is submitted to Data Quality for data verification.
	//
	// - CHECKING_CONDITION: The instance is performing branch condition verification.
	//
	// - FAILURE: The instance failed to run.
	//
	// - SUCCESS: The instance is run successfully.
	//
	// example:
	//
	// WAIT_TIME
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduling type of the instance node. Valid values:
	//
	// - NORMAL(0): a normal scheduling node. The node is scheduled on a daily basis.
	//
	// - MANUAL(1): a manual node. The node is not scheduled on a daily basis.
	//
	// - PAUSE(2): a paused node. The node is scheduled on a daily basis, but is set to failed when scheduling starts.
	//
	// - SKIP(3): a dry-run node. The node is scheduled on a daily basis, but is set to successful when scheduling starts.
	//
	// - SKIP_UNCHOOSE(4): a node that is not selected in a temporary workflow. This type of node exists only in temporary workflows and is set to successful when scheduling starts.
	//
	// - SKIP_CYCLE(5): a weekly or monthly node that has not reached its run cycle. The node is scheduled on a daily basis, but is set to successful when scheduling starts.
	//
	// - CONDITION_UNCHOOSE(6): a downstream node that is not selected by an upstream branch (IF) node. The node is directly set to dry-run.
	//
	// - REALTIME_DEPRECATED(7): an expired periodic instance generated in real time. This type of node is directly set to successful.
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
