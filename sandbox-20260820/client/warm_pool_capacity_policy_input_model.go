// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmPoolCapacityPolicyInput interface {
	dara.Model
	String() string
	GoString() string
	SetMinReadySandboxes(v int64) *WarmPoolCapacityPolicyInput
	GetMinReadySandboxes() *int64
	SetScheduledPolicies(v []*WarmPoolScheduledPolicyInput) *WarmPoolCapacityPolicyInput
	GetScheduledPolicies() []*WarmPoolScheduledPolicyInput
}

type WarmPoolCapacityPolicyInput struct {
	// The base target number of ready sandboxes, which is required and can be set to 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MinReadySandboxes *int64 `json:"minReadySandboxes,omitempty" xml:"minReadySandboxes,omitempty"`
	// The list of scheduled capacity rules, which is optional.
	ScheduledPolicies []*WarmPoolScheduledPolicyInput `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s WarmPoolCapacityPolicyInput) String() string {
	return dara.Prettify(s)
}

func (s WarmPoolCapacityPolicyInput) GoString() string {
	return s.String()
}

func (s *WarmPoolCapacityPolicyInput) GetMinReadySandboxes() *int64 {
	return s.MinReadySandboxes
}

func (s *WarmPoolCapacityPolicyInput) GetScheduledPolicies() []*WarmPoolScheduledPolicyInput {
	return s.ScheduledPolicies
}

func (s *WarmPoolCapacityPolicyInput) SetMinReadySandboxes(v int64) *WarmPoolCapacityPolicyInput {
	s.MinReadySandboxes = &v
	return s
}

func (s *WarmPoolCapacityPolicyInput) SetScheduledPolicies(v []*WarmPoolScheduledPolicyInput) *WarmPoolCapacityPolicyInput {
	s.ScheduledPolicies = v
	return s
}

func (s *WarmPoolCapacityPolicyInput) Validate() error {
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
