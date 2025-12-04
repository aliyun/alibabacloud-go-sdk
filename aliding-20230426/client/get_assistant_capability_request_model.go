// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssistantCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *GetAssistantCapabilityRequest
	GetAssistantId() *string
	SetExtLoginUser(v *GetAssistantCapabilityRequestExtLoginUser) *GetAssistantCapabilityRequest
	GetExtLoginUser() *GetAssistantCapabilityRequestExtLoginUser
	SetMessages(v []*GetAssistantCapabilityRequestMessages) *GetAssistantCapabilityRequest
	GetMessages() []*GetAssistantCapabilityRequestMessages
	SetOriginalAssistantId(v string) *GetAssistantCapabilityRequest
	GetOriginalAssistantId() *string
	SetProtocol(v string) *GetAssistantCapabilityRequest
	GetProtocol() *string
	SetSourceIdOfOriginalAssistantId(v string) *GetAssistantCapabilityRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *GetAssistantCapabilityRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetThreadId(v string) *GetAssistantCapabilityRequest
	GetThreadId() *string
	SetTimeout(v int32) *GetAssistantCapabilityRequest
	GetTimeout() *int32
}

type GetAssistantCapabilityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId  *string                                    `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	ExtLoginUser *GetAssistantCapabilityRequestExtLoginUser `json:"extLoginUser,omitempty" xml:"extLoginUser,omitempty" type:"Struct"`
	// This parameter is required.
	Messages []*GetAssistantCapabilityRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// originalAssistantId1
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// cfp
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
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
	// threadId
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// 5000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s GetAssistantCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequest) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *GetAssistantCapabilityRequest) GetExtLoginUser() *GetAssistantCapabilityRequestExtLoginUser {
	return s.ExtLoginUser
}

func (s *GetAssistantCapabilityRequest) GetMessages() []*GetAssistantCapabilityRequestMessages {
	return s.Messages
}

func (s *GetAssistantCapabilityRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *GetAssistantCapabilityRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetAssistantCapabilityRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *GetAssistantCapabilityRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *GetAssistantCapabilityRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetAssistantCapabilityRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetAssistantCapabilityRequest) SetAssistantId(v string) *GetAssistantCapabilityRequest {
	s.AssistantId = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetExtLoginUser(v *GetAssistantCapabilityRequestExtLoginUser) *GetAssistantCapabilityRequest {
	s.ExtLoginUser = v
	return s
}

func (s *GetAssistantCapabilityRequest) SetMessages(v []*GetAssistantCapabilityRequestMessages) *GetAssistantCapabilityRequest {
	s.Messages = v
	return s
}

func (s *GetAssistantCapabilityRequest) SetOriginalAssistantId(v string) *GetAssistantCapabilityRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetProtocol(v string) *GetAssistantCapabilityRequest {
	s.Protocol = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetSourceIdOfOriginalAssistantId(v string) *GetAssistantCapabilityRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetSourceTypeOfOriginalAssistantId(v string) *GetAssistantCapabilityRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetThreadId(v string) *GetAssistantCapabilityRequest {
	s.ThreadId = &v
	return s
}

func (s *GetAssistantCapabilityRequest) SetTimeout(v int32) *GetAssistantCapabilityRequest {
	s.Timeout = &v
	return s
}

func (s *GetAssistantCapabilityRequest) Validate() error {
	if s.ExtLoginUser != nil {
		if err := s.ExtLoginUser.Validate(); err != nil {
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

type GetAssistantCapabilityRequestExtLoginUser struct {
	// example:
	//
	// mozi
	ExtLoginUserDomain *string `json:"extLoginUserDomain,omitempty" xml:"extLoginUserDomain,omitempty"`
	// example:
	//
	// outeruserId123
	ExtLoginUserId *string `json:"extLoginUserId,omitempty" xml:"extLoginUserId,omitempty"`
	// example:
	//
	// 外部游客1
	ExtLoginUserName *string `json:"extLoginUserName,omitempty" xml:"extLoginUserName,omitempty"`
}

func (s GetAssistantCapabilityRequestExtLoginUser) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestExtLoginUser) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestExtLoginUser) GetExtLoginUserDomain() *string {
	return s.ExtLoginUserDomain
}

func (s *GetAssistantCapabilityRequestExtLoginUser) GetExtLoginUserId() *string {
	return s.ExtLoginUserId
}

func (s *GetAssistantCapabilityRequestExtLoginUser) GetExtLoginUserName() *string {
	return s.ExtLoginUserName
}

func (s *GetAssistantCapabilityRequestExtLoginUser) SetExtLoginUserDomain(v string) *GetAssistantCapabilityRequestExtLoginUser {
	s.ExtLoginUserDomain = &v
	return s
}

func (s *GetAssistantCapabilityRequestExtLoginUser) SetExtLoginUserId(v string) *GetAssistantCapabilityRequestExtLoginUser {
	s.ExtLoginUserId = &v
	return s
}

func (s *GetAssistantCapabilityRequestExtLoginUser) SetExtLoginUserName(v string) *GetAssistantCapabilityRequestExtLoginUser {
	s.ExtLoginUserName = &v
	return s
}

func (s *GetAssistantCapabilityRequestExtLoginUser) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessages struct {
	Content *GetAssistantCapabilityRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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

func (s GetAssistantCapabilityRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessages) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessages) GetContent() *GetAssistantCapabilityRequestMessagesContent {
	return s.Content
}

func (s *GetAssistantCapabilityRequestMessages) GetContentDesc() *string {
	return s.ContentDesc
}

func (s *GetAssistantCapabilityRequestMessages) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *GetAssistantCapabilityRequestMessages) GetRole() *string {
	return s.Role
}

func (s *GetAssistantCapabilityRequestMessages) SetContent(v *GetAssistantCapabilityRequestMessagesContent) *GetAssistantCapabilityRequestMessages {
	s.Content = v
	return s
}

func (s *GetAssistantCapabilityRequestMessages) SetContentDesc(v string) *GetAssistantCapabilityRequestMessages {
	s.ContentDesc = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessages) SetCreateAt(v int64) *GetAssistantCapabilityRequestMessages {
	s.CreateAt = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessages) SetRole(v string) *GetAssistantCapabilityRequestMessages {
	s.Role = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessages) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssistantCapabilityRequestMessagesContent struct {
	CardCallback *GetAssistantCapabilityRequestMessagesContentCardCallback `json:"cardCallback,omitempty" xml:"cardCallback,omitempty" type:"Struct"`
	DingCard     *GetAssistantCapabilityRequestMessagesContentDingCard     `json:"dingCard,omitempty" xml:"dingCard,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DingNormalCard *GetAssistantCapabilityRequestMessagesContentDingNormalCard `json:"dingNormalCard,omitempty" xml:"dingNormalCard,omitempty" type:"Struct"`
	Markdown       *GetAssistantCapabilityRequestMessagesContentMarkdown       `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	StructView     *GetAssistantCapabilityRequestMessagesContentStructView     `json:"structView,omitempty" xml:"structView,omitempty" type:"Struct"`
	Text           *GetAssistantCapabilityRequestMessagesContentText           `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 枚举字段，可为：text,markdown,cardCallback,dingCard,agentArtifact,dingNormalCard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetCardCallback() *GetAssistantCapabilityRequestMessagesContentCardCallback {
	return s.CardCallback
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetDingCard() *GetAssistantCapabilityRequestMessagesContentDingCard {
	return s.DingCard
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetDingNormalCard() *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	return s.DingNormalCard
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetMarkdown() *GetAssistantCapabilityRequestMessagesContentMarkdown {
	return s.Markdown
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetStructView() *GetAssistantCapabilityRequestMessagesContentStructView {
	return s.StructView
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetText() *GetAssistantCapabilityRequestMessagesContentText {
	return s.Text
}

func (s *GetAssistantCapabilityRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetCardCallback(v *GetAssistantCapabilityRequestMessagesContentCardCallback) *GetAssistantCapabilityRequestMessagesContent {
	s.CardCallback = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetDingCard(v *GetAssistantCapabilityRequestMessagesContentDingCard) *GetAssistantCapabilityRequestMessagesContent {
	s.DingCard = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetDingNormalCard(v *GetAssistantCapabilityRequestMessagesContentDingNormalCard) *GetAssistantCapabilityRequestMessagesContent {
	s.DingNormalCard = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetMarkdown(v *GetAssistantCapabilityRequestMessagesContentMarkdown) *GetAssistantCapabilityRequestMessagesContent {
	s.Markdown = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetStructView(v *GetAssistantCapabilityRequestMessagesContentStructView) *GetAssistantCapabilityRequestMessagesContent {
	s.StructView = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetText(v *GetAssistantCapabilityRequestMessagesContentText) *GetAssistantCapabilityRequestMessagesContent {
	s.Text = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) SetType(v string) *GetAssistantCapabilityRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContent) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentCardCallback struct {
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

func (s GetAssistantCapabilityRequestMessagesContentCardCallback) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentCardCallback) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentCardCallback) GetContent() *string {
	return s.Content
}

func (s *GetAssistantCapabilityRequestMessagesContentCardCallback) GetRelatedMessageId() *string {
	return s.RelatedMessageId
}

func (s *GetAssistantCapabilityRequestMessagesContentCardCallback) SetContent(v string) *GetAssistantCapabilityRequestMessagesContentCardCallback {
	s.Content = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentCardCallback) SetRelatedMessageId(v string) *GetAssistantCapabilityRequestMessagesContentCardCallback {
	s.RelatedMessageId = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentCardCallback) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentDingCard struct {
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

func (s GetAssistantCapabilityRequestMessagesContentDingCard) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingCard) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) GetContent() *string {
	return s.Content
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) GetContentType() *string {
	return s.ContentType
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) GetFinished() *bool {
	return s.Finished
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) SetContent(v string) *GetAssistantCapabilityRequestMessagesContentDingCard {
	s.Content = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) SetContentType(v string) *GetAssistantCapabilityRequestMessagesContentDingCard {
	s.ContentType = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) SetFinished(v bool) *GetAssistantCapabilityRequestMessagesContentDingCard {
	s.Finished = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) SetTemplateId(v string) *GetAssistantCapabilityRequestMessagesContentDingCard {
	s.TemplateId = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingCard) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentDingNormalCard struct {
	// example:
	//
	// {}
	CardData *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// example:
	//
	// templateId1
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// {}
	CardUpdateOptions *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	DynamicDataSourceConfigs []*GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	PrivateData map[string]interface{} `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCard) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCard) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) GetCardData() *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData {
	return s.CardData
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) GetCardUpdateOptions() *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions {
	return s.CardUpdateOptions
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) GetDynamicDataSourceConfigs() []*GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	return s.DynamicDataSourceConfigs
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) GetPrivateData() map[string]interface{} {
	return s.PrivateData
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) SetCardData(v *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	s.CardData = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) SetCardTemplateId(v string) *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	s.CardTemplateId = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) SetCardUpdateOptions(v *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	s.CardUpdateOptions = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) SetDynamicDataSourceConfigs(v []*GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	s.DynamicDataSourceConfigs = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) SetPrivateData(v map[string]interface{}) *GetAssistantCapabilityRequestMessagesContentDingNormalCard {
	s.PrivateData = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCard) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData struct {
	CardParamMap interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) GetCardParamMap() interface{} {
	return s.CardParamMap
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) SetCardParamMap(v interface{}) *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData {
	s.CardParamMap = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardData) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions struct {
	// example:
	//
	// {}
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// example:
	//
	// {}
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdateCardDataByKey() *bool {
	return s.UpdateCardDataByKey
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) GetUpdatePrivateDataByKey() *bool {
	return s.UpdatePrivateDataByKey
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdateCardDataByKey(v bool) *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardCardUpdateOptions) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs struct {
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
	PullConfig *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetConstParams() map[string]interface{} {
	return s.ConstParams
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetDynamicDataSourceId() *string {
	return s.DynamicDataSourceId
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) GetPullConfig() *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	return s.PullConfig
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetConstParams(v map[string]interface{}) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) SetPullConfig(v *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigs) Validate() error {
	if s.PullConfig != nil {
		if err := s.PullConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig struct {
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

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetPullStrategy() *string {
	return s.PullStrategy
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentDingNormalCardDynamicDataSourceConfigsPullConfig) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentMarkdown struct {
	// example:
	//
	// 1. markdown内容
	//
	// 2. markdown内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentMarkdown) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentMarkdown) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentMarkdown) GetValue() *string {
	return s.Value
}

func (s *GetAssistantCapabilityRequestMessagesContentMarkdown) SetValue(v string) *GetAssistantCapabilityRequestMessagesContentMarkdown {
	s.Value = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentMarkdown) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentStructView struct {
	Parts []*GetAssistantCapabilityRequestMessagesContentStructViewParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructView) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructView) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructView) GetParts() []*GetAssistantCapabilityRequestMessagesContentStructViewParts {
	return s.Parts
}

func (s *GetAssistantCapabilityRequestMessagesContentStructView) SetParts(v []*GetAssistantCapabilityRequestMessagesContentStructViewParts) *GetAssistantCapabilityRequestMessagesContentStructView {
	s.Parts = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructView) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentStructViewParts struct {
	Append *bool `json:"append,omitempty" xml:"append,omitempty"`
	// example:
	//
	// {}
	DataPart *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart `json:"dataPart,omitempty" xml:"dataPart,omitempty" type:"Struct"`
	Finish   *bool                                                                `json:"finish,omitempty" xml:"finish,omitempty"`
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
	ReasonPart *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart `json:"reasonPart,omitempty" xml:"reasonPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	RecommendPart *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart `json:"recommendPart,omitempty" xml:"recommendPart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ReferencePart *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart `json:"referencePart,omitempty" xml:"referencePart,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TextPart *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart `json:"textPart,omitempty" xml:"textPart,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// textPart
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewParts) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewParts) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetAppend() *bool {
	return s.Append
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetDataPart() *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart {
	return s.DataPart
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetFinish() *bool {
	return s.Finish
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetPartDesc() *string {
	return s.PartDesc
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetPartId() *string {
	return s.PartId
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetReasonPart() *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart {
	return s.ReasonPart
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetRecommendPart() *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart {
	return s.RecommendPart
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetReferencePart() *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart {
	return s.ReferencePart
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetTextPart() *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart {
	return s.TextPart
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) GetType() *string {
	return s.Type
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetAppend(v bool) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.Append = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetDataPart(v *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.DataPart = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetFinish(v bool) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.Finish = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetPartDesc(v string) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.PartDesc = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetPartId(v string) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.PartId = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetReasonPart(v *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.ReasonPart = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetRecommendPart(v *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.RecommendPart = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetReferencePart(v *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.ReferencePart = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetTextPart(v *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.TextPart = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) SetType(v string) *GetAssistantCapabilityRequestMessagesContentStructViewParts {
	s.Type = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewParts) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) GetData() interface{} {
	return s.Data
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) SetData(v interface{}) *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart {
	s.Data = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsDataPart) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart struct {
	// example:
	//
	// 123123
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) GetReason() *string {
	return s.Reason
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) SetReason(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart {
	s.Reason = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReasonPart) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart struct {
	Recommends []*GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) GetRecommends() []*GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends {
	return s.Recommends
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) SetRecommends(v []*GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart {
	s.Recommends = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPart) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) GetText() *string {
	return s.Text
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) GetUrl() *string {
	return s.Url
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) SetMobileUrl(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.MobileUrl = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) SetText(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Text = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) SetUrl(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends {
	s.Url = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsRecommendPartRecommends) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart struct {
	References []*GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) GetReferences() []*GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	return s.References
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) SetReferences(v []*GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart {
	s.References = v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePart) Validate() error {
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

type GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences struct {
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

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetIndex() *string {
	return s.Index
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetName() *string {
	return s.Name
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceCode() *string {
	return s.SourceCode
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetSourceIcon() *string {
	return s.SourceIcon
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetSummary() *string {
	return s.Summary
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetTitle() *string {
	return s.Title
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) GetUrl() *string {
	return s.Url
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetIndex(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Index = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetName(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Name = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceCode(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceCode = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetSourceIcon(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.SourceIcon = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetSummary(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Summary = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetTitle(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Title = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) SetUrl(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences {
	s.Url = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsReferencePartReferences) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart struct {
	// example:
	//
	// 123123
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) GetText() *string {
	return s.Text
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) SetText(v string) *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart {
	s.Text = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentStructViewPartsTextPart) Validate() error {
	return dara.Validate(s)
}

type GetAssistantCapabilityRequestMessagesContentText struct {
	// example:
	//
	// 你好！
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAssistantCapabilityRequestMessagesContentText) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityRequestMessagesContentText) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityRequestMessagesContentText) GetValue() *string {
	return s.Value
}

func (s *GetAssistantCapabilityRequestMessagesContentText) SetValue(v string) *GetAssistantCapabilityRequestMessagesContentText {
	s.Value = &v
	return s
}

func (s *GetAssistantCapabilityRequestMessagesContentText) Validate() error {
	return dara.Validate(s)
}
