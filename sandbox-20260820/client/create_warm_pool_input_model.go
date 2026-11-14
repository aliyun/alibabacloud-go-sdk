// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarmPoolInput interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityPolicy(v *WarmPoolCapacityPolicyInput) *CreateWarmPoolInput
	GetCapacityPolicy() *WarmPoolCapacityPolicyInput
	SetTemplateID(v string) *CreateWarmPoolInput
	GetTemplateID() *string
}

type CreateWarmPoolInput struct {
	// The warm pool capacity policy, which includes the base target capacity and optional scheduled rules.
	//
	// This parameter is required.
	CapacityPolicy *WarmPoolCapacityPolicyInput `json:"capacityPolicy,omitempty" xml:"capacityPolicy,omitempty"`
	// The template ID, which specifies the template associated with the warm pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl-example1234567890
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s CreateWarmPoolInput) String() string {
	return dara.Prettify(s)
}

func (s CreateWarmPoolInput) GoString() string {
	return s.String()
}

func (s *CreateWarmPoolInput) GetCapacityPolicy() *WarmPoolCapacityPolicyInput {
	return s.CapacityPolicy
}

func (s *CreateWarmPoolInput) GetTemplateID() *string {
	return s.TemplateID
}

func (s *CreateWarmPoolInput) SetCapacityPolicy(v *WarmPoolCapacityPolicyInput) *CreateWarmPoolInput {
	s.CapacityPolicy = v
	return s
}

func (s *CreateWarmPoolInput) SetTemplateID(v string) *CreateWarmPoolInput {
	s.TemplateID = &v
	return s
}

func (s *CreateWarmPoolInput) Validate() error {
	if s.CapacityPolicy != nil {
		if err := s.CapacityPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
