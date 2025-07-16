// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *InvokeAssistantRequest
	GetAssistantId() *string
	SetMessages(v []*InvokeAssistantRequestMessages) *InvokeAssistantRequest
	GetMessages() []*InvokeAssistantRequestMessages
	SetOriginalAssistantId(v string) *InvokeAssistantRequest
	GetOriginalAssistantId() *string
	SetSessionId(v string) *InvokeAssistantRequest
	GetSessionId() *string
	SetSourceIdOfOriginalAssistantId(v string) *InvokeAssistantRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *InvokeAssistantRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetStream(v bool) *InvokeAssistantRequest
	GetStream() *bool
}

type InvokeAssistantRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	Messages []*InvokeAssistantRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// assistantId2
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// sessionId1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s InvokeAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequest) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *InvokeAssistantRequest) GetMessages() []*InvokeAssistantRequestMessages {
	return s.Messages
}

func (s *InvokeAssistantRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *InvokeAssistantRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *InvokeAssistantRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *InvokeAssistantRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *InvokeAssistantRequest) GetStream() *bool {
	return s.Stream
}

func (s *InvokeAssistantRequest) SetAssistantId(v string) *InvokeAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *InvokeAssistantRequest) SetMessages(v []*InvokeAssistantRequestMessages) *InvokeAssistantRequest {
	s.Messages = v
	return s
}

func (s *InvokeAssistantRequest) SetOriginalAssistantId(v string) *InvokeAssistantRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *InvokeAssistantRequest) SetSessionId(v string) *InvokeAssistantRequest {
	s.SessionId = &v
	return s
}

func (s *InvokeAssistantRequest) SetSourceIdOfOriginalAssistantId(v string) *InvokeAssistantRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *InvokeAssistantRequest) SetSourceTypeOfOriginalAssistantId(v string) *InvokeAssistantRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *InvokeAssistantRequest) SetStream(v bool) *InvokeAssistantRequest {
	s.Stream = &v
	return s
}

func (s *InvokeAssistantRequest) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessages struct {
	Content *InvokeAssistantRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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

func (s InvokeAssistantRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessages) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessages) GetContent() *InvokeAssistantRequestMessagesContent {
	return s.Content
}

func (s *InvokeAssistantRequestMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *InvokeAssistantRequestMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *InvokeAssistantRequestMessages) GetRole() *string {
	return s.Role
}

func (s *InvokeAssistantRequestMessages) SetContent(v *InvokeAssistantRequestMessagesContent) *InvokeAssistantRequestMessages {
	s.Content = v
	return s
}

func (s *InvokeAssistantRequestMessages) SetContentDesc(v string) *InvokeAssistantRequestMessages {
	s.ContentDesc = &v
	return s
}

func (s *InvokeAssistantRequestMessages) SetCreateAt(v int64) *InvokeAssistantRequestMessages {
	s.CreateAt = &v
	return s
}

func (s *InvokeAssistantRequestMessages) SetRole(v string) *InvokeAssistantRequestMessages {
	s.Role = &v
	return s
}

func (s *InvokeAssistantRequestMessages) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContent struct {
	CardCallback *InvokeAssistantRequestMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *InvokeAssistantRequestMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *InvokeAssistantRequestMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *InvokeAssistantRequestMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *InvokeAssistantRequestMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *InvokeAssistantRequestMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAssistantRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContent) GetCardCallback() *InvokeAssistantRequestMessagesContentCardCallback {
	return s.CardCallback
}

func (s *InvokeAssistantRequestMessagesContent) GetDingCard() *InvokeAssistantRequestMessagesContentDingCard {
	return s.DingCard
}

func (s *InvokeAssistantRequestMessagesContent) GetDingNormalCard() *InvokeAssistantRequestMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *InvokeAssistantRequestMessagesContent) GetMarkdown() *InvokeAssistantRequestMessagesContentMarkdown {
	return s.Markdown
}

func (s *InvokeAssistantRequestMessagesContent) GetStructView() *InvokeAssistantRequestMessagesContentStructView {
	return s.StructView
}

func (s *InvokeAssistantRequestMessagesContent) GetText() *InvokeAssistantRequestMessagesContentText {
	return s.Text
}

func (s *InvokeAssistantRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *InvokeAssistantRequestMessagesContent) SetCardCallback(v *InvokeAssistantRequestMessagesContentCardCallback) *InvokeAssistantRequestMessagesContent {
	s.CardCallback = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetDingCard(v *InvokeAssistantRequestMessagesContentDingCard) *InvokeAssistantRequestMessagesContent {
	s.DingCard = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetDingNormalCard(v *InvokeAssistantRequestMessagesContentDingNormalCard) *InvokeAssistantRequestMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetMarkdown(v *InvokeAssistantRequestMessagesContentMarkdown) *InvokeAssistantRequestMessagesContent {
	s.Markdown = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetStructView(v *InvokeAssistantRequestMessagesContentStructView) *InvokeAssistantRequestMessagesContent {
	s.StructView = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetText(v *InvokeAssistantRequestMessagesContentText) *InvokeAssistantRequestMessagesContent {
	s.Text = v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) SetType(v string) *InvokeAssistantRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContent) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentCardCallback struct {
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

func (s InvokeAssistantRequestMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *InvokeAssistantRequestMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *InvokeAssistantRequestMessagesContentCardCallback) SetContent(v string) *InvokeAssistantRequestMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentCardCallback) SetRelatedMessageId(v string) *InvokeAssistantRequestMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingCard struct {
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

func (s InvokeAssistantRequestMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *InvokeAssistantRequestMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *InvokeAssistantRequestMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *InvokeAssistantRequestMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *InvokeAssistantRequestMessagesContentDingCard) SetContent(v string) *InvokeAssistantRequestMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingCard) SetContentType(v string) *InvokeAssistantRequestMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingCard) SetFinished(v bool) *InvokeAssistantRequestMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingCard) SetTemplateId(v string) *InvokeAssistantRequestMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *InvokeAssistantRequestMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) GetCardData() *InvokeAssistantRequestMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) GetCardUpdateOptions() *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) SetCardData(v *InvokeAssistantRequestMessagesContentDingNormalCardCardData) *InvokeAssistantRequestMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) SetCardTemplateId(v string) *InvokeAssistantRequestMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) SetCardUpdateOptions(v *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) *InvokeAssistantRequestMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) *InvokeAssistantRequestMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *InvokeAssistantRequestMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCard) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingNormalCardCardData struct {
	// example:
	//
	// {}
	CardParamMap map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardData) GetCardParamMap() map[string]interface{} {
	return s.CardParamMap
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardData) SetCardParamMap(v map[string]interface{}) *InvokeAssistantRequestMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *InvokeAssistantRequestMessagesContentMarkdown) SetValue(v string) *InvokeAssistantRequestMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructView struct {
	Parts []*InvokeAssistantRequestMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s InvokeAssistantRequestMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructView) GetParts() []*InvokeAssistantRequestMessagesContentStructViewParts {
	return s.Parts
}

func (s *InvokeAssistantRequestMessagesContentStructView) SetParts(v []*InvokeAssistantRequestMessagesContentStructViewParts) *InvokeAssistantRequestMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructView) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *InvokeAssistantRequestMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                         `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *InvokeAssistantRequestMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetDataPart() *InvokeAssistantRequestMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetReasonPart() *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetRecommendPart() *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetReferencePart() *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetTextPart() *InvokeAssistantRequestMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetAppend(v bool) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetDataPart(v *InvokeAssistantRequestMessagesContentStructViewPartsDataPart) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetFinish(v bool) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetPartDesc(v string) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetPartId(v string) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetReasonPart(v *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetRecommendPart(v *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetReferencePart(v *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetTextPart(v *InvokeAssistantRequestMessagesContentStructViewPartsTextPart) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) SetType(v string) *InvokeAssistantRequestMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewParts) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsDataPart) SetData(v interface{}) *InvokeAssistantRequestMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) SetReason(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) GetRecommends() []*InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsReferencePart struct {
	References []*InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) GetReferences() []*InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) SetReferences(v []*InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsTextPart) SetText(v string) *InvokeAssistantRequestMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type InvokeAssistantRequestMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAssistantRequestMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantRequestMessagesContentText) GoString() string {
	return s.String()
}

func (s *InvokeAssistantRequestMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *InvokeAssistantRequestMessagesContentText) SetValue(v string) *InvokeAssistantRequestMessagesContentText {
	s.Value = &v
	return s
}

func (s *InvokeAssistantRequestMessagesContentText) Validate() error {
	return dara.Validate(s)
}
