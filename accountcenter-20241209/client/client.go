// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnterpriseOrgQueryLoadTreeRequest struct {
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	// example:
	//
	// true
	LoadOrgOnly  *bool   `json:"LoadOrgOnly,omitempty" xml:"LoadOrgOnly,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CF20ED94-D406-512F-9798-4E1F65720BF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetEncryptTicket(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetLoadOrgOnly(v bool) *EnterpriseOrgQueryLoadTreeRequest {
	s.LoadOrgOnly = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedEcId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedLeId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedNbId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.RequestId = &v
	return s
}

type EnterpriseOrgQueryLoadTreeResponseBody struct {
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A93073FC-1E86-58BA-AB83-54DA6BDB4F03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TreeDto *string `json:"TreeDto,omitempty" xml:"TreeDto,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetCode(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetMessage(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetSuccess(v bool) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Success = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetTreeDto(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.TreeDto = &v
	return s
}

type EnterpriseOrgQueryLoadTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseOrgQueryLoadTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgQueryLoadTreeResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetStatusCode(v int32) *EnterpriseOrgQueryLoadTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetBody(v *EnterpriseOrgQueryLoadTreeResponseBody) *EnterpriseOrgQueryLoadTreeResponse {
	s.Body = v
	return s
}

type EnterpriseRegisterAccountRequest struct {
	// example:
	//
	// 资方支付平台
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	EncryptPassword *string `json:"EncryptPassword,omitempty" xml:"EncryptPassword,omitempty"`
	EncryptTicket   *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	LoginEmail      *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// example:
	//
	// 668514d8083f058f78f7199a
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	OrientedEcId   *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId   *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId   *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A93073FC-1E86-58BA-AB83-54DA6BDB4F03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
	SiteNick         *string `json:"SiteNick,omitempty" xml:"SiteNick,omitempty"`
}

func (s EnterpriseRegisterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountRequest) SetAlias(v string) *EnterpriseRegisterAccountRequest {
	s.Alias = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptPassword(v string) *EnterpriseRegisterAccountRequest {
	s.EncryptPassword = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptTicket(v string) *EnterpriseRegisterAccountRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetLoginEmail(v string) *EnterpriseRegisterAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrganizationId(v string) *EnterpriseRegisterAccountRequest {
	s.OrganizationId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedEcId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedLeId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedNbId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetRequestId(v string) *EnterpriseRegisterAccountRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetShowCompleteInfo(v bool) *EnterpriseRegisterAccountRequest {
	s.ShowCompleteInfo = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetSiteNick(v string) *EnterpriseRegisterAccountRequest {
	s.SiteNick = &v
	return s
}

type EnterpriseRegisterAccountResponseBody struct {
	AccountInfo *EnterpriseRegisterAccountResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The operation is not allowed. Channel state (RELEASED) does not meet expectations (ANSWERED).
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BDFCF081-7BCD-52D5-9D82-0F58D96EFD92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponseBody) SetAccountInfo(v *EnterpriseRegisterAccountResponseBodyAccountInfo) *EnterpriseRegisterAccountResponseBody {
	s.AccountInfo = v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetCode(v string) *EnterpriseRegisterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetMessage(v string) *EnterpriseRegisterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetRequestId(v string) *EnterpriseRegisterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetSuccess(v bool) *EnterpriseRegisterAccountResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRegisterAccountResponseBodyAccountInfo struct {
	EnterpriseLicenseNo *string `json:"EnterpriseLicenseNo,omitempty" xml:"EnterpriseLicenseNo,omitempty"`
	// example:
	//
	// 海南屿可网络科技有限公司
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	// example:
	//
	// 195529
	LoginId *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	// example:
	//
	// 5190216604405754
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseLicenseNo(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.EnterpriseLicenseNo = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseName(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.EnterpriseName = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetLoginId(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.LoginId = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetPk(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.Pk = &v
	return s
}

type EnterpriseRegisterAccountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRegisterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRegisterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponse) SetHeaders(v map[string]*string) *EnterpriseRegisterAccountResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRegisterAccountResponse) SetStatusCode(v int32) *EnterpriseRegisterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRegisterAccountResponse) SetBody(v *EnterpriseRegisterAccountResponseBody) *EnterpriseRegisterAccountResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("accountcenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTreeWithOptions(request *EnterpriseOrgQueryLoadTreeRequest, runtime *util.RuntimeOptions) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoadOrgOnly)) {
		query["LoadOrgOnly"] = request.LoadOrgOnly
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseOrgQueryLoadTree"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTree(request *EnterpriseOrgQueryLoadTreeRequest) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.EnterpriseOrgQueryLoadTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccountWithOptions(request *EnterpriseRegisterAccountRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRegisterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		query["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptPassword)) {
		query["EncryptPassword"] = request.EncryptPassword
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoginEmail)) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SiteNick)) {
		query["SiteNick"] = request.SiteNick
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRegisterAccount"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccount(request *EnterpriseRegisterAccountRequest) (_result *EnterpriseRegisterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.EnterpriseRegisterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
