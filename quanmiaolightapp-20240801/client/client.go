// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GenerateBroadcastNewsRequest struct {
	// This parameter is required.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
}

func (s GenerateBroadcastNewsRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsRequest) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsRequest) SetPrompt(v string) *GenerateBroadcastNewsRequest {
	s.Prompt = &v
	return s
}

type GenerateBroadcastNewsResponseBody struct {
	// example:
	//
	// xx
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *GenerateBroadcastNewsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateBroadcastNewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBody) SetCode(v string) *GenerateBroadcastNewsResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetData(v *GenerateBroadcastNewsResponseBodyData) *GenerateBroadcastNewsResponseBody {
	s.Data = v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetHttpStatusCode(v int32) *GenerateBroadcastNewsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetMessage(v string) *GenerateBroadcastNewsResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetRequestId(v string) *GenerateBroadcastNewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetSuccess(v bool) *GenerateBroadcastNewsResponseBody {
	s.Success = &v
	return s
}

type GenerateBroadcastNewsResponseBodyData struct {
	HotTopicSummaries []*GenerateBroadcastNewsResponseBodyDataHotTopicSummaries `json:"hotTopicSummaries,omitempty" xml:"hotTopicSummaries,omitempty" type:"Repeated"`
	// example:
	//
	// 2bb0ea82dafd48a8817fadc4c90e2b52
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string                                     `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Text   *string                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage  *GenerateBroadcastNewsResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GenerateBroadcastNewsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyData) SetHotTopicSummaries(v []*GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) *GenerateBroadcastNewsResponseBodyData {
	s.HotTopicSummaries = v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetSessionId(v string) *GenerateBroadcastNewsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetTaskId(v string) *GenerateBroadcastNewsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetText(v string) *GenerateBroadcastNewsResponseBodyData {
	s.Text = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetUsage(v *GenerateBroadcastNewsResponseBodyDataUsage) *GenerateBroadcastNewsResponseBodyData {
	s.Usage = v
	return s
}

type GenerateBroadcastNewsResponseBodyDataHotTopicSummaries struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_08
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 1000000
	HotValue *float64 `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	// example:
	//
	// 1458tb3bjo7531kap42a
	Id     *string                                                         `json:"id,omitempty" xml:"id,omitempty"`
	Images []*GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetCategory(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Category = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotTopic(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotTopic = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotTopicVersion(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotTopicVersion = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotValue(v float64) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotValue = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetId(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Id = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetImages(v []*GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Images = v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetTextSummary(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.TextSummary = &v
	return s
}

type GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages struct {
	// example:
	//
	// http://xxx.com/xxx.jpeg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) SetUrl(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages {
	s.Url = &v
	return s
}

type GenerateBroadcastNewsResponseBodyDataUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 3
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataUsage) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetInputTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetOutputTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetTotalTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

type GenerateBroadcastNewsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateBroadcastNewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateBroadcastNewsResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateBroadcastNewsResponse) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponse) SetHeaders(v map[string]*string) *GenerateBroadcastNewsResponse {
	s.Headers = v
	return s
}

func (s *GenerateBroadcastNewsResponse) SetStatusCode(v int32) *GenerateBroadcastNewsResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateBroadcastNewsResponse) SetBody(v *GenerateBroadcastNewsResponseBody) *GenerateBroadcastNewsResponse {
	s.Body = v
	return s
}

type GenerateOutputFormatRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// This parameter is required.
	Tags []*GenerateOutputFormatRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s GenerateOutputFormatRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatRequest) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatRequest) SetBusinessType(v string) *GenerateOutputFormatRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetContent(v string) *GenerateOutputFormatRequest {
	s.Content = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetExtraInfo(v string) *GenerateOutputFormatRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetTags(v []*GenerateOutputFormatRequestTags) *GenerateOutputFormatRequest {
	s.Tags = v
	return s
}

func (s *GenerateOutputFormatRequest) SetTaskDescription(v string) *GenerateOutputFormatRequest {
	s.TaskDescription = &v
	return s
}

type GenerateOutputFormatRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s GenerateOutputFormatRequestTags) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatRequestTags) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatRequestTags) SetTagDefinePrompt(v string) *GenerateOutputFormatRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *GenerateOutputFormatRequestTags) SetTagName(v string) *GenerateOutputFormatRequestTags {
	s.TagName = &v
	return s
}

type GenerateOutputFormatShrinkRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// This parameter is required.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s GenerateOutputFormatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatShrinkRequest) SetBusinessType(v string) *GenerateOutputFormatShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetContent(v string) *GenerateOutputFormatShrinkRequest {
	s.Content = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetExtraInfo(v string) *GenerateOutputFormatShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetTagsShrink(v string) *GenerateOutputFormatShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetTaskDescription(v string) *GenerateOutputFormatShrinkRequest {
	s.TaskDescription = &v
	return s
}

type GenerateOutputFormatResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GenerateOutputFormatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateOutputFormatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponseBody) SetCode(v string) *GenerateOutputFormatResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetData(v *GenerateOutputFormatResponseBodyData) *GenerateOutputFormatResponseBody {
	s.Data = v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetHttpStatusCode(v int32) *GenerateOutputFormatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetMessage(v string) *GenerateOutputFormatResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetRequestId(v string) *GenerateOutputFormatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetSuccess(v bool) *GenerateOutputFormatResponseBody {
	s.Success = &v
	return s
}

type GenerateOutputFormatResponseBodyData struct {
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
}

func (s GenerateOutputFormatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponseBodyData) SetOutputFormat(v string) *GenerateOutputFormatResponseBodyData {
	s.OutputFormat = &v
	return s
}

type GenerateOutputFormatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOutputFormatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOutputFormatResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateOutputFormatResponse) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponse) SetHeaders(v map[string]*string) *GenerateOutputFormatResponse {
	s.Headers = v
	return s
}

func (s *GenerateOutputFormatResponse) SetStatusCode(v int32) *GenerateOutputFormatResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOutputFormatResponse) SetBody(v *GenerateOutputFormatResponseBody) *GenerateOutputFormatResponse {
	s.Body = v
	return s
}

type GetTagMiningAnalysisTaskRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTagMiningAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskRequest) SetTaskId(v string) *GetTagMiningAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

type GetTagMiningAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTagMiningAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// DataNotExists
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetCode(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetData(v *GetTagMiningAnalysisTaskResponseBodyData) *GetTagMiningAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetHttpStatusCode(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetMessage(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetRequestId(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetSuccess(v bool) *GetTagMiningAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyData struct {
	ErrorCode    *string                                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Results      []*GetTagMiningAnalysisTaskResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNIN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetErrorCode(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetResults(v []*GetTagMiningAnalysisTaskResponseBodyDataResults) *GetTagMiningAnalysisTaskResponseBodyData {
	s.Results = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetStatus(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyDataResults struct {
	// example:
	//
	// 1
	CustomId *string                                                 `json:"customId,omitempty" xml:"customId,omitempty"`
	Header   *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload  *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetCustomId(v string) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.CustomId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetHeader(v *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.Header = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetPayload(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.Payload = v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsHeader struct {
	// example:
	//
	// DataNotExists
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetErrorCode(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.ErrorCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetErrorMessage(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.ErrorMessage = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetEvent(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.Event = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetRequestId(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.RequestId = &v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayload struct {
	Output *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) SetOutput(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload {
	s.Output = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) SetUsage(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload {
	s.Usage = v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput struct {
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) SetText(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput {
	s.Text = &v
	return s
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage struct {
	// example:
	//
	// 100
	InputToken *int64 `json:"inputToken,omitempty" xml:"inputToken,omitempty"`
	// example:
	//
	// 200
	OutputToken *int64 `json:"outputToken,omitempty" xml:"outputToken,omitempty"`
	// example:
	//
	// 300
	TotalToken *int64 `json:"totalToken,omitempty" xml:"totalToken,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetInputToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.InputToken = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetOutputToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.OutputToken = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetTotalToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.TotalToken = &v
	return s
}

type GetTagMiningAnalysisTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagMiningAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetTagMiningAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponse) SetStatusCode(v int32) *GetTagMiningAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponse) SetBody(v *GetTagMiningAnalysisTaskResponseBody) *GetTagMiningAnalysisTaskResponse {
	s.Body = v
	return s
}

type GetVideoAnalysisConfigResponseBody struct {
	// example:
	//
	// xx
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoAnalysisConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoAnalysisConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponseBody) SetCode(v string) *GetVideoAnalysisConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetData(v *GetVideoAnalysisConfigResponseBodyData) *GetVideoAnalysisConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetHttpStatusCode(v int32) *GetVideoAnalysisConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetMessage(v string) *GetVideoAnalysisConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetRequestId(v string) *GetVideoAnalysisConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetSuccess(v bool) *GetVideoAnalysisConfigResponseBody {
	s.Success = &v
	return s
}

type GetVideoAnalysisConfigResponseBodyData struct {
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s GetVideoAnalysisConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponseBodyData) SetAsyncConcurrency(v int32) *GetVideoAnalysisConfigResponseBodyData {
	s.AsyncConcurrency = &v
	return s
}

type GetVideoAnalysisConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoAnalysisConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoAnalysisConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponse) SetHeaders(v map[string]*string) *GetVideoAnalysisConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVideoAnalysisConfigResponse) SetStatusCode(v int32) *GetVideoAnalysisConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoAnalysisConfigResponse) SetBody(v *GetVideoAnalysisConfigResponseBody) *GetVideoAnalysisConfigResponse {
	s.Body = v
	return s
}

type GetVideoAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVideoAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskRequest) SetTaskId(v string) *GetVideoAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

type GetVideoAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBody) SetCode(v string) *GetVideoAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetData(v *GetVideoAnalysisTaskResponseBodyData) *GetVideoAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetVideoAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetMessage(v string) *GetVideoAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetRequestId(v string) *GetVideoAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetSuccess(v bool) *GetVideoAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/mjp-test-default.
	ErrorMessage *string                                      `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Header       *GetVideoAnalysisTaskResponseBodyDataHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload      *GetVideoAnalysisTaskResponseBodyDataPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId      *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskRunInfo *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo `json:"taskRunInfo,omitempty" xml:"taskRunInfo,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetHeader(v *GetVideoAnalysisTaskResponseBodyDataHeader) *GetVideoAnalysisTaskResponseBodyData {
	s.Header = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetPayload(v *GetVideoAnalysisTaskResponseBodyDataPayload) *GetVideoAnalysisTaskResponseBodyData {
	s.Payload = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskId(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskRunInfo(v *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskRunInfo = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskStatus(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataHeader struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Deduct task already success,Please do not resubmit.token \\"369e8f2c-d283-424a-96c4-c83efe08c89e\\"
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// TIMEOUT_CLOSE_ORDER
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxx
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// d5c38cf6-a4bf-4a57-a697-9f449926f0c9
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 6e223291-729b-4e84-9271-c13ada1a776b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 215045f817272303448235204efdef
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataHeader) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataHeader) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetErrorCode(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetErrorMessage(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetEvent(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.Event = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetEventInfo(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.EventInfo = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetSessionId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.SessionId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetTaskId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.TaskId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetTraceId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.TraceId = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayload struct {
	Output *GetVideoAnalysisTaskResponseBodyDataPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *GetVideoAnalysisTaskResponseBodyDataPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayload) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayload) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) SetOutput(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) *GetVideoAnalysisTaskResponseBodyDataPayload {
	s.Output = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) *GetVideoAnalysisTaskResponseBodyDataPayload {
	s.Usage = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutput struct {
	ResultJsonFileUrl              *string                                                                          `json:"resultJsonFileUrl,omitempty" xml:"resultJsonFileUrl,omitempty"`
	VideoAnalysisResult            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult            `json:"videoAnalysisResult,omitempty" xml:"videoAnalysisResult,omitempty" type:"Struct"`
	VideoCaptionResult             *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult             `json:"videoCaptionResult,omitempty" xml:"videoCaptionResult,omitempty" type:"Struct"`
	VideoGenerateResult            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult            `json:"videoGenerateResult,omitempty" xml:"videoGenerateResult,omitempty" type:"Struct"`
	VideoGenerateResults           []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults         `json:"videoGenerateResults,omitempty" xml:"videoGenerateResults,omitempty" type:"Repeated"`
	VideoMindMappingGenerateResult *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult `json:"videoMindMappingGenerateResult,omitempty" xml:"videoMindMappingGenerateResult,omitempty" type:"Struct"`
	VideoTitleGenerateResult       *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult       `json:"videoTitleGenerateResult,omitempty" xml:"videoTitleGenerateResult,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetResultJsonFileUrl(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.ResultJsonFileUrl = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoAnalysisResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoAnalysisResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoCaptionResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoCaptionResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoGenerateResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoGenerateResults(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoGenerateResults = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoMindMappingGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoMindMappingGenerateResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoTitleGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoTitleGenerateResult = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// xxx
	Text                     *string                                                                                         `json:"text,omitempty" xml:"text,omitempty"`
	Usage                    *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage                      `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoShotAnalysisResults []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults `json:"videoShotAnalysisResults,omitempty" xml:"videoShotAnalysisResults,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetVideoShotAnalysisResults(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.VideoShotAnalysisResults = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage struct {
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2024-10-05 06:22:00
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetEndTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.EndTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetStartTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.StartTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.Text = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	VideoCaptions    []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) SetVideoCaptions(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult {
	s.VideoCaptions = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions struct {
	// example:
	//
	// 1736129678000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:01
	EndTimeFormat *string `json:"endTimeFormat,omitempty" xml:"endTimeFormat,omitempty"`
	// example:
	//
	// 00:01
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2025-01-07 11:52:06
	StartTimeFormat *string `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetEndTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetEndTimeFormat(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTimeFormat = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetStartTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetStartTimeFormat(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTimeFormat = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.Text = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool   `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32  `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce      *bool   `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	ReasonText       *string `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	// example:
	//
	// xxx
	Text  *string                                                                    `json:"text,omitempty" xml:"text,omitempty"`
	Usage *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetIndex(v int32) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Index = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetModelId(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ModelId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetModelReduce(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetReasonText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ReasonText = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Usage = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage struct {
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

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults struct {
	GenerateFinished *bool                                                                       `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32                                                                      `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string                                                                     `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ReasonText       *string                                                                     `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text             *string                                                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetIndex(v int32) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Index = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetModelId(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.ModelId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetReasonText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.ReasonText = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Usage = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished  *bool                                                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text              *string                                                                                             `json:"text,omitempty" xml:"text,omitempty"`
	Usage             *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage               `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoMindMappings []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings `json:"videoMindMappings,omitempty" xml:"videoMindMappings,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetVideoMindMappings(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.VideoMindMappings = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage struct {
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

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings struct {
	ChildNodes []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetChildNodes(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.ChildNodes = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.Name = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes struct {
	ChildNodes []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                                 `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetChildNodes(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.ChildNodes = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.Name = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	s.Name = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// xxxx
	Text  *string                                                                         `json:"text,omitempty" xml:"text,omitempty"`
	Usage *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.Usage = v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage struct {
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataPayloadUsage struct {
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.TotalTokens = &v
	return s
}

type GetVideoAnalysisTaskResponseBodyDataTaskRunInfo struct {
	// example:
	//
	// true
	ConcurrentChargeEnable *bool `json:"concurrentChargeEnable,omitempty" xml:"concurrentChargeEnable,omitempty"`
	// example:
	//
	// 1
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) SetConcurrentChargeEnable(v bool) *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo {
	s.ConcurrentChargeEnable = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) SetResponseTime(v int64) *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo {
	s.ResponseTime = &v
	return s
}

type GetVideoAnalysisTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetVideoAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoAnalysisTaskResponse) SetStatusCode(v int32) *GetVideoAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponse) SetBody(v *GetVideoAnalysisTaskResponseBody) *GetVideoAnalysisTaskResponse {
	s.Body = v
	return s
}

type ListHotTopicSummariesRequest struct {
	// example:
	//
	// xx
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// JlroP3CjgQh5PQDlH3ArzADkBTPZgVqo+64jhZRglNq0mEYoV5SlGb/Juvo8CdfYE9rlwEr2pIJQwdaYotak9g==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListHotTopicSummariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesRequest) SetCategory(v string) *ListHotTopicSummariesRequest {
	s.Category = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetHotTopic(v string) *ListHotTopicSummariesRequest {
	s.HotTopic = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetHotTopicVersion(v string) *ListHotTopicSummariesRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetMaxResults(v int32) *ListHotTopicSummariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetNextToken(v string) *ListHotTopicSummariesRequest {
	s.NextToken = &v
	return s
}

type ListHotTopicSummariesResponseBody struct {
	// example:
	//
	// xx
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListHotTopicSummariesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// JlroP3CjgQh5PQDlH3ArzADkBTPZgVqo+64jhZRglNq0mEYoV5SlGb/Juvo8CdfYE9rlwEr2pIJQwdaYotak9g==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListHotTopicSummariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBody) SetCode(v string) *ListHotTopicSummariesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetData(v []*ListHotTopicSummariesResponseBodyData) *ListHotTopicSummariesResponseBody {
	s.Data = v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetHttpStatusCode(v int32) *ListHotTopicSummariesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetMaxResults(v int32) *ListHotTopicSummariesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetMessage(v string) *ListHotTopicSummariesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetNextToken(v string) *ListHotTopicSummariesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetRequestId(v string) *ListHotTopicSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetSuccess(v bool) *ListHotTopicSummariesResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetTotalCount(v int32) *ListHotTopicSummariesResponseBody {
	s.TotalCount = &v
	return s
}

type ListHotTopicSummariesResponseBodyData struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 1000000
	HotValue *float64 `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	// example:
	//
	// db5dc5b3d8954a30b65ba700c9dda3bb
	Id      *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	News    []*ListHotTopicSummariesResponseBodyDataNews  `json:"news,omitempty" xml:"news,omitempty" type:"Repeated"`
	Summary *ListHotTopicSummariesResponseBodyDataSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// xx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyData) SetCategory(v string) *ListHotTopicSummariesResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotTopic(v string) *ListHotTopicSummariesResponseBodyData {
	s.HotTopic = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotTopicVersion(v string) *ListHotTopicSummariesResponseBodyData {
	s.HotTopicVersion = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotValue(v float64) *ListHotTopicSummariesResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetId(v string) *ListHotTopicSummariesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetNews(v []*ListHotTopicSummariesResponseBodyDataNews) *ListHotTopicSummariesResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetSummary(v *ListHotTopicSummariesResponseBodyDataSummary) *ListHotTopicSummariesResponseBodyData {
	s.Summary = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetTextSummary(v string) *ListHotTopicSummariesResponseBodyData {
	s.TextSummary = &v
	return s
}

type ListHotTopicSummariesResponseBodyDataNews struct {
	Comments []*ListHotTopicSummariesResponseBodyDataNewsComments `json:"comments,omitempty" xml:"comments,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-10 15:32:00
	PubTime *string `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// url
	//
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataNews) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetComments(v []*ListHotTopicSummariesResponseBodyDataNewsComments) *ListHotTopicSummariesResponseBodyDataNews {
	s.Comments = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetContent(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetPubTime(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetTitle(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetUrl(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Url = &v
	return s
}

type ListHotTopicSummariesResponseBodyDataNewsComments struct {
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataNewsComments) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataNewsComments) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataNewsComments) SetText(v string) *ListHotTopicSummariesResponseBodyDataNewsComments {
	s.Text = &v
	return s
}

type ListHotTopicSummariesResponseBodyDataSummary struct {
	Summaries []*ListHotTopicSummariesResponseBodyDataSummarySummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
}

func (s ListHotTopicSummariesResponseBodyDataSummary) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataSummary) SetSummaries(v []*ListHotTopicSummariesResponseBodyDataSummarySummaries) *ListHotTopicSummariesResponseBodyDataSummary {
	s.Summaries = v
	return s
}

type ListHotTopicSummariesResponseBodyDataSummarySummaries struct {
	// example:
	//
	// xx
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataSummarySummaries) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataSummarySummaries) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) SetSummary(v string) *ListHotTopicSummariesResponseBodyDataSummarySummaries {
	s.Summary = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) SetTitle(v string) *ListHotTopicSummariesResponseBodyDataSummarySummaries {
	s.Title = &v
	return s
}

type ListHotTopicSummariesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotTopicSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotTopicSummariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponse) SetHeaders(v map[string]*string) *ListHotTopicSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListHotTopicSummariesResponse) SetStatusCode(v int32) *ListHotTopicSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotTopicSummariesResponse) SetBody(v *ListHotTopicSummariesResponseBody) *ListHotTopicSummariesResponse {
	s.Body = v
	return s
}

type RunHotTopicChatRequest struct {
	Category        *string   `json:"category,omitempty" xml:"category,omitempty"`
	GenerateOptions []*string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string   `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	HotTopics       []*string `json:"hotTopics,omitempty" xml:"hotTopics,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ImageCount *int32                            `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	Messages   []*RunHotTopicChatRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId             *string                                              `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	Prompt                        *string                                              `json:"prompt,omitempty" xml:"prompt,omitempty"`
	StepForBroadcastContentConfig *RunHotTopicChatRequestStepForBroadcastContentConfig `json:"stepForBroadcastContentConfig,omitempty" xml:"stepForBroadcastContentConfig,omitempty" type:"Struct"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunHotTopicChatRequest) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequest) SetCategory(v string) *RunHotTopicChatRequest {
	s.Category = &v
	return s
}

func (s *RunHotTopicChatRequest) SetGenerateOptions(v []*string) *RunHotTopicChatRequest {
	s.GenerateOptions = v
	return s
}

func (s *RunHotTopicChatRequest) SetHotTopicVersion(v string) *RunHotTopicChatRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatRequest) SetHotTopics(v []*string) *RunHotTopicChatRequest {
	s.HotTopics = v
	return s
}

func (s *RunHotTopicChatRequest) SetImageCount(v int32) *RunHotTopicChatRequest {
	s.ImageCount = &v
	return s
}

func (s *RunHotTopicChatRequest) SetMessages(v []*RunHotTopicChatRequestMessages) *RunHotTopicChatRequest {
	s.Messages = v
	return s
}

func (s *RunHotTopicChatRequest) SetModelCustomPromptTemplate(v string) *RunHotTopicChatRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunHotTopicChatRequest) SetModelId(v string) *RunHotTopicChatRequest {
	s.ModelId = &v
	return s
}

func (s *RunHotTopicChatRequest) SetOriginalSessionId(v string) *RunHotTopicChatRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunHotTopicChatRequest) SetPrompt(v string) *RunHotTopicChatRequest {
	s.Prompt = &v
	return s
}

func (s *RunHotTopicChatRequest) SetStepForBroadcastContentConfig(v *RunHotTopicChatRequestStepForBroadcastContentConfig) *RunHotTopicChatRequest {
	s.StepForBroadcastContentConfig = v
	return s
}

func (s *RunHotTopicChatRequest) SetTaskId(v string) *RunHotTopicChatRequest {
	s.TaskId = &v
	return s
}

type RunHotTopicChatRequestMessages struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-12-10 18:51:29
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunHotTopicChatRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatRequestMessages) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestMessages) SetContent(v string) *RunHotTopicChatRequestMessages {
	s.Content = &v
	return s
}

func (s *RunHotTopicChatRequestMessages) SetCreateTime(v string) *RunHotTopicChatRequestMessages {
	s.CreateTime = &v
	return s
}

func (s *RunHotTopicChatRequestMessages) SetRole(v string) *RunHotTopicChatRequestMessages {
	s.Role = &v
	return s
}

type RunHotTopicChatRequestStepForBroadcastContentConfig struct {
	Categories            []*string                                                                   `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	CustomHotValueWeights []*RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights `json:"customHotValueWeights,omitempty" xml:"customHotValueWeights,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TopicCount *int32 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfig) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfig) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetCategories(v []*string) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.Categories = v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetCustomHotValueWeights(v []*RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.CustomHotValueWeights = v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetTopicCount(v int32) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.TopicCount = &v
	return s
}

type RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights struct {
	// example:
	//
	// comments
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) SetDimension(v string) *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights {
	s.Dimension = &v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) SetWeight(v int32) *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights {
	s.Weight = &v
	return s
}

type RunHotTopicChatShrinkRequest struct {
	Category              *string `json:"category,omitempty" xml:"category,omitempty"`
	GenerateOptionsShrink *string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	HotTopicsShrink *string `json:"hotTopics,omitempty" xml:"hotTopics,omitempty"`
	// example:
	//
	// 1
	ImageCount     *int32  `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	MessagesShrink *string `json:"messages,omitempty" xml:"messages,omitempty"`
	// example:
	//
	// xx
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId                   *string `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	Prompt                              *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	StepForBroadcastContentConfigShrink *string `json:"stepForBroadcastContentConfig,omitempty" xml:"stepForBroadcastContentConfig,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunHotTopicChatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatShrinkRequest) SetCategory(v string) *RunHotTopicChatShrinkRequest {
	s.Category = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetGenerateOptionsShrink(v string) *RunHotTopicChatShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetHotTopicVersion(v string) *RunHotTopicChatShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetHotTopicsShrink(v string) *RunHotTopicChatShrinkRequest {
	s.HotTopicsShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetImageCount(v int32) *RunHotTopicChatShrinkRequest {
	s.ImageCount = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetMessagesShrink(v string) *RunHotTopicChatShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetModelCustomPromptTemplate(v string) *RunHotTopicChatShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetModelId(v string) *RunHotTopicChatShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetOriginalSessionId(v string) *RunHotTopicChatShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetPrompt(v string) *RunHotTopicChatShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetStepForBroadcastContentConfigShrink(v string) *RunHotTopicChatShrinkRequest {
	s.StepForBroadcastContentConfigShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetTaskId(v string) *RunHotTopicChatShrinkRequest {
	s.TaskId = &v
	return s
}

type RunHotTopicChatResponseBody struct {
	Header  *RunHotTopicChatResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunHotTopicChatResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 04DA1A52-4E51-56CB-BA64-FDDA0B53BAE8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunHotTopicChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBody) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBody) SetHeader(v *RunHotTopicChatResponseBodyHeader) *RunHotTopicChatResponseBody {
	s.Header = v
	return s
}

func (s *RunHotTopicChatResponseBody) SetPayload(v *RunHotTopicChatResponseBodyPayload) *RunHotTopicChatResponseBody {
	s.Payload = v
	return s
}

func (s *RunHotTopicChatResponseBody) SetRequestId(v string) *RunHotTopicChatResponseBody {
	s.RequestId = &v
	return s
}

type RunHotTopicChatResponseBodyHeader struct {
	// example:
	//
	// InvalidParam
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xx
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// xxx
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

func (s RunHotTopicChatResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyHeader) SetErrorCode(v string) *RunHotTopicChatResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetErrorMessage(v string) *RunHotTopicChatResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetEvent(v string) *RunHotTopicChatResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetEventInfo(v string) *RunHotTopicChatResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetSessionId(v string) *RunHotTopicChatResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetTaskId(v string) *RunHotTopicChatResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetTraceId(v string) *RunHotTopicChatResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunHotTopicChatResponseBodyPayload struct {
	Output *RunHotTopicChatResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunHotTopicChatResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunHotTopicChatResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayload) SetOutput(v *RunHotTopicChatResponseBodyPayloadOutput) *RunHotTopicChatResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayload) SetUsage(v *RunHotTopicChatResponseBodyPayloadUsage) *RunHotTopicChatResponseBodyPayload {
	s.Usage = v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutput struct {
	Articles          []*RunHotTopicChatResponseBodyPayloadOutputArticles          `json:"articles,omitempty" xml:"articles,omitempty" type:"Repeated"`
	HotTopicSummaries []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries `json:"hotTopicSummaries,omitempty" xml:"hotTopicSummaries,omitempty" type:"Repeated"`
	MultimodalMedias  []*RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias  `json:"multimodalMedias,omitempty" xml:"multimodalMedias,omitempty" type:"Repeated"`
	RecommendQueries  []*string                                                    `json:"recommendQueries,omitempty" xml:"recommendQueries,omitempty" type:"Repeated"`
	SearchQuery       *string                                                      `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetArticles(v []*RunHotTopicChatResponseBodyPayloadOutputArticles) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetHotTopicSummaries(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) *RunHotTopicChatResponseBodyPayloadOutput {
	s.HotTopicSummaries = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetMultimodalMedias(v []*RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) *RunHotTopicChatResponseBodyPayloadOutput {
	s.MultimodalMedias = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetRecommendQueries(v []*string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.RecommendQueries = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetSearchQuery(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetText(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutputArticles struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-22 16:45:06
	PubTime          *string  `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	Score            *float64 `json:"score,omitempty" xml:"score,omitempty"`
	SearchSourceName *string  `json:"searchSourceName,omitempty" xml:"searchSourceName,omitempty"`
	Select           *bool    `json:"select,omitempty" xml:"select,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputArticles) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetContent(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetScore(v float64) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Score = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSearchSourceName(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.SearchSourceName = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSelect(v bool) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Select = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetTitle(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries struct {
	// example:
	//
	// 100000
	CustomHotValue *float64 `json:"customHotValue,omitempty" xml:"customHotValue,omitempty"`
	// example:
	//
	// xxx
	CustomTextSummary *string `json:"customTextSummary,omitempty" xml:"customTextSummary,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_08
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 100000
	HotValue *float64                                                           `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	Images   []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	News     []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews   `json:"news,omitempty" xml:"news,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetCustomHotValue(v float64) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.CustomHotValue = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetCustomTextSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.CustomTextSummary = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotTopic(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotTopic = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotTopicVersion(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotValue(v float64) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotValue = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetImages(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.Images = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetNews(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.News = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetTextSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.TextSummary = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages struct {
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages {
	s.Url = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews struct {
	// example:
	//
	// xxx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) SetTitle(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews {
	s.Title = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews {
	s.Url = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias struct {
	// example:
	//
	// http://xxxx
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// example:
	//
	// image
	MediaType *string  `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	SortScore *float64 `json:"sortScore,omitempty" xml:"sortScore,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetFileUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetMediaType(v string) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetSortScore(v float64) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.SortScore = &v
	return s
}

type RunHotTopicChatResponseBodyPayloadUsage struct {
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

func (s RunHotTopicChatResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetInputTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunHotTopicChatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunHotTopicChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunHotTopicChatResponse) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicChatResponse) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponse) SetHeaders(v map[string]*string) *RunHotTopicChatResponse {
	s.Headers = v
	return s
}

func (s *RunHotTopicChatResponse) SetStatusCode(v int32) *RunHotTopicChatResponse {
	s.StatusCode = &v
	return s
}

func (s *RunHotTopicChatResponse) SetBody(v *RunHotTopicChatResponseBody) *RunHotTopicChatResponse {
	s.Body = v
	return s
}

type RunHotTopicSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-16_8
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// This parameter is required.
	StepForCustomSummaryStyleConfig *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig `json:"stepForCustomSummaryStyleConfig,omitempty" xml:"stepForCustomSummaryStyleConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	TopicIds []*string `json:"topicIds,omitempty" xml:"topicIds,omitempty" type:"Repeated"`
}

func (s RunHotTopicSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryRequest) SetHotTopicVersion(v string) *RunHotTopicSummaryRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicSummaryRequest) SetStepForCustomSummaryStyleConfig(v *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) *RunHotTopicSummaryRequest {
	s.StepForCustomSummaryStyleConfig = v
	return s
}

func (s *RunHotTopicSummaryRequest) SetTopicIds(v []*string) *RunHotTopicSummaryRequest {
	s.TopicIds = v
	return s
}

type RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig struct {
	// example:
	//
	// 2
	SummaryImageCount *int32 `json:"summaryImageCount,omitempty" xml:"summaryImageCount,omitempty"`
	// example:
	//
	// qwen-max
	SummaryModel *string `json:"summaryModel,omitempty" xml:"summaryModel,omitempty"`
	// example:
	//
	// xxxx
	SummaryPrompt *string `json:"summaryPrompt,omitempty" xml:"summaryPrompt,omitempty"`
}

func (s RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryImageCount(v int32) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryImageCount = &v
	return s
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryModel(v string) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryModel = &v
	return s
}

func (s *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig) SetSummaryPrompt(v string) *RunHotTopicSummaryRequestStepForCustomSummaryStyleConfig {
	s.SummaryPrompt = &v
	return s
}

type RunHotTopicSummaryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-16_8
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// This parameter is required.
	StepForCustomSummaryStyleConfigShrink *string `json:"stepForCustomSummaryStyleConfig,omitempty" xml:"stepForCustomSummaryStyleConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	TopicIdsShrink *string `json:"topicIds,omitempty" xml:"topicIds,omitempty"`
}

func (s RunHotTopicSummaryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryShrinkRequest) SetHotTopicVersion(v string) *RunHotTopicSummaryShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicSummaryShrinkRequest) SetStepForCustomSummaryStyleConfigShrink(v string) *RunHotTopicSummaryShrinkRequest {
	s.StepForCustomSummaryStyleConfigShrink = &v
	return s
}

func (s *RunHotTopicSummaryShrinkRequest) SetTopicIdsShrink(v string) *RunHotTopicSummaryShrinkRequest {
	s.TopicIdsShrink = &v
	return s
}

type RunHotTopicSummaryResponseBody struct {
	Header  *RunHotTopicSummaryResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunHotTopicSummaryResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunHotTopicSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBody) SetHeader(v *RunHotTopicSummaryResponseBodyHeader) *RunHotTopicSummaryResponseBody {
	s.Header = v
	return s
}

func (s *RunHotTopicSummaryResponseBody) SetPayload(v *RunHotTopicSummaryResponseBodyPayload) *RunHotTopicSummaryResponseBody {
	s.Payload = v
	return s
}

func (s *RunHotTopicSummaryResponseBody) SetRequestId(v string) *RunHotTopicSummaryResponseBody {
	s.RequestId = &v
	return s
}

type RunHotTopicSummaryResponseBodyHeader struct {
	// example:
	//
	// AccessForbidden
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxxxx
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunHotTopicSummaryResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetErrorCode(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetErrorMessage(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetEvent(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetSessionId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetTaskId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetTraceId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunHotTopicSummaryResponseBodyPayload struct {
	Output *RunHotTopicSummaryResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunHotTopicSummaryResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunHotTopicSummaryResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayload) SetOutput(v *RunHotTopicSummaryResponseBodyPayloadOutput) *RunHotTopicSummaryResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayload) SetUsage(v *RunHotTopicSummaryResponseBodyPayloadUsage) *RunHotTopicSummaryResponseBodyPayload {
	s.Usage = v
	return s
}

type RunHotTopicSummaryResponseBodyPayloadOutput struct {
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	TopicId *string `json:"topicId,omitempty" xml:"topicId,omitempty"`
}

func (s RunHotTopicSummaryResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) SetText(v string) *RunHotTopicSummaryResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) SetTopicId(v string) *RunHotTopicSummaryResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

type RunHotTopicSummaryResponseBodyPayloadUsage struct {
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

func (s RunHotTopicSummaryResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetInputTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunHotTopicSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunHotTopicSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunHotTopicSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s RunHotTopicSummaryResponse) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponse) SetHeaders(v map[string]*string) *RunHotTopicSummaryResponse {
	s.Headers = v
	return s
}

func (s *RunHotTopicSummaryResponse) SetStatusCode(v int32) *RunHotTopicSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *RunHotTopicSummaryResponse) SetBody(v *RunHotTopicSummaryResponseBody) *RunHotTopicSummaryResponse {
	s.Body = v
	return s
}

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
	CustomLimitation *string `json:"customLimitation,omitempty" xml:"customLimitation,omitempty"`
	CustomPrompt     *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	InputExample     *string `json:"inputExample,omitempty" xml:"inputExample,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId        *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	OutputExample  *string `json:"outputExample,omitempty" xml:"outputExample,omitempty"`
	SourceMaterial *string `json:"sourceMaterial,omitempty" xml:"sourceMaterial,omitempty"`
	WritingType    *string `json:"writingType,omitempty" xml:"writingType,omitempty"`
}

func (s RunMarketingInformationWritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMarketingInformationWritingRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingRequest) SetCustomLimitation(v string) *RunMarketingInformationWritingRequest {
	s.CustomLimitation = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetCustomPrompt(v string) *RunMarketingInformationWritingRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetInputExample(v string) *RunMarketingInformationWritingRequest {
	s.InputExample = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetModelId(v string) *RunMarketingInformationWritingRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetOutputExample(v string) *RunMarketingInformationWritingRequest {
	s.OutputExample = &v
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

type RunNetworkContentAuditRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string                              `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*RunNetworkContentAuditRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunNetworkContentAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditRequest) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditRequest) SetBusinessType(v string) *RunNetworkContentAuditRequest {
	s.BusinessType = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetContent(v string) *RunNetworkContentAuditRequest {
	s.Content = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetExtraInfo(v string) *RunNetworkContentAuditRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetModelId(v string) *RunNetworkContentAuditRequest {
	s.ModelId = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetOutputFormat(v string) *RunNetworkContentAuditRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetTags(v []*RunNetworkContentAuditRequestTags) *RunNetworkContentAuditRequest {
	s.Tags = v
	return s
}

func (s *RunNetworkContentAuditRequest) SetTaskDescription(v string) *RunNetworkContentAuditRequest {
	s.TaskDescription = &v
	return s
}

type RunNetworkContentAuditRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunNetworkContentAuditRequestTags) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditRequestTags) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditRequestTags) SetTagDefinePrompt(v string) *RunNetworkContentAuditRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunNetworkContentAuditRequestTags) SetTagName(v string) *RunNetworkContentAuditRequestTags {
	s.TagName = &v
	return s
}

type RunNetworkContentAuditShrinkRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	TagsShrink   *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunNetworkContentAuditShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditShrinkRequest) SetBusinessType(v string) *RunNetworkContentAuditShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetContent(v string) *RunNetworkContentAuditShrinkRequest {
	s.Content = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetExtraInfo(v string) *RunNetworkContentAuditShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetModelId(v string) *RunNetworkContentAuditShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetOutputFormat(v string) *RunNetworkContentAuditShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetTagsShrink(v string) *RunNetworkContentAuditShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetTaskDescription(v string) *RunNetworkContentAuditShrinkRequest {
	s.TaskDescription = &v
	return s
}

type RunNetworkContentAuditResponseBody struct {
	Header  *RunNetworkContentAuditResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunNetworkContentAuditResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunNetworkContentAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponseBody) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBody) SetHeader(v *RunNetworkContentAuditResponseBodyHeader) *RunNetworkContentAuditResponseBody {
	s.Header = v
	return s
}

func (s *RunNetworkContentAuditResponseBody) SetPayload(v *RunNetworkContentAuditResponseBodyPayload) *RunNetworkContentAuditResponseBody {
	s.Payload = v
	return s
}

func (s *RunNetworkContentAuditResponseBody) SetRequestId(v string) *RunNetworkContentAuditResponseBody {
	s.RequestId = &v
	return s
}

type RunNetworkContentAuditResponseBodyHeader struct {
	// example:
	//
	// AccessForbidden
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxxxx
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunNetworkContentAuditResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetErrorCode(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetErrorMessage(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetEvent(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetSessionId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetTaskId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetTraceId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunNetworkContentAuditResponseBodyPayload struct {
	Output *RunNetworkContentAuditResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunNetworkContentAuditResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunNetworkContentAuditResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayload) SetOutput(v *RunNetworkContentAuditResponseBodyPayloadOutput) *RunNetworkContentAuditResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayload) SetUsage(v *RunNetworkContentAuditResponseBodyPayloadUsage) *RunNetworkContentAuditResponseBodyPayload {
	s.Usage = v
	return s
}

type RunNetworkContentAuditResponseBodyPayloadOutput struct {
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunNetworkContentAuditResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayloadOutput) SetText(v string) *RunNetworkContentAuditResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunNetworkContentAuditResponseBodyPayloadUsage struct {
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

func (s RunNetworkContentAuditResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetInputTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunNetworkContentAuditResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNetworkContentAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNetworkContentAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s RunNetworkContentAuditResponse) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponse) SetHeaders(v map[string]*string) *RunNetworkContentAuditResponse {
	s.Headers = v
	return s
}

func (s *RunNetworkContentAuditResponse) SetStatusCode(v int32) *RunNetworkContentAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNetworkContentAuditResponse) SetBody(v *RunNetworkContentAuditResponseBody) *RunNetworkContentAuditResponse {
	s.Body = v
	return s
}

type RunScriptChatRequest struct {
	// This parameter is required.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunScriptChatRequest) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatRequest) GoString() string {
	return s.String()
}

func (s *RunScriptChatRequest) SetPrompt(v string) *RunScriptChatRequest {
	s.Prompt = &v
	return s
}

func (s *RunScriptChatRequest) SetTaskId(v string) *RunScriptChatRequest {
	s.TaskId = &v
	return s
}

type RunScriptChatResponseBody struct {
	// example:
	//
	// true
	End     *bool                             `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptChatResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptChatResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBody) SetEnd(v bool) *RunScriptChatResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptChatResponseBody) SetHeader(v *RunScriptChatResponseBodyHeader) *RunScriptChatResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptChatResponseBody) SetPayload(v *RunScriptChatResponseBodyPayload) *RunScriptChatResponseBody {
	s.Payload = v
	return s
}

type RunScriptChatResponseBodyHeader struct {
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
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F8A35034-EDCF-5C50-95A5-1044316F36E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 147648697127_914847410985_1730600302167
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptChatResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyHeader) SetErrorCode(v string) *RunScriptChatResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetErrorMessage(v string) *RunScriptChatResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetEvent(v string) *RunScriptChatResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetEventInfo(v string) *RunScriptChatResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetRequestId(v string) *RunScriptChatResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetSessionId(v string) *RunScriptChatResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetTaskId(v string) *RunScriptChatResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetTraceId(v string) *RunScriptChatResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunScriptChatResponseBodyPayload struct {
	Output *RunScriptChatResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptChatResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptChatResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayload) SetOutput(v *RunScriptChatResponseBodyPayloadOutput) *RunScriptChatResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptChatResponseBodyPayload) SetUsage(v *RunScriptChatResponseBodyPayloadUsage) *RunScriptChatResponseBodyPayload {
	s.Usage = v
	return s
}

type RunScriptChatResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptChatResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayloadOutput) SetText(v string) *RunScriptChatResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunScriptChatResponseBodyPayloadUsage struct {
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

func (s RunScriptChatResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunScriptChatResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptChatResponse) String() string {
	return tea.Prettify(s)
}

func (s RunScriptChatResponse) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponse) SetHeaders(v map[string]*string) *RunScriptChatResponse {
	s.Headers = v
	return s
}

func (s *RunScriptChatResponse) SetStatusCode(v int32) *RunScriptChatResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptChatResponse) SetBody(v *RunScriptChatResponseBody) *RunScriptChatResponse {
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

type RunScriptRefineRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunScriptRefineRequest) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineRequest) GoString() string {
	return s.String()
}

func (s *RunScriptRefineRequest) SetTaskId(v string) *RunScriptRefineRequest {
	s.TaskId = &v
	return s
}

type RunScriptRefineResponseBody struct {
	End     *bool                               `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptRefineResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptRefineResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptRefineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBody) SetEnd(v bool) *RunScriptRefineResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptRefineResponseBody) SetHeader(v *RunScriptRefineResponseBodyHeader) *RunScriptRefineResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptRefineResponseBody) SetPayload(v *RunScriptRefineResponseBodyPayload) *RunScriptRefineResponseBody {
	s.Payload = v
	return s
}

type RunScriptRefineResponseBodyHeader struct {
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
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F8A35034-EDCF-5C50-95A5-1044316F36E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 17dc8bcd-f34a-46d1-a7a3-0fa3d1ce3824
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 14356391-6c6c-40d5-b80a-8ecd03b69d72
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptRefineResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyHeader) SetErrorCode(v string) *RunScriptRefineResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetErrorMessage(v string) *RunScriptRefineResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetEvent(v string) *RunScriptRefineResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetEventInfo(v string) *RunScriptRefineResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetRequestId(v string) *RunScriptRefineResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetSessionId(v string) *RunScriptRefineResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetTaskId(v string) *RunScriptRefineResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetTraceId(v string) *RunScriptRefineResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunScriptRefineResponseBodyPayload struct {
	Output *RunScriptRefineResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptRefineResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptRefineResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayload) SetOutput(v *RunScriptRefineResponseBodyPayloadOutput) *RunScriptRefineResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptRefineResponseBodyPayload) SetUsage(v *RunScriptRefineResponseBodyPayloadUsage) *RunScriptRefineResponseBodyPayload {
	s.Usage = v
	return s
}

type RunScriptRefineResponseBodyPayloadOutput struct {
	Content []map[string]*string `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Outline *string              `json:"outline,omitempty" xml:"outline,omitempty"`
	Role    *string              `json:"role,omitempty" xml:"role,omitempty"`
	Scene   *string              `json:"scene,omitempty" xml:"scene,omitempty"`
	Summary *string              `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptRefineResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetContent(v []map[string]*string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Content = v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetOutline(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Outline = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetRole(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Role = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetScene(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Scene = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetSummary(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Summary = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetText(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunScriptRefineResponseBodyPayloadUsage struct {
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

func (s RunScriptRefineResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunScriptRefineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptRefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptRefineResponse) String() string {
	return tea.Prettify(s)
}

func (s RunScriptRefineResponse) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponse) SetHeaders(v map[string]*string) *RunScriptRefineResponse {
	s.Headers = v
	return s
}

func (s *RunScriptRefineResponse) SetStatusCode(v int32) *RunScriptRefineResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptRefineResponse) SetBody(v *RunScriptRefineResponseBody) *RunScriptRefineResponse {
	s.Body = v
	return s
}

type RunStyleWritingRequest struct {
	LearningSamples    []*string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty" type:"Repeated"`
	ProcessStage       *string   `json:"processStage,omitempty" xml:"processStage,omitempty"`
	ReferenceMaterials []*string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty" type:"Repeated"`
	StyleFeature       *string   `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	UseSearch          *bool     `json:"useSearch,omitempty" xml:"useSearch,omitempty"`
	WritingTheme       *string   `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
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

func (s *RunStyleWritingRequest) SetProcessStage(v string) *RunStyleWritingRequest {
	s.ProcessStage = &v
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

func (s *RunStyleWritingRequest) SetUseSearch(v bool) *RunStyleWritingRequest {
	s.UseSearch = &v
	return s
}

func (s *RunStyleWritingRequest) SetWritingTheme(v string) *RunStyleWritingRequest {
	s.WritingTheme = &v
	return s
}

type RunStyleWritingShrinkRequest struct {
	LearningSamplesShrink    *string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty"`
	ProcessStage             *string `json:"processStage,omitempty" xml:"processStage,omitempty"`
	ReferenceMaterialsShrink *string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty"`
	StyleFeature             *string `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	UseSearch                *bool   `json:"useSearch,omitempty" xml:"useSearch,omitempty"`
	WritingTheme             *string `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
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

func (s *RunStyleWritingShrinkRequest) SetProcessStage(v string) *RunStyleWritingShrinkRequest {
	s.ProcessStage = &v
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

func (s *RunStyleWritingShrinkRequest) SetUseSearch(v bool) *RunStyleWritingShrinkRequest {
	s.UseSearch = &v
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

type RunTagMiningAnalysisRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string                            `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*RunTagMiningAnalysisRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunTagMiningAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisRequest) SetBusinessType(v string) *RunTagMiningAnalysisRequest {
	s.BusinessType = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetContent(v string) *RunTagMiningAnalysisRequest {
	s.Content = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetExtraInfo(v string) *RunTagMiningAnalysisRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetModelId(v string) *RunTagMiningAnalysisRequest {
	s.ModelId = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetOutputFormat(v string) *RunTagMiningAnalysisRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetTags(v []*RunTagMiningAnalysisRequestTags) *RunTagMiningAnalysisRequest {
	s.Tags = v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetTaskDescription(v string) *RunTagMiningAnalysisRequest {
	s.TaskDescription = &v
	return s
}

type RunTagMiningAnalysisRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunTagMiningAnalysisRequestTags) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisRequestTags) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisRequestTags) SetTagDefinePrompt(v string) *RunTagMiningAnalysisRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunTagMiningAnalysisRequestTags) SetTagName(v string) *RunTagMiningAnalysisRequestTags {
	s.TagName = &v
	return s
}

type RunTagMiningAnalysisShrinkRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	TagsShrink   *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunTagMiningAnalysisShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisShrinkRequest) SetBusinessType(v string) *RunTagMiningAnalysisShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetContent(v string) *RunTagMiningAnalysisShrinkRequest {
	s.Content = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetExtraInfo(v string) *RunTagMiningAnalysisShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetModelId(v string) *RunTagMiningAnalysisShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetOutputFormat(v string) *RunTagMiningAnalysisShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetTagsShrink(v string) *RunTagMiningAnalysisShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetTaskDescription(v string) *RunTagMiningAnalysisShrinkRequest {
	s.TaskDescription = &v
	return s
}

type RunTagMiningAnalysisResponseBody struct {
	Header  *RunTagMiningAnalysisResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunTagMiningAnalysisResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunTagMiningAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBody) SetHeader(v *RunTagMiningAnalysisResponseBodyHeader) *RunTagMiningAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunTagMiningAnalysisResponseBody) SetPayload(v *RunTagMiningAnalysisResponseBodyPayload) *RunTagMiningAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunTagMiningAnalysisResponseBody) SetRequestId(v string) *RunTagMiningAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type RunTagMiningAnalysisResponseBodyHeader struct {
	// example:
	//
	// AccessForbidden
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxxxx
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunTagMiningAnalysisResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetErrorCode(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetEvent(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetSessionId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetTaskId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetTraceId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunTagMiningAnalysisResponseBodyPayload struct {
	Output *RunTagMiningAnalysisResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunTagMiningAnalysisResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunTagMiningAnalysisResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayload) SetOutput(v *RunTagMiningAnalysisResponseBodyPayloadOutput) *RunTagMiningAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayload) SetUsage(v *RunTagMiningAnalysisResponseBodyPayloadUsage) *RunTagMiningAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

type RunTagMiningAnalysisResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunTagMiningAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayloadOutput) SetText(v string) *RunTagMiningAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunTagMiningAnalysisResponseBodyPayloadUsage struct {
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

func (s RunTagMiningAnalysisResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunTagMiningAnalysisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTagMiningAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTagMiningAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunTagMiningAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponse) SetHeaders(v map[string]*string) *RunTagMiningAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetStatusCode(v int32) *RunTagMiningAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetBody(v *RunTagMiningAnalysisResponseBody) *RunTagMiningAnalysisResponse {
	s.Body = v
	return s
}

type RunVideoAnalysisRequest struct {
	FaceIdentitySimilarityMinScore *float32                                  `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethod              *RunVideoAnalysisRequestFrameSampleMethod `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty" type:"Struct"`
	GenerateOptions                []*string                                 `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// english
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
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
	OriginalSessionId *string  `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	SnapshotInterval  *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval *int32 `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string                                    `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TextProcessTasks               []*RunVideoAnalysisRequestTextProcessTasks `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty" type:"Repeated"`
	VideoExtraInfo                 *string                                    `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string                                    `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId               *string                              `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRoles                 []*RunVideoAnalysisRequestVideoRoles `json:"videoRoles,omitempty" xml:"videoRoles,omitempty" type:"Repeated"`
	VideoShotFaceIdentityCount *int32                               `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
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

func (s *RunVideoAnalysisRequest) SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetFrameSampleMethod(v *RunVideoAnalysisRequestFrameSampleMethod) *RunVideoAnalysisRequest {
	s.FrameSampleMethod = v
	return s
}

func (s *RunVideoAnalysisRequest) SetGenerateOptions(v []*string) *RunVideoAnalysisRequest {
	s.GenerateOptions = v
	return s
}

func (s *RunVideoAnalysisRequest) SetLanguage(v string) *RunVideoAnalysisRequest {
	s.Language = &v
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

func (s *RunVideoAnalysisRequest) SetSnapshotInterval(v float64) *RunVideoAnalysisRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetSplitInterval(v int32) *RunVideoAnalysisRequest {
	s.SplitInterval = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetTaskId(v string) *RunVideoAnalysisRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetTextProcessTasks(v []*RunVideoAnalysisRequestTextProcessTasks) *RunVideoAnalysisRequest {
	s.TextProcessTasks = v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoExtraInfo(v string) *RunVideoAnalysisRequest {
	s.VideoExtraInfo = &v
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

func (s *RunVideoAnalysisRequest) SetVideoRoles(v []*RunVideoAnalysisRequestVideoRoles) *RunVideoAnalysisRequest {
	s.VideoRoles = v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoUrl(v string) *RunVideoAnalysisRequest {
	s.VideoUrl = &v
	return s
}

type RunVideoAnalysisRequestFrameSampleMethod struct {
	Interval   *float64 `json:"interval,omitempty" xml:"interval,omitempty"`
	MethodName *string  `json:"methodName,omitempty" xml:"methodName,omitempty"`
	Pixel      *int32   `json:"pixel,omitempty" xml:"pixel,omitempty"`
}

func (s RunVideoAnalysisRequestFrameSampleMethod) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisRequestFrameSampleMethod) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetInterval(v float64) *RunVideoAnalysisRequestFrameSampleMethod {
	s.Interval = &v
	return s
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetMethodName(v string) *RunVideoAnalysisRequestFrameSampleMethod {
	s.MethodName = &v
	return s
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetPixel(v int32) *RunVideoAnalysisRequestFrameSampleMethod {
	s.Pixel = &v
	return s
}

type RunVideoAnalysisRequestTextProcessTasks struct {
	ModelCustomPromptTemplate   *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	ModelId                     *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s RunVideoAnalysisRequestTextProcessTasks) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisRequestTextProcessTasks) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelId(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelId = &v
	return s
}

type RunVideoAnalysisRequestVideoRoles struct {
	RoleInfo *string   `json:"roleInfo,omitempty" xml:"roleInfo,omitempty"`
	RoleName *string   `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Urls     []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisRequestVideoRoles) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisRequestVideoRoles) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestVideoRoles) SetRoleInfo(v string) *RunVideoAnalysisRequestVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoRoles) SetRoleName(v string) *RunVideoAnalysisRequestVideoRoles {
	s.RoleName = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoRoles) SetUrls(v []*string) *RunVideoAnalysisRequestVideoRoles {
	s.Urls = v
	return s
}

type RunVideoAnalysisShrinkRequest struct {
	FaceIdentitySimilarityMinScore *float32 `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethodShrink        *string  `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty"`
	GenerateOptionsShrink          *string  `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// english
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
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
	OriginalSessionId *string  `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	SnapshotInterval  *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval *int32 `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TextProcessTasksShrink         *string `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty"`
	VideoExtraInfo                 *string `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId               *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRolesShrink           *string `json:"videoRoles,omitempty" xml:"videoRoles,omitempty"`
	VideoShotFaceIdentityCount *int32  `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
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

func (s *RunVideoAnalysisShrinkRequest) SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisShrinkRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetFrameSampleMethodShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.FrameSampleMethodShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetLanguage(v string) *RunVideoAnalysisShrinkRequest {
	s.Language = &v
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

func (s *RunVideoAnalysisShrinkRequest) SetSnapshotInterval(v float64) *RunVideoAnalysisShrinkRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetSplitInterval(v int32) *RunVideoAnalysisShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetTaskId(v string) *RunVideoAnalysisShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetTextProcessTasksShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.TextProcessTasksShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoExtraInfo(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoExtraInfo = &v
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

func (s *RunVideoAnalysisShrinkRequest) SetVideoRolesShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoRolesShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisShrinkRequest {
	s.VideoShotFaceIdentityCount = &v
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
	Usage  *RunVideoAnalysisResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
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

func (s *RunVideoAnalysisResponseBodyPayload) SetUsage(v *RunVideoAnalysisResponseBodyPayloadUsage) *RunVideoAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutput struct {
	// example:
	//
	// http://
	ResultJsonFileUrl              *string                                                                  `json:"resultJsonFileUrl,omitempty" xml:"resultJsonFileUrl,omitempty"`
	VideoAnalysisResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult            `json:"videoAnalysisResult,omitempty" xml:"videoAnalysisResult,omitempty" type:"Struct"`
	VideoCaptionResult             *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult             `json:"videoCaptionResult,omitempty" xml:"videoCaptionResult,omitempty" type:"Struct"`
	VideoGenerateResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult            `json:"videoGenerateResult,omitempty" xml:"videoGenerateResult,omitempty" type:"Struct"`
	VideoGenerateResults           []*RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults         `json:"videoGenerateResults,omitempty" xml:"videoGenerateResults,omitempty" type:"Repeated"`
	VideoMindMappingGenerateResult *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult `json:"videoMindMappingGenerateResult,omitempty" xml:"videoMindMappingGenerateResult,omitempty" type:"Struct"`
	VideoShotSnapshotResult        *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult        `json:"videoShotSnapshotResult,omitempty" xml:"videoShotSnapshotResult,omitempty" type:"Struct"`
	VideoTitleGenerateResult       *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult       `json:"videoTitleGenerateResult,omitempty" xml:"videoTitleGenerateResult,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetResultJsonFileUrl(v string) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.ResultJsonFileUrl = &v
	return s
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

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoGenerateResults(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoGenerateResults = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoMindMappingGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoMindMappingGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoShotSnapshotResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoShotSnapshotResult = v
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
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// qwen-vl-max
	ModelId                  *string                                                                                 `json:"modelId,omitempty" xml:"modelId,omitempty"`
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

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.ModelId = &v
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
	GenerateFinished *bool  `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// qwen-max
	ModelId     *string                                                            `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce *bool                                                              `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	ReasonText  *string                                                            `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text        *string                                                            `json:"text,omitempty" xml:"text,omitempty"`
	Usage       *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
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

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetIndex(v int32) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Index = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetReasonText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ReasonText = &v
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

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults struct {
	GenerateFinished *bool                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32                                                              `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string                                                             `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ReasonText       *string                                                             `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text             *string                                                             `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetIndex(v int32) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Index = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetReasonText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.ReasonText = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Usage = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.TotalTokens = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// true
	ModelId           *string                                                                                     `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce       *bool                                                                                       `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
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

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.ModelReduce = &v
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

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult struct {
	VideoShots []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots `json:"videoShots,omitempty" xml:"videoShots,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) SetVideoShots(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult {
	s.VideoShots = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots struct {
	EndTime         *int64                                                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EndTimeFormat   *string                                                                                     `json:"endTimeFormat,omitempty" xml:"endTimeFormat,omitempty"`
	StartTime       *int64                                                                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StartTimeFormat *string                                                                                     `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	VideoSnapshots  []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots `json:"videoSnapshots,omitempty" xml:"videoSnapshots,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetEndTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.EndTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetStartTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.StartTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetVideoSnapshots(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.VideoSnapshots = v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) SetUrl(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots {
	s.Url = &v
	return s
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// qwen-max
	ModelId     *string                                                                 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce *bool                                                                   `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	Text        *string                                                                 `json:"text,omitempty" xml:"text,omitempty"`
	Usage       *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
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

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.ModelReduce = &v
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

type RunVideoAnalysisResponseBodyPayloadUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
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

type SubmitTagMiningAnalysisTaskRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string   `json:"businessType,omitempty" xml:"businessType,omitempty"`
	Contents     []*string `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string                                   `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*SubmitTagMiningAnalysisTaskRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetBusinessType(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetContents(v []*string) *SubmitTagMiningAnalysisTaskRequest {
	s.Contents = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetModelId(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetTags(v []*SubmitTagMiningAnalysisTaskRequestTags) *SubmitTagMiningAnalysisTaskRequest {
	s.Tags = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetUrl(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.Url = &v
	return s
}

type SubmitTagMiningAnalysisTaskRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskRequestTags) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskRequestTags) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) SetTagDefinePrompt(v string) *SubmitTagMiningAnalysisTaskRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) SetTagName(v string) *SubmitTagMiningAnalysisTaskRequestTags {
	s.TagName = &v
	return s
}

type SubmitTagMiningAnalysisTaskShrinkRequest struct {
	// example:
	//
	// clueMining
	BusinessType   *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	ContentsShrink *string `json:"contents,omitempty" xml:"contents,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	TagsShrink   *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetBusinessType(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetContentsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetTagsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetUrl(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.Url = &v
	return s
}

type SubmitTagMiningAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitTagMiningAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetCode(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetData(v *SubmitTagMiningAnalysisTaskResponseBodyData) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetMessage(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetRequestId(v string) *SubmitTagMiningAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitTagMiningAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitTagMiningAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitTagMiningAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitTagMiningAnalysisTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTagMiningAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitTagMiningAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetBody(v *SubmitTagMiningAnalysisTaskResponseBody) *SubmitTagMiningAnalysisTaskResponse {
	s.Body = v
	return s
}

type SubmitVideoAnalysisTaskRequest struct {
	FaceIdentitySimilarityMinScore *float32                                         `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethod              *SubmitVideoAnalysisTaskRequestFrameSampleMethod `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty" type:"Struct"`
	GenerateOptions                []*string                                        `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// chinese
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
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
	// 2
	SnapshotInterval *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval                  *int32                                            `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	TextProcessTasks               []*SubmitVideoAnalysisTaskRequestTextProcessTasks `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty" type:"Repeated"`
	VideoExtraInfo                 *string                                           `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string                                           `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max-latest
	VideoModelId               *string                                     `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRoles                 []*SubmitVideoAnalysisTaskRequestVideoRoles `json:"videoRoles,omitempty" xml:"videoRoles,omitempty" type:"Repeated"`
	VideoShotFaceIdentityCount *int32                                      `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequest) SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetFrameSampleMethod(v *SubmitVideoAnalysisTaskRequestFrameSampleMethod) *SubmitVideoAnalysisTaskRequest {
	s.FrameSampleMethod = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetGenerateOptions(v []*string) *SubmitVideoAnalysisTaskRequest {
	s.GenerateOptions = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetLanguage(v string) *SubmitVideoAnalysisTaskRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelId(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetSplitInterval(v int32) *SubmitVideoAnalysisTaskRequest {
	s.SplitInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetTextProcessTasks(v []*SubmitVideoAnalysisTaskRequestTextProcessTasks) *SubmitVideoAnalysisTaskRequest {
	s.TextProcessTasks = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoModelId(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoRoles(v []*SubmitVideoAnalysisTaskRequestVideoRoles) *SubmitVideoAnalysisTaskRequest {
	s.VideoRoles = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoUrl(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoUrl = &v
	return s
}

type SubmitVideoAnalysisTaskRequestFrameSampleMethod struct {
	// example:
	//
	// 2
	Interval *float64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// standard
	MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	// example:
	//
	// 768
	Pixel *int32 `json:"pixel,omitempty" xml:"pixel,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestFrameSampleMethod) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestFrameSampleMethod) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetInterval(v float64) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.Interval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetMethodName(v string) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.MethodName = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetPixel(v int32) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.Pixel = &v
	return s
}

type SubmitVideoAnalysisTaskRequestTextProcessTasks struct {
	ModelCustomPromptTemplate   *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	ModelId                     *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestTextProcessTasks) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestTextProcessTasks) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelId(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelId = &v
	return s
}

type SubmitVideoAnalysisTaskRequestVideoRoles struct {
	RoleInfo *string   `json:"roleInfo,omitempty" xml:"roleInfo,omitempty"`
	RoleName *string   `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Urls     []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s SubmitVideoAnalysisTaskRequestVideoRoles) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestVideoRoles) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetRoleInfo(v string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetRoleName(v string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.RoleName = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetUrls(v []*string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.Urls = v
	return s
}

type SubmitVideoAnalysisTaskShrinkRequest struct {
	FaceIdentitySimilarityMinScore *float32 `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethodShrink        *string  `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty"`
	GenerateOptionsShrink          *string  `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// chinese
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
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
	// 2
	SnapshotInterval *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval                  *int32  `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	TextProcessTasksShrink         *string `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty"`
	VideoExtraInfo                 *string `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max-latest
	VideoModelId               *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRolesShrink           *string `json:"videoRoles,omitempty" xml:"videoRoles,omitempty"`
	VideoShotFaceIdentityCount *int32  `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s SubmitVideoAnalysisTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetFrameSampleMethodShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.FrameSampleMethodShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetGenerateOptionsShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetLanguage(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskShrinkRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetSplitInterval(v int32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetTextProcessTasksShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.TextProcessTasksShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoRolesShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoRolesShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoUrl(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoUrl = &v
	return s
}

type SubmitVideoAnalysisTaskResponseBody struct {
	// example:
	//
	// xx
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitVideoAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetCode(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetData(v *SubmitVideoAnalysisTaskResponseBodyData) *SubmitVideoAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitVideoAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetMessage(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetRequestId(v string) *SubmitVideoAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitVideoAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitVideoAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitVideoAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitVideoAnalysisTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitVideoAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoAnalysisTaskResponse) SetStatusCode(v int32) *SubmitVideoAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponse) SetBody(v *SubmitVideoAnalysisTaskResponseBody) *SubmitVideoAnalysisTaskResponse {
	s.Body = v
	return s
}

type UpdateVideoAnalysisConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s UpdateVideoAnalysisConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoAnalysisConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigRequest) SetAsyncConcurrency(v int32) *UpdateVideoAnalysisConfigRequest {
	s.AsyncConcurrency = &v
	return s
}

type UpdateVideoAnalysisConfigResponseBody struct {
	// example:
	//
	// xx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoAnalysisConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoAnalysisConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetCode(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetHttpStatusCode(v int32) *UpdateVideoAnalysisConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetMessage(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetRequestId(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetSuccess(v bool) *UpdateVideoAnalysisConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateVideoAnalysisConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoAnalysisConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoAnalysisConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoAnalysisConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigResponse) SetHeaders(v map[string]*string) *UpdateVideoAnalysisConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoAnalysisConfigResponse) SetStatusCode(v int32) *UpdateVideoAnalysisConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponse) SetBody(v *UpdateVideoAnalysisConfigResponseBody) *UpdateVideoAnalysisConfigResponse {
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
// 新闻播报-抽取分类获取播报热点
//
// @param request - GenerateBroadcastNewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateBroadcastNewsResponse
func (client *Client) GenerateBroadcastNewsWithOptions(workspaceId *string, request *GenerateBroadcastNewsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateBroadcastNewsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateBroadcastNews"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/GenerateBroadcastNews"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateBroadcastNewsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateBroadcastNewsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 新闻播报-抽取分类获取播报热点
//
// @param request - GenerateBroadcastNewsRequest
//
// @return GenerateBroadcastNewsResponse
func (client *Client) GenerateBroadcastNews(workspaceId *string, request *GenerateBroadcastNewsRequest) (_result *GenerateBroadcastNewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateBroadcastNewsResponse{}
	_body, _err := client.GenerateBroadcastNewsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-标签挖掘-获取示例输出格式
//
// @param tmpReq - GenerateOutputFormatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOutputFormatResponse
func (client *Client) GenerateOutputFormatWithOptions(workspaceId *string, tmpReq *GenerateOutputFormatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateOutputFormatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateOutputFormatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["businessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDescription)) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateOutputFormat"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/generateOutputFormat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateOutputFormatResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateOutputFormatResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-标签挖掘-获取示例输出格式
//
// @param request - GenerateOutputFormatRequest
//
// @return GenerateOutputFormatResponse
func (client *Client) GenerateOutputFormat(workspaceId *string, request *GenerateOutputFormatRequest) (_result *GenerateOutputFormatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateOutputFormatResponse{}
	_body, _err := client.GenerateOutputFormatWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取挖掘分析任务结果
//
// @param request - GetTagMiningAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagMiningAnalysisTaskResponse
func (client *Client) GetTagMiningAnalysisTaskWithOptions(workspaceId *string, request *GetTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTagMiningAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagMiningAnalysisTask"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/getTagMiningAnalysisTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTagMiningAnalysisTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTagMiningAnalysisTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取挖掘分析任务结果
//
// @param request - GetTagMiningAnalysisTaskRequest
//
// @return GetTagMiningAnalysisTaskResponse
func (client *Client) GetTagMiningAnalysisTask(workspaceId *string, request *GetTagMiningAnalysisTaskRequest) (_result *GetTagMiningAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTagMiningAnalysisTaskResponse{}
	_body, _err := client.GetTagMiningAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频理解-获取配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoAnalysisConfigResponse
func (client *Client) GetVideoAnalysisConfigWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVideoAnalysisConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoAnalysisConfig"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/getVideoAnalysisConfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVideoAnalysisConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVideoAnalysisConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 视频理解-获取配置
//
// @return GetVideoAnalysisConfigResponse
func (client *Client) GetVideoAnalysisConfig(workspaceId *string) (_result *GetVideoAnalysisConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoAnalysisConfigResponse{}
	_body, _err := client.GetVideoAnalysisConfigWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-获取视频理解异步任务结果
//
// @param request - GetVideoAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoAnalysisTaskResponse
func (client *Client) GetVideoAnalysisTaskWithOptions(workspaceId *string, request *GetVideoAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVideoAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoAnalysisTask"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/getVideoAnalysisTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVideoAnalysisTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVideoAnalysisTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-获取视频理解异步任务结果
//
// @param request - GetVideoAnalysisTaskRequest
//
// @return GetVideoAnalysisTaskResponse
func (client *Client) GetVideoAnalysisTask(workspaceId *string, request *GetVideoAnalysisTaskRequest) (_result *GetVideoAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoAnalysisTaskResponse{}
	_body, _err := client.GetVideoAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-新闻播报-获取热点话题摘要列表
//
// @param request - ListHotTopicSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotTopicSummariesResponse
func (client *Client) ListHotTopicSummariesWithOptions(workspaceId *string, request *ListHotTopicSummariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHotTopicSummariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.HotTopic)) {
		body["hotTopic"] = request.HotTopic
	}

	if !tea.BoolValue(util.IsUnset(request.HotTopicVersion)) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotTopicSummaries"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/listHotTopicSummaries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListHotTopicSummariesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListHotTopicSummariesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-新闻播报-获取热点话题摘要列表
//
// @param request - ListHotTopicSummariesRequest
//
// @return ListHotTopicSummariesResponse
func (client *Client) ListHotTopicSummaries(workspaceId *string, request *ListHotTopicSummariesRequest) (_result *ListHotTopicSummariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHotTopicSummariesResponse{}
	_body, _err := client.ListHotTopicSummariesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-热点播报-问答
//
// @param tmpReq - RunHotTopicChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicChatResponse
func (client *Client) RunHotTopicChatWithOptions(workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunHotTopicChatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunHotTopicChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GenerateOptions)) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, tea.String("generateOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.HotTopics)) {
		request.HotTopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotTopics, tea.String("hotTopics"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Messages)) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, tea.String("messages"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StepForBroadcastContentConfig)) {
		request.StepForBroadcastContentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForBroadcastContentConfig, tea.String("stepForBroadcastContentConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateOptionsShrink)) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HotTopicVersion)) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HotTopicsShrink)) {
		body["hotTopics"] = request.HotTopicsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCount)) {
		body["imageCount"] = request.ImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.MessagesShrink)) {
		body["messages"] = request.MessagesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCustomPromptTemplate)) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalSessionId)) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.StepForBroadcastContentConfigShrink)) {
		body["stepForBroadcastContentConfig"] = request.StepForBroadcastContentConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunHotTopicChat"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runHotTopicChat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunHotTopicChatResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunHotTopicChatResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-热点播报-问答
//
// @param request - RunHotTopicChatRequest
//
// @return RunHotTopicChatResponse
func (client *Client) RunHotTopicChat(workspaceId *string, request *RunHotTopicChatRequest) (_result *RunHotTopicChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunHotTopicChatResponse{}
	_body, _err := client.RunHotTopicChatWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-热点播报-热点摘要生成
//
// @param tmpReq - RunHotTopicSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicSummaryResponse
func (client *Client) RunHotTopicSummaryWithOptions(workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunHotTopicSummaryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunHotTopicSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StepForCustomSummaryStyleConfig)) {
		request.StepForCustomSummaryStyleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForCustomSummaryStyleConfig, tea.String("stepForCustomSummaryStyleConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TopicIds)) {
		request.TopicIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIds, tea.String("topicIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotTopicVersion)) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.StepForCustomSummaryStyleConfigShrink)) {
		body["stepForCustomSummaryStyleConfig"] = request.StepForCustomSummaryStyleConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TopicIdsShrink)) {
		body["topicIds"] = request.TopicIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunHotTopicSummary"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runHotTopicSummary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunHotTopicSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunHotTopicSummaryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-热点播报-热点摘要生成
//
// @param request - RunHotTopicSummaryRequest
//
// @return RunHotTopicSummaryResponse
func (client *Client) RunHotTopicSummary(workspaceId *string, request *RunHotTopicSummaryRequest) (_result *RunHotTopicSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunHotTopicSummaryResponse{}
	_body, _err := client.RunHotTopicSummaryWithOptions(workspaceId, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunMarketingInformationExtractResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunMarketingInformationExtractResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if !tea.BoolValue(util.IsUnset(request.CustomLimitation)) {
		body["customLimitation"] = request.CustomLimitation
	}

	if !tea.BoolValue(util.IsUnset(request.CustomPrompt)) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.InputExample)) {
		body["inputExample"] = request.InputExample
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputExample)) {
		body["outputExample"] = request.OutputExample
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunMarketingInformationWritingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunMarketingInformationWritingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 轻应用-网络内容审核
//
// @param tmpReq - RunNetworkContentAuditRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNetworkContentAuditResponse
func (client *Client) RunNetworkContentAuditWithOptions(workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunNetworkContentAuditResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunNetworkContentAuditShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["businessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["outputFormat"] = request.OutputFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDescription)) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunNetworkContentAudit"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runNetworkContentAudit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunNetworkContentAuditResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunNetworkContentAuditResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-网络内容审核
//
// @param request - RunNetworkContentAuditRequest
//
// @return RunNetworkContentAuditResponse
func (client *Client) RunNetworkContentAudit(workspaceId *string, request *RunNetworkContentAuditRequest) (_result *RunNetworkContentAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunNetworkContentAuditResponse{}
	_body, _err := client.RunNetworkContentAuditWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 长剧本创作
//
// @param request - RunScriptChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptChatResponse
func (client *Client) RunScriptChatWithOptions(workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunScriptChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunScriptChat"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runScriptChat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunScriptChatResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunScriptChatResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 长剧本创作
//
// @param request - RunScriptChatRequest
//
// @return RunScriptChatResponse
func (client *Client) RunScriptChat(workspaceId *string, request *RunScriptChatRequest) (_result *RunScriptChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptChatResponse{}
	_body, _err := client.RunScriptChatWithOptions(workspaceId, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunScriptContinueResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunScriptContinueResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunScriptPlanningResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunScriptPlanningResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 剧本对话内容的整理
//
// @param request - RunScriptRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptRefineResponse
func (client *Client) RunScriptRefineWithOptions(workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunScriptRefineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunScriptRefine"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runScriptRefine"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunScriptRefineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunScriptRefineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 剧本对话内容的整理
//
// @param request - RunScriptRefineRequest
//
// @return RunScriptRefineResponse
func (client *Client) RunScriptRefine(workspaceId *string, request *RunScriptRefineRequest) (_result *RunScriptRefineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptRefineResponse{}
	_body, _err := client.RunScriptRefineWithOptions(workspaceId, request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ProcessStage)) {
		body["processStage"] = request.ProcessStage
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceMaterialsShrink)) {
		body["referenceMaterials"] = request.ReferenceMaterialsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StyleFeature)) {
		body["styleFeature"] = request.StyleFeature
	}

	if !tea.BoolValue(util.IsUnset(request.UseSearch)) {
		body["useSearch"] = request.UseSearch
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunStyleWritingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunStyleWritingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 轻应用-标签挖掘
//
// @param tmpReq - RunTagMiningAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTagMiningAnalysisResponse
func (client *Client) RunTagMiningAnalysisWithOptions(workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunTagMiningAnalysisResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunTagMiningAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["businessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["outputFormat"] = request.OutputFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDescription)) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunTagMiningAnalysis"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/runTagMiningAnalysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunTagMiningAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunTagMiningAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-标签挖掘
//
// @param request - RunTagMiningAnalysisRequest
//
// @return RunTagMiningAnalysisResponse
func (client *Client) RunTagMiningAnalysis(workspaceId *string, request *RunTagMiningAnalysisRequest) (_result *RunTagMiningAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunTagMiningAnalysisResponse{}
	_body, _err := client.RunTagMiningAnalysisWithOptions(workspaceId, request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.FrameSampleMethod)) {
		request.FrameSampleMethodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FrameSampleMethod, tea.String("frameSampleMethod"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GenerateOptions)) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, tea.String("generateOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TextProcessTasks)) {
		request.TextProcessTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextProcessTasks, tea.String("textProcessTasks"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoRoles)) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, tea.String("videoRoles"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceIdentitySimilarityMinScore)) {
		body["faceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.FrameSampleMethodShrink)) {
		body["frameSampleMethod"] = request.FrameSampleMethodShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateOptionsShrink)) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
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

	if !tea.BoolValue(util.IsUnset(request.SnapshotInterval)) {
		body["snapshotInterval"] = request.SnapshotInterval
	}

	if !tea.BoolValue(util.IsUnset(request.SplitInterval)) {
		body["splitInterval"] = request.SplitInterval
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TextProcessTasksShrink)) {
		body["textProcessTasks"] = request.TextProcessTasksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VideoExtraInfo)) {
		body["videoExtraInfo"] = request.VideoExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelCustomPromptTemplate)) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelId)) {
		body["videoModelId"] = request.VideoModelId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoRolesShrink)) {
		body["videoRoles"] = request.VideoRolesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VideoShotFaceIdentityCount)) {
		body["videoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunVideoAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunVideoAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Summary:
//
// 轻应用-标签挖掘
//
// @param tmpReq - SubmitTagMiningAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTagMiningAnalysisTaskResponse
func (client *Client) SubmitTagMiningAnalysisTaskWithOptions(workspaceId *string, tmpReq *SubmitTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTagMiningAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTagMiningAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Contents)) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, tea.String("contents"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["businessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.ContentsShrink)) {
		body["contents"] = request.ContentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["outputFormat"] = request.OutputFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDescription)) {
		body["taskDescription"] = request.TaskDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTagMiningAnalysisTask"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/submitTagMiningAnalysisTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitTagMiningAnalysisTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitTagMiningAnalysisTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-标签挖掘
//
// @param request - SubmitTagMiningAnalysisTaskRequest
//
// @return SubmitTagMiningAnalysisTaskResponse
func (client *Client) SubmitTagMiningAnalysisTask(workspaceId *string, request *SubmitTagMiningAnalysisTaskRequest) (_result *SubmitTagMiningAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTagMiningAnalysisTaskResponse{}
	_body, _err := client.SubmitTagMiningAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 轻应用-提交视频理解任务
//
// @param tmpReq - SubmitVideoAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoAnalysisTaskResponse
func (client *Client) SubmitVideoAnalysisTaskWithOptions(workspaceId *string, tmpReq *SubmitVideoAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitVideoAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitVideoAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FrameSampleMethod)) {
		request.FrameSampleMethodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FrameSampleMethod, tea.String("frameSampleMethod"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GenerateOptions)) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, tea.String("generateOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TextProcessTasks)) {
		request.TextProcessTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextProcessTasks, tea.String("textProcessTasks"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoRoles)) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, tea.String("videoRoles"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceIdentitySimilarityMinScore)) {
		body["faceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.FrameSampleMethodShrink)) {
		body["frameSampleMethod"] = request.FrameSampleMethodShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateOptionsShrink)) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
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

	if !tea.BoolValue(util.IsUnset(request.SnapshotInterval)) {
		body["snapshotInterval"] = request.SnapshotInterval
	}

	if !tea.BoolValue(util.IsUnset(request.SplitInterval)) {
		body["splitInterval"] = request.SplitInterval
	}

	if !tea.BoolValue(util.IsUnset(request.TextProcessTasksShrink)) {
		body["textProcessTasks"] = request.TextProcessTasksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VideoExtraInfo)) {
		body["videoExtraInfo"] = request.VideoExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelCustomPromptTemplate)) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoModelId)) {
		body["videoModelId"] = request.VideoModelId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoRolesShrink)) {
		body["videoRoles"] = request.VideoRolesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VideoShotFaceIdentityCount)) {
		body["videoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["videoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitVideoAnalysisTask"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/submitVideoAnalysisTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitVideoAnalysisTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitVideoAnalysisTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 轻应用-提交视频理解任务
//
// @param request - SubmitVideoAnalysisTaskRequest
//
// @return SubmitVideoAnalysisTaskResponse
func (client *Client) SubmitVideoAnalysisTask(workspaceId *string, request *SubmitVideoAnalysisTaskRequest) (_result *SubmitVideoAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitVideoAnalysisTaskResponse{}
	_body, _err := client.SubmitVideoAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频理解-更新配置
//
// @param request - UpdateVideoAnalysisConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoAnalysisConfigResponse
func (client *Client) UpdateVideoAnalysisConfigWithOptions(workspaceId *string, request *UpdateVideoAnalysisConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateVideoAnalysisConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncConcurrency)) {
		body["asyncConcurrency"] = request.AsyncConcurrency
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVideoAnalysisConfig"),
		Version:     tea.String("2024-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/updateVideoAnalysisConfig"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateVideoAnalysisConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateVideoAnalysisConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 视频理解-更新配置
//
// @param request - UpdateVideoAnalysisConfigRequest
//
// @return UpdateVideoAnalysisConfigResponse
func (client *Client) UpdateVideoAnalysisConfig(workspaceId *string, request *UpdateVideoAnalysisConfigRequest) (_result *UpdateVideoAnalysisConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVideoAnalysisConfigResponse{}
	_body, _err := client.UpdateVideoAnalysisConfigWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
