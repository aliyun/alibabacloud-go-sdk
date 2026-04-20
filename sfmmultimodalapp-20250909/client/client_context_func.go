// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 指令创建
//
// @param tmpReq - CreateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommandResponse
func (client *Client) CreateCommandWithContext(ctx context.Context, tmpReq *CreateCommandRequest, runtime *dara.RuntimeOptions) (_result *CreateCommandResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithContext(ctx context.Context, tmpReq *CreateMemoryRequest, runtime *dara.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmAppResponse
func (client *Client) CreateMmAppWithContext(ctx context.Context, tmpReq *CreateMmAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProfileResponse
func (client *Client) CreateProfileWithContext(ctx context.Context, tmpReq *CreateProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateProfileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommandWithContext(ctx context.Context, request *DeleteCommandRequest, runtime *dara.RuntimeOptions) (_result *DeleteCommandResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithContext(ctx context.Context, request *DeleteMemoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmAppResponse
func (client *Client) DeleteMmAppWithContext(ctx context.Context, request *DeleteMmAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProfileResponse
func (client *Client) DeleteProfileWithContext(ctx context.Context, request *DeleteProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteProfileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommandResponse
func (client *Client) DescribeCommandWithContext(ctx context.Context, request *DescribeCommandRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommandResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMmAppResponse
func (client *Client) DescribeMmAppWithContext(ctx context.Context, request *DescribeMmAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommandResponse
func (client *Client) ListCommandWithContext(ctx context.Context, request *ListCommandRequest, runtime *dara.RuntimeOptions) (_result *ListCommandResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmAppResponse
func (client *Client) ListMmAppWithContext(ctx context.Context, request *ListMmAppRequest, runtime *dara.RuntimeOptions) (_result *ListMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedMmAppResponse
func (client *Client) ListPublishedMmAppWithContext(ctx context.Context, request *ListPublishedMmAppRequest, runtime *dara.RuntimeOptions) (_result *ListPublishedMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - MmAppBindingMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MmAppBindingMcpResponse
func (client *Client) MmAppBindingMcpWithContext(ctx context.Context, tmpReq *MmAppBindingMcpRequest, runtime *dara.RuntimeOptions) (_result *MmAppBindingMcpResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - MmAppBindingRagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MmAppBindingRagResponse
func (client *Client) MmAppBindingRagWithContext(ctx context.Context, tmpReq *MmAppBindingRagRequest, runtime *dara.RuntimeOptions) (_result *MmAppBindingRagResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchMemoryConfigResponse
func (client *Client) PatchMemoryConfigWithContext(ctx context.Context, request *PatchMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *PatchMemoryConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishMmAppResponse
func (client *Client) PublishMmAppWithContext(ctx context.Context, request *PublishMmAppRequest, runtime *dara.RuntimeOptions) (_result *PublishMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemoryConfigResponse
func (client *Client) QueryMemoryConfigWithContext(ctx context.Context, request *QueryMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryMemoryConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemoryListResponse
func (client *Client) QueryMemoryListWithContext(ctx context.Context, request *QueryMemoryListRequest, runtime *dara.RuntimeOptions) (_result *QueryMemoryListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProfileResponse
func (client *Client) QueryProfileWithContext(ctx context.Context, request *QueryProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryProfileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserProfileResponse
func (client *Client) QueryUserProfileWithContext(ctx context.Context, request *QueryUserProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryUserProfileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommandResponse
func (client *Client) UpdateCommandWithContext(ctx context.Context, tmpReq *UpdateCommandRequest, runtime *dara.RuntimeOptions) (_result *UpdateCommandResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithContext(ctx context.Context, tmpReq *UpdateMemoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateMmAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppResponse
func (client *Client) UpdateMmAppWithContext(ctx context.Context, tmpReq *UpdateMmAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppMemoryResponse
func (client *Client) UpdateMmAppMemoryWithContext(ctx context.Context, request *UpdateMmAppMemoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagResponse
func (client *Client) UpdateMmAppRagWithContext(ctx context.Context, request *UpdateMmAppRagRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagConfigResponse
func (client *Client) UpdateMmAppRagConfigWithContext(ctx context.Context, request *UpdateMmAppRagConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateMmAppRagWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppRagWeightResponse
func (client *Client) UpdateMmAppRagWeightWithContext(ctx context.Context, tmpReq *UpdateMmAppRagWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppRagWeightResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmAppTransitionResponse
func (client *Client) UpdateMmAppTransitionWithContext(ctx context.Context, request *UpdateMmAppTransitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateMmAppTransitionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProfileResponse
func (client *Client) UpdateProfileWithContext(ctx context.Context, tmpReq *UpdateProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateProfileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
