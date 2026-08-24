// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateVirusScanTaskRequest
	GetEndTime() *int64
	SetHighRiskOperation(v string) *CreateVirusScanTaskRequest
	GetHighRiskOperation() *string
	SetLowRiskOperation(v string) *CreateVirusScanTaskRequest
	GetLowRiskOperation() *string
	SetMatchMode(v string) *CreateVirusScanTaskRequest
	GetMatchMode() *string
	SetMaxCpuUsage(v int64) *CreateVirusScanTaskRequest
	GetMaxCpuUsage() *int64
	SetMidRiskOperation(v string) *CreateVirusScanTaskRequest
	GetMidRiskOperation() *string
	SetPerformanceMode(v string) *CreateVirusScanTaskRequest
	GetPerformanceMode() *string
	SetScanMode(v string) *CreateVirusScanTaskRequest
	GetScanMode() *string
	SetScanPath(v []*string) *CreateVirusScanTaskRequest
	GetScanPath() []*string
	SetScanTargets(v []*string) *CreateVirusScanTaskRequest
	GetScanTargets() []*string
	SetTaskDescription(v string) *CreateVirusScanTaskRequest
	GetTaskDescription() *string
	SetUserGroupIds(v []*string) *CreateVirusScanTaskRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateVirusScanTaskRequest
	GetWhitelist() []*string
}

type CreateVirusScanTaskRequest struct {
	// The task expiration time, in seconds-level UNIX timestamp. After this time, endpoints no longer pull and execute this task. If this parameter is not specified or the specified time is earlier than the current time, the value defaults to the current time plus 24 hours.
	//
	// example:
	//
	// 1786377600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: applies only to users in specified user groups. UserGroupIds is required when this value is specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The maximum percentage of endpoint CPU usage during scanning. Valid values: 0 to 100. If this parameter is not specified or is set to 0, the default value is determined by PerformanceMode: 50 for SecurityFirst, 30 for Balance, and 15 for ExperienceFirst.
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
	// The scan performance pattern. Valid values:
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
	// The scan path scope. Valid values:
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
	// The task description. The description can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, spaces, periods (.), commas (,), semicolons (;), forward slashes (/), at signs (@), hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// Full scan for R&D department
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The collection of user group IDs to which the task applies. This parameter is required when MatchMode is set to UserGroupNormal and cannot be specified when MatchMode is set to UserGroupAll. At least 1 and at most 100 IDs can be specified. Duplicate values are not allowed.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempt users. Users in this list do not execute this scan task. A maximum of 1000 users can be specified. Duplicate values are not allowed.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateVirusScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateVirusScanTaskRequest) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *CreateVirusScanTaskRequest) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *CreateVirusScanTaskRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateVirusScanTaskRequest) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *CreateVirusScanTaskRequest) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *CreateVirusScanTaskRequest) GetPerformanceMode() *string {
	return s.PerformanceMode
}

func (s *CreateVirusScanTaskRequest) GetScanMode() *string {
	return s.ScanMode
}

func (s *CreateVirusScanTaskRequest) GetScanPath() []*string {
	return s.ScanPath
}

func (s *CreateVirusScanTaskRequest) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *CreateVirusScanTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *CreateVirusScanTaskRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateVirusScanTaskRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateVirusScanTaskRequest) SetEndTime(v int64) *CreateVirusScanTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetHighRiskOperation(v string) *CreateVirusScanTaskRequest {
	s.HighRiskOperation = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetLowRiskOperation(v string) *CreateVirusScanTaskRequest {
	s.LowRiskOperation = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetMatchMode(v string) *CreateVirusScanTaskRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetMaxCpuUsage(v int64) *CreateVirusScanTaskRequest {
	s.MaxCpuUsage = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetMidRiskOperation(v string) *CreateVirusScanTaskRequest {
	s.MidRiskOperation = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetPerformanceMode(v string) *CreateVirusScanTaskRequest {
	s.PerformanceMode = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetScanMode(v string) *CreateVirusScanTaskRequest {
	s.ScanMode = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetScanPath(v []*string) *CreateVirusScanTaskRequest {
	s.ScanPath = v
	return s
}

func (s *CreateVirusScanTaskRequest) SetScanTargets(v []*string) *CreateVirusScanTaskRequest {
	s.ScanTargets = v
	return s
}

func (s *CreateVirusScanTaskRequest) SetTaskDescription(v string) *CreateVirusScanTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetUserGroupIds(v []*string) *CreateVirusScanTaskRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateVirusScanTaskRequest) SetWhitelist(v []*string) *CreateVirusScanTaskRequest {
	s.Whitelist = v
	return s
}

func (s *CreateVirusScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
