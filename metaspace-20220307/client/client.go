// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyCoordinationWithCodeRequest struct {
	// example:
	//
	// PA3MU***
	CoordinationCode *string `json:"CoordinationCode,omitempty" xml:"CoordinationCode,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// A8B35215993FBF283F28D61******
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyCoordinationWithCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCoordinationWithCodeRequest) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeRequest) SetCoordinationCode(v string) *ApplyCoordinationWithCodeRequest {
	s.CoordinationCode = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetLoginRegionId(v string) *ApplyCoordinationWithCodeRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetLoginToken(v string) *ApplyCoordinationWithCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetSessionId(v string) *ApplyCoordinationWithCodeRequest {
	s.SessionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetUuid(v string) *ApplyCoordinationWithCodeRequest {
	s.Uuid = &v
	return s
}

type ApplyCoordinationWithCodeResponseBody struct {
	Data *ApplyCoordinationWithCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCoordinationWithCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponseBody) SetData(v *ApplyCoordinationWithCodeResponseBodyData) *ApplyCoordinationWithCodeResponseBody {
	s.Data = v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBody) SetRequestId(v string) *ApplyCoordinationWithCodeResponseBody {
	s.RequestId = &v
	return s
}

type ApplyCoordinationWithCodeResponseBodyData struct {
	// example:
	//
	// co-0ad0f3p4n2******
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// example:
	//
	// COORDINATING
	CoordinateStatus *string `json:"CoordinateStatus,omitempty" xml:"CoordinateStatus,omitempty"`
	// example:
	//
	// DQpbRGVza3RvcF0NCkZvcmNlVGxzVHlwZT0xDQ******
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// example:
	//
	// 1126819517******
	CoordinatorAliUid *int64 `json:"CoordinatorAliUid,omitempty" xml:"CoordinatorAliUid,omitempty"`
	// example:
	//
	// bob
	CoordinatorUserId *string `json:"CoordinatorUserId,omitempty" xml:"CoordinatorUserId,omitempty"`
	// example:
	//
	// 1126819517******
	OwnerAliUid *int64 `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
	// example:
	//
	// alice
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// ecd-3vv4mf8zxg******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// demo-desktop
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ApplyCoordinationWithCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinateStatus(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinateStatus = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinateTicket(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinateTicket = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinatorAliUid(v int64) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinatorAliUid = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinatorUserId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinatorUserId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetOwnerAliUid(v int64) *ApplyCoordinationWithCodeResponseBodyData {
	s.OwnerAliUid = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetOwnerUserId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceName(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceRegionId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceType(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceType = &v
	return s
}

type ApplyCoordinationWithCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCoordinationWithCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCoordinationWithCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponse) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponse) SetHeaders(v map[string]*string) *ApplyCoordinationWithCodeResponse {
	s.Headers = v
	return s
}

func (s *ApplyCoordinationWithCodeResponse) SetStatusCode(v int32) *ApplyCoordinationWithCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponse) SetBody(v *ApplyCoordinationWithCodeResponseBody) *ApplyCoordinationWithCodeResponse {
	s.Body = v
	return s
}

type EndAllCoordinationByOwnerRequest struct {
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// ecd-68a7ddrt0******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EndAllCoordinationByOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s EndAllCoordinationByOwnerRequest) GoString() string {
	return s.String()
}

func (s *EndAllCoordinationByOwnerRequest) SetLoginRegionId(v string) *EndAllCoordinationByOwnerRequest {
	s.LoginRegionId = &v
	return s
}

func (s *EndAllCoordinationByOwnerRequest) SetLoginToken(v string) *EndAllCoordinationByOwnerRequest {
	s.LoginToken = &v
	return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceId(v string) *EndAllCoordinationByOwnerRequest {
	s.ResourceId = &v
	return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceRegionId(v string) *EndAllCoordinationByOwnerRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *EndAllCoordinationByOwnerRequest) SetResourceType(v string) *EndAllCoordinationByOwnerRequest {
	s.ResourceType = &v
	return s
}

func (s *EndAllCoordinationByOwnerRequest) SetSessionId(v string) *EndAllCoordinationByOwnerRequest {
	s.SessionId = &v
	return s
}

type EndAllCoordinationByOwnerResponseBody struct {
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndAllCoordinationByOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndAllCoordinationByOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *EndAllCoordinationByOwnerResponseBody) SetRequestId(v string) *EndAllCoordinationByOwnerResponseBody {
	s.RequestId = &v
	return s
}

type EndAllCoordinationByOwnerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EndAllCoordinationByOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndAllCoordinationByOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s EndAllCoordinationByOwnerResponse) GoString() string {
	return s.String()
}

func (s *EndAllCoordinationByOwnerResponse) SetHeaders(v map[string]*string) *EndAllCoordinationByOwnerResponse {
	s.Headers = v
	return s
}

func (s *EndAllCoordinationByOwnerResponse) SetStatusCode(v int32) *EndAllCoordinationByOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *EndAllCoordinationByOwnerResponse) SetBody(v *EndAllCoordinationByOwnerResponseBody) *EndAllCoordinationByOwnerResponse {
	s.Body = v
	return s
}

type GenerateCoordinationCodeRequest struct {
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// ecd-68a7ddrt0******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// demo-desktop
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GenerateCoordinationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeRequest) SetLoginRegionId(v string) *GenerateCoordinationCodeRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetLoginToken(v string) *GenerateCoordinationCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceId(v string) *GenerateCoordinationCodeRequest {
	s.ResourceId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceName(v string) *GenerateCoordinationCodeRequest {
	s.ResourceName = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceRegionId(v string) *GenerateCoordinationCodeRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceType(v string) *GenerateCoordinationCodeRequest {
	s.ResourceType = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetSessionId(v string) *GenerateCoordinationCodeRequest {
	s.SessionId = &v
	return s
}

type GenerateCoordinationCodeResponseBody struct {
	// example:
	//
	// PA3MU***
	CoordinationCode *string `json:"CoordinationCode,omitempty" xml:"CoordinationCode,omitempty"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCoordinationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponseBody) SetCoordinationCode(v string) *GenerateCoordinationCodeResponseBody {
	s.CoordinationCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) SetRequestId(v string) *GenerateCoordinationCodeResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCoordinationCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCoordinationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCoordinationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponse) SetHeaders(v map[string]*string) *GenerateCoordinationCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetStatusCode(v int32) *GenerateCoordinationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetBody(v *GenerateCoordinationCodeResponseBody) *GenerateCoordinationCodeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("metaspace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 用协同码发起协同
//
// @param request - ApplyCoordinationWithCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyCoordinationWithCodeResponse
func (client *Client) ApplyCoordinationWithCodeWithOptions(request *ApplyCoordinationWithCodeRequest, runtime *util.RuntimeOptions) (_result *ApplyCoordinationWithCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoordinationCode)) {
		body["CoordinationCode"] = request.CoordinationCode
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyCoordinationWithCode"),
		Version:     tea.String("2022-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyCoordinationWithCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用协同码发起协同
//
// @param request - ApplyCoordinationWithCodeRequest
//
// @return ApplyCoordinationWithCodeResponse
func (client *Client) ApplyCoordinationWithCode(request *ApplyCoordinationWithCodeRequest) (_result *ApplyCoordinationWithCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyCoordinationWithCodeResponse{}
	_body, _err := client.ApplyCoordinationWithCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Owner主动结束本次协同，同步失效协同码
//
// @param request - EndAllCoordinationByOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndAllCoordinationByOwnerResponse
func (client *Client) EndAllCoordinationByOwnerWithOptions(request *EndAllCoordinationByOwnerRequest, runtime *util.RuntimeOptions) (_result *EndAllCoordinationByOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		body["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EndAllCoordinationByOwner"),
		Version:     tea.String("2022-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EndAllCoordinationByOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Owner主动结束本次协同，同步失效协同码
//
// @param request - EndAllCoordinationByOwnerRequest
//
// @return EndAllCoordinationByOwnerResponse
func (client *Client) EndAllCoordinationByOwner(request *EndAllCoordinationByOwnerRequest) (_result *EndAllCoordinationByOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EndAllCoordinationByOwnerResponse{}
	_body, _err := client.EndAllCoordinationByOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成协同码
//
// @param request - GenerateCoordinationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCodeWithOptions(request *GenerateCoordinationCodeRequest, runtime *util.RuntimeOptions) (_result *GenerateCoordinationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		body["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCoordinationCode"),
		Version:     tea.String("2022-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成协同码
//
// @param request - GenerateCoordinationCodeRequest
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCode(request *GenerateCoordinationCodeRequest) (_result *GenerateCoordinationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.GenerateCoordinationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
