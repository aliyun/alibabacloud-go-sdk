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

type ScanMsgListRequest struct {
	Biz          *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	ChannelType  *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ScanSequence *int32  `json:"ScanSequence,omitempty" xml:"ScanSequence,omitempty"`
	SequenceId   *int64  `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ScanMsgListRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanMsgListRequest) GoString() string {
	return s.String()
}

func (s *ScanMsgListRequest) SetBiz(v string) *ScanMsgListRequest {
	s.Biz = &v
	return s
}

func (s *ScanMsgListRequest) SetChannelType(v string) *ScanMsgListRequest {
	s.ChannelType = &v
	return s
}

func (s *ScanMsgListRequest) SetCorpId(v string) *ScanMsgListRequest {
	s.CorpId = &v
	return s
}

func (s *ScanMsgListRequest) SetScanSequence(v int32) *ScanMsgListRequest {
	s.ScanSequence = &v
	return s
}

func (s *ScanMsgListRequest) SetSequenceId(v int64) *ScanMsgListRequest {
	s.SequenceId = &v
	return s
}

func (s *ScanMsgListRequest) SetSize(v int32) *ScanMsgListRequest {
	s.Size = &v
	return s
}

func (s *ScanMsgListRequest) SetUserId(v string) *ScanMsgListRequest {
	s.UserId = &v
	return s
}

type ScanMsgListResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ExtMsg    *string                          `json:"ExtMsg,omitempty" xml:"ExtMsg,omitempty"`
	Module    []*ScanMsgListResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	Msg       *string                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScanMsgListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanMsgListResponseBody) GoString() string {
	return s.String()
}

func (s *ScanMsgListResponseBody) SetCode(v int32) *ScanMsgListResponseBody {
	s.Code = &v
	return s
}

func (s *ScanMsgListResponseBody) SetExtMsg(v string) *ScanMsgListResponseBody {
	s.ExtMsg = &v
	return s
}

func (s *ScanMsgListResponseBody) SetModule(v []*ScanMsgListResponseBodyModule) *ScanMsgListResponseBody {
	s.Module = v
	return s
}

func (s *ScanMsgListResponseBody) SetMsg(v string) *ScanMsgListResponseBody {
	s.Msg = &v
	return s
}

func (s *ScanMsgListResponseBody) SetRequestId(v string) *ScanMsgListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanMsgListResponseBody) SetSuccess(v bool) *ScanMsgListResponseBody {
	s.Success = &v
	return s
}

type ScanMsgListResponseBodyModule struct {
	Biz         *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	HaveRedDot  *bool   `json:"HaveRedDot,omitempty" xml:"HaveRedDot,omitempty"`
	SendTime    *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SequenceId  *int64  `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	ShowMsg     *bool   `json:"ShowMsg,omitempty" xml:"ShowMsg,omitempty"`
	UnReadCount *int32  `json:"UnReadCount,omitempty" xml:"UnReadCount,omitempty"`
}

func (s ScanMsgListResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s ScanMsgListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ScanMsgListResponseBodyModule) SetBiz(v string) *ScanMsgListResponseBodyModule {
	s.Biz = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetContent(v string) *ScanMsgListResponseBodyModule {
	s.Content = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetHaveRedDot(v bool) *ScanMsgListResponseBodyModule {
	s.HaveRedDot = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetSendTime(v string) *ScanMsgListResponseBodyModule {
	s.SendTime = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetSequenceId(v int64) *ScanMsgListResponseBodyModule {
	s.SequenceId = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetShowMsg(v bool) *ScanMsgListResponseBodyModule {
	s.ShowMsg = &v
	return s
}

func (s *ScanMsgListResponseBodyModule) SetUnReadCount(v int32) *ScanMsgListResponseBodyModule {
	s.UnReadCount = &v
	return s
}

type ScanMsgListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanMsgListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScanMsgListResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanMsgListResponse) GoString() string {
	return s.String()
}

func (s *ScanMsgListResponse) SetHeaders(v map[string]*string) *ScanMsgListResponse {
	s.Headers = v
	return s
}

func (s *ScanMsgListResponse) SetStatusCode(v int32) *ScanMsgListResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanMsgListResponse) SetBody(v *ScanMsgListResponseBody) *ScanMsgListResponse {
	s.Body = v
	return s
}

type TakeAccessTokenRequest struct {
	AppKey    *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s TakeAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenRequest) SetAppKey(v string) *TakeAccessTokenRequest {
	s.AppKey = &v
	return s
}

func (s *TakeAccessTokenRequest) SetAppSecret(v string) *TakeAccessTokenRequest {
	s.AppSecret = &v
	return s
}

type TakeAccessTokenResponseBody struct {
	Code      *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data      *TakeAccessTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                          `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TakeAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponseBody) SetCode(v string) *TakeAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetData(v *TakeAccessTokenResponseBodyData) *TakeAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *TakeAccessTokenResponseBody) SetMessage(v string) *TakeAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetRequestId(v string) *TakeAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetSuccess(v string) *TakeAccessTokenResponseBody {
	s.Success = &v
	return s
}

type TakeAccessTokenResponseBodyData struct {
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Expire      *int64  `json:"expire,omitempty" xml:"expire,omitempty"`
}

func (s TakeAccessTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponseBodyData) SetAccessToken(v string) *TakeAccessTokenResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *TakeAccessTokenResponseBodyData) SetExpire(v int64) *TakeAccessTokenResponseBodyData {
	s.Expire = &v
	return s
}

type TakeAccessTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TakeAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TakeAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponse) SetHeaders(v map[string]*string) *TakeAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *TakeAccessTokenResponse) SetStatusCode(v int32) *TakeAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TakeAccessTokenResponse) SetBody(v *TakeAccessTokenResponseBody) *TakeAccessTokenResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("btripopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ScanMsgList(request *ScanMsgListRequest) (_result *ScanMsgListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScanMsgListResponse{}
	_body, _err := client.ScanMsgListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScanMsgListWithOptions(request *ScanMsgListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScanMsgListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Biz)) {
		body["Biz"] = request.Biz
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["CorpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanSequence)) {
		body["ScanSequence"] = request.ScanSequence
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceId)) {
		body["SequenceId"] = request.SequenceId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanMsgList"),
		Version:     tea.String("2022-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ScanMsgList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanMsgListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TakeAccessToken(request *TakeAccessTokenRequest) (_result *TakeAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TakeAccessTokenResponse{}
	_body, _err := client.TakeAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TakeAccessTokenWithOptions(request *TakeAccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TakeAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["app_key"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TakeAccessToken"),
		Version:     tea.String("2022-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/btrip/open/access-token/take"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TakeAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
