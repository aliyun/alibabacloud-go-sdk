// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanScheduledStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *UpdateVulScanScheduledStrategyResponseBody
	GetCreateTime() *int64
	SetCustomMatchGroup(v []*UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) *UpdateVulScanScheduledStrategyResponseBody
	GetCustomMatchGroup() []*UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup
	SetLastTriggerTime(v int64) *UpdateVulScanScheduledStrategyResponseBody
	GetLastTriggerTime() *int64
	SetMatchMode(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *UpdateVulScanScheduledStrategyResponseBody
	GetMatchTargetIds() []*string
	SetPriority(v int32) *UpdateVulScanScheduledStrategyResponseBody
	GetPriority() *int32
	SetRequestId(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetRequestId() *string
	SetScanBeginTime(v int32) *UpdateVulScanScheduledStrategyResponseBody
	GetScanBeginTime() *int32
	SetScanEndTime(v int32) *UpdateVulScanScheduledStrategyResponseBody
	GetScanEndTime() *int32
	SetScanFrequency(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetScanFrequency() *string
	SetScanInterval(v int32) *UpdateVulScanScheduledStrategyResponseBody
	GetScanInterval() *int32
	SetStatus(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetStatus() *string
	SetStrategyDescription(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetStrategyDescription() *string
	SetStrategyId(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetStrategyId() *string
	SetStrategyName(v string) *UpdateVulScanScheduledStrategyResponseBody
	GetStrategyName() *string
	SetWhitelist(v []*string) *UpdateVulScanScheduledStrategyResponseBody
	GetWhitelist() []*string
}

type UpdateVulScanScheduledStrategyResponseBody struct {
	// The time when the policy was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1786291200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure. If no organizational structure is configured, an empty list is returned.
	CustomMatchGroup []*UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The time when the policy last triggered a scan. The value is a UNIX timestamp in seconds. If the policy has never triggered a scan, 0 is returned.
	//
	// example:
	//
	// 1786291200
	LastTriggerTime *int64 `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
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
	// The IDs of the user groups to which the policy applies. If MatchMode is set to UserGroupAll, an empty list is returned.
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
	// The start hour when a scan can be triggered. The value must be an integer on the hour. Valid values: 0 to 23, inclusive. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *int32 `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour when a scan can be triggered. The value must be an integer on the hour. Valid values: 1 to 24, exclusive. The value must be greater than the value of ScanBeginTime. This field is not a timestamp.
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
	// The policy description.
	//
	// example:
	//
	// Execute vulnerability scanning on R&D department endpoints every Sunday morning
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
	// Weekly vulnerability scanning for the R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The user whitelist. Users in the whitelist are excluded from the scan of this policy. If no whitelist is configured, an empty list is returned.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateVulScanScheduledStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanScheduledStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetCustomMatchGroup() []*UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetLastTriggerTime() *int64 {
	return s.LastTriggerTime
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetScanBeginTime() *int32 {
	return s.ScanBeginTime
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetScanEndTime() *int32 {
	return s.ScanEndTime
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetScanInterval() *int32 {
	return s.ScanInterval
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateVulScanScheduledStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetCreateTime(v int64) *UpdateVulScanScheduledStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetCustomMatchGroup(v []*UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) *UpdateVulScanScheduledStrategyResponseBody {
	s.CustomMatchGroup = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetLastTriggerTime(v int64) *UpdateVulScanScheduledStrategyResponseBody {
	s.LastTriggerTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetMatchMode(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetMatchTargetIds(v []*string) *UpdateVulScanScheduledStrategyResponseBody {
	s.MatchTargetIds = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetPriority(v int32) *UpdateVulScanScheduledStrategyResponseBody {
	s.Priority = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetRequestId(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetScanBeginTime(v int32) *UpdateVulScanScheduledStrategyResponseBody {
	s.ScanBeginTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetScanEndTime(v int32) *UpdateVulScanScheduledStrategyResponseBody {
	s.ScanEndTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetScanFrequency(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.ScanFrequency = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetScanInterval(v int32) *UpdateVulScanScheduledStrategyResponseBody {
	s.ScanInterval = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetStatus(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetStrategyDescription(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.StrategyDescription = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetStrategyId(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetStrategyName(v string) *UpdateVulScanScheduledStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) SetWhitelist(v []*string) *UpdateVulScanScheduledStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBody) Validate() error {
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

type UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup struct {
	// The organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The ID of the identity source.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetGroup(v []*string) *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.Group = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) SetIdpId(v string) *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponseBodyCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}
