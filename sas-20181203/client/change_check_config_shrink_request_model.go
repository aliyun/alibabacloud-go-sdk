// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedCheck(v []*ChangeCheckConfigShrinkRequestAddedCheck) *ChangeCheckConfigShrinkRequest
	GetAddedCheck() []*ChangeCheckConfigShrinkRequestAddedCheck
	SetConfigRequirementIdsShrink(v string) *ChangeCheckConfigShrinkRequest
	GetConfigRequirementIdsShrink() *string
	SetConfigStandardIdsShrink(v string) *ChangeCheckConfigShrinkRequest
	GetConfigStandardIdsShrink() *string
	SetConfigure(v string) *ChangeCheckConfigShrinkRequest
	GetConfigure() *string
	SetCycleDays(v []*int32) *ChangeCheckConfigShrinkRequest
	GetCycleDays() []*int32
	SetEnableAddCheck(v bool) *ChangeCheckConfigShrinkRequest
	GetEnableAddCheck() *bool
	SetEnableAutoCheck(v bool) *ChangeCheckConfigShrinkRequest
	GetEnableAutoCheck() *bool
	SetEndTime(v int32) *ChangeCheckConfigShrinkRequest
	GetEndTime() *int32
	SetRegionId(v string) *ChangeCheckConfigShrinkRequest
	GetRegionId() *string
	SetRemovedCheck(v []*ChangeCheckConfigShrinkRequestRemovedCheck) *ChangeCheckConfigShrinkRequest
	GetRemovedCheck() []*ChangeCheckConfigShrinkRequestRemovedCheck
	SetStandardIds(v []*int64) *ChangeCheckConfigShrinkRequest
	GetStandardIds() []*int64
	SetStartTime(v int32) *ChangeCheckConfigShrinkRequest
	GetStartTime() *int32
	SetSystemConfig(v bool) *ChangeCheckConfigShrinkRequest
	GetSystemConfig() *bool
	SetVendors(v []*string) *ChangeCheckConfigShrinkRequest
	GetVendors() []*string
}

type ChangeCheckConfigShrinkRequest struct {
	// The list of check items that you want to add to the policy.
	//
	// >  If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect.
	AddedCheck []*ChangeCheckConfigShrinkRequestAddedCheck `json:"AddedCheck,omitempty" xml:"AddedCheck,omitempty" type:"Repeated"`
	// The requirement IDs that you want to specify for the check policy.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the requirement ID. If the ConfigStandardIds parameter is configured, this parameter does not take effect.
	ConfigRequirementIdsShrink *string `json:"ConfigRequirementIds,omitempty" xml:"ConfigRequirementIds,omitempty"`
	// The standard IDs that you want to specify for the check policy.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the standard ID.
	ConfigStandardIdsShrink *string `json:"ConfigStandardIds,omitempty" xml:"ConfigStandardIds,omitempty"`
	// The configuration of the check item. Valid value:
	//
	// 	- **all**: Add all check items.
	//
	// example:
	//
	// all
	Configure *string `json:"Configure,omitempty" xml:"Configure,omitempty"`
	// The days in a week on which a check is performed.
	CycleDays []*int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty" type:"Repeated"`
	// Specifies whether to check the new check items in the selected requirement item. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	EnableAddCheck *bool `json:"EnableAddCheck,omitempty" xml:"EnableAddCheck,omitempty"`
	// Specifies whether to enable the automatic periodical check feature. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	EnableAutoCheck *bool `json:"EnableAutoCheck,omitempty" xml:"EnableAutoCheck,omitempty"`
	// The end time of the check. The value specifies a point in time in a day. The time period that is specified by the start time and end time must be one of the following time periods:
	//
	// 	- **00:00 to 06:00:*	- If you set the StartTime parameter to 0, you must set the EndTime parameter to 6.
	//
	// 	- **06:00 to 12:00**: If you set the StartTime parameter to 6, you must set the EndTime parameter to 12.
	//
	// 	- **12:00 to 18:00**: If you set the StartTime parameter to 12, you must set the EndTime parameter to 18.
	//
	// 	- **18:00 to 24:00:*	- If you set the StartTime parameter to 18, you must set the EndTime parameter to 24.
	//
	// example:
	//
	// 6
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID of the bastion host to query.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of the check items that you want to remove from the policy.
	//
	// >  If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect.
	RemovedCheck []*ChangeCheckConfigShrinkRequestRemovedCheck `json:"RemovedCheck,omitempty" xml:"RemovedCheck,omitempty" type:"Repeated"`
	// An array that consists of the information about the check item.
	StandardIds []*int64 `json:"StandardIds,omitempty" xml:"StandardIds,omitempty" type:"Repeated"`
	// The start time of the check. The value specifies a point in time in a day.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to use the configuration automatically generated by the system. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SystemConfig *bool `json:"SystemConfig,omitempty" xml:"SystemConfig,omitempty"`
	// The cloud service providers.
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s ChangeCheckConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigShrinkRequest) GetAddedCheck() []*ChangeCheckConfigShrinkRequestAddedCheck {
	return s.AddedCheck
}

func (s *ChangeCheckConfigShrinkRequest) GetConfigRequirementIdsShrink() *string {
	return s.ConfigRequirementIdsShrink
}

func (s *ChangeCheckConfigShrinkRequest) GetConfigStandardIdsShrink() *string {
	return s.ConfigStandardIdsShrink
}

func (s *ChangeCheckConfigShrinkRequest) GetConfigure() *string {
	return s.Configure
}

func (s *ChangeCheckConfigShrinkRequest) GetCycleDays() []*int32 {
	return s.CycleDays
}

func (s *ChangeCheckConfigShrinkRequest) GetEnableAddCheck() *bool {
	return s.EnableAddCheck
}

func (s *ChangeCheckConfigShrinkRequest) GetEnableAutoCheck() *bool {
	return s.EnableAutoCheck
}

func (s *ChangeCheckConfigShrinkRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ChangeCheckConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeCheckConfigShrinkRequest) GetRemovedCheck() []*ChangeCheckConfigShrinkRequestRemovedCheck {
	return s.RemovedCheck
}

func (s *ChangeCheckConfigShrinkRequest) GetStandardIds() []*int64 {
	return s.StandardIds
}

func (s *ChangeCheckConfigShrinkRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ChangeCheckConfigShrinkRequest) GetSystemConfig() *bool {
	return s.SystemConfig
}

func (s *ChangeCheckConfigShrinkRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *ChangeCheckConfigShrinkRequest) SetAddedCheck(v []*ChangeCheckConfigShrinkRequestAddedCheck) *ChangeCheckConfigShrinkRequest {
	s.AddedCheck = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetConfigRequirementIdsShrink(v string) *ChangeCheckConfigShrinkRequest {
	s.ConfigRequirementIdsShrink = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetConfigStandardIdsShrink(v string) *ChangeCheckConfigShrinkRequest {
	s.ConfigStandardIdsShrink = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetConfigure(v string) *ChangeCheckConfigShrinkRequest {
	s.Configure = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetCycleDays(v []*int32) *ChangeCheckConfigShrinkRequest {
	s.CycleDays = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetEnableAddCheck(v bool) *ChangeCheckConfigShrinkRequest {
	s.EnableAddCheck = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetEnableAutoCheck(v bool) *ChangeCheckConfigShrinkRequest {
	s.EnableAutoCheck = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetEndTime(v int32) *ChangeCheckConfigShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetRegionId(v string) *ChangeCheckConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetRemovedCheck(v []*ChangeCheckConfigShrinkRequestRemovedCheck) *ChangeCheckConfigShrinkRequest {
	s.RemovedCheck = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetStandardIds(v []*int64) *ChangeCheckConfigShrinkRequest {
	s.StandardIds = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetStartTime(v int32) *ChangeCheckConfigShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetSystemConfig(v bool) *ChangeCheckConfigShrinkRequest {
	s.SystemConfig = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetVendors(v []*string) *ChangeCheckConfigShrinkRequest {
	s.Vendors = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) Validate() error {
	if s.AddedCheck != nil {
		for _, item := range s.AddedCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemovedCheck != nil {
		for _, item := range s.RemovedCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeCheckConfigShrinkRequestAddedCheck struct {
	// The ID of the check item.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the ID of the check item.
	//
	// example:
	//
	// 5
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The section ID of the check item.
	//
	// example:
	//
	// 69
	SectionId *int64 `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
}

func (s ChangeCheckConfigShrinkRequestAddedCheck) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigShrinkRequestAddedCheck) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigShrinkRequestAddedCheck) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ChangeCheckConfigShrinkRequestAddedCheck) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ChangeCheckConfigShrinkRequestAddedCheck) SetCheckId(v int64) *ChangeCheckConfigShrinkRequestAddedCheck {
	s.CheckId = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequestAddedCheck) SetSectionId(v int64) *ChangeCheckConfigShrinkRequestAddedCheck {
	s.SectionId = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequestAddedCheck) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckConfigShrinkRequestRemovedCheck struct {
	// The ID of the check item.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the ID of the check item.
	//
	// example:
	//
	// 19
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The section ID of the check item.
	//
	// example:
	//
	// 69
	SectionId *int64 `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
}

func (s ChangeCheckConfigShrinkRequestRemovedCheck) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigShrinkRequestRemovedCheck) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigShrinkRequestRemovedCheck) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ChangeCheckConfigShrinkRequestRemovedCheck) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ChangeCheckConfigShrinkRequestRemovedCheck) SetCheckId(v int64) *ChangeCheckConfigShrinkRequestRemovedCheck {
	s.CheckId = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequestRemovedCheck) SetSectionId(v int64) *ChangeCheckConfigShrinkRequestRemovedCheck {
	s.SectionId = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequestRemovedCheck) Validate() error {
	return dara.Validate(s)
}
