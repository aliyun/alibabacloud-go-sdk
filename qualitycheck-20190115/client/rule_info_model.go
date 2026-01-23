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
	SetScoreNum(v int32) *RuleInfo
	GetScoreNum() *int32
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
	AutoReview               *int32              `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryNameList []*string           `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	CheckType                *int64              `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Comments                 *string             `json:"Comments,omitempty" xml:"Comments,omitempty"`
	ConfigType               *int32              `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	CreateEmpName            *string             `json:"CreateEmpName,omitempty" xml:"CreateEmpName,omitempty"`
	CreateEmpid              *string             `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	CreateTime               *string             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deny                     *int32              `json:"Deny,omitempty" xml:"Deny,omitempty"`
	Dialogues                []*RuleTestDialogue `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	Effective                *int32              `json:"Effective,omitempty" xml:"Effective,omitempty"`
	EffectiveEndTime         *string             `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	EffectiveStartTime       *string             `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	EndTime                  *string             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExternalProperty         *int32              `json:"ExternalProperty,omitempty" xml:"ExternalProperty,omitempty"`
	FullCycle                *int32              `json:"FullCycle,omitempty" xml:"FullCycle,omitempty"`
	GraphFlow                interface{}         `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	IsDelete                 *int32              `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	IsOnline                 *int32              `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	Lambda                   *string             `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	LastUpdateEmpName        *string             `json:"LastUpdateEmpName,omitempty" xml:"LastUpdateEmpName,omitempty"`
	LastUpdateEmpid          *string             `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	LastUpdateTime           *string             `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Level                    *int32              `json:"Level,omitempty" xml:"Level,omitempty"`
	Meet                     *int32              `json:"Meet,omitempty" xml:"Meet,omitempty"`
	ModifyType               *int32              `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	Name                     *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMode            *int32              `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	PreqRule                 *RuleInfoPreqRule   `json:"PreqRule,omitempty" xml:"PreqRule,omitempty" type:"Struct"`
	QualityCheckType         *int32              `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	Rid                      *string             `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleCategoryName         *string             `json:"RuleCategoryName,omitempty" xml:"RuleCategoryName,omitempty"`
	RuleScoreType            *int32              `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	RuleType                 *int32              `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SchemeCheckType          *SchemeCheckType    `json:"SchemeCheckType,omitempty" xml:"SchemeCheckType,omitempty"`
	SchemeId                 *int64              `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeName               *string             `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	SchemeRuleMappingId      *int64              `json:"SchemeRuleMappingId,omitempty" xml:"SchemeRuleMappingId,omitempty"`
	ScoreDeleted             *bool               `json:"ScoreDeleted,omitempty" xml:"ScoreDeleted,omitempty"`
	ScoreId                  *int64              `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName                *string             `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	ScoreNum                 *int32              `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType             *int32              `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType         *int32              `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreSubId               *int64              `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName             *string             `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	ScoreType                *int32              `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	SortIndex                *int32              `json:"SortIndex,omitempty" xml:"SortIndex,omitempty"`
	StartTime                *string             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                   *int32              `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetType               *int32              `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskFlowId               *int64              `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowType             *int32              `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
	Triggers                 []*string           `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
	Type                     *int32              `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                   *string             `json:"Weight,omitempty" xml:"Weight,omitempty"`
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

func (s *RuleInfo) GetScoreNum() *int32 {
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

func (s *RuleInfo) SetScoreNum(v int32) *RuleInfo {
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
