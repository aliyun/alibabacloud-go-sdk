// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutElasticConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetMinInstances(v int64) *PutElasticConfigInput
	GetMinInstances() *int64
	SetResidentPoolId(v string) *PutElasticConfigInput
	GetResidentPoolId() *string
	SetScalingPolicies(v []*ScalingPolicy) *PutElasticConfigInput
	GetScalingPolicies() []*ScalingPolicy
	SetScheduledPolicies(v []*ScheduledPolicy) *PutElasticConfigInput
	GetScheduledPolicies() []*ScheduledPolicy
}

type PutElasticConfigInput struct {
	MinInstances      *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	ResidentPoolId    *string            `json:"residentPoolId,omitempty" xml:"residentPoolId,omitempty"`
	ScalingPolicies   []*ScalingPolicy   `json:"scalingPolicies" xml:"scalingPolicies" type:"Repeated"`
	ScheduledPolicies []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s PutElasticConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutElasticConfigInput) GoString() string {
	return s.String()
}

func (s *PutElasticConfigInput) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *PutElasticConfigInput) GetResidentPoolId() *string {
	return s.ResidentPoolId
}

func (s *PutElasticConfigInput) GetScalingPolicies() []*ScalingPolicy {
	return s.ScalingPolicies
}

func (s *PutElasticConfigInput) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *PutElasticConfigInput) SetMinInstances(v int64) *PutElasticConfigInput {
	s.MinInstances = &v
	return s
}

func (s *PutElasticConfigInput) SetResidentPoolId(v string) *PutElasticConfigInput {
	s.ResidentPoolId = &v
	return s
}

func (s *PutElasticConfigInput) SetScalingPolicies(v []*ScalingPolicy) *PutElasticConfigInput {
	s.ScalingPolicies = v
	return s
}

func (s *PutElasticConfigInput) SetScheduledPolicies(v []*ScheduledPolicy) *PutElasticConfigInput {
	s.ScheduledPolicies = v
	return s
}

func (s *PutElasticConfigInput) Validate() error {
	if s.ScalingPolicies != nil {
		for _, item := range s.ScalingPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduledPolicies != nil {
		for _, item := range s.ScheduledPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
