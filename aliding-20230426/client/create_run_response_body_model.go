// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*CreateRunResponseBodyMessages) *CreateRunResponseBody
	GetMessages() []*CreateRunResponseBodyMessages
	SetRequestId(v string) *CreateRunResponseBody
	GetRequestId() *string
	SetRun(v *CreateRunResponseBodyRun) *CreateRunResponseBody
	GetRun() *CreateRunResponseBodyRun
}

type CreateRunResponseBody struct {
	Messages []*CreateRunResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Run       *CreateRunResponseBodyRun `json:"run,omitempty" xml:"run,omitempty" type:"Struct"`
}

func (s CreateRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBody) GetMessages() []*CreateRunResponseBodyMessages {
	return s.Messages
}

func (s *CreateRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRunResponseBody) GetRun() *CreateRunResponseBodyRun {
	return s.Run
}

func (s *CreateRunResponseBody) SetMessages(v []*CreateRunResponseBodyMessages) *CreateRunResponseBody {
	s.Messages = v
	return s
}

func (s *CreateRunResponseBody) SetRequestId(v string) *CreateRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRunResponseBody) SetRun(v *CreateRunResponseBodyRun) *CreateRunResponseBody {
	s.Run = v
	return s
}

func (s *CreateRunResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Run != nil {
		if err := s.Run.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessages struct {
	Content *CreateRunResponseBodyMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 这是一张小猫钓鱼图
	ContentDesc   *string                                     `json:"contentDesc,omitempty" xml:"contentDesc,omitempty"`
	ContentStruct *CreateRunResponseBodyMessagesContentStruct `json:"contentStruct,omitempty" xml:"contentStruct,omitempty" type:"Struct"`
	// example:
	//
	// 1642448000000
	CreateAt *int64 `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// example:
	//
	// messageId1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// runId1
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// threadId1
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateRunResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessages) GetContent() *CreateRunResponseBodyMessagesContent {
	return s.Content
}

func (s *CreateRunResponseBodyMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *CreateRunResponseBodyMessages) GetContentStruct() *CreateRunResponseBodyMessagesContentStruct {
	return s.ContentStruct
}

func (s *CreateRunResponseBodyMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *CreateRunResponseBodyMessages) GetId() *string {
	return s.Id
}

func (s *CreateRunResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *CreateRunResponseBodyMessages) GetRunId() *string {
	return s.RunId
}

func (s *CreateRunResponseBodyMessages) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateRunResponseBodyMessages) SetContent(v *CreateRunResponseBodyMessagesContent) *CreateRunResponseBodyMessages {
	s.Content = v
	return s
}

func (s *CreateRunResponseBodyMessages) SetContentDesc(v string) *CreateRunResponseBodyMessages {
	s.ContentDesc = &v
	return s
}

func (s *CreateRunResponseBodyMessages) SetContentStruct(v *CreateRunResponseBodyMessagesContentStruct) *CreateRunResponseBodyMessages {
	s.ContentStruct = v
	return s
}

func (s *CreateRunResponseBodyMessages) SetCreateAt(v int64) *CreateRunResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *CreateRunResponseBodyMessages) SetId(v string) *CreateRunResponseBodyMessages {
	s.Id = &v
	return s
}

func (s *CreateRunResponseBodyMessages) SetRole(v string) *CreateRunResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *CreateRunResponseBodyMessages) SetRunId(v string) *CreateRunResponseBodyMessages {
	s.RunId = &v
	return s
}

func (s *CreateRunResponseBodyMessages) SetThreadId(v string) *CreateRunResponseBodyMessages {
	s.ThreadId = &v
	return s
}

func (s *CreateRunResponseBodyMessages) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	if s.ContentStruct != nil {
		if err := s.ContentStruct.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContent struct {
	CardCallback *CreateRunResponseBodyMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *CreateRunResponseBodyMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *CreateRunResponseBodyMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *CreateRunResponseBodyMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *CreateRunResponseBodyMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *CreateRunResponseBodyMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRunResponseBodyMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContent) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContent) GetCardCallback() *CreateRunResponseBodyMessagesContentCardCallback {
	return s.CardCallback
}

func (s *CreateRunResponseBodyMessagesContent) GetDingCard() *CreateRunResponseBodyMessagesContentDingCard {
	return s.DingCard
}

func (s *CreateRunResponseBodyMessagesContent) GetDingNormalCard() *CreateRunResponseBodyMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *CreateRunResponseBodyMessagesContent) GetMarkdown() *CreateRunResponseBodyMessagesContentMarkdown {
	return s.Markdown
}

func (s *CreateRunResponseBodyMessagesContent) GetStructView() *CreateRunResponseBodyMessagesContentStructView {
	return s.StructView
}

func (s *CreateRunResponseBodyMessagesContent) GetText() *CreateRunResponseBodyMessagesContentText {
	return s.Text
}

func (s *CreateRunResponseBodyMessagesContent) GetType() *string {
	return s.Type
}

func (s *CreateRunResponseBodyMessagesContent) SetCardCallback(v *CreateRunResponseBodyMessagesContentCardCallback) *CreateRunResponseBodyMessagesContent {
	s.CardCallback = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetDingCard(v *CreateRunResponseBodyMessagesContentDingCard) *CreateRunResponseBodyMessagesContent {
	s.DingCard = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetDingNormalCard(v *CreateRunResponseBodyMessagesContentDingNormalCard) *CreateRunResponseBodyMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetMarkdown(v *CreateRunResponseBodyMessagesContentMarkdown) *CreateRunResponseBodyMessagesContent {
	s.Markdown = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetStructView(v *CreateRunResponseBodyMessagesContentStructView) *CreateRunResponseBodyMessagesContent {
	s.StructView = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetText(v *CreateRunResponseBodyMessagesContentText) *CreateRunResponseBodyMessagesContent {
	s.Text = v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) SetType(v string) *CreateRunResponseBodyMessagesContent {
	s.Type = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContent) Validate() error {
	if s.CardCallback != nil {
		if err := s.CardCallback.Validate(); err != nil {
			return err
		}
	}
	if s.DingCard != nil {
		if err := s.DingCard.Validate(); err != nil {
			return err
		}
	}
	if s.DingNormalCard != nil {
		if err := s.DingNormalCard.Validate(); err != nil {
			return err
		}
	}
	if s.Markdown != nil {
		if err := s.Markdown.Validate(); err != nil {
			return err
		}
	}
	if s.StructView != nil {
		if err := s.StructView.Validate(); err != nil {
			return err
		}
	}
	if s.Text != nil {
		if err := s.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentCardCallback struct {
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// aliding_messageId123
	RelatedMessageId *string `json:"relatedMessageId,omitempty" xml:"relatedMessageId,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *CreateRunResponseBodyMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *CreateRunResponseBodyMessagesContentCardCallback) SetContent(v string) *CreateRunResponseBodyMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentCardCallback) SetRelatedMessageId(v string) *CreateRunResponseBodyMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentDingCard struct {
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// basic_card_schema
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// example:
	//
	// true
	Finished *bool `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// templateId123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *CreateRunResponseBodyMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *CreateRunResponseBodyMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *CreateRunResponseBodyMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateRunResponseBodyMessagesContentDingCard) SetContent(v string) *CreateRunResponseBodyMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingCard) SetContentType(v string) *CreateRunResponseBodyMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingCard) SetFinished(v bool) *CreateRunResponseBodyMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingCard) SetTemplateId(v string) *CreateRunResponseBodyMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *CreateRunResponseBodyMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) GetCardData() *CreateRunResponseBodyMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) GetCardUpdateOptions() *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) SetCardData(v *CreateRunResponseBodyMessagesContentDingNormalCardCardData) *CreateRunResponseBodyMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) SetCardTemplateId(v string) *CreateRunResponseBodyMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) SetCardUpdateOptions(v *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) *CreateRunResponseBodyMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) *CreateRunResponseBodyMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *CreateRunResponseBodyMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCard) Validate() error {
	if s.CardData != nil {
		if err := s.CardData.Validate(); err != nil {
			return err
		}
	}
	if s.CardUpdateOptions != nil {
		if err := s.CardUpdateOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DynamicDataSourceConfigs != nil {
		for _, item := range s.DynamicDataSourceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentDingNormalCardCardData struct {
	CardParamMap interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardData) GetCardParamMap() interface{} {
	return s.CardParamMap
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardData) SetCardParamMap(v interface{}) *CreateRunResponseBodyMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
	// example:
	//
	// {}
	ConstParams map[string]interface{} `json:"constParams,omitempty" xml:"constParams,omitempty"`
	// example:
	//
	// dynamicDataSourceId1
	DynamicDataSourceId *string `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	// example:
	//
	// {}
	PullConfig *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	if s.PullConfig != nil {
		if err := s.PullConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
	// example:
	//
	// 3
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// NONE
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	// example:
	//
	// SECONDS
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *CreateRunResponseBodyMessagesContentMarkdown) SetValue(v string) *CreateRunResponseBodyMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructView struct {
	Parts []*CreateRunResponseBodyMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructView) GetParts() []*CreateRunResponseBodyMessagesContentStructViewParts {
	return s.Parts
}

func (s *CreateRunResponseBodyMessagesContentStructView) SetParts(v []*CreateRunResponseBodyMessagesContentStructViewParts) *CreateRunResponseBodyMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructView) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *CreateRunResponseBodyMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                        `json:"finish,omitempty" xml:"finish,omitempty"`
	// example:
	//
	// 这是正文内容部分
	PartDesc *string `json:"partDesc,omitempty" xml:"partDesc,omitempty"`
	// example:
	//
	// artifactId123
	PartId *string `json:"partId,omitempty" xml:"partId,omitempty"`
	// example:
	//
	// {}
	ReasonPart *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *CreateRunResponseBodyMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetDataPart() *CreateRunResponseBodyMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetReasonPart() *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetRecommendPart() *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetReferencePart() *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetTextPart() *CreateRunResponseBodyMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetAppend(v bool) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetDataPart(v *CreateRunResponseBodyMessagesContentStructViewPartsDataPart) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetFinish(v bool) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetPartDesc(v string) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetPartId(v string) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetReasonPart(v *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetRecommendPart(v *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetReferencePart(v *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetTextPart(v *CreateRunResponseBodyMessagesContentStructViewPartsTextPart) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) SetType(v string) *CreateRunResponseBodyMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewParts) Validate() error {
	if s.DataPart != nil {
		if err := s.DataPart.Validate(); err != nil {
			return err
		}
	}
	if s.ReasonPart != nil {
		if err := s.ReasonPart.Validate(); err != nil {
			return err
		}
	}
	if s.RecommendPart != nil {
		if err := s.RecommendPart.Validate(); err != nil {
			return err
		}
	}
	if s.ReferencePart != nil {
		if err := s.ReferencePart.Validate(); err != nil {
			return err
		}
	}
	if s.TextPart != nil {
		if err := s.TextPart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsDataPart) SetData(v interface{}) *CreateRunResponseBodyMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) SetReason(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) GetRecommends() []*CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPart) Validate() error {
	if s.Recommends != nil {
		for _, item := range s.Recommends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructViewPartsReferencePart struct {
	References []*CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) GetReferences() []*CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) SetReferences(v []*CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePart) Validate() error {
	if s.References != nil {
		for _, item := range s.References {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences struct {
	// example:
	//
	// 0
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// mcp是....
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ata
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	SourceIcon *string `json:"sourceIcon,omitempty" xml:"sourceIcon,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// 《mcp原理介绍》
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://taobao.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsTextPart) SetText(v string) *CreateRunResponseBodyMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentText) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *CreateRunResponseBodyMessagesContentText) SetValue(v string) *CreateRunResponseBodyMessagesContentText {
	s.Value = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentText) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStruct struct {
	Parts []*CreateRunResponseBodyMessagesContentStructParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStruct) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStruct) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStruct) GetParts() []*CreateRunResponseBodyMessagesContentStructParts {
	return s.Parts
}

func (s *CreateRunResponseBodyMessagesContentStruct) SetParts(v []*CreateRunResponseBodyMessagesContentStructParts) *CreateRunResponseBodyMessagesContentStruct {
	s.Parts = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStruct) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *CreateRunResponseBodyMessagesContentStructPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                    `json:"finish,omitempty" xml:"finish,omitempty"`
	// example:
	//
	// 这是正文内容部分
	PartDesc *string `json:"partDesc,omitempty" xml:"partDesc,omitempty"`
	// example:
	//
	// artifactId123
	PartId *string `json:"partId,omitempty" xml:"partId,omitempty"`
	// example:
	//
	// {}
	ReasonPart *CreateRunResponseBodyMessagesContentStructPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *CreateRunResponseBodyMessagesContentStructPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *CreateRunResponseBodyMessagesContentStructPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *CreateRunResponseBodyMessagesContentStructPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructParts) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructParts) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetAppend() *bool {
	return s.Append
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetDataPart() *CreateRunResponseBodyMessagesContentStructPartsDataPart {
	return s.DataPart
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetFinish() *bool {
	return s.Finish
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetPartId() *string {
	return s.PartId
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetReasonPart() *CreateRunResponseBodyMessagesContentStructPartsReasonPart {
	return s.ReasonPart
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetRecommendPart() *CreateRunResponseBodyMessagesContentStructPartsRecommendPart {
	return s.RecommendPart
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetReferencePart() *CreateRunResponseBodyMessagesContentStructPartsReferencePart {
	return s.ReferencePart
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetTextPart() *CreateRunResponseBodyMessagesContentStructPartsTextPart {
	return s.TextPart
}

func (s *CreateRunResponseBodyMessagesContentStructParts) GetType() *string {
	return s.Type
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetAppend(v bool) *CreateRunResponseBodyMessagesContentStructParts {
	s.Append = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetDataPart(v *CreateRunResponseBodyMessagesContentStructPartsDataPart) *CreateRunResponseBodyMessagesContentStructParts {
	s.DataPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetFinish(v bool) *CreateRunResponseBodyMessagesContentStructParts {
	s.Finish = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetPartDesc(v string) *CreateRunResponseBodyMessagesContentStructParts {
	s.PartDesc = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetPartId(v string) *CreateRunResponseBodyMessagesContentStructParts {
	s.PartId = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetReasonPart(v *CreateRunResponseBodyMessagesContentStructPartsReasonPart) *CreateRunResponseBodyMessagesContentStructParts {
	s.ReasonPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetRecommendPart(v *CreateRunResponseBodyMessagesContentStructPartsRecommendPart) *CreateRunResponseBodyMessagesContentStructParts {
	s.RecommendPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetReferencePart(v *CreateRunResponseBodyMessagesContentStructPartsReferencePart) *CreateRunResponseBodyMessagesContentStructParts {
	s.ReferencePart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetTextPart(v *CreateRunResponseBodyMessagesContentStructPartsTextPart) *CreateRunResponseBodyMessagesContentStructParts {
	s.TextPart = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) SetType(v string) *CreateRunResponseBodyMessagesContentStructParts {
	s.Type = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructParts) Validate() error {
	if s.DataPart != nil {
		if err := s.DataPart.Validate(); err != nil {
			return err
		}
	}
	if s.ReasonPart != nil {
		if err := s.ReasonPart.Validate(); err != nil {
			return err
		}
	}
	if s.RecommendPart != nil {
		if err := s.RecommendPart.Validate(); err != nil {
			return err
		}
	}
	if s.ReferencePart != nil {
		if err := s.ReferencePart.Validate(); err != nil {
			return err
		}
	}
	if s.TextPart != nil {
		if err := s.TextPart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsDataPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *CreateRunResponseBodyMessagesContentStructPartsDataPart) SetData(v interface{}) *CreateRunResponseBodyMessagesContentStructPartsDataPart {
	s.Data = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsReasonPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReasonPart) SetReason(v string) *CreateRunResponseBodyMessagesContentStructPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructPartsRecommendPart struct {
	Recommends []*CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPart) GetRecommends() []*CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPart) SetRecommends(v []*CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) *CreateRunResponseBodyMessagesContentStructPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPart) Validate() error {
	if s.Recommends != nil {
		for _, item := range s.Recommends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) SetMobileUrl(v string) *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) SetText(v string) *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) SetUrl(v string) *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructPartsReferencePart struct {
	References []*CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsReferencePart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePart) GetReferences() []*CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	return s.References
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePart) SetReferences(v []*CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) *CreateRunResponseBodyMessagesContentStructPartsReferencePart {
	s.References = v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePart) Validate() error {
	if s.References != nil {
		for _, item := range s.References {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences struct {
	// example:
	//
	// 0
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// mcp是....
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ata
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	SourceIcon *string `json:"sourceIcon,omitempty" xml:"sourceIcon,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// 《mcp原理介绍》
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://taobao.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetIndex(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetName(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetSourceCode(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetSourceIcon(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetSummary(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetTitle(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) SetUrl(v string) *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyMessagesContentStructPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateRunResponseBodyMessagesContentStructPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyMessagesContentStructPartsTextPart) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyMessagesContentStructPartsTextPart) GetText() *string {
	return s.Text
}

func (s *CreateRunResponseBodyMessagesContentStructPartsTextPart) SetText(v string) *CreateRunResponseBodyMessagesContentStructPartsTextPart {
	s.Text = &v
	return s
}

func (s *CreateRunResponseBodyMessagesContentStructPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type CreateRunResponseBodyRun struct {
	CancelledAt  *int64  `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64  `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreateAt     *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	ExpiresAt    *int64  `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	StartedAt    *int64  `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateRunResponseBodyRun) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBodyRun) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBodyRun) GetCancelledAt() *int64 {
	return s.CancelledAt
}

func (s *CreateRunResponseBodyRun) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *CreateRunResponseBodyRun) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *CreateRunResponseBodyRun) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *CreateRunResponseBodyRun) GetFailedAt() *int64 {
	return s.FailedAt
}

func (s *CreateRunResponseBodyRun) GetId() *string {
	return s.Id
}

func (s *CreateRunResponseBodyRun) GetLastErrorMsg() *string {
	return s.LastErrorMsg
}

func (s *CreateRunResponseBodyRun) GetStartedAt() *int64 {
	return s.StartedAt
}

func (s *CreateRunResponseBodyRun) GetStatus() *string {
	return s.Status
}

func (s *CreateRunResponseBodyRun) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateRunResponseBodyRun) SetCancelledAt(v int64) *CreateRunResponseBodyRun {
	s.CancelledAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetCompletedAt(v int64) *CreateRunResponseBodyRun {
	s.CompletedAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetCreateAt(v int64) *CreateRunResponseBodyRun {
	s.CreateAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetExpiresAt(v int64) *CreateRunResponseBodyRun {
	s.ExpiresAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetFailedAt(v int64) *CreateRunResponseBodyRun {
	s.FailedAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetId(v string) *CreateRunResponseBodyRun {
	s.Id = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetLastErrorMsg(v string) *CreateRunResponseBodyRun {
	s.LastErrorMsg = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetStartedAt(v int64) *CreateRunResponseBodyRun {
	s.StartedAt = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetStatus(v string) *CreateRunResponseBodyRun {
	s.Status = &v
	return s
}

func (s *CreateRunResponseBodyRun) SetThreadId(v string) *CreateRunResponseBodyRun {
	s.ThreadId = &v
	return s
}

func (s *CreateRunResponseBodyRun) Validate() error {
	return dara.Validate(s)
}
