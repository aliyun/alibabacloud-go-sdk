// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRolloutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatch(v *UpdateServiceRolloutRequestBatch) *UpdateServiceRolloutRequest
	GetBatch() *UpdateServiceRolloutRequestBatch
	SetPartition(v *UpdateServiceRolloutRequestPartition) *UpdateServiceRolloutRequest
	GetPartition() *UpdateServiceRolloutRequestPartition
	SetPaused(v bool) *UpdateServiceRolloutRequest
	GetPaused() *bool
}

type UpdateServiceRolloutRequest struct {
	// **Batch*	-
	//
	// 	- Type: object
	//
	// 	- Required: no
	//
	// 	- Description: The batch rollout configuration. Used to adjust batch policy parameters. Mutually exclusive with Partition.
	Batch *UpdateServiceRolloutRequestBatch `json:"Batch,omitempty" xml:"Batch,omitempty" type:"Struct"`
	// **Partition*	-
	//
	// 	- Type: object
	//
	// 	- Required: no
	//
	// 	- Description: The partition rollout configuration. Used to adjust partition policy parameters. Mutually exclusive with Batch.
	Partition *UpdateServiceRolloutRequestPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	// Specifies whether to pause the rollout. A value of `true` indicates that the rollout is paused. A value of `false` indicates that the rollout is resumed.
	//
	// example:
	//
	// true
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
}

func (s UpdateServiceRolloutRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRolloutRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRolloutRequest) GetBatch() *UpdateServiceRolloutRequestBatch {
	return s.Batch
}

func (s *UpdateServiceRolloutRequest) GetPartition() *UpdateServiceRolloutRequestPartition {
	return s.Partition
}

func (s *UpdateServiceRolloutRequest) GetPaused() *bool {
	return s.Paused
}

func (s *UpdateServiceRolloutRequest) SetBatch(v *UpdateServiceRolloutRequestBatch) *UpdateServiceRolloutRequest {
	s.Batch = v
	return s
}

func (s *UpdateServiceRolloutRequest) SetPartition(v *UpdateServiceRolloutRequestPartition) *UpdateServiceRolloutRequest {
	s.Partition = v
	return s
}

func (s *UpdateServiceRolloutRequest) SetPaused(v bool) *UpdateServiceRolloutRequest {
	s.Paused = &v
	return s
}

func (s *UpdateServiceRolloutRequest) Validate() error {
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

type UpdateServiceRolloutRequestBatch struct {
	// The batch size. The number of replicas to update in each batch. Integer or percentage formats are supported. Default value: `"25%"`.
	//
	// example:
	//
	// 25%
	BatchSize *string `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The batch interval. The wait time between two rollouts. Supported units: `s` (seconds), `m` (minutes), and `h` (hours).
	//
	// example:
	//
	// 60s
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s UpdateServiceRolloutRequestBatch) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRolloutRequestBatch) GoString() string {
	return s.String()
}

func (s *UpdateServiceRolloutRequestBatch) GetBatchSize() *string {
	return s.BatchSize
}

func (s *UpdateServiceRolloutRequestBatch) GetInterval() *string {
	return s.Interval
}

func (s *UpdateServiceRolloutRequestBatch) SetBatchSize(v string) *UpdateServiceRolloutRequestBatch {
	s.BatchSize = &v
	return s
}

func (s *UpdateServiceRolloutRequestBatch) SetInterval(v string) *UpdateServiceRolloutRequestBatch {
	s.Interval = &v
	return s
}

func (s *UpdateServiceRolloutRequestBatch) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRolloutRequestPartition struct {
	// **Partition*	-
	//
	// 	- Type: string
	//
	// 	- Required: yes
	//
	// 	- Description: The partition value. Two formats are supported:
	//
	//   1. Integer format: such as "5", which indicates that 5 old replicas are retained.
	//
	//   2. Percentage format: such as "50%", which indicates that 50% of the old replicas are retained.
	//
	//   Adjustment strategies:
	//
	//   - Increase the value: Roll back to the old version and increase the proportion of old replicas.
	//
	//   - Decrease the value: Proceed to the new version and decrease the proportion of old replicas.
	//
	//   - Set to "0%": Complete the full update and switch all replicas to the new version.
	//
	// 	- Example: 30%
	//
	// example:
	//
	// 30%
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s UpdateServiceRolloutRequestPartition) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRolloutRequestPartition) GoString() string {
	return s.String()
}

func (s *UpdateServiceRolloutRequestPartition) GetPartition() *string {
	return s.Partition
}

func (s *UpdateServiceRolloutRequestPartition) SetPartition(v string) *UpdateServiceRolloutRequestPartition {
	s.Partition = &v
	return s
}

func (s *UpdateServiceRolloutRequestPartition) Validate() error {
	return dara.Validate(s)
}
