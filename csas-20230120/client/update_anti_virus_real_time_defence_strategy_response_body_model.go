// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntiVirusRealTimeDefenceStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHighRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetHighRiskOperation() *string
	SetLowRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetLowRiskOperation() *string
	SetMatchMode(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetMatchMode() *string
	SetMidRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetMidRiskOperation() *string
	SetRequestId(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetRequestId() *string
	SetScanTargets(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetScanTargets() []*string
	SetStatus(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetStatus() *string
	SetStrategyId(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetStrategyId() *string
	SetUserGroupIds(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
	GetWhitelist() []*string
}

type UpdateAntiVirusRealTimeDefenceStrategyResponseBody struct {
	// The action to take on high-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file. Quarantine is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The action to take on low-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file.
	//
	// - **None**: Takes no action. None is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups. An empty string is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The action to take on medium-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file. Notify is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// Notify
	MidRiskOperation *string `json:"MidRiskOperation,omitempty" xml:"MidRiskOperation,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The collection of virus types to be handled by real-time defense. An empty list is returned if no real-time defense policy has been configured.
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// The enabling status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled. This value is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The real-time defense policy ID. An empty string is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// av-rtd-2f5c8e1a7b94****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The collection of user group IDs to which the policy applies. An empty list is returned when MatchMode is set to UserGroupAll.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The exception user list. Users in this list are excluded from real-time defense. An empty list is returned if no exception users are configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateAntiVirusRealTimeDefenceStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetHighRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.HighRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetLowRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.LowRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetMatchMode(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetMidRiskOperation(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MidRiskOperation = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetRequestId(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetScanTargets(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.ScanTargets = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetStatus(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetStrategyId(v string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetUserGroupIds(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.UserGroupIds = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) SetWhitelist(v []*string) *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
