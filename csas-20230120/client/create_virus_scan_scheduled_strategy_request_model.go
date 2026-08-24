// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHighRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest
	GetHighRiskOperation() *string
	SetLowRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest
	GetLowRiskOperation() *string
	SetMatchMode(v string) *CreateVirusScanScheduledStrategyRequest
	GetMatchMode() *string
	SetMaxCpuUsage(v int64) *CreateVirusScanScheduledStrategyRequest
	GetMaxCpuUsage() *int64
	SetMidRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest
	GetMidRiskOperation() *string
	SetPerformanceMode(v string) *CreateVirusScanScheduledStrategyRequest
	GetPerformanceMode() *string
	SetPriority(v int32) *CreateVirusScanScheduledStrategyRequest
	GetPriority() *int32
	SetScanBeginTime(v int64) *CreateVirusScanScheduledStrategyRequest
	GetScanBeginTime() *int64
	SetScanEndTime(v int64) *CreateVirusScanScheduledStrategyRequest
	GetScanEndTime() *int64
	SetScanFrequency(v string) *CreateVirusScanScheduledStrategyRequest
	GetScanFrequency() *string
	SetScanInterval(v int64) *CreateVirusScanScheduledStrategyRequest
	GetScanInterval() *int64
	SetScanMode(v string) *CreateVirusScanScheduledStrategyRequest
	GetScanMode() *string
	SetScanPath(v []*string) *CreateVirusScanScheduledStrategyRequest
	GetScanPath() []*string
	SetScanTargets(v []*string) *CreateVirusScanScheduledStrategyRequest
	GetScanTargets() []*string
	SetStatus(v string) *CreateVirusScanScheduledStrategyRequest
	GetStatus() *string
	SetStrategyDescription(v string) *CreateVirusScanScheduledStrategyRequest
	GetStrategyDescription() *string
	SetStrategyName(v string) *CreateVirusScanScheduledStrategyRequest
	GetStrategyName() *string
	SetUserGroupIds(v []*string) *CreateVirusScanScheduledStrategyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateVirusScanScheduledStrategyRequest
	GetWhitelist() []*string
}

type CreateVirusScanScheduledStrategyRequest struct {
	// The action to take on high-risk virus files. Valid values:
	//
	// - **Quarantine**: quarantine quarantined file.
	//
	// - **Notify**: report an alert only without taking action on quarantined file.
	//
	// This parameter is required.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The action to take on low-risk virus files. Valid values:
	//
	// - **Quarantine**: quarantine quarantined file.
	//
	// - **Notify**: report an alert only without taking action on quarantined file.
	//
	// - **None**: take no action.
	//
	// This parameter is required.
	//
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// The matching method for the effective scope. Valid values:
	//
	// - **UserGroupAll**: the policy takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: the policy takes effect only for users in specified user groups. UserGroupIds is required when this value is specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The maximum percentage of terminal CPU usage during scanning. Valid values: 0 to 100. If this parameter is not specified or is set to 0, the default value based on PerformanceMode is used: 50 for SecurityFirst, 30 for Balance, and 15 for ExperienceFirst.
	//
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// The action to take on medium-risk virus files. Valid values:
	//
	// - **Quarantine**: quarantine quarantined file.
	//
	// - **Notify**: report an alert only without taking action on quarantined file.
	//
	// This parameter is required.
	//
	// example:
	//
	// Notify
	MidRiskOperation *string `json:"MidRiskOperation,omitempty" xml:"MidRiskOperation,omitempty"`
	// The scan performance mode. Valid values:
	//
	// - **SecurityFirst**: security first. The default CPU usage limit is 50%.
	//
	// - **Balance**: balanced. The default CPU usage limit is 30%.
	//
	// - **ExperienceFirst**: experience first. The default CPU usage limit is 15%.
	//
	// This parameter is required.
	//
	// example:
	//
	// Balance
	PerformanceMode *string `json:"PerformanceMode,omitempty" xml:"PerformanceMode,omitempty"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The start hour for triggering scans, specified as a whole hour. Valid values: 0 to 23 (inclusive). This field is not a timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ScanBeginTime *int64 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour for triggering scans, specified as a whole hour. Valid values: 1 to 24 (exclusive of the specified hour). The value must be greater than ScanBeginTime. Scan tasks generated by each trigger expire at this hour on the same day. This field is not a timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	ScanEndTime *int64 `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
	// The unit of the trigger cycle. Valid values:
	//
	// - **day**: by day.
	//
	// - **week**: by week.
	//
	// This parameter is required.
	//
	// example:
	//
	// week
	ScanFrequency *string `json:"ScanFrequency,omitempty" xml:"ScanFrequency,omitempty"`
	// The interval number of the trigger cycle. This parameter works together with ScanFrequency to determine the trigger cycle. Valid values: 1 to 30. For example, if ScanFrequency is set to week and ScanInterval is set to 1, the scan is triggered once a week.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ScanInterval *int64 `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
	// The path scope of the scan. Valid values:
	//
	// - **Quick**: quick scan. Only system critical directories and common risk locations are scanned.
	//
	// - **Full**: full disk scan.
	//
	// - **Custom**: custom path scan. ScanPath is required when this value is specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// Quick
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The collection of custom scan paths. This parameter is required when ScanMode is set to Custom and cannot be specified when ScanMode is set to Quick or Full. A maximum of 100 paths can be specified. Duplicate values are not allowed.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The collection of virus types to be handled in this scan. At least one type must be specified. Duplicate values are not allowed.
	//
	// This parameter is required.
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// The enabling status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// After the policy is enabled, it immediately participates in periodic scheduling. When the policy is disabled, it is only saved and does not trigger scans.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The policy description. The description can contain Chinese characters, uppercase and lowercase letters, digits, spaces, periods (.), commas (,), semicolons (;), forward slashes (/), at signs (@), hyphens (-), and underscores (_).
	//
	// example:
	//
	// Full disk scan for R&D department terminals every Sunday at midnight
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The policy name. The name can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// WeeklyScanForRDDept
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The collection of user group IDs for the effective scope. This parameter is required when MatchMode is set to UserGroupNormal and cannot be specified when MatchMode is set to UserGroupAll. At least 1 and at most 100 IDs can be specified. Duplicate values are not allowed.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempt users. Users in this list are excluded from the scan triggered by this policy. A maximum of 1000 users can be specified. Duplicate values are not allowed.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateVirusScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateVirusScanScheduledStrategyRequest) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *CreateVirusScanScheduledStrategyRequest) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *CreateVirusScanScheduledStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateVirusScanScheduledStrategyRequest) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *CreateVirusScanScheduledStrategyRequest) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *CreateVirusScanScheduledStrategyRequest) GetPerformanceMode() *string {
	return s.PerformanceMode
}

func (s *CreateVirusScanScheduledStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanBeginTime() *int64 {
	return s.ScanBeginTime
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanEndTime() *int64 {
	return s.ScanEndTime
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanInterval() *int64 {
	return s.ScanInterval
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanMode() *string {
	return s.ScanMode
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanPath() []*string {
	return s.ScanPath
}

func (s *CreateVirusScanScheduledStrategyRequest) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *CreateVirusScanScheduledStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateVirusScanScheduledStrategyRequest) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *CreateVirusScanScheduledStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateVirusScanScheduledStrategyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateVirusScanScheduledStrategyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateVirusScanScheduledStrategyRequest) SetHighRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest {
	s.HighRiskOperation = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetLowRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest {
	s.LowRiskOperation = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetMatchMode(v string) *CreateVirusScanScheduledStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetMaxCpuUsage(v int64) *CreateVirusScanScheduledStrategyRequest {
	s.MaxCpuUsage = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetMidRiskOperation(v string) *CreateVirusScanScheduledStrategyRequest {
	s.MidRiskOperation = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetPerformanceMode(v string) *CreateVirusScanScheduledStrategyRequest {
	s.PerformanceMode = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetPriority(v int32) *CreateVirusScanScheduledStrategyRequest {
	s.Priority = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanBeginTime(v int64) *CreateVirusScanScheduledStrategyRequest {
	s.ScanBeginTime = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanEndTime(v int64) *CreateVirusScanScheduledStrategyRequest {
	s.ScanEndTime = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanFrequency(v string) *CreateVirusScanScheduledStrategyRequest {
	s.ScanFrequency = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanInterval(v int64) *CreateVirusScanScheduledStrategyRequest {
	s.ScanInterval = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanMode(v string) *CreateVirusScanScheduledStrategyRequest {
	s.ScanMode = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanPath(v []*string) *CreateVirusScanScheduledStrategyRequest {
	s.ScanPath = v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetScanTargets(v []*string) *CreateVirusScanScheduledStrategyRequest {
	s.ScanTargets = v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetStatus(v string) *CreateVirusScanScheduledStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetStrategyDescription(v string) *CreateVirusScanScheduledStrategyRequest {
	s.StrategyDescription = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetStrategyName(v string) *CreateVirusScanScheduledStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetUserGroupIds(v []*string) *CreateVirusScanScheduledStrategyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) SetWhitelist(v []*string) *CreateVirusScanScheduledStrategyRequest {
	s.Whitelist = v
	return s
}

func (s *CreateVirusScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
