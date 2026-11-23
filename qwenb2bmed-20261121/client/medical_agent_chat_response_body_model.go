// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*MedicalAgentChatResponseBodyAudioFrames) *MedicalAgentChatResponseBody
	GetAudioFrames() []*MedicalAgentChatResponseBodyAudioFrames
	SetCommunication(v *MedicalAgentChatResponseBodyCommunication) *MedicalAgentChatResponseBody
	GetCommunication() *MedicalAgentChatResponseBodyCommunication
	SetData(v *MedicalAgentChatResponseBodyData) *MedicalAgentChatResponseBody
	GetData() *MedicalAgentChatResponseBodyData
	SetError(v *MedicalAgentChatResponseBodyError) *MedicalAgentChatResponseBody
	GetError() *MedicalAgentChatResponseBodyError
	SetErrorCode(v interface{}) *MedicalAgentChatResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *MedicalAgentChatResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *MedicalAgentChatResponseBody
	GetRequestId() *string
}

type MedicalAgentChatResponseBody struct {
	// 可选的聚合音频帧列表。若返回音频，正常帧可含 type=audio、Base64 字符串 data 和布尔值 last，正常末帧的 data 可为空字符串。失败帧可含 type=error、code、requestId、last 和可选 message。各帧字段依类型而定。
	AudioFrames []*MedicalAgentChatResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// 关联原请求及会话的标识。
	Communication *MedicalAgentChatResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// 本次应答的业务数据，正文位于 messages。
	Data *MedicalAgentChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 可选的结构化错误信息。返回时读取 code、message 和 requestId。
	Error *MedicalAgentChatResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// 业务状态码。0 表示业务成功，非 0 时结合 errorMsg 或 error 查看原因。流式还应处理 event:error。
	//
	// example:
	//
	// 0
	ErrorCode interface{} `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 业务状态的文字说明。失败时用于辅助排查，程序分支应依据状态码。
	//
	// example:
	//
	// success
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 可选的业务请求 ID。联系支持时可同时提供响应头 x-acs-request-id 中的阿里云请求 ID。
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MedicalAgentChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBody) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBody) GetAudioFrames() []*MedicalAgentChatResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *MedicalAgentChatResponseBody) GetCommunication() *MedicalAgentChatResponseBodyCommunication {
	return s.Communication
}

func (s *MedicalAgentChatResponseBody) GetData() *MedicalAgentChatResponseBodyData {
	return s.Data
}

func (s *MedicalAgentChatResponseBody) GetError() *MedicalAgentChatResponseBodyError {
	return s.Error
}

func (s *MedicalAgentChatResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *MedicalAgentChatResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MedicalAgentChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBody) SetAudioFrames(v []*MedicalAgentChatResponseBodyAudioFrames) *MedicalAgentChatResponseBody {
	s.AudioFrames = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetCommunication(v *MedicalAgentChatResponseBodyCommunication) *MedicalAgentChatResponseBody {
	s.Communication = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetData(v *MedicalAgentChatResponseBodyData) *MedicalAgentChatResponseBody {
	s.Data = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetError(v *MedicalAgentChatResponseBodyError) *MedicalAgentChatResponseBody {
	s.Error = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetErrorCode(v interface{}) *MedicalAgentChatResponseBody {
	s.ErrorCode = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetErrorMsg(v string) *MedicalAgentChatResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MedicalAgentChatResponseBody) SetRequestId(v string) *MedicalAgentChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBody) Validate() error {
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

type MedicalAgentChatResponseBodyAudioFrames struct {
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
	// AAEC
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// true 表示本路音频的最后一帧。
	Last *bool `json:"last,omitempty" xml:"last,omitempty"`
	// 音频处理失败时按实际错误返回。
	//
	// example:
	//
	// 音频生成失败
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 用于关联音频错误与本次业务请求。
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// audio 表示音频，error 表示处理失败。
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAgentChatResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetCode(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetData(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetLast(v bool) *MedicalAgentChatResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetMessage(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetRequestId(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetType(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyCommunication struct {
	// 对应请求中的 communication.requestId，便于匹配并发请求。
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 业务响应帧标识。一次完整 JSON 聚合时可为 0。SSE 按事件到达顺序处理，responseId 无需连续。
	//
	// example:
	//
	// 1
	ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
	// 对应请求中的 communication.sessionId。
	//
	// example:
	//
	// med-session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s MedicalAgentChatResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *MedicalAgentChatResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalAgentChatResponseBodyCommunication) SetRequestId(v string) *MedicalAgentChatResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) SetResponseId(v int64) *MedicalAgentChatResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) SetSessionId(v string) *MedicalAgentChatResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyData struct {
	// 可选业务调试信息，可为任意 JSON 值（含 null）。是否返回、返回时机和内容按实际执行结果，不保证每轮或结束事件必有；不用于判断业务完成。
	//
	// example:
	//
	// {}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// 响应的可选补充信息。使用图片等资源时可读取 inputResources 中返回的对应资源信息。
	ExtraInfo *MedicalAgentChatResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// 当前响应中的消息列表。读取消息的 mimeType、status 和 content。完整正文可通过 mimeType=multi_load/iframe 且 status=complete 的消息取得。
	Messages []*MedicalAgentChatResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalAgentChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *MedicalAgentChatResponseBodyData) GetExtraInfo() *MedicalAgentChatResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *MedicalAgentChatResponseBodyData) GetMessages() []*MedicalAgentChatResponseBodyDataMessages {
	return s.Messages
}

func (s *MedicalAgentChatResponseBodyData) SetDebug(v interface{}) *MedicalAgentChatResponseBodyData {
	s.Debug = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) SetExtraInfo(v *MedicalAgentChatResponseBodyDataExtraInfo) *MedicalAgentChatResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) SetMessages(v []*MedicalAgentChatResponseBodyDataMessages) *MedicalAgentChatResponseBodyData {
	s.Messages = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) Validate() error {
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

type MedicalAgentChatResponseBodyDataExtraInfo struct {
	// 可选的资源信息数组，通过 contentIndex 与请求中的内容块对应。未返回资源信息时可省略。
	InputResources []*MedicalAgentChatResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s MedicalAgentChatResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) GetInputResources() []*MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) SetInputResources(v []*MedicalAgentChatResponseBodyDataExtraInfoInputResources) *MedicalAgentChatResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) Validate() error {
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

type MedicalAgentChatResponseBodyDataExtraInfoInputResources struct {
	// 原请求本轮 content 数组的索引，从 0 开始。文本块也占用索引。
	//
	// example:
	//
	// 1
	ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
	// 上游提供时返回资源文件名。
	//
	// example:
	//
	// report.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 返回的输入资源标识，存在时可用于在业务记录中引用该资源。
	//
	// example:
	//
	// ws_123
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 资源类型，例如图片引用。
	//
	// example:
	//
	// image
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 返回的资源访问地址。客户端按实际返回值使用。
	//
	// example:
	//
	// https://example.com/report.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetFileName(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetType(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetUrl(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyDataMessages struct {
	// 消息正文，具体结构随 mimeType 变化。流式快照可能包含到当前为止的完整正文，收到新快照时更新原内容，避免重复拼接。
	//
	// example:
	//
	// 血红蛋白用于携带氧气。解读报告时可结合报告上的参考范围。
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 可选的医疗进度或正文卡片信息，按所在消息的 mimeType 与实际返回类型解析。只需要正文时可忽略展示字段。
	MetaData *MedicalAgentChatResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// 该消息的内容类型，用于选择解析和展示方式。multi_load/iframe 表示正文；遇到应用不支持的类型，可跳过该消息。
	//
	// example:
	//
	// multi_load/iframe
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// 当前消息的处理状态。complete 表示这条消息已完成；流式请求是否结束以 SSE 结束事件为准。
	//
	// example:
	//
	// complete
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetMetaData() *MedicalAgentChatResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetMetaData(v *MedicalAgentChatResponseBodyDataMessagesMetaData) *MedicalAgentChatResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetMimeType(v string) *MedicalAgentChatResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetStatus(v string) *MedicalAgentChatResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyDataMessagesMetaData struct {
	// 状态栏进度消息可返回的标题。
	//
	// example:
	//
	// 问题分析中
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// 结构由元数据类型决定。bar_thinking 可含 body、status；bar_ref_source_inline 可含 query_list、docs、status。检索词与资料数组可能为空，资料对象字段随来源变化。
	//
	// example:
	//
	// {"body":"正在分析用户的问题。","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 正文卡片列表，可为空或省略。正文中的 [(sourceSeq)] 与相同 sourceSeq 的卡片关联，卡片 type 和 content 按实际返回解析，不把普通资料列表自行当作卡片。
	MultiLoad []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// 用于关联同一进度模块或正文卡片的可选标识，按实际值匹配状态更新。
	//
	// example:
	//
	// bar_thinking_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// 用于区分进度模块，例如实际返回的 bar_thinking 或 bar_ref_source_inline；不保证每轮都有这些模块。
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetMultiLoad() []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetBarTitle(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetMultiLoad(v []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetType(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) Validate() error {
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

type MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad struct {
	// 可选，结构随卡片类型变化。图文卡片可含资源列表。
	//
	// example:
	//
	// {"list":[{"image_url":"https://example.com/image.jpg","image_width":"800","image_height":"600"}]}
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

func (s MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyError struct {
	// 业务错误码，客户端可根据此字段判断错误类型。
	//
	// example:
	//
	// INVALID_REQUEST
	Code interface{} `json:"code,omitempty" xml:"code,omitempty"`
	// 错误原因的文字说明，便于排查后调整请求。
	//
	// example:
	//
	// 请求参数无效
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 失败请求的业务请求 ID，联系支持时可一并提供。
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MedicalAgentChatResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyError) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *MedicalAgentChatResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *MedicalAgentChatResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyError) SetCode(v interface{}) *MedicalAgentChatResponseBodyError {
	s.Code = v
	return s
}

func (s *MedicalAgentChatResponseBodyError) SetMessage(v string) *MedicalAgentChatResponseBodyError {
	s.Message = &v
	return s
}

func (s *MedicalAgentChatResponseBodyError) SetRequestId(v string) *MedicalAgentChatResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyError) Validate() error {
	return dara.Validate(s)
}
