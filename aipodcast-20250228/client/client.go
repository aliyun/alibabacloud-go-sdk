// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PodcastTaskResultQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddri
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskResultQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskResultQueryRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryRequest) SetTaskId(v string) *PodcastTaskResultQueryRequest {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskResultQueryRequest) SetWorkspaceId(v string) *PodcastTaskResultQueryRequest {
	s.WorkspaceId = &v
	return s
}

type PodcastTaskResultQueryResponseBody struct {
	// example:
	//
	// "success"
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *PodcastTaskResultQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// "success"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C38F034D-7F36-531C-95AC-0C752F80E840
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PodcastTaskResultQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskResultQueryResponseBody) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponseBody) SetCode(v string) *PodcastTaskResultQueryResponseBody {
	s.Code = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetData(v *PodcastTaskResultQueryResponseBodyData) *PodcastTaskResultQueryResponseBody {
	s.Data = v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetHttpStatusCode(v string) *PodcastTaskResultQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetMessage(v string) *PodcastTaskResultQueryResponseBody {
	s.Message = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetRequestId(v string) *PodcastTaskResultQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetSuccess(v bool) *PodcastTaskResultQueryResponseBody {
	s.Success = &v
	return s
}

type PodcastTaskResultQueryResponseBodyData struct {
	// example:
	//
	// {
	//
	//             "audio": "http://note-ai-file.oss-cn-beijing.aliyuncs.com/202503241702148295/audio.mp3?OSSAccessKeyId=LTAI5tPLWJfJHNkZbfnQv245&Expires=1742810788&Signature=b5p83nh443Gr7foqdvgrI4%2FKZVM%3D",
	//
	//             "script": "http://note-ai-file.oss-cn-beijing.aliyuncs.com/202503241702148295/script.txt?OSSAccessKeyId=LTAI5tPLWJfJHNkZbfnQv245&Expires=1742810622&Signature=TBBdikHzOWW3YqDw3sNMTXiMo6A%3D"
	//
	// }
	ResultUrl interface{} `json:"resultUrl,omitempty" xml:"resultUrl,omitempty"`
	Script    *string     `json:"script,omitempty" xml:"script,omitempty"`
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCEEDED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s PodcastTaskResultQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskResultQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponseBodyData) SetResultUrl(v interface{}) *PodcastTaskResultQueryResponseBodyData {
	s.ResultUrl = v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetScript(v string) *PodcastTaskResultQueryResponseBodyData {
	s.Script = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetTaskId(v string) *PodcastTaskResultQueryResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetTaskStatus(v string) *PodcastTaskResultQueryResponseBodyData {
	s.TaskStatus = &v
	return s
}

type PodcastTaskResultQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PodcastTaskResultQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PodcastTaskResultQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskResultQueryResponse) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponse) SetHeaders(v map[string]*string) *PodcastTaskResultQueryResponse {
	s.Headers = v
	return s
}

func (s *PodcastTaskResultQueryResponse) SetStatusCode(v int32) *PodcastTaskResultQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *PodcastTaskResultQueryResponse) SetBody(v *PodcastTaskResultQueryResponseBody) *PodcastTaskResultQueryResponse {
	s.Body = v
	return s
}

type PodcastTaskSubmitRequest struct {
	// example:
	//
	// 2
	Counts   *int32    `json:"counts,omitempty" xml:"counts,omitempty"`
	FileUrls []*string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty" type:"Repeated"`
	Text     *string   `json:"text,omitempty" xml:"text,omitempty"`
	Topic    *string   `json:"topic,omitempty" xml:"topic,omitempty"`
	Voices   []*string `json:"voices,omitempty" xml:"voices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskSubmitRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitRequest) SetCounts(v int32) *PodcastTaskSubmitRequest {
	s.Counts = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetFileUrls(v []*string) *PodcastTaskSubmitRequest {
	s.FileUrls = v
	return s
}

func (s *PodcastTaskSubmitRequest) SetText(v string) *PodcastTaskSubmitRequest {
	s.Text = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetTopic(v string) *PodcastTaskSubmitRequest {
	s.Topic = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetVoices(v []*string) *PodcastTaskSubmitRequest {
	s.Voices = v
	return s
}

func (s *PodcastTaskSubmitRequest) SetWorkspaceId(v string) *PodcastTaskSubmitRequest {
	s.WorkspaceId = &v
	return s
}

type PodcastTaskSubmitShrinkRequest struct {
	// example:
	//
	// 2
	Counts         *int32  `json:"counts,omitempty" xml:"counts,omitempty"`
	FileUrlsShrink *string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty"`
	Text           *string `json:"text,omitempty" xml:"text,omitempty"`
	Topic          *string `json:"topic,omitempty" xml:"topic,omitempty"`
	VoicesShrink   *string `json:"voices,omitempty" xml:"voices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskSubmitShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskSubmitShrinkRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitShrinkRequest) SetCounts(v int32) *PodcastTaskSubmitShrinkRequest {
	s.Counts = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetFileUrlsShrink(v string) *PodcastTaskSubmitShrinkRequest {
	s.FileUrlsShrink = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetText(v string) *PodcastTaskSubmitShrinkRequest {
	s.Text = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetTopic(v string) *PodcastTaskSubmitShrinkRequest {
	s.Topic = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetVoicesShrink(v string) *PodcastTaskSubmitShrinkRequest {
	s.VoicesShrink = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetWorkspaceId(v string) *PodcastTaskSubmitShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type PodcastTaskSubmitResponseBody struct {
	// example:
	//
	// "success"
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *PodcastTaskSubmitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// "success"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9CE5B91A-6E6B-55FB-A1AF-037DF01C84B3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PodcastTaskSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponseBody) SetCode(v string) *PodcastTaskSubmitResponseBody {
	s.Code = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetData(v *PodcastTaskSubmitResponseBodyData) *PodcastTaskSubmitResponseBody {
	s.Data = v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetHttpStatusCode(v string) *PodcastTaskSubmitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetMessage(v string) *PodcastTaskSubmitResponseBody {
	s.Message = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetRequestId(v string) *PodcastTaskSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *PodcastTaskSubmitResponseBody) SetSuccess(v bool) *PodcastTaskSubmitResponseBody {
	s.Success = &v
	return s
}

type PodcastTaskSubmitResponseBodyData struct {
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCEEDED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s PodcastTaskSubmitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskSubmitResponseBodyData) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponseBodyData) SetTaskId(v string) *PodcastTaskSubmitResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskSubmitResponseBodyData) SetTaskStatus(v string) *PodcastTaskSubmitResponseBodyData {
	s.TaskStatus = &v
	return s
}

type PodcastTaskSubmitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PodcastTaskSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PodcastTaskSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s PodcastTaskSubmitResponse) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponse) SetHeaders(v map[string]*string) *PodcastTaskSubmitResponse {
	s.Headers = v
	return s
}

func (s *PodcastTaskSubmitResponse) SetStatusCode(v int32) *PodcastTaskSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *PodcastTaskSubmitResponse) SetBody(v *PodcastTaskSubmitResponseBody) *PodcastTaskSubmitResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aipodcast"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// ai播客生成任务结果查询
//
// @param request - PodcastTaskResultQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PodcastTaskResultQueryResponse
func (client *Client) PodcastTaskResultQueryWithOptions(request *PodcastTaskResultQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PodcastTaskResultQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PodcastTaskResultQuery"),
		Version:     tea.String("2025-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/podcast/task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PodcastTaskResultQueryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PodcastTaskResultQueryResponse{}
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
// ai播客生成任务结果查询
//
// @param request - PodcastTaskResultQueryRequest
//
// @return PodcastTaskResultQueryResponse
func (client *Client) PodcastTaskResultQuery(request *PodcastTaskResultQueryRequest) (_result *PodcastTaskResultQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PodcastTaskResultQueryResponse{}
	_body, _err := client.PodcastTaskResultQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ai播客生成任务提交
//
// @param tmpReq - PodcastTaskSubmitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PodcastTaskSubmitResponse
func (client *Client) PodcastTaskSubmitWithOptions(tmpReq *PodcastTaskSubmitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PodcastTaskSubmitResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PodcastTaskSubmitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FileUrls)) {
		request.FileUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileUrls, tea.String("fileUrls"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Voices)) {
		request.VoicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Voices, tea.String("voices"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Counts)) {
		body["counts"] = request.Counts
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrlsShrink)) {
		body["fileUrls"] = request.FileUrlsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.VoicesShrink)) {
		body["voices"] = request.VoicesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PodcastTaskSubmit"),
		Version:     tea.String("2025-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/podcast/task/submit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PodcastTaskSubmitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PodcastTaskSubmitResponse{}
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
// ai播客生成任务提交
//
// @param request - PodcastTaskSubmitRequest
//
// @return PodcastTaskSubmitResponse
func (client *Client) PodcastTaskSubmit(request *PodcastTaskSubmitRequest) (_result *PodcastTaskSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PodcastTaskSubmitResponse{}
	_body, _err := client.PodcastTaskSubmitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
