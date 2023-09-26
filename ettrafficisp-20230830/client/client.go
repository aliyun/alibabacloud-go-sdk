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

type AkDisableRequest struct {
	AccessKeyId   *string   `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeyName *string   `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	Permissions   []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s AkDisableRequest) String() string {
	return tea.Prettify(s)
}

func (s AkDisableRequest) GoString() string {
	return s.String()
}

func (s *AkDisableRequest) SetAccessKeyId(v string) *AkDisableRequest {
	s.AccessKeyId = &v
	return s
}

func (s *AkDisableRequest) SetAccessKeyName(v string) *AkDisableRequest {
	s.AccessKeyName = &v
	return s
}

func (s *AkDisableRequest) SetPermissions(v []*string) *AkDisableRequest {
	s.Permissions = v
	return s
}

type AkDisableResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AkDisableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AkDisableResponseBody) GoString() string {
	return s.String()
}

func (s *AkDisableResponseBody) SetCode(v string) *AkDisableResponseBody {
	s.Code = &v
	return s
}

func (s *AkDisableResponseBody) SetMessage(v string) *AkDisableResponseBody {
	s.Message = &v
	return s
}

func (s *AkDisableResponseBody) SetRequestId(v string) *AkDisableResponseBody {
	s.RequestId = &v
	return s
}

type AkDisableResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AkDisableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AkDisableResponse) String() string {
	return tea.Prettify(s)
}

func (s AkDisableResponse) GoString() string {
	return s.String()
}

func (s *AkDisableResponse) SetHeaders(v map[string]*string) *AkDisableResponse {
	s.Headers = v
	return s
}

func (s *AkDisableResponse) SetStatusCode(v int32) *AkDisableResponse {
	s.StatusCode = &v
	return s
}

func (s *AkDisableResponse) SetBody(v *AkDisableResponseBody) *AkDisableResponse {
	s.Body = v
	return s
}

type AkEnableRequest struct {
	AccessKeyId   *string   `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeyName *string   `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	Permissions   []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s AkEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s AkEnableRequest) GoString() string {
	return s.String()
}

func (s *AkEnableRequest) SetAccessKeyId(v string) *AkEnableRequest {
	s.AccessKeyId = &v
	return s
}

func (s *AkEnableRequest) SetAccessKeyName(v string) *AkEnableRequest {
	s.AccessKeyName = &v
	return s
}

func (s *AkEnableRequest) SetPermissions(v []*string) *AkEnableRequest {
	s.Permissions = v
	return s
}

type AkEnableResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AkEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AkEnableResponseBody) GoString() string {
	return s.String()
}

func (s *AkEnableResponseBody) SetCode(v string) *AkEnableResponseBody {
	s.Code = &v
	return s
}

func (s *AkEnableResponseBody) SetMessage(v string) *AkEnableResponseBody {
	s.Message = &v
	return s
}

func (s *AkEnableResponseBody) SetRequestId(v string) *AkEnableResponseBody {
	s.RequestId = &v
	return s
}

type AkEnableResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AkEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AkEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s AkEnableResponse) GoString() string {
	return s.String()
}

func (s *AkEnableResponse) SetHeaders(v map[string]*string) *AkEnableResponse {
	s.Headers = v
	return s
}

func (s *AkEnableResponse) SetStatusCode(v int32) *AkEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *AkEnableResponse) SetBody(v *AkEnableResponseBody) *AkEnableResponse {
	s.Body = v
	return s
}

type AkGenerateRequest struct {
	AccessKeyName *string   `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	Permissions   []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	UserId        *int64    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AkGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AkGenerateRequest) GoString() string {
	return s.String()
}

func (s *AkGenerateRequest) SetAccessKeyName(v string) *AkGenerateRequest {
	s.AccessKeyName = &v
	return s
}

func (s *AkGenerateRequest) SetPermissions(v []*string) *AkGenerateRequest {
	s.Permissions = v
	return s
}

func (s *AkGenerateRequest) SetUserId(v int64) *AkGenerateRequest {
	s.UserId = &v
	return s
}

type AkGenerateResponseBody struct {
	Code      *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *AkGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AkGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AkGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AkGenerateResponseBody) SetCode(v string) *AkGenerateResponseBody {
	s.Code = &v
	return s
}

func (s *AkGenerateResponseBody) SetData(v *AkGenerateResponseBodyData) *AkGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AkGenerateResponseBody) SetMessage(v string) *AkGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *AkGenerateResponseBody) SetRequestId(v string) *AkGenerateResponseBody {
	s.RequestId = &v
	return s
}

type AkGenerateResponseBodyData struct {
	// Access Key ID
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeyName   *string `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AkGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AkGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AkGenerateResponseBodyData) SetAccessKeyId(v string) *AkGenerateResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *AkGenerateResponseBodyData) SetAccessKeyName(v string) *AkGenerateResponseBodyData {
	s.AccessKeyName = &v
	return s
}

func (s *AkGenerateResponseBodyData) SetAccessKeySecret(v string) *AkGenerateResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *AkGenerateResponseBodyData) SetUserId(v string) *AkGenerateResponseBodyData {
	s.UserId = &v
	return s
}

type AkGenerateResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AkGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AkGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AkGenerateResponse) GoString() string {
	return s.String()
}

func (s *AkGenerateResponse) SetHeaders(v map[string]*string) *AkGenerateResponse {
	s.Headers = v
	return s
}

func (s *AkGenerateResponse) SetStatusCode(v int32) *AkGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AkGenerateResponse) SetBody(v *AkGenerateResponseBody) *AkGenerateResponse {
	s.Body = v
	return s
}

type AkListPageRequest struct {
	Page  *int32 `json:"page,omitempty" xml:"page,omitempty"`
	Size  *int32 `json:"size,omitempty" xml:"size,omitempty"`
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s AkListPageRequest) String() string {
	return tea.Prettify(s)
}

func (s AkListPageRequest) GoString() string {
	return s.String()
}

func (s *AkListPageRequest) SetPage(v int32) *AkListPageRequest {
	s.Page = &v
	return s
}

func (s *AkListPageRequest) SetSize(v int32) *AkListPageRequest {
	s.Size = &v
	return s
}

func (s *AkListPageRequest) SetStart(v int32) *AkListPageRequest {
	s.Start = &v
	return s
}

type AkListPageResponseBody struct {
	Code      *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *AkListPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AkListPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AkListPageResponseBody) GoString() string {
	return s.String()
}

func (s *AkListPageResponseBody) SetCode(v string) *AkListPageResponseBody {
	s.Code = &v
	return s
}

func (s *AkListPageResponseBody) SetData(v *AkListPageResponseBodyData) *AkListPageResponseBody {
	s.Data = v
	return s
}

func (s *AkListPageResponseBody) SetMessage(v string) *AkListPageResponseBody {
	s.Message = &v
	return s
}

func (s *AkListPageResponseBody) SetRequestId(v string) *AkListPageResponseBody {
	s.RequestId = &v
	return s
}

type AkListPageResponseBodyData struct {
	Count       *int32                                `json:"count,omitempty" xml:"count,omitempty"`
	CurrentPage *int32                                `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	DataList    []*AkListPageResponseBodyDataDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	TotalPage   *int32                                `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s AkListPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AkListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *AkListPageResponseBodyData) SetCount(v int32) *AkListPageResponseBodyData {
	s.Count = &v
	return s
}

func (s *AkListPageResponseBodyData) SetCurrentPage(v int32) *AkListPageResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *AkListPageResponseBodyData) SetDataList(v []*AkListPageResponseBodyDataDataList) *AkListPageResponseBodyData {
	s.DataList = v
	return s
}

func (s *AkListPageResponseBodyData) SetTotalPage(v int32) *AkListPageResponseBodyData {
	s.TotalPage = &v
	return s
}

type AkListPageResponseBodyDataDataList struct {
	AccessKeyId     *string                                          `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeyName   *string                                          `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	AccessKeySecret *string                                          `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Active          *int32                                           `json:"active,omitempty" xml:"active,omitempty"`
	GmtCreate       *string                                          `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModify       *string                                          `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	LastCallTime    *string                                          `json:"lastCallTime,omitempty" xml:"lastCallTime,omitempty"`
	Permissions     []*AkListPageResponseBodyDataDataListPermissions `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s AkListPageResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s AkListPageResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *AkListPageResponseBodyDataDataList) SetAccessKeyId(v string) *AkListPageResponseBodyDataDataList {
	s.AccessKeyId = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetAccessKeyName(v string) *AkListPageResponseBodyDataDataList {
	s.AccessKeyName = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetAccessKeySecret(v string) *AkListPageResponseBodyDataDataList {
	s.AccessKeySecret = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetActive(v int32) *AkListPageResponseBodyDataDataList {
	s.Active = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetGmtCreate(v string) *AkListPageResponseBodyDataDataList {
	s.GmtCreate = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetGmtModify(v string) *AkListPageResponseBodyDataDataList {
	s.GmtModify = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetLastCallTime(v string) *AkListPageResponseBodyDataDataList {
	s.LastCallTime = &v
	return s
}

func (s *AkListPageResponseBodyDataDataList) SetPermissions(v []*AkListPageResponseBodyDataDataListPermissions) *AkListPageResponseBodyDataDataList {
	s.Permissions = v
	return s
}

type AkListPageResponseBodyDataDataListPermissions struct {
	AuthTime       *string `json:"authTime,omitempty" xml:"authTime,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Group          *string `json:"group,omitempty" xml:"group,omitempty"`
	PermissionCode *string `json:"permissionCode,omitempty" xml:"permissionCode,omitempty"`
	PermissionName *string `json:"permissionName,omitempty" xml:"permissionName,omitempty"`
}

func (s AkListPageResponseBodyDataDataListPermissions) String() string {
	return tea.Prettify(s)
}

func (s AkListPageResponseBodyDataDataListPermissions) GoString() string {
	return s.String()
}

func (s *AkListPageResponseBodyDataDataListPermissions) SetAuthTime(v string) *AkListPageResponseBodyDataDataListPermissions {
	s.AuthTime = &v
	return s
}

func (s *AkListPageResponseBodyDataDataListPermissions) SetDescription(v string) *AkListPageResponseBodyDataDataListPermissions {
	s.Description = &v
	return s
}

func (s *AkListPageResponseBodyDataDataListPermissions) SetGroup(v string) *AkListPageResponseBodyDataDataListPermissions {
	s.Group = &v
	return s
}

func (s *AkListPageResponseBodyDataDataListPermissions) SetPermissionCode(v string) *AkListPageResponseBodyDataDataListPermissions {
	s.PermissionCode = &v
	return s
}

func (s *AkListPageResponseBodyDataDataListPermissions) SetPermissionName(v string) *AkListPageResponseBodyDataDataListPermissions {
	s.PermissionName = &v
	return s
}

type AkListPageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AkListPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AkListPageResponse) String() string {
	return tea.Prettify(s)
}

func (s AkListPageResponse) GoString() string {
	return s.String()
}

func (s *AkListPageResponse) SetHeaders(v map[string]*string) *AkListPageResponse {
	s.Headers = v
	return s
}

func (s *AkListPageResponse) SetStatusCode(v int32) *AkListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *AkListPageResponse) SetBody(v *AkListPageResponseBody) *AkListPageResponse {
	s.Body = v
	return s
}

type AkUpdateRequest struct {
	AccessKeyId   *string   `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeyName *string   `json:"accessKeyName,omitempty" xml:"accessKeyName,omitempty"`
	Permissions   []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s AkUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s AkUpdateRequest) GoString() string {
	return s.String()
}

func (s *AkUpdateRequest) SetAccessKeyId(v string) *AkUpdateRequest {
	s.AccessKeyId = &v
	return s
}

func (s *AkUpdateRequest) SetAccessKeyName(v string) *AkUpdateRequest {
	s.AccessKeyName = &v
	return s
}

func (s *AkUpdateRequest) SetPermissions(v []*string) *AkUpdateRequest {
	s.Permissions = v
	return s
}

type AkUpdateResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AkUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AkUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *AkUpdateResponseBody) SetCode(v string) *AkUpdateResponseBody {
	s.Code = &v
	return s
}

func (s *AkUpdateResponseBody) SetMessage(v string) *AkUpdateResponseBody {
	s.Message = &v
	return s
}

func (s *AkUpdateResponseBody) SetRequestId(v string) *AkUpdateResponseBody {
	s.RequestId = &v
	return s
}

type AkUpdateResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AkUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AkUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s AkUpdateResponse) GoString() string {
	return s.String()
}

func (s *AkUpdateResponse) SetHeaders(v map[string]*string) *AkUpdateResponse {
	s.Headers = v
	return s
}

func (s *AkUpdateResponse) SetStatusCode(v int32) *AkUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *AkUpdateResponse) SetBody(v *AkUpdateResponseBody) *AkUpdateResponse {
	s.Body = v
	return s
}

type DimGroupResponseBody struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DimGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DimGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DimGroupResponseBody) SetCode(v string) *DimGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DimGroupResponseBody) SetData(v interface{}) *DimGroupResponseBody {
	s.Data = v
	return s
}

func (s *DimGroupResponseBody) SetMessage(v string) *DimGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DimGroupResponseBody) SetRequestId(v string) *DimGroupResponseBody {
	s.RequestId = &v
	return s
}

type DimGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DimGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DimGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DimGroupResponse) GoString() string {
	return s.String()
}

func (s *DimGroupResponse) SetHeaders(v map[string]*string) *DimGroupResponse {
	s.Headers = v
	return s
}

func (s *DimGroupResponse) SetStatusCode(v int32) *DimGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DimGroupResponse) SetBody(v *DimGroupResponseBody) *DimGroupResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ettrafficisp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AkDisableWithOptions(request *AkDisableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AkDisableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeyName)) {
		body["accessKeyName"] = request.AccessKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AkDisable"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/ak/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AkDisableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AkDisable(request *AkDisableRequest) (_result *AkDisableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AkDisableResponse{}
	_body, _err := client.AkDisableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AkEnableWithOptions(request *AkEnableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AkEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeyName)) {
		body["accessKeyName"] = request.AccessKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AkEnable"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/ak/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AkEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AkEnable(request *AkEnableRequest) (_result *AkEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AkEnableResponse{}
	_body, _err := client.AkEnableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AkGenerateWithOptions(request *AkGenerateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AkGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyName)) {
		body["accessKeyName"] = request.AccessKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AkGenerate"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/ak/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AkGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AkGenerate(request *AkGenerateRequest) (_result *AkGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AkGenerateResponse{}
	_body, _err := client.AkGenerateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AkListPageWithOptions(request *AkListPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AkListPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AkListPage"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/ak/listPage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AkListPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AkListPage(request *AkListPageRequest) (_result *AkListPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AkListPageResponse{}
	_body, _err := client.AkListPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AkUpdateWithOptions(request *AkUpdateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AkUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeyName)) {
		body["accessKeyName"] = request.AccessKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AkUpdate"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/ak/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AkUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AkUpdate(request *AkUpdateRequest) (_result *AkUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AkUpdateResponse{}
	_body, _err := client.AkUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DimGroupWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DimGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DimGroup"),
		Version:     tea.String("2023-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/console/dim/group"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DimGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DimGroup() (_result *DimGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DimGroupResponse{}
	_body, _err := client.DimGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
