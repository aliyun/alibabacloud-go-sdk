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
	// The conflict snapshot hash used to prevent concurrent dirty overwrites during confirmation. Obtain this value from the response of a previous dryRun=true call.
	//
	// You do not need to specify this parameter in the following cases: no conflict exists, the request is a dry run (dryRun=true), or overwrite=false.
	//
	// When dryRun=false and overwrite=true, if this parameter is not specified or the value has expired, the backend returns accepted=false with a new conflict preview. Perform a dry run again to confirm the new conflict.
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
	// The list of consumer IDs to bind to the rule.
	//
	// example:
	//
	// 1001,1002,1003
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run without persisting or applying the configuration. A dry run checks whether conflicting rules exist on the bound consumers. For example, a consumer that already has a daily calendar quota cannot have another daily calendar quota rule added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting on conflict. If overwriting is allowed, the conflicting principals (consumers) are unbound from the old rule and bound to the new rule.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The period multiplier.
	//
	// example:
	//
	// 10
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// The period type. Valid values: day (calendar day), week (calendar week), and month (calendar month).
	//
	// This parameter is required.
	//
	// example:
	//
	// week
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The quota dimension or throttling type. Currently, only token is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// The total available quota per period.
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
	// 团队规则
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The time zone for the calendar period, in UTC+x format.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The reset period alignment type. Currently, only calendar alignment is supported, which means windowAlignment="calendar".
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
