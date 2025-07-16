// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *CreateMessageRequest
	GetAssistantId() *string
	SetMessages(v []*CreateMessageRequestMessages) *CreateMessageRequest
	GetMessages() []*CreateMessageRequestMessages
	SetOriginalAssistantId(v string) *CreateMessageRequest
	GetOriginalAssistantId() *string
	SetSourceIdOfOriginalAssistantId(v string) *CreateMessageRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *CreateMessageRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetThreadId(v string) *CreateMessageRequest
	GetThreadId() *string
}

type CreateMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	Messages []*CreateMessageRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// assistantId
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// threadId123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *CreateMessageRequest) GetMessages() []*CreateMessageRequestMessages {
	return s.Messages
}

func (s *CreateMessageRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *CreateMessageRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *CreateMessageRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *CreateMessageRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateMessageRequest) SetAssistantId(v string) *CreateMessageRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateMessageRequest) SetMessages(v []*CreateMessageRequestMessages) *CreateMessageRequest {
	s.Messages = v
	return s
}

func (s *CreateMessageRequest) SetOriginalAssistantId(v string) *CreateMessageRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *CreateMessageRequest) SetSourceIdOfOriginalAssistantId(v string) *CreateMessageRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *CreateMessageRequest) SetSourceTypeOfOriginalAssistantId(v string) *CreateMessageRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *CreateMessageRequest) SetThreadId(v string) *CreateMessageRequest {
	s.ThreadId = &v
	return s
}

func (s *CreateMessageRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessages struct {
	Content *CreateMessageRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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

func (s CreateMessageRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessages) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessages) GetContent() *CreateMessageRequestMessagesContent {
	return s.Content
}

func (s *CreateMessageRequestMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *CreateMessageRequestMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *CreateMessageRequestMessages) GetRole() *string {
	return s.Role
}

func (s *CreateMessageRequestMessages) SetContent(v *CreateMessageRequestMessagesContent) *CreateMessageRequestMessages {
	s.Content = v
	return s
}

func (s *CreateMessageRequestMessages) SetContentDesc(v string) *CreateMessageRequestMessages {
	s.ContentDesc = &v
	return s
}

func (s *CreateMessageRequestMessages) SetCreateAt(v int64) *CreateMessageRequestMessages {
	s.CreateAt = &v
	return s
}

func (s *CreateMessageRequestMessages) SetRole(v string) *CreateMessageRequestMessages {
	s.Role = &v
	return s
}

func (s *CreateMessageRequestMessages) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContent struct {
	CardCallback *CreateMessageRequestMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *CreateMessageRequestMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *CreateMessageRequestMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *CreateMessageRequestMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *CreateMessageRequestMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *CreateMessageRequestMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMessageRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContent) GetCardCallback() *CreateMessageRequestMessagesContentCardCallback {
	return s.CardCallback
}

func (s *CreateMessageRequestMessagesContent) GetDingCard() *CreateMessageRequestMessagesContentDingCard {
	return s.DingCard
}

func (s *CreateMessageRequestMessagesContent) GetDingNormalCard() *CreateMessageRequestMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *CreateMessageRequestMessagesContent) GetMarkdown() *CreateMessageRequestMessagesContentMarkdown {
	return s.Markdown
}

func (s *CreateMessageRequestMessagesContent) GetStructView() *CreateMessageRequestMessagesContentStructView {
	return s.StructView
}

func (s *CreateMessageRequestMessagesContent) GetText() *CreateMessageRequestMessagesContentText {
	return s.Text
}

func (s *CreateMessageRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *CreateMessageRequestMessagesContent) SetCardCallback(v *CreateMessageRequestMessagesContentCardCallback) *CreateMessageRequestMessagesContent {
	s.CardCallback = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetDingCard(v *CreateMessageRequestMessagesContentDingCard) *CreateMessageRequestMessagesContent {
	s.DingCard = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetDingNormalCard(v *CreateMessageRequestMessagesContentDingNormalCard) *CreateMessageRequestMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetMarkdown(v *CreateMessageRequestMessagesContentMarkdown) *CreateMessageRequestMessagesContent {
	s.Markdown = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetStructView(v *CreateMessageRequestMessagesContentStructView) *CreateMessageRequestMessagesContent {
	s.StructView = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetText(v *CreateMessageRequestMessagesContentText) *CreateMessageRequestMessagesContent {
	s.Text = v
	return s
}

func (s *CreateMessageRequestMessagesContent) SetType(v string) *CreateMessageRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *CreateMessageRequestMessagesContent) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentCardCallback struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliding_messageId123
	RelatedMessageId *string `json:"relatedMessageId,omitempty" xml:"relatedMessageId,omitempty"`
}

func (s CreateMessageRequestMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *CreateMessageRequestMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *CreateMessageRequestMessagesContentCardCallback) SetContent(v string) *CreateMessageRequestMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *CreateMessageRequestMessagesContentCardCallback) SetRelatedMessageId(v string) *CreateMessageRequestMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *CreateMessageRequestMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingCard struct {
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

func (s CreateMessageRequestMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *CreateMessageRequestMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *CreateMessageRequestMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *CreateMessageRequestMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateMessageRequestMessagesContentDingCard) SetContent(v string) *CreateMessageRequestMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingCard) SetContentType(v string) *CreateMessageRequestMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingCard) SetFinished(v bool) *CreateMessageRequestMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingCard) SetTemplateId(v string) *CreateMessageRequestMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *CreateMessageRequestMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s CreateMessageRequestMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) GetCardData() *CreateMessageRequestMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) GetCardUpdateOptions() *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) SetCardData(v *CreateMessageRequestMessagesContentDingNormalCardCardData) *CreateMessageRequestMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) SetCardTemplateId(v string) *CreateMessageRequestMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) SetCardUpdateOptions(v *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) *CreateMessageRequestMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) *CreateMessageRequestMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *CreateMessageRequestMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCard) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingNormalCardCardData struct {
	// example:
	//
	// {}
	CardParamMap map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateMessageRequestMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardData) GetCardParamMap() map[string]interface{} {
	return s.CardParamMap
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardData) SetCardParamMap(v map[string]interface{}) *CreateMessageRequestMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *CreateMessageRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMessageRequestMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *CreateMessageRequestMessagesContentMarkdown) SetValue(v string) *CreateMessageRequestMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *CreateMessageRequestMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructView struct {
	Parts []*CreateMessageRequestMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s CreateMessageRequestMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructView) GetParts() []*CreateMessageRequestMessagesContentStructViewParts {
	return s.Parts
}

func (s *CreateMessageRequestMessagesContentStructView) SetParts(v []*CreateMessageRequestMessagesContentStructViewParts) *CreateMessageRequestMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructView) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *CreateMessageRequestMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                       `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *CreateMessageRequestMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *CreateMessageRequestMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *CreateMessageRequestMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *CreateMessageRequestMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMessageRequestMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetDataPart() *CreateMessageRequestMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetReasonPart() *CreateMessageRequestMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetRecommendPart() *CreateMessageRequestMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetReferencePart() *CreateMessageRequestMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetTextPart() *CreateMessageRequestMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *CreateMessageRequestMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetAppend(v bool) *CreateMessageRequestMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetDataPart(v *CreateMessageRequestMessagesContentStructViewPartsDataPart) *CreateMessageRequestMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetFinish(v bool) *CreateMessageRequestMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetPartDesc(v string) *CreateMessageRequestMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetPartId(v string) *CreateMessageRequestMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetReasonPart(v *CreateMessageRequestMessagesContentStructViewPartsReasonPart) *CreateMessageRequestMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetRecommendPart(v *CreateMessageRequestMessagesContentStructViewPartsRecommendPart) *CreateMessageRequestMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetReferencePart(v *CreateMessageRequestMessagesContentStructViewPartsReferencePart) *CreateMessageRequestMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetTextPart(v *CreateMessageRequestMessagesContentStructViewPartsTextPart) *CreateMessageRequestMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) SetType(v string) *CreateMessageRequestMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewParts) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *CreateMessageRequestMessagesContentStructViewPartsDataPart) SetData(v interface{}) *CreateMessageRequestMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReasonPart) SetReason(v string) *CreateMessageRequestMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPart) GetRecommends() []*CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) *CreateMessageRequestMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsReferencePart struct {
	References []*CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePart) GetReferences() []*CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePart) SetReferences(v []*CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) *CreateMessageRequestMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateMessageRequestMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *CreateMessageRequestMessagesContentStructViewPartsTextPart) SetText(v string) *CreateMessageRequestMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *CreateMessageRequestMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type CreateMessageRequestMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMessageRequestMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageRequestMessagesContentText) GoString() string {
	return s.String()
}

func (s *CreateMessageRequestMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *CreateMessageRequestMessagesContentText) SetValue(v string) *CreateMessageRequestMessagesContentText {
	s.Value = &v
	return s
}

func (s *CreateMessageRequestMessagesContentText) Validate() error {
	return dara.Validate(s)
}
