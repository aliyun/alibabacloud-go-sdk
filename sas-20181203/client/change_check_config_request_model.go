// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedCheck(v []*ChangeCheckConfigRequestAddedCheck) *ChangeCheckConfigRequest
	GetAddedCheck() []*ChangeCheckConfigRequestAddedCheck
	SetConfigRequirementIds(v *ChangeCheckConfigRequestConfigRequirementIds) *ChangeCheckConfigRequest
	GetConfigRequirementIds() *ChangeCheckConfigRequestConfigRequirementIds
	SetConfigStandardIds(v *ChangeCheckConfigRequestConfigStandardIds) *ChangeCheckConfigRequest
	GetConfigStandardIds() *ChangeCheckConfigRequestConfigStandardIds
	SetConfigure(v string) *ChangeCheckConfigRequest
	GetConfigure() *string
	SetCycleDays(v []*int32) *ChangeCheckConfigRequest
	GetCycleDays() []*int32
	SetEnableAddCheck(v bool) *ChangeCheckConfigRequest
	GetEnableAddCheck() *bool
	SetEnableAutoCheck(v bool) *ChangeCheckConfigRequest
	GetEnableAutoCheck() *bool
	SetEndTime(v int32) *ChangeCheckConfigRequest
	GetEndTime() *int32
	SetRegionId(v string) *ChangeCheckConfigRequest
	GetRegionId() *string
	SetRemovedCheck(v []*ChangeCheckConfigRequestRemovedCheck) *ChangeCheckConfigRequest
	GetRemovedCheck() []*ChangeCheckConfigRequestRemovedCheck
	SetResourceDirectoryAccountId(v int64) *ChangeCheckConfigRequest
	GetResourceDirectoryAccountId() *int64
	SetStandardIds(v []*int64) *ChangeCheckConfigRequest
	GetStandardIds() []*int64
	SetStartTime(v int32) *ChangeCheckConfigRequest
	GetStartTime() *int32
	SetSystemConfig(v bool) *ChangeCheckConfigRequest
	GetSystemConfig() *bool
	SetVendors(v []*string) *ChangeCheckConfigRequest
	GetVendors() []*string
}

type ChangeCheckConfigRequest struct {
	// The list of check items that you want to add to the policy.
	//
	// >  If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect.
	AddedCheck []*ChangeCheckConfigRequestAddedCheck `json:"AddedCheck,omitempty" xml:"AddedCheck,omitempty" type:"Repeated"`
	// The requirement IDs that you want to specify for the check policy.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the requirement ID. If the ConfigStandardIds parameter is configured, this parameter does not take effect.
	ConfigRequirementIds *ChangeCheckConfigRequestConfigRequirementIds `json:"ConfigRequirementIds,omitempty" xml:"ConfigRequirementIds,omitempty" type:"Struct"`
	// The standard IDs that you want to specify for the check policy.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the standard ID.
	ConfigStandardIds *ChangeCheckConfigRequestConfigStandardIds `json:"ConfigStandardIds,omitempty" xml:"ConfigStandardIds,omitempty" type:"Struct"`
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
	RemovedCheck []*ChangeCheckConfigRequestRemovedCheck `json:"RemovedCheck,omitempty" xml:"RemovedCheck,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
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

func (s ChangeCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequest) GetAddedCheck() []*ChangeCheckConfigRequestAddedCheck {
	return s.AddedCheck
}

func (s *ChangeCheckConfigRequest) GetConfigRequirementIds() *ChangeCheckConfigRequestConfigRequirementIds {
	return s.ConfigRequirementIds
}

func (s *ChangeCheckConfigRequest) GetConfigStandardIds() *ChangeCheckConfigRequestConfigStandardIds {
	return s.ConfigStandardIds
}

func (s *ChangeCheckConfigRequest) GetConfigure() *string {
	return s.Configure
}

func (s *ChangeCheckConfigRequest) GetCycleDays() []*int32 {
	return s.CycleDays
}

func (s *ChangeCheckConfigRequest) GetEnableAddCheck() *bool {
	return s.EnableAddCheck
}

func (s *ChangeCheckConfigRequest) GetEnableAutoCheck() *bool {
	return s.EnableAutoCheck
}

func (s *ChangeCheckConfigRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ChangeCheckConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeCheckConfigRequest) GetRemovedCheck() []*ChangeCheckConfigRequestRemovedCheck {
	return s.RemovedCheck
}

func (s *ChangeCheckConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ChangeCheckConfigRequest) GetStandardIds() []*int64 {
	return s.StandardIds
}

func (s *ChangeCheckConfigRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ChangeCheckConfigRequest) GetSystemConfig() *bool {
	return s.SystemConfig
}

func (s *ChangeCheckConfigRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *ChangeCheckConfigRequest) SetAddedCheck(v []*ChangeCheckConfigRequestAddedCheck) *ChangeCheckConfigRequest {
	s.AddedCheck = v
	return s
}

func (s *ChangeCheckConfigRequest) SetConfigRequirementIds(v *ChangeCheckConfigRequestConfigRequirementIds) *ChangeCheckConfigRequest {
	s.ConfigRequirementIds = v
	return s
}

func (s *ChangeCheckConfigRequest) SetConfigStandardIds(v *ChangeCheckConfigRequestConfigStandardIds) *ChangeCheckConfigRequest {
	s.ConfigStandardIds = v
	return s
}

func (s *ChangeCheckConfigRequest) SetConfigure(v string) *ChangeCheckConfigRequest {
	s.Configure = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetCycleDays(v []*int32) *ChangeCheckConfigRequest {
	s.CycleDays = v
	return s
}

func (s *ChangeCheckConfigRequest) SetEnableAddCheck(v bool) *ChangeCheckConfigRequest {
	s.EnableAddCheck = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetEnableAutoCheck(v bool) *ChangeCheckConfigRequest {
	s.EnableAutoCheck = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetEndTime(v int32) *ChangeCheckConfigRequest {
	s.EndTime = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetRegionId(v string) *ChangeCheckConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetRemovedCheck(v []*ChangeCheckConfigRequestRemovedCheck) *ChangeCheckConfigRequest {
	s.RemovedCheck = v
	return s
}

func (s *ChangeCheckConfigRequest) SetResourceDirectoryAccountId(v int64) *ChangeCheckConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetStandardIds(v []*int64) *ChangeCheckConfigRequest {
	s.StandardIds = v
	return s
}

func (s *ChangeCheckConfigRequest) SetStartTime(v int32) *ChangeCheckConfigRequest {
	s.StartTime = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetSystemConfig(v bool) *ChangeCheckConfigRequest {
	s.SystemConfig = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetVendors(v []*string) *ChangeCheckConfigRequest {
	s.Vendors = v
	return s
}

func (s *ChangeCheckConfigRequest) Validate() error {
	if s.AddedCheck != nil {
		for _, item := range s.AddedCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigRequirementIds != nil {
		if err := s.ConfigRequirementIds.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigStandardIds != nil {
		if err := s.ConfigStandardIds.Validate(); err != nil {
			return err
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

type ChangeCheckConfigRequestAddedCheck struct {
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

func (s ChangeCheckConfigRequestAddedCheck) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequestAddedCheck) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequestAddedCheck) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ChangeCheckConfigRequestAddedCheck) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ChangeCheckConfigRequestAddedCheck) SetCheckId(v int64) *ChangeCheckConfigRequestAddedCheck {
	s.CheckId = &v
	return s
}

func (s *ChangeCheckConfigRequestAddedCheck) SetSectionId(v int64) *ChangeCheckConfigRequestAddedCheck {
	s.SectionId = &v
	return s
}

func (s *ChangeCheckConfigRequestAddedCheck) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckConfigRequestConfigRequirementIds struct {
	// The requirement IDs that you want to add to the policy.
	AddIds []*int64 `json:"AddIds,omitempty" xml:"AddIds,omitempty" type:"Repeated"`
	// The requirement IDs that you want to remove from the policy.
	RemoveIds []*int64 `json:"RemoveIds,omitempty" xml:"RemoveIds,omitempty" type:"Repeated"`
}

func (s ChangeCheckConfigRequestConfigRequirementIds) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequestConfigRequirementIds) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequestConfigRequirementIds) GetAddIds() []*int64 {
	return s.AddIds
}

func (s *ChangeCheckConfigRequestConfigRequirementIds) GetRemoveIds() []*int64 {
	return s.RemoveIds
}

func (s *ChangeCheckConfigRequestConfigRequirementIds) SetAddIds(v []*int64) *ChangeCheckConfigRequestConfigRequirementIds {
	s.AddIds = v
	return s
}

func (s *ChangeCheckConfigRequestConfigRequirementIds) SetRemoveIds(v []*int64) *ChangeCheckConfigRequestConfigRequirementIds {
	s.RemoveIds = v
	return s
}

func (s *ChangeCheckConfigRequestConfigRequirementIds) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckConfigRequestConfigStandardIds struct {
	// The standard IDs that you want to add to the policy.
	AddIds []*int64 `json:"AddIds,omitempty" xml:"AddIds,omitempty" type:"Repeated"`
	// The standard IDs that you want to remove from the policy.
	RemoveIds []*int64 `json:"RemoveIds,omitempty" xml:"RemoveIds,omitempty" type:"Repeated"`
}

func (s ChangeCheckConfigRequestConfigStandardIds) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequestConfigStandardIds) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequestConfigStandardIds) GetAddIds() []*int64 {
	return s.AddIds
}

func (s *ChangeCheckConfigRequestConfigStandardIds) GetRemoveIds() []*int64 {
	return s.RemoveIds
}

func (s *ChangeCheckConfigRequestConfigStandardIds) SetAddIds(v []*int64) *ChangeCheckConfigRequestConfigStandardIds {
	s.AddIds = v
	return s
}

func (s *ChangeCheckConfigRequestConfigStandardIds) SetRemoveIds(v []*int64) *ChangeCheckConfigRequestConfigStandardIds {
	s.RemoveIds = v
	return s
}

func (s *ChangeCheckConfigRequestConfigStandardIds) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckConfigRequestRemovedCheck struct {
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

func (s ChangeCheckConfigRequestRemovedCheck) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequestRemovedCheck) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequestRemovedCheck) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ChangeCheckConfigRequestRemovedCheck) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ChangeCheckConfigRequestRemovedCheck) SetCheckId(v int64) *ChangeCheckConfigRequestRemovedCheck {
	s.CheckId = &v
	return s
}

func (s *ChangeCheckConfigRequestRemovedCheck) SetSectionId(v int64) *ChangeCheckConfigRequestRemovedCheck {
	s.SectionId = &v
	return s
}

func (s *ChangeCheckConfigRequestRemovedCheck) Validate() error {
	return dara.Validate(s)
}
