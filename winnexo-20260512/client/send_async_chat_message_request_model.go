// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SendAsyncChatMessageRequest
	GetAppToken() *string
	SetAutoPreference(v string) *SendAsyncChatMessageRequest
	GetAutoPreference() *string
	SetContent(v string) *SendAsyncChatMessageRequest
	GetContent() *string
	SetContentType(v string) *SendAsyncChatMessageRequest
	GetContentType() *string
	SetContextItems(v []*SendAsyncChatMessageRequestContextItems) *SendAsyncChatMessageRequest
	GetContextItems() []*SendAsyncChatMessageRequestContextItems
	SetDigitalEmployeeName(v []*string) *SendAsyncChatMessageRequest
	GetDigitalEmployeeName() []*string
	SetDirectChat(v bool) *SendAsyncChatMessageRequest
	GetDirectChat() *bool
	SetDirectMaxHistoryRounds(v interface{}) *SendAsyncChatMessageRequest
	GetDirectMaxHistoryRounds() interface{}
	SetEnableWebSearch(v bool) *SendAsyncChatMessageRequest
	GetEnableWebSearch() *bool
	SetFiles(v []*SendAsyncChatMessageRequestFiles) *SendAsyncChatMessageRequest
	GetFiles() []*SendAsyncChatMessageRequestFiles
	SetLocale(v string) *SendAsyncChatMessageRequest
	GetLocale() *string
	SetModel(v string) *SendAsyncChatMessageRequest
	GetModel() *string
	SetOpsUserCode(v string) *SendAsyncChatMessageRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SendAsyncChatMessageRequest
	GetOpsUserId() *string
	SetReuseLastSession(v bool) *SendAsyncChatMessageRequest
	GetReuseLastSession() *bool
	SetSegments(v []*SendAsyncChatMessageRequestSegments) *SendAsyncChatMessageRequest
	GetSegments() []*SendAsyncChatMessageRequestSegments
	SetSessionId(v string) *SendAsyncChatMessageRequest
	GetSessionId() *string
	SetStream(v bool) *SendAsyncChatMessageRequest
	GetStream() *bool
	SetTaskExecution(v *SendAsyncChatMessageRequestTaskExecution) *SendAsyncChatMessageRequest
	GetTaskExecution() *SendAsyncChatMessageRequestTaskExecution
	SetTenantId(v string) *SendAsyncChatMessageRequest
	GetTenantId() *string
	SetThinkingEnabled(v bool) *SendAsyncChatMessageRequest
	GetThinkingEnabled() *bool
	SetThinkingLevel(v string) *SendAsyncChatMessageRequest
	GetThinkingLevel() *string
	SetWorkMode(v string) *SendAsyncChatMessageRequest
	GetWorkMode() *string
}

type SendAsyncChatMessageRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// AUTO 偏好：none=默认、performance=效果优先、credit=成本优先；缺省或 null 继承会话，空白重置为 none
	//
	// example:
	//
	// credit
	AutoPreference *string `json:"autoPreference,omitempty" xml:"autoPreference,omitempty"`
	// The body of the user message.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The message type. Valid values: Text and Markdown.
	//
	// example:
	//
	// Text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The references and annotations for the current turn.
	//
	// example:
	//
	// []
	ContextItems []*SendAsyncChatMessageRequestContextItems `json:"contextItems,omitempty" xml:"contextItems,omitempty" type:"Repeated"`
	// The list of digital employee names. For compatibility with the old format, a single string can also be passed.
	//
	// example:
	//
	// ["pcitc-magent"]
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// **[Deprecated]*	- Use workMode=direct instead. If set to true, winnexo_chat is retained and processed in request-level direct mode.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The maximum number of history rounds referenced in the current turn for the direct mode. Valid values: integers from 0 to 100. Default value: 0. Invalid values fall back to the original single-turn path. This parameter is ignored in non-direct modes.
	//
	// example:
	//
	// 1
	DirectMaxHistoryRounds interface{} `json:"directMaxHistoryRounds,omitempty" xml:"directMaxHistoryRounds,omitempty"`
	// Specifies whether to enable web search. Default value: false. In the task execution scenario (when taskExecution is passed), the task configuration prevails.
	//
	// example:
	//
	// false
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The list of file references. Each item is an object, and fileId is required (returned by uploadChatFile).
	Files []*SendAsyncChatMessageRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The language for the current turn. Valid values: zh-CN and en-US. This parameter is used for API error messages and passed through to CodeAgent. If omitted, the current language behavior is retained and user settings are not modified.
	//
	// example:
	//
	// en-US
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The model tier. Valid values: quick, standard, flagship, auto, or other currently available model names (case-insensitive). If not specified, the model tier of the session is inherited. For a new session, the default value is standard.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// Specifies whether to reuse the most recent session under the digital employee when sessionId is not provided (CLI scenario). Default value: false, which means a new session is created.
	//
	// example:
	//
	// false
	ReuseLastSession *bool `json:"reuseLastSession,omitempty" xml:"reuseLastSession,omitempty"`
	// The structured segments in the message, such as session or object mentions.
	//
	// example:
	//
	// []
	Segments []*SendAsyncChatMessageRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// The session ID. If not specified, a new session is created.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to use streaming generation. This API always generates background content in streaming mode and writes it to the message stream. The value of this parameter does not change the response structure.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. If passed, the request is processed through the task execution pipeline.
	TaskExecution *SendAsyncChatMessageRequestTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 本轮显式覆盖是否思考；AUTO 未传时由决策模型判断，固定档未传时沿用个人设置，最终遵守模型能力与租户开关
	//
	// example:
	//
	// true
	ThinkingEnabled *bool `json:"thinkingEnabled,omitempty" xml:"thinkingEnabled,omitempty"`
	// 本轮显式覆盖思考强度；AUTO 未传时由决策模型判断，固定档未传时沿用个人设置，最终遵守模型能力与租户开关
	//
	// example:
	//
	// high
	ThinkingLevel *string `json:"thinkingLevel,omitempty" xml:"thinkingLevel,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A), work (deep work, default), and direct (direct connection mode, which does not start the sandbox, uses independent direct history controlled by directMaxHistoryRounds, and is equivalent to directChat=true). The ask and work modes are session-level. For continued chats, the session mode is inherited by default. To switch modes, create a new session or fork. In multi-digital employee or task execution scenarios, if ask is passed, it takes effect as work. If directChat is set to true, this parameter is ignored.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
}

func (s SendAsyncChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequest) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SendAsyncChatMessageRequest) GetAutoPreference() *string {
	return s.AutoPreference
}

func (s *SendAsyncChatMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendAsyncChatMessageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendAsyncChatMessageRequest) GetContextItems() []*SendAsyncChatMessageRequestContextItems {
	return s.ContextItems
}

func (s *SendAsyncChatMessageRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *SendAsyncChatMessageRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendAsyncChatMessageRequest) GetDirectMaxHistoryRounds() interface{} {
	return s.DirectMaxHistoryRounds
}

func (s *SendAsyncChatMessageRequest) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendAsyncChatMessageRequest) GetFiles() []*SendAsyncChatMessageRequestFiles {
	return s.Files
}

func (s *SendAsyncChatMessageRequest) GetLocale() *string {
	return s.Locale
}

func (s *SendAsyncChatMessageRequest) GetModel() *string {
	return s.Model
}

func (s *SendAsyncChatMessageRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SendAsyncChatMessageRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SendAsyncChatMessageRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendAsyncChatMessageRequest) GetSegments() []*SendAsyncChatMessageRequestSegments {
	return s.Segments
}

func (s *SendAsyncChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendAsyncChatMessageRequest) GetStream() *bool {
	return s.Stream
}

func (s *SendAsyncChatMessageRequest) GetTaskExecution() *SendAsyncChatMessageRequestTaskExecution {
	return s.TaskExecution
}

func (s *SendAsyncChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendAsyncChatMessageRequest) GetThinkingEnabled() *bool {
	return s.ThinkingEnabled
}

func (s *SendAsyncChatMessageRequest) GetThinkingLevel() *string {
	return s.ThinkingLevel
}

func (s *SendAsyncChatMessageRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendAsyncChatMessageRequest) SetAppToken(v string) *SendAsyncChatMessageRequest {
	s.AppToken = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetAutoPreference(v string) *SendAsyncChatMessageRequest {
	s.AutoPreference = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetContent(v string) *SendAsyncChatMessageRequest {
	s.Content = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetContentType(v string) *SendAsyncChatMessageRequest {
	s.ContentType = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetContextItems(v []*SendAsyncChatMessageRequestContextItems) *SendAsyncChatMessageRequest {
	s.ContextItems = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetDigitalEmployeeName(v []*string) *SendAsyncChatMessageRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetDirectChat(v bool) *SendAsyncChatMessageRequest {
	s.DirectChat = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetDirectMaxHistoryRounds(v interface{}) *SendAsyncChatMessageRequest {
	s.DirectMaxHistoryRounds = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetEnableWebSearch(v bool) *SendAsyncChatMessageRequest {
	s.EnableWebSearch = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetFiles(v []*SendAsyncChatMessageRequestFiles) *SendAsyncChatMessageRequest {
	s.Files = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetLocale(v string) *SendAsyncChatMessageRequest {
	s.Locale = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetModel(v string) *SendAsyncChatMessageRequest {
	s.Model = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetOpsUserCode(v string) *SendAsyncChatMessageRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetOpsUserId(v string) *SendAsyncChatMessageRequest {
	s.OpsUserId = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetReuseLastSession(v bool) *SendAsyncChatMessageRequest {
	s.ReuseLastSession = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetSegments(v []*SendAsyncChatMessageRequestSegments) *SendAsyncChatMessageRequest {
	s.Segments = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetSessionId(v string) *SendAsyncChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetStream(v bool) *SendAsyncChatMessageRequest {
	s.Stream = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetTaskExecution(v *SendAsyncChatMessageRequestTaskExecution) *SendAsyncChatMessageRequest {
	s.TaskExecution = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetTenantId(v string) *SendAsyncChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetThinkingEnabled(v bool) *SendAsyncChatMessageRequest {
	s.ThinkingEnabled = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetThinkingLevel(v string) *SendAsyncChatMessageRequest {
	s.ThinkingLevel = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetWorkMode(v string) *SendAsyncChatMessageRequest {
	s.WorkMode = &v
	return s
}

func (s *SendAsyncChatMessageRequest) Validate() error {
	if s.ContextItems != nil {
		for _, item := range s.ContextItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Segments != nil {
		for _, item := range s.Segments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskExecution != nil {
		if err := s.TaskExecution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendAsyncChatMessageRequestContextItems struct {
	// The annotation content entered by the user.
	//
	// example:
	//
	// Please verify the data for the East China region
	Annotation *string `json:"annotation,omitempty" xml:"annotation,omitempty"`
	// The list of attachments associated with the output annotation. This parameter is only applicable when type is set to output_annotation.
	//
	// example:
	//
	// []
	AnnotationFiles []*SendAsyncChatMessageRequestContextItemsAnnotationFiles `json:"annotationFiles,omitempty" xml:"annotationFiles,omitempty" type:"Repeated"`
	// The content corresponding to the reference or annotation. For output_annotation, this parameter is optional, but annotationFiles must be provided.
	//
	// example:
	//
	// Sales revenue for this month
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The ID of the referenced or annotated output file.
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// The name of the referenced or annotated output file.
	//
	// example:
	//
	// Sales data.xlsx
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The type of the referenced or annotated output file.
	//
	// example:
	//
	// xlsx
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// The ID of the referenced or annotated message.
	//
	// example:
	//
	// exampleMessageId
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The name of the sender of the referenced or annotated message.
	//
	// example:
	//
	// Sales Assistant
	SenderName *string `json:"senderName,omitempty" xml:"senderName,omitempty"`
	// The context type. Valid values: message_reference, message_selection, message_annotation, output_selection, and output_annotation.
	//
	// example:
	//
	// message_annotation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendAsyncChatMessageRequestContextItems) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestContextItems) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestContextItems) GetAnnotation() *string {
	return s.Annotation
}

func (s *SendAsyncChatMessageRequestContextItems) GetAnnotationFiles() []*SendAsyncChatMessageRequestContextItemsAnnotationFiles {
	return s.AnnotationFiles
}

func (s *SendAsyncChatMessageRequestContextItems) GetContent() *string {
	return s.Content
}

func (s *SendAsyncChatMessageRequestContextItems) GetFileId() *string {
	return s.FileId
}

func (s *SendAsyncChatMessageRequestContextItems) GetFileName() *string {
	return s.FileName
}

func (s *SendAsyncChatMessageRequestContextItems) GetFileType() *string {
	return s.FileType
}

func (s *SendAsyncChatMessageRequestContextItems) GetMessageId() *string {
	return s.MessageId
}

func (s *SendAsyncChatMessageRequestContextItems) GetSenderName() *string {
	return s.SenderName
}

func (s *SendAsyncChatMessageRequestContextItems) GetType() *string {
	return s.Type
}

func (s *SendAsyncChatMessageRequestContextItems) SetAnnotation(v string) *SendAsyncChatMessageRequestContextItems {
	s.Annotation = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetAnnotationFiles(v []*SendAsyncChatMessageRequestContextItemsAnnotationFiles) *SendAsyncChatMessageRequestContextItems {
	s.AnnotationFiles = v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetContent(v string) *SendAsyncChatMessageRequestContextItems {
	s.Content = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetFileId(v string) *SendAsyncChatMessageRequestContextItems {
	s.FileId = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetFileName(v string) *SendAsyncChatMessageRequestContextItems {
	s.FileName = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetFileType(v string) *SendAsyncChatMessageRequestContextItems {
	s.FileType = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetMessageId(v string) *SendAsyncChatMessageRequestContextItems {
	s.MessageId = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetSenderName(v string) *SendAsyncChatMessageRequestContextItems {
	s.SenderName = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) SetType(v string) *SendAsyncChatMessageRequestContextItems {
	s.Type = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItems) Validate() error {
	if s.AnnotationFiles != nil {
		for _, item := range s.AnnotationFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendAsyncChatMessageRequestContextItemsAnnotationFiles struct {
	// The file name of the attachment.
	//
	// example:
	//
	// Supplementary notes.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The public access URL of the attachment. This parameter is optional.
	//
	// example:
	//
	// https://example.com/public/file
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID of the attachment. This parameter is required. It is used to verify the ownership of the attachment and provide it for session processing.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The internal access URL of the attachment. This parameter is optional.
	//
	// example:
	//
	// https://example.com/file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
}

func (s SendAsyncChatMessageRequestContextItemsAnnotationFiles) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestContextItemsAnnotationFiles) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) GetFileName() *string {
	return s.FileName
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) SetFileName(v string) *SendAsyncChatMessageRequestContextItemsAnnotationFiles {
	s.FileName = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) SetFilePublicUrl(v string) *SendAsyncChatMessageRequestContextItemsAnnotationFiles {
	s.FilePublicUrl = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) SetFileRecordId(v string) *SendAsyncChatMessageRequestContextItemsAnnotationFiles {
	s.FileRecordId = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) SetFileUrl(v string) *SendAsyncChatMessageRequestContextItemsAnnotationFiles {
	s.FileUrl = &v
	return s
}

func (s *SendAsyncChatMessageRequestContextItemsAnnotationFiles) Validate() error {
	return dara.Validate(s)
}

type SendAsyncChatMessageRequestFiles struct {
	// The file ID, returned by uploadChatFile.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// The file type.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendAsyncChatMessageRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestFiles) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestFiles) GetFileId() *string {
	return s.FileId
}

func (s *SendAsyncChatMessageRequestFiles) GetType() *string {
	return s.Type
}

func (s *SendAsyncChatMessageRequestFiles) SetFileId(v string) *SendAsyncChatMessageRequestFiles {
	s.FileId = &v
	return s
}

func (s *SendAsyncChatMessageRequestFiles) SetType(v string) *SendAsyncChatMessageRequestFiles {
	s.Type = &v
	return s
}

func (s *SendAsyncChatMessageRequestFiles) Validate() error {
	return dara.Validate(s)
}

type SendAsyncChatMessageRequestSegments struct {
	// The display content of the segment.
	//
	// example:
	//
	// Previous session
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The category of the mentioned object.
	//
	// example:
	//
	// conversations
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// The name of the referenced object.
	//
	// example:
	//
	// Previous session
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the referenced object.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type of the referenced object.
	//
	// example:
	//
	// conversation
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The session type.
	//
	// example:
	//
	// personal
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// The segment type.
	//
	// example:
	//
	// mention
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendAsyncChatMessageRequestSegments) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestSegments) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestSegments) GetContent() *string {
	return s.Content
}

func (s *SendAsyncChatMessageRequestSegments) GetMentionType() *string {
	return s.MentionType
}

func (s *SendAsyncChatMessageRequestSegments) GetName() *string {
	return s.Name
}

func (s *SendAsyncChatMessageRequestSegments) GetObjectId() *string {
	return s.ObjectId
}

func (s *SendAsyncChatMessageRequestSegments) GetObjectType() *string {
	return s.ObjectType
}

func (s *SendAsyncChatMessageRequestSegments) GetSessionType() *string {
	return s.SessionType
}

func (s *SendAsyncChatMessageRequestSegments) GetType() *string {
	return s.Type
}

func (s *SendAsyncChatMessageRequestSegments) SetContent(v string) *SendAsyncChatMessageRequestSegments {
	s.Content = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetMentionType(v string) *SendAsyncChatMessageRequestSegments {
	s.MentionType = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetName(v string) *SendAsyncChatMessageRequestSegments {
	s.Name = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetObjectId(v string) *SendAsyncChatMessageRequestSegments {
	s.ObjectId = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetObjectType(v string) *SendAsyncChatMessageRequestSegments {
	s.ObjectType = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetSessionType(v string) *SendAsyncChatMessageRequestSegments {
	s.SessionType = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) SetType(v string) *SendAsyncChatMessageRequestSegments {
	s.Type = &v
	return s
}

func (s *SendAsyncChatMessageRequestSegments) Validate() error {
	return dara.Validate(s)
}

type SendAsyncChatMessageRequestTaskExecution struct {
	// The billing ID.
	//
	// example:
	//
	// exampleBillingId
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// Specifies whether to enable web search.
	//
	// example:
	//
	// true
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// The name of the digital employee.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The list of associated skill codes.
	//
	// example:
	//
	// ["code"]
	SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// string_value
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The task understanding content.
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the user to which the task belongs.
	//
	// example:
	//
	// exampleUserId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendAsyncChatMessageRequestTaskExecution) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestTaskExecution) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetBillingId() *string {
	return s.BillingId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetSkillCodes() []*string {
	return s.SkillCodes
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskId() *string {
	return s.TaskId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskName() *string {
	return s.TaskName
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTenantId() *string {
	return s.TenantId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetUserId() *string {
	return s.UserId
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetBillingId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.BillingId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetEnableWebSearch(v bool) *SendAsyncChatMessageRequestTaskExecution {
	s.EnableWebSearch = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetExecutionId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.ExecutionId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetOperatingObjectName(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.OperatingObjectName = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetSkillCodes(v []*string) *SendAsyncChatMessageRequestTaskExecution {
	s.SkillCodes = v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskName(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskName = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskUnderstand(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskUnderstand = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTenantId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TenantId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetUserId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.UserId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) Validate() error {
	return dara.Validate(s)
}
