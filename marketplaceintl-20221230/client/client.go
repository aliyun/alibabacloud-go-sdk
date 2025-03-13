// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeSellerInstancesRequest struct {
	// example:
	//
	// 5000002763123
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// OPENED
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5322460655123456
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSellerInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSellerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesRequest) SetInstanceId(v int64) *DescribeSellerInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetInstanceStatus(v string) *DescribeSellerInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetPageIndex(v int32) *DescribeSellerInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetPageSize(v int64) *DescribeSellerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetUserId(v int64) *DescribeSellerInstancesRequest {
	s.UserId = &v
	return s
}

type DescribeSellerInstancesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// fatal
	//
	// example:
	//
	// False
	Fatal *bool `json:"Fatal,omitempty" xml:"Fatal,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeSellerInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 103
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSellerInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSellerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponseBody) SetCode(v string) *DescribeSellerInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetCount(v int64) *DescribeSellerInstancesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetFatal(v bool) *DescribeSellerInstancesResponseBody {
	s.Fatal = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetMessage(v string) *DescribeSellerInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetPageNumber(v int64) *DescribeSellerInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetPageSize(v int64) *DescribeSellerInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetRequestId(v string) *DescribeSellerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetResult(v []*DescribeSellerInstancesResponseBodyResult) *DescribeSellerInstancesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetSuccess(v bool) *DescribeSellerInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetVersion(v string) *DescribeSellerInstancesResponseBody {
	s.Version = &v
	return s
}

type DescribeSellerInstancesResponseBodyResult struct {
	// example:
	//
	// {\\"authUrl\\":\\"https://marketplace.alibabacloud.com/\\"}
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty"`
	// example:
	//
	// 1
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// sgcmgj000356
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 1741752000000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// example:
	//
	// {\\"userName\\":\\"marketplace\\"}
	HostInfo *string `json:"HostInfo,omitempty" xml:"HostInfo,omitempty"`
	// example:
	//
	// {\\"userName\\":\\"marketplace\\"}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// example:
	//
	// 5000002763123
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// OPENED
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// 5322460655123456
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSellerInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSellerInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponseBodyResult) SetAppInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.AppInfo = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetChargeType(v int32) *DescribeSellerInstancesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetCommodityCode(v string) *DescribeSellerInstancesResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetCreatedOn(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.CreatedOn = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetHostInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.HostInfo = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.Info = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInstanceId(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInstanceStatus(v string) *DescribeSellerInstancesResponseBodyResult {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetUserId(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.UserId = &v
	return s
}

type DescribeSellerInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSellerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSellerInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSellerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponse) SetHeaders(v map[string]*string) *DescribeSellerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSellerInstancesResponse) SetStatusCode(v int32) *DescribeSellerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSellerInstancesResponse) SetBody(v *DescribeSellerInstancesResponseBody) *DescribeSellerInstancesResponse {
	s.Body = v
	return s
}

type NoticeInstanceUserRequest struct {
	// example:
	//
	// 5000000264872
	InstanceId  *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NoticeParam *string `json:"NoticeParam,omitempty" xml:"NoticeParam,omitempty"`
	// example:
	//
	// 1
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
}

func (s NoticeInstanceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s NoticeInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserRequest) SetInstanceId(v int64) *NoticeInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *NoticeInstanceUserRequest) SetNoticeParam(v string) *NoticeInstanceUserRequest {
	s.NoticeParam = &v
	return s
}

func (s *NoticeInstanceUserRequest) SetNoticeType(v int64) *NoticeInstanceUserRequest {
	s.NoticeType = &v
	return s
}

type NoticeInstanceUserResponseBody struct {
	AccessDeniedDetail *NoticeInstanceUserResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NoticeInstanceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NoticeInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponseBody) SetAccessDeniedDetail(v *NoticeInstanceUserResponseBodyAccessDeniedDetail) *NoticeInstanceUserResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetCode(v string) *NoticeInstanceUserResponseBody {
	s.Code = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetMessage(v string) *NoticeInstanceUserResponseBody {
	s.Message = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetRequestId(v string) *NoticeInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetResult(v bool) *NoticeInstanceUserResponseBody {
	s.Result = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetSuccess(v bool) *NoticeInstanceUserResponseBody {
	s.Success = &v
	return s
}

type NoticeInstanceUserResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s NoticeInstanceUserResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s NoticeInstanceUserResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthAction(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetPolicyType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type NoticeInstanceUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NoticeInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NoticeInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s NoticeInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponse) SetHeaders(v map[string]*string) *NoticeInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *NoticeInstanceUserResponse) SetStatusCode(v int32) *NoticeInstanceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *NoticeInstanceUserResponse) SetBody(v *NoticeInstanceUserResponseBody) *NoticeInstanceUserResponse {
	s.Body = v
	return s
}

type PushMeteringDataRequest struct {
	// example:
	//
	// 2023-01-11 10:31:00
	GmtCreate    *string                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MeteringData []*PushMeteringDataRequestMeteringData `json:"MeteringData,omitempty" xml:"MeteringData,omitempty" type:"Repeated"`
}

func (s PushMeteringDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) SetGmtCreate(v string) *PushMeteringDataRequest {
	s.GmtCreate = &v
	return s
}

func (s *PushMeteringDataRequest) SetMeteringData(v []*PushMeteringDataRequestMeteringData) *PushMeteringDataRequest {
	s.MeteringData = v
	return s
}

type PushMeteringDataRequestMeteringData struct {
	// example:
	//
	// 1666854480406
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// gtm-cn-20p314k5h05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test001
	MeteringAssist *string `json:"MeteringAssist,omitempty" xml:"MeteringAssist,omitempty"`
	// example:
	//
	// {"VirtualCpu":10}
	MeteringEntity *string `json:"MeteringEntity,omitempty" xml:"MeteringEntity,omitempty"`
	// example:
	//
	// 1662284820000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PushMeteringDataRequestMeteringData) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequestMeteringData) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequestMeteringData) SetEndTime(v int64) *PushMeteringDataRequestMeteringData {
	s.EndTime = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetInstanceId(v string) *PushMeteringDataRequestMeteringData {
	s.InstanceId = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringAssist(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringAssist = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringEntity(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringEntity = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetStartTime(v int64) *PushMeteringDataRequestMeteringData {
	s.StartTime = &v
	return s
}

type PushMeteringDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// parameter \\"service\\" can not be blank.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// false
	ForceFatal *bool `json:"ForceFatal,omitempty" xml:"ForceFatal,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) SetCode(v string) *PushMeteringDataResponseBody {
	s.Code = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetDynamicMessage(v string) *PushMeteringDataResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetForceFatal(v bool) *PushMeteringDataResponseBody {
	s.ForceFatal = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetMessage(v string) *PushMeteringDataResponseBody {
	s.Message = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetResult(v bool) *PushMeteringDataResponseBody {
	s.Result = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetSuccess(v bool) *PushMeteringDataResponseBody {
	s.Success = &v
	return s
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMeteringDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("marketplaceintl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 卖家查询实例列表
//
// @param request - DescribeSellerInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSellerInstancesResponse
func (client *Client) DescribeSellerInstancesWithOptions(request *DescribeSellerInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeSellerInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSellerInstances"),
		Version:     tea.String("2022-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSellerInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSellerInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 卖家查询实例列表
//
// @param request - DescribeSellerInstancesRequest
//
// @return DescribeSellerInstancesResponse
func (client *Client) DescribeSellerInstances(request *DescribeSellerInstancesRequest) (_result *DescribeSellerInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSellerInstancesResponse{}
	_body, _err := client.DescribeSellerInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// isv推送实例消息给用户
//
// @param request - NoticeInstanceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NoticeInstanceUserResponse
func (client *Client) NoticeInstanceUserWithOptions(request *NoticeInstanceUserRequest, runtime *util.RuntimeOptions) (_result *NoticeInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeParam)) {
		body["NoticeParam"] = request.NoticeParam
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("NoticeInstanceUser"),
		Version:     tea.String("2022-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &NoticeInstanceUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &NoticeInstanceUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// isv推送实例消息给用户
//
// @param request - NoticeInstanceUserRequest
//
// @return NoticeInstanceUserResponse
func (client *Client) NoticeInstanceUser(request *NoticeInstanceUserRequest) (_result *NoticeInstanceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NoticeInstanceUserResponse{}
	_body, _err := client.NoticeInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际云市场推送计量数据
//
// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *util.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtCreate)) {
		body["GmtCreate"] = request.GmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.MeteringData)) {
		body["MeteringData"] = request.MeteringData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMeteringData"),
		Version:     tea.String("2022-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PushMeteringDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PushMeteringDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 国际云市场推送计量数据
//
// @param request - PushMeteringDataRequest
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
