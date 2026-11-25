// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *GetVulScanScheduledStrategyResponseBody
	GetCreateTime() *int64
	SetCustomMatchGroup(v []*GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) *GetVulScanScheduledStrategyResponseBody
	GetCustomMatchGroup() []*GetVulScanScheduledStrategyResponseBodyCustomMatchGroup
	SetLastTriggerTime(v int64) *GetVulScanScheduledStrategyResponseBody
	GetLastTriggerTime() *int64
	SetMatchMode(v string) *GetVulScanScheduledStrategyResponseBody
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *GetVulScanScheduledStrategyResponseBody
	GetMatchTargetIds() []*string
	SetPriority(v int32) *GetVulScanScheduledStrategyResponseBody
	GetPriority() *int32
	SetRequestId(v string) *GetVulScanScheduledStrategyResponseBody
	GetRequestId() *string
	SetScanBeginTime(v int32) *GetVulScanScheduledStrategyResponseBody
	GetScanBeginTime() *int32
	SetScanEndTime(v int32) *GetVulScanScheduledStrategyResponseBody
	GetScanEndTime() *int32
	SetScanFrequency(v string) *GetVulScanScheduledStrategyResponseBody
	GetScanFrequency() *string
	SetScanInterval(v int32) *GetVulScanScheduledStrategyResponseBody
	GetScanInterval() *int32
	SetStatus(v string) *GetVulScanScheduledStrategyResponseBody
	GetStatus() *string
	SetStrategyDescription(v string) *GetVulScanScheduledStrategyResponseBody
	GetStrategyDescription() *string
	SetStrategyId(v string) *GetVulScanScheduledStrategyResponseBody
	GetStrategyId() *string
	SetStrategyName(v string) *GetVulScanScheduledStrategyResponseBody
	GetStrategyName() *string
	SetWhitelist(v []*string) *GetVulScanScheduledStrategyResponseBody
	GetWhitelist() []*string
}

type GetVulScanScheduledStrategyResponseBody struct {
	// The time when the policy was created. The value is a UNIX timestamp accurate to seconds.
	//
	// example:
	//
	// 1786291200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure. If no organizational structure is configured, an empty list is returned.
	CustomMatchGroup []*GetVulScanScheduledStrategyResponseBodyCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The time when the policy last triggered a scan. The value is a UNIX timestamp accurate to seconds. If the policy has never been triggered, 0 is returned.
	//
	// example:
	//
	// 1786291200
	LastTriggerTime *int64 `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in the specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of effective user group IDs. If MatchMode is set to UserGroupAll, an empty list is returned.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: 1 to 100.
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
	// The start hour when a scan is allowed to be triggered. The value must be an integer on the hour. Valid values: 0 to 23. The specified hour is included. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *int32 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour when a scan is allowed to be triggered. The value must be an integer on the hour. Valid values: 1 to 24. The specified hour is excluded and must be greater than the value of ScanBeginTime. This field is not a timestamp.
	//
	// example:
	//
	// 6
	ScanEndTime *int32 `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
	// The unit of the trigger cycle. Valid values:
	//
	// - **day**: Day.
	//
	// - **week**: Week.
	//
	// example:
	//
	// week
	ScanFrequency *string `json:"ScanFrequency,omitempty" xml:"ScanFrequency,omitempty"`
	// The interval of the trigger cycle. This parameter and ScanFrequency determine the trigger cycle. Valid values: 1 to 30. For example, if ScanFrequency is set to week and ScanInterval is set to 1, a scan is performed once a week.
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
	// The description of the policy.
	//
	// example:
	//
	// Execute vulnerability scanning for R&D department endpoints every Sunday morning
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The ID of the scheduled vulnerability scanning policy.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Weekly vulnerability scanning for the R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The whitelist of users. Scans based on this policy are not performed on the users in the whitelist. If no whitelist is configured, an empty list is returned.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s GetVulScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulScanScheduledStrategyResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetVulScanScheduledStrategyResponseBody) GetCustomMatchGroup() []*GetVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *GetVulScanScheduledStrategyResponseBody) GetLastTriggerTime() *int64 {
	return s.LastTriggerTime
}

func (s *GetVulScanScheduledStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetVulScanScheduledStrategyResponseBody) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *GetVulScanScheduledStrategyResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetVulScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulScanScheduledStrategyResponseBody) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *GetVulScanScheduledStrategyResponseBody) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *GetVulScanScheduledStrategyResponseBody) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *GetVulScanScheduledStrategyResponseBody) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *GetVulScanScheduledStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVulScanScheduledStrategyResponseBody) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *GetVulScanScheduledStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetVulScanScheduledStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *GetVulScanScheduledStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *GetVulScanScheduledStrategyResponseBody) SetCreateTime(v int64) *GetVulScanScheduledStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetCustomMatchGroup(v []*GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) *GetVulScanScheduledStrategyResponseBody {
	s.CustomMatchGroup = v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetLastTriggerTime(v int64) *GetVulScanScheduledStrategyResponseBody {
	s.LastTriggerTime = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetMatchMode(v string) *GetVulScanScheduledStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetMatchTargetIds(v []*string) *GetVulScanScheduledStrategyResponseBody {
	s.MatchTargetIds = v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetPriority(v int32) *GetVulScanScheduledStrategyResponseBody {
	s.Priority = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetRequestId(v string) *GetVulScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetScanBeginTime(v int32) *GetVulScanScheduledStrategyResponseBody {
	s.ScanBeginTime = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetScanEndTime(v int32) *GetVulScanScheduledStrategyResponseBody {
	s.ScanEndTime = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetScanFrequency(v string) *GetVulScanScheduledStrategyResponseBody {
	s.ScanFrequency = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetScanInterval(v int32) *GetVulScanScheduledStrategyResponseBody {
	s.ScanInterval = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetStatus(v string) *GetVulScanScheduledStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetStrategyDescription(v string) *GetVulScanScheduledStrategyResponseBody {
	s.StrategyDescription = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetStrategyId(v string) *GetVulScanScheduledStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetStrategyName(v string) *GetVulScanScheduledStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) SetWhitelist(v []*string) *GetVulScanScheduledStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBody) Validate() error {
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

type GetVulScanScheduledStrategyResponseBodyCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The ID of the identity source.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetGroup(v []*string) *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.Group = v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetIdpId(v string) *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponseBodyCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
