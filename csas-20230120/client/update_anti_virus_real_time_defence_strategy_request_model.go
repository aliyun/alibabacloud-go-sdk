// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntiVirusRealTimeDefenceStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHighRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetHighRiskOperation() *string
	SetLowRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetLowRiskOperation() *string
	SetMatchMode(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetMatchMode() *string
	SetMaxCpuUsage(v int64) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetMaxCpuUsage() *int64
	SetMidRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetMidRiskOperation() *string
	SetScanTargets(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetScanTargets() []*string
	SetStatus(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetStatus() *string
	SetUserGroupIds(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest
	GetWhitelist() []*string
}

type UpdateAntiVirusRealTimeDefenceStrategyRequest struct {
	// The action to take on high-risk virus files. Required when configuring the real-time defense policy for the first time. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The action to take on low-risk virus files. Required when configuring the real-time defense policy for the first time. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file.
	//
	// - **None**: Takes no action.
	//
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// The matching mode for the effective scope. Required when configuring the real-time defense policy for the first time. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups. UserGroupIds is required in this case.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The maximum percentage of endpoint CPU that real-time defense can consume. Valid values: 0 to 100. When configuring for the first time, the value is stored as 0 but takes effect as 30.
	//
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// The action to take on medium-risk virus files. Required when configuring the real-time defense policy for the first time. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file.
	//
	// example:
	//
	// Notify
	MidRiskOperation *string `json:"MidRiskOperation,omitempty" xml:"MidRiskOperation,omitempty"`
	// The collection of virus types to be handled by real-time defense. Duplicates are not allowed. Required when configuring the real-time defense policy for the first time. When the policy already exists, this parameter performs a full replacement. The collection you pass in replaces the existing configuration.
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// The enabling status. Required when configuring the real-time defense policy for the first time. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of user group IDs to which the policy applies. Required when MatchMode is set to UserGroupNormal. Not allowed when MatchMode is set to UserGroupAll. At least 1 and at most 100 entries are allowed. Duplicates are not allowed. When MatchMode is UserGroupNormal, you must pass in the complete user group collection on every call, even when modifying only other parameters.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The exception user list. Users in this list are excluded from real-time defense. A maximum of 1000 entries are allowed. Duplicates are not allowed. This parameter performs a full replacement. The list you pass in replaces the existing list.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateAntiVirusRealTimeDefenceStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntiVirusRealTimeDefenceStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetHighRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.HighRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetLowRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.LowRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetMatchMode(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetMaxCpuUsage(v int64) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.MaxCpuUsage = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetMidRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.MidRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetScanTargets(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.ScanTargets = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetStatus(v string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.Status = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetUserGroupIds(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) SetWhitelist(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyRequest {
	s.Whitelist = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyRequest) Validate() error {
	return dara.Validate(s)
}
