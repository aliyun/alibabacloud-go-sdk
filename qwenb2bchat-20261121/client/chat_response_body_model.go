// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*ChatResponseBodyAudioFrames) *ChatResponseBody
	GetAudioFrames() []*ChatResponseBodyAudioFrames
	SetCommunication(v *ChatResponseBodyCommunication) *ChatResponseBody
	GetCommunication() *ChatResponseBodyCommunication
	SetData(v *ChatResponseBodyData) *ChatResponseBody
	GetData() *ChatResponseBodyData
	SetError(v *ChatResponseBodyError) *ChatResponseBody
	GetError() *ChatResponseBodyError
	SetErrorCode(v interface{}) *ChatResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *ChatResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
}

type ChatResponseBody struct {
	// 开通 HTTP 音频输出后，完整 JSON 可在 audioFrames 中返回音频帧。type 为 audio 时，data 为 Base64 音频数据，last 为布尔末帧标识。type 为 error 时可含 code、requestId 和 message。无音频时可为空或省略。
	AudioFrames []*ChatResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// 用于关联响应帧与原业务请求和会话的信息。
	Communication *ChatResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// 可选业务对象，字段按实际响应出现。单帧 HTTP 音频响应的 data 也可能直接为 Base64 字符串；多态音频完整形状见接口说明。客户端不得假定固定字段全部存在。
	Data *ChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 可选错误对象，出现时按失败处理。准备阶段可用错误HTTP状态+JSON返回；已识别业务错误可保留HTTP200，SSE首帧提交后用错误终帧表达。优先结合HTTP X-Request-Id关联。
	Error *ChatResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// 可选字符串或数值，成功业务响应可为数字0；非零或失败码按错误处理。未返回本字段不自动判失败，须综合 error、HTTP 和流终帧。
	//
	// example:
	//
	// 0
	ErrorCode interface{} `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 与响应顶层业务状态码对应的说明。
	//
	// example:
	//
	// success
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 响应顶层可选的业务请求关联号。POP 请求号由响应头 x-acs-request-id 提供。
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetAudioFrames() []*ChatResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *ChatResponseBody) GetCommunication() *ChatResponseBodyCommunication {
	return s.Communication
}

func (s *ChatResponseBody) GetData() *ChatResponseBodyData {
	return s.Data
}

func (s *ChatResponseBody) GetError() *ChatResponseBodyError {
	return s.Error
}

func (s *ChatResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *ChatResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) SetAudioFrames(v []*ChatResponseBodyAudioFrames) *ChatResponseBody {
	s.AudioFrames = v
	return s
}

func (s *ChatResponseBody) SetCommunication(v *ChatResponseBodyCommunication) *ChatResponseBody {
	s.Communication = v
	return s
}

func (s *ChatResponseBody) SetData(v *ChatResponseBodyData) *ChatResponseBody {
	s.Data = v
	return s
}

func (s *ChatResponseBody) SetError(v *ChatResponseBodyError) *ChatResponseBody {
	s.Error = v
	return s
}

func (s *ChatResponseBody) SetErrorCode(v interface{}) *ChatResponseBody {
	s.ErrorCode = v
	return s
}

func (s *ChatResponseBody) SetErrorMsg(v string) *ChatResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
	if s.AudioFrames != nil {
		for _, item := range s.AudioFrames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Communication != nil {
		if err := s.Communication.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatResponseBodyAudioFrames struct {
	// 音频处理失败时可能提供。
	//
	// example:
	//
	// DOWNSTREAM_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 音频帧的 Base64 数据。末帧的值可以为空字符串。
	//
	// example:
	//
	// SUQz
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// true 表示本路音频的最后一帧。
	Last *bool `json:"last,omitempty" xml:"last,omitempty"`
	// 音频处理失败时按实际错误返回。
	//
	// example:
	//
	// 语音合成暂不可用
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 用于关联音频错误与本次业务请求。
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// audio 表示音频，error 表示处理失败。
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChatResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *ChatResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *ChatResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *ChatResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyAudioFrames) SetCode(v string) *ChatResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetData(v string) *ChatResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetLast(v bool) *ChatResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetMessage(v string) *ChatResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetRequestId(v string) *ChatResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetType(v string) *ChatResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyCommunication struct {
	// 当前响应所属的业务请求 ID，返回时与请求中的 communication.requestId 一致。
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 可选整数，业务响应帧标识；网关多帧聚合可为0。不代表可对所有网络帧排序的序列号，也不是会话或请求ID。
	//
	// example:
	//
	// 1
	ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
	// 当前响应所属的客户会话 ID。
	//
	// example:
	//
	// session-example-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *ChatResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatResponseBodyCommunication) SetRequestId(v string) *ChatResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyCommunication) SetResponseId(v int64) *ChatResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *ChatResponseBodyCommunication) SetSessionId(v string) *ChatResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyData struct {
	// 可选业务调试信息，可省略或为 null。在支持的联调场景中可能包含 intentLabel；具体结构和返回时机按实际响应读取。
	//
	// example:
	//
	// {"intentLabel":"image_search"}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// 业务响应附带的可选扩展信息。inputResources 是本轮输入资源的处理结果数组，元素可含 contentIndex、type、resourceId、url 和 fileName。contentIndex 是原始消息 content 数组的零起始下标，文本块也计数。
	ExtraInfo *ChatResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// 可选数组，可能含多条正文、过程或卡片消息。multi_load/iframe 文本为累计快照，不将每帧全文重复拼接；不要固定只取 messages[0]。
	Messages []*ChatResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s ChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *ChatResponseBodyData) GetExtraInfo() *ChatResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *ChatResponseBodyData) GetMessages() []*ChatResponseBodyDataMessages {
	return s.Messages
}

func (s *ChatResponseBodyData) SetDebug(v interface{}) *ChatResponseBodyData {
	s.Debug = v
	return s
}

func (s *ChatResponseBodyData) SetExtraInfo(v *ChatResponseBodyDataExtraInfo) *ChatResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *ChatResponseBodyData) SetMessages(v []*ChatResponseBodyDataMessages) *ChatResponseBodyData {
	s.Messages = v
	return s
}

func (s *ChatResponseBodyData) Validate() error {
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
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

type ChatResponseBodyDataExtraInfo struct {
	// 仅在响应提供本轮输入资源处理结果时返回。
	InputResources []*ChatResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s ChatResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataExtraInfo) GetInputResources() []*ChatResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *ChatResponseBodyDataExtraInfo) SetInputResources(v []*ChatResponseBodyDataExtraInfoInputResources) *ChatResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *ChatResponseBodyDataExtraInfo) Validate() error {
	if s.InputResources != nil {
		for _, item := range s.InputResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyDataExtraInfoInputResources struct {
	// 从零开始，按原始 content 数组计数，包含文本块。
	//
	// example:
	//
	// 1
	ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
	// 有文件名时返回。
	//
	// example:
	//
	// cat.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 本轮输入资源的标识，用于关联对应资源。
	//
	// example:
	//
	// ws_example_001
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 原始内容块对应的资源类型。
	//
	// example:
	//
	// image
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 本轮输入资源的可访问地址。
	//
	// example:
	//
	// https://cdn.example.com/uploads/cat.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ChatResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *ChatResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetFileName(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetType(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetUrl(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyDataMessages struct {
	// 可选多态内容，按 mimeType 解析。multi_load/iframe 常为累计文本快照；后帧更新当前内容。过程、卡片按各自类型处理，不把任意 content 强转字符串。
	//
	// example:
	//
	// 图片如下[(image_waterfall_1)]
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 可选、可扩展的消息元数据。正文消息可通过 multiLoad 提供卡片；过程消息可按类型包含 sourceSeq、type、barTitle、content，用于关联和展示相应内容。
	MetaData *ChatResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// 消息类型，按此值解析正文、进度及可选信号和追问。常见类型见接入说明。
	//
	// example:
	//
	// multi_load/iframe
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// 可选消息状态，如 processing、complete；complete 只结束该消息，不替代 SSE event:complete。HTTP200 或单条消息complete都不足以判定整次流成功。
	//
	// example:
	//
	// complete
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ChatResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessages) GetMetaData() *ChatResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *ChatResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *ChatResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *ChatResponseBodyDataMessages) SetContent(v interface{}) *ChatResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessages) SetMetaData(v *ChatResponseBodyDataMessagesMetaData) *ChatResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *ChatResponseBodyDataMessages) SetMimeType(v string) *ChatResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *ChatResponseBodyDataMessages) SetStatus(v string) *ChatResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *ChatResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatResponseBodyDataMessagesMetaData struct {
	// 用于展示当前过程的上下文标题。处理状态见对应消息的 status。
	//
	// example:
	//
	// 问题分析
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// 数据结构随模块类型变化。bar/progress 中，思考过程使用 body，工具和结束提示使用 title，过程中检索资料或图片分别使用 docs、images。完整键和示例见接入说明。
	//
	// example:
	//
	// {"body":"正在分析问题","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 与正文关联的结构化内容列表，例如图文卡片，按消息类型返回。
	MultiLoad []*ChatResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// 按本次请求和 sourceSeq 关联同一模块的更新。新的完整快照替换旧快照。
	//
	// example:
	//
	// thinking-1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// 例如 bar_thinking、bar_tool、bar_think_finish、bar_ref_source_inline 或 bar_image_inline。按实际类型选择展示组件。
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChatResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *ChatResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessagesMetaData) GetMultiLoad() []*ChatResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *ChatResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *ChatResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataMessagesMetaData) SetBarTitle(v string) *ChatResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetContent(v interface{}) *ChatResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetMultiLoad(v []*ChatResponseBodyDataMessagesMetaDataMultiLoad) *ChatResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *ChatResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetType(v string) *ChatResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) Validate() error {
	if s.MultiLoad != nil {
		for _, item := range s.MultiLoad {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyDataMessagesMetaDataMultiLoad struct {
	// 可选，结构随卡片类型变化。图文卡片可含资源列表。
	//
	// example:
	//
	// {"list":[{"imageUrl":"https://images.example.com/cat.png"}]}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 与正文中的卡片引用关联。相同标识代表同一卡片。
	//
	// example:
	//
	// image_waterfall_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// 按类型选择卡片展示组件。
	//
	// example:
	//
	// image_waterfall
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChatResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyError struct {
	// 可选的错误码，可为字符串或数值，表示服务或 Agent 的具体错误类型。程序应按 code 判断，不依赖 message 全文；数值业务码 500 不代表 HTTP 状态一定为 500。
	//
	// example:
	//
	// INVALID_REQUEST
	Code interface{} `json:"code,omitempty" xml:"code,omitempty"`
	// 业务错误的说明文本。
	//
	// example:
	//
	// 请求参数无效
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 失败调用的业务请求关联号，用于定位问题。
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyError) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *ChatResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyError) SetCode(v interface{}) *ChatResponseBodyError {
	s.Code = v
	return s
}

func (s *ChatResponseBodyError) SetMessage(v string) *ChatResponseBodyError {
	s.Message = &v
	return s
}

func (s *ChatResponseBodyError) SetRequestId(v string) *ChatResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyError) Validate() error {
	return dara.Validate(s)
}
