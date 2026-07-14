// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupDiscoverRule interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ObserveGroupDiscoverRule
	GetEnabled() *bool
	SetEntityType(v string) *ObserveGroupDiscoverRule
	GetEntityType() *string
	SetEntityTypes(v []*string) *ObserveGroupDiscoverRule
	GetEntityTypes() []*string
	SetGmtCreate(v int64) *ObserveGroupDiscoverRule
	GetGmtCreate() *int64
	SetInstanceIds(v []*string) *ObserveGroupDiscoverRule
	GetInstanceIds() []*string
	SetNameRules(v *ObserveGroupDiscoverRuleNameRules) *ObserveGroupDiscoverRule
	GetNameRules() *ObserveGroupDiscoverRuleNameRules
	SetRegionIds(v []*string) *ObserveGroupDiscoverRule
	GetRegionIds() []*string
	SetResourceGroupId(v string) *ObserveGroupDiscoverRule
	GetResourceGroupId() *string
	SetRuleId(v string) *ObserveGroupDiscoverRule
	GetRuleId() *string
	SetRuleType(v string) *ObserveGroupDiscoverRule
	GetRuleType() *string
	SetScope(v string) *ObserveGroupDiscoverRule
	GetScope() *string
	SetSpl(v string) *ObserveGroupDiscoverRule
	GetSpl() *string
	SetTagRules(v *ObserveGroupDiscoverRuleTagRules) *ObserveGroupDiscoverRule
	GetTagRules() *ObserveGroupDiscoverRuleTagRules
	SetUserId(v string) *ObserveGroupDiscoverRule
	GetUserId() *string
}

type ObserveGroupDiscoverRule struct {
	// Indicates whether the rule is enabled. If set to false, the data plane skips this rule and does not perform matching, tagging, or delivery.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The entity type (legacy). Retained for backward compatibility. Use entityTypes instead.
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The list of entity types. A single rule can match multiple types, such as acs.ecs.instance, acs.rds.instance, and acs.arms.service.
	EntityTypes []*string `json:"entityTypes,omitempty" xml:"entityTypes,omitempty" type:"Repeated"`
	// The time when the rule was created, in UNIX millisecond timestamp format. This value is used for display in the console.
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The list of manually specified instance IDs in enumeration mode, including instances synchronized manually in version 1.0.
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	// The name matching rules.
	NameRules *ObserveGroupDiscoverRuleNameRules `json:"nameRules,omitempty" xml:"nameRules,omitempty" type:"Struct"`
	// The list of region IDs used for filtering by region.
	RegionIds []*string `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
	// The resource group ID used for filtering.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The stable rule ID used as an anchor for editing, deleting, and enabling or disabling operations. Format: dr-<16-character hash>.
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The matching method. Valid values: byTag, byResourceGroup, byInstanceName, byManual, and bySpl.
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// The applicable scope. Valid values: all (all entity types, exclusive) and entity (specified entity types).
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The full SPL expression for advanced configuration. If this parameter is not empty, it takes precedence over other filter fields.
	Spl *string `json:"spl,omitempty" xml:"spl,omitempty"`
	// The tag matching rules.
	TagRules *ObserveGroupDiscoverRuleTagRules `json:"tagRules,omitempty" xml:"tagRules,omitempty" type:"Struct"`
	// The UID of the user to whom the rule belongs.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ObserveGroupDiscoverRule) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRule) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *ObserveGroupDiscoverRule) GetEntityType() *string {
	return s.EntityType
}

func (s *ObserveGroupDiscoverRule) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *ObserveGroupDiscoverRule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ObserveGroupDiscoverRule) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ObserveGroupDiscoverRule) GetNameRules() *ObserveGroupDiscoverRuleNameRules {
	return s.NameRules
}

func (s *ObserveGroupDiscoverRule) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *ObserveGroupDiscoverRule) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ObserveGroupDiscoverRule) GetRuleId() *string {
	return s.RuleId
}

func (s *ObserveGroupDiscoverRule) GetRuleType() *string {
	return s.RuleType
}

func (s *ObserveGroupDiscoverRule) GetScope() *string {
	return s.Scope
}

func (s *ObserveGroupDiscoverRule) GetSpl() *string {
	return s.Spl
}

func (s *ObserveGroupDiscoverRule) GetTagRules() *ObserveGroupDiscoverRuleTagRules {
	return s.TagRules
}

func (s *ObserveGroupDiscoverRule) GetUserId() *string {
	return s.UserId
}

func (s *ObserveGroupDiscoverRule) SetEnabled(v bool) *ObserveGroupDiscoverRule {
	s.Enabled = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetEntityType(v string) *ObserveGroupDiscoverRule {
	s.EntityType = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetEntityTypes(v []*string) *ObserveGroupDiscoverRule {
	s.EntityTypes = v
	return s
}

func (s *ObserveGroupDiscoverRule) SetGmtCreate(v int64) *ObserveGroupDiscoverRule {
	s.GmtCreate = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetInstanceIds(v []*string) *ObserveGroupDiscoverRule {
	s.InstanceIds = v
	return s
}

func (s *ObserveGroupDiscoverRule) SetNameRules(v *ObserveGroupDiscoverRuleNameRules) *ObserveGroupDiscoverRule {
	s.NameRules = v
	return s
}

func (s *ObserveGroupDiscoverRule) SetRegionIds(v []*string) *ObserveGroupDiscoverRule {
	s.RegionIds = v
	return s
}

func (s *ObserveGroupDiscoverRule) SetResourceGroupId(v string) *ObserveGroupDiscoverRule {
	s.ResourceGroupId = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetRuleId(v string) *ObserveGroupDiscoverRule {
	s.RuleId = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetRuleType(v string) *ObserveGroupDiscoverRule {
	s.RuleType = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetScope(v string) *ObserveGroupDiscoverRule {
	s.Scope = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetSpl(v string) *ObserveGroupDiscoverRule {
	s.Spl = &v
	return s
}

func (s *ObserveGroupDiscoverRule) SetTagRules(v *ObserveGroupDiscoverRuleTagRules) *ObserveGroupDiscoverRule {
	s.TagRules = v
	return s
}

func (s *ObserveGroupDiscoverRule) SetUserId(v string) *ObserveGroupDiscoverRule {
	s.UserId = &v
	return s
}

func (s *ObserveGroupDiscoverRule) Validate() error {
	if s.NameRules != nil {
		if err := s.NameRules.Validate(); err != nil {
			return err
		}
	}
	if s.TagRules != nil {
		if err := s.TagRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObserveGroupDiscoverRuleNameRules struct {
	// The name matching logic.
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The name condition list.
	Tags []*ObserveGroupDiscoverRuleNameRulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ObserveGroupDiscoverRuleNameRules) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRuleNameRules) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRuleNameRules) GetOp() *string {
	return s.Op
}

func (s *ObserveGroupDiscoverRuleNameRules) GetTags() []*ObserveGroupDiscoverRuleNameRulesTags {
	return s.Tags
}

func (s *ObserveGroupDiscoverRuleNameRules) SetOp(v string) *ObserveGroupDiscoverRuleNameRules {
	s.Op = &v
	return s
}

func (s *ObserveGroupDiscoverRuleNameRules) SetTags(v []*ObserveGroupDiscoverRuleNameRulesTags) *ObserveGroupDiscoverRuleNameRules {
	s.Tags = v
	return s
}

func (s *ObserveGroupDiscoverRuleNameRules) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ObserveGroupDiscoverRuleNameRulesTags struct {
	// The matching operation.
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The matching value list.
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s ObserveGroupDiscoverRuleNameRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRuleNameRulesTags) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRuleNameRulesTags) GetOp() *string {
	return s.Op
}

func (s *ObserveGroupDiscoverRuleNameRulesTags) GetTagValues() []*string {
	return s.TagValues
}

func (s *ObserveGroupDiscoverRuleNameRulesTags) SetOp(v string) *ObserveGroupDiscoverRuleNameRulesTags {
	s.Op = &v
	return s
}

func (s *ObserveGroupDiscoverRuleNameRulesTags) SetTagValues(v []*string) *ObserveGroupDiscoverRuleNameRulesTags {
	s.TagValues = v
	return s
}

func (s *ObserveGroupDiscoverRuleNameRulesTags) Validate() error {
	return dara.Validate(s)
}

type ObserveGroupDiscoverRuleTagRules struct {
	// The tag matching logic.
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The tag condition list.
	Tags []*ObserveGroupDiscoverRuleTagRulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ObserveGroupDiscoverRuleTagRules) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRuleTagRules) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRuleTagRules) GetOp() *string {
	return s.Op
}

func (s *ObserveGroupDiscoverRuleTagRules) GetTags() []*ObserveGroupDiscoverRuleTagRulesTags {
	return s.Tags
}

func (s *ObserveGroupDiscoverRuleTagRules) SetOp(v string) *ObserveGroupDiscoverRuleTagRules {
	s.Op = &v
	return s
}

func (s *ObserveGroupDiscoverRuleTagRules) SetTags(v []*ObserveGroupDiscoverRuleTagRulesTags) *ObserveGroupDiscoverRuleTagRules {
	s.Tags = v
	return s
}

func (s *ObserveGroupDiscoverRuleTagRules) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ObserveGroupDiscoverRuleTagRulesTags struct {
	// The matching operation.
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The tag key.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value list.
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s ObserveGroupDiscoverRuleTagRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRuleTagRulesTags) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) GetOp() *string {
	return s.Op
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) GetTagValues() []*string {
	return s.TagValues
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) SetOp(v string) *ObserveGroupDiscoverRuleTagRulesTags {
	s.Op = &v
	return s
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) SetTagKey(v string) *ObserveGroupDiscoverRuleTagRulesTags {
	s.TagKey = &v
	return s
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) SetTagValues(v []*string) *ObserveGroupDiscoverRuleTagRulesTags {
	s.TagValues = v
	return s
}

func (s *ObserveGroupDiscoverRuleTagRulesTags) Validate() error {
	return dara.Validate(s)
}
