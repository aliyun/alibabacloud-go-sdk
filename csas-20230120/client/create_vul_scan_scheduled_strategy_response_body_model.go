// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateVulScanScheduledStrategyResponseBody
	GetCreateTime() *int64
	SetCustomMatchGroup(v []*CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) *CreateVulScanScheduledStrategyResponseBody
	GetCustomMatchGroup() []*CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup
	SetLastTriggerTime(v int64) *CreateVulScanScheduledStrategyResponseBody
	GetLastTriggerTime() *int64
	SetMatchMode(v string) *CreateVulScanScheduledStrategyResponseBody
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *CreateVulScanScheduledStrategyResponseBody
	GetMatchTargetIds() []*string
	SetPriority(v int32) *CreateVulScanScheduledStrategyResponseBody
	GetPriority() *int32
	SetRequestId(v string) *CreateVulScanScheduledStrategyResponseBody
	GetRequestId() *string
	SetScanBeginTime(v int32) *CreateVulScanScheduledStrategyResponseBody
	GetScanBeginTime() *int32
	SetScanEndTime(v int32) *CreateVulScanScheduledStrategyResponseBody
	GetScanEndTime() *int32
	SetScanFrequency(v string) *CreateVulScanScheduledStrategyResponseBody
	GetScanFrequency() *string
	SetScanInterval(v int32) *CreateVulScanScheduledStrategyResponseBody
	GetScanInterval() *int32
	SetStatus(v string) *CreateVulScanScheduledStrategyResponseBody
	GetStatus() *string
	SetStrategyDescription(v string) *CreateVulScanScheduledStrategyResponseBody
	GetStrategyDescription() *string
	SetStrategyId(v string) *CreateVulScanScheduledStrategyResponseBody
	GetStrategyId() *string
	SetStrategyName(v string) *CreateVulScanScheduledStrategyResponseBody
	GetStrategyName() *string
	SetWhitelist(v []*string) *CreateVulScanScheduledStrategyResponseBody
	GetWhitelist() []*string
}

type CreateVulScanScheduledStrategyResponseBody struct {
	// The time when the policy was created, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1786291200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure. An empty list is returned if the policy is not configured by organizational structure.
	CustomMatchGroup []*CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The time when the policy last triggered a scan, in seconds-level UNIX timestamp. The value 0 is returned if the policy has never been triggered.
	//
	// example:
	//
	// 1786291200
	LastTriggerTime *int64 `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: The policy takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: The policy takes effect only for users in specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The IDs of the user groups for which the policy takes effect. An empty list is returned when MatchMode is set to UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start hour during which the scan can be triggered. The value is an integer hour. Valid values: 0 to 23, inclusive. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *int32 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour during which the scan can be triggered. The value is an integer hour. Valid values: 1 to 24, exclusive of the specified hour. The value must be greater than ScanBeginTime. This field is not a timestamp.
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
	// The interval number of the trigger cycle, which determines the trigger cycle together with ScanFrequency. Valid values: 1 to 30. For example, if ScanFrequency is set to week and ScanInterval is set to 1, the scan is triggered once a week.
	//
	// example:
	//
	// 1
	ScanInterval *int32 `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
	// The enabling status. Valid values:
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
	// Weekly_Vulnerability_Scanning_RD_Dept
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The list of exempt users. Users in this list are excluded from the scan of this policy. An empty list is returned if no exempt users are configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateVulScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetCustomMatchGroup() []*CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetLastTriggerTime() *int64 {
	return s.LastTriggerTime
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateVulScanScheduledStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetCreateTime(v int64) *CreateVulScanScheduledStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetCustomMatchGroup(v []*CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) *CreateVulScanScheduledStrategyResponseBody {
	s.CustomMatchGroup = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetLastTriggerTime(v int64) *CreateVulScanScheduledStrategyResponseBody {
	s.LastTriggerTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetMatchMode(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetMatchTargetIds(v []*string) *CreateVulScanScheduledStrategyResponseBody {
	s.MatchTargetIds = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetPriority(v int32) *CreateVulScanScheduledStrategyResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetRequestId(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetScanBeginTime(v int32) *CreateVulScanScheduledStrategyResponseBody {
	s.ScanBeginTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetScanEndTime(v int32) *CreateVulScanScheduledStrategyResponseBody {
	s.ScanEndTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetScanFrequency(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.ScanFrequency = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetScanInterval(v int32) *CreateVulScanScheduledStrategyResponseBody {
	s.ScanInterval = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetStatus(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetStrategyDescription(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.StrategyDescription = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetStrategyId(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetStrategyName(v string) *CreateVulScanScheduledStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) SetWhitelist(v []*string) *CreateVulScanScheduledStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBody) Validate() error {
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

type CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The identity provider ID.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetGroup(v []*string) *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.Group = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetIdpId(v string) *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponseBodyCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
