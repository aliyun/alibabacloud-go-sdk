// This file is auto-generated, don't edit it. Thanks.
package streamVoice

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iVoiceCommand interface {
  dara.Model
  String() string
  GoString() string
  SetCommunication(v *VoiceCommandCommunication) *VoiceCommand
  GetCommunication() *VoiceCommandCommunication 
  SetData(v *VoiceCommandData) *VoiceCommand
  GetData() *VoiceCommandData 
  SetParams(v *VoiceCommandParams) *VoiceCommand
  GetParams() *VoiceCommandParams 
  SetHeader(v *VoiceCommandHeader) *VoiceCommand
  GetHeader() *VoiceCommandHeader 
  SetPayload(v *VoiceCommandPayload) *VoiceCommand
  GetPayload() *VoiceCommandPayload 
}

type VoiceCommand struct {
  // 本轮语音交互的请求关联号、会话及应用和用户身份信息。
  // 
  // example:
  // 
  // {"appId":"your-app","requestId":"voice-request-001","sessionId":"voice-session-001","historyOwner":"GATEWAY"}
  Communication *VoiceCommandCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
  // run-task 可选对象，无默认值；传入时只支持 messages，不能为 null。供 Agent 使用的业务输入，语音音频另通过二进制帧上传。
  // 
  // example:
  // 
  // {"messages":[{"role":"user","content":"请结合我的语音回答。"}]}
  Data *VoiceCommandData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // run-task 可选对象，省略为无客户扩展；显式传入不能为 null。支持 directAgent、deepthink 及已接入 Agent 的扩展，字段和值类型依 Agent 契约；语音参数放在 payload.parameters。本接口不支持通过 enableTextImageMixed 启用图文混排，传 true 不生效；历史托管方放在 communication.historyOwner。
  // 
  // example:
  // 
  // {"deepthink":"true"}
  Params *VoiceCommandParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
  // 控制帧头，包含客户端要执行的语音交互动作。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"action":"run-task"}
  Header *VoiceCommandHeader `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
  // run-task 必须提供 taskGroup=audio、task=voice 和 parameters.asr（除兼容保留的 retryOf 分支，该分支当前不可用）。finish-audio/cancel-task 的 payload 可省略、为 null 或空对象。
  // 
  // example:
  // 
  // {"taskGroup":"audio","task":"voice","parameters":{"asr":{"format":"pcm","sampleRate":16000}}}
  Payload *VoiceCommandPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s VoiceCommand) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommand) GoString() string {
  return s.String()
}

func (s *VoiceCommand) GetCommunication() *VoiceCommandCommunication  {
  return s.Communication
}

func (s *VoiceCommand) GetData() *VoiceCommandData  {
  return s.Data
}

func (s *VoiceCommand) GetParams() *VoiceCommandParams  {
  return s.Params
}

func (s *VoiceCommand) GetHeader() *VoiceCommandHeader  {
  return s.Header
}

func (s *VoiceCommand) GetPayload() *VoiceCommandPayload  {
  return s.Payload
}

func (s *VoiceCommand) SetCommunication(v *VoiceCommandCommunication) *VoiceCommand {
  s.Communication = v
  return s
}

func (s *VoiceCommand) SetData(v *VoiceCommandData) *VoiceCommand {
  s.Data = v
  return s
}

func (s *VoiceCommand) SetParams(v *VoiceCommandParams) *VoiceCommand {
  s.Params = v
  return s
}

func (s *VoiceCommand) SetHeader(v *VoiceCommandHeader) *VoiceCommand {
  s.Header = v
  return s
}

func (s *VoiceCommand) SetPayload(v *VoiceCommandPayload) *VoiceCommand {
  s.Payload = v
  return s
}

func (s *VoiceCommand) Validate() error {
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
  if s.Params != nil {
    if err := s.Params.Validate(); err != nil {
      return err
    }
  }
  if s.Header != nil {
    if err := s.Header.Validate(); err != nil {
      return err
    }
  }
  if s.Payload != nil {
    if err := s.Payload.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceCommandCommunication struct {
  // run-task 必填，本轮客户业务请求 ID，匹配 [A-Za-z0-9._:-]{1,128}。不能使用 gw-ingress- 保留前缀。下一轮需新建连接并使用独立 ID；finish-audio/cancel-task 沿用当前轮次且不另传 communication。与可选握手 X-Request-Id 不同。
  // 
  // example:
  // 
  // voice-request-001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // 每个 run-task 必填。客户会话 ID，匹配 [A-Za-z0-9._-]{1,128}，同一会话续聊时复用；用于会话历史关联，不是订单号或业务请求 ID。
  // 
  // example:
  // 
  // voice-session-001
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // 每个 run-task 必填的业务应用 ID，不从握手 query 读取。须符合业务网关凭据的应用绑定或委托配置；每条连接只接受一个 run-task；下一轮需新建连接，即使复用 sessionId 也要提供本字段。非空，最多 256 个 UTF-16 代码单元，无首尾空白、控制字符或非法 Unicode。
  // 
  // example:
  // 
  // your-app
  AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
  // 每个 run-task 必填，区分大小写，只能是 CLIENT 或 GATEWAY，无默认值。须与该会话首次调用一致。CLIENT 历史放在 payload.input.history；GATEWAY 由服务端托管，不接受客户提交非空历史。
  // 
  // example:
  // 
  // GATEWAY
  HistoryOwner *string `json:"historyOwner,omitempty" xml:"historyOwner,omitempty"`
  // 可选，无默认用户。省略时匿名，不启用用户记忆；即使凭据绑定固定用户也不要求必须传入。若显式传入，须符合固定绑定；必须为非空字符串，最多 256 个 UTF-16 代码单元，无首尾空白、控制字符或非法 Unicode。不能为 null。
  // 
  // example:
  // 
  // customer-user-001
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s VoiceCommandCommunication) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandCommunication) GoString() string {
  return s.String()
}

func (s *VoiceCommandCommunication) GetRequestId() *string  {
  return s.RequestId
}

func (s *VoiceCommandCommunication) GetSessionId() *string  {
  return s.SessionId
}

func (s *VoiceCommandCommunication) GetAppId() *string  {
  return s.AppId
}

func (s *VoiceCommandCommunication) GetHistoryOwner() *string  {
  return s.HistoryOwner
}

func (s *VoiceCommandCommunication) GetUserId() *string  {
  return s.UserId
}

func (s *VoiceCommandCommunication) SetRequestId(v string) *VoiceCommandCommunication {
  s.RequestId = &v
  return s
}

func (s *VoiceCommandCommunication) SetSessionId(v string) *VoiceCommandCommunication {
  s.SessionId = &v
  return s
}

func (s *VoiceCommandCommunication) SetAppId(v string) *VoiceCommandCommunication {
  s.AppId = &v
  return s
}

func (s *VoiceCommandCommunication) SetHistoryOwner(v string) *VoiceCommandCommunication {
  s.HistoryOwner = &v
  return s
}

func (s *VoiceCommandCommunication) SetUserId(v string) *VoiceCommandCommunication {
  s.UserId = &v
  return s
}

func (s *VoiceCommandCommunication) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandData struct {
  // run-task 可选，省略或空数组表示没有附加图文消息。传入时只包含本轮一条 user 消息，历史放入 payload.input.history；不要重复提交音频转写。
  // 
  // example:
  // 
  // [{"role":"user","content":"请结合我的语音回答。"}]
  Messages []*VoiceCommandDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s VoiceCommandData) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandData) GoString() string {
  return s.String()
}

func (s *VoiceCommandData) GetMessages() []*VoiceCommandDataMessages  {
  return s.Messages
}

func (s *VoiceCommandData) SetMessages(v []*VoiceCommandDataMessages) *VoiceCommandData {
  s.Messages = v
  return s
}

func (s *VoiceCommandData) Validate() error {
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

type VoiceCommandDataMessages struct {
  // 本轮消息的角色，传 user。CLIENT 托管的历史消息放在 payload.input.history。
  // 
  // example:
  // 
  // user
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  // 本轮 user 消息的内容，为非空文本或所接入 Agent 支持的内容块数组；文本/图片块格式沿用 Chat。音频不放入此字段，另发 WSS 二进制帧。
  // 
  // example:
  // 
  // 请结合我的语音回答。
  Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
}

func (s VoiceCommandDataMessages) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandDataMessages) GoString() string {
  return s.String()
}

func (s *VoiceCommandDataMessages) GetRole() *string  {
  return s.Role
}

func (s *VoiceCommandDataMessages) GetContent() interface{}  {
  return s.Content
}

func (s *VoiceCommandDataMessages) SetRole(v string) *VoiceCommandDataMessages {
  s.Role = &v
  return s
}

func (s *VoiceCommandDataMessages) SetContent(v interface{}) *VoiceCommandDataMessages {
  s.Content = v
  return s
}

func (s *VoiceCommandDataMessages) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandParams struct {
  // 可选的本轮 Agent 选择器；省略使用租户默认路由，显式值须在已授权 Agent 范围内。无统一公开枚举，请使用接入方提供的 Agent 名称。
  // 
  // example:
  // 
  // <已授权的Agent名称>
  DirectAgent *string `json:"directAgent,omitempty" xml:"directAgent,omitempty"`
  // 可选的 Agent 参数，无统一网关默认值。支持本参数的 Agent 使用字符串 true 开启，具体值和默认行为依所接入 Agent 契约；不要误写为控制帧顶层字段。
  // 
  // example:
  // 
  // true
  Deepthink *string `json:"deepthink,omitempty" xml:"deepthink,omitempty"`
}

func (s VoiceCommandParams) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandParams) GoString() string {
  return s.String()
}

func (s *VoiceCommandParams) GetDirectAgent() *string  {
  return s.DirectAgent
}

func (s *VoiceCommandParams) GetDeepthink() *string  {
  return s.Deepthink
}

func (s *VoiceCommandParams) SetDirectAgent(v string) *VoiceCommandParams {
  s.DirectAgent = &v
  return s
}

func (s *VoiceCommandParams) SetDeepthink(v string) *VoiceCommandParams {
  s.Deepthink = &v
  return s
}

func (s *VoiceCommandParams) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandHeader struct {
  // 所有控制帧必填，无默认值；区分大小写，只接受 run-task、finish-audio、cancel-task。run-task 启动一轮交互；finish-audio 表示上行音频发送结束；cancel-task 取消当前轮次。没有 continue-task 或 finish-task；Ping/Pong 使用 WebSocket 原生控制帧。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // run-task
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s VoiceCommandHeader) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandHeader) GoString() string {
  return s.String()
}

func (s *VoiceCommandHeader) GetAction() *string  {
  return s.Action
}

func (s *VoiceCommandHeader) SetAction(v string) *VoiceCommandHeader {
  s.Action = &v
  return s
}

func (s *VoiceCommandHeader) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandPayload struct {
  // run-task 必填字符串，固定为 audio，区分大小写，无默认值。
  // 
  // example:
  // 
  // audio
  TaskGroup *string `json:"taskGroup,omitempty" xml:"taskGroup,omitempty"`
  // run-task 必填字符串，固定为 voice，区分大小写，无默认值。
  // 
  // example:
  // 
  // voice
  Task *string `json:"task,omitempty" xml:"task,omitempty"`
  // 可选的保留字段，字符串格式为 [A-Za-z0-9._:-]{1,128}。当前不支持引用先前请求重试；即使值符合格式，也返回 RETRY_EXPIRED，不会重试音频或继续之前的任务。正常调用请省略，不要用于自动续传。
  // 
  // example:
  // 
  // previous-voice-request-001
  RetryOf *string `json:"retryOf,omitempty" xml:"retryOf,omitempty"`
  // 普通 run-task 必填对象，只包含 asr 和可选 tts。asr 无默认值；tts 可省略或 null，使用默认合成设置。
  // 
  // example:
  // 
  // {"asr":{"format":"pcm","sampleRate":16000}}
  Parameters *VoiceCommandPayloadParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
  // 可选对象，可省略或 null；只包含 context 和 history。context 辅助 ASR，history 是 Agent 业务历史，两者用途不同。
  // 
  // example:
  // 
  // {"context":[{"role":"user","content":[{"type":"input_text","text":"千问语音对话"}]}]}
  Input *VoiceCommandPayloadInput `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
}

func (s VoiceCommandPayload) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayload) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayload) GetTaskGroup() *string  {
  return s.TaskGroup
}

func (s *VoiceCommandPayload) GetTask() *string  {
  return s.Task
}

func (s *VoiceCommandPayload) GetRetryOf() *string  {
  return s.RetryOf
}

func (s *VoiceCommandPayload) GetParameters() *VoiceCommandPayloadParameters  {
  return s.Parameters
}

func (s *VoiceCommandPayload) GetInput() *VoiceCommandPayloadInput  {
  return s.Input
}

func (s *VoiceCommandPayload) SetTaskGroup(v string) *VoiceCommandPayload {
  s.TaskGroup = &v
  return s
}

func (s *VoiceCommandPayload) SetTask(v string) *VoiceCommandPayload {
  s.Task = &v
  return s
}

func (s *VoiceCommandPayload) SetRetryOf(v string) *VoiceCommandPayload {
  s.RetryOf = &v
  return s
}

func (s *VoiceCommandPayload) SetParameters(v *VoiceCommandPayloadParameters) *VoiceCommandPayload {
  s.Parameters = v
  return s
}

func (s *VoiceCommandPayload) SetInput(v *VoiceCommandPayloadInput) *VoiceCommandPayload {
  s.Input = v
  return s
}

func (s *VoiceCommandPayload) Validate() error {
  if s.Parameters != nil {
    if err := s.Parameters.Validate(); err != nil {
      return err
    }
  }
  if s.Input != nil {
    if err := s.Input.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceCommandPayloadParameters struct {
  // 普通 run-task 必填对象；format 和 sampleRate 必填且无默认值，languageHints 可选。上传与这些参数匹配的音频二进制帧，不使用 Base64 JSON 音频字段。
  // 
  // example:
  // 
  // {"format":"pcm","sampleRate":16000}
  Asr *VoiceCommandPayloadParametersAsr `json:"asr,omitempty" xml:"asr,omitempty" type:"Struct"`
  // 可选对象，可省略或传 null，使用默认合成设置。默认 format=mp3、sampleRate=22050、volume=50、rate=1.0，voice 使用租户默认音色；显式字段仍须符合租户许可。
  // 
  // example:
  // 
  // {"format":"mp3","sampleRate":22050,"volume":50,"rate":1.0}
  Tts *VoiceCommandPayloadParametersTts `json:"tts,omitempty" xml:"tts,omitempty" type:"Struct"`
}

func (s VoiceCommandPayloadParameters) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadParameters) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadParameters) GetAsr() *VoiceCommandPayloadParametersAsr  {
  return s.Asr
}

func (s *VoiceCommandPayloadParameters) GetTts() *VoiceCommandPayloadParametersTts  {
  return s.Tts
}

func (s *VoiceCommandPayloadParameters) SetAsr(v *VoiceCommandPayloadParametersAsr) *VoiceCommandPayloadParameters {
  s.Asr = v
  return s
}

func (s *VoiceCommandPayloadParameters) SetTts(v *VoiceCommandPayloadParametersTts) *VoiceCommandPayloadParameters {
  s.Tts = v
  return s
}

func (s *VoiceCommandPayloadParameters) Validate() error {
  if s.Asr != nil {
    if err := s.Asr.Validate(); err != nil {
      return err
    }
  }
  if s.Tts != nil {
    if err := s.Tts.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type VoiceCommandPayloadParametersAsr struct {
  // ASR 必填字符串，无默认值；支持 pcm、wav 或 opus，具体值须符合租户许可。音频实际编码须与声明一致，不能用文件后缀代替编码。
  // 
  // example:
  // 
  // pcm
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  // ASR 必填整数，单位 Hz，无默认值；须为租户许可的采样率并与实际音频一致，例如 16000。
  // 
  // example:
  // 
  // 16000
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // 可选字符串数组，无默认提示；最多 4 项，每项非空白、最多 16 个 UTF-16 代码单元。用于辅助识别，非业务用户身份。
  // 
  // example:
  // 
  // ["zh"]
  LanguageHints []*string `json:"languageHints,omitempty" xml:"languageHints,omitempty" type:"Repeated"`
}

func (s VoiceCommandPayloadParametersAsr) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadParametersAsr) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadParametersAsr) GetFormat() *string  {
  return s.Format
}

func (s *VoiceCommandPayloadParametersAsr) GetSampleRate() *int  {
  return s.SampleRate
}

func (s *VoiceCommandPayloadParametersAsr) GetLanguageHints() []*string  {
  return s.LanguageHints
}

func (s *VoiceCommandPayloadParametersAsr) SetFormat(v string) *VoiceCommandPayloadParametersAsr {
  s.Format = &v
  return s
}

func (s *VoiceCommandPayloadParametersAsr) SetSampleRate(v int) *VoiceCommandPayloadParametersAsr {
  s.SampleRate = &v
  return s
}

func (s *VoiceCommandPayloadParametersAsr) SetLanguageHints(v []*string) *VoiceCommandPayloadParametersAsr {
  s.LanguageHints = v
  return s
}

func (s *VoiceCommandPayloadParametersAsr) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandPayloadParametersTts struct {
  // 可选字符串，只支持 mp3 或 pcm；省略或 null 时为 mp3。客户端按该格式解码随后返回的二进制音频。
  // 
  // example:
  // 
  // mp3
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  // 可选整数，单位 Hz；默认 22050，可选 8000、16000、22050、24000、44100、48000。客户端按本轮格式及采样率解码音频。
  // 
  // example:
  // 
  // 22050
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // 可选字符串，须为租户许可的音色；省略或 null 时使用租户默认音色。没有跨租户固定的默认音色名称。
  // 
  // example:
  // 
  // <租户已许可的音色>
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  // 可选整数，范围 [0,100]，默认 50；不是小数比例。
  // 
  // example:
  // 
  // 50
  Volume *int `json:"volume,omitempty" xml:"volume,omitempty"`
  // 可选数值，范围 [0.5,2.0]，默认 1.0，表示语速倍率。
  // 
  // example:
  // 
  // 1.0
  Rate *float32 `json:"rate,omitempty" xml:"rate,omitempty"`
}

func (s VoiceCommandPayloadParametersTts) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadParametersTts) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadParametersTts) GetFormat() *string  {
  return s.Format
}

func (s *VoiceCommandPayloadParametersTts) GetSampleRate() *int  {
  return s.SampleRate
}

func (s *VoiceCommandPayloadParametersTts) GetVoice() *string  {
  return s.Voice
}

func (s *VoiceCommandPayloadParametersTts) GetVolume() *int  {
  return s.Volume
}

func (s *VoiceCommandPayloadParametersTts) GetRate() *float32  {
  return s.Rate
}

func (s *VoiceCommandPayloadParametersTts) SetFormat(v string) *VoiceCommandPayloadParametersTts {
  s.Format = &v
  return s
}

func (s *VoiceCommandPayloadParametersTts) SetSampleRate(v int) *VoiceCommandPayloadParametersTts {
  s.SampleRate = &v
  return s
}

func (s *VoiceCommandPayloadParametersTts) SetVoice(v string) *VoiceCommandPayloadParametersTts {
  s.Voice = &v
  return s
}

func (s *VoiceCommandPayloadParametersTts) SetVolume(v int) *VoiceCommandPayloadParametersTts {
  s.Volume = &v
  return s
}

func (s *VoiceCommandPayloadParametersTts) SetRate(v float32) *VoiceCommandPayloadParametersTts {
  s.Rate = &v
  return s
}

func (s *VoiceCommandPayloadParametersTts) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandPayloadInput struct {
  // 可选数组，省略为无 ASR 上下文。最多 10 条，user/assistant 各最多 5 条，assistant 必须紧跟 user。每条 content 文本块顺序合并后最多 400 个 UTF-16 代码单元。
  // 
  // example:
  // 
  // [{"role":"user","content":[{"type":"input_text","text":"千问语音对话"}]}]
  Context []*VoiceCommandPayloadInputContext `json:"context,omitempty" xml:"context,omitempty" type:"Repeated"`
  // CLIENT 模式可选历史数组，省略或空数组为无历史；仅含此前完整 user/assistant 对，交替并以 assistant 结束，不含本轮问题。GATEWAY 模式不能提交非空历史。本轮附加图文放 data.messages，音频走二进制帧。
  // 
  // example:
  // 
  // [{"role":"user","content":"你好"},{"role":"assistant","content":"你好，有什么可以帮助你？"}]
  History []*VoiceCommandPayloadInputHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
}

func (s VoiceCommandPayloadInput) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadInput) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadInput) GetContext() []*VoiceCommandPayloadInputContext  {
  return s.Context
}

func (s *VoiceCommandPayloadInput) GetHistory() []*VoiceCommandPayloadInputHistory  {
  return s.History
}

func (s *VoiceCommandPayloadInput) SetContext(v []*VoiceCommandPayloadInputContext) *VoiceCommandPayloadInput {
  s.Context = v
  return s
}

func (s *VoiceCommandPayloadInput) SetHistory(v []*VoiceCommandPayloadInputHistory) *VoiceCommandPayloadInput {
  s.History = v
  return s
}

func (s *VoiceCommandPayloadInput) Validate() error {
  if s.Context != nil {
    for _, item := range s.Context {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.History != nil {
    for _, item := range s.History {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type VoiceCommandPayloadInputContext struct {
  // 上下文条目必填字符串，只接受 user 或 assistant，区分大小写；assistant 必须紧跟同一轮 user。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // user
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  // 上下文条目必填的非空数组，同一条消息的各文本块按顺序合并后用于识别。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // [{"type":"input_text","text":"千问语音对话"}]
  Content []*VoiceCommandPayloadInputContextContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s VoiceCommandPayloadInputContext) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadInputContext) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadInputContext) GetRole() *string  {
  return s.Role
}

func (s *VoiceCommandPayloadInputContext) GetContent() []*VoiceCommandPayloadInputContextContent  {
  return s.Content
}

func (s *VoiceCommandPayloadInputContext) SetRole(v string) *VoiceCommandPayloadInputContext {
  s.Role = &v
  return s
}

func (s *VoiceCommandPayloadInputContext) SetContent(v []*VoiceCommandPayloadInputContextContent) *VoiceCommandPayloadInputContext {
  s.Content = v
  return s
}

func (s *VoiceCommandPayloadInputContext) Validate() error {
  if s.Content != nil {
    for _, item := range s.Content {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type VoiceCommandPayloadInputContextContent struct {
  // 上下文内容块必填字符串；user 消息固定 input_text，assistant 消息固定 text。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // input_text
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 上下文内容块必填的非空文本，用于辅助识别；计入该条消息合计 400 个 UTF-16 代码单元限制。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 千问语音对话
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s VoiceCommandPayloadInputContextContent) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadInputContextContent) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadInputContextContent) GetType() *string  {
  return s.Type
}

func (s *VoiceCommandPayloadInputContextContent) GetText() *string  {
  return s.Text
}

func (s *VoiceCommandPayloadInputContextContent) SetType(v string) *VoiceCommandPayloadInputContextContent {
  s.Type = &v
  return s
}

func (s *VoiceCommandPayloadInputContextContent) SetText(v string) *VoiceCommandPayloadInputContextContent {
  s.Text = &v
  return s
}

func (s *VoiceCommandPayloadInputContextContent) Validate() error {
  return dara.Validate(s)
}

type VoiceCommandPayloadInputHistory struct {
  // 历史条目必填字符串，只接受 user 或 assistant，按 user、assistant 成对交替并以 assistant 结束。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // user
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  // 历史条目必填，为非空字符串或非空内容块数组。user 历史支持文本/图片，assistant 历史只接受文本；内容块格式沿用 Chat。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 你好
  Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
}

func (s VoiceCommandPayloadInputHistory) String() string {
  return dara.Prettify(s)
}

func (s VoiceCommandPayloadInputHistory) GoString() string {
  return s.String()
}

func (s *VoiceCommandPayloadInputHistory) GetRole() *string  {
  return s.Role
}

func (s *VoiceCommandPayloadInputHistory) GetContent() interface{}  {
  return s.Content
}

func (s *VoiceCommandPayloadInputHistory) SetRole(v string) *VoiceCommandPayloadInputHistory {
  s.Role = &v
  return s
}

func (s *VoiceCommandPayloadInputHistory) SetContent(v interface{}) *VoiceCommandPayloadInputHistory {
  s.Content = v
  return s
}

func (s *VoiceCommandPayloadInputHistory) Validate() error {
  return dara.Validate(s)
}

