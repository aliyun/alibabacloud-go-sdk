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
	client.Endpoint, _err = client.GetEndpoint(dara.String("sfmmultimodalapp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 指令创建
//
// @param tmpReq - CreateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommandResponse
func (client *Client) CreateCommandWithOptions(tmpReq *CreateCommandRequest, runtime *dara.RuntimeOptions) (_result *CreateCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ToolExamples) {
		request.ToolExamplesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToolExamples, dara.String("ToolExamples"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ToolParams) {
		request.ToolParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToolParams, dara.String("ToolParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ToolDescription) {
		query["ToolDescription"] = request.ToolDescription
	}

	if !dara.IsNil(request.ToolExamplesShrink) {
		query["ToolExamples"] = request.ToolExamplesShrink
	}

	if !dara.IsNil(request.ToolName) {
		query["ToolName"] = request.ToolName
	}

	if !dara.IsNil(request.ToolParamsShrink) {
		query["ToolParams"] = request.ToolParamsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCommand"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指令创建
//
// @param request - CreateCommandRequest
//
// @return CreateCommandResponse
func (client *Client) CreateCommand(request *CreateCommandRequest) (_result *CreateCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCommandResponse{}
	_body, _err := client.CreateCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增用户记忆
//
// @param tmpReq - CreateMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithOptions(tmpReq *CreateMemoryRequest, runtime *dara.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMemoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetaData) {
		request.MetaDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetaData, dara.String("MetaData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoUpdate) {
		query["AutoUpdate"] = request.AutoUpdate
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
	}

	if !dara.IsNil(request.MessagesJson) {
		query["MessagesJson"] = request.MessagesJson
	}

	if !dara.IsNil(request.MetaDataShrink) {
		query["MetaData"] = request.MetaDataShrink
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemory"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增用户记忆
//
// @param request - CreateMemoryRequest
//
// @return CreateMemoryResponse
func (client *Client) CreateMemory(request *CreateMemoryRequest) (_result *CreateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMemoryResponse{}
	_body, _err := client.CreateMemoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建多模态应用
//
// @param tmpReq - CreateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmAppResponse
func (client *Client) CreateMmAppWithOptions(tmpReq *CreateMmAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMmAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BindingConfig) {
		request.BindingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BindingConfig, dara.String("BindingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ConversationConfig) {
		request.ConversationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConversationConfig, dara.String("ConversationConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ModelConfig) {
		request.ModelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelConfig, dara.String("ModelConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BindingConfigShrink) {
		query["BindingConfig"] = request.BindingConfigShrink
	}

	if !dara.IsNil(request.ConversationConfigShrink) {
		query["ConversationConfig"] = request.ConversationConfigShrink
	}

	if !dara.IsNil(request.ModelConfigShrink) {
		query["ModelConfig"] = request.ModelConfigShrink
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建多模态应用
//
// @param request - CreateMmAppRequest
//
// @return CreateMmAppResponse
func (client *Client) CreateMmApp(request *CreateMmAppRequest) (_result *CreateMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMmAppResponse{}
	_body, _err := client.CreateMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户画像配置
//
// @param tmpReq - CreateProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProfileResponse
func (client *Client) CreateProfileWithOptions(tmpReq *CreateProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AttributesShrink) {
		query["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProfile"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户画像配置
//
// @param request - CreateProfileRequest
//
// @return CreateProfileResponse
func (client *Client) CreateProfile(request *CreateProfileRequest) (_result *CreateProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProfileResponse{}
	_body, _err := client.CreateProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指令
//
// @param request - DeleteCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommandWithOptions(request *DeleteCommandRequest, runtime *dara.RuntimeOptions) (_result *DeleteCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.ToolId) {
		query["ToolId"] = request.ToolId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCommand"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指令
//
// @param request - DeleteCommandRequest
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommand(request *DeleteCommandRequest) (_result *DeleteCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCommandResponse{}
	_body, _err := client.DeleteCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户记忆
//
// @param request - DeleteMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(request *DeleteMemoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.MemoryNodeId) {
		query["MemoryNodeId"] = request.MemoryNodeId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户记忆
//
// @param request - DeleteMemoryRequest
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(request *DeleteMemoryRequest) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除多模态应用
//
// @param request - DeleteMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmAppResponse
func (client *Client) DeleteMmAppWithOptions(request *DeleteMmAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除多模态应用
//
// @param request - DeleteMmAppRequest
//
// @return DeleteMmAppResponse
func (client *Client) DeleteMmApp(request *DeleteMmAppRequest) (_result *DeleteMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMmAppResponse{}
	_body, _err := client.DeleteMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户画像配置
//
// @param request - DeleteProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProfileResponse
func (client *Client) DeleteProfileWithOptions(request *DeleteProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProfile"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户画像配置
//
// @param request - DeleteProfileRequest
//
// @return DeleteProfileResponse
func (client *Client) DeleteProfile(request *DeleteProfileRequest) (_result *DeleteProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProfileResponse{}
	_body, _err := client.DeleteProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指令详情
//
// @param request - DescribeCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommandResponse
func (client *Client) DescribeCommandWithOptions(request *DescribeCommandRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.ToolId) {
		query["ToolId"] = request.ToolId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCommand"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指令详情
//
// @param request - DescribeCommandRequest
//
// @return DescribeCommandResponse
func (client *Client) DescribeCommand(request *DescribeCommandRequest) (_result *DescribeCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCommandResponse{}
	_body, _err := client.DescribeCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态应用详情
//
// @param request - DescribeMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMmAppResponse
func (client *Client) DescribeMmAppWithOptions(request *DescribeMmAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态应用详情
//
// @param request - DescribeMmAppRequest
//
// @return DescribeMmAppResponse
func (client *Client) DescribeMmApp(request *DescribeMmAppRequest) (_result *DescribeMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMmAppResponse{}
	_body, _err := client.DescribeMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指令列表
//
// @param request - ListCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommandResponse
func (client *Client) ListCommandWithOptions(request *ListCommandRequest, runtime *dara.RuntimeOptions) (_result *ListCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ToolName) {
		query["ToolName"] = request.ToolName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommand"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指令列表
//
// @param request - ListCommandRequest
//
// @return ListCommandResponse
func (client *Client) ListCommand(request *ListCommandRequest) (_result *ListCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCommandResponse{}
	_body, _err := client.ListCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取多模态应用列表
//
// @param request - ListMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmAppResponse
func (client *Client) ListMmAppWithOptions(request *ListMmAppRequest, runtime *dara.RuntimeOptions) (_result *ListMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取多模态应用列表
//
// @param request - ListMmAppRequest
//
// @return ListMmAppResponse
func (client *Client) ListMmApp(request *ListMmAppRequest) (_result *ListMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMmAppResponse{}
	_body, _err := client.ListMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用发布列表
//
// @param request - ListPublishedMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedMmAppResponse
func (client *Client) ListPublishedMmAppWithOptions(request *ListPublishedMmAppRequest, runtime *dara.RuntimeOptions) (_result *ListPublishedMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用发布列表
//
// @param request - ListPublishedMmAppRequest
//
// @return ListPublishedMmAppResponse
func (client *Client) ListPublishedMmApp(request *ListPublishedMmAppRequest) (_result *ListPublishedMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPublishedMmAppResponse{}
	_body, _err := client.ListPublishedMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态应用绑定MCP
//
// @param tmpReq - MmAppBindingMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MmAppBindingMcpResponse
func (client *Client) MmAppBindingMcpWithOptions(tmpReq *MmAppBindingMcpRequest, runtime *dara.RuntimeOptions) (_result *MmAppBindingMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MmAppBindingMcpShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Mcps) {
		request.McpsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Mcps, dara.String("Mcps"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.McpsShrink) {
		query["Mcps"] = request.McpsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MmAppBindingMcp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MmAppBindingMcpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态应用绑定MCP
//
// @param request - MmAppBindingMcpRequest
//
// @return MmAppBindingMcpResponse
func (client *Client) MmAppBindingMcp(request *MmAppBindingMcpRequest) (_result *MmAppBindingMcpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MmAppBindingMcpResponse{}
	_body, _err := client.MmAppBindingMcpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态应用绑定知识库
//
// @param tmpReq - MmAppBindingRagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MmAppBindingRagResponse
func (client *Client) MmAppBindingRagWithOptions(tmpReq *MmAppBindingRagRequest, runtime *dara.RuntimeOptions) (_result *MmAppBindingRagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MmAppBindingRagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.KnowledgeBaseCodeList) {
		request.KnowledgeBaseCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KnowledgeBaseCodeList, dara.String("KnowledgeBaseCodeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.KnowledgeBaseCodeListShrink) {
		query["KnowledgeBaseCodeList"] = request.KnowledgeBaseCodeListShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MmAppBindingRag"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MmAppBindingRagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态应用绑定知识库
//
// @param request - MmAppBindingRagRequest
//
// @return MmAppBindingRagResponse
func (client *Client) MmAppBindingRag(request *MmAppBindingRagRequest) (_result *MmAppBindingRagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MmAppBindingRagResponse{}
	_body, _err := client.MmAppBindingRagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更用户记忆配置
//
// @param request - PatchMemoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchMemoryConfigResponse
func (client *Client) PatchMemoryConfigWithOptions(request *PatchMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *PatchMemoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoUpdate) {
		query["AutoUpdate"] = request.AutoUpdate
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchMemoryConfig"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PatchMemoryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更用户记忆配置
//
// @param request - PatchMemoryConfigRequest
//
// @return PatchMemoryConfigResponse
func (client *Client) PatchMemoryConfig(request *PatchMemoryConfigRequest) (_result *PatchMemoryConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PatchMemoryConfigResponse{}
	_body, _err := client.PatchMemoryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态应用发布
//
// @param request - PublishMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishMmAppResponse
func (client *Client) PublishMmAppWithOptions(request *PublishMmAppRequest, runtime *dara.RuntimeOptions) (_result *PublishMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态应用发布
//
// @param request - PublishMmAppRequest
//
// @return PublishMmAppResponse
func (client *Client) PublishMmApp(request *PublishMmAppRequest) (_result *PublishMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishMmAppResponse{}
	_body, _err := client.PublishMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户记忆配置
//
// @param request - QueryMemoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemoryConfigResponse
func (client *Client) QueryMemoryConfigWithOptions(request *QueryMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryMemoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMemoryConfig"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMemoryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户记忆配置
//
// @param request - QueryMemoryConfigRequest
//
// @return QueryMemoryConfigResponse
func (client *Client) QueryMemoryConfig(request *QueryMemoryConfigRequest) (_result *QueryMemoryConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMemoryConfigResponse{}
	_body, _err := client.QueryMemoryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户记忆列表
//
// @param request - QueryMemoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemoryListResponse
func (client *Client) QueryMemoryListWithOptions(request *QueryMemoryListRequest, runtime *dara.RuntimeOptions) (_result *QueryMemoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMemoryList"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMemoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户记忆列表
//
// @param request - QueryMemoryListRequest
//
// @return QueryMemoryListResponse
func (client *Client) QueryMemoryList(request *QueryMemoryListRequest) (_result *QueryMemoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMemoryListResponse{}
	_body, _err := client.QueryMemoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户画像配置
//
// @param request - QueryProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProfileResponse
func (client *Client) QueryProfileWithOptions(request *QueryProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProfile"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户画像配置
//
// @param request - QueryProfileRequest
//
// @return QueryProfileResponse
func (client *Client) QueryProfile(request *QueryProfileRequest) (_result *QueryProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryProfileResponse{}
	_body, _err := client.QueryProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户画像
//
// @param request - QueryUserProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserProfileResponse
func (client *Client) QueryUserProfileWithOptions(request *QueryUserProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryUserProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserProfile"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户画像
//
// @param request - QueryUserProfileRequest
//
// @return QueryUserProfileResponse
func (client *Client) QueryUserProfile(request *QueryUserProfileRequest) (_result *QueryUserProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserProfileResponse{}
	_body, _err := client.QueryUserProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指令更新
//
// @param tmpReq - UpdateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommandResponse
func (client *Client) UpdateCommandWithOptions(tmpReq *UpdateCommandRequest, runtime *dara.RuntimeOptions) (_result *UpdateCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ToolExamples) {
		request.ToolExamplesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToolExamples, dara.String("ToolExamples"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ToolParams) {
		request.ToolParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToolParams, dara.String("ToolParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ToolDescription) {
		query["ToolDescription"] = request.ToolDescription
	}

	if !dara.IsNil(request.ToolExamplesShrink) {
		query["ToolExamples"] = request.ToolExamplesShrink
	}

	if !dara.IsNil(request.ToolId) {
		query["ToolId"] = request.ToolId
	}

	if !dara.IsNil(request.ToolName) {
		query["ToolName"] = request.ToolName
	}

	if !dara.IsNil(request.ToolParamsShrink) {
		query["ToolParams"] = request.ToolParamsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCommand"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指令更新
//
// @param request - UpdateCommandRequest
//
// @return UpdateCommandResponse
func (client *Client) UpdateCommand(request *UpdateCommandRequest) (_result *UpdateCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCommandResponse{}
	_body, _err := client.UpdateCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户记忆
//
// @param tmpReq - UpdateMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(tmpReq *UpdateMemoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMemoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetaData) {
		request.MetaDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetaData, dara.String("MetaData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.MemoryNodeId) {
		query["MemoryNodeId"] = request.MemoryNodeId
	}

	if !dara.IsNil(request.MetaDataShrink) {
		query["MetaData"] = request.MetaDataShrink
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户记忆
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态应用更新
//
// @param tmpReq - UpdateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppResponse
func (client *Client) UpdateMmAppWithOptions(tmpReq *UpdateMmAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMmAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BindingConfig) {
		request.BindingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BindingConfig, dara.String("BindingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ConversationConfig) {
		request.ConversationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConversationConfig, dara.String("ConversationConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ModelConfig) {
		request.ModelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelConfig, dara.String("ModelConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BindingConfigShrink) {
		query["BindingConfig"] = request.BindingConfigShrink
	}

	if !dara.IsNil(request.ConversationConfigShrink) {
		query["ConversationConfig"] = request.ConversationConfigShrink
	}

	if !dara.IsNil(request.ModelConfigShrink) {
		query["ModelConfig"] = request.ModelConfigShrink
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmApp"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态应用更新
//
// @param request - UpdateMmAppRequest
//
// @return UpdateMmAppResponse
func (client *Client) UpdateMmApp(request *UpdateMmAppRequest) (_result *UpdateMmAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppResponse{}
	_body, _err := client.UpdateMmAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改多模态应用长期记忆开关
//
// @param request - UpdateMmAppMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppMemoryResponse
func (client *Client) UpdateMmAppMemoryWithOptions(request *UpdateMmAppMemoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmAppMemory"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改多模态应用长期记忆开关
//
// @param request - UpdateMmAppMemoryRequest
//
// @return UpdateMmAppMemoryResponse
func (client *Client) UpdateMmAppMemory(request *UpdateMmAppMemoryRequest) (_result *UpdateMmAppMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppMemoryResponse{}
	_body, _err := client.UpdateMmAppMemoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改知识库开关
//
// @param request - UpdateMmAppRagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagResponse
func (client *Client) UpdateMmAppRagWithOptions(request *UpdateMmAppRagRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmAppRag"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppRagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改知识库开关
//
// @param request - UpdateMmAppRagRequest
//
// @return UpdateMmAppRagResponse
func (client *Client) UpdateMmAppRag(request *UpdateMmAppRagRequest) (_result *UpdateMmAppRagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppRagResponse{}
	_body, _err := client.UpdateMmAppRagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改知识库配置
//
// @param request - UpdateMmAppRagConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagConfigResponse
func (client *Client) UpdateMmAppRagConfigWithOptions(request *UpdateMmAppRagConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PromptStrategy) {
		query["PromptStrategy"] = request.PromptStrategy
	}

	if !dara.IsNil(request.RetrieveMaxLength) {
		query["RetrieveMaxLength"] = request.RetrieveMaxLength
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmAppRagConfig"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppRagConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改知识库配置
//
// @param request - UpdateMmAppRagConfigRequest
//
// @return UpdateMmAppRagConfigResponse
func (client *Client) UpdateMmAppRagConfig(request *UpdateMmAppRagConfigRequest) (_result *UpdateMmAppRagConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppRagConfigResponse{}
	_body, _err := client.UpdateMmAppRagConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改知识库权重
//
// @param tmpReq - UpdateMmAppRagWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagWeightResponse
func (client *Client) UpdateMmAppRagWeightWithOptions(tmpReq *UpdateMmAppRagWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagWeightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMmAppRagWeightShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RankWeights) {
		request.RankWeightsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RankWeights, dara.String("RankWeights"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.RankWeightsShrink) {
		query["RankWeights"] = request.RankWeightsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmAppRagWeight"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppRagWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改知识库权重
//
// @param request - UpdateMmAppRagWeightRequest
//
// @return UpdateMmAppRagWeightResponse
func (client *Client) UpdateMmAppRagWeight(request *UpdateMmAppRagWeightRequest) (_result *UpdateMmAppRagWeightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppRagWeightResponse{}
	_body, _err := client.UpdateMmAppRagWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用承接语开关
//
// @param request - UpdateMmAppTransitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppTransitionResponse
func (client *Client) UpdateMmAppTransitionWithOptions(request *UpdateMmAppTransitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppTransitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmAppTransition"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmAppTransitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用承接语开关
//
// @param request - UpdateMmAppTransitionRequest
//
// @return UpdateMmAppTransitionResponse
func (client *Client) UpdateMmAppTransition(request *UpdateMmAppTransitionRequest) (_result *UpdateMmAppTransitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMmAppTransitionResponse{}
	_body, _err := client.UpdateMmAppTransitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更用户画像配置
//
// @param tmpReq - UpdateProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProfileResponse
func (client *Client) UpdateProfileWithOptions(tmpReq *UpdateProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttributesOperations) {
		request.AttributesOperationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributesOperations, dara.String("AttributesOperations"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AttributesOperationsShrink) {
		query["AttributesOperations"] = request.AttributesOperationsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.UserDefinedId) {
		query["UserDefinedId"] = request.UserDefinedId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProfile"),
		Version:     dara.String("2025-09-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更用户画像配置
//
// @param request - UpdateProfileRequest
//
// @return UpdateProfileResponse
func (client *Client) UpdateProfile(request *UpdateProfileRequest) (_result *UpdateProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProfileResponse{}
	_body, _err := client.UpdateProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
