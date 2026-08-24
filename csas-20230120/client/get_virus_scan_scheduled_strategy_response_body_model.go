// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetVirusScanScheduledStrategyResponseBody
	GetCreateTime() *string
	SetCustomMatchGroup(v []*GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) *GetVirusScanScheduledStrategyResponseBody
	GetCustomMatchGroup() []*GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup
	SetHighRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody
	GetHighRiskOperation() *string
	SetLastTriggerTime(v string) *GetVirusScanScheduledStrategyResponseBody
	GetLastTriggerTime() *string
	SetLowRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody
	GetLowRiskOperation() *string
	SetMatchMode(v string) *GetVirusScanScheduledStrategyResponseBody
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *GetVirusScanScheduledStrategyResponseBody
	GetMatchTargetIds() []*string
	SetMaxCpuUsage(v int64) *GetVirusScanScheduledStrategyResponseBody
	GetMaxCpuUsage() *int64
	SetMidRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody
	GetMidRiskOperation() *string
	SetPerformanceMode(v string) *GetVirusScanScheduledStrategyResponseBody
	GetPerformanceMode() *string
	SetPriority(v int32) *GetVirusScanScheduledStrategyResponseBody
	GetPriority() *int32
	SetRequestId(v string) *GetVirusScanScheduledStrategyResponseBody
	GetRequestId() *string
	SetScanBeginTime(v int32) *GetVirusScanScheduledStrategyResponseBody
	GetScanBeginTime() *int32
	SetScanEndTime(v int32) *GetVirusScanScheduledStrategyResponseBody
	GetScanEndTime() *int32
	SetScanFrequency(v string) *GetVirusScanScheduledStrategyResponseBody
	GetScanFrequency() *string
	SetScanInterval(v int32) *GetVirusScanScheduledStrategyResponseBody
	GetScanInterval() *int32
	SetScanMode(v string) *GetVirusScanScheduledStrategyResponseBody
	GetScanMode() *string
	SetScanPath(v []*string) *GetVirusScanScheduledStrategyResponseBody
	GetScanPath() []*string
	SetScanTargets(v []*string) *GetVirusScanScheduledStrategyResponseBody
	GetScanTargets() []*string
	SetStatus(v string) *GetVirusScanScheduledStrategyResponseBody
	GetStatus() *string
	SetStrategyDescription(v string) *GetVirusScanScheduledStrategyResponseBody
	GetStrategyDescription() *string
	SetStrategyId(v string) *GetVirusScanScheduledStrategyResponseBody
	GetStrategyId() *string
	SetStrategyName(v string) *GetVirusScanScheduledStrategyResponseBody
	GetStrategyName() *string
	SetWhitelist(v []*string) *GetVirusScanScheduledStrategyResponseBody
	GetWhitelist() []*string
}

type GetVirusScanScheduledStrategyResponseBody struct {
	// The time when the policy was created, in the format of yyyy-MM-dd HH:mm:ss in the UTC+8 time zone.
	//
	// example:
	//
	// 2026-08-21 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure.
	CustomMatchGroup []*GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The action to take on high-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Only report an alert without taking action on quarantined file.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The time when the policy last triggered a scan, in the format of yyyy-MM-dd HH:mm:ss in the UTC+8 time zone. An empty string is returned if the policy has never been triggered.
	//
	// example:
	//
	// 2026-08-21 01:00:03
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The action to take on low-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Only report an alert without taking action on quarantined file.
	//
	// - **None**: No action.
	//
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// The matching method for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of user group IDs within the effective scope. An empty list is returned when MatchMode is set to UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The maximum percentage of endpoint CPU usage allowed during the scan.
	//
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// The action to take on medium-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantine quarantined file.
	//
	// - **Notify**: Only report an alert without taking action on quarantined file.
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
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The interval number of the trigger cycle, which determines the trigger cycle together with ScanFrequency. For example, if ScanFrequency is set to week and ScanInterval is set to 1, the scan is triggered once a week.
	//
	// example:
	//
	// 1
	ScanInterval *int32 `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
	// The scan path scope. Valid values:
	//
	// - **Quick**: Quick scan. Only scans critical system directories and common risk locations.
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
	// Full disk scan for R&D department endpoints every Sunday at midnight
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
	// Weekly scan for R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The list of exempted users. Users in this list are excluded from the scan policy. An empty list is returned if no exemptions are configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s GetVirusScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetCustomMatchGroup() []*GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetLastTriggerTime() *string {
	return s.LastTriggerTime
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetPerformanceMode() *string {
	return s.PerformanceMode
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanMode() *string {
	return s.ScanMode
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanPath() []*string {
	return s.ScanPath
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *GetVirusScanScheduledStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetCreateTime(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetCustomMatchGroup(v []*GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) *GetVirusScanScheduledStrategyResponseBody {
	s.CustomMatchGroup = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetHighRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.HighRiskOperation = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetLastTriggerTime(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.LastTriggerTime = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetLowRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.LowRiskOperation = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetMatchMode(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetMatchTargetIds(v []*string) *GetVirusScanScheduledStrategyResponseBody {
	s.MatchTargetIds = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetMaxCpuUsage(v int64) *GetVirusScanScheduledStrategyResponseBody {
	s.MaxCpuUsage = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetMidRiskOperation(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.MidRiskOperation = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetPerformanceMode(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.PerformanceMode = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetPriority(v int32) *GetVirusScanScheduledStrategyResponseBody {
	s.Priority = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetRequestId(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanBeginTime(v int32) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanBeginTime = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanEndTime(v int32) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanEndTime = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanFrequency(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanFrequency = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanInterval(v int32) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanInterval = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanMode(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanMode = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanPath(v []*string) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanPath = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetScanTargets(v []*string) *GetVirusScanScheduledStrategyResponseBody {
	s.ScanTargets = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetStatus(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetStrategyDescription(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.StrategyDescription = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetStrategyId(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetStrategyName(v string) *GetVirusScanScheduledStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) SetWhitelist(v []*string) *GetVirusScanScheduledStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBody) Validate() error {
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

type GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The identity provider ID.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) SetGroup(v []*string) *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.Group = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) SetIdpId(v string) *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponseBodyCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
