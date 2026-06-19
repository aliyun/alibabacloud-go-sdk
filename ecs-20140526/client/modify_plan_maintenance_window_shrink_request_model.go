// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlanMaintenanceWindowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifyPlanMaintenanceWindowShrinkRequest
	GetEnable() *bool
	SetMinMaintenanceInterval(v int32) *ModifyPlanMaintenanceWindowShrinkRequest
	GetMinMaintenanceInterval() *int32
	SetPlanWindowId(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetPlanWindowId() *string
	SetPlanWindowName(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetRegionId() *string
	SetSupportMaintenanceAction(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetSupportMaintenanceAction() *string
	SetTargetResourceShrink(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetTargetResourceShrink() *string
	SetTimePeriodShrink(v string) *ModifyPlanMaintenanceWindowShrinkRequest
	GetTimePeriodShrink() *string
}

type ModifyPlanMaintenanceWindowShrinkRequest struct {
	// Specifies whether to enable or disable the O&M window. Leave this parameter empty if no modification is needed.
	//
	// example:
	//
	// false
	Enable                 *bool  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MinMaintenanceInterval *int32 `json:"MinMaintenanceInterval,omitempty" xml:"MinMaintenanceInterval,omitempty"`
	// The ID of the O&M window to modify. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// pw-bp1au1w8v8a1yer65g5k
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// The name of the O&M window. Leave this parameter empty if no modification is needed.
	//
	// example:
	//
	// WIndowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The region ID of the instance. You can call DescribeRegions to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The supported O&M actions. Leave this parameter empty if no modification is needed.
	//
	// example:
	//
	// Reboot
	SupportMaintenanceAction *string `json:"SupportMaintenanceAction,omitempty" xml:"SupportMaintenanceAction,omitempty"`
	// The resource to which the O&M window applies. Leave this parameter empty if no modification is needed.
	TargetResourceShrink *string `json:"TargetResource,omitempty" xml:"TargetResource,omitempty"`
	// The recurring cycle of the O&M window. Leave this parameter empty if no modification is needed.
	TimePeriodShrink *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
}

func (s ModifyPlanMaintenanceWindowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetMinMaintenanceInterval() *int32 {
	return s.MinMaintenanceInterval
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetSupportMaintenanceAction() *string {
	return s.SupportMaintenanceAction
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetTargetResourceShrink() *string {
	return s.TargetResourceShrink
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) GetTimePeriodShrink() *string {
	return s.TimePeriodShrink
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetEnable(v bool) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.Enable = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetMinMaintenanceInterval(v int32) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.MinMaintenanceInterval = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetPlanWindowId(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.PlanWindowId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetPlanWindowName(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.PlanWindowName = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetRegionId(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetSupportMaintenanceAction(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.SupportMaintenanceAction = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetTargetResourceShrink(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.TargetResourceShrink = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) SetTimePeriodShrink(v string) *ModifyPlanMaintenanceWindowShrinkRequest {
	s.TimePeriodShrink = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
