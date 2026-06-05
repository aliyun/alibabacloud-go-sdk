// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("websitebuild"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分配Supabase实例
//
// @param request - AllocateSupabaseForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateSupabaseForAdminResponse
func (client *Client) AllocateSupabaseForAdminWithOptions(request *AllocateSupabaseForAdminRequest, runtime *dara.RuntimeOptions) (_result *AllocateSupabaseForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateSupabaseForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateSupabaseForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分配Supabase实例
//
// @param request - AllocateSupabaseForAdminRequest
//
// @return AllocateSupabaseForAdminResponse
func (client *Client) AllocateSupabaseForAdmin(request *AllocateSupabaseForAdminRequest) (_result *AllocateSupabaseForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateSupabaseForAdminResponse{}
	_body, _err := client.AllocateSupabaseForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量校验资源计量
//
// @param request - BatchCheckResourceMeasureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCheckResourceMeasureResponse
func (client *Client) BatchCheckResourceMeasureWithOptions(request *BatchCheckResourceMeasureRequest, runtime *dara.RuntimeOptions) (_result *BatchCheckResourceMeasureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BelongId) {
		query["BelongId"] = request.BelongId
	}

	if !dara.IsNil(request.BelongIdType) {
		query["BelongIdType"] = request.BelongIdType
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EspBizId) {
		query["EspBizId"] = request.EspBizId
	}

	if !dara.IsNil(request.OrderComponentParams) {
		query["OrderComponentParams"] = request.OrderComponentParams
	}

	if !dara.IsNil(request.ResourceCheckItems) {
		query["ResourceCheckItems"] = request.ResourceCheckItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCheckResourceMeasure"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCheckResourceMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量校验资源计量
//
// @param request - BatchCheckResourceMeasureRequest
//
// @return BatchCheckResourceMeasureResponse
func (client *Client) BatchCheckResourceMeasure(request *BatchCheckResourceMeasureRequest) (_result *BatchCheckResourceMeasureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCheckResourceMeasureResponse{}
	_body, _err := client.BatchCheckResourceMeasureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Bind Application Domain
//
// @param request - BindAppDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAppDomainResponse
func (client *Client) BindAppDomainWithOptions(request *BindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *BindAppDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAppDomain"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAppDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Bind Application Domain
//
// @param request - BindAppDomainRequest
//
// @return BindAppDomainResponse
func (client *Client) BindAppDomain(request *BindAppDomainRequest) (_result *BindAppDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAppDomainResponse{}
	_body, _err := client.BindAppDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验资源计量
//
// @param request - CheckResourceMeasureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourceMeasureResponse
func (client *Client) CheckResourceMeasureWithOptions(request *CheckResourceMeasureRequest, runtime *dara.RuntimeOptions) (_result *CheckResourceMeasureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BelongId) {
		query["BelongId"] = request.BelongId
	}

	if !dara.IsNil(request.BelongIdType) {
		query["BelongIdType"] = request.BelongIdType
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EspBizId) {
		query["EspBizId"] = request.EspBizId
	}

	if !dara.IsNil(request.OrderComponentParams) {
		query["OrderComponentParams"] = request.OrderComponentParams
	}

	if !dara.IsNil(request.ResourceCode) {
		query["ResourceCode"] = request.ResourceCode
	}

	if !dara.IsNil(request.ResourceValue) {
		query["ResourceValue"] = request.ResourceValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckResourceMeasure"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckResourceMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验资源计量
//
// @param request - CheckResourceMeasureRequest
//
// @return CheckResourceMeasureResponse
func (client *Client) CheckResourceMeasure(request *CheckResourceMeasureRequest) (_result *CheckResourceMeasureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckResourceMeasureResponse{}
	_body, _err := client.CheckResourceMeasureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验用户资源计量
//
// @param request - CheckUserResourceMeasureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUserResourceMeasureResponse
func (client *Client) CheckUserResourceMeasureWithOptions(request *CheckUserResourceMeasureRequest, runtime *dara.RuntimeOptions) (_result *CheckUserResourceMeasureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BelongId) {
		query["BelongId"] = request.BelongId
	}

	if !dara.IsNil(request.BelongIdType) {
		query["BelongIdType"] = request.BelongIdType
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EspBizId) {
		query["EspBizId"] = request.EspBizId
	}

	if !dara.IsNil(request.OrderComponentParams) {
		query["OrderComponentParams"] = request.OrderComponentParams
	}

	if !dara.IsNil(request.ResourceCode) {
		query["ResourceCode"] = request.ResourceCode
	}

	if !dara.IsNil(request.ResourceValue) {
		query["ResourceValue"] = request.ResourceValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUserResourceMeasure"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUserResourceMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验用户资源计量
//
// @param request - CheckUserResourceMeasureRequest
//
// @return CheckUserResourceMeasureResponse
func (client *Client) CheckUserResourceMeasure(request *CheckUserResourceMeasureRequest) (_result *CheckUserResourceMeasureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUserResourceMeasureResponse{}
	_body, _err := client.CheckUserResourceMeasureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 复制插件配置
//
// @param request - CopyAppPluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyAppPluginConfigResponse
func (client *Client) CopyAppPluginConfigWithOptions(request *CopyAppPluginConfigRequest, runtime *dara.RuntimeOptions) (_result *CopyAppPluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceBizId) {
		query["SourceBizId"] = request.SourceBizId
	}

	if !dara.IsNil(request.TargetBizId) {
		query["TargetBizId"] = request.TargetBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyAppPluginConfig"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyAppPluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制插件配置
//
// @param request - CopyAppPluginConfigRequest
//
// @return CopyAppPluginConfigResponse
func (client *Client) CopyAppPluginConfig(request *CopyAppPluginConfigRequest) (_result *CopyAppPluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyAppPluginConfigResponse{}
	_body, _err := client.CopyAppPluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异步发起AI员工对话
//
// @param request - CreateAIStaffChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAIStaffChatResponse
func (client *Client) CreateAIStaffChatWithOptions(request *CreateAIStaffChatRequest, runtime *dara.RuntimeOptions) (_result *CreateAIStaffChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ChatId) {
		body["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.MetaData) {
		bodyFlat["MetaData"] = request.MetaData
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAIStaffChat"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAIStaffChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异步发起AI员工对话
//
// @param request - CreateAIStaffChatRequest
//
// @return CreateAIStaffChatResponse
func (client *Client) CreateAIStaffChat(request *CreateAIStaffChatRequest) (_result *CreateAIStaffChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAIStaffChatResponse{}
	_body, _err := client.CreateAIStaffChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI员工会话
//
// @param request - CreateAIStaffConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAIStaffConversationResponse
func (client *Client) CreateAIStaffConversationWithOptions(request *CreateAIStaffConversationRequest, runtime *dara.RuntimeOptions) (_result *CreateAIStaffConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAIStaffConversation"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAIStaffConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI员工会话
//
// @param request - CreateAIStaffConversationRequest
//
// @return CreateAIStaffConversationResponse
func (client *Client) CreateAIStaffConversation(request *CreateAIStaffConversationRequest) (_result *CreateAIStaffConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAIStaffConversationResponse{}
	_body, _err := client.CreateAIStaffConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用助手智能体
//
// @param request - CreateAppAssistantAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppAssistantAgentResponse
func (client *Client) CreateAppAssistantAgentWithOptions(request *CreateAppAssistantAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateAppAssistantAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.PlatformType) {
		query["PlatformType"] = request.PlatformType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppAssistantAgent"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppAssistantAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用助手智能体
//
// @param request - CreateAppAssistantAgentRequest
//
// @return CreateAppAssistantAgentResponse
func (client *Client) CreateAppAssistantAgent(request *CreateAppAssistantAgentRequest) (_result *CreateAppAssistantAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppAssistantAgentResponse{}
	_body, _err := client.CreateAppAssistantAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成应用助手SSO免登
//
// @param request - CreateAppAssistantAgentSsoLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppAssistantAgentSsoLoginResponse
func (client *Client) CreateAppAssistantAgentSsoLoginWithOptions(request *CreateAppAssistantAgentSsoLoginRequest, runtime *dara.RuntimeOptions) (_result *CreateAppAssistantAgentSsoLoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.PlatformType) {
		query["PlatformType"] = request.PlatformType
	}

	if !dara.IsNil(request.TargetUrl) {
		query["TargetUrl"] = request.TargetUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppAssistantAgentSsoLogin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppAssistantAgentSsoLoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成应用助手SSO免登
//
// @param request - CreateAppAssistantAgentSsoLoginRequest
//
// @return CreateAppAssistantAgentSsoLoginResponse
func (client *Client) CreateAppAssistantAgentSsoLogin(request *CreateAppAssistantAgentSsoLoginRequest) (_result *CreateAppAssistantAgentSsoLoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppAssistantAgentSsoLoginResponse{}
	_body, _err := client.CreateAppAssistantAgentSsoLoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 万小智发起AI对话
//
// @param request - CreateAppChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppChatResponse
func (client *Client) CreateAppChatWithSSE(request *CreateAppChatRequest, runtime *dara.RuntimeOptions, _yield chan *CreateAppChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createAppChatWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// 万小智发起AI对话
//
// @param request - CreateAppChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppChatResponse
func (client *Client) CreateAppChatWithOptions(request *CreateAppChatRequest, runtime *dara.RuntimeOptions) (_result *CreateAppChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["BotId"] = request.BotId
	}

	if !dara.IsNil(request.ChatId) {
		body["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppChat"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateAppChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 万小智发起AI对话
//
// @param request - CreateAppChatRequest
//
// @return CreateAppChatResponse
func (client *Client) CreateAppChat(request *CreateAppChatRequest) (_result *CreateAppChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppChatResponse{}
	_body, _err := client.CreateAppChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a website instance
//
// @param tmpReq - CreateAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInstanceResponse
func (client *Client) CreateAppInstanceWithOptions(tmpReq *CreateAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateAction) {
		query["CreateAction"] = request.CreateAction
	}

	if !dara.IsNil(request.DeployArea) {
		query["DeployArea"] = request.DeployArea
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a website instance
//
// @param request - CreateAppInstanceRequest
//
// @return CreateAppInstanceResponse
func (client *Client) CreateAppInstance(request *CreateAppInstanceRequest) (_result *CreateAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppInstanceResponse{}
	_body, _err := client.CreateAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SSO ticket
//
// @param request - CreateAppInstanceTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInstanceTicketResponse
func (client *Client) CreateAppInstanceTicketWithOptions(request *CreateAppInstanceTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInstanceTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppInstanceTicket"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppInstanceTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SSO ticket
//
// @param request - CreateAppInstanceTicketRequest
//
// @return CreateAppInstanceTicketResponse
func (client *Client) CreateAppInstanceTicket(request *CreateAppInstanceTicketRequest) (_result *CreateAppInstanceTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppInstanceTicketResponse{}
	_body, _err := client.CreateAppInstanceTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用实例创建 LLM 网关 API-KEY
//
// @param request - CreateAppLlmApiKeyForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppLlmApiKeyForPartnerResponse
func (client *Client) CreateAppLlmApiKeyForPartnerWithOptions(request *CreateAppLlmApiKeyForPartnerRequest, runtime *dara.RuntimeOptions) (_result *CreateAppLlmApiKeyForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IpWhiteList) {
		body["IpWhiteList"] = request.IpWhiteList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppLlmApiKeyForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppLlmApiKeyForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用实例创建 LLM 网关 API-KEY
//
// @param request - CreateAppLlmApiKeyForPartnerRequest
//
// @return CreateAppLlmApiKeyForPartnerResponse
func (client *Client) CreateAppLlmApiKeyForPartner(request *CreateAppLlmApiKeyForPartnerRequest) (_result *CreateAppLlmApiKeyForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppLlmApiKeyForPartnerResponse{}
	_body, _err := client.CreateAppLlmApiKeyForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Activate the Wanxiaozhi Inspiration Value service
//
// @param request - CreateAppTokenServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppTokenServiceResponse
func (client *Client) CreateAppTokenServiceWithOptions(request *CreateAppTokenServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateAppTokenServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateAction) {
		query["CreateAction"] = request.CreateAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppTokenService"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppTokenServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Activate the Wanxiaozhi Inspiration Value service
//
// @param request - CreateAppTokenServiceRequest
//
// @return CreateAppTokenServiceResponse
func (client *Client) CreateAppTokenService(request *CreateAppTokenServiceRequest) (_result *CreateAppTokenServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppTokenServiceResponse{}
	_body, _err := client.CreateAppTokenServiceWithOptions(request, runtime)
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
func (client *Client) CreateLogoTaskWithOptions(request *CreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLogoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogoVersion) {
		query["LogoVersion"] = request.LogoVersion
	}

	if !dara.IsNil(request.NegativePrompt) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
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
// 创建素材中心文件夹
//
// @param request - CreateMaterialDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMaterialDirectoryResponse
func (client *Client) CreateMaterialDirectoryWithOptions(request *CreateMaterialDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateMaterialDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		query["ParentDirectoryId"] = request.ParentDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMaterialDirectory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMaterialDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建素材中心文件夹
//
// @param request - CreateMaterialDirectoryRequest
//
// @return CreateMaterialDirectoryResponse
func (client *Client) CreateMaterialDirectory(request *CreateMaterialDirectoryRequest) (_result *CreateMaterialDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMaterialDirectoryResponse{}
	_body, _err := client.CreateMaterialDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete the SSL certificate of a domain
//
// @param request - DeleteAppDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppDomainCertificateResponse
func (client *Client) DeleteAppDomainCertificateWithOptions(request *DeleteAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppDomainCertificate"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete the SSL certificate of a domain
//
// @param request - DeleteAppDomainCertificateRequest
//
// @return DeleteAppDomainCertificateResponse
func (client *Client) DeleteAppDomainCertificate(request *DeleteAppDomainCertificateRequest) (_result *DeleteAppDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppDomainCertificateResponse{}
	_body, _err := client.DeleteAppDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete the domain redirection rules
//
// @param request - DeleteAppDomainRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppDomainRedirectResponse
func (client *Client) DeleteAppDomainRedirectWithOptions(request *DeleteAppDomainRedirectRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainRedirectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppDomainRedirect"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppDomainRedirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete the domain redirection rules
//
// @param request - DeleteAppDomainRedirectRequest
//
// @return DeleteAppDomainRedirectResponse
func (client *Client) DeleteAppDomainRedirect(request *DeleteAppDomainRedirectRequest) (_result *DeleteAppDomainRedirectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppDomainRedirectResponse{}
	_body, _err := client.DeleteAppDomainRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteAppInstanceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstanceFileResponse
func (client *Client) DeleteAppInstanceFileWithOptions(request *DeleteAppInstanceFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInstanceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppInstanceFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppInstanceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteAppInstanceFileRequest
//
// @return DeleteAppInstanceFileResponse
func (client *Client) DeleteAppInstanceFile(request *DeleteAppInstanceFileRequest) (_result *DeleteAppInstanceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppInstanceFileResponse{}
	_body, _err := client.DeleteAppInstanceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Supabase密钥
//
// @param request - DeleteAppSupabaseSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppSupabaseSecretsResponse
func (client *Client) DeleteAppSupabaseSecretsWithOptions(request *DeleteAppSupabaseSecretsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppSupabaseSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.KeysJson) {
		query["KeysJson"] = request.KeysJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppSupabaseSecrets"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppSupabaseSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Supabase密钥
//
// @param request - DeleteAppSupabaseSecretsRequest
//
// @return DeleteAppSupabaseSecretsResponse
func (client *Client) DeleteAppSupabaseSecrets(request *DeleteAppSupabaseSecretsRequest) (_result *DeleteAppSupabaseSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppSupabaseSecretsResponse{}
	_body, _err := client.DeleteAppSupabaseSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除素材中心文件夹
//
// @param request - DeleteMaterialDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaterialDirectoryResponse
func (client *Client) DeleteMaterialDirectoryWithOptions(request *DeleteMaterialDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaterialDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaterialDirectory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaterialDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除素材中心文件夹
//
// @param request - DeleteMaterialDirectoryRequest
//
// @return DeleteMaterialDirectoryResponse
func (client *Client) DeleteMaterialDirectory(request *DeleteMaterialDirectoryRequest) (_result *DeleteMaterialDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMaterialDirectoryResponse{}
	_body, _err := client.DeleteMaterialDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除素材生产任务
//
// @param tmpReq - DeleteMaterialTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaterialTaskResponse
func (client *Client) DeleteMaterialTaskWithOptions(tmpReq *DeleteMaterialTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaterialTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMaterialTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaterialTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaterialTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除素材生产任务
//
// @param request - DeleteMaterialTaskRequest
//
// @return DeleteMaterialTaskResponse
func (client *Client) DeleteMaterialTask(request *DeleteMaterialTaskRequest) (_result *DeleteMaterialTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMaterialTaskResponse{}
	_body, _err := client.DeleteMaterialTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the DNS resolution records of a domain
//
// @param request - DescribeAppDomainDnsRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppDomainDnsRecordResponse
func (client *Client) DescribeAppDomainDnsRecordWithOptions(request *DescribeAppDomainDnsRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppDomainDnsRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Purpose) {
		query["Purpose"] = request.Purpose
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppDomainDnsRecord"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppDomainDnsRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the DNS resolution records of a domain
//
// @param request - DescribeAppDomainDnsRecordRequest
//
// @return DescribeAppDomainDnsRecordResponse
func (client *Client) DescribeAppDomainDnsRecord(request *DescribeAppDomainDnsRecordRequest) (_result *DescribeAppDomainDnsRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppDomainDnsRecordResponse{}
	_body, _err := client.DescribeAppDomainDnsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DispatchConsoleAPIForPartner
//
// @param request - DispatchConsoleAPIForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DispatchConsoleAPIForPartnerResponse
func (client *Client) DispatchConsoleAPIForPartnerWithOptions(request *DispatchConsoleAPIForPartnerRequest, runtime *dara.RuntimeOptions) (_result *DispatchConsoleAPIForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LiveToken) {
		query["LiveToken"] = request.LiveToken
	}

	if !dara.IsNil(request.Operation) {
		query["Operation"] = request.Operation
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SiteHost) {
		query["SiteHost"] = request.SiteHost
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DispatchConsoleAPIForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DispatchConsoleAPIForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DispatchConsoleAPIForPartner
//
// @param request - DispatchConsoleAPIForPartnerRequest
//
// @return DispatchConsoleAPIForPartnerResponse
func (client *Client) DispatchConsoleAPIForPartner(request *DispatchConsoleAPIForPartnerRequest) (_result *DispatchConsoleAPIForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DispatchConsoleAPIForPartnerResponse{}
	_body, _err := client.DispatchConsoleAPIForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑插件配置
//
// @param request - EditPluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditPluginConfigResponse
func (client *Client) EditPluginConfigWithOptions(request *EditPluginConfigRequest, runtime *dara.RuntimeOptions) (_result *EditPluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.PluginConfig) {
		query["PluginConfig"] = request.PluginConfig
	}

	if !dara.IsNil(request.PluginDesc) {
		query["PluginDesc"] = request.PluginDesc
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditPluginConfig"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditPluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑插件配置
//
// @param request - EditPluginConfigRequest
//
// @return EditPluginConfigResponse
func (client *Client) EditPluginConfig(request *EditPluginConfigRequest) (_result *EditPluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditPluginConfigResponse{}
	_body, _err := client.EditPluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出素材文件
//
// @param tmpReq - ExportMaterialFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportMaterialFileResponse
func (client *Client) ExportMaterialFileWithOptions(tmpReq *ExportMaterialFileRequest, runtime *dara.RuntimeOptions) (_result *ExportMaterialFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportMaterialFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportMaterialFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportMaterialFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出素材文件
//
// @param request - ExportMaterialFileRequest
//
// @return ExportMaterialFileResponse
func (client *Client) ExportMaterialFile(request *ExportMaterialFileRequest) (_result *ExportMaterialFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportMaterialFileResponse{}
	_body, _err := client.ExportMaterialFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI员工站点预览地址
//
// @param request - GetAIStaffPreviewUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIStaffPreviewUrlResponse
func (client *Client) GetAIStaffPreviewUrlWithOptions(request *GetAIStaffPreviewUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAIStaffPreviewUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Restart) {
		body["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAIStaffPreviewUrl"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAIStaffPreviewUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI员工站点预览地址
//
// @param request - GetAIStaffPreviewUrlRequest
//
// @return GetAIStaffPreviewUrlResponse
func (client *Client) GetAIStaffPreviewUrl(request *GetAIStaffPreviewUrlRequest) (_result *GetAIStaffPreviewUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAIStaffPreviewUrlResponse{}
	_body, _err := client.GetAIStaffPreviewUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代码工作区详情
//
// @param request - GetAppCodeWorkspaceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppCodeWorkspaceDetailResponse
func (client *Client) GetAppCodeWorkspaceDetailWithOptions(request *GetAppCodeWorkspaceDetailRequest, runtime *dara.RuntimeOptions) (_result *GetAppCodeWorkspaceDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppCodeWorkspaceDetail"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppCodeWorkspaceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代码工作区详情
//
// @param request - GetAppCodeWorkspaceDetailRequest
//
// @return GetAppCodeWorkspaceDetailResponse
func (client *Client) GetAppCodeWorkspaceDetail(request *GetAppCodeWorkspaceDetailRequest) (_result *GetAppCodeWorkspaceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppCodeWorkspaceDetailResponse{}
	_body, _err := client.GetAppCodeWorkspaceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取对话详情
//
// @param request - GetAppConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConversationResponse
func (client *Client) GetAppConversationWithOptions(request *GetAppConversationRequest, runtime *dara.RuntimeOptions) (_result *GetAppConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		query["BotId"] = request.BotId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppConversation"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取对话详情
//
// @param request - GetAppConversationRequest
//
// @return GetAppConversationResponse
func (client *Client) GetAppConversation(request *GetAppConversationRequest) (_result *GetAppConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppConversationResponse{}
	_body, _err := client.GetAppConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取会话锁定状态
//
// @param request - GetAppConversationLockStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConversationLockStatusResponse
func (client *Client) GetAppConversationLockStatusWithOptions(request *GetAppConversationLockStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAppConversationLockStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppConversationLockStatus"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppConversationLockStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话锁定状态
//
// @param request - GetAppConversationLockStatusRequest
//
// @return GetAppConversationLockStatusResponse
func (client *Client) GetAppConversationLockStatus(request *GetAppConversationLockStatusRequest) (_result *GetAppConversationLockStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppConversationLockStatusResponse{}
	_body, _err := client.GetAppConversationLockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据库表结构
//
// @param request - GetAppDatabaseTableSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppDatabaseTableSchemasResponse
func (client *Client) GetAppDatabaseTableSchemasWithOptions(request *GetAppDatabaseTableSchemasRequest, runtime *dara.RuntimeOptions) (_result *GetAppDatabaseTableSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppDatabaseTableSchemas"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppDatabaseTableSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据库表结构
//
// @param request - GetAppDatabaseTableSchemasRequest
//
// @return GetAppDatabaseTableSchemasResponse
func (client *Client) GetAppDatabaseTableSchemas(request *GetAppDatabaseTableSchemasRequest) (_result *GetAppDatabaseTableSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppDatabaseTableSchemasResponse{}
	_body, _err := client.GetAppDatabaseTableSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取文件及修改时间
//
// @param request - GetAppFileContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppFileContentResponse
func (client *Client) GetAppFileContentWithOptions(request *GetAppFileContentRequest, runtime *dara.RuntimeOptions) (_result *GetAppFileContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppFileContent"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppFileContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取文件及修改时间
//
// @param request - GetAppFileContentRequest
//
// @return GetAppFileContentResponse
func (client *Client) GetAppFileContent(request *GetAppFileContentRequest) (_result *GetAppFileContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppFileContentResponse{}
	_body, _err := client.GetAppFileContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Application Instance Details
//
// @param request - GetAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceResponse
func (client *Client) GetAppInstanceWithOptions(request *GetAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Application Instance Details
//
// @param request - GetAppInstanceRequest
//
// @return GetAppInstanceResponse
func (client *Client) GetAppInstance(request *GetAppInstanceRequest) (_result *GetAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppInstanceResponse{}
	_body, _err := client.GetAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetAppInstanceEntitlement
//
// @param request - GetAppInstanceEntitlementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceEntitlementResponse
func (client *Client) GetAppInstanceEntitlementWithOptions(request *GetAppInstanceEntitlementRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceEntitlementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstanceEntitlement"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceEntitlementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetAppInstanceEntitlement
//
// @param request - GetAppInstanceEntitlementRequest
//
// @return GetAppInstanceEntitlementResponse
func (client *Client) GetAppInstanceEntitlement(request *GetAppInstanceEntitlementRequest) (_result *GetAppInstanceEntitlementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppInstanceEntitlementResponse{}
	_body, _err := client.GetAppInstanceEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query application instance information
//
// @param request - GetAppInstanceForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceForAdminResponse
func (client *Client) GetAppInstanceForAdminWithOptions(request *GetAppInstanceForAdminRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstanceForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query application instance information
//
// @param request - GetAppInstanceForAdminRequest
//
// @return GetAppInstanceForAdminResponse
func (client *Client) GetAppInstanceForAdmin(request *GetAppInstanceForAdminRequest) (_result *GetAppInstanceForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppInstanceForAdminResponse{}
	_body, _err := client.GetAppInstanceForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query instance details
//
// @param request - GetAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceForPartnerResponse
func (client *Client) GetAppInstanceForPartnerWithOptions(request *GetAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query instance details
//
// @param request - GetAppInstanceForPartnerRequest
//
// @return GetAppInstanceForPartnerResponse
func (client *Client) GetAppInstanceForPartner(request *GetAppInstanceForPartnerRequest) (_result *GetAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppInstanceForPartnerResponse{}
	_body, _err := client.GetAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用临时短链
//
// @param request - GetAppInstanceTempShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceTempShortUrlResponse
func (client *Client) GetAppInstanceTempShortUrlWithOptions(request *GetAppInstanceTempShortUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceTempShortUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstanceTempShortUrl"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceTempShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用临时短链
//
// @param request - GetAppInstanceTempShortUrlRequest
//
// @return GetAppInstanceTempShortUrlResponse
func (client *Client) GetAppInstanceTempShortUrl(request *GetAppInstanceTempShortUrlRequest) (_result *GetAppInstanceTempShortUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppInstanceTempShortUrlResponse{}
	_body, _err := client.GetAppInstanceTempShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生码-获取插件配置信息
//
// @param request - GetAppPluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppPluginConfigResponse
func (client *Client) GetAppPluginConfigWithOptions(request *GetAppPluginConfigRequest, runtime *dara.RuntimeOptions) (_result *GetAppPluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.PluginId) {
		body["PluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppPluginConfig"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppPluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生码-获取插件配置信息
//
// @param request - GetAppPluginConfigRequest
//
// @return GetAppPluginConfigResponse
func (client *Client) GetAppPluginConfig(request *GetAppPluginConfigRequest) (_result *GetAppPluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppPluginConfigResponse{}
	_body, _err := client.GetAppPluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布状态查询
//
// @param request - GetAppPublishStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppPublishStatusResponse
func (client *Client) GetAppPublishStatusWithOptions(request *GetAppPublishStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAppPublishStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeployOrderId) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	if !dara.IsNil(request.WebsiteDomain) {
		query["WebsiteDomain"] = request.WebsiteDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppPublishStatus"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppPublishStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布状态查询
//
// @param request - GetAppPublishStatusRequest
//
// @return GetAppPublishStatusResponse
func (client *Client) GetAppPublishStatus(request *GetAppPublishStatusRequest) (_result *GetAppPublishStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppPublishStatusResponse{}
	_body, _err := client.GetAppPublishStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询万小智推荐商品
//
// @param request - GetAppRecommendedCommoditiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppRecommendedCommoditiesResponse
func (client *Client) GetAppRecommendedCommoditiesWithOptions(request *GetAppRecommendedCommoditiesRequest, runtime *dara.RuntimeOptions) (_result *GetAppRecommendedCommoditiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.ResourceConditions) {
		query["ResourceConditions"] = request.ResourceConditions
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppRecommendedCommodities"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppRecommendedCommoditiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询万小智推荐商品
//
// @param request - GetAppRecommendedCommoditiesRequest
//
// @return GetAppRecommendedCommoditiesResponse
func (client *Client) GetAppRecommendedCommodities(request *GetAppRecommendedCommoditiesRequest) (_result *GetAppRecommendedCommoditiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppRecommendedCommoditiesResponse{}
	_body, _err := client.GetAppRecommendedCommoditiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 需求查询
//
// @param request - GetAppRequirementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppRequirementResponse
func (client *Client) GetAppRequirementWithOptions(request *GetAppRequirementRequest, runtime *dara.RuntimeOptions) (_result *GetAppRequirementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppRequirement"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppRequirementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 需求查询
//
// @param request - GetAppRequirementRequest
//
// @return GetAppRequirementResponse
func (client *Client) GetAppRequirement(request *GetAppRequirementRequest) (_result *GetAppRequirementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppRequirementResponse{}
	_body, _err := client.GetAppRequirementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取沙箱预览地址
//
// @param request - GetAppSandboxPreviewUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSandboxPreviewUrlResponse
func (client *Client) GetAppSandboxPreviewUrlWithOptions(request *GetAppSandboxPreviewUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAppSandboxPreviewUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Restart) {
		body["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSandboxPreviewUrl"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSandboxPreviewUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取沙箱预览地址
//
// @param request - GetAppSandboxPreviewUrlRequest
//
// @return GetAppSandboxPreviewUrlResponse
func (client *Client) GetAppSandboxPreviewUrl(request *GetAppSandboxPreviewUrlRequest) (_result *GetAppSandboxPreviewUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSandboxPreviewUrlResponse{}
	_body, _err := client.GetAppSandboxPreviewUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询SEO索引状态
//
// @param request - GetAppSeoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSeoStatusResponse
func (client *Client) GetAppSeoStatusWithOptions(request *GetAppSeoStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAppSeoStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.SeType) {
		query["SeType"] = request.SeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSeoStatus"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSeoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询SEO索引状态
//
// @param request - GetAppSeoStatusRequest
//
// @return GetAppSeoStatusResponse
func (client *Client) GetAppSeoStatus(request *GetAppSeoStatusRequest) (_result *GetAppSeoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSeoStatusResponse{}
	_body, _err := client.GetAppSeoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SEO索引图表
//
// @param request - GetAppSeoTrendsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSeoTrendsResponse
func (client *Client) GetAppSeoTrendsWithOptions(request *GetAppSeoTrendsRequest, runtime *dara.RuntimeOptions) (_result *GetAppSeoTrendsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.SeType) {
		query["SeType"] = request.SeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSeoTrends"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSeoTrendsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SEO索引图表
//
// @param request - GetAppSeoTrendsRequest
//
// @return GetAppSeoTrendsResponse
func (client *Client) GetAppSeoTrends(request *GetAppSeoTrendsRequest) (_result *GetAppSeoTrendsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSeoTrendsResponse{}
	_body, _err := client.GetAppSeoTrendsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取站点地图
//
// @param request - GetAppSitemapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSitemapResponse
func (client *Client) GetAppSitemapWithOptions(request *GetAppSitemapRequest, runtime *dara.RuntimeOptions) (_result *GetAppSitemapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.SeType) {
		query["SeType"] = request.SeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSitemap"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSitemapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取站点地图
//
// @param request - GetAppSitemapRequest
//
// @return GetAppSitemapResponse
func (client *Client) GetAppSitemap(request *GetAppSitemapRequest) (_result *GetAppSitemapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSitemapResponse{}
	_body, _err := client.GetAppSitemapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Supabase认证设置查询
//
// @param request - GetAppSupabaseAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSupabaseAuthConfigResponse
func (client *Client) GetAppSupabaseAuthConfigWithOptions(request *GetAppSupabaseAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *GetAppSupabaseAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSupabaseAuthConfig"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSupabaseAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Supabase认证设置查询
//
// @param request - GetAppSupabaseAuthConfigRequest
//
// @return GetAppSupabaseAuthConfigResponse
func (client *Client) GetAppSupabaseAuthConfig(request *GetAppSupabaseAuthConfigRequest) (_result *GetAppSupabaseAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSupabaseAuthConfigResponse{}
	_body, _err := client.GetAppSupabaseAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Supabase实例信息
//
// @param request - GetAppSupabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSupabaseInstanceResponse
func (client *Client) GetAppSupabaseInstanceWithOptions(request *GetAppSupabaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetAppSupabaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSupabaseInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSupabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Supabase实例信息
//
// @param request - GetAppSupabaseInstanceRequest
//
// @return GetAppSupabaseInstanceResponse
func (client *Client) GetAppSupabaseInstance(request *GetAppSupabaseInstanceRequest) (_result *GetAppSupabaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSupabaseInstanceResponse{}
	_body, _err := client.GetAppSupabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase密钥
//
// @param request - GetAppSupabaseSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSupabaseSecretsResponse
func (client *Client) GetAppSupabaseSecretsWithOptions(request *GetAppSupabaseSecretsRequest, runtime *dara.RuntimeOptions) (_result *GetAppSupabaseSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSupabaseSecrets"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSupabaseSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase密钥
//
// @param request - GetAppSupabaseSecretsRequest
//
// @return GetAppSupabaseSecretsResponse
func (client *Client) GetAppSupabaseSecrets(request *GetAppSupabaseSecretsRequest) (_result *GetAppSupabaseSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSupabaseSecretsResponse{}
	_body, _err := client.GetAppSupabaseSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模板详情查询
//
// @param request - GetAppTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppTemplateResponse
func (client *Client) GetAppTemplateWithOptions(request *GetAppTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAppTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppTemplate"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模板详情查询
//
// @param request - GetAppTemplateRequest
//
// @return GetAppTemplateResponse
func (client *Client) GetAppTemplate(request *GetAppTemplateRequest) (_result *GetAppTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppTemplateResponse{}
	_body, _err := client.GetAppTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询万小智灵感值服务
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppTokenServiceResponse
func (client *Client) GetAppTokenServiceWithOptions(runtime *dara.RuntimeOptions) (_result *GetAppTokenServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppTokenService"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppTokenServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询万小智灵感值服务
//
// @return GetAppTokenServiceResponse
func (client *Client) GetAppTokenService() (_result *GetAppTokenServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppTokenServiceResponse{}
	_body, _err := client.GetAppTokenServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作区目录结构
//
// @param request - GetAppWorkspaceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppWorkspaceDirectoryResponse
func (client *Client) GetAppWorkspaceDirectoryWithOptions(request *GetAppWorkspaceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *GetAppWorkspaceDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Deep) {
		body["Deep"] = request.Deep
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppWorkspaceDirectory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppWorkspaceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作区目录结构
//
// @param request - GetAppWorkspaceDirectoryRequest
//
// @return GetAppWorkspaceDirectoryResponse
func (client *Client) GetAppWorkspaceDirectory(request *GetAppWorkspaceDirectoryRequest) (_result *GetAppWorkspaceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppWorkspaceDirectoryResponse{}
	_body, _err := client.GetAppWorkspaceDirectoryWithOptions(request, runtime)
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
func (client *Client) GetCreateLogoTaskWithOptions(request *GetCreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCreateLogoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
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
// 提供给服务商的域名查询接口
//
// @param request - GetDomainInfoForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainInfoForPartnerResponse
func (client *Client) GetDomainInfoForPartnerWithOptions(request *GetDomainInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetDomainInfoForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomainInfoForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainInfoForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提供给服务商的域名查询接口
//
// @param request - GetDomainInfoForPartnerRequest
//
// @return GetDomainInfoForPartnerResponse
func (client *Client) GetDomainInfoForPartner(request *GetDomainInfoForPartnerRequest) (_result *GetDomainInfoForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDomainInfoForPartnerResponse{}
	_body, _err := client.GetDomainInfoForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名备案信息
//
// @param request - GetIcpFilingInfoForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIcpFilingInfoForPartnerResponse
func (client *Client) GetIcpFilingInfoForPartnerWithOptions(request *GetIcpFilingInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetIcpFilingInfoForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIcpFilingInfoForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIcpFilingInfoForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名备案信息
//
// @param request - GetIcpFilingInfoForPartnerRequest
//
// @return GetIcpFilingInfoForPartnerResponse
func (client *Client) GetIcpFilingInfoForPartner(request *GetIcpFilingInfoForPartnerRequest) (_result *GetIcpFilingInfoForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIcpFilingInfoForPartnerResponse{}
	_body, _err := client.GetIcpFilingInfoForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询LLM Proxy配置
//
// @param request - GetLlmProxyConfigForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLlmProxyConfigForAdminResponse
func (client *Client) GetLlmProxyConfigForAdminWithOptions(request *GetLlmProxyConfigForAdminRequest, runtime *dara.RuntimeOptions) (_result *GetLlmProxyConfigForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Capability) {
		query["Capability"] = request.Capability
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLlmProxyConfigForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLlmProxyConfigForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询LLM Proxy配置
//
// @param request - GetLlmProxyConfigForAdminRequest
//
// @return GetLlmProxyConfigForAdminResponse
func (client *Client) GetLlmProxyConfigForAdmin(request *GetLlmProxyConfigForAdminRequest) (_result *GetLlmProxyConfigForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLlmProxyConfigForAdminResponse{}
	_body, _err := client.GetLlmProxyConfigForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取小程序授权链接
//
// @param request - GetMiniAppAuthUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiniAppAuthUrlResponse
func (client *Client) GetMiniAppAuthUrlWithOptions(request *GetMiniAppAuthUrlRequest, runtime *dara.RuntimeOptions) (_result *GetMiniAppAuthUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.RedirectUri) {
		query["RedirectUri"] = request.RedirectUri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMiniAppAuthUrl"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMiniAppAuthUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取小程序授权链接
//
// @param request - GetMiniAppAuthUrlRequest
//
// @return GetMiniAppAuthUrlResponse
func (client *Client) GetMiniAppAuthUrl(request *GetMiniAppAuthUrlRequest) (_result *GetMiniAppAuthUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMiniAppAuthUrlResponse{}
	_body, _err := client.GetMiniAppAuthUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询站点绑定的小程序
//
// @param tmpReq - GetMiniAppBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiniAppBindingResponse
func (client *Client) GetMiniAppBindingWithOptions(tmpReq *GetMiniAppBindingRequest, runtime *dara.RuntimeOptions) (_result *GetMiniAppBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMiniAppBindingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SettingKeys) {
		request.SettingKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SettingKeys, dara.String("SettingKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.SettingKeysShrink) {
		query["SettingKeys"] = request.SettingKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMiniAppBinding"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMiniAppBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询站点绑定的小程序
//
// @param request - GetMiniAppBindingRequest
//
// @return GetMiniAppBindingResponse
func (client *Client) GetMiniAppBinding(request *GetMiniAppBindingRequest) (_result *GetMiniAppBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMiniAppBindingResponse{}
	_body, _err := client.GetMiniAppBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件查询应用实例绑定的小程序
//
// @param request - GetMiniAppBindingForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiniAppBindingForAdminResponse
func (client *Client) GetMiniAppBindingForAdminWithOptions(request *GetMiniAppBindingForAdminRequest, runtime *dara.RuntimeOptions) (_result *GetMiniAppBindingForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.PlatformAppid) {
		query["PlatformAppid"] = request.PlatformAppid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMiniAppBindingForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMiniAppBindingForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件查询应用实例绑定的小程序
//
// @param request - GetMiniAppBindingForAdminRequest
//
// @return GetMiniAppBindingForAdminResponse
func (client *Client) GetMiniAppBindingForAdmin(request *GetMiniAppBindingForAdminRequest) (_result *GetMiniAppBindingForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMiniAppBindingForAdminResponse{}
	_body, _err := client.GetMiniAppBindingForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过授权码得到accessToken
//
// @param request - GetUserAccessTokenForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAccessTokenForPartnerResponse
func (client *Client) GetUserAccessTokenForPartnerWithOptions(request *GetUserAccessTokenForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetUserAccessTokenForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteHost) {
		query["SiteHost"] = request.SiteHost
	}

	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAccessTokenForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAccessTokenForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过授权码得到accessToken
//
// @param request - GetUserAccessTokenForPartnerRequest
//
// @return GetUserAccessTokenForPartnerResponse
func (client *Client) GetUserAccessTokenForPartner(request *GetUserAccessTokenForPartnerRequest) (_result *GetUserAccessTokenForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserAccessTokenForPartnerResponse{}
	_body, _err := client.GetUserAccessTokenForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴获取用户SLR角色授权临时凭证
//
// @param request - GetUserTmpIdentityForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTmpIdentityForPartnerResponse
func (client *Client) GetUserTmpIdentityForPartnerWithOptions(request *GetUserTmpIdentityForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetUserTmpIdentityForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthPurpose) {
		query["AuthPurpose"] = request.AuthPurpose
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserTmpIdentityForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserTmpIdentityForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴获取用户SLR角色授权临时凭证
//
// @param request - GetUserTmpIdentityForPartnerRequest
//
// @return GetUserTmpIdentityForPartnerResponse
func (client *Client) GetUserTmpIdentityForPartner(request *GetUserTmpIdentityForPartnerRequest) (_result *GetUserTmpIdentityForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserTmpIdentityForPartnerResponse{}
	_body, _err := client.GetUserTmpIdentityForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查AccessToken
//
// @param request - IntrospectAppInstanceTicketForPreviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntrospectAppInstanceTicketForPreviewResponse
func (client *Client) IntrospectAppInstanceTicketForPreviewWithOptions(request *IntrospectAppInstanceTicketForPreviewRequest, runtime *dara.RuntimeOptions) (_result *IntrospectAppInstanceTicketForPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntrospectAppInstanceTicketForPreview"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntrospectAppInstanceTicketForPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查AccessToken
//
// @param request - IntrospectAppInstanceTicketForPreviewRequest
//
// @return IntrospectAppInstanceTicketForPreviewResponse
func (client *Client) IntrospectAppInstanceTicketForPreview(request *IntrospectAppInstanceTicketForPreviewRequest) (_result *IntrospectAppInstanceTicketForPreviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IntrospectAppInstanceTicketForPreviewResponse{}
	_body, _err := client.IntrospectAppInstanceTicketForPreviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI员工对话增量SSE事件
//
// @param request - ListAIStaffChatEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIStaffChatEventsResponse
func (client *Client) ListAIStaffChatEventsWithOptions(request *ListAIStaffChatEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAIStaffChatEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.LastEventId) {
		body["LastEventId"] = request.LastEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIStaffChatEvents"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIStaffChatEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI员工对话增量SSE事件
//
// @param request - ListAIStaffChatEventsRequest
//
// @return ListAIStaffChatEventsResponse
func (client *Client) ListAIStaffChatEvents(request *ListAIStaffChatEventsRequest) (_result *ListAIStaffChatEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIStaffChatEventsResponse{}
	_body, _err := client.ListAIStaffChatEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询AI员工对话消息列表
//
// @param request - ListAIStaffChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIStaffChatMessagesResponse
func (client *Client) ListAIStaffChatMessagesWithOptions(request *ListAIStaffChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListAIStaffChatMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartCreateTime) {
		body["StartCreateTime"] = request.StartCreateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIStaffChatMessages"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIStaffChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询AI员工对话消息列表
//
// @param request - ListAIStaffChatMessagesRequest
//
// @return ListAIStaffChatMessagesResponse
func (client *Client) ListAIStaffChatMessages(request *ListAIStaffChatMessagesRequest) (_result *ListAIStaffChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIStaffChatMessagesResponse{}
	_body, _err := client.ListAIStaffChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用助手智能体列表
//
// @param request - ListAppAssistantAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppAssistantAgentsResponse
func (client *Client) ListAppAssistantAgentsWithOptions(request *ListAppAssistantAgentsRequest, runtime *dara.RuntimeOptions) (_result *ListAppAssistantAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.PlatformType) {
		query["PlatformType"] = request.PlatformType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppAssistantAgents"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppAssistantAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用助手智能体列表
//
// @param request - ListAppAssistantAgentsRequest
//
// @return ListAppAssistantAgentsResponse
func (client *Client) ListAppAssistantAgents(request *ListAppAssistantAgentsRequest) (_result *ListAppAssistantAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppAssistantAgentsResponse{}
	_body, _err := client.ListAppAssistantAgentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定聊天的消息列表
//
// @param request - ListAppChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppChatMessagesResponse
func (client *Client) ListAppChatMessagesWithOptions(request *ListAppChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListAppChatMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		query["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SectionId) {
		query["SectionId"] = request.SectionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppChatMessages"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定聊天的消息列表
//
// @param request - ListAppChatMessagesRequest
//
// @return ListAppChatMessagesResponse
func (client *Client) ListAppChatMessages(request *ListAppChatMessagesRequest) (_result *ListAppChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppChatMessagesResponse{}
	_body, _err := client.ListAppChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListAppCommoditySpecificationsForPartner is deprecated, please use WebsiteBuild::2025-04-29::ListAppCommoditySpecificationsV2ForPartner instead.
//
// Summary:
//
// 获取商品配置信息
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppCommoditySpecificationsForPartnerResponse
// Deprecated
func (client *Client) ListAppCommoditySpecificationsForPartnerWithOptions(runtime *dara.RuntimeOptions) (_result *ListAppCommoditySpecificationsForPartnerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppCommoditySpecificationsForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppCommoditySpecificationsForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListAppCommoditySpecificationsForPartner is deprecated, please use WebsiteBuild::2025-04-29::ListAppCommoditySpecificationsV2ForPartner instead.
//
// Summary:
//
// 获取商品配置信息
//
// @return ListAppCommoditySpecificationsForPartnerResponse
// Deprecated
func (client *Client) ListAppCommoditySpecificationsForPartner() (_result *ListAppCommoditySpecificationsForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppCommoditySpecificationsForPartnerResponse{}
	_body, _err := client.ListAppCommoditySpecificationsForPartnerWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 网站信息查询
//
// @param request - ListAppCommoditySpecificationsV2ForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppCommoditySpecificationsV2ForPartnerResponse
func (client *Client) ListAppCommoditySpecificationsV2ForPartnerWithOptions(request *ListAppCommoditySpecificationsV2ForPartnerRequest, runtime *dara.RuntimeOptions) (_result *ListAppCommoditySpecificationsV2ForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppCommoditySpecificationsV2ForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppCommoditySpecificationsV2ForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网站信息查询
//
// @param request - ListAppCommoditySpecificationsV2ForPartnerRequest
//
// @return ListAppCommoditySpecificationsV2ForPartnerResponse
func (client *Client) ListAppCommoditySpecificationsV2ForPartner(request *ListAppCommoditySpecificationsV2ForPartnerRequest) (_result *ListAppCommoditySpecificationsV2ForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppCommoditySpecificationsV2ForPartnerResponse{}
	_body, _err := client.ListAppCommoditySpecificationsV2ForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询对话消息列表
//
// @param request - ListAppConversationMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppConversationMessagesResponse
func (client *Client) ListAppConversationMessagesWithOptions(request *ListAppConversationMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListAppConversationMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartCreateTime) {
		query["StartCreateTime"] = request.StartCreateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppConversationMessages"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppConversationMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询对话消息列表
//
// @param request - ListAppConversationMessagesRequest
//
// @return ListAppConversationMessagesResponse
func (client *Client) ListAppConversationMessages(request *ListAppConversationMessagesRequest) (_result *ListAppConversationMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppConversationMessagesResponse{}
	_body, _err := client.ListAppConversationMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索对话列表
//
// @param request - ListAppConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppConversationsResponse
func (client *Client) ListAppConversationsWithOptions(request *ListAppConversationsRequest, runtime *dara.RuntimeOptions) (_result *ListAppConversationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		query["BotId"] = request.BotId
	}

	if !dara.IsNil(request.EndModifyTime) {
		query["EndModifyTime"] = request.EndModifyTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.StartModifyTime) {
		query["StartModifyTime"] = request.StartModifyTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppConversations"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索对话列表
//
// @param request - ListAppConversationsRequest
//
// @return ListAppConversationsResponse
func (client *Client) ListAppConversations(request *ListAppConversationsRequest) (_result *ListAppConversationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppConversationsResponse{}
	_body, _err := client.ListAppConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of domain redirection rules
//
// @param request - ListAppDomainRedirectRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppDomainRedirectRecordsResponse
func (client *Client) ListAppDomainRedirectRecordsWithOptions(request *ListAppDomainRedirectRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListAppDomainRedirectRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppDomainRedirectRecords"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppDomainRedirectRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of domain redirection rules
//
// @param request - ListAppDomainRedirectRecordsRequest
//
// @return ListAppDomainRedirectRecordsResponse
func (client *Client) ListAppDomainRedirectRecords(request *ListAppDomainRedirectRecordsRequest) (_result *ListAppDomainRedirectRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppDomainRedirectRecordsResponse{}
	_body, _err := client.ListAppDomainRedirectRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List all domain names under the application instance
//
// @param request - ListAppInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstanceDomainsResponse
func (client *Client) ListAppInstanceDomainsWithOptions(request *ListAppInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainKeyword) {
		query["DomainKeyword"] = request.DomainKeyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstanceDomains"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List all domain names under the application instance
//
// @param request - ListAppInstanceDomainsRequest
//
// @return ListAppInstanceDomainsResponse
func (client *Client) ListAppInstanceDomains(request *ListAppInstanceDomainsRequest) (_result *ListAppInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppInstanceDomainsResponse{}
	_body, _err := client.ListAppInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Website Instance List Query
//
// @param tmpReq - ListAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstancesResponse
func (client *Client) ListAppInstancesWithOptions(tmpReq *ListAppInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAppInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.EndTimeBegin) {
		query["EndTimeBegin"] = request.EndTimeBegin
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstances"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Website Instance List Query
//
// @param request - ListAppInstancesRequest
//
// @return ListAppInstancesResponse
func (client *Client) ListAppInstances(request *ListAppInstancesRequest) (_result *ListAppInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.ListAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取码农插件配置列表
//
// @param request - ListAppPluginConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppPluginConfigsResponse
func (client *Client) ListAppPluginConfigsWithOptions(request *ListAppPluginConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListAppPluginConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppPluginConfigs"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppPluginConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取码农插件配置列表
//
// @param request - ListAppPluginConfigsRequest
//
// @return ListAppPluginConfigsResponse
func (client *Client) ListAppPluginConfigs(request *ListAppPluginConfigsRequest) (_result *ListAppPluginConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppPluginConfigsResponse{}
	_body, _err := client.ListAppPluginConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用插件列表
//
// @param request - ListAppPluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppPluginsResponse
func (client *Client) ListAppPluginsWithOptions(request *ListAppPluginsRequest, runtime *dara.RuntimeOptions) (_result *ListAppPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Phase) {
		query["Phase"] = request.Phase
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppPlugins"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用插件列表
//
// @param request - ListAppPluginsRequest
//
// @return ListAppPluginsResponse
func (client *Client) ListAppPlugins(request *ListAppPluginsRequest) (_result *ListAppPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppPluginsResponse{}
	_body, _err := client.ListAppPluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布历史查询
//
// @param request - ListAppPublishHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppPublishHistoryResponse
func (client *Client) ListAppPublishHistoryWithOptions(request *ListAppPublishHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListAppPublishHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WebsiteDomain) {
		query["WebsiteDomain"] = request.WebsiteDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppPublishHistory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppPublishHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布历史查询
//
// @param request - ListAppPublishHistoryRequest
//
// @return ListAppPublishHistoryResponse
func (client *Client) ListAppPublishHistory(request *ListAppPublishHistoryRequest) (_result *ListAppPublishHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppPublishHistoryResponse{}
	_body, _err := client.ListAppPublishHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 字典列表查询
//
// @param request - ListAppTemplateDictsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppTemplateDictsResponse
func (client *Client) ListAppTemplateDictsWithOptions(request *ListAppTemplateDictsRequest, runtime *dara.RuntimeOptions) (_result *ListAppTemplateDictsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DictType) {
		query["DictType"] = request.DictType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppTemplateDicts"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppTemplateDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 字典列表查询
//
// @param request - ListAppTemplateDictsRequest
//
// @return ListAppTemplateDictsResponse
func (client *Client) ListAppTemplateDicts(request *ListAppTemplateDictsRequest) (_result *ListAppTemplateDictsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppTemplateDictsResponse{}
	_body, _err := client.ListAppTemplateDictsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模板列表查询
//
// @param request - ListAppTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppTemplatesResponse
func (client *Client) ListAppTemplatesWithOptions(request *ListAppTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListAppTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ColorScheme) {
		query["ColorScheme"] = request.ColorScheme
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductVersion) {
		query["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppTemplates"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模板列表查询
//
// @param request - ListAppTemplatesRequest
//
// @return ListAppTemplatesResponse
func (client *Client) ListAppTemplates(request *ListAppTemplatesRequest) (_result *ListAppTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.ListAppTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付宝ISV插件配置
//
// @param request - ListIsvPaymentPluginConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIsvPaymentPluginConfigsResponse
func (client *Client) ListIsvPaymentPluginConfigsWithOptions(request *ListIsvPaymentPluginConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListIsvPaymentPluginConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIsvPaymentPluginConfigs"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIsvPaymentPluginConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付宝ISV插件配置
//
// @param request - ListIsvPaymentPluginConfigsRequest
//
// @return ListIsvPaymentPluginConfigsResponse
func (client *Client) ListIsvPaymentPluginConfigs(request *ListIsvPaymentPluginConfigsRequest) (_result *ListIsvPaymentPluginConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIsvPaymentPluginConfigsResponse{}
	_body, _err := client.ListIsvPaymentPluginConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify the configuration of a building instance
//
// @param request - ModifyAppInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppInstanceSpecResponse
func (client *Client) ModifyAppInstanceSpecWithOptions(request *ModifyAppInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeployArea) {
		query["DeployArea"] = request.DeployArea
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppInstanceSpec"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify the configuration of a building instance
//
// @param request - ModifyAppInstanceSpecRequest
//
// @return ModifyAppInstanceSpecResponse
func (client *Client) ModifyAppInstanceSpec(request *ModifyAppInstanceSpecRequest) (_result *ModifyAppInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppInstanceSpecResponse{}
	_body, _err := client.ModifyAppInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改素材中心文件夹
//
// @param request - ModifyMaterialDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaterialDirectoryResponse
func (client *Client) ModifyMaterialDirectoryWithOptions(request *ModifyMaterialDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaterialDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaterialDirectory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaterialDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改素材中心文件夹
//
// @param request - ModifyMaterialDirectoryRequest
//
// @return ModifyMaterialDirectoryResponse
func (client *Client) ModifyMaterialDirectory(request *ModifyMaterialDirectoryRequest) (_result *ModifyMaterialDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaterialDirectoryResponse{}
	_body, _err := client.ModifyMaterialDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改素材文件
//
// @param request - ModifyMaterialFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaterialFileResponse
func (client *Client) ModifyMaterialFileWithOptions(request *ModifyMaterialFileRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaterialFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaterialFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaterialFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改素材文件
//
// @param request - ModifyMaterialFileRequest
//
// @return ModifyMaterialFileResponse
func (client *Client) ModifyMaterialFile(request *ModifyMaterialFileRequest) (_result *ModifyMaterialFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaterialFileResponse{}
	_body, _err := client.ModifyMaterialFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改素材文件状态
//
// @param tmpReq - ModifyMaterialFileStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaterialFileStatusResponse
func (client *Client) ModifyMaterialFileStatusWithOptions(tmpReq *ModifyMaterialFileStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaterialFileStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyMaterialFileStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaterialFileStatus"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaterialFileStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改素材文件状态
//
// @param request - ModifyMaterialFileStatusRequest
//
// @return ModifyMaterialFileStatusResponse
func (client *Client) ModifyMaterialFileStatus(request *ModifyMaterialFileStatusRequest) (_result *ModifyMaterialFileStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaterialFileStatusResponse{}
	_body, _err := client.ModifyMaterialFileStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动素材中心文件夹
//
// @param request - MoveMaterialDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveMaterialDirectoryResponse
func (client *Client) MoveMaterialDirectoryWithOptions(request *MoveMaterialDirectoryRequest, runtime *dara.RuntimeOptions) (_result *MoveMaterialDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		query["ParentDirectoryId"] = request.ParentDirectoryId
	}

	if !dara.IsNil(request.SortNum) {
		query["SortNum"] = request.SortNum
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveMaterialDirectory"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveMaterialDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动素材中心文件夹
//
// @param request - MoveMaterialDirectoryRequest
//
// @return MoveMaterialDirectoryResponse
func (client *Client) MoveMaterialDirectory(request *MoveMaterialDirectoryRequest) (_result *MoveMaterialDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveMaterialDirectoryResponse{}
	_body, _err := client.MoveMaterialDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动素材文件
//
// @param tmpReq - MoveMaterialFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveMaterialFileResponse
func (client *Client) MoveMaterialFileWithOptions(tmpReq *MoveMaterialFileRequest, runtime *dara.RuntimeOptions) (_result *MoveMaterialFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MoveMaterialFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveMaterialFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveMaterialFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动素材文件
//
// @param request - MoveMaterialFileRequest
//
// @return MoveMaterialFileResponse
func (client *Client) MoveMaterialFile(request *MoveMaterialFileRequest) (_result *MoveMaterialFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveMaterialFileResponse{}
	_body, _err := client.MoveMaterialFileWithOptions(request, runtime)
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
func (client *Client) OperateAppInstanceForPartnerWithOptions(request *OperateAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) OperateAppServiceForPartnerWithOptions(request *OperateAppServiceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppServiceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppServiceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.OperateAppServiceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 切换模板点赞统计
//
// @param request - OperateAppTemplateLikeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppTemplateLikeResponse
func (client *Client) OperateAppTemplateLikeWithOptions(request *OperateAppTemplateLikeRequest, runtime *dara.RuntimeOptions) (_result *OperateAppTemplateLikeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Liked) {
		query["Liked"] = request.Liked
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppTemplateLike"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppTemplateLikeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 切换模板点赞统计
//
// @param request - OperateAppTemplateLikeRequest
//
// @return OperateAppTemplateLikeResponse
func (client *Client) OperateAppTemplateLike(request *OperateAppTemplateLikeRequest) (_result *OperateAppTemplateLikeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppTemplateLikeResponse{}
	_body, _err := client.OperateAppTemplateLikeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用Supabase操作
//
// @param request - OperateSupabaseForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateSupabaseForAdminResponse
func (client *Client) OperateSupabaseForAdminWithOptions(request *OperateSupabaseForAdminRequest, runtime *dara.RuntimeOptions) (_result *OperateSupabaseForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.ExecuteSql) {
		query["ExecuteSql"] = request.ExecuteSql
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OrderByClause) {
		query["OrderByClause"] = request.OrderByClause
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WhereClause) {
		query["WhereClause"] = request.WhereClause
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateSupabaseForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateSupabaseForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用Supabase操作
//
// @param request - OperateSupabaseForAdminRequest
//
// @return OperateSupabaseForAdminResponse
func (client *Client) OperateSupabaseForAdmin(request *OperateSupabaseForAdminRequest) (_result *OperateSupabaseForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateSupabaseForAdminResponse{}
	_body, _err := client.OperateSupabaseForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布应用实例
//
// @param request - PublishAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishAppInstanceResponse
func (client *Client) PublishAppInstanceWithOptions(request *PublishAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *PublishAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeployChannel) {
		query["DeployChannel"] = request.DeployChannel
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.LogicalNumber) {
		query["LogicalNumber"] = request.LogicalNumber
	}

	if !dara.IsNil(request.PublishNumber) {
		query["PublishNumber"] = request.PublishNumber
	}

	if !dara.IsNil(request.WeappAction) {
		query["WeappAction"] = request.WeappAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishAppInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布应用实例
//
// @param request - PublishAppInstanceRequest
//
// @return PublishAppInstanceResponse
func (client *Client) PublishAppInstance(request *PublishAppInstanceRequest) (_result *PublishAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishAppInstanceResponse{}
	_body, _err := client.PublishAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送资源计量数据
//
// @param request - PushResourceMeasureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourceMeasureResponse
func (client *Client) PushResourceMeasureWithOptions(request *PushResourceMeasureRequest, runtime *dara.RuntimeOptions) (_result *PushResourceMeasureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.BelongId) {
		query["BelongId"] = request.BelongId
	}

	if !dara.IsNil(request.BelongIdType) {
		query["BelongIdType"] = request.BelongIdType
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MeasureData) {
		query["MeasureData"] = request.MeasureData
	}

	if !dara.IsNil(request.MetaData) {
		query["MetaData"] = request.MetaData
	}

	if !dara.IsNil(request.ResourceCode) {
		query["ResourceCode"] = request.ResourceCode
	}

	if !dara.IsNil(request.UseTime) {
		query["UseTime"] = request.UseTime
	}

	if !dara.IsNil(request.UseType) {
		query["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushResourceMeasure"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResourceMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送资源计量数据
//
// @param request - PushResourceMeasureRequest
//
// @return PushResourceMeasureResponse
func (client *Client) PushResourceMeasure(request *PushResourceMeasureRequest) (_result *PushResourceMeasureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushResourceMeasureResponse{}
	_body, _err := client.PushResourceMeasureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询灵感值获取明细
//
// @param request - QueryInspirationAccountDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInspirationAccountDetailsResponse
func (client *Client) QueryInspirationAccountDetailsWithOptions(request *QueryInspirationAccountDetailsRequest, runtime *dara.RuntimeOptions) (_result *QueryInspirationAccountDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInspirationAccountDetails"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInspirationAccountDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询灵感值获取明细
//
// @param request - QueryInspirationAccountDetailsRequest
//
// @return QueryInspirationAccountDetailsResponse
func (client *Client) QueryInspirationAccountDetails(request *QueryInspirationAccountDetailsRequest) (_result *QueryInspirationAccountDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInspirationAccountDetailsResponse{}
	_body, _err := client.QueryInspirationAccountDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询灵感值余额总览
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInspirationBalanceResponse
func (client *Client) QueryInspirationBalanceWithOptions(runtime *dara.RuntimeOptions) (_result *QueryInspirationBalanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInspirationBalance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInspirationBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询灵感值余额总览
//
// @return QueryInspirationBalanceResponse
func (client *Client) QueryInspirationBalance() (_result *QueryInspirationBalanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInspirationBalanceResponse{}
	_body, _err := client.QueryInspirationBalanceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询灵感值消耗明细
//
// @param request - QueryInspirationConsumeRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInspirationConsumeRecordsResponse
func (client *Client) QueryInspirationConsumeRecordsWithOptions(request *QueryInspirationConsumeRecordsRequest, runtime *dara.RuntimeOptions) (_result *QueryInspirationConsumeRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInspirationConsumeRecords"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInspirationConsumeRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询灵感值消耗明细
//
// @param request - QueryInspirationConsumeRecordsRequest
//
// @return QueryInspirationConsumeRecordsResponse
func (client *Client) QueryInspirationConsumeRecords(request *QueryInspirationConsumeRecordsRequest) (_result *QueryInspirationConsumeRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInspirationConsumeRecordsResponse{}
	_body, _err := client.QueryInspirationConsumeRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材中心文件夹树结构
//
// @param request - QueryMaterialDirectoryTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialDirectoryTreeResponse
func (client *Client) QueryMaterialDirectoryTreeWithOptions(request *QueryMaterialDirectoryTreeRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialDirectoryTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.HiddenPublic) {
		query["HiddenPublic"] = request.HiddenPublic
	}

	if !dara.IsNil(request.Root) {
		query["Root"] = request.Root
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialDirectoryTree"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialDirectoryTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材中心文件夹树结构
//
// @param request - QueryMaterialDirectoryTreeRequest
//
// @return QueryMaterialDirectoryTreeResponse
func (client *Client) QueryMaterialDirectoryTree(request *QueryMaterialDirectoryTreeRequest) (_result *QueryMaterialDirectoryTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialDirectoryTreeResponse{}
	_body, _err := client.QueryMaterialDirectoryTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材文件详情
//
// @param request - QueryMaterialFileDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialFileDetailResponse
func (client *Client) QueryMaterialFileDetailWithOptions(request *QueryMaterialFileDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialFileDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialFileDetail"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialFileDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材文件详情
//
// @param request - QueryMaterialFileDetailRequest
//
// @return QueryMaterialFileDetailResponse
func (client *Client) QueryMaterialFileDetail(request *QueryMaterialFileDetailRequest) (_result *QueryMaterialFileDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialFileDetailResponse{}
	_body, _err := client.QueryMaterialFileDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Material File List
//
// @param tmpReq - QueryMaterialFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialFileListResponse
func (client *Client) QueryMaterialFileListWithOptions(tmpReq *QueryMaterialFileListRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMaterialFileListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SuffixList) {
		request.SuffixListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SuffixList, dara.String("SuffixList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TypeList) {
		request.TypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TypeList, dara.String("TypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxFileSize) {
		query["MaxFileSize"] = request.MaxFileSize
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MinFileSize) {
		query["MinFileSize"] = request.MinFileSize
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	if !dara.IsNil(request.SuffixListShrink) {
		query["SuffixList"] = request.SuffixListShrink
	}

	if !dara.IsNil(request.TypeListShrink) {
		query["TypeList"] = request.TypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialFileList"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Material File List
//
// @param request - QueryMaterialFileListRequest
//
// @return QueryMaterialFileListResponse
func (client *Client) QueryMaterialFileList(request *QueryMaterialFileListRequest) (_result *QueryMaterialFileListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialFileListResponse{}
	_body, _err := client.QueryMaterialFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材中心文件概要信息
//
// @param tmpReq - QueryMaterialFileSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialFileSummaryInfoResponse
func (client *Client) QueryMaterialFileSummaryInfoWithOptions(tmpReq *QueryMaterialFileSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialFileSummaryInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMaterialFileSummaryInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TypeList) {
		request.TypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TypeList, dara.String("TypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	if !dara.IsNil(request.TypeListShrink) {
		query["TypeList"] = request.TypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialFileSummaryInfo"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialFileSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材中心文件概要信息
//
// @param request - QueryMaterialFileSummaryInfoRequest
//
// @return QueryMaterialFileSummaryInfoResponse
func (client *Client) QueryMaterialFileSummaryInfo(request *QueryMaterialFileSummaryInfoRequest) (_result *QueryMaterialFileSummaryInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialFileSummaryInfoResponse{}
	_body, _err := client.QueryMaterialFileSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材生产任务详情
//
// @param request - QueryMaterialTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialTaskDetailResponse
func (client *Client) QueryMaterialTaskDetailWithOptions(request *QueryMaterialTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialTaskDetail"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材生产任务详情
//
// @param request - QueryMaterialTaskDetailRequest
//
// @return QueryMaterialTaskDetailResponse
func (client *Client) QueryMaterialTaskDetail(request *QueryMaterialTaskDetailRequest) (_result *QueryMaterialTaskDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialTaskDetailResponse{}
	_body, _err := client.QueryMaterialTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材生产任务列表
//
// @param tmpReq - QueryMaterialTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMaterialTaskListResponse
func (client *Client) QueryMaterialTaskListWithOptions(tmpReq *QueryMaterialTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryMaterialTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMaterialTaskListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskTypeList) {
		request.TaskTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypeList, dara.String("TaskTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	if !dara.IsNil(request.TaskTypeListShrink) {
		query["TaskTypeList"] = request.TaskTypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMaterialTaskList"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMaterialTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材生产任务列表
//
// @param request - QueryMaterialTaskListRequest
//
// @return QueryMaterialTaskListResponse
func (client *Client) QueryMaterialTaskList(request *QueryMaterialTaskListRequest) (_result *QueryMaterialTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMaterialTaskListResponse{}
	_body, _err := client.QueryMaterialTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase Auth配置信息
//
// @param request - QuerySupabaseAuthConfigsForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySupabaseAuthConfigsForAdminResponse
func (client *Client) QuerySupabaseAuthConfigsForAdminWithOptions(request *QuerySupabaseAuthConfigsForAdminRequest, runtime *dara.RuntimeOptions) (_result *QuerySupabaseAuthConfigsForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySupabaseAuthConfigsForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySupabaseAuthConfigsForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase Auth配置信息
//
// @param request - QuerySupabaseAuthConfigsForAdminRequest
//
// @return QuerySupabaseAuthConfigsForAdminResponse
func (client *Client) QuerySupabaseAuthConfigsForAdmin(request *QuerySupabaseAuthConfigsForAdminRequest) (_result *QuerySupabaseAuthConfigsForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySupabaseAuthConfigsForAdminResponse{}
	_body, _err := client.QuerySupabaseAuthConfigsForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase配置信息
//
// @param request - QuerySupabaseConfigsForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySupabaseConfigsForAdminResponse
func (client *Client) QuerySupabaseConfigsForAdminWithOptions(request *QuerySupabaseConfigsForAdminRequest, runtime *dara.RuntimeOptions) (_result *QuerySupabaseConfigsForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySupabaseConfigsForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySupabaseConfigsForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase配置信息
//
// @param request - QuerySupabaseConfigsForAdminRequest
//
// @return QuerySupabaseConfigsForAdminResponse
func (client *Client) QuerySupabaseConfigsForAdmin(request *QuerySupabaseConfigsForAdminRequest) (_result *QuerySupabaseConfigsForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySupabaseConfigsForAdminResponse{}
	_body, _err := client.QuerySupabaseConfigsForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase实例信息
//
// @param request - QuerySupabaseInstanceInfoForAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySupabaseInstanceInfoForAdminResponse
func (client *Client) QuerySupabaseInstanceInfoForAdminWithOptions(request *QuerySupabaseInstanceInfoForAdminRequest, runtime *dara.RuntimeOptions) (_result *QuerySupabaseInstanceInfoForAdminResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySupabaseInstanceInfoForAdmin"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySupabaseInstanceInfoForAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase实例信息
//
// @param request - QuerySupabaseInstanceInfoForAdminRequest
//
// @return QuerySupabaseInstanceInfoForAdminResponse
func (client *Client) QuerySupabaseInstanceInfoForAdmin(request *QuerySupabaseInstanceInfoForAdminRequest) (_result *QuerySupabaseInstanceInfoForAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySupabaseInstanceInfoForAdminResponse{}
	_body, _err := client.QuerySupabaseInstanceInfoForAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新连接AI对话
//
// @param request - ReconnectAppChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReconnectAppChatResponse
func (client *Client) ReconnectAppChatWithSSE(request *ReconnectAppChatRequest, runtime *dara.RuntimeOptions, _yield chan *ReconnectAppChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.reconnectAppChatWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// 重新连接AI对话
//
// @param request - ReconnectAppChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReconnectAppChatResponse
func (client *Client) ReconnectAppChatWithOptions(request *ReconnectAppChatRequest, runtime *dara.RuntimeOptions) (_result *ReconnectAppChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		query["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.LastEventId) {
		query["LastEventId"] = request.LastEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReconnectAppChat"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &ReconnectAppChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新连接AI对话
//
// @param request - ReconnectAppChatRequest
//
// @return ReconnectAppChatResponse
func (client *Client) ReconnectAppChat(request *ReconnectAppChatRequest) (_result *ReconnectAppChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReconnectAppChatResponse{}
	_body, _err := client.ReconnectAppChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Refresh ticket
//
// @param request - RefreshAppInstanceTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAppInstanceTicketResponse
func (client *Client) RefreshAppInstanceTicketWithOptions(request *RefreshAppInstanceTicketRequest, runtime *dara.RuntimeOptions) (_result *RefreshAppInstanceTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshAppInstanceTicket"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshAppInstanceTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Refresh ticket
//
// @param request - RefreshAppInstanceTicketRequest
//
// @return RefreshAppInstanceTicketResponse
func (client *Client) RefreshAppInstanceTicket(request *RefreshAppInstanceTicketRequest) (_result *RefreshAppInstanceTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshAppInstanceTicketResponse{}
	_body, _err := client.RefreshAppInstanceTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 渠道业务退款接口
//
// @param request - RefundAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundAppInstanceForPartnerResponse
func (client *Client) RefundAppInstanceForPartnerWithOptions(request *RefundAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *RefundAppInstanceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RefundReason) {
		query["RefundReason"] = request.RefundReason
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渠道业务退款接口
//
// @param request - RefundAppInstanceForPartnerRequest
//
// @return RefundAppInstanceForPartnerResponse
func (client *Client) RefundAppInstanceForPartner(request *RefundAppInstanceForPartnerRequest) (_result *RefundAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefundAppInstanceForPartnerResponse{}
	_body, _err := client.RefundAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Renewal of website building instance
//
// @param request - RenewAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppInstanceResponse
func (client *Client) RenewAppInstanceWithOptions(request *RenewAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAppInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Renewal of website building instance
//
// @param request - RenewAppInstanceRequest
//
// @return RenewAppInstanceResponse
func (client *Client) RenewAppInstance(request *RenewAppInstanceRequest) (_result *RenewAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAppInstanceResponse{}
	_body, _err := client.RenewAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 续期/刷新沙箱环境
//
// @param request - RenewAppSandboxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppSandboxResponse
func (client *Client) RenewAppSandboxWithOptions(request *RenewAppSandboxRequest, runtime *dara.RuntimeOptions) (_result *RenewAppSandboxResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAppSandbox"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAppSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 续期/刷新沙箱环境
//
// @param request - RenewAppSandboxRequest
//
// @return RenewAppSandboxResponse
func (client *Client) RenewAppSandbox(request *RenewAppSandboxRequest) (_result *RenewAppSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAppSandboxResponse{}
	_body, _err := client.RenewAppSandboxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 代码快照回滚
//
// @param request - RollbackAppCodeSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackAppCodeSnapshotResponse
func (client *Client) RollbackAppCodeSnapshotWithOptions(request *RollbackAppCodeSnapshotRequest, runtime *dara.RuntimeOptions) (_result *RollbackAppCodeSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SiteId) {
		query["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TargetLogicalNumber) {
		query["TargetLogicalNumber"] = request.TargetLogicalNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackAppCodeSnapshot"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackAppCodeSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 代码快照回滚
//
// @param request - RollbackAppCodeSnapshotRequest
//
// @return RollbackAppCodeSnapshotResponse
func (client *Client) RollbackAppCodeSnapshot(request *RollbackAppCodeSnapshotRequest) (_result *RollbackAppCodeSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackAppCodeSnapshotResponse{}
	_body, _err := client.RollbackAppCodeSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 回滚应用实例发布
//
// @param request - RollbackAppInstancePublishRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackAppInstancePublishResponse
func (client *Client) RollbackAppInstancePublishWithOptions(request *RollbackAppInstancePublishRequest, runtime *dara.RuntimeOptions) (_result *RollbackAppInstancePublishResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeployChannel) {
		query["DeployChannel"] = request.DeployChannel
	}

	if !dara.IsNil(request.PublishNumber) {
		query["PublishNumber"] = request.PublishNumber
	}

	if !dara.IsNil(request.QuickRollback) {
		query["QuickRollback"] = request.QuickRollback
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackAppInstancePublish"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackAppInstancePublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回滚应用实例发布
//
// @param request - RollbackAppInstancePublishRequest
//
// @return RollbackAppInstancePublishResponse
func (client *Client) RollbackAppInstancePublish(request *RollbackAppInstancePublishRequest) (_result *RollbackAppInstancePublishResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackAppInstancePublishResponse{}
	_body, _err := client.RollbackAppInstancePublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存需求
//
// @param request - SaveAppRequirementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAppRequirementResponse
func (client *Client) SaveAppRequirementWithOptions(request *SaveAppRequirementRequest, runtime *dara.RuntimeOptions) (_result *SaveAppRequirementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prd) {
		body["Prd"] = request.Prd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAppRequirement"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAppRequirementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存需求
//
// @param request - SaveAppRequirementRequest
//
// @return SaveAppRequirementResponse
func (client *Client) SaveAppRequirement(request *SaveAppRequirementRequest) (_result *SaveAppRequirementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveAppRequirementResponse{}
	_body, _err := client.SaveAppRequirementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存Supabase密钥
//
// @param request - SaveAppSupabaseSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAppSupabaseSecretsResponse
func (client *Client) SaveAppSupabaseSecretsWithOptions(request *SaveAppSupabaseSecretsRequest, runtime *dara.RuntimeOptions) (_result *SaveAppSupabaseSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.SecretsJson) {
		query["SecretsJson"] = request.SecretsJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAppSupabaseSecrets"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAppSupabaseSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存Supabase密钥
//
// @param request - SaveAppSupabaseSecretsRequest
//
// @return SaveAppSupabaseSecretsResponse
func (client *Client) SaveAppSupabaseSecrets(request *SaveAppSupabaseSecretsRequest) (_result *SaveAppSupabaseSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveAppSupabaseSecretsResponse{}
	_body, _err := client.SaveAppSupabaseSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片检索
//
// @param tmpReq - SearchImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageResponse
func (client *Client) SearchImageWithOptions(tmpReq *SearchImageRequest, runtime *dara.RuntimeOptions) (_result *SearchImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ColorHex) {
		query["ColorHex"] = request.ColorHex
	}

	if !dara.IsNil(request.HasPerson) {
		query["HasPerson"] = request.HasPerson
	}

	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !dara.IsNil(request.ImageRatio) {
		query["ImageRatio"] = request.ImageRatio
	}

	if !dara.IsNil(request.MaxHeight) {
		query["MaxHeight"] = request.MaxHeight
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MaxWidth) {
		query["MaxWidth"] = request.MaxWidth
	}

	if !dara.IsNil(request.MinHeight) {
		query["MinHeight"] = request.MinHeight
	}

	if !dara.IsNil(request.MinWidth) {
		query["MinWidth"] = request.MinWidth
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImage"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片检索
//
// @param request - SearchImageRequest
//
// @return SearchImageResponse
func (client *Client) SearchImage(request *SearchImageRequest) (_result *SearchImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageResponse{}
	_body, _err := client.SearchImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Set the SSL certificate for a domain
//
// @param request - SetAppDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppDomainCertificateResponse
func (client *Client) SetAppDomainCertificateWithOptions(request *SetAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetAppDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificateType) {
		query["CertificateType"] = request.CertificateType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppDomainCertificate"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Set the SSL certificate for a domain
//
// @param request - SetAppDomainCertificateRequest
//
// @return SetAppDomainCertificateResponse
func (client *Client) SetAppDomainCertificate(request *SetAppDomainCertificateRequest) (_result *SetAppDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAppDomainCertificateResponse{}
	_body, _err := client.SetAppDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交SEO索引
//
// @param request - SubmitAppSeoIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAppSeoIndexResponse
func (client *Client) SubmitAppSeoIndexWithOptions(request *SubmitAppSeoIndexRequest, runtime *dara.RuntimeOptions) (_result *SubmitAppSeoIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.SeType) {
		query["SeType"] = request.SeType
	}

	if !dara.IsNil(request.SubmitLater) {
		query["SubmitLater"] = request.SubmitLater
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAppSeoIndex"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAppSeoIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交SEO索引
//
// @param request - SubmitAppSeoIndexRequest
//
// @return SubmitAppSeoIndexResponse
func (client *Client) SubmitAppSeoIndex(request *SubmitAppSeoIndexRequest) (_result *SubmitAppSeoIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAppSeoIndexResponse{}
	_body, _err := client.SubmitAppSeoIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交素材生产任务
//
// @param request - SubmitMaterialTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMaterialTaskResponse
func (client *Client) SubmitMaterialTaskWithOptions(request *SubmitMaterialTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitMaterialTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskParam) {
		query["TaskParam"] = request.TaskParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMaterialTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMaterialTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交素材生产任务
//
// @param request - SubmitMaterialTaskRequest
//
// @return SubmitMaterialTaskResponse
func (client *Client) SubmitMaterialTask(request *SubmitMaterialTaskRequest) (_result *SubmitMaterialTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMaterialTaskResponse{}
	_body, _err := client.SubmitMaterialTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 切换到指定对话
//
// @param request - SwitchAppConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchAppConversationResponse
func (client *Client) SwitchAppConversationWithOptions(request *SwitchAppConversationRequest, runtime *dara.RuntimeOptions) (_result *SwitchAppConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BotId) {
		query["BotId"] = request.BotId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchAppConversation"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchAppConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 切换到指定对话
//
// @param request - SwitchAppConversationRequest
//
// @return SwitchAppConversationResponse
func (client *Client) SwitchAppConversation(request *SwitchAppConversationRequest) (_result *SwitchAppConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchAppConversationResponse{}
	_body, _err := client.SwitchAppConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴同步应用实例
//
// @param tmpReq - SyncAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartnerWithOptions(tmpReq *SyncAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncAppInstanceForPartnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppInstance) {
		request.AppInstanceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInstance, dara.String("AppInstance"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceShrink) {
		query["AppInstance"] = request.AppInstanceShrink
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	if !dara.IsNil(request.SourceBizId) {
		query["SourceBizId"] = request.SourceBizId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴同步应用实例
//
// @param request - SyncAppInstanceForPartnerRequest
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartner(request *SyncAppInstanceForPartnerRequest) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.SyncAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Unbind Application Domain
//
// @param request - UnbindAppDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAppDomainResponse
func (client *Client) UnbindAppDomainWithOptions(request *UnbindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *UnbindAppDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAppDomain"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAppDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unbind Application Domain
//
// @param request - UnbindAppDomainRequest
//
// @return UnbindAppDomainResponse
func (client *Client) UnbindAppDomain(request *UnbindAppDomainRequest) (_result *UnbindAppDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAppDomainResponse{}
	_body, _err := client.UnbindAppDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新消息内容
//
// @param request - UpdateAppChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppChatMessageResponse
func (client *Client) UpdateAppChatMessageWithOptions(request *UpdateAppChatMessageRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddedMetaData) {
		query["AddedMetaData"] = request.AddedMetaData
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppChatMessage"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消息内容
//
// @param request - UpdateAppChatMessageRequest
//
// @return UpdateAppChatMessageResponse
func (client *Client) UpdateAppChatMessage(request *UpdateAppChatMessageRequest) (_result *UpdateAppChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppChatMessageResponse{}
	_body, _err := client.UpdateAppChatMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑JSX代码
//
// @param request - UpdateAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppCodeResponse
func (client *Client) UpdateAppCodeWithOptions(request *UpdateAppCodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppCode"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑JSX代码
//
// @param request - UpdateAppCodeRequest
//
// @return UpdateAppCodeResponse
func (client *Client) UpdateAppCode(request *UpdateAppCodeRequest) (_result *UpdateAppCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppCodeResponse{}
	_body, _err := client.UpdateAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文件
//
// @param request - UpdateAppFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppFileResponse
func (client *Client) UpdateAppFileWithOptions(request *UpdateAppFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文件
//
// @param request - UpdateAppFileRequest
//
// @return UpdateAppFileResponse
func (client *Client) UpdateAppFile(request *UpdateAppFileRequest) (_result *UpdateAppFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppFileResponse{}
	_body, _err := client.UpdateAppFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 建站实例变配
//
// @param tmpReq - UpdateAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppInstanceResponse
func (client *Client) UpdateAppInstanceWithOptions(tmpReq *UpdateAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAppInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeployArea) {
		query["DeployArea"] = request.DeployArea
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.IconUrl) {
		query["IconUrl"] = request.IconUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.SiteVersion) {
		query["SiteVersion"] = request.SiteVersion
	}

	if !dara.IsNil(request.ThumbnailUrl) {
		query["ThumbnailUrl"] = request.ThumbnailUrl
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppInstance"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 建站实例变配
//
// @param request - UpdateAppInstanceRequest
//
// @return UpdateAppInstanceResponse
func (client *Client) UpdateAppInstance(request *UpdateAppInstanceRequest) (_result *UpdateAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppInstanceResponse{}
	_body, _err := client.UpdateAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update SEO Index Status
//
// Description:
//
// # WanXiaoZhi 2.0 AI Conversation
//
// @param request - UpdateAppSeoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppSeoStatusResponse
func (client *Client) UpdateAppSeoStatusWithOptions(request *UpdateAppSeoStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppSeoStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.SeType) {
		query["SeType"] = request.SeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppSeoStatus"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppSeoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update SEO Index Status
//
// Description:
//
// # WanXiaoZhi 2.0 AI Conversation
//
// @param request - UpdateAppSeoStatusRequest
//
// @return UpdateAppSeoStatusResponse
func (client *Client) UpdateAppSeoStatus(request *UpdateAppSeoStatusRequest) (_result *UpdateAppSeoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppSeoStatusResponse{}
	_body, _err := client.UpdateAppSeoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Supabase认证配置更新
//
// @param request - UpdateAppSupabaseAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppSupabaseAuthConfigResponse
func (client *Client) UpdateAppSupabaseAuthConfigWithOptions(request *UpdateAppSupabaseAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppSupabaseAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ConfigsJson) {
		query["ConfigsJson"] = request.ConfigsJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppSupabaseAuthConfig"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppSupabaseAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Supabase认证配置更新
//
// @param request - UpdateAppSupabaseAuthConfigRequest
//
// @return UpdateAppSupabaseAuthConfigResponse
func (client *Client) UpdateAppSupabaseAuthConfig(request *UpdateAppSupabaseAuthConfigRequest) (_result *UpdateAppSupabaseAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppSupabaseAuthConfigResponse{}
	_body, _err := client.UpdateAppSupabaseAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Supabase密钥
//
// @param request - UpdateAppSupabaseSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppSupabaseSecretResponse
func (client *Client) UpdateAppSupabaseSecretWithOptions(request *UpdateAppSupabaseSecretRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppSupabaseSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.SecretKey) {
		query["SecretKey"] = request.SecretKey
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.SecretType) {
		query["SecretType"] = request.SecretType
	}

	if !dara.IsNil(request.SecretValue) {
		query["SecretValue"] = request.SecretValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppSupabaseSecret"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppSupabaseSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Supabase密钥
//
// @param request - UpdateAppSupabaseSecretRequest
//
// @return UpdateAppSupabaseSecretResponse
func (client *Client) UpdateAppSupabaseSecret(request *UpdateAppSupabaseSecretRequest) (_result *UpdateAppSupabaseSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppSupabaseSecretResponse{}
	_body, _err := client.UpdateAppSupabaseSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新绑定小程序信息
//
// @param request - UpdateMiniAppBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMiniAppBindingResponse
func (client *Client) UpdateMiniAppBindingWithOptions(request *UpdateMiniAppBindingRequest, runtime *dara.RuntimeOptions) (_result *UpdateMiniAppBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.SettingKey) {
		query["SettingKey"] = request.SettingKey
	}

	if !dara.IsNil(request.SettingValue) {
		query["SettingValue"] = request.SettingValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMiniAppBinding"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMiniAppBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新绑定小程序信息
//
// @param request - UpdateMiniAppBindingRequest
//
// @return UpdateMiniAppBindingResponse
func (client *Client) UpdateMiniAppBinding(request *UpdateMiniAppBindingRequest) (_result *UpdateMiniAppBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMiniAppBindingResponse{}
	_body, _err := client.UpdateMiniAppBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传到站点根目录
//
// @param request - UploadAppSiteValidationFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAppSiteValidationFileResponse
func (client *Client) UploadAppSiteValidationFileWithOptions(request *UploadAppSiteValidationFileRequest, runtime *dara.RuntimeOptions) (_result *UploadAppSiteValidationFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.File) {
		query["File"] = request.File
	}

	if !dara.IsNil(request.FileContent) {
		query["FileContent"] = request.FileContent
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.SiteHost) {
		query["SiteHost"] = request.SiteHost
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadAppSiteValidationFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadAppSiteValidationFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传到站点根目录
//
// @param request - UploadAppSiteValidationFileRequest
//
// @return UploadAppSiteValidationFileResponse
func (client *Client) UploadAppSiteValidationFile(request *UploadAppSiteValidationFileRequest) (_result *UploadAppSiteValidationFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadAppSiteValidationFileResponse{}
	_body, _err := client.UploadAppSiteValidationFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传素材文件
//
// @param request - UploadMaterialFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMaterialFileResponse
func (client *Client) UploadMaterialFileWithOptions(request *UploadMaterialFileRequest, runtime *dara.RuntimeOptions) (_result *UploadMaterialFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMaterialFile"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMaterialFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传素材文件
//
// @param request - UploadMaterialFileRequest
//
// @return UploadMaterialFileResponse
func (client *Client) UploadMaterialFile(request *UploadMaterialFileRequest) (_result *UploadMaterialFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadMaterialFileResponse{}
	_body, _err := client.UploadMaterialFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) createAppChatWithSSE_opYieldFunc(_yield chan *CreateAppChatResponse, _yieldErr chan error, request *CreateAppChatRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["BotId"] = request.BotId
	}

	if !dara.IsNil(request.ChatId) {
		body["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppChat"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.StringValue(resp.Event.Data)
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}

func (client *Client) reconnectAppChatWithSSE_opYieldFunc(_yield chan *ReconnectAppChatResponse, _yieldErr chan error, request *ReconnectAppChatRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		query["ChatId"] = request.ChatId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.LastEventId) {
		query["LastEventId"] = request.LastEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReconnectAppChat"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.StringValue(resp.Event.Data)
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
