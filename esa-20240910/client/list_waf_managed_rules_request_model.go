// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackType(v int32) *ListWafManagedRulesRequest
	GetAttackType() *int32
	SetId(v int64) *ListWafManagedRulesRequest
	GetId() *int64
	SetInstanceId(v string) *ListWafManagedRulesRequest
	GetInstanceId() *string
	SetLanguage(v string) *ListWafManagedRulesRequest
	GetLanguage() *string
	SetManagedRuleset(v *ListWafManagedRulesRequestManagedRuleset) *ListWafManagedRulesRequest
	GetManagedRuleset() *ListWafManagedRulesRequestManagedRuleset
	SetPageNumber(v int32) *ListWafManagedRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafManagedRulesRequest
	GetPageSize() *int32
	SetProtectionLevel(v int32) *ListWafManagedRulesRequest
	GetProtectionLevel() *int32
	SetQueryArgs(v *ListWafManagedRulesRequestQueryArgs) *ListWafManagedRulesRequest
	GetQueryArgs() *ListWafManagedRulesRequestQueryArgs
	SetSiteId(v int64) *ListWafManagedRulesRequest
	GetSiteId() *int64
}

type ListWafManagedRulesRequest struct {
	// The attack type of the vulnerability prevention event. Valid values:
	//
	// - SQL injection
	//
	// - cross-site scripting (XSS)
	//
	// - code execute
	//
	// - CRLF
	//
	// - local file inclusion (LFI)
	//
	// - remote file inclusion (RFI)
	//
	// - webshell
	//
	// - cross-site request forgery
	//
	// - Others
	//
	// - SEMA
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The ID of the WAF rule.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The WAF instance ID.
	//
	// example:
	//
	// esa-site-awmmx25y2igw
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language type. The response is returned in the specified language. Valid values:
	//
	// - **en**: English.
	//
	// - **zh**: Chinese.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The managed ruleset configuration in JSON string format.
	//
	// Contains the AttackType, ProtectionLevel, Action, and ManagedRules subfields. When ProtectionLevel is set to -1 (custom mode), specify the status and action for each rule through the ManagedRules array.
	ManagedRuleset *ListWafManagedRulesRequestManagedRuleset `json:"ManagedRuleset,omitempty" xml:"ManagedRuleset,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The currently saved protection level, which represents the existing configuration state in the database.
	//
	// Valid values: -1 (custom mode), 1 (loose), 2 (medium), 3 (strict), 4 (super strict).
	//
	// Difference from ManagedRuleset.ProtectionLevel: this parameter indicates the currently effective configuration, while ManagedRuleset.ProtectionLevel indicates the target value being passed in.
	//
	// example:
	//
	// 1
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The query conditions.
	//
	// example:
	//
	// {\\"Status\\":\\"\\",\\"ProtectionLevels\\":[2,1],\\"Action\\":\\"\\",\\"IdNameLike\\":\\"\\"}
	QueryArgs *ListWafManagedRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequest) GetAttackType() *int32 {
	return s.AttackType
}

func (s *ListWafManagedRulesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListWafManagedRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWafManagedRulesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListWafManagedRulesRequest) GetManagedRuleset() *ListWafManagedRulesRequestManagedRuleset {
	return s.ManagedRuleset
}

func (s *ListWafManagedRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafManagedRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafManagedRulesRequest) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *ListWafManagedRulesRequest) GetQueryArgs() *ListWafManagedRulesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListWafManagedRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafManagedRulesRequest) SetAttackType(v int32) *ListWafManagedRulesRequest {
	s.AttackType = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetId(v int64) *ListWafManagedRulesRequest {
	s.Id = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetInstanceId(v string) *ListWafManagedRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetLanguage(v string) *ListWafManagedRulesRequest {
	s.Language = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetManagedRuleset(v *ListWafManagedRulesRequestManagedRuleset) *ListWafManagedRulesRequest {
	s.ManagedRuleset = v
	return s
}

func (s *ListWafManagedRulesRequest) SetPageNumber(v int32) *ListWafManagedRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetPageSize(v int32) *ListWafManagedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetProtectionLevel(v int32) *ListWafManagedRulesRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetQueryArgs(v *ListWafManagedRulesRequestQueryArgs) *ListWafManagedRulesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafManagedRulesRequest) SetSiteId(v int64) *ListWafManagedRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafManagedRulesRequest) Validate() error {
	if s.ManagedRuleset != nil {
		if err := s.ManagedRuleset.Validate(); err != nil {
			return err
		}
	}
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWafManagedRulesRequestManagedRuleset struct {
	// The unified action when ProtectionLevel is greater than 0. This parameter cannot be empty in this case.
	//
	// Common valid values: monitor, deny, js, captcha. The actual available values depend on the instance quota.
	//
	// example:
	//
	// monitor
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The attack type encoding. The value cannot be 0.
	//
	// Example values: 11 (SQL injection), 12 (XSS), 13 (code execute), 14 (CRLF), 15 (local file inclusion (LFI)), 16 (remote file inclusion (RFI)), 17 (WebShell), 22 (command injection), 26 (SSRF), 27 (path traversal), 28 (protocol violation), 31 (scanner behavior).
	//
	// example:
	//
	// 11
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The rule configuration list in custom mode. This parameter is used only when ProtectionLevel is set to -1.
	//
	// Each element contains Id, Status, and Action, which are used to specify the enabled status and action for each managed rule.
	ManagedRules []*ListWafManagedRulesRequestManagedRulesetManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	// The protection level within the ruleset.
	//
	// Valid values: -1 (custom mode, specify each rule through ManagedRules), 1 (loose), 2 (medium), 3 (strict), 4 (super strict).
	//
	// When the value is -1, ManagedRules cannot be empty. When the value is greater than 0, Action cannot be empty.
	//
	// example:
	//
	// -1
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
}

func (s ListWafManagedRulesRequestManagedRuleset) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequestManagedRuleset) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequestManagedRuleset) GetAction() *string {
	return s.Action
}

func (s *ListWafManagedRulesRequestManagedRuleset) GetAttackType() *int32 {
	return s.AttackType
}

func (s *ListWafManagedRulesRequestManagedRuleset) GetManagedRules() []*ListWafManagedRulesRequestManagedRulesetManagedRules {
	return s.ManagedRules
}

func (s *ListWafManagedRulesRequestManagedRuleset) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *ListWafManagedRulesRequestManagedRuleset) SetAction(v string) *ListWafManagedRulesRequestManagedRuleset {
	s.Action = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRuleset) SetAttackType(v int32) *ListWafManagedRulesRequestManagedRuleset {
	s.AttackType = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRuleset) SetManagedRules(v []*ListWafManagedRulesRequestManagedRulesetManagedRules) *ListWafManagedRulesRequestManagedRuleset {
	s.ManagedRules = v
	return s
}

func (s *ListWafManagedRulesRequestManagedRuleset) SetProtectionLevel(v int32) *ListWafManagedRulesRequestManagedRuleset {
	s.ProtectionLevel = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRuleset) Validate() error {
	if s.ManagedRules != nil {
		for _, item := range s.ManagedRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWafManagedRulesRequestManagedRulesetManagedRules struct {
	// The action for a single rule. This parameter takes effect only in custom mode (ProtectionLevel = -1).
	//
	// Common valid values: monitor, deny, js, captcha. The actual available values depend on the instance quota.
	//
	// example:
	//
	// js
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The unique ID of a single managed rule.
	//
	// example:
	//
	// 20611349
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule enabled status.
	//
	// Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafManagedRulesRequestManagedRulesetManagedRules) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequestManagedRulesetManagedRules) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) GetAction() *string {
	return s.Action
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) GetId() *int64 {
	return s.Id
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) GetStatus() *string {
	return s.Status
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) SetAction(v string) *ListWafManagedRulesRequestManagedRulesetManagedRules {
	s.Action = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) SetId(v int64) *ListWafManagedRulesRequestManagedRulesetManagedRules {
	s.Id = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) SetStatus(v string) *ListWafManagedRulesRequestManagedRulesetManagedRules {
	s.Status = &v
	return s
}

func (s *ListWafManagedRulesRequestManagedRulesetManagedRules) Validate() error {
	return dara.Validate(s)
}

type ListWafManagedRulesRequestQueryArgs struct {
	// The action.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Fuzzy match by rule ID or rule name.
	//
	// example:
	//
	// example
	IdNameLike *string `json:"IdNameLike,omitempty" xml:"IdNameLike,omitempty"`
	// The list of rule protection levels.
	ProtectionLevels []*int32 `json:"ProtectionLevels,omitempty" xml:"ProtectionLevels,omitempty" type:"Repeated"`
	// The status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafManagedRulesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequestQueryArgs) GetAction() *string {
	return s.Action
}

func (s *ListWafManagedRulesRequestQueryArgs) GetIdNameLike() *string {
	return s.IdNameLike
}

func (s *ListWafManagedRulesRequestQueryArgs) GetProtectionLevels() []*int32 {
	return s.ProtectionLevels
}

func (s *ListWafManagedRulesRequestQueryArgs) GetStatus() *string {
	return s.Status
}

func (s *ListWafManagedRulesRequestQueryArgs) SetAction(v string) *ListWafManagedRulesRequestQueryArgs {
	s.Action = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetIdNameLike(v string) *ListWafManagedRulesRequestQueryArgs {
	s.IdNameLike = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetProtectionLevels(v []*int32) *ListWafManagedRulesRequestQueryArgs {
	s.ProtectionLevels = v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetStatus(v string) *ListWafManagedRulesRequestQueryArgs {
	s.Status = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
