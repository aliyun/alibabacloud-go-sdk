// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlanMaintenanceWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifyPlanMaintenanceWindowRequest
	GetEnable() *bool
	SetMinMaintenanceInterval(v int32) *ModifyPlanMaintenanceWindowRequest
	GetMinMaintenanceInterval() *int32
	SetPlanWindowId(v string) *ModifyPlanMaintenanceWindowRequest
	GetPlanWindowId() *string
	SetPlanWindowName(v string) *ModifyPlanMaintenanceWindowRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *ModifyPlanMaintenanceWindowRequest
	GetRegionId() *string
	SetSupportMaintenanceAction(v string) *ModifyPlanMaintenanceWindowRequest
	GetSupportMaintenanceAction() *string
	SetTargetResource(v *ModifyPlanMaintenanceWindowRequestTargetResource) *ModifyPlanMaintenanceWindowRequest
	GetTargetResource() *ModifyPlanMaintenanceWindowRequestTargetResource
	SetTimePeriod(v *ModifyPlanMaintenanceWindowRequestTimePeriod) *ModifyPlanMaintenanceWindowRequest
	GetTimePeriod() *ModifyPlanMaintenanceWindowRequestTimePeriod
}

type ModifyPlanMaintenanceWindowRequest struct {
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
	TargetResource *ModifyPlanMaintenanceWindowRequestTargetResource `json:"TargetResource,omitempty" xml:"TargetResource,omitempty" type:"Struct"`
	// The recurring cycle of the O&M window. Leave this parameter empty if no modification is needed.
	TimePeriod *ModifyPlanMaintenanceWindowRequestTimePeriod `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty" type:"Struct"`
}

func (s ModifyPlanMaintenanceWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyPlanMaintenanceWindowRequest) GetMinMaintenanceInterval() *int32 {
	return s.MinMaintenanceInterval
}

func (s *ModifyPlanMaintenanceWindowRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *ModifyPlanMaintenanceWindowRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *ModifyPlanMaintenanceWindowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPlanMaintenanceWindowRequest) GetSupportMaintenanceAction() *string {
	return s.SupportMaintenanceAction
}

func (s *ModifyPlanMaintenanceWindowRequest) GetTargetResource() *ModifyPlanMaintenanceWindowRequestTargetResource {
	return s.TargetResource
}

func (s *ModifyPlanMaintenanceWindowRequest) GetTimePeriod() *ModifyPlanMaintenanceWindowRequestTimePeriod {
	return s.TimePeriod
}

func (s *ModifyPlanMaintenanceWindowRequest) SetEnable(v bool) *ModifyPlanMaintenanceWindowRequest {
	s.Enable = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetMinMaintenanceInterval(v int32) *ModifyPlanMaintenanceWindowRequest {
	s.MinMaintenanceInterval = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetPlanWindowId(v string) *ModifyPlanMaintenanceWindowRequest {
	s.PlanWindowId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetPlanWindowName(v string) *ModifyPlanMaintenanceWindowRequest {
	s.PlanWindowName = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetRegionId(v string) *ModifyPlanMaintenanceWindowRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetSupportMaintenanceAction(v string) *ModifyPlanMaintenanceWindowRequest {
	s.SupportMaintenanceAction = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetTargetResource(v *ModifyPlanMaintenanceWindowRequestTargetResource) *ModifyPlanMaintenanceWindowRequest {
	s.TargetResource = v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) SetTimePeriod(v *ModifyPlanMaintenanceWindowRequestTimePeriod) *ModifyPlanMaintenanceWindowRequest {
	s.TimePeriod = v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequest) Validate() error {
	if s.TargetResource != nil {
		if err := s.TargetResource.Validate(); err != nil {
			return err
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPlanMaintenanceWindowRequestTargetResource struct {
	// The ID of the resource group to which the O&M window applies. This parameter is required only when Scope is set to ResourceGroup.
	//
	// example:
	//
	// rg-acfmy4cc27vsvia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type for the O&M window configuration.
	//
	// example:
	//
	// Tag
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tags to which the O&M window applies. This parameter is required only when Scope is set to Tag.
	Tags []*ModifyPlanMaintenanceWindowRequestTargetResourceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ModifyPlanMaintenanceWindowRequestTargetResource) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowRequestTargetResource) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) GetScope() *string {
	return s.Scope
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) GetTags() []*ModifyPlanMaintenanceWindowRequestTargetResourceTags {
	return s.Tags
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) SetResourceGroupId(v string) *ModifyPlanMaintenanceWindowRequestTargetResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) SetScope(v string) *ModifyPlanMaintenanceWindowRequestTargetResource {
	s.Scope = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) SetTags(v []*ModifyPlanMaintenanceWindowRequestTargetResourceTags) *ModifyPlanMaintenanceWindowRequestTargetResource {
	s.Tags = v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResource) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPlanMaintenanceWindowRequestTargetResourceTags struct {
	// The key of the tag to which the O&M window applies.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to which the O&M window applies.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyPlanMaintenanceWindowRequestTargetResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowRequestTargetResourceTags) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResourceTags) GetKey() *string {
	return s.Key
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResourceTags) GetValue() *string {
	return s.Value
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResourceTags) SetKey(v string) *ModifyPlanMaintenanceWindowRequestTargetResourceTags {
	s.Key = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResourceTags) SetValue(v string) *ModifyPlanMaintenanceWindowRequestTargetResourceTags {
	s.Value = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTargetResourceTags) Validate() error {
	return dara.Validate(s)
}

type ModifyPlanMaintenanceWindowRequestTimePeriod struct {
	// The cycle type. Valid values: Daily and Weekly.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The time ranges of the O&M window recurring cycle (UTC time zone).
	RangeList []*ModifyPlanMaintenanceWindowRequestTimePeriodRangeList `json:"RangeList,omitempty" xml:"RangeList,omitempty" type:"Repeated"`
}

func (s ModifyPlanMaintenanceWindowRequestTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowRequestTimePeriod) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriod) GetRangeList() []*ModifyPlanMaintenanceWindowRequestTimePeriodRangeList {
	return s.RangeList
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriod) SetPeriodUnit(v string) *ModifyPlanMaintenanceWindowRequestTimePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriod) SetRangeList(v []*ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) *ModifyPlanMaintenanceWindowRequestTimePeriod {
	s.RangeList = v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriod) Validate() error {
	if s.RangeList != nil {
		for _, item := range s.RangeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPlanMaintenanceWindowRequestTimePeriodRangeList struct {
	// The end time of the O&M window.
	//
	// - If PeriodUnit is set to Weekly, the format is Monday,22:00. Monday can be replaced with Tuesday, Wednesday, Thursday, Friday, Saturday, or Sunday.
	//
	// - If PeriodUnit is set to Daily, the format is 22:00.
	//
	// - The comma delimiter separates two parts. The first part represents the hour, with valid values from 00 to 23. The second part represents the minute, which currently supports only 00.
	//
	// example:
	//
	// Tuesday,03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the O&M window.
	//
	// - If PeriodUnit is set to Weekly, the format is Monday,22:00. Monday can be replaced with Tuesday, Wednesday, Thursday, Friday, Saturday, or Sunday.
	//
	// - If PeriodUnit is set to Daily, the format is 22:00.
	//
	// - The comma delimiter separates two parts. The first part represents the hour, with valid values from 00 to 23. The second part represents the minute, which currently supports only 00.
	//
	// example:
	//
	// Monday,22:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) SetEndTime(v string) *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList {
	s.EndTime = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) SetStartTime(v string) *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList {
	s.StartTime = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowRequestTimePeriodRangeList) Validate() error {
	return dara.Validate(s)
}
