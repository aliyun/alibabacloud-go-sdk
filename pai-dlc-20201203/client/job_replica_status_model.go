// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobReplicaStatus interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int32) *JobReplicaStatus
	GetActive() *int32
	SetDequeued(v int32) *JobReplicaStatus
	GetDequeued() *int32
	SetEstimatedAutoScalingSpec(v *AutoScalingSpec) *JobReplicaStatus
	GetEstimatedAutoScalingSpec() *AutoScalingSpec
	SetEstimatedPodCount(v int64) *JobReplicaStatus
	GetEstimatedPodCount() *int64
	SetEstimatedResourceConfig(v *ResourceConfig) *JobReplicaStatus
	GetEstimatedResourceConfig() *ResourceConfig
	SetQueuing(v int32) *JobReplicaStatus
	GetQueuing() *int32
	SetType(v string) *JobReplicaStatus
	GetType() *string
}

type JobReplicaStatus struct {
	Active   *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	Dequeued *int32 `json:"Dequeued,omitempty" xml:"Dequeued,omitempty"`
	// if can be null:
	// true
	EstimatedAutoScalingSpec *AutoScalingSpec `json:"EstimatedAutoScalingSpec,omitempty" xml:"EstimatedAutoScalingSpec,omitempty"`
	EstimatedPodCount        *int64           `json:"EstimatedPodCount,omitempty" xml:"EstimatedPodCount,omitempty"`
	// if can be null:
	// true
	EstimatedResourceConfig *ResourceConfig `json:"EstimatedResourceConfig,omitempty" xml:"EstimatedResourceConfig,omitempty"`
	Queuing                 *int32          `json:"Queuing,omitempty" xml:"Queuing,omitempty"`
	Type                    *string         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s JobReplicaStatus) String() string {
	return dara.Prettify(s)
}

func (s JobReplicaStatus) GoString() string {
	return s.String()
}

func (s *JobReplicaStatus) GetActive() *int32 {
	return s.Active
}

func (s *JobReplicaStatus) GetDequeued() *int32 {
	return s.Dequeued
}

func (s *JobReplicaStatus) GetEstimatedAutoScalingSpec() *AutoScalingSpec {
	return s.EstimatedAutoScalingSpec
}

func (s *JobReplicaStatus) GetEstimatedPodCount() *int64 {
	return s.EstimatedPodCount
}

func (s *JobReplicaStatus) GetEstimatedResourceConfig() *ResourceConfig {
	return s.EstimatedResourceConfig
}

func (s *JobReplicaStatus) GetQueuing() *int32 {
	return s.Queuing
}

func (s *JobReplicaStatus) GetType() *string {
	return s.Type
}

func (s *JobReplicaStatus) SetActive(v int32) *JobReplicaStatus {
	s.Active = &v
	return s
}

func (s *JobReplicaStatus) SetDequeued(v int32) *JobReplicaStatus {
	s.Dequeued = &v
	return s
}

func (s *JobReplicaStatus) SetEstimatedAutoScalingSpec(v *AutoScalingSpec) *JobReplicaStatus {
	s.EstimatedAutoScalingSpec = v
	return s
}

func (s *JobReplicaStatus) SetEstimatedPodCount(v int64) *JobReplicaStatus {
	s.EstimatedPodCount = &v
	return s
}

func (s *JobReplicaStatus) SetEstimatedResourceConfig(v *ResourceConfig) *JobReplicaStatus {
	s.EstimatedResourceConfig = v
	return s
}

func (s *JobReplicaStatus) SetQueuing(v int32) *JobReplicaStatus {
	s.Queuing = &v
	return s
}

func (s *JobReplicaStatus) SetType(v string) *JobReplicaStatus {
	s.Type = &v
	return s
}

func (s *JobReplicaStatus) Validate() error {
	if s.EstimatedAutoScalingSpec != nil {
		if err := s.EstimatedAutoScalingSpec.Validate(); err != nil {
			return err
		}
	}
	if s.EstimatedResourceConfig != nil {
		if err := s.EstimatedResourceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
