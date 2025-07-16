// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*ListMessageResponseBodyMessages) *ListMessageResponseBody
	GetMessages() []*ListMessageResponseBodyMessages
	SetRequestId(v string) *ListMessageResponseBody
	GetRequestId() *string
}

type ListMessageResponseBody struct {
	Messages []*ListMessageResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBody) GetMessages() []*ListMessageResponseBodyMessages {
	return s.Messages
}

func (s *ListMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageResponseBody) SetMessages(v []*ListMessageResponseBodyMessages) *ListMessageResponseBody {
	s.Messages = v
	return s
}

func (s *ListMessageResponseBody) SetRequestId(v string) *ListMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessages struct {
	Content *ListMessageResponseBodyMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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

func (s ListMessageResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessages) GetContent() *ListMessageResponseBodyMessagesContent {
	return s.Content
}

func (s *ListMessageResponseBodyMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *ListMessageResponseBodyMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *ListMessageResponseBodyMessages) GetId() *string {
	return s.Id
}

func (s *ListMessageResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *ListMessageResponseBodyMessages) GetRunId() *string {
	return s.RunId
}

func (s *ListMessageResponseBodyMessages) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListMessageResponseBodyMessages) SetContent(v *ListMessageResponseBodyMessagesContent) *ListMessageResponseBodyMessages {
	s.Content = v
	return s
}

func (s *ListMessageResponseBodyMessages) SetContentDesc(v string) *ListMessageResponseBodyMessages {
	s.ContentDesc = &v
	return s
}

func (s *ListMessageResponseBodyMessages) SetCreateAt(v int64) *ListMessageResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *ListMessageResponseBodyMessages) SetId(v string) *ListMessageResponseBodyMessages {
	s.Id = &v
	return s
}

func (s *ListMessageResponseBodyMessages) SetRole(v string) *ListMessageResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *ListMessageResponseBodyMessages) SetRunId(v string) *ListMessageResponseBodyMessages {
	s.RunId = &v
	return s
}

func (s *ListMessageResponseBodyMessages) SetThreadId(v string) *ListMessageResponseBodyMessages {
	s.ThreadId = &v
	return s
}

func (s *ListMessageResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContent struct {
	CardCallback *ListMessageResponseBodyMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *ListMessageResponseBodyMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *ListMessageResponseBodyMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *ListMessageResponseBodyMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *ListMessageResponseBodyMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *ListMessageResponseBodyMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMessageResponseBodyMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContent) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContent) GetCardCallback() *ListMessageResponseBodyMessagesContentCardCallback {
	return s.CardCallback
}

func (s *ListMessageResponseBodyMessagesContent) GetDingCard() *ListMessageResponseBodyMessagesContentDingCard {
	return s.DingCard
}

func (s *ListMessageResponseBodyMessagesContent) GetDingNormalCard() *ListMessageResponseBodyMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *ListMessageResponseBodyMessagesContent) GetMarkdown() *ListMessageResponseBodyMessagesContentMarkdown {
	return s.Markdown
}

func (s *ListMessageResponseBodyMessagesContent) GetStructView() *ListMessageResponseBodyMessagesContentStructView {
	return s.StructView
}

func (s *ListMessageResponseBodyMessagesContent) GetText() *ListMessageResponseBodyMessagesContentText {
	return s.Text
}

func (s *ListMessageResponseBodyMessagesContent) GetType() *string {
	return s.Type
}

func (s *ListMessageResponseBodyMessagesContent) SetCardCallback(v *ListMessageResponseBodyMessagesContentCardCallback) *ListMessageResponseBodyMessagesContent {
	s.CardCallback = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetDingCard(v *ListMessageResponseBodyMessagesContentDingCard) *ListMessageResponseBodyMessagesContent {
	s.DingCard = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetDingNormalCard(v *ListMessageResponseBodyMessagesContentDingNormalCard) *ListMessageResponseBodyMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetMarkdown(v *ListMessageResponseBodyMessagesContentMarkdown) *ListMessageResponseBodyMessagesContent {
	s.Markdown = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetStructView(v *ListMessageResponseBodyMessagesContentStructView) *ListMessageResponseBodyMessagesContent {
	s.StructView = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetText(v *ListMessageResponseBodyMessagesContentText) *ListMessageResponseBodyMessagesContent {
	s.Text = v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) SetType(v string) *ListMessageResponseBodyMessagesContent {
	s.Type = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContent) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentCardCallback struct {
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// aliding_messageId123
	RelatedMessageId *string `json:"relatedMessageId,omitempty" xml:"relatedMessageId,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *ListMessageResponseBodyMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *ListMessageResponseBodyMessagesContentCardCallback) SetContent(v string) *ListMessageResponseBodyMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentCardCallback) SetRelatedMessageId(v string) *ListMessageResponseBodyMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingCard struct {
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

func (s ListMessageResponseBodyMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *ListMessageResponseBodyMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *ListMessageResponseBodyMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *ListMessageResponseBodyMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMessageResponseBodyMessagesContentDingCard) SetContent(v string) *ListMessageResponseBodyMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingCard) SetContentType(v string) *ListMessageResponseBodyMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingCard) SetFinished(v bool) *ListMessageResponseBodyMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingCard) SetTemplateId(v string) *ListMessageResponseBodyMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *ListMessageResponseBodyMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) GetCardData() *ListMessageResponseBodyMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) GetCardUpdateOptions() *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) GetPrivateData() map[string]map[string]interface{} {
	return s.PrivateData
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) SetCardData(v *ListMessageResponseBodyMessagesContentDingNormalCardCardData) *ListMessageResponseBodyMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) SetCardTemplateId(v string) *ListMessageResponseBodyMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) SetCardUpdateOptions(v *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) *ListMessageResponseBodyMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) *ListMessageResponseBodyMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) SetPrivateData(v map[string]map[string]interface{}) *ListMessageResponseBodyMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCard) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingNormalCardCardData struct {
	// example:
	//
	// {}
	CardParamMap map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardData) GetCardParamMap() map[string]interface{} {
	return s.CardParamMap
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardData) SetCardParamMap(v map[string]interface{}) *ListMessageResponseBodyMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *ListMessageResponseBodyMessagesContentMarkdown) SetValue(v string) *ListMessageResponseBodyMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructView struct {
	Parts []*ListMessageResponseBodyMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s ListMessageResponseBodyMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructView) GetParts() []*ListMessageResponseBodyMessagesContentStructViewParts {
	return s.Parts
}

func (s *ListMessageResponseBodyMessagesContentStructView) SetParts(v []*ListMessageResponseBodyMessagesContentStructViewParts) *ListMessageResponseBodyMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructView) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *ListMessageResponseBodyMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                          `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *ListMessageResponseBodyMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetDataPart() *ListMessageResponseBodyMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetReasonPart() *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetRecommendPart() *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetReferencePart() *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetTextPart() *ListMessageResponseBodyMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetAppend(v bool) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetDataPart(v *ListMessageResponseBodyMessagesContentStructViewPartsDataPart) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetFinish(v bool) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetPartDesc(v string) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetPartId(v string) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetReasonPart(v *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetRecommendPart(v *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetReferencePart(v *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetTextPart(v *ListMessageResponseBodyMessagesContentStructViewPartsTextPart) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) SetType(v string) *ListMessageResponseBodyMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewParts) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsDataPart) SetData(v interface{}) *ListMessageResponseBodyMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) SetReason(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) GetRecommends() []*ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPart) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsReferencePart struct {
	References []*ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) GetReferences() []*ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) SetReferences(v []*ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePart) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsTextPart) SetText(v string) *ListMessageResponseBodyMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type ListMessageResponseBodyMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMessageResponseBodyMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyMessagesContentText) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *ListMessageResponseBodyMessagesContentText) SetValue(v string) *ListMessageResponseBodyMessagesContentText {
	s.Value = &v
	return s
}

func (s *ListMessageResponseBodyMessagesContentText) Validate() error {
	return dara.Validate(s)
}
