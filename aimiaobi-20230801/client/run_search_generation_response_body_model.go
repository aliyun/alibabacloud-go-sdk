// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunSearchGenerationResponseBodyHeader) *RunSearchGenerationResponseBody
	GetHeader() *RunSearchGenerationResponseBodyHeader
	SetPayload(v *RunSearchGenerationResponseBodyPayload) *RunSearchGenerationResponseBody
	GetPayload() *RunSearchGenerationResponseBodyPayload
	SetRequestId(v string) *RunSearchGenerationResponseBody
	GetRequestId() *string
}

type RunSearchGenerationResponseBody struct {
	Header  *RunSearchGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunSearchGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// xx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunSearchGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBody) GetHeader() *RunSearchGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunSearchGenerationResponseBody) GetPayload() *RunSearchGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunSearchGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSearchGenerationResponseBody) SetHeader(v *RunSearchGenerationResponseBodyHeader) *RunSearchGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunSearchGenerationResponseBody) SetPayload(v *RunSearchGenerationResponseBodyPayload) *RunSearchGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunSearchGenerationResponseBody) SetRequestId(v string) *RunSearchGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchGenerationResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// xx
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// example:
	//
	// 1
	ResponseTime *int64 `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// example:
	//
	// x
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// x
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xx
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunSearchGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunSearchGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunSearchGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunSearchGenerationResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunSearchGenerationResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunSearchGenerationResponseBodyHeader) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *RunSearchGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunSearchGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunSearchGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunSearchGenerationResponseBodyHeader) SetErrorCode(v string) *RunSearchGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetErrorMessage(v string) *RunSearchGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetEvent(v string) *RunSearchGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetEventInfo(v string) *RunSearchGenerationResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetOriginSessionId(v string) *RunSearchGenerationResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetResponseTime(v int64) *RunSearchGenerationResponseBodyHeader {
	s.ResponseTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetSessionId(v string) *RunSearchGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetTaskId(v string) *RunSearchGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) SetTraceId(v string) *RunSearchGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayload struct {
	Output *RunSearchGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunSearchGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayload) GetOutput() *RunSearchGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunSearchGenerationResponseBodyPayload) GetUsage() *RunSearchGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunSearchGenerationResponseBodyPayload) SetOutput(v *RunSearchGenerationResponseBodyPayloadOutput) *RunSearchGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayload) SetUsage(v *RunSearchGenerationResponseBodyPayloadUsage) *RunSearchGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutput struct {
	AgentContext *RunSearchGenerationResponseBodyPayloadOutputAgentContext `json:"AgentContext,omitempty" xml:"AgentContext,omitempty" type:"Struct"`
	Messages     []*RunSearchGenerationResponseBodyPayloadOutputMessages   `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutput) GetAgentContext() *RunSearchGenerationResponseBodyPayloadOutputAgentContext {
	return s.AgentContext
}

func (s *RunSearchGenerationResponseBodyPayloadOutput) GetMessages() []*RunSearchGenerationResponseBodyPayloadOutputMessages {
	return s.Messages
}

func (s *RunSearchGenerationResponseBodyPayloadOutput) SetAgentContext(v *RunSearchGenerationResponseBodyPayloadOutputAgentContext) *RunSearchGenerationResponseBodyPayloadOutput {
	s.AgentContext = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutput) SetMessages(v []*RunSearchGenerationResponseBodyPayloadOutputMessages) *RunSearchGenerationResponseBodyPayloadOutput {
	s.Messages = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutput) Validate() error {
	if s.AgentContext != nil {
		if err := s.AgentContext.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContext struct {
	BizContext *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext `json:"BizContext,omitempty" xml:"BizContext,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContext) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContext) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContext) GetBizContext() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	return s.BizContext
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContext) SetBizContext(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) *RunSearchGenerationResponseBodyPayloadOutputAgentContext {
	s.BizContext = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContext) Validate() error {
	if s.BizContext != nil {
		if err := s.BizContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext struct {
	// example:
	//
	// 您想了解关于xx的哪些信息？
	AskUser         *string   `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	AskUserKeywords []*string `json:"AskUserKeywords,omitempty" xml:"AskUserKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// start
	CurrentStep      *string                                                                             `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	GeneratedContent *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent `json:"GeneratedContent,omitempty" xml:"GeneratedContent,omitempty" type:"Struct"`
	ModelId          *string                                                                             `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// search
	NextStep                 *string   `json:"NextStep,omitempty" xml:"NextStep,omitempty"`
	RecommendSearchQueryList []*string `json:"RecommendSearchQueryList,omitempty" xml:"RecommendSearchQueryList,omitempty" type:"Repeated"`
	SearchKeywords           []*string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty" type:"Repeated"`
	SearchQueryList          []*string `json:"SearchQueryList,omitempty" xml:"SearchQueryList,omitempty" type:"Repeated"`
	// example:
	//
	// searchQuery
	SupplementDataType *string `json:"SupplementDataType,omitempty" xml:"SupplementDataType,omitempty"`
	// example:
	//
	// true
	SupplementEnable *bool                                                                             `json:"SupplementEnable,omitempty" xml:"SupplementEnable,omitempty"`
	TokenCalculate   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate `json:"TokenCalculate,omitempty" xml:"TokenCalculate,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetAskUser() *string {
	return s.AskUser
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetAskUserKeywords() []*string {
	return s.AskUserKeywords
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetGeneratedContent() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	return s.GeneratedContent
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetModelId() *string {
	return s.ModelId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetNextStep() *string {
	return s.NextStep
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetRecommendSearchQueryList() []*string {
	return s.RecommendSearchQueryList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetSearchKeywords() []*string {
	return s.SearchKeywords
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetSearchQueryList() []*string {
	return s.SearchQueryList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetSupplementDataType() *string {
	return s.SupplementDataType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetSupplementEnable() *bool {
	return s.SupplementEnable
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) GetTokenCalculate() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	return s.TokenCalculate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetAskUser(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.AskUser = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetAskUserKeywords(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.AskUserKeywords = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetCurrentStep(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.CurrentStep = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetGeneratedContent(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.GeneratedContent = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetModelId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.ModelId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetNextStep(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.NextStep = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetRecommendSearchQueryList(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.RecommendSearchQueryList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetSearchKeywords(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.SearchKeywords = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetSearchQueryList(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.SearchQueryList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetSupplementDataType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.SupplementDataType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetSupplementEnable(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.SupplementEnable = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) SetTokenCalculate(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext {
	s.TokenCalculate = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContext) Validate() error {
	if s.GeneratedContent != nil {
		if err := s.GeneratedContent.Validate(); err != nil {
			return err
		}
	}
	if s.TokenCalculate != nil {
		if err := s.TokenCalculate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent struct {
	AudioSearchResult  *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult  `json:"AudioSearchResult,omitempty" xml:"AudioSearchResult,omitempty" type:"Struct"`
	ClusterTopicResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult `json:"ClusterTopicResult,omitempty" xml:"ClusterTopicResult,omitempty" type:"Struct"`
	ExcerptResult      *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult      `json:"ExcerptResult,omitempty" xml:"ExcerptResult,omitempty" type:"Struct"`
	ImageSearchResult  *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult  `json:"ImageSearchResult,omitempty" xml:"ImageSearchResult,omitempty" type:"Struct"`
	NewsElementResult  *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult  `json:"NewsElementResult,omitempty" xml:"NewsElementResult,omitempty" type:"Struct"`
	TextGenerateResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult `json:"TextGenerateResult,omitempty" xml:"TextGenerateResult,omitempty" type:"Struct"`
	TextSearchResult   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult   `json:"TextSearchResult,omitempty" xml:"TextSearchResult,omitempty" type:"Struct"`
	TimelineResult     *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult     `json:"TimelineResult,omitempty" xml:"TimelineResult,omitempty" type:"Struct"`
	VideoSearchResult  *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult  `json:"VideoSearchResult,omitempty" xml:"VideoSearchResult,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetAudioSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult {
	return s.AudioSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetClusterTopicResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult {
	return s.ClusterTopicResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetExcerptResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	return s.ExcerptResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetImageSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult {
	return s.ImageSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetNewsElementResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult {
	return s.NewsElementResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetTextGenerateResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	return s.TextGenerateResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetTextSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult {
	return s.TextSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetTimelineResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	return s.TimelineResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) GetVideoSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult {
	return s.VideoSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetAudioSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.AudioSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetClusterTopicResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.ClusterTopicResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetExcerptResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.ExcerptResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetImageSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.ImageSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetNewsElementResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.NewsElementResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetTextGenerateResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.TextGenerateResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetTextSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.TextSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetTimelineResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.TimelineResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) SetVideoSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent {
	s.VideoSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContent) Validate() error {
	if s.AudioSearchResult != nil {
		if err := s.AudioSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.ClusterTopicResult != nil {
		if err := s.ClusterTopicResult.Validate(); err != nil {
			return err
		}
	}
	if s.ExcerptResult != nil {
		if err := s.ExcerptResult.Validate(); err != nil {
			return err
		}
	}
	if s.ImageSearchResult != nil {
		if err := s.ImageSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.NewsElementResult != nil {
		if err := s.NewsElementResult.Validate(); err != nil {
			return err
		}
	}
	if s.TextGenerateResult != nil {
		if err := s.TextGenerateResult.Validate(); err != nil {
			return err
		}
	}
	if s.TextSearchResult != nil {
		if err := s.TextSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.TimelineResult != nil {
		if err := s.TimelineResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoSearchResult != nil {
		if err := s.VideoSearchResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult struct {
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// http://xxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xxx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *string `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) GetTraceabilityId() *string {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) SetTraceabilityId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xxx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.9
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 2
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentAudioSearchResultSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult struct {
	ClusterTopics []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics `json:"ClusterTopics,omitempty" xml:"ClusterTopics,omitempty" type:"Repeated"`
	// example:
	//
	// true
	GenerateFinished *bool `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	// example:
	//
	// xx
	TextGenerate *string `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) GetClusterTopics() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	return s.ClusterTopics
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) SetClusterTopics(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult {
	s.ClusterTopics = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResult) Validate() error {
	if s.ClusterTopics != nil {
		for _, item := range s.ClusterTopics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics struct {
	AudioSearchResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult `json:"AudioSearchResult,omitempty" xml:"AudioSearchResult,omitempty" type:"Struct"`
	ImageSearchResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult `json:"ImageSearchResult,omitempty" xml:"ImageSearchResult,omitempty" type:"Struct"`
	TextSearchResult  *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult  `json:"TextSearchResult,omitempty" xml:"TextSearchResult,omitempty" type:"Struct"`
	// example:
	//
	// xx
	Topic             *string                                                                                                                             `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VideoSearchResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult `json:"VideoSearchResult,omitempty" xml:"VideoSearchResult,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GetAudioSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult {
	return s.AudioSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GetImageSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult {
	return s.ImageSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GetTextSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult {
	return s.TextSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GetTopic() *string {
	return s.Topic
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) GetVideoSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult {
	return s.VideoSearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) SetAudioSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	s.AudioSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) SetImageSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	s.ImageSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) SetTextSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	s.TextSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) SetTopic(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	s.Topic = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) SetVideoSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics {
	s.VideoSearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopics) Validate() error {
	if s.AudioSearchResult != nil {
		if err := s.AudioSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.ImageSearchResult != nil {
		if err := s.ImageSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.TextSearchResult != nil {
		if err := s.TextSearchResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoSearchResult != nil {
		if err := s.VideoSearchResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult struct {
	// example:
	//
	// 1
	Current      *int32                                                                                                                                          `json:"Current,omitempty" xml:"Current,omitempty"`
	SearchResult *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) GetSearchResult() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) SetSearchResult(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResult) Validate() error {
	if s.SearchResult != nil {
		if err := s.SearchResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// http://xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xxx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 1
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsAudioSearchResultSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult struct {
	// example:
	//
	// 1
	Current      *int32                                                                                                                                            `json:"Current,omitempty" xml:"Current,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult struct {
	Article *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsImageSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult struct {
	// example:
	//
	// 1
	Current      *int32                                                                                                                                           `json:"Current,omitempty" xml:"Current,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid          *string                                                                                                                                                          `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	MultimodalMedias []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-04-04 08:39:09
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetMultimodalMedias() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias {
	return s.MultimodalMedias
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetMultimodalMedias(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.MultimodalMedias = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResult) Validate() error {
	if s.MultimodalMedias != nil {
		for _, item := range s.MultimodalMedias {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias struct {
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsTextSearchResultSearchResultMultimodalMedias) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult struct {
	// example:
	//
	// 1
	Current      *int32                                                                                                                                            `json:"Current,omitempty" xml:"Current,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.9
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentClusterTopicResultClusterTopicsVideoSearchResultSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	// example:
	//
	// concise
	GenerateLevel      *string                                                                                                        `json:"GenerateLevel,omitempty" xml:"GenerateLevel,omitempty"`
	ReasonTextGenerate *string                                                                                                        `json:"ReasonTextGenerate,omitempty" xml:"ReasonTextGenerate,omitempty"`
	SearchResult       []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	TextGenerate *string `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GetGenerateLevel() *string {
	return s.GenerateLevel
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GetReasonTextGenerate() *string {
	return s.ReasonTextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) SetGenerateLevel(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	s.GenerateLevel = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) SetReasonTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	s.ReasonTextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult struct {
	Chunks []*string `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	Excerpt          *string                                                                                                                        `json:"Excerpt,omitempty" xml:"Excerpt,omitempty"`
	MultimodalMedias []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-04-04 08:39:09
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 0.99
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// xx
	Summary                         *string                                                                                                                                       `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextGenerateMultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList `json:"TextGenerateMultimodalMediaList,omitempty" xml:"TextGenerateMultimodalMediaList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *int32 `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetChunks() []*string {
	return s.Chunks
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetExcerpt() *string {
	return s.Excerpt
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetMultimodalMedias() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias {
	return s.MultimodalMedias
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetScore() *float32 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetSelect() *bool {
	return s.Select
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetTextGenerateMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList {
	return s.TextGenerateMultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetTraceabilityId() *int32 {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetChunks(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Chunks = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetExcerpt(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Excerpt = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetMultimodalMedias(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.MultimodalMedias = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetScore(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetSelect(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Select = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetTextGenerateMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.TextGenerateMultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetTraceabilityId(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResult) Validate() error {
	if s.MultimodalMedias != nil {
		for _, item := range s.MultimodalMedias {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextGenerateMultimodalMediaList != nil {
		for _, item := range s.TextGenerateMultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias struct {
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultMultimodalMedias) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList struct {
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 1
	End                 *int32                                                                                                                                                           `json:"End,omitempty" xml:"End,omitempty"`
	MultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList `json:"MultimodalMediaList,omitempty" xml:"MultimodalMediaList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) GetEnd() *int32 {
	return s.End
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) GetMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList {
	return s.MultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) GetStart() *int32 {
	return s.Start
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) SetEnd(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList {
	s.End = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) SetMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList {
	s.MultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) SetStart(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList {
	s.Start = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaList) Validate() error {
	if s.MultimodalMediaList != nil {
		for _, item := range s.MultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList struct {
	Article *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaList) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xxx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentExcerptResultSearchResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult struct {
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult struct {
	Article *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *string `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) GetTraceabilityId() *string {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) SetTraceabilityId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentImageSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult struct {
	// example:
	//
	// true
	GenerateFinished       *bool                                                                                                                        `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	NewsElementArticleList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList `json:"NewsElementArticleList,omitempty" xml:"NewsElementArticleList,omitempty" type:"Repeated"`
	// example:
	//
	// x
	TextGenerate *string `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) GetNewsElementArticleList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList {
	return s.NewsElementArticleList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) SetNewsElementArticleList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult {
	s.NewsElementArticleList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResult) Validate() error {
	if s.NewsElementArticleList != nil {
		for _, item := range s.NewsElementArticleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList struct {
	Article         *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle           `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	NewsElementList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList `json:"NewsElementList,omitempty" xml:"NewsElementList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	TextGenerate *string `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) GetNewsElementList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList {
	return s.NewsElementList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) SetNewsElementList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList {
	s.NewsElementList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleList) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.NewsElementList != nil {
		for _, item := range s.NewsElementList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle struct {
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2023-04-04 08:39:09
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 0.99
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetScore() *float32 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetSelect() *bool {
	return s.Select
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetScore(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetSelect(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Select = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList struct {
	// example:
	//
	// task-started
	Event *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Struct"`
	// example:
	//
	// xx
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// xx
	People *string `json:"People,omitempty" xml:"People,omitempty"`
	// example:
	//
	// 时间
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) GetEvent() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent {
	return s.Event
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) GetLocation() *string {
	return s.Location
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) GetPeople() *string {
	return s.People
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) GetTime() *string {
	return s.Time
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) SetEvent(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList {
	s.Event = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) SetLocation(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList {
	s.Location = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) SetPeople(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList {
	s.People = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) SetTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList {
	s.Time = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementList) Validate() error {
	if s.Event != nil {
		if err := s.Event.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent struct {
	CauseList   []*string `json:"CauseList,omitempty" xml:"CauseList,omitempty" type:"Repeated"`
	ProcessList []*string `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
	ResultList  []*string `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) GetCauseList() []*string {
	return s.CauseList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) GetProcessList() []*string {
	return s.ProcessList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) GetResultList() []*string {
	return s.ResultList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) SetCauseList(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent {
	s.CauseList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) SetProcessList(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent {
	s.ProcessList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) SetResultList(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent {
	s.ResultList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentNewsElementResultNewsElementArticleListNewsElementListEvent) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	// example:
	//
	// concise
	GenerateLevel              *string                                                                                                                           `json:"GenerateLevel,omitempty" xml:"GenerateLevel,omitempty"`
	GenerateTraceability       *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability         `json:"GenerateTraceability,omitempty" xml:"GenerateTraceability,omitempty" type:"Struct"`
	MultimodalSearchResultList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList `json:"MultimodalSearchResultList,omitempty" xml:"MultimodalSearchResultList,omitempty" type:"Repeated"`
	ReasonTextGenerate         *string                                                                                                                           `json:"ReasonTextGenerate,omitempty" xml:"ReasonTextGenerate,omitempty"`
	ReferenceList              []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList              `json:"ReferenceList,omitempty" xml:"ReferenceList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	TextGenerate                    *string                                                                                                                                `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
	TextGenerateMultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList `json:"TextGenerateMultimodalMediaList,omitempty" xml:"TextGenerateMultimodalMediaList,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetGenerateLevel() *string {
	return s.GenerateLevel
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetGenerateTraceability() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability {
	return s.GenerateTraceability
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetMultimodalSearchResultList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	return s.MultimodalSearchResultList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetReasonTextGenerate() *string {
	return s.ReasonTextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetReferenceList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	return s.ReferenceList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) GetTextGenerateMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList {
	return s.TextGenerateMultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetGenerateLevel(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.GenerateLevel = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetGenerateTraceability(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.GenerateTraceability = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetMultimodalSearchResultList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.MultimodalSearchResultList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetReasonTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.ReasonTextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetReferenceList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.ReferenceList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) SetTextGenerateMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult {
	s.TextGenerateMultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResult) Validate() error {
	if s.GenerateTraceability != nil {
		if err := s.GenerateTraceability.Validate(); err != nil {
			return err
		}
	}
	if s.MultimodalSearchResultList != nil {
		for _, item := range s.MultimodalSearchResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReferenceList != nil {
		for _, item := range s.ReferenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextGenerateMultimodalMediaList != nil {
		for _, item := range s.TextGenerateMultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability struct {
	Coordinates []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	// example:
	//
	// 0.9
	Duplicate *float64 `json:"Duplicate,omitempty" xml:"Duplicate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) GetCoordinates() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates {
	return s.Coordinates
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) GetDuplicate() *float64 {
	return s.Duplicate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) SetCoordinates(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability {
	s.Coordinates = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) SetDuplicate(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability {
	s.Duplicate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceability) Validate() error {
	if s.Coordinates != nil {
		for _, item := range s.Coordinates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates struct {
	GenerateCoordinate *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate `json:"GenerateCoordinate,omitempty" xml:"GenerateCoordinate,omitempty" type:"Struct"`
	NewsCoordinate     *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate     `json:"NewsCoordinate,omitempty" xml:"NewsCoordinate,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) GetGenerateCoordinate() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	return s.GenerateCoordinate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) GetNewsCoordinate() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate {
	return s.NewsCoordinate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) SetGenerateCoordinate(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates {
	s.GenerateCoordinate = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) SetNewsCoordinate(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates {
	s.NewsCoordinate = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinates) Validate() error {
	if s.GenerateCoordinate != nil {
		if err := s.GenerateCoordinate.Validate(); err != nil {
			return err
		}
	}
	if s.NewsCoordinate != nil {
		if err := s.NewsCoordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate struct {
	// example:
	//
	// 1
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 1
	Z *int32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetX() *int32 {
	return s.X
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetY() *int32 {
	return s.Y
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetZ() *int32 {
	return s.Z
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetX(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.X = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetY(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.Y = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetZ(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.Z = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesGenerateCoordinate) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate struct {
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 1
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 1
	Z *int32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) GetX() *int32 {
	return s.X
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) GetY() *int32 {
	return s.Y
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) GetZ() *int32 {
	return s.Z
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) SetX(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.X = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) SetY(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.Y = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) SetZ(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.Z = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultGenerateTraceabilityCoordinatesNewsCoordinate) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList struct {
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// xx
	SearchQuery  *string                                                                                                                                       `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// realtime
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 时间脉络-时间
	TimelineDateStr *string `json:"TimelineDateStr,omitempty" xml:"TimelineDateStr,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetSearchType() *string {
	return s.SearchType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetTimelineDateStr() *string {
	return s.TimelineDateStr
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetSearchQuery(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.SearchQuery = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetSearchType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.SearchType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetTimelineDateStr(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.TimelineDateStr = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultList) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.1
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultMultimodalSearchResultListSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList struct {
	Chunks []*string `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2023-04-04 08:39:09
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 0.99
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// 新华社
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *int32 `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetChunks() []*string {
	return s.Chunks
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetScore() *float32 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSelect() *bool {
	return s.Select
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSource() *string {
	return s.Source
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetTraceabilityId() *int32 {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetChunks(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Chunks = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetScore(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSelect(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Select = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Source = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetTraceabilityId(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultReferenceList) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList struct {
	// example:
	//
	// 1
	End                 *int32                                                                                                                                                    `json:"End,omitempty" xml:"End,omitempty"`
	MultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList `json:"MultimodalMediaList,omitempty" xml:"MultimodalMediaList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) GetEnd() *int32 {
	return s.End
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) GetMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList {
	return s.MultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) GetStart() *int32 {
	return s.Start
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) SetEnd(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList {
	s.End = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) SetMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList {
	s.MultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) SetStart(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList {
	s.Start = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaList) Validate() error {
	if s.MultimodalMediaList != nil {
		for _, item := range s.MultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList struct {
	Article *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaList) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextGenerateResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult struct {
	Current      *int32                                                                                                            `json:"Current,omitempty" xml:"Current,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	Size         *int32                                                                                                            `json:"Size,omitempty" xml:"Size,omitempty"`
	Total        *int32                                                                                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) GetCurrent() *int32 {
	return s.Current
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) GetSize() *int32 {
	return s.Size
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) GetTotal() *int32 {
	return s.Total
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) SetCurrent(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult {
	s.Current = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) SetSize(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult {
	s.Size = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) SetTotal(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult {
	s.Total = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult struct {
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-11-25 14:25:59
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// xxx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *string `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetTraceabilityId() *string {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetTraceabilityId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTextSearchResultSearchResult) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult struct {
	// example:
	//
	// true
	GenerateFinished           *bool                                                                                                                         `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	GenerateTraceability       *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability         `json:"GenerateTraceability,omitempty" xml:"GenerateTraceability,omitempty" type:"Struct"`
	MultimodalSearchResultList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList `json:"MultimodalSearchResultList,omitempty" xml:"MultimodalSearchResultList,omitempty" type:"Repeated"`
	ReasonTextGenerate         *string                                                                                                                       `json:"ReasonTextGenerate,omitempty" xml:"ReasonTextGenerate,omitempty"`
	ReferenceList              []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList              `json:"ReferenceList,omitempty" xml:"ReferenceList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	TextGenerate                    *string                                                                                                                            `json:"TextGenerate,omitempty" xml:"TextGenerate,omitempty"`
	TextGenerateMultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList `json:"TextGenerateMultimodalMediaList,omitempty" xml:"TextGenerateMultimodalMediaList,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetGenerateTraceability() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability {
	return s.GenerateTraceability
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetMultimodalSearchResultList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList {
	return s.MultimodalSearchResultList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetReasonTextGenerate() *string {
	return s.ReasonTextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetReferenceList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	return s.ReferenceList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetTextGenerate() *string {
	return s.TextGenerate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) GetTextGenerateMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList {
	return s.TextGenerateMultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetGenerateTraceability(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.GenerateTraceability = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetMultimodalSearchResultList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.MultimodalSearchResultList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetReasonTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.ReasonTextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetReferenceList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.ReferenceList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetTextGenerate(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.TextGenerate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) SetTextGenerateMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult {
	s.TextGenerateMultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResult) Validate() error {
	if s.GenerateTraceability != nil {
		if err := s.GenerateTraceability.Validate(); err != nil {
			return err
		}
	}
	if s.MultimodalSearchResultList != nil {
		for _, item := range s.MultimodalSearchResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReferenceList != nil {
		for _, item := range s.ReferenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextGenerateMultimodalMediaList != nil {
		for _, item := range s.TextGenerateMultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability struct {
	Coordinates []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	// example:
	//
	// 0.9
	Duplicate *float64 `json:"Duplicate,omitempty" xml:"Duplicate,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) GetCoordinates() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates {
	return s.Coordinates
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) GetDuplicate() *float64 {
	return s.Duplicate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) SetCoordinates(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability {
	s.Coordinates = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) SetDuplicate(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability {
	s.Duplicate = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceability) Validate() error {
	if s.Coordinates != nil {
		for _, item := range s.Coordinates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates struct {
	GenerateCoordinate *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate `json:"GenerateCoordinate,omitempty" xml:"GenerateCoordinate,omitempty" type:"Struct"`
	NewsCoordinate     *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate     `json:"NewsCoordinate,omitempty" xml:"NewsCoordinate,omitempty" type:"Struct"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) GetGenerateCoordinate() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	return s.GenerateCoordinate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) GetNewsCoordinate() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate {
	return s.NewsCoordinate
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) SetGenerateCoordinate(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates {
	s.GenerateCoordinate = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) SetNewsCoordinate(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates {
	s.NewsCoordinate = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinates) Validate() error {
	if s.GenerateCoordinate != nil {
		if err := s.GenerateCoordinate.Validate(); err != nil {
			return err
		}
	}
	if s.NewsCoordinate != nil {
		if err := s.NewsCoordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate struct {
	// example:
	//
	// 1
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 1
	Z *int32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetX() *int32 {
	return s.X
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetY() *int32 {
	return s.Y
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) GetZ() *int32 {
	return s.Z
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetX(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.X = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetY(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.Y = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) SetZ(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate {
	s.Z = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesGenerateCoordinate) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate struct {
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 1
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 1
	Z *int32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) GetX() *int32 {
	return s.X
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) GetY() *int32 {
	return s.Y
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) GetZ() *int32 {
	return s.Z
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) SetX(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.X = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) SetY(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.Y = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) SetZ(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate {
	s.Z = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultGenerateTraceabilityCoordinatesNewsCoordinate) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList struct {
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-11
	TimelineDateStr *string `json:"TimelineDateStr,omitempty" xml:"TimelineDateStr,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) GetTimelineDateStr() *string {
	return s.TimelineDateStr
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) SetTimelineDateStr(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList {
	s.TimelineDateStr = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultList) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.99
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultMultimodalSearchResultListSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList struct {
	Chunks []*string `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2023-04-04 08:39:09
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 0.99
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// 新华社
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *int32 `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetChunks() []*string {
	return s.Chunks
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetScore() *float32 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSelect() *bool {
	return s.Select
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSource() *string {
	return s.Source
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetTraceabilityId() *int32 {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetChunks(v []*string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Chunks = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetPubTime(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetScore(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSearchSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSearchSourceType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSelect(v bool) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Select = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSource(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Source = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetTraceabilityId(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultReferenceList) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList struct {
	// example:
	//
	// 1
	End                 *int32                                                                                                                                                `json:"End,omitempty" xml:"End,omitempty"`
	MultimodalMediaList []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList `json:"MultimodalMediaList,omitempty" xml:"MultimodalMediaList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) GetEnd() *int32 {
	return s.End
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) GetMultimodalMediaList() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList {
	return s.MultimodalMediaList
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) GetStart() *int32 {
	return s.Start
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) SetEnd(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList {
	s.End = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) SetMultimodalMediaList(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList {
	s.MultimodalMediaList = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) SetStart(v int32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList {
	s.Start = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaList) Validate() error {
	if s.MultimodalMediaList != nil {
		for _, item := range s.MultimodalMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList struct {
	Article *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) GetMediaType() *string {
	return s.MediaType
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) SetMediaType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList {
	s.MediaType = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaList) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xxxx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentTimelineResultTextGenerateMultimodalMediaListMultimodalMediaListArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult struct {
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult struct {
	Article   *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle     `json:"Article,omitempty" xml:"Article,omitempty" type:"Struct"`
	ClipInfos []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 1
	TraceabilityId *string `json:"TraceabilityId,omitempty" xml:"TraceabilityId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GetArticle() *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	return s.Article
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GetClipInfos() []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	return s.ClipInfos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) GetTraceabilityId() *string {
	return s.TraceabilityId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) SetArticle(v *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	s.Article = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) SetClipInfos(v []*RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	s.ClipInfos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) SetFileUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) SetTraceabilityId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult {
	s.TraceabilityId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResult) Validate() error {
	if s.Article != nil {
		if err := s.Article.Validate(); err != nil {
			return err
		}
	}
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle struct {
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetDocId(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetSearchSourceName(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetSummary(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetTitle(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) SetUrl(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultArticle) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos struct {
	// example:
	//
	// 1
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.8
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// asr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GetScore() *float64 {
	return s.Score
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GetText() *string {
	return s.Text
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) GetType() *string {
	return s.Type
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) SetFrom(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	s.From = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) SetScore(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) SetText(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	s.Text = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) SetTo(v float64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	s.To = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) SetType(v string) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos {
	s.Type = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextGeneratedContentVideoSearchResultSearchResultClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate struct {
	FirstTokenTime *float32 `json:"FirstTokenTime,omitempty" xml:"FirstTokenTime,omitempty"`
	OutputAvgTime  *float32 `json:"OutputAvgTime,omitempty" xml:"OutputAvgTime,omitempty"`
	SearchTime     *float32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
	Time           *float32 `json:"Time,omitempty" xml:"Time,omitempty"`
	TotalTokens    *int64   `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GetFirstTokenTime() *float32 {
	return s.FirstTokenTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GetOutputAvgTime() *float32 {
	return s.OutputAvgTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GetSearchTime() *float32 {
	return s.SearchTime
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GetTime() *float32 {
	return s.Time
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) SetFirstTokenTime(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	s.FirstTokenTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) SetOutputAvgTime(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	s.OutputAvgTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) SetSearchTime(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	s.SearchTime = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) SetTime(v float32) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	s.Time = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) SetTotalTokens(v int64) *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate {
	s.TotalTokens = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputAgentContextBizContextTokenCalculate) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputMessages struct {
	Clarifications *bool `json:"Clarifications,omitempty" xml:"Clarifications,omitempty"`
	// example:
	//
	// xx
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GenerateFinished *bool   `json:"GenerateFinished,omitempty" xml:"GenerateFinished,omitempty"`
	// example:
	//
	// xx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// generateStartStatement
	NodeCode      *string   `json:"NodeCode,omitempty" xml:"NodeCode,omitempty"`
	SearchQueries []*string `json:"SearchQueries,omitempty" xml:"SearchQueries,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	SearchQuery  *string                                                             `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	SearchResult []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessages) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessages) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetClarifications() *bool {
	return s.Clarifications
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetId() *string {
	return s.Id
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetNodeCode() *string {
	return s.NodeCode
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetSearchQueries() []*string {
	return s.SearchQueries
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) GetSearchResult() []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetClarifications(v bool) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.Clarifications = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetContent(v string) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetGenerateFinished(v bool) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.GenerateFinished = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetId(v string) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.Id = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetNodeCode(v string) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.NodeCode = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetSearchQueries(v []*string) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.SearchQueries = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetSearchQuery(v string) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.SearchQuery = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) SetSearchResult(v []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) *RunSearchGenerationResponseBodyPayloadOutputMessages {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessages) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult struct {
	Audios []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios `json:"Audios,omitempty" xml:"Audios,omitempty" type:"Repeated"`
	Images []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	MultimodalSearchQuery *string                                                                   `json:"MultimodalSearchQuery,omitempty" xml:"MultimodalSearchQuery,omitempty"`
	Texts                 []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	Videos                []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GetAudios() []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios {
	return s.Audios
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GetImages() []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages {
	return s.Images
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GetMultimodalSearchQuery() *string {
	return s.MultimodalSearchQuery
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GetTexts() []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts {
	return s.Texts
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) GetVideos() []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos {
	return s.Videos
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) SetAudios(v []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	s.Audios = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) SetImages(v []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	s.Images = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) SetMultimodalSearchQuery(v string) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	s.MultimodalSearchQuery = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) SetTexts(v []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	s.Texts = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) SetVideos(v []*RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult {
	s.Videos = v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResult) Validate() error {
	if s.Audios != nil {
		for _, item := range s.Audios {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Videos != nil {
		for _, item := range s.Videos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios struct {
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultAudios) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages struct {
	// example:
	//
	// xx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultImages) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts struct {
	// example:
	//
	// xx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) SetDocUuid(v string) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultTexts) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos struct {
	// example:
	//
	// 1
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) GetMediaId() *string {
	return s.MediaId
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) SetMediaId(v string) *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos {
	s.MediaId = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadOutputMessagesSearchResultVideos) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 3
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunSearchGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunSearchGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunSearchGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunSearchGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunSearchGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
