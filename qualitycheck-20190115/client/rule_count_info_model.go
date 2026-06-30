// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuleCountInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReview(v int32) *RuleCountInfo
	GetAutoReview() *int32
	SetBusinessCategoryBasicInfoList(v []*BusinessCategoryBasicInfo) *RuleCountInfo
	GetBusinessCategoryBasicInfoList() []*BusinessCategoryBasicInfo
	SetBusinessCategoryNameList(v []*string) *RuleCountInfo
	GetBusinessCategoryNameList() []*string
	SetBusinessRange(v []*int32) *RuleCountInfo
	GetBusinessRange() []*int32
	SetCheckNumber(v int64) *RuleCountInfo
	GetCheckNumber() *int64
	SetComments(v string) *RuleCountInfo
	GetComments() *string
	SetCreateEmpName(v string) *RuleCountInfo
	GetCreateEmpName() *string
	SetCreateEmpid(v string) *RuleCountInfo
	GetCreateEmpid() *string
	SetCreateTime(v string) *RuleCountInfo
	GetCreateTime() *string
	SetDeny(v int32) *RuleCountInfo
	GetDeny() *int32
	SetEffective(v int32) *RuleCountInfo
	GetEffective() *int32
	SetEffectiveEndTime(v string) *RuleCountInfo
	GetEffectiveEndTime() *string
	SetEffectiveStartTime(v string) *RuleCountInfo
	GetEffectiveStartTime() *string
	SetEndTime(v string) *RuleCountInfo
	GetEndTime() *string
	SetFullCycle(v int32) *RuleCountInfo
	GetFullCycle() *int32
	SetGraphFlow(v interface{}) *RuleCountInfo
	GetGraphFlow() interface{}
	SetHitNumber(v int64) *RuleCountInfo
	GetHitNumber() *int64
	SetHitRate(v float32) *RuleCountInfo
	GetHitRate() *float32
	SetHitRealViolationRate(v float32) *RuleCountInfo
	GetHitRealViolationRate() *float32
	SetIsDelete(v int32) *RuleCountInfo
	GetIsDelete() *int32
	SetIsSelect(v bool) *RuleCountInfo
	GetIsSelect() *bool
	SetJobName(v string) *RuleCountInfo
	GetJobName() *string
	SetLastUpdateEmpName(v string) *RuleCountInfo
	GetLastUpdateEmpName() *string
	SetLastUpdateEmpid(v string) *RuleCountInfo
	GetLastUpdateEmpid() *string
	SetLastUpdateTime(v string) *RuleCountInfo
	GetLastUpdateTime() *string
	SetName(v string) *RuleCountInfo
	GetName() *string
	SetOperationMode(v int32) *RuleCountInfo
	GetOperationMode() *int32
	SetPreReviewNumber(v int64) *RuleCountInfo
	GetPreReviewNumber() *int64
	SetProblemNumber(v int64) *RuleCountInfo
	GetProblemNumber() *int64
	SetQualityCheckType(v int32) *RuleCountInfo
	GetQualityCheckType() *int32
	SetRealViolationNumber(v int32) *RuleCountInfo
	GetRealViolationNumber() *int32
	SetReviewAccuracyRate(v float32) *RuleCountInfo
	GetReviewAccuracyRate() *float32
	SetReviewNumber(v int64) *RuleCountInfo
	GetReviewNumber() *int64
	SetReviewRate(v float32) *RuleCountInfo
	GetReviewRate() *float32
	SetReviewStatusName(v string) *RuleCountInfo
	GetReviewStatusName() *string
	SetRid(v int64) *RuleCountInfo
	GetRid() *int64
	SetRuleScoreSingleType(v int32) *RuleCountInfo
	GetRuleScoreSingleType() *int32
	SetRuleScoreType(v int32) *RuleCountInfo
	GetRuleScoreType() *int32
	SetRuleType(v int32) *RuleCountInfo
	GetRuleType() *int32
	SetScoreSubId(v int64) *RuleCountInfo
	GetScoreSubId() *int64
	SetStartTime(v string) *RuleCountInfo
	GetStartTime() *string
	SetStatus(v int32) *RuleCountInfo
	GetStatus() *int32
	SetTargetType(v int32) *RuleCountInfo
	GetTargetType() *int32
	SetType(v int32) *RuleCountInfo
	GetType() *int32
	SetTypeName(v string) *RuleCountInfo
	GetTypeName() *string
	SetUnReviewNumber(v int64) *RuleCountInfo
	GetUnReviewNumber() *int64
	SetUserGroup(v string) *RuleCountInfo
	GetUserGroup() *string
}

type RuleCountInfo struct {
	// The review option.
	//
	// - 1: Manual review
	//
	// - 3: Automatic review
	//
	// example:
	//
	// 1
	AutoReview *int32 `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	// A list of basic information about business categories.
	BusinessCategoryBasicInfoList []*BusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfoList,omitempty" xml:"BusinessCategoryBasicInfoList,omitempty" type:"Repeated"`
	// A list of business category names.
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	// The business scope.
	BusinessRange []*int32 `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty" type:"Repeated"`
	// The number of checked sessions.
	//
	// example:
	//
	// 1
	CheckNumber *int64 `json:"CheckNumber,omitempty" xml:"CheckNumber,omitempty"`
	// The comments.
	//
	// example:
	//
	// 测试
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
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
	// The creation time.
	//
	// example:
	//
	// 1615133575000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// You can set this parameter to 1 to reject rule edits, or to 0 or null to allow rule edits.
	//
	// example:
	//
	// 1
	Deny *int32 `json:"Deny,omitempty" xml:"Deny,omitempty"`
	// Indicates if the rule is active. \\`0\\`: No. \\`1\\`: Yes.
	//
	// example:
	//
	// 1
	Effective *int32 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The effective end time.
	//
	// example:
	//
	// 1662685868850
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The effective start time.
	//
	// example:
	//
	// 1662685868850
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1650092585176
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates if the rule is active for the entire lifecycle. \\`0\\`: No. \\`1\\`: Yes.
	//
	// example:
	//
	// 1
	FullCycle *int32 `json:"FullCycle,omitempty" xml:"FullCycle,omitempty"`
	// The flow canvas, in JSON format.
	//
	// example:
	//
	// 较复杂，忽略
	GraphFlow interface{} `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	// The number of hit sessions.
	//
	// example:
	//
	// 1
	HitNumber *int64 `json:"HitNumber,omitempty" xml:"HitNumber,omitempty"`
	// The hit rate.
	//
	// example:
	//
	// 1
	HitRate *float32 `json:"HitRate,omitempty" xml:"HitRate,omitempty"`
	// The hit rate of actual violations.
	//
	// example:
	//
	// 1
	HitRealViolationRate *float32 `json:"HitRealViolationRate,omitempty" xml:"HitRealViolationRate,omitempty"`
	// Indicates if the rule is deleted.
	//
	// example:
	//
	// 1
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Indicates if the rule is selected.
	//
	// example:
	//
	// true
	IsSelect *bool `json:"IsSelect,omitempty" xml:"IsSelect,omitempty"`
	// The scheduled task name.
	//
	// example:
	//
	// job-1-20221012-105943
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The name of the user who last updated the rule.
	//
	// example:
	//
	// 张三
	LastUpdateEmpName *string `json:"LastUpdateEmpName,omitempty" xml:"LastUpdateEmpName,omitempty"`
	// The employee ID of the user who last updated the rule.
	//
	// example:
	//
	// 1
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 1648200901000
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The rule name.
	//
	// example:
	//
	// 0801转封装测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation mode.
	//
	// example:
	//
	// 1
	OperationMode *int32 `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	// The number of items pending review.
	//
	// example:
	//
	// 1
	PreReviewNumber *int64 `json:"PreReviewNumber,omitempty" xml:"PreReviewNumber,omitempty"`
	// The number of problems found.
	//
	// example:
	//
	// 1
	ProblemNumber *int64 `json:"ProblemNumber,omitempty" xml:"ProblemNumber,omitempty"`
	// The quality check type.
	//
	// - 0: Offline
	//
	// - 1: Real-time
	//
	// example:
	//
	// 1
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// The number of actual violations after review.
	//
	// example:
	//
	// 1
	RealViolationNumber *int32 `json:"RealViolationNumber,omitempty" xml:"RealViolationNumber,omitempty"`
	// The review accuracy rate.
	//
	// example:
	//
	// 1
	ReviewAccuracyRate *float32 `json:"ReviewAccuracyRate,omitempty" xml:"ReviewAccuracyRate,omitempty"`
	// The number of reviewed items.
	//
	// example:
	//
	// 1
	ReviewNumber *int64 `json:"ReviewNumber,omitempty" xml:"ReviewNumber,omitempty"`
	// The review rate.
	//
	// example:
	//
	// 1
	ReviewRate *float32 `json:"ReviewRate,omitempty" xml:"ReviewRate,omitempty"`
	// The review status name.
	//
	// example:
	//
	// 通过
	ReviewStatusName *string `json:"ReviewStatusName,omitempty" xml:"ReviewStatusName,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 123
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// The scoring type. \\`8\\`: No score is set.
	//
	// example:
	//
	// 1
	RuleScoreSingleType *int32 `json:"RuleScoreSingleType,omitempty" xml:"RuleScoreSingleType,omitempty"`
	// Specifies if a score is calculated.
	//
	// - 1: No
	//
	// - 3: Yes
	//
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// The rule origin. For example, a built-in rule or a user-created rule.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The sub-item ID.
	//
	// example:
	//
	// 1
	ScoreSubId *int64 `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1662685868850
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule.
	//
	// - 0: Pending
	//
	// - 1: Active
	//
	// - 2: Expired
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule category.
	//
	// - 10: Standard
	//
	// - 11: Flow-based
	//
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The rule type.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type name.
	//
	// example:
	//
	// 全部类别
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The number of unreviewed items.
	//
	// example:
	//
	// 1
	UnReviewNumber *int64 `json:"UnReviewNumber,omitempty" xml:"UnReviewNumber,omitempty"`
	// The user group.
	//
	// example:
	//
	// xxxx
	UserGroup *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s RuleCountInfo) String() string {
	return dara.Prettify(s)
}

func (s RuleCountInfo) GoString() string {
	return s.String()
}

func (s *RuleCountInfo) GetAutoReview() *int32 {
	return s.AutoReview
}

func (s *RuleCountInfo) GetBusinessCategoryBasicInfoList() []*BusinessCategoryBasicInfo {
	return s.BusinessCategoryBasicInfoList
}

func (s *RuleCountInfo) GetBusinessCategoryNameList() []*string {
	return s.BusinessCategoryNameList
}

func (s *RuleCountInfo) GetBusinessRange() []*int32 {
	return s.BusinessRange
}

func (s *RuleCountInfo) GetCheckNumber() *int64 {
	return s.CheckNumber
}

func (s *RuleCountInfo) GetComments() *string {
	return s.Comments
}

func (s *RuleCountInfo) GetCreateEmpName() *string {
	return s.CreateEmpName
}

func (s *RuleCountInfo) GetCreateEmpid() *string {
	return s.CreateEmpid
}

func (s *RuleCountInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RuleCountInfo) GetDeny() *int32 {
	return s.Deny
}

func (s *RuleCountInfo) GetEffective() *int32 {
	return s.Effective
}

func (s *RuleCountInfo) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *RuleCountInfo) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *RuleCountInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *RuleCountInfo) GetFullCycle() *int32 {
	return s.FullCycle
}

func (s *RuleCountInfo) GetGraphFlow() interface{} {
	return s.GraphFlow
}

func (s *RuleCountInfo) GetHitNumber() *int64 {
	return s.HitNumber
}

func (s *RuleCountInfo) GetHitRate() *float32 {
	return s.HitRate
}

func (s *RuleCountInfo) GetHitRealViolationRate() *float32 {
	return s.HitRealViolationRate
}

func (s *RuleCountInfo) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *RuleCountInfo) GetIsSelect() *bool {
	return s.IsSelect
}

func (s *RuleCountInfo) GetJobName() *string {
	return s.JobName
}

func (s *RuleCountInfo) GetLastUpdateEmpName() *string {
	return s.LastUpdateEmpName
}

func (s *RuleCountInfo) GetLastUpdateEmpid() *string {
	return s.LastUpdateEmpid
}

func (s *RuleCountInfo) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *RuleCountInfo) GetName() *string {
	return s.Name
}

func (s *RuleCountInfo) GetOperationMode() *int32 {
	return s.OperationMode
}

func (s *RuleCountInfo) GetPreReviewNumber() *int64 {
	return s.PreReviewNumber
}

func (s *RuleCountInfo) GetProblemNumber() *int64 {
	return s.ProblemNumber
}

func (s *RuleCountInfo) GetQualityCheckType() *int32 {
	return s.QualityCheckType
}

func (s *RuleCountInfo) GetRealViolationNumber() *int32 {
	return s.RealViolationNumber
}

func (s *RuleCountInfo) GetReviewAccuracyRate() *float32 {
	return s.ReviewAccuracyRate
}

func (s *RuleCountInfo) GetReviewNumber() *int64 {
	return s.ReviewNumber
}

func (s *RuleCountInfo) GetReviewRate() *float32 {
	return s.ReviewRate
}

func (s *RuleCountInfo) GetReviewStatusName() *string {
	return s.ReviewStatusName
}

func (s *RuleCountInfo) GetRid() *int64 {
	return s.Rid
}

func (s *RuleCountInfo) GetRuleScoreSingleType() *int32 {
	return s.RuleScoreSingleType
}

func (s *RuleCountInfo) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *RuleCountInfo) GetRuleType() *int32 {
	return s.RuleType
}

func (s *RuleCountInfo) GetScoreSubId() *int64 {
	return s.ScoreSubId
}

func (s *RuleCountInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *RuleCountInfo) GetStatus() *int32 {
	return s.Status
}

func (s *RuleCountInfo) GetTargetType() *int32 {
	return s.TargetType
}

func (s *RuleCountInfo) GetType() *int32 {
	return s.Type
}

func (s *RuleCountInfo) GetTypeName() *string {
	return s.TypeName
}

func (s *RuleCountInfo) GetUnReviewNumber() *int64 {
	return s.UnReviewNumber
}

func (s *RuleCountInfo) GetUserGroup() *string {
	return s.UserGroup
}

func (s *RuleCountInfo) SetAutoReview(v int32) *RuleCountInfo {
	s.AutoReview = &v
	return s
}

func (s *RuleCountInfo) SetBusinessCategoryBasicInfoList(v []*BusinessCategoryBasicInfo) *RuleCountInfo {
	s.BusinessCategoryBasicInfoList = v
	return s
}

func (s *RuleCountInfo) SetBusinessCategoryNameList(v []*string) *RuleCountInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *RuleCountInfo) SetBusinessRange(v []*int32) *RuleCountInfo {
	s.BusinessRange = v
	return s
}

func (s *RuleCountInfo) SetCheckNumber(v int64) *RuleCountInfo {
	s.CheckNumber = &v
	return s
}

func (s *RuleCountInfo) SetComments(v string) *RuleCountInfo {
	s.Comments = &v
	return s
}

func (s *RuleCountInfo) SetCreateEmpName(v string) *RuleCountInfo {
	s.CreateEmpName = &v
	return s
}

func (s *RuleCountInfo) SetCreateEmpid(v string) *RuleCountInfo {
	s.CreateEmpid = &v
	return s
}

func (s *RuleCountInfo) SetCreateTime(v string) *RuleCountInfo {
	s.CreateTime = &v
	return s
}

func (s *RuleCountInfo) SetDeny(v int32) *RuleCountInfo {
	s.Deny = &v
	return s
}

func (s *RuleCountInfo) SetEffective(v int32) *RuleCountInfo {
	s.Effective = &v
	return s
}

func (s *RuleCountInfo) SetEffectiveEndTime(v string) *RuleCountInfo {
	s.EffectiveEndTime = &v
	return s
}

func (s *RuleCountInfo) SetEffectiveStartTime(v string) *RuleCountInfo {
	s.EffectiveStartTime = &v
	return s
}

func (s *RuleCountInfo) SetEndTime(v string) *RuleCountInfo {
	s.EndTime = &v
	return s
}

func (s *RuleCountInfo) SetFullCycle(v int32) *RuleCountInfo {
	s.FullCycle = &v
	return s
}

func (s *RuleCountInfo) SetGraphFlow(v interface{}) *RuleCountInfo {
	s.GraphFlow = v
	return s
}

func (s *RuleCountInfo) SetHitNumber(v int64) *RuleCountInfo {
	s.HitNumber = &v
	return s
}

func (s *RuleCountInfo) SetHitRate(v float32) *RuleCountInfo {
	s.HitRate = &v
	return s
}

func (s *RuleCountInfo) SetHitRealViolationRate(v float32) *RuleCountInfo {
	s.HitRealViolationRate = &v
	return s
}

func (s *RuleCountInfo) SetIsDelete(v int32) *RuleCountInfo {
	s.IsDelete = &v
	return s
}

func (s *RuleCountInfo) SetIsSelect(v bool) *RuleCountInfo {
	s.IsSelect = &v
	return s
}

func (s *RuleCountInfo) SetJobName(v string) *RuleCountInfo {
	s.JobName = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateEmpName(v string) *RuleCountInfo {
	s.LastUpdateEmpName = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateEmpid(v string) *RuleCountInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateTime(v string) *RuleCountInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *RuleCountInfo) SetName(v string) *RuleCountInfo {
	s.Name = &v
	return s
}

func (s *RuleCountInfo) SetOperationMode(v int32) *RuleCountInfo {
	s.OperationMode = &v
	return s
}

func (s *RuleCountInfo) SetPreReviewNumber(v int64) *RuleCountInfo {
	s.PreReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetProblemNumber(v int64) *RuleCountInfo {
	s.ProblemNumber = &v
	return s
}

func (s *RuleCountInfo) SetQualityCheckType(v int32) *RuleCountInfo {
	s.QualityCheckType = &v
	return s
}

func (s *RuleCountInfo) SetRealViolationNumber(v int32) *RuleCountInfo {
	s.RealViolationNumber = &v
	return s
}

func (s *RuleCountInfo) SetReviewAccuracyRate(v float32) *RuleCountInfo {
	s.ReviewAccuracyRate = &v
	return s
}

func (s *RuleCountInfo) SetReviewNumber(v int64) *RuleCountInfo {
	s.ReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetReviewRate(v float32) *RuleCountInfo {
	s.ReviewRate = &v
	return s
}

func (s *RuleCountInfo) SetReviewStatusName(v string) *RuleCountInfo {
	s.ReviewStatusName = &v
	return s
}

func (s *RuleCountInfo) SetRid(v int64) *RuleCountInfo {
	s.Rid = &v
	return s
}

func (s *RuleCountInfo) SetRuleScoreSingleType(v int32) *RuleCountInfo {
	s.RuleScoreSingleType = &v
	return s
}

func (s *RuleCountInfo) SetRuleScoreType(v int32) *RuleCountInfo {
	s.RuleScoreType = &v
	return s
}

func (s *RuleCountInfo) SetRuleType(v int32) *RuleCountInfo {
	s.RuleType = &v
	return s
}

func (s *RuleCountInfo) SetScoreSubId(v int64) *RuleCountInfo {
	s.ScoreSubId = &v
	return s
}

func (s *RuleCountInfo) SetStartTime(v string) *RuleCountInfo {
	s.StartTime = &v
	return s
}

func (s *RuleCountInfo) SetStatus(v int32) *RuleCountInfo {
	s.Status = &v
	return s
}

func (s *RuleCountInfo) SetTargetType(v int32) *RuleCountInfo {
	s.TargetType = &v
	return s
}

func (s *RuleCountInfo) SetType(v int32) *RuleCountInfo {
	s.Type = &v
	return s
}

func (s *RuleCountInfo) SetTypeName(v string) *RuleCountInfo {
	s.TypeName = &v
	return s
}

func (s *RuleCountInfo) SetUnReviewNumber(v int64) *RuleCountInfo {
	s.UnReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetUserGroup(v string) *RuleCountInfo {
	s.UserGroup = &v
	return s
}

func (s *RuleCountInfo) Validate() error {
	if s.BusinessCategoryBasicInfoList != nil {
		for _, item := range s.BusinessCategoryBasicInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
