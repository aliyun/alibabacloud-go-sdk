// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanScheduledStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVulScanScheduledStrategiesResponseBody
	GetRequestId() *string
	SetStrategies(v []*ListVulScanScheduledStrategiesResponseBodyStrategies) *ListVulScanScheduledStrategiesResponseBody
	GetStrategies() []*ListVulScanScheduledStrategiesResponseBodyStrategies
	SetTotalNum(v int64) *ListVulScanScheduledStrategiesResponseBody
	GetTotalNum() *int64
}

type ListVulScanScheduledStrategiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scheduled vulnerability scan policies.
	Strategies []*ListVulScanScheduledStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	// The total number of scheduled vulnerability scan policies that match the query conditions.
	//
	// example:
	//
	// 37
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListVulScanScheduledStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanScheduledStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVulScanScheduledStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVulScanScheduledStrategiesResponseBody) GetStrategies() []*ListVulScanScheduledStrategiesResponseBodyStrategies {
	return s.Strategies
}

func (s *ListVulScanScheduledStrategiesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListVulScanScheduledStrategiesResponseBody) SetRequestId(v string) *ListVulScanScheduledStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBody) SetStrategies(v []*ListVulScanScheduledStrategiesResponseBodyStrategies) *ListVulScanScheduledStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBody) SetTotalNum(v int64) *ListVulScanScheduledStrategiesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBody) Validate() error {
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

type ListVulScanScheduledStrategiesResponseBodyStrategies struct {
	// The time when the policy was created, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1786291200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure. An empty list is returned if the scope is not configured by organizational structure.
	CustomMatchGroup []*ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The time when the policy last triggered a scan, in seconds-level UNIX timestamp. The value 0 is returned if the policy has never been triggered.
	//
	// example:
	//
	// 1786291200
	LastTriggerTime *int64 `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The matching mode of the effective scope. Valid values:
	//
	// - **UserGroupAll**: Takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Takes effect only for users in specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of user group IDs that the policy takes effect for. An empty list is returned when MatchMode is set to UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The start hour during which scans can be triggered. The value is an integer representing the hour of the day. Valid values: 0 to 23 (inclusive). This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *int32 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour during which scans can be triggered. The value is an integer representing the hour of the day. Valid values: 1 to 24 (exclusive). The value must be greater than ScanBeginTime. This field is not a timestamp.
	//
	// example:
	//
	// 6
	ScanEndTime *int32 `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
	// The unit of the trigger cycle. Valid values:
	//
	// - **day**: by day.
	//
	// - **week**: by week.
	//
	// example:
	//
	// week
	ScanFrequency *string `json:"ScanFrequency,omitempty" xml:"ScanFrequency,omitempty"`
	// The interval number of the trigger cycle. This parameter works together with ScanFrequency to determine the trigger cycle. Valid values: 1 to 30. For example, if ScanFrequency is set to week and ScanInterval is set to 1, the scan is triggered once a week.
	//
	// example:
	//
	// 1
	ScanInterval *int32 `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
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
	// Execute vulnerability scanning on R&D department endpoints every Sunday at midnight
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The ID of the scheduled vulnerability scan policy.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Weekly vulnerability scanning for R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The list of exempt usernames. Users in this list are excluded from the scan of this policy. An empty list is returned if no exemptions are configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListVulScanScheduledStrategiesResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanScheduledStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetCustomMatchGroup() []*ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetLastTriggerTime() *int64 {
	return s.LastTriggerTime
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetStatus() *string {
	return s.Status
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetCreateTime(v int64) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.CreateTime = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetCustomMatchGroup(v []*ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.CustomMatchGroup = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetLastTriggerTime(v int64) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.LastTriggerTime = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetMatchMode(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.MatchMode = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetMatchTargetIds(v []*string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.MatchTargetIds = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetPriority(v int32) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.Priority = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetScanBeginTime(v int32) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.ScanBeginTime = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetScanEndTime(v int32) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.ScanEndTime = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetScanFrequency(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.ScanFrequency = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetScanInterval(v int32) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.ScanInterval = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetStatus(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.Status = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetStrategyDescription(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyDescription = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetStrategyId(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetStrategyName(v string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) SetWhitelist(v []*string) *ListVulScanScheduledStrategiesResponseBodyStrategies {
	s.Whitelist = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategies) Validate() error {
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

type ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The identity provider ID.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) SetGroup(v []*string) *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	s.Group = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) SetIdpId(v string) *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponseBodyStrategiesCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
