// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*MedicalModelInferenceResponseBodyAudioFrames) *MedicalModelInferenceResponseBody
	GetAudioFrames() []*MedicalModelInferenceResponseBodyAudioFrames
	SetCommunication(v *MedicalModelInferenceResponseBodyCommunication) *MedicalModelInferenceResponseBody
	GetCommunication() *MedicalModelInferenceResponseBodyCommunication
	SetData(v *MedicalModelInferenceResponseBodyData) *MedicalModelInferenceResponseBody
	GetData() *MedicalModelInferenceResponseBodyData
	SetError(v *MedicalModelInferenceResponseBodyError) *MedicalModelInferenceResponseBody
	GetError() *MedicalModelInferenceResponseBodyError
	SetErrorCode(v interface{}) *MedicalModelInferenceResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *MedicalModelInferenceResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *MedicalModelInferenceResponseBody
	GetRequestId() *string
}

type MedicalModelInferenceResponseBody struct {
	// 可选的聚合音频帧列表。若返回音频，正常帧可含 type=audio、Base64 字符串 data 和布尔值 last，正常末帧的 data 可为空字符串。失败帧可含 type=error、code、requestId、last 和可选 message。各帧字段依类型而定。
	AudioFrames []*MedicalModelInferenceResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// 关联原请求及会话的标识。
	Communication *MedicalModelInferenceResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// 本次应答的业务数据，正文位于 messages。
	Data *MedicalModelInferenceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 可选的结构化错误信息。返回时读取 code、message 和 requestId。
	Error *MedicalModelInferenceResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
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

func (s MedicalModelInferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBody) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBody) GetAudioFrames() []*MedicalModelInferenceResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *MedicalModelInferenceResponseBody) GetCommunication() *MedicalModelInferenceResponseBodyCommunication {
	return s.Communication
}

func (s *MedicalModelInferenceResponseBody) GetData() *MedicalModelInferenceResponseBodyData {
	return s.Data
}

func (s *MedicalModelInferenceResponseBody) GetError() *MedicalModelInferenceResponseBodyError {
	return s.Error
}

func (s *MedicalModelInferenceResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *MedicalModelInferenceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MedicalModelInferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBody) SetAudioFrames(v []*MedicalModelInferenceResponseBodyAudioFrames) *MedicalModelInferenceResponseBody {
	s.AudioFrames = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetCommunication(v *MedicalModelInferenceResponseBodyCommunication) *MedicalModelInferenceResponseBody {
	s.Communication = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetData(v *MedicalModelInferenceResponseBodyData) *MedicalModelInferenceResponseBody {
	s.Data = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetError(v *MedicalModelInferenceResponseBodyError) *MedicalModelInferenceResponseBody {
	s.Error = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetErrorCode(v interface{}) *MedicalModelInferenceResponseBody {
	s.ErrorCode = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetErrorMsg(v string) *MedicalModelInferenceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetRequestId(v string) *MedicalModelInferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBody) Validate() error {
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

type MedicalModelInferenceResponseBodyAudioFrames struct {
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

func (s MedicalModelInferenceResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetCode(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetData(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetLast(v bool) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetMessage(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetRequestId(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetType(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyCommunication struct {
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

func (s MedicalModelInferenceResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetRequestId(v string) *MedicalModelInferenceResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetResponseId(v int64) *MedicalModelInferenceResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetSessionId(v string) *MedicalModelInferenceResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyData struct {
	// 可选业务调试信息，可为任意 JSON 值（含 null）。是否返回、返回时机和内容按实际执行结果，不保证每轮或结束事件必有；不用于判断业务完成。
	//
	// example:
	//
	// {}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// 响应的可选补充信息。使用图片等资源时可读取 inputResources 中返回的对应资源信息。
	ExtraInfo *MedicalModelInferenceResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// 当前响应中的消息列表。读取消息的 mimeType、status 和 content。完整正文可通过 mimeType=multi_load/iframe 且 status=complete 的消息取得。
	Messages []*MedicalModelInferenceResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalModelInferenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *MedicalModelInferenceResponseBodyData) GetExtraInfo() *MedicalModelInferenceResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *MedicalModelInferenceResponseBodyData) GetMessages() []*MedicalModelInferenceResponseBodyDataMessages {
	return s.Messages
}

func (s *MedicalModelInferenceResponseBodyData) SetDebug(v interface{}) *MedicalModelInferenceResponseBodyData {
	s.Debug = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) SetExtraInfo(v *MedicalModelInferenceResponseBodyDataExtraInfo) *MedicalModelInferenceResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) SetMessages(v []*MedicalModelInferenceResponseBodyDataMessages) *MedicalModelInferenceResponseBodyData {
	s.Messages = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) Validate() error {
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

type MedicalModelInferenceResponseBodyDataExtraInfo struct {
	// 可选的资源信息数组，通过 contentIndex 与请求中的内容块对应。未返回资源信息时可省略。
	InputResources []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s MedicalModelInferenceResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) GetInputResources() []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) SetInputResources(v []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources) *MedicalModelInferenceResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) Validate() error {
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

type MedicalModelInferenceResponseBodyDataExtraInfoInputResources struct {
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

func (s MedicalModelInferenceResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetFileName(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetType(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetUrl(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyDataMessages struct {
	// 消息正文，具体结构随 mimeType 变化。流式快照可能包含到当前为止的完整正文，收到新快照时更新原内容，避免重复拼接。
	//
	// example:
	//
	// 血红蛋白用于携带氧气。解读报告时可结合报告上的参考范围。
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 可选消息附加信息，按实际返回处理。当前医疗模型不承诺返回医疗 Agent 的检索进度、资料来源或业务卡片。
	MetaData *MedicalModelInferenceResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
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

func (s MedicalModelInferenceResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetMetaData() *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetMetaData(v *MedicalModelInferenceResponseBodyDataMessagesMetaData) *MedicalModelInferenceResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetMimeType(v string) *MedicalModelInferenceResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetStatus(v string) *MedicalModelInferenceResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalModelInferenceResponseBodyDataMessagesMetaData struct {
	// 可选的消息附加标题，存在时按实际值使用。
	//
	// example:
	//
	// 问题分析中
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// 可选元数据内容，按实际返回类型解释。当前医疗模型不保证返回进度模块。
	//
	// example:
	//
	// {"body":"正在分析用户的问题。","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 为兼容已有响应形状保留的可选列表，可为空或省略；当前医疗模型不保证生成 Agent 业务卡片。
	MultiLoad []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// 可选的消息关联标识，存在时按实际值使用；本模型接口不保证返回。
	//
	// example:
	//
	// bar_thinking_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// 可选元数据类型，存在时按实际类型解释；不据此推断模型提供 Agent 过程能力。
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetMultiLoad() []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetBarTitle(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetMultiLoad(v []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetType(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) Validate() error {
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

type MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad struct {
	// 可选 JSON 内容，结构随实际附加项类型变化，不保证存在固定业务卡片。
	//
	// example:
	//
	// {"list":[{"image_url":"https://example.com/image.jpg","image_width":"800","image_height":"600"}]}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// 可选的附加项关联标识，存在时按实际值解释。
	//
	// example:
	//
	// image_waterfall_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// 可选的附加项类型，遇到未识别类型可忽略该附加项。
	//
	// example:
	//
	// image_waterfall
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyError struct {
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

func (s MedicalModelInferenceResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyError) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *MedicalModelInferenceResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *MedicalModelInferenceResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyError) SetCode(v interface{}) *MedicalModelInferenceResponseBodyError {
	s.Code = v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) SetMessage(v string) *MedicalModelInferenceResponseBodyError {
	s.Message = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) SetRequestId(v string) *MedicalModelInferenceResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) Validate() error {
	return dara.Validate(s)
}
