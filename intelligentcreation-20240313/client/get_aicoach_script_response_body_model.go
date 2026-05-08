// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetAICoachScriptResponseBody
	GetAgentId() *string
	SetAppendQuestionFlag(v bool) *GetAICoachScriptResponseBody
	GetAppendQuestionFlag() *bool
	SetAssessmentScope(v string) *GetAICoachScriptResponseBody
	GetAssessmentScope() *string
	SetCheckCheatConfig(v *GetAICoachScriptResponseBodyCheckCheatConfig) *GetAICoachScriptResponseBody
	GetCheckCheatConfig() *GetAICoachScriptResponseBodyCheckCheatConfig
	SetClosingRemarks(v string) *GetAICoachScriptResponseBody
	GetClosingRemarks() *string
	SetCompleteStrategy(v *GetAICoachScriptResponseBodyCompleteStrategy) *GetAICoachScriptResponseBody
	GetCompleteStrategy() *GetAICoachScriptResponseBodyCompleteStrategy
	SetCoverUrl(v string) *GetAICoachScriptResponseBody
	GetCoverUrl() *string
	SetCustomReplyRules(v []*GetAICoachScriptResponseBodyCustomReplyRules) *GetAICoachScriptResponseBody
	GetCustomReplyRules() []*GetAICoachScriptResponseBodyCustomReplyRules
	SetDialogueInputTextLimit(v int32) *GetAICoachScriptResponseBody
	GetDialogueInputTextLimit() *int32
	SetDialogueTextFlag(v bool) *GetAICoachScriptResponseBody
	GetDialogueTextFlag() *bool
	SetDialogueTipFlag(v bool) *GetAICoachScriptResponseBody
	GetDialogueTipFlag() *bool
	SetDialogueVoiceLimit(v int32) *GetAICoachScriptResponseBody
	GetDialogueVoiceLimit() *int32
	SetEvaluateReportFlag(v bool) *GetAICoachScriptResponseBody
	GetEvaluateReportFlag() *bool
	SetExpressiveness(v map[string]*int32) *GetAICoachScriptResponseBody
	GetExpressiveness() map[string]*int32
	SetExpressivenessList(v []*GetAICoachScriptResponseBodyExpressivenessList) *GetAICoachScriptResponseBody
	GetExpressivenessList() []*GetAICoachScriptResponseBodyExpressivenessList
	SetGifDynamicUrl(v string) *GetAICoachScriptResponseBody
	GetGifDynamicUrl() *string
	SetGifStaticUrl(v string) *GetAICoachScriptResponseBody
	GetGifStaticUrl() *string
	SetGmtCreate(v string) *GetAICoachScriptResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetAICoachScriptResponseBody
	GetGmtModified() *string
	SetInitiator(v string) *GetAICoachScriptResponseBody
	GetInitiator() *string
	SetInteractionInputTypes(v []*string) *GetAICoachScriptResponseBody
	GetInteractionInputTypes() []*string
	SetInteractionType(v int32) *GetAICoachScriptResponseBody
	GetInteractionType() *int32
	SetIntroduce(v string) *GetAICoachScriptResponseBody
	GetIntroduce() *string
	SetName(v string) *GetAICoachScriptResponseBody
	GetName() *string
	SetOpeningRemarks(v string) *GetAICoachScriptResponseBody
	GetOpeningRemarks() *string
	SetOrderAckFlag(v bool) *GetAICoachScriptResponseBody
	GetOrderAckFlag() *bool
	SetPointDeductionRuleList(v []*GetAICoachScriptResponseBodyPointDeductionRuleList) *GetAICoachScriptResponseBody
	GetPointDeductionRuleList() []*GetAICoachScriptResponseBodyPointDeductionRuleList
	SetPoints(v []*GetAICoachScriptResponseBodyPoints) *GetAICoachScriptResponseBody
	GetPoints() []*GetAICoachScriptResponseBodyPoints
	SetRequestId(v string) *GetAICoachScriptResponseBody
	GetRequestId() *string
	SetSampleDialogueList(v []*GetAICoachScriptResponseBodySampleDialogueList) *GetAICoachScriptResponseBody
	GetSampleDialogueList() []*GetAICoachScriptResponseBodySampleDialogueList
	SetScoreConfig(v *GetAICoachScriptResponseBodyScoreConfig) *GetAICoachScriptResponseBody
	GetScoreConfig() *GetAICoachScriptResponseBodyScoreConfig
	SetScriptRecordId(v string) *GetAICoachScriptResponseBody
	GetScriptRecordId() *string
	SetSparringTipContent(v string) *GetAICoachScriptResponseBody
	GetSparringTipContent() *string
	SetSparringTipTitle(v string) *GetAICoachScriptResponseBody
	GetSparringTipTitle() *string
	SetStatus(v int32) *GetAICoachScriptResponseBody
	GetStatus() *int32
	SetStudentThinkTimeFlag(v bool) *GetAICoachScriptResponseBody
	GetStudentThinkTimeFlag() *bool
	SetStudentThinkTimeLimit(v int32) *GetAICoachScriptResponseBody
	GetStudentThinkTimeLimit() *int32
	SetType(v int32) *GetAICoachScriptResponseBody
	GetType() *int32
	SetVoiceId(v string) *GetAICoachScriptResponseBody
	GetVoiceId() *string
	SetVoiceLanguage(v string) *GetAICoachScriptResponseBody
	GetVoiceLanguage() *string
	SetWeights(v *GetAICoachScriptResponseBodyWeights) *GetAICoachScriptResponseBody
	GetWeights() *GetAICoachScriptResponseBodyWeights
}

type GetAICoachScriptResponseBody struct {
	AgentId            *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppendQuestionFlag *bool   `json:"appendQuestionFlag,omitempty" xml:"appendQuestionFlag,omitempty"`
	// example:
	//
	// point
	AssessmentScope  *string                                       `json:"assessmentScope,omitempty" xml:"assessmentScope,omitempty"`
	CheckCheatConfig *GetAICoachScriptResponseBodyCheckCheatConfig `json:"checkCheatConfig,omitempty" xml:"checkCheatConfig,omitempty" type:"Struct"`
	ClosingRemarks   *string                                       `json:"closingRemarks,omitempty" xml:"closingRemarks,omitempty"`
	CompleteStrategy *GetAICoachScriptResponseBodyCompleteStrategy `json:"completeStrategy,omitempty" xml:"completeStrategy,omitempty" type:"Struct"`
	// example:
	//
	// https://demo.com
	CoverUrl         *string                                         `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CustomReplyRules []*GetAICoachScriptResponseBodyCustomReplyRules `json:"customReplyRules,omitempty" xml:"customReplyRules,omitempty" type:"Repeated"`
	// example:
	//
	// 500
	DialogueInputTextLimit *int32 `json:"dialogueInputTextLimit,omitempty" xml:"dialogueInputTextLimit,omitempty"`
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
	// 30
	DialogueVoiceLimit *int32 `json:"dialogueVoiceLimit,omitempty" xml:"dialogueVoiceLimit,omitempty"`
	// example:
	//
	// true
	EvaluateReportFlag *bool                                             `json:"evaluateReportFlag,omitempty" xml:"evaluateReportFlag,omitempty"`
	Expressiveness     map[string]*int32                                 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	ExpressivenessList []*GetAICoachScriptResponseBodyExpressivenessList `json:"expressivenessList,omitempty" xml:"expressivenessList,omitempty" type:"Repeated"`
	GifDynamicUrl      *string                                           `json:"gifDynamicUrl,omitempty" xml:"gifDynamicUrl,omitempty"`
	GifStaticUrl       *string                                           `json:"gifStaticUrl,omitempty" xml:"gifStaticUrl,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// coach
	Initiator             *string   `json:"initiator,omitempty" xml:"initiator,omitempty"`
	InteractionInputTypes []*string `json:"interactionInputTypes,omitempty" xml:"interactionInputTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InteractionType *int32 `json:"interactionType,omitempty" xml:"interactionType,omitempty"`
	// example:
	//
	// demo
	Introduce *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// example:
	//
	// demo
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	// example:
	//
	// true
	OrderAckFlag           *bool                                                 `json:"orderAckFlag,omitempty" xml:"orderAckFlag,omitempty"`
	PointDeductionRuleList []*GetAICoachScriptResponseBodyPointDeductionRuleList `json:"pointDeductionRuleList,omitempty" xml:"pointDeductionRuleList,omitempty" type:"Repeated"`
	Points                 []*GetAICoachScriptResponseBodyPoints                 `json:"points,omitempty" xml:"points,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RequestId          *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SampleDialogueList []*GetAICoachScriptResponseBodySampleDialogueList `json:"sampleDialogueList,omitempty" xml:"sampleDialogueList,omitempty" type:"Repeated"`
	ScoreConfig        *GetAICoachScriptResponseBodyScoreConfig          `json:"scoreConfig,omitempty" xml:"scoreConfig,omitempty" type:"Struct"`
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
	// 100
	StudentThinkTimeLimit *int32 `json:"studentThinkTimeLimit,omitempty" xml:"studentThinkTimeLimit,omitempty"`
	// example:
	//
	// 1
	Type          *int32                               `json:"type,omitempty" xml:"type,omitempty"`
	VoiceId       *string                              `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
	VoiceLanguage *string                              `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	Weights       *GetAICoachScriptResponseBodyWeights `json:"weights,omitempty" xml:"weights,omitempty" type:"Struct"`
}

func (s GetAICoachScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAICoachScriptResponseBody) GetAppendQuestionFlag() *bool {
	return s.AppendQuestionFlag
}

func (s *GetAICoachScriptResponseBody) GetAssessmentScope() *string {
	return s.AssessmentScope
}

func (s *GetAICoachScriptResponseBody) GetCheckCheatConfig() *GetAICoachScriptResponseBodyCheckCheatConfig {
	return s.CheckCheatConfig
}

func (s *GetAICoachScriptResponseBody) GetClosingRemarks() *string {
	return s.ClosingRemarks
}

func (s *GetAICoachScriptResponseBody) GetCompleteStrategy() *GetAICoachScriptResponseBodyCompleteStrategy {
	return s.CompleteStrategy
}

func (s *GetAICoachScriptResponseBody) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetAICoachScriptResponseBody) GetCustomReplyRules() []*GetAICoachScriptResponseBodyCustomReplyRules {
	return s.CustomReplyRules
}

func (s *GetAICoachScriptResponseBody) GetDialogueInputTextLimit() *int32 {
	return s.DialogueInputTextLimit
}

func (s *GetAICoachScriptResponseBody) GetDialogueTextFlag() *bool {
	return s.DialogueTextFlag
}

func (s *GetAICoachScriptResponseBody) GetDialogueTipFlag() *bool {
	return s.DialogueTipFlag
}

func (s *GetAICoachScriptResponseBody) GetDialogueVoiceLimit() *int32 {
	return s.DialogueVoiceLimit
}

func (s *GetAICoachScriptResponseBody) GetEvaluateReportFlag() *bool {
	return s.EvaluateReportFlag
}

func (s *GetAICoachScriptResponseBody) GetExpressiveness() map[string]*int32 {
	return s.Expressiveness
}

func (s *GetAICoachScriptResponseBody) GetExpressivenessList() []*GetAICoachScriptResponseBodyExpressivenessList {
	return s.ExpressivenessList
}

func (s *GetAICoachScriptResponseBody) GetGifDynamicUrl() *string {
	return s.GifDynamicUrl
}

func (s *GetAICoachScriptResponseBody) GetGifStaticUrl() *string {
	return s.GifStaticUrl
}

func (s *GetAICoachScriptResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAICoachScriptResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAICoachScriptResponseBody) GetInitiator() *string {
	return s.Initiator
}

func (s *GetAICoachScriptResponseBody) GetInteractionInputTypes() []*string {
	return s.InteractionInputTypes
}

func (s *GetAICoachScriptResponseBody) GetInteractionType() *int32 {
	return s.InteractionType
}

func (s *GetAICoachScriptResponseBody) GetIntroduce() *string {
	return s.Introduce
}

func (s *GetAICoachScriptResponseBody) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBody) GetOpeningRemarks() *string {
	return s.OpeningRemarks
}

func (s *GetAICoachScriptResponseBody) GetOrderAckFlag() *bool {
	return s.OrderAckFlag
}

func (s *GetAICoachScriptResponseBody) GetPointDeductionRuleList() []*GetAICoachScriptResponseBodyPointDeductionRuleList {
	return s.PointDeductionRuleList
}

func (s *GetAICoachScriptResponseBody) GetPoints() []*GetAICoachScriptResponseBodyPoints {
	return s.Points
}

func (s *GetAICoachScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachScriptResponseBody) GetSampleDialogueList() []*GetAICoachScriptResponseBodySampleDialogueList {
	return s.SampleDialogueList
}

func (s *GetAICoachScriptResponseBody) GetScoreConfig() *GetAICoachScriptResponseBodyScoreConfig {
	return s.ScoreConfig
}

func (s *GetAICoachScriptResponseBody) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *GetAICoachScriptResponseBody) GetSparringTipContent() *string {
	return s.SparringTipContent
}

func (s *GetAICoachScriptResponseBody) GetSparringTipTitle() *string {
	return s.SparringTipTitle
}

func (s *GetAICoachScriptResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetAICoachScriptResponseBody) GetStudentThinkTimeFlag() *bool {
	return s.StudentThinkTimeFlag
}

func (s *GetAICoachScriptResponseBody) GetStudentThinkTimeLimit() *int32 {
	return s.StudentThinkTimeLimit
}

func (s *GetAICoachScriptResponseBody) GetType() *int32 {
	return s.Type
}

func (s *GetAICoachScriptResponseBody) GetVoiceId() *string {
	return s.VoiceId
}

func (s *GetAICoachScriptResponseBody) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *GetAICoachScriptResponseBody) GetWeights() *GetAICoachScriptResponseBodyWeights {
	return s.Weights
}

func (s *GetAICoachScriptResponseBody) SetAgentId(v string) *GetAICoachScriptResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetAppendQuestionFlag(v bool) *GetAICoachScriptResponseBody {
	s.AppendQuestionFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetAssessmentScope(v string) *GetAICoachScriptResponseBody {
	s.AssessmentScope = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCheckCheatConfig(v *GetAICoachScriptResponseBodyCheckCheatConfig) *GetAICoachScriptResponseBody {
	s.CheckCheatConfig = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetClosingRemarks(v string) *GetAICoachScriptResponseBody {
	s.ClosingRemarks = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCompleteStrategy(v *GetAICoachScriptResponseBodyCompleteStrategy) *GetAICoachScriptResponseBody {
	s.CompleteStrategy = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCoverUrl(v string) *GetAICoachScriptResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCustomReplyRules(v []*GetAICoachScriptResponseBodyCustomReplyRules) *GetAICoachScriptResponseBody {
	s.CustomReplyRules = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueInputTextLimit(v int32) *GetAICoachScriptResponseBody {
	s.DialogueInputTextLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueTextFlag(v bool) *GetAICoachScriptResponseBody {
	s.DialogueTextFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueTipFlag(v bool) *GetAICoachScriptResponseBody {
	s.DialogueTipFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueVoiceLimit(v int32) *GetAICoachScriptResponseBody {
	s.DialogueVoiceLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetEvaluateReportFlag(v bool) *GetAICoachScriptResponseBody {
	s.EvaluateReportFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetExpressiveness(v map[string]*int32) *GetAICoachScriptResponseBody {
	s.Expressiveness = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetExpressivenessList(v []*GetAICoachScriptResponseBodyExpressivenessList) *GetAICoachScriptResponseBody {
	s.ExpressivenessList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGifDynamicUrl(v string) *GetAICoachScriptResponseBody {
	s.GifDynamicUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGifStaticUrl(v string) *GetAICoachScriptResponseBody {
	s.GifStaticUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGmtCreate(v string) *GetAICoachScriptResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGmtModified(v string) *GetAICoachScriptResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInitiator(v string) *GetAICoachScriptResponseBody {
	s.Initiator = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInteractionInputTypes(v []*string) *GetAICoachScriptResponseBody {
	s.InteractionInputTypes = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInteractionType(v int32) *GetAICoachScriptResponseBody {
	s.InteractionType = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetIntroduce(v string) *GetAICoachScriptResponseBody {
	s.Introduce = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetName(v string) *GetAICoachScriptResponseBody {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetOpeningRemarks(v string) *GetAICoachScriptResponseBody {
	s.OpeningRemarks = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetOrderAckFlag(v bool) *GetAICoachScriptResponseBody {
	s.OrderAckFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetPointDeductionRuleList(v []*GetAICoachScriptResponseBodyPointDeductionRuleList) *GetAICoachScriptResponseBody {
	s.PointDeductionRuleList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetPoints(v []*GetAICoachScriptResponseBodyPoints) *GetAICoachScriptResponseBody {
	s.Points = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetRequestId(v string) *GetAICoachScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSampleDialogueList(v []*GetAICoachScriptResponseBodySampleDialogueList) *GetAICoachScriptResponseBody {
	s.SampleDialogueList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetScoreConfig(v *GetAICoachScriptResponseBodyScoreConfig) *GetAICoachScriptResponseBody {
	s.ScoreConfig = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetScriptRecordId(v string) *GetAICoachScriptResponseBody {
	s.ScriptRecordId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSparringTipContent(v string) *GetAICoachScriptResponseBody {
	s.SparringTipContent = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSparringTipTitle(v string) *GetAICoachScriptResponseBody {
	s.SparringTipTitle = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStatus(v int32) *GetAICoachScriptResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStudentThinkTimeFlag(v bool) *GetAICoachScriptResponseBody {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStudentThinkTimeLimit(v int32) *GetAICoachScriptResponseBody {
	s.StudentThinkTimeLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetType(v int32) *GetAICoachScriptResponseBody {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetVoiceId(v string) *GetAICoachScriptResponseBody {
	s.VoiceId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetVoiceLanguage(v string) *GetAICoachScriptResponseBody {
	s.VoiceLanguage = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetWeights(v *GetAICoachScriptResponseBodyWeights) *GetAICoachScriptResponseBody {
	s.Weights = v
	return s
}

func (s *GetAICoachScriptResponseBody) Validate() error {
	if s.CheckCheatConfig != nil {
		if err := s.CheckCheatConfig.Validate(); err != nil {
			return err
		}
	}
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
	if s.ExpressivenessList != nil {
		for _, item := range s.ExpressivenessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PointDeductionRuleList != nil {
		for _, item := range s.PointDeductionRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Points != nil {
		for _, item := range s.Points {
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

type GetAICoachScriptResponseBodyCheckCheatConfig struct {
	CheckImage *bool `json:"checkImage,omitempty" xml:"checkImage,omitempty"`
	CheckVoice *bool `json:"checkVoice,omitempty" xml:"checkVoice,omitempty"`
}

func (s GetAICoachScriptResponseBodyCheckCheatConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCheckCheatConfig) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) GetCheckImage() *bool {
	return s.CheckImage
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) GetCheckVoice() *bool {
	return s.CheckVoice
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) SetCheckImage(v bool) *GetAICoachScriptResponseBodyCheckCheatConfig {
	s.CheckImage = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) SetCheckVoice(v bool) *GetAICoachScriptResponseBodyCheckCheatConfig {
	s.CheckVoice = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyCompleteStrategy struct {
	// example:
	//
	// 5
	AbnormalQuitSessionExpired *int32 `json:"abnormalQuitSessionExpired,omitempty" xml:"abnormalQuitSessionExpired,omitempty"`
	// example:
	//
	// true
	AbnormalQuitSessionExpiredFlag *bool `json:"abnormalQuitSessionExpiredFlag,omitempty" xml:"abnormalQuitSessionExpiredFlag,omitempty"`
	// example:
	//
	// true
	ClickCompleteAutoEnd *bool `json:"clickCompleteAutoEnd,omitempty" xml:"clickCompleteAutoEnd,omitempty"`
	// example:
	//
	// 15
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// true
	DurationFlag *bool `json:"durationFlag,omitempty" xml:"durationFlag,omitempty"`
	// example:
	//
	// true
	FullCoverageAutoEnd *bool `json:"fullCoverageAutoEnd,omitempty" xml:"fullCoverageAutoEnd,omitempty"`
}

func (s GetAICoachScriptResponseBodyCompleteStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCompleteStrategy) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetAbnormalQuitSessionExpired() *int32 {
	return s.AbnormalQuitSessionExpired
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetAbnormalQuitSessionExpiredFlag() *bool {
	return s.AbnormalQuitSessionExpiredFlag
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetClickCompleteAutoEnd() *bool {
	return s.ClickCompleteAutoEnd
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetDuration() *int32 {
	return s.Duration
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetDurationFlag() *bool {
	return s.DurationFlag
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) GetFullCoverageAutoEnd() *bool {
	return s.FullCoverageAutoEnd
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetAbnormalQuitSessionExpired(v int32) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.AbnormalQuitSessionExpired = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetAbnormalQuitSessionExpiredFlag(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.AbnormalQuitSessionExpiredFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetClickCompleteAutoEnd(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.ClickCompleteAutoEnd = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetDuration(v int32) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.Duration = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetDurationFlag(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.DurationFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetFullCoverageAutoEnd(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.FullCoverageAutoEnd = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyCustomReplyRules struct {
	Action        *GetAICoachScriptResponseBodyCustomReplyRulesAction        `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
	Logic         *string                                                    `json:"logic,omitempty" xml:"logic,omitempty"`
	MainCondition *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition `json:"mainCondition,omitempty" xml:"mainCondition,omitempty" type:"Struct"`
	Priority      *int32                                                     `json:"priority,omitempty" xml:"priority,omitempty"`
	SubCondition  *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition  `json:"subCondition,omitempty" xml:"subCondition,omitempty" type:"Struct"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRules) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRules) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) GetAction() *GetAICoachScriptResponseBodyCustomReplyRulesAction {
	return s.Action
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) GetLogic() *string {
	return s.Logic
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) GetMainCondition() *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition {
	return s.MainCondition
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) GetPriority() *int32 {
	return s.Priority
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) GetSubCondition() *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition {
	return s.SubCondition
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) SetAction(v *GetAICoachScriptResponseBodyCustomReplyRulesAction) *GetAICoachScriptResponseBodyCustomReplyRules {
	s.Action = v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) SetLogic(v string) *GetAICoachScriptResponseBodyCustomReplyRules {
	s.Logic = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) SetMainCondition(v *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) *GetAICoachScriptResponseBodyCustomReplyRules {
	s.MainCondition = v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) SetPriority(v int32) *GetAICoachScriptResponseBodyCustomReplyRules {
	s.Priority = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) SetSubCondition(v *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) *GetAICoachScriptResponseBodyCustomReplyRules {
	s.SubCondition = v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRules) Validate() error {
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

type GetAICoachScriptResponseBodyCustomReplyRulesAction struct {
	Parameters *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Type       *string                                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesAction) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesAction) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesAction) GetParameters() *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters {
	return s.Parameters
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesAction) GetType() *string {
	return s.Type
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesAction) SetParameters(v *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) *GetAICoachScriptResponseBodyCustomReplyRulesAction {
	s.Parameters = v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesAction) SetType(v string) *GetAICoachScriptResponseBodyCustomReplyRulesAction {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesAction) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachScriptResponseBodyCustomReplyRulesActionParameters struct {
	AssessPointId *string `json:"assessPointId,omitempty" xml:"assessPointId,omitempty"`
	CustomContent *string `json:"customContent,omitempty" xml:"customContent,omitempty"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) GetAssessPointId() *string {
	return s.AssessPointId
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) GetCustomContent() *string {
	return s.CustomContent
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) SetAssessPointId(v string) *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters {
	s.AssessPointId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) SetCustomContent(v string) *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters {
	s.CustomContent = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesActionParameters) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyCustomReplyRulesMainCondition struct {
	Parameters *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Type       *string                                                              `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) GetParameters() *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters {
	return s.Parameters
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) GetType() *string {
	return s.Type
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) SetParameters(v *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition {
	s.Parameters = v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) SetType(v string) *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainCondition) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters struct {
	AssessPointId *string `json:"assessPointId,omitempty" xml:"assessPointId,omitempty"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) GetAssessPointId() *string {
	return s.AssessPointId
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) SetAssessPointId(v string) *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters {
	s.AssessPointId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesMainConditionParameters) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyCustomReplyRulesSubCondition struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) GetType() *string {
	return s.Type
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) SetType(v string) *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCustomReplyRulesSubCondition) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyExpressivenessList struct {
	Desc             *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	ExpressivenessId *string `json:"expressivenessId,omitempty" xml:"expressivenessId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Rule             *string `json:"rule,omitempty" xml:"rule,omitempty"`
	Type             *string `json:"type,omitempty" xml:"type,omitempty"`
	Weight           *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyExpressivenessList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyExpressivenessList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetDesc() *string {
	return s.Desc
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetExpressivenessId() *string {
	return s.ExpressivenessId
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetRule() *string {
	return s.Rule
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetType() *string {
	return s.Type
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetDesc(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Desc = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetEnabled(v bool) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Enabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetExpressivenessId(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.ExpressivenessId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetName(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetRule(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Rule = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetType(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetWeight(v int32) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Weight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyPointDeductionRuleList struct {
	// example:
	//
	// demo
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	PunishmentTypes []*string `json:"punishmentTypes,omitempty" xml:"punishmentTypes,omitempty" type:"Repeated"`
	RuleValue       *string   `json:"ruleValue,omitempty" xml:"ruleValue,omitempty"`
	// example:
	//
	// 90
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointDeductionRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointDeductionRuleList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) GetDescription() *string {
	return s.Description
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) GetPunishmentTypes() []*string {
	return s.PunishmentTypes
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) GetRuleValue() *string {
	return s.RuleValue
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetDescription(v string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.Description = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetPunishmentTypes(v []*string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.PunishmentTypes = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetRuleValue(v string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.RuleValue = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetWeight(v int32) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.Weight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyPoints struct {
	AnswerList    []*GetAICoachScriptResponseBodyPointsAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	KnowledgeList []*string                                       `json:"knowledgeList,omitempty" xml:"knowledgeList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
	// example:
	//
	// test
	QuestionDescription *string `json:"questionDescription,omitempty" xml:"questionDescription,omitempty"`
	// example:
	//
	// 1
	SortNo *int32 `json:"sortNo,omitempty" xml:"sortNo,omitempty"`
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPoints) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPoints) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPoints) GetAnswerList() []*GetAICoachScriptResponseBodyPointsAnswerList {
	return s.AnswerList
}

func (s *GetAICoachScriptResponseBodyPoints) GetKnowledgeList() []*string {
	return s.KnowledgeList
}

func (s *GetAICoachScriptResponseBodyPoints) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyPoints) GetPointId() *string {
	return s.PointId
}

func (s *GetAICoachScriptResponseBodyPoints) GetQuestionDescription() *string {
	return s.QuestionDescription
}

func (s *GetAICoachScriptResponseBodyPoints) GetSortNo() *int32 {
	return s.SortNo
}

func (s *GetAICoachScriptResponseBodyPoints) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachScriptResponseBodyPoints) SetAnswerList(v []*GetAICoachScriptResponseBodyPointsAnswerList) *GetAICoachScriptResponseBodyPoints {
	s.AnswerList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetKnowledgeList(v []*string) *GetAICoachScriptResponseBodyPoints {
	s.KnowledgeList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetName(v string) *GetAICoachScriptResponseBodyPoints {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetPointId(v string) *GetAICoachScriptResponseBodyPoints {
	s.PointId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetQuestionDescription(v string) *GetAICoachScriptResponseBodyPoints {
	s.QuestionDescription = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetSortNo(v int32) *GetAICoachScriptResponseBodyPoints {
	s.SortNo = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetWeight(v int32) *GetAICoachScriptResponseBodyPoints {
	s.Weight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) Validate() error {
	if s.AnswerList != nil {
		for _, item := range s.AnswerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachScriptResponseBodyPointsAnswerList struct {
	AnswerValues   []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValues `json:"answerValues,omitempty" xml:"answerValues,omitempty" type:"Repeated"`
	EnabledKeyword *bool                                                       `json:"enabledKeyword,omitempty" xml:"enabledKeyword,omitempty"`
	Name           *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	NameList       []*string                                                   `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	Operators      *string                                                     `json:"operators,omitempty" xml:"operators,omitempty"`
	Parameters     []*GetAICoachScriptResponseBodyPointsAnswerListParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// example:
	//
	// normalKnowledge
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetAnswerValues() []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	return s.AnswerValues
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetEnabledKeyword() *bool {
	return s.EnabledKeyword
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetNameList() []*string {
	return s.NameList
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetOperators() *string {
	return s.Operators
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetParameters() []*GetAICoachScriptResponseBodyPointsAnswerListParameters {
	return s.Parameters
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetType() *string {
	return s.Type
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetAnswerValues(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.AnswerValues = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetEnabledKeyword(v bool) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.EnabledKeyword = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetNameList(v []*string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.NameList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetOperators(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Operators = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetParameters(v []*GetAICoachScriptResponseBodyPointsAnswerListParameters) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Parameters = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetType(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Weight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) Validate() error {
	if s.AnswerValues != nil {
		for _, item := range s.AnswerValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValues struct {
	AnswerName    *string                                                                  `json:"answerName,omitempty" xml:"answerName,omitempty"`
	AnswerWeight  *int32                                                                   `json:"answerWeight,omitempty" xml:"answerWeight,omitempty"`
	KeywordValues []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues `json:"keywordValues,omitempty" xml:"keywordValues,omitempty" type:"Repeated"`
	KeywordWeight *int32                                                                   `json:"keywordWeight,omitempty" xml:"keywordWeight,omitempty"`
	ScoringRules  []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules  `json:"scoringRules,omitempty" xml:"scoringRules,omitempty" type:"Repeated"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GetAnswerName() *string {
	return s.AnswerName
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GetAnswerWeight() *int32 {
	return s.AnswerWeight
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GetKeywordValues() []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues {
	return s.KeywordValues
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GetKeywordWeight() *int32 {
	return s.KeywordWeight
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GetScoringRules() []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules {
	return s.ScoringRules
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetAnswerName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.AnswerName = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetAnswerWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.AnswerWeight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetKeywordValues(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.KeywordValues = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetKeywordWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.KeywordWeight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetScoringRules(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.ScoringRules = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) Validate() error {
	if s.KeywordValues != nil {
		for _, item := range s.KeywordValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScoringRules != nil {
		for _, item := range s.ScoringRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Weight *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) SetWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues {
	s.Weight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyPointsAnswerListParameters struct {
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListParameters) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) GetValue() *string {
	return s.Value
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListParameters {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) SetValue(v string) *GetAICoachScriptResponseBodyPointsAnswerListParameters {
	s.Value = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodySampleDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// coach
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachScriptResponseBodySampleDialogueList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodySampleDialogueList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) GetMessage() *string {
	return s.Message
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) GetRole() *string {
	return s.Role
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) SetMessage(v string) *GetAICoachScriptResponseBodySampleDialogueList {
	s.Message = &v
	return s
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) SetRole(v string) *GetAICoachScriptResponseBodySampleDialogueList {
	s.Role = &v
	return s
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyScoreConfig struct {
	Enabled      *bool                                            `json:"enabled,omitempty" xml:"enabled,omitempty"`
	LevelEnabled *bool                                            `json:"levelEnabled,omitempty" xml:"levelEnabled,omitempty"`
	Levels       []*GetAICoachScriptResponseBodyScoreConfigLevels `json:"levels,omitempty" xml:"levels,omitempty" type:"Repeated"`
	PassScore    *string                                          `json:"passScore,omitempty" xml:"passScore,omitempty"`
}

func (s GetAICoachScriptResponseBodyScoreConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyScoreConfig) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyScoreConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAICoachScriptResponseBodyScoreConfig) GetLevelEnabled() *bool {
	return s.LevelEnabled
}

func (s *GetAICoachScriptResponseBodyScoreConfig) GetLevels() []*GetAICoachScriptResponseBodyScoreConfigLevels {
	return s.Levels
}

func (s *GetAICoachScriptResponseBodyScoreConfig) GetPassScore() *string {
	return s.PassScore
}

func (s *GetAICoachScriptResponseBodyScoreConfig) SetEnabled(v bool) *GetAICoachScriptResponseBodyScoreConfig {
	s.Enabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfig) SetLevelEnabled(v bool) *GetAICoachScriptResponseBodyScoreConfig {
	s.LevelEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfig) SetLevels(v []*GetAICoachScriptResponseBodyScoreConfigLevels) *GetAICoachScriptResponseBodyScoreConfig {
	s.Levels = v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfig) SetPassScore(v string) *GetAICoachScriptResponseBodyScoreConfig {
	s.PassScore = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfig) Validate() error {
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

type GetAICoachScriptResponseBodyScoreConfigLevels struct {
	Max  *int32  `json:"max,omitempty" xml:"max,omitempty"`
	Min  *int32  `json:"min,omitempty" xml:"min,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachScriptResponseBodyScoreConfigLevels) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyScoreConfigLevels) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) GetMax() *int32 {
	return s.Max
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) GetMin() *int32 {
	return s.Min
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) GetName() *string {
	return s.Name
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) SetMax(v int32) *GetAICoachScriptResponseBodyScoreConfigLevels {
	s.Max = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) SetMin(v int32) *GetAICoachScriptResponseBodyScoreConfigLevels {
	s.Min = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) SetName(v string) *GetAICoachScriptResponseBodyScoreConfigLevels {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyScoreConfigLevels) Validate() error {
	return dara.Validate(s)
}

type GetAICoachScriptResponseBodyWeights struct {
	// example:
	//
	// 10
	AbilityEvaluation *int32 `json:"abilityEvaluation,omitempty" xml:"abilityEvaluation,omitempty"`
	// example:
	//
	// false
	AbilityEvaluationEnabled *bool `json:"abilityEvaluationEnabled,omitempty" xml:"abilityEvaluationEnabled,omitempty"`
	// example:
	//
	// 10
	AssessmentPoint        *int32 `json:"assessmentPoint,omitempty" xml:"assessmentPoint,omitempty"`
	AssessmentPointEnabled *bool  `json:"assessmentPointEnabled,omitempty" xml:"assessmentPointEnabled,omitempty"`
	CustomReplyRuleEnabled *bool  `json:"customReplyRuleEnabled,omitempty" xml:"customReplyRuleEnabled,omitempty"`
	// example:
	//
	// 10
	Expressiveness *int32 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	// example:
	//
	// true
	ExpressivenessEnabled *bool `json:"expressivenessEnabled,omitempty" xml:"expressivenessEnabled,omitempty"`
	// example:
	//
	// 10
	PointDeductionRule *int32 `json:"pointDeductionRule,omitempty" xml:"pointDeductionRule,omitempty"`
	// example:
	//
	// true
	PointDeductionRuleEnabled          *bool `json:"pointDeductionRuleEnabled,omitempty" xml:"pointDeductionRuleEnabled,omitempty"`
	SimilarPronunciationScoringEnabled *bool `json:"similarPronunciationScoringEnabled,omitempty" xml:"similarPronunciationScoringEnabled,omitempty"`
	// example:
	//
	// 10
	Standard *int32 `json:"standard,omitempty" xml:"standard,omitempty"`
	// example:
	//
	// true
	StandardEnabled *bool `json:"standardEnabled,omitempty" xml:"standardEnabled,omitempty"`
}

func (s GetAICoachScriptResponseBodyWeights) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponseBodyWeights) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyWeights) GetAbilityEvaluation() *int32 {
	return s.AbilityEvaluation
}

func (s *GetAICoachScriptResponseBodyWeights) GetAbilityEvaluationEnabled() *bool {
	return s.AbilityEvaluationEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetAssessmentPoint() *int32 {
	return s.AssessmentPoint
}

func (s *GetAICoachScriptResponseBodyWeights) GetAssessmentPointEnabled() *bool {
	return s.AssessmentPointEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetCustomReplyRuleEnabled() *bool {
	return s.CustomReplyRuleEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetExpressiveness() *int32 {
	return s.Expressiveness
}

func (s *GetAICoachScriptResponseBodyWeights) GetExpressivenessEnabled() *bool {
	return s.ExpressivenessEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetPointDeductionRule() *int32 {
	return s.PointDeductionRule
}

func (s *GetAICoachScriptResponseBodyWeights) GetPointDeductionRuleEnabled() *bool {
	return s.PointDeductionRuleEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetSimilarPronunciationScoringEnabled() *bool {
	return s.SimilarPronunciationScoringEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) GetStandard() *int32 {
	return s.Standard
}

func (s *GetAICoachScriptResponseBodyWeights) GetStandardEnabled() *bool {
	return s.StandardEnabled
}

func (s *GetAICoachScriptResponseBodyWeights) SetAbilityEvaluation(v int32) *GetAICoachScriptResponseBodyWeights {
	s.AbilityEvaluation = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAbilityEvaluationEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.AbilityEvaluationEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAssessmentPoint(v int32) *GetAICoachScriptResponseBodyWeights {
	s.AssessmentPoint = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAssessmentPointEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.AssessmentPointEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetCustomReplyRuleEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.CustomReplyRuleEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetExpressiveness(v int32) *GetAICoachScriptResponseBodyWeights {
	s.Expressiveness = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetExpressivenessEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.ExpressivenessEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetPointDeductionRule(v int32) *GetAICoachScriptResponseBodyWeights {
	s.PointDeductionRule = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetPointDeductionRuleEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.PointDeductionRuleEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetSimilarPronunciationScoringEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.SimilarPronunciationScoringEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetStandard(v int32) *GetAICoachScriptResponseBodyWeights {
	s.Standard = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetStandardEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.StandardEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) Validate() error {
	return dara.Validate(s)
}
