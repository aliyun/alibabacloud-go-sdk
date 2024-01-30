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

type CheckUserAuthResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserAuthResponseBody) SetCode(v string) *CheckUserAuthResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUserAuthResponseBody) SetData(v string) *CheckUserAuthResponseBody {
	s.Data = &v
	return s
}

func (s *CheckUserAuthResponseBody) SetMessage(v string) *CheckUserAuthResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUserAuthResponseBody) SetRequestId(v string) *CheckUserAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserAuthResponseBody) SetSuccess(v bool) *CheckUserAuthResponseBody {
	s.Success = &v
	return s
}

type CheckUserAuthResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckUserAuthResponse) SetHeaders(v map[string]*string) *CheckUserAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckUserAuthResponse) SetStatusCode(v int32) *CheckUserAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserAuthResponse) SetBody(v *CheckUserAuthResponseBody) *CheckUserAuthResponse {
	s.Body = v
	return s
}

type ListResourcesByTagRequest struct {
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TagKeys    []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TagValues  []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListResourcesByTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagRequest) SetPageNumber(v int32) *ListResourcesByTagRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesByTagRequest) SetPageSize(v int32) *ListResourcesByTagRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesByTagRequest) SetTagKeys(v []*string) *ListResourcesByTagRequest {
	s.TagKeys = v
	return s
}

func (s *ListResourcesByTagRequest) SetTagValues(v []*string) *ListResourcesByTagRequest {
	s.TagValues = v
	return s
}

type ListResourcesByTagResponseBody struct {
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNumber *int32  `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s ListResourcesByTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBody) SetData(v string) *ListResourcesByTagResponseBody {
	s.Data = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetPageNumber(v int32) *ListResourcesByTagResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetPageSize(v int32) *ListResourcesByTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetRequestId(v string) *ListResourcesByTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetTotalNumber(v int32) *ListResourcesByTagResponseBody {
	s.TotalNumber = &v
	return s
}

type ListResourcesByTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesByTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesByTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponse) SetHeaders(v map[string]*string) *ListResourcesByTagResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesByTagResponse) SetStatusCode(v int32) *ListResourcesByTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesByTagResponse) SetBody(v *ListResourcesByTagResponseBody) *ListResourcesByTagResponse {
	s.Body = v
	return s
}

type ListTagsAllResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsAllResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsAllResponseBody) SetData(v string) *ListTagsAllResponseBody {
	s.Data = &v
	return s
}

func (s *ListTagsAllResponseBody) SetRequestId(v string) *ListTagsAllResponseBody {
	s.RequestId = &v
	return s
}

type ListTagsAllResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsAllResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsAllResponse) GoString() string {
	return s.String()
}

func (s *ListTagsAllResponse) SetHeaders(v map[string]*string) *ListTagsAllResponse {
	s.Headers = v
	return s
}

func (s *ListTagsAllResponse) SetStatusCode(v int32) *ListTagsAllResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsAllResponse) SetBody(v *ListTagsAllResponseBody) *ListTagsAllResponse {
	s.Body = v
	return s
}

type ListTagsByResourceRequest struct {
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s ListTagsByResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsByResourceRequest) GoString() string {
	return s.String()
}

func (s *ListTagsByResourceRequest) SetResourceIds(v []*string) *ListTagsByResourceRequest {
	s.ResourceIds = v
	return s
}

type ListTagsByResourceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsByResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsByResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsByResourceResponseBody) SetData(v string) *ListTagsByResourceResponseBody {
	s.Data = &v
	return s
}

func (s *ListTagsByResourceResponseBody) SetRequestId(v string) *ListTagsByResourceResponseBody {
	s.RequestId = &v
	return s
}

type ListTagsByResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsByResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsByResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsByResourceResponse) GoString() string {
	return s.String()
}

func (s *ListTagsByResourceResponse) SetHeaders(v map[string]*string) *ListTagsByResourceResponse {
	s.Headers = v
	return s
}

func (s *ListTagsByResourceResponse) SetStatusCode(v int32) *ListTagsByResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsByResourceResponse) SetBody(v *ListTagsByResourceResponseBody) *ListTagsByResourceResponse {
	s.Body = v
	return s
}

type MetaDatabaseSearchRequest struct {
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	Start     *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s MetaDatabaseSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s MetaDatabaseSearchRequest) GoString() string {
	return s.String()
}

func (s *MetaDatabaseSearchRequest) SetLimit(v int64) *MetaDatabaseSearchRequest {
	s.Limit = &v
	return s
}

func (s *MetaDatabaseSearchRequest) SetSearchKey(v string) *MetaDatabaseSearchRequest {
	s.SearchKey = &v
	return s
}

func (s *MetaDatabaseSearchRequest) SetStart(v int64) *MetaDatabaseSearchRequest {
	s.Start = &v
	return s
}

type MetaDatabaseSearchResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MetaDatabaseSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MetaDatabaseSearchResponseBody) GoString() string {
	return s.String()
}

func (s *MetaDatabaseSearchResponseBody) SetCode(v string) *MetaDatabaseSearchResponseBody {
	s.Code = &v
	return s
}

func (s *MetaDatabaseSearchResponseBody) SetData(v string) *MetaDatabaseSearchResponseBody {
	s.Data = &v
	return s
}

func (s *MetaDatabaseSearchResponseBody) SetMessage(v string) *MetaDatabaseSearchResponseBody {
	s.Message = &v
	return s
}

func (s *MetaDatabaseSearchResponseBody) SetRequestId(v string) *MetaDatabaseSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *MetaDatabaseSearchResponseBody) SetSuccess(v bool) *MetaDatabaseSearchResponseBody {
	s.Success = &v
	return s
}

type MetaDatabaseSearchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MetaDatabaseSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MetaDatabaseSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s MetaDatabaseSearchResponse) GoString() string {
	return s.String()
}

func (s *MetaDatabaseSearchResponse) SetHeaders(v map[string]*string) *MetaDatabaseSearchResponse {
	s.Headers = v
	return s
}

func (s *MetaDatabaseSearchResponse) SetStatusCode(v int32) *MetaDatabaseSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *MetaDatabaseSearchResponse) SetBody(v *MetaDatabaseSearchResponseBody) *MetaDatabaseSearchResponse {
	s.Body = v
	return s
}

type MetaSearchRequest struct {
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	Start     *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s MetaSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s MetaSearchRequest) GoString() string {
	return s.String()
}

func (s *MetaSearchRequest) SetLimit(v int64) *MetaSearchRequest {
	s.Limit = &v
	return s
}

func (s *MetaSearchRequest) SetSearchKey(v string) *MetaSearchRequest {
	s.SearchKey = &v
	return s
}

func (s *MetaSearchRequest) SetStart(v int64) *MetaSearchRequest {
	s.Start = &v
	return s
}

type MetaSearchResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MetaSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MetaSearchResponseBody) GoString() string {
	return s.String()
}

func (s *MetaSearchResponseBody) SetCode(v string) *MetaSearchResponseBody {
	s.Code = &v
	return s
}

func (s *MetaSearchResponseBody) SetData(v string) *MetaSearchResponseBody {
	s.Data = &v
	return s
}

func (s *MetaSearchResponseBody) SetMessage(v string) *MetaSearchResponseBody {
	s.Message = &v
	return s
}

func (s *MetaSearchResponseBody) SetRequestId(v string) *MetaSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *MetaSearchResponseBody) SetSuccess(v bool) *MetaSearchResponseBody {
	s.Success = &v
	return s
}

type MetaSearchResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MetaSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MetaSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s MetaSearchResponse) GoString() string {
	return s.String()
}

func (s *MetaSearchResponse) SetHeaders(v map[string]*string) *MetaSearchResponse {
	s.Headers = v
	return s
}

func (s *MetaSearchResponse) SetStatusCode(v int32) *MetaSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *MetaSearchResponse) SetBody(v *MetaSearchResponseBody) *MetaSearchResponse {
	s.Body = v
	return s
}

type UpdateUserAuthRequest struct {
	GrantState *int64 `json:"GrantState,omitempty" xml:"GrantState,omitempty"`
}

func (s UpdateUserAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserAuthRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthRequest) SetGrantState(v int64) *UpdateUserAuthRequest {
	s.GrantState = &v
	return s
}

type UpdateUserAuthResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthResponseBody) SetCode(v string) *UpdateUserAuthResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserAuthResponseBody) SetData(v string) *UpdateUserAuthResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserAuthResponseBody) SetMessage(v string) *UpdateUserAuthResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserAuthResponseBody) SetRequestId(v string) *UpdateUserAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserAuthResponseBody) SetSuccess(v bool) *UpdateUserAuthResponseBody {
	s.Success = &v
	return s
}

type UpdateUserAuthResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserAuthResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthResponse) SetHeaders(v map[string]*string) *UpdateUserAuthResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserAuthResponse) SetStatusCode(v int32) *UpdateUserAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserAuthResponse) SetBody(v *UpdateUserAuthResponseBody) *UpdateUserAuthResponse {
	s.Body = v
	return s
}

type YaochiSuggestRequest struct {
	SuggestText *string `json:"SuggestText,omitempty" xml:"SuggestText,omitempty"`
}

func (s YaochiSuggestRequest) String() string {
	return tea.Prettify(s)
}

func (s YaochiSuggestRequest) GoString() string {
	return s.String()
}

func (s *YaochiSuggestRequest) SetSuggestText(v string) *YaochiSuggestRequest {
	s.SuggestText = &v
	return s
}

type YaochiSuggestResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s YaochiSuggestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s YaochiSuggestResponseBody) GoString() string {
	return s.String()
}

func (s *YaochiSuggestResponseBody) SetCode(v string) *YaochiSuggestResponseBody {
	s.Code = &v
	return s
}

func (s *YaochiSuggestResponseBody) SetData(v string) *YaochiSuggestResponseBody {
	s.Data = &v
	return s
}

func (s *YaochiSuggestResponseBody) SetMessage(v string) *YaochiSuggestResponseBody {
	s.Message = &v
	return s
}

func (s *YaochiSuggestResponseBody) SetRequestId(v string) *YaochiSuggestResponseBody {
	s.RequestId = &v
	return s
}

func (s *YaochiSuggestResponseBody) SetSuccess(v bool) *YaochiSuggestResponseBody {
	s.Success = &v
	return s
}

type YaochiSuggestResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *YaochiSuggestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s YaochiSuggestResponse) String() string {
	return tea.Prettify(s)
}

func (s YaochiSuggestResponse) GoString() string {
	return s.String()
}

func (s *YaochiSuggestResponse) SetHeaders(v map[string]*string) *YaochiSuggestResponse {
	s.Headers = v
	return s
}

func (s *YaochiSuggestResponse) SetStatusCode(v int32) *YaochiSuggestResponse {
	s.StatusCode = &v
	return s
}

func (s *YaochiSuggestResponse) SetBody(v *YaochiSuggestResponseBody) *YaochiSuggestResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dms-yaochi-portal"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckUserAuthWithOptions(runtime *util.RuntimeOptions) (_result *CheckUserAuthResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckUserAuth"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserAuth() (_result *CheckUserAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserAuthResponse{}
	_body, _err := client.CheckUserAuthWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesByTagWithOptions(request *ListResourcesByTagRequest, runtime *util.RuntimeOptions) (_result *ListResourcesByTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["TagKeys"] = request.TagKeys
	}

	if !tea.BoolValue(util.IsUnset(request.TagValues)) {
		query["TagValues"] = request.TagValues
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcesByTag"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesByTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourcesByTag(request *ListResourcesByTagRequest) (_result *ListResourcesByTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesByTagResponse{}
	_body, _err := client.ListResourcesByTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsAllWithOptions(runtime *util.RuntimeOptions) (_result *ListTagsAllResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListTagsAll"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagsAll() (_result *ListTagsAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsAllResponse{}
	_body, _err := client.ListTagsAllWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsByResourceWithOptions(request *ListTagsByResourceRequest, runtime *util.RuntimeOptions) (_result *ListTagsByResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagsByResource"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsByResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagsByResource(request *ListTagsByResourceRequest) (_result *ListTagsByResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsByResourceResponse{}
	_body, _err := client.ListTagsByResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MetaDatabaseSearchWithOptions(request *MetaDatabaseSearchRequest, runtime *util.RuntimeOptions) (_result *MetaDatabaseSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MetaDatabaseSearch"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MetaDatabaseSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MetaDatabaseSearch(request *MetaDatabaseSearchRequest) (_result *MetaDatabaseSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MetaDatabaseSearchResponse{}
	_body, _err := client.MetaDatabaseSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MetaSearchWithOptions(request *MetaSearchRequest, runtime *util.RuntimeOptions) (_result *MetaSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MetaSearch"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MetaSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MetaSearch(request *MetaSearchRequest) (_result *MetaSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MetaSearchResponse{}
	_body, _err := client.MetaSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserAuthWithOptions(request *UpdateUserAuthRequest, runtime *util.RuntimeOptions) (_result *UpdateUserAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GrantState)) {
		query["GrantState"] = request.GrantState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserAuth"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserAuth(request *UpdateUserAuthRequest) (_result *UpdateUserAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserAuthResponse{}
	_body, _err := client.UpdateUserAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) YaochiSuggestWithOptions(request *YaochiSuggestRequest, runtime *util.RuntimeOptions) (_result *YaochiSuggestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SuggestText)) {
		query["SuggestText"] = request.SuggestText
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("YaochiSuggest"),
		Version:     tea.String("2023-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &YaochiSuggestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) YaochiSuggest(request *YaochiSuggestRequest) (_result *YaochiSuggestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &YaochiSuggestResponse{}
	_body, _err := client.YaochiSuggestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
