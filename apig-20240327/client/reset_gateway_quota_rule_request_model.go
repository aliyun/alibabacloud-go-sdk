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
	// The conflict snapshot hash, which is used to prevent concurrent dirty overwrites when overwriting is confirmed. You can obtain this value by first executing a request with dryRun=true and extracting it from the returned data.
	//
	// This parameter is not required in the following scenarios: no conflicts exist; the request is only a dry run with dryRun=true; overwrite=false, meaning overwriting is not confirmed.
	//
	// If dryRun=false and overwrite=true but this parameter is not provided or has expired and does not match, the backend returns accepted=false along with a new conflict preview. You must perform a dry run again to confirm the new conflicts.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Specifies whether to perform a dry run only without saving to the database or delivering actual configurations. This is used to check whether conflicting rules exist on the bound subjects. A single subject cannot be configured with two natural period quotas of the same period. For example, a subject that already has a natural daily quota cannot be added with another natural daily quota rule.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting when a conflict occurs. Note that if overwriting is allowed, the conflicting subjects (consumers or consumer groups) are unbound from the old rule and bound to the new rule.
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
	// The period type. Natural periods support daily, weekly, and monthly statistics, and valid values are day, week, and month. Custom periods support only daily statistics, and the fixed value is day.
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
	// The time zone corresponding to the natural period, in UTC+x format.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The period type for the reset. Natural periods (windowAlignment=calendar) and custom periods (windowAlignment=epoch) are supported.
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
