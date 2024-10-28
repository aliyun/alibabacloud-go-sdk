// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RecognizeAudioRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://xx.com/abc.mp3
	AudioFileUrl *string `json:"AudioFileUrl,omitempty" xml:"AudioFileUrl,omitempty"`
	// example:
	//
	// https://abc.edu.org.cn/en/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableCallBack *bool `json:"EnableCallBack,omitempty" xml:"EnableCallBack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 433c1361-0f6e-48fc-ad51
	OuterBizId *string `json:"OuterBizId,omitempty" xml:"OuterBizId,omitempty"`
}

func (s RecognizeAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAudioRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAudioRequest) SetAudioFileUrl(v string) *RecognizeAudioRequest {
	s.AudioFileUrl = &v
	return s
}

func (s *RecognizeAudioRequest) SetCallbackUrl(v string) *RecognizeAudioRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RecognizeAudioRequest) SetEnableCallBack(v bool) *RecognizeAudioRequest {
	s.EnableCallBack = &v
	return s
}

func (s *RecognizeAudioRequest) SetOuterBizId(v string) *RecognizeAudioRequest {
	s.OuterBizId = &v
	return s
}

type RecognizeAudioResponseBody struct {
	ItemList []*RecognizeAudioResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 38CD0881-BC7B-5ADB-A3EB-FF813927D09A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3ab5c082-2c0e-4f39-b12f-057dd5e373ff
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecognizeAudioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAudioResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAudioResponseBody) SetItemList(v []*RecognizeAudioResponseBodyItemList) *RecognizeAudioResponseBody {
	s.ItemList = v
	return s
}

func (s *RecognizeAudioResponseBody) SetRequestId(v string) *RecognizeAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeAudioResponseBody) SetTaskId(v string) *RecognizeAudioResponseBody {
	s.TaskId = &v
	return s
}

type RecognizeAudioResponseBodyItemList struct {
	// example:
	//
	// 0
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1230
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// How are you
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeAudioResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAudioResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *RecognizeAudioResponseBodyItemList) SetBeginTime(v string) *RecognizeAudioResponseBodyItemList {
	s.BeginTime = &v
	return s
}

func (s *RecognizeAudioResponseBodyItemList) SetEndTime(v string) *RecognizeAudioResponseBodyItemList {
	s.EndTime = &v
	return s
}

func (s *RecognizeAudioResponseBodyItemList) SetText(v string) *RecognizeAudioResponseBodyItemList {
	s.Text = &v
	return s
}

type RecognizeAudioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeAudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAudioResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAudioResponse) SetHeaders(v map[string]*string) *RecognizeAudioResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAudioResponse) SetStatusCode(v int32) *RecognizeAudioResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAudioResponse) SetBody(v *RecognizeAudioResponseBody) *RecognizeAudioResponse {
	s.Body = v
	return s
}

type SubmitEvaluationTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://xx.com/abc.mp3
	AudioUrl *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://abc.edu.org.cn/en/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// As flame of fire we gather, as skyful of stars we scatter.
	MaterialText *string `json:"MaterialText,omitempty" xml:"MaterialText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 433c1361-0f6e-48fc-ad51
	OuterBizId *string `json:"OuterBizId,omitempty" xml:"OuterBizId,omitempty"`
	// This parameter is required.
	SuggestedAnswer *string `json:"SuggestedAnswer,omitempty" xml:"SuggestedAnswer,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EN_TO_ZH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitEvaluationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitEvaluationTaskRequest) SetAudioUrl(v string) *SubmitEvaluationTaskRequest {
	s.AudioUrl = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetCallbackUrl(v string) *SubmitEvaluationTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetMaterialText(v string) *SubmitEvaluationTaskRequest {
	s.MaterialText = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetOuterBizId(v string) *SubmitEvaluationTaskRequest {
	s.OuterBizId = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetSuggestedAnswer(v string) *SubmitEvaluationTaskRequest {
	s.SuggestedAnswer = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetText(v string) *SubmitEvaluationTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitEvaluationTaskRequest) SetType(v string) *SubmitEvaluationTaskRequest {
	s.Type = &v
	return s
}

type SubmitEvaluationTaskResponseBody struct {
	// example:
	//
	// BA3BB7D0-E098-5F0C-AF25-0BEFAEE7D1F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8fb75c55-98b5-4b82-ae67-5dbb9a0646cc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitEvaluationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEvaluationTaskResponseBody) SetRequestId(v string) *SubmitEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEvaluationTaskResponseBody) SetTaskId(v string) *SubmitEvaluationTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitEvaluationTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEvaluationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitEvaluationTaskResponse) SetHeaders(v map[string]*string) *SubmitEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitEvaluationTaskResponse) SetStatusCode(v int32) *SubmitEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEvaluationTaskResponse) SetBody(v *SubmitEvaluationTaskResponseBody) *SubmitEvaluationTaskResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("eduinterpreting"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 英语口译语音文件识别成英文内容
//
// @param request - RecognizeAudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeAudioResponse
func (client *Client) RecognizeAudioWithOptions(request *RecognizeAudioRequest, runtime *util.RuntimeOptions) (_result *RecognizeAudioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableCallBack)) {
		query["EnableCallBack"] = request.EnableCallBack
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioFileUrl)) {
		body["AudioFileUrl"] = request.AudioFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBizId)) {
		body["OuterBizId"] = request.OuterBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAudio"),
		Version:     tea.String("2024-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAudioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 英语口译语音文件识别成英文内容
//
// @param request - RecognizeAudioRequest
//
// @return RecognizeAudioResponse
func (client *Client) RecognizeAudio(request *RecognizeAudioRequest) (_result *RecognizeAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAudioResponse{}
	_body, _err := client.RecognizeAudioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 口译评测任务
//
// @param request - SubmitEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEvaluationTaskResponse
func (client *Client) SubmitEvaluationTaskWithOptions(request *SubmitEvaluationTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitEvaluationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioUrl)) {
		body["AudioUrl"] = request.AudioUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialText)) {
		body["MaterialText"] = request.MaterialText
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBizId)) {
		body["OuterBizId"] = request.OuterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.SuggestedAnswer)) {
		body["SuggestedAnswer"] = request.SuggestedAnswer
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitEvaluationTask"),
		Version:     tea.String("2024-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 口译评测任务
//
// @param request - SubmitEvaluationTaskRequest
//
// @return SubmitEvaluationTaskResponse
func (client *Client) SubmitEvaluationTask(request *SubmitEvaluationTaskRequest) (_result *SubmitEvaluationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitEvaluationTaskResponse{}
	_body, _err := client.SubmitEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
