// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTaskResultRequest struct {
	Debug  *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) SetDebug(v bool) *GetTaskResultRequest {
	s.Debug = &v
	return s
}

func (s *GetTaskResultRequest) SetTaskId(v string) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetTaskResultResponseBody struct {
	BizDuration *int32                           `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *GetTaskResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	SolveTime   *int64                           `json:"SolveTime,omitempty" xml:"SolveTime,omitempty"`
	StatusCode  *int64                           `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusText  *string                          `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	TaskId      *string                          `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) SetBizDuration(v int32) *GetTaskResultResponseBody {
	s.BizDuration = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetResult(v *GetTaskResultResponseBodyResult) *GetTaskResultResponseBody {
	s.Result = v
	return s
}

func (s *GetTaskResultResponseBody) SetSolveTime(v int64) *GetTaskResultResponseBody {
	s.SolveTime = &v
	return s
}

func (s *GetTaskResultResponseBody) SetStatusCode(v int64) *GetTaskResultResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResultResponseBody) SetStatusText(v string) *GetTaskResultResponseBody {
	s.StatusText = &v
	return s
}

func (s *GetTaskResultResponseBody) SetTaskId(v string) *GetTaskResultResponseBody {
	s.TaskId = &v
	return s
}

type GetTaskResultResponseBodyResult struct {
	Paragraphs []*GetTaskResultResponseBodyResultParagraphs `json:"Paragraphs,omitempty" xml:"Paragraphs,omitempty" type:"Repeated"`
	Sentences  []*GetTaskResultResponseBodyResultSentences  `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
	Words      []*GetTaskResultResponseBodyResultWords      `json:"Words,omitempty" xml:"Words,omitempty" type:"Repeated"`
}

func (s GetTaskResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyResult) SetParagraphs(v []*GetTaskResultResponseBodyResultParagraphs) *GetTaskResultResponseBodyResult {
	s.Paragraphs = v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetSentences(v []*GetTaskResultResponseBodyResultSentences) *GetTaskResultResponseBodyResult {
	s.Sentences = v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetWords(v []*GetTaskResultResponseBodyResultWords) *GetTaskResultResponseBodyResult {
	s.Words = v
	return s
}

type GetTaskResultResponseBodyResultParagraphs struct {
	BeginTIme *int64  `json:"BeginTIme,omitempty" xml:"BeginTIme,omitempty"`
	ChannelId *int64  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetTaskResultResponseBodyResultParagraphs) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyResultParagraphs) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyResultParagraphs) SetBeginTIme(v int64) *GetTaskResultResponseBodyResultParagraphs {
	s.BeginTIme = &v
	return s
}

func (s *GetTaskResultResponseBodyResultParagraphs) SetChannelId(v int64) *GetTaskResultResponseBodyResultParagraphs {
	s.ChannelId = &v
	return s
}

func (s *GetTaskResultResponseBodyResultParagraphs) SetEndTime(v int64) *GetTaskResultResponseBodyResultParagraphs {
	s.EndTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResultParagraphs) SetText(v string) *GetTaskResultResponseBodyResultParagraphs {
	s.Text = &v
	return s
}

type GetTaskResultResponseBodyResultSentences struct {
	BeginTime       *int64   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ChannelId       *int64   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EmotionValue    *float32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	EndTime         *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SilenceDuration *int32   `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeakerId       *string  `json:"SpeakerId,omitempty" xml:"SpeakerId,omitempty"`
	SpeechRate      *int32   `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Text            *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetTaskResultResponseBodyResultSentences) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyResultSentences) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyResultSentences) SetBeginTime(v int64) *GetTaskResultResponseBodyResultSentences {
	s.BeginTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetChannelId(v int64) *GetTaskResultResponseBodyResultSentences {
	s.ChannelId = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetEmotionValue(v float32) *GetTaskResultResponseBodyResultSentences {
	s.EmotionValue = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetEndTime(v int64) *GetTaskResultResponseBodyResultSentences {
	s.EndTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetSilenceDuration(v int32) *GetTaskResultResponseBodyResultSentences {
	s.SilenceDuration = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetSpeakerId(v string) *GetTaskResultResponseBodyResultSentences {
	s.SpeakerId = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetSpeechRate(v int32) *GetTaskResultResponseBodyResultSentences {
	s.SpeechRate = &v
	return s
}

func (s *GetTaskResultResponseBodyResultSentences) SetText(v string) *GetTaskResultResponseBodyResultSentences {
	s.Text = &v
	return s
}

type GetTaskResultResponseBodyResultWords struct {
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ChannelId *int64  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Word      *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s GetTaskResultResponseBodyResultWords) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyResultWords) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyResultWords) SetBeginTime(v int64) *GetTaskResultResponseBodyResultWords {
	s.BeginTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResultWords) SetChannelId(v int64) *GetTaskResultResponseBodyResultWords {
	s.ChannelId = &v
	return s
}

func (s *GetTaskResultResponseBodyResultWords) SetEndTime(v int64) *GetTaskResultResponseBodyResultWords {
	s.EndTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResultWords) SetWord(v string) *GetTaskResultResponseBodyResultWords {
	s.Word = &v
	return s
}

type GetTaskResultResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponse) SetHeaders(v map[string]*string) *GetTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResultResponse) SetBody(v *GetTaskResultResponseBody) *GetTaskResultResponse {
	s.Body = v
	return s
}

type SubmitTaskRequest struct {
	Debug *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Task  *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s SubmitTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTaskRequest) SetDebug(v bool) *SubmitTaskRequest {
	s.Debug = &v
	return s
}

func (s *SubmitTaskRequest) SetTask(v string) *SubmitTaskRequest {
	s.Task = &v
	return s
}

type SubmitTaskResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCode *int64  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusText *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTaskResponseBody) SetRequestId(v string) *SubmitTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTaskResponseBody) SetStatusCode(v int64) *SubmitTaskResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitTaskResponseBody) SetStatusText(v string) *SubmitTaskResponseBody {
	s.StatusText = &v
	return s
}

func (s *SubmitTaskResponseBody) SetTaskId(v string) *SubmitTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTaskResponse) SetHeaders(v map[string]*string) *SubmitTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTaskResponse) SetBody(v *SubmitTaskResponseBody) *SubmitTaskResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("speechfiletranscriberlite"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetTaskResultWithOptions(request *GetTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskResult"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskResult(request *GetTaskResultRequest) (_result *GetTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResultResponse{}
	_body, _err := client.GetTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTaskWithOptions(request *SubmitTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		query["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		query["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTask"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTask(request *SubmitTaskRequest) (_result *SubmitTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTaskResponse{}
	_body, _err := client.SubmitTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
