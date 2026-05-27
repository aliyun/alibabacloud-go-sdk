// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建对话
//
// @param request - CreateChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatResponse
func (client *Client) CreateChatWithSSECtx(ctx context.Context, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CreateChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createChatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 创建对话
//
// @param request - CreateChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatResponse
func (client *Client) CreateChatWithContext(ctx context.Context, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.DigitalEmployeeName) {
		body["digitalEmployeeName"] = request.DigitalEmployeeName
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChat"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建DigitalEmployee
//
// @param request - CreateDigitalEmployeeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDigitalEmployeeResponse
func (client *Client) CreateDigitalEmployeeWithContext(ctx context.Context, request *CreateDigitalEmployeeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDigitalEmployeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.DefaultRule) {
		body["defaultRule"] = request.DefaultRule
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Knowledges) {
		body["knowledges"] = request.Knowledges
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RoleArn) {
		body["roleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDigitalEmployee"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digital-employee"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDigitalEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建技能
//
// @param request - CreateDigitalEmployeeSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDigitalEmployeeSkillResponse
func (client *Client) CreateDigitalEmployeeSkillWithContext(ctx context.Context, name *string, request *CreateDigitalEmployeeSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDigitalEmployeeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.SkillName) {
		body["skillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDigitalEmployeeSkill"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDigitalEmployeeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 MCP 服务
//
// @param request - CreateMcpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpServiceResponse
func (client *Client) CreateMcpServiceWithContext(ctx context.Context, name *string, request *CreateMcpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMcpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Connection) {
		body["connection"] = request.Connection
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.McpServiceName) {
		body["mcpServiceName"] = request.McpServiceName
	}

	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpService"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/mcpService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会话
//
// @param request - CreateThreadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateThreadResponse
func (client *Client) CreateThreadWithContext(ctx context.Context, name *string, request *CreateThreadRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateThreadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateThread"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/thread"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateThreadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建票据
//
// @param request - CreateTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, request *CreateTicketRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenExpirationTime) {
		query["accessTokenExpirationTime"] = request.AccessTokenExpirationTime
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["expirationTime"] = request.ExpirationTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tickets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除DigitalEmployee
//
// @param request - DeleteDigitalEmployeeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDigitalEmployeeResponse
func (client *Client) DeleteDigitalEmployeeWithContext(ctx context.Context, name *string, request *DeleteDigitalEmployeeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDigitalEmployeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDigitalEmployee"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digital-employee/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDigitalEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除技能
//
// @param request - DeleteDigitalEmployeeSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDigitalEmployeeSkillResponse
func (client *Client) DeleteDigitalEmployeeSkillWithContext(ctx context.Context, name *string, skillName *string, request *DeleteDigitalEmployeeSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDigitalEmployeeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDigitalEmployeeSkill"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDigitalEmployeeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 MCP 服务
//
// @param request - DeleteMcpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpServiceResponse
func (client *Client) DeleteMcpServiceWithContext(ctx context.Context, name *string, mcpServiceName *string, request *DeleteMcpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMcpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpService"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/mcpService/" + dara.PercentEncode(dara.StringValue(mcpServiceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会话
//
// @param request - DeleteThreadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteThreadResponse
func (client *Client) DeleteThreadWithContext(ctx context.Context, name *string, threadId *string, request *DeleteThreadRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteThreadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteThread"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/thread/" + dara.PercentEncode(dara.StringValue(threadId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteThreadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预览远端 MCP 工具列表
//
// @param request - FetchRemoteMcpToolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchRemoteMcpToolsResponse
func (client *Client) FetchRemoteMcpToolsWithContext(ctx context.Context, request *FetchRemoteMcpToolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FetchRemoteMcpToolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Connection) {
		body["connection"] = request.Connection
	}

	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchRemoteMcpTools"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/mcptools"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchRemoteMcpToolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载小型产物文件
//
// @param request - GetArtifactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactResponse
func (client *Client) GetArtifactWithContext(ctx context.Context, name *string, request *GetArtifactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactPath) {
		query["artifactPath"] = request.ArtifactPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArtifact"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/artifact"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("binary"),
	}
	res := &GetArtifactResponse{}
	callApiTmp, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp := dara.ToMap(callApiTmp)
	if !dara.IsNil(tmp["body"]) {
		respBody := dara.ToReader(tmp["body"])
		res.Body = respBody
	}

	if !dara.IsNil(tmp["headers"]) {
		respHeaders := dara.ToMap(tmp["headers"])
		res.Headers = openapiutil.StringifyMapValue(respHeaders)
	}

	if !dara.IsNil(tmp["statusCode"]) {
		statusCode := dara.ForceInt(tmp["statusCode"])
		res.StatusCode = dara.ToInt32(dara.Int(statusCode))
	}

	_result = res
	return _result, _err
}

// Summary:
//
// 查询 DigitalEmployee
//
// @param request - GetDigitalEmployeeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDigitalEmployeeResponse
func (client *Client) GetDigitalEmployeeWithContext(ctx context.Context, name *string, request *GetDigitalEmployeeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDigitalEmployeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDigitalEmployee"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digital-employee/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDigitalEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取技能详情
//
// @param request - GetDigitalEmployeeSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDigitalEmployeeSkillResponse
func (client *Client) GetDigitalEmployeeSkillWithContext(ctx context.Context, name *string, skillName *string, request *GetDigitalEmployeeSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDigitalEmployeeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDigitalEmployeeSkill"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDigitalEmployeeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 MCP 服务
//
// @param request - GetMcpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpServiceResponse
func (client *Client) GetMcpServiceWithContext(ctx context.Context, name *string, mcpServiceName *string, request *GetMcpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMcpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpService"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/mcpService/" + dara.PercentEncode(dara.StringValue(mcpServiceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话
//
// @param request - GetThreadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetThreadResponse
func (client *Client) GetThreadWithContext(ctx context.Context, name *string, threadId *string, request *GetThreadRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetThreadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetThread"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/thread/" + dara.PercentEncode(dara.StringValue(threadId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetThreadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话数据
//
// @param request - GetThreadDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetThreadDataResponse
func (client *Client) GetThreadDataWithContext(ctx context.Context, name *string, threadId *string, request *GetThreadDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetThreadDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetThreadData"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/thread/" + dara.PercentEncode(dara.StringValue(threadId)) + "/data"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetThreadDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出产物文件
//
// @param request - ListArtifactsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactsResponse
func (client *Client) ListArtifactsWithContext(ctx context.Context, name *string, request *ListArtifactsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactPath) {
		query["artifactPath"] = request.ArtifactPath
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListArtifacts"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/artifacts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出技能版本
//
// @param request - ListDigitalEmployeeSkillVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDigitalEmployeeSkillVersionsResponse
func (client *Client) ListDigitalEmployeeSkillVersionsWithContext(ctx context.Context, name *string, skillName *string, request *ListDigitalEmployeeSkillVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDigitalEmployeeSkillVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDigitalEmployeeSkillVersions"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDigitalEmployeeSkillVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出技能
//
// @param request - ListDigitalEmployeeSkillsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDigitalEmployeeSkillsResponse
func (client *Client) ListDigitalEmployeeSkillsWithContext(ctx context.Context, name *string, request *ListDigitalEmployeeSkillsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDigitalEmployeeSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SkillName) {
		query["skillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDigitalEmployeeSkills"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skills"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDigitalEmployeeSkillsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出资源DigitalEmployee
//
// @param tmpReq - ListDigitalEmployeesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDigitalEmployeesResponse
func (client *Client) ListDigitalEmployeesWithContext(ctx context.Context, tmpReq *ListDigitalEmployeesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDigitalEmployeesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDigitalEmployeesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EmployeeType) {
		query["employeeType"] = request.EmployeeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDigitalEmployees"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digital-employee"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDigitalEmployeesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数字员工下的 MCP 服务列表
//
// @param request - ListMcpServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpServicesResponse
func (client *Client) ListMcpServicesWithContext(ctx context.Context, name *string, request *ListMcpServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMcpServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpServices"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/mcpServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出会话
//
// @param tmpReq - ListThreadsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListThreadsResponse
func (client *Client) ListThreadsWithContext(ctx context.Context, name *string, tmpReq *ListThreadsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListThreadsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListThreadsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.IncludeMission) {
		query["includeMission"] = request.IncludeMission
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.ThreadId) {
		query["threadId"] = request.ThreadId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListThreads"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/threads"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListThreadsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新UpdateDigitalEmployee
//
// @param request - UpdateDigitalEmployeeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDigitalEmployeeResponse
func (client *Client) UpdateDigitalEmployeeWithContext(ctx context.Context, name *string, request *UpdateDigitalEmployeeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDigitalEmployeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.DefaultRule) {
		body["defaultRule"] = request.DefaultRule
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Knowledges) {
		body["knowledges"] = request.Knowledges
	}

	if !dara.IsNil(request.RoleArn) {
		body["roleArn"] = request.RoleArn
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDigitalEmployee"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digital-employee/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDigitalEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新技能
//
// @param request - UpdateDigitalEmployeeSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDigitalEmployeeSkillResponse
func (client *Client) UpdateDigitalEmployeeSkillWithContext(ctx context.Context, name *string, skillName *string, request *UpdateDigitalEmployeeSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDigitalEmployeeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDigitalEmployeeSkill"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDigitalEmployeeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 MCP 服务
//
// @param request - UpdateMcpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpServiceResponse
func (client *Client) UpdateMcpServiceWithContext(ctx context.Context, name *string, mcpServiceName *string, request *UpdateMcpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMcpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Connection) {
		body["connection"] = request.Connection
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpService"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/mcpService/" + dara.PercentEncode(dara.StringValue(mcpServiceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会话
//
// @param request - UpdateThreadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateThreadResponse
func (client *Client) UpdateThreadWithContext(ctx context.Context, name *string, threadId *string, request *UpdateThreadRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateThreadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateThread"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/digitalEmployee/" + dara.PercentEncode(dara.StringValue(name)) + "/thread/" + dara.PercentEncode(dara.StringValue(threadId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateThreadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) createChatWithSSECtx_opYieldFunc(_yield chan *CreateChatResponse, _yieldErr chan error, ctx context.Context, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.DigitalEmployeeName) {
		body["digitalEmployeeName"] = request.DigitalEmployeeName
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChat"),
		Version:     dara.String("2026-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
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
