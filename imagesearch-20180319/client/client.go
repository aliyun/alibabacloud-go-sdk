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

type AddItemRequest struct {
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s AddItemRequest) String() string {
	return tea.Prettify(s)
}

func (s AddItemRequest) GoString() string {
	return s.String()
}

func (s *AddItemRequest) SetInstanceName(v string) *AddItemRequest {
	s.InstanceName = &v
	return s
}

type AddItemResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddItemResponseBody) SetCode(v int32) *AddItemResponseBody {
	s.Code = &v
	return s
}

func (s *AddItemResponseBody) SetMessage(v string) *AddItemResponseBody {
	s.Message = &v
	return s
}

func (s *AddItemResponseBody) SetRequestId(v string) *AddItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddItemResponseBody) SetSuccess(v bool) *AddItemResponseBody {
	s.Success = &v
	return s
}

type AddItemResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddItemResponse) String() string {
	return tea.Prettify(s)
}

func (s AddItemResponse) GoString() string {
	return s.String()
}

func (s *AddItemResponse) SetHeaders(v map[string]*string) *AddItemResponse {
	s.Headers = v
	return s
}

func (s *AddItemResponse) SetStatusCode(v int32) *AddItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddItemResponse) SetBody(v *AddItemResponseBody) *AddItemResponse {
	s.Body = v
	return s
}

type DeleteItemRequest struct {
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s DeleteItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteItemRequest) SetInstanceName(v string) *DeleteItemRequest {
	s.InstanceName = &v
	return s
}

type DeleteItemResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteItemResponseBody) SetCode(v int32) *DeleteItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteItemResponseBody) SetMessage(v string) *DeleteItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteItemResponseBody) SetRequestId(v string) *DeleteItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteItemResponseBody) SetSuccess(v bool) *DeleteItemResponseBody {
	s.Success = &v
	return s
}

type DeleteItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteItemResponse) SetHeaders(v map[string]*string) *DeleteItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteItemResponse) SetStatusCode(v int32) *DeleteItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteItemResponse) SetBody(v *DeleteItemResponseBody) *DeleteItemResponse {
	s.Body = v
	return s
}

type SearchItemRequest struct {
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s SearchItemRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchItemRequest) GoString() string {
	return s.String()
}

func (s *SearchItemRequest) SetInstanceName(v string) *SearchItemRequest {
	s.InstanceName = &v
	return s
}

type SearchItemResponseBody struct {
	Auctions  *SearchItemResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Struct"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Head      *SearchItemResponseBodyHead     `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PicInfo   *SearchItemResponseBodyPicInfo  `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBody) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBody) SetAuctions(v *SearchItemResponseBodyAuctions) *SearchItemResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchItemResponseBody) SetCode(v int32) *SearchItemResponseBody {
	s.Code = &v
	return s
}

func (s *SearchItemResponseBody) SetHead(v *SearchItemResponseBodyHead) *SearchItemResponseBody {
	s.Head = v
	return s
}

func (s *SearchItemResponseBody) SetMessage(v string) *SearchItemResponseBody {
	s.Message = &v
	return s
}

func (s *SearchItemResponseBody) SetPicInfo(v *SearchItemResponseBodyPicInfo) *SearchItemResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchItemResponseBody) SetRequestId(v string) *SearchItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchItemResponseBody) SetSuccess(v bool) *SearchItemResponseBody {
	s.Success = &v
	return s
}

type SearchItemResponseBodyAuctions struct {
	Auction []*SearchItemResponseBodyAuctionsAuction `json:"Auction,omitempty" xml:"Auction,omitempty" type:"Repeated"`
}

func (s SearchItemResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyAuctions) SetAuction(v []*SearchItemResponseBodyAuctionsAuction) *SearchItemResponseBodyAuctions {
	s.Auction = v
	return s
}

type SearchItemResponseBodyAuctionsAuction struct {
	CatId          *string `json:"CatId,omitempty" xml:"CatId,omitempty"`
	CustContent    *string `json:"CustContent,omitempty" xml:"CustContent,omitempty"`
	ItemId         *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	PicName        *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
}

func (s SearchItemResponseBodyAuctionsAuction) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyAuctionsAuction) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyAuctionsAuction) SetCatId(v string) *SearchItemResponseBodyAuctionsAuction {
	s.CatId = &v
	return s
}

func (s *SearchItemResponseBodyAuctionsAuction) SetCustContent(v string) *SearchItemResponseBodyAuctionsAuction {
	s.CustContent = &v
	return s
}

func (s *SearchItemResponseBodyAuctionsAuction) SetItemId(v string) *SearchItemResponseBodyAuctionsAuction {
	s.ItemId = &v
	return s
}

func (s *SearchItemResponseBodyAuctionsAuction) SetPicName(v string) *SearchItemResponseBodyAuctionsAuction {
	s.PicName = &v
	return s
}

func (s *SearchItemResponseBodyAuctionsAuction) SetSortExprValues(v string) *SearchItemResponseBodyAuctionsAuction {
	s.SortExprValues = &v
	return s
}

type SearchItemResponseBodyHead struct {
	DocsFound  *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchItemResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyHead) SetDocsFound(v int32) *SearchItemResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchItemResponseBodyHead) SetDocsReturn(v int32) *SearchItemResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchItemResponseBodyHead) SetSearchTime(v int32) *SearchItemResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchItemResponseBodyPicInfo struct {
	AllCategory *SearchItemResponseBodyPicInfoAllCategory `json:"AllCategory,omitempty" xml:"AllCategory,omitempty" type:"Struct"`
	Category    *string                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	Region      *string                                   `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchItemResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyPicInfo) SetAllCategory(v *SearchItemResponseBodyPicInfoAllCategory) *SearchItemResponseBodyPicInfo {
	s.AllCategory = v
	return s
}

func (s *SearchItemResponseBodyPicInfo) SetCategory(v string) *SearchItemResponseBodyPicInfo {
	s.Category = &v
	return s
}

func (s *SearchItemResponseBodyPicInfo) SetRegion(v string) *SearchItemResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchItemResponseBodyPicInfoAllCategory struct {
	Category []*SearchItemResponseBodyPicInfoAllCategoryCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s SearchItemResponseBodyPicInfoAllCategory) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyPicInfoAllCategory) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyPicInfoAllCategory) SetCategory(v []*SearchItemResponseBodyPicInfoAllCategoryCategory) *SearchItemResponseBodyPicInfoAllCategory {
	s.Category = v
	return s
}

type SearchItemResponseBodyPicInfoAllCategoryCategory struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchItemResponseBodyPicInfoAllCategoryCategory) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponseBodyPicInfoAllCategoryCategory) GoString() string {
	return s.String()
}

func (s *SearchItemResponseBodyPicInfoAllCategoryCategory) SetId(v string) *SearchItemResponseBodyPicInfoAllCategoryCategory {
	s.Id = &v
	return s
}

func (s *SearchItemResponseBodyPicInfoAllCategoryCategory) SetName(v string) *SearchItemResponseBodyPicInfoAllCategoryCategory {
	s.Name = &v
	return s
}

type SearchItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchItemResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchItemResponse) GoString() string {
	return s.String()
}

func (s *SearchItemResponse) SetHeaders(v map[string]*string) *SearchItemResponse {
	s.Headers = v
	return s
}

func (s *SearchItemResponse) SetStatusCode(v int32) *SearchItemResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchItemResponse) SetBody(v *SearchItemResponseBody) *SearchItemResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddItem(request *AddItemRequest) (_result *AddItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddItemResponse{}
	_body, _err := client.AddItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddItemWithOptions(request *AddItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddItem"),
		Version:     tea.String("2018-03-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/item/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteItem(request *DeleteItemRequest) (_result *DeleteItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteItemResponse{}
	_body, _err := client.DeleteItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteItemWithOptions(request *DeleteItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteItem"),
		Version:     tea.String("2018-03-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/item/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchItem(request *SearchItemRequest) (_result *SearchItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchItemResponse{}
	_body, _err := client.SearchItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchItemWithOptions(request *SearchItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchItem"),
		Version:     tea.String("2018-03-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/item/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
