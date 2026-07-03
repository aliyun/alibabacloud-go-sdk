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
	// The list of principal (consumer) IDs to bind.
	//
	// example:
	//
	// cs-001,cs-002
	AddIds []*string `json:"addIds,omitempty" xml:"addIds,omitempty" type:"Repeated"`
	// The conflict snapshot hash, used to prevent concurrent dirty overwrites when confirming an overwrite. Obtain this value from the response of a prior dryRun=true call.
	//
	// This parameter is not required in the following cases: no conflicts exist, the call is a dry run (dryRun=true), or overwrite is set to false.
	//
	// When dryRun is set to false and overwrite is set to true, if this parameter is missing or the value has expired and no longer matches, the backend returns accepted=false with a new conflict preview. Perform the dry run again to confirm the new conflicts.
	//
	// example:
	//
	// f8f44dc6cf369a017d56b7197eb4fb5ac4bbb6b09a92b9b41999541f50xxxxxx
	ConflictHash *string `json:"conflictHash,omitempty" xml:"conflictHash,omitempty"`
	// Deprecated
	//
	// The list of consumer group IDs. This parameter is not supported.
	//
	// example:
	//
	// group1,group2
	ConsumerGroupIds []*string `json:"consumerGroupIds,omitempty" xml:"consumerGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run only without persisting or applying the configuration. A dry run checks whether conflicting rules exist on the bound consumers. For example, a consumer that already has a calendar-day quota cannot have another calendar-day quota rule added.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to allow overwriting when a conflict exists. If overwriting is allowed, the conflicting principals (consumers) are unbound from the old rule and bound to the new rule.
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The updated total available quota.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// The list of principal (consumer) IDs to unbind.
	//
	// example:
	//
	// cs003,cs-004
	RemoveIds []*string `json:"removeIds,omitempty" xml:"removeIds,omitempty" type:"Repeated"`
	// The updated rule name.
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
