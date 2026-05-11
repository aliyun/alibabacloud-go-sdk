// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRolloutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatch(v *CreateServiceRolloutRequestBatch) *CreateServiceRolloutRequest
	GetBatch() *CreateServiceRolloutRequestBatch
	SetPartition(v *CreateServiceRolloutRequestPartition) *CreateServiceRolloutRequest
	GetPartition() *CreateServiceRolloutRequestPartition
	SetPaused(v bool) *CreateServiceRolloutRequest
	GetPaused() *bool
}

type CreateServiceRolloutRequest struct {
	Batch     *CreateServiceRolloutRequestBatch     `json:"Batch,omitempty" xml:"Batch,omitempty" type:"Struct"`
	Partition *CreateServiceRolloutRequestPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	// example:
	//
	// False
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
}

func (s CreateServiceRolloutRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRolloutRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRolloutRequest) GetBatch() *CreateServiceRolloutRequestBatch {
	return s.Batch
}

func (s *CreateServiceRolloutRequest) GetPartition() *CreateServiceRolloutRequestPartition {
	return s.Partition
}

func (s *CreateServiceRolloutRequest) GetPaused() *bool {
	return s.Paused
}

func (s *CreateServiceRolloutRequest) SetBatch(v *CreateServiceRolloutRequestBatch) *CreateServiceRolloutRequest {
	s.Batch = v
	return s
}

func (s *CreateServiceRolloutRequest) SetPartition(v *CreateServiceRolloutRequestPartition) *CreateServiceRolloutRequest {
	s.Partition = v
	return s
}

func (s *CreateServiceRolloutRequest) SetPaused(v bool) *CreateServiceRolloutRequest {
	s.Paused = &v
	return s
}

func (s *CreateServiceRolloutRequest) Validate() error {
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

type CreateServiceRolloutRequestBatch struct {
	// example:
	//
	// 1
	BatchSize *string `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// example:
	//
	// 5m
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s CreateServiceRolloutRequestBatch) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRolloutRequestBatch) GoString() string {
	return s.String()
}

func (s *CreateServiceRolloutRequestBatch) GetBatchSize() *string {
	return s.BatchSize
}

func (s *CreateServiceRolloutRequestBatch) GetInterval() *string {
	return s.Interval
}

func (s *CreateServiceRolloutRequestBatch) SetBatchSize(v string) *CreateServiceRolloutRequestBatch {
	s.BatchSize = &v
	return s
}

func (s *CreateServiceRolloutRequestBatch) SetInterval(v string) *CreateServiceRolloutRequestBatch {
	s.Interval = &v
	return s
}

func (s *CreateServiceRolloutRequestBatch) Validate() error {
	return dara.Validate(s)
}

type CreateServiceRolloutRequestPartition struct {
	// example:
	//
	// 1
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s CreateServiceRolloutRequestPartition) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRolloutRequestPartition) GoString() string {
	return s.String()
}

func (s *CreateServiceRolloutRequestPartition) GetPartition() *string {
	return s.Partition
}

func (s *CreateServiceRolloutRequestPartition) SetPartition(v string) *CreateServiceRolloutRequestPartition {
	s.Partition = &v
	return s
}

func (s *CreateServiceRolloutRequestPartition) Validate() error {
	return dara.Validate(s)
}
