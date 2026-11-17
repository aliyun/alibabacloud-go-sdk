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
	// The conflict snapshot hash, which is used to prevent concurrent dirty overwrites when you confirm to overwrite. You can obtain this value by first executing a dry run with dryRun=true and extracting it from the response.
	//
	// You do not need to specify this parameter in the following scenarios: no conflicts exist; you are only performing a dry run with dryRun=true; or overwrite=false, meaning you do not confirm to overwrite.
	//
	// If dryRun=false and overwrite=true, but this parameter is not provided or has expired and does not match, the backend returns accepted=false along with a new conflict preview. You must perform a dry run again to confirm the new conflicts.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Specifies whether to perform a dry run only. A dry run does not deliver the actual configuration. This parameter is used to check whether conflicting rules exist on the bound subjects. A single subject cannot be configured with two calendar period quotas of the same period. For example, a subject that already has a calendar day quota cannot be assigned another calendar day quota rule.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting when conflicts exist. Note that if you allow overwriting, the conflicting subjects (consumers or consumer groups) are unbound from the old rule and bound to the new rule.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The period multiplier, which specifies the number of periods after which the quota is reset. This parameter is required when the rule uses a custom period. Valid values: 1 to 60.
	//
	// example:
	//
	// 1
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// The period type. For a calendar period, statistics can be collected by day, week, or month. Valid values: day, week, and month. For a custom period, statistics can only be collected by day. Valid value: day.
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
	// The period type after the reset. Valid values:
	//
	// - calendar: Calendar period.
	//
	// - epoch: Custom period.
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
