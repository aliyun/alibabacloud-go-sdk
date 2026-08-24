// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntiVirusRealTimeDefenceStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHighRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetHighRiskOperation() *string
	SetLowRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetLowRiskOperation() *string
	SetMatchMode(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetMatchTargetIds() []*string
	SetMaxCpuUsage(v int64) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetMaxCpuUsage() *int64
	SetMidRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetMidRiskOperation() *string
	SetRequestId(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetRequestId() *string
	SetScanTargets(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetScanTargets() []*string
	SetStatus(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetStatus() *string
	SetStrategyId(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetStrategyId() *string
	SetWhitelist(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody
	GetWhitelist() []*string
}

type GetAntiVirusRealTimeDefenceStrategyResponseBody struct {
	// The action taken on high-risk virus files. Valid values:
	//
	// - **Quarantine**: Quarantines quarantined file.
	//
	// - **Notify**: Reports an alert only without taking action on quarantined file. Quarantine is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// The action taken on low-risk virus files. Valid values:
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
	// The matching mode of the effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Applies only to users in specified user groups. An empty string is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of user group IDs to which the policy applies. An empty list is returned when MatchMode is set to UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The maximum percentage of endpoint CPU that real-time defense can use. The default value 30 is returned if a policy has been configured but this parameter is not separately set. 0 is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// The action taken on medium-risk virus files. Valid values:
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
	// The collection of virus types that the real-time defense handles. An empty list is returned if no real-time defense policy has been configured.
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// The enabling status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled. Disabled is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the real-time defense policy. An empty string is returned if no real-time defense policy has been configured.
	//
	// example:
	//
	// av-rtd-2f5c8e1a7b94****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The list of exempted usernames. Users in this list are not subject to real-time defense. An empty list is returned if no exemption is configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s GetAntiVirusRealTimeDefenceStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAntiVirusRealTimeDefenceStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetHighRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.HighRiskOperation = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetLowRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.LowRiskOperation = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetMatchMode(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetMatchTargetIds(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MatchTargetIds = v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetMaxCpuUsage(v int64) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MaxCpuUsage = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetMidRiskOperation(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.MidRiskOperation = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetRequestId(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetScanTargets(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.ScanTargets = v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetStatus(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.Status = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetStrategyId(v string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) SetWhitelist(v []*string) *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	s.Whitelist = v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
