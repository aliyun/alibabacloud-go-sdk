// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmPoolCapacityPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetMinReadySandboxes(v int64) *WarmPoolCapacityPolicy
	GetMinReadySandboxes() *int64
	SetScheduledPolicies(v []*WarmPoolScheduledPolicy) *WarmPoolCapacityPolicy
	GetScheduledPolicies() []*WarmPoolScheduledPolicy
}

type WarmPoolCapacityPolicy struct {
	// The base target number of ready sandboxes. This parameter is required and can be set to 0.
	//
	// example:
	//
	// 0
	MinReadySandboxes *int64 `json:"minReadySandboxes,omitempty" xml:"minReadySandboxes,omitempty"`
	// The list of scheduled capacity rules. This parameter is optional.
	ScheduledPolicies []*WarmPoolScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s WarmPoolCapacityPolicy) String() string {
	return dara.Prettify(s)
}

func (s WarmPoolCapacityPolicy) GoString() string {
	return s.String()
}

func (s *WarmPoolCapacityPolicy) GetMinReadySandboxes() *int64 {
	return s.MinReadySandboxes
}

func (s *WarmPoolCapacityPolicy) GetScheduledPolicies() []*WarmPoolScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *WarmPoolCapacityPolicy) SetMinReadySandboxes(v int64) *WarmPoolCapacityPolicy {
	s.MinReadySandboxes = &v
	return s
}

func (s *WarmPoolCapacityPolicy) SetScheduledPolicies(v []*WarmPoolScheduledPolicy) *WarmPoolCapacityPolicy {
	s.ScheduledPolicies = v
	return s
}

func (s *WarmPoolCapacityPolicy) Validate() error {
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
