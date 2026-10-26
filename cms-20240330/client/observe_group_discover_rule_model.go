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
	// 规则是否启用；false 时数据面跳过该规则（不匹配 / 不打标 / 不下发）。
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// （旧）单实体类型，向后兼容，优先用 entityTypes。
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// 实体类型列表，支持一条规则跨多类型匹配（如 acs.ecs.instance/acs.rds.instance/acs.arms.service）。
	EntityTypes []*string `json:"entityTypes,omitempty" xml:"entityTypes,omitempty" type:"Repeated"`
	// 规则创建时间（Unix 毫秒时间戳），控制台展示用。
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 手动指定的实例 ID 列表（枚举型，含 1.0 手动实例同步）。
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	// 名称匹配规则。
	NameRules *ObserveGroupDiscoverRuleNameRules `json:"nameRules,omitempty" xml:"nameRules,omitempty" type:"Struct"`
	// 按地域过滤的 region ID 列表。
	RegionIds []*string `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
	// 按资源组过滤。
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 规则稳定 ID（编辑/删除/启停锚点），格式 dr-<16位hash>。
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// 匹配方式：byTag / byResourceGroup / byInstanceName / byManual / bySpl。
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// 适用范围：all（全部实体类型，独占）/ entity（指定实体类型）。
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 完整 SPL 表达式（高级配置）；非空时优先生效，覆盖其它过滤字段。
	Spl *string `json:"spl,omitempty" xml:"spl,omitempty"`
	// 标签匹配规则。
	TagRules *ObserveGroupDiscoverRuleTagRules `json:"tagRules,omitempty" xml:"tagRules,omitempty" type:"Struct"`
	// 规则所属用户 UID。
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
	// 名称匹配逻辑。
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// 名称条件列表。
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
	// 匹配操作。
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// 匹配值列表。
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
	// 标签匹配逻辑。
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// 标签条件列表。
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
	// 匹配操作。
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// 标签键。
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// 标签值列表。
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
