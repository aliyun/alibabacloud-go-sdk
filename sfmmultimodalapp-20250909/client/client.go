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
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建多模态应用
//
// @param tmpReq - CreateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmAppResponse
func (client *Client) CreateMmAppWithOptions(tmpReq *CreateMmAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMmAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除指令
//
// @param request - DeleteCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommandWithOptions(request *DeleteCommandRequest, runtime *dara.RuntimeOptions) (_result *DeleteCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除多模态应用
//
// @param request - DeleteMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmAppResponse
func (client *Client) DeleteMmAppWithOptions(request *DeleteMmAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMmAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 指令详情
//
// @param request - DescribeCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommandResponse
func (client *Client) DescribeCommandWithOptions(request *DescribeCommandRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 多模态应用发布
//
// @param request - PublishMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishMmAppResponse
func (client *Client) PublishMmAppWithOptions(request *PublishMmAppRequest, runtime *dara.RuntimeOptions) (_result *PublishMmAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 指令更新
//
// @param tmpReq - UpdateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommandResponse
func (client *Client) UpdateCommandWithOptions(tmpReq *UpdateCommandRequest, runtime *dara.RuntimeOptions) (_result *UpdateCommandResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 多模态应用更新
//
// @param tmpReq - UpdateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppResponse
func (client *Client) UpdateMmAppWithOptions(tmpReq *UpdateMmAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
