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
	SetConfigId(v string) *ChangeCheckConfigShrinkRequest
	GetConfigId() *string
	SetConfigRequirementIdsShrink(v string) *ChangeCheckConfigShrinkRequest
	GetConfigRequirementIdsShrink() *string
	SetConfigStandardIdsShrink(v string) *ChangeCheckConfigShrinkRequest
	GetConfigStandardIdsShrink() *string
	SetConfigure(v string) *ChangeCheckConfigShrinkRequest
	GetConfigure() *string
	SetCycleDays(v []*int32) *ChangeCheckConfigShrinkRequest
	GetCycleDays() []*int32
	SetDaysOfMonth(v []*int32) *ChangeCheckConfigShrinkRequest
	GetDaysOfMonth() []*int32
	SetDryRun(v bool) *ChangeCheckConfigShrinkRequest
	GetDryRun() *bool
	SetEnableAddCheck(v bool) *ChangeCheckConfigShrinkRequest
	GetEnableAddCheck() *bool
	SetEnableAutoCheck(v bool) *ChangeCheckConfigShrinkRequest
	GetEnableAutoCheck() *bool
	SetEndTime(v int32) *ChangeCheckConfigShrinkRequest
	GetEndTime() *int32
	SetEndTimeSeconds(v int32) *ChangeCheckConfigShrinkRequest
	GetEndTimeSeconds() *int32
	SetFrequency(v string) *ChangeCheckConfigShrinkRequest
	GetFrequency() *string
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
	SetStartTimeSeconds(v int32) *ChangeCheckConfigShrinkRequest
	GetStartTimeSeconds() *int32
	SetStrategyName(v string) *ChangeCheckConfigShrinkRequest
	GetStrategyName() *string
	SetSystemConfig(v bool) *ChangeCheckConfigShrinkRequest
	GetSystemConfig() *bool
	SetVendors(v []*string) *ChangeCheckConfigShrinkRequest
	GetVendors() []*string
}

type ChangeCheckConfigShrinkRequest struct {
	// The list of check items to be added to the policy.
	//
	// <notice> If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect. </notice>
	AddedCheck []*ChangeCheckConfigShrinkRequestAddedCheck `json:"AddedCheck,omitempty" xml:"AddedCheck,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 22222f91-57af-4907-87f1-f9046da*****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configures the check policy by using the IDs of the specified regulations.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the regulation IDs. If the ConfigStandardIds parameter is configured, this parameter does not take effect.
	ConfigRequirementIdsShrink *string `json:"ConfigRequirementIds,omitempty" xml:"ConfigRequirementIds,omitempty"`
	// Configures the check policy by using the IDs of the specified standards.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the standard IDs.
	ConfigStandardIdsShrink *string `json:"ConfigStandardIds,omitempty" xml:"ConfigStandardIds,omitempty"`
	// The field configuration. Valid values:
	//
	// - **all**: Adds all check items.
	//
	// example:
	//
	// all
	Configure *string `json:"Configure,omitempty" xml:"Configure,omitempty"`
	// The periodic check time.
	CycleDays []*int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty" type:"Repeated"`
	// The list of days of the month when the policy cycle is monthly.
	DaysOfMonth []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - true: Only checks the request without performing the actual operation.
	//
	// - false: Executes the request normally.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to check new check items by default when new check items are added to the selected regulations. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableAddCheck *bool `json:"EnableAddCheck,omitempty" xml:"EnableAddCheck,omitempty"`
	// Specifies whether to enable automatic periodic checks. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableAutoCheck *bool `json:"EnableAutoCheck,omitempty" xml:"EnableAutoCheck,omitempty"`
	// The end time of the check time period, which indicates the hour of the day. The start time and end time must be selected from the following time intervals. Valid values: 6, 12, 18, and 24.
	//
	// - **0~6**: If the start time is 0, the end time must be set to 6 on the same day.
	//
	// - **6~12**: If the start time is 6, the end time must be set to 12 on the same day.
	//
	// - **12~18**: If the start time is 12, the end time must be set to 18 on the same day.
	//
	// - **18~24**: If the start time is 18, the end time must be set to 24 on the same day.
	//
	// example:
	//
	// 6
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end time of the check, in seconds from 00:00:00 of the current day. Valid values: 1 to 86400, where 86400 indicates 00:00:00 of the next day. This parameter is optional. If this parameter is not provided, the current value of the policy is retained. This parameter must be paired with StartTimeSeconds. If only one of them is provided, the system returns IllegalParam (paramName=scheduleTime). If both this parameter and StartTime/EndTime are provided, the second-level parameters take precedence, and the hour-level parameters are ignored. The time window span must not be less than 60 seconds. This parameter is not accepted for the system default policy (ConfigId=default) and will be rejected if provided.
	//
	// example:
	//
	// 32399
	EndTimeSeconds *int32 `json:"EndTimeSeconds,omitempty" xml:"EndTimeSeconds,omitempty"`
	// The frequency type. Valid values:
	//
	// - DAILY: Daily.
	//
	// - WEEKLY: Weekly.
	//
	// - MONTHLY: Monthly.
	//
	// - 6HOURLY: Every 6 hours.
	//
	// example:
	//
	// DAILY
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The region of the Security Center instance. Valid values:
	//
	// - **cn-hangzhou**: Center
	//
	// - **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of check items to be removed from the policy.
	//
	// <notice> If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect. </notice>
	RemovedCheck []*ChangeCheckConfigShrinkRequestRemovedCheck `json:"RemovedCheck,omitempty" xml:"RemovedCheck,omitempty" type:"Repeated"`
	// The ID of the member accounts in the resource folder, which is an Alibaba Cloud account ID.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// This parameter is deprecated and can be ignored.
	StandardIds []*int64 `json:"StandardIds,omitempty" xml:"StandardIds,omitempty" type:"Repeated"`
	// The start time of the check time period, which indicates the hour of the day. The start time and end time must be selected from the following time intervals. Valid values: 0, 6, 12, and 18.
	//
	// - **0~6**: If the start time is 0, the end time must be set to 6 on the same day.
	//
	// - **6~12**: If the start time is 6, the end time must be set to 12 on the same day.
	//
	// - **12~18**: If the start time is 12, the end time must be set to 18 on the same day.
	//
	// - **18~24**: If the start time is 18, the end time must be set to 24 on the same day.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start time of the check, in seconds from 00:00:00 of the current day. Valid values: 0 to 86399, which correspond to 00:00:00 to 23:59:59. This parameter is optional. If this parameter is not provided, the current value of the policy is retained. This parameter must be paired with EndTimeSeconds. If only one of them is provided, the system returns IllegalParam (paramName=scheduleTime). If both this parameter and StartTime/EndTime are provided, the second-level parameters take precedence, and the hour-level parameters are ignored. If the two sets of parameters are inconsistent, the server records a WARN log. The time window span must not be less than 60 seconds. The frequency of the system default policy (ConfigId=default) is always 6HOURLY. This parameter is not accepted for the default policy and will be rejected if provided. Under a multi-policy frequency, the actual scan execution time is exactly the second specified by this parameter.
	//
	// example:
	//
	// 25200
	StartTimeSeconds *int32 `json:"StartTimeSeconds,omitempty" xml:"StartTimeSeconds,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// Specifies whether to use the automatically generated system configuration. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	SystemConfig *bool `json:"SystemConfig,omitempty" xml:"SystemConfig,omitempty"`
	// The list of cloud providers.
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

func (s *ChangeCheckConfigShrinkRequest) GetConfigId() *string {
	return s.ConfigId
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

func (s *ChangeCheckConfigShrinkRequest) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *ChangeCheckConfigShrinkRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *ChangeCheckConfigShrinkRequest) GetEndTimeSeconds() *int32 {
	return s.EndTimeSeconds
}

func (s *ChangeCheckConfigShrinkRequest) GetFrequency() *string {
	return s.Frequency
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

func (s *ChangeCheckConfigShrinkRequest) GetStartTimeSeconds() *int32 {
	return s.StartTimeSeconds
}

func (s *ChangeCheckConfigShrinkRequest) GetStrategyName() *string {
	return s.StrategyName
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

func (s *ChangeCheckConfigShrinkRequest) SetConfigId(v string) *ChangeCheckConfigShrinkRequest {
	s.ConfigId = &v
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

func (s *ChangeCheckConfigShrinkRequest) SetDaysOfMonth(v []*int32) *ChangeCheckConfigShrinkRequest {
	s.DaysOfMonth = v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetDryRun(v bool) *ChangeCheckConfigShrinkRequest {
	s.DryRun = &v
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

func (s *ChangeCheckConfigShrinkRequest) SetEndTimeSeconds(v int32) *ChangeCheckConfigShrinkRequest {
	s.EndTimeSeconds = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetFrequency(v string) *ChangeCheckConfigShrinkRequest {
	s.Frequency = &v
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

func (s *ChangeCheckConfigShrinkRequest) SetStartTimeSeconds(v int32) *ChangeCheckConfigShrinkRequest {
	s.StartTimeSeconds = &v
	return s
}

func (s *ChangeCheckConfigShrinkRequest) SetStrategyName(v string) *ChangeCheckConfigShrinkRequest {
	s.StrategyName = &v
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
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the check item ID.
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
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the check item ID.
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
