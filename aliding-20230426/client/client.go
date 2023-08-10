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

type CreateSheetHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetHeaders) SetAccountContext(v *CreateSheetHeadersAccountContext) *CreateSheetHeaders {
	s.AccountContext = v
	return s
}

type CreateSheetHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateSheetHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSheetHeadersAccountContext) SetAccountId(v string) *CreateSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

type CreateSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSheetShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetShrinkHeaders) SetAccountContextShrink(v string) *CreateSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type CreateSheetRequest struct {
	Name          *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CreateSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkbookId    *string                          `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetRequest) SetName(v string) *CreateSheetRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetRequest) SetTenantContext(v *CreateSheetRequestTenantContext) *CreateSheetRequest {
	s.TenantContext = v
	return s
}

func (s *CreateSheetRequest) SetWorkbookId(v string) *CreateSheetRequest {
	s.WorkbookId = &v
	return s
}

type CreateSheetRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateSheetRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateSheetRequestTenantContext) SetTenantId(v string) *CreateSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

type CreateSheetShrinkRequest struct {
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkbookId          *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkRequest) SetName(v string) *CreateSheetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetTenantContextShrink(v string) *CreateSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetWorkbookId(v string) *CreateSheetShrinkRequest {
	s.WorkbookId = &v
	return s
}

type CreateSheetResponseBody struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSheetResponseBody) SetId(v string) *CreateSheetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSheetResponseBody) SetName(v string) *CreateSheetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSheetResponseBody) SetRequestId(v string) *CreateSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSheetResponseBody) SetVisibility(v string) *CreateSheetResponseBody {
	s.Visibility = &v
	return s
}

type CreateSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSheetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponse) GoString() string {
	return s.String()
}

func (s *CreateSheetResponse) SetHeaders(v map[string]*string) *CreateSheetResponse {
	s.Headers = v
	return s
}

func (s *CreateSheetResponse) SetStatusCode(v int32) *CreateSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSheetResponse) SetBody(v *CreateSheetResponseBody) *CreateSheetResponse {
	s.Body = v
	return s
}

type InsertRowsBeforeHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertRowsBeforeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertRowsBeforeHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeHeaders) SetAccountContext(v *InsertRowsBeforeHeadersAccountContext) *InsertRowsBeforeHeaders {
	s.AccountContext = v
	return s
}

type InsertRowsBeforeHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertRowsBeforeHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeadersAccountContext) SetAccountId(v string) *InsertRowsBeforeHeadersAccountContext {
	s.AccountId = &v
	return s
}

type InsertRowsBeforeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertRowsBeforeShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeShrinkHeaders) SetAccountContextShrink(v string) *InsertRowsBeforeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type InsertRowsBeforeRequest struct {
	Row           *int64                                `json:"Row,omitempty" xml:"Row,omitempty"`
	RowCount      *int64                                `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	SheetId       *string                               `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *InsertRowsBeforeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkbookId    *string                               `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertRowsBeforeRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequest) SetRow(v int64) *InsertRowsBeforeRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetRowCount(v int64) *InsertRowsBeforeRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetSheetId(v string) *InsertRowsBeforeRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetTenantContext(v *InsertRowsBeforeRequestTenantContext) *InsertRowsBeforeRequest {
	s.TenantContext = v
	return s
}

func (s *InsertRowsBeforeRequest) SetWorkbookId(v string) *InsertRowsBeforeRequest {
	s.WorkbookId = &v
	return s
}

type InsertRowsBeforeRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertRowsBeforeRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequestTenantContext) SetTenantId(v string) *InsertRowsBeforeRequestTenantContext {
	s.TenantId = &v
	return s
}

type InsertRowsBeforeShrinkRequest struct {
	Row                 *int64  `json:"Row,omitempty" xml:"Row,omitempty"`
	RowCount            *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	SheetId             *string `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkbookId          *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertRowsBeforeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkRequest) SetRow(v int64) *InsertRowsBeforeShrinkRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetRowCount(v int64) *InsertRowsBeforeShrinkRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetSheetId(v string) *InsertRowsBeforeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetTenantContextShrink(v string) *InsertRowsBeforeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetWorkbookId(v string) *InsertRowsBeforeShrinkRequest {
	s.WorkbookId = &v
	return s
}

type InsertRowsBeforeResponseBody struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InsertRowsBeforeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponseBody) SetId(v string) *InsertRowsBeforeResponseBody {
	s.Id = &v
	return s
}

func (s *InsertRowsBeforeResponseBody) SetRequestId(v string) *InsertRowsBeforeResponseBody {
	s.RequestId = &v
	return s
}

type InsertRowsBeforeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertRowsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertRowsBeforeResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponse) SetHeaders(v map[string]*string) *InsertRowsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertRowsBeforeResponse) SetStatusCode(v int32) *InsertRowsBeforeResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertRowsBeforeResponse) SetBody(v *InsertRowsBeforeResponseBody) *InsertRowsBeforeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliding"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateSheetWithOptions(tmpReq *CreateSheetRequest, tmpHeader *CreateSheetHeaders, runtime *util.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkbookId)) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSheet"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/createSheet"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSheetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSheet(request *CreateSheetRequest) (_result *CreateSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSheetHeaders{}
	_result = &CreateSheetResponse{}
	_body, _err := client.CreateSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRowsBeforeWithOptions(tmpReq *InsertRowsBeforeRequest, tmpHeader *InsertRowsBeforeHeaders, runtime *util.RuntimeOptions) (_result *InsertRowsBeforeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertRowsBeforeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertRowsBeforeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["Row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["RowCount"] = request.RowCount
	}

	if !tea.BoolValue(util.IsUnset(request.SheetId)) {
		body["SheetId"] = request.SheetId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkbookId)) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertRowsBefore"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/insertRowsBefore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRowsBefore(request *InsertRowsBeforeRequest) (_result *InsertRowsBeforeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertRowsBeforeHeaders{}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.InsertRowsBeforeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
