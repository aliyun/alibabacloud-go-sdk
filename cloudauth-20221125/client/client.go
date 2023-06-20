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

type EntElementVerifyRequest struct {
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InfoVerifyType    *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
	LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	LegalPersonName   *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	LicenseNo         *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId    *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	SceneCode         *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyRequest) GoString() string {
	return s.String()
}

func (s *EntElementVerifyRequest) SetEntName(v string) *EntElementVerifyRequest {
	s.EntName = &v
	return s
}

func (s *EntElementVerifyRequest) SetInfoVerifyType(v string) *EntElementVerifyRequest {
	s.InfoVerifyType = &v
	return s
}

func (s *EntElementVerifyRequest) SetLegalPersonCertNo(v string) *EntElementVerifyRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *EntElementVerifyRequest) SetLegalPersonName(v string) *EntElementVerifyRequest {
	s.LegalPersonName = &v
	return s
}

func (s *EntElementVerifyRequest) SetLicenseNo(v string) *EntElementVerifyRequest {
	s.LicenseNo = &v
	return s
}

func (s *EntElementVerifyRequest) SetMerchantBizId(v string) *EntElementVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EntElementVerifyRequest) SetMerchantUserId(v string) *EntElementVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EntElementVerifyRequest) SetSceneCode(v string) *EntElementVerifyRequest {
	s.SceneCode = &v
	return s
}

func (s *EntElementVerifyRequest) SetUserAuthorization(v string) *EntElementVerifyRequest {
	s.UserAuthorization = &v
	return s
}

type EntElementVerifyResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntElementVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponseBody) SetCode(v string) *EntElementVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetMessage(v string) *EntElementVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetRequestId(v string) *EntElementVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetResult(v *EntElementVerifyResponseBodyResult) *EntElementVerifyResponseBody {
	s.Result = v
	return s
}

type EntElementVerifyResponseBodyResult struct {
	BizCode      *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	ReasonCode   *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntElementVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponseBodyResult) SetBizCode(v string) *EntElementVerifyResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonCode(v string) *EntElementVerifyResponseBodyResult {
	s.ReasonCode = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyResponseBodyResult {
	s.ReasonDetail = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetStatus(v string) *EntElementVerifyResponseBodyResult {
	s.Status = &v
	return s
}

type EntElementVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EntElementVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EntElementVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponse) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponse) SetHeaders(v map[string]*string) *EntElementVerifyResponse {
	s.Headers = v
	return s
}

func (s *EntElementVerifyResponse) SetStatusCode(v int32) *EntElementVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EntElementVerifyResponse) SetBody(v *EntElementVerifyResponseBody) *EntElementVerifyResponse {
	s.Body = v
	return s
}

type EntVerifyRequest struct {
	AccountNo         *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InfoVerifyType    *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
	LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	LegalPersonMobile *string `json:"LegalPersonMobile,omitempty" xml:"LegalPersonMobile,omitempty"`
	LegalPersonName   *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	LicenseNo         *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId    *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	RiskModelVersion  *string `json:"RiskModelVersion,omitempty" xml:"RiskModelVersion,omitempty"`
	RiskVerifyType    *string `json:"RiskVerifyType,omitempty" xml:"RiskVerifyType,omitempty"`
	SceneCode         *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyRequest) GoString() string {
	return s.String()
}

func (s *EntVerifyRequest) SetAccountNo(v string) *EntVerifyRequest {
	s.AccountNo = &v
	return s
}

func (s *EntVerifyRequest) SetEntName(v string) *EntVerifyRequest {
	s.EntName = &v
	return s
}

func (s *EntVerifyRequest) SetInfoVerifyType(v string) *EntVerifyRequest {
	s.InfoVerifyType = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonCertNo(v string) *EntVerifyRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonMobile(v string) *EntVerifyRequest {
	s.LegalPersonMobile = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonName(v string) *EntVerifyRequest {
	s.LegalPersonName = &v
	return s
}

func (s *EntVerifyRequest) SetLicenseNo(v string) *EntVerifyRequest {
	s.LicenseNo = &v
	return s
}

func (s *EntVerifyRequest) SetMerchantBizId(v string) *EntVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EntVerifyRequest) SetMerchantUserId(v string) *EntVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EntVerifyRequest) SetRiskModelVersion(v string) *EntVerifyRequest {
	s.RiskModelVersion = &v
	return s
}

func (s *EntVerifyRequest) SetRiskVerifyType(v string) *EntVerifyRequest {
	s.RiskVerifyType = &v
	return s
}

func (s *EntVerifyRequest) SetSceneCode(v string) *EntVerifyRequest {
	s.SceneCode = &v
	return s
}

func (s *EntVerifyRequest) SetUserAuthorization(v string) *EntVerifyRequest {
	s.UserAuthorization = &v
	return s
}

type EntVerifyResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBody) SetCode(v string) *EntVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EntVerifyResponseBody) SetMessage(v string) *EntVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EntVerifyResponseBody) SetRequestId(v string) *EntVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntVerifyResponseBody) SetResult(v *EntVerifyResponseBodyResult) *EntVerifyResponseBody {
	s.Result = v
	return s
}

type EntVerifyResponseBodyResult struct {
	RiskVerifyResult *EntVerifyResponseBodyResultRiskVerifyResult `json:"RiskVerifyResult,omitempty" xml:"RiskVerifyResult,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResult) SetRiskVerifyResult(v *EntVerifyResponseBodyResultRiskVerifyResult) *EntVerifyResponseBodyResult {
	s.RiskVerifyResult = v
	return s
}

type EntVerifyResponseBodyResultRiskVerifyResult struct {
	Found        *bool                                                      `json:"Found,omitempty" xml:"Found,omitempty"`
	ModelResults []*EntVerifyResponseBodyResultRiskVerifyResultModelResults `json:"ModelResults,omitempty" xml:"ModelResults,omitempty" type:"Repeated"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetFound(v bool) *EntVerifyResponseBodyResultRiskVerifyResult {
	s.Found = &v
	return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetModelResults(v []*EntVerifyResponseBodyResultRiskVerifyResultModelResults) *EntVerifyResponseBodyResultRiskVerifyResult {
	s.ModelResults = v
	return s
}

type EntVerifyResponseBodyResultRiskVerifyResultModelResults struct {
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetModelName(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
	s.ModelName = &v
	return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetResult(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
	s.Result = &v
	return s
}

type EntVerifyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EntVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EntVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponse) GoString() string {
	return s.String()
}

func (s *EntVerifyResponse) SetHeaders(v map[string]*string) *EntVerifyResponse {
	s.Headers = v
	return s
}

func (s *EntVerifyResponse) SetStatusCode(v int32) *EntVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EntVerifyResponse) SetBody(v *EntVerifyResponseBody) *EntVerifyResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) EntElementVerifyWithOptions(request *EntElementVerifyRequest, runtime *util.RuntimeOptions) (_result *EntElementVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		query["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.InfoVerifyType)) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntElementVerify"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EntElementVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EntElementVerify(request *EntElementVerifyRequest) (_result *EntElementVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntElementVerifyResponse{}
	_body, _err := client.EntElementVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EntVerifyWithOptions(request *EntVerifyRequest, runtime *util.RuntimeOptions) (_result *EntVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		query["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.InfoVerifyType)) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonMobile)) {
		query["LegalPersonMobile"] = request.LegalPersonMobile
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskModelVersion)) {
		query["RiskModelVersion"] = request.RiskModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RiskVerifyType)) {
		query["RiskVerifyType"] = request.RiskVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntVerify"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EntVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EntVerify(request *EntVerifyRequest) (_result *EntVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntVerifyResponse{}
	_body, _err := client.EntVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
