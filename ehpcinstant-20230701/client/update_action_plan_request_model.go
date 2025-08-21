// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateActionPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *UpdateActionPlanRequest
	GetActionPlanId() *string
	SetDesiredCapacity(v float32) *UpdateActionPlanRequest
	GetDesiredCapacity() *float32
	SetEnabled(v string) *UpdateActionPlanRequest
	GetEnabled() *string
}

type UpdateActionPlanRequest struct {
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// example:
	//
	// 1000
	DesiredCapacity *float32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s UpdateActionPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateActionPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateActionPlanRequest) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *UpdateActionPlanRequest) GetDesiredCapacity() *float32 {
	return s.DesiredCapacity
}

func (s *UpdateActionPlanRequest) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateActionPlanRequest) SetActionPlanId(v string) *UpdateActionPlanRequest {
	s.ActionPlanId = &v
	return s
}

func (s *UpdateActionPlanRequest) SetDesiredCapacity(v float32) *UpdateActionPlanRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *UpdateActionPlanRequest) SetEnabled(v string) *UpdateActionPlanRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateActionPlanRequest) Validate() error {
	return dara.Validate(s)
}
