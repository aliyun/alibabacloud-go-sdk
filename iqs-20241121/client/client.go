// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccountInfoManageRequest struct {
	AccountId       *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	QuarkKey        *string `json:"quarkKey,omitempty" xml:"quarkKey,omitempty"`
	QuarkName       *string `json:"quarkName,omitempty" xml:"quarkName,omitempty"`
	TestQps         *int32  `json:"testQps,omitempty" xml:"testQps,omitempty"`
	TestQueryPerDay *int32  `json:"testQueryPerDay,omitempty" xml:"testQueryPerDay,omitempty"`
}

func (s AccountInfoManageRequest) String() string {
	return tea.Prettify(s)
}

func (s AccountInfoManageRequest) GoString() string {
	return s.String()
}

func (s *AccountInfoManageRequest) SetAccountId(v string) *AccountInfoManageRequest {
	s.AccountId = &v
	return s
}

func (s *AccountInfoManageRequest) SetName(v string) *AccountInfoManageRequest {
	s.Name = &v
	return s
}

func (s *AccountInfoManageRequest) SetQuarkKey(v string) *AccountInfoManageRequest {
	s.QuarkKey = &v
	return s
}

func (s *AccountInfoManageRequest) SetQuarkName(v string) *AccountInfoManageRequest {
	s.QuarkName = &v
	return s
}

func (s *AccountInfoManageRequest) SetTestQps(v int32) *AccountInfoManageRequest {
	s.TestQps = &v
	return s
}

func (s *AccountInfoManageRequest) SetTestQueryPerDay(v int32) *AccountInfoManageRequest {
	s.TestQueryPerDay = &v
	return s
}

type ExpendExpiredTimeRequest struct {
	AccountId     *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	TestStartTime *string `json:"testStartTime,omitempty" xml:"testStartTime,omitempty"`
}

func (s ExpendExpiredTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpendExpiredTimeRequest) GoString() string {
	return s.String()
}

func (s *ExpendExpiredTimeRequest) SetAccountId(v string) *ExpendExpiredTimeRequest {
	s.AccountId = &v
	return s
}

func (s *ExpendExpiredTimeRequest) SetTestStartTime(v string) *ExpendExpiredTimeRequest {
	s.TestStartTime = &v
	return s
}

type IrAccountEntity struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName       *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountStatus     *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	AdjustedNormalQps *int32  `json:"adjustedNormalQps,omitempty" xml:"adjustedNormalQps,omitempty"`
	Configuration     *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	CreateTime        *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted         *int32  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	ModifiedTime      *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	QuarkKey          *string `json:"quarkKey,omitempty" xml:"quarkKey,omitempty"`
	QuarkOsr          *string `json:"quarkOsr,omitempty" xml:"quarkOsr,omitempty"`
	QuarkUsername     *string `json:"quarkUsername,omitempty" xml:"quarkUsername,omitempty"`
	SearchType        *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	TestQps           *int32  `json:"testQps,omitempty" xml:"testQps,omitempty"`
	TestQueryPerDay   *int32  `json:"testQueryPerDay,omitempty" xml:"testQueryPerDay,omitempty"`
	TestStartTime     *string `json:"testStartTime,omitempty" xml:"testStartTime,omitempty"`
}

func (s IrAccountEntity) String() string {
	return tea.Prettify(s)
}

func (s IrAccountEntity) GoString() string {
	return s.String()
}

func (s *IrAccountEntity) SetAccountId(v string) *IrAccountEntity {
	s.AccountId = &v
	return s
}

func (s *IrAccountEntity) SetAccountName(v string) *IrAccountEntity {
	s.AccountName = &v
	return s
}

func (s *IrAccountEntity) SetAccountStatus(v string) *IrAccountEntity {
	s.AccountStatus = &v
	return s
}

func (s *IrAccountEntity) SetAdjustedNormalQps(v int32) *IrAccountEntity {
	s.AdjustedNormalQps = &v
	return s
}

func (s *IrAccountEntity) SetConfiguration(v string) *IrAccountEntity {
	s.Configuration = &v
	return s
}

func (s *IrAccountEntity) SetCreateTime(v string) *IrAccountEntity {
	s.CreateTime = &v
	return s
}

func (s *IrAccountEntity) SetId(v int64) *IrAccountEntity {
	s.Id = &v
	return s
}

func (s *IrAccountEntity) SetIsDeleted(v int32) *IrAccountEntity {
	s.IsDeleted = &v
	return s
}

func (s *IrAccountEntity) SetModifiedTime(v string) *IrAccountEntity {
	s.ModifiedTime = &v
	return s
}

func (s *IrAccountEntity) SetQuarkKey(v string) *IrAccountEntity {
	s.QuarkKey = &v
	return s
}

func (s *IrAccountEntity) SetQuarkOsr(v string) *IrAccountEntity {
	s.QuarkOsr = &v
	return s
}

func (s *IrAccountEntity) SetQuarkUsername(v string) *IrAccountEntity {
	s.QuarkUsername = &v
	return s
}

func (s *IrAccountEntity) SetSearchType(v string) *IrAccountEntity {
	s.SearchType = &v
	return s
}

func (s *IrAccountEntity) SetTestQps(v int32) *IrAccountEntity {
	s.TestQps = &v
	return s
}

func (s *IrAccountEntity) SetTestQueryPerDay(v int32) *IrAccountEntity {
	s.TestQueryPerDay = &v
	return s
}

func (s *IrAccountEntity) SetTestStartTime(v string) *IrAccountEntity {
	s.TestStartTime = &v
	return s
}

type IrAccountPageQuery struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s IrAccountPageQuery) String() string {
	return tea.Prettify(s)
}

func (s IrAccountPageQuery) GoString() string {
	return s.String()
}

func (s *IrAccountPageQuery) SetName(v string) *IrAccountPageQuery {
	s.Name = &v
	return s
}

type ManageAccountResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ManageAccountResult) String() string {
	return tea.Prettify(s)
}

func (s ManageAccountResult) GoString() string {
	return s.String()
}

func (s *ManageAccountResult) SetSuccess(v bool) *ManageAccountResult {
	s.Success = &v
	return s
}

type OperationResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OperationResult) String() string {
	return tea.Prettify(s)
}

func (s OperationResult) GoString() string {
	return s.String()
}

func (s *OperationResult) SetSuccess(v bool) *OperationResult {
	s.Success = &v
	return s
}

type ExpandSearchExpiredTimeRequest struct {
	Body *ExpendExpiredTimeRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandSearchExpiredTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpandSearchExpiredTimeRequest) GoString() string {
	return s.String()
}

func (s *ExpandSearchExpiredTimeRequest) SetBody(v *ExpendExpiredTimeRequest) *ExpandSearchExpiredTimeRequest {
	s.Body = v
	return s
}

type ExpandSearchExpiredTimeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandSearchExpiredTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpandSearchExpiredTimeResponse) GoString() string {
	return s.String()
}

func (s *ExpandSearchExpiredTimeResponse) SetHeaders(v map[string]*string) *ExpandSearchExpiredTimeResponse {
	s.Headers = v
	return s
}

func (s *ExpandSearchExpiredTimeResponse) SetStatusCode(v int32) *ExpandSearchExpiredTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpandSearchExpiredTimeResponse) SetBody(v *OperationResult) *ExpandSearchExpiredTimeResponse {
	s.Body = v
	return s
}

type ManageSearchAccountInfoRequest struct {
	Body *AccountInfoManageRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSearchAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageSearchAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *ManageSearchAccountInfoRequest) SetBody(v *AccountInfoManageRequest) *ManageSearchAccountInfoRequest {
	s.Body = v
	return s
}

type ManageSearchAccountInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSearchAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageSearchAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *ManageSearchAccountInfoResponse) SetHeaders(v map[string]*string) *ManageSearchAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *ManageSearchAccountInfoResponse) SetStatusCode(v int32) *ManageSearchAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageSearchAccountInfoResponse) SetBody(v *OperationResult) *ManageSearchAccountInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 延长测试时间
//
// @param request - ExpandSearchExpiredTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandSearchExpiredTimeResponse
func (client *Client) ExpandSearchExpiredTimeWithOptions(request *ExpandSearchExpiredTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExpandSearchExpiredTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpandSearchExpiredTime"),
		Version:     tea.String("2024-11-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/expendExpiredTime"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ExpandSearchExpiredTimeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ExpandSearchExpiredTimeResponse{}
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
// 延长测试时间
//
// @param request - ExpandSearchExpiredTimeRequest
//
// @return ExpandSearchExpiredTimeResponse
func (client *Client) ExpandSearchExpiredTime(request *ExpandSearchExpiredTimeRequest) (_result *ExpandSearchExpiredTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpandSearchExpiredTimeResponse{}
	_body, _err := client.ExpandSearchExpiredTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理智搜用户
//
// @param request - ManageSearchAccountInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageSearchAccountInfoResponse
func (client *Client) ManageSearchAccountInfoWithOptions(request *ManageSearchAccountInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ManageSearchAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageSearchAccountInfo"),
		Version:     tea.String("2024-11-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ManageSearchAccountInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ManageSearchAccountInfoResponse{}
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
// 管理智搜用户
//
// @param request - ManageSearchAccountInfoRequest
//
// @return ManageSearchAccountInfoResponse
func (client *Client) ManageSearchAccountInfo(request *ManageSearchAccountInfoRequest) (_result *ManageSearchAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ManageSearchAccountInfoResponse{}
	_body, _err := client.ManageSearchAccountInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
