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
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541fxxxxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// example:
	//
	// group1,group2
	ConsumerGroupIds []*string `json:"consumerGroupIds,omitempty" xml:"consumerGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1001,1002,1003
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// week
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 团队规则
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
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
