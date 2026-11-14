// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarmPoolInput interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityPolicy(v *WarmPoolCapacityPolicyInput) *UpdateWarmPoolInput
	GetCapacityPolicy() *WarmPoolCapacityPolicyInput
}

type UpdateWarmPoolInput struct {
	// The warm pool capacity policy, which includes the base target capacity and optional scheduled rules.
	//
	// This parameter is required.
	CapacityPolicy *WarmPoolCapacityPolicyInput `json:"capacityPolicy,omitempty" xml:"capacityPolicy,omitempty"`
}

func (s UpdateWarmPoolInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarmPoolInput) GoString() string {
	return s.String()
}

func (s *UpdateWarmPoolInput) GetCapacityPolicy() *WarmPoolCapacityPolicyInput {
	return s.CapacityPolicy
}

func (s *UpdateWarmPoolInput) SetCapacityPolicy(v *WarmPoolCapacityPolicyInput) *UpdateWarmPoolInput {
	s.CapacityPolicy = v
	return s
}

func (s *UpdateWarmPoolInput) Validate() error {
	if s.CapacityPolicy != nil {
		if err := s.CapacityPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
