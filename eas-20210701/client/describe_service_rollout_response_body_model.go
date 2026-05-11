// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceRolloutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceRolloutResponseBody
	GetRequestId() *string
	SetRollout(v *DescribeServiceRolloutResponseBodyRollout) *DescribeServiceRolloutResponseBody
	GetRollout() *DescribeServiceRolloutResponseBodyRollout
}

type DescribeServiceRolloutResponseBody struct {
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rollout   *DescribeServiceRolloutResponseBodyRollout `json:"Rollout,omitempty" xml:"Rollout,omitempty" type:"Struct"`
}

func (s DescribeServiceRolloutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceRolloutResponseBody) GetRollout() *DescribeServiceRolloutResponseBodyRollout {
	return s.Rollout
}

func (s *DescribeServiceRolloutResponseBody) SetRequestId(v string) *DescribeServiceRolloutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceRolloutResponseBody) SetRollout(v *DescribeServiceRolloutResponseBodyRollout) *DescribeServiceRolloutResponseBody {
	s.Rollout = v
	return s
}

func (s *DescribeServiceRolloutResponseBody) Validate() error {
	if s.Rollout != nil {
		if err := s.Rollout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceRolloutResponseBodyRollout struct {
	Status   *DescribeServiceRolloutResponseBodyRolloutStatus   `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	Strategy *DescribeServiceRolloutResponseBodyRolloutStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeServiceRolloutResponseBodyRollout) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBodyRollout) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBodyRollout) GetStatus() *DescribeServiceRolloutResponseBodyRolloutStatus {
	return s.Status
}

func (s *DescribeServiceRolloutResponseBodyRollout) GetStrategy() *DescribeServiceRolloutResponseBodyRolloutStrategy {
	return s.Strategy
}

func (s *DescribeServiceRolloutResponseBodyRollout) SetStatus(v *DescribeServiceRolloutResponseBodyRolloutStatus) *DescribeServiceRolloutResponseBodyRollout {
	s.Status = v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRollout) SetStrategy(v *DescribeServiceRolloutResponseBodyRolloutStrategy) *DescribeServiceRolloutResponseBodyRollout {
	s.Strategy = v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRollout) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceRolloutResponseBodyRolloutStatus struct {
	// example:
	//
	// service-abc123-v1
	CurrentRevision *string `json:"CurrentRevision,omitempty" xml:"CurrentRevision,omitempty"`
	// example:
	//
	// 2026/05/08 16:10:56
	NextBatchStartTime *string `json:"NextBatchStartTime,omitempty" xml:"NextBatchStartTime,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 10
	TotalReplicas *int32 `json:"TotalReplicas,omitempty" xml:"TotalReplicas,omitempty"`
	// example:
	//
	// service-abc123-v2
	UpdateRevision *string `json:"UpdateRevision,omitempty" xml:"UpdateRevision,omitempty"`
	// example:
	//
	// 5
	UpdatedReplicas *int32 `json:"UpdatedReplicas,omitempty" xml:"UpdatedReplicas,omitempty"`
}

func (s DescribeServiceRolloutResponseBodyRolloutStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBodyRolloutStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetCurrentRevision() *string {
	return s.CurrentRevision
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetNextBatchStartTime() *string {
	return s.NextBatchStartTime
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetPhase() *string {
	return s.Phase
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetTotalReplicas() *int32 {
	return s.TotalReplicas
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetUpdateRevision() *string {
	return s.UpdateRevision
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) GetUpdatedReplicas() *int32 {
	return s.UpdatedReplicas
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetCurrentRevision(v string) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.CurrentRevision = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetNextBatchStartTime(v string) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.NextBatchStartTime = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetPhase(v string) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.Phase = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetTotalReplicas(v int32) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.TotalReplicas = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetUpdateRevision(v string) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.UpdateRevision = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) SetUpdatedReplicas(v int32) *DescribeServiceRolloutResponseBodyRolloutStatus {
	s.UpdatedReplicas = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceRolloutResponseBodyRolloutStrategy struct {
	Batch     *DescribeServiceRolloutResponseBodyRolloutStrategyBatch     `json:"Batch,omitempty" xml:"Batch,omitempty" type:"Struct"`
	Partition *DescribeServiceRolloutResponseBodyRolloutStrategyPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategy) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategy) GetBatch() *DescribeServiceRolloutResponseBodyRolloutStrategyBatch {
	return s.Batch
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategy) GetPartition() *DescribeServiceRolloutResponseBodyRolloutStrategyPartition {
	return s.Partition
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategy) SetBatch(v *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) *DescribeServiceRolloutResponseBodyRolloutStrategy {
	s.Batch = v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategy) SetPartition(v *DescribeServiceRolloutResponseBodyRolloutStrategyPartition) *DescribeServiceRolloutResponseBodyRolloutStrategy {
	s.Partition = v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategy) Validate() error {
	if s.Batch != nil {
		if err := s.Batch.Validate(); err != nil {
			return err
		}
	}
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceRolloutResponseBodyRolloutStrategyBatch struct {
	// example:
	//
	// 1
	BatchSize *string `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// example:
	//
	// 5m
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategyBatch) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategyBatch) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) GetBatchSize() *string {
	return s.BatchSize
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) GetInterval() *string {
	return s.Interval
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) SetBatchSize(v string) *DescribeServiceRolloutResponseBodyRolloutStrategyBatch {
	s.BatchSize = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) SetInterval(v string) *DescribeServiceRolloutResponseBodyRolloutStrategyBatch {
	s.Interval = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyBatch) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceRolloutResponseBodyRolloutStrategyPartition struct {
	// example:
	//
	// 50%
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategyPartition) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponseBodyRolloutStrategyPartition) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyPartition) GetPartition() *string {
	return s.Partition
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyPartition) SetPartition(v string) *DescribeServiceRolloutResponseBodyRolloutStrategyPartition {
	s.Partition = &v
	return s
}

func (s *DescribeServiceRolloutResponseBodyRolloutStrategyPartition) Validate() error {
	return dara.Validate(s)
}
