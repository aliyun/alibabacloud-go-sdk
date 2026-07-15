// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayQuotaRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConflictHash(v string) *AddGatewayQuotaRuleRequest
	GetConflictHash() *string
	SetConsumerGroupIds(v []*string) *AddGatewayQuotaRuleRequest
	GetConsumerGroupIds() []*string
	SetConsumerIds(v []*string) *AddGatewayQuotaRuleRequest
	GetConsumerIds() []*string
	SetDryRun(v bool) *AddGatewayQuotaRuleRequest
	GetDryRun() *bool
	SetOverwrite(v bool) *AddGatewayQuotaRuleRequest
	GetOverwrite() *bool
	SetPeriodMultiplier(v int64) *AddGatewayQuotaRuleRequest
	GetPeriodMultiplier() *int64
	SetPeriodType(v string) *AddGatewayQuotaRuleRequest
	GetPeriodType() *string
	SetQuotaDimension(v string) *AddGatewayQuotaRuleRequest
	GetQuotaDimension() *string
	SetQuotaLimit(v int64) *AddGatewayQuotaRuleRequest
	GetQuotaLimit() *int64
	SetRuleName(v string) *AddGatewayQuotaRuleRequest
	GetRuleName() *string
	SetTimezone(v string) *AddGatewayQuotaRuleRequest
	GetTimezone() *string
	SetWindowAlignment(v string) *AddGatewayQuotaRuleRequest
	GetWindowAlignment() *string
}

type AddGatewayQuotaRuleRequest struct {
	// The conflict snapshot hash, used to prevent concurrent dirty overwrites during confirmation. Obtain this value from the response of a previous dry run (dryRun=true).
	//
	// This parameter is not required in the following cases: no conflicts exist, the request is a dry run (dryRun=true), or overwrite is set to false.
	//
	// If dryRun is set to false and overwrite is set to true but this parameter is not specified or the value has expired, the system returns accepted=false with a new conflict preview. Perform a new dry run to confirm the updated conflicts.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Deprecated
	//
	// The list of consumer group IDs. This parameter is not supported.
	//
	// example:
	//
	// group1,group2
	ConsumerGroupIds []*string `json:"consumerGroupIds,omitempty" xml:"consumerGroupIds,omitempty" type:"Repeated"`
	// The list of consumer IDs to bind to the rule. You can specify up to 1,000 consumers in a single request.
	//
	// example:
	//
	// 1001,1002,1003
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run without applying the configuration. A dry run checks whether conflicting rules exist on the bound consumers. For example, a consumer that already has a calendar-day quota rule cannot have another calendar-day quota rule added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting when conflicts exist. If overwriting is allowed, the conflicting consumers are unbound from the old rule and bound to the new rule.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The period multiplier, which specifies the number of periods after which the quota resets. This parameter is required for custom period rules. Minimum value: 1. Maximum value: 60.
	//
	// example:
	//
	// 10
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// The period unit. For calendar periods, the value can be day, week, or month. For custom periods, only day is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// week
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The quota dimension or throttling type. Valid values: token and credit. The credit quota applies only to dedicated instances of version 2.1.19 or later.
	//
	// This parameter is required.
	//
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// The total available quota per period (the limit).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// team-rule
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The time zone for the calendar period, in UTC+x format.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The reset period type. Valid values:
	//
	// - calendar: calendar period. The period starts from the beginning of a calendar day, week, or month.
	//
	// - epoch: custom period. The period starts from the time the rule is applied. The custom period applies only to dedicated instances of version 2.1.19 or later.
	//
	// example:
	//
	// calendar
	WindowAlignment *string `json:"windowAlignment,omitempty" xml:"windowAlignment,omitempty"`
}

func (s AddGatewayQuotaRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleRequest) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *AddGatewayQuotaRuleRequest) GetConsumerGroupIds() []*string {
	return s.ConsumerGroupIds
}

func (s *AddGatewayQuotaRuleRequest) GetConsumerIds() []*string {
	return s.ConsumerIds
}

func (s *AddGatewayQuotaRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddGatewayQuotaRuleRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *AddGatewayQuotaRuleRequest) GetPeriodMultiplier() *int64 {
	return s.PeriodMultiplier
}

func (s *AddGatewayQuotaRuleRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *AddGatewayQuotaRuleRequest) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *AddGatewayQuotaRuleRequest) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *AddGatewayQuotaRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddGatewayQuotaRuleRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *AddGatewayQuotaRuleRequest) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *AddGatewayQuotaRuleRequest) SetConflictHash(v string) *AddGatewayQuotaRuleRequest {
	s.ConflictHash = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetConsumerGroupIds(v []*string) *AddGatewayQuotaRuleRequest {
	s.ConsumerGroupIds = v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetConsumerIds(v []*string) *AddGatewayQuotaRuleRequest {
	s.ConsumerIds = v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetDryRun(v bool) *AddGatewayQuotaRuleRequest {
	s.DryRun = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetOverwrite(v bool) *AddGatewayQuotaRuleRequest {
	s.Overwrite = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetPeriodMultiplier(v int64) *AddGatewayQuotaRuleRequest {
	s.PeriodMultiplier = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetPeriodType(v string) *AddGatewayQuotaRuleRequest {
	s.PeriodType = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetQuotaDimension(v string) *AddGatewayQuotaRuleRequest {
	s.QuotaDimension = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetQuotaLimit(v int64) *AddGatewayQuotaRuleRequest {
	s.QuotaLimit = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetRuleName(v string) *AddGatewayQuotaRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetTimezone(v string) *AddGatewayQuotaRuleRequest {
	s.Timezone = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) SetWindowAlignment(v string) *AddGatewayQuotaRuleRequest {
	s.WindowAlignment = &v
	return s
}

func (s *AddGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
