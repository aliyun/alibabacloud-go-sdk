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
	SetClientToken(v string) *ChangeCheckConfigRequest
	GetClientToken() *string
	SetConfigId(v string) *ChangeCheckConfigRequest
	GetConfigId() *string
	SetConfigRequirementIds(v *ChangeCheckConfigRequestConfigRequirementIds) *ChangeCheckConfigRequest
	GetConfigRequirementIds() *ChangeCheckConfigRequestConfigRequirementIds
	SetConfigStandardIds(v *ChangeCheckConfigRequestConfigStandardIds) *ChangeCheckConfigRequest
	GetConfigStandardIds() *ChangeCheckConfigRequestConfigStandardIds
	SetConfigure(v string) *ChangeCheckConfigRequest
	GetConfigure() *string
	SetCycleDays(v []*int32) *ChangeCheckConfigRequest
	GetCycleDays() []*int32
	SetDaysOfMonth(v []*int32) *ChangeCheckConfigRequest
	GetDaysOfMonth() []*int32
	SetDryRun(v bool) *ChangeCheckConfigRequest
	GetDryRun() *bool
	SetEnableAddCheck(v bool) *ChangeCheckConfigRequest
	GetEnableAddCheck() *bool
	SetEnableAutoCheck(v bool) *ChangeCheckConfigRequest
	GetEnableAutoCheck() *bool
	SetEndTime(v int32) *ChangeCheckConfigRequest
	GetEndTime() *int32
	SetEndTimeSeconds(v int32) *ChangeCheckConfigRequest
	GetEndTimeSeconds() *int32
	SetFrequency(v string) *ChangeCheckConfigRequest
	GetFrequency() *string
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
	SetStartTimeSeconds(v int32) *ChangeCheckConfigRequest
	GetStartTimeSeconds() *int32
	SetStrategyName(v string) *ChangeCheckConfigRequest
	GetStrategyName() *string
	SetSystemConfig(v bool) *ChangeCheckConfigRequest
	GetSystemConfig() *bool
	SetVendors(v []*string) *ChangeCheckConfigRequest
	GetVendors() []*string
}

type ChangeCheckConfigRequest struct {
	// The list of check items to be added to the policy.
	//
	// <notice> If the ConfigStandardIds or ConfigRequirementIds parameter is configured, this parameter does not take effect. </notice>
	AddedCheck []*ChangeCheckConfigRequestAddedCheck `json:"AddedCheck,omitempty" xml:"AddedCheck,omitempty" type:"Repeated"`
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
	ConfigRequirementIds *ChangeCheckConfigRequestConfigRequirementIds `json:"ConfigRequirementIds,omitempty" xml:"ConfigRequirementIds,omitempty" type:"Struct"`
	// Configures the check policy by using the IDs of the specified standards.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the standard IDs.
	ConfigStandardIds *ChangeCheckConfigRequestConfigStandardIds `json:"ConfigStandardIds,omitempty" xml:"ConfigStandardIds,omitempty" type:"Struct"`
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
	RemovedCheck []*ChangeCheckConfigRequestRemovedCheck `json:"RemovedCheck,omitempty" xml:"RemovedCheck,omitempty" type:"Repeated"`
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

func (s ChangeCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigRequest) GetAddedCheck() []*ChangeCheckConfigRequestAddedCheck {
	return s.AddedCheck
}

func (s *ChangeCheckConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ChangeCheckConfigRequest) GetConfigId() *string {
	return s.ConfigId
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

func (s *ChangeCheckConfigRequest) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *ChangeCheckConfigRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *ChangeCheckConfigRequest) GetEndTimeSeconds() *int32 {
	return s.EndTimeSeconds
}

func (s *ChangeCheckConfigRequest) GetFrequency() *string {
	return s.Frequency
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

func (s *ChangeCheckConfigRequest) GetStartTimeSeconds() *int32 {
	return s.StartTimeSeconds
}

func (s *ChangeCheckConfigRequest) GetStrategyName() *string {
	return s.StrategyName
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

func (s *ChangeCheckConfigRequest) SetClientToken(v string) *ChangeCheckConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetConfigId(v string) *ChangeCheckConfigRequest {
	s.ConfigId = &v
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

func (s *ChangeCheckConfigRequest) SetDaysOfMonth(v []*int32) *ChangeCheckConfigRequest {
	s.DaysOfMonth = v
	return s
}

func (s *ChangeCheckConfigRequest) SetDryRun(v bool) *ChangeCheckConfigRequest {
	s.DryRun = &v
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

func (s *ChangeCheckConfigRequest) SetEndTimeSeconds(v int32) *ChangeCheckConfigRequest {
	s.EndTimeSeconds = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetFrequency(v string) *ChangeCheckConfigRequest {
	s.Frequency = &v
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

func (s *ChangeCheckConfigRequest) SetStartTimeSeconds(v int32) *ChangeCheckConfigRequest {
	s.StartTimeSeconds = &v
	return s
}

func (s *ChangeCheckConfigRequest) SetStrategyName(v string) *ChangeCheckConfigRequest {
	s.StrategyName = &v
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
	// The list of regulation IDs to be added to the policy.
	AddIds []*int64 `json:"AddIds,omitempty" xml:"AddIds,omitempty" type:"Repeated"`
	// The list of regulation IDs to be removed from the policy.
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
	// The list of standard IDs to be added to the policy.
	AddIds []*int64 `json:"AddIds,omitempty" xml:"AddIds,omitempty" type:"Repeated"`
	// The list of standard IDs to be removed from the policy.
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
