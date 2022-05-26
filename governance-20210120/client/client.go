// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnrollAccountRequest struct {
	// 账号名称前缀
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// 注册账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// 基线项配置数组
	BaselineItems []*EnrollAccountRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// 账号显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 父资源夹ID
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// 结算账号ID
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// RegionId
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnrollAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountRequest) GoString() string {
	return s.String()
}

func (s *EnrollAccountRequest) SetAccountNamePrefix(v string) *EnrollAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *EnrollAccountRequest) SetAccountUid(v int64) *EnrollAccountRequest {
	s.AccountUid = &v
	return s
}

func (s *EnrollAccountRequest) SetBaselineItems(v []*EnrollAccountRequestBaselineItems) *EnrollAccountRequest {
	s.BaselineItems = v
	return s
}

func (s *EnrollAccountRequest) SetDisplayName(v string) *EnrollAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *EnrollAccountRequest) SetFolderId(v string) *EnrollAccountRequest {
	s.FolderId = &v
	return s
}

func (s *EnrollAccountRequest) SetPayerAccountUid(v int64) *EnrollAccountRequest {
	s.PayerAccountUid = &v
	return s
}

func (s *EnrollAccountRequest) SetRegionId(v string) *EnrollAccountRequest {
	s.RegionId = &v
	return s
}

type EnrollAccountRequestBaselineItems struct {
	// 基线项配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 基线项名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否跳过基线项
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// 基线项版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s EnrollAccountRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *EnrollAccountRequestBaselineItems) SetConfig(v string) *EnrollAccountRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetName(v string) *EnrollAccountRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetSkip(v bool) *EnrollAccountRequestBaselineItems {
	s.Skip = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetVersion(v string) *EnrollAccountRequestBaselineItems {
	s.Version = &v
	return s
}

type EnrollAccountResponseBody struct {
	// 注册账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnrollAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountResponseBody) GoString() string {
	return s.String()
}

func (s *EnrollAccountResponseBody) SetAccountUid(v int64) *EnrollAccountResponseBody {
	s.AccountUid = &v
	return s
}

func (s *EnrollAccountResponseBody) SetRequestId(v string) *EnrollAccountResponseBody {
	s.RequestId = &v
	return s
}

type EnrollAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnrollAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnrollAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountResponse) GoString() string {
	return s.String()
}

func (s *EnrollAccountResponse) SetHeaders(v map[string]*string) *EnrollAccountResponse {
	s.Headers = v
	return s
}

func (s *EnrollAccountResponse) SetStatusCode(v int32) *EnrollAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *EnrollAccountResponse) SetBody(v *EnrollAccountResponseBody) *EnrollAccountResponse {
	s.Body = v
	return s
}

type GetEnrolledAccountRequest struct {
	// 账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// RegionId
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEnrolledAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountRequest) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountRequest) SetAccountUid(v int64) *GetEnrolledAccountRequest {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountRequest) SetRegionId(v string) *GetEnrolledAccountRequest {
	s.RegionId = &v
	return s
}

type GetEnrolledAccountResponseBody struct {
	// 账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 账号显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 错误信息
	ErrorInfo *GetEnrolledAccountResponseBodyErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// 父资源夹ID
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// 是否初始化完成
	Initialized *bool `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// 注册账号时的输入参数
	Inputs *GetEnrolledAccountResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// 所属的Master账号ID
	MasterAccountUid *int64 `json:"MasterAccountUid,omitempty" xml:"MasterAccountUid,omitempty"`
	// 结算账号ID
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// 基线实施进度
	Progress []*GetEnrolledAccountResponseBodyProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 账号注册状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetEnrolledAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBody) SetAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetCreateTime(v string) *GetEnrolledAccountResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetDisplayName(v string) *GetEnrolledAccountResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetErrorInfo(v *GetEnrolledAccountResponseBodyErrorInfo) *GetEnrolledAccountResponseBody {
	s.ErrorInfo = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetFolderId(v string) *GetEnrolledAccountResponseBody {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInitialized(v bool) *GetEnrolledAccountResponseBody {
	s.Initialized = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInputs(v *GetEnrolledAccountResponseBodyInputs) *GetEnrolledAccountResponseBody {
	s.Inputs = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetMasterAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.MasterAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.PayerAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetProgress(v []*GetEnrolledAccountResponseBodyProgress) *GetEnrolledAccountResponseBody {
	s.Progress = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetRequestId(v string) *GetEnrolledAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetStatus(v string) *GetEnrolledAccountResponseBody {
	s.Status = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetUpdateTime(v string) *GetEnrolledAccountResponseBody {
	s.UpdateTime = &v
	return s
}

type GetEnrolledAccountResponseBodyErrorInfo struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误处理建议
	Recommend *string `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEnrolledAccountResponseBodyErrorInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyErrorInfo) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetCode(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Code = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetMessage(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Message = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRecommend(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Recommend = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRequestId(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.RequestId = &v
	return s
}

type GetEnrolledAccountResponseBodyInputs struct {
	// 账号名称前缀
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// 账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// 基线项配置数组
	BaselineItems []*GetEnrolledAccountResponseBodyInputsBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// 账号展示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 父资源夹ID
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// 结算账号ID
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountNamePrefix(v string) *GetEnrolledAccountResponseBodyInputs {
	s.AccountNamePrefix = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetBaselineItems(v []*GetEnrolledAccountResponseBodyInputsBaselineItems) *GetEnrolledAccountResponseBodyInputs {
	s.BaselineItems = v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetDisplayName(v string) *GetEnrolledAccountResponseBodyInputs {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetFolderId(v string) *GetEnrolledAccountResponseBodyInputs {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.PayerAccountUid = &v
	return s
}

type GetEnrolledAccountResponseBodyInputsBaselineItems struct {
	// 基线项配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 基线项名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否跳过基线项
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// 基线项版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetConfig(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Config = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetName(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetSkip(v bool) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Skip = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetVersion(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Version = &v
	return s
}

type GetEnrolledAccountResponseBodyProgress struct {
	// 基线项名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 基线项实施状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEnrolledAccountResponseBodyProgress) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyProgress) SetName(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyProgress) SetStatus(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Status = &v
	return s
}

type GetEnrolledAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnrolledAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnrolledAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponse) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponse) SetHeaders(v map[string]*string) *GetEnrolledAccountResponse {
	s.Headers = v
	return s
}

func (s *GetEnrolledAccountResponse) SetStatusCode(v int32) *GetEnrolledAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnrolledAccountResponse) SetBody(v *GetEnrolledAccountResponseBody) *GetEnrolledAccountResponse {
	s.Body = v
	return s
}

type ListEnrolledAccountsRequest struct {
	// 每页的最大数据条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询返回结果下一页的令牌。首次调用API不需要NextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// RegionId
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnrolledAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsRequest) SetMaxResults(v int32) *ListEnrolledAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetNextToken(v string) *ListEnrolledAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetRegionId(v string) *ListEnrolledAccountsRequest {
	s.RegionId = &v
	return s
}

type ListEnrolledAccountsResponseBody struct {
	// 账号列表
	EnrolledAccounts []*ListEnrolledAccountsResponseBodyEnrolledAccounts `json:"EnrolledAccounts,omitempty" xml:"EnrolledAccounts,omitempty" type:"Repeated"`
	// 查询返回结果下一页的令牌
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnrolledAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBody) SetEnrolledAccounts(v []*ListEnrolledAccountsResponseBodyEnrolledAccounts) *ListEnrolledAccountsResponseBody {
	s.EnrolledAccounts = v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetNextToken(v string) *ListEnrolledAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetRequestId(v string) *ListEnrolledAccountsResponseBody {
	s.RequestId = &v
	return s
}

type ListEnrolledAccountsResponseBodyEnrolledAccounts struct {
	// 账号ID
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 账号显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 父资源夹ID
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// 结算账号ID
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// 创建状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.AccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetCreateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetDisplayName(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetFolderId(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.FolderId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetPayerAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.PayerAccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetStatus(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.Status = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetUpdateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.UpdateTime = &v
	return s
}

type ListEnrolledAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEnrolledAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnrolledAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponse) SetHeaders(v map[string]*string) *ListEnrolledAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListEnrolledAccountsResponse) SetStatusCode(v int32) *ListEnrolledAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnrolledAccountsResponse) SetBody(v *ListEnrolledAccountsResponseBody) *ListEnrolledAccountsResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("governance"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) EnrollAccountWithOptions(request *EnrollAccountRequest, runtime *util.RuntimeOptions) (_result *EnrollAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNamePrefix)) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.AccountUid)) {
		query["AccountUid"] = request.AccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineItems)) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountUid)) {
		query["PayerAccountUid"] = request.PayerAccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnrollAccount"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnrollAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnrollAccount(request *EnrollAccountRequest) (_result *EnrollAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnrollAccountResponse{}
	_body, _err := client.EnrollAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnrolledAccountWithOptions(request *GetEnrolledAccountRequest, runtime *util.RuntimeOptions) (_result *GetEnrolledAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountUid)) {
		query["AccountUid"] = request.AccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnrolledAccount"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnrolledAccount(request *GetEnrolledAccountRequest) (_result *GetEnrolledAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.GetEnrolledAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnrolledAccountsWithOptions(request *ListEnrolledAccountsRequest, runtime *util.RuntimeOptions) (_result *ListEnrolledAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnrolledAccounts"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnrolledAccounts(request *ListEnrolledAccountsRequest) (_result *ListEnrolledAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.ListEnrolledAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
