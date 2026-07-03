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
	// The conflict snapshot hash, used to prevent concurrent dirty overwrites during confirmation. Obtain this value from the response data of a previous dryRun=true call.
	//
	// This parameter is not required in the following cases: no conflict exists, the request is a dry run (dryRun=true), or overwrite=false (overwrite not confirmed).
	//
	// When dryRun=false and overwrite=true, if this parameter is not provided or the value has expired and no longer matches, the backend returns accepted=false with a new conflict preview. You must perform a dry run again to confirm the new conflict.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Specifies whether to perform only a dry run without delivering the actual configuration. A dry run checks whether conflicting rules exist on the bound consumers. For example, a consumer that already has a calendar-day quota cannot have another calendar-day quota rule added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting on conflict. If overwriting is allowed, the conflicting entity (consumer) is unbound from the old rule and bound to the new rule.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The period multiplier, indicating how many periods elapse before the quota resets. Returned when the rule uses a custom period. Minimum value: 1. Maximum value: 60.
	//
	// example:
	//
	// 1
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// The period type. Valid values: day, week, or month.
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
	// The time zone corresponding to the calendar period (UTC+x format).
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The period type after the reset. Currently only calendar periods are supported, which means windowAlignment="calendar".
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
