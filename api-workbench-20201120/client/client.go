// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
  openapiutil  "github.com/alibabacloud-go/openapi-util/service"
  endpointutil  "github.com/alibabacloud-go/endpoint-util/service"
  "github.com/alibabacloud-go/tea/tea"
)

type TestStruct struct {
  TestString *TestStructTestString `json:"testString,omitempty" xml:"testString,omitempty" type:"Struct"`
}

func (s TestStruct) String() string {
  return tea.Prettify(s)
}

func (s TestStruct) GoString() string {
  return s.String()
}

func (s *TestStruct) SetTestString(v *TestStructTestString) *TestStruct {
  s.TestString = v
  return s
}

type TestStructTestString struct {
  Abc []byte `json:"abc,omitempty" xml:"abc,omitempty"`
}

func (s TestStructTestString) String() string {
  return tea.Prettify(s)
}

func (s TestStructTestString) GoString() string {
  return s.String()
}

func (s *TestStructTestString) SetAbc(v []byte) *TestStructTestString {
  s.Abc = v
  return s
}

type GetApiMetaRequest struct {
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
  ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetApiMetaRequest) String() string {
  return tea.Prettify(s)
}

func (s GetApiMetaRequest) GoString() string {
  return s.String()
}

func (s *GetApiMetaRequest) SetApiName(v string) *GetApiMetaRequest {
  s.ApiName = &v
  return s
}

func (s *GetApiMetaRequest) SetProductName(v string) *GetApiMetaRequest {
  s.ProductName = &v
  return s
}

func (s *GetApiMetaRequest) SetVersion(v string) *GetApiMetaRequest {
  s.Version = &v
  return s
}

type GetApiMetaResponseBody struct {
  Api *string `json:"api,omitempty" xml:"api,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  ErrorCodes *string `json:"error_codes,omitempty" xml:"error_codes,omitempty"`
  ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  Params *string `json:"params,omitempty" xml:"params,omitempty"`
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  Request *string `json:"request,omitempty" xml:"request,omitempty"`
  Response *string `json:"response,omitempty" xml:"response,omitempty"`
  ResponseDoc *string `json:"responseDoc,omitempty" xml:"responseDoc,omitempty"`
  Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
  Timeout *int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
  Title *string `json:"title,omitempty" xml:"title,omitempty"`
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetApiMetaResponseBody) String() string {
  return tea.Prettify(s)
}

func (s GetApiMetaResponseBody) GoString() string {
  return s.String()
}

func (s *GetApiMetaResponseBody) SetApi(v string) *GetApiMetaResponseBody {
  s.Api = &v
  return s
}

func (s *GetApiMetaResponseBody) SetDescription(v string) *GetApiMetaResponseBody {
  s.Description = &v
  return s
}

func (s *GetApiMetaResponseBody) SetErrorCodes(v string) *GetApiMetaResponseBody {
  s.ErrorCodes = &v
  return s
}

func (s *GetApiMetaResponseBody) SetExtraInfo(v string) *GetApiMetaResponseBody {
  s.ExtraInfo = &v
  return s
}

func (s *GetApiMetaResponseBody) SetMethod(v string) *GetApiMetaResponseBody {
  s.Method = &v
  return s
}

func (s *GetApiMetaResponseBody) SetParams(v string) *GetApiMetaResponseBody {
  s.Params = &v
  return s
}

func (s *GetApiMetaResponseBody) SetPath(v string) *GetApiMetaResponseBody {
  s.Path = &v
  return s
}

func (s *GetApiMetaResponseBody) SetProtocol(v string) *GetApiMetaResponseBody {
  s.Protocol = &v
  return s
}

func (s *GetApiMetaResponseBody) SetRequest(v string) *GetApiMetaResponseBody {
  s.Request = &v
  return s
}

func (s *GetApiMetaResponseBody) SetResponse(v string) *GetApiMetaResponseBody {
  s.Response = &v
  return s
}

func (s *GetApiMetaResponseBody) SetResponseDoc(v string) *GetApiMetaResponseBody {
  s.ResponseDoc = &v
  return s
}

func (s *GetApiMetaResponseBody) SetSummary(v string) *GetApiMetaResponseBody {
  s.Summary = &v
  return s
}

func (s *GetApiMetaResponseBody) SetTimeout(v int64) *GetApiMetaResponseBody {
  s.Timeout = &v
  return s
}

func (s *GetApiMetaResponseBody) SetTitle(v string) *GetApiMetaResponseBody {
  s.Title = &v
  return s
}

func (s *GetApiMetaResponseBody) SetVersion(v string) *GetApiMetaResponseBody {
  s.Version = &v
  return s
}

type GetApiMetaResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  Body *GetApiMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApiMetaResponse) String() string {
  return tea.Prettify(s)
}

func (s GetApiMetaResponse) GoString() string {
  return s.String()
}

func (s *GetApiMetaResponse) SetHeaders(v map[string]*string) *GetApiMetaResponse {
  s.Headers = v
  return s
}

func (s *GetApiMetaResponse) SetStatusCode(v int32) *GetApiMetaResponse {
  s.StatusCode = &v
  return s
}

func (s *GetApiMetaResponse) SetBody(v *GetApiMetaResponseBody) *GetApiMetaResponse {
  s.Body = v
  return s
}

type MakeCodeRequest struct {
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
  ApiStyle *string `json:"apiStyle,omitempty" xml:"apiStyle,omitempty"`
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
  Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
  Params *string `json:"params,omitempty" xml:"params,omitempty"`
  Product *string `json:"product,omitempty" xml:"product,omitempty"`
  SdkType *string `json:"sdkType,omitempty" xml:"sdkType,omitempty"`
}

func (s MakeCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s MakeCodeRequest) GoString() string {
  return s.String()
}

func (s *MakeCodeRequest) SetApiName(v string) *MakeCodeRequest {
  s.ApiName = &v
  return s
}

func (s *MakeCodeRequest) SetApiStyle(v string) *MakeCodeRequest {
  s.ApiStyle = &v
  return s
}

func (s *MakeCodeRequest) SetApiVersion(v string) *MakeCodeRequest {
  s.ApiVersion = &v
  return s
}

func (s *MakeCodeRequest) SetEndpoint(v string) *MakeCodeRequest {
  s.Endpoint = &v
  return s
}

func (s *MakeCodeRequest) SetMeta(v string) *MakeCodeRequest {
  s.Meta = &v
  return s
}

func (s *MakeCodeRequest) SetParams(v string) *MakeCodeRequest {
  s.Params = &v
  return s
}

func (s *MakeCodeRequest) SetProduct(v string) *MakeCodeRequest {
  s.Product = &v
  return s
}

func (s *MakeCodeRequest) SetSdkType(v string) *MakeCodeRequest {
  s.SdkType = &v
  return s
}

type MakeCodeResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  SdkDemos map[string]*string `json:"sdkDemos,omitempty" xml:"sdkDemos,omitempty"`
}

func (s MakeCodeResponseBody) String() string {
  return tea.Prettify(s)
}

func (s MakeCodeResponseBody) GoString() string {
  return s.String()
}

func (s *MakeCodeResponseBody) SetRequestId(v string) *MakeCodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *MakeCodeResponseBody) SetSdkDemos(v map[string]*string) *MakeCodeResponseBody {
  s.SdkDemos = v
  return s
}

type MakeCodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  Body *MakeCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s MakeCodeResponse) GoString() string {
  return s.String()
}

func (s *MakeCodeResponse) SetHeaders(v map[string]*string) *MakeCodeResponse {
  s.Headers = v
  return s
}

func (s *MakeCodeResponse) SetStatusCode(v int32) *MakeCodeResponse {
  s.StatusCode = &v
  return s
}

func (s *MakeCodeResponse) SetBody(v *MakeCodeResponseBody) *MakeCodeResponse {
  s.Body = v
  return s
}

type SearchProductRequest struct {
  BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
  Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SearchProductRequest) String() string {
  return tea.Prettify(s)
}

func (s SearchProductRequest) GoString() string {
  return s.String()
}

func (s *SearchProductRequest) SetBizType(v string) *SearchProductRequest {
  s.BizType = &v
  return s
}

func (s *SearchProductRequest) SetPage(v int32) *SearchProductRequest {
  s.Page = &v
  return s
}

func (s *SearchProductRequest) SetPageSize(v int32) *SearchProductRequest {
  s.PageSize = &v
  return s
}

func (s *SearchProductRequest) SetQuery(v string) *SearchProductRequest {
  s.Query = &v
  return s
}

func (s *SearchProductRequest) SetToken(v string) *SearchProductRequest {
  s.Token = &v
  return s
}

type SearchProductResponseBody struct {
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  Data *SearchProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SearchProductResponseBody) String() string {
  return tea.Prettify(s)
}

func (s SearchProductResponseBody) GoString() string {
  return s.String()
}

func (s *SearchProductResponseBody) SetCode(v int32) *SearchProductResponseBody {
  s.Code = &v
  return s
}

func (s *SearchProductResponseBody) SetData(v *SearchProductResponseBodyData) *SearchProductResponseBody {
  s.Data = v
  return s
}

func (s *SearchProductResponseBody) SetRequestId(v string) *SearchProductResponseBody {
  s.RequestId = &v
  return s
}

type SearchProductResponseBodyData struct {
  List []*SearchProductResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
  Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
  Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty"`
  PerPage *int64 `json:"perPage,omitempty" xml:"perPage,omitempty"`
  RealTotal *int64 `json:"real_total,omitempty" xml:"real_total,omitempty"`
  RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
  Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchProductResponseBodyData) String() string {
  return tea.Prettify(s)
}

func (s SearchProductResponseBodyData) GoString() string {
  return s.String()
}

func (s *SearchProductResponseBodyData) SetList(v []*SearchProductResponseBodyDataList) *SearchProductResponseBodyData {
  s.List = v
  return s
}

func (s *SearchProductResponseBodyData) SetPage(v int64) *SearchProductResponseBodyData {
  s.Page = &v
  return s
}

func (s *SearchProductResponseBodyData) SetPages(v int64) *SearchProductResponseBodyData {
  s.Pages = &v
  return s
}

func (s *SearchProductResponseBodyData) SetPerPage(v int64) *SearchProductResponseBodyData {
  s.PerPage = &v
  return s
}

func (s *SearchProductResponseBodyData) SetRealTotal(v int64) *SearchProductResponseBodyData {
  s.RealTotal = &v
  return s
}

func (s *SearchProductResponseBodyData) SetRequestId(v string) *SearchProductResponseBodyData {
  s.RequestId = &v
  return s
}

func (s *SearchProductResponseBodyData) SetTotal(v int64) *SearchProductResponseBodyData {
  s.Total = &v
  return s
}

type SearchProductResponseBodyDataList struct     {
  BelongGroup *string `json:"belong_group,omitempty" xml:"belong_group,omitempty"`
  DefaultVersion *string `json:"default_version,omitempty" xml:"default_version,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  LocationCode *string `json:"location_code,omitempty" xml:"location_code,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  SearchSummary *string `json:"search_summary,omitempty" xml:"search_summary,omitempty"`
  ShortName *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
  ShowNameCn *string `json:"show_name_cn,omitempty" xml:"show_name_cn,omitempty"`
  ShowNameEn *string `json:"show_name_en,omitempty" xml:"show_name_en,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SearchProductResponseBodyDataList) String() string {
  return tea.Prettify(s)
}

func (s SearchProductResponseBodyDataList) GoString() string {
  return s.String()
}

func (s *SearchProductResponseBodyDataList) SetBelongGroup(v string) *SearchProductResponseBodyDataList {
  s.BelongGroup = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetDefaultVersion(v string) *SearchProductResponseBodyDataList {
  s.DefaultVersion = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetDescription(v string) *SearchProductResponseBodyDataList {
  s.Description = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetId(v string) *SearchProductResponseBodyDataList {
  s.Id = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetLocationCode(v string) *SearchProductResponseBodyDataList {
  s.LocationCode = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetName(v string) *SearchProductResponseBodyDataList {
  s.Name = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetSearchSummary(v string) *SearchProductResponseBodyDataList {
  s.SearchSummary = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetShortName(v string) *SearchProductResponseBodyDataList {
  s.ShortName = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetShowNameCn(v string) *SearchProductResponseBodyDataList {
  s.ShowNameCn = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetShowNameEn(v string) *SearchProductResponseBodyDataList {
  s.ShowNameEn = &v
  return s
}

func (s *SearchProductResponseBodyDataList) SetStatus(v string) *SearchProductResponseBodyDataList {
  s.Status = &v
  return s
}

type SearchProductResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  Body *SearchProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchProductResponse) String() string {
  return tea.Prettify(s)
}

func (s SearchProductResponse) GoString() string {
  return s.String()
}

func (s *SearchProductResponse) SetHeaders(v map[string]*string) *SearchProductResponse {
  s.Headers = v
  return s
}

func (s *SearchProductResponse) SetStatusCode(v int32) *SearchProductResponse {
  s.StatusCode = &v
  return s
}

func (s *SearchProductResponse) SetBody(v *SearchProductResponseBody) *SearchProductResponse {
  s.Body = v
  return s
}

type TestOpenApiRequestRequest struct {
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
  Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
  Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s TestOpenApiRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s TestOpenApiRequestRequest) GoString() string {
  return s.String()
}

func (s *TestOpenApiRequestRequest) SetApiName(v string) *TestOpenApiRequestRequest {
  s.ApiName = &v
  return s
}

func (s *TestOpenApiRequestRequest) SetApiVersion(v string) *TestOpenApiRequestRequest {
  s.ApiVersion = &v
  return s
}

func (s *TestOpenApiRequestRequest) SetMeta(v string) *TestOpenApiRequestRequest {
  s.Meta = &v
  return s
}

func (s *TestOpenApiRequestRequest) SetParams(v map[string]interface{}) *TestOpenApiRequestRequest {
  s.Params = v
  return s
}

func (s *TestOpenApiRequestRequest) SetProduct(v string) *TestOpenApiRequestRequest {
  s.Product = &v
  return s
}

type TestOpenApiRequestResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TestOpenApiRequestResponseBody) String() string {
  return tea.Prettify(s)
}

func (s TestOpenApiRequestResponseBody) GoString() string {
  return s.String()
}

func (s *TestOpenApiRequestResponseBody) SetRequestId(v string) *TestOpenApiRequestResponseBody {
  s.RequestId = &v
  return s
}

func (s *TestOpenApiRequestResponseBody) SetHeaders(v map[string]*string) *TestOpenApiRequestResponseBody {
  s.Headers = v
  return s
}

func (s *TestOpenApiRequestResponseBody) SetResult(v map[string]interface{}) *TestOpenApiRequestResponseBody {
  s.Result = v
  return s
}

type TestOpenApiRequestResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  Body *TestOpenApiRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestOpenApiRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s TestOpenApiRequestResponse) GoString() string {
  return s.String()
}

func (s *TestOpenApiRequestResponse) SetHeaders(v map[string]*string) *TestOpenApiRequestResponse {
  s.Headers = v
  return s
}

func (s *TestOpenApiRequestResponse) SetStatusCode(v int32) *TestOpenApiRequestResponse {
  s.StatusCode = &v
  return s
}

func (s *TestOpenApiRequestResponse) SetBody(v *TestOpenApiRequestResponseBody) *TestOpenApiRequestResponse {
  s.Body = v
  return s
}

type Client struct {
  openapi.Client
}

func NewClient(config *openapi.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *openapi.Config)(_err error) {
  _err = client.Client.Init(config  )
  if _err != nil {
    return _err
  }
  client.EndpointRule = tea.String("")
  _err = client.CheckConfig(config)
  if _err != nil {
    return _err
  }
  client.Endpoint, _err = client.GetEndpoint(tea.String("api-workbench"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
  if _err != nil {
    return _err
  }

  return nil
}



func (client *Client) GetEndpoint (productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
  if !tea.BoolValue(util.Empty(endpoint)) {
    _result = endpoint
    return _result , _err
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

func (client *Client) GetApiMeta (request *GetApiMetaRequest) (_result *GetApiMetaResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  _result = &GetApiMetaResponse{}
  _body, _err := client.GetApiMetaWithOptions(request, headers, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetApiMetaWithOptions (request *GetApiMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetApiMetaResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  query := map[string]interface{}{}
  if !tea.BoolValue(util.IsUnset(request.ApiName)) {
    query["apiName"] = request.ApiName
  }

  if !tea.BoolValue(util.IsUnset(request.ProductName)) {
    query["productName"] = request.ProductName
  }

  if !tea.BoolValue(util.IsUnset(request.Version)) {
    query["version"] = request.Version
  }

  req := &openapi.OpenApiRequest{
    Headers: headers,
    Query: openapiutil.Query(query),
  }
  params := &openapi.Params{
    Action: tea.String("GetApiMeta"),
    Version: tea.String("2020-11-20"),
    Protocol: tea.String("HTTPS"),
    Pathname: tea.String("/openapi/product/apiInfo"),
    Method: tea.String("GET"),
    AuthType: tea.String("AK"),
    Style: tea.String("ROA"),
    ReqBodyType: tea.String("json"),
    BodyType: tea.String("json"),
  }
  _result = &GetApiMetaResponse{}
  _body, _err := client.CallApi(params, req, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) MakeCode (request *MakeCodeRequest) (_result *MakeCodeResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  _result = &MakeCodeResponse{}
  _body, _err := client.MakeCodeWithOptions(request, headers, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) MakeCodeWithOptions (request *MakeCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MakeCodeResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  body := map[string]interface{}{}
  if !tea.BoolValue(util.IsUnset(request.ApiName)) {
    body["apiName"] = request.ApiName
  }

  if !tea.BoolValue(util.IsUnset(request.ApiStyle)) {
    body["apiStyle"] = request.ApiStyle
  }

  if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
    body["apiVersion"] = request.ApiVersion
  }

  if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
    body["endpoint"] = request.Endpoint
  }

  if !tea.BoolValue(util.IsUnset(request.Meta)) {
    body["meta"] = request.Meta
  }

  if !tea.BoolValue(util.IsUnset(request.Params)) {
    body["params"] = request.Params
  }

  if !tea.BoolValue(util.IsUnset(request.Product)) {
    body["product"] = request.Product
  }

  if !tea.BoolValue(util.IsUnset(request.SdkType)) {
    body["sdkType"] = request.SdkType
  }

  req := &openapi.OpenApiRequest{
    Headers: headers,
    Body: openapiutil.ParseToMap(body),
  }
  params := &openapi.Params{
    Action: tea.String("MakeCode"),
    Version: tea.String("2020-11-20"),
    Protocol: tea.String("HTTPS"),
    Pathname: tea.String("/openapi/product/makeCode"),
    Method: tea.String("POST"),
    AuthType: tea.String("AK"),
    Style: tea.String("ROA"),
    ReqBodyType: tea.String("json"),
    BodyType: tea.String("json"),
  }
  _result = &MakeCodeResponse{}
  _body, _err := client.CallApi(params, req, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SearchProduct (request *SearchProductRequest) (_result *SearchProductResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  _result = &SearchProductResponse{}
  _body, _err := client.SearchProductWithOptions(request, headers, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SearchProductWithOptions (request *SearchProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchProductResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  query := map[string]interface{}{}
  if !tea.BoolValue(util.IsUnset(request.BizType)) {
    query["BizType"] = request.BizType
  }

  if !tea.BoolValue(util.IsUnset(request.Page)) {
    query["Page"] = request.Page
  }

  if !tea.BoolValue(util.IsUnset(request.PageSize)) {
    query["PageSize"] = request.PageSize
  }

  if !tea.BoolValue(util.IsUnset(request.Query)) {
    query["Query"] = request.Query
  }

  if !tea.BoolValue(util.IsUnset(request.Token)) {
    query["Token"] = request.Token
  }

  req := &openapi.OpenApiRequest{
    Headers: headers,
    Query: openapiutil.Query(query),
  }
  params := &openapi.Params{
    Action: tea.String("SearchProduct"),
    Version: tea.String("2020-11-20"),
    Protocol: tea.String("HTTPS"),
    Pathname: tea.String("/openapi/product/search"),
    Method: tea.String("GET"),
    AuthType: tea.String("AK"),
    Style: tea.String("ROA"),
    ReqBodyType: tea.String("json"),
    BodyType: tea.String("json"),
  }
  _result = &SearchProductResponse{}
  _body, _err := client.CallApi(params, req, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) TestOpenApiRequest (request *TestOpenApiRequestRequest) (_result *TestOpenApiRequestResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  _result = &TestOpenApiRequestResponse{}
  _body, _err := client.TestOpenApiRequestWithOptions(request, headers, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) TestOpenApiRequestWithOptions (request *TestOpenApiRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TestOpenApiRequestResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  body := map[string]interface{}{}
  if !tea.BoolValue(util.IsUnset(request.ApiName)) {
    body["apiName"] = request.ApiName
  }

  if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
    body["apiVersion"] = request.ApiVersion
  }

  if !tea.BoolValue(util.IsUnset(request.Meta)) {
    body["meta"] = request.Meta
  }

  if !tea.BoolValue(util.IsUnset(request.Params)) {
    body["params"] = request.Params
  }

  if !tea.BoolValue(util.IsUnset(request.Product)) {
    body["product"] = request.Product
  }

  req := &openapi.OpenApiRequest{
    Headers: headers,
    Body: openapiutil.ParseToMap(body),
  }
  params := &openapi.Params{
    Action: tea.String("TestOpenApiRequest"),
    Version: tea.String("2020-11-20"),
    Protocol: tea.String("HTTPS"),
    Pathname: tea.String("/openapi/product/openApiRequest"),
    Method: tea.String("POST"),
    AuthType: tea.String("AK"),
    Style: tea.String("ROA"),
    ReqBodyType: tea.String("json"),
    BodyType: tea.String("json"),
  }
  _result = &TestOpenApiRequestResponse{}
  _body, _err := client.CallApi(params, req, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

