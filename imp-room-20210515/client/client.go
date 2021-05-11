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

type GetLoginTokenRequest struct {
	// AppId
	AppId         *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *GetLoginTokenRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) SetAppId(v string) *GetLoginTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetLoginTokenRequest) SetRequestParams(v *GetLoginTokenRequestRequestParams) *GetLoginTokenRequest {
	s.RequestParams = v
	return s
}

type GetLoginTokenRequestRequestParams struct {
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// AppKey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GetLoginTokenRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequestRequestParams) SetAppUid(v string) *GetLoginTokenRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *GetLoginTokenRequestRequestParams) SetAppKey(v string) *GetLoginTokenRequestRequestParams {
	s.AppKey = &v
	return s
}

func (s *GetLoginTokenRequestRequestParams) SetDeviceId(v string) *GetLoginTokenRequestRequestParams {
	s.DeviceId = &v
	return s
}

type GetLoginTokenShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetLoginTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenShrinkRequest) SetAppId(v string) *GetLoginTokenShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetRequestParamsShrink(v string) *GetLoginTokenShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetLoginTokenResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *GetLoginTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetCode(v string) *GetLoginTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetMessage(v string) *GetLoginTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetResult(v *GetLoginTokenResponseBodyResult) *GetLoginTokenResponseBody {
	s.Result = v
	return s
}

type GetLoginTokenResponseBodyResult struct {
	// 登录Tokon
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 更新Token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// 登录Token过期时间
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
}

func (s GetLoginTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyResult) SetAccessToken(v string) *GetLoginTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *GetLoginTokenResponseBodyResult) SetRefreshToken(v string) *GetLoginTokenResponseBodyResult {
	s.RefreshToken = &v
	return s
}

func (s *GetLoginTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetLoginTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

type GetLoginTokenResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

type CreateRoomRequest struct {
	Request *CreateRoomRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s CreateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequest) SetRequest(v *CreateRoomRequestRequest) *CreateRoomRequest {
	s.Request = v
	return s
}

type CreateRoomRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 业务类型
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间公告
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 拓展字段
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateRoomRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequestRequest) SetDomain(v string) *CreateRoomRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateRoomRequestRequest) SetBizType(v string) *CreateRoomRequestRequest {
	s.BizType = &v
	return s
}

func (s *CreateRoomRequestRequest) SetTemplateId(v string) *CreateRoomRequestRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetRoomId(v string) *CreateRoomRequestRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetTitle(v string) *CreateRoomRequestRequest {
	s.Title = &v
	return s
}

func (s *CreateRoomRequestRequest) SetNotice(v string) *CreateRoomRequestRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomRequestRequest) SetOwnerId(v string) *CreateRoomRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetExtension(v map[string]*string) *CreateRoomRequestRequest {
	s.Extension = v
	return s
}

type CreateRoomResponseBody struct {
	// Id of the request
	RequestId       *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result          *CreateRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResponseSuccess *bool                         `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s CreateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBody) SetRequestId(v string) *CreateRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoomResponseBody) SetResult(v *CreateRoomResponseBodyResult) *CreateRoomResponseBody {
	s.Result = v
	return s
}

func (s *CreateRoomResponseBody) SetResponseSuccess(v bool) *CreateRoomResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateRoomResponseBody) SetErrorCode(v string) *CreateRoomResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRoomResponseBody) SetErrorMsg(v string) *CreateRoomResponseBody {
	s.ErrorMsg = &v
	return s
}

type CreateRoomResponseBodyResult struct {
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s CreateRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBodyResult) SetRoomId(v string) *CreateRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

type CreateRoomResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateRoomResponse) SetHeaders(v map[string]*string) *CreateRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateRoomResponse) SetBody(v *CreateRoomResponseBody) *CreateRoomResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imp-room"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetLoginTokenWithOptions(tmpReq *GetLoginTokenRequest, runtime *util.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetLoginTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLoginToken"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoomWithOptions(request *CreateRoomRequest, runtime *util.RuntimeOptions) (_result *CreateRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRoom"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRoom(request *CreateRoomRequest) (_result *CreateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoomResponse{}
	_body, _err := client.CreateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
