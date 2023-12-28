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

type RevokeFeedbackRequest struct {
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RevokeFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeFeedbackRequest) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackRequest) SetSampleType(v string) *RevokeFeedbackRequest {
	s.SampleType = &v
	return s
}

func (s *RevokeFeedbackRequest) SetValue(v string) *RevokeFeedbackRequest {
	s.Value = &v
	return s
}

type RevokeFeedbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackResponseBody) SetCode(v string) *RevokeFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeFeedbackResponseBody) SetMessage(v string) *RevokeFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeFeedbackResponseBody) SetRequestId(v string) *RevokeFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type RevokeFeedbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeFeedbackResponse) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackResponse) SetHeaders(v map[string]*string) *RevokeFeedbackResponse {
	s.Headers = v
	return s
}

func (s *RevokeFeedbackResponse) SetStatusCode(v int32) *RevokeFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeFeedbackResponse) SetBody(v *RevokeFeedbackResponseBody) *RevokeFeedbackResponse {
	s.Body = v
	return s
}

type SendFeedbackRequest struct {
	RiskLabel  *string `json:"RiskLabel,omitempty" xml:"RiskLabel,omitempty"`
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SendFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFeedbackRequest) GoString() string {
	return s.String()
}

func (s *SendFeedbackRequest) SetRiskLabel(v string) *SendFeedbackRequest {
	s.RiskLabel = &v
	return s
}

func (s *SendFeedbackRequest) SetSampleType(v string) *SendFeedbackRequest {
	s.SampleType = &v
	return s
}

func (s *SendFeedbackRequest) SetValue(v string) *SendFeedbackRequest {
	s.Value = &v
	return s
}

type SendFeedbackResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *SendFeedbackResponseBody) SetCode(v int32) *SendFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *SendFeedbackResponseBody) SetMessage(v string) *SendFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *SendFeedbackResponseBody) SetRequestId(v string) *SendFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type SendFeedbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SendFeedbackResponse) GoString() string {
	return s.String()
}

func (s *SendFeedbackResponse) SetHeaders(v map[string]*string) *SendFeedbackResponse {
	s.Headers = v
	return s
}

func (s *SendFeedbackResponse) SetStatusCode(v int32) *SendFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFeedbackResponse) SetBody(v *SendFeedbackResponseBody) *SendFeedbackResponse {
	s.Body = v
	return s
}

type UploadSampleApiRequest struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DataValue  *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	Service    *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s UploadSampleApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadSampleApiRequest) GoString() string {
	return s.String()
}

func (s *UploadSampleApiRequest) SetDataType(v string) *UploadSampleApiRequest {
	s.DataType = &v
	return s
}

func (s *UploadSampleApiRequest) SetDataValue(v string) *UploadSampleApiRequest {
	s.DataValue = &v
	return s
}

func (s *UploadSampleApiRequest) SetSampleType(v string) *UploadSampleApiRequest {
	s.SampleType = &v
	return s
}

func (s *UploadSampleApiRequest) SetService(v string) *UploadSampleApiRequest {
	s.Service = &v
	return s
}

type UploadSampleApiResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadSampleApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadSampleApiResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSampleApiResponseBody) SetCode(v string) *UploadSampleApiResponseBody {
	s.Code = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetMessage(v string) *UploadSampleApiResponseBody {
	s.Message = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetRequestId(v string) *UploadSampleApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetSuccess(v string) *UploadSampleApiResponseBody {
	s.Success = &v
	return s
}

type UploadSampleApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadSampleApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadSampleApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadSampleApiResponse) GoString() string {
	return s.String()
}

func (s *UploadSampleApiResponse) SetHeaders(v map[string]*string) *UploadSampleApiResponse {
	s.Headers = v
	return s
}

func (s *UploadSampleApiResponse) SetStatusCode(v int32) *UploadSampleApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSampleApiResponse) SetBody(v *UploadSampleApiResponseBody) *UploadSampleApiResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("safconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) RevokeFeedbackWithOptions(request *RevokeFeedbackRequest, runtime *util.RuntimeOptions) (_result *RevokeFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SampleType)) {
		body["SampleType"] = request.SampleType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeFeedback"),
		Version:     tea.String("2021-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeFeedback(request *RevokeFeedbackRequest) (_result *RevokeFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeFeedbackResponse{}
	_body, _err := client.RevokeFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendFeedbackWithOptions(request *SendFeedbackRequest, runtime *util.RuntimeOptions) (_result *SendFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RiskLabel)) {
		query["RiskLabel"] = request.RiskLabel
	}

	if !tea.BoolValue(util.IsUnset(request.SampleType)) {
		query["SampleType"] = request.SampleType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendFeedback"),
		Version:     tea.String("2021-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendFeedback(request *SendFeedbackRequest) (_result *SendFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendFeedbackResponse{}
	_body, _err := client.SendFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadSampleApiWithOptions(request *UploadSampleApiRequest, runtime *util.RuntimeOptions) (_result *UploadSampleApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DataValue)) {
		query["DataValue"] = request.DataValue
	}

	if !tea.BoolValue(util.IsUnset(request.SampleType)) {
		query["SampleType"] = request.SampleType
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadSampleApi"),
		Version:     tea.String("2021-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadSampleApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadSampleApi(request *UploadSampleApiRequest) (_result *UploadSampleApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadSampleApiResponse{}
	_body, _err := client.UploadSampleApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
