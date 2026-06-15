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
	// Specifies whether to enable the maintenance window.
	//
	// - **true**: Enables the maintenance window.
	//
	// - **false**: Disables the maintenance window.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable                 *bool  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MinMaintenanceInterval *int32 `json:"MinMaintenanceInterval,omitempty" xml:"MinMaintenanceInterval,omitempty"`
	// The name of the maintenance window. The name can be up to 200 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// WIndowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maintenance operation supported by the maintenance window.
	//
	// This parameter is required.
	//
	// example:
	//
	// Reboot
	SupportMaintenanceAction *string `json:"SupportMaintenanceAction,omitempty" xml:"SupportMaintenanceAction,omitempty"`
	// The resources to which the maintenance window applies.
	//
	// This parameter is required.
	TargetResource *CreatePlanMaintenanceWindowRequestTargetResource `json:"TargetResource,omitempty" xml:"TargetResource,omitempty" type:"Struct"`
	// The recurring schedule for the maintenance window.
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
	// The ID of the resource group. This parameter is required if `Scope` is set to `ResourceGroup`.
	//
	// example:
	//
	// rg-aekzhm7pmnvcbty
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scope of resources to which the maintenance window applies.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tag
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tags of the resources to which the maintenance window applies. This parameter is required if `Scope` is set to `Tag`.
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
	// The tag key.
	//
	// example:
	//
	// vms_qualification_孙总身份证_e5590864-1fef-4db2-b2a7-bd2d657fed43.png
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	// Specifies how often the maintenance window recurs.
	//
	// This parameter is required.
	//
	// example:
	//
	// Weekly
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The time ranges of the recurring maintenance window. All times are in UTC.
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
	// The end time of the maintenance window.
	//
	// - If `PeriodUnit` is set to `Weekly`, use the format `Day,HH:mm`. Valid values for `Day` are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
	//
	// - If `PeriodUnit` is set to `Daily`, use the format `HH:mm`.
	//
	// - The time is in `HH:mm` format, where `HH` is the hour (00-23) and `mm` is the minute. Only `00` is supported for the minute.
	//
	// example:
	//
	// Tuesday,03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// - If `PeriodUnit` is set to `Weekly`, use the format `Day,HH:mm`. Valid values for `Day` are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
	//
	// - If `PeriodUnit` is set to `Daily`, use the format `HH:mm`.
	//
	// - The time is in `HH:mm` format, where `HH` is the hour (00-23) and `mm` is the minute. Only `00` is supported for the minute.
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
