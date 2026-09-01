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
	SetClientToken(v string) *ChangeCheckConfigShrinkRequest
	GetClientToken() *string
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
	SetResourceDirectoryAccountId(v int64) *ChangeCheckConfigShrinkRequest
	GetResourceDirectoryAccountId() *int64
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
	// The list of check items to add to the policy.
	//
	// <notice> If ConfigStandardIds or ConfigRequirementIds is specified, this parameter does not take effect.
	AddedCheck []*ChangeCheckConfigShrinkRequestAddedCheck `json:"AddedCheck,omitempty" xml:"AddedCheck,omitempty" type:"Repeated"`
	// The client token used to ensure request idempotency. Use a different token for each request. Only ASCII characters are supported. The token can be up to 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Configures the check policy by specifying requirement IDs.
	//
	// > Call [ListCheckResult](~~ListCheckResult~~) to obtain requirement IDs. If ConfigStandardIds is specified, this parameter does not take effect.
	ConfigRequirementIdsShrink *string `json:"ConfigRequirementIds,omitempty" xml:"ConfigRequirementIds,omitempty"`
	// Configures the check policy by specifying standard IDs.
	//
	// > Call [ListCheckResult](~~ListCheckResult~~) to obtain standard IDs.
	ConfigStandardIdsShrink *string `json:"ConfigStandardIds,omitempty" xml:"ConfigStandardIds,omitempty"`
	// The field configuration. Valid values:
	//
	// - **all:*	- Adds all check items.
	//
	// example:
	//
	// all
	Configure *string `json:"Configure,omitempty" xml:"Configure,omitempty"`
	// The scheduled check days.
	CycleDays []*int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty" type:"Repeated"`
	// Specifies whether to automatically include newly added check items from the selected requirements. Valid values:
	//
	// - **true:*	- Enabled.
	//
	// - **false:*	- Disabled.
	//
	// example:
	//
	// false
	EnableAddCheck *bool `json:"EnableAddCheck,omitempty" xml:"EnableAddCheck,omitempty"`
	// Specifies whether to enable automatic scheduled checks. Valid values:
	//
	// - **true:*	- Enabled.
	//
	// - **false:*	- Disabled.
	//
	// example:
	//
	// true
	EnableAutoCheck *bool `json:"EnableAutoCheck,omitempty" xml:"EnableAutoCheck,omitempty"`
	// The end hour of the check time window, expressed as an hour of the day. The start and end times must fall within one of the following time ranges. Valid values: 6, 12, 18, 24.
	//
	// - **0~6:*	- If the start time is 0, set the end time to 6.
	//
	// - **6~12:*	- If the start time is 6, set the end time to 12.
	//
	// - **12~18:*	- If the start time is 12, set the end time to 18.
	//
	// - **18~24:*	- If the start time is 18, set the end time to 24.
	//
	// example:
	//
	// 6
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region of the Security Center instance. Valid values:
	//
	// - **cn-hangzhou:*	- China (Hangzhou)
	//
	// - **ap-southeast-1:*	- Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of check items to remove from the policy.
	//
	// <notice> If ConfigStandardIds or ConfigRequirementIds is specified, this parameter does not take effect.
	RemovedCheck []*ChangeCheckConfigShrinkRequestRemovedCheck `json:"RemovedCheck,omitempty" xml:"RemovedCheck,omitempty" type:"Repeated"`
	// The ID of the resource directory member accounts (Alibaba Cloud account).
	//
	// > Call [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// This parameter is deprecated. You do not need to configure it.
	StandardIds []*int64 `json:"StandardIds,omitempty" xml:"StandardIds,omitempty" type:"Repeated"`
	// The start hour of the check time window, expressed as an hour of the day. The start and end times must fall within one of the following time ranges. Valid values: 0, 6, 12, 18.
	//
	// - **0~6:*	- If the start time is 0, set the end time to 6.
	//
	// - **6~12:*	- If the start time is 6, set the end time to 12.
	//
	// - **12~18:*	- If the start time is 12, set the end time to 18.
	//
	// - **18~24:*	- If the start time is 18, set the end time to 24.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to use the system-generated configuration. Valid values:
	//
	// - **true:*	- Yes.
	//
	// - **false:*	- No.
	//
	// example:
	//
	// true
	SystemConfig *bool `json:"SystemConfig,omitempty" xml:"SystemConfig,omitempty"`
	// The list of cloud vendors.
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

func (s *ChangeCheckConfigShrinkRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *ChangeCheckConfigShrinkRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *ChangeCheckConfigShrinkRequest) SetClientToken(v string) *ChangeCheckConfigShrinkRequest {
	s.ClientToken = &v
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

func (s *ChangeCheckConfigShrinkRequest) SetResourceDirectoryAccountId(v int64) *ChangeCheckConfigShrinkRequest {
	s.ResourceDirectoryAccountId = &v
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
	// > Call [ListCheckResult](~~ListCheckResult~~) to obtain check item IDs.
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
	// > Call [ListCheckResult](~~ListCheckResult~~) to obtain check item IDs.
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
