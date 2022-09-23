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

type AddImageRequest struct {
	CategoryId    *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop          *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IntAttr       *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicContent    *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName       *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId     *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StrAttr       *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetCategoryId(v int32) *AddImageRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageRequest) SetCrop(v bool) *AddImageRequest {
	s.Crop = &v
	return s
}

func (s *AddImageRequest) SetCustomContent(v string) *AddImageRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageRequest) SetInstanceName(v string) *AddImageRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageRequest) SetIntAttr(v int32) *AddImageRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageRequest) SetPicContent(v string) *AddImageRequest {
	s.PicContent = &v
	return s
}

func (s *AddImageRequest) SetPicName(v string) *AddImageRequest {
	s.PicName = &v
	return s
}

func (s *AddImageRequest) SetProductId(v string) *AddImageRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageRequest) SetRegion(v string) *AddImageRequest {
	s.Region = &v
	return s
}

func (s *AddImageRequest) SetStrAttr(v string) *AddImageRequest {
	s.StrAttr = &v
	return s
}

type AddImageResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PicInfo   *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetCode(v int32) *AddImageResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageResponseBody) SetMessage(v string) *AddImageResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageResponseBody) SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponseBodyPicInfo struct {
	CategoryId *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s AddImageResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetInstanceName(v string) *DeleteImageRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

func (s *DeleteImageRequest) SetProductId(v string) *DeleteImageRequest {
	s.ProductId = &v
	return s
}

type DeleteImageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type SearchImageRequest struct {
	CategoryId   *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop         *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicContent   *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicName      *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Start        *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageRequest) GoString() string {
	return s.String()
}

func (s *SearchImageRequest) SetCategoryId(v int32) *SearchImageRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageRequest) SetCrop(v bool) *SearchImageRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageRequest) SetFilter(v string) *SearchImageRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageRequest) SetInstanceName(v string) *SearchImageRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageRequest) SetNum(v int32) *SearchImageRequest {
	s.Num = &v
	return s
}

func (s *SearchImageRequest) SetPicContent(v string) *SearchImageRequest {
	s.PicContent = &v
	return s
}

func (s *SearchImageRequest) SetPicName(v string) *SearchImageRequest {
	s.PicName = &v
	return s
}

func (s *SearchImageRequest) SetProductId(v string) *SearchImageRequest {
	s.ProductId = &v
	return s
}

func (s *SearchImageRequest) SetRegion(v string) *SearchImageRequest {
	s.Region = &v
	return s
}

func (s *SearchImageRequest) SetStart(v int32) *SearchImageRequest {
	s.Start = &v
	return s
}

func (s *SearchImageRequest) SetType(v string) *SearchImageRequest {
	s.Type = &v
	return s
}

type SearchImageResponseBody struct {
	Auctions  []*SearchImageResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Head      *SearchImageResponseBodyHead       `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	Msg       *string                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	PicInfo   *SearchImageResponseBodyPicInfo    `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBody) SetAuctions(v []*SearchImageResponseBodyAuctions) *SearchImageResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageResponseBody) SetCode(v int32) *SearchImageResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageResponseBody) SetHead(v *SearchImageResponseBodyHead) *SearchImageResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageResponseBody) SetMsg(v string) *SearchImageResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageResponseBody) SetPicInfo(v *SearchImageResponseBodyPicInfo) *SearchImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageResponseBody) SetRequestId(v string) *SearchImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageResponseBody) SetSuccess(v bool) *SearchImageResponseBody {
	s.Success = &v
	return s
}

type SearchImageResponseBodyAuctions struct {
	CategoryId     *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CustomContent  *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	IntAttr        *int32  `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	PicName        *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	ProductId      *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	StrAttr        *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
}

func (s SearchImageResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyAuctions) SetCategoryId(v int32) *SearchImageResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetCustomContent(v string) *SearchImageResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetIntAttr(v int32) *SearchImageResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetPicName(v string) *SearchImageResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetProductId(v string) *SearchImageResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetSortExprValues(v string) *SearchImageResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageResponseBodyAuctions) SetStrAttr(v string) *SearchImageResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

type SearchImageResponseBodyHead struct {
	DocsFound  *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyHead) SetDocsFound(v int32) *SearchImageResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageResponseBodyHead) SetDocsReturn(v int32) *SearchImageResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageResponseBodyHead) SetSearchTime(v int32) *SearchImageResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchImageResponseBodyPicInfo struct {
	AllCategories []*SearchImageResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	CategoryId    *int32                                         `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Region        *string                                        `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyPicInfo) SetAllCategories(v []*SearchImageResponseBodyPicInfoAllCategories) *SearchImageResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageResponseBodyPicInfo) SetRegion(v string) *SearchImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchImageResponseBodyPicInfoAllCategories struct {
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

type SearchImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponse) GoString() string {
	return s.String()
}

func (s *SearchImageResponse) SetHeaders(v map[string]*string) *SearchImageResponse {
	s.Headers = v
	return s
}

func (s *SearchImageResponse) SetStatusCode(v int32) *SearchImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageResponse) SetBody(v *SearchImageResponseBody) *SearchImageResponse {
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

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageWithOptions(request *AddImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		body["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		body["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		body["StrAttr"] = request.StrAttr
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2019-03-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2019-03-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImage(request *SearchImageRequest) (_result *SearchImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchImageResponse{}
	_body, _err := client.SearchImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageWithOptions(request *SearchImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		body["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImage"),
		Version:     tea.String("2019-03-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
