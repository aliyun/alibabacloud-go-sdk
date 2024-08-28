// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTokenRequest struct {
	// This parameter is required.
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// This parameter is required.
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetAppKey(v string) *GetTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetTokenRequest) SetAppSecret(v string) *GetTokenRequest {
	s.AppSecret = &v
	return s
}

type GetTokenResponseBody struct {
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorData *GetTokenResponseBodyErrorData `json:"errorData,omitempty" xml:"errorData,omitempty" type:"Struct"`
	ErrorMsg  *string                        `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD708D6-1A8C-5CF9-85B8-D620F314F1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// confirmed
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetErrorCode(v string) *GetTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTokenResponseBody) SetErrorData(v *GetTokenResponseBodyErrorData) *GetTokenResponseBody {
	s.ErrorData = v
	return s
}

func (s *GetTokenResponseBody) SetErrorMsg(v string) *GetTokenResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetStatus(v int32) *GetTokenResponseBody {
	s.Status = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

type GetTokenResponseBodyData struct {
	// example:
	//
	// 7200
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1724130275
	GenerateTime *string `json:"generateTime,omitempty" xml:"generateTime,omitempty"`
	// token
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6ImRpc3RyaWJ1dGlvbl9rZXlpZCJ9.eyJqdGkiOiI2cDQwZDctSDQ0dUJicEJkYTZadzdBIiwiaWF0IjoxNzI0MzE2MzM1LCJleHAiOjE3MjQzMjM1MzUsIm5iZiI6MTcyNDMxNjI3NSwiYXBwS2V5IjoiNjE3NzgxZDQxM2FmNGRlZGFiNzkifQ.XtjSM7qVbESvt7n31RtD5Pp6854IVyGMEco4aEruMDMkrXHkcdZejyecKFx3RdsCldZPgvowc5EOl44c3JJfm6DENH4M6BRBSc90eqXYcD_xVJ9FhDWzK9O6OnkvR7HX1t-qqMdikLviM1w1G0DGMLaasvZ8MPMewL8k6NnvOSGePwUhb-m5IZ14VYv7BPO2dp8Jh00qNSJQrmNiWWzJUtK_xllNr3LKQ7cIVtPgFUckvRDw9Hal5oACXSRdkZFOAGlFSjpB_BbTpe5vc-AJ8K89nRD53sIy9YyVQClV_HH7PrXxFFJgReGvNsnP1h9gf55q86IzOQBkj9vGm2zXdw
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetExpireTime(v string) *GetTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetGenerateTime(v string) *GetTokenResponseBodyData {
	s.GenerateTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetToken(v string) *GetTokenResponseBodyData {
	s.Token = &v
	return s
}

type GetTokenResponseBodyErrorData struct {
	// example:
	//
	// 1001
	RawErrorCode *string `json:"rawErrorCode,omitempty" xml:"rawErrorCode,omitempty"`
	RawErrorMsg  *string `json:"rawErrorMsg,omitempty" xml:"rawErrorMsg,omitempty"`
}

func (s GetTokenResponseBodyErrorData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyErrorData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyErrorData) SetRawErrorCode(v string) *GetTokenResponseBodyErrorData {
	s.RawErrorCode = &v
	return s
}

func (s *GetTokenResponseBodyErrorData) SetRawErrorMsg(v string) *GetTokenResponseBodyErrorData {
	s.RawErrorMsg = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type SearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6ImRpc3RyaWJ1dGlvbl9rZXlpZCJ9.eyJqdGkiOiI2cDQwZDctSDQ0dUJicEJkYTZadzdBIiwiaWF0IjoxNzI0MzE2MzM1LCJleHAiOjE3MjQzMjM1MzUsIm5iZiI6MTcyNDMxNjI3NSwiYXBwS2V5IjoiNjE3NzgxZDQxM2FmNGRlZGFiNzkifQ.XtjSM7qVbESvt7n31RtD5Pp6854IVyGMEco4aEruMDMkrXHkcdZejyecKFx3RdsCldZPgvowc5EOl44c3JJfm6DENH4M6BRBSc90eqXYcD_xVJ9FhDWzK9O6OnkvR7HX1t-qqMdikLviM1w1G0DGMLaasvZ8MPMewL8k6NnvOSGePwUhb-m5IZ14VYv7BPO2dp8Jh00qNSJQrmNiWWzJUtK_xllNr3LKQ7cIVtPgFUckvRDw9Hal5oACXSRdkZFOAGlFSjpB_BbTpe5vc-AJ8K89nRD53sIy9YyVQClV_HH7PrXxFFJgReGvNsnP1h9gf55q86IzOQBkj9vGm2zXdw
	XAcsAirticketAccessToken *string `json:"xAcsAirticketAccessToken,omitempty" xml:"xAcsAirticketAccessToken,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"xAcsAirticketLanguage,omitempty" xml:"xAcsAirticketLanguage,omitempty"`
}

func (s SearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchHeaders) GoString() string {
	return s.String()
}

func (s *SearchHeaders) SetCommonHeaders(v map[string]*string) *SearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchHeaders) SetXAcsAirticketAccessToken(v string) *SearchHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *SearchHeaders) SetXAcsAirticketLanguage(v string) *SearchHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type SearchRequest struct {
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "cabinClass": "ALL_CABIN",
	//
	//   "passengerTypeQuantity": {
	//
	//     "ADT": 1,
	//
	//     "INFANT": 0,
	//
	//     "CHD": 0
	//
	//   },
	//
	//   "searchMode": 0,
	//
	//   "searchOdInfoList": [
	//
	//     {
	//
	//       "arrCityCode": "BJS",
	//
	//       "depCityCode": "JNG",
	//
	//       "depDate": "2024-08-14 00:00:00"
	//
	//     }
	//
	//   ],
	//
	//   "searchSource": "gd",
	//
	//   "tripType": 1
	//
	// }
	SearchParam *string `json:"searchParam,omitempty" xml:"searchParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// “1”
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// “1”
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// “121343”
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequest) GoString() string {
	return s.String()
}

func (s *SearchRequest) SetScene(v string) *SearchRequest {
	s.Scene = &v
	return s
}

func (s *SearchRequest) SetSearchParam(v string) *SearchRequest {
	s.SearchParam = &v
	return s
}

func (s *SearchRequest) SetSource(v string) *SearchRequest {
	s.Source = &v
	return s
}

func (s *SearchRequest) SetTerminal(v string) *SearchRequest {
	s.Terminal = &v
	return s
}

func (s *SearchRequest) SetUserId(v string) *SearchRequest {
	s.UserId = &v
	return s
}

type SearchResponseBody struct {
	// This parameter is required.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	ErrorCode *string                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorData *SearchResponseBodyErrorData `json:"errorData,omitempty" xml:"errorData,omitempty" type:"Struct"`
	ErrorMsg  *string                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 71ad3e90-53f8-445b-be9d-df91039cba47
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResponseBody) SetData(v string) *SearchResponseBody {
	s.Data = &v
	return s
}

func (s *SearchResponseBody) SetErrorCode(v string) *SearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchResponseBody) SetErrorData(v *SearchResponseBodyErrorData) *SearchResponseBody {
	s.ErrorData = v
	return s
}

func (s *SearchResponseBody) SetErrorMsg(v string) *SearchResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchResponseBody) SetRequestId(v string) *SearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResponseBody) SetSuccess(v bool) *SearchResponseBody {
	s.Success = &v
	return s
}

type SearchResponseBodyErrorData struct {
	// example:
	//
	// 1001
	RawErrorCode *string `json:"rawErrorCode,omitempty" xml:"rawErrorCode,omitempty"`
	RawErrorMsg  *string `json:"rawErrorMsg,omitempty" xml:"rawErrorMsg,omitempty"`
}

func (s SearchResponseBodyErrorData) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyErrorData) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyErrorData) SetRawErrorCode(v string) *SearchResponseBodyErrorData {
	s.RawErrorCode = &v
	return s
}

func (s *SearchResponseBodyErrorData) SetRawErrorMsg(v string) *SearchResponseBodyErrorData {
	s.RawErrorMsg = &v
	return s
}

type SearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResponse) GoString() string {
	return s.String()
}

func (s *SearchResponse) SetHeaders(v map[string]*string) *SearchResponse {
	s.Headers = v
	return s
}

func (s *SearchResponse) SetStatusCode(v int32) *SearchResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("trafficfxopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["appSecret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2024-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/distribution/trade/getToken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建token
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分销报价接口
//
// @param request - SearchRequest
//
// @param headers - SearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResponse
func (client *Client) SearchWithOptions(request *SearchRequest, headers *SearchHeaders, runtime *util.RuntimeOptions) (_result *SearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParam)) {
		body["searchParam"] = request.SearchParam
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Terminal)) {
		body["terminal"] = request.Terminal
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["xAcsAirticketAccessToken"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["xAcsAirticketLanguage"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Search"),
		Version:     tea.String("2024-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/distribution/trade/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分销报价接口
//
// @param request - SearchRequest
//
// @return SearchResponse
func (client *Client) Search(request *SearchRequest) (_result *SearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchHeaders{}
	_result = &SearchResponse{}
	_body, _err := client.SearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
