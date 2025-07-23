// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDiscoverRule interface {
  dara.Model
  String() string
  GoString() string
  SetAnnotations(v []*EntityDiscoverRuleAnnotations) *EntityDiscoverRule
  GetAnnotations() []*EntityDiscoverRuleAnnotations 
  SetEntityTypes(v []*string) *EntityDiscoverRule
  GetEntityTypes() []*string 
  SetFieldRules(v []*EntityDiscoverRuleFieldRules) *EntityDiscoverRule
  GetFieldRules() []*EntityDiscoverRuleFieldRules 
  SetInstanceIds(v []*string) *EntityDiscoverRule
  GetInstanceIds() []*string 
  SetIpMatchRule(v []*EntityDiscoverRuleIpMatchRule) *EntityDiscoverRule
  GetIpMatchRule() []*EntityDiscoverRuleIpMatchRule 
  SetLabels(v []*EntityDiscoverRuleLabels) *EntityDiscoverRule
  GetLabels() []*EntityDiscoverRuleLabels 
  SetRegionIds(v []*string) *EntityDiscoverRule
  GetRegionIds() []*string 
  SetResourceGroupId(v string) *EntityDiscoverRule
  GetResourceGroupId() *string 
  SetTags(v []*EntityDiscoverRuleTags) *EntityDiscoverRule
  GetTags() []*EntityDiscoverRuleTags 
}

type EntityDiscoverRule struct {
  Annotations []*EntityDiscoverRuleAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
  EntityTypes []*string `json:"entityTypes,omitempty" xml:"entityTypes,omitempty" type:"Repeated"`
  FieldRules []*EntityDiscoverRuleFieldRules `json:"fieldRules,omitempty" xml:"fieldRules,omitempty" type:"Repeated"`
  InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
  IpMatchRule []*EntityDiscoverRuleIpMatchRule `json:"ipMatchRule,omitempty" xml:"ipMatchRule,omitempty" type:"Repeated"`
  Labels []*EntityDiscoverRuleLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
  RegionIds []*string `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
  Tags []*EntityDiscoverRuleTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s EntityDiscoverRule) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRule) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRule) GetAnnotations() []*EntityDiscoverRuleAnnotations  {
  return s.Annotations
}

func (s *EntityDiscoverRule) GetEntityTypes() []*string  {
  return s.EntityTypes
}

func (s *EntityDiscoverRule) GetFieldRules() []*EntityDiscoverRuleFieldRules  {
  return s.FieldRules
}

func (s *EntityDiscoverRule) GetInstanceIds() []*string  {
  return s.InstanceIds
}

func (s *EntityDiscoverRule) GetIpMatchRule() []*EntityDiscoverRuleIpMatchRule  {
  return s.IpMatchRule
}

func (s *EntityDiscoverRule) GetLabels() []*EntityDiscoverRuleLabels  {
  return s.Labels
}

func (s *EntityDiscoverRule) GetRegionIds() []*string  {
  return s.RegionIds
}

func (s *EntityDiscoverRule) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EntityDiscoverRule) GetTags() []*EntityDiscoverRuleTags  {
  return s.Tags
}

func (s *EntityDiscoverRule) SetAnnotations(v []*EntityDiscoverRuleAnnotations) *EntityDiscoverRule {
  s.Annotations = v
  return s
}

func (s *EntityDiscoverRule) SetEntityTypes(v []*string) *EntityDiscoverRule {
  s.EntityTypes = v
  return s
}

func (s *EntityDiscoverRule) SetFieldRules(v []*EntityDiscoverRuleFieldRules) *EntityDiscoverRule {
  s.FieldRules = v
  return s
}

func (s *EntityDiscoverRule) SetInstanceIds(v []*string) *EntityDiscoverRule {
  s.InstanceIds = v
  return s
}

func (s *EntityDiscoverRule) SetIpMatchRule(v []*EntityDiscoverRuleIpMatchRule) *EntityDiscoverRule {
  s.IpMatchRule = v
  return s
}

func (s *EntityDiscoverRule) SetLabels(v []*EntityDiscoverRuleLabels) *EntityDiscoverRule {
  s.Labels = v
  return s
}

func (s *EntityDiscoverRule) SetRegionIds(v []*string) *EntityDiscoverRule {
  s.RegionIds = v
  return s
}

func (s *EntityDiscoverRule) SetResourceGroupId(v string) *EntityDiscoverRule {
  s.ResourceGroupId = &v
  return s
}

func (s *EntityDiscoverRule) SetTags(v []*EntityDiscoverRuleTags) *EntityDiscoverRule {
  s.Tags = v
  return s
}

func (s *EntityDiscoverRule) Validate() error {
  return dara.Validate(s)
}

type EntityDiscoverRuleAnnotations struct {
  Op *string `json:"op,omitempty" xml:"op,omitempty"`
  TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
  TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s EntityDiscoverRuleAnnotations) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRuleAnnotations) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRuleAnnotations) GetOp() *string  {
  return s.Op
}

func (s *EntityDiscoverRuleAnnotations) GetTagKey() *string  {
  return s.TagKey
}

func (s *EntityDiscoverRuleAnnotations) GetTagValues() []*string  {
  return s.TagValues
}

func (s *EntityDiscoverRuleAnnotations) SetOp(v string) *EntityDiscoverRuleAnnotations {
  s.Op = &v
  return s
}

func (s *EntityDiscoverRuleAnnotations) SetTagKey(v string) *EntityDiscoverRuleAnnotations {
  s.TagKey = &v
  return s
}

func (s *EntityDiscoverRuleAnnotations) SetTagValues(v []*string) *EntityDiscoverRuleAnnotations {
  s.TagValues = v
  return s
}

func (s *EntityDiscoverRuleAnnotations) Validate() error {
  return dara.Validate(s)
}

type EntityDiscoverRuleFieldRules struct {
  FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
  FieldValues []*string `json:"fieldValues,omitempty" xml:"fieldValues,omitempty" type:"Repeated"`
  Op *string `json:"op,omitempty" xml:"op,omitempty"`
}

func (s EntityDiscoverRuleFieldRules) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRuleFieldRules) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRuleFieldRules) GetFieldKey() *string  {
  return s.FieldKey
}

func (s *EntityDiscoverRuleFieldRules) GetFieldValues() []*string  {
  return s.FieldValues
}

func (s *EntityDiscoverRuleFieldRules) GetOp() *string  {
  return s.Op
}

func (s *EntityDiscoverRuleFieldRules) SetFieldKey(v string) *EntityDiscoverRuleFieldRules {
  s.FieldKey = &v
  return s
}

func (s *EntityDiscoverRuleFieldRules) SetFieldValues(v []*string) *EntityDiscoverRuleFieldRules {
  s.FieldValues = v
  return s
}

func (s *EntityDiscoverRuleFieldRules) SetOp(v string) *EntityDiscoverRuleFieldRules {
  s.Op = &v
  return s
}

func (s *EntityDiscoverRuleFieldRules) Validate() error {
  return dara.Validate(s)
}

type EntityDiscoverRuleIpMatchRule struct {
  IpCIDR *string `json:"ipCIDR,omitempty" xml:"ipCIDR,omitempty"`
  IpFieldKey *string `json:"ipFieldKey,omitempty" xml:"ipFieldKey,omitempty"`
}

func (s EntityDiscoverRuleIpMatchRule) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRuleIpMatchRule) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRuleIpMatchRule) GetIpCIDR() *string  {
  return s.IpCIDR
}

func (s *EntityDiscoverRuleIpMatchRule) GetIpFieldKey() *string  {
  return s.IpFieldKey
}

func (s *EntityDiscoverRuleIpMatchRule) SetIpCIDR(v string) *EntityDiscoverRuleIpMatchRule {
  s.IpCIDR = &v
  return s
}

func (s *EntityDiscoverRuleIpMatchRule) SetIpFieldKey(v string) *EntityDiscoverRuleIpMatchRule {
  s.IpFieldKey = &v
  return s
}

func (s *EntityDiscoverRuleIpMatchRule) Validate() error {
  return dara.Validate(s)
}

type EntityDiscoverRuleLabels struct {
  Op *string `json:"op,omitempty" xml:"op,omitempty"`
  TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
  TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s EntityDiscoverRuleLabels) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRuleLabels) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRuleLabels) GetOp() *string  {
  return s.Op
}

func (s *EntityDiscoverRuleLabels) GetTagKey() *string  {
  return s.TagKey
}

func (s *EntityDiscoverRuleLabels) GetTagValues() []*string  {
  return s.TagValues
}

func (s *EntityDiscoverRuleLabels) SetOp(v string) *EntityDiscoverRuleLabels {
  s.Op = &v
  return s
}

func (s *EntityDiscoverRuleLabels) SetTagKey(v string) *EntityDiscoverRuleLabels {
  s.TagKey = &v
  return s
}

func (s *EntityDiscoverRuleLabels) SetTagValues(v []*string) *EntityDiscoverRuleLabels {
  s.TagValues = v
  return s
}

func (s *EntityDiscoverRuleLabels) Validate() error {
  return dara.Validate(s)
}

type EntityDiscoverRuleTags struct {
  Op *string `json:"op,omitempty" xml:"op,omitempty"`
  TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
  TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s EntityDiscoverRuleTags) String() string {
  return dara.Prettify(s)
}

func (s EntityDiscoverRuleTags) GoString() string {
  return s.String()
}

func (s *EntityDiscoverRuleTags) GetOp() *string  {
  return s.Op
}

func (s *EntityDiscoverRuleTags) GetTagKey() *string  {
  return s.TagKey
}

func (s *EntityDiscoverRuleTags) GetTagValues() []*string  {
  return s.TagValues
}

func (s *EntityDiscoverRuleTags) SetOp(v string) *EntityDiscoverRuleTags {
  s.Op = &v
  return s
}

func (s *EntityDiscoverRuleTags) SetTagKey(v string) *EntityDiscoverRuleTags {
  s.TagKey = &v
  return s
}

func (s *EntityDiscoverRuleTags) SetTagValues(v []*string) *EntityDiscoverRuleTags {
  s.TagValues = v
  return s
}

func (s *EntityDiscoverRuleTags) Validate() error {
  return dara.Validate(s)
}

