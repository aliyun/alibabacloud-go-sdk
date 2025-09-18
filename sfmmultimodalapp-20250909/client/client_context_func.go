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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
