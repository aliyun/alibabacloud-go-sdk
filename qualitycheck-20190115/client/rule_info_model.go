// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReview(v int32) *RuleInfo
	GetAutoReview() *int32
	SetBusinessCategoryNameList(v []*string) *RuleInfo
	GetBusinessCategoryNameList() []*string
	SetCheckType(v int64) *RuleInfo
	GetCheckType() *int64
	SetComments(v string) *RuleInfo
	GetComments() *string
	SetConfigType(v int32) *RuleInfo
	GetConfigType() *int32
	SetCreateEmpName(v string) *RuleInfo
	GetCreateEmpName() *string
	SetCreateEmpid(v string) *RuleInfo
	GetCreateEmpid() *string
	SetCreateTime(v string) *RuleInfo
	GetCreateTime() *string
	SetDeny(v int32) *RuleInfo
	GetDeny() *int32
	SetDialogues(v []*RuleTestDialogue) *RuleInfo
	GetDialogues() []*RuleTestDialogue
	SetEffective(v int32) *RuleInfo
	GetEffective() *int32
	SetEffectiveEndTime(v string) *RuleInfo
	GetEffectiveEndTime() *string
	SetEffectiveStartTime(v string) *RuleInfo
	GetEffectiveStartTime() *string
	SetEndTime(v string) *RuleInfo
	GetEndTime() *string
	SetExternalProperty(v int32) *RuleInfo
	GetExternalProperty() *int32
	SetFullCycle(v int32) *RuleInfo
	GetFullCycle() *int32
	SetGraphFlow(v interface{}) *RuleInfo
	GetGraphFlow() interface{}
	SetIsDelete(v int32) *RuleInfo
	GetIsDelete() *int32
	SetIsOnline(v int32) *RuleInfo
	GetIsOnline() *int32
	SetLambda(v string) *RuleInfo
	GetLambda() *string
	SetLastUpdateEmpName(v string) *RuleInfo
	GetLastUpdateEmpName() *string
	SetLastUpdateEmpid(v string) *RuleInfo
	GetLastUpdateEmpid() *string
	SetLastUpdateTime(v string) *RuleInfo
	GetLastUpdateTime() *string
	SetLevel(v int32) *RuleInfo
	GetLevel() *int32
	SetMeet(v int32) *RuleInfo
	GetMeet() *int32
	SetModifyType(v int32) *RuleInfo
	GetModifyType() *int32
	SetName(v string) *RuleInfo
	GetName() *string
	SetOperationMode(v int32) *RuleInfo
	GetOperationMode() *int32
	SetPreqRule(v *RuleInfoPreqRule) *RuleInfo
	GetPreqRule() *RuleInfoPreqRule
	SetQualityCheckType(v int32) *RuleInfo
	GetQualityCheckType() *int32
	SetRid(v string) *RuleInfo
	GetRid() *string
	SetRuleCategoryName(v string) *RuleInfo
	GetRuleCategoryName() *string
	SetRuleScoreType(v int32) *RuleInfo
	GetRuleScoreType() *int32
	SetRuleType(v int32) *RuleInfo
	GetRuleType() *int32
	SetSchemeCheckType(v *SchemeCheckType) *RuleInfo
	GetSchemeCheckType() *SchemeCheckType
	SetSchemeId(v int64) *RuleInfo
	GetSchemeId() *int64
	SetSchemeName(v string) *RuleInfo
	GetSchemeName() *string
	SetSchemeRuleMappingId(v int64) *RuleInfo
	GetSchemeRuleMappingId() *int64
	SetScoreDeleted(v bool) *RuleInfo
	GetScoreDeleted() *bool
	SetScoreId(v int64) *RuleInfo
	GetScoreId() *int64
	SetScoreName(v string) *RuleInfo
	GetScoreName() *string
	SetScoreNum(v float32) *RuleInfo
	GetScoreNum() *float32
	SetScoreNumType(v int32) *RuleInfo
	GetScoreNumType() *int32
	SetScoreRuleHitType(v int32) *RuleInfo
	GetScoreRuleHitType() *int32
	SetScoreSubId(v int64) *RuleInfo
	GetScoreSubId() *int64
	SetScoreSubName(v string) *RuleInfo
	GetScoreSubName() *string
	SetScoreType(v int32) *RuleInfo
	GetScoreType() *int32
	SetSortIndex(v int32) *RuleInfo
	GetSortIndex() *int32
	SetStartTime(v string) *RuleInfo
	GetStartTime() *string
	SetStatus(v int32) *RuleInfo
	GetStatus() *int32
	SetTargetType(v int32) *RuleInfo
	GetTargetType() *int32
	SetTaskFlowId(v int64) *RuleInfo
	GetTaskFlowId() *int64
	SetTaskFlowType(v int32) *RuleInfo
	GetTaskFlowType() *int32
	SetTriggers(v []*string) *RuleInfo
	GetTriggers() []*string
	SetType(v int32) *RuleInfo
	GetType() *int32
	SetWeight(v string) *RuleInfo
	GetWeight() *string
}

type RuleInfo struct {
	// The review option. This parameter is used for compatibility with the v4.0 protocol. Valid values: `1` (Manual review) and `3` (Automatic review).
	//
	// example:
	//
	// 1
	AutoReview *int32 `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	// A list of business category names.
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	// The check type. Valid values: `0` (Service compliance check), `1` (Service attitude check), `2` (Service professionalism check), `3` (Customer attitude check), and `4` (Service process correctness check).
	//
	// example:
	//
	// 3
	CheckType *int64 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// Remarks on the rule.
	//
	// example:
	//
	// 违规
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The configuration type. This parameter is used for compatibility with the v4.0 protocol. Valid values: `1` (Simple condition configuration) and `2` (Advanced configuration). Default: `1`.
	//
	// example:
	//
	// 1
	ConfigType *int32 `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// 张三
	CreateEmpName *string `json:"CreateEmpName,omitempty" xml:"CreateEmpName,omitempty"`
	// The employee ID of the creator.
	//
	// example:
	//
	// 1
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// The creation time. This value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1641277321000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether editing the rule is denied. Valid values: `1` (denied), `0` or `null` (allowed).
	//
	// example:
	//
	// 1
	Deny *int32 `json:"Deny,omitempty" xml:"Deny,omitempty"`
	// A list of test dialogues.
	Dialogues []*RuleTestDialogue `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	// Indicates whether the rule is in effect. Valid values: `0` (No) and `1` (Yes).
	//
	// example:
	//
	// 1
	Effective *int32 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The effective end time of the rule.
	//
	// example:
	//
	// OperationMode
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The effective start time of the rule.
	//
	// example:
	//
	// OperationMode
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The end time. This value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1641277321000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The external property.
	//
	// example:
	//
	// 1
	ExternalProperty *int32 `json:"ExternalProperty,omitempty" xml:"ExternalProperty,omitempty"`
	// Indicates whether the rule is effective throughout its lifecycle. Valid values: `0` (No) and `1` (Yes).
	//
	// example:
	//
	// 1
	FullCycle *int32 `json:"FullCycle,omitempty" xml:"FullCycle,omitempty"`
	// The graph flow.
	//
	// example:
	//
	// {}
	GraphFlow interface{} `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	// Indicates whether the rule is deleted.
	//
	// example:
	//
	// 1
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Indicates whether the rule is online.
	//
	// example:
	//
	// 1
	IsOnline *int32 `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	// The conditional expression, such as `a&&b`.
	//
	// example:
	//
	// a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	// The name of the employee who last updated the rule.
	//
	// example:
	//
	// 张三
	LastUpdateEmpName *string `json:"LastUpdateEmpName,omitempty" xml:"LastUpdateEmpName,omitempty"`
	// The ID of the employee who last updated the rule.
	//
	// example:
	//
	// 1
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// The last update time. This value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1641277321000
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The severity level of the rule. Valid values: `0` (Severe violation), `1` (Moderate violation), and `2` (Minor violation).
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// Indicates whether the conditions are met.
	//
	// example:
	//
	// 1
	Meet *int32 `json:"Meet,omitempty" xml:"Meet,omitempty"`
	// Indicates whether the rule has been modified within the quality check scheme.
	//
	// example:
	//
	// 1
	ModifyType *int32 `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The rule name.
	//
	// example:
	//
	// 开头语规则
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation mode.
	//
	// example:
	//
	// 1
	OperationMode *int32            `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	PreqRule      *RuleInfoPreqRule `json:"PreqRule,omitempty" xml:"PreqRule,omitempty" type:"Struct"`
	// The quality check type. This parameter is used for compatibility with the v4.0 protocol. Valid values: `0` (offline quality check) and `1` (real-time quality check).
	//
	// example:
	//
	// 1
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// The name of the rule category.
	//
	// example:
	//
	// 正向规则
	RuleCategoryName *string `json:"RuleCategoryName,omitempty" xml:"RuleCategoryName,omitempty"`
	// Specifies whether to score the rule. This parameter is used for compatibility with the v4.0 protocol. Valid values: `1` (Do not score) and `3` (Score).
	//
	// example:
	//
	// 3
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// The rule type. This parameter is used for compatibility with the v4.0 protocol. Valid values: `0` (Default) and `1` (User-created).
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The quality check dimension object.
	SchemeCheckType *SchemeCheckType `json:"SchemeCheckType,omitempty" xml:"SchemeCheckType,omitempty"`
	// The ID of the quality check scheme to which the rule belongs.
	//
	// example:
	//
	// 1
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// The name of the quality check scheme.
	//
	// example:
	//
	// 通用方案
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The mapping ID between the quality check scheme and the rule.
	//
	// example:
	//
	// 1
	SchemeRuleMappingId *int64 `json:"SchemeRuleMappingId,omitempty" xml:"SchemeRuleMappingId,omitempty"`
	// Indicates whether the scoring item is deleted. A deleted item may be displayed as grayed out.
	//
	// example:
	//
	// true
	ScoreDeleted *bool `json:"ScoreDeleted,omitempty" xml:"ScoreDeleted,omitempty"`
	// The ID of the main scoring item.
	//
	// example:
	//
	// 1
	ScoreId *int64 `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	// The name of the main scoring item.
	//
	// example:
	//
	// 违规
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	// The score value.
	//
	// example:
	//
	// 1
	ScoreNum *float32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// The scoring method. This parameter is used for compatibility with the v4.0 protocol. Valid values: `0` (Add/deduct points when the rule is triggered) and `1` (Assign a one-time score when the rule is triggered).
	//
	// example:
	//
	// 1
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// The scoring trigger. Valid values: `0` (Score when a node is hit).
	//
	// example:
	//
	// 0
	ScoreRuleHitType *int32 `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	// The ID of the scoring subitem.
	//
	// example:
	//
	// 1
	ScoreSubId *int64 `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	// The name of the scoring subitem.
	//
	// example:
	//
	// 1
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	// Indicates whether to add or deduct points.
	//
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	// The sort order of the rule in the quality check dimension.
	//
	// example:
	//
	// 1
	SortIndex *int32 `json:"SortIndex,omitempty" xml:"SortIndex,omitempty"`
	// The start time. This value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1641277321000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule category. This parameter is used for compatibility with the v4.0 protocol. Valid values: `10` (General rule for quality check schemes) and `11` (Flow rule for quality check schemes).
	//
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the flowchart.
	//
	// example:
	//
	// 111111111
	TaskFlowId *int64 `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	// The type of the flowchart. This parameter is deprecated. The default value is `1`.
	//
	// example:
	//
	// 1
	TaskFlowType *int32 `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
	// A list of triggers.
	Triggers []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
	// The rule category. This parameter is used for compatibility with the v4.0 protocol. Valid values: `2` (Public opinion monitoring) and `3` (Business).
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the rule.
	//
	// example:
	//
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RuleInfo) String() string {
	return dara.Prettify(s)
}

func (s RuleInfo) GoString() string {
	return s.String()
}

func (s *RuleInfo) GetAutoReview() *int32 {
	return s.AutoReview
}

func (s *RuleInfo) GetBusinessCategoryNameList() []*string {
	return s.BusinessCategoryNameList
}

func (s *RuleInfo) GetCheckType() *int64 {
	return s.CheckType
}

func (s *RuleInfo) GetComments() *string {
	return s.Comments
}

func (s *RuleInfo) GetConfigType() *int32 {
	return s.ConfigType
}

func (s *RuleInfo) GetCreateEmpName() *string {
	return s.CreateEmpName
}

func (s *RuleInfo) GetCreateEmpid() *string {
	return s.CreateEmpid
}

func (s *RuleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RuleInfo) GetDeny() *int32 {
	return s.Deny
}

func (s *RuleInfo) GetDialogues() []*RuleTestDialogue {
	return s.Dialogues
}

func (s *RuleInfo) GetEffective() *int32 {
	return s.Effective
}

func (s *RuleInfo) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *RuleInfo) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *RuleInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *RuleInfo) GetExternalProperty() *int32 {
	return s.ExternalProperty
}

func (s *RuleInfo) GetFullCycle() *int32 {
	return s.FullCycle
}

func (s *RuleInfo) GetGraphFlow() interface{} {
	return s.GraphFlow
}

func (s *RuleInfo) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *RuleInfo) GetIsOnline() *int32 {
	return s.IsOnline
}

func (s *RuleInfo) GetLambda() *string {
	return s.Lambda
}

func (s *RuleInfo) GetLastUpdateEmpName() *string {
	return s.LastUpdateEmpName
}

func (s *RuleInfo) GetLastUpdateEmpid() *string {
	return s.LastUpdateEmpid
}

func (s *RuleInfo) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *RuleInfo) GetLevel() *int32 {
	return s.Level
}

func (s *RuleInfo) GetMeet() *int32 {
	return s.Meet
}

func (s *RuleInfo) GetModifyType() *int32 {
	return s.ModifyType
}

func (s *RuleInfo) GetName() *string {
	return s.Name
}

func (s *RuleInfo) GetOperationMode() *int32 {
	return s.OperationMode
}

func (s *RuleInfo) GetPreqRule() *RuleInfoPreqRule {
	return s.PreqRule
}

func (s *RuleInfo) GetQualityCheckType() *int32 {
	return s.QualityCheckType
}

func (s *RuleInfo) GetRid() *string {
	return s.Rid
}

func (s *RuleInfo) GetRuleCategoryName() *string {
	return s.RuleCategoryName
}

func (s *RuleInfo) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *RuleInfo) GetRuleType() *int32 {
	return s.RuleType
}

func (s *RuleInfo) GetSchemeCheckType() *SchemeCheckType {
	return s.SchemeCheckType
}

func (s *RuleInfo) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *RuleInfo) GetSchemeName() *string {
	return s.SchemeName
}

func (s *RuleInfo) GetSchemeRuleMappingId() *int64 {
	return s.SchemeRuleMappingId
}

func (s *RuleInfo) GetScoreDeleted() *bool {
	return s.ScoreDeleted
}

func (s *RuleInfo) GetScoreId() *int64 {
	return s.ScoreId
}

func (s *RuleInfo) GetScoreName() *string {
	return s.ScoreName
}

func (s *RuleInfo) GetScoreNum() *float32 {
	return s.ScoreNum
}

func (s *RuleInfo) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *RuleInfo) GetScoreRuleHitType() *int32 {
	return s.ScoreRuleHitType
}

func (s *RuleInfo) GetScoreSubId() *int64 {
	return s.ScoreSubId
}

func (s *RuleInfo) GetScoreSubName() *string {
	return s.ScoreSubName
}

func (s *RuleInfo) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *RuleInfo) GetSortIndex() *int32 {
	return s.SortIndex
}

func (s *RuleInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *RuleInfo) GetStatus() *int32 {
	return s.Status
}

func (s *RuleInfo) GetTargetType() *int32 {
	return s.TargetType
}

func (s *RuleInfo) GetTaskFlowId() *int64 {
	return s.TaskFlowId
}

func (s *RuleInfo) GetTaskFlowType() *int32 {
	return s.TaskFlowType
}

func (s *RuleInfo) GetTriggers() []*string {
	return s.Triggers
}

func (s *RuleInfo) GetType() *int32 {
	return s.Type
}

func (s *RuleInfo) GetWeight() *string {
	return s.Weight
}

func (s *RuleInfo) SetAutoReview(v int32) *RuleInfo {
	s.AutoReview = &v
	return s
}

func (s *RuleInfo) SetBusinessCategoryNameList(v []*string) *RuleInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *RuleInfo) SetCheckType(v int64) *RuleInfo {
	s.CheckType = &v
	return s
}

func (s *RuleInfo) SetComments(v string) *RuleInfo {
	s.Comments = &v
	return s
}

func (s *RuleInfo) SetConfigType(v int32) *RuleInfo {
	s.ConfigType = &v
	return s
}

func (s *RuleInfo) SetCreateEmpName(v string) *RuleInfo {
	s.CreateEmpName = &v
	return s
}

func (s *RuleInfo) SetCreateEmpid(v string) *RuleInfo {
	s.CreateEmpid = &v
	return s
}

func (s *RuleInfo) SetCreateTime(v string) *RuleInfo {
	s.CreateTime = &v
	return s
}

func (s *RuleInfo) SetDeny(v int32) *RuleInfo {
	s.Deny = &v
	return s
}

func (s *RuleInfo) SetDialogues(v []*RuleTestDialogue) *RuleInfo {
	s.Dialogues = v
	return s
}

func (s *RuleInfo) SetEffective(v int32) *RuleInfo {
	s.Effective = &v
	return s
}

func (s *RuleInfo) SetEffectiveEndTime(v string) *RuleInfo {
	s.EffectiveEndTime = &v
	return s
}

func (s *RuleInfo) SetEffectiveStartTime(v string) *RuleInfo {
	s.EffectiveStartTime = &v
	return s
}

func (s *RuleInfo) SetEndTime(v string) *RuleInfo {
	s.EndTime = &v
	return s
}

func (s *RuleInfo) SetExternalProperty(v int32) *RuleInfo {
	s.ExternalProperty = &v
	return s
}

func (s *RuleInfo) SetFullCycle(v int32) *RuleInfo {
	s.FullCycle = &v
	return s
}

func (s *RuleInfo) SetGraphFlow(v interface{}) *RuleInfo {
	s.GraphFlow = v
	return s
}

func (s *RuleInfo) SetIsDelete(v int32) *RuleInfo {
	s.IsDelete = &v
	return s
}

func (s *RuleInfo) SetIsOnline(v int32) *RuleInfo {
	s.IsOnline = &v
	return s
}

func (s *RuleInfo) SetLambda(v string) *RuleInfo {
	s.Lambda = &v
	return s
}

func (s *RuleInfo) SetLastUpdateEmpName(v string) *RuleInfo {
	s.LastUpdateEmpName = &v
	return s
}

func (s *RuleInfo) SetLastUpdateEmpid(v string) *RuleInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *RuleInfo) SetLastUpdateTime(v string) *RuleInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *RuleInfo) SetLevel(v int32) *RuleInfo {
	s.Level = &v
	return s
}

func (s *RuleInfo) SetMeet(v int32) *RuleInfo {
	s.Meet = &v
	return s
}

func (s *RuleInfo) SetModifyType(v int32) *RuleInfo {
	s.ModifyType = &v
	return s
}

func (s *RuleInfo) SetName(v string) *RuleInfo {
	s.Name = &v
	return s
}

func (s *RuleInfo) SetOperationMode(v int32) *RuleInfo {
	s.OperationMode = &v
	return s
}

func (s *RuleInfo) SetPreqRule(v *RuleInfoPreqRule) *RuleInfo {
	s.PreqRule = v
	return s
}

func (s *RuleInfo) SetQualityCheckType(v int32) *RuleInfo {
	s.QualityCheckType = &v
	return s
}

func (s *RuleInfo) SetRid(v string) *RuleInfo {
	s.Rid = &v
	return s
}

func (s *RuleInfo) SetRuleCategoryName(v string) *RuleInfo {
	s.RuleCategoryName = &v
	return s
}

func (s *RuleInfo) SetRuleScoreType(v int32) *RuleInfo {
	s.RuleScoreType = &v
	return s
}

func (s *RuleInfo) SetRuleType(v int32) *RuleInfo {
	s.RuleType = &v
	return s
}

func (s *RuleInfo) SetSchemeCheckType(v *SchemeCheckType) *RuleInfo {
	s.SchemeCheckType = v
	return s
}

func (s *RuleInfo) SetSchemeId(v int64) *RuleInfo {
	s.SchemeId = &v
	return s
}

func (s *RuleInfo) SetSchemeName(v string) *RuleInfo {
	s.SchemeName = &v
	return s
}

func (s *RuleInfo) SetSchemeRuleMappingId(v int64) *RuleInfo {
	s.SchemeRuleMappingId = &v
	return s
}

func (s *RuleInfo) SetScoreDeleted(v bool) *RuleInfo {
	s.ScoreDeleted = &v
	return s
}

func (s *RuleInfo) SetScoreId(v int64) *RuleInfo {
	s.ScoreId = &v
	return s
}

func (s *RuleInfo) SetScoreName(v string) *RuleInfo {
	s.ScoreName = &v
	return s
}

func (s *RuleInfo) SetScoreNum(v float32) *RuleInfo {
	s.ScoreNum = &v
	return s
}

func (s *RuleInfo) SetScoreNumType(v int32) *RuleInfo {
	s.ScoreNumType = &v
	return s
}

func (s *RuleInfo) SetScoreRuleHitType(v int32) *RuleInfo {
	s.ScoreRuleHitType = &v
	return s
}

func (s *RuleInfo) SetScoreSubId(v int64) *RuleInfo {
	s.ScoreSubId = &v
	return s
}

func (s *RuleInfo) SetScoreSubName(v string) *RuleInfo {
	s.ScoreSubName = &v
	return s
}

func (s *RuleInfo) SetScoreType(v int32) *RuleInfo {
	s.ScoreType = &v
	return s
}

func (s *RuleInfo) SetSortIndex(v int32) *RuleInfo {
	s.SortIndex = &v
	return s
}

func (s *RuleInfo) SetStartTime(v string) *RuleInfo {
	s.StartTime = &v
	return s
}

func (s *RuleInfo) SetStatus(v int32) *RuleInfo {
	s.Status = &v
	return s
}

func (s *RuleInfo) SetTargetType(v int32) *RuleInfo {
	s.TargetType = &v
	return s
}

func (s *RuleInfo) SetTaskFlowId(v int64) *RuleInfo {
	s.TaskFlowId = &v
	return s
}

func (s *RuleInfo) SetTaskFlowType(v int32) *RuleInfo {
	s.TaskFlowType = &v
	return s
}

func (s *RuleInfo) SetTriggers(v []*string) *RuleInfo {
	s.Triggers = v
	return s
}

func (s *RuleInfo) SetType(v int32) *RuleInfo {
	s.Type = &v
	return s
}

func (s *RuleInfo) SetWeight(v string) *RuleInfo {
	s.Weight = &v
	return s
}

func (s *RuleInfo) Validate() error {
	if s.Dialogues != nil {
		for _, item := range s.Dialogues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreqRule != nil {
		if err := s.PreqRule.Validate(); err != nil {
			return err
		}
	}
	if s.SchemeCheckType != nil {
		if err := s.SchemeCheckType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RuleInfoPreqRule struct {
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s RuleInfoPreqRule) String() string {
	return dara.Prettify(s)
}

func (s RuleInfoPreqRule) GoString() string {
	return s.String()
}

func (s *RuleInfoPreqRule) GetRid() *string {
	return s.Rid
}

func (s *RuleInfoPreqRule) SetRid(v string) *RuleInfoPreqRule {
	s.Rid = &v
	return s
}

func (s *RuleInfoPreqRule) Validate() error {
	return dara.Validate(s)
}
