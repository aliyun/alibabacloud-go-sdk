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

type GetOwnerAgreementInstanceRequest struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
}

func (s GetOwnerAgreementInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerAgreementInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetOwnerAgreementInstanceRequest) SetAliyunUid(v string) *GetOwnerAgreementInstanceRequest {
	s.AliyunUid = &v
	return s
}

type GetOwnerAgreementInstanceResponseBody struct {
	Data      *GetOwnerAgreementInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                                    `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMsg    *string                                    `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOwnerAgreementInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerAgreementInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnerAgreementInstanceResponseBody) SetData(v *GetOwnerAgreementInstanceResponseBodyData) *GetOwnerAgreementInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBody) SetErrCode(v string) *GetOwnerAgreementInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBody) SetErrMsg(v string) *GetOwnerAgreementInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBody) SetRequestId(v string) *GetOwnerAgreementInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBody) SetSuccess(v bool) *GetOwnerAgreementInstanceResponseBody {
	s.Success = &v
	return s
}

type GetOwnerAgreementInstanceResponseBodyData struct {
	AgreementPropertyRoleDTOList []*GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList `json:"AgreementPropertyRoleDTOList,omitempty" xml:"AgreementPropertyRoleDTOList,omitempty" type:"Repeated"`
	Name                         *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid                          *string                                                                  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Uid                          *string                                                                  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetOwnerAgreementInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerAgreementInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOwnerAgreementInstanceResponseBodyData) SetAgreementPropertyRoleDTOList(v []*GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList) *GetOwnerAgreementInstanceResponseBodyData {
	s.AgreementPropertyRoleDTOList = v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBodyData) SetName(v string) *GetOwnerAgreementInstanceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBodyData) SetPid(v string) *GetOwnerAgreementInstanceResponseBodyData {
	s.Pid = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponseBodyData) SetUid(v string) *GetOwnerAgreementInstanceResponseBodyData {
	s.Uid = &v
	return s
}

type GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList struct {
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList) GoString() string {
	return s.String()
}

func (s *GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList) SetAgreementCode(v string) *GetOwnerAgreementInstanceResponseBodyDataAgreementPropertyRoleDTOList {
	s.AgreementCode = &v
	return s
}

type GetOwnerAgreementInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOwnerAgreementInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOwnerAgreementInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerAgreementInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetOwnerAgreementInstanceResponse) SetHeaders(v map[string]*string) *GetOwnerAgreementInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetOwnerAgreementInstanceResponse) SetStatusCode(v int32) *GetOwnerAgreementInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOwnerAgreementInstanceResponse) SetBody(v *GetOwnerAgreementInstanceResponseBody) *GetOwnerAgreementInstanceResponse {
	s.Body = v
	return s
}

type GetPartnerByUidRequest struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
}

func (s GetPartnerByUidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerByUidRequest) GoString() string {
	return s.String()
}

func (s *GetPartnerByUidRequest) SetAliyunUid(v string) *GetPartnerByUidRequest {
	s.AliyunUid = &v
	return s
}

type GetPartnerByUidResponseBody struct {
	Data      *GetPartnerByUidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                          `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMsg    *string                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPartnerByUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerByUidResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartnerByUidResponseBody) SetData(v *GetPartnerByUidResponseBodyData) *GetPartnerByUidResponseBody {
	s.Data = v
	return s
}

func (s *GetPartnerByUidResponseBody) SetErrCode(v string) *GetPartnerByUidResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetPartnerByUidResponseBody) SetErrMsg(v string) *GetPartnerByUidResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetPartnerByUidResponseBody) SetRequestId(v string) *GetPartnerByUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPartnerByUidResponseBody) SetSuccess(v bool) *GetPartnerByUidResponseBody {
	s.Success = &v
	return s
}

type GetPartnerByUidResponseBodyData struct {
	AgreementPropertyRoleDTOList []*GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList `json:"AgreementPropertyRoleDTOList,omitempty" xml:"AgreementPropertyRoleDTOList,omitempty" type:"Repeated"`
	Name                         *string                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid                          *string                                                        `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Uid                          *string                                                        `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetPartnerByUidResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerByUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPartnerByUidResponseBodyData) SetAgreementPropertyRoleDTOList(v []*GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList) *GetPartnerByUidResponseBodyData {
	s.AgreementPropertyRoleDTOList = v
	return s
}

func (s *GetPartnerByUidResponseBodyData) SetName(v string) *GetPartnerByUidResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPartnerByUidResponseBodyData) SetPid(v string) *GetPartnerByUidResponseBodyData {
	s.Pid = &v
	return s
}

func (s *GetPartnerByUidResponseBodyData) SetUid(v string) *GetPartnerByUidResponseBodyData {
	s.Uid = &v
	return s
}

type GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList struct {
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList) GoString() string {
	return s.String()
}

func (s *GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList) SetAgreementCode(v string) *GetPartnerByUidResponseBodyDataAgreementPropertyRoleDTOList {
	s.AgreementCode = &v
	return s
}

type GetPartnerByUidResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPartnerByUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPartnerByUidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerByUidResponse) GoString() string {
	return s.String()
}

func (s *GetPartnerByUidResponse) SetHeaders(v map[string]*string) *GetPartnerByUidResponse {
	s.Headers = v
	return s
}

func (s *GetPartnerByUidResponse) SetStatusCode(v int32) *GetPartnerByUidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartnerByUidResponse) SetBody(v *GetPartnerByUidResponseBody) *GetPartnerByUidResponse {
	s.Body = v
	return s
}

type QueryFusionOrderListRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndUserId *int64  `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryFusionOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFusionOrderListRequest) GoString() string {
	return s.String()
}

func (s *QueryFusionOrderListRequest) SetEndTime(v string) *QueryFusionOrderListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryFusionOrderListRequest) SetEndUserId(v int64) *QueryFusionOrderListRequest {
	s.EndUserId = &v
	return s
}

func (s *QueryFusionOrderListRequest) SetOrderId(v int64) *QueryFusionOrderListRequest {
	s.OrderId = &v
	return s
}

func (s *QueryFusionOrderListRequest) SetPageNo(v int32) *QueryFusionOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFusionOrderListRequest) SetPageSize(v int32) *QueryFusionOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFusionOrderListRequest) SetStartTime(v string) *QueryFusionOrderListRequest {
	s.StartTime = &v
	return s
}

type QueryFusionOrderListResponseBody struct {
	Code     *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data     []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg      *string   `json:"Msg,omitempty" xml:"Msg,omitempty"`
	PageNo   *int32    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFusionOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFusionOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFusionOrderListResponseBody) SetCode(v string) *QueryFusionOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetData(v []*string) *QueryFusionOrderListResponseBody {
	s.Data = v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetMsg(v string) *QueryFusionOrderListResponseBody {
	s.Msg = &v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetPageNo(v int32) *QueryFusionOrderListResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetPageSize(v int32) *QueryFusionOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetRequestId(v string) *QueryFusionOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFusionOrderListResponseBody) SetTotal(v int32) *QueryFusionOrderListResponseBody {
	s.Total = &v
	return s
}

type QueryFusionOrderListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFusionOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFusionOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFusionOrderListResponse) GoString() string {
	return s.String()
}

func (s *QueryFusionOrderListResponse) SetHeaders(v map[string]*string) *QueryFusionOrderListResponse {
	s.Headers = v
	return s
}

func (s *QueryFusionOrderListResponse) SetStatusCode(v int32) *QueryFusionOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFusionOrderListResponse) SetBody(v *QueryFusionOrderListResponseBody) *QueryFusionOrderListResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("agency.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("agency.aliyuncs.com"),
		"ap-south-1":                  tea.String("agency.aliyuncs.com"),
		"ap-southeast-2":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-3":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-5":              tea.String("agency.aliyuncs.com"),
		"cn-beijing":                  tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("agency.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("agency.aliyuncs.com"),
		"cn-chengdu":                  tea.String("agency.aliyuncs.com"),
		"cn-edge-1":                   tea.String("agency.aliyuncs.com"),
		"cn-fujian":                   tea.String("agency.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("agency.aliyuncs.com"),
		"cn-hongkong":                 tea.String("agency.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("agency.aliyuncs.com"),
		"cn-huhehaote":                tea.String("agency.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("agency.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-qingdao":                  tea.String("agency.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai":                 tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("agency.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("agency.aliyuncs.com"),
		"cn-wuhan":                    tea.String("agency.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("agency.aliyuncs.com"),
		"cn-yushanfang":               tea.String("agency.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("agency.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("agency.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("agency.aliyuncs.com"),
		"eu-central-1":                tea.String("agency.aliyuncs.com"),
		"eu-west-1":                   tea.String("agency.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("agency.aliyuncs.com"),
		"me-east-1":                   tea.String("agency.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("agency.aliyuncs.com"),
		"us-east-1":                   tea.String("agency.aliyuncs.com"),
		"us-west-1":                   tea.String("agency.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("agency"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetOwnerAgreementInstanceWithOptions(request *GetOwnerAgreementInstanceRequest, runtime *util.RuntimeOptions) (_result *GetOwnerAgreementInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		query["AliyunUid"] = request.AliyunUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOwnerAgreementInstance"),
		Version:     tea.String("2021-06-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOwnerAgreementInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOwnerAgreementInstance(request *GetOwnerAgreementInstanceRequest) (_result *GetOwnerAgreementInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOwnerAgreementInstanceResponse{}
	_body, _err := client.GetOwnerAgreementInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPartnerByUidWithOptions(request *GetPartnerByUidRequest, runtime *util.RuntimeOptions) (_result *GetPartnerByUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		query["AliyunUid"] = request.AliyunUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPartnerByUid"),
		Version:     tea.String("2021-06-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPartnerByUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPartnerByUid(request *GetPartnerByUidRequest) (_result *GetPartnerByUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPartnerByUidResponse{}
	_body, _err := client.GetPartnerByUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFusionOrderListWithOptions(request *QueryFusionOrderListRequest, runtime *util.RuntimeOptions) (_result *QueryFusionOrderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFusionOrderList"),
		Version:     tea.String("2021-06-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFusionOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFusionOrderList(request *QueryFusionOrderListRequest) (_result *QueryFusionOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFusionOrderListResponse{}
	_body, _err := client.QueryFusionOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
