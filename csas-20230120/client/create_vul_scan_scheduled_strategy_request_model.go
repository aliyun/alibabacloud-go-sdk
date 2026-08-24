// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchMode(v string) *CreateVulScanScheduledStrategyRequest
	GetMatchMode() *string
	SetPriority(v int32) *CreateVulScanScheduledStrategyRequest
	GetPriority() *int32
	SetScanBeginTime(v string) *CreateVulScanScheduledStrategyRequest
	GetScanBeginTime() *string
	SetScanEndTime(v string) *CreateVulScanScheduledStrategyRequest
	GetScanEndTime() *string
	SetScanFrequency(v string) *CreateVulScanScheduledStrategyRequest
	GetScanFrequency() *string
	SetScanInterval(v string) *CreateVulScanScheduledStrategyRequest
	GetScanInterval() *string
	SetStatus(v string) *CreateVulScanScheduledStrategyRequest
	GetStatus() *string
	SetStrategyDescription(v string) *CreateVulScanScheduledStrategyRequest
	GetStrategyDescription() *string
	SetStrategyName(v string) *CreateVulScanScheduledStrategyRequest
	GetStrategyName() *string
	SetUserGroupIds(v []*string) *CreateVulScanScheduledStrategyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateVulScanScheduledStrategyRequest
	GetWhitelist() []*string
}

type CreateVulScanScheduledStrategyRequest struct {
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: The policy takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: The policy takes effect only for users in specified user groups. In this case, UserGroupIds is required.
	//
	// This parameter is required.
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
	// The start hour during which the scan can be triggered. The value is an integer hour. Valid values: 0 to 23, inclusive. This field is not a timestamp.
	//
	// example:
	//
	// 1
	ScanBeginTime *string `json:"ScanBeginTime,omitempty" xml:"ScanBeginTime,omitempty"`
	// The end hour during which the scan can be triggered. The value is an integer hour. Valid values: 1 to 24, exclusive of the specified hour. The value must be greater than ScanBeginTime. This field is not a timestamp.
	//
	// example:
	//
	// 6
	ScanEndTime *string `json:"ScanEndTime,omitempty" xml:"ScanEndTime,omitempty"`
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
	// Execute vulnerability scanning on R&D department endpoints every Sunday at midnight
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The policy name. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Weekly_Vulnerability_Scanning_RD_Dept
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The IDs of the user groups for which the policy takes effect. This parameter is required when MatchMode is set to UserGroupNormal and must not be specified when MatchMode is set to UserGroupAll. The list must contain at least 1 and at most 100 entries. Duplicate entries are not allowed.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempt users. Users in this list are excluded from the scan of this policy. The list can contain up to 1000 entries. Duplicate entries are not allowed.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateVulScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateVulScanScheduledStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateVulScanScheduledStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVulScanScheduledStrategyRequest) GetScanBeginTime() *string {
	return s.ScanBeginTime
}

func (s *CreateVulScanScheduledStrategyRequest) GetScanEndTime() *string {
	return s.ScanEndTime
}

func (s *CreateVulScanScheduledStrategyRequest) GetScanFrequency() *string {
	return s.ScanFrequency
}

func (s *CreateVulScanScheduledStrategyRequest) GetScanInterval() *string {
	return s.ScanInterval
}

func (s *CreateVulScanScheduledStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateVulScanScheduledStrategyRequest) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *CreateVulScanScheduledStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateVulScanScheduledStrategyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateVulScanScheduledStrategyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateVulScanScheduledStrategyRequest) SetMatchMode(v string) *CreateVulScanScheduledStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetPriority(v int32) *CreateVulScanScheduledStrategyRequest {
	s.Priority = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetScanBeginTime(v string) *CreateVulScanScheduledStrategyRequest {
	s.ScanBeginTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetScanEndTime(v string) *CreateVulScanScheduledStrategyRequest {
	s.ScanEndTime = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetScanFrequency(v string) *CreateVulScanScheduledStrategyRequest {
	s.ScanFrequency = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetScanInterval(v string) *CreateVulScanScheduledStrategyRequest {
	s.ScanInterval = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetStatus(v string) *CreateVulScanScheduledStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetStrategyDescription(v string) *CreateVulScanScheduledStrategyRequest {
	s.StrategyDescription = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetStrategyName(v string) *CreateVulScanScheduledStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetUserGroupIds(v []*string) *CreateVulScanScheduledStrategyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) SetWhitelist(v []*string) *CreateVulScanScheduledStrategyRequest {
	s.Whitelist = v
	return s
}

func (s *CreateVulScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
