// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlanMaintenanceWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *CreatePlanMaintenanceWindowRequest
	GetEnable() *bool
	SetMinMaintenanceInterval(v int32) *CreatePlanMaintenanceWindowRequest
	GetMinMaintenanceInterval() *int32
	SetPlanWindowName(v string) *CreatePlanMaintenanceWindowRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *CreatePlanMaintenanceWindowRequest
	GetRegionId() *string
	SetSupportMaintenanceAction(v string) *CreatePlanMaintenanceWindowRequest
	GetSupportMaintenanceAction() *string
	SetTargetResource(v *CreatePlanMaintenanceWindowRequestTargetResource) *CreatePlanMaintenanceWindowRequest
	GetTargetResource() *CreatePlanMaintenanceWindowRequestTargetResource
	SetTimePeriod(v *CreatePlanMaintenanceWindowRequestTimePeriod) *CreatePlanMaintenanceWindowRequest
	GetTimePeriod() *CreatePlanMaintenanceWindowRequestTimePeriod
}

type CreatePlanMaintenanceWindowRequest struct {
	// Specifies whether to enable the window. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The minimum interval between two scheduled O&M events for the same instance. Valid values: 1 to 48. Unit: hours.
	//
	// example:
	//
	// 24
	MinMaintenanceInterval *int32 `json:"MinMaintenanceInterval,omitempty" xml:"MinMaintenanceInterval,omitempty"`
	// The name of the O&M window. You can customize the name. The name can be up to 200 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// WIndowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The region ID. This parameter is required for Alibaba Cloud SDK routing. However, specifying any region ID for this operation yields the same result. The O&M window applies to all ECS instances across the network that meet the resource filtering conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The O&M operations supported by the O&M window.
	//
	// This parameter is required.
	//
	// example:
	//
	// Reboot
	SupportMaintenanceAction *string `json:"SupportMaintenanceAction,omitempty" xml:"SupportMaintenanceAction,omitempty"`
	// The resources to which the O&M window applies.
	//
	// This parameter is required.
	TargetResource *CreatePlanMaintenanceWindowRequestTargetResource `json:"TargetResource,omitempty" xml:"TargetResource,omitempty" type:"Struct"`
	// The recurrence cycle of the window.
	//
	// This parameter is required.
	TimePeriod *CreatePlanMaintenanceWindowRequestTimePeriod `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty" type:"Struct"`
}

func (s CreatePlanMaintenanceWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowRequest) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreatePlanMaintenanceWindowRequest) GetMinMaintenanceInterval() *int32 {
	return s.MinMaintenanceInterval
}

func (s *CreatePlanMaintenanceWindowRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *CreatePlanMaintenanceWindowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePlanMaintenanceWindowRequest) GetSupportMaintenanceAction() *string {
	return s.SupportMaintenanceAction
}

func (s *CreatePlanMaintenanceWindowRequest) GetTargetResource() *CreatePlanMaintenanceWindowRequestTargetResource {
	return s.TargetResource
}

func (s *CreatePlanMaintenanceWindowRequest) GetTimePeriod() *CreatePlanMaintenanceWindowRequestTimePeriod {
	return s.TimePeriod
}

func (s *CreatePlanMaintenanceWindowRequest) SetEnable(v bool) *CreatePlanMaintenanceWindowRequest {
	s.Enable = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetMinMaintenanceInterval(v int32) *CreatePlanMaintenanceWindowRequest {
	s.MinMaintenanceInterval = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetPlanWindowName(v string) *CreatePlanMaintenanceWindowRequest {
	s.PlanWindowName = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetRegionId(v string) *CreatePlanMaintenanceWindowRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetSupportMaintenanceAction(v string) *CreatePlanMaintenanceWindowRequest {
	s.SupportMaintenanceAction = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetTargetResource(v *CreatePlanMaintenanceWindowRequestTargetResource) *CreatePlanMaintenanceWindowRequest {
	s.TargetResource = v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) SetTimePeriod(v *CreatePlanMaintenanceWindowRequestTimePeriod) *CreatePlanMaintenanceWindowRequest {
	s.TimePeriod = v
	return s
}

func (s *CreatePlanMaintenanceWindowRequest) Validate() error {
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

type CreatePlanMaintenanceWindowRequestTargetResource struct {
	// The ID of the resource group. This parameter is required when Scope is set to ResourceGroup.
	//
	// example:
	//
	// rg-aekzhm7pmnvcbty
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of resources configured for the O&M window.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tag
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tags to which the O&M window applies. This parameter is required when Scope is set to Tag.
	Tags []*CreatePlanMaintenanceWindowRequestTargetResourceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreatePlanMaintenanceWindowRequestTargetResource) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowRequestTargetResource) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) GetScope() *string {
	return s.Scope
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) GetTags() []*CreatePlanMaintenanceWindowRequestTargetResourceTags {
	return s.Tags
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) SetResourceGroupId(v string) *CreatePlanMaintenanceWindowRequestTargetResource {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) SetScope(v string) *CreatePlanMaintenanceWindowRequestTargetResource {
	s.Scope = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) SetTags(v []*CreatePlanMaintenanceWindowRequestTargetResourceTags) *CreatePlanMaintenanceWindowRequestTargetResource {
	s.Tags = v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTargetResource) Validate() error {
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

type CreatePlanMaintenanceWindowRequestTargetResourceTags struct {
	// The key of the tag to which the window applies.
	//
	// example:
	//
	// vms_qualification_Mr. Sun ID card_e5590864-1fef-4db2-b2a7-bd2d657fed43.png
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to which the window applies.
	//
	// example:
	//
	// 21.137.18.60
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePlanMaintenanceWindowRequestTargetResourceTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowRequestTargetResourceTags) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowRequestTargetResourceTags) GetKey() *string {
	return s.Key
}

func (s *CreatePlanMaintenanceWindowRequestTargetResourceTags) GetValue() *string {
	return s.Value
}

func (s *CreatePlanMaintenanceWindowRequestTargetResourceTags) SetKey(v string) *CreatePlanMaintenanceWindowRequestTargetResourceTags {
	s.Key = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTargetResourceTags) SetValue(v string) *CreatePlanMaintenanceWindowRequestTargetResourceTags {
	s.Value = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTargetResourceTags) Validate() error {
	return dara.Validate(s)
}

type CreatePlanMaintenanceWindowRequestTimePeriod struct {
	// The type of the recurrence cycle. Valid values: Daily and Weekly.
	//
	// This parameter is required.
	//
	// example:
	//
	// Weekly
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The time ranges of the O&M window recurrence cycle in the UTC time zone.
	//
	// This parameter is required.
	RangeList []*CreatePlanMaintenanceWindowRequestTimePeriodRangeList `json:"RangeList,omitempty" xml:"RangeList,omitempty" type:"Repeated"`
}

func (s CreatePlanMaintenanceWindowRequestTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowRequestTimePeriod) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriod) GetRangeList() []*CreatePlanMaintenanceWindowRequestTimePeriodRangeList {
	return s.RangeList
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriod) SetPeriodUnit(v string) *CreatePlanMaintenanceWindowRequestTimePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriod) SetRangeList(v []*CreatePlanMaintenanceWindowRequestTimePeriodRangeList) *CreatePlanMaintenanceWindowRequestTimePeriod {
	s.RangeList = v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriod) Validate() error {
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

type CreatePlanMaintenanceWindowRequestTimePeriodRangeList struct {
	// The end time of the O&M window.
	//
	// - If PeriodUnit is set to Weekly, the format is Monday,22:00. You can replace Monday with Tuesday, Wednesday, Thursday, Friday, Saturday, or Sunday.
	//
	// - If PeriodUnit is set to Daily, the format is 22:00.
	//
	// - The comma (,) is used as a delimiter. The first part represents the hour, which ranges from 00 to 23. The second part represents the minute. Currently, only 00 is supported.
	//
	// example:
	//
	// Tuesday,03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the O&M window.
	//
	// - If PeriodUnit is set to Weekly, the format is Monday,22:00. You can replace Monday with Tuesday, Wednesday, Thursday, Friday, Saturday, or Sunday.
	//
	// - If PeriodUnit is set to Daily, the format is 22:00.
	//
	// - The comma (,) is used as a delimiter. The first part represents the hour, which ranges from 00 to 23. The second part represents the minute. Currently, only 00 is supported.
	//
	// example:
	//
	// Monday,22:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreatePlanMaintenanceWindowRequestTimePeriodRangeList) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowRequestTimePeriodRangeList) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriodRangeList) GetEndTime() *string {
	return s.EndTime
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriodRangeList) GetStartTime() *string {
	return s.StartTime
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriodRangeList) SetEndTime(v string) *CreatePlanMaintenanceWindowRequestTimePeriodRangeList {
	s.EndTime = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriodRangeList) SetStartTime(v string) *CreatePlanMaintenanceWindowRequestTimePeriodRangeList {
	s.StartTime = &v
	return s
}

func (s *CreatePlanMaintenanceWindowRequestTimePeriodRangeList) Validate() error {
	return dara.Validate(s)
}
