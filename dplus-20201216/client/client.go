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

type CreateImageAmazonTaskRequest struct {
	Gif          *bool     `json:"Gif,omitempty" xml:"Gif,omitempty"`
	ImgUrlList   []*string `json:"ImgUrlList,omitempty" xml:"ImgUrlList,omitempty" type:"Repeated"`
	TemplateMode *string   `json:"TemplateMode,omitempty" xml:"TemplateMode,omitempty"`
	TextList     []*string `json:"TextList,omitempty" xml:"TextList,omitempty" type:"Repeated"`
}

func (s CreateImageAmazonTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskRequest) SetGif(v bool) *CreateImageAmazonTaskRequest {
	s.Gif = &v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetImgUrlList(v []*string) *CreateImageAmazonTaskRequest {
	s.ImgUrlList = v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetTemplateMode(v string) *CreateImageAmazonTaskRequest {
	s.TemplateMode = &v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetTextList(v []*string) *CreateImageAmazonTaskRequest {
	s.TextList = v
	return s
}

type CreateImageAmazonTaskShrinkRequest struct {
	Gif              *bool   `json:"Gif,omitempty" xml:"Gif,omitempty"`
	ImgUrlListShrink *string `json:"ImgUrlList,omitempty" xml:"ImgUrlList,omitempty"`
	TemplateMode     *string `json:"TemplateMode,omitempty" xml:"TemplateMode,omitempty"`
	TextListShrink   *string `json:"TextList,omitempty" xml:"TextList,omitempty"`
}

func (s CreateImageAmazonTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskShrinkRequest) SetGif(v bool) *CreateImageAmazonTaskShrinkRequest {
	s.Gif = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetImgUrlListShrink(v string) *CreateImageAmazonTaskShrinkRequest {
	s.ImgUrlListShrink = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetTemplateMode(v string) *CreateImageAmazonTaskShrinkRequest {
	s.TemplateMode = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetTextListShrink(v string) *CreateImageAmazonTaskShrinkRequest {
	s.TextListShrink = &v
	return s
}

type CreateImageAmazonTaskResponseBody struct {
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s CreateImageAmazonTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskResponseBody) SetCode(v int64) *CreateImageAmazonTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetData(v int64) *CreateImageAmazonTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetMessage(v string) *CreateImageAmazonTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetRequestId(v string) *CreateImageAmazonTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetSuccessResponse(v bool) *CreateImageAmazonTaskResponseBody {
	s.SuccessResponse = &v
	return s
}

type CreateImageAmazonTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageAmazonTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageAmazonTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskResponse) SetHeaders(v map[string]*string) *CreateImageAmazonTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageAmazonTaskResponse) SetBody(v *CreateImageAmazonTaskResponseBody) *CreateImageAmazonTaskResponse {
	s.Body = v
	return s
}

type GetTaskResultRequest struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) SetTaskId(v int64) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetTaskResultResponseBody struct {
	Code    *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) SetCode(v int64) *GetTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResultResponseBody) SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetMessage(v string) *GetTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccessResponse(v bool) *GetTaskResultResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTaskResultResponseBodyData struct {
	Resuslt    *string `json:"Resuslt,omitempty" xml:"Resuslt,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyData) SetResuslt(v string) *GetTaskResultResponseBodyData {
	s.Resuslt = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetStatus(v int64) *GetTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetStatusName(v string) *GetTaskResultResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskId(v int64) *GetTaskResultResponseBodyData {
	s.TaskId = &v
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

type GetTaskStatusRequest struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetTaskId(v int64) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	Code    *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetCode(v int64) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetData(v *GetTaskStatusResponseBodyData) *GetTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccessResponse(v bool) *GetTaskStatusResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTaskStatusResponseBodyData struct {
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBodyData) SetStatus(v int64) *GetTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBodyData) SetStatusName(v string) *GetTaskStatusResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *GetTaskStatusResponseBodyData) SetTaskId(v int64) *GetTaskStatusResponseBodyData {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateImageAmazonTaskWithOptions(tmpReq *CreateImageAmazonTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageAmazonTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateImageAmazonTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImgUrlList)) {
		request.ImgUrlListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImgUrlList, tea.String("ImgUrlList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TextList)) {
		request.TextListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextList, tea.String("TextList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["Gif"] = request.Gif
	query["ImgUrlList"] = request.ImgUrlListShrink
	query["TemplateMode"] = request.TemplateMode
	query["TextList"] = request.TextListShrink
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageAmazonTask"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageAmazonTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageAmazonTask(request *CreateImageAmazonTaskRequest) (_result *CreateImageAmazonTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageAmazonTaskResponse{}
	_body, _err := client.CreateImageAmazonTaskWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskResult"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
