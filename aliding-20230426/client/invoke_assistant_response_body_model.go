// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*InvokeAssistantResponseBodyMessages) *InvokeAssistantResponseBody
	GetMessages() []*InvokeAssistantResponseBodyMessages
	SetRequestId(v string) *InvokeAssistantResponseBody
	GetRequestId() *string
	SetSessionId(v string) *InvokeAssistantResponseBody
	GetSessionId() *string
	SetStreamEnd(v bool) *InvokeAssistantResponseBody
	GetStreamEnd() *bool
}

type InvokeAssistantResponseBody struct {
	Messages []*InvokeAssistantResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// sessionId1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// true
	StreamEnd *bool `json:"streamEnd,omitempty" xml:"streamEnd,omitempty"`
}

func (s InvokeAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBody) GetMessages() []*InvokeAssistantResponseBodyMessages {
	return s.Messages
}

func (s *InvokeAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeAssistantResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *InvokeAssistantResponseBody) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *InvokeAssistantResponseBody) SetMessages(v []*InvokeAssistantResponseBodyMessages) *InvokeAssistantResponseBody {
	s.Messages = v
	return s
}

func (s *InvokeAssistantResponseBody) SetRequestId(v string) *InvokeAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeAssistantResponseBody) SetSessionId(v string) *InvokeAssistantResponseBody {
	s.SessionId = &v
	return s
}

func (s *InvokeAssistantResponseBody) SetStreamEnd(v bool) *InvokeAssistantResponseBody {
	s.StreamEnd = &v
	return s
}

func (s *InvokeAssistantResponseBody) Validate() error {
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

type InvokeAssistantResponseBodyMessages struct {
	Content *InvokeAssistantResponseBodyMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s InvokeAssistantResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessages) GetContent() *InvokeAssistantResponseBodyMessagesContent {
	return s.Content
}

func (s *InvokeAssistantResponseBodyMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *InvokeAssistantResponseBodyMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *InvokeAssistantResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *InvokeAssistantResponseBodyMessages) SetContent(v *InvokeAssistantResponseBodyMessagesContent) *InvokeAssistantResponseBodyMessages {
	s.Content = v
	return s
}

func (s *InvokeAssistantResponseBodyMessages) SetContentDesc(v string) *InvokeAssistantResponseBodyMessages {
	s.ContentDesc = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessages) SetCreateAt(v int64) *InvokeAssistantResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessages) SetRole(v string) *InvokeAssistantResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessages) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeAssistantResponseBodyMessagesContent struct {
	CardCallback *InvokeAssistantResponseBodyMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *InvokeAssistantResponseBodyMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *InvokeAssistantResponseBodyMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *InvokeAssistantResponseBodyMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *InvokeAssistantResponseBodyMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *InvokeAssistantResponseBodyMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContent) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetCardCallback() *InvokeAssistantResponseBodyMessagesContentCardCallback {
	return s.CardCallback
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetDingCard() *InvokeAssistantResponseBodyMessagesContentDingCard {
	return s.DingCard
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetDingNormalCard() *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetMarkdown() *InvokeAssistantResponseBodyMessagesContentMarkdown {
	return s.Markdown
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetStructView() *InvokeAssistantResponseBodyMessagesContentStructView {
	return s.StructView
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetText() *InvokeAssistantResponseBodyMessagesContentText {
	return s.Text
}

func (s *InvokeAssistantResponseBodyMessagesContent) GetType() *string {
	return s.Type
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetCardCallback(v *InvokeAssistantResponseBodyMessagesContentCardCallback) *InvokeAssistantResponseBodyMessagesContent {
	s.CardCallback = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetDingCard(v *InvokeAssistantResponseBodyMessagesContentDingCard) *InvokeAssistantResponseBodyMessagesContent {
	s.DingCard = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetDingNormalCard(v *InvokeAssistantResponseBodyMessagesContentDingNormalCard) *InvokeAssistantResponseBodyMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetMarkdown(v *InvokeAssistantResponseBodyMessagesContentMarkdown) *InvokeAssistantResponseBodyMessagesContent {
	s.Markdown = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetStructView(v *InvokeAssistantResponseBodyMessagesContentStructView) *InvokeAssistantResponseBodyMessagesContent {
	s.StructView = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetText(v *InvokeAssistantResponseBodyMessagesContentText) *InvokeAssistantResponseBodyMessagesContent {
	s.Text = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) SetType(v string) *InvokeAssistantResponseBodyMessagesContent {
	s.Type = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContent) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentCardCallback struct {
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// aliding_messageId123
	RelatedMessageId *string `json:"relatedMessageId,omitempty" xml:"relatedMessageId,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *InvokeAssistantResponseBodyMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *InvokeAssistantResponseBodyMessagesContentCardCallback) SetContent(v string) *InvokeAssistantResponseBodyMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentCardCallback) SetRelatedMessageId(v string) *InvokeAssistantResponseBodyMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentDingCard struct {
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

func (s InvokeAssistantResponseBodyMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) SetContent(v string) *InvokeAssistantResponseBodyMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) SetContentType(v string) *InvokeAssistantResponseBodyMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) SetFinished(v bool) *InvokeAssistantResponseBodyMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) SetTemplateId(v string) *InvokeAssistantResponseBodyMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) GetCardData() *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) GetCardUpdateOptions() *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) SetCardData(v *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) SetCardTemplateId(v string) *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) SetCardUpdateOptions(v *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *InvokeAssistantResponseBodyMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCard) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData struct {
	CardParamMap interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) GetCardParamMap() interface{} {
	return s.CardParamMap
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) SetCardParamMap(v interface{}) *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	if s.PullConfig != nil {
		if err := s.PullConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *InvokeAssistantResponseBodyMessagesContentMarkdown) SetValue(v string) *InvokeAssistantResponseBodyMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentStructView struct {
	Parts []*InvokeAssistantResponseBodyMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructView) GetParts() []*InvokeAssistantResponseBodyMessagesContentStructViewParts {
	return s.Parts
}

func (s *InvokeAssistantResponseBodyMessagesContentStructView) SetParts(v []*InvokeAssistantResponseBodyMessagesContentStructViewParts) *InvokeAssistantResponseBodyMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructView) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                              `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetDataPart() *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetReasonPart() *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetRecommendPart() *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetReferencePart() *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetTextPart() *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetAppend(v bool) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetDataPart(v *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetFinish(v bool) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetPartDesc(v string) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetPartId(v string) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetReasonPart(v *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetRecommendPart(v *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetReferencePart(v *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetTextPart(v *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) SetType(v string) *InvokeAssistantResponseBodyMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewParts) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) SetData(v interface{}) *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) SetReason(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) GetRecommends() []*InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPart) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart struct {
	References []*InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) GetReferences() []*InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) SetReferences(v []*InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePart) Validate() error {
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

type InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) SetText(v string) *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantResponseBodyMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAssistantResponseBodyMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponseBodyMessagesContentText) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponseBodyMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *InvokeAssistantResponseBodyMessagesContentText) SetValue(v string) *InvokeAssistantResponseBodyMessagesContentText {
	s.Value = &v
	return s
}

func (s *InvokeAssistantResponseBodyMessagesContentText) Validate() error {
	return dara.Validate(s)
}
