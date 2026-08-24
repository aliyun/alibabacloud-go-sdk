// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanScheduledStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVirusScanScheduledStrategiesResponseBody
	GetRequestId() *string
	SetStrategies(v []*ListVirusScanScheduledStrategiesResponseBodyStrategies) *ListVirusScanScheduledStrategiesResponseBody
	GetStrategies() []*ListVirusScanScheduledStrategiesResponseBodyStrategies
	SetTotalNum(v int32) *ListVirusScanScheduledStrategiesResponseBody
	GetTotalNum() *int32
}

type ListVirusScanScheduledStrategiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scheduled virus scan policies.
	Strategies []*ListVirusScanScheduledStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	// The total number of scheduled virus scan policies.
	//
	// example:
	//
	// 100
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListVirusScanScheduledStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanScheduledStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanScheduledStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanScheduledStrategiesResponseBody) GetStrategies() []*ListVirusScanScheduledStrategiesResponseBodyStrategies {
	return s.Strategies
}

func (s *ListVirusScanScheduledStrategiesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListVirusScanScheduledStrategiesResponseBody) SetRequestId(v string) *ListVirusScanScheduledStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBody) SetStrategies(v []*ListVirusScanScheduledStrategiesResponseBodyStrategies) *ListVirusScanScheduledStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBody) SetTotalNum(v int32) *ListVirusScanScheduledStrategiesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBody) Validate() error {
	if s.Strategies != nil {
		for _, item := range s.Strategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanScheduledStrategiesResponseBodyStrategies struct {
	// The time when the policy was created, in the format of yyyy-MM-dd HH:mm:ss. The time is in UTC+8.
	//
	// example:
	//
	// 2026-08-21 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure.
	CustomMatchGroup []*ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The action to take on high-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Report an alert only without taking action on quarantined file.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The time when the policy last triggered a scan, in the format of yyyy-MM-dd HH:mm:ss. The time is in UTC+8. An empty string is returned if the policy has never been triggered.
	//
	// example:
	//
	// 2026-08-21 01:00:03
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The action to take on low-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Report an alert only without taking action on quarantined file.
	//
	// - **None**: Take no action.
	//
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of user group IDs to which the policy applies. An empty list is returned when MatchMode is set to UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The maximum percentage of terminal CPU usage allowed during scanning.
	//
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// The action to take on medium-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Report an alert only without taking action on quarantined file.
	//
	// example:
	//
	// Notify
	MidRiskOperation *string `json:"MidRiskOperation,omitempty" xml:"MidRiskOperation,omitempty"`
	// The scan performance schema pattern. Valid values:
	//
	// - **SecurityFirst**: Security first. The default CPU usage upper limit is 50%.
	//
	// - **Balance**: Balanced. The default CPU usage upper limit is 30%.
	//
	// - **ExperienceFirst**: Experience first. The default CPU usage upper limit is 15%.
	//
	// example:
	//
	// Balance
	PerformanceMode *string `json:"PerformanceMode,omitempty" xml:"PerformanceMode,omitempty"`
	// The policy priority. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The start hour during which the scan can be triggered. The value is a whole hour number ranging from 0 to 23, inclusive. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *int32 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour during which the scan can be triggered. The value is a whole hour number ranging from 1 to 24, exclusive of the specified hour, and must be greater than ScanBeginTime. The scan task generated by each trigger expires at this hour on the same day. This field is not a timestamp.
	//
	// example:
	//
	// 6
	ScanEndTime *int32 `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
	// The unit of the trigger cycle. Valid values:
	//
	// - **day**: By day.
	//
	// - **week**: By week.
	//
	// example:
	//
	// week
	ScanFrequency *string `json:"ScanFrequency,omitempty" xml:"ScanFrequency,omitempty"`
	// The interval number of the trigger cycle. This parameter works together with ScanFrequency to determine the trigger cycle. For example, if ScanFrequency is set to week and ScanInterval is set to 1, the scan is triggered once a week.
	//
	// example:
	//
	// 1
	ScanInterval *int32 `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
	// The scan path scope. Valid values:
	//
	// - **Quick**: Quick scan. Only critical system directories and common risk locations are scanned.
	//
	// - **Full**: Full disk scan.
	//
	// - **Custom**: Custom path scan.
	//
	// example:
	//
	// Quick
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The collection of custom scan paths.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The collection of virus types to be handled in this scan.
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// The enabled status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The policy description.
	//
	// example:
	//
	// Full disk scan on R&D department terminals every Sunday at midnight
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The ID of the scheduled virus scan policy.
	//
	// example:
	//
	// vc-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Weekly_Scan_DevTeam
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The list of exempted users. Users in this list are excluded from the scan performed by this policy. An empty list is returned if no exemptions are configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListVirusScanScheduledStrategiesResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanScheduledStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetCustomMatchGroup() []*ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetLastTriggerTime() *string {
	return s.LastTriggerTime
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetPerformanceMode() *string {
	return s.PerformanceMode
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanMode() *string {
	return s.ScanMode
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanPath() []*string {
	return s.ScanPath
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetStatus() *string {
	return s.Status
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetCreateTime(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.CreateTime = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetCustomMatchGroup(v []*ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.CustomMatchGroup = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetHighRiskOperation(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.HighRiskOperation = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetLastTriggerTime(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.LastTriggerTime = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetLowRiskOperation(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.LowRiskOperation = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetMatchMode(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.MatchMode = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetMatchTargetIds(v []*string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.MatchTargetIds = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetMaxCpuUsage(v int64) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.MaxCpuUsage = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetMidRiskOperation(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.MidRiskOperation = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetPerformanceMode(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.PerformanceMode = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetPriority(v int32) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.Priority = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanBeginTime(v int32) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanBeginTime = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanEndTime(v int32) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanEndTime = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanFrequency(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanFrequency = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanInterval(v int32) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanInterval = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanMode(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanMode = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanPath(v []*string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanPath = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetScanTargets(v []*string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.ScanTargets = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetStatus(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.Status = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetStrategyDescription(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyDescription = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetStrategyId(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetStrategyName(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) SetWhitelist(v []*string) *ListVirusScanScheduledStrategiesResponseBodyStrategies {
	s.Whitelist = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategies) Validate() error {
	if s.CustomMatchGroup != nil {
		for _, item := range s.CustomMatchGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The identity provider ID.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) SetGroup(v []*string) *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	s.Group = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) SetIdpId(v string) *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
