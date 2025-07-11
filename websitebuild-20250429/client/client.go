// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateLogoTaskRequest struct {
	LogoVersion    *string `json:"LogoVersion,omitempty" xml:"LogoVersion,omitempty"`
	NegativePrompt *string `json:"NegativePrompt,omitempty" xml:"NegativePrompt,omitempty"`
	// example:
	//
	// {\\"ehpcutilParam\\":\\"sched job_submit --commandline \\\\\\"/ehpcdata/data/usersTest/huangqiaoyi-1725933699384/huangqiaoyi-1725933699384.slurm\\\\\\" --runasuser TestGfjnSimworks\\"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CreateLogoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskRequest) SetLogoVersion(v string) *CreateLogoTaskRequest {
	s.LogoVersion = &v
	return s
}

func (s *CreateLogoTaskRequest) SetNegativePrompt(v string) *CreateLogoTaskRequest {
	s.NegativePrompt = &v
	return s
}

func (s *CreateLogoTaskRequest) SetParameters(v string) *CreateLogoTaskRequest {
	s.Parameters = &v
	return s
}

func (s *CreateLogoTaskRequest) SetPrompt(v string) *CreateLogoTaskRequest {
	s.Prompt = &v
	return s
}

type CreateLogoTaskResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// c3r127e325at9yd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLogoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskResponseBody) SetErrorCode(v string) *CreateLogoTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetErrorMsg(v string) *CreateLogoTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetRequestId(v string) *CreateLogoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetSuccess(v bool) *CreateLogoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetTaskId(v string) *CreateLogoTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateLogoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskResponse) SetHeaders(v map[string]*string) *CreateLogoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLogoTaskResponse) SetStatusCode(v int32) *CreateLogoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogoTaskResponse) SetBody(v *CreateLogoTaskResponseBody) *CreateLogoTaskResponse {
	s.Body = v
	return s
}

type GetCreateLogoTaskRequest struct {
	// example:
	//
	// 20051349
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCreateLogoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreateLogoTaskRequest) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskRequest) SetTaskId(v string) *GetCreateLogoTaskRequest {
	s.TaskId = &v
	return s
}

type GetCreateLogoTaskResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Task    *GetCreateLogoTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetCreateLogoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreateLogoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponseBody) SetErrorCode(v string) *GetCreateLogoTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetErrorMsg(v string) *GetCreateLogoTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetRequestId(v string) *GetCreateLogoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetSuccess(v bool) *GetCreateLogoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetTask(v *GetCreateLogoTaskResponseBodyTask) *GetCreateLogoTaskResponseBody {
	s.Task = v
	return s
}

type GetCreateLogoTaskResponseBodyTask struct {
	// example:
	//
	// 604860995
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESS
	TaskStatus *string   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Urls       []*string `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s GetCreateLogoTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetCreateLogoTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponseBodyTask) SetTaskId(v string) *GetCreateLogoTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetCreateLogoTaskResponseBodyTask) SetTaskStatus(v string) *GetCreateLogoTaskResponseBodyTask {
	s.TaskStatus = &v
	return s
}

func (s *GetCreateLogoTaskResponseBodyTask) SetUrls(v []*string) *GetCreateLogoTaskResponseBodyTask {
	s.Urls = v
	return s
}

type GetCreateLogoTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateLogoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateLogoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreateLogoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponse) SetHeaders(v map[string]*string) *GetCreateLogoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetCreateLogoTaskResponse) SetStatusCode(v int32) *GetCreateLogoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateLogoTaskResponse) SetBody(v *GetCreateLogoTaskResponseBody) *GetCreateLogoTaskResponse {
	s.Body = v
	return s
}

type OperateAppInstanceForPartnerRequest struct {
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// SERVICE_DELIVERY_PROCESS
	OperateEvent *string `json:"OperateEvent,omitempty" xml:"OperateEvent,omitempty"`
}

func (s OperateAppInstanceForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAppInstanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerRequest) SetExtend(v string) *OperateAppInstanceForPartnerRequest {
	s.Extend = &v
	return s
}

func (s *OperateAppInstanceForPartnerRequest) SetOperateEvent(v string) *OperateAppInstanceForPartnerRequest {
	s.OperateEvent = &v
	return s
}

type OperateAppInstanceForPartnerResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateAppInstanceForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAppInstanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerResponseBody) SetErrorCode(v string) *OperateAppInstanceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetErrorMsg(v string) *OperateAppInstanceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetRequestId(v string) *OperateAppInstanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetSuccess(v string) *OperateAppInstanceForPartnerResponseBody {
	s.Success = &v
	return s
}

type OperateAppInstanceForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppInstanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppInstanceForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAppInstanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerResponse) SetHeaders(v map[string]*string) *OperateAppInstanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateAppInstanceForPartnerResponse) SetStatusCode(v int32) *OperateAppInstanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponse) SetBody(v *OperateAppInstanceForPartnerResponseBody) *OperateAppInstanceForPartnerResponse {
	s.Body = v
	return s
}

type OperateAppServiceForPartnerRequest struct {
	// example:
	//
	// WS00001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// {\\"user_device_id\\":\\"6bef45cb0c76de284d24de074c088b73\\"}\\n
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// SERVICE_FINISH
	OperateEvent *string `json:"OperateEvent,omitempty" xml:"OperateEvent,omitempty"`
	// example:
	//
	// WEBSITE_DESIGN
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s OperateAppServiceForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAppServiceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerRequest) SetBizId(v string) *OperateAppServiceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetExtend(v string) *OperateAppServiceForPartnerRequest {
	s.Extend = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetOperateEvent(v string) *OperateAppServiceForPartnerRequest {
	s.OperateEvent = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetServiceType(v string) *OperateAppServiceForPartnerRequest {
	s.ServiceType = &v
	return s
}

type OperateAppServiceForPartnerResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateAppServiceForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAppServiceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerResponseBody) SetErrorCode(v string) *OperateAppServiceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetErrorMsg(v string) *OperateAppServiceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetRequestId(v string) *OperateAppServiceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetSuccess(v bool) *OperateAppServiceForPartnerResponseBody {
	s.Success = &v
	return s
}

type OperateAppServiceForPartnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppServiceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppServiceForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAppServiceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerResponse) SetHeaders(v map[string]*string) *OperateAppServiceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateAppServiceForPartnerResponse) SetStatusCode(v int32) *OperateAppServiceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppServiceForPartnerResponse) SetBody(v *OperateAppServiceForPartnerResponseBody) *OperateAppServiceForPartnerResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("websitebuild"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 提交创建Logo任务
//
// @param request - CreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTaskWithOptions(request *CreateLogoTaskRequest, runtime *util.RuntimeOptions) (_result *CreateLogoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogoVersion)) {
		query["LogoVersion"] = request.LogoVersion
	}

	if !tea.BoolValue(util.IsUnset(request.NegativePrompt)) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogoTask"),
		Version:     tea.String("2025-04-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交创建Logo任务
//
// @param request - CreateLogoTaskRequest
//
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTask(request *CreateLogoTaskRequest) (_result *CreateLogoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Logo创建任务
//
// @param request - GetCreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTaskWithOptions(request *GetCreateLogoTaskRequest, runtime *util.RuntimeOptions) (_result *GetCreateLogoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreateLogoTask"),
		Version:     tea.String("2025-04-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Logo创建任务
//
// @param request - GetCreateLogoTaskRequest
//
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTask(request *GetCreateLogoTaskRequest) (_result *GetCreateLogoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.GetCreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用
//
// @param request - OperateAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartnerWithOptions(request *OperateAppInstanceForPartnerRequest, runtime *util.RuntimeOptions) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.OperateEvent)) {
		query["OperateEvent"] = request.OperateEvent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateAppInstanceForPartner"),
		Version:     tea.String("2025-04-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用
//
// @param request - OperateAppInstanceForPartnerRequest
//
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartner(request *OperateAppInstanceForPartnerRequest) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.OperateAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用服务
//
// @param request - OperateAppServiceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartnerWithOptions(request *OperateAppServiceForPartnerRequest, runtime *util.RuntimeOptions) (_result *OperateAppServiceForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.OperateEvent)) {
		query["OperateEvent"] = request.OperateEvent
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateAppServiceForPartner"),
		Version:     tea.String("2025-04-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用服务
//
// @param request - OperateAppServiceForPartnerRequest
//
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartner(request *OperateAppServiceForPartnerRequest) (_result *OperateAppServiceForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.OperateAppServiceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
