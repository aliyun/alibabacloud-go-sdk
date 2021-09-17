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

type GetDeviceIdByIdentityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceIdByIdentityHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceIdByIdentityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceIdByIdentityHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetAuthorization(v string) *GetDeviceIdByIdentityHeaders {
	s.Authorization = &v
	return s
}

type GetDeviceIdByIdentityRequest struct {
	// 认证类型，MAC/SN
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// 编码类型，获取猫精的设备标识的途径有多种，不同途径对应不同的编码类型： PACKAGE_NAME：apk包名 SKILL_ID：技能id
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// 认证标识，具体的MAC地址或者SN号
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	// 产品唯一标志符productKey，从AI开放平台中获取
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// 编码类型对应的值，例如：编码类型是SKILLID，其值就为webhook服务中得到的skillId；编码类似是PACKAGENAME，其值就为对应客户端app的packageName。
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
}

func (s GetDeviceIdByIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityType(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeType(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityId(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityId = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetProductKey(v string) *GetDeviceIdByIdentityRequest {
	s.ProductKey = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeKey(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeKey = &v
	return s
}

type GetDeviceIdByIdentityResponseBody struct {
	// 返回的错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回的错误码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回result
	Result *GetDeviceIdByIdentityResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceIdByIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBody) SetMessage(v string) *GetDeviceIdByIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetCode(v int32) *GetDeviceIdByIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetResult(v *GetDeviceIdByIdentityResponseBodyResult) *GetDeviceIdByIdentityResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetRequestId(v string) *GetDeviceIdByIdentityResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceIdByIdentityResponseBodyResult struct {
	// 设备信息对应的openId
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// 组织id及归一id列表
	DeviceUnionIds []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s GetDeviceIdByIdentityResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceOpenId(v string) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceUnionIds(v []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

type GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds struct {
	// 组织id，
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// 组织id对应的归一id
	DeviceUnionId *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

type GetDeviceIdByIdentityResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceIdByIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceIdByIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceIdByIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponse) SetHeaders(v map[string]*string) *GetDeviceIdByIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceIdByIdentityResponse) SetBody(v *GetDeviceIdByIdentityResponseBody) *GetDeviceIdByIdentityResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetDeviceIdByIdentity(request *GetDeviceIdByIdentityRequest) (_result *GetDeviceIdByIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceIdByIdentityHeaders{}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.GetDeviceIdByIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceIdByIdentityWithOptions(request *GetDeviceIdByIdentityRequest, headers *GetDeviceIdByIdentityHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceIdByIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		query["IdentityType"] = request.IdentityType
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		query["IdentityId"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeKey)) {
		query["EncodeKey"] = request.EncodeKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = headers.XAcsAligenieAccessToken
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = headers.Authorization
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceIdByIdentity"),
		Version:     tea.String("ssp_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ssp/getDeviceIdByIdentity"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
