// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetOpenApiListRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpenApiListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiListRequest) GoString() string {
	return s.String()
}

func (s *GetOpenApiListRequest) SetRequestId(v string) *GetOpenApiListRequest {
	s.RequestId = &v
	return s
}

type GetOpenApiListResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OpenApiString *string `json:"OpenApiString,omitempty" xml:"OpenApiString,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpenApiListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenApiListResponseBody) SetCode(v string) *GetOpenApiListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetMessage(v string) *GetOpenApiListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetOpenApiString(v string) *GetOpenApiListResponseBody {
	s.OpenApiString = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetRequestId(v string) *GetOpenApiListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetSuccess(v bool) *GetOpenApiListResponseBody {
	s.Success = &v
	return s
}

type GetOpenApiListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenApiListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiListResponse) GoString() string {
	return s.String()
}

func (s *GetOpenApiListResponse) SetHeaders(v map[string]*string) *GetOpenApiListResponse {
	s.Headers = v
	return s
}

func (s *GetOpenApiListResponse) SetStatusCode(v int32) *GetOpenApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenApiListResponse) SetBody(v *GetOpenApiListResponseBody) *GetOpenApiListResponse {
	s.Body = v
	return s
}

type ListConsoleProductResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListConsoleProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListConsoleProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsoleProductResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponseBody) SetCode(v string) *ListConsoleProductResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsoleProductResponseBody) SetData(v []*ListConsoleProductResponseBodyData) *ListConsoleProductResponseBody {
	s.Data = v
	return s
}

func (s *ListConsoleProductResponseBody) SetRequestId(v string) *ListConsoleProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsoleProductResponseBody) SetSuccess(v bool) *ListConsoleProductResponseBody {
	s.Success = &v
	return s
}

type ListConsoleProductResponseBodyData struct {
	BelongedCategory    *string   `json:"BelongedCategory,omitempty" xml:"BelongedCategory,omitempty"`
	Categories          []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	ChannelLinks        []*string `json:"ChannelLinks,omitempty" xml:"ChannelLinks,omitempty" type:"Repeated"`
	DocId               *string   `json:"DocId,omitempty" xml:"DocId,omitempty"`
	Keywords            []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Names               *string   `json:"Names,omitempty" xml:"Names,omitempty"`
	Pinyin              *string   `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	ProductId           *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RelatedConsoleAppId *string   `json:"RelatedConsoleAppId,omitempty" xml:"RelatedConsoleAppId,omitempty"`
	RelatedPipId        *string   `json:"RelatedPipId,omitempty" xml:"RelatedPipId,omitempty"`
	ShowInNav           *bool     `json:"ShowInNav,omitempty" xml:"ShowInNav,omitempty"`
	SupportedAccounts   []*string `json:"SupportedAccounts,omitempty" xml:"SupportedAccounts,omitempty" type:"Repeated"`
	SupportedChannels   []*string `json:"SupportedChannels,omitempty" xml:"SupportedChannels,omitempty" type:"Repeated"`
	Tag                 *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagExpireTime       *string   `json:"TagExpireTime,omitempty" xml:"TagExpireTime,omitempty"`
}

func (s ListConsoleProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsoleProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponseBodyData) SetBelongedCategory(v string) *ListConsoleProductResponseBodyData {
	s.BelongedCategory = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetCategories(v []*string) *ListConsoleProductResponseBodyData {
	s.Categories = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetChannelLinks(v []*string) *ListConsoleProductResponseBodyData {
	s.ChannelLinks = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetDocId(v string) *ListConsoleProductResponseBodyData {
	s.DocId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetKeywords(v []*string) *ListConsoleProductResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetNames(v string) *ListConsoleProductResponseBodyData {
	s.Names = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetPinyin(v string) *ListConsoleProductResponseBodyData {
	s.Pinyin = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetProductId(v string) *ListConsoleProductResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetRelatedConsoleAppId(v string) *ListConsoleProductResponseBodyData {
	s.RelatedConsoleAppId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetRelatedPipId(v string) *ListConsoleProductResponseBodyData {
	s.RelatedPipId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetShowInNav(v bool) *ListConsoleProductResponseBodyData {
	s.ShowInNav = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetSupportedAccounts(v []*string) *ListConsoleProductResponseBodyData {
	s.SupportedAccounts = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetSupportedChannels(v []*string) *ListConsoleProductResponseBodyData {
	s.SupportedChannels = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetTag(v string) *ListConsoleProductResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetTagExpireTime(v string) *ListConsoleProductResponseBodyData {
	s.TagExpireTime = &v
	return s
}

type ListConsoleProductResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsoleProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsoleProductResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsoleProductResponse) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponse) SetHeaders(v map[string]*string) *ListConsoleProductResponse {
	s.Headers = v
	return s
}

func (s *ListConsoleProductResponse) SetStatusCode(v int32) *ListConsoleProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsoleProductResponse) SetBody(v *ListConsoleProductResponseBody) *ListConsoleProductResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("consolecs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - GetOpenApiListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenApiListResponse
func (client *Client) GetOpenApiListWithOptions(request *GetOpenApiListRequest, runtime *util.RuntimeOptions) (_result *GetOpenApiListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpenApiList"),
		Version:     tea.String("2016-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenApiListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOpenApiListRequest
//
// @return GetOpenApiListResponse
func (client *Client) GetOpenApiList(request *GetOpenApiListRequest) (_result *GetOpenApiListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenApiListResponse{}
	_body, _err := client.GetOpenApiListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListConsoleProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsoleProductResponse
func (client *Client) ListConsoleProductWithOptions(runtime *util.RuntimeOptions) (_result *ListConsoleProductResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListConsoleProduct"),
		Version:     tea.String("2016-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsoleProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return ListConsoleProductResponse
func (client *Client) ListConsoleProduct() (_result *ListConsoleProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConsoleProductResponse{}
	_body, _err := client.ListConsoleProductWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
