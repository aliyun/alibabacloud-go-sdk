// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAgAccountRequest struct {
	// This parameter is required.
	LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAgAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAgAccountRequest) SetLoginEmail(v string) *CreateAgAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *CreateAgAccountRequest) SetType(v string) *CreateAgAccountRequest {
	s.Type = &v
	return s
}

type CreateAgAccountResponseBody struct {
	AgRelationDto *CreateAgAccountResponseBodyAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Struct"`
	// example:
	//
	// LOGIN_EMAIL_ALREADY_EXIST
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// this email has already been used.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 28B37990-2326-5A65-9F71-8251590F4779
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBody) SetAgRelationDto(v *CreateAgAccountResponseBodyAgRelationDto) *CreateAgAccountResponseBody {
	s.AgRelationDto = v
	return s
}

func (s *CreateAgAccountResponseBody) SetCode(v string) *CreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetMessage(v string) *CreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetRequestId(v string) *CreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetSuccess(v bool) *CreateAgAccountResponseBody {
	s.Success = &v
	return s
}

type CreateAgAccountResponseBodyAgRelationDto struct {
	// MPK（UID）
	//
	// example:
	//
	// 1387058806890955
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// example:
	//
	// 1441498519728474
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// example:
	//
	// BID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAgAccountResponseBodyAgRelationDto) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponseBodyAgRelationDto) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetMpk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetPk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Pk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetType(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Type = &v
	return s
}

type CreateAgAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponse) SetHeaders(v map[string]*string) *CreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAgAccountResponse) SetStatusCode(v int32) *CreateAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgAccountResponse) SetBody(v *CreateAgAccountResponseBody) *CreateAgAccountResponse {
	s.Body = v
	return s
}

type GetAgRelationRequest struct {
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetAgRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgRelationRequest) GoString() string {
	return s.String()
}

func (s *GetAgRelationRequest) SetPk(v string) *GetAgRelationRequest {
	s.Pk = &v
	return s
}

type GetAgRelationResponseBody struct {
	AgRelationDto *GetAgRelationResponseBodyAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Struct"`
	Code          *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponseBody) SetAgRelationDto(v *GetAgRelationResponseBodyAgRelationDto) *GetAgRelationResponseBody {
	s.AgRelationDto = v
	return s
}

func (s *GetAgRelationResponseBody) SetCode(v string) *GetAgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgRelationResponseBody) SetMessage(v string) *GetAgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgRelationResponseBody) SetRequestId(v string) *GetAgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgRelationResponseBody) SetSuccess(v bool) *GetAgRelationResponseBody {
	s.Success = &v
	return s
}

type GetAgRelationResponseBodyAgRelationDto struct {
	Mpk  *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	Pk   *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAgRelationResponseBodyAgRelationDto) String() string {
	return tea.Prettify(s)
}

func (s GetAgRelationResponseBodyAgRelationDto) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetMpk(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetPk(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Pk = &v
	return s
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetType(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Type = &v
	return s
}

type GetAgRelationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgRelationResponse) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponse) SetHeaders(v map[string]*string) *GetAgRelationResponse {
	s.Headers = v
	return s
}

func (s *GetAgRelationResponse) SetStatusCode(v int32) *GetAgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgRelationResponse) SetBody(v *GetAgRelationResponseBody) *GetAgRelationResponse {
	s.Body = v
	return s
}

type GetRamBindRequest struct {
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetRamBindRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRamBindRequest) GoString() string {
	return s.String()
}

func (s *GetRamBindRequest) SetPk(v string) *GetRamBindRequest {
	s.Pk = &v
	return s
}

type GetRamBindResponseBody struct {
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RamBindDto *GetRamBindResponseBodyRamBindDto `json:"RamBindDto,omitempty" xml:"RamBindDto,omitempty" type:"Struct"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRamBindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRamBindResponseBody) GoString() string {
	return s.String()
}

func (s *GetRamBindResponseBody) SetCode(v string) *GetRamBindResponseBody {
	s.Code = &v
	return s
}

func (s *GetRamBindResponseBody) SetMessage(v string) *GetRamBindResponseBody {
	s.Message = &v
	return s
}

func (s *GetRamBindResponseBody) SetRamBindDto(v *GetRamBindResponseBodyRamBindDto) *GetRamBindResponseBody {
	s.RamBindDto = v
	return s
}

func (s *GetRamBindResponseBody) SetRequestId(v string) *GetRamBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRamBindResponseBody) SetSuccess(v bool) *GetRamBindResponseBody {
	s.Success = &v
	return s
}

type GetRamBindResponseBodyRamBindDto struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRamBindResponseBodyRamBindDto) String() string {
	return tea.Prettify(s)
}

func (s GetRamBindResponseBodyRamBindDto) GoString() string {
	return s.String()
}

func (s *GetRamBindResponseBodyRamBindDto) SetRoleName(v string) *GetRamBindResponseBodyRamBindDto {
	s.RoleName = &v
	return s
}

type GetRamBindResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRamBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRamBindResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRamBindResponse) GoString() string {
	return s.String()
}

func (s *GetRamBindResponse) SetHeaders(v map[string]*string) *GetRamBindResponse {
	s.Headers = v
	return s
}

func (s *GetRamBindResponse) SetStatusCode(v int32) *GetRamBindResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRamBindResponse) SetBody(v *GetRamBindResponseBody) *GetRamBindResponse {
	s.Body = v
	return s
}

type PaginateAgRelationsRequest struct {
	// This parameter is required.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryCount *bool  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
}

func (s PaginateAgRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PaginateAgRelationsRequest) GoString() string {
	return s.String()
}

func (s *PaginateAgRelationsRequest) SetPageNo(v int64) *PaginateAgRelationsRequest {
	s.PageNo = &v
	return s
}

func (s *PaginateAgRelationsRequest) SetPageSize(v int32) *PaginateAgRelationsRequest {
	s.PageSize = &v
	return s
}

func (s *PaginateAgRelationsRequest) SetQueryCount(v bool) *PaginateAgRelationsRequest {
	s.QueryCount = &v
	return s
}

type PaginateAgRelationsResponseBody struct {
	AgRelationDtos *PaginateAgRelationsResponseBodyAgRelationDtos `json:"AgRelationDtos,omitempty" xml:"AgRelationDtos,omitempty" type:"Struct"`
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo         *int64                                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PaginateAgRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PaginateAgRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *PaginateAgRelationsResponseBody) SetAgRelationDtos(v *PaginateAgRelationsResponseBodyAgRelationDtos) *PaginateAgRelationsResponseBody {
	s.AgRelationDtos = v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetCode(v string) *PaginateAgRelationsResponseBody {
	s.Code = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetMessage(v string) *PaginateAgRelationsResponseBody {
	s.Message = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetPageNo(v int64) *PaginateAgRelationsResponseBody {
	s.PageNo = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetPageSize(v int32) *PaginateAgRelationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetRequestId(v string) *PaginateAgRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetSuccess(v bool) *PaginateAgRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *PaginateAgRelationsResponseBody) SetTotalCount(v int64) *PaginateAgRelationsResponseBody {
	s.TotalCount = &v
	return s
}

type PaginateAgRelationsResponseBodyAgRelationDtos struct {
	AgRelationDto []*PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Repeated"`
}

func (s PaginateAgRelationsResponseBodyAgRelationDtos) String() string {
	return tea.Prettify(s)
}

func (s PaginateAgRelationsResponseBodyAgRelationDtos) GoString() string {
	return s.String()
}

func (s *PaginateAgRelationsResponseBodyAgRelationDtos) SetAgRelationDto(v []*PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) *PaginateAgRelationsResponseBodyAgRelationDtos {
	s.AgRelationDto = v
	return s
}

type PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto struct {
	Mpk  *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	Pk   *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) String() string {
	return tea.Prettify(s)
}

func (s PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) GoString() string {
	return s.String()
}

func (s *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) SetMpk(v string) *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) SetPk(v string) *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto {
	s.Pk = &v
	return s
}

func (s *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto) SetType(v string) *PaginateAgRelationsResponseBodyAgRelationDtosAgRelationDto {
	s.Type = &v
	return s
}

type PaginateAgRelationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PaginateAgRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PaginateAgRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PaginateAgRelationsResponse) GoString() string {
	return s.String()
}

func (s *PaginateAgRelationsResponse) SetHeaders(v map[string]*string) *PaginateAgRelationsResponse {
	s.Headers = v
	return s
}

func (s *PaginateAgRelationsResponse) SetStatusCode(v int32) *PaginateAgRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *PaginateAgRelationsResponse) SetBody(v *PaginateAgRelationsResponseBody) *PaginateAgRelationsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliyunid-ag"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建ag账号
//
// @param request - CreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccountWithOptions(request *CreateAgAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAgAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoginEmail)) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAgAccount"),
		Version:     tea.String("2018-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAgAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAgAccountResponse{}
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
// 创建ag账号
//
// @param request - CreateAgAccountRequest
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccount(request *CreateAgAccountRequest) (_result *CreateAgAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAgRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgRelationResponse
func (client *Client) GetAgRelationWithOptions(request *GetAgRelationRequest, runtime *util.RuntimeOptions) (_result *GetAgRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgRelation"),
		Version:     tea.String("2018-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAgRelationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAgRelationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetAgRelationRequest
//
// @return GetAgRelationResponse
func (client *Client) GetAgRelation(request *GetAgRelationRequest) (_result *GetAgRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgRelationResponse{}
	_body, _err := client.GetAgRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRamBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRamBindResponse
func (client *Client) GetRamBindWithOptions(request *GetRamBindRequest, runtime *util.RuntimeOptions) (_result *GetRamBindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRamBind"),
		Version:     tea.String("2018-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRamBindResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRamBindResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetRamBindRequest
//
// @return GetRamBindResponse
func (client *Client) GetRamBind(request *GetRamBindRequest) (_result *GetRamBindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRamBindResponse{}
	_body, _err := client.GetRamBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PaginateAgRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PaginateAgRelationsResponse
func (client *Client) PaginateAgRelationsWithOptions(request *PaginateAgRelationsRequest, runtime *util.RuntimeOptions) (_result *PaginateAgRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCount)) {
		query["QueryCount"] = request.QueryCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PaginateAgRelations"),
		Version:     tea.String("2018-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PaginateAgRelationsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PaginateAgRelationsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - PaginateAgRelationsRequest
//
// @return PaginateAgRelationsResponse
func (client *Client) PaginateAgRelations(request *PaginateAgRelationsRequest) (_result *PaginateAgRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PaginateAgRelationsResponse{}
	_body, _err := client.PaginateAgRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
