// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListQueuesResponseBody
	GetClusterId() *string
	SetPageNumber(v int32) *ListQueuesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQueuesResponseBody
	GetPageSize() *int32
	SetQueues(v []*ListQueuesResponseBodyQueues) *ListQueuesResponseBody
	GetQueues() []*ListQueuesResponseBodyQueues
	SetRequestId(v string) *ListQueuesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQueuesResponseBody
	GetTotalCount() *int32
}

type ListQueuesResponseBody struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the queues.
	Queues []*ListQueuesResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C6E5005C-00B0-4F27-98BB-95AB88016C22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListQueuesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQueuesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQueuesResponseBody) GetQueues() []*ListQueuesResponseBodyQueues {
	return s.Queues
}

func (s *ListQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueuesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQueuesResponseBody) SetClusterId(v string) *ListQueuesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListQueuesResponseBody) SetPageNumber(v int32) *ListQueuesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListQueuesResponseBody) SetPageSize(v int32) *ListQueuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListQueuesResponseBody) SetQueues(v []*ListQueuesResponseBodyQueues) *ListQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueuesResponseBody) SetTotalCount(v int32) *ListQueuesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQueuesResponseBody) Validate() error {
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueuesResponseBodyQueues struct {
	// The hardware configurations of the compute nodes that are added in auto scale-outs. Up to five nodes are displayed.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// The time when the queue was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2023-11-10T02:04:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether auto scale-in is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether auto scale-out is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 100
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes that are added to the queue in each auto scale-out.
	//
	// example:
	//
	// 1
	MaxCountPerCycle *int32 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes that the queue must contain.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The statistics about the compute nodes in the queue.
	Nodes *ListQueuesResponseBodyQueuesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The queue name.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The total number of vCPUs that are used by all compute nodes in the queue.
	//
	// example:
	//
	// 24
	TotalCores *int32 `json:"TotalCores,omitempty" xml:"TotalCores,omitempty"`
	// The time when the queue was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-04-25T02:02:32
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The vSwitches that can be used for added nodes during auto scale-outs. Up to three vSwitches are displayed.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueues) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueues) GetComputeNodes() []*NodeTemplate {
	return s.ComputeNodes
}

func (s *ListQueuesResponseBodyQueues) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQueuesResponseBodyQueues) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *ListQueuesResponseBodyQueues) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *ListQueuesResponseBodyQueues) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *ListQueuesResponseBodyQueues) GetMaxCountPerCycle() *int32 {
	return s.MaxCountPerCycle
}

func (s *ListQueuesResponseBodyQueues) GetMinCount() *int32 {
	return s.MinCount
}

func (s *ListQueuesResponseBodyQueues) GetNodes() *ListQueuesResponseBodyQueuesNodes {
	return s.Nodes
}

func (s *ListQueuesResponseBodyQueues) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueuesResponseBodyQueues) GetTotalCores() *int32 {
	return s.TotalCores
}

func (s *ListQueuesResponseBodyQueues) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListQueuesResponseBodyQueues) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListQueuesResponseBodyQueues) SetComputeNodes(v []*NodeTemplate) *ListQueuesResponseBodyQueues {
	s.ComputeNodes = v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetCreateTime(v string) *ListQueuesResponseBodyQueues {
	s.CreateTime = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetEnableScaleIn(v bool) *ListQueuesResponseBodyQueues {
	s.EnableScaleIn = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetEnableScaleOut(v bool) *ListQueuesResponseBodyQueues {
	s.EnableScaleOut = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMaxCount(v int32) *ListQueuesResponseBodyQueues {
	s.MaxCount = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMaxCountPerCycle(v int32) *ListQueuesResponseBodyQueues {
	s.MaxCountPerCycle = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMinCount(v int32) *ListQueuesResponseBodyQueues {
	s.MinCount = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetNodes(v *ListQueuesResponseBodyQueuesNodes) *ListQueuesResponseBodyQueues {
	s.Nodes = v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetQueueName(v string) *ListQueuesResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetTotalCores(v int32) *ListQueuesResponseBodyQueues {
	s.TotalCores = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetUpdateTime(v string) *ListQueuesResponseBodyQueues {
	s.UpdateTime = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetVSwitchIds(v []*string) *ListQueuesResponseBodyQueues {
	s.VSwitchIds = v
	return s
}

func (s *ListQueuesResponseBodyQueues) Validate() error {
	if s.ComputeNodes != nil {
		for _, item := range s.ComputeNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		if err := s.Nodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueuesResponseBodyQueuesNodes struct {
	// The number of compute nodes that are not ready.
	//
	// example:
	//
	// 2
	CreatingCounts *int32 `json:"CreatingCounts,omitempty" xml:"CreatingCounts,omitempty"`
	// The number of malfunctioning compute nodes.
	//
	// example:
	//
	// 0
	ExceptionCounts *int32 `json:"ExceptionCounts,omitempty" xml:"ExceptionCounts,omitempty"`
	// The number of running compute nodes.
	//
	// example:
	//
	// 1
	RunningCounts *int32 `json:"RunningCounts,omitempty" xml:"RunningCounts,omitempty"`
}

func (s ListQueuesResponseBodyQueuesNodes) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesNodes) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesNodes) GetCreatingCounts() *int32 {
	return s.CreatingCounts
}

func (s *ListQueuesResponseBodyQueuesNodes) GetExceptionCounts() *int32 {
	return s.ExceptionCounts
}

func (s *ListQueuesResponseBodyQueuesNodes) GetRunningCounts() *int32 {
	return s.RunningCounts
}

func (s *ListQueuesResponseBodyQueuesNodes) SetCreatingCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.CreatingCounts = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesNodes) SetExceptionCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.ExceptionCounts = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesNodes) SetRunningCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.RunningCounts = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesNodes) Validate() error {
	return dara.Validate(s)
}
