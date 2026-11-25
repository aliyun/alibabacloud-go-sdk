// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchMode(v string) *UpdateVulScanScheduledStrategyRequest
	GetMatchMode() *string
	SetPriority(v int32) *UpdateVulScanScheduledStrategyRequest
	GetPriority() *int32
	SetScanBeginTime(v string) *UpdateVulScanScheduledStrategyRequest
	GetScanBeginTime() *string
	SetScanEndTime(v string) *UpdateVulScanScheduledStrategyRequest
	GetScanEndTime() *string
	SetScanFrequency(v string) *UpdateVulScanScheduledStrategyRequest
	GetScanFrequency() *string
	SetScanInterval(v string) *UpdateVulScanScheduledStrategyRequest
	GetScanInterval() *string
	SetStatus(v string) *UpdateVulScanScheduledStrategyRequest
	GetStatus() *string
	SetStrategyDescription(v string) *UpdateVulScanScheduledStrategyRequest
	GetStrategyDescription() *string
	SetStrategyId(v string) *UpdateVulScanScheduledStrategyRequest
	GetStrategyId() *string
	SetStrategyName(v string) *UpdateVulScanScheduledStrategyRequest
	GetStrategyName() *string
	SetUserGroupIds(v []*string) *UpdateVulScanScheduledStrategyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateVulScanScheduledStrategyRequest
	GetWhitelist() []*string
}

type UpdateVulScanScheduledStrategyRequest struct {
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups. If this value is specified, UserGroupIds is required.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The start hour when a scan can be triggered. The value must be an integer on the hour. Valid values: 0 to 23, inclusive. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *string `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour when a scan can be triggered. The value must be an integer on the hour. Valid values: 1 to 24, exclusive. The value must be greater than the value of ScanBeginTime. This field is not a timestamp.
	//
	// example:
	//
	// 6
	ScanEndTime *string `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
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
	ScanInterval *string `json:"ScanInterval,omitempty" xml:"ScanInterval,omitempty"`
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
	// The ID of the scheduled vulnerability scan policy to be modified. You can obtain the value from the following operations:
	//
	// - [ListVulScanScheduledStrategies](~~ListVulScanScheduledStrategies~~): Queries scheduled vulnerability scan policies.
	//
	// - [CreateVulScanScheduledStrategy](~~CreateVulScanScheduledStrategy~~): Creates a scheduled vulnerability scan policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// Weekly vulnerability scanning for the R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The IDs of the user groups to which the policy applies. This parameter is required when MatchMode is set to UserGroupNormal and cannot be specified when MatchMode is set to UserGroupAll. You can specify 1 to 100 user group IDs, and the IDs cannot be duplicated. The specified IDs overwrite the existing user groups of the policy.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The user whitelist. Users in the whitelist are excluded from the scan of this policy. The whitelist can contain up to 1,000 users and cannot contain duplicates. This parameter overwrites the existing whitelist of the policy.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateVulScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateVulScanScheduledStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateVulScanScheduledStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateVulScanScheduledStrategyRequest) GetScanBeginTime() *string {
	return s.ScanBeginTime
}

func (s *UpdateVulScanScheduledStrategyRequest) GetScanEndTime() *string {
	return s.ScanEndTime
}

func (s *UpdateVulScanScheduledStrategyRequest) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *UpdateVulScanScheduledStrategyRequest) GetScanInterval() *string {
	return s.ScanInterval
}

func (s *UpdateVulScanScheduledStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateVulScanScheduledStrategyRequest) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *UpdateVulScanScheduledStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateVulScanScheduledStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateVulScanScheduledStrategyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateVulScanScheduledStrategyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateVulScanScheduledStrategyRequest) SetMatchMode(v string) *UpdateVulScanScheduledStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetPriority(v int32) *UpdateVulScanScheduledStrategyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetScanBeginTime(v string) *UpdateVulScanScheduledStrategyRequest {
	s.ScanBeginTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetScanEndTime(v string) *UpdateVulScanScheduledStrategyRequest {
	s.ScanEndTime = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetScanFrequency(v string) *UpdateVulScanScheduledStrategyRequest {
	s.ScanFrequency = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetScanInterval(v string) *UpdateVulScanScheduledStrategyRequest {
	s.ScanInterval = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetStatus(v string) *UpdateVulScanScheduledStrategyRequest {
	s.Status = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetStrategyDescription(v string) *UpdateVulScanScheduledStrategyRequest {
	s.StrategyDescription = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetStrategyId(v string) *UpdateVulScanScheduledStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetStrategyName(v string) *UpdateVulScanScheduledStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetUserGroupIds(v []*string) *UpdateVulScanScheduledStrategyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) SetWhitelist(v []*string) *UpdateVulScanScheduledStrategyRequest {
	s.Whitelist = v
	return s
}

func (s *UpdateVulScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
