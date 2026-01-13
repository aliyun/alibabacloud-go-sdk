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
	SetIntervalMinutes(v int32) *UpdateActionPlanRequest
	GetIntervalMinutes() *int32
}

type UpdateActionPlanRequest struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// The expected scale of resources for the execution plan. If the ResourceType parameter is set to VcpuCapacity, the execution plan is expected to have 10000 vCPUs.
	//
	// example:
	//
	// 1000
	DesiredCapacity *float32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Whether to enable the execution plan. Valid values:
	//
	// 	- true: enables the execution plan.
	//
	// 	- false: The execution plan is disabled.
	//
	//     **
	//
	//     **Note:*	- After an execution plan is disabled, the created Instant jobs are not automatically managed by the execution plan.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 60
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
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

func (s *UpdateActionPlanRequest) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
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

func (s *UpdateActionPlanRequest) SetIntervalMinutes(v int32) *UpdateActionPlanRequest {
	s.IntervalMinutes = &v
	return s
}

func (s *UpdateActionPlanRequest) Validate() error {
	return dara.Validate(s)
}
