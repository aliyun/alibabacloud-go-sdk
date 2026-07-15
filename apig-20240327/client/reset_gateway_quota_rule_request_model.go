// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGatewayQuotaRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConflictHash(v string) *ResetGatewayQuotaRuleRequest
	GetConflictHash() *string
	SetDryRun(v bool) *ResetGatewayQuotaRuleRequest
	GetDryRun() *bool
	SetOverwrite(v bool) *ResetGatewayQuotaRuleRequest
	GetOverwrite() *bool
	SetPeriodMultiplier(v int64) *ResetGatewayQuotaRuleRequest
	GetPeriodMultiplier() *int64
	SetPeriodType(v string) *ResetGatewayQuotaRuleRequest
	GetPeriodType() *string
	SetQuotaLimit(v int64) *ResetGatewayQuotaRuleRequest
	GetQuotaLimit() *int64
	SetTimezone(v string) *ResetGatewayQuotaRuleRequest
	GetTimezone() *string
	SetWindowAlignment(v string) *ResetGatewayQuotaRuleRequest
	GetWindowAlignment() *string
}

type ResetGatewayQuotaRuleRequest struct {
	// The conflict snapshot hash, used to prevent concurrent dirty overwrites during a confirmed overwrite. Obtain this value from the response of a previous dry run (dryRun=true).
	//
	// You do not need to specify this parameter in the following cases: no conflicts exist, you are performing a dry run (dryRun=true), or you set overwrite to false.
	//
	// If dryRun is set to false and overwrite is set to true but this parameter is not specified or the value has expired, the backend returns accepted=false with a new conflict preview. You must perform a new dry run to confirm the updated conflicts.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Specifies whether to perform only a dry run without persisting or applying the configuration. A dry run checks whether conflicting rules exist on the bound consumers. For example, a consumer that already has a calendar-day quota cannot have another calendar-day quota rule added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting when conflicts exist. If overwriting is allowed, the conflicting entity (consumer) is unbound from the old rule and bound to the new rule.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The period multiplier, which specifies the number of periods after which the quota is reset. This parameter is returned when the rule uses a custom period. Minimum value: 1. Maximum value: 60.
	//
	// example:
	//
	// 1
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// The period type. Calendar periods support daily, weekly, and monthly statistics. Valid values: day, week, and month. Custom periods support only daily statistics. The value is fixed to day.
	//
	// example:
	//
	// week
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The total available quota per period after the reset.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// The time zone for the calendar period, in UTC+x format.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The period alignment type after the reset. Valid values:
	//
	// - calendar: calendar period.
	//
	// - epoch: custom period. Custom periods are supported only on dedicated instances running version 2.1.19 or later.
	//
	// example:
	//
	// calendar
	WindowAlignment *string `json:"windowAlignment,omitempty" xml:"windowAlignment,omitempty"`
}

func (s ResetGatewayQuotaRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleRequest) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleRequest) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *ResetGatewayQuotaRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResetGatewayQuotaRuleRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ResetGatewayQuotaRuleRequest) GetPeriodMultiplier() *int64 {
	return s.PeriodMultiplier
}

func (s *ResetGatewayQuotaRuleRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *ResetGatewayQuotaRuleRequest) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *ResetGatewayQuotaRuleRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *ResetGatewayQuotaRuleRequest) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *ResetGatewayQuotaRuleRequest) SetConflictHash(v string) *ResetGatewayQuotaRuleRequest {
	s.ConflictHash = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetDryRun(v bool) *ResetGatewayQuotaRuleRequest {
	s.DryRun = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetOverwrite(v bool) *ResetGatewayQuotaRuleRequest {
	s.Overwrite = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetPeriodMultiplier(v int64) *ResetGatewayQuotaRuleRequest {
	s.PeriodMultiplier = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetPeriodType(v string) *ResetGatewayQuotaRuleRequest {
	s.PeriodType = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetQuotaLimit(v int64) *ResetGatewayQuotaRuleRequest {
	s.QuotaLimit = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetTimezone(v string) *ResetGatewayQuotaRuleRequest {
	s.Timezone = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) SetWindowAlignment(v string) *ResetGatewayQuotaRuleRequest {
	s.WindowAlignment = &v
	return s
}

func (s *ResetGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
