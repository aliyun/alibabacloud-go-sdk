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

type DataRateLimitTestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DataRateLimitTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataRateLimitTestResponseBody) GoString() string {
	return s.String()
}

func (s *DataRateLimitTestResponseBody) SetRequestId(v string) *DataRateLimitTestResponseBody {
	s.RequestId = &v
	return s
}

type DataRateLimitTestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DataRateLimitTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DataRateLimitTestResponse) String() string {
	return tea.Prettify(s)
}

func (s DataRateLimitTestResponse) GoString() string {
	return s.String()
}

func (s *DataRateLimitTestResponse) SetHeaders(v map[string]*string) *DataRateLimitTestResponse {
	s.Headers = v
	return s
}

func (s *DataRateLimitTestResponse) SetStatusCode(v int32) *DataRateLimitTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DataRateLimitTestResponse) SetBody(v *DataRateLimitTestResponseBody) *DataRateLimitTestResponse {
	s.Body = v
	return s
}

type NormalRpcHsfApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NormalRpcHsfApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NormalRpcHsfApiResponseBody) GoString() string {
	return s.String()
}

func (s *NormalRpcHsfApiResponseBody) SetRequestId(v string) *NormalRpcHsfApiResponseBody {
	s.RequestId = &v
	return s
}

type NormalRpcHsfApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *NormalRpcHsfApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NormalRpcHsfApiResponse) String() string {
	return tea.Prettify(s)
}

func (s NormalRpcHsfApiResponse) GoString() string {
	return s.String()
}

func (s *NormalRpcHsfApiResponse) SetHeaders(v map[string]*string) *NormalRpcHsfApiResponse {
	s.Headers = v
	return s
}

func (s *NormalRpcHsfApiResponse) SetStatusCode(v int32) *NormalRpcHsfApiResponse {
	s.StatusCode = &v
	return s
}

func (s *NormalRpcHsfApiResponse) SetBody(v *NormalRpcHsfApiResponseBody) *NormalRpcHsfApiResponse {
	s.Body = v
	return s
}

type NormalRpcHttpApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NormalRpcHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NormalRpcHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *NormalRpcHttpApiResponseBody) SetRequestId(v string) *NormalRpcHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type NormalRpcHttpApiResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *NormalRpcHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NormalRpcHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s NormalRpcHttpApiResponse) GoString() string {
	return s.String()
}

func (s *NormalRpcHttpApiResponse) SetHeaders(v map[string]*string) *NormalRpcHttpApiResponse {
	s.Headers = v
	return s
}

func (s *NormalRpcHttpApiResponse) SetStatusCode(v int32) *NormalRpcHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *NormalRpcHttpApiResponse) SetBody(v *NormalRpcHttpApiResponseBody) *NormalRpcHttpApiResponse {
	s.Body = v
	return s
}

type RpcDataUploadRequest struct {
	LargeParam *string `json:"largeParam,omitempty" xml:"largeParam,omitempty"`
	Query1     *string `json:"query1,omitempty" xml:"query1,omitempty"`
	Query2     *int64  `json:"query2,omitempty" xml:"query2,omitempty"`
}

func (s RpcDataUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadRequest) GoString() string {
	return s.String()
}

func (s *RpcDataUploadRequest) SetLargeParam(v string) *RpcDataUploadRequest {
	s.LargeParam = &v
	return s
}

func (s *RpcDataUploadRequest) SetQuery1(v string) *RpcDataUploadRequest {
	s.Query1 = &v
	return s
}

func (s *RpcDataUploadRequest) SetQuery2(v int64) *RpcDataUploadRequest {
	s.Query2 = &v
	return s
}

type RpcDataUploadResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Headers    map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
	Params     map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Speed      *string                `json:"speed,omitempty" xml:"speed,omitempty"`
	TotalBytes *int64                 `json:"totalBytes,omitempty" xml:"totalBytes,omitempty"`
	TotalTime  *int64                 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	Url        *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RpcDataUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadResponseBody) GoString() string {
	return s.String()
}

func (s *RpcDataUploadResponseBody) SetRequestId(v string) *RpcDataUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *RpcDataUploadResponseBody) SetHeaders(v map[string]interface{}) *RpcDataUploadResponseBody {
	s.Headers = v
	return s
}

func (s *RpcDataUploadResponseBody) SetParams(v map[string]interface{}) *RpcDataUploadResponseBody {
	s.Params = v
	return s
}

func (s *RpcDataUploadResponseBody) SetSpeed(v string) *RpcDataUploadResponseBody {
	s.Speed = &v
	return s
}

func (s *RpcDataUploadResponseBody) SetTotalBytes(v int64) *RpcDataUploadResponseBody {
	s.TotalBytes = &v
	return s
}

func (s *RpcDataUploadResponseBody) SetTotalTime(v int64) *RpcDataUploadResponseBody {
	s.TotalTime = &v
	return s
}

func (s *RpcDataUploadResponseBody) SetUrl(v string) *RpcDataUploadResponseBody {
	s.Url = &v
	return s
}

type RpcDataUploadResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RpcDataUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RpcDataUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadResponse) GoString() string {
	return s.String()
}

func (s *RpcDataUploadResponse) SetHeaders(v map[string]*string) *RpcDataUploadResponse {
	s.Headers = v
	return s
}

func (s *RpcDataUploadResponse) SetStatusCode(v int32) *RpcDataUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *RpcDataUploadResponse) SetBody(v *RpcDataUploadResponseBody) *RpcDataUploadResponse {
	s.Body = v
	return s
}

type RpcDataUploadAndDownloadRequest struct {
	Query1 *string `json:"query1,omitempty" xml:"query1,omitempty"`
}

func (s RpcDataUploadAndDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadAndDownloadRequest) GoString() string {
	return s.String()
}

func (s *RpcDataUploadAndDownloadRequest) SetQuery1(v string) *RpcDataUploadAndDownloadRequest {
	s.Query1 = &v
	return s
}

type RpcDataUploadAndDownloadResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RpcDataUploadAndDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadAndDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *RpcDataUploadAndDownloadResponseBody) SetRequestId(v string) *RpcDataUploadAndDownloadResponseBody {
	s.RequestId = &v
	return s
}

type RpcDataUploadAndDownloadResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RpcDataUploadAndDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RpcDataUploadAndDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadAndDownloadResponse) GoString() string {
	return s.String()
}

func (s *RpcDataUploadAndDownloadResponse) SetHeaders(v map[string]*string) *RpcDataUploadAndDownloadResponse {
	s.Headers = v
	return s
}

func (s *RpcDataUploadAndDownloadResponse) SetStatusCode(v int32) *RpcDataUploadAndDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *RpcDataUploadAndDownloadResponse) SetBody(v *RpcDataUploadAndDownloadResponseBody) *RpcDataUploadAndDownloadResponse {
	s.Body = v
	return s
}

type RpcDataUploadTestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RpcDataUploadTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadTestResponseBody) GoString() string {
	return s.String()
}

func (s *RpcDataUploadTestResponseBody) SetRequestId(v string) *RpcDataUploadTestResponseBody {
	s.RequestId = &v
	return s
}

type RpcDataUploadTestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RpcDataUploadTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RpcDataUploadTestResponse) String() string {
	return tea.Prettify(s)
}

func (s RpcDataUploadTestResponse) GoString() string {
	return s.String()
}

func (s *RpcDataUploadTestResponse) SetHeaders(v map[string]*string) *RpcDataUploadTestResponse {
	s.Headers = v
	return s
}

func (s *RpcDataUploadTestResponse) SetStatusCode(v int32) *RpcDataUploadTestResponse {
	s.StatusCode = &v
	return s
}

func (s *RpcDataUploadTestResponse) SetBody(v *RpcDataUploadTestResponseBody) *RpcDataUploadTestResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("ft.aliyuncs.com"),
		"ap-south-1":                  tea.String("ft.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-2":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-3":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-5":              tea.String("ft.aliyuncs.com"),
		"cn-beijing":                  tea.String("ft.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ft.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ft.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ft.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ft.aliyuncs.com"),
		"cn-chengdu":                  tea.String("ft.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ft.aliyuncs.com"),
		"cn-fujian":                   tea.String("ft.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ft.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ft.aliyuncs.com"),
		"cn-huhehaote":                tea.String("ft.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ft.aliyuncs.com"),
		"cn-qingdao":                  tea.String("ft.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ft.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ft.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ft.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ft.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ft.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ft.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ft.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("ft.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ft.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ft.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ft.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ft.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ft.aliyuncs.com"),
		"eu-central-1":                tea.String("ft.aliyuncs.com"),
		"eu-west-1":                   tea.String("ft.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ft.aliyuncs.com"),
		"me-east-1":                   tea.String("ft.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ft.aliyuncs.com"),
		"us-west-1":                   tea.String("ft.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ft"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DataRateLimitTestWithOptions(runtime *util.RuntimeOptions) (_result *DataRateLimitTestResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DataRateLimitTest"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DataRateLimitTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DataRateLimitTest() (_result *DataRateLimitTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DataRateLimitTestResponse{}
	_body, _err := client.DataRateLimitTestWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NormalRpcHsfApiWithOptions(runtime *util.RuntimeOptions) (_result *NormalRpcHsfApiResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("NormalRpcHsfApi"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NormalRpcHsfApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NormalRpcHsfApi() (_result *NormalRpcHsfApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NormalRpcHsfApiResponse{}
	_body, _err := client.NormalRpcHsfApiWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NormalRpcHttpApiWithOptions(runtime *util.RuntimeOptions) (_result *NormalRpcHttpApiResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("NormalRpcHttpApi"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NormalRpcHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NormalRpcHttpApi() (_result *NormalRpcHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NormalRpcHttpApiResponse{}
	_body, _err := client.NormalRpcHttpApiWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RpcDataUploadWithOptions(request *RpcDataUploadRequest, runtime *util.RuntimeOptions) (_result *RpcDataUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query1)) {
		query["query1"] = request.Query1
	}

	if !tea.BoolValue(util.IsUnset(request.Query2)) {
		query["query2"] = request.Query2
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LargeParam)) {
		body["largeParam"] = request.LargeParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RpcDataUpload"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RpcDataUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RpcDataUpload(request *RpcDataUploadRequest) (_result *RpcDataUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RpcDataUploadResponse{}
	_body, _err := client.RpcDataUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RpcDataUploadAndDownloadWithOptions(request *RpcDataUploadAndDownloadRequest, runtime *util.RuntimeOptions) (_result *RpcDataUploadAndDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query1)) {
		query["query1"] = request.Query1
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RpcDataUploadAndDownload"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RpcDataUploadAndDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RpcDataUploadAndDownload(request *RpcDataUploadAndDownloadRequest) (_result *RpcDataUploadAndDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RpcDataUploadAndDownloadResponse{}
	_body, _err := client.RpcDataUploadAndDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RpcDataUploadTestWithOptions(runtime *util.RuntimeOptions) (_result *RpcDataUploadTestResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RpcDataUploadTest"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RpcDataUploadTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RpcDataUploadTest() (_result *RpcDataUploadTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RpcDataUploadTestResponse{}
	_body, _err := client.RpcDataUploadTestWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
