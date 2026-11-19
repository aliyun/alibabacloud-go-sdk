// This file is auto-generated, don't edit it. Thanks.
package streamVoice

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iVoiceEvent interface {
  dara.Model
  String() string
  GoString() string
  SetHeader(v *VoiceEventHeader) *VoiceEvent
  GetHeader() *VoiceEventHeader 
  SetCommunication(v *VoiceEventCommunication) *VoiceEvent
  GetCommunication() *VoiceEventCommunication 
  SetPayload(v *VoiceEventPayload) *VoiceEvent
  GetPayload() *VoiceEventPayload 
  SetError(v *VoiceEventError) *VoiceEvent
  GetError() *VoiceEventError 
}

type VoiceEvent struct {
  // 服务端事件头，包含事件类型及适用时的错误信息。
  // 
  // example:
  // 
  // {"event":"task-started"}
  Header *VoiceEventHeader `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
  // 服务端事件与当前语音交互的关联信息。
  // 
  // example:
  // 
  // {"requestId":"voice-request-001"}
  Communication *VoiceEventCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
  // 事件数据，按 header.event 区分。task-started 有 stage；result-generated 有 output；task-failed 有 stage/retryable；heartbeat 和 task-finished 的 payload 为 {}。字段未出现不等于其值为 0 或 false。
  // 
  // example:
  // 
  // {"stage":"asr"}
  Payload *VoiceEventPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
  // 可选的业务错误详情。发生业务失败时可返回，与 header 中的语音阶段错误信息配合处理；成功事件省略。
  // 
  // example:
  // 
  // {"code":"INVALID_REQUEST","message":"invalid request","requestId":"voice-request-001"}
  Error *VoiceEventError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
}

func (s VoiceEvent) String() string {
  return dara.Prettify(s)
}

func (s VoiceEvent) GoString() string {
  return s.String()
}

func (s *VoiceEvent) GetHeader() *VoiceEventHeader  {
  return s.Header
}

func (s *VoiceEvent) GetCommunication() *VoiceEventCommunication  {
  return s.Communication
}

func (s *VoiceEvent) GetPayload() *VoiceEventPayload  {
  return s.Payload
}

func (s *VoiceEvent) GetError() *VoiceEventError  {
  return s.Error
}

func (s *VoiceEvent) SetHeader(v *VoiceEventHeader) *VoiceEvent {
  s.Header = v
  return s
}

func (s *VoiceEvent) SetCommunication(v *VoiceEventCommunication) *VoiceEvent {
  s.Communication = v
  return s
}

func (s *VoiceEvent) SetPayload(v *VoiceEventPayload) *VoiceEvent {
  s.Payload = v
  return s
}

func (s *VoiceEvent) SetError(v *VoiceEventError) *VoiceEvent {
  s.Error = v
  return s
}

func (s *VoiceEvent) Validate() error {
  if s.Header != nil {
    if err := s.Header.Validate(); err != nil {
      return err
    }
  }
  if s.Communication != nil {
    if err := s.Communication.Validate(); err != nil {
      return err
    }
  }
  if s.Payload != nil {
    if err := s.Payload.Validate(); err != nil {
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

type VoiceEventHeader struct {
  // 必填事件类型，只可能为 task-started、result-generated、heartbeat、task-finished、task-failed。整轮完成以 task-finished 为准，不以 Agent complete 或单个 TTS 事件代替。
  // 
  // example:
  // 
  // task-started
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  // task-failed 时返回的语音阶段错误码，字符串；例如 INVALID_REQUEST、ASR_FAILED、AGENT_FAILED、TTS_FAILED，其他准入/协议错误见接口说明。若同时存在 Agent 业务错误，另看顶层 error.code，分别处理两层错误。
  // 
  // example:
  // 
  // INVALID_REQUEST
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // task-failed 时返回的可读失败说明；非失败事件省略，不依赖固定文案做程序判断。
  // 
  // example:
  // 
  // invalid request
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

func (s VoiceEventHeader) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventHeader) GoString() string {
  return s.String()
}

func (s *VoiceEventHeader) GetEvent() *string  {
  return s.Event
}

func (s *VoiceEventHeader) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *VoiceEventHeader) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *VoiceEventHeader) SetEvent(v string) *VoiceEventHeader {
  s.Event = &v
  return s
}

func (s *VoiceEventHeader) SetErrorCode(v string) *VoiceEventHeader {
  s.ErrorCode = &v
  return s
}

func (s *VoiceEventHeader) SetErrorMessage(v string) *VoiceEventHeader {
  s.ErrorMessage = &v
  return s
}

func (s *VoiceEventHeader) Validate() error {
  return dara.Validate(s)
}

type VoiceEventCommunication struct {
  // 对应当前 run-task 的 communication.requestId，用于轮次关联；无法解析出合法请求号时，整个 communication 可省略。不能用音频帧或事件数量替代关联号。
  // 
  // example:
  // 
  // voice-request-001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s VoiceEventCommunication) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventCommunication) GoString() string {
  return s.String()
}

func (s *VoiceEventCommunication) GetRequestId() *string  {
  return s.RequestId
}

func (s *VoiceEventCommunication) SetRequestId(v string) *VoiceEventCommunication {
  s.RequestId = &v
  return s
}

func (s *VoiceEventCommunication) Validate() error {
  return dara.Validate(s)
}

type VoiceEventPayload struct {
  // 条件返回；task-started 时为 asr，task-failed 时为 request/asr/agent/tts。result-generated 的阶段在 payload.output.stage 中。
  // 
  // example:
  // 
  // asr
  Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
  // 仅 task-failed 返回的布尔值；表示是否适合新建一轮重试，不表示可以重放原签名或使用不可用的 retryOf。其他事件无默认值。
  // 
  // example:
  // 
  // false
  Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
  // 仅 result-generated 返回；按 stage/type 识别具体结构。Agent 完整响应和兼容进度可能同时出现；使用 response 作为回答主数据，避免与 text/delta/messages 重复拼接。
  // 
  // example:
  // 
  // {"stage":"asr","text":"你好","final":true}
  Output *VoiceEventPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
}

func (s VoiceEventPayload) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayload) GoString() string {
  return s.String()
}

func (s *VoiceEventPayload) GetStage() *string  {
  return s.Stage
}

func (s *VoiceEventPayload) GetRetryable() *bool  {
  return s.Retryable
}

func (s *VoiceEventPayload) GetOutput() *VoiceEventPayloadOutput  {
  return s.Output
}

func (s *VoiceEventPayload) SetStage(v string) *VoiceEventPayload {
  s.Stage = &v
  return s
}

func (s *VoiceEventPayload) SetRetryable(v bool) *VoiceEventPayload {
  s.Retryable = &v
  return s
}

func (s *VoiceEventPayload) SetOutput(v *VoiceEventPayloadOutput) *VoiceEventPayload {
  s.Output = v
  return s
}

func (s *VoiceEventPayload) Validate() error {
  if s.Output != nil {
    if err := s.Output.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceEventPayloadOutput struct {
  // result-generated 的输出阶段，asr=语音识别、agent=Agent 回答、tts=语音合成。字段可用性取决于该阶段及 type。
  // 
  // example:
  // 
  // asr
  Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
  // 按阶段解释：Agent 完整帧=response、兼容状态进度=progress；TTS 为 sentence-begin/sentence-synthesis/sentence-end。ASR/兼容文字按 text/final 等字段解析；不要假设所有事件都返回全部 output 字段。
  // 
  // example:
  // 
  // response
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 仅 Agent type=response 返回，message 或 complete。complete 仅表示 Agent 流结束，可能仍需等待 TTS 音频；整轮结束以 task-finished 为准。
  // 
  // example:
  // 
  // message
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  // asr 阶段为最终转写。Agent 兼容终稿使用累计完整回答。
  // 
  // example:
  // 
  // 你好
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  // 仅兼容增量事件返回，使用完整 response 的客户端可以忽略本字段。
  // 
  // example:
  // 
  // 你好
  Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
  // ASR 最终转写及兼容 Agent 终稿为 true，兼容 Agent 文字增量为 false。
  // 
  // example:
  // 
  // true
  Final *bool `json:"final,omitempty" xml:"final,omitempty"`
  // type=progress 时返回状态栏消息批次，不拼接到正式回答。结构与 response.data.messages 中的状态栏消息一致。
  // 
  // example:
  // 
  // [{"mimeType":"bar/progress","status":"processing","content":"正在处理"}]
  Messages []interface{} `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
  // 仅 Agent type=response 时返回，与 Chat 流式帧一致的响应对象。complete 可携带最终正文，也可以为 {}；还可包含所用 Agent 的公开业务扩展。按字段说明读取正文、消息和错误信息。
  // 
  // example:
  // 
  // {"communication":{"requestId":"voice-request-001","sessionId":"voice-session-001","appId":"your-app","responseId":1},"data":{"messages":[{"mimeType":"multi_load/iframe","status":"complete","content":"你好，有什么可以帮助你？"}]},"errorCode":0,"errorMsg":""}
  Response *VoiceEventPayloadOutputResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
  // 仅 TTS 句子事件返回的关联对象；同一句 index 一致。sentence-synthesis 后紧跟该句的一帧二进制音频，按本轮 TTS format/sampleRate 解码；句子 text 仅在 sentence-begin 返回。
  // 
  // example:
  // 
  // {"index":0,"text":"你好"}
  Sentence *VoiceEventPayloadOutputSentence `json:"sentence,omitempty" xml:"sentence,omitempty" type:"Struct"`
}

func (s VoiceEventPayloadOutput) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutput) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutput) GetStage() *string  {
  return s.Stage
}

func (s *VoiceEventPayloadOutput) GetType() *string  {
  return s.Type
}

func (s *VoiceEventPayloadOutput) GetEvent() *string  {
  return s.Event
}

func (s *VoiceEventPayloadOutput) GetText() *string  {
  return s.Text
}

func (s *VoiceEventPayloadOutput) GetDelta() *string  {
  return s.Delta
}

func (s *VoiceEventPayloadOutput) GetFinal() *bool  {
  return s.Final
}

func (s *VoiceEventPayloadOutput) GetMessages() []interface{}  {
  return s.Messages
}

func (s *VoiceEventPayloadOutput) GetResponse() *VoiceEventPayloadOutputResponse  {
  return s.Response
}

func (s *VoiceEventPayloadOutput) GetSentence() *VoiceEventPayloadOutputSentence  {
  return s.Sentence
}

func (s *VoiceEventPayloadOutput) SetStage(v string) *VoiceEventPayloadOutput {
  s.Stage = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetType(v string) *VoiceEventPayloadOutput {
  s.Type = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetEvent(v string) *VoiceEventPayloadOutput {
  s.Event = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetText(v string) *VoiceEventPayloadOutput {
  s.Text = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetDelta(v string) *VoiceEventPayloadOutput {
  s.Delta = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetFinal(v bool) *VoiceEventPayloadOutput {
  s.Final = &v
  return s
}

func (s *VoiceEventPayloadOutput) SetMessages(v []interface{}) *VoiceEventPayloadOutput {
  s.Messages = v
  return s
}

func (s *VoiceEventPayloadOutput) SetResponse(v *VoiceEventPayloadOutputResponse) *VoiceEventPayloadOutput {
  s.Response = v
  return s
}

func (s *VoiceEventPayloadOutput) SetSentence(v *VoiceEventPayloadOutputSentence) *VoiceEventPayloadOutput {
  s.Sentence = v
  return s
}

func (s *VoiceEventPayloadOutput) Validate() error {
  if s.Response != nil {
    if err := s.Response.Validate(); err != nil {
      return err
    }
  }
  if s.Sentence != nil {
    if err := s.Sentence.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceEventPayloadOutputResponse struct {
  // 用于关联此 Agent 响应的业务请求。
  // 
  // example:
  // 
  // voice-request-001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // 与外层语音帧的 communication 独立，来自对应 Chat 响应。
  // 
  // example:
  // 
  // {"requestId":"voice-request-001","sessionId":"voice-session-001","appId":"your-app","responseId":1}
  Communication *VoiceEventPayloadOutputResponseCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
  // 可携带正文快照、图文卡片、过程消息和业务扩展。
  // 
  // example:
  // 
  // {"messages":[{"mimeType":"multi_load/iframe","status":"complete","content":"你好"}]}
  Data *VoiceEventPayloadOutputResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // 该 Agent 响应中的业务状态。
  // 
  // example:
  // 
  // 0
  ErrorCode interface{} `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // 与 errorCode 对应的说明。
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s VoiceEventPayloadOutputResponse) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponse) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponse) GetRequestId() *string  {
  return s.RequestId
}

func (s *VoiceEventPayloadOutputResponse) GetCommunication() *VoiceEventPayloadOutputResponseCommunication  {
  return s.Communication
}

func (s *VoiceEventPayloadOutputResponse) GetData() *VoiceEventPayloadOutputResponseData  {
  return s.Data
}

func (s *VoiceEventPayloadOutputResponse) GetErrorCode() interface{}  {
  return s.ErrorCode
}

func (s *VoiceEventPayloadOutputResponse) GetErrorMsg() *string  {
  return s.ErrorMsg
}

func (s *VoiceEventPayloadOutputResponse) SetRequestId(v string) *VoiceEventPayloadOutputResponse {
  s.RequestId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponse) SetCommunication(v *VoiceEventPayloadOutputResponseCommunication) *VoiceEventPayloadOutputResponse {
  s.Communication = v
  return s
}

func (s *VoiceEventPayloadOutputResponse) SetData(v *VoiceEventPayloadOutputResponseData) *VoiceEventPayloadOutputResponse {
  s.Data = v
  return s
}

func (s *VoiceEventPayloadOutputResponse) SetErrorCode(v interface{}) *VoiceEventPayloadOutputResponse {
  s.ErrorCode = v
  return s
}

func (s *VoiceEventPayloadOutputResponse) SetErrorMsg(v string) *VoiceEventPayloadOutputResponse {
  s.ErrorMsg = &v
  return s
}

func (s *VoiceEventPayloadOutputResponse) Validate() error {
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
  return nil
}

type VoiceEventPayloadOutputResponseCommunication struct {
  // Agent 业务响应关联号。
  // 
  // example:
  // 
  // voice-request-001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // 当前 Agent 响应所属的客户会话。
  // 
  // example:
  // 
  // voice-session-001
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // 当前响应所属的业务应用。
  // 
  // example:
  // 
  // your-app
  AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
  // 该值不代表网关发送帧的顺序。
  // 
  // example:
  // 
  // 1
  ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
}

func (s VoiceEventPayloadOutputResponseCommunication) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseCommunication) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseCommunication) GetRequestId() *string  {
  return s.RequestId
}

func (s *VoiceEventPayloadOutputResponseCommunication) GetSessionId() *string  {
  return s.SessionId
}

func (s *VoiceEventPayloadOutputResponseCommunication) GetAppId() *string  {
  return s.AppId
}

func (s *VoiceEventPayloadOutputResponseCommunication) GetResponseId() *int64  {
  return s.ResponseId
}

func (s *VoiceEventPayloadOutputResponseCommunication) SetRequestId(v string) *VoiceEventPayloadOutputResponseCommunication {
  s.RequestId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseCommunication) SetSessionId(v string) *VoiceEventPayloadOutputResponseCommunication {
  s.SessionId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseCommunication) SetAppId(v string) *VoiceEventPayloadOutputResponseCommunication {
  s.AppId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseCommunication) SetResponseId(v int64) *VoiceEventPayloadOutputResponseCommunication {
  s.ResponseId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseCommunication) Validate() error {
  return dara.Validate(s)
}

type VoiceEventPayloadOutputResponseData struct {
  // 文字、卡片和状态栏消息按 mimeType 处理。
  // 
  // example:
  // 
  // [{"mimeType":"multi_load/iframe","status":"complete","content":"你好"}]
  Messages []*VoiceEventPayloadOutputResponseDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
  // 可选的调试信息，仅在 Agent 响应提供时出现，结构由所用 Agent 的公开约定确定。不保证存在或固定结构，不能作为正文或整轮完成的判断依据。
  // 
  // example:
  // 
  // {}
  Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
  // 可含 inputResources 等资源处理结果。
  // 
  // example:
  // 
  // {"inputResources":[]}
  ExtraInfo *VoiceEventPayloadOutputResponseDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
}

func (s VoiceEventPayloadOutputResponseData) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseData) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseData) GetMessages() []*VoiceEventPayloadOutputResponseDataMessages  {
  return s.Messages
}

func (s *VoiceEventPayloadOutputResponseData) GetDebug() interface{}  {
  return s.Debug
}

func (s *VoiceEventPayloadOutputResponseData) GetExtraInfo() *VoiceEventPayloadOutputResponseDataExtraInfo  {
  return s.ExtraInfo
}

func (s *VoiceEventPayloadOutputResponseData) SetMessages(v []*VoiceEventPayloadOutputResponseDataMessages) *VoiceEventPayloadOutputResponseData {
  s.Messages = v
  return s
}

func (s *VoiceEventPayloadOutputResponseData) SetDebug(v interface{}) *VoiceEventPayloadOutputResponseData {
  s.Debug = v
  return s
}

func (s *VoiceEventPayloadOutputResponseData) SetExtraInfo(v *VoiceEventPayloadOutputResponseDataExtraInfo) *VoiceEventPayloadOutputResponseData {
  s.ExtraInfo = v
  return s
}

func (s *VoiceEventPayloadOutputResponseData) Validate() error {
  if s.Messages != nil {
    for _, item := range s.Messages {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.ExtraInfo != nil {
    if err := s.ExtraInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceEventPayloadOutputResponseDataMessages struct {
  // 例如 multi_load/iframe、bar/progress。
  // 
  // example:
  // 
  // multi_load/iframe
  MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
  // 例如 processing、complete，不代表整轮完成。
  // 
  // example:
  // 
  // complete
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // 文本或结构化内容，格式由 mimeType 决定。
  // 
  // example:
  // 
  // 你好
  Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
  // 结构随消息类型变化，可含卡片、状态栏过程信息及关联标识。
  // 
  // example:
  // 
  // {"sourceSeq":"card-001","type":"bar_thinking","barTitle":"思考中"}
  MetaData *VoiceEventPayloadOutputResponseDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
}

func (s VoiceEventPayloadOutputResponseDataMessages) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseDataMessages) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseDataMessages) GetMimeType() *string  {
  return s.MimeType
}

func (s *VoiceEventPayloadOutputResponseDataMessages) GetStatus() *string  {
  return s.Status
}

func (s *VoiceEventPayloadOutputResponseDataMessages) GetContent() interface{}  {
  return s.Content
}

func (s *VoiceEventPayloadOutputResponseDataMessages) GetMetaData() *VoiceEventPayloadOutputResponseDataMessagesMetaData  {
  return s.MetaData
}

func (s *VoiceEventPayloadOutputResponseDataMessages) SetMimeType(v string) *VoiceEventPayloadOutputResponseDataMessages {
  s.MimeType = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessages) SetStatus(v string) *VoiceEventPayloadOutputResponseDataMessages {
  s.Status = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessages) SetContent(v interface{}) *VoiceEventPayloadOutputResponseDataMessages {
  s.Content = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessages) SetMetaData(v *VoiceEventPayloadOutputResponseDataMessagesMetaData) *VoiceEventPayloadOutputResponseDataMessages {
  s.MetaData = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessages) Validate() error {
  if s.MetaData != nil {
    if err := s.MetaData.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceEventPayloadOutputResponseDataMessagesMetaData struct {
  // 用于匹配正文中的卡片挂载标记，或区分状态栏模块。
  // 
  // example:
  // 
  // card-001
  SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
  // 由 Agent 提供，例如 bar_thinking 或 bar_tool。
  // 
  // example:
  // 
  // bar_thinking
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // Agent 返回时可用于展示。
  // 
  // example:
  // 
  // 思考中
  BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
  // 格式取决于 metaData.type。
  // 
  // example:
  // 
  // 正在整理回答
  Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
  // 卡片结构随类型变化，按关联标识挂载。
  // 
  // example:
  // 
  // [{"type":"example-card","sourceSeq":"card-001","content":{}}]
  MultiLoad []*VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
}

func (s VoiceEventPayloadOutputResponseDataMessagesMetaData) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseDataMessagesMetaData) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) GetSourceSeq() *string  {
  return s.SourceSeq
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) GetType() *string  {
  return s.Type
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) GetBarTitle() *string  {
  return s.BarTitle
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) GetContent() interface{}  {
  return s.Content
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) GetMultiLoad() []*VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad  {
  return s.MultiLoad
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) SetSourceSeq(v string) *VoiceEventPayloadOutputResponseDataMessagesMetaData {
  s.SourceSeq = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) SetType(v string) *VoiceEventPayloadOutputResponseDataMessagesMetaData {
  s.Type = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) SetBarTitle(v string) *VoiceEventPayloadOutputResponseDataMessagesMetaData {
  s.BarTitle = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) SetContent(v interface{}) *VoiceEventPayloadOutputResponseDataMessagesMetaData {
  s.Content = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) SetMultiLoad(v []*VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) *VoiceEventPayloadOutputResponseDataMessagesMetaData {
  s.MultiLoad = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaData) Validate() error {
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

type VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad struct {
  // 按类型选择卡片展示组件。
  // 
  // example:
  // 
  // example-card
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 与正文中的卡片引用关联。相同标识代表同一卡片。
  // 
  // example:
  // 
  // card-001
  SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
  // 可选，结构随卡片类型变化。图文卡片可含资源列表。
  // 
  // example:
  // 
  // {}
  Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
}

func (s VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) GetType() *string  {
  return s.Type
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) GetSourceSeq() *string  {
  return s.SourceSeq
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) GetContent() interface{}  {
  return s.Content
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) SetType(v string) *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad {
  s.Type = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad {
  s.SourceSeq = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad {
  s.Content = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataMessagesMetaDataMultiLoad) Validate() error {
  return dara.Validate(s)
}

type VoiceEventPayloadOutputResponseDataExtraInfo struct {
  // 仅在 Agent 返回资源处理结果时出现。
  // 
  // example:
  // 
  // [{"contentIndex":1,"type":"image","resourceId":"resource-example-001","url":"https://example.com/image.png","fileName":"image.png"}]
  InputResources []*VoiceEventPayloadOutputResponseDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s VoiceEventPayloadOutputResponseDataExtraInfo) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseDataExtraInfo) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfo) GetInputResources() []*VoiceEventPayloadOutputResponseDataExtraInfoInputResources  {
  return s.InputResources
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfo) SetInputResources(v []*VoiceEventPayloadOutputResponseDataExtraInfoInputResources) *VoiceEventPayloadOutputResponseDataExtraInfo {
  s.InputResources = v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfo) Validate() error {
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

type VoiceEventPayloadOutputResponseDataExtraInfoInputResources struct {
  // 以本轮原始 content 数组的顺序计数，从零开始，文本项也计数。
  // 
  // example:
  // 
  // 1
  ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
  // 用于区分本轮输入资源的类型。
  // 
  // example:
  // 
  // image
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 资源服务返回的标识。
  // 
  // example:
  // 
  // resource-example-001
  ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
  // 资源处理成功时返回的文件 URL。
  // 
  // example:
  // 
  // https://example.com/image.png
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // 可选的资源文件名。
  // 
  // example:
  // 
  // image.png
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s VoiceEventPayloadOutputResponseDataExtraInfoInputResources) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GetType() *string  {
  return s.Type
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GetResourceId() *string  {
  return s.ResourceId
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GetUrl() *string  {
  return s.Url
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) GetFileName() *string  {
  return s.FileName
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) SetContentIndex(v int32) *VoiceEventPayloadOutputResponseDataExtraInfoInputResources {
  s.ContentIndex = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) SetType(v string) *VoiceEventPayloadOutputResponseDataExtraInfoInputResources {
  s.Type = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) SetResourceId(v string) *VoiceEventPayloadOutputResponseDataExtraInfoInputResources {
  s.ResourceId = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) SetUrl(v string) *VoiceEventPayloadOutputResponseDataExtraInfoInputResources {
  s.Url = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) SetFileName(v string) *VoiceEventPayloadOutputResponseDataExtraInfoInputResources {
  s.FileName = &v
  return s
}

func (s *VoiceEventPayloadOutputResponseDataExtraInfoInputResources) Validate() error {
  return dara.Validate(s)
}

type VoiceEventPayloadOutputSentence struct {
  // TTS 句子事件返回的整数索引，同一句的多个事件共用该值，用于关联随后音频；不能作为整轮完成标记。
  // 
  // example:
  // 
  // 0
  Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
  // 仅 TTS sentence-begin 返回的待合成文本；sentence-synthesis/sentence-end 不保证存在。
  // 
  // example:
  // 
  // 你好
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s VoiceEventPayloadOutputSentence) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventPayloadOutputSentence) GoString() string {
  return s.String()
}

func (s *VoiceEventPayloadOutputSentence) GetIndex() *int32  {
  return s.Index
}

func (s *VoiceEventPayloadOutputSentence) GetText() *string  {
  return s.Text
}

func (s *VoiceEventPayloadOutputSentence) SetIndex(v int32) *VoiceEventPayloadOutputSentence {
  s.Index = &v
  return s
}

func (s *VoiceEventPayloadOutputSentence) SetText(v string) *VoiceEventPayloadOutputSentence {
  s.Text = &v
  return s
}

func (s *VoiceEventPayloadOutputSentence) Validate() error {
  return dara.Validate(s)
}

type VoiceEventError struct {
  // 业务错误码，可为字符串或数值，具体类型由所用 Agent 的公开接口约定确定。它与 header.errorCode 的语音阶段错误码含义不同，应分别处理。
  // 
  // example:
  // 
  // INVALID_REQUEST
  Code interface{} `json:"code,omitempty" xml:"code,omitempty"`
  // 公开业务错误说明，可读文本；不要依赖固定全文做程序判断。
  // 
  // example:
  // 
  // invalid request
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // 失败调用的业务请求关联号，用于排查；不是握手平台 RequestId、订单号或实例 ID。
  // 
  // example:
  // 
  // voice-request-001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s VoiceEventError) String() string {
  return dara.Prettify(s)
}

func (s VoiceEventError) GoString() string {
  return s.String()
}

func (s *VoiceEventError) GetCode() interface{}  {
  return s.Code
}

func (s *VoiceEventError) GetMessage() *string  {
  return s.Message
}

func (s *VoiceEventError) GetRequestId() *string  {
  return s.RequestId
}

func (s *VoiceEventError) SetCode(v interface{}) *VoiceEventError {
  s.Code = v
  return s
}

func (s *VoiceEventError) SetMessage(v string) *VoiceEventError {
  s.Message = &v
  return s
}

func (s *VoiceEventError) SetRequestId(v string) *VoiceEventError {
  s.RequestId = &v
  return s
}

func (s *VoiceEventError) Validate() error {
  return dara.Validate(s)
}

