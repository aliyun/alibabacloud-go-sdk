// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachScriptPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListAICoachScriptPageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAICoachScriptPageResponseBody
	GetErrorMessage() *string
	SetList(v []*ListAICoachScriptPageResponseBodyList) *ListAICoachScriptPageResponseBody
	GetList() []*ListAICoachScriptPageResponseBodyList
	SetRequestId(v string) *ListAICoachScriptPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAICoachScriptPageResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAICoachScriptPageResponseBody
	GetTotal() *int32
}

type ListAICoachScriptPageResponseBody struct {
	// example:
	//
	// PARAM_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*ListAICoachScriptPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAICoachScriptPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAICoachScriptPageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAICoachScriptPageResponseBody) GetList() []*ListAICoachScriptPageResponseBodyList {
	return s.List
}

func (s *ListAICoachScriptPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAICoachScriptPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAICoachScriptPageResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAICoachScriptPageResponseBody) SetErrorCode(v string) *ListAICoachScriptPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetErrorMessage(v string) *ListAICoachScriptPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetList(v []*ListAICoachScriptPageResponseBodyList) *ListAICoachScriptPageResponseBody {
	s.List = v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetRequestId(v string) *ListAICoachScriptPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetSuccess(v bool) *ListAICoachScriptPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetTotal(v int32) *ListAICoachScriptPageResponseBody {
	s.Total = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyList struct {
	// example:
	//
	// true
	AppendQuestionFlag *string `json:"appendQuestionFlag,omitempty" xml:"appendQuestionFlag,omitempty"`
	// example:
	//
	// point
	AssessmentScope  *string                                                `json:"assessmentScope,omitempty" xml:"assessmentScope,omitempty"`
	ClosingRemarks   *string                                                `json:"closingRemarks,omitempty" xml:"closingRemarks,omitempty"`
	CompleteStrategy *ListAICoachScriptPageResponseBodyListCompleteStrategy `json:"completeStrategy,omitempty" xml:"completeStrategy,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx/cover.png
	CoverUrl         *string                                                  `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CustomReplyRules []*ListAICoachScriptPageResponseBodyListCustomReplyRules `json:"customReplyRules,omitempty" xml:"customReplyRules,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DialogueTextFlag *bool `json:"dialogueTextFlag,omitempty" xml:"dialogueTextFlag,omitempty"`
	// example:
	//
	// true
	DialogueTipFlag *bool `json:"dialogueTipFlag,omitempty" xml:"dialogueTipFlag,omitempty"`
	// example:
	//
	// true
	EvaluateReportFlag *bool              `json:"evaluateReportFlag,omitempty" xml:"evaluateReportFlag,omitempty"`
	Expressiveness     map[string]*string `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	// example:
	//
	// https://xxx.gif
	GifDynamicUrl *string `json:"gifDynamicUrl,omitempty" xml:"gifDynamicUrl,omitempty"`
	// example:
	//
	// https://xxx.gif
	GifStaticUrl *string `json:"gifStaticUrl,omitempty" xml:"gifStaticUrl,omitempty"`
	// example:
	//
	// 2024-12-25 14:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-12-25 14:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// student
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// example:
	//
	// 4
	InteractionType *string `json:"interactionType,omitempty" xml:"interactionType,omitempty"`
	Introduce       *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// example:
	//
	// prod-ydsf
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	// example:
	//
	// true
	OrderAckFlag       *bool                                                      `json:"orderAckFlag,omitempty" xml:"orderAckFlag,omitempty"`
	SampleDialogueList []*ListAICoachScriptPageResponseBodyListSampleDialogueList `json:"sampleDialogueList,omitempty" xml:"sampleDialogueList,omitempty" type:"Repeated"`
	ScoreConfig        *ListAICoachScriptPageResponseBodyListScoreConfig          `json:"scoreConfig,omitempty" xml:"scoreConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ScriptRecordId     *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	SparringTipContent *string `json:"sparringTipContent,omitempty" xml:"sparringTipContent,omitempty"`
	SparringTipTitle   *string `json:"sparringTipTitle,omitempty" xml:"sparringTipTitle,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	StudentThinkTimeFlag *bool `json:"studentThinkTimeFlag,omitempty" xml:"studentThinkTimeFlag,omitempty"`
	// example:
	//
	// 1
	Type    *int32                                        `json:"type,omitempty" xml:"type,omitempty"`
	Weights *ListAICoachScriptPageResponseBodyListWeights `json:"weights,omitempty" xml:"weights,omitempty" type:"Struct"`
}

func (s ListAICoachScriptPageResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyList) GetAppendQuestionFlag() *string {
	return s.AppendQuestionFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetAssessmentScope() *string {
	return s.AssessmentScope
}

func (s *ListAICoachScriptPageResponseBodyList) GetClosingRemarks() *string {
	return s.ClosingRemarks
}

func (s *ListAICoachScriptPageResponseBodyList) GetCompleteStrategy() *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	return s.CompleteStrategy
}

func (s *ListAICoachScriptPageResponseBodyList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListAICoachScriptPageResponseBodyList) GetCustomReplyRules() []*ListAICoachScriptPageResponseBodyListCustomReplyRules {
	return s.CustomReplyRules
}

func (s *ListAICoachScriptPageResponseBodyList) GetDialogueTextFlag() *bool {
	return s.DialogueTextFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetDialogueTipFlag() *bool {
	return s.DialogueTipFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetEvaluateReportFlag() *bool {
	return s.EvaluateReportFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetExpressiveness() map[string]*string {
	return s.Expressiveness
}

func (s *ListAICoachScriptPageResponseBodyList) GetGifDynamicUrl() *string {
	return s.GifDynamicUrl
}

func (s *ListAICoachScriptPageResponseBodyList) GetGifStaticUrl() *string {
	return s.GifStaticUrl
}

func (s *ListAICoachScriptPageResponseBodyList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAICoachScriptPageResponseBodyList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAICoachScriptPageResponseBodyList) GetInitiator() *string {
	return s.Initiator
}

func (s *ListAICoachScriptPageResponseBodyList) GetInteractionType() *string {
	return s.InteractionType
}

func (s *ListAICoachScriptPageResponseBodyList) GetIntroduce() *string {
	return s.Introduce
}

func (s *ListAICoachScriptPageResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListAICoachScriptPageResponseBodyList) GetOpeningRemarks() *string {
	return s.OpeningRemarks
}

func (s *ListAICoachScriptPageResponseBodyList) GetOrderAckFlag() *bool {
	return s.OrderAckFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetSampleDialogueList() []*ListAICoachScriptPageResponseBodyListSampleDialogueList {
	return s.SampleDialogueList
}

func (s *ListAICoachScriptPageResponseBodyList) GetScoreConfig() *ListAICoachScriptPageResponseBodyListScoreConfig {
	return s.ScoreConfig
}

func (s *ListAICoachScriptPageResponseBodyList) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *ListAICoachScriptPageResponseBodyList) GetSparringTipContent() *string {
	return s.SparringTipContent
}

func (s *ListAICoachScriptPageResponseBodyList) GetSparringTipTitle() *string {
	return s.SparringTipTitle
}

func (s *ListAICoachScriptPageResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListAICoachScriptPageResponseBodyList) GetStudentThinkTimeFlag() *bool {
	return s.StudentThinkTimeFlag
}

func (s *ListAICoachScriptPageResponseBodyList) GetType() *int32 {
	return s.Type
}

func (s *ListAICoachScriptPageResponseBodyList) GetWeights() *ListAICoachScriptPageResponseBodyListWeights {
	return s.Weights
}

func (s *ListAICoachScriptPageResponseBodyList) SetAppendQuestionFlag(v string) *ListAICoachScriptPageResponseBodyList {
	s.AppendQuestionFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetAssessmentScope(v string) *ListAICoachScriptPageResponseBodyList {
	s.AssessmentScope = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetClosingRemarks(v string) *ListAICoachScriptPageResponseBodyList {
	s.ClosingRemarks = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetCompleteStrategy(v *ListAICoachScriptPageResponseBodyListCompleteStrategy) *ListAICoachScriptPageResponseBodyList {
	s.CompleteStrategy = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetCoverUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetCustomReplyRules(v []*ListAICoachScriptPageResponseBodyListCustomReplyRules) *ListAICoachScriptPageResponseBodyList {
	s.CustomReplyRules = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetDialogueTextFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.DialogueTextFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetDialogueTipFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.DialogueTipFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetEvaluateReportFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.EvaluateReportFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetExpressiveness(v map[string]*string) *ListAICoachScriptPageResponseBodyList {
	s.Expressiveness = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGifDynamicUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.GifDynamicUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGifStaticUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.GifStaticUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGmtCreate(v string) *ListAICoachScriptPageResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGmtModified(v string) *ListAICoachScriptPageResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetInitiator(v string) *ListAICoachScriptPageResponseBodyList {
	s.Initiator = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetInteractionType(v string) *ListAICoachScriptPageResponseBodyList {
	s.InteractionType = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetIntroduce(v string) *ListAICoachScriptPageResponseBodyList {
	s.Introduce = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetName(v string) *ListAICoachScriptPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetOpeningRemarks(v string) *ListAICoachScriptPageResponseBodyList {
	s.OpeningRemarks = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetOrderAckFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.OrderAckFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSampleDialogueList(v []*ListAICoachScriptPageResponseBodyListSampleDialogueList) *ListAICoachScriptPageResponseBodyList {
	s.SampleDialogueList = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetScoreConfig(v *ListAICoachScriptPageResponseBodyListScoreConfig) *ListAICoachScriptPageResponseBodyList {
	s.ScoreConfig = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetScriptRecordId(v string) *ListAICoachScriptPageResponseBodyList {
	s.ScriptRecordId = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSparringTipContent(v string) *ListAICoachScriptPageResponseBodyList {
	s.SparringTipContent = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSparringTipTitle(v string) *ListAICoachScriptPageResponseBodyList {
	s.SparringTipTitle = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetStatus(v int32) *ListAICoachScriptPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetStudentThinkTimeFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetType(v int32) *ListAICoachScriptPageResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetWeights(v *ListAICoachScriptPageResponseBodyListWeights) *ListAICoachScriptPageResponseBodyList {
	s.Weights = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) Validate() error {
	if s.CompleteStrategy != nil {
		if err := s.CompleteStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.CustomReplyRules != nil {
		for _, item := range s.CustomReplyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SampleDialogueList != nil {
		for _, item := range s.SampleDialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScoreConfig != nil {
		if err := s.ScoreConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Weights != nil {
		if err := s.Weights.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyListCompleteStrategy struct {
	// example:
	//
	// true
	ClickCompleteAutoEnd *bool `json:"clickCompleteAutoEnd,omitempty" xml:"clickCompleteAutoEnd,omitempty"`
	// example:
	//
	// 75
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// false
	FullCoverageAutoEnd *bool `json:"fullCoverageAutoEnd,omitempty" xml:"fullCoverageAutoEnd,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCompleteStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCompleteStrategy) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) GetClickCompleteAutoEnd() *bool {
	return s.ClickCompleteAutoEnd
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) GetDuration() *int32 {
	return s.Duration
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) GetFullCoverageAutoEnd() *bool {
	return s.FullCoverageAutoEnd
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetClickCompleteAutoEnd(v bool) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.ClickCompleteAutoEnd = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetDuration(v int32) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.Duration = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetFullCoverageAutoEnd(v bool) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.FullCoverageAutoEnd = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListCustomReplyRules struct {
	Action *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
	// example:
	//
	// and
	//
	// or
	Logic         *string                                                             `json:"logic,omitempty" xml:"logic,omitempty"`
	MainCondition *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition `json:"mainCondition,omitempty" xml:"mainCondition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// 1
	SortNo       *int32                                                             `json:"sortNo,omitempty" xml:"sortNo,omitempty"`
	SubCondition *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition `json:"subCondition,omitempty" xml:"subCondition,omitempty" type:"Struct"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRules) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRules) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetAction() *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction {
	return s.Action
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetLogic() *string {
	return s.Logic
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetMainCondition() *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition {
	return s.MainCondition
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetSortNo() *int32 {
	return s.SortNo
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) GetSubCondition() *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition {
	return s.SubCondition
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetAction(v *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.Action = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetLogic(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.Logic = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetMainCondition(v *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.MainCondition = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetPriority(v int32) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.Priority = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetSortNo(v int32) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.SortNo = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) SetSubCondition(v *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) *ListAICoachScriptPageResponseBodyListCustomReplyRules {
	s.SubCondition = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRules) Validate() error {
	if s.Action != nil {
		if err := s.Action.Validate(); err != nil {
			return err
		}
	}
	if s.MainCondition != nil {
		if err := s.MainCondition.Validate(); err != nil {
			return err
		}
	}
	if s.SubCondition != nil {
		if err := s.SubCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyListCustomReplyRulesAction struct {
	Parameters *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Type       *string                                                                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) GetParameters() *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters {
	return s.Parameters
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) GetType() *string {
	return s.Type
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) SetParameters(v *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction {
	s.Parameters = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) SetType(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesAction) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters struct {
	// example:
	//
	// 1
	AssessPoint   *string `json:"assessPoint,omitempty" xml:"assessPoint,omitempty"`
	CustomContent *string `json:"customContent,omitempty" xml:"customContent,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) GetAssessPoint() *string {
	return s.AssessPoint
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) GetCustomContent() *string {
	return s.CustomContent
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) SetAssessPoint(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters {
	s.AssessPoint = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) SetCustomContent(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters {
	s.CustomContent = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesActionParameters) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition struct {
	Parameters *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Type       *string                                                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) GetParameters() *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters {
	return s.Parameters
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) GetType() *string {
	return s.Type
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) SetParameters(v *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition {
	s.Parameters = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) SetType(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainCondition) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters struct {
	// example:
	//
	// 1
	AssessPoint *string `json:"assessPoint,omitempty" xml:"assessPoint,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) GetAssessPoint() *string {
	return s.AssessPoint
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) SetAssessPoint(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters {
	s.AssessPoint = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesMainConditionParameters) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) GetType() *string {
	return s.Type
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) SetType(v string) *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCustomReplyRulesSubCondition) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListSampleDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// student
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListSampleDialogueList) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListSampleDialogueList) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) GetMessage() *string {
	return s.Message
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) GetRole() *string {
	return s.Role
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) SetMessage(v string) *ListAICoachScriptPageResponseBodyListSampleDialogueList {
	s.Message = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) SetRole(v string) *ListAICoachScriptPageResponseBodyListSampleDialogueList {
	s.Role = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListScoreConfig struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// true
	LevelEnabled *bool                                                     `json:"levelEnabled,omitempty" xml:"levelEnabled,omitempty"`
	Levels       []*ListAICoachScriptPageResponseBodyListScoreConfigLevels `json:"levels,omitempty" xml:"levels,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	PassScore *int32 `json:"passScore,omitempty" xml:"passScore,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListScoreConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListScoreConfig) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) GetLevelEnabled() *bool {
	return s.LevelEnabled
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) GetLevels() []*ListAICoachScriptPageResponseBodyListScoreConfigLevels {
	return s.Levels
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) GetPassScore() *int32 {
	return s.PassScore
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetEnabled(v bool) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.Enabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetLevelEnabled(v bool) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.LevelEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetLevels(v []*ListAICoachScriptPageResponseBodyListScoreConfigLevels) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.Levels = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetPassScore(v int32) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.PassScore = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) Validate() error {
	if s.Levels != nil {
		for _, item := range s.Levels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICoachScriptPageResponseBodyListScoreConfigLevels struct {
	// example:
	//
	// 80
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 60
	Min  *int32  `json:"min,omitempty" xml:"min,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListScoreConfigLevels) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListScoreConfigLevels) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) GetMax() *int32 {
	return s.Max
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) GetMin() *int32 {
	return s.Min
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) GetName() *string {
	return s.Name
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) SetMax(v int32) *ListAICoachScriptPageResponseBodyListScoreConfigLevels {
	s.Max = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) SetMin(v int32) *ListAICoachScriptPageResponseBodyListScoreConfigLevels {
	s.Min = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) SetName(v string) *ListAICoachScriptPageResponseBodyListScoreConfigLevels {
	s.Name = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfigLevels) Validate() error {
	return dara.Validate(s)
}

type ListAICoachScriptPageResponseBodyListWeights struct {
	// example:
	//
	// 50
	AssessmentPoint *int32 `json:"assessmentPoint,omitempty" xml:"assessmentPoint,omitempty"`
	// example:
	//
	// true
	AssessmentPointEnabled *bool `json:"assessmentPointEnabled,omitempty" xml:"assessmentPointEnabled,omitempty"`
	// example:
	//
	// true
	CustomReplyRuleEnabled *bool `json:"customReplyRuleEnabled,omitempty" xml:"customReplyRuleEnabled,omitempty"`
	// example:
	//
	// 30
	Expressiveness *int32 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	// example:
	//
	// true
	ExpressivenessEnabled *bool `json:"expressivenessEnabled,omitempty" xml:"expressivenessEnabled,omitempty"`
	// example:
	//
	// 20
	PointDeductionRule *int32 `json:"pointDeductionRule,omitempty" xml:"pointDeductionRule,omitempty"`
	// example:
	//
	// true
	PointDeductionRuleEnabled *bool `json:"pointDeductionRuleEnabled,omitempty" xml:"pointDeductionRuleEnabled,omitempty"`
	// example:
	//
	// true
	SimilarPronunciationScoringEnabled *bool `json:"similarPronunciationScoringEnabled,omitempty" xml:"similarPronunciationScoringEnabled,omitempty"`
	// example:
	//
	// 20
	Standard *int32 `json:"standard,omitempty" xml:"standard,omitempty"`
	// example:
	//
	// true
	StandardEnabled *bool `json:"standardEnabled,omitempty" xml:"standardEnabled,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListWeights) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListWeights) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetAssessmentPoint() *int32 {
	return s.AssessmentPoint
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetAssessmentPointEnabled() *bool {
	return s.AssessmentPointEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetCustomReplyRuleEnabled() *bool {
	return s.CustomReplyRuleEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetExpressiveness() *int32 {
	return s.Expressiveness
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetExpressivenessEnabled() *bool {
	return s.ExpressivenessEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetPointDeductionRule() *int32 {
	return s.PointDeductionRule
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetPointDeductionRuleEnabled() *bool {
	return s.PointDeductionRuleEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetSimilarPronunciationScoringEnabled() *bool {
	return s.SimilarPronunciationScoringEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetStandard() *int32 {
	return s.Standard
}

func (s *ListAICoachScriptPageResponseBodyListWeights) GetStandardEnabled() *bool {
	return s.StandardEnabled
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetAssessmentPoint(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.AssessmentPoint = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetAssessmentPointEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.AssessmentPointEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetCustomReplyRuleEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.CustomReplyRuleEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetExpressiveness(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.Expressiveness = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetExpressivenessEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.ExpressivenessEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetPointDeductionRule(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.PointDeductionRule = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetPointDeductionRuleEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.PointDeductionRuleEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetSimilarPronunciationScoringEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.SimilarPronunciationScoringEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetStandard(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.Standard = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetStandardEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.StandardEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) Validate() error {
	return dara.Validate(s)
}
