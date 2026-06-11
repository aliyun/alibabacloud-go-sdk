// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddIds(v []*string) *UpdateGatewayQuotaRuleRequest
	GetAddIds() []*string
	SetConflictHash(v string) *UpdateGatewayQuotaRuleRequest
	GetConflictHash() *string
	SetConsumerGroupIds(v []*string) *UpdateGatewayQuotaRuleRequest
	GetConsumerGroupIds() []*string
	SetDryRun(v bool) *UpdateGatewayQuotaRuleRequest
	GetDryRun() *bool
	SetOverwrite(v bool) *UpdateGatewayQuotaRuleRequest
	GetOverwrite() *bool
	SetQuotaLimit(v int64) *UpdateGatewayQuotaRuleRequest
	GetQuotaLimit() *int64
	SetRemoveIds(v []*string) *UpdateGatewayQuotaRuleRequest
	GetRemoveIds() []*string
	SetRuleName(v string) *UpdateGatewayQuotaRuleRequest
	GetRuleName() *string
}

type UpdateGatewayQuotaRuleRequest struct {
	// example:
	//
	// cs-001,cs-002
	AddIds []*string `json:"addIds,omitempty" xml:"addIds,omitempty" type:"Repeated"`
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541f50xxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// example:
	//
	// group1,group2
	ConsumerGroupIds []*string `json:"consumerGroupIds,omitempty" xml:"consumerGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DryRun    *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// example:
	//
	// cs003,cs-004
	RemoveIds []*string `json:"removeIds,omitempty" xml:"removeIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// team-rule
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s UpdateGatewayQuotaRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleRequest) GetAddIds() []*string {
	return s.AddIds
}

func (s *UpdateGatewayQuotaRuleRequest) GetConflictHash() *string {
	return s.ConflictHash
}

func (s *UpdateGatewayQuotaRuleRequest) GetConsumerGroupIds() []*string {
	return s.ConsumerGroupIds
}

func (s *UpdateGatewayQuotaRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateGatewayQuotaRuleRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *UpdateGatewayQuotaRuleRequest) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *UpdateGatewayQuotaRuleRequest) GetRemoveIds() []*string {
	return s.RemoveIds
}

func (s *UpdateGatewayQuotaRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateGatewayQuotaRuleRequest) SetAddIds(v []*string) *UpdateGatewayQuotaRuleRequest {
	s.AddIds = v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetConflictHash(v string) *UpdateGatewayQuotaRuleRequest {
	s.ConflictHash = &v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetConsumerGroupIds(v []*string) *UpdateGatewayQuotaRuleRequest {
	s.ConsumerGroupIds = v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetDryRun(v bool) *UpdateGatewayQuotaRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetOverwrite(v bool) *UpdateGatewayQuotaRuleRequest {
	s.Overwrite = &v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetQuotaLimit(v int64) *UpdateGatewayQuotaRuleRequest {
	s.QuotaLimit = &v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetRemoveIds(v []*string) *UpdateGatewayQuotaRuleRequest {
	s.RemoveIds = v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) SetRuleName(v string) *UpdateGatewayQuotaRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
