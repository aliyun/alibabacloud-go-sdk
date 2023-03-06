// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchJobCheckRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s BatchJobCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckRequest) GoString() string {
	return s.String()
}

func (s *BatchJobCheckRequest) SetRegionId(v string) *BatchJobCheckRequest {
	s.RegionId = &v
	return s
}

func (s *BatchJobCheckRequest) SetSubmitId(v string) *BatchJobCheckRequest {
	s.SubmitId = &v
	return s
}

type BatchJobCheckResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchJobCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJobCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBody) SetCode(v int32) *BatchJobCheckResponseBody {
	s.Code = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetData(v *BatchJobCheckResponseBodyData) *BatchJobCheckResponseBody {
	s.Data = v
	return s
}

func (s *BatchJobCheckResponseBody) SetErrCode(v string) *BatchJobCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetMessage(v string) *BatchJobCheckResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetRequestId(v string) *BatchJobCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetSuccess(v bool) *BatchJobCheckResponseBody {
	s.Success = &v
	return s
}

type BatchJobCheckResponseBodyData struct {
	ConfigId    *string                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ErrTaskList []*BatchJobCheckResponseBodyDataErrTaskList `json:"ErrTaskList,omitempty" xml:"ErrTaskList,omitempty" type:"Repeated"`
	FailedCount *int32                                      `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FinishCount *int32                                      `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	FolderId    *string                                     `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	TaskCount   *int32                                      `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TaskStatus  *string                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s BatchJobCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyData) SetConfigId(v string) *BatchJobCheckResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetErrTaskList(v []*BatchJobCheckResponseBodyDataErrTaskList) *BatchJobCheckResponseBodyData {
	s.ErrTaskList = v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFailedCount(v int32) *BatchJobCheckResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFinishCount(v int32) *BatchJobCheckResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFolderId(v string) *BatchJobCheckResponseBodyData {
	s.FolderId = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetTaskCount(v int32) *BatchJobCheckResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetTaskStatus(v string) *BatchJobCheckResponseBodyData {
	s.TaskStatus = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskList struct {
	ProductList []*BatchJobCheckResponseBodyDataErrTaskListProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	UserId      *int64                                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskList) SetProductList(v []*BatchJobCheckResponseBodyDataErrTaskListProductList) *BatchJobCheckResponseBodyDataErrTaskList {
	s.ProductList = v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskList) SetUserId(v int64) *BatchJobCheckResponseBodyDataErrTaskList {
	s.UserId = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskListProductList struct {
	LogList     []*BatchJobCheckResponseBodyDataErrTaskListProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	ProductCode *string                                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductList) SetLogList(v []*BatchJobCheckResponseBodyDataErrTaskListProductListLogList) *BatchJobCheckResponseBodyDataErrTaskListProductList {
	s.LogList = v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductList) SetProductCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductList {
	s.ProductCode = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskListProductListLogList struct {
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LogCode             *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogStoreNamePattern *string `json:"LogStoreNamePattern,omitempty" xml:"LogStoreNamePattern,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProjectNamePattern  *string `json:"ProjectNamePattern,omitempty" xml:"ProjectNamePattern,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductListLogList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetErrorCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ErrorCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetLogCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.LogCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetLogStoreNamePattern(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.LogStoreNamePattern = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetProductCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ProductCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetProjectNamePattern(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ProjectNamePattern = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetRegionCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.RegionCode = &v
	return s
}

type BatchJobCheckResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJobCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJobCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponse) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponse) SetHeaders(v map[string]*string) *BatchJobCheckResponse {
	s.Headers = v
	return s
}

func (s *BatchJobCheckResponse) SetStatusCode(v int32) *BatchJobCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJobCheckResponse) SetBody(v *BatchJobCheckResponseBody) *BatchJobCheckResponse {
	s.Body = v
	return s
}

type BatchJobSubmitRequest struct {
	JsonConfig *string `json:"JsonConfig,omitempty" xml:"JsonConfig,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchJobSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitRequest) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitRequest) SetJsonConfig(v string) *BatchJobSubmitRequest {
	s.JsonConfig = &v
	return s
}

func (s *BatchJobSubmitRequest) SetRegionId(v string) *BatchJobSubmitRequest {
	s.RegionId = &v
	return s
}

type BatchJobSubmitResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchJobSubmitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJobSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBody) SetCode(v int32) *BatchJobSubmitResponseBody {
	s.Code = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetData(v *BatchJobSubmitResponseBodyData) *BatchJobSubmitResponseBody {
	s.Data = v
	return s
}

func (s *BatchJobSubmitResponseBody) SetErrCode(v string) *BatchJobSubmitResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetMessage(v string) *BatchJobSubmitResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetRequestId(v string) *BatchJobSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetSuccess(v bool) *BatchJobSubmitResponseBody {
	s.Success = &v
	return s
}

type BatchJobSubmitResponseBodyData struct {
	ConfigId   *string                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigList []*BatchJobSubmitResponseBodyDataConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	SubmitId   *string                                     `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	TaskCount  *int32                                      `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s BatchJobSubmitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyData) SetConfigId(v string) *BatchJobSubmitResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetConfigList(v []*BatchJobSubmitResponseBodyDataConfigList) *BatchJobSubmitResponseBodyData {
	s.ConfigList = v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetSubmitId(v string) *BatchJobSubmitResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetTaskCount(v int32) *BatchJobSubmitResponseBodyData {
	s.TaskCount = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigList struct {
	ProductList []*BatchJobSubmitResponseBodyDataConfigListProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	UserId      *int64                                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigList) SetProductList(v []*BatchJobSubmitResponseBodyDataConfigListProductList) *BatchJobSubmitResponseBodyDataConfigList {
	s.ProductList = v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigList) SetUserId(v int64) *BatchJobSubmitResponseBodyDataConfigList {
	s.UserId = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigListProductList struct {
	LogList     []*BatchJobSubmitResponseBodyDataConfigListProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	ProductCode *string                                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigListProductList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigListProductList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductList) SetLogList(v []*BatchJobSubmitResponseBodyDataConfigListProductListLogList) *BatchJobSubmitResponseBodyDataConfigListProductList {
	s.LogList = v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductList) SetProductCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductList {
	s.ProductCode = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigListProductListLogList struct {
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LogCode             *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogStoreNamePattern *string `json:"LogStoreNamePattern,omitempty" xml:"LogStoreNamePattern,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProjectNamePattern  *string `json:"ProjectNamePattern,omitempty" xml:"ProjectNamePattern,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigListProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigListProductListLogList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetErrorCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ErrorCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetLogCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.LogCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetLogStoreNamePattern(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.LogStoreNamePattern = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetProductCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ProductCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetProjectNamePattern(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ProjectNamePattern = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetRegionCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.RegionCode = &v
	return s
}

type BatchJobSubmitResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJobSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJobSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponse) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponse) SetHeaders(v map[string]*string) *BatchJobSubmitResponse {
	s.Headers = v
	return s
}

func (s *BatchJobSubmitResponse) SetStatusCode(v int32) *BatchJobSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJobSubmitResponse) SetBody(v *BatchJobSubmitResponseBody) *BatchJobSubmitResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ReceiveUid  *int64  `json:"ReceiveUid,omitempty" xml:"ReceiveUid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetChannelType(v int32) *SendMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendMessageRequest) SetReceiveUid(v int64) *SendMessageRequest {
	s.ReceiveUid = &v
	return s
}

func (s *SendMessageRequest) SetRegionId(v string) *SendMessageRequest {
	s.RegionId = &v
	return s
}

type SendMessageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCode(v int32) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetData(v bool) *SendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageResponseBody) SetErrCode(v string) *SendMessageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) SetSuccess(v bool) *SendMessageResponseBody {
	s.Success = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloud-siem"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchJobCheckWithOptions(request *BatchJobCheckRequest, runtime *util.RuntimeOptions) (_result *BatchJobCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitId)) {
		body["SubmitId"] = request.SubmitId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJobCheck"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJobCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJobCheck(request *BatchJobCheckRequest) (_result *BatchJobCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJobCheckResponse{}
	_body, _err := client.BatchJobCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchJobSubmitWithOptions(request *BatchJobSubmitRequest, runtime *util.RuntimeOptions) (_result *BatchJobSubmitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonConfig)) {
		body["JsonConfig"] = request.JsonConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJobSubmit"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJobSubmitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJobSubmit(request *BatchJobSubmitRequest) (_result *BatchJobSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJobSubmitResponse{}
	_body, _err := client.BatchJobSubmitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveUid)) {
		body["ReceiveUid"] = request.ReceiveUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
