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

type CheckVerificationRequest struct {
	// The verification code.
	Code                 *string `json:"Code,omitempty" xml:"Code,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service ID that is displayed in the Phone Number Verification Service console.
	ServiceSid *string `json:"ServiceSid,omitempty" xml:"ServiceSid,omitempty"`
	// The mobile phone number of the recipient. You must add the country code to the beginning of the mobile phone number. Example: 6212345\*\*\*\*01.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The unique authentication ID that is returned by calling the StartVerification operation.
	VerificationId *string `json:"VerificationId,omitempty" xml:"VerificationId,omitempty"`
}

func (s CheckVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVerificationRequest) GoString() string {
	return s.String()
}

func (s *CheckVerificationRequest) SetCode(v string) *CheckVerificationRequest {
	s.Code = &v
	return s
}

func (s *CheckVerificationRequest) SetOwnerId(v int64) *CheckVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckVerificationRequest) SetResourceOwnerAccount(v string) *CheckVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckVerificationRequest) SetResourceOwnerId(v int64) *CheckVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckVerificationRequest) SetServiceSid(v string) *CheckVerificationRequest {
	s.ServiceSid = &v
	return s
}

func (s *CheckVerificationRequest) SetTo(v string) *CheckVerificationRequest {
	s.To = &v
	return s
}

func (s *CheckVerificationRequest) SetVerificationId(v string) *CheckVerificationRequest {
	s.VerificationId = &v
	return s
}

type CheckVerificationResponseBody struct {
	// The HTTP status code that was returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that was returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The data that was returned for the successful request. Example: "Model": { "phoneNumber": "", "channel": "", "verificationId": "", "status": "approved" },
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The ID of the request. An ID is a unique identifier that Alibaba Cloud generates for a request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVerificationResponseBody) SetCode(v string) *CheckVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CheckVerificationResponseBody) SetMessage(v string) *CheckVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CheckVerificationResponseBody) SetModel(v map[string]interface{}) *CheckVerificationResponseBody {
	s.Model = v
	return s
}

func (s *CheckVerificationResponseBody) SetRequestId(v string) *CheckVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVerificationResponseBody) SetSuccess(v string) *CheckVerificationResponseBody {
	s.Success = &v
	return s
}

type CheckVerificationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckVerificationResponse) GoString() string {
	return s.String()
}

func (s *CheckVerificationResponse) SetHeaders(v map[string]*string) *CheckVerificationResponse {
	s.Headers = v
	return s
}

func (s *CheckVerificationResponse) SetStatusCode(v int32) *CheckVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVerificationResponse) SetBody(v *CheckVerificationResponseBody) *CheckVerificationResponse {
	s.Body = v
	return s
}

type SearchVerificationRequest struct {
	// The verification code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of the page to return. Pages start from page 1.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when a text message is sent, in milliseconds. You can query text messages that were sent within the last 30 days.
	//
	// Example: 1677600000000.
	SendDate *int64 `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The service ID that is displayed in the Phone Number Verification Service console.
	ServiceSid *string `json:"ServiceSid,omitempty" xml:"ServiceSid,omitempty"`
	// The mobile phone number of the recipient. You must add the country code to the beginning of the mobile phone number. Example: 6212345\*\*\*\*01.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The unique authentication ID that is returned by calling the StartVerification operation.
	VerificationId *string `json:"VerificationId,omitempty" xml:"VerificationId,omitempty"`
}

func (s SearchVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchVerificationRequest) GoString() string {
	return s.String()
}

func (s *SearchVerificationRequest) SetCode(v string) *SearchVerificationRequest {
	s.Code = &v
	return s
}

func (s *SearchVerificationRequest) SetCurrentPage(v int64) *SearchVerificationRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchVerificationRequest) SetOwnerId(v int64) *SearchVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchVerificationRequest) SetPageSize(v int64) *SearchVerificationRequest {
	s.PageSize = &v
	return s
}

func (s *SearchVerificationRequest) SetResourceOwnerAccount(v string) *SearchVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchVerificationRequest) SetResourceOwnerId(v int64) *SearchVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchVerificationRequest) SetSendDate(v int64) *SearchVerificationRequest {
	s.SendDate = &v
	return s
}

func (s *SearchVerificationRequest) SetServiceSid(v string) *SearchVerificationRequest {
	s.ServiceSid = &v
	return s
}

func (s *SearchVerificationRequest) SetTo(v string) *SearchVerificationRequest {
	s.To = &v
	return s
}

func (s *SearchVerificationRequest) SetVerificationId(v string) *SearchVerificationRequest {
	s.VerificationId = &v
	return s
}

type SearchVerificationResponseBody struct {
	// The HTTP status code that was returned for the request.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that was returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The data that was returned for the request. Example: "Model": { "records": \[ { "sendDate":, "channel": "", "serviceSid": "", "to": "", "updatedDate":, "verificationId": "", "status": "" } ], "pageNo": , "totalPage": 1, "pageSize": 20, "totalCount": 1, }
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The ID of the request. An ID is a unique identifier that Alibaba Cloud generates for a request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SearchVerificationResponseBody) SetCode(v string) *SearchVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *SearchVerificationResponseBody) SetMessage(v string) *SearchVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *SearchVerificationResponseBody) SetModel(v map[string]interface{}) *SearchVerificationResponseBody {
	s.Model = v
	return s
}

func (s *SearchVerificationResponseBody) SetRequestId(v string) *SearchVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchVerificationResponseBody) SetSuccess(v bool) *SearchVerificationResponseBody {
	s.Success = &v
	return s
}

type SearchVerificationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchVerificationResponse) GoString() string {
	return s.String()
}

func (s *SearchVerificationResponse) SetHeaders(v map[string]*string) *SearchVerificationResponse {
	s.Headers = v
	return s
}

func (s *SearchVerificationResponse) SetStatusCode(v int32) *SearchVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchVerificationResponse) SetBody(v *SearchVerificationResponseBody) *SearchVerificationResponse {
	s.Body = v
	return s
}

type StartVerificationRequest struct {
	// The channels that you can use for verification.
	//
	// Valid values:
	//
	// *   Voice
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   SMS
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   WhatsApp
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Channel              *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service ID that is displayed in the Phone Number Verification Service console.
	ServiceSid *string `json:"ServiceSid,omitempty" xml:"ServiceSid,omitempty"`
	// The mobile phone number of the recipient. You must add the country code to the beginning of the mobile phone number. Example: 6212345\*\*\*\*01.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s StartVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVerificationRequest) GoString() string {
	return s.String()
}

func (s *StartVerificationRequest) SetChannel(v string) *StartVerificationRequest {
	s.Channel = &v
	return s
}

func (s *StartVerificationRequest) SetOwnerId(v int64) *StartVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *StartVerificationRequest) SetResourceOwnerAccount(v string) *StartVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartVerificationRequest) SetResourceOwnerId(v int64) *StartVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartVerificationRequest) SetServiceSid(v string) *StartVerificationRequest {
	s.ServiceSid = &v
	return s
}

func (s *StartVerificationRequest) SetTo(v string) *StartVerificationRequest {
	s.To = &v
	return s
}

type StartVerificationResponseBody struct {
	// The HTTP status code that was returned for the request.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that was returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The data that was returned only if the request was successful. Example: "Model": { "verifyCode": "", "verificationId": "", "status": "" }
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The ID of the request. An ID is a unique identifier that Alibaba Cloud generates for a request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *StartVerificationResponseBody) SetCode(v string) *StartVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *StartVerificationResponseBody) SetMessage(v string) *StartVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *StartVerificationResponseBody) SetModel(v map[string]interface{}) *StartVerificationResponseBody {
	s.Model = v
	return s
}

func (s *StartVerificationResponseBody) SetRequestId(v string) *StartVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartVerificationResponseBody) SetSuccess(v bool) *StartVerificationResponseBody {
	s.Success = &v
	return s
}

type StartVerificationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVerificationResponse) GoString() string {
	return s.String()
}

func (s *StartVerificationResponse) SetHeaders(v map[string]*string) *StartVerificationResponse {
	s.Headers = v
	return s
}

func (s *StartVerificationResponse) SetStatusCode(v int32) *StartVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartVerificationResponse) SetBody(v *StartVerificationResponseBody) *StartVerificationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dypnsapi-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckVerificationWithOptions(request *CheckVerificationRequest, runtime *util.RuntimeOptions) (_result *CheckVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSid)) {
		query["ServiceSid"] = request.ServiceSid
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationId)) {
		query["VerificationId"] = request.VerificationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckVerification"),
		Version:     tea.String("2017-07-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckVerification(request *CheckVerificationRequest) (_result *CheckVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckVerificationResponse{}
	_body, _err := client.CheckVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchVerificationWithOptions(request *SearchVerificationRequest, runtime *util.RuntimeOptions) (_result *SearchVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDate)) {
		query["SendDate"] = request.SendDate
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSid)) {
		query["ServiceSid"] = request.ServiceSid
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationId)) {
		query["VerificationId"] = request.VerificationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchVerification"),
		Version:     tea.String("2017-07-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchVerification(request *SearchVerificationRequest) (_result *SearchVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchVerificationResponse{}
	_body, _err := client.SearchVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVerificationWithOptions(request *StartVerificationRequest, runtime *util.RuntimeOptions) (_result *StartVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSid)) {
		query["ServiceSid"] = request.ServiceSid
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartVerification"),
		Version:     tea.String("2017-07-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVerification(request *StartVerificationRequest) (_result *StartVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVerificationResponse{}
	_body, _err := client.StartVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
