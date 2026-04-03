// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMinInstances(v int64) *ScalingConfig
	GetMinInstances() *int64
	SetScheduledPolicies(v []*ScheduledPolicy) *ScalingConfig
	GetScheduledPolicies() []*ScheduledPolicy
}

type ScalingConfig struct {
	// example:
	//
	// 2
	MinInstances      *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	ScheduledPolicies []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s ScalingConfig) String() string {
	return dara.Prettify(s)
}

func (s ScalingConfig) GoString() string {
	return s.String()
}

func (s *ScalingConfig) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *ScalingConfig) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *ScalingConfig) SetMinInstances(v int64) *ScalingConfig {
	s.MinInstances = &v
	return s
}

func (s *ScalingConfig) SetScheduledPolicies(v []*ScheduledPolicy) *ScalingConfig {
	s.ScheduledPolicies = v
	return s
}

func (s *ScalingConfig) Validate() error {
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
