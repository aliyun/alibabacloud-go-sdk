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
	// This parameter is required.
	AttackType     *int32                                    `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	Id             *int64                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId     *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Language       *string                                   `json:"Language,omitempty" xml:"Language,omitempty"`
	ManagedRuleset *ListWafManagedRulesRequestManagedRuleset `json:"ManagedRuleset,omitempty" xml:"ManagedRuleset,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize        *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProtectionLevel *int32                               `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	QueryArgs       *ListWafManagedRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
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
	Action          *string                                                 `json:"Action,omitempty" xml:"Action,omitempty"`
	AttackType      *int32                                                  `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	ManagedRules    []*ListWafManagedRulesRequestManagedRulesetManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	ProtectionLevel *int32                                                  `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
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
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
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
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// example
	IdNameLike       *string  `json:"IdNameLike,omitempty" xml:"IdNameLike,omitempty"`
	ProtectionLevels []*int32 `json:"ProtectionLevels,omitempty" xml:"ProtectionLevels,omitempty" type:"Repeated"`
	Status           *string  `json:"Status,omitempty" xml:"Status,omitempty"`
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
