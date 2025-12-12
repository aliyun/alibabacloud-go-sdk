// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskAssignRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTaskAssignRulesResponseBody
	GetCode() *string
	SetCount(v int32) *ListTaskAssignRulesResponseBody
	GetCount() *int32
	SetData(v *ListTaskAssignRulesResponseBodyData) *ListTaskAssignRulesResponseBody
	GetData() *ListTaskAssignRulesResponseBodyData
	SetMessage(v string) *ListTaskAssignRulesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTaskAssignRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskAssignRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTaskAssignRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskAssignRulesResponseBody
	GetSuccess() *bool
}

type ListTaskAssignRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 23
	Count *int32                               `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListTaskAssignRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskAssignRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTaskAssignRulesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListTaskAssignRulesResponseBody) GetData() *ListTaskAssignRulesResponseBodyData {
	return s.Data
}

func (s *ListTaskAssignRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTaskAssignRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskAssignRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskAssignRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskAssignRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskAssignRulesResponseBody) SetCode(v string) *ListTaskAssignRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetCount(v int32) *ListTaskAssignRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetData(v *ListTaskAssignRulesResponseBodyData) *ListTaskAssignRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetMessage(v string) *ListTaskAssignRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageNumber(v int32) *ListTaskAssignRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageSize(v int32) *ListTaskAssignRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetRequestId(v string) *ListTaskAssignRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetSuccess(v bool) *ListTaskAssignRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyData struct {
	TaskAssignRuleInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo `json:"TaskAssignRuleInfo,omitempty" xml:"TaskAssignRuleInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyData) GetTaskAssignRuleInfo() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	return s.TaskAssignRuleInfo
}

func (s *ListTaskAssignRulesResponseBodyData) SetTaskAssignRuleInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) *ListTaskAssignRulesResponseBodyData {
	s.TaskAssignRuleInfo = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyData) Validate() error {
	if s.TaskAssignRuleInfo != nil {
		for _, item := range s.TaskAssignRuleInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo struct {
	Agents *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Struct"`
	// example:
	//
	// XX
	AgentsStr *string `json:"AgentsStr,omitempty" xml:"AgentsStr,omitempty"`
	// example:
	//
	// 0
	AssignmentType *int32 `json:"AssignmentType,omitempty" xml:"AssignmentType,omitempty"`
	// example:
	//
	// 39600
	CallTimeEnd *int64 `json:"CallTimeEnd,omitempty" xml:"CallTimeEnd,omitempty"`
	// example:
	//
	// 39600
	CallTimeStart *int64 `json:"CallTimeStart,omitempty" xml:"CallTimeStart,omitempty"`
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 2019-07-12T14:47Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 400
	DurationMax *int32 `json:"DurationMax,omitempty" xml:"DurationMax,omitempty"`
	// example:
	//
	// 100
	DurationMin *int32 `json:"DurationMin,omitempty" xml:"DurationMin,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1
	Priority  *int32                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Reviewers *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers `json:"Reviewers,omitempty" xml:"Reviewers,omitempty" type:"Struct"`
	// example:
	//
	// 23
	RuleId       *int64                                                             `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName     *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Rules        *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules        `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	SamplingMode *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode `json:"SamplingMode,omitempty" xml:"SamplingMode,omitempty" type:"Struct"`
	SkillGroups  *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups  `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Struct"`
	// example:
	//
	// XX
	SkillGroupsStr *string `json:"SkillGroupsStr,omitempty" xml:"SkillGroupsStr,omitempty"`
	// example:
	//
	// 2019-07-12T14:47Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetAgents() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents {
	return s.Agents
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetAgentsStr() *string {
	return s.AgentsStr
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetAssignmentType() *int32 {
	return s.AssignmentType
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetCallTimeEnd() *int64 {
	return s.CallTimeEnd
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetCallTimeStart() *int64 {
	return s.CallTimeStart
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetCallType() *int32 {
	return s.CallType
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetDurationMax() *int32 {
	return s.DurationMax
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetDurationMin() *int32 {
	return s.DurationMin
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetEnabled() *int32 {
	return s.Enabled
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetReviewers() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers {
	return s.Reviewers
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetRules() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules {
	return s.Rules
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetSamplingMode() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	return s.SamplingMode
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetSkillGroups() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups {
	return s.SkillGroups
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetSkillGroupsStr() *string {
	return s.SkillGroupsStr
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Agents = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgentsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AgentsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAssignmentType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AssignmentType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeEnd(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeEnd = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeStart(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeStart = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCreateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMax(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMax = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMin(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMin = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetEnabled(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Enabled = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetPriority(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Priority = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetReviewers(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Reviewers = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleId(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRules(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Rules = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSamplingMode(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SamplingMode = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroups(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroups = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroupsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroupsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetUpdateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) Validate() error {
	if s.Agents != nil {
		if err := s.Agents.Validate(); err != nil {
			return err
		}
	}
	if s.Reviewers != nil {
		if err := s.Reviewers.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	if s.SamplingMode != nil {
		if err := s.SamplingMode.Validate(); err != nil {
			return err
		}
	}
	if s.SkillGroups != nil {
		if err := s.SkillGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents struct {
	Agent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) GetAgent() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	return s.Agent
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) SetAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents {
	s.Agent = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) Validate() error {
	if s.Agent != nil {
		for _, item := range s.Agent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent struct {
	// example:
	//
	// 202526561358712105
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) GetAgentId() *string {
	return s.AgentId
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) GetAgentName() *string {
	return s.AgentName
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) Validate() error {
	return dara.Validate(s)
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers struct {
	Reviewer []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer `json:"Reviewer,omitempty" xml:"Reviewer,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) GetReviewer() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	return s.Reviewer
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) SetReviewer(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers {
	s.Reviewer = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) Validate() error {
	if s.Reviewer != nil {
		for _, item := range s.Reviewer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer struct {
	// example:
	//
	// 2337235457340978
	ReviewerId   *string `json:"ReviewerId,omitempty" xml:"ReviewerId,omitempty"`
	ReviewerName *string `json:"ReviewerName,omitempty" xml:"ReviewerName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) GetReviewerId() *string {
	return s.ReviewerId
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) GetReviewerName() *string {
	return s.ReviewerName
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) Validate() error {
	return dara.Validate(s)
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules struct {
	RuleBasicInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) GetRuleBasicInfo() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	return s.RuleBasicInfo
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) SetRuleBasicInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules {
	s.RuleBasicInfo = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) Validate() error {
	if s.RuleBasicInfo != nil {
		for _, item := range s.RuleBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2312
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) GetName() *string {
	return s.Name
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) GetRid() *string {
	return s.Rid
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Name = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetRid(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) Validate() error {
	return dara.Validate(s)
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode struct {
	// example:
	//
	// 60
	AnyNumberOfDraws *int32 `json:"AnyNumberOfDraws,omitempty" xml:"AnyNumberOfDraws,omitempty"`
	// example:
	//
	// true
	Designated *bool `json:"Designated,omitempty" xml:"Designated,omitempty"`
	// example:
	//
	// 0
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 30
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 20
	NumberOfDraws *int32 `json:"NumberOfDraws,omitempty" xml:"NumberOfDraws,omitempty"`
	// example:
	//
	// 0.1
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// example:
	//
	// 5
	RandomInspectionNumber *int32                                                                               `json:"RandomInspectionNumber,omitempty" xml:"RandomInspectionNumber,omitempty"`
	SamplingModeAgents     *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents `json:"SamplingModeAgents,omitempty" xml:"SamplingModeAgents,omitempty" type:"Struct"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetAnyNumberOfDraws() *int32 {
	return s.AnyNumberOfDraws
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetDesignated() *bool {
	return s.Designated
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetDimension() *int32 {
	return s.Dimension
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetLimit() *int32 {
	return s.Limit
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetNumberOfDraws() *int32 {
	return s.NumberOfDraws
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetProportion() *float32 {
	return s.Proportion
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetRandomInspectionNumber() *int32 {
	return s.RandomInspectionNumber
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GetSamplingModeAgents() *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents {
	return s.SamplingModeAgents
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetAnyNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.AnyNumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDesignated(v bool) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Designated = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDimension(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Dimension = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetLimit(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Limit = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.NumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetProportion(v float32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Proportion = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetRandomInspectionNumber(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.RandomInspectionNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetSamplingModeAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.SamplingModeAgents = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) Validate() error {
	if s.SamplingModeAgents != nil {
		if err := s.SamplingModeAgents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents struct {
	SamplingModeAgent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent `json:"SamplingModeAgent,omitempty" xml:"SamplingModeAgent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) GetSamplingModeAgent() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	return s.SamplingModeAgent
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) SetSamplingModeAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents {
	s.SamplingModeAgent = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) Validate() error {
	if s.SamplingModeAgent != nil {
		for _, item := range s.SamplingModeAgent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent struct {
	// example:
	//
	// 123
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// zhangsan
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) GetAgentId() *string {
	return s.AgentId
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) GetAgentName() *string {
	return s.AgentName
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) Validate() error {
	return dara.Validate(s)
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups struct {
	SkillGroup []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) GetSkillGroup() []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	return s.SkillGroup
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) SetSkillGroup(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups {
	s.SkillGroup = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) Validate() error {
	if s.SkillGroup != nil {
		for _, item := range s.SkillGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup struct {
	// example:
	//
	// XXX
	SkillId   *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) GetSkillId() *string {
	return s.SkillId
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) GetSkillName() *string {
	return s.SkillName
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) Validate() error {
	return dara.Validate(s)
}
