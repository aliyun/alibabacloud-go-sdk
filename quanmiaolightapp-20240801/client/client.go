// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RunMarketingInformationExtractRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	ExtractType  *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId         *string   `json:"modelId,omitempty" xml:"modelId,omitempty"`
	SourceMaterials []*string `json:"sourceMaterials,omitempty" xml:"sourceMaterials,omitempty" type:"Repeated"`
}

func (s RunMarketingInformationExtractRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractRequest) SetCustomPrompt(v string) *RunMarketingInformationExtractRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetExtractType(v string) *RunMarketingInformationExtractRequest {
	s.ExtractType = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetModelId(v string) *RunMarketingInformationExtractRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetSourceMaterials(v []*string) *RunMarketingInformationExtractRequest {
	s.SourceMaterials = v
	return s
}

type RunMarketingInformationExtractShrinkRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	ExtractType  *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId               *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	SourceMaterialsShrink *string `json:"sourceMaterials,omitempty" xml:"sourceMaterials,omitempty"`
}

func (s RunMarketingInformationExtractShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractShrinkRequest) SetCustomPrompt(v string) *RunMarketingInformationExtractShrinkRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetExtractType(v string) *RunMarketingInformationExtractShrinkRequest {
	s.ExtractType = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetModelId(v string) *RunMarketingInformationExtractShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetSourceMaterialsShrink(v string) *RunMarketingInformationExtractShrinkRequest {
	s.SourceMaterialsShrink = &v
	return s
}

type RunMarketingInformationExtractResponseBody struct {
	// example:
	//
	// {\\"TimeZone\\": \\"Asia/Shanghai\\", \\"DateTime\\": \\"2024-03-07T17:00:09+08:00\\"}
	End     *bool                                              `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunMarketingInformationExtractResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunMarketingInformationExtractResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunMarketingInformationExtractResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBody) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBody) SetEnd(v bool) *RunMarketingInformationExtractResponseBody {
	s.End = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBody) SetHeader(v *RunMarketingInformationExtractResponseBodyHeader) *RunMarketingInformationExtractResponseBody {
	s.Header = v
	return s
}

func (s *RunMarketingInformationExtractResponseBody) SetPayload(v *RunMarketingInformationExtractResponseBodyPayload) *RunMarketingInformationExtractResponseBody {
	s.Payload = v
	return s
}

type RunMarketingInformationExtractResponseBodyHeader struct {
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F08C71C0-9399-548C-838B-1DA01DE211B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hojazg0u9lvytjc17ebgzzj2u4zukgh122tfg7wj1e6a1vcowy1ewzinauxriai9atcr6r323mm9ddbr0bg5m61ij8hxnf8664tstlfkfol6m8luc4shs3gums7l46uauyy0xndqmhdjtdon6coyhb4x17bo762bg9e3tb2geufg2
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 12826092918145
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunMarketingInformationExtractResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetEvent(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetEventInfo(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetRequestId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetSessionId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetTaskId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetTraceId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunMarketingInformationExtractResponseBodyPayload struct {
	Output *RunMarketingInformationExtractResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunMarketingInformationExtractResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunMarketingInformationExtractResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayload) SetOutput(v *RunMarketingInformationExtractResponseBodyPayloadOutput) *RunMarketingInformationExtractResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayload) SetUsage(v *RunMarketingInformationExtractResponseBodyPayloadUsage) *RunMarketingInformationExtractResponseBodyPayload {
	s.Usage = v
	return s
}

type RunMarketingInformationExtractResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunMarketingInformationExtractResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayloadOutput) SetText(v string) *RunMarketingInformationExtractResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunMarketingInformationExtractResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunMarketingInformationExtractResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetInputTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunMarketingInformationExtractResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunMarketingInformationExtractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMarketingInformationExtractResponse) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationExtractResponse) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponse) SetHeaders(v map[string]*string) *RunMarketingInformationExtractResponse {
	s.Headers = v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetStatusCode(v int32) *RunMarketingInformationExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetBody(v *RunMarketingInformationExtractResponseBody) *RunMarketingInformationExtractResponse {
	s.Body = v
	return s
}

type RunMarketingInformationWritingRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId        *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	SourceMaterial *string `json:"sourceMaterial,omitempty" xml:"sourceMaterial,omitempty"`
	WritingType    *string `json:"writingType,omitempty" xml:"writingType,omitempty"`
}

func (s RunMarketingInformationWritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingRequest) SetCustomPrompt(v string) *RunMarketingInformationWritingRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetModelId(v string) *RunMarketingInformationWritingRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetSourceMaterial(v string) *RunMarketingInformationWritingRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetWritingType(v string) *RunMarketingInformationWritingRequest {
	s.WritingType = &v
	return s
}

type RunMarketingInformationWritingResponseBody struct {
	// example:
	//
	// 2024-06-21T10:29:52+08:00
	End     *bool                                              `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunMarketingInformationWritingResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunMarketingInformationWritingResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunMarketingInformationWritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBody) SetEnd(v bool) *RunMarketingInformationWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBody) SetHeader(v *RunMarketingInformationWritingResponseBodyHeader) *RunMarketingInformationWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunMarketingInformationWritingResponseBody) SetPayload(v *RunMarketingInformationWritingResponseBodyPayload) *RunMarketingInformationWritingResponseBody {
	s.Payload = v
	return s
}

type RunMarketingInformationWritingResponseBodyHeader struct {
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 436BC5AE-0573-59D8-9803-6B5FDCD3BBA1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// uqubxgqzlnf4exfektij032lgb3yvix678p232n56387aqxdo4uutdt9wstqzovvz6j3ho7wnbgye785u79yn5q3euqmsvzmqdn3nmfq2826oscjvsi43kof8b8uxufpp1x97jjukk6jd3183hy8ni6hqpskuhuascpd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 13312125943232
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 213e20e517049392478441155e8b2a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunMarketingInformationWritingResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetEvent(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetEventInfo(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetRequestId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetSessionId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetTaskId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetTraceId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunMarketingInformationWritingResponseBodyPayload struct {
	Output *RunMarketingInformationWritingResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunMarketingInformationWritingResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunMarketingInformationWritingResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayload) SetOutput(v *RunMarketingInformationWritingResponseBodyPayloadOutput) *RunMarketingInformationWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayload) SetUsage(v *RunMarketingInformationWritingResponseBodyPayloadUsage) *RunMarketingInformationWritingResponseBodyPayload {
	s.Usage = v
	return s
}

type RunMarketingInformationWritingResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunMarketingInformationWritingResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) SetText(v string) *RunMarketingInformationWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunMarketingInformationWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunMarketingInformationWritingResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunMarketingInformationWritingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunMarketingInformationWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMarketingInformationWritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingResponse) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponse) SetHeaders(v map[string]*string) *RunMarketingInformationWritingResponse {
	s.Headers = v
	return s
}

func (s *RunMarketingInformationWritingResponse) SetStatusCode(v int32) *RunMarketingInformationWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMarketingInformationWritingResponse) SetBody(v *RunMarketingInformationWritingResponseBody) *RunMarketingInformationWritingResponse {
	s.Body = v
	return s
}

type RunScriptContinueRequest struct {
	// example:
	//
	// 一队全副武装的执法人员和消防员闯入了一间明显已被遗弃多日、门窗紧闭并用胶带封死的公寓，面对着屋内令人作呕的恶臭和门厅里的混乱场面，他们似乎在寻找某种隐藏的真相或危险源，而一封日期为16号的信件成为了揭开谜团的关键线索，随着便衣探员深入探索，一系列封闭的房间暗示着这里曾发生过不为人知的秘密事件。
	ScriptSummary *string `json:"scriptSummary,omitempty" xml:"scriptSummary,omitempty"`
	// example:
	//
	// 悬疑，都市，惊悚
	ScriptTypeKeyword *string `json:"scriptTypeKeyword,omitempty" xml:"scriptTypeKeyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 门厅一片狼藉。朝向天井的窗户开着。公寓门突然被撞开了。\n一名便衣探员、两名穿制服的警察和几位消防员———也身着工作服———进来，四下张望。他们都戴着手套以及盖住口鼻的面罩。在他们身后，门房和他妻子也挤进门厅。他们都捂着鼻子。门房的另一只手里拿着一叠信件和促销广告单。他们身后，跟着一位女邻居。\n便衣探员（对门房和邻居）：请在外面等候。\n他向一名警察示意，警察正忙着把好奇的旁观者请出门外。\n警察（对门房，指着那一叠信件）：最近的一封是哪天的？\n门房（查对信件）：最近的一封似乎是16号的......等一下......\n便衣探员想打开左侧的门，却是徒劳。门用胶带封上了。\n便衣探员（对消防员）：你来试一下好吗？\n消防员摆弄门的时候，便衣探员进了卧室隔壁的餐厅。他迅速打开窗，转身，想经过对开门进左侧的房间。这两扇门也锁着，门缝被贴上了胶带。他右转进入起居室，也打开了窗户
	UserProvidedContent *string `json:"userProvidedContent,omitempty" xml:"userProvidedContent,omitempty"`
}

func (s RunScriptContinueRequest) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueRequest) GoString() string {
	return s.String()
}

func (s *RunScriptContinueRequest) SetScriptSummary(v string) *RunScriptContinueRequest {
	s.ScriptSummary = &v
	return s
}

func (s *RunScriptContinueRequest) SetScriptTypeKeyword(v string) *RunScriptContinueRequest {
	s.ScriptTypeKeyword = &v
	return s
}

func (s *RunScriptContinueRequest) SetUserProvidedContent(v string) *RunScriptContinueRequest {
	s.UserProvidedContent = &v
	return s
}

type RunScriptContinueResponseBody struct {
	End     *bool                                 `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptContinueResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptContinueResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptContinueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBody) SetEnd(v bool) *RunScriptContinueResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptContinueResponseBody) SetHeader(v *RunScriptContinueResponseBodyHeader) *RunScriptContinueResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptContinueResponseBody) SetPayload(v *RunScriptContinueResponseBodyPayload) *RunScriptContinueResponseBody {
	s.Payload = v
	return s
}

type RunScriptContinueResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 0EB27AE3-CA53-5FAE-83C6-EE66CA4DF5DF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptContinueResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyHeader) SetErrorCode(v string) *RunScriptContinueResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetErrorMessage(v string) *RunScriptContinueResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetEvent(v string) *RunScriptContinueResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetEventInfo(v string) *RunScriptContinueResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetRequestId(v string) *RunScriptContinueResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetSessionId(v string) *RunScriptContinueResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetTaskId(v string) *RunScriptContinueResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetTraceId(v string) *RunScriptContinueResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunScriptContinueResponseBodyPayload struct {
	Output *RunScriptContinueResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptContinueResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptContinueResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayload) SetOutput(v *RunScriptContinueResponseBodyPayloadOutput) *RunScriptContinueResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptContinueResponseBodyPayload) SetUsage(v *RunScriptContinueResponseBodyPayloadUsage) *RunScriptContinueResponseBodyPayload {
	s.Usage = v
	return s
}

type RunScriptContinueResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptContinueResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayloadOutput) SetText(v string) *RunScriptContinueResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunScriptContinueResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunScriptContinueResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunScriptContinueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptContinueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptContinueResponse) String() string {
	return tea.Prettify(s)
}

func (s RunScriptContinueResponse) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponse) SetHeaders(v map[string]*string) *RunScriptContinueResponse {
	s.Headers = v
	return s
}

func (s *RunScriptContinueResponse) SetStatusCode(v int32) *RunScriptContinueResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptContinueResponse) SetBody(v *RunScriptContinueResponseBody) *RunScriptContinueResponse {
	s.Body = v
	return s
}

type RunScriptPlanningRequest struct {
	// example:
	//
	// 故事尽可能狗血
	AdditionalNote  *string `json:"additionalNote,omitempty" xml:"additionalNote,omitempty"`
	DialogueInScene *bool   `json:"dialogueInScene,omitempty" xml:"dialogueInScene,omitempty"`
	PlotConflict    *bool   `json:"plotConflict,omitempty" xml:"plotConflict,omitempty"`
	// example:
	//
	// 都市战神
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// example:
	//
	// 3
	ScriptShotCount *int32 `json:"scriptShotCount,omitempty" xml:"scriptShotCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 在一个宁静的小镇上，每个家庭都在同一天收到一个神秘的、没有标记的包裹。
	ScriptSummary *string `json:"scriptSummary,omitempty" xml:"scriptSummary,omitempty"`
	// example:
	//
	// 现代，都市，爱情，玄幻
	ScriptTypeKeyword *string `json:"scriptTypeKeyword,omitempty" xml:"scriptTypeKeyword,omitempty"`
}

func (s RunScriptPlanningRequest) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningRequest) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningRequest) SetAdditionalNote(v string) *RunScriptPlanningRequest {
	s.AdditionalNote = &v
	return s
}

func (s *RunScriptPlanningRequest) SetDialogueInScene(v bool) *RunScriptPlanningRequest {
	s.DialogueInScene = &v
	return s
}

func (s *RunScriptPlanningRequest) SetPlotConflict(v bool) *RunScriptPlanningRequest {
	s.PlotConflict = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptName(v string) *RunScriptPlanningRequest {
	s.ScriptName = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptShotCount(v int32) *RunScriptPlanningRequest {
	s.ScriptShotCount = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptSummary(v string) *RunScriptPlanningRequest {
	s.ScriptSummary = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptTypeKeyword(v string) *RunScriptPlanningRequest {
	s.ScriptTypeKeyword = &v
	return s
}

type RunScriptPlanningResponseBody struct {
	End     *bool                                 `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptPlanningResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptPlanningResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptPlanningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBody) SetEnd(v bool) *RunScriptPlanningResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptPlanningResponseBody) SetHeader(v *RunScriptPlanningResponseBodyHeader) *RunScriptPlanningResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptPlanningResponseBody) SetPayload(v *RunScriptPlanningResponseBodyPayload) *RunScriptPlanningResponseBody {
	s.Payload = v
	return s
}

type RunScriptPlanningResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 0EB27AE3-CA53-5FAE-83C6-EE66CA4DF5DF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptPlanningResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyHeader) SetErrorCode(v string) *RunScriptPlanningResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetErrorMessage(v string) *RunScriptPlanningResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetEvent(v string) *RunScriptPlanningResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetEventInfo(v string) *RunScriptPlanningResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetRequestId(v string) *RunScriptPlanningResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetSessionId(v string) *RunScriptPlanningResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetTaskId(v string) *RunScriptPlanningResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetTraceId(v string) *RunScriptPlanningResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunScriptPlanningResponseBodyPayload struct {
	Output *RunScriptPlanningResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptPlanningResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptPlanningResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayload) SetOutput(v *RunScriptPlanningResponseBodyPayloadOutput) *RunScriptPlanningResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptPlanningResponseBodyPayload) SetUsage(v *RunScriptPlanningResponseBodyPayloadUsage) *RunScriptPlanningResponseBodyPayload {
	s.Usage = v
	return s
}

type RunScriptPlanningResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptPlanningResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayloadOutput) SetText(v string) *RunScriptPlanningResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunScriptPlanningResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunScriptPlanningResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunScriptPlanningResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptPlanningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptPlanningResponse) String() string {
	return tea.Prettify(s)
}

func (s RunScriptPlanningResponse) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponse) SetHeaders(v map[string]*string) *RunScriptPlanningResponse {
	s.Headers = v
	return s
}

func (s *RunScriptPlanningResponse) SetStatusCode(v int32) *RunScriptPlanningResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptPlanningResponse) SetBody(v *RunScriptPlanningResponseBody) *RunScriptPlanningResponse {
	s.Body = v
	return s
}

type RunStyleWritingRequest struct {
	// This parameter is required.
	LearningSamples []*string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty" type:"Repeated"`
	// This parameter is required.
	ReferenceMaterials []*string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty" type:"Repeated"`
	StyleFeature       *string   `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	// This parameter is required.
	WritingTheme *string `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
}

func (s RunStyleWritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingRequest) GoString() string {
	return s.String()
}

func (s *RunStyleWritingRequest) SetLearningSamples(v []*string) *RunStyleWritingRequest {
	s.LearningSamples = v
	return s
}

func (s *RunStyleWritingRequest) SetReferenceMaterials(v []*string) *RunStyleWritingRequest {
	s.ReferenceMaterials = v
	return s
}

func (s *RunStyleWritingRequest) SetStyleFeature(v string) *RunStyleWritingRequest {
	s.StyleFeature = &v
	return s
}

func (s *RunStyleWritingRequest) SetWritingTheme(v string) *RunStyleWritingRequest {
	s.WritingTheme = &v
	return s
}

type RunStyleWritingShrinkRequest struct {
	// This parameter is required.
	LearningSamplesShrink *string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty"`
	// This parameter is required.
	ReferenceMaterialsShrink *string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty"`
	StyleFeature             *string `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	// This parameter is required.
	WritingTheme *string `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
}

func (s RunStyleWritingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStyleWritingShrinkRequest) SetLearningSamplesShrink(v string) *RunStyleWritingShrinkRequest {
	s.LearningSamplesShrink = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetReferenceMaterialsShrink(v string) *RunStyleWritingShrinkRequest {
	s.ReferenceMaterialsShrink = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetStyleFeature(v string) *RunStyleWritingShrinkRequest {
	s.StyleFeature = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetWritingTheme(v string) *RunStyleWritingShrinkRequest {
	s.WritingTheme = &v
	return s
}

type RunStyleWritingResponseBody struct {
	// example:
	//
	// true
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// {"event":"task-progress-start-generating","sessionId":"3cd10828-0e42-471c-8f1a-931cde20b035","taskId":"d3be9981-ca2d-4e17-bf31-1c0a628e9f99","traceId":"66bef4a7f5d61ff3c43f3b710574e175"}
	Header  *RunStyleWritingResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunStyleWritingResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunStyleWritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBody) SetEnd(v bool) *RunStyleWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunStyleWritingResponseBody) SetHeader(v *RunStyleWritingResponseBodyHeader) *RunStyleWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunStyleWritingResponseBody) SetPayload(v *RunStyleWritingResponseBodyPayload) *RunStyleWritingResponseBody {
	s.Payload = v
	return s
}

type RunStyleWritingResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-progress-start-generating
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 0EB27AE3-CA53-5FAE-83C6-EE66CA4DF5DF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunStyleWritingResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyHeader) SetErrorCode(v string) *RunStyleWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetErrorMessage(v string) *RunStyleWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetEvent(v string) *RunStyleWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetEventInfo(v string) *RunStyleWritingResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetRequestId(v string) *RunStyleWritingResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetSessionId(v string) *RunStyleWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetTaskId(v string) *RunStyleWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetTraceId(v string) *RunStyleWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunStyleWritingResponseBodyPayload struct {
	Output *RunStyleWritingResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// {
	//
	//         "inputTokens": 1816,
	//
	//         "outputTokens": 96,
	//
	//         "totalTokens": 1912
	//
	//     }
	Usage *RunStyleWritingResponseBodyPayloadUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunStyleWritingResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayload) SetOutput(v *RunStyleWritingResponseBodyPayloadOutput) *RunStyleWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStyleWritingResponseBodyPayload) SetUsage(v *RunStyleWritingResponseBodyPayloadUsage) *RunStyleWritingResponseBodyPayload {
	s.Usage = v
	return s
}

type RunStyleWritingResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunStyleWritingResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayloadOutput) SetText(v string) *RunStyleWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunStyleWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunStyleWritingResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunStyleWritingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunStyleWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStyleWritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunStyleWritingResponse) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponse) SetHeaders(v map[string]*string) *RunStyleWritingResponse {
	s.Headers = v
	return s
}

func (s *RunStyleWritingResponse) SetStatusCode(v int32) *RunStyleWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStyleWritingResponse) SetBody(v *RunStyleWritingResponseBody) *RunStyleWritingResponse {
	s.Body = v
	return s
}

type RunVideoAnalysisRequest struct {
	GenerateOptions           []*string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	ModelCustomPromptTemplate *string   `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// PlotDetail
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId *string `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s RunVideoAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequest) SetGenerateOptions(v []*string) *RunVideoAnalysisRequest {
	s.GenerateOptions = v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelId(v string) *RunVideoAnalysisRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetOriginalSessionId(v string) *RunVideoAnalysisRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetTaskId(v string) *RunVideoAnalysisRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoModelId(v string) *RunVideoAnalysisRequest {
	s.VideoModelId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoUrl(v string) *RunVideoAnalysisRequest {
	s.VideoUrl = &v
	return s
}

type RunVideoAnalysisShrinkRequest struct {
	GenerateOptionsShrink     *string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// PlotDetail
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId *string `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s RunVideoAnalysisShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisShrinkRequest) SetGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelId(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetOriginalSessionId(v string) *RunVideoAnalysisShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetTaskId(v string) *RunVideoAnalysisShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoModelId(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoModelId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoUrl(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoUrl = &v
	return s
}

type RunVideoAnalysisResponseBody struct {
	Header  *RunVideoAnalysisResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunVideoAnalysisResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunVideoAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBody) SetHeader(v *RunVideoAnalysisResponseBodyHeader) *RunVideoAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunVideoAnalysisResponseBody) SetPayload(v *RunVideoAnalysisResponseBodyPayload) *RunVideoAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunVideoAnalysisResponseBody) SetRequestId(v string) *RunVideoAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type RunVideoAnalysisResponseBodyHeader struct {
	// example:
	//
	// InvalidParam
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-progress-start-generating
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// xxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunVideoAnalysisResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyHeader) SetErrorCode(v string) *RunVideoAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunVideoAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetEvent(v string) *RunVideoAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetEventInfo(v string) *RunVideoAnalysisResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetSessionId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetTaskId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetTraceId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunVideoAnalysisResponseBodyPayload struct {
	Output *RunVideoAnalysisResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayload) SetOutput(v *RunVideoAnalysisResponseBodyPayloadOutput) *RunVideoAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutput struct {
	VideoAnalysisResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult            `json:"videoAnalysisResult,omitempty" xml:"videoAnalysisResult,omitempty" type:"Struct"`
	VideoCaptionResult             *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult             `json:"videoCaptionResult,omitempty" xml:"videoCaptionResult,omitempty" type:"Struct"`
	VideoGenerateResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult            `json:"videoGenerateResult,omitempty" xml:"videoGenerateResult,omitempty" type:"Struct"`
	VideoMindMappingGenerateResult *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult `json:"videoMindMappingGenerateResult,omitempty" xml:"videoMindMappingGenerateResult,omitempty" type:"Struct"`
	VideoTitleGenerateResult       *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult       `json:"videoTitleGenerateResult,omitempty" xml:"videoTitleGenerateResult,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoAnalysisResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoAnalysisResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoCaptionResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoCaptionResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoMindMappingGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoMindMappingGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoTitleGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoTitleGenerateResult = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult struct {
	// example:
	//
	// true
	GenerateFinished         *bool                                                                                   `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text                     *string                                                                                 `json:"text,omitempty" xml:"text,omitempty"`
	Usage                    *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage                      `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoShotAnalysisResults []*RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults `json:"videoShotAnalysisResults,omitempty" xml:"videoShotAnalysisResults,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetVideoShotAnalysisResults(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.VideoShotAnalysisResults = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.TotalTokens = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults struct {
	// example:
	//
	// 10000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1000
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.Text = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                                       `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	VideoCaptions    []*RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) SetVideoCaptions(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult {
	s.VideoCaptions = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:01
	EndTimeFormat *string `json:"endTimeFormat,omitempty" xml:"endTimeFormat,omitempty"`
	// example:
	//
	// 0
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 00:01
	StartTimeFormat *string `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	// example:
	//
	// xxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetEndTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetStartTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.Text = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                              `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text             *string                                                            `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Usage = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished  *bool                                                                                       `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text              *string                                                                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage             *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage               `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoMindMappings []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings `json:"videoMindMappings,omitempty" xml:"videoMindMappings,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetVideoMindMappings(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.VideoMindMappings = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings struct {
	ChildNodes []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                               `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetChildNodes(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.ChildNodes = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.Name = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes struct {
	ChildNodes []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                         `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetChildNodes(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.ChildNodes = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.Name = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	s.Name = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                                   `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text             *string                                                                 `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.Usage = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type RunVideoAnalysisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunVideoAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunVideoAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponse) SetHeaders(v map[string]*string) *RunVideoAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunVideoAnalysisResponse) SetStatusCode(v int32) *RunVideoAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunVideoAnalysisResponse) SetBody(v *RunVideoAnalysisResponseBody) *RunVideoAnalysisResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("quanmiaolightapp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营销信息抽取服务
//
// @param tmpReq - RunMarketingInformationExtractRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationExtractResponse
func (client *Client) RunMarketingInformationExtractWithOptions(workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunMarketingInformationExtractResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunMarketingInformationExtractShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SourceMaterials)) {
		request.SourceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMaterials, tea.String("sourceMaterials"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPrompt)) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.ExtractType)) {
		body["extractType"] = request.ExtractType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceMaterialsShrink)) {
		body["sourceMaterials"] = request.SourceMaterialsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunMarketingInformationExtract"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationExtract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunMarketingInformationExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销信息抽取服务
//
// @param request - RunMarketingInformationExtractRequest
//
// @return RunMarketingInformationExtractResponse
func (client *Client) RunMarketingInformationExtract(workspaceId *string, request *RunMarketingInformationExtractRequest) (_result *RunMarketingInformationExtractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunMarketingInformationExtractResponse{}
	_body, _err := client.RunMarketingInformationExtractWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营销文案写作服务
//
// @param request - RunMarketingInformationWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationWritingResponse
func (client *Client) RunMarketingInformationWritingWithOptions(workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunMarketingInformationWritingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPrompt)) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceMaterial)) {
		body["sourceMaterial"] = request.SourceMaterial
	}

	if !tea.BoolValue(util.IsUnset(request.WritingType)) {
		body["writingType"] = request.WritingType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunMarketingInformationWriting"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationWriting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunMarketingInformationWritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销文案写作服务
//
// @param request - RunMarketingInformationWritingRequest
//
// @return RunMarketingInformationWritingResponse
func (client *Client) RunMarketingInformationWriting(workspaceId *string, request *RunMarketingInformationWritingRequest) (_result *RunMarketingInformationWritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunMarketingInformationWritingResponse{}
	_body, _err := client.RunMarketingInformationWritingWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 剧本续写
//
// @param request - RunScriptContinueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptContinueResponse
func (client *Client) RunScriptContinueWithOptions(workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunScriptContinueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScriptSummary)) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptTypeKeyword)) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvidedContent)) {
		body["userProvidedContent"] = request.UserProvidedContent
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunScriptContinue"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runScriptContinue"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunScriptContinueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剧本续写
//
// @param request - RunScriptContinueRequest
//
// @return RunScriptContinueResponse
func (client *Client) RunScriptContinue(workspaceId *string, request *RunScriptContinueRequest) (_result *RunScriptContinueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptContinueResponse{}
	_body, _err := client.RunScriptContinueWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 剧本策划
//
// @param request - RunScriptPlanningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptPlanningResponse
func (client *Client) RunScriptPlanningWithOptions(workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunScriptPlanningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalNote)) {
		body["additionalNote"] = request.AdditionalNote
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueInScene)) {
		body["dialogueInScene"] = request.DialogueInScene
	}

	if !tea.BoolValue(util.IsUnset(request.PlotConflict)) {
		body["plotConflict"] = request.PlotConflict
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptName)) {
		body["scriptName"] = request.ScriptName
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptShotCount)) {
		body["scriptShotCount"] = request.ScriptShotCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptSummary)) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptTypeKeyword)) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunScriptPlanning"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runScriptPlanning"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunScriptPlanningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剧本策划
//
// @param request - RunScriptPlanningRequest
//
// @return RunScriptPlanningResponse
func (client *Client) RunScriptPlanning(workspaceId *string, request *RunScriptPlanningRequest) (_result *RunScriptPlanningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptPlanningResponse{}
	_body, _err := client.RunScriptPlanningWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文体学习和写作推理服务
//
// @param tmpReq - RunStyleWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleWritingResponse
func (client *Client) RunStyleWritingWithOptions(workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunStyleWritingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunStyleWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LearningSamples)) {
		request.LearningSamplesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LearningSamples, tea.String("learningSamples"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceMaterials)) {
		request.ReferenceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceMaterials, tea.String("referenceMaterials"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LearningSamplesShrink)) {
		body["learningSamples"] = request.LearningSamplesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceMaterialsShrink)) {
		body["referenceMaterials"] = request.ReferenceMaterialsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StyleFeature)) {
		body["styleFeature"] = request.StyleFeature
	}

	if !tea.BoolValue(util.IsUnset(request.WritingTheme)) {
		body["writingTheme"] = request.WritingTheme
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunStyleWriting"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runStyleWriting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunStyleWritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文体学习和写作推理服务
//
// @param request - RunStyleWritingRequest
//
// @return RunStyleWritingResponse
func (client *Client) RunStyleWriting(workspaceId *string, request *RunStyleWritingRequest) (_result *RunStyleWritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunStyleWritingResponse{}
	_body, _err := client.RunStyleWritingWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-视频理解
//
// @param tmpReq - RunVideoAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoAnalysisResponse
func (client *Client) RunVideoAnalysisWithOptions(workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunVideoAnalysisResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunVideoAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GenerateOptions)) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, tea.String("generateOptions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateOptionsShrink)) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCustomPromptTemplate)) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCustomPromptTemplateId)) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalSessionId)) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelCustomPromptTemplate)) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelId)) {
		body["videoModelId"] = request.VideoModelId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["videoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunVideoAnalysis"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runVideoAnalysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunVideoAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-视频理解
//
// @param request - RunVideoAnalysisRequest
//
// @return RunVideoAnalysisResponse
func (client *Client) RunVideoAnalysis(workspaceId *string, request *RunVideoAnalysisRequest) (_result *RunVideoAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunVideoAnalysisResponse{}
	_body, _err := client.RunVideoAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
