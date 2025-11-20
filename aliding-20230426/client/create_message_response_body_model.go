// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*CreateMessageResponseBodyMessages) *CreateMessageResponseBody
	GetMessages() []*CreateMessageResponseBodyMessages
	SetRequestId(v string) *CreateMessageResponseBody
	GetRequestId() *string
}

type CreateMessageResponseBody struct {
	Messages []*CreateMessageResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBody) GetMessages() []*CreateMessageResponseBodyMessages {
	return s.Messages
}

func (s *CreateMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMessageResponseBody) SetMessages(v []*CreateMessageResponseBodyMessages) *CreateMessageResponseBody {
	s.Messages = v
	return s
}

func (s *CreateMessageResponseBody) SetRequestId(v string) *CreateMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageResponseBody) Validate() error {
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

type CreateMessageResponseBodyMessages struct {
	Content *CreateMessageResponseBodyMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 这是一张小猫钓鱼图
	ContentDesc *string `json:"contentDesc,omitempty" xml:"contentDesc,omitempty"`
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

func (s CreateMessageResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessages) GetContent() *CreateMessageResponseBodyMessagesContent {
	return s.Content
}

func (s *CreateMessageResponseBodyMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *CreateMessageResponseBodyMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *CreateMessageResponseBodyMessages) GetId() *string {
	return s.Id
}

func (s *CreateMessageResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *CreateMessageResponseBodyMessages) GetRunId() *string {
	return s.RunId
}

func (s *CreateMessageResponseBodyMessages) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateMessageResponseBodyMessages) SetContent(v *CreateMessageResponseBodyMessagesContent) *CreateMessageResponseBodyMessages {
	s.Content = v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetContentDesc(v string) *CreateMessageResponseBodyMessages {
	s.ContentDesc = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetCreateAt(v int64) *CreateMessageResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetId(v string) *CreateMessageResponseBodyMessages {
	s.Id = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetRole(v string) *CreateMessageResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetRunId(v string) *CreateMessageResponseBodyMessages {
	s.RunId = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) SetThreadId(v string) *CreateMessageResponseBodyMessages {
	s.ThreadId = &v
	return s
}

func (s *CreateMessageResponseBodyMessages) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMessageResponseBodyMessagesContent struct {
	CardCallback *CreateMessageResponseBodyMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *CreateMessageResponseBodyMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *CreateMessageResponseBodyMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *CreateMessageResponseBodyMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *CreateMessageResponseBodyMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *CreateMessageResponseBodyMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContent) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContent) GetCardCallback() *CreateMessageResponseBodyMessagesContentCardCallback {
	return s.CardCallback
}

func (s *CreateMessageResponseBodyMessagesContent) GetDingCard() *CreateMessageResponseBodyMessagesContentDingCard {
	return s.DingCard
}

func (s *CreateMessageResponseBodyMessagesContent) GetDingNormalCard() *CreateMessageResponseBodyMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *CreateMessageResponseBodyMessagesContent) GetMarkdown() *CreateMessageResponseBodyMessagesContentMarkdown {
	return s.Markdown
}

func (s *CreateMessageResponseBodyMessagesContent) GetStructView() *CreateMessageResponseBodyMessagesContentStructView {
	return s.StructView
}

func (s *CreateMessageResponseBodyMessagesContent) GetText() *CreateMessageResponseBodyMessagesContentText {
	return s.Text
}

func (s *CreateMessageResponseBodyMessagesContent) GetType() *string {
	return s.Type
}

func (s *CreateMessageResponseBodyMessagesContent) SetCardCallback(v *CreateMessageResponseBodyMessagesContentCardCallback) *CreateMessageResponseBodyMessagesContent {
	s.CardCallback = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetDingCard(v *CreateMessageResponseBodyMessagesContentDingCard) *CreateMessageResponseBodyMessagesContent {
	s.DingCard = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetDingNormalCard(v *CreateMessageResponseBodyMessagesContentDingNormalCard) *CreateMessageResponseBodyMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetMarkdown(v *CreateMessageResponseBodyMessagesContentMarkdown) *CreateMessageResponseBodyMessagesContent {
	s.Markdown = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetStructView(v *CreateMessageResponseBodyMessagesContentStructView) *CreateMessageResponseBodyMessagesContent {
	s.StructView = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetText(v *CreateMessageResponseBodyMessagesContentText) *CreateMessageResponseBodyMessagesContent {
	s.Text = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) SetType(v string) *CreateMessageResponseBodyMessagesContent {
	s.Type = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContent) Validate() error {
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

type CreateMessageResponseBodyMessagesContentCardCallback struct {
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// aliding_messageId123
	RelatedMessageId *string `json:"relatedMessageId,omitempty" xml:"relatedMessageId,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *CreateMessageResponseBodyMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *CreateMessageResponseBodyMessagesContentCardCallback) SetContent(v string) *CreateMessageResponseBodyMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentCardCallback) SetRelatedMessageId(v string) *CreateMessageResponseBodyMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentDingCard struct {
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

func (s CreateMessageResponseBodyMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) SetContent(v string) *CreateMessageResponseBodyMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) SetContentType(v string) *CreateMessageResponseBodyMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) SetFinished(v bool) *CreateMessageResponseBodyMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) SetTemplateId(v string) *CreateMessageResponseBodyMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *CreateMessageResponseBodyMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) GetCardData() *CreateMessageResponseBodyMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) GetCardUpdateOptions() *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) SetCardData(v *CreateMessageResponseBodyMessagesContentDingNormalCardCardData) *CreateMessageResponseBodyMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) SetCardTemplateId(v string) *CreateMessageResponseBodyMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) SetCardUpdateOptions(v *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) *CreateMessageResponseBodyMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) *CreateMessageResponseBodyMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *CreateMessageResponseBodyMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCard) Validate() error {
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

type CreateMessageResponseBodyMessagesContentDingNormalCardCardData struct {
	// example:
	//
	// {}
	CardParamMap map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardData) GetCardParamMap() map[string]interface{} {
	return s.CardParamMap
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardData) SetCardParamMap(v map[string]interface{}) *CreateMessageResponseBodyMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	if s.PullConfig != nil {
		if err := s.PullConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *CreateMessageResponseBodyMessagesContentMarkdown) SetValue(v string) *CreateMessageResponseBodyMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentStructView struct {
	Parts []*CreateMessageResponseBodyMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s CreateMessageResponseBodyMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructView) GetParts() []*CreateMessageResponseBodyMessagesContentStructViewParts {
	return s.Parts
}

func (s *CreateMessageResponseBodyMessagesContentStructView) SetParts(v []*CreateMessageResponseBodyMessagesContentStructViewParts) *CreateMessageResponseBodyMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructView) Validate() error {
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

type CreateMessageResponseBodyMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                            `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetDataPart() *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetReasonPart() *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetRecommendPart() *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetReferencePart() *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetTextPart() *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetAppend(v bool) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetDataPart(v *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetFinish(v bool) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetPartDesc(v string) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetPartId(v string) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetReasonPart(v *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetRecommendPart(v *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetReferencePart(v *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetTextPart(v *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) SetType(v string) *CreateMessageResponseBodyMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewParts) Validate() error {
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

type CreateMessageResponseBodyMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) SetData(v interface{}) *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) SetReason(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) GetRecommends() []*CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPart) Validate() error {
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

type CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart struct {
	References []*CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) GetReferences() []*CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) SetReferences(v []*CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePart) Validate() error {
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

type CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) SetText(v string) *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageResponseBodyMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMessageResponseBodyMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponseBodyMessagesContentText) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBodyMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *CreateMessageResponseBodyMessagesContentText) SetValue(v string) *CreateMessageResponseBodyMessagesContentText {
	s.Value = &v
	return s
}

func (s *CreateMessageResponseBodyMessagesContentText) Validate() error {
	return dara.Validate(s)
}
