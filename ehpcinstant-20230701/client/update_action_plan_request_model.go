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
	// The expected resource scale of the execution plan. If ResourceType is set to VcpuCapacity, this field indicates that the execution plan expects to maintain 10,000 vCPUs.
	//
	// example:
	//
	// 1000
	DesiredCapacity *float32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether to enable the execution plan. Valid values:
	//
	// - true: Enables the execution plan.
	//
	// - false: Disables the execution plan.
	//
	// 	Notice: After the execution plan is disabled, already created Instant jobs are no longer automatically managed by the execution plan.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The scheduling interval of the execution plan. Unit: minutes. Valid values: 10 to 10080. Default value: 60.
	//
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
