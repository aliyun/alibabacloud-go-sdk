// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("AgentCore")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentcore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Writes messages or text to a memory store. If sync is omitted or set to true, the operation synchronously waits for extraction and returns only memories, which can be an empty array. If sync is set to false, the operation returns only a taskId, and you can query the result by calling GetMemoryTask.
//
// Description:
//
// Writes messages or text to a memory store to fetch long-term memory to retrieve and use later.
//
// - If `sync` is omitted or set to `true`: waits for the fetch to complete and returns `memories`. Returns an empty array if no memories are fetched.
//
// - If `sync` is set to `false`: returns a `taskId`. You can query the fetch progress and results by calling [GetMemoryTask](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMemoryTask).
//
// @param tmpReq - AddMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemoriesResponse
func (client *Client) AddMemoriesWithOptions(workspaceId *string, memoryStoreName *string, tmpReq *AddMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddMemoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMemories"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Writes messages or text to a memory store. If sync is omitted or set to true, the operation synchronously waits for extraction and returns only memories, which can be an empty array. If sync is set to false, the operation returns only a taskId, and you can query the result by calling GetMemoryTask.
//
// Description:
//
// Writes messages or text to a memory store to fetch long-term memory to retrieve and use later.
//
// - If `sync` is omitted or set to `true`: waits for the fetch to complete and returns `memories`. Returns an empty array if no memories are fetched.
//
// - If `sync` is set to `false`: returns a `taskId`. You can query the fetch progress and results by calling [GetMemoryTask](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMemoryTask).
//
// @param request - AddMemoriesRequest
//
// @return AddMemoriesResponse
func (client *Client) AddMemories(workspaceId *string, memoryStoreName *string, request *AddMemoriesRequest) (_result *AddMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMemoriesResponse{}
	_body, _err := client.AddMemoriesWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple models in a specified workspace. If any model is in use, the entire batch request fails.
//
// Description:
//
// Deletes multiple models in a specified workspace. If any model is in use, the entire batch request fails.
//
// @param tmpReq - BatchDeleteModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteModelsResponse
func (client *Client) BatchDeleteModelsWithOptions(workspaceId *string, tmpReq *BatchDeleteModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeleteModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteModelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteModels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models/actions/batch-delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple models in a specified workspace. If any model is in use, the entire batch request fails.
//
// Description:
//
// Deletes multiple models in a specified workspace. If any model is in use, the entire batch request fails.
//
// @param request - BatchDeleteModelsRequest
//
// @return BatchDeleteModelsResponse
func (client *Client) BatchDeleteModels(workspaceId *string, request *BatchDeleteModelsRequest) (_result *BatchDeleteModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteModelsResponse{}
	_body, _err := client.BatchDeleteModelsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads multiple skill ZIP packages through OSS and returns the processing result for each skill.
//
// Description:
//
// Imports multiple skill ZIP packages that have been uploaded to OSS into a workspace in batches, and returns the processing result for each skill.
//
// OSS is used here to temporarily store the ZIP packages to be imported. First, call [GetSkillImportFileUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetSkillImportFileUrl) to obtain the upload URL and upload the files, and then call this operation. To check the package content and name conflicts in advance, call [PrecheckSkillUploadViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/PrecheckSkillUploadViaOss).
//
// @param tmpReq - BatchUploadSkillsViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUploadSkillsViaOssResponse
func (client *Client) BatchUploadSkillsViaOssWithOptions(workspaceId *string, tmpReq *BatchUploadSkillsViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchUploadSkillsViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUploadSkillsViaOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUploadSkillsViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skill-actions/batch-upload-via-oss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUploadSkillsViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads multiple skill ZIP packages through OSS and returns the processing result for each skill.
//
// Description:
//
// Imports multiple skill ZIP packages that have been uploaded to OSS into a workspace in batches, and returns the processing result for each skill.
//
// OSS is used here to temporarily store the ZIP packages to be imported. First, call [GetSkillImportFileUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetSkillImportFileUrl) to obtain the upload URL and upload the files, and then call this operation. To check the package content and name conflicts in advance, call [PrecheckSkillUploadViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/PrecheckSkillUploadViaOss).
//
// @param request - BatchUploadSkillsViaOssRequest
//
// @return BatchUploadSkillsViaOssResponse
func (client *Client) BatchUploadSkillsViaOss(workspaceId *string, request *BatchUploadSkillsViaOssRequest) (_result *BatchUploadSkillsViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUploadSkillsViaOssResponse{}
	_body, _err := client.BatchUploadSkillsViaOssWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts an MCP server to free-edit mode.
//
// Description:
//
// Converts an MCP server installed from a marketplace template to free-edit mode so that you can directly modify the server configuration. After the conversion, the original source and tag information is retained, and the template editing restrictions are removed. The server no longer appears on the use case page associated with the template.
//
// @param request - ConvertMcpToFreeEditRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertMcpToFreeEditResponse
func (client *Client) ConvertMcpToFreeEditWithOptions(workspaceId *string, mcpServerId *string, request *ConvertMcpToFreeEditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConvertMcpToFreeEditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertMcpToFreeEdit"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId)) + "/convert-to-free-edit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertMcpToFreeEditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts an MCP server to free-edit mode.
//
// Description:
//
// Converts an MCP server installed from a marketplace template to free-edit mode so that you can directly modify the server configuration. After the conversion, the original source and tag information is retained, and the template editing restrictions are removed. The server no longer appears on the use case page associated with the template.
//
// @param request - ConvertMcpToFreeEditRequest
//
// @return ConvertMcpToFreeEditResponse
func (client *Client) ConvertMcpToFreeEdit(workspaceId *string, mcpServerId *string, request *ConvertMcpToFreeEditRequest) (_result *ConvertMcpToFreeEditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConvertMcpToFreeEditResponse{}
	_body, _err := client.ConvertMcpToFreeEditWithOptions(workspaceId, mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IM channel for a specified agent and binds it to a publicly accessible service endpoint.
//
// Description:
//
// Creates an IM channel for an agent to route messages from instant messaging platforms to the agent.
//
// During creation, a publicly accessible service endpoint (`ServiceEndpoint`) is bound to the channel. After the channel is created, you can call [GetAgentIMChannel](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentIMChannel) to query its configuration and status.
//
// @param tmpReq - CreateAgentIMChannelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentIMChannelResponse
func (client *Client) CreateAgentIMChannelWithOptions(workspaceId *string, agentId *string, tmpReq *CreateAgentIMChannelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentIMChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentIMChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentIMChannel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentIMChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IM channel for a specified agent and binds it to a publicly accessible service endpoint.
//
// Description:
//
// Creates an IM channel for an agent to route messages from instant messaging platforms to the agent.
//
// During creation, a publicly accessible service endpoint (`ServiceEndpoint`) is bound to the channel. After the channel is created, you can call [GetAgentIMChannel](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentIMChannel) to query its configuration and status.
//
// @param request - CreateAgentIMChannelRequest
//
// @return CreateAgentIMChannelResponse
func (client *Client) CreateAgentIMChannel(workspaceId *string, agentId *string, request *CreateAgentIMChannelRequest) (_result *CreateAgentIMChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentIMChannelResponse{}
	_body, _err := client.CreateAgentIMChannelWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AgentSpec in a specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
//
// Description:
//
// Creates an agent template (AgentSpec) and generates the first draft version for editing reusable agent configurations. The creation fails if a template with the same name already exists in the same workspace. After creation, you can query the template and version information by calling [GetAgentSpec](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentSpec).
//
// @param tmpReq - CreateAgentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentSpecResponse
func (client *Client) CreateAgentSpecWithOptions(workspaceId *string, tmpReq *CreateAgentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentSpec"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AgentSpec in a specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
//
// Description:
//
// Creates an agent template (AgentSpec) and generates the first draft version for editing reusable agent configurations. The creation fails if a template with the same name already exists in the same workspace. After creation, you can query the template and version information by calling [GetAgentSpec](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentSpec).
//
// @param request - CreateAgentSpecRequest
//
// @return CreateAgentSpecResponse
func (client *Client) CreateAgentSpec(workspaceId *string, request *CreateAgentSpecRequest) (_result *CreateAgentSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentSpecResponse{}
	_body, _err := client.CreateAgentSpecWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist and must not have a draft currently being edited.
//
// Description:
//
// Creates a new draft version for an existing agent template (AgentSpec) to modify and publish the next version of the configuration.
//
// The template must already exist and must not have a draft currently being edited.
//
// @param tmpReq - CreateAgentSpecVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentSpecVersionResponse
func (client *Client) CreateAgentSpecVersionWithOptions(workspaceId *string, agentSpecName *string, tmpReq *CreateAgentSpecVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentSpecVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentSpecVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentSpecVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentSpecVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist and must not have a draft currently being edited.
//
// Description:
//
// Creates a new draft version for an existing agent template (AgentSpec) to modify and publish the next version of the configuration.
//
// The template must already exist and must not have a draft currently being edited.
//
// @param request - CreateAgentSpecVersionRequest
//
// @return CreateAgentSpecVersionResponse
func (client *Client) CreateAgentSpecVersion(workspaceId *string, agentSpecName *string, request *CreateAgentSpecVersionRequest) (_result *CreateAgentSpecVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentSpecVersionResponse{}
	_body, _err := client.CreateAgentSpecVersionWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Saves an agent version.
//
// Description:
//
// Saves the currently effective configuration of a specified managed or custom agent as an immutable fixed version, and archives the version template and artifacts. This operation only saves the version and does not deploy an instance. Other types of agents, such as external agents, do not support versions and return AgentVersionSubtypeUnsupported. The latest release of the agent must be effective, meaning latestSpecVersion is equal to effectiveSpecVersion. The system generates an incrementing version number, agentVersion, such as v1 and v2, for each agent. The version number is unique within the agent, and the sequence numbers of deleted versions are not reused. This version number is used for querying, deleting, and referencing versions in service endpoints. Each agent can retain up to 50 undeleted versions. If the limit is reached, AgentVersionQuotaExceeded is returned. You can continue to save versions after deleting unused ones. Duplicate submissions with the same versionName and release serial number return the existing version and are not subject to the limit. Submissions with the same versionName but a different release serial number return AgentVersionAlreadyExists.
//
// @param tmpReq - CreateAgentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentVersionResponse
func (client *Client) CreateAgentVersionWithOptions(workspaceId *string, agentId *string, tmpReq *CreateAgentVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves an agent version.
//
// Description:
//
// Saves the currently effective configuration of a specified managed or custom agent as an immutable fixed version, and archives the version template and artifacts. This operation only saves the version and does not deploy an instance. Other types of agents, such as external agents, do not support versions and return AgentVersionSubtypeUnsupported. The latest release of the agent must be effective, meaning latestSpecVersion is equal to effectiveSpecVersion. The system generates an incrementing version number, agentVersion, such as v1 and v2, for each agent. The version number is unique within the agent, and the sequence numbers of deleted versions are not reused. This version number is used for querying, deleting, and referencing versions in service endpoints. Each agent can retain up to 50 undeleted versions. If the limit is reached, AgentVersionQuotaExceeded is returned. You can continue to save versions after deleting unused ones. Duplicate submissions with the same versionName and release serial number return the existing version and are not subject to the limit. Submissions with the same versionName but a different release serial number return AgentVersionAlreadyExists.
//
// @param request - CreateAgentVersionRequest
//
// @return CreateAgentVersionResponse
func (client *Client) CreateAgentVersion(workspaceId *string, agentId *string, request *CreateAgentVersionRequest) (_result *CreateAgentVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentVersionResponse{}
	_body, _err := client.CreateAgentVersionWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an apiKey or mcpHeader credential in a workspace for agents or MCP to access external services. The credential content is submitted as a JSON string, and masked results are returned when queried.
//
// Description:
//
// Creates an apiKey or mcpHeader credential in a workspace for agents or MCP to access external services.
//
// ### Credential types
//
// - `apiKey`: `credentialMetadata` is a JSON string that contains only a non-empty `apiKey`. Do not submit `resourceScope` and `resourceRefs`.
//
// - `mcpHeader`: `credentialMetadata` is a JSON string that contains a non-empty `headers` array. Each item contains a non-empty `name` and `value` to configure MCP request headers.
//
// ### MCP header scope
//
// You must specify `resourceScope` when creating an `mcpHeader` credential:
//
// - `ALL`: Applies to all MCP services. Omit `resourceRefs`.
//
// - `SPECIFIED`: Applies to specified MCP services. `resourceRefs` must be a non-empty array. For each item, set `resourceType=mcpServer` and specify `resourceId`.
//
// ### Subsequent queries
//
// Use [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential) or [ListCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListCredentials) to query credentials. For `apiKey`, a masked value is returned. For `mcpHeader`, only header names are returned, and header values are not returned.
//
// @param tmpReq - CreateCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialResponse
func (client *Client) CreateCredentialWithOptions(workspaceId *string, tmpReq *CreateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCredentialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/credentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an apiKey or mcpHeader credential in a workspace for agents or MCP to access external services. The credential content is submitted as a JSON string, and masked results are returned when queried.
//
// Description:
//
// Creates an apiKey or mcpHeader credential in a workspace for agents or MCP to access external services.
//
// ### Credential types
//
// - `apiKey`: `credentialMetadata` is a JSON string that contains only a non-empty `apiKey`. Do not submit `resourceScope` and `resourceRefs`.
//
// - `mcpHeader`: `credentialMetadata` is a JSON string that contains a non-empty `headers` array. Each item contains a non-empty `name` and `value` to configure MCP request headers.
//
// ### MCP header scope
//
// You must specify `resourceScope` when creating an `mcpHeader` credential:
//
// - `ALL`: Applies to all MCP services. Omit `resourceRefs`.
//
// - `SPECIFIED`: Applies to specified MCP services. `resourceRefs` must be a non-empty array. For each item, set `resourceType=mcpServer` and specify `resourceId`.
//
// ### Subsequent queries
//
// Use [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential) or [ListCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListCredentials) to query credentials. For `apiKey`, a masked value is returned. For `mcpHeader`, only header names are returned, and header values are not returned.
//
// @param request - CreateCredentialRequest
//
// @return CreateCredentialResponse
func (client *Client) CreateCredential(workspaceId *string, request *CreateCredentialRequest) (_result *CreateCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCredentialResponse{}
	_body, _err := client.CreateCredentialWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom code or image agent in a specified workspace, supports code package, ACR image, and custom image deployments, and executes the deployment asynchronously.
//
// Description:
//
// Deploys a user-provided code package or container image as an agent in a workspace. This corresponds to the "Custom Code / Image" creation method in the console.
//
// ### Preparations
//
// Prepare a writable workspace and submit name, roleArn, artifact, and runtime. In runtime, specify both compute and sessionPolicy. roleArn is the RAM role used by the agent to access your cloud resources.
//
// ### Select a deployment method
//
// | Deployment method | Required configurations |
//
// | --- | --- |
//
// | ACR image | artifact.port, and type=acr, instanceId, namespace, repo, and tag in artifact.container. |
//
// | Custom image | artifact.port, artifact.container.type=custom, and the complete customImage. The clientToken query parameter is required. |
//
// | Code package | artifact.port, and sourceUri, languageRuntime, and entrypoint in artifact.code. |
//
// artifact.code and artifact.container are mutually exclusive. The service listening port ranges from 1 to 65535. You can specify a startup command for a container image by using artifact.container.entrypoint. If this parameter is omitted, the default startup method of the image is used.
//
// ### Code package
//
// - ZIP and TAR.GZ formats are supported. The compressed package size ranges from 1 to 500,000,000 bytes. Dependencies must be included in the package. The platform does not install dependencies or perform online builds.
//
// - Invoke [ListAgentCodeRuntimes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListAgentCodeRuntimes) to query the currently active `languageRuntime`.
//
// - For local files, first invoke [GetAgentCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageUploadUrl) and upload the file directly by using the upload method, URL, and request headers returned in the response. Then, commit `sourceUri`.
//
// - `sourceUri` supports `oss://` URLs or HTTPS object URLs of standard OSS public domain names, including signed URLs. Private OSS objects must allow the workspace role to read them. You can first invoke [GetAgentCodePackageOssRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageOssRamAuthorizeUrl) to complete the authorization.
//
// - Read permissions, objects, and sizes are checked upon commit, and a fixed copy is saved. If the transfer failed, the agent configuration is not created. Archive integrity and decompression are checked when the instance starts.
//
// - `entrypoint` is the startup command executed in the `/code` folder, such as `python main.py`. The command does not automatically execute shell pipelines or expand environment variables. For complex startup logic, write a script in the package.
//
// Sample request body for a code package (the runtime environment is subject to the query result):
//
// ```json
//
// {
//
//	"name": "code-agent",
//
//	"roleArn": "acs:ram::1234567890123456:role/AgentRuntimeRole",
//
//	"artifact": {
//
//	  "port": 8000,
//
//	  "code": {
//
//	    "sourceUri": "oss://example-code-bucket/releases/agent.zip",
//
//	    "languageRuntime": "python3.12",
//
//	    "entrypoint": "python main.py"
//
//	  }
//
//	},
//
//	"runtime": {
//
//	  "compute": {
//
//	    "computeClass": "STANDARD"
//
//	  },
//
//	  "sessionPolicy": {
//
//	    "type": "AFFINITY"
//
//	  }
//
//	}
//
// }
//
// ```
//
// ### Custom image
//
// Specify the complete image address in customImage.image. Tags or digests are supported. Do not include protocols or user information. Configure the image pull network separately in customImage.pullNetwork:
//
// | Pull method | Configuration |
//
// | --- | --- |
//
// | Internet | type=public. Do not specify VPC fields. |
//
// | VPC | type=vpc. Specify vpcId, vSwitchId, and securityGroupId. |
//
// If repository authentication is required, specify both customImage.username and customImage.password. For anonymous access, omit both. The password is not returned in the response. The image pull network and the agent runtime network are configured separately.
//
// useImageCache=true applies only to ACR images. For custom images and code packages, set this parameter to false or omit it.
//
// ### Runtime configurations
//
// Omitted fields in runtime.hpa use the following default values:
//
// | Field | Default value | Valid values and description |
//
// | --- | --- | --- |
//
// | enabled | false | Specifies whether to enable elastic scaling. |
//
// | minSandboxCount | 1 | The minimum number of instances. Valid values: 0 to 100. |
//
// | maxSandboxCount | 10 | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | maxConcurrentSessionsPerSandbox | 20 | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | sessionTtlSeconds | 600 | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If sessionPolicy.type=ISOLATED_HEADER_FIELD, you must provide a valid HTTP request header name in headerName. Do not specify this field for other policies.
//
// If network is omitted, the default access policy on the agent side is used. Actual access is still restricted by the workspace network. When submitting access items, specify the corresponding enabled field.
//
// The total number of items in ossMounts and agenticFsMounts cannot exceed 10. The mount target mountPath must be an absolute path in the container, such as /data/train. It cannot conflict with platform reserved directories, be duplicated, or form parent-child directory overrides.
//
// File mounts are used to allow the agent to access persistent data. When using AgenticFS, you can call [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl) to complete the access point authorization first.
//
// The path for an OSS mount must be a non-empty directory prefix in the bucket. The server for an AgenticFS mount is the access point domain name, without a protocol, port, or path. The path must be an existing non-empty relative directory in the access point. You cannot mount the root directory of the access point. If readOnly is omitted, the default value is false.
//
// Environment variable names must be non-empty and unique in the list. Values must be strings, and empty strings are allowed. When configuring healthCheck, you must specify port. Omitting the health check means no health check is configured. When configuring logging.sls, specify both non-empty project and logstore.
//
// ### Query results and retries
//
// A successful response indicates that the request is accepted. Save the agentId and call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent). Check latestVersionStatus, latestExecutionError, status, and sandboxPhaseCounts to confirm the configuration processing and instance status.
//
// The clientToken query parameter is required for custom image submissions, and optional for ACR image and code package submissions. The token must be 1 to 128 characters in length and can contain letters, digits, periods, underscores, colons, and hyphens. Reuse the original token to retry the same submission. If the parameters are the same, the existing acceptance result is returned. Use a new token for different submission content.
//
// @param tmpReq - CreateCustomAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomAgentResponse
func (client *Client) CreateCustomAgentWithOptions(workspaceId *string, tmpReq *CreateCustomAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom code or image agent in a specified workspace, supports code package, ACR image, and custom image deployments, and executes the deployment asynchronously.
//
// Description:
//
// Deploys a user-provided code package or container image as an agent in a workspace. This corresponds to the "Custom Code / Image" creation method in the console.
//
// ### Preparations
//
// Prepare a writable workspace and submit name, roleArn, artifact, and runtime. In runtime, specify both compute and sessionPolicy. roleArn is the RAM role used by the agent to access your cloud resources.
//
// ### Select a deployment method
//
// | Deployment method | Required configurations |
//
// | --- | --- |
//
// | ACR image | artifact.port, and type=acr, instanceId, namespace, repo, and tag in artifact.container. |
//
// | Custom image | artifact.port, artifact.container.type=custom, and the complete customImage. The clientToken query parameter is required. |
//
// | Code package | artifact.port, and sourceUri, languageRuntime, and entrypoint in artifact.code. |
//
// artifact.code and artifact.container are mutually exclusive. The service listening port ranges from 1 to 65535. You can specify a startup command for a container image by using artifact.container.entrypoint. If this parameter is omitted, the default startup method of the image is used.
//
// ### Code package
//
// - ZIP and TAR.GZ formats are supported. The compressed package size ranges from 1 to 500,000,000 bytes. Dependencies must be included in the package. The platform does not install dependencies or perform online builds.
//
// - Invoke [ListAgentCodeRuntimes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListAgentCodeRuntimes) to query the currently active `languageRuntime`.
//
// - For local files, first invoke [GetAgentCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageUploadUrl) and upload the file directly by using the upload method, URL, and request headers returned in the response. Then, commit `sourceUri`.
//
// - `sourceUri` supports `oss://` URLs or HTTPS object URLs of standard OSS public domain names, including signed URLs. Private OSS objects must allow the workspace role to read them. You can first invoke [GetAgentCodePackageOssRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageOssRamAuthorizeUrl) to complete the authorization.
//
// - Read permissions, objects, and sizes are checked upon commit, and a fixed copy is saved. If the transfer failed, the agent configuration is not created. Archive integrity and decompression are checked when the instance starts.
//
// - `entrypoint` is the startup command executed in the `/code` folder, such as `python main.py`. The command does not automatically execute shell pipelines or expand environment variables. For complex startup logic, write a script in the package.
//
// Sample request body for a code package (the runtime environment is subject to the query result):
//
// ```json
//
// {
//
//	"name": "code-agent",
//
//	"roleArn": "acs:ram::1234567890123456:role/AgentRuntimeRole",
//
//	"artifact": {
//
//	  "port": 8000,
//
//	  "code": {
//
//	    "sourceUri": "oss://example-code-bucket/releases/agent.zip",
//
//	    "languageRuntime": "python3.12",
//
//	    "entrypoint": "python main.py"
//
//	  }
//
//	},
//
//	"runtime": {
//
//	  "compute": {
//
//	    "computeClass": "STANDARD"
//
//	  },
//
//	  "sessionPolicy": {
//
//	    "type": "AFFINITY"
//
//	  }
//
//	}
//
// }
//
// ```
//
// ### Custom image
//
// Specify the complete image address in customImage.image. Tags or digests are supported. Do not include protocols or user information. Configure the image pull network separately in customImage.pullNetwork:
//
// | Pull method | Configuration |
//
// | --- | --- |
//
// | Internet | type=public. Do not specify VPC fields. |
//
// | VPC | type=vpc. Specify vpcId, vSwitchId, and securityGroupId. |
//
// If repository authentication is required, specify both customImage.username and customImage.password. For anonymous access, omit both. The password is not returned in the response. The image pull network and the agent runtime network are configured separately.
//
// useImageCache=true applies only to ACR images. For custom images and code packages, set this parameter to false or omit it.
//
// ### Runtime configurations
//
// Omitted fields in runtime.hpa use the following default values:
//
// | Field | Default value | Valid values and description |
//
// | --- | --- | --- |
//
// | enabled | false | Specifies whether to enable elastic scaling. |
//
// | minSandboxCount | 1 | The minimum number of instances. Valid values: 0 to 100. |
//
// | maxSandboxCount | 10 | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | maxConcurrentSessionsPerSandbox | 20 | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | sessionTtlSeconds | 600 | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If sessionPolicy.type=ISOLATED_HEADER_FIELD, you must provide a valid HTTP request header name in headerName. Do not specify this field for other policies.
//
// If network is omitted, the default access policy on the agent side is used. Actual access is still restricted by the workspace network. When submitting access items, specify the corresponding enabled field.
//
// The total number of items in ossMounts and agenticFsMounts cannot exceed 10. The mount target mountPath must be an absolute path in the container, such as /data/train. It cannot conflict with platform reserved directories, be duplicated, or form parent-child directory overrides.
//
// File mounts are used to allow the agent to access persistent data. When using AgenticFS, you can call [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl) to complete the access point authorization first.
//
// The path for an OSS mount must be a non-empty directory prefix in the bucket. The server for an AgenticFS mount is the access point domain name, without a protocol, port, or path. The path must be an existing non-empty relative directory in the access point. You cannot mount the root directory of the access point. If readOnly is omitted, the default value is false.
//
// Environment variable names must be non-empty and unique in the list. Values must be strings, and empty strings are allowed. When configuring healthCheck, you must specify port. Omitting the health check means no health check is configured. When configuring logging.sls, specify both non-empty project and logstore.
//
// ### Query results and retries
//
// A successful response indicates that the request is accepted. Save the agentId and call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent). Check latestVersionStatus, latestExecutionError, status, and sandboxPhaseCounts to confirm the configuration processing and instance status.
//
// The clientToken query parameter is required for custom image submissions, and optional for ACR image and code package submissions. The token must be 1 to 128 characters in length and can contain letters, digits, periods, underscores, colons, and hyphens. Reuse the original token to retry the same submission. If the parameters are the same, the existing acceptance result is returned. Use a new token for different submission content.
//
// @param request - CreateCustomAgentRequest
//
// @return CreateCustomAgentResponse
func (client *Client) CreateCustomAgent(workspaceId *string, request *CreateCustomAgentRequest) (_result *CreateCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCustomAgentResponse{}
	_body, _err := client.CreateCustomAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an agent that runs in an external environment, saves the configurations for the platform model, instructions, tools, and skills, and asynchronously processes the bootstrap configuration.
//
// Description:
//
// Registers an agent hosted in the user\\"s runtime environment, which corresponds to the "Managed Agent" in the console. After registration, you can continue to obtain the bootstrap information and complete the connection.
//
// ### Configure a managed agent
//
// Provide name, and select the configuration based on the model source:
//
// | modelSource | Model configuration |
//
// | --- | --- |
//
// | PLATFORM | Uses a platform model. Provide model.modelConnectionId and model.modelName. |
//
// | RUNTIME | The model is managed by the external runtime environment. Do not submit model. |
//
// If modelSource is omitted but model is provided, it is processed as PLATFORM. If both are omitted, the agent cannot be created.
//
// Configure instructions, MCP tools, and skills through instruction, tools, and skills. You can also reference an agent template through template.aiRegistry. The external process is deployed by the caller. This operation does not create a platform runtime instance.
//
// ### Bootstrap sequence
//
// **Step 1:*	- After creation, save the agentId and call [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent) to query the configuration processing result.
//
// **Step 2:*	- After the latest configuration takes effect, call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) to obtain the available bootstrap networks.
//
// **Step 3:*	- Select an available network and call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain the bootstrap token and CMS configuration. The external process then completes the bootstrap.
//
// clientToken is a reserved query parameter and currently does not provide idempotency guarantees.
//
// @param tmpReq - CreateExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalAgentResponse
func (client *Client) CreateExternalAgentWithOptions(workspaceId *string, tmpReq *CreateExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExternalAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateExternalAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExternalAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExternalAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an agent that runs in an external environment, saves the configurations for the platform model, instructions, tools, and skills, and asynchronously processes the bootstrap configuration.
//
// Description:
//
// Registers an agent hosted in the user\\"s runtime environment, which corresponds to the "Managed Agent" in the console. After registration, you can continue to obtain the bootstrap information and complete the connection.
//
// ### Configure a managed agent
//
// Provide name, and select the configuration based on the model source:
//
// | modelSource | Model configuration |
//
// | --- | --- |
//
// | PLATFORM | Uses a platform model. Provide model.modelConnectionId and model.modelName. |
//
// | RUNTIME | The model is managed by the external runtime environment. Do not submit model. |
//
// If modelSource is omitted but model is provided, it is processed as PLATFORM. If both are omitted, the agent cannot be created.
//
// Configure instructions, MCP tools, and skills through instruction, tools, and skills. You can also reference an agent template through template.aiRegistry. The external process is deployed by the caller. This operation does not create a platform runtime instance.
//
// ### Bootstrap sequence
//
// **Step 1:*	- After creation, save the agentId and call [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent) to query the configuration processing result.
//
// **Step 2:*	- After the latest configuration takes effect, call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) to obtain the available bootstrap networks.
//
// **Step 3:*	- Select an available network and call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain the bootstrap token and CMS configuration. The external process then completes the bootstrap.
//
// clientToken is a reserved query parameter and currently does not provide idempotency guarantees.
//
// @param request - CreateExternalAgentRequest
//
// @return CreateExternalAgentResponse
func (client *Client) CreateExternalAgent(workspaceId *string, request *CreateExternalAgentRequest) (_result *CreateExternalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExternalAgentResponse{}
	_body, _err := client.CreateExternalAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues an access token for a specified network to a managed agent with effective configurations and returns the CMS configuration.
//
// Description:
//
// Issues an access token and connection configurations for a registered managed agent whose configurations have taken effect to complete the onboarding in the runtime environment where the agent resides.
//
// ### Prerequisites and parameters
//
// The managed agent must be in the running state, and its latest configurations must have taken effect. Call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) first to query available networks.
//
// The `networkType` parameter is passed as a query parameter. Valid values: `INTRANET` (internal network) and `INTERNET` (Internet). The selected endpoint must be available.
//
// ### Use the response
//
// Pass the `bootstrapToken` to the external runtime process to complete the onboarding. The `tokenFingerprint` is used to identify the token and cannot be used as a substitute for the token. The `cms` field contains the observability reporting configurations for the corresponding network. Configure the endpoint, workspace, and license key as provided in the response.
//
// @param request - CreateExternalAgentBootstrapTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalAgentBootstrapTokenResponse
func (client *Client) CreateExternalAgentBootstrapTokenWithOptions(workspaceId *string, agentId *string, request *CreateExternalAgentBootstrapTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExternalAgentBootstrapTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkType) {
		query["networkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExternalAgentBootstrapToken"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/bootstrap/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExternalAgentBootstrapTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues an access token for a specified network to a managed agent with effective configurations and returns the CMS configuration.
//
// Description:
//
// Issues an access token and connection configurations for a registered managed agent whose configurations have taken effect to complete the onboarding in the runtime environment where the agent resides.
//
// ### Prerequisites and parameters
//
// The managed agent must be in the running state, and its latest configurations must have taken effect. Call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) first to query available networks.
//
// The `networkType` parameter is passed as a query parameter. Valid values: `INTRANET` (internal network) and `INTERNET` (Internet). The selected endpoint must be available.
//
// ### Use the response
//
// Pass the `bootstrapToken` to the external runtime process to complete the onboarding. The `tokenFingerprint` is used to identify the token and cannot be used as a substitute for the token. The `cms` field contains the observability reporting configurations for the corresponding network. Configure the endpoint, workspace, and license key as provided in the response.
//
// @param request - CreateExternalAgentBootstrapTokenRequest
//
// @return CreateExternalAgentBootstrapTokenResponse
func (client *Client) CreateExternalAgentBootstrapToken(workspaceId *string, agentId *string, request *CreateExternalAgentBootstrapTokenRequest) (_result *CreateExternalAgentBootstrapTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExternalAgentBootstrapTokenResponse{}
	_body, _err := client.CreateExternalAgentBootstrapTokenWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds an external identity provider such as DingTalk, Lark, or a custom identity source to a workspace for enterprise account logon and user synchronization.
//
// Description:
//
// Integrates an enterprise identity source with a workspace to allow application users to log on with enterprise accounts, or to synchronize users from the enterprise address book to the workspace.
//
// DingTalk, Lark, and custom identity sources (CustomSSO) are supported. A workspace can be bound to only one identity source. To change the integration method, you must unbind the current identity source first.
//
// ### Select an integration method
//
// | Integration method | Scenario | Task completed by this operation |
//
// | --- | --- | --- |
//
// | DingTalk or Lark | The enterprise already uses DingTalk or Lark and wants to integrate enterprise accounts and the address book. | Saves the enterprise application credentials and establishes a binding for subsequent logon and address book synchronization configuration. |
//
// | Custom identity source (CustomSSO) | The enterprise wants to integrate its own identity authentication and user management system. | Creates a shared user pool to prepare for subsequent SAML single sign-on (SSO) and SCIM user synchronization configuration. |
//
// A user pool stores application users and identity integration configurations. SAML handles enterprise account logon. SCIM provisions users from the enterprise identity system to the user pool. You can configure these two capabilities separately.
//
// ### Bind DingTalk or Lark
//
// Specify the enterprise application credentials in metadata. DingTalk requires appKey, appSecret, and corpId. Lark requires appId and appSecret.
//
// Logon is enabled by default. A user synchronization task is executed by default upon the first binding. If you want to select departments before synchronization, set syncOnBind to false.
//
// Subsequent address book change synchronization is controlled by syncEnabled, which is disabled by default. When enabled, you must also provide encryptKey and verificationToken.
//
// ### Bind a custom identity source
//
// When you create a CustomSSO binding, you only need to specify the identity source type:
//
// ```http
//
// POST /workspaces/{workspaceId}/identity-providers
//
// {"identityProviderType":"CustomSSO"}
//
// ```
//
// This call only creates a user pool. Logon and user synchronization are disabled by default. Do not specify metadata, loginEnabled, or syncEnabled during creation.
//
// After creation, call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to query the progress. If WAITING_CONFIGURATION is returned, the user pool is created and you can proceed with the integration configuration:
//
// - **Integrate enterprise account logon**: Call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the SAML integration information of AgentCore and configure it in the enterprise identity system. Then, call [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) to save the SAML configuration of the enterprise identity system and enable logon as needed.
//
// - **Integrate user synchronization**: Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to obtain the SCIM receiving endpoint and credentials, and configure them in the enterprise identity system. After the external configuration is complete, call [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM) to enable AgentCore user synchronization.
//
// The two processes are independent. You can select one or both based on your business requirements.
//
// @param tmpReq - CreateIdentityProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProviderWithOptions(workspaceId *string, tmpReq *CreateIdentityProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentityProvider"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an external identity provider such as DingTalk, Lark, or a custom identity source to a workspace for enterprise account logon and user synchronization.
//
// Description:
//
// Integrates an enterprise identity source with a workspace to allow application users to log on with enterprise accounts, or to synchronize users from the enterprise address book to the workspace.
//
// DingTalk, Lark, and custom identity sources (CustomSSO) are supported. A workspace can be bound to only one identity source. To change the integration method, you must unbind the current identity source first.
//
// ### Select an integration method
//
// | Integration method | Scenario | Task completed by this operation |
//
// | --- | --- | --- |
//
// | DingTalk or Lark | The enterprise already uses DingTalk or Lark and wants to integrate enterprise accounts and the address book. | Saves the enterprise application credentials and establishes a binding for subsequent logon and address book synchronization configuration. |
//
// | Custom identity source (CustomSSO) | The enterprise wants to integrate its own identity authentication and user management system. | Creates a shared user pool to prepare for subsequent SAML single sign-on (SSO) and SCIM user synchronization configuration. |
//
// A user pool stores application users and identity integration configurations. SAML handles enterprise account logon. SCIM provisions users from the enterprise identity system to the user pool. You can configure these two capabilities separately.
//
// ### Bind DingTalk or Lark
//
// Specify the enterprise application credentials in metadata. DingTalk requires appKey, appSecret, and corpId. Lark requires appId and appSecret.
//
// Logon is enabled by default. A user synchronization task is executed by default upon the first binding. If you want to select departments before synchronization, set syncOnBind to false.
//
// Subsequent address book change synchronization is controlled by syncEnabled, which is disabled by default. When enabled, you must also provide encryptKey and verificationToken.
//
// ### Bind a custom identity source
//
// When you create a CustomSSO binding, you only need to specify the identity source type:
//
// ```http
//
// POST /workspaces/{workspaceId}/identity-providers
//
// {"identityProviderType":"CustomSSO"}
//
// ```
//
// This call only creates a user pool. Logon and user synchronization are disabled by default. Do not specify metadata, loginEnabled, or syncEnabled during creation.
//
// After creation, call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to query the progress. If WAITING_CONFIGURATION is returned, the user pool is created and you can proceed with the integration configuration:
//
// - **Integrate enterprise account logon**: Call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the SAML integration information of AgentCore and configure it in the enterprise identity system. Then, call [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) to save the SAML configuration of the enterprise identity system and enable logon as needed.
//
// - **Integrate user synchronization**: Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to obtain the SCIM receiving endpoint and credentials, and configure them in the enterprise identity system. After the external configuration is complete, call [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM) to enable AgentCore user synchronization.
//
// The two processes are independent. You can select one or both based on your business requirements.
//
// @param request - CreateIdentityProviderRequest
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProvider(workspaceId *string, request *CreateIdentityProviderRequest) (_result *CreateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CreateIdentityProviderWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Explicitly issues a SCIM credential and returns the plaintext token only in the current response. A maximum of two credentials can exist in the same pool. This operation does not automatically revoke old credentials or enable AC user synchronization.
//
// Description:
//
// Creates a SCIM access token to authorize an enterprise identity system to send user provisioning requests to AgentCore. For example, when you plan to replace the token used by the enterprise identity system, you can create a new token first and revoke the old token after the switchover is complete.
//
// SCIM is a user provisioning protocol through which an enterprise identity system sends user information by using a receiving endpoint and an access token.
//
// ### Before you begin
//
// A `CustomSSO` binding is created, and the SCIM receiving endpoint is enabled by calling [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM).
//
// A maximum of two SCIM credentials can be retained in the same user pool. When the limit is reached, query existing credentials by calling [ListIdentityProviderSCIMCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderSCIMCredentials) and revoke the credentials that are no longer in use.
//
// ### What to do next
//
// **Step 1:*	- Save the plaintext of the `token` in the response. The token is returned only upon creation. You can only query the credential summary later.
//
// **Step 2:*	- Configure the receiving endpoint and the new token in the enterprise identity system.
//
// **Step 3:*	- After you confirm that the enterprise identity system uses the new token, revoke the old token by calling [DeleteIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteIdentityProviderSCIMCredential) as needed.
//
// If the creation response is lost, query the credential list to verify the creation result to avoid duplicate creation and occupying the credential quota. The user synchronization switch is configured separately by calling [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM).
//
// @param request - CreateIdentityProviderSCIMCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderSCIMCredentialResponse
func (client *Client) CreateIdentityProviderSCIMCredentialWithOptions(workspaceId *string, request *CreateIdentityProviderSCIMCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderSCIMCredentialResponse, _err error) {
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
		Action:      dara.String("CreateIdentityProviderSCIMCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning/credentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderSCIMCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Explicitly issues a SCIM credential and returns the plaintext token only in the current response. A maximum of two credentials can exist in the same pool. This operation does not automatically revoke old credentials or enable AC user synchronization.
//
// Description:
//
// Creates a SCIM access token to authorize an enterprise identity system to send user provisioning requests to AgentCore. For example, when you plan to replace the token used by the enterprise identity system, you can create a new token first and revoke the old token after the switchover is complete.
//
// SCIM is a user provisioning protocol through which an enterprise identity system sends user information by using a receiving endpoint and an access token.
//
// ### Before you begin
//
// A `CustomSSO` binding is created, and the SCIM receiving endpoint is enabled by calling [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM).
//
// A maximum of two SCIM credentials can be retained in the same user pool. When the limit is reached, query existing credentials by calling [ListIdentityProviderSCIMCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderSCIMCredentials) and revoke the credentials that are no longer in use.
//
// ### What to do next
//
// **Step 1:*	- Save the plaintext of the `token` in the response. The token is returned only upon creation. You can only query the credential summary later.
//
// **Step 2:*	- Configure the receiving endpoint and the new token in the enterprise identity system.
//
// **Step 3:*	- After you confirm that the enterprise identity system uses the new token, revoke the old token by calling [DeleteIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteIdentityProviderSCIMCredential) as needed.
//
// If the creation response is lost, query the credential list to verify the creation result to avoid duplicate creation and occupying the credential quota. The user synchronization switch is configured separately by calling [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM).
//
// @param request - CreateIdentityProviderSCIMCredentialRequest
//
// @return CreateIdentityProviderSCIMCredentialResponse
func (client *Client) CreateIdentityProviderSCIMCredential(workspaceId *string, request *CreateIdentityProviderSCIMCredentialRequest) (_result *CreateIdentityProviderSCIMCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdentityProviderSCIMCredentialResponse{}
	_body, _err := client.CreateIdentityProviderSCIMCredentialWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a manual synchronization for an identity provider.
//
// Description:
//
// Synchronizes users from an enterprise identity provider to a workspace based on the saved department scope. This is suitable for importing users for the first time or manually updating workspace users after adjusting the synchronization scope.
//
// ### Prerequisites
//
// - The identity provider is configured and its status is `READY`. If the previous synchronization failed and the status is `SYNC_FAILED`, you can also re-initiate it.
//
// - The user pool is ready, and no identity provider operations are in progress.
//
// - Manual synchronization uses the saved `syncScope`. You do not need to enable `syncEnabled`, which is used for subsequent address book change synchronization.
//
// ### Steps for synchronizing by department
//
// **Step 1:*	- Query departments by calling [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments) and select the scope to synchronize.
//
// **Step 2:*	- Save the `syncScope` by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider), and then call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to wait for the configuration processing to complete.
//
// **Step 3:*	- Call this operation to initiate the synchronization and save the returned `taskId`.
//
// **Step 4:*	- Query the progress by calling [GetIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSyncJob). After the synchronization is complete, view the workspace user list.
//
// If you want to select departments before the initial synchronization, set `syncOnBind=false` when creating the identity provider, and then perform the preceding steps.
//
// ### Synchronization results and re-execution
//
// Synchronization is subject to the maximum number of users allowed in the workspace. Narrowing the synchronization scope does not actively delete users that have already been synchronized.
//
// Each successful initiation creates a new synchronization task. If a request times out or the response is lost, check the identity provider status by calling [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) before deciding whether to resubmit the request.
//
// @param request - CreateIdentityProviderSyncJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderSyncJobResponse
func (client *Client) CreateIdentityProviderSyncJobWithOptions(workspaceId *string, identityProviderType *string, request *CreateIdentityProviderSyncJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderSyncJobResponse, _err error) {
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
		Action:      dara.String("CreateIdentityProviderSyncJob"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType)) + "/sync-jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a manual synchronization for an identity provider.
//
// Description:
//
// Synchronizes users from an enterprise identity provider to a workspace based on the saved department scope. This is suitable for importing users for the first time or manually updating workspace users after adjusting the synchronization scope.
//
// ### Prerequisites
//
// - The identity provider is configured and its status is `READY`. If the previous synchronization failed and the status is `SYNC_FAILED`, you can also re-initiate it.
//
// - The user pool is ready, and no identity provider operations are in progress.
//
// - Manual synchronization uses the saved `syncScope`. You do not need to enable `syncEnabled`, which is used for subsequent address book change synchronization.
//
// ### Steps for synchronizing by department
//
// **Step 1:*	- Query departments by calling [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments) and select the scope to synchronize.
//
// **Step 2:*	- Save the `syncScope` by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider), and then call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to wait for the configuration processing to complete.
//
// **Step 3:*	- Call this operation to initiate the synchronization and save the returned `taskId`.
//
// **Step 4:*	- Query the progress by calling [GetIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSyncJob). After the synchronization is complete, view the workspace user list.
//
// If you want to select departments before the initial synchronization, set `syncOnBind=false` when creating the identity provider, and then perform the preceding steps.
//
// ### Synchronization results and re-execution
//
// Synchronization is subject to the maximum number of users allowed in the workspace. Narrowing the synchronization scope does not actively delete users that have already been synchronized.
//
// Each successful initiation creates a new synchronization task. If a request times out or the response is lost, check the identity provider status by calling [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) before deciding whether to resubmit the request.
//
// @param request - CreateIdentityProviderSyncJobRequest
//
// @return CreateIdentityProviderSyncJobResponse
func (client *Client) CreateIdentityProviderSyncJob(workspaceId *string, identityProviderType *string, request *CreateIdentityProviderSyncJobRequest) (_result *CreateIdentityProviderSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdentityProviderSyncJobResponse{}
	_body, _err := client.CreateIdentityProviderSyncJobWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or reuses a cache for an ACR Enterprise Edition image in a workspace, and returns the cache status upon acceptance.
//
// Description:
//
// Pre-builds or reuses an ACR Enterprise Edition image cache for custom code or image agents for subsequent deployment and startup.
//
// ### Request parameters
//
// `regionId`, `workspaceId`, and `clientToken` are passed as query parameters. The region must match the entry endpoint and the workspace.
//
// Specify `acrInstanceId`, `namespace`, `repository`, and `tag` in the JSON request body. For `repository`, specify only the repository name without repeating the namespace. The source repository must have pull authorization for the workspace resource role. You can obtain the authorization URL by calling [GetWorkspaceAcrRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAcrRamAuthorizeUrl).
//
// ### Transfer network
//
// Both `accessInternet` and `accessVpc` must be provided, and at least one must be set to `true`. The server validates this based on the workspace network conditions. These parameters control the source reading and transfer for the current operation and do not modify the runtime network of the agent.
//
// ### Acceptance and readiness
//
// A successful response returns the current cache snapshot. `PENDING` indicates that the cache is being prepared, `READY` indicates that the cache is ready, and `FAILED` indicates a failure. For failures, check the `error` field. Save the `imageId` and call [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache) to query subsequent results. If the image digest has not been parsed, `digest` may be empty.
//
// Currently, `clientToken` is only validated to ensure it is not empty. It does not guarantee deduplication by identifier or recovery of the original acceptance result. If the result is not confirmed, query the existing cache first.
//
// @param tmpReq - CreateImageCacheRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageCacheResponse
func (client *Client) CreateImageCacheWithOptions(tmpReq *CreateImageCacheRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateImageCacheShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageCache"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/image-caches"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or reuses a cache for an ACR Enterprise Edition image in a workspace, and returns the cache status upon acceptance.
//
// Description:
//
// Pre-builds or reuses an ACR Enterprise Edition image cache for custom code or image agents for subsequent deployment and startup.
//
// ### Request parameters
//
// `regionId`, `workspaceId`, and `clientToken` are passed as query parameters. The region must match the entry endpoint and the workspace.
//
// Specify `acrInstanceId`, `namespace`, `repository`, and `tag` in the JSON request body. For `repository`, specify only the repository name without repeating the namespace. The source repository must have pull authorization for the workspace resource role. You can obtain the authorization URL by calling [GetWorkspaceAcrRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAcrRamAuthorizeUrl).
//
// ### Transfer network
//
// Both `accessInternet` and `accessVpc` must be provided, and at least one must be set to `true`. The server validates this based on the workspace network conditions. These parameters control the source reading and transfer for the current operation and do not modify the runtime network of the agent.
//
// ### Acceptance and readiness
//
// A successful response returns the current cache snapshot. `PENDING` indicates that the cache is being prepared, `READY` indicates that the cache is ready, and `FAILED` indicates a failure. For failures, check the `error` field. Save the `imageId` and call [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache) to query subsequent results. If the image digest has not been parsed, `digest` may be empty.
//
// Currently, `clientToken` is only validated to ensure it is not empty. It does not guarantee deduplication by identifier or recovery of the original acceptance result. If the result is not confirmed, query the existing cache first.
//
// @param request - CreateImageCacheRequest
//
// @return CreateImageCacheResponse
func (client *Client) CreateImageCache(request *CreateImageCacheRequest) (_result *CreateImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CreateImageCacheWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a managed Harness agent in a specified workspace, configures the harness, model, skills, and runtime resources, and deploys the agent asynchronously.
//
// Description:
//
// Creates an agent by using a platform-managed runtime framework. This corresponds to the "Managed Harness" creation method in the console, where you can configure models, instructions, and skills.
//
// ### Required configurations
//
// Submit `name`, `model`, and `runtime`. The `runtime` parameter must include both `compute` and `sessionPolicy`. The agent is deployed and run by the platform. You do not need to provide a code package or container image.
//
// ### Harness and model
//
// | Harness | Model and connector configurations |
//
// | --- | --- |
//
// | `qwenpaw` | If you omit `harness` during creation, this type is used by default. Set `model.modelConnectionId` to the workspace model connection ID, and set `model.modelName` to the name of the model created under this connection. |
//
// | `qodercli` | The qodercli connector must be enabled in the workspace. Set `model.modelConnectionId` to the model identifier returned by [ListConnectorModels](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectorModels). You can omit `model.modelName`. |
//
// `harness` is an object, such as `{"type":"qwenpaw"}`. For qodercli, set `harness.configuration.connectorServiceAccountKey` to the account key record ID returned by [ListConnectors](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectors), not the original service account key. You can omit this parameter if only one key record exists. You must specify this parameter if multiple records exist. `connectorServiceAccountName` is the account name. Do not specify these two binding fields for qwenpaw.
//
// The harness type cannot be changed after creation.
//
// ### Instructions, templates, and skills
//
// `skills` uses the skill names in the workspace AI registry. If `sourceType=REFERENCE`, versions are referenced by `versionSelector`: use `LABEL` with a label name, or `VERSION` with a specific version number. If the selector is omitted, `LABEL/latest` is used. If `sourceType=STATIC`, static files bundled with the package are used, and no version selector is provided.
//
// `template.aiRegistry` specifies the agent template name and an optional version. If you omit the version during creation, the published version is resolved and fixed. When a template is used, `instruction`, `promptVariables`, and `subAgents` are derived from the template, and the corresponding fields in the current request do not take effect. Explicit skills and template skills are recorded separately. You can query the target and effective versions by using `configuredSkills` and `skills`.
//
// If you do not use a template, you can set instructions by using `instruction`, declare variables by using `promptVariables`, and configure up to 20 sub-agents by using `subAgents`. The name and instruction of a sub-agent cannot be empty. Each sub-agent must have a unique name. Tools currently use `tools[].type=MCP`.
//
// ### Runtime configurations and mounts
//
// Omitted fields in `runtime.hpa` use the following default values:
//
// | Field | Default value | Valid values and description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether to enable elastic scaling. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. Valid values: 0 to 100. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If `sessionPolicy.type=ISOLATED_HEADER_FIELD`, you must provide `headerName`. Do not specify this parameter for other policies.
//
// If you provide `network`, you must specify both `accessInternet.enabled` and `accessVpc.enabled`. Actual access is restricted by the workspace network configurations.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. The mount target `mountPath` uses an absolute path in the container, such as `/data/train`. The path cannot conflict with platform-reserved directories, and mount targets cannot be duplicated or form parent-child directory overwrites.
//
// File mounts are used to allow the agent to access persistent data. When you use AgenticFS, you can first complete the access point authorization by using [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl).
//
// The `path` for an OSS mount must be a non-empty directory prefix in a bucket. The `server` for an AgenticFS mount is the access point domain name, without a protocol, port, or path. The `path` is an existing non-empty relative directory in the access point. You cannot mount the root directory of the access point. If `readOnly` is omitted, the default value is `false`.
//
// `environment.variables` configures environment variables, and `environment.credentialReferences` references workspace credentials. Environment variable names must be non-empty and unique, and values must be strings.
//
// ### Query deployment results
//
// A successful response indicates that the request is accepted. Save the `agentId` and use [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query `latestVersionStatus`, `status`, and `sandboxPhaseCounts` to confirm the configuration processing and instance status. `clientToken` is a reserved query parameter and does not provide idempotency guarantees.
//
// @param tmpReq - CreateManagedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateManagedAgentResponse
func (client *Client) CreateManagedAgentWithOptions(workspaceId *string, tmpReq *CreateManagedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateManagedAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateManagedAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateManagedAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateManagedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a managed Harness agent in a specified workspace, configures the harness, model, skills, and runtime resources, and deploys the agent asynchronously.
//
// Description:
//
// Creates an agent by using a platform-managed runtime framework. This corresponds to the "Managed Harness" creation method in the console, where you can configure models, instructions, and skills.
//
// ### Required configurations
//
// Submit `name`, `model`, and `runtime`. The `runtime` parameter must include both `compute` and `sessionPolicy`. The agent is deployed and run by the platform. You do not need to provide a code package or container image.
//
// ### Harness and model
//
// | Harness | Model and connector configurations |
//
// | --- | --- |
//
// | `qwenpaw` | If you omit `harness` during creation, this type is used by default. Set `model.modelConnectionId` to the workspace model connection ID, and set `model.modelName` to the name of the model created under this connection. |
//
// | `qodercli` | The qodercli connector must be enabled in the workspace. Set `model.modelConnectionId` to the model identifier returned by [ListConnectorModels](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectorModels). You can omit `model.modelName`. |
//
// `harness` is an object, such as `{"type":"qwenpaw"}`. For qodercli, set `harness.configuration.connectorServiceAccountKey` to the account key record ID returned by [ListConnectors](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectors), not the original service account key. You can omit this parameter if only one key record exists. You must specify this parameter if multiple records exist. `connectorServiceAccountName` is the account name. Do not specify these two binding fields for qwenpaw.
//
// The harness type cannot be changed after creation.
//
// ### Instructions, templates, and skills
//
// `skills` uses the skill names in the workspace AI registry. If `sourceType=REFERENCE`, versions are referenced by `versionSelector`: use `LABEL` with a label name, or `VERSION` with a specific version number. If the selector is omitted, `LABEL/latest` is used. If `sourceType=STATIC`, static files bundled with the package are used, and no version selector is provided.
//
// `template.aiRegistry` specifies the agent template name and an optional version. If you omit the version during creation, the published version is resolved and fixed. When a template is used, `instruction`, `promptVariables`, and `subAgents` are derived from the template, and the corresponding fields in the current request do not take effect. Explicit skills and template skills are recorded separately. You can query the target and effective versions by using `configuredSkills` and `skills`.
//
// If you do not use a template, you can set instructions by using `instruction`, declare variables by using `promptVariables`, and configure up to 20 sub-agents by using `subAgents`. The name and instruction of a sub-agent cannot be empty. Each sub-agent must have a unique name. Tools currently use `tools[].type=MCP`.
//
// ### Runtime configurations and mounts
//
// Omitted fields in `runtime.hpa` use the following default values:
//
// | Field | Default value | Valid values and description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether to enable elastic scaling. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. Valid values: 0 to 100. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If `sessionPolicy.type=ISOLATED_HEADER_FIELD`, you must provide `headerName`. Do not specify this parameter for other policies.
//
// If you provide `network`, you must specify both `accessInternet.enabled` and `accessVpc.enabled`. Actual access is restricted by the workspace network configurations.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. The mount target `mountPath` uses an absolute path in the container, such as `/data/train`. The path cannot conflict with platform-reserved directories, and mount targets cannot be duplicated or form parent-child directory overwrites.
//
// File mounts are used to allow the agent to access persistent data. When you use AgenticFS, you can first complete the access point authorization by using [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl).
//
// The `path` for an OSS mount must be a non-empty directory prefix in a bucket. The `server` for an AgenticFS mount is the access point domain name, without a protocol, port, or path. The `path` is an existing non-empty relative directory in the access point. You cannot mount the root directory of the access point. If `readOnly` is omitted, the default value is `false`.
//
// `environment.variables` configures environment variables, and `environment.credentialReferences` references workspace credentials. Environment variable names must be non-empty and unique, and values must be strings.
//
// ### Query deployment results
//
// A successful response indicates that the request is accepted. Save the `agentId` and use [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query `latestVersionStatus`, `status`, and `sandboxPhaseCounts` to confirm the configuration processing and instance status. `clientToken` is a reserved query parameter and does not provide idempotency guarantees.
//
// @param request - CreateManagedAgentRequest
//
// @return CreateManagedAgentResponse
func (client *Client) CreateManagedAgent(workspaceId *string, request *CreateManagedAgentRequest) (_result *CreateManagedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateManagedAgentResponse{}
	_body, _err := client.CreateManagedAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an MCP service in a specified workspace. The creation process is asynchronous. You can check whether the service is ready by using the returned status or a query operation.
//
// Description:
//
// Creates an MCP service in a workspace for agents to use the tools provided by the service.
//
// The creation process is executed asynchronously. After you submit the request, call the [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) operation to check whether the service is ready.
//
// To deploy the service by using a local code package, first call the [GetMcpCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcpCodePackageUploadUrl) operation to obtain the upload URL and upload the file. Then, fill in the returned code package token in the deployment configuration.
//
// The `clientToken` parameter does not provide idempotency guarantees. If the request result is not confirmed, query existing services before you decide whether to recreate the service.
//
// @param tmpReq - CreateMcpRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpResponse
func (client *Client) CreateMcpWithOptions(workspaceId *string, tmpReq *CreateMcpRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMcpShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcp"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an MCP service in a specified workspace. The creation process is asynchronous. You can check whether the service is ready by using the returned status or a query operation.
//
// Description:
//
// Creates an MCP service in a workspace for agents to use the tools provided by the service.
//
// The creation process is executed asynchronously. After you submit the request, call the [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) operation to check whether the service is ready.
//
// To deploy the service by using a local code package, first call the [GetMcpCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcpCodePackageUploadUrl) operation to obtain the upload URL and upload the file. Then, fill in the returned code package token in the deployment configuration.
//
// The `clientToken` parameter does not provide idempotency guarantees. If the request result is not confirmed, query existing services before you decide whether to recreate the service.
//
// @param request - CreateMcpRequest
//
// @return CreateMcpResponse
func (client *Client) CreateMcp(workspaceId *string, request *CreateMcpRequest) (_result *CreateMcpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMcpResponse{}
	_body, _err := client.CreateMcpWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a long-term memory store in a specified workspace.
//
// Description:
//
// Creates a long-term memory store (MemoryStore) in a workspace to organize the long-term memories of an agent and their source sessions.
//
// After creation, you can write messages or text by using [AddMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/AddMemories) and retrieve memories by semantics by using [SearchMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/SearchMemories).
//
// @param tmpReq - CreateMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryStoreResponse
func (client *Client) CreateMemoryStoreWithOptions(workspaceId *string, tmpReq *CreateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMemoryStoreShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemoryStore"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a long-term memory store in a specified workspace.
//
// Description:
//
// Creates a long-term memory store (MemoryStore) in a workspace to organize the long-term memories of an agent and their source sessions.
//
// After creation, you can write messages or text by using [AddMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/AddMemories) and retrieve memories by semantics by using [SearchMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/SearchMemories).
//
// @param request - CreateMemoryStoreRequest
//
// @return CreateMemoryStoreResponse
func (client *Client) CreateMemoryStore(workspaceId *string, request *CreateMemoryStoreRequest) (_result *CreateMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryStoreResponse{}
	_body, _err := client.CreateMemoryStoreWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model configuration under a model connection in a specified workspace for model creation.
//
// Description:
//
// Creates a model configuration under a model connection in a workspace for model creation, allowing an agent to select and invoke the model.
//
// @param tmpReq - CreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithOptions(workspaceId *string, tmpReq *CreateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model configuration under a model connection in a specified workspace for model creation.
//
// Description:
//
// Creates a model configuration under a model connection in a workspace for model creation, allowing an agent to select and invoke the model.
//
// @param request - CreateModelRequest
//
// @return CreateModelResponse
func (client *Client) CreateModel(workspaceId *string, request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials for model creation.
//
// Description:
//
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials for model creation.
//
// @param tmpReq - CreateModelConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelConnectionResponse
func (client *Client) CreateModelConnectionWithOptions(workspaceId *string, tmpReq *CreateModelConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateModelConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelConnection"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/model-connections"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials for model creation.
//
// Description:
//
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials for model creation.
//
// @param request - CreateModelConnectionRequest
//
// @return CreateModelConnectionResponse
func (client *Client) CreateModelConnection(workspaceId *string, request *CreateModelConnectionRequest) (_result *CreateModelConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelConnectionResponse{}
	_body, _err := client.CreateModelConnectionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates permission rules.
//
// Description:
//
// Creates access rules in batches for specified resources in a workspace to define which users or agents can or cannot access the resources.
//
// @param tmpReq - CreatePermissionRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePermissionRulesResponse
func (client *Client) CreatePermissionRulesWithOptions(workspaceId *string, tmpReq *CreatePermissionRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePermissionRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePermissionRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePermissionRules"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/permission-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePermissionRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates permission rules.
//
// Description:
//
// Creates access rules in batches for specified resources in a workspace to define which users or agents can or cannot access the resources.
//
// @param request - CreatePermissionRulesRequest
//
// @return CreatePermissionRulesResponse
func (client *Client) CreatePermissionRules(workspaceId *string, request *CreatePermissionRulesRequest) (_result *CreatePermissionRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePermissionRulesResponse{}
	_body, _err := client.CreatePermissionRulesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a skill in a specified workspace and generates an editable draft version. You can derive the draft from an existing version, or specify a target version and a commit message.
//
// Description:
//
// Creates a skill and its draft in a workspace for editing skill instructions and related resources. You can derive the draft from an existing version, or specify a target version and a commit message.
//
// @param tmpReq - CreateSkillDraftRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillDraftResponse
func (client *Client) CreateSkillDraftWithOptions(workspaceId *string, tmpReq *CreateSkillDraftRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSkillDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSkillDraftShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillDraft"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a skill in a specified workspace and generates an editable draft version. You can derive the draft from an existing version, or specify a target version and a commit message.
//
// Description:
//
// Creates a skill and its draft in a workspace for editing skill instructions and related resources. You can derive the draft from an existing version, or specify a target version and a commit message.
//
// @param request - CreateSkillDraftRequest
//
// @return CreateSkillDraftResponse
func (client *Client) CreateSkillDraft(workspaceId *string, request *CreateSkillDraftRequest) (_result *CreateSkillDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSkillDraftResponse{}
	_body, _err := client.CreateSkillDraftWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a team in a specified workspace and configures user members and agent members at the same time. Exactly one user member must have the ADMIN role, and agent members can only have the LEADER or WORKER role.
//
// Description:
//
// Creates a team in a specified workspace and configures user members and agent members at the same time.
//
// @param tmpReq - CreateTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTeamResponse
func (client *Client) CreateTeamWithOptions(workspaceId *string, tmpReq *CreateTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTeamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTeam"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/teams"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a team in a specified workspace and configures user members and agent members at the same time. Exactly one user member must have the ADMIN role, and agent members can only have the LEADER or WORKER role.
//
// Description:
//
// Creates a team in a specified workspace and configures user members and agent members at the same time.
//
// @param request - CreateTeamRequest
//
// @return CreateTeamResponse
func (client *Client) CreateTeam(workspaceId *string, request *CreateTeamRequest) (_result *CreateTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTeamResponse{}
	_body, _err := client.CreateTeamWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user in a specified workspace. The username must be unique within the workspace, can contain only lowercase letters, digits, and hyphens, and cannot be a reserved name such as manager, admin, or any name starting with worker-. If the password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
//
// Description:
//
// Creates a user in a workspace for member collaboration and application access, which is used when configuring team members and user logons.
//
// @param tmpReq - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(workspaceId *string, tmpReq *CreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user in a specified workspace. The username must be unique within the workspace, can contain only lowercase letters, digits, and hyphens, and cannot be a reserved name such as manager, admin, or any name starting with worker-. If the password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
//
// Description:
//
// Creates a user in a workspace for member collaboration and application access, which is used when configuring team members and user logons.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(workspaceId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AgentCore workspace as a resource scope for organizing agents, models, Model Context Protocol (MCP) services, and collaboration configurations.
//
// Description:
//
// Creates an AgentCore workspace as a resource scope for organizing agents, models, Model Context Protocol (MCP) services, and collaboration configurations.
//
// After the workspace is created, you can use the returned workspace ID to call [GetWorkspace](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspace) to query the details, and then manage agents and associated resources in the workspace. For network configuration fields and their dependencies, see the request parameters table.
//
// @param tmpReq - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(tmpReq *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AgentCore workspace as a resource scope for organizing agents, models, Model Context Protocol (MCP) services, and collaboration configurations.
//
// Description:
//
// Creates an AgentCore workspace as a resource scope for organizing agents, models, Model Context Protocol (MCP) services, and collaboration configurations.
//
// After the workspace is created, you can use the returned workspace ID to call [GetWorkspace](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspace) to query the details, and then manage agents and associated resources in the workspace. For network configuration fields and their dependencies, see the request parameters table.
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes a specified model through a published model connection to verify whether the model invocation pipeline is available.
//
// Description:
//
// Invokes a specified model through a published model connection to verify whether the model invocation pipeline is available.
//
// @param tmpReq - DebugModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugModelResponse
func (client *Client) DebugModelWithOptions(workspaceId *string, modelId *string, tmpReq *DebugModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DebugModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DebugModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugModel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelId)) + "/actions/debug"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes a specified model through a published model connection to verify whether the model invocation pipeline is available.
//
// Description:
//
// Invokes a specified model through a published model connection to verify whether the model invocation pipeline is available.
//
// @param request - DebugModelRequest
//
// @return DebugModelResponse
func (client *Client) DebugModel(workspaceId *string, modelId *string, request *DebugModelRequest) (_result *DebugModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DebugModelResponse{}
	_body, _err := client.DebugModelWithOptions(workspaceId, modelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access application.
//
// Description:
//
// Deletes a specified access application and removes the access configuration of the application from the workspace.
//
// @param request - DeleteAccessApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessApplicationResponse
func (client *Client) DeleteAccessApplicationWithOptions(workspaceId *string, applicationId *string, request *DeleteAccessApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAccessApplicationResponse, _err error) {
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
		Action:      dara.String("DeleteAccessApplication"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/access-applications/" + dara.PercentEncode(dara.StringValue(applicationId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access application.
//
// Description:
//
// Deletes a specified access application and removes the access configuration of the application from the workspace.
//
// @param request - DeleteAccessApplicationRequest
//
// @return DeleteAccessApplicationResponse
func (client *Client) DeleteAccessApplication(workspaceId *string, applicationId *string, request *DeleteAccessApplicationRequest) (_result *DeleteAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAccessApplicationResponse{}
	_body, _err := client.DeleteAccessApplicationWithOptions(workspaceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the IM channel of a specified agent.
//
// Description:
//
// Deletes the IM channel of a specified agent.
//
// @param request - DeleteAgentIMChannelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentIMChannelResponse
func (client *Client) DeleteAgentIMChannelWithOptions(workspaceId *string, agentId *string, imChannelId *string, request *DeleteAgentIMChannelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentIMChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentIMChannel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels/" + dara.PercentEncode(dara.StringValue(imChannelId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentIMChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the IM channel of a specified agent.
//
// Description:
//
// Deletes the IM channel of a specified agent.
//
// @param request - DeleteAgentIMChannelRequest
//
// @return DeleteAgentIMChannelResponse
func (client *Client) DeleteAgentIMChannel(workspaceId *string, agentId *string, imChannelId *string, request *DeleteAgentIMChannelRequest) (_result *DeleteAgentIMChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentIMChannelResponse{}
	_body, _err := client.DeleteAgentIMChannelWithOptions(workspaceId, agentId, imChannelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified AgentSpec and all its versions and metadata. This operation is irreversible.
//
// Description:
//
// Deletes the specified agent template and all its versions and metadata.
//
// This operation is irreversible.
//
// @param request - DeleteAgentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentSpecResponse
func (client *Client) DeleteAgentSpecWithOptions(workspaceId *string, agentSpecName *string, request *DeleteAgentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentSpecResponse, _err error) {
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
		Action:      dara.String("DeleteAgentSpec"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified AgentSpec and all its versions and metadata. This operation is irreversible.
//
// Description:
//
// Deletes the specified agent template and all its versions and metadata.
//
// This operation is irreversible.
//
// @param request - DeleteAgentSpecRequest
//
// @return DeleteAgentSpecResponse
func (client *Client) DeleteAgentSpec(workspaceId *string, agentSpecName *string, request *DeleteAgentSpecRequest) (_result *DeleteAgentSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentSpecResponse{}
	_body, _err := client.DeleteAgentSpecWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the draft version of the specified AgentSpec that is currently being edited and clears the draft version pointer.
//
// Description:
//
// Deletes the draft version of the agent template (AgentSpec) that is currently being edited.
//
// @param request - DeleteAgentSpecVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentSpecVersionResponse
func (client *Client) DeleteAgentSpecVersionWithOptions(workspaceId *string, agentSpecName *string, request *DeleteAgentSpecVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentSpecVersionResponse, _err error) {
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
		Action:      dara.String("DeleteAgentSpecVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/draft"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentSpecVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the draft version of the specified AgentSpec that is currently being edited and clears the draft version pointer.
//
// Description:
//
// Deletes the draft version of the agent template (AgentSpec) that is currently being edited.
//
// @param request - DeleteAgentSpecVersionRequest
//
// @return DeleteAgentSpecVersionResponse
func (client *Client) DeleteAgentSpecVersion(workspaceId *string, agentSpecName *string, request *DeleteAgentSpecVersionRequest) (_result *DeleteAgentSpecVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentSpecVersionResponse{}
	_body, _err := client.DeleteAgentSpecVersionWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a saved fixed version of a specified managed or custom agent.
//
// Description:
//
// Deletes a saved fixed version of a specified managed or custom agent. You cannot delete versions that are still referenced by service endpoints or the system-maintained latest version.
//
// @param request - DeleteAgentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentVersionResponse
func (client *Client) DeleteAgentVersionWithOptions(workspaceId *string, agentId *string, agentVersion *string, request *DeleteAgentVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/versions/" + dara.PercentEncode(dara.StringValue(agentVersion))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a saved fixed version of a specified managed or custom agent.
//
// Description:
//
// Deletes a saved fixed version of a specified managed or custom agent. You cannot delete versions that are still referenced by service endpoints or the system-maintained latest version.
//
// @param request - DeleteAgentVersionRequest
//
// @return DeleteAgentVersionResponse
func (client *Client) DeleteAgentVersion(workspaceId *string, agentId *string, agentVersion *string, request *DeleteAgentVersionRequest) (_result *DeleteAgentVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentVersionResponse{}
	_body, _err := client.DeleteAgentVersionWithOptions(workspaceId, agentId, agentVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a credential in a specified workspace and deletes the ciphertext hosted in the credential service. After deletion, agents bound to the credential can no longer obtain the 
//
// Description:
//
// Deletes an access credential that is no longer used in a workspace.
//
// If the credential is still bound to an MCP service, dissociate it before deletion.
//
// @param request - DeleteCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithOptions(workspaceId *string, credentialId *string, request *DeleteCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/credentials/" + dara.PercentEncode(dara.StringValue(credentialId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a credential in a specified workspace and deletes the ciphertext hosted in the credential service. After deletion, agents bound to the credential can no longer obtain the credential content.
//
// Description:
//
// Deletes an access credential that is no longer used in a workspace.
//
// If the credential is still bound to an MCP service, dissociate it before deletion.
//
// @param request - DeleteCredentialRequest
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredential(workspaceId *string, credentialId *string, request *DeleteCredentialRequest) (_result *DeleteCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCredentialResponse{}
	_body, _err := client.DeleteCredentialWithOptions(workspaceId, credentialId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom code or image agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// Description:
//
// Deletes a custom code or image agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// ### Before you begin
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion result
//
// A successful response indicates that the deletion request is accepted, and the response summary typically returns `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// Call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query the deletion progress. After the deletion is complete, a detail query returns a resource not found error. Runtime resources are cleaned up in the background.
//
// @param request - DeleteCustomAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAgentResponse
func (client *Client) DeleteCustomAgentWithOptions(workspaceId *string, agentId *string, request *DeleteCustomAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCustomAgentResponse, _err error) {
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
		Action:      dara.String("DeleteCustomAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom code or image agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// Description:
//
// Deletes a custom code or image agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// ### Before you begin
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion result
//
// A successful response indicates that the deletion request is accepted, and the response summary typically returns `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// Call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query the deletion progress. After the deletion is complete, a detail query returns a resource not found error. Runtime resources are cleaned up in the background.
//
// @param request - DeleteCustomAgentRequest
//
// @return DeleteCustomAgentResponse
func (client *Client) DeleteCustomAgent(workspaceId *string, agentId *string, request *DeleteCustomAgentRequest) (_result *DeleteCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCustomAgentResponse{}
	_body, _err := client.DeleteCustomAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a managed agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// Description:
//
// Deletes a managed agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// ### Prerequisites
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion results
//
// A successful response indicates that the deletion request has been accepted, and the response summary is typically `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// You can query the deletion progress by calling [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent). After the deletion is complete, a details query returns a resource not found error. The processes and runtime environment of the managed agent are managed by the caller.
//
// @param request - DeleteExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExternalAgentResponse
func (client *Client) DeleteExternalAgentWithOptions(workspaceId *string, agentId *string, request *DeleteExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExternalAgentResponse, _err error) {
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
		Action:      dara.String("DeleteExternalAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExternalAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a managed agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// Description:
//
// Deletes a managed agent in a specified workspace and asynchronously cleans up its associated platform resources.
//
// ### Prerequisites
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion results
//
// A successful response indicates that the deletion request has been accepted, and the response summary is typically `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// You can query the deletion progress by calling [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent). After the deletion is complete, a details query returns a resource not found error. The processes and runtime environment of the managed agent are managed by the caller.
//
// @param request - DeleteExternalAgentRequest
//
// @return DeleteExternalAgentResponse
func (client *Client) DeleteExternalAgent(workspaceId *string, agentId *string, request *DeleteExternalAgentRequest) (_result *DeleteExternalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExternalAgentResponse{}
	_body, _err := client.DeleteExternalAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disconnects the identity source integration of a workspace. For CustomSSO, this operation clears the shared user pool and related SAML and SCIM resources. To disable only logon or synchronization, use the corresponding API operations.
//
// Description:
//
// Disconnects the association between a workspace and an enterprise identity source. This operation is used to terminate identity federation or clear the current attach configuration before you attach another identity source.
//
// Disconnecting clears the AgentCore users and associated configurations imported from the identity source. For `CustomSSO`, it also clears the shared user pool and SAML and SCIM access resources. Applications created by users in the enterprise identity system must still be managed on the enterprise side.
//
// ### Preparations before disconnecting
//
// Remove the reference relationships between the related users and collaboration teams first. If teams still reference these users, the API operation rejects the disconnection request.
//
// If you only want to temporarily disable a specific capability, you can retain the identity source binding. Use [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) to disable logon, or use [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM) to disable SCIM synchronization.
//
// ### Verify the disconnection result
//
// The disconnection is executed in the background. After the request is accepted, `DISCONNECTING` is returned. You can query the progress by calling [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider):
//
// - If the query returns `IdentityProviderNotFound`, the binding has been removed and the disconnection is complete.
//
// - If `DISCONNECT_FAILED` is returned, the cleanup failed. Some resources may have already been cleared. You can initiate the disconnection again.
//
// You can bind a new identity source to the workspace only after the disconnection is complete.
//
// @param request - DeleteIdentityProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProviderWithOptions(workspaceId *string, identityProviderType *string, request *DeleteIdentityProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdentityProvider"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disconnects the identity source integration of a workspace. For CustomSSO, this operation clears the shared user pool and related SAML and SCIM resources. To disable only logon or synchronization, use the corresponding API operations.
//
// Description:
//
// Disconnects the association between a workspace and an enterprise identity source. This operation is used to terminate identity federation or clear the current attach configuration before you attach another identity source.
//
// Disconnecting clears the AgentCore users and associated configurations imported from the identity source. For `CustomSSO`, it also clears the shared user pool and SAML and SCIM access resources. Applications created by users in the enterprise identity system must still be managed on the enterprise side.
//
// ### Preparations before disconnecting
//
// Remove the reference relationships between the related users and collaboration teams first. If teams still reference these users, the API operation rejects the disconnection request.
//
// If you only want to temporarily disable a specific capability, you can retain the identity source binding. Use [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) to disable logon, or use [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM) to disable SCIM synchronization.
//
// ### Verify the disconnection result
//
// The disconnection is executed in the background. After the request is accepted, `DISCONNECTING` is returned. You can query the progress by calling [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider):
//
// - If the query returns `IdentityProviderNotFound`, the binding has been removed and the disconnection is complete.
//
// - If `DISCONNECT_FAILED` is returned, the cleanup failed. Some resources may have already been cleared. You can initiate the disconnection again.
//
// You can bind a new identity source to the workspace only after the disconnection is complete.
//
// @param request - DeleteIdentityProviderRequest
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProvider(workspaceId *string, identityProviderType *string, request *DeleteIdentityProviderRequest) (_result *DeleteIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.DeleteIdentityProviderWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a specified SCIM credential. This operation does not trigger automatic creation of a new token, modify the AC synchronization switch, or disable SAML logon.
//
// Description:
//
// Revokes a SCIM access token to deactivate an old token or clean up credentials that are no longer in use or have expired.
//
// SCIM tokens are used to authorize enterprise identity systems to send user provisioning requests to user pools. After revocation, the enterprise identity system can no longer use the token to access the SCIM endpoint.
//
// ### Usage
//
// **Step 1:*	- Query the credential list by calling [ListIdentityProviderSCIMCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderSCIMCredentials) to obtain the target credentialId.
//
// **Step 2:*	- If the enterprise identity system is still using the token, configure another valid token for it first.
//
// **Step 3:*	- Call this operation to revoke the target credential, and then refresh the list to confirm the result.
//
// This operation only revokes the specified token. Synchronized users, other tokens, logon configurations, and the user synchronization switch are retained. To disable user synchronization, use [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM).
//
// @param request - DeleteIdentityProviderSCIMCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderSCIMCredentialResponse
func (client *Client) DeleteIdentityProviderSCIMCredentialWithOptions(workspaceId *string, credentialId *string, request *DeleteIdentityProviderSCIMCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderSCIMCredentialResponse, _err error) {
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
		Action:      dara.String("DeleteIdentityProviderSCIMCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning/credentials/" + dara.PercentEncode(dara.StringValue(credentialId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdentityProviderSCIMCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a specified SCIM credential. This operation does not trigger automatic creation of a new token, modify the AC synchronization switch, or disable SAML logon.
//
// Description:
//
// Revokes a SCIM access token to deactivate an old token or clean up credentials that are no longer in use or have expired.
//
// SCIM tokens are used to authorize enterprise identity systems to send user provisioning requests to user pools. After revocation, the enterprise identity system can no longer use the token to access the SCIM endpoint.
//
// ### Usage
//
// **Step 1:*	- Query the credential list by calling [ListIdentityProviderSCIMCredentials](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderSCIMCredentials) to obtain the target credentialId.
//
// **Step 2:*	- If the enterprise identity system is still using the token, configure another valid token for it first.
//
// **Step 3:*	- Call this operation to revoke the target credential, and then refresh the list to confirm the result.
//
// This operation only revokes the specified token. Synchronized users, other tokens, logon configurations, and the user synchronization switch are retained. To disable user synchronization, use [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM).
//
// @param request - DeleteIdentityProviderSCIMCredentialRequest
//
// @return DeleteIdentityProviderSCIMCredentialResponse
func (client *Client) DeleteIdentityProviderSCIMCredential(workspaceId *string, credentialId *string, request *DeleteIdentityProviderSCIMCredentialRequest) (_result *DeleteIdentityProviderSCIMCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIdentityProviderSCIMCredentialResponse{}
	_body, _err := client.DeleteIdentityProviderSCIMCredentialWithOptions(workspaceId, credentialId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the image cache association in a workspace and checks for agent references and instance usage before submission.
//
// Description:
//
// Deletes the image cache association in a workspace and checks for agent references and instance usage before submission.
//
// ### Deletion conditions
//
// The deletion request is rejected if the cache is still referenced by an agent with the cache switch enabled, or if instances are still using it or resources are being released. You can first query `agentCount`, `canDelete`, and `deleteBlockReason` by calling [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache).
//
// ### Query results
//
// The `imageId` in the response corresponds to this request. The current implementation of `deletionPending` returns `false`. To confirm whether the cache is still visible, call [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache) or [ListImageCaches](https://api.aliyun.com/document/AgentCore/2026-08-04/ListImageCaches).
//
// This operation deletes the cache association in the specified workspace, but does not delete the source ACR repository or the source image.
//
// @param request - DeleteImageCacheRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageCacheResponse
func (client *Client) DeleteImageCacheWithOptions(imageId *string, request *DeleteImageCacheRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImageCache"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/image-caches/" + dara.PercentEncode(dara.StringValue(imageId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the image cache association in a workspace and checks for agent references and instance usage before submission.
//
// Description:
//
// Deletes the image cache association in a workspace and checks for agent references and instance usage before submission.
//
// ### Deletion conditions
//
// The deletion request is rejected if the cache is still referenced by an agent with the cache switch enabled, or if instances are still using it or resources are being released. You can first query `agentCount`, `canDelete`, and `deleteBlockReason` by calling [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache).
//
// ### Query results
//
// The `imageId` in the response corresponds to this request. The current implementation of `deletionPending` returns `false`. To confirm whether the cache is still visible, call [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache) or [ListImageCaches](https://api.aliyun.com/document/AgentCore/2026-08-04/ListImageCaches).
//
// This operation deletes the cache association in the specified workspace, but does not delete the source ACR repository or the source image.
//
// @param request - DeleteImageCacheRequest
//
// @return DeleteImageCacheResponse
func (client *Client) DeleteImageCache(imageId *string, request *DeleteImageCacheRequest) (_result *DeleteImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DeleteImageCacheWithOptions(imageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a managed Harness agent in a specified workspace and asynchronously cleans up its platform-associated resources.
//
// Description:
//
// Deletes a managed Harness agent in a specified workspace and asynchronously cleans up its platform-associated resources.
//
// ### Prerequisites
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion results
//
// A successful response indicates that the deletion request has been accepted, and the response summary is typically `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// You can call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query the deletion progress. After the deletion is complete, a details query returns a resource not found error. Runtime resources are cleaned up in the background.
//
// @param request - DeleteManagedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteManagedAgentResponse
func (client *Client) DeleteManagedAgentWithOptions(workspaceId *string, agentId *string, request *DeleteManagedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteManagedAgentResponse, _err error) {
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
		Action:      dara.String("DeleteManagedAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteManagedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a managed Harness agent in a specified workspace and asynchronously cleans up its platform-associated resources.
//
// Description:
//
// Deletes a managed Harness agent in a specified workspace and asynchronously cleans up its platform-associated resources.
//
// ### Prerequisites
//
// The workspace must be writable. If the agent is still associated with a team, dissociate it from the team first.
//
// ### Deletion results
//
// A successful response indicates that the deletion request has been accepted, and the response summary is typically `status=Deleting`. If the agent is already being deleted, a subsequent request returns the current summary.
//
// You can call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query the deletion progress. After the deletion is complete, a details query returns a resource not found error. Runtime resources are cleaned up in the background.
//
// @param request - DeleteManagedAgentRequest
//
// @return DeleteManagedAgentResponse
func (client *Client) DeleteManagedAgent(workspaceId *string, agentId *string, request *DeleteManagedAgentRequest) (_result *DeleteManagedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteManagedAgentResponse{}
	_body, _err := client.DeleteManagedAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified MCP service. The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
//
// Description:
//
// Deletes a specified MCP service.
//
// The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
//
// @param request - DeleteMcpRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcpWithOptions(mcpServerId *string, workspaceId *string, request *DeleteMcpRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMcpResponse, _err error) {
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
		Action:      dara.String("DeleteMcp"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified MCP service. The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
//
// Description:
//
// Deletes a specified MCP service.
//
// The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
//
// @param request - DeleteMcpRequest
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcp(mcpServerId *string, workspaceId *string, request *DeleteMcpRequest) (_result *DeleteMcpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMcpResponse{}
	_body, _err := client.DeleteMcpWithOptions(mcpServerId, workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a long-term memory by memory ID.
//
// Description:
//
// Deletes a long-term memory by memory ID.
//
// @param request - DeleteMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspaceId *string, memoryStoreName *string, memoryId *string, request *DeleteMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
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
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Deletes a long-term memory by memory ID.
//
// Description:
//
// Deletes a long-term memory by memory ID.
//
// @param request - DeleteMemoryRequest
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspaceId *string, memoryStoreName *string, memoryId *string, request *DeleteMemoryRequest) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspaceId, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously deletes a MemoryStore and all its long-term memories.
//
// Description:
//
// Asynchronously deletes a specified long-term memory store (MemoryStore) and all its long-term memories.
//
// @param request - DeleteMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryStoreResponse
func (client *Client) DeleteMemoryStoreWithOptions(workspaceId *string, memoryStoreName *string, request *DeleteMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryStoreResponse, _err error) {
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
		Action:      dara.String("DeleteMemoryStore"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously deletes a MemoryStore and all its long-term memories.
//
// Description:
//
// Asynchronously deletes a specified long-term memory store (MemoryStore) and all its long-term memories.
//
// @param request - DeleteMemoryStoreRequest
//
// @return DeleteMemoryStoreResponse
func (client *Client) DeleteMemoryStore(workspaceId *string, memoryStoreName *string, request *DeleteMemoryStoreRequest) (_result *DeleteMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryStoreResponse{}
	_body, _err := client.DeleteMemoryStoreWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model in a specified workspace. Models that are in use cannot be deleted.
//
// Description:
//
// Deletes a model in a specified workspace. Models that are in use cannot be deleted.
//
// @param request - DeleteModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithOptions(workspaceId *string, modelId *string, request *DeleteModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model in a specified workspace. Models that are in use cannot be deleted.
//
// Description:
//
// Deletes a model in a specified workspace. Models that are in use cannot be deleted.
//
// @param request - DeleteModelRequest
//
// @return DeleteModelResponse
func (client *Client) DeleteModel(workspaceId *string, modelId *string, request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(workspaceId, modelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an asynchronous deletion task for a specified model connection. The connection cannot be deleted if it has associated models or runtime references.
//
// Description:
//
// Submits an asynchronous deletion task for a specified model connection. The connection cannot be deleted if it has associated models or runtime references.
//
// @param request - DeleteModelConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelConnectionResponse
func (client *Client) DeleteModelConnectionWithOptions(workspaceId *string, connectionId *string, request *DeleteModelConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelConnection"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/model-connections/" + dara.PercentEncode(dara.StringValue(connectionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an asynchronous deletion task for a specified model connection. The connection cannot be deleted if it has associated models or runtime references.
//
// Description:
//
// Submits an asynchronous deletion task for a specified model connection. The connection cannot be deleted if it has associated models or runtime references.
//
// @param request - DeleteModelConnectionRequest
//
// @return DeleteModelConnectionResponse
func (client *Client) DeleteModelConnection(workspaceId *string, connectionId *string, request *DeleteModelConnectionRequest) (_result *DeleteModelConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelConnectionResponse{}
	_body, _err := client.DeleteModelConnectionWithOptions(workspaceId, connectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a permission rule.
//
// Description:
//
// Deletes an access rule in a workspace and revokes the configuration of the rule. Repeated deletions still return a success response.
//
// @param request - DeletePermissionRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePermissionRuleResponse
func (client *Client) DeletePermissionRuleWithOptions(workspaceId *string, permissionRuleId *string, request *DeletePermissionRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePermissionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePermissionRule"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/permission-rules/" + dara.PercentEncode(dara.StringValue(permissionRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePermissionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a permission rule.
//
// Description:
//
// Deletes an access rule in a workspace and revokes the configuration of the rule. Repeated deletions still return a success response.
//
// @param request - DeletePermissionRuleRequest
//
// @return DeletePermissionRuleResponse
func (client *Client) DeletePermissionRule(workspaceId *string, permissionRuleId *string, request *DeletePermissionRuleRequest) (_result *DeletePermissionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePermissionRuleResponse{}
	_body, _err := client.DeletePermissionRuleWithOptions(workspaceId, permissionRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a skill and its related version data from the specified workspace. The deletion operation is irreversible.
//
// Description:
//
// Deletes a skill and its related version data from the specified workspace.
//
// The deletion operation is irreversible.
//
// @param request - DeleteSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkillWithOptions(workspaceId *string, skillName *string, request *DeleteSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSkillResponse, _err error) {
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
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a skill and its related version data from the specified workspace. The deletion operation is irreversible.
//
// Description:
//
// Deletes a skill and its related version data from the specified workspace.
//
// The deletion operation is irreversible.
//
// @param request - DeleteSkillRequest
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkill(workspaceId *string, skillName *string, request *DeleteSkillRequest) (_result *DeleteSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSkillResponse{}
	_body, _err := client.DeleteSkillWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the draft version of the specified skill that is currently being edited.
//
// Description:
//
// Deletes the draft of a skill that is being edited in normal version mode.
//
// In `HEAD` mode, a persistent draft is used, and the delete draft operation is not supported. To modify the content, continue editing the draft. To remove the entire skill, use [DeleteSkill](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteSkill).
//
// @param request - DeleteSkillDraftRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillDraftResponse
func (client *Client) DeleteSkillDraftWithOptions(workspaceId *string, skillName *string, request *DeleteSkillDraftRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSkillDraftResponse, _err error) {
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
		Action:      dara.String("DeleteSkillDraft"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/draft"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the draft version of the specified skill that is currently being edited.
//
// Description:
//
// Deletes the draft of a skill that is being edited in normal version mode.
//
// In `HEAD` mode, a persistent draft is used, and the delete draft operation is not supported. To modify the content, continue editing the draft. To remove the entire skill, use [DeleteSkill](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteSkill).
//
// @param request - DeleteSkillDraftRequest
//
// @return DeleteSkillDraftResponse
func (client *Client) DeleteSkillDraft(workspaceId *string, skillName *string, request *DeleteSkillDraftRequest) (_result *DeleteSkillDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSkillDraftResponse{}
	_body, _err := client.DeleteSkillDraftWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a team in a specified workspace. Deleting a team does not delete the users and agents within it, but only removes the membership associations.
//
// Description:
//
// Deletes a team in a specified workspace.
//
// @param request - DeleteTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTeamResponse
func (client *Client) DeleteTeamWithOptions(workspaceId *string, teamId *string, request *DeleteTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTeam"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/teams/" + dara.PercentEncode(dara.StringValue(teamId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a team in a specified workspace. Deleting a team does not delete the users and agents within it, but only removes the membership associations.
//
// Description:
//
// Deletes a team in a specified workspace.
//
// @param request - DeleteTeamRequest
//
// @return DeleteTeamResponse
func (client *Client) DeleteTeam(workspaceId *string, teamId *string, request *DeleteTeamRequest) (_result *DeleteTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTeamResponse{}
	_body, _err := client.DeleteTeamWithOptions(workspaceId, teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user from a specified workspace. A user cannot be deleted if they are still a member of any team. Remove the user from the team first.
//
// Description:
//
// Deletes a user from a specified workspace.
//
// @param request - DeleteUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(workspaceId *string, agentCoreUserId *string, request *DeleteUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users/" + dara.PercentEncode(dara.StringValue(agentCoreUserId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user from a specified workspace. A user cannot be deleted if they are still a member of any team. Remove the user from the team first.
//
// Description:
//
// Deletes a user from a specified workspace.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(workspaceId *string, agentCoreUserId *string, request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(workspaceId, agentCoreUserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace that is no longer in use and ends the management lifecycle of the workspace.
//
// Description:
//
// Deletes a workspace that is no longer in use and ends the management lifecycle of the workspace. The API operation processes repeated deletion requests for the same workspace in the `Deleting` or `Deleted` state based on idempotence.
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceId *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace that is no longer in use and ends the management lifecycle of the workspace.
//
// Description:
//
// Deletes a workspace that is no longer in use and ends the management lifecycle of the workspace. The API operation processes repeated deletion requests for the same workspace in the `Deleting` or `Deleted` state based on idempotence.
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceId *string, request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the QoderCLI connector in a specified workspace. The connector can be disabled only when no agent is using it. After being disabled, the connector and its service account key configuration are removed.
//
// Description:
//
// Disables the QoderCLI connector in a workspace.
//
// @param request - DisableConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableConnectorResponse
func (client *Client) DisableConnectorWithOptions(workspaceId *string, connectorName *string, request *DisableConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableConnectorResponse, _err error) {
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
		Action:      dara.String("DisableConnector"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName)) + "/actions/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the QoderCLI connector in a specified workspace. The connector can be disabled only when no agent is using it. After being disabled, the connector and its service account key configuration are removed.
//
// Description:
//
// Disables the QoderCLI connector in a workspace.
//
// @param request - DisableConnectorRequest
//
// @return DisableConnectorResponse
func (client *Client) DisableConnector(workspaceId *string, connectorName *string, request *DisableConnectorRequest) (_result *DisableConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableConnectorResponse{}
	_body, _err := client.DisableConnectorWithOptions(workspaceId, connectorName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the OSS presigned download URL for a specified AgentSpec to download the AgentSpec ZIP package.
//
// Description:
//
// Retrieves the OSS presigned download URL for a specified agent template to download the agent template ZIP package.
//
// @param request - DownloadAgentSpecViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadAgentSpecViaOssResponse
func (client *Client) DownloadAgentSpecViaOssWithOptions(workspaceId *string, agentSpecName *string, request *DownloadAgentSpecViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DownloadAgentSpecViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpecVersion) {
		query["agentSpecVersion"] = request.AgentSpecVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadAgentSpecViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/actions/download-via-oss"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadAgentSpecViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the OSS presigned download URL for a specified AgentSpec to download the AgentSpec ZIP package.
//
// Description:
//
// Retrieves the OSS presigned download URL for a specified agent template to download the agent template ZIP package.
//
// @param request - DownloadAgentSpecViaOssRequest
//
// @return DownloadAgentSpecViaOssResponse
func (client *Client) DownloadAgentSpecViaOss(workspaceId *string, agentSpecName *string, request *DownloadAgentSpecViaOssRequest) (_result *DownloadAgentSpecViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DownloadAgentSpecViaOssResponse{}
	_body, _err := client.DownloadAgentSpecViaOssWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the OSS pre-signed download URL for a specified skill version to download the corresponding skill ZIP package.
//
// Description:
//
// Retrieves the OSS pre-signed download URL for a specified skill version to download the corresponding skill ZIP package.
//
// @param request - DownloadSkillVersionViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadSkillVersionViaOssResponse
func (client *Client) DownloadSkillVersionViaOssWithOptions(workspaceId *string, skillName *string, skillVersion *string, request *DownloadSkillVersionViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DownloadSkillVersionViaOssResponse, _err error) {
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
		Action:      dara.String("DownloadSkillVersionViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion)) + "/actions/download-via-oss"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadSkillVersionViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the OSS pre-signed download URL for a specified skill version to download the corresponding skill ZIP package.
//
// Description:
//
// Retrieves the OSS pre-signed download URL for a specified skill version to download the corresponding skill ZIP package.
//
// @param request - DownloadSkillVersionViaOssRequest
//
// @return DownloadSkillVersionViaOssResponse
func (client *Client) DownloadSkillVersionViaOss(workspaceId *string, skillName *string, skillVersion *string, request *DownloadSkillVersionViaOssRequest) (_result *DownloadSkillVersionViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DownloadSkillVersionViaOssResponse{}
	_body, _err := client.DownloadSkillVersionViaOssWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the QoderCLI connector in a specified workspace. The request must provide a site, an API key, and at least one service account key. The service verifies all account keys before saving the configuration.
//
// Description:
//
// Enables the QoderCLI connector in a workspace. Before enabling, verify the credentials by using [VerifyConnector](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyConnector).
//
// @param tmpReq - EnableConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableConnectorResponse
func (client *Client) EnableConnectorWithOptions(workspaceId *string, connectorName *string, tmpReq *EnableConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnableConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableConnector"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName)) + "/actions/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the QoderCLI connector in a specified workspace. The request must provide a site, an API key, and at least one service account key. The service verifies all account keys before saving the configuration.
//
// Description:
//
// Enables the QoderCLI connector in a workspace. Before enabling, verify the credentials by using [VerifyConnector](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyConnector).
//
// @param request - EnableConnectorRequest
//
// @return EnableConnectorResponse
func (client *Client) EnableConnector(workspaceId *string, connectorName *string, request *EnableConnectorRequest) (_result *EnableConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableConnectorResponse{}
	_body, _err := client.EnableConnectorWithOptions(workspaceId, connectorName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a custom image preparation task by using a submitted clientToken and queries the current progress.
//
// Description:
//
// Uses the `clientToken` of a create or update request to retrieve a custom image preparation task, which allows you to continue querying the processing progress after a request response is lost.
//
// ### Search scope
//
// Pass the `clientToken` used when submitting [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent) as a query parameter, and use the original workspace.
//
// The preparation task can be retrieved only if the submission result includes `imagePreparation`. If no corresponding preparation task exists, a resource not found error is returned.
//
// ### Query phases
//
// | `phase` | Description |
//
// | --- | --- |
//
// | `QUEUED`, `CHECKING` | Queued or checking the source. |
//
// | `COPYING` | Copying the image. |
//
// | `STARTING` | Starting the target version. |
//
// | `HEALTH_CHECKING` | Waiting for the target version instance to be ready. |
//
// | `SUCCEEDED` | The target version deployment succeeded. |
//
// | `FAILED` | The preparation or deployment failed. Check `failedPhase` and `error`. |
//
// | `CANCELLED` | The target version has been replaced by a subsequent configuration or canceled. |
//
// After the task is found, you can use the `preparationId` to call [GetCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgentImagePreparation) to continuously query the progress.
//
// @param request - FindCustomAgentImagePreparationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindCustomAgentImagePreparationResponse
func (client *Client) FindCustomAgentImagePreparationWithOptions(workspaceId *string, request *FindCustomAgentImagePreparationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FindCustomAgentImagePreparationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindCustomAgentImagePreparation"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-image-preparations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FindCustomAgentImagePreparationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom image preparation task by using a submitted clientToken and queries the current progress.
//
// Description:
//
// Uses the `clientToken` of a create or update request to retrieve a custom image preparation task, which allows you to continue querying the processing progress after a request response is lost.
//
// ### Search scope
//
// Pass the `clientToken` used when submitting [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent) as a query parameter, and use the original workspace.
//
// The preparation task can be retrieved only if the submission result includes `imagePreparation`. If no corresponding preparation task exists, a resource not found error is returned.
//
// ### Query phases
//
// | `phase` | Description |
//
// | --- | --- |
//
// | `QUEUED`, `CHECKING` | Queued or checking the source. |
//
// | `COPYING` | Copying the image. |
//
// | `STARTING` | Starting the target version. |
//
// | `HEALTH_CHECKING` | Waiting for the target version instance to be ready. |
//
// | `SUCCEEDED` | The target version deployment succeeded. |
//
// | `FAILED` | The preparation or deployment failed. Check `failedPhase` and `error`. |
//
// | `CANCELLED` | The target version has been replaced by a subsequent configuration or canceled. |
//
// After the task is found, you can use the `preparationId` to call [GetCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgentImagePreparation) to continuously query the progress.
//
// @param request - FindCustomAgentImagePreparationRequest
//
// @return FindCustomAgentImagePreparationResponse
func (client *Client) FindCustomAgentImagePreparation(workspaceId *string, request *FindCustomAgentImagePreparationRequest) (_result *FindCustomAgentImagePreparationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindCustomAgentImagePreparationResponse{}
	_body, _err := client.FindCustomAgentImagePreparationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Forcibly publishes a specified skill version by bypassing the standard review process.
//
// Description:
//
// Forcibly publishes a specified skill version by bypassing the standard review process.
//
// @param tmpReq - ForcePublishSkillVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForcePublishSkillVersionResponse
func (client *Client) ForcePublishSkillVersionWithOptions(workspaceId *string, skillName *string, skillVersion *string, tmpReq *ForcePublishSkillVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ForcePublishSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ForcePublishSkillVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForcePublishSkillVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion)) + "/actions/force-publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForcePublishSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forcibly publishes a specified skill version by bypassing the standard review process.
//
// Description:
//
// Forcibly publishes a specified skill version by bypassing the standard review process.
//
// @param request - ForcePublishSkillVersionRequest
//
// @return ForcePublishSkillVersionResponse
func (client *Client) ForcePublishSkillVersion(workspaceId *string, skillName *string, skillVersion *string, request *ForcePublishSkillVersionRequest) (_result *ForcePublishSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ForcePublishSkillVersionResponse{}
	_body, _err := client.ForcePublishSkillVersionWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access application.
//
// Description:
//
// Queries the configuration of an access application in a workspace to view the application access information and user source settings.
//
// The query result does not include the plaintext of the application secret. The secret is returned only when the application is created.
//
// @param request - GetAccessApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessApplicationResponse
func (client *Client) GetAccessApplicationWithOptions(workspaceId *string, applicationId *string, request *GetAccessApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAccessApplicationResponse, _err error) {
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
		Action:      dara.String("GetAccessApplication"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/access-applications/" + dara.PercentEncode(dara.StringValue(applicationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access application.
//
// Description:
//
// Queries the configuration of an access application in a workspace to view the application access information and user source settings.
//
// The query result does not include the plaintext of the application secret. The secret is returned only when the application is created.
//
// @param request - GetAccessApplicationRequest
//
// @return GetAccessApplicationResponse
func (client *Client) GetAccessApplication(workspaceId *string, applicationId *string, request *GetAccessApplicationRequest) (_result *GetAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAccessApplicationResponse{}
	_body, _err := client.GetAccessApplicationWithOptions(workspaceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the time-limited read-only download URL for the currently saved fixed code package of a custom code or image agent.
//
// Description:
//
// Retrieves the download URL for the currently saved code package of a custom code or image agent, which can be used to download, inspect, or reuse the fixed code package.
//
// ### Download scope
//
// Locates the fixed code package referenced by the currently saved configuration by using `workspaceId` and `agentId`. This applies to code packages uploaded locally or imported from OSS. Arbitrary bucket names, object keys, or historical version parameters are not accepted.
//
// The current configuration must be deployed as a code package and the fixed package must be available. You do not need to wait for the running instance to be ready. An error is returned if the object has been changed, is missing, or is unreadable.
//
// ### Use the response
//
// Use `downloadUrl` to download the file. The `fileName`, `fileSize`, `format`, `etag`, and `crc64` fields describe the corresponding fixed package. The unit of `fileSize` is bytes. The `resourceVersionId` field identifies the saved configuration version corresponding to this download.
//
// The URL is valid for up to 10 minutes and does not exceed the validity period of the underlying credentials, subject to `expiresAt`. Retrieve the URL again after it expires.
//
// @param request - GetAgentCodePackageDownloadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentCodePackageDownloadUrlResponse
func (client *Client) GetAgentCodePackageDownloadUrlWithOptions(workspaceId *string, agentId *string, request *GetAgentCodePackageDownloadUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentCodePackageDownloadUrlResponse, _err error) {
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
		Action:      dara.String("GetAgentCodePackageDownloadUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/code-package/download-url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentCodePackageDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the time-limited read-only download URL for the currently saved fixed code package of a custom code or image agent.
//
// Description:
//
// Retrieves the download URL for the currently saved code package of a custom code or image agent, which can be used to download, inspect, or reuse the fixed code package.
//
// ### Download scope
//
// Locates the fixed code package referenced by the currently saved configuration by using `workspaceId` and `agentId`. This applies to code packages uploaded locally or imported from OSS. Arbitrary bucket names, object keys, or historical version parameters are not accepted.
//
// The current configuration must be deployed as a code package and the fixed package must be available. You do not need to wait for the running instance to be ready. An error is returned if the object has been changed, is missing, or is unreadable.
//
// ### Use the response
//
// Use `downloadUrl` to download the file. The `fileName`, `fileSize`, `format`, `etag`, and `crc64` fields describe the corresponding fixed package. The unit of `fileSize` is bytes. The `resourceVersionId` field identifies the saved configuration version corresponding to this download.
//
// The URL is valid for up to 10 minutes and does not exceed the validity period of the underlying credentials, subject to `expiresAt`. Retrieve the URL again after it expires.
//
// @param request - GetAgentCodePackageDownloadUrlRequest
//
// @return GetAgentCodePackageDownloadUrlResponse
func (client *Client) GetAgentCodePackageDownloadUrl(workspaceId *string, agentId *string, request *GetAgentCodePackageDownloadUrlRequest) (_result *GetAgentCodePackageDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentCodePackageDownloadUrlResponse{}
	_body, _err := client.GetAgentCodePackageDownloadUrlWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a RAM authorization link for the specified code package source bucket for the workspace resource role.
//
// Description:
//
// Generates an authorization link for the workspace resource role to read the code package source bucket, which is used to complete authorization before deploying an existing OSS code package.
//
// ### Authorization steps
//
// **Step 1:*	- Provide the `ossBucketName` to obtain the `authorizeUrl` for the current workspace resource role.
//
// **Step 2:*	- Open the returned link and confirm the authorization.
//
// **Step 3:*	- Call [VerifyAgentCodePackageOssRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyAgentCodePackageOssRamAuthorization) to query the policy binding status. To verify a specific code package object, call [ValidateAgentCodePackageFromOss](https://api.aliyun.com/document/AgentCore/2026-08-04/ValidateAgentCodePackageFromOss).
//
// ### Authorization scope
//
// The role is selected by the server and cannot be overridden by the caller. This authorization reuses the `OSSObjectRestoreCustom` template, which includes permissions such as reading, writing, and deleting objects, and is not a read-only authorization. Existing authorizations for the same role and bucket can be reused.
//
// Obtaining the link itself does not complete the authorization, transfer the code package, or modify the default OSS binding of the workspace.
//
// @param request - GetAgentCodePackageOssRamAuthorizeUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentCodePackageOssRamAuthorizeUrlResponse
func (client *Client) GetAgentCodePackageOssRamAuthorizeUrlWithOptions(workspaceId *string, request *GetAgentCodePackageOssRamAuthorizeUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentCodePackageOssRamAuthorizeUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssBucketName) {
		query["ossBucketName"] = request.OssBucketName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentCodePackageOssRamAuthorizeUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-code-packages/oss/authorize"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentCodePackageOssRamAuthorizeUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a RAM authorization link for the specified code package source bucket for the workspace resource role.
//
// Description:
//
// Generates an authorization link for the workspace resource role to read the code package source bucket, which is used to complete authorization before deploying an existing OSS code package.
//
// ### Authorization steps
//
// **Step 1:*	- Provide the `ossBucketName` to obtain the `authorizeUrl` for the current workspace resource role.
//
// **Step 2:*	- Open the returned link and confirm the authorization.
//
// **Step 3:*	- Call [VerifyAgentCodePackageOssRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyAgentCodePackageOssRamAuthorization) to query the policy binding status. To verify a specific code package object, call [ValidateAgentCodePackageFromOss](https://api.aliyun.com/document/AgentCore/2026-08-04/ValidateAgentCodePackageFromOss).
//
// ### Authorization scope
//
// The role is selected by the server and cannot be overridden by the caller. This authorization reuses the `OSSObjectRestoreCustom` template, which includes permissions such as reading, writing, and deleting objects, and is not a read-only authorization. Existing authorizations for the same role and bucket can be reused.
//
// Obtaining the link itself does not complete the authorization, transfer the code package, or modify the default OSS binding of the workspace.
//
// @param request - GetAgentCodePackageOssRamAuthorizeUrlRequest
//
// @return GetAgentCodePackageOssRamAuthorizeUrlResponse
func (client *Client) GetAgentCodePackageOssRamAuthorizeUrl(workspaceId *string, request *GetAgentCodePackageOssRamAuthorizeUrlRequest) (_result *GetAgentCodePackageOssRamAuthorizeUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentCodePackageOssRamAuthorizeUrlResponse{}
	_body, _err := client.GetAgentCodePackageOssRamAuthorizeUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a time-limited upload URL for a local code package, and returns the upload method, request headers, and the sourceUri for subsequent deployment.
//
// Description:
//
// Retrieves a time-limited direct upload URL and upload requirements for a local code package. After the upload is complete, use the returned `sourceUri` to submit the agent deployment.
//
// ### Request requirements
//
// | Field | Requirement |
//
// | --- | --- |
//
// | `fileName` | The name of the ZIP or TAR.GZ file without the folder. Maximum length: 255 characters. |
//
// | `format` | `zip` or `tar.gz`. Must match the file name extension. |
//
// | `fileSize` | The size of the compressed package in bytes. Valid values: 1 to 500000000. |
//
// ### Upload and deployment
//
// **Step 1:*	- Use the `uploadMethod`, `uploadUrl`, and `headers` from the response to upload the file directly. This operation only generates the URL and does not upload the file content.
//
// **Step 2:*	- After the upload is complete, specify the `sourceUri` in the `artifact.code.sourceUri` parameter of the [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent) operation. You must also provide the `languageRuntime`, `entrypoint`, and service port.
//
// **Step 3:*	- During creation or update, the actual object and its size are verified again. The configuration is saved after the object is transferred to a fixed replica.
//
// The upload URL is valid for up to 10 minutes and does not exceed the expiration time of the underlying temporary credentials. The exact expiration time is subject to the `expiresAt` value. If the URL expires, obtain a new one and use the original link and request headers from the response.
//
// @param tmpReq - GetAgentCodePackageUploadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentCodePackageUploadUrlResponse
func (client *Client) GetAgentCodePackageUploadUrlWithOptions(workspaceId *string, tmpReq *GetAgentCodePackageUploadUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentCodePackageUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentCodePackageUploadUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentCodePackageUploadUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-code-packages/upload-url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentCodePackageUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a time-limited upload URL for a local code package, and returns the upload method, request headers, and the sourceUri for subsequent deployment.
//
// Description:
//
// Retrieves a time-limited direct upload URL and upload requirements for a local code package. After the upload is complete, use the returned `sourceUri` to submit the agent deployment.
//
// ### Request requirements
//
// | Field | Requirement |
//
// | --- | --- |
//
// | `fileName` | The name of the ZIP or TAR.GZ file without the folder. Maximum length: 255 characters. |
//
// | `format` | `zip` or `tar.gz`. Must match the file name extension. |
//
// | `fileSize` | The size of the compressed package in bytes. Valid values: 1 to 500000000. |
//
// ### Upload and deployment
//
// **Step 1:*	- Use the `uploadMethod`, `uploadUrl`, and `headers` from the response to upload the file directly. This operation only generates the URL and does not upload the file content.
//
// **Step 2:*	- After the upload is complete, specify the `sourceUri` in the `artifact.code.sourceUri` parameter of the [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent) operation. You must also provide the `languageRuntime`, `entrypoint`, and service port.
//
// **Step 3:*	- During creation or update, the actual object and its size are verified again. The configuration is saved after the object is transferred to a fixed replica.
//
// The upload URL is valid for up to 10 minutes and does not exceed the expiration time of the underlying temporary credentials. The exact expiration time is subject to the `expiresAt` value. If the URL expires, obtain a new one and use the original link and request headers from the response.
//
// @param request - GetAgentCodePackageUploadUrlRequest
//
// @return GetAgentCodePackageUploadUrlResponse
func (client *Client) GetAgentCodePackageUploadUrl(workspaceId *string, request *GetAgentCodePackageUploadUrlRequest) (_result *GetAgentCodePackageUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentCodePackageUploadUrlResponse{}
	_body, _err := client.GetAgentCodePackageUploadUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified agent IM channel.
//
// Description:
//
// Queries the details of a specified agent IM channel.
//
// @param request - GetAgentIMChannelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentIMChannelResponse
func (client *Client) GetAgentIMChannelWithOptions(workspaceId *string, agentId *string, imChannelId *string, request *GetAgentIMChannelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentIMChannelResponse, _err error) {
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
		Action:      dara.String("GetAgentIMChannel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels/" + dara.PercentEncode(dara.StringValue(imChannelId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentIMChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified agent IM channel.
//
// Description:
//
// Queries the details of a specified agent IM channel.
//
// @param request - GetAgentIMChannelRequest
//
// @return GetAgentIMChannelResponse
func (client *Client) GetAgentIMChannel(workspaceId *string, agentId *string, imChannelId *string, request *GetAgentIMChannelRequest) (_result *GetAgentIMChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentIMChannelResponse{}
	_body, _err := client.GetAgentIMChannelWithOptions(workspaceId, agentId, imChannelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the governance details of a specified AgentSpec, including basic information, governance pointers, and summaries of all versions.
//
// Description:
//
// Queries the basic information, version status, and summaries of all versions of an agent template (AgentSpec) for viewing the template and selecting a version.
//
// @param request - GetAgentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSpecResponse
func (client *Client) GetAgentSpecWithOptions(workspaceId *string, agentSpecName *string, request *GetAgentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpecResponse, _err error) {
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
		Action:      dara.String("GetAgentSpec"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the governance details of a specified AgentSpec, including basic information, governance pointers, and summaries of all versions.
//
// Description:
//
// Queries the basic information, version status, and summaries of all versions of an agent template (AgentSpec) for viewing the template and selecting a version.
//
// @param request - GetAgentSpecRequest
//
// @return GetAgentSpecResponse
func (client *Client) GetAgentSpec(workspaceId *string, agentSpecName *string, request *GetAgentSpecRequest) (_result *GetAgentSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentSpecResponse{}
	_body, _err := client.GetAgentSpecWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API operation to complete the import.
//
// Description:
//
// Retrieves the upload URL for an agent template (AgentSpec) ZIP package to import templates from local files.
//
// Files are transferred through Object Storage Service (OSS). First, upload the file by using the returned URL, and then call the [UploadAgentSpecViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadAgentSpecViaOss) API operation to complete the import.
//
// @param request - GetAgentSpecImportFileUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSpecImportFileUrlResponse
func (client *Client) GetAgentSpecImportFileUrlWithOptions(workspaceId *string, request *GetAgentSpecImportFileUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpecImportFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		query["contentType"] = request.ContentType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentSpecImportFileUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-spec-actions/get-import-file-url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSpecImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API operation to complete the import.
//
// Description:
//
// Retrieves the upload URL for an agent template (AgentSpec) ZIP package to import templates from local files.
//
// Files are transferred through Object Storage Service (OSS). First, upload the file by using the returned URL, and then call the [UploadAgentSpecViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadAgentSpecViaOss) API operation to complete the import.
//
// @param request - GetAgentSpecImportFileUrlRequest
//
// @return GetAgentSpecImportFileUrlResponse
func (client *Client) GetAgentSpecImportFileUrl(workspaceId *string, request *GetAgentSpecImportFileUrlRequest) (_result *GetAgentSpecImportFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentSpecImportFileUrlResponse{}
	_body, _err := client.GetAgentSpecImportFileUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing the latest version.
//
// Description:
//
// Queries the full content of the latest version of an agent template (AgentSpec) for viewing or editing the latest version.
//
// @param request - GetAgentSpecLatestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSpecLatestResponse
func (client *Client) GetAgentSpecLatestWithOptions(workspaceId *string, agentSpecName *string, request *GetAgentSpecLatestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpecLatestResponse, _err error) {
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
		Action:      dara.String("GetAgentSpecLatest"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/latest"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSpecLatestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing the latest version.
//
// Description:
//
// Queries the full content of the latest version of an agent template (AgentSpec) for viewing or editing the latest version.
//
// @param request - GetAgentSpecLatestRequest
//
// @return GetAgentSpecLatestResponse
func (client *Client) GetAgentSpecLatest(workspaceId *string, agentSpecName *string, request *GetAgentSpecLatestRequest) (_result *GetAgentSpecLatestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentSpecLatestResponse{}
	_body, _err := client.GetAgentSpecLatestWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
//
// Description:
//
// Queries the full content of a specified agent template version, including the manifest content, resource files, and metadata.
//
// @param request - GetAgentSpecVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSpecVersionResponse
func (client *Client) GetAgentSpecVersionWithOptions(workspaceId *string, agentSpecName *string, agentSpecVersion *string, request *GetAgentSpecVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpecVersionResponse, _err error) {
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
		Action:      dara.String("GetAgentSpecVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/versions/" + dara.PercentEncode(dara.StringValue(agentSpecVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSpecVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
//
// Description:
//
// Queries the full content of a specified agent template version, including the manifest content, resource files, and metadata.
//
// @param request - GetAgentSpecVersionRequest
//
// @return GetAgentSpecVersionResponse
func (client *Client) GetAgentSpecVersion(workspaceId *string, agentSpecName *string, agentSpecVersion *string, request *GetAgentSpecVersionRequest) (_result *GetAgentSpecVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentSpecVersionResponse{}
	_body, _err := client.GetAgentSpecVersionWithOptions(workspaceId, agentSpecName, agentSpecVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an agent version.
//
// Description:
//
// Queries the details of a saved fixed version for a specified managed or custom agent. The configuration field returns the agent configuration at the time the version was saved.
//
// @param request - GetAgentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentVersionResponse
func (client *Client) GetAgentVersionWithOptions(workspaceId *string, agentId *string, agentVersion *string, request *GetAgentVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentVersionResponse, _err error) {
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
		Action:      dara.String("GetAgentVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/versions/" + dara.PercentEncode(dara.StringValue(agentVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an agent version.
//
// Description:
//
// Queries the details of a saved fixed version for a specified managed or custom agent. The configuration field returns the agent configuration at the time the version was saved.
//
// @param request - GetAgentVersionRequest
//
// @return GetAgentVersionResponse
func (client *Client) GetAgentVersion(workspaceId *string, agentId *string, agentVersion *string, request *GetAgentVersionRequest) (_result *GetAgentVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentVersionResponse{}
	_body, _err := client.GetAgentVersionWithOptions(workspaceId, agentId, agentVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified credential and returns the list of agents bound to the credential. The credential content is returned in a masked format.
//
// Description:
//
// Queries the details of a specified credential and returns the list of agents bound to the credential.
//
// ### Response content
//
// Queries the credential type, description, application scope, and binding information. Supports two types: `apiKey` and `mcpHeader`.
//
// - `apiKey`: The key value is returned in a masked format.
//
// - `mcpHeader`: Returns the configured header name without the header value. The `resourceType` of the applicable resource is `mcpServer`.
//
// To modify the configuration, call [UpdateCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCredential).
//
// @param request - GetCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithOptions(workspaceId *string, credentialId *string, request *GetCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
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
		Action:      dara.String("GetCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/credentials/" + dara.PercentEncode(dara.StringValue(credentialId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified credential and returns the list of agents bound to the credential. The credential content is returned in a masked format.
//
// Description:
//
// Queries the details of a specified credential and returns the list of agents bound to the credential.
//
// ### Response content
//
// Queries the credential type, description, application scope, and binding information. Supports two types: `apiKey` and `mcpHeader`.
//
// - `apiKey`: The key value is returned in a masked format.
//
// - `mcpHeader`: Returns the configured header name without the header value. The `resourceType` of the applicable resource is `mcpServer`.
//
// To modify the configuration, call [UpdateCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCredential).
//
// @param request - GetCredentialRequest
//
// @return GetCredentialResponse
func (client *Client) GetCredential(workspaceId *string, credentialId *string, request *GetCredentialRequest) (_result *GetCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCredentialResponse{}
	_body, _err := client.GetCredentialWithOptions(workspaceId, credentialId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest deployment configuration, version processing results, code package or image information, and running instance status of a custom code or image agent.
//
// Description:
//
// Queries the configuration, artifacts, and running status of a custom code or image agent to confirm the processing and effective results of creation or updates.
//
// ### Query configuration processing results
//
// After creation or update, you can poll this operation. The returned configuration belongs to the latest submitted version and may not have taken effect yet.
//
// | Field | Description |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest expected configuration version. |
//
// | `effectiveSpecVersion` | The currently effective configuration version. |
//
// | `latestVersionStatus` | The processing status of the latest expected version. |
//
// | `latestExecutionError` | The reason why the latest version is waiting for a retry or failed to execute, and whether automatic retry is enabled. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | The snapshot of the number of running instances in each phase. |
//
// `pending`, `processing`, and `waiting_retry` indicate that processing is continuing. `succeeded` indicates that the configuration is processed successfully. If the value is `failed`, troubleshoot the issue based on `latestExecutionError`. Instance status and configuration processing status are queried separately.
//
// ### Code packages and images
//
// In code package mode, `artifact.code.sourceUri` is the saved fixed package address, and `codePackageInfo` provides the package information. To download the current fixed package, call [GetAgentCodePackageDownloadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageDownloadUrl).
//
// In image mode, `customImage` is the most recently saved custom source, and `currentImage` is the image source of the effective version. The two may be different. You can use `imagePreparation.preparationId` to call [GetCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgentImagePreparation) to query the preparation and deployment phases of the current custom image. `imageCachePreparation.statusAtAcceptance` only indicates the cache status at the time of acceptance.
//
// The response does not return the image repository password. If you query an agent that does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
//
// @param request - GetCustomAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomAgentResponse
func (client *Client) GetCustomAgentWithOptions(workspaceId *string, agentId *string, request *GetCustomAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomAgentResponse, _err error) {
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
		Action:      dara.String("GetCustomAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest deployment configuration, version processing results, code package or image information, and running instance status of a custom code or image agent.
//
// Description:
//
// Queries the configuration, artifacts, and running status of a custom code or image agent to confirm the processing and effective results of creation or updates.
//
// ### Query configuration processing results
//
// After creation or update, you can poll this operation. The returned configuration belongs to the latest submitted version and may not have taken effect yet.
//
// | Field | Description |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest expected configuration version. |
//
// | `effectiveSpecVersion` | The currently effective configuration version. |
//
// | `latestVersionStatus` | The processing status of the latest expected version. |
//
// | `latestExecutionError` | The reason why the latest version is waiting for a retry or failed to execute, and whether automatic retry is enabled. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | The snapshot of the number of running instances in each phase. |
//
// `pending`, `processing`, and `waiting_retry` indicate that processing is continuing. `succeeded` indicates that the configuration is processed successfully. If the value is `failed`, troubleshoot the issue based on `latestExecutionError`. Instance status and configuration processing status are queried separately.
//
// ### Code packages and images
//
// In code package mode, `artifact.code.sourceUri` is the saved fixed package address, and `codePackageInfo` provides the package information. To download the current fixed package, call [GetAgentCodePackageDownloadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageDownloadUrl).
//
// In image mode, `customImage` is the most recently saved custom source, and `currentImage` is the image source of the effective version. The two may be different. You can use `imagePreparation.preparationId` to call [GetCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgentImagePreparation) to query the preparation and deployment phases of the current custom image. `imageCachePreparation.statusAtAcceptance` only indicates the cache status at the time of acceptance.
//
// The response does not return the image repository password. If you query an agent that does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
//
// @param request - GetCustomAgentRequest
//
// @return GetCustomAgentResponse
func (client *Client) GetCustomAgent(workspaceId *string, agentId *string, request *GetCustomAgentRequest) (_result *GetCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomAgentResponse{}
	_body, _err := client.GetCustomAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the preparation and deployment progress of the current submission for a custom image by preparation task ID.
//
// Description:
//
// Queries the processing progress, such as image checking, copying, and starting, by preparation task ID to track the result of the current custom image submission.
//
// ### Obtain the task identifier
//
// Pass the value of `imagePreparation.preparationId` from the response of creating or updating a custom code or image agent as the `preparationId` parameter without modification. It identifies the task for the current submission, not the running instance ID.
//
// ### Determine completion
//
// | `phase` | Meaning |
//
// | --- | --- |
//
// | `QUEUED`, `CHECKING` | Queued or checking the source. |
//
// | `COPYING` | Copying the image. |
//
// | `STARTING` | Starting the target version. |
//
// | `HEALTH_CHECKING` | Waiting for the target version instance to be ready. |
//
// | `SUCCEEDED` | The target version is deployed successfully. |
//
// | `FAILED` | The preparation or deployment failed. Check `failedPhase` and `error`. |
//
// | `CANCELLED` | The target version has been replaced by a subsequent configuration or canceled. |
//
// The query only reads the status and does not recreate the copy task. `SUCCEEDED` applies to the current target version. Seeing only old instances running does not indicate that the current submission is complete. If the configuration version is replaced by a subsequent submission, the current task may return `CANCELLED`.
//
// If the create or update response is lost, you can call [FindCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/FindCustomAgentImagePreparation) with the `clientToken` used at that time to retrieve the preparation task.
//
// @param request - GetCustomAgentImagePreparationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomAgentImagePreparationResponse
func (client *Client) GetCustomAgentImagePreparationWithOptions(workspaceId *string, preparationId *string, request *GetCustomAgentImagePreparationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomAgentImagePreparationResponse, _err error) {
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
		Action:      dara.String("GetCustomAgentImagePreparation"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-image-preparations/" + dara.PercentEncode(dara.StringValue(preparationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomAgentImagePreparationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the preparation and deployment progress of the current submission for a custom image by preparation task ID.
//
// Description:
//
// Queries the processing progress, such as image checking, copying, and starting, by preparation task ID to track the result of the current custom image submission.
//
// ### Obtain the task identifier
//
// Pass the value of `imagePreparation.preparationId` from the response of creating or updating a custom code or image agent as the `preparationId` parameter without modification. It identifies the task for the current submission, not the running instance ID.
//
// ### Determine completion
//
// | `phase` | Meaning |
//
// | --- | --- |
//
// | `QUEUED`, `CHECKING` | Queued or checking the source. |
//
// | `COPYING` | Copying the image. |
//
// | `STARTING` | Starting the target version. |
//
// | `HEALTH_CHECKING` | Waiting for the target version instance to be ready. |
//
// | `SUCCEEDED` | The target version is deployed successfully. |
//
// | `FAILED` | The preparation or deployment failed. Check `failedPhase` and `error`. |
//
// | `CANCELLED` | The target version has been replaced by a subsequent configuration or canceled. |
//
// The query only reads the status and does not recreate the copy task. `SUCCEEDED` applies to the current target version. Seeing only old instances running does not indicate that the current submission is complete. If the configuration version is replaced by a subsequent submission, the current task may return `CANCELLED`.
//
// If the create or update response is lost, you can call [FindCustomAgentImagePreparation](https://api.aliyun.com/document/AgentCore/2026-08-04/FindCustomAgentImagePreparation) with the `clientToken` used at that time to retrieve the preparation task.
//
// @param request - GetCustomAgentImagePreparationRequest
//
// @return GetCustomAgentImagePreparationResponse
func (client *Client) GetCustomAgentImagePreparation(workspaceId *string, preparationId *string, request *GetCustomAgentImagePreparationRequest) (_result *GetCustomAgentImagePreparationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomAgentImagePreparationResponse{}
	_body, _err := client.GetCustomAgentImagePreparationWithOptions(workspaceId, preparationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration, model source, version processing status, and external runtime status of a managed agent.
//
// Description:
//
// Queries the platform configuration, version processing status, and external runtime status of a managed agent to verify access and configuration effectiveness.
//
// ### Configuration processing and runtime status
//
// `latestSpecVersion` and `latestVersionStatus` indicate the latest submitted configuration and its processing status. `status` represents the platform resource status, and `externalAgentStatus` represents the external runtime status. A successful configuration acceptance does not mean that the external process is connected.
//
// If `modelSource` is `PLATFORM`, the returned model configuration is used. `RUNTIME` indicates that the model is managed by the external environment. `configuredSkills` indicates explicit skill configurations, and `skills` represents the aggregated skill information.
//
// ### External access
//
// After the configuration takes effect, call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) to query available networks, and then call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain access information. If the agent does not exist, has been deleted, or does not belong to the workspace, a resource not found error is returned.
//
// @param request - GetExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalAgentResponse
func (client *Client) GetExternalAgentWithOptions(workspaceId *string, agentId *string, request *GetExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExternalAgentResponse, _err error) {
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
		Action:      dara.String("GetExternalAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExternalAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration, model source, version processing status, and external runtime status of a managed agent.
//
// Description:
//
// Queries the platform configuration, version processing status, and external runtime status of a managed agent to verify access and configuration effectiveness.
//
// ### Configuration processing and runtime status
//
// `latestSpecVersion` and `latestVersionStatus` indicate the latest submitted configuration and its processing status. `status` represents the platform resource status, and `externalAgentStatus` represents the external runtime status. A successful configuration acceptance does not mean that the external process is connected.
//
// If `modelSource` is `PLATFORM`, the returned model configuration is used. `RUNTIME` indicates that the model is managed by the external environment. `configuredSkills` indicates explicit skill configurations, and `skills` represents the aggregated skill information.
//
// ### External access
//
// After the configuration takes effect, call [GetExternalAgentBootstrapOptions](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgentBootstrapOptions) to query available networks, and then call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain access information. If the agent does not exist, has been deleted, or does not belong to the workspace, a resource not found error is returned.
//
// @param request - GetExternalAgentRequest
//
// @return GetExternalAgentResponse
func (client *Client) GetExternalAgent(workspaceId *string, agentId *string, request *GetExternalAgentRequest) (_result *GetExternalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExternalAgentResponse{}
	_body, _err := client.GetExternalAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available internal and public network access options for a managed agent.
//
// Description:
//
// Queries the available public and internal network access methods for a managed agent to select the network used for subsequent connections.
//
// ### Prerequisites
//
// The managed agent must be in the running state, and the latest desired configuration must be applied. If the configuration is still being created, updated, or not fully applied, query the status first by calling [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent).
//
// ### Select an access network
//
// In the response, the `networkType` in `networkOptions` is `INTRANET` or `INTERNET`, and `available` indicates whether the corresponding access endpoint is available. Select a network with `available=true` and call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain the corresponding access information.
//
// @param request - GetExternalAgentBootstrapOptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalAgentBootstrapOptionsResponse
func (client *Client) GetExternalAgentBootstrapOptionsWithOptions(workspaceId *string, agentId *string, request *GetExternalAgentBootstrapOptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExternalAgentBootstrapOptionsResponse, _err error) {
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
		Action:      dara.String("GetExternalAgentBootstrapOptions"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/bootstrap/options"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExternalAgentBootstrapOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available internal and public network access options for a managed agent.
//
// Description:
//
// Queries the available public and internal network access methods for a managed agent to select the network used for subsequent connections.
//
// ### Prerequisites
//
// The managed agent must be in the running state, and the latest desired configuration must be applied. If the configuration is still being created, updated, or not fully applied, query the status first by calling [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent).
//
// ### Select an access network
//
// In the response, the `networkType` in `networkOptions` is `INTRANET` or `INTERNET`, and `available` indicates whether the corresponding access endpoint is available. Select a network with `available=true` and call [CreateExternalAgentBootstrapToken](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateExternalAgentBootstrapToken) to obtain the corresponding access information.
//
// @param request - GetExternalAgentBootstrapOptionsRequest
//
// @return GetExternalAgentBootstrapOptionsResponse
func (client *Client) GetExternalAgentBootstrapOptions(workspaceId *string, agentId *string, request *GetExternalAgentBootstrapOptionsRequest) (_result *GetExternalAgentBootstrapOptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExternalAgentBootstrapOptionsResponse{}
	_body, _err := client.GetExternalAgentBootstrapOptionsWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding and expected configuration of a specified identity provider. When SAML is configured for CustomSSO, this operation reads back the actual upstream SSO configuration. The SCIM status is queried through a separate API operation.
//
// Description:
//
// Queries the enterprise identity providers bound to a workspace to view the access configuration, logon and user synchronization switches, and configuration processing progress.
//
// For a custom identity provider (`CustomSSO`), this API operation also reads the saved SAML logon configuration for viewing or modifying enterprise single sign-on (SSO) settings.
//
// ### How to interpret the query results
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `identityProviderType` | Indicates whether DingTalk, Lark, or a custom identity provider is currently connected. |
//
// | `loginEnabled` | Indicates whether enterprise account logon is enabled. |
//
// | `saml` | The saved SAML configuration of the enterprise identity system, used for backfilling when editing the configuration. |
//
// | `configuration` | The actual SAML configuration read from the identity service, used to verify whether the logon configuration has taken effect. |
//
// | `status` | The processing progress to attach the identity provider or apply the logon configuration. |
//
// If `CustomSSO` returns `WAITING_CONFIGURATION`, it indicates that the user pool is ready and the SAML logon configuration is pending. In this case, you can call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the access information of AgentCore and continue configuring enterprise single sign-on.
//
// After saving the logon configuration, you can query `status` to track the processing progress. `READY` indicates that the configuration processing is complete. To check whether logon is enabled, you also need to verify `loginEnabled` and `configuration`.
//
// ### Query user synchronization status
//
// SAML logon and SCIM user synchronization are configured separately. SCIM is used to pre-provision users from the enterprise identity system to the user pool. You can call [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM) to query the receiver configuration and synchronization task results.
//
// @param request - GetIdentityProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProviderWithOptions(workspaceId *string, identityProviderType *string, request *GetIdentityProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
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
		Action:      dara.String("GetIdentityProvider"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding and expected configuration of a specified identity provider. When SAML is configured for CustomSSO, this operation reads back the actual upstream SSO configuration. The SCIM status is queried through a separate API operation.
//
// Description:
//
// Queries the enterprise identity providers bound to a workspace to view the access configuration, logon and user synchronization switches, and configuration processing progress.
//
// For a custom identity provider (`CustomSSO`), this API operation also reads the saved SAML logon configuration for viewing or modifying enterprise single sign-on (SSO) settings.
//
// ### How to interpret the query results
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `identityProviderType` | Indicates whether DingTalk, Lark, or a custom identity provider is currently connected. |
//
// | `loginEnabled` | Indicates whether enterprise account logon is enabled. |
//
// | `saml` | The saved SAML configuration of the enterprise identity system, used for backfilling when editing the configuration. |
//
// | `configuration` | The actual SAML configuration read from the identity service, used to verify whether the logon configuration has taken effect. |
//
// | `status` | The processing progress to attach the identity provider or apply the logon configuration. |
//
// If `CustomSSO` returns `WAITING_CONFIGURATION`, it indicates that the user pool is ready and the SAML logon configuration is pending. In this case, you can call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the access information of AgentCore and continue configuring enterprise single sign-on.
//
// After saving the logon configuration, you can query `status` to track the processing progress. `READY` indicates that the configuration processing is complete. To check whether logon is enabled, you also need to verify `loginEnabled` and `configuration`.
//
// ### Query user synchronization status
//
// SAML logon and SCIM user synchronization are configured separately. SCIM is used to pre-provision users from the enterprise identity system to the user pool. You can call [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM) to query the receiver configuration and synchronization task results.
//
// @param request - GetIdentityProviderRequest
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProvider(workspaceId *string, identityProviderType *string, request *GetIdentityProviderRequest) (_result *GetIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.GetIdentityProviderWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the actual SCIM receiving configuration, the expected AgentCore synchronization switch, and the status of the most recent independent SCIM task. SAML configuration is not required beforehand.
//
// Description:
//
// Queries the access configuration and processing progress of SCIM user synchronization to determine whether the receiving end is ready, whether AgentCore user synchronization is enabled, and whether the most recent synchronization operation is completed.
//
// SCIM is a user provisioning protocol. The enterprise identity system sends user information to the user pool through the receiving URL and access token. After AgentCore user synchronization is enabled, users in the user pool are imported into the workspace.
//
// ### How to interpret the query results
//
// | Field | Business meaning |
//
// | --- | --- |
//
// | `configuration` | The actual status and receiving URL of the SCIM receiving endpoint. |
//
// | `syncEnabled` | The expected switch for AgentCore user synchronization. |
//
// | `operationStatus` | The execution result of the most recent operation to enable or disable synchronization. It is an empty string `""` if no historical tasks exist. |
//
// | `status` | The SCIM access state derived from the receiving endpoint configuration, credentials, and operation results. |
//
// ### Proceed based on the status
//
// | `status` | Current stage | Next step |
//
// | --- | --- | --- |
//
// | `UNCONFIGURED` | The receiving end is not ready. | Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to prepare for access. |
//
// | `WAITING_CONFIGURATION` | The URL and valid credentials are available. Waiting for the enterprise-side configuration to be completed and synchronization to be enabled. | Configure the URL and token in the enterprise identity system, and then call [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM). |
//
// | `UPDATING` | Synchronization is being enabled or disabled. | Continue to query this operation and wait for the operation to complete. |
//
// | `READY` | Synchronization is enabled, the current enabling task is successful, and the receiving end and credentials are still available. | You can view the synchronized workspace users. |
//
// | `DISABLED` | Synchronization is disabled. Existing configurations, credentials, and users are retained. | To re-establish access, check and prepare the receiving end first, and then enable synchronization. |
//
// | `UPDATE_FAILED` | The operation failed, or the configuration and credentials after enabling are no longer available. | Check the operation results and access configuration in combination with `operationStatus`. |
//
// This operation queries the SCIM status. The SAML enterprise account logon configuration is queried through [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider). `READY` reflects the result of the current enabling operation. The user assignment and provisioning status in the enterprise identity system must still be verified on the enterprise side.
//
// @param request - GetIdentityProviderSCIMRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderSCIMResponse
func (client *Client) GetIdentityProviderSCIMWithOptions(workspaceId *string, request *GetIdentityProviderSCIMRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderSCIMResponse, _err error) {
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
		Action:      dara.String("GetIdentityProviderSCIM"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderSCIMResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the actual SCIM receiving configuration, the expected AgentCore synchronization switch, and the status of the most recent independent SCIM task. SAML configuration is not required beforehand.
//
// Description:
//
// Queries the access configuration and processing progress of SCIM user synchronization to determine whether the receiving end is ready, whether AgentCore user synchronization is enabled, and whether the most recent synchronization operation is completed.
//
// SCIM is a user provisioning protocol. The enterprise identity system sends user information to the user pool through the receiving URL and access token. After AgentCore user synchronization is enabled, users in the user pool are imported into the workspace.
//
// ### How to interpret the query results
//
// | Field | Business meaning |
//
// | --- | --- |
//
// | `configuration` | The actual status and receiving URL of the SCIM receiving endpoint. |
//
// | `syncEnabled` | The expected switch for AgentCore user synchronization. |
//
// | `operationStatus` | The execution result of the most recent operation to enable or disable synchronization. It is an empty string `""` if no historical tasks exist. |
//
// | `status` | The SCIM access state derived from the receiving endpoint configuration, credentials, and operation results. |
//
// ### Proceed based on the status
//
// | `status` | Current stage | Next step |
//
// | --- | --- | --- |
//
// | `UNCONFIGURED` | The receiving end is not ready. | Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to prepare for access. |
//
// | `WAITING_CONFIGURATION` | The URL and valid credentials are available. Waiting for the enterprise-side configuration to be completed and synchronization to be enabled. | Configure the URL and token in the enterprise identity system, and then call [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM). |
//
// | `UPDATING` | Synchronization is being enabled or disabled. | Continue to query this operation and wait for the operation to complete. |
//
// | `READY` | Synchronization is enabled, the current enabling task is successful, and the receiving end and credentials are still available. | You can view the synchronized workspace users. |
//
// | `DISABLED` | Synchronization is disabled. Existing configurations, credentials, and users are retained. | To re-establish access, check and prepare the receiving end first, and then enable synchronization. |
//
// | `UPDATE_FAILED` | The operation failed, or the configuration and credentials after enabling are no longer available. | Check the operation results and access configuration in combination with `operationStatus`. |
//
// This operation queries the SCIM status. The SAML enterprise account logon configuration is queried through [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider). `READY` reflects the result of the current enabling operation. The user assignment and provisioning status in the enterprise identity system must still be verified on the enterprise side.
//
// @param request - GetIdentityProviderSCIMRequest
//
// @return GetIdentityProviderSCIMResponse
func (client *Client) GetIdentityProviderSCIM(workspaceId *string, request *GetIdentityProviderSCIMRequest) (_result *GetIdentityProviderSCIMResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityProviderSCIMResponse{}
	_body, _err := client.GetIdentityProviderSCIMWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SAML service provider (SP) information of CustomSSO, including the Entity ID, ACS URL, and SP metadata XML, for users to configure an external IdP.
//
// Description:
//
// Retrieves the SAML service provider (SP) access information of AgentCore to configure single sign-on (SSO) in an enterprise identity provider (IdP).
//
// An SP is a service provider that receives logon assertions. An IdP is an enterprise identity provider that verifies user identities and issues logon assertions. This operation returns the SP information on the AgentCore side.
//
// ### Before you begin
//
// The workspace is bound to CustomSSO, and the shared user pool is initialized. The protocol parameter is required and must be set to SAML.
//
// You can call the [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) operation to query the binding information. After the user pool is available, you can obtain the SP information and then complete the SAML configuration.
//
// ### Request example
//
// ```http
//
// GET /workspaces/{workspaceId}/identity-providers/CustomSSO/service-provider?protocol=SAML
//
// ```
//
// ### Response information
//
// | Field | Purpose |
//
// | --- | --- |
//
// | protocol | The access protocol. This operation returns SAML. |
//
// | entityId | The SP identifier of AgentCore. Configure it as the service provider Entity ID in the enterprise IdP. |
//
// | acsUrl | The URL where AgentCore receives SAML logon assertions. Configure it as the ACS URL in the enterprise IdP. |
//
// | metadataXml | The SP metadata XML. Import it to configure the SP in an enterprise IdP that supports metadata import. |
//
// ### Configuration flow
//
// **Step 1:*	- Call this operation to obtain the SP information, and configure the SP identifier and the URL to accept assertions of AgentCore in the enterprise IdP.
//
// **Step 2:*	- Export the metadata XML from the enterprise IdP and save it to AgentCore by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider).
//
// **Step 3:*	- After the configuration is complete, enable single sign-on logon by setting loginEnabled as needed.
//
// @param request - GetIdentityProviderServiceProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderServiceProviderResponse
func (client *Client) GetIdentityProviderServiceProviderWithOptions(workspaceId *string, identityProviderType *string, request *GetIdentityProviderServiceProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderServiceProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Protocol) {
		query["protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProviderServiceProvider"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType)) + "/service-provider"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderServiceProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SAML service provider (SP) information of CustomSSO, including the Entity ID, ACS URL, and SP metadata XML, for users to configure an external IdP.
//
// Description:
//
// Retrieves the SAML service provider (SP) access information of AgentCore to configure single sign-on (SSO) in an enterprise identity provider (IdP).
//
// An SP is a service provider that receives logon assertions. An IdP is an enterprise identity provider that verifies user identities and issues logon assertions. This operation returns the SP information on the AgentCore side.
//
// ### Before you begin
//
// The workspace is bound to CustomSSO, and the shared user pool is initialized. The protocol parameter is required and must be set to SAML.
//
// You can call the [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) operation to query the binding information. After the user pool is available, you can obtain the SP information and then complete the SAML configuration.
//
// ### Request example
//
// ```http
//
// GET /workspaces/{workspaceId}/identity-providers/CustomSSO/service-provider?protocol=SAML
//
// ```
//
// ### Response information
//
// | Field | Purpose |
//
// | --- | --- |
//
// | protocol | The access protocol. This operation returns SAML. |
//
// | entityId | The SP identifier of AgentCore. Configure it as the service provider Entity ID in the enterprise IdP. |
//
// | acsUrl | The URL where AgentCore receives SAML logon assertions. Configure it as the ACS URL in the enterprise IdP. |
//
// | metadataXml | The SP metadata XML. Import it to configure the SP in an enterprise IdP that supports metadata import. |
//
// ### Configuration flow
//
// **Step 1:*	- Call this operation to obtain the SP information, and configure the SP identifier and the URL to accept assertions of AgentCore in the enterprise IdP.
//
// **Step 2:*	- Export the metadata XML from the enterprise IdP and save it to AgentCore by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider).
//
// **Step 3:*	- After the configuration is complete, enable single sign-on logon by setting loginEnabled as needed.
//
// @param request - GetIdentityProviderServiceProviderRequest
//
// @return GetIdentityProviderServiceProviderResponse
func (client *Client) GetIdentityProviderServiceProvider(workspaceId *string, identityProviderType *string, request *GetIdentityProviderServiceProviderRequest) (_result *GetIdentityProviderServiceProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityProviderServiceProviderResponse{}
	_body, _err := client.GetIdentityProviderServiceProviderWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an identity provider synchronization task.
//
// Description:
//
// Queries the progress, number of processed users, and failure causes of a user synchronization task to track whether manual synchronization is complete.
//
// Use the `taskId` returned by [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) and keep the same workspace and identity provider.
//
// ### Determine synchronization results
//
// | `status` | Meaning and what to do next |
//
// | --- | --- |
//
// | `PENDING`, `RUNNING` | The node is not yet complete. You can continue to query. |
//
// | `PAUSED` | The node is paused and has not ended. |
//
// | `SUCCESS` | The synchronization is complete. You can view the workspace users. |
//
// | `FAILED` | The synchronization failed. Check `failureReason`. |
//
// | `CANCELED` | The node is canceled. |
//
// `syncedUserCount` is the number of processed users, including updated existing users. It is not the number of newly added users and cannot be directly used as a completion percentage. `null` is returned when the count is not yet available.
//
// @param request - GetIdentityProviderSyncJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderSyncJobResponse
func (client *Client) GetIdentityProviderSyncJobWithOptions(workspaceId *string, identityProviderType *string, taskId *string, request *GetIdentityProviderSyncJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderSyncJobResponse, _err error) {
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
		Action:      dara.String("GetIdentityProviderSyncJob"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType)) + "/sync-jobs/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an identity provider synchronization task.
//
// Description:
//
// Queries the progress, number of processed users, and failure causes of a user synchronization task to track whether manual synchronization is complete.
//
// Use the `taskId` returned by [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) and keep the same workspace and identity provider.
//
// ### Determine synchronization results
//
// | `status` | Meaning and what to do next |
//
// | --- | --- |
//
// | `PENDING`, `RUNNING` | The node is not yet complete. You can continue to query. |
//
// | `PAUSED` | The node is paused and has not ended. |
//
// | `SUCCESS` | The synchronization is complete. You can view the workspace users. |
//
// | `FAILED` | The synchronization failed. Check `failureReason`. |
//
// | `CANCELED` | The node is canceled. |
//
// `syncedUserCount` is the number of processed users, including updated existing users. It is not the number of newly added users and cannot be directly used as a completion percentage. `null` is returned when the count is not yet available.
//
// @param request - GetIdentityProviderSyncJobRequest
//
// @return GetIdentityProviderSyncJobResponse
func (client *Client) GetIdentityProviderSyncJob(workspaceId *string, identityProviderType *string, taskId *string, request *GetIdentityProviderSyncJobRequest) (_result *GetIdentityProviderSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityProviderSyncJobResponse{}
	_body, _err := client.GetIdentityProviderSyncJobWithOptions(workspaceId, identityProviderType, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the synchronization scope of an identity provider.
//
// Description:
//
// Queries the saved user synchronization scope, checks whether the selected departments are valid, and views the estimated number of departments and users involved. You can use this operation to verify the scope before initiating a manual synchronization.
//
// ### View synchronization scope
//
// | `syncScope.mode` | Synchronization scope |
//
// | --- | --- |
//
// | `AllAuthorized` | The address book scope authorized for the enterprise application. The specific scope is determined by the enterprise\\"s authorization. |
//
// | `SelectedDepartments` | The saved scope of specified departments. Use `selectedDepartments` to view the selected departments and their validity. |
//
// This operation reads the saved configuration. After adjusting the department selection, save the changes by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) before querying the estimated results. To browse the complete department tree, use [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments).
//
// ### Use user count estimation
//
// You can use `expandedDepartmentCount` and `estimatedUserCount` only when `estimateStatus=Ready` and the counts are not `null`. Counts that have not been obtained return `null` and must not be treated as 0.
//
// After confirming the scope, initiate the synchronization by calling [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob).
//
// @param request - GetIdentityProviderSyncScopeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderSyncScopeResponse
func (client *Client) GetIdentityProviderSyncScopeWithOptions(workspaceId *string, identityProviderType *string, request *GetIdentityProviderSyncScopeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderSyncScopeResponse, _err error) {
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
		Action:      dara.String("GetIdentityProviderSyncScope"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType)) + "/sync-scope"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderSyncScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the synchronization scope of an identity provider.
//
// Description:
//
// Queries the saved user synchronization scope, checks whether the selected departments are valid, and views the estimated number of departments and users involved. You can use this operation to verify the scope before initiating a manual synchronization.
//
// ### View synchronization scope
//
// | `syncScope.mode` | Synchronization scope |
//
// | --- | --- |
//
// | `AllAuthorized` | The address book scope authorized for the enterprise application. The specific scope is determined by the enterprise\\"s authorization. |
//
// | `SelectedDepartments` | The saved scope of specified departments. Use `selectedDepartments` to view the selected departments and their validity. |
//
// This operation reads the saved configuration. After adjusting the department selection, save the changes by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider) before querying the estimated results. To browse the complete department tree, use [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments).
//
// ### Use user count estimation
//
// You can use `expandedDepartmentCount` and `estimatedUserCount` only when `estimateStatus=Ready` and the counts are not `null`. Counts that have not been obtained return `null` and must not be treated as 0.
//
// After confirming the scope, initiate the synchronization by calling [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob).
//
// @param request - GetIdentityProviderSyncScopeRequest
//
// @return GetIdentityProviderSyncScopeResponse
func (client *Client) GetIdentityProviderSyncScope(workspaceId *string, identityProviderType *string, request *GetIdentityProviderSyncScopeRequest) (_result *GetIdentityProviderSyncScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityProviderSyncScopeResponse{}
	_body, _err := client.GetIdentityProviderSyncScopeWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status, digest, size, and deletion conditions of an image cache in a specified workspace.
//
// Description:
//
// Queries the creation status and usage of an image cache to determine whether the cache is available and meets the deletion conditions.
//
// ### Cache status
//
// | `status` | Meaning |
//
// | --- | --- |
//
// | `PENDING` | Preparing. Continue to query. |
//
// | `READY` | Ready. |
//
// | `FAILED` | Preparation failed. Troubleshoot based on the `error` field. |
//
// | `DELETING` | Deleting. |
//
// | `DISABLED` | Disabled. |
//
// The `digest` is provided after the source resolution is completed. The unit of `cacheSizeBytes` is bytes, and it can be empty when the cache is not yet available.
//
// ### Usage and deletion
//
// `agentCount` indicates the number of agents that have enabled the cache switch and reference this cache. `canDelete` and `deleteBlockReason` indicate whether the cache can be deleted and the reason for blocking deletion. Deletion requests will recheck the usage.
//
// The query scope of this operation is determined by `regionId`, `workspaceId`, and `imageId`.
//
// @param request - GetImageCacheRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageCacheResponse
func (client *Client) GetImageCacheWithOptions(imageId *string, request *GetImageCacheRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageCache"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/image-caches/" + dara.PercentEncode(dara.StringValue(imageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status, digest, size, and deletion conditions of an image cache in a specified workspace.
//
// Description:
//
// Queries the creation status and usage of an image cache to determine whether the cache is available and meets the deletion conditions.
//
// ### Cache status
//
// | `status` | Meaning |
//
// | --- | --- |
//
// | `PENDING` | Preparing. Continue to query. |
//
// | `READY` | Ready. |
//
// | `FAILED` | Preparation failed. Troubleshoot based on the `error` field. |
//
// | `DELETING` | Deleting. |
//
// | `DISABLED` | Disabled. |
//
// The `digest` is provided after the source resolution is completed. The unit of `cacheSizeBytes` is bytes, and it can be empty when the cache is not yet available.
//
// ### Usage and deletion
//
// `agentCount` indicates the number of agents that have enabled the cache switch and reference this cache. `canDelete` and `deleteBlockReason` indicate whether the cache can be deleted and the reason for blocking deletion. Deletion requests will recheck the usage.
//
// The query scope of this operation is determined by `regionId`, `workspaceId`, and `imageId`.
//
// @param request - GetImageCacheRequest
//
// @return GetImageCacheResponse
func (client *Client) GetImageCache(imageId *string, request *GetImageCacheRequest) (_result *GetImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageCacheResponse{}
	_body, _err := client.GetImageCacheWithOptions(imageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest configuration, Harness and skill versions, configuration processing status, and running instance statistics of a managed Harness agent.
//
// Description:
//
// Queries the configuration, skill versions, and running status of a managed Harness agent to view the current configuration and its processing results.
//
// ### Configuration and running status
//
// This operation returns the latest submitted configuration, which may still be processing. Use the following fields to determine the update progress and running status:
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest submitted desired configuration version. |
//
// | `latestVersionStatus` | The processing status of the desired version. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | The current snapshot of the number of instances in each phase. |
//
// Successful configuration processing does not mean that all old instances have been replaced. To view specific instances or active sessions, call [ListSandboxes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxes) and [ListSandboxSessions](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxSessions) respectively.
//
// ### Harness and skill versions
//
// `harness.type` indicates the runtime framework determined at creation. For qodercli, `connectorServiceAccountKey` returns the ID of the bound account key record and does not return the original key.
//
// `configuredSkills` refers to explicitly configured skills. `skills` includes the aggregated results of template and explicit configurations. `resolvedVersion` is the current target version, and `appliedVersion` is the version currently in effect at runtime. These two may differ during an update. `fromTemplate=true` indicates that the skill comes from a template.
//
// If you query an agent that does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
//
// @param request - GetManagedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedAgentResponse
func (client *Client) GetManagedAgentWithOptions(workspaceId *string, agentId *string, request *GetManagedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetManagedAgentResponse, _err error) {
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
		Action:      dara.String("GetManagedAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest configuration, Harness and skill versions, configuration processing status, and running instance statistics of a managed Harness agent.
//
// Description:
//
// Queries the configuration, skill versions, and running status of a managed Harness agent to view the current configuration and its processing results.
//
// ### Configuration and running status
//
// This operation returns the latest submitted configuration, which may still be processing. Use the following fields to determine the update progress and running status:
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest submitted desired configuration version. |
//
// | `latestVersionStatus` | The processing status of the desired version. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | The current snapshot of the number of instances in each phase. |
//
// Successful configuration processing does not mean that all old instances have been replaced. To view specific instances or active sessions, call [ListSandboxes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxes) and [ListSandboxSessions](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxSessions) respectively.
//
// ### Harness and skill versions
//
// `harness.type` indicates the runtime framework determined at creation. For qodercli, `connectorServiceAccountKey` returns the ID of the bound account key record and does not return the original key.
//
// `configuredSkills` refers to explicitly configured skills. `skills` includes the aggregated results of template and explicit configurations. `resolvedVersion` is the current target version, and `appliedVersion` is the version currently in effect at runtime. These two may differ during an update. `fromTemplate=true` indicates that the skill comes from a template.
//
// If you query an agent that does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
//
// @param request - GetManagedAgentRequest
//
// @return GetManagedAgentResponse
func (client *Client) GetManagedAgent(workspaceId *string, agentId *string, request *GetManagedAgentRequest) (_result *GetManagedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetManagedAgentResponse{}
	_body, _err := client.GetManagedAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified MCP service, such as the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// Queries the details of a specified MCP service, such as the address, type, status, authentication configuration, and protocol.
//
// @param request - GetMcpRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpResponse
func (client *Client) GetMcpWithOptions(workspaceId *string, mcpServerId *string, request *GetMcpRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMcpResponse, _err error) {
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
		Action:      dara.String("GetMcp"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified MCP service, such as the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// Queries the details of a specified MCP service, such as the address, type, status, authentication configuration, and protocol.
//
// @param request - GetMcpRequest
//
// @return GetMcpResponse
func (client *Client) GetMcp(workspaceId *string, mcpServerId *string, request *GetMcpRequest) (_result *GetMcpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMcpResponse{}
	_body, _err := client.GetMcpWithOptions(workspaceId, mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns a time-limited presigned upload URL and a temporary code package token. After the upload is complete, the token can be used to create or update a code deployment.
//
// Description:
//
// Retrieves a temporary upload URL for a local MCP code package to create or update an MCP service of the code deployment type.
//
// ### Upload and deployment steps
//
// **Step 1:*	- Prepare a ZIP code package, enter the file name in `fileName`, and call this operation to obtain the upload information.
//
// **Step 2:*	- Use the returned `uploadMethod`, `uploadUrl`, and `headers` to upload the file content. The package size cannot exceed `maxSize`. The current limit is 50 MiB (52,428,800 bytes).
//
// **Step 3:*	- After the upload is complete, submit the `codePackageToken` to the code deployment configuration in [CreateMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateMcp) or [UpdateMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateMcp).
//
// **Step 4:*	- After submitting the deployment, query the service status by calling [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp).
//
// ### Validity period
//
// For the expiration time of the upload URL, see `expiresAt`. For the expiration time of the temporary code package, see `artifactExpiresAt`. Complete the upload and deployment within the corresponding periods. If the upload URL expires, retrieve it again.
//
// @param tmpReq - GetMcpCodePackageUploadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpCodePackageUploadUrlResponse
func (client *Client) GetMcpCodePackageUploadUrlWithOptions(workspaceId *string, tmpReq *GetMcpCodePackageUploadUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMcpCodePackageUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMcpCodePackageUploadUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpCodePackageUploadUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-server-actions/get-code-package-upload-url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpCodePackageUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a time-limited presigned upload URL and a temporary code package token. After the upload is complete, the token can be used to create or update a code deployment.
//
// Description:
//
// Retrieves a temporary upload URL for a local MCP code package to create or update an MCP service of the code deployment type.
//
// ### Upload and deployment steps
//
// **Step 1:*	- Prepare a ZIP code package, enter the file name in `fileName`, and call this operation to obtain the upload information.
//
// **Step 2:*	- Use the returned `uploadMethod`, `uploadUrl`, and `headers` to upload the file content. The package size cannot exceed `maxSize`. The current limit is 50 MiB (52,428,800 bytes).
//
// **Step 3:*	- After the upload is complete, submit the `codePackageToken` to the code deployment configuration in [CreateMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateMcp) or [UpdateMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateMcp).
//
// **Step 4:*	- After submitting the deployment, query the service status by calling [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp).
//
// ### Validity period
//
// For the expiration time of the upload URL, see `expiresAt`. For the expiration time of the temporary code package, see `artifactExpiresAt`. Complete the upload and deployment within the corresponding periods. If the upload URL expires, retrieve it again.
//
// @param request - GetMcpCodePackageUploadUrlRequest
//
// @return GetMcpCodePackageUploadUrlResponse
func (client *Client) GetMcpCodePackageUploadUrl(workspaceId *string, request *GetMcpCodePackageUploadUrlRequest) (_result *GetMcpCodePackageUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMcpCodePackageUploadUrlResponse{}
	_body, _err := client.GetMcpCodePackageUploadUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an MCP marketplace template.
//
// Description:
//
// Queries the capabilities, parameters, and version information of an MCP marketplace template to help you select a template and prepare for installation and configuration.
//
// Returns the current template version and the installation form schema.
//
// @param request - GetMcpMarketItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpMarketItemResponse
func (client *Client) GetMcpMarketItemWithOptions(workspaceId *string, marketItemId *string, request *GetMcpMarketItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMcpMarketItemResponse, _err error) {
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
		Action:      dara.String("GetMcpMarketItem"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-market/items/" + dara.PercentEncode(dara.StringValue(marketItemId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpMarketItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an MCP marketplace template.
//
// Description:
//
// Queries the capabilities, parameters, and version information of an MCP marketplace template to help you select a template and prepare for installation and configuration.
//
// Returns the current template version and the installation form schema.
//
// @param request - GetMcpMarketItemRequest
//
// @return GetMcpMarketItemResponse
func (client *Client) GetMcpMarketItem(workspaceId *string, marketItemId *string, request *GetMcpMarketItemRequest) (_result *GetMcpMarketItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMcpMarketItemResponse{}
	_body, _err := client.GetMcpMarketItemWithOptions(workspaceId, marketItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a long-term memory by memory ID.
//
// Description:
//
// Queries a long-term memory by memory ID.
//
// @param request - GetMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspaceId *string, memoryStoreName *string, memoryId *string, request *GetMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
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
		Action:      dara.String("GetMemory"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a long-term memory by memory ID.
//
// Description:
//
// Queries a long-term memory by memory ID.
//
// @param request - GetMemoryRequest
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspaceId *string, memoryStoreName *string, memoryId *string, request *GetMemoryRequest) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspaceId, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the public information and current status of a MemoryStore in a specified workspace.
//
// Description:
//
// Queries the configuration and current status of a long-term memory store (MemoryStore) for viewing memory store information.
//
// @param request - GetMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryStoreResponse
func (client *Client) GetMemoryStoreWithOptions(workspaceId *string, memoryStoreName *string, request *GetMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryStoreResponse, _err error) {
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
		Action:      dara.String("GetMemoryStore"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the public information and current status of a MemoryStore in a specified workspace.
//
// Description:
//
// Queries the configuration and current status of a long-term memory store (MemoryStore) for viewing memory store information.
//
// @param request - GetMemoryStoreRequest
//
// @return GetMemoryStoreResponse
func (client *Client) GetMemoryStore(workspaceId *string, memoryStoreName *string, request *GetMemoryStoreRequest) (_result *GetMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryStoreResponse{}
	_body, _err := client.GetMemoryStoreWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the task ID returned by the asynchronous AddMemories operation.
//
// Description:
//
// Queries the progress and results of a memory extraction task by using the `taskId` returned by the asynchronous [AddMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/AddMemories) operation.
//
// @param request - GetMemoryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryTaskResponse
func (client *Client) GetMemoryTaskWithOptions(workspaceId *string, memoryStoreName *string, taskId *string, request *GetMemoryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryTaskResponse, _err error) {
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
		Action:      dara.String("GetMemoryTask"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task ID returned by the asynchronous AddMemories operation.
//
// Description:
//
// Queries the progress and results of a memory extraction task by using the `taskId` returned by the asynchronous [AddMemories](https://api.aliyun.com/document/AgentCore/2026-08-04/AddMemories) operation.
//
// @param request - GetMemoryTaskRequest
//
// @return GetMemoryTaskResponse
func (client *Client) GetMemoryTask(workspaceId *string, memoryStoreName *string, taskId *string, request *GetMemoryTaskRequest) (_result *GetMemoryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryTaskResponse{}
	_body, _err := client.GetMemoryTaskWithOptions(workspaceId, memoryStoreName, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration and region of a model in a specified workspace.
//
// Description:
//
// Queries the detailed configuration and region of a model in a specified workspace.
//
// @param request - GetModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelResponse
func (client *Client) GetModelWithOptions(workspaceId *string, modelId *string, request *GetModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelResponse, _err error) {
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
		Action:      dara.String("GetModel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration and region of a model in a specified workspace.
//
// Description:
//
// Queries the detailed configuration and region of a model in a specified workspace.
//
// @param request - GetModelRequest
//
// @return GetModelResponse
func (client *Client) GetModel(workspaceId *string, modelId *string, request *GetModelRequest) (_result *GetModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelResponse{}
	_body, _err := client.GetModelWithOptions(workspaceId, modelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration, credential configuration status, publish status, and region of a specified model connection.
//
// Description:
//
// Queries the detailed configuration, credential configuration status, publish status, and region of a specified model connection.
//
// @param request - GetModelConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelConnectionResponse
func (client *Client) GetModelConnectionWithOptions(workspaceId *string, connectionId *string, request *GetModelConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelConnectionResponse, _err error) {
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
		Action:      dara.String("GetModelConnection"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/model-connections/" + dara.PercentEncode(dara.StringValue(connectionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration, credential configuration status, publish status, and region of a specified model connection.
//
// Description:
//
// Queries the detailed configuration, credential configuration status, publish status, and region of a specified model connection.
//
// @param request - GetModelConnectionRequest
//
// @return GetModelConnectionResponse
func (client *Client) GetModelConnection(workspaceId *string, connectionId *string, request *GetModelConnectionRequest) (_result *GetModelConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelConnectionResponse{}
	_body, _err := client.GetModelConnectionWithOptions(workspaceId, connectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overall result of a specified sandbox rebuild task. Always use the initial operationId returned by RebuildSandbox without tracking the rebuild task and subsequent creation tasks separately. Only the SUCCEEDED status indicates completion. Do not treat UNKNOWN, missing status, or query exceptions as success, and do not automatically resubmit the rebuild.
//
// Description:
//
// Always use the initial operationId returned by RebuildSandbox for queries. Magic aggregates and returns the rebuildOutcome result from Runtime Manager, so the frontend does not need to track Claim tasks. Only SUCCEEDED indicates overall completion. UNKNOWN, missing fields, or query exceptions must not be treated as success, and the rebuild must not be automatically resubmitted. You can specify the endpoint runtime version by using serviceEndpointId. If omitted, the default latest runtime identity is used.
//
// @param request - GetSandboxRebuildOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSandboxRebuildOperationResponse
func (client *Client) GetSandboxRebuildOperationWithOptions(workspaceId *string, agentId *string, sandboxId *string, operationId *string, request *GetSandboxRebuildOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSandboxRebuildOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceEndpointId) {
		query["serviceEndpointId"] = request.ServiceEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSandboxRebuildOperation"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/rebuild-operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSandboxRebuildOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overall result of a specified sandbox rebuild task. Always use the initial operationId returned by RebuildSandbox without tracking the rebuild task and subsequent creation tasks separately. Only the SUCCEEDED status indicates completion. Do not treat UNKNOWN, missing status, or query exceptions as success, and do not automatically resubmit the rebuild.
//
// Description:
//
// Always use the initial operationId returned by RebuildSandbox for queries. Magic aggregates and returns the rebuildOutcome result from Runtime Manager, so the frontend does not need to track Claim tasks. Only SUCCEEDED indicates overall completion. UNKNOWN, missing fields, or query exceptions must not be treated as success, and the rebuild must not be automatically resubmitted. You can specify the endpoint runtime version by using serviceEndpointId. If omitted, the default latest runtime identity is used.
//
// @param request - GetSandboxRebuildOperationRequest
//
// @return GetSandboxRebuildOperationResponse
func (client *Client) GetSandboxRebuildOperation(workspaceId *string, agentId *string, sandboxId *string, operationId *string, request *GetSandboxRebuildOperationRequest) (_result *GetSandboxRebuildOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSandboxRebuildOperationResponse{}
	_body, _err := client.GetSandboxRebuildOperationWithOptions(workspaceId, agentId, sandboxId, operationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified service endpoint, including the target route, access URLs, authentication configuration, and current status.
//
// Description:
//
// Queries the service endpoint of an agent or collaboration component to obtain stable access URLs, authentication configuration, and current status.
//
// A service endpoint associates access URLs with a specified agent version or workspace collaboration component. You can use the target route and URL list in the response to verify the access target.
//
// @param request - GetServiceEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEndpointResponse
func (client *Client) GetServiceEndpointWithOptions(workspaceId *string, serviceEndpointId *string, request *GetServiceEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceEndpointResponse, _err error) {
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
		Action:      dara.String("GetServiceEndpoint"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/service-endpoints/" + dara.PercentEncode(dara.StringValue(serviceEndpointId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified service endpoint, including the target route, access URLs, authentication configuration, and current status.
//
// Description:
//
// Queries the service endpoint of an agent or collaboration component to obtain stable access URLs, authentication configuration, and current status.
//
// A service endpoint associates access URLs with a specified agent version or workspace collaboration component. You can use the target route and URL list in the response to verify the access target.
//
// @param request - GetServiceEndpointRequest
//
// @return GetServiceEndpointResponse
func (client *Client) GetServiceEndpoint(workspaceId *string, serviceEndpointId *string, request *GetServiceEndpointRequest) (_result *GetServiceEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceEndpointResponse{}
	_body, _err := client.GetServiceEndpointWithOptions(workspaceId, serviceEndpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the active API key and its invocation method for a service endpoint in a specified workspace.
//
// Description:
//
// Retrieves the active API key and invocation method of a service endpoint for external applications to access the endpoint.
//
// The endpoint must have `API_KEY` authentication enabled and access credentials ready. Securely store the returned API key.
//
// @param request - GetServiceEndpointApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEndpointApiKeyResponse
func (client *Client) GetServiceEndpointApiKeyWithOptions(workspaceId *string, serviceEndpointId *string, request *GetServiceEndpointApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceEndpointApiKeyResponse, _err error) {
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
		Action:      dara.String("GetServiceEndpointApiKey"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/service-endpoints/" + dara.PercentEncode(dara.StringValue(serviceEndpointId)) + "/api-key/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceEndpointApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the active API key and its invocation method for a service endpoint in a specified workspace.
//
// Description:
//
// Retrieves the active API key and invocation method of a service endpoint for external applications to access the endpoint.
//
// The endpoint must have `API_KEY` authentication enabled and access credentials ready. Securely store the returned API key.
//
// @param request - GetServiceEndpointApiKeyRequest
//
// @return GetServiceEndpointApiKeyResponse
func (client *Client) GetServiceEndpointApiKey(workspaceId *string, serviceEndpointId *string, request *GetServiceEndpointApiKeyRequest) (_result *GetServiceEndpointApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceEndpointApiKeyResponse{}
	_body, _err := client.GetServiceEndpointApiKeyWithOptions(workspaceId, serviceEndpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information, version status, tags, visibility scope, and version list of a specified skill.
//
// Description:
//
// Queries the basic information, version status, tags, visibility scope, and version list of a specified skill.
//
// @param request - GetSkillDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetailWithOptions(workspaceId *string, skillName *string, request *GetSkillDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSkillDetailResponse, _err error) {
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
		Action:      dara.String("GetSkillDetail"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information, version status, tags, visibility scope, and version list of a specified skill.
//
// Description:
//
// Queries the basic information, version status, tags, visibility scope, and version list of a specified skill.
//
// @param request - GetSkillDetailRequest
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetail(workspaceId *string, skillName *string, request *GetSkillDetailRequest) (_result *GetSkillDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSkillDetailResponse{}
	_body, _err := client.GetSkillDetailWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the Object Storage Service (OSS) presigned upload URL and object name required to import a skill ZIP package. After the upload is complete, you can call the skill OSS upload API operation to complete the import.
//
// Description:
//
// Retrieves the upload URL for a skill ZIP package to import a skill from a local file.
//
// Files are transferred through OSS. Upload the file by using the returned URL, and then call [UploadSkillViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadSkillViaOss) to complete the import. To check the package content and conflicts, call [PrecheckSkillUploadViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/PrecheckSkillUploadViaOss) first.
//
// @param request - GetSkillImportFileUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillImportFileUrlResponse
func (client *Client) GetSkillImportFileUrlWithOptions(workspaceId *string, request *GetSkillImportFileUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSkillImportFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		query["contentType"] = request.ContentType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillImportFileUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skill-actions/get-import-file-url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the Object Storage Service (OSS) presigned upload URL and object name required to import a skill ZIP package. After the upload is complete, you can call the skill OSS upload API operation to complete the import.
//
// Description:
//
// Retrieves the upload URL for a skill ZIP package to import a skill from a local file.
//
// Files are transferred through OSS. Upload the file by using the returned URL, and then call [UploadSkillViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadSkillViaOss) to complete the import. To check the package content and conflicts, call [PrecheckSkillUploadViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/PrecheckSkillUploadViaOss) first.
//
// @param request - GetSkillImportFileUrlRequest
//
// @return GetSkillImportFileUrlResponse
func (client *Client) GetSkillImportFileUrl(workspaceId *string, request *GetSkillImportFileUrlRequest) (_result *GetSkillImportFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSkillImportFileUrlResponse{}
	_body, _err := client.GetSkillImportFileUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the complete details of a specified skill version, including version metadata, skill content, and associated resources.
//
// Description:
//
// Queries the complete details of a specified skill version, including version metadata, skill content, and associated resources.
//
// @param request - GetSkillVersionDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillVersionDetailResponse
func (client *Client) GetSkillVersionDetailWithOptions(workspaceId *string, skillName *string, skillVersion *string, request *GetSkillVersionDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSkillVersionDetailResponse, _err error) {
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
		Action:      dara.String("GetSkillVersionDetail"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillVersionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the complete details of a specified skill version, including version metadata, skill content, and associated resources.
//
// Description:
//
// Queries the complete details of a specified skill version, including version metadata, skill content, and associated resources.
//
// @param request - GetSkillVersionDetailRequest
//
// @return GetSkillVersionDetailResponse
func (client *Client) GetSkillVersionDetail(workspaceId *string, skillName *string, skillVersion *string, request *GetSkillVersionDetailRequest) (_result *GetSkillVersionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSkillVersionDetailResponse{}
	_body, _err := client.GetSkillVersionDetailWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified team, including the complete attributes and team roles of each user and agent member in the team.
//
// Description:
//
// Queries the details of a specified team, including the complete attributes and team roles of each user and agent member in the team.
//
// @param request - GetTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTeamResponse
func (client *Client) GetTeamWithOptions(workspaceId *string, teamId *string, request *GetTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTeamResponse, _err error) {
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
		Action:      dara.String("GetTeam"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/teams/" + dara.PercentEncode(dara.StringValue(teamId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified team, including the complete attributes and team roles of each user and agent member in the team.
//
// Description:
//
// Queries the details of a specified team, including the complete attributes and team roles of each user and agent member in the team.
//
// @param request - GetTeamRequest
//
// @return GetTeamResponse
func (client *Client) GetTeam(workspaceId *string, teamId *string, request *GetTeamRequest) (_result *GetTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTeamResponse{}
	_body, _err := client.GetTeamWithOptions(workspaceId, teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific user in a specified workspace. Returns an error if the user does not exist.
//
// Description:
//
// Queries the details of a specific user in a specified workspace.
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(workspaceId *string, agentCoreUserId *string, request *GetUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
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
		Action:      dara.String("GetUser"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users/" + dara.PercentEncode(dara.StringValue(agentCoreUserId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific user in a specified workspace. Returns an error if the user does not exist.
//
// Description:
//
// Queries the details of a specific user in a specified workspace.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(workspaceId *string, agentCoreUserId *string, request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(workspaceId, agentCoreUserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS workspace, AI registry namespace, and current network policies.
//
// Description:
//
// Queries the basic information, associated resources, and network configurations of a workspace for status viewing or configuration verification.
//
// The returned information includes the lifecycle status, associated observability workspace, AI registry namespace, and current network policies.
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceId *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS workspace, AI registry namespace, and current network policies.
//
// Description:
//
// Queries the basic information, associated resources, and network configurations of a workspace for status viewing or configuration verification.
//
// The returned information includes the lifecycle status, associated observability workspace, AI registry namespace, and current network policies.
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceId *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a RAM authorization link for the workspace resource role to pull images from a specified ACR Enterprise Edition repository.
//
// Description:
//
// Generates an authorization link for the workspace resource role to pull images from a specified ACR Enterprise Edition repository, allowing you to complete access authorization before image deployment.
//
// ### Request parameters
//
// The `workspaceId` parameter specifies the workspace. The `acrInstanceId`, `namespace`, and `repository` parameters are passed as query parameters. The `repository` parameter specifies the repository name without the namespace.
//
// ### Complete authorization
//
// Open the `authorizeUrl` in the response and confirm the authorization. Then, call [VerifyWorkspaceAcrRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyWorkspaceAcrRamAuthorization) to check the binding status. The role is selected by the server. The `roleName` and `roleSource` parameters indicate the selected role and its source, but do not indicate that the authorization is complete.
//
// This authorization is used for ACR image deployment of custom agents and reading image cache sources. Obtaining the link itself does not pull images or create caches.
//
// @param request - GetWorkspaceAcrRamAuthorizeUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceAcrRamAuthorizeUrlResponse
func (client *Client) GetWorkspaceAcrRamAuthorizeUrlWithOptions(workspaceId *string, request *GetWorkspaceAcrRamAuthorizeUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceAcrRamAuthorizeUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		query["acrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Repository) {
		query["repository"] = request.Repository
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceAcrRamAuthorizeUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/acr/authorize"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceAcrRamAuthorizeUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a RAM authorization link for the workspace resource role to pull images from a specified ACR Enterprise Edition repository.
//
// Description:
//
// Generates an authorization link for the workspace resource role to pull images from a specified ACR Enterprise Edition repository, allowing you to complete access authorization before image deployment.
//
// ### Request parameters
//
// The `workspaceId` parameter specifies the workspace. The `acrInstanceId`, `namespace`, and `repository` parameters are passed as query parameters. The `repository` parameter specifies the repository name without the namespace.
//
// ### Complete authorization
//
// Open the `authorizeUrl` in the response and confirm the authorization. Then, call [VerifyWorkspaceAcrRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyWorkspaceAcrRamAuthorization) to check the binding status. The role is selected by the server. The `roleName` and `roleSource` parameters indicate the selected role and its source, but do not indicate that the authorization is complete.
//
// This authorization is used for ACR image deployment of custom agents and reading image cache sources. Obtaining the link itself does not pull images or create caches.
//
// @param request - GetWorkspaceAcrRamAuthorizeUrlRequest
//
// @return GetWorkspaceAcrRamAuthorizeUrlResponse
func (client *Client) GetWorkspaceAcrRamAuthorizeUrl(workspaceId *string, request *GetWorkspaceAcrRamAuthorizeUrlRequest) (_result *GetWorkspaceAcrRamAuthorizeUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceAcrRamAuthorizeUrlResponse{}
	_body, _err := client.GetWorkspaceAcrRamAuthorizeUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the RAM authorization page URL to grant the workspace resource role access to a specified AgenticFS access point before configuring NAS storage mounting for an agent.
//
// Description:
//
// Retrieves the AgenticFS mount authorization link to enable agents in the workspace to access a specified file space.
//
// [AgenticFS](https://www.alibabacloud.com/help/en/nas/product-overview/agenticfs-overview) is a storage type of Apsara File Storage NAS designed for AI agents. It can be used to persistently store files used by agents. AgenticSpace is the file space within it, and an access point provides the mount entry.
//
// ### Authorization and mounting steps
//
// **Step 1:*	- Prepare a file system, an AgenticSpace, and an access point in NAS, and obtain the file system ID, access point ID, and domain name.
//
// **Step 2:*	- Submit `fileSystemId`, `accessPointId`, and `server` as query parameters. For `server`, enter the access point domain name returned by the NAS `ListAccessPoints` operation, without the protocol, port, or path.
//
// **Step 3:*	- Open the returned `authorizeUrl` to complete the authorization for the workspace resource role.
//
// **Step 4:*	- Check the authorization status by calling [VerifyWorkspaceAgenticFsMountRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyWorkspaceAgenticFsMountRamAuthorization).
//
// **Step 5:*	- When creating or updating a custom code or image agent, or a managed Harness agent, enter the access point domain name, the relative directory `path`, and the container mount path `mountPath` in `agenticFsMounts`.
//
// The access point must be located in the VPC associated with the workspace, and VPC access must be enabled for the agent. The mount directory must already exist. Use a non-empty relative directory within the access point for `path`, and use an absolute path within the container for `mountPath`.
//
// @param request - GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse
func (client *Client) GetWorkspaceAgenticFsMountRamAuthorizeUrlWithOptions(workspaceId *string, request *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["accessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["fileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Server) {
		query["server"] = request.Server
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceAgenticFsMountRamAuthorizeUrl"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agentic-fs/authorize"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the RAM authorization page URL to grant the workspace resource role access to a specified AgenticFS access point before configuring NAS storage mounting for an agent.
//
// Description:
//
// Retrieves the AgenticFS mount authorization link to enable agents in the workspace to access a specified file space.
//
// [AgenticFS](https://www.alibabacloud.com/help/en/nas/product-overview/agenticfs-overview) is a storage type of Apsara File Storage NAS designed for AI agents. It can be used to persistently store files used by agents. AgenticSpace is the file space within it, and an access point provides the mount entry.
//
// ### Authorization and mounting steps
//
// **Step 1:*	- Prepare a file system, an AgenticSpace, and an access point in NAS, and obtain the file system ID, access point ID, and domain name.
//
// **Step 2:*	- Submit `fileSystemId`, `accessPointId`, and `server` as query parameters. For `server`, enter the access point domain name returned by the NAS `ListAccessPoints` operation, without the protocol, port, or path.
//
// **Step 3:*	- Open the returned `authorizeUrl` to complete the authorization for the workspace resource role.
//
// **Step 4:*	- Check the authorization status by calling [VerifyWorkspaceAgenticFsMountRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyWorkspaceAgenticFsMountRamAuthorization).
//
// **Step 5:*	- When creating or updating a custom code or image agent, or a managed Harness agent, enter the access point domain name, the relative directory `path`, and the container mount path `mountPath` in `agenticFsMounts`.
//
// The access point must be located in the VPC associated with the workspace, and VPC access must be enabled for the agent. The mount directory must already exist. Use a non-empty relative directory within the access point for `path`, and use an absolute path within the container for `mountPath`.
//
// @param request - GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest
//
// @return GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse
func (client *Client) GetWorkspaceAgenticFsMountRamAuthorizeUrl(workspaceId *string, request *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) (_result *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse{}
	_body, _err := client.GetWorkspaceAgenticFsMountRamAuthorizeUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the activation status and configuration of a specified component in a workspace. The collaboration value corresponds to the Collaboration component displayed in the console.
//
// Description:
//
// Queries the status and configuration of the workspace collaboration component to check whether the component is activated and whether the configuration takes effect.
//
// Use `pluginName=collaboration` to specify the collaboration component. To activate the component, call [InstallWorkspacePlugin](https://api.aliyun.com/document/AgentCore/2026-08-04/InstallWorkspacePlugin).
//
// @param request - GetWorkspacePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspacePluginResponse
func (client *Client) GetWorkspacePluginWithOptions(workspaceId *string, pluginName *string, request *GetWorkspacePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspacePluginResponse, _err error) {
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
		Action:      dara.String("GetWorkspacePlugin"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/plugins/" + dara.PercentEncode(dara.StringValue(pluginName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspacePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the activation status and configuration of a specified component in a workspace. The collaboration value corresponds to the Collaboration component displayed in the console.
//
// Description:
//
// Queries the status and configuration of the workspace collaboration component to check whether the component is activated and whether the configuration takes effect.
//
// Use `pluginName=collaboration` to specify the collaboration component. To activate the component, call [InstallWorkspacePlugin](https://api.aliyun.com/document/AgentCore/2026-08-04/InstallWorkspacePlugin).
//
// @param request - GetWorkspacePluginRequest
//
// @return GetWorkspacePluginResponse
func (client *Client) GetWorkspacePlugin(workspaceId *string, pluginName *string, request *GetWorkspacePluginRequest) (_result *GetWorkspacePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspacePluginResponse{}
	_body, _err := client.GetWorkspacePluginWithOptions(workspaceId, pluginName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an MCP marketplace template.
//
// Description:
//
// Creates an MCP service by using a marketplace template to reuse the tool capabilities and configurations provided by the template.
//
// First, call [GetMcpMarketItem](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcpMarketItem) to view the template versions and required parameters, and then submit the configuration based on the selected version. After the service is created, call [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) to query the service status.
//
// @param tmpReq - InstallMcpMarketItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallMcpMarketItemResponse
func (client *Client) InstallMcpMarketItemWithOptions(workspaceId *string, marketItemId *string, tmpReq *InstallMcpMarketItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallMcpMarketItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallMcpMarketItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["templateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallMcpMarketItem"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-market/items/" + dara.PercentEncode(dara.StringValue(marketItemId)) + "/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallMcpMarketItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an MCP marketplace template.
//
// Description:
//
// Creates an MCP service by using a marketplace template to reuse the tool capabilities and configurations provided by the template.
//
// First, call [GetMcpMarketItem](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcpMarketItem) to view the template versions and required parameters, and then submit the configuration based on the selected version. After the service is created, call [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) to query the service status.
//
// @param request - InstallMcpMarketItemRequest
//
// @return InstallMcpMarketItemResponse
func (client *Client) InstallMcpMarketItem(workspaceId *string, marketItemId *string, request *InstallMcpMarketItemRequest) (_result *InstallMcpMarketItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallMcpMarketItemResponse{}
	_body, _err := client.InstallMcpMarketItemWithOptions(workspaceId, marketItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a component for a workspace. This operation corresponds to enabling the collaboration component in the console. After the component is enabled, you can query the component status and current configurations.
//
// Description:
//
// Enables the collaboration component for a workspace and configures initial settings such as the network.
//
// Set pluginName=collaboration to specify the collaboration component. The enabling process is executed asynchronously. After submission, call [GetWorkspacePlugin](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspacePlugin) to check the component status and processing result.
//
// @param tmpReq - InstallWorkspacePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallWorkspacePluginResponse
func (client *Client) InstallWorkspacePluginWithOptions(workspaceId *string, pluginName *string, tmpReq *InstallWorkspacePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallWorkspacePluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallWorkspacePluginShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallWorkspacePlugin"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/plugins/" + dara.PercentEncode(dara.StringValue(pluginName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallWorkspacePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a component for a workspace. This operation corresponds to enabling the collaboration component in the console. After the component is enabled, you can query the component status and current configurations.
//
// Description:
//
// Enables the collaboration component for a workspace and configures initial settings such as the network.
//
// Set pluginName=collaboration to specify the collaboration component. The enabling process is executed asynchronously. After submission, call [GetWorkspacePlugin](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspacePlugin) to check the component status and processing result.
//
// @param request - InstallWorkspacePluginRequest
//
// @return InstallWorkspacePluginResponse
func (client *Client) InstallWorkspacePlugin(workspaceId *string, pluginName *string, request *InstallWorkspacePluginRequest) (_result *InstallWorkspacePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallWorkspacePluginResponse{}
	_body, _err := client.InstallWorkspacePluginWithOptions(workspaceId, pluginName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the identity, name, and architecture of language runtime environments currently available for code package deployment by page.
//
// Description:
//
// Queries the available runtime environments for code package deployment of custom code or image agents, which is used to populate artifact.code.languageRuntime.
//
// ### Select a runtime environment
//
// Copy the languageRuntime value from the response exactly as is into the artifact.code.languageRuntime parameter of [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent). The catalog is dynamically provided by the current platform version, and internal base image addresses are not returned.
//
// ### Pagination rules
//
// If maxResults is omitted or set to 0, the default value is 10. If the value is greater than 100, it is processed as 100. The value cannot be negative. Omit nextToken for the first request. For subsequent requests, use the token returned in the previous response and keep the workspace unchanged. An empty token indicates that the query is complete. If the token becomes invalid due to changes in the runtime environment catalog, restart the query from the first page.
//
// Regular updates, restarts, and scale-outs of existing agents continue to use their saved runtime images. When you create an agent or switch the language runtime environment, the catalog currently returned by this operation prevails.
//
// @param request - ListAgentCodeRuntimesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentCodeRuntimesResponse
func (client *Client) ListAgentCodeRuntimesWithOptions(workspaceId *string, request *ListAgentCodeRuntimesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentCodeRuntimesResponse, _err error) {
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
		Action:      dara.String("ListAgentCodeRuntimes"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-code-runtimes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentCodeRuntimesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the identity, name, and architecture of language runtime environments currently available for code package deployment by page.
//
// Description:
//
// Queries the available runtime environments for code package deployment of custom code or image agents, which is used to populate artifact.code.languageRuntime.
//
// ### Select a runtime environment
//
// Copy the languageRuntime value from the response exactly as is into the artifact.code.languageRuntime parameter of [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent). The catalog is dynamically provided by the current platform version, and internal base image addresses are not returned.
//
// ### Pagination rules
//
// If maxResults is omitted or set to 0, the default value is 10. If the value is greater than 100, it is processed as 100. The value cannot be negative. Omit nextToken for the first request. For subsequent requests, use the token returned in the previous response and keep the workspace unchanged. An empty token indicates that the query is complete. If the token becomes invalid due to changes in the runtime environment catalog, restart the query from the first page.
//
// Regular updates, restarts, and scale-outs of existing agents continue to use their saved runtime images. When you create an agent or switch the language runtime environment, the catalog currently returned by this operation prevails.
//
// @param request - ListAgentCodeRuntimesRequest
//
// @return ListAgentCodeRuntimesResponse
func (client *Client) ListAgentCodeRuntimes(workspaceId *string, request *ListAgentCodeRuntimesRequest) (_result *ListAgentCodeRuntimesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentCodeRuntimesResponse{}
	_body, _err := client.ListAgentCodeRuntimesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IM channels for a specified agent.
//
// Description:
//
// Queries the IM channels for a specified agent.
//
// @param request - ListAgentIMChannelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentIMChannelsResponse
func (client *Client) ListAgentIMChannelsWithOptions(workspaceId *string, agentId *string, request *ListAgentIMChannelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentIMChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["channelType"] = request.ChannelType
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentIMChannels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentIMChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IM channels for a specified agent.
//
// Description:
//
// Queries the IM channels for a specified agent.
//
// @param request - ListAgentIMChannelsRequest
//
// @return ListAgentIMChannelsResponse
func (client *Client) ListAgentIMChannels(workspaceId *string, agentId *string, request *ListAgentIMChannelsRequest) (_result *ListAgentIMChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentIMChannelsResponse{}
	_body, _err := client.ListAgentIMChannelsWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query of AgentSpecs in a specified workspace, and supports name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// Performs a paging query of Agent templates in a specified workspace, and supports name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// @param request - ListAgentSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentSpecsResponse
func (client *Client) ListAgentSpecsWithOptions(workspaceId *string, request *ListAgentSpecsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentSpecsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpecName) {
		query["agentSpecName"] = request.AgentSpecName
	}

	if !dara.IsNil(request.BizTag) {
		query["bizTag"] = request.BizTag
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Owner) {
		query["owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNo) {
		query["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.WithCapabilities) {
		query["withCapabilities"] = request.WithCapabilities
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentSpecs"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query of AgentSpecs in a specified workspace, and supports name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// Performs a paging query of Agent templates in a specified workspace, and supports name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// @param request - ListAgentSpecsRequest
//
// @return ListAgentSpecsResponse
func (client *Client) ListAgentSpecs(workspaceId *string, request *ListAgentSpecsRequest) (_result *ListAgentSpecsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentSpecsResponse{}
	_body, _err := client.ListAgentSpecsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the teams to which specified agents belong. Specify the agent IDs to query by using agentIds. The response includes the membership information of each agent in its teams, such as the team ID, team name, and the role of the agent in the team.
//
// Description:
//
// Queries the teams to which specified agents belong.
//
// @param tmpReq - ListAgentTeamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentTeamsResponse
func (client *Client) ListAgentTeamsWithOptions(workspaceId *string, tmpReq *ListAgentTeamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentTeamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAgentTeamsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentTeams"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-team-memberships"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentTeamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the teams to which specified agents belong. Specify the agent IDs to query by using agentIds. The response includes the membership information of each agent in its teams, such as the team ID, team name, and the role of the agent in the team.
//
// Description:
//
// Queries the teams to which specified agents belong.
//
// @param request - ListAgentTeamsRequest
//
// @return ListAgentTeamsResponse
func (client *Client) ListAgentTeams(workspaceId *string, request *ListAgentTeamsRequest) (_result *ListAgentTeamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentTeamsResponse{}
	_body, _err := client.ListAgentTeamsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of agent versions.
//
// Description:
//
// Performs a paged query for saved fixed versions of a specified managed or custom agent. The paging results are returned in descending order of creation time and do not include the latest version under system maintenance. List items do not return the configuration. To view the configuration, invoke GetAgentVersion.
//
// @param request - ListAgentVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentVersionsResponse
func (client *Client) ListAgentVersionsWithOptions(workspaceId *string, agentId *string, request *ListAgentVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentVersionsResponse, _err error) {
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
		Action:      dara.String("ListAgentVersions"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of agent versions.
//
// Description:
//
// Performs a paged query for saved fixed versions of a specified managed or custom agent. The paging results are returned in descending order of creation time and do not include the latest version under system maintenance. List items do not return the configuration. To view the configuration, invoke GetAgentVersion.
//
// @param request - ListAgentVersionsRequest
//
// @return ListAgentVersionsResponse
func (client *Client) ListAgentVersions(workspaceId *string, agentId *string, request *ListAgentVersionsRequest) (_result *ListAgentVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentVersionsResponse{}
	_body, _err := client.ListAgentVersionsWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the official and enterprise models available for the enabled QoderCLI connector. The service retrieves models in real time by using the saved API key and organization ID, and deduplicates them by model ID.
//
// Description:
//
// Queries the official and enterprise models available for the enabled QoderCLI connector.
//
// Queries the list of models available for a specified connector. Pagination is supported.
//
// @param request - ListConnectorModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectorModelsResponse
func (client *Client) ListConnectorModelsWithOptions(workspaceId *string, connectorName *string, request *ListConnectorModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectorModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorKeyId) {
		query["connectorKeyId"] = request.ConnectorKeyId
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
		Action:      dara.String("ListConnectorModels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName)) + "/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectorModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the official and enterprise models available for the enabled QoderCLI connector. The service retrieves models in real time by using the saved API key and organization ID, and deduplicates them by model ID.
//
// Description:
//
// Queries the official and enterprise models available for the enabled QoderCLI connector.
//
// Queries the list of models available for a specified connector. Pagination is supported.
//
// @param request - ListConnectorModelsRequest
//
// @return ListConnectorModelsResponse
func (client *Client) ListConnectorModels(workspaceId *string, connectorName *string, request *ListConnectorModelsRequest) (_result *ListConnectorModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectorModelsResponse{}
	_body, _err := client.ListConnectorModelsWithOptions(workspaceId, connectorName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of connectors supported by a specified workspace, along with their status, the number of bound agents, and configurations. The current version returns the QoderCLI connector.
//
// Description:
//
// Queries the connectors supported by the workspace to view their activation status, configurations, and the number of bound agents.
//
// @param request - ListConnectorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectorsResponse
func (client *Client) ListConnectorsWithOptions(workspaceId *string, request *ListConnectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectorsResponse, _err error) {
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
		Action:      dara.String("ListConnectors"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of connectors supported by a specified workspace, along with their status, the number of bound agents, and configurations. The current version returns the QoderCLI connector.
//
// Description:
//
// Queries the connectors supported by the workspace to view their activation status, configurations, and the number of bound agents.
//
// @param request - ListConnectorsRequest
//
// @return ListConnectorsResponse
func (client *Client) ListConnectors(workspaceId *string, request *ListConnectorsRequest) (_result *ListConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectorsResponse{}
	_body, _err := client.ListConnectorsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query for credentials in a specified workspace. Filters credentials by type by using credentialType, performs fuzzy filtering by credential name by using nameLike, specifies the maximum number of records per page for paging by using maxResults, and retrieves the next page by using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Performs a paged query for credentials in a specified workspace.
//
// ### Query and paging
//
// Filters credentials by `apiKey` or `mcpHeader` type by using `credentialType`, and performs fuzzy filtering by name by using `nameLike`.
//
// Uses `maxResults` to specify the number of entries per page for paging, and uses `nextToken` in the response to retrieve the next page. The returned credential content is desensitized. To query the details of a single credential, call [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential).
//
// @param request - ListCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithOptions(workspaceId *string, request *ListCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialType) {
		query["credentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentials"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/credentials"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for credentials in a specified workspace. Filters credentials by type by using credentialType, performs fuzzy filtering by credential name by using nameLike, specifies the maximum number of records per page for paging by using maxResults, and retrieves the next page by using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Performs a paged query for credentials in a specified workspace.
//
// ### Query and paging
//
// Filters credentials by `apiKey` or `mcpHeader` type by using `credentialType`, and performs fuzzy filtering by name by using `nameLike`.
//
// Uses `maxResults` to specify the number of entries per page for paging, and uses `nextToken` in the response to retrieve the next page. The returned credential content is desensitized. To query the details of a single credential, call [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential).
//
// @param request - ListCredentialsRequest
//
// @return ListCredentialsResponse
func (client *Client) ListCredentials(workspaceId *string, request *ListCredentialsRequest) (_result *ListCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCredentialsResponse{}
	_body, _err := client.ListCredentialsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve the identity, deployment method, and status summary of custom code or image agents in a workspace, with paging support.
//
// Description:
//
// Performs a paged query to retrieve the identity, deployment method, and status summary of custom code or image agents in a workspace, with paging support.
//
// ### Paged query
//
// If `maxResults` is omitted or set to `0`, the default value is 20 and the maximum value is 100. Omit `nextToken` in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates the end of paging.
//
// ### Response content
//
// This operation returns the summary for the current page. It does not return the complete deployment configurations or errors that occur when you execute the latest configurations. Use `items[].agentId` to invoke [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query code packages, image sources, results after you execute the version, and running instance statuses.
//
// @param request - ListCustomAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentsResponse
func (client *Client) ListCustomAgentsWithOptions(workspaceId *string, request *ListCustomAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCustomAgentsResponse, _err error) {
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
		Action:      dara.String("ListCustomAgents"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve the identity, deployment method, and status summary of custom code or image agents in a workspace, with paging support.
//
// Description:
//
// Performs a paged query to retrieve the identity, deployment method, and status summary of custom code or image agents in a workspace, with paging support.
//
// ### Paged query
//
// If `maxResults` is omitted or set to `0`, the default value is 20 and the maximum value is 100. Omit `nextToken` in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates the end of paging.
//
// ### Response content
//
// This operation returns the summary for the current page. It does not return the complete deployment configurations or errors that occur when you execute the latest configurations. Use `items[].agentId` to invoke [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query code packages, image sources, results after you execute the version, and running instance statuses.
//
// @param request - ListCustomAgentsRequest
//
// @return ListCustomAgentsResponse
func (client *Client) ListCustomAgents(workspaceId *string, request *ListCustomAgentsRequest) (_result *ListCustomAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCustomAgentsResponse{}
	_body, _err := client.ListCustomAgentsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve the identity, configuration versions, and status summary of managed agents in a workspace.
//
// Description:
//
// Performs a paging query to retrieve the identity, configuration versions, and status summary of managed agents in a workspace.
//
// ### Paging query
//
// If `maxResults` is omitted or set to `0`, the default value is 20, and the maximum value is 100. Omit `nextToken` in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates that there are no more pages.
//
// ### Response content
//
// This operation returns the summary of managed agents. To query the complete models, instructions, tools, skills, and accessed state, call [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent).
//
// @param request - ListExternalAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalAgentsResponse
func (client *Client) ListExternalAgentsWithOptions(workspaceId *string, request *ListExternalAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExternalAgentsResponse, _err error) {
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
		Action:      dara.String("ListExternalAgents"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve the identity, configuration versions, and status summary of managed agents in a workspace.
//
// Description:
//
// Performs a paging query to retrieve the identity, configuration versions, and status summary of managed agents in a workspace.
//
// ### Paging query
//
// If `maxResults` is omitted or set to `0`, the default value is 20, and the maximum value is 100. Omit `nextToken` in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates that there are no more pages.
//
// ### Response content
//
// This operation returns the summary of managed agents. To query the complete models, instructions, tools, skills, and accessed state, call [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent).
//
// @param request - ListExternalAgentsRequest
//
// @return ListExternalAgentsResponse
func (client *Client) ListExternalAgents(workspaceId *string, request *ListExternalAgentsRequest) (_result *ListExternalAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExternalAgentsResponse{}
	_body, _err := client.ListExternalAgentsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the departments of an identity provider by paging.
//
// Description:
//
// Queries the list of DingTalk or Lark departments by parent department and paging. This is used when selecting or adjusting the user synchronization scope.
//
// ### Before you begin
//
// A DingTalk or Lark identity provider is bound to the workspace, and the user pool, application credentials, and authorization to read contacts are prepared.
//
// ### Query methods
//
// - Omit `parentDepartmentId` to query the root level. Pass a department ID to query its direct child departments.
//
// - `hasChildren` indicates whether a department can be expanded. After expansion, use the department ID to query the next level.
//
// - `nextToken` is used to continue reading the next page under the same parent department. An empty string indicates that there are no more pages.
//
// ### Total count and pagination
//
// This operation does not calculate the total number of departments. Currently, `totalCount` returns JSON `null`, which does not mean the number of departments is 0. Whether to continue pagination depends on `nextToken`.
//
// ### Request examples
//
// ```http
//
// GET /workspaces/{workspaceId}/identity-providers/{identityProviderType}/departments?maxResults=50
//
// ```
//
// ### Response examples
//
// The following shows the response structure for an empty directory page. `totalCount` retains the `null` value actually returned by the operation.
//
// ```json
//
// {
//
//	"requestId": "01234567-89AB-CDEF-0123-456789ABCDEF",
//
//	"message": "success",
//
//	"httpStatusCode": 200,
//
//	"code": "SUCCESS",
//
//	"success": true,
//
//	"maxResults": 50,
//
//	"nextToken": "",
//
//	"totalCount": null,
//
//	"items": []
//
// }
//
// ```
//
// ### What to do next
//
// After selecting the departments to synchronize, save the synchronization scope by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider), and then call [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) to initiate synchronization as needed.
//
// @param request - ListIdentityProviderDepartmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProviderDepartmentsResponse
func (client *Client) ListIdentityProviderDepartmentsWithOptions(workspaceId *string, identityProviderType *string, request *ListIdentityProviderDepartmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIdentityProviderDepartmentsResponse, _err error) {
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

	if !dara.IsNil(request.ParentDepartmentId) {
		query["parentDepartmentId"] = request.ParentDepartmentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProviderDepartments"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType)) + "/departments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProviderDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the departments of an identity provider by paging.
//
// Description:
//
// Queries the list of DingTalk or Lark departments by parent department and paging. This is used when selecting or adjusting the user synchronization scope.
//
// ### Before you begin
//
// A DingTalk or Lark identity provider is bound to the workspace, and the user pool, application credentials, and authorization to read contacts are prepared.
//
// ### Query methods
//
// - Omit `parentDepartmentId` to query the root level. Pass a department ID to query its direct child departments.
//
// - `hasChildren` indicates whether a department can be expanded. After expansion, use the department ID to query the next level.
//
// - `nextToken` is used to continue reading the next page under the same parent department. An empty string indicates that there are no more pages.
//
// ### Total count and pagination
//
// This operation does not calculate the total number of departments. Currently, `totalCount` returns JSON `null`, which does not mean the number of departments is 0. Whether to continue pagination depends on `nextToken`.
//
// ### Request examples
//
// ```http
//
// GET /workspaces/{workspaceId}/identity-providers/{identityProviderType}/departments?maxResults=50
//
// ```
//
// ### Response examples
//
// The following shows the response structure for an empty directory page. `totalCount` retains the `null` value actually returned by the operation.
//
// ```json
//
// {
//
//	"requestId": "01234567-89AB-CDEF-0123-456789ABCDEF",
//
//	"message": "success",
//
//	"httpStatusCode": 200,
//
//	"code": "SUCCESS",
//
//	"success": true,
//
//	"maxResults": 50,
//
//	"nextToken": "",
//
//	"totalCount": null,
//
//	"items": []
//
// }
//
// ```
//
// ### What to do next
//
// After selecting the departments to synchronize, save the synchronization scope by calling [UpdateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProvider), and then call [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) to initiate synchronization as needed.
//
// @param request - ListIdentityProviderDepartmentsRequest
//
// @return ListIdentityProviderDepartmentsResponse
func (client *Client) ListIdentityProviderDepartments(workspaceId *string, identityProviderType *string, request *ListIdentityProviderDepartmentsRequest) (_result *ListIdentityProviderDepartmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIdentityProviderDepartmentsResponse{}
	_body, _err := client.ListIdentityProviderDepartmentsWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the complete summary list of up to two SCIM credentials in a shared user pool, without returning plaintext tokens or supporting pagination.
//
// Description:
//
// Queries the list of SCIM access credentials to check which tokens are still valid, when they expire, and to select credentials for revocation.
//
// SCIM tokens are used to authorize enterprise identity systems to send user provisioning requests to the user pool. This operation returns credential summaries, including the credential ID, status, creation time, and expiration time. The plaintext token is returned only upon creation.
//
// ### How to use the query results
//
// - Credentials with the Enabled status that have not expired can continue to be used.
//
// - To obtain a new token, call [CreateIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSCIMCredential).
//
// - To revoke a token, pass the target credentialId to [DeleteIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteIdentityProviderSCIMCredential).
//
// A single user pool can retain up to two SCIM credentials. The list may include expired or disabled credentials. Clean up unused credentials before creating a new token.
//
// ### Return method
//
// This operation returns the complete array of credentials at once without using pagination parameters. If no credentials exist, the following is returned:
//
// ```json
//
// {"data":[]}
//
// ```
//
// @param request - ListIdentityProviderSCIMCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProviderSCIMCredentialsResponse
func (client *Client) ListIdentityProviderSCIMCredentialsWithOptions(workspaceId *string, request *ListIdentityProviderSCIMCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIdentityProviderSCIMCredentialsResponse, _err error) {
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
		Action:      dara.String("ListIdentityProviderSCIMCredentials"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning/credentials"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProviderSCIMCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the complete summary list of up to two SCIM credentials in a shared user pool, without returning plaintext tokens or supporting pagination.
//
// Description:
//
// Queries the list of SCIM access credentials to check which tokens are still valid, when they expire, and to select credentials for revocation.
//
// SCIM tokens are used to authorize enterprise identity systems to send user provisioning requests to the user pool. This operation returns credential summaries, including the credential ID, status, creation time, and expiration time. The plaintext token is returned only upon creation.
//
// ### How to use the query results
//
// - Credentials with the Enabled status that have not expired can continue to be used.
//
// - To obtain a new token, call [CreateIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSCIMCredential).
//
// - To revoke a token, pass the target credentialId to [DeleteIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/DeleteIdentityProviderSCIMCredential).
//
// A single user pool can retain up to two SCIM credentials. The list may include expired or disabled credentials. Clean up unused credentials before creating a new token.
//
// ### Return method
//
// This operation returns the complete array of credentials at once without using pagination parameters. If no credentials exist, the following is returned:
//
// ```json
//
// {"data":[]}
//
// ```
//
// @param request - ListIdentityProviderSCIMCredentialsRequest
//
// @return ListIdentityProviderSCIMCredentialsResponse
func (client *Client) ListIdentityProviderSCIMCredentials(workspaceId *string, request *ListIdentityProviderSCIMCredentialsRequest) (_result *ListIdentityProviderSCIMCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIdentityProviderSCIMCredentialsResponse{}
	_body, _err := client.ListIdentityProviderSCIMCredentialsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the summary of existing identity provider bindings for a workspace, including CustomSSO. Returns local binding information without querying the actual upstream SAML or SCIM configurations.
//
// Description:
//
// Queries the enterprise identity providers currently connected to a workspace. This operation is used to display the identity provider list or to check whether a binding already exists before configuring a new connection.
//
// A workspace can be bound to at most one external identity provider: DingTalk, Lark, or a custom identity provider (`CustomSSO`). If no identity provider is bound, the returned `items` is an empty array.
//
// ### What you can view in the list
//
// The list returns the identity provider type, binding status, and logon and user synchronization switches. When the status of `CustomSSO` is `WAITING_CONFIGURATION`, it indicates that the user pool is ready and you can proceed to complete the integration configuration.
//
// This operation returns a binding summary. To view specific integration configurations, use the corresponding operations:
//
// - [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider): Views the identity provider details and SAML logon configuration.
//
// - [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider): Retrieves the SAML integration information required to configure AgentCore in the enterprise identity system.
//
// - [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM): Views the SCIM endpoint configuration and user synchronization status.
//
// ### Pagination
//
// Currently, at most one binding is returned. `nextToken` should be omitted or passed as an empty string. The default value of `maxResults` is 10. Valid values: 1 to 100.
//
// @param request - ListIdentityProvidersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProvidersWithOptions(workspaceId *string, request *ListIdentityProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
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
		Action:      dara.String("ListIdentityProviders"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary of existing identity provider bindings for a workspace, including CustomSSO. Returns local binding information without querying the actual upstream SAML or SCIM configurations.
//
// Description:
//
// Queries the enterprise identity providers currently connected to a workspace. This operation is used to display the identity provider list or to check whether a binding already exists before configuring a new connection.
//
// A workspace can be bound to at most one external identity provider: DingTalk, Lark, or a custom identity provider (`CustomSSO`). If no identity provider is bound, the returned `items` is an empty array.
//
// ### What you can view in the list
//
// The list returns the identity provider type, binding status, and logon and user synchronization switches. When the status of `CustomSSO` is `WAITING_CONFIGURATION`, it indicates that the user pool is ready and you can proceed to complete the integration configuration.
//
// This operation returns a binding summary. To view specific integration configurations, use the corresponding operations:
//
// - [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider): Views the identity provider details and SAML logon configuration.
//
// - [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider): Retrieves the SAML integration information required to configure AgentCore in the enterprise identity system.
//
// - [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM): Views the SCIM endpoint configuration and user synchronization status.
//
// ### Pagination
//
// Currently, at most one binding is returned. `nextToken` should be omitted or passed as an empty string. The default value of `maxResults` is 10. Valid values: 1 to 100.
//
// @param request - ListIdentityProvidersRequest
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProviders(workspaceId *string, request *ListIdentityProvidersRequest) (_result *ListIdentityProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.ListIdentityProvidersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query for image caches in a workspace. Supports paging and filtering by image source and cache status.
//
// Description:
//
// Performs a paged query for image caches in a workspace. Supports paging and filtering by image source and cache status.
//
// ### Filter conditions
//
// Provide regionId and workspaceId in the query parameters. You can filter by ACR instance, registry, namespace, repository, tag, digest, and status. Use imageRefContains to search for image address fragments.
//
// ### Pagination rules
//
// The default value of maxResults is 20, and the valid values are 1 to 100. Omit nextToken for the first request. For subsequent requests, use the token returned in the response as is, and keep the workspace, region, page size, and all filter conditions unchanged. An empty token indicates that there are no more pages.
//
// ### Response content
//
// Each item includes imageId, source, digest, status, size, reference count, and deletion conditions. PENDING only indicates that the cache is being prepared. For the cause of failure, see error. You can query a single cache by calling [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache).
//
// @param request - ListImageCachesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageCachesResponse
func (client *Client) ListImageCachesWithOptions(request *ListImageCachesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListImageCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		query["acrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.Digest) {
		query["digest"] = request.Digest
	}

	if !dara.IsNil(request.ImageRefContains) {
		query["imageRefContains"] = request.ImageRefContains
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.Registry) {
		query["registry"] = request.Registry
	}

	if !dara.IsNil(request.Repository) {
		query["repository"] = request.Repository
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["tag"] = request.Tag
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageCaches"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/image-caches"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for image caches in a workspace. Supports paging and filtering by image source and cache status.
//
// Description:
//
// Performs a paged query for image caches in a workspace. Supports paging and filtering by image source and cache status.
//
// ### Filter conditions
//
// Provide regionId and workspaceId in the query parameters. You can filter by ACR instance, registry, namespace, repository, tag, digest, and status. Use imageRefContains to search for image address fragments.
//
// ### Pagination rules
//
// The default value of maxResults is 20, and the valid values are 1 to 100. Omit nextToken for the first request. For subsequent requests, use the token returned in the response as is, and keep the workspace, region, page size, and all filter conditions unchanged. An empty token indicates that there are no more pages.
//
// ### Response content
//
// Each item includes imageId, source, digest, status, size, reference count, and deletion conditions. PENDING only indicates that the cache is being prepared. For the cause of failure, see error. You can query a single cache by calling [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache).
//
// @param request - ListImageCachesRequest
//
// @return ListImageCachesResponse
func (client *Client) ListImageCaches(request *ListImageCachesRequest) (_result *ListImageCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImageCachesResponse{}
	_body, _err := client.ListImageCachesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query for the identity, harness, configuration version, and status summary of managed Harness agents in a workspace.
//
// Description:
//
// Performs a paging query for the identity, harness, configuration version, and status summary of managed Harness agents in a workspace.
//
// ### Paging query
//
// If maxResults is omitted or set to 0, the default value is 20, and the maximum value is 100. Omit nextToken in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates that there are no more pages.
//
// ### Response content
//
// The list returns summary information, where harness.type indicates the runtime framework. For complete configurations, skill targets and effective versions, and instance status statistics, call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent).
//
// @param request - ListManagedAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedAgentsResponse
func (client *Client) ListManagedAgentsWithOptions(workspaceId *string, request *ListManagedAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListManagedAgentsResponse, _err error) {
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
		Action:      dara.String("ListManagedAgents"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManagedAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query for the identity, harness, configuration version, and status summary of managed Harness agents in a workspace.
//
// Description:
//
// Performs a paging query for the identity, harness, configuration version, and status summary of managed Harness agents in a workspace.
//
// ### Paging query
//
// If maxResults is omitted or set to 0, the default value is 20, and the maximum value is 100. Omit nextToken in the first request. For subsequent requests, pass the token returned in the previous response as is, and keep the same workspace and filter conditions. An empty token indicates that there are no more pages.
//
// ### Response content
//
// The list returns summary information, where harness.type indicates the runtime framework. For complete configurations, skill targets and effective versions, and instance status statistics, call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent).
//
// @param request - ListManagedAgentsRequest
//
// @return ListManagedAgentsResponse
func (client *Client) ListManagedAgents(workspaceId *string, request *ListManagedAgentsRequest) (_result *ListManagedAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListManagedAgentsResponse{}
	_body, _err := client.ListManagedAgentsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries MCP marketplace templates.
//
// Description:
//
// Queries official MCP marketplace templates available for installation to help select required tool services. Supports filtering by keyword, usage tag, and MCP type.
//
// @param request - ListMcpMarketItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpMarketItemsResponse
func (client *Client) ListMcpMarketItemsWithOptions(workspaceId *string, request *ListMcpMarketItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMcpMarketItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.McpType) {
		query["mcpType"] = request.McpType
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficialTag) {
		query["officialTag"] = request.OfficialTag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpMarketItems"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-market/items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpMarketItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries MCP marketplace templates.
//
// Description:
//
// Queries official MCP marketplace templates available for installation to help select required tool services. Supports filtering by keyword, usage tag, and MCP type.
//
// @param request - ListMcpMarketItemsRequest
//
// @return ListMcpMarketItemsResponse
func (client *Client) ListMcpMarketItems(workspaceId *string, request *ListMcpMarketItemsRequest) (_result *ListMcpMarketItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMcpMarketItemsResponse{}
	_body, _err := client.ListMcpMarketItemsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of tools exposed by a specified MCP service and their input and output schemas.
//
// Description:
//
// Queries the list of tools exposed by a specified MCP service and their input and output schemas.
//
// @param request - ListMcpToolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpToolsResponse
func (client *Client) ListMcpToolsWithOptions(workspaceId *string, mcpServerId *string, request *ListMcpToolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMcpToolsResponse, _err error) {
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
		Action:      dara.String("ListMcpTools"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId)) + "/tools"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpToolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of tools exposed by a specified MCP service and their input and output schemas.
//
// Description:
//
// Queries the list of tools exposed by a specified MCP service and their input and output schemas.
//
// @param request - ListMcpToolsRequest
//
// @return ListMcpToolsResponse
func (client *Client) ListMcpTools(workspaceId *string, mcpServerId *string, request *ListMcpToolsRequest) (_result *ListMcpToolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMcpToolsResponse{}
	_body, _err := client.ListMcpToolsWithOptions(workspaceId, mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query for MCP services and their configurations and status in a specified workspace.
//
// Description:
//
// Performs a paging query for MCP services and their configurations and status in a specified workspace.
//
// @param request - ListMcpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpsResponse
func (client *Client) ListMcpsWithOptions(workspaceId *string, request *ListMcpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMcpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomTag) {
		query["customTag"] = request.CustomTag
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

	if !dara.IsNil(request.OfficialTag) {
		query["officialTag"] = request.OfficialTag
	}

	if !dara.IsNil(request.SearchType) {
		query["searchType"] = request.SearchType
	}

	if !dara.IsNil(request.UsageActive) {
		query["usageActive"] = request.UsageActive
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcps"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query for MCP services and their configurations and status in a specified workspace.
//
// Description:
//
// Performs a paging query for MCP services and their configurations and status in a specified workspace.
//
// @param request - ListMcpsRequest
//
// @return ListMcpsResponse
func (client *Client) ListMcps(workspaceId *string, request *ListMcpsRequest) (_result *ListMcpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMcpsResponse{}
	_body, _err := client.ListMcpsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries long-term memory with paging and supports independent filtering by UserId, AgentId, and SessionId.
//
// Description:
//
// Queries long-term memory with paging and supports filtering by `UserId`, `AgentId`, and `SessionId`.
//
// @param request - ListMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoriesResponse
func (client *Client) ListMemoriesWithOptions(workspaceId *string, memoryStoreName *string, request *ListMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemories"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries long-term memory with paging and supports independent filtering by UserId, AgentId, and SessionId.
//
// Description:
//
// Queries long-term memory with paging and supports filtering by `UserId`, `AgentId`, and `SessionId`.
//
// @param request - ListMemoriesRequest
//
// @return ListMemoriesResponse
func (client *Client) ListMemories(workspaceId *string, memoryStoreName *string, request *ListMemoriesRequest) (_result *ListMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoriesResponse{}
	_body, _err := client.ListMemoriesWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries raw messages by exact scope with paging. SessionId is required, and at least one of UserId and AgentId must be provided.
//
// Description:
//
// Queries raw messages of a specific session in the memory store with paging to view the session content used for generating long-term memory.
//
// `SessionId` is required, and at least one of `UserId` and `AgentId` must be provided to identify the target session.
//
// @param request - ListMemorySessionMessagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemorySessionMessagesResponse
func (client *Client) ListMemorySessionMessagesWithOptions(workspaceId *string, memoryStoreName *string, request *ListMemorySessionMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemorySessionMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemorySessionMessages"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemorySessionMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries raw messages by exact scope with paging. SessionId is required, and at least one of UserId and AgentId must be provided.
//
// Description:
//
// Queries raw messages of a specific session in the memory store with paging to view the session content used for generating long-term memory.
//
// `SessionId` is required, and at least one of `UserId` and `AgentId` must be provided to identify the target session.
//
// @param request - ListMemorySessionMessagesRequest
//
// @return ListMemorySessionMessagesResponse
func (client *Client) ListMemorySessionMessages(workspaceId *string, memoryStoreName *string, request *ListMemorySessionMessagesRequest) (_result *ListMemorySessionMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemorySessionMessagesResponse{}
	_body, _err := client.ListMemorySessionMessagesWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query for sessions in a MemoryStore, supporting paging and filtering by user ID and agent ID independently.
//
// Description:
//
// Performs a paged query for source sessions in a long-term memory store (MemoryStore), supporting paging and filtering by user ID and agent ID respectively.
//
// @param request - ListMemorySessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemorySessionsResponse
func (client *Client) ListMemorySessionsWithOptions(workspaceId *string, memoryStoreName *string, request *ListMemorySessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemorySessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemorySessions"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemorySessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for sessions in a MemoryStore, supporting paging and filtering by user ID and agent ID independently.
//
// Description:
//
// Performs a paged query for source sessions in a long-term memory store (MemoryStore), supporting paging and filtering by user ID and agent ID respectively.
//
// @param request - ListMemorySessionsRequest
//
// @return ListMemorySessionsResponse
func (client *Client) ListMemorySessions(workspaceId *string, memoryStoreName *string, request *ListMemorySessionsRequest) (_result *ListMemorySessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemorySessionsResponse{}
	_body, _err := client.ListMemorySessionsWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of MemoryStores in a specified workspace by using paging.
//
// Description:
//
// Performs a paged query of the long-term memory stores (MemoryStores) in a workspace by using paging for selection or management.
//
// @param request - ListMemoryStoresRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryStoresResponse
func (client *Client) ListMemoryStoresWithOptions(workspaceId *string, request *ListMemoryStoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryStoresResponse, _err error) {
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
		Action:      dara.String("ListMemoryStores"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of MemoryStores in a specified workspace by using paging.
//
// Description:
//
// Performs a paged query of the long-term memory stores (MemoryStores) in a workspace by using paging for selection or management.
//
// @param request - ListMemoryStoresRequest
//
// @return ListMemoryStoresResponse
func (client *Client) ListMemoryStores(workspaceId *string, request *ListMemoryStoresRequest) (_result *ListMemoryStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryStoresResponse{}
	_body, _err := client.ListMemoryStoresWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries model connections in a specified workspace by using paging. Supports filtering by name, provider type, and invocation protocol.
//
// Description:
//
// Queries model connections in a workspace by using paging to view configured model services.
//
// - Use `Name` to filter by name, and `SearchType` to specify exact match or fuzzy match.
//
// - You can filter by model provider type and invocation protocol.
//
// @param request - ListModelConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelConnectionsResponse
func (client *Client) ListModelConnectionsWithOptions(workspaceId *string, request *ListModelConnectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeModels) {
		query["includeModels"] = request.IncludeModels
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

	if !dara.IsNil(request.Protocol) {
		query["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ProviderType) {
		query["providerType"] = request.ProviderType
	}

	if !dara.IsNil(request.SearchType) {
		query["searchType"] = request.SearchType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelConnections"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/model-connections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries model connections in a specified workspace by using paging. Supports filtering by name, provider type, and invocation protocol.
//
// Description:
//
// Queries model connections in a workspace by using paging to view configured model services.
//
// - Use `Name` to filter by name, and `SearchType` to specify exact match or fuzzy match.
//
// - You can filter by model provider type and invocation protocol.
//
// @param request - ListModelConnectionsRequest
//
// @return ListModelConnectionsResponse
func (client *Client) ListModelConnections(workspaceId *string, request *ListModelConnectionsRequest) (_result *ListModelConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelConnectionsResponse{}
	_body, _err := client.ListModelConnectionsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query for models in a specified workspace and supports filtering by model connection and model name.
//
// Description:
//
// Performs a paging query for models in a specified workspace and supports filtering by model connection and model name.
//
// @param request - ListModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
func (client *Client) ListModelsWithOptions(workspaceId *string, request *ListModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionId) {
		query["connectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModelName) {
		query["modelName"] = request.ModelName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query for models in a specified workspace and supports filtering by model connection and model name.
//
// Description:
//
// Performs a paging query for models in a specified workspace and supports filtering by model connection and model name.
//
// @param request - ListModelsRequest
//
// @return ListModelsResponse
func (client *Client) ListModels(workspaceId *string, request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries permission rules.
//
// Description:
//
// Performs a paged query of access rules by resource, supporting paging to check existing rules and their sources.
//
// @param request - ListPermissionRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionRulesResponse
func (client *Client) ListPermissionRulesWithOptions(workspaceId *string, request *ListPermissionRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionRulesResponse, _err error) {
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

	if !dara.IsNil(request.PrincipalId) {
		query["principalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalName) {
		query["principalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		query["principalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Target) {
		query["target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionRules"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/permission-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries permission rules.
//
// Description:
//
// Performs a paged query of access rules by resource, supporting paging to check existing rules and their sources.
//
// @param request - ListPermissionRulesRequest
//
// @return ListPermissionRulesResponse
func (client *Client) ListPermissionRules(workspaceId *string, request *ListPermissionRulesRequest) (_result *ListPermissionRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionRulesResponse{}
	_body, _err := client.ListPermissionRulesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the providers supported by the built-in model catalog of AgentCore and their default service configurations.
//
// Description:
//
// Queries the providers supported by the built-in model catalog of AgentCore and their default service configurations.
//
// @param request - ListPredefinedModelProvidersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPredefinedModelProvidersResponse
func (client *Client) ListPredefinedModelProvidersWithOptions(request *ListPredefinedModelProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPredefinedModelProvidersResponse, _err error) {
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
		Action:      dara.String("ListPredefinedModelProviders"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/model-catalog/providers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPredefinedModelProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the providers supported by the built-in model catalog of AgentCore and their default service configurations.
//
// Description:
//
// Queries the providers supported by the built-in model catalog of AgentCore and their default service configurations.
//
// @param request - ListPredefinedModelProvidersRequest
//
// @return ListPredefinedModelProvidersResponse
func (client *Client) ListPredefinedModelProviders(request *ListPredefinedModelProvidersRequest) (_result *ListPredefinedModelProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPredefinedModelProvidersResponse{}
	_body, _err := client.ListPredefinedModelProvidersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the models and their capabilities for a specified provider in the AgentCore built-in model catalog.
//
// Description:
//
// Queries the models and their capabilities for a specified provider in the AgentCore built-in model catalog.
//
// @param request - ListPredefinedModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPredefinedModelsResponse
func (client *Client) ListPredefinedModelsWithOptions(providerType *string, request *ListPredefinedModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPredefinedModelsResponse, _err error) {
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
		Action:      dara.String("ListPredefinedModels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/model-catalog/providers/" + dara.PercentEncode(dara.StringValue(providerType)) + "/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPredefinedModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the models and their capabilities for a specified provider in the AgentCore built-in model catalog.
//
// Description:
//
// Queries the models and their capabilities for a specified provider in the AgentCore built-in model catalog.
//
// @param request - ListPredefinedModelsRequest
//
// @return ListPredefinedModelsResponse
func (client *Client) ListPredefinedModels(providerType *string, request *ListPredefinedModelsRequest) (_result *ListPredefinedModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPredefinedModelsResponse{}
	_body, _err := client.ListPredefinedModelsWithOptions(providerType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve the current active sessions and their sources in a specified agent instance.
//
// Description:
//
// Queries the active sessions in a specified running instance to view the current mapping between sessions and instances.
//
// ### Query scope
//
// Locates the instance by using workspaceId, agentId, and sandboxId. You can obtain the sandboxId from [ListSandboxes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxes). This operation queries the current active sessions.
//
// ### Pagination
//
// The default value of maxResults is 20. Valid values: 1 to 100. Omit nextToken in the first request. For subsequent requests, use the returned token as is, and keep the workspace, agent, instance, and page size unchanged. An empty token indicates that no more pages are available.
//
// ### Session sources
//
// The source can be API, CONSOLE_DEBUG, EXTERNAL_CHANNEL, or UNKNOWN. For external channels, the channel type is additionally returned. If no source information is provided, UNKNOWN is returned.
//
// @param request - ListSandboxSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSandboxSessionsResponse
func (client *Client) ListSandboxSessionsWithOptions(workspaceId *string, agentId *string, sandboxId *string, request *ListSandboxSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSandboxSessionsResponse, _err error) {
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

	if !dara.IsNil(request.ServiceEndpointId) {
		query["serviceEndpointId"] = request.ServiceEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSandboxSessions"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSandboxSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve the current active sessions and their sources in a specified agent instance.
//
// Description:
//
// Queries the active sessions in a specified running instance to view the current mapping between sessions and instances.
//
// ### Query scope
//
// Locates the instance by using workspaceId, agentId, and sandboxId. You can obtain the sandboxId from [ListSandboxes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxes). This operation queries the current active sessions.
//
// ### Pagination
//
// The default value of maxResults is 20. Valid values: 1 to 100. Omit nextToken in the first request. For subsequent requests, use the returned token as is, and keep the workspace, agent, instance, and page size unchanged. An empty token indicates that no more pages are available.
//
// ### Session sources
//
// The source can be API, CONSOLE_DEBUG, EXTERNAL_CHANNEL, or UNKNOWN. For external channels, the channel type is additionally returned. If no source information is provided, UNKNOWN is returned.
//
// @param request - ListSandboxSessionsRequest
//
// @return ListSandboxSessionsResponse
func (client *Client) ListSandboxSessions(workspaceId *string, agentId *string, sandboxId *string, request *ListSandboxSessionsRequest) (_result *ListSandboxSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSandboxSessionsResponse{}
	_body, _err := client.ListSandboxSessionsWithOptions(workspaceId, agentId, sandboxId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query on the running instances of an agent, and supports filtering by Sandbox ID and current valid session ID fragments.
//
// Description:
//
// Queries the running instances (sandboxes) of an agent to view instance distribution and current running status.
//
// ### Filter instances
//
// `searchText` matches Sandbox ID fragments, and `sessionId` matches current valid session ID fragments. Both are case-insensitive. If both are submitted, they are combined for filtering.
//
// ### Pagination and status
//
// The default value of `maxResults` is 20. Valid values: 1 to 100. Omit `nextToken` for the first request. For subsequent requests, use the response token as is, and keep `workspaceId`, `agentId`, filter conditions, and `maxResults` unchanged. An empty token indicates that the query is complete.
//
// `phase` indicates the instance phase, `activeSessionCount` indicates the number of current active sessions, and `maxConcurrentSessions` indicates the configured concurrency limit for a single instance. You can query the active sessions of an instance by calling [ListSandboxSessions](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxSessions).
//
// @param request - ListSandboxesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSandboxesResponse
func (client *Client) ListSandboxesWithOptions(workspaceId *string, agentId *string, request *ListSandboxesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSandboxesResponse, _err error) {
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

	if !dara.IsNil(request.SearchText) {
		query["searchText"] = request.SearchText
	}

	if !dara.IsNil(request.ServiceEndpointId) {
		query["serviceEndpointId"] = request.ServiceEndpointId
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSandboxes"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/sandboxes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSandboxesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query on the running instances of an agent, and supports filtering by Sandbox ID and current valid session ID fragments.
//
// Description:
//
// Queries the running instances (sandboxes) of an agent to view instance distribution and current running status.
//
// ### Filter instances
//
// `searchText` matches Sandbox ID fragments, and `sessionId` matches current valid session ID fragments. Both are case-insensitive. If both are submitted, they are combined for filtering.
//
// ### Pagination and status
//
// The default value of `maxResults` is 20. Valid values: 1 to 100. Omit `nextToken` for the first request. For subsequent requests, use the response token as is, and keep `workspaceId`, `agentId`, filter conditions, and `maxResults` unchanged. An empty token indicates that the query is complete.
//
// `phase` indicates the instance phase, `activeSessionCount` indicates the number of current active sessions, and `maxConcurrentSessions` indicates the configured concurrency limit for a single instance. You can query the active sessions of an instance by calling [ListSandboxSessions](https://api.aliyun.com/document/AgentCore/2026-08-04/ListSandboxSessions).
//
// @param request - ListSandboxesRequest
//
// @return ListSandboxesResponse
func (client *Client) ListSandboxes(workspaceId *string, agentId *string, request *ListSandboxesRequest) (_result *ListSandboxesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSandboxesResponse{}
	_body, _err := client.ListSandboxesWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service endpoints in a specified workspace with paging, and supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// Queries service endpoints in a workspace with paging to find the access endpoints for agents or collaboration components.
//
// - Use targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status to filter the results.
//
// - Use maxResults to set the number of entries per page. If omitted, the default value is 20.
//
// - Use the returned nextToken to query the next page.
//
// @param request - ListServiceEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceEndpointsResponse
func (client *Client) ListServiceEndpointsWithOptions(workspaceId *string, request *ListServiceEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		query["agentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.CollaborationComponent) {
		query["collaborationComponent"] = request.CollaborationComponent
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceBindingId) {
		query["resourceBindingId"] = request.ResourceBindingId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TargetType) {
		query["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceEndpoints"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/service-endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service endpoints in a specified workspace with paging, and supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// Queries service endpoints in a workspace with paging to find the access endpoints for agents or collaboration components.
//
// - Use targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status to filter the results.
//
// - Use maxResults to set the number of entries per page. If omitted, the default value is 20.
//
// - Use the returned nextToken to query the next page.
//
// @param request - ListServiceEndpointsRequest
//
// @return ListServiceEndpointsResponse
func (client *Client) ListServiceEndpoints(workspaceId *string, request *ListServiceEndpointsRequest) (_result *ListServiceEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceEndpointsResponse{}
	_body, _err := client.ListServiceEndpointsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the reference relationships of a skill.
//
// Description:
//
// Queries the reference relationships of a specified skill to verify its usage before modifying or cleaning up skill assets.
//
// @param request - ListSkillReferencesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillReferencesResponse
func (client *Client) ListSkillReferencesWithOptions(workspaceId *string, skillName *string, request *ListSkillReferencesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSkillReferencesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SelectorType) {
		query["selectorType"] = request.SelectorType
	}

	if !dara.IsNil(request.SelectorValue) {
		query["selectorValue"] = request.SelectorValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillReferences"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/references"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillReferencesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reference relationships of a skill.
//
// Description:
//
// Queries the reference relationships of a specified skill to verify its usage before modifying or cleaning up skill assets.
//
// @param request - ListSkillReferencesRequest
//
// @return ListSkillReferencesResponse
func (client *Client) ListSkillReferences(workspaceId *string, skillName *string, request *ListSkillReferencesRequest) (_result *ListSkillReferencesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSkillReferencesResponse{}
	_body, _err := client.ListSkillReferencesWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of skills in a specified workspace and returns the basic information, version status, and paging information of the skills.
//
// Description:
//
// Performs a paged query of skills in a specified workspace and returns the basic information, version status, and paging information of the skills.
//
// @param request - ListSkillsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(workspaceId *string, request *ListSkillsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
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

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Owner) {
		query["owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNo) {
		query["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.SkillName) {
		query["skillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of skills in a specified workspace and returns the basic information, version status, and paging information of the skills.
//
// Description:
//
// Performs a paged query of skills in a specified workspace and returns the basic information, version status, and paging information of the skills.
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(workspaceId *string, request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query of teams in a specified workspace. You can use nameLike to filter teams by fuzzy match on the team name, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. The member information in the list includes only the member ID, name, and team role.
//
// Description:
//
// Performs a paging query of teams in a specified workspace.
//
// @param request - ListTeamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamsResponse
func (client *Client) ListTeamsWithOptions(workspaceId *string, request *ListTeamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
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

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeams"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/teams"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query of teams in a specified workspace. You can use nameLike to filter teams by fuzzy match on the team name, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. The member information in the list includes only the member ID, name, and team role.
//
// Description:
//
// Performs a paging query of teams in a specified workspace.
//
// @param request - ListTeamsRequest
//
// @return ListTeamsResponse
func (client *Client) ListTeams(workspaceId *string, request *ListTeamsRequest) (_result *ListTeamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTeamsResponse{}
	_body, _err := client.ListTeamsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query for users in a specified workspace. You can filter users by fuzzy match on the username by using nameLike, specify the maximum number of records per paging by using maxResults, and retrieve the next paging by using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Queries the application user directory in a workspace for viewing users and selecting team members.
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(workspaceId *string, request *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["email"] = request.Email
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for users in a specified workspace. You can filter users by fuzzy match on the username by using nameLike, specify the maximum number of records per paging by using maxResults, and retrieve the next paging by using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Queries the application user directory in a workspace for viewing users and selecting team members.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(workspaceId *string, request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of workspaces under the current tenant. By default, the list does not return soft-deleted records whose status is Deleted. The server applies stable sorting to the results in chronological order of creation during paging.
//
// Description:
//
// Performs a paged query of workspaces accessible to the current account for selecting a workspace or viewing a workspace overview.
//
// - Use `nameLike` to filter by name using fuzzy match.
//
// - Use `maxResults` to configure the Settings for the number of entries per page. If this parameter is omitted or set to `0`, the default value is 20.
//
// - Use `nextToken` to query the next page during paging, or use `skip` to skip a specified number of workspaces.
//
// The list is sorted in the order of creation. By default, deleted workspaces (with the `Deleted` status) are not returned.
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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

	if !dara.IsNil(request.Skip) {
		query["skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of workspaces under the current tenant. By default, the list does not return soft-deleted records whose status is Deleted. The server applies stable sorting to the results in chronological order of creation during paging.
//
// Description:
//
// Performs a paged query of workspaces accessible to the current account for selecting a workspace or viewing a workspace overview.
//
// - Use `nameLike` to filter by name using fuzzy match.
//
// - Use `maxResults` to configure the Settings for the number of entries per page. If this parameter is omitted or set to `0`, the default value is 20.
//
// - Use `nextToken` to query the next page during paging, or use `skip` to skip a specified number of workspaces.
//
// The list is sorted in the order of creation. By default, deleted workspaces (with the `Deleted` status) are not returned.
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes the online version of a specified skill offline so that it is no longer used as the online version.
//
// Description:
//
// Takes the online version of a specified skill offline so that it is no longer used as the online version.
//
// @param tmpReq - OfflineSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineSkillResponse
func (client *Client) OfflineSkillWithOptions(workspaceId *string, skillName *string, tmpReq *OfflineSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflineSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineSkill"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/actions/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes the online version of a specified skill offline so that it is no longer used as the online version.
//
// Description:
//
// Takes the online version of a specified skill offline so that it is no longer used as the online version.
//
// @param request - OfflineSkillRequest
//
// @return OfflineSkillResponse
func (client *Client) OfflineSkill(workspaceId *string, skillName *string, request *OfflineSkillRequest) (_result *OfflineSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineSkillResponse{}
	_body, _err := client.OfflineSkillWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the specified skill version as the online version.
//
// Description:
//
// Sets the specified skill version as the online version.
//
// @param tmpReq - OnlineSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineSkillResponse
func (client *Client) OnlineSkillWithOptions(workspaceId *string, skillName *string, tmpReq *OnlineSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OnlineSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineSkill"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/actions/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the specified skill version as the online version.
//
// Description:
//
// Sets the specified skill version as the online version.
//
// @param request - OnlineSkillRequest
//
// @return OnlineSkillResponse
func (client *Client) OnlineSkill(workspaceId *string, skillName *string, request *OnlineSkillRequest) (_result *OnlineSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineSkillResponse{}
	_body, _err := client.OnlineSkillWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Parses and checks one or more skill ZIP packages uploaded to OSS, and returns the names, versions, and conflict check results.
//
// Description:
//
// Checks one or more skill ZIP packages uploaded to OSS for skill names, versions, and import conflicts.
//
// After the check passes, call [UploadSkillViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadSkillViaOss) or [BatchUploadSkillsViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/BatchUploadSkillsViaOss) to import the skills into the workspace.
//
// @param tmpReq - PrecheckSkillUploadViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckSkillUploadViaOssResponse
func (client *Client) PrecheckSkillUploadViaOssWithOptions(workspaceId *string, tmpReq *PrecheckSkillUploadViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PrecheckSkillUploadViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PrecheckSkillUploadViaOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrecheckSkillUploadViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skill-actions/precheck-upload-via-oss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrecheckSkillUploadViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses and checks one or more skill ZIP packages uploaded to OSS, and returns the names, versions, and conflict check results.
//
// Description:
//
// Checks one or more skill ZIP packages uploaded to OSS for skill names, versions, and import conflicts.
//
// After the check passes, call [UploadSkillViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/UploadSkillViaOss) or [BatchUploadSkillsViaOss](https://api.aliyun.com/document/AgentCore/2026-08-04/BatchUploadSkillsViaOss) to import the skills into the workspace.
//
// @param request - PrecheckSkillUploadViaOssRequest
//
// @return PrecheckSkillUploadViaOssResponse
func (client *Client) PrecheckSkillUploadViaOss(workspaceId *string, request *PrecheckSkillUploadViaOssRequest) (_result *PrecheckSkillUploadViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PrecheckSkillUploadViaOssResponse{}
	_body, _err := client.PrecheckSkillUploadViaOssWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Prepares the SCIM receiving configuration and returns the provisioning URL and credential summary. Issues a token if no available credentials exist and returns the plaintext only in this response, without automatically enabling AgentCore user synchronization.
//
// Description:
//
// Prepares the receiving endpoint and access token for integrating an enterprise identity system with SCIM user synchronization. The enterprise identity system uses this information to provision users to the shared user pool of AgentCore.
//
// SCIM is a user provisioning protocol. The enterprise identity system first sends users to the receiving endpoint, and then AgentCore synchronizes the users from the user pool to the workspace.
//
// ### Before you begin
//
// A `CustomSSO` binding has been created by using [CreateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProvider), and the shared user pool is ready. SCIM integration can be performed independently of SAML logon configuration.
//
// ### What this operation prepares
//
// This operation enables the SCIM receiving endpoint and returns the receiving URL and credential summary:
//
// | Current credential status | Returned result |
//
// | --- | --- |
//
// | An unexpired credential with the status of `Enabled` already exists | Reuses the existing credential. The `token` field returns an empty string `""`. Use the previously saved token. |
//
// | No available credentials exist | Creates a new credential. The `token` field returns the plaintext of the new token, which is returned only in this response. |
//
// The plaintext of an existing token cannot be queried again. If the token is lost, you can create a new token by calling [CreateIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSCIMCredential). A maximum of two credentials can be retained for the same user pool. When the limit is reached, you must delete unused credentials first.
//
// ### Subsequent configuration
//
// 1. Configure the returned receiving URL and valid token in the SCIM user provisioning settings of the enterprise identity system.
//
// 2. Complete the connection test and user assignment in the enterprise identity system.
//
// 3. Enable AgentCore user synchronization by calling [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM), and then query the synchronization result by calling [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM).
//
// @param request - PrepareIdentityProviderSCIMRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrepareIdentityProviderSCIMResponse
func (client *Client) PrepareIdentityProviderSCIMWithOptions(workspaceId *string, request *PrepareIdentityProviderSCIMRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PrepareIdentityProviderSCIMResponse, _err error) {
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
		Action:      dara.String("PrepareIdentityProviderSCIM"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning/prepare"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PrepareIdentityProviderSCIMResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prepares the SCIM receiving configuration and returns the provisioning URL and credential summary. Issues a token if no available credentials exist and returns the plaintext only in this response, without automatically enabling AgentCore user synchronization.
//
// Description:
//
// Prepares the receiving endpoint and access token for integrating an enterprise identity system with SCIM user synchronization. The enterprise identity system uses this information to provision users to the shared user pool of AgentCore.
//
// SCIM is a user provisioning protocol. The enterprise identity system first sends users to the receiving endpoint, and then AgentCore synchronizes the users from the user pool to the workspace.
//
// ### Before you begin
//
// A `CustomSSO` binding has been created by using [CreateIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProvider), and the shared user pool is ready. SCIM integration can be performed independently of SAML logon configuration.
//
// ### What this operation prepares
//
// This operation enables the SCIM receiving endpoint and returns the receiving URL and credential summary:
//
// | Current credential status | Returned result |
//
// | --- | --- |
//
// | An unexpired credential with the status of `Enabled` already exists | Reuses the existing credential. The `token` field returns an empty string `""`. Use the previously saved token. |
//
// | No available credentials exist | Creates a new credential. The `token` field returns the plaintext of the new token, which is returned only in this response. |
//
// The plaintext of an existing token cannot be queried again. If the token is lost, you can create a new token by calling [CreateIdentityProviderSCIMCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSCIMCredential). A maximum of two credentials can be retained for the same user pool. When the limit is reached, you must delete unused credentials first.
//
// ### Subsequent configuration
//
// 1. Configure the returned receiving URL and valid token in the SCIM user provisioning settings of the enterprise identity system.
//
// 2. Complete the connection test and user assignment in the enterprise identity system.
//
// 3. Enable AgentCore user synchronization by calling [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM), and then query the synchronization result by calling [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM).
//
// @param request - PrepareIdentityProviderSCIMRequest
//
// @return PrepareIdentityProviderSCIMResponse
func (client *Client) PrepareIdentityProviderSCIM(workspaceId *string, request *PrepareIdentityProviderSCIMRequest) (_result *PrepareIdentityProviderSCIMResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PrepareIdentityProviderSCIMResponse{}
	_body, _err := client.PrepareIdentityProviderSCIMWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a specified skill version and changes its status to published.
//
// Description:
//
// Publishes a specified skill version and changes its status to published.
//
// @param tmpReq - PublishSkillVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishSkillVersionResponse
func (client *Client) PublishSkillVersionWithOptions(workspaceId *string, skillName *string, skillVersion *string, tmpReq *PublishSkillVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PublishSkillVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishSkillVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion)) + "/actions/publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified skill version and changes its status to published.
//
// Description:
//
// Publishes a specified skill version and changes its status to published.
//
// @param request - PublishSkillVersionRequest
//
// @return PublishSkillVersionResponse
func (client *Client) PublishSkillVersion(workspaceId *string, skillName *string, skillVersion *string, request *PublishSkillVersionRequest) (_result *PublishSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishSkillVersionResponse{}
	_body, _err := client.PublishSkillVersionWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a request to rebuild a single sandbox image for a specified managed agent. Existing sessions may be interrupted. An HTTP 200 response only indicates that the request is accepted. Use the returned operationId to call GetSandboxRebuildOperation to query the overall result. Only a SUCCEEDED status indicates that the rebuild is complete. Do not automatically resubmit the rebuild request if the query times out, returns UNKNOWN, or the result is unclear.
//
// Description:
//
// Submits a request to rebuild a single sandbox image for a specified managed agent. Existing sessions may be interrupted. An HTTP 200 response only indicates that the request is successfully submitted and accepted. You must use the returned operationId to call GetSandboxRebuildOperation. Only an overall SUCCEEDED status indicates that the rebuild is complete. Do not automatically resubmit the request if the query times out, returns UNKNOWN, or the result is unclear. You can specify an endpoint runtime version by using serviceEndpointId. If this parameter is omitted, the default latest runtime identity is used.
//
// @param request - RebuildSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildSandboxResponse
func (client *Client) RebuildSandboxWithOptions(workspaceId *string, agentId *string, sandboxId *string, request *RebuildSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebuildSandboxResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceEndpointId) {
		query["serviceEndpointId"] = request.ServiceEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildSandbox"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/rebuild"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a request to rebuild a single sandbox image for a specified managed agent. Existing sessions may be interrupted. An HTTP 200 response only indicates that the request is accepted. Use the returned operationId to call GetSandboxRebuildOperation to query the overall result. Only a SUCCEEDED status indicates that the rebuild is complete. Do not automatically resubmit the rebuild request if the query times out, returns UNKNOWN, or the result is unclear.
//
// Description:
//
// Submits a request to rebuild a single sandbox image for a specified managed agent. Existing sessions may be interrupted. An HTTP 200 response only indicates that the request is successfully submitted and accepted. You must use the returned operationId to call GetSandboxRebuildOperation. Only an overall SUCCEEDED status indicates that the rebuild is complete. Do not automatically resubmit the request if the query times out, returns UNKNOWN, or the result is unclear. You can specify an endpoint runtime version by using serviceEndpointId. If this parameter is omitted, the default latest runtime identity is used.
//
// @param request - RebuildSandboxRequest
//
// @return RebuildSandboxResponse
func (client *Client) RebuildSandbox(workspaceId *string, agentId *string, sandboxId *string, request *RebuildSandboxRequest) (_result *RebuildSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebuildSandboxResponse{}
	_body, _err := client.RebuildSandboxWithOptions(workspaceId, agentId, sandboxId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reverts a specified skill version to an editable draft for further modification.
//
// Description:
//
// Reverts a specified skill version to an editable draft for further modification.
//
// @param tmpReq - RedraftSkillVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedraftSkillVersionResponse
func (client *Client) RedraftSkillVersionWithOptions(workspaceId *string, skillName *string, skillVersion *string, tmpReq *RedraftSkillVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RedraftSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RedraftSkillVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedraftSkillVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion)) + "/actions/redraft"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedraftSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reverts a specified skill version to an editable draft for further modification.
//
// Description:
//
// Reverts a specified skill version to an editable draft for further modification.
//
// @param request - RedraftSkillVersionRequest
//
// @return RedraftSkillVersionResponse
func (client *Client) RedraftSkillVersion(workspaceId *string, skillName *string, skillVersion *string, request *RedraftSkillVersionRequest) (_result *RedraftSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RedraftSkillVersionResponse{}
	_body, _err := client.RedraftSkillVersionWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the logon password of a specified user. Specify the user by using agentCoreUserId or username. You must specify at least one of them. This operation only supports resetting the passwords of users who use local password authentication for the workspace. If you do not specify the password parameter, the server generates a random password and returns it in the response.
//
// Description:
//
// Resets the logon password of a specified user.
//
// @param tmpReq - ResetUserPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithOptions(workspaceId *string, tmpReq *ResetUserPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResetUserPasswordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserPassword"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users/actions/reset-password"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the logon password of a specified user. Specify the user by using agentCoreUserId or username. You must specify at least one of them. This operation only supports resetting the passwords of users who use local password authentication for the workspace. If you do not specify the password parameter, the server generates a random password and returns it in the response.
//
// Description:
//
// Resets the logon password of a specified user.
//
// @param request - ResetUserPasswordRequest
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPassword(workspaceId *string, request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resubmits the creation request for an ACR image cache in a specified workspace.
//
// Description:
//
// Resubmits the creation request for an existing image cache to recreate a cache that has not been successfully created.
//
// ### Call scope
//
// Pass `regionId`, `workspaceId`, and `clientToken` as query parameters, and use `imageId` to locate the cache. This operation applies to ACR source caches and re-verifies the source repository authorization before submission.
//
// ### Query retry results
//
// A successful response indicates that a retry acceptance snapshot is returned, not that the creation is complete. Query the status through [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache). `READY` indicates that the cache is ready. If the status is `FAILED`, check the `error` field.
//
// The `clientToken` parameter currently only verifies that the value is not empty and does not guarantee deduplication or recovery of the original acceptance result based on this identifier.
//
// @param tmpReq - RetryImageCacheRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryImageCacheResponse
func (client *Client) RetryImageCacheWithOptions(imageId *string, tmpReq *RetryImageCacheRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RetryImageCacheShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryImageCache"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/image-caches/" + dara.PercentEncode(dara.StringValue(imageId)) + "/retry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resubmits the creation request for an ACR image cache in a specified workspace.
//
// Description:
//
// Resubmits the creation request for an existing image cache to recreate a cache that has not been successfully created.
//
// ### Call scope
//
// Pass `regionId`, `workspaceId`, and `clientToken` as query parameters, and use `imageId` to locate the cache. This operation applies to ACR source caches and re-verifies the source repository authorization before submission.
//
// ### Query retry results
//
// A successful response indicates that a retry acceptance snapshot is returned, not that the creation is complete. Query the status through [GetImageCache](https://api.aliyun.com/document/AgentCore/2026-08-04/GetImageCache). `READY` indicates that the cache is ready. If the status is `FAILED`, check the `error` field.
//
// The `clientToken` parameter currently only verifies that the value is not empty and does not guarantee deduplication or recovery of the original acceptance result based on this identifier.
//
// @param request - RetryImageCacheRequest
//
// @return RetryImageCacheResponse
func (client *Client) RetryImageCache(imageId *string, request *RetryImageCacheRequest) (_result *RetryImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryImageCacheResponse{}
	_body, _err := client.RetryImageCacheWithOptions(imageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves long-term memory from MemoryStore based on semantics.
//
// Description:
//
// Retrieves relevant memories semantically from the long-term memory store (MemoryStore) to help agents obtain the information required for subsequent tasks.
//
// @param tmpReq - SearchMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemoriesWithOptions(workspaceId *string, memoryStoreName *string, tmpReq *SearchMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchMemoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMemories"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves long-term memory from MemoryStore based on semantics.
//
// Description:
//
// Retrieves relevant memories semantically from the long-term memory store (MemoryStore) to help agents obtain the information required for subsequent tasks.
//
// @param request - SearchMemoriesRequest
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemories(workspaceId *string, memoryStoreName *string, request *SearchMemoriesRequest) (_result *SearchMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchMemoriesResponse{}
	_body, _err := client.SearchMemoriesWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a specified draft version of an AgentSpec for review. If no release pipeline is configured, the version is directly published online.
//
// Description:
//
// Submits a specified draft version of an agent template for review.
//
// If no release pipeline is configured, the version is directly published online.
//
// @param tmpReq - SubmitAgentSpecVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAgentSpecVersionResponse
func (client *Client) SubmitAgentSpecVersionWithOptions(workspaceId *string, agentSpecName *string, agentSpecVersion *string, tmpReq *SubmitAgentSpecVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitAgentSpecVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitAgentSpecVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAgentSpecVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName)) + "/versions/" + dara.PercentEncode(dara.StringValue(agentSpecVersion)) + "/actions/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAgentSpecVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a specified draft version of an AgentSpec for review. If no release pipeline is configured, the version is directly published online.
//
// Description:
//
// Submits a specified draft version of an agent template for review.
//
// If no release pipeline is configured, the version is directly published online.
//
// @param request - SubmitAgentSpecVersionRequest
//
// @return SubmitAgentSpecVersionResponse
func (client *Client) SubmitAgentSpecVersion(workspaceId *string, agentSpecName *string, agentSpecVersion *string, request *SubmitAgentSpecVersionRequest) (_result *SubmitAgentSpecVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitAgentSpecVersionResponse{}
	_body, _err := client.SubmitAgentSpecVersionWithOptions(workspaceId, agentSpecName, agentSpecVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a specified draft version of a skill for review.
//
// Description:
//
// Submits a specified draft version of a skill for review.
//
// @param tmpReq - SubmitSkillVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSkillVersionResponse
func (client *Client) SubmitSkillVersionWithOptions(workspaceId *string, skillName *string, skillVersion *string, tmpReq *SubmitSkillVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitSkillVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSkillVersion"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/versions/" + dara.PercentEncode(dara.StringValue(skillVersion)) + "/actions/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a specified draft version of a skill for review.
//
// Description:
//
// Submits a specified draft version of a skill for review.
//
// @param request - SubmitSkillVersionRequest
//
// @return SubmitSkillVersionResponse
func (client *Client) SubmitSkillVersion(workspaceId *string, skillName *string, skillVersion *string, request *SubmitSkillVersionRequest) (_result *SubmitSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitSkillVersionResponse{}
	_body, _err := client.SubmitSkillVersionWithOptions(workspaceId, skillName, skillVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access application.
//
// Description:
//
// Modifies the configuration of an access application in a workspace to adjust application integration and identity source settings.
//
// @param tmpReq - UpdateAccessApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessApplicationResponse
func (client *Client) UpdateAccessApplicationWithOptions(workspaceId *string, applicationId *string, tmpReq *UpdateAccessApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccessApplication"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/access-applications/" + dara.PercentEncode(dara.StringValue(applicationId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an access application.
//
// Description:
//
// Modifies the configuration of an access application in a workspace to adjust application integration and identity source settings.
//
// @param request - UpdateAccessApplicationRequest
//
// @return UpdateAccessApplicationResponse
func (client *Client) UpdateAccessApplication(workspaceId *string, applicationId *string, request *UpdateAccessApplicationRequest) (_result *UpdateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAccessApplicationResponse{}
	_body, _err := client.UpdateAccessApplicationWithOptions(workspaceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel. At least one updatable field must be provided.
//
// Description:
//
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel.
//
// At least one updatable field must be provided.
//
// @param tmpReq - UpdateAgentIMChannelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentIMChannelResponse
func (client *Client) UpdateAgentIMChannelWithOptions(workspaceId *string, agentId *string, imChannelId *string, tmpReq *UpdateAgentIMChannelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentIMChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAgentIMChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentIMChannel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels/" + dara.PercentEncode(dara.StringValue(imChannelId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentIMChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel. At least one updatable field must be provided.
//
// Description:
//
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel.
//
// At least one updatable field must be provided.
//
// @param request - UpdateAgentIMChannelRequest
//
// @return UpdateAgentIMChannelResponse
func (client *Client) UpdateAgentIMChannel(workspaceId *string, agentId *string, imChannelId *string, request *UpdateAgentIMChannelRequest) (_result *UpdateAgentIMChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentIMChannelResponse{}
	_body, _err := client.UpdateAgentIMChannelWithOptions(workspaceId, agentId, imChannelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Fully updates the channel credentials for a specified IM channel. Secrets are not returned in the response.
//
// Description:
//
// Replaces the access credentials for an IM channel to rotate keys or correct authentication configurations.
//
// Submit the complete new credentials. The response does not return plaintext secrets.
//
// @param tmpReq - UpdateAgentIMChannelCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentIMChannelCredentialResponse
func (client *Client) UpdateAgentIMChannelCredentialWithOptions(workspaceId *string, agentId *string, imChannelId *string, tmpReq *UpdateAgentIMChannelCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentIMChannelCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAgentIMChannelCredentialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentIMChannelCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agents/" + dara.PercentEncode(dara.StringValue(agentId)) + "/im-channels/" + dara.PercentEncode(dara.StringValue(imChannelId)) + "/actions/update-credential"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentIMChannelCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Fully updates the channel credentials for a specified IM channel. Secrets are not returned in the response.
//
// Description:
//
// Replaces the access credentials for an IM channel to rotate keys or correct authentication configurations.
//
// Submit the complete new credentials. The response does not return plaintext secrets.
//
// @param request - UpdateAgentIMChannelCredentialRequest
//
// @return UpdateAgentIMChannelCredentialResponse
func (client *Client) UpdateAgentIMChannelCredential(workspaceId *string, agentId *string, imChannelId *string, request *UpdateAgentIMChannelCredentialRequest) (_result *UpdateAgentIMChannelCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentIMChannelCredentialResponse{}
	_body, _err := client.UpdateAgentIMChannelCredentialWithOptions(workspaceId, agentId, imChannelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the business tags, version tags, and visibility scope of a specified AgentSpec. Fields not provided remain unchanged.
//
// Description:
//
// Updates the business tags, version tags, and visibility scope of a specified agent template. Fields not provided remain unchanged.
//
// @param tmpReq - UpdateAgentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentSpecResponse
func (client *Client) UpdateAgentSpecWithOptions(workspaceId *string, agentSpecName *string, tmpReq *UpdateAgentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAgentSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentSpec"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-specs/" + dara.PercentEncode(dara.StringValue(agentSpecName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the business tags, version tags, and visibility scope of a specified AgentSpec. Fields not provided remain unchanged.
//
// Description:
//
// Updates the business tags, version tags, and visibility scope of a specified agent template. Fields not provided remain unchanged.
//
// @param request - UpdateAgentSpecRequest
//
// @return UpdateAgentSpecResponse
func (client *Client) UpdateAgentSpec(workspaceId *string, agentSpecName *string, request *UpdateAgentSpecRequest) (_result *UpdateAgentSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentSpecResponse{}
	_body, _err := client.UpdateAgentSpecWithOptions(workspaceId, agentSpecName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the API key or service account key collection for an enabled QoderCLI connector. When serviceAccountKeys is submitted, the service uses it as the complete updated key collection to add, rename, rotate, and delete keys.
//
// Description:
//
// Updates the API key or service account key collection for an enabled QoderCLI connector.
//
// When updating service account keys, existing keys are matched by their record `id`.
//
// @param tmpReq - UpdateConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnectorWithOptions(workspaceId *string, connectorName *string, tmpReq *UpdateConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnector"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName)) + "/actions/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the API key or service account key collection for an enabled QoderCLI connector. When serviceAccountKeys is submitted, the service uses it as the complete updated key collection to add, rename, rotate, and delete keys.
//
// Description:
//
// Updates the API key or service account key collection for an enabled QoderCLI connector.
//
// When updating service account keys, existing keys are matched by their record `id`.
//
// @param request - UpdateConnectorRequest
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnector(workspaceId *string, connectorName *string, request *UpdateConnectorRequest) (_result *UpdateConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConnectorResponse{}
	_body, _err := client.UpdateConnectorWithOptions(workspaceId, connectorName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the content, description, or MCP header application scope of a specified credential. Unspecified configurations retain their original values. The credential name and type cannot be modified.
//
// Description:
//
// Updates the content, description, or MCP header application scope of a specified credential.
//
// ### Update rules
//
// You must modify at least one of the credential content, description, or application scope. When modifying only the application scope, you do not need to submit credentialMetadata or description.
//
// - Omit credentialMetadata: The original credential content is retained. If provided, submit a JSON string according to the original credential type.
//
// - Omit description: The original description is retained.
//
// - The credential name and credentialType cannot be modified.
//
// ### MCP header application scope
//
// Only mcpHeader supports updating the application scope. For apiKey, do not submit resourceScope or resourceRefs.
//
// - Omit both the scope and the resource list: The original configuration is retained.
//
// - Submit resourceRefs: You must also submit resourceScope.
//
// - resourceScope=ALL: Omit resourceRefs.
//
// - resourceScope=SPECIFIED: Submit resourceRefs. An empty array [] is allowed during updates. For each item, specify resourceType=mcpServer and resourceId.
//
// ### Query results
//
// After the update, call [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential) to query the current configuration and the masked credential content.
//
// @param tmpReq - UpdateCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredentialWithOptions(workspaceId *string, credentialId *string, tmpReq *UpdateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCredentialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredential"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/credentials/" + dara.PercentEncode(dara.StringValue(credentialId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the content, description, or MCP header application scope of a specified credential. Unspecified configurations retain their original values. The credential name and type cannot be modified.
//
// Description:
//
// Updates the content, description, or MCP header application scope of a specified credential.
//
// ### Update rules
//
// You must modify at least one of the credential content, description, or application scope. When modifying only the application scope, you do not need to submit credentialMetadata or description.
//
// - Omit credentialMetadata: The original credential content is retained. If provided, submit a JSON string according to the original credential type.
//
// - Omit description: The original description is retained.
//
// - The credential name and credentialType cannot be modified.
//
// ### MCP header application scope
//
// Only mcpHeader supports updating the application scope. For apiKey, do not submit resourceScope or resourceRefs.
//
// - Omit both the scope and the resource list: The original configuration is retained.
//
// - Submit resourceRefs: You must also submit resourceScope.
//
// - resourceScope=ALL: Omit resourceRefs.
//
// - resourceScope=SPECIFIED: Submit resourceRefs. An empty array [] is allowed during updates. For each item, specify resourceType=mcpServer and resourceId.
//
// ### Query results
//
// After the update, call [GetCredential](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCredential) to query the current configuration and the masked credential content.
//
// @param request - UpdateCredentialRequest
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredential(workspaceId *string, credentialId *string, request *UpdateCredentialRequest) (_result *UpdateCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCredentialResponse{}
	_body, _err := client.UpdateCredentialWithOptions(workspaceId, credentialId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a custom code or image agent in a specified workspace and executes the update asynchronously. Omitted top-level fields retain their original configurations. You can call GetCustomAgent to query the update result.
//
// Description:
//
// Updates the deployment artifact or runtime configuration of a custom code or image agent to release new code, images, or configurations.
//
// ### Update rules
//
// Updates a custom code or image agent in a specified workspace. You must provide at least one top-level field to update. Omitted top-level fields retain their original configurations.
//
// The update rules for each configuration are as follows.
//
// | Configuration | Update method and precautions |
//
// | --- | --- |
//
// | `name` | Omit to retain the original name. The value cannot be `null` or an empty string. |
//
// | `description` | Omit to retain the original description. Pass `null` to clear the description. |
//
// | `artifact` | When submitting, specify `port` and the required fields for the selected source. To retain a custom startup command, also specify `entrypoint`. `code` and `container` are mutually exclusive. |
//
// | `runtime` | When submitting, specify both `compute` and `sessionPolicy`. To retain the elastic policy, submit the complete `hpa`. |
//
// | `environment` | The environment variable list is replaced as a whole. You must include all variables that need to be retained. |
//
// | `network` | When submitting, omitted access items are treated as allowed. To retain access restrictions, specify both `accessInternet` and `accessVpc`. |
//
// | `healthCheck` | When submitting, you must specify `port`. Other unspecified configurations use default values. Pass `null` to remove the health check. |
//
// | `logging` | When submitting `logging.sls`, specify non-empty `project` and `logstore` at the same time. |
//
// | `protocol`, `ossMounts`, `agenticFsMounts` | When submitting, the corresponding list is replaced as a whole. Pass `[]` to clear the corresponding list. `agenticFsMounts` does not accept `null`. |
//
// ### Deployment artifacts
//
// Three deployment methods are supported: code packages, ACR images, and custom images. When updating or switching the deployment method, provide the configuration according to the target method.
//
// | Deployment method | Configuration requirements |
//
// | --- | --- |
//
// | ACR image | When updating the image tag, also provide `artifact.port`, as well as `type=acr`, `instanceId`, `namespace`, `repo`, and `tag` in `artifact.container`. To retain a custom startup command, also provide `entrypoint`. |
//
// | Custom image | Set `artifact.container.type=custom`, provide the complete `customImage`, and provide `clientToken` in the query parameters. When switching to this method, also specify `artifact.port`. Do not include ACR instance, namespace, repository, and tag fields in `container`. |
//
// | Code package | Provide `artifact.port`, as well as `sourceUri`, `languageRuntime`, and `entrypoint` in `artifact.code`. Do not commit `container` at the same time. |
//
// **Custom image**
//
// - `customImage` must include the complete image address `image` and the independent pull network `pullNetwork`. The image address does not include a protocol or user information, and supports addresses with tags or digests.
//
// - For public network pulling, set `pullNetwork.type=public` and do not include VPC fields. For VPC pulling, set `type=vpc` and provide `vpcId`, `vSwitchId`, and `securityGroupId` at the same time. The image pull network and the agent runtime network are configured separately.
//
// - When repository authentication is required, provide non-empty `username` and `password` at the same time. For anonymous access, omit both. The response does not return the password.
//
// - Omitting the entire `customImage` retains the saved image source. Commit the complete object when updating the custom image source. You do not need to resubmit this object when only adjusting the runtime configuration.
//
// - Custom images and code packages do not support enabling `useImageCache`. The value in the final configuration must be `false`. When this field is omitted for a custom image, it is treated as `false`. When switching to a code package, if the original configuration is `true`, you must explicitly change it to `false`. When updating other sources, omitting this field retains the original value.
//
// **Code package**
//
// - Supports ZIP or TAR.GZ. The compressed package size cannot exceed `500000000` bytes. Dependencies are packaged by the invoker. The platform does not execute online builds or install dependencies.
//
// - Query the active `languageRuntime` values by invoking [ListAgentCodeRuntimes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListAgentCodeRuntimes).
//
// - `sourceUri` supports `oss://` addresses or HTTPS object addresses of standard OSS public domain name, including signed URLs.
//
// - For local code packages, first invoke [GetAgentCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageUploadUrl), upload the file according to the response information, and then commit the returned `sourceUri`.
//
// - The new source is verified for read permissions, objects, and size before acceptance, and is transferred to a fixed copy on the platform. If the transfer failed, the current agent configuration is not saved. Common updates can reuse the saved fixed code package address.
//
// ### Runtime configuration and elasticity
//
// Omitting the entire `runtime` retains the original runtime configuration. When submitting `runtime`, you must specify both `compute` and `sessionPolicy`.
//
// If you submit `runtime` but omit `hpa`, or omit some fields in `hpa`, the unspecified items use the following default values. **To retain the existing elastic policy, submit the complete `hpa`.**
//
// | Field | Default value | Description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether elastic scaling is enabled. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances, which cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. |
//
// The upper and lower limits of the number of instances are configuration values. Query the current number of instances through `sandboxPhaseCounts`.
//
// When `sessionPolicy.type=ISOLATED_HEADER_FIELD`, you must specify a valid HTTP request header name `headerName`. Do not specify `headerName` for other session policies.
//
// ### Network and mounts
//
// Omitting the entire `network` retains the original configuration. When submitting this object, omitted access items are treated as allowed on the agent side, and actual access is restricted by the workspace network configuration.
//
// To retain the original restrictions, submit both `accessInternet.enabled` and `accessVpc.enabled` at the same time. Do not submit only one of them.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. When updating either list, you must include all mount items that need to be retained in the list. The two lists are updated separately. You can use `agenticFsMounts=[]` to clear AgenticFS mounts, but you cannot use `null` to clear them.
//
// In this operation, ACR images, custom images, and code packages all use the same user mount target rules: `mountPath` is the absolute path in the container, such as `/data/train`, and must not conflict with platform reserved directories. Mount targets must not be duplicated or form parent-child directory overwrites.
//
// ### Environment variables
//
// Omitting the entire `environment` retains the original configuration. When submitting `environment.variables`, the original list is replaced as a whole with the current list. You must include all variables that need to be retained.
//
// The following methods all clear the environment variables configured by the user:
//
// - `environment.variables=[]`.
//
// - `environment=null`.
//
// - Submit `environment` but omit `variables`, such as `environment={}`.
//
// Variable names must be non-empty and unique in the list. Variable values must be strings, and empty strings are allowed.
//
// ### Health checks and logs
//
// Omitting the entire `healthCheck` retains the original configuration. Pass `null` to remove the health check. When submitting this object, you must specify `port`. Other unspecified items use the following default values.
//
// | Field | Default value |
//
// | --- | --- |
//
// | `path` | `/healthz` |
//
// | `intervalSeconds` | `10` seconds |
//
// | `timeoutSeconds` | `2` seconds |
//
// | `unhealthyThreshold` | `3` times |
//
// The health check port ranges from 1 to 65535. Health Check Interval, timeout period, and failure threshold must be positive integers, and the timeout period must be less than Health Check Interval. Do not submit `healthyThreshold`.
//
// When configuring log delivery, `logging.sls.project` and `logging.sls.logstore` must be provided at the same time and cannot be empty.
//
// ### Asynchronous results and retries
//
// `success=true` indicates that the update request has been accepted. Call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query the configuration processing result.
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest expected configuration version. |
//
// | `effectiveSpecVersion` | The effective configuration version. |
//
// | `latestVersionStatus` | The processing status of the latest expected configuration version. |
//
// | `latestExecutionError` | The reason why the latest configuration is waiting for a retry or fails to be executed, and whether to automatically retry. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | A snapshot of the number of instances counted by phase. |
//
// When `latestVersionStatus` is `pending`, `processing`, or `waiting_retry`, you can continue to query later. `succeeded` indicates that the configuration is processed successfully. When the value is `failed`, troubleshoot the configuration based on `latestExecutionError`.
//
// When explicitly submitting `customImage` or explicitly setting `artifact.container.type=custom`, you must provide `clientToken` in the query parameters. You can also optionally pass this parameter when updating ACR images or code packages to reuse the acceptance result of the same submission. The identifier is 1 to 128 characters in length and can contain letters, digits, periods, underscores, colons, and hyphens.
//
// When retrying the same submission, reuse the original `clientToken`. The service returns the existing acceptance result based on the tenant, workspace, and this identifier. Use a new identifier when modifying the submission content.
//
// @param tmpReq - UpdateCustomAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomAgentResponse
func (client *Client) UpdateCustomAgentWithOptions(workspaceId *string, agentId *string, tmpReq *UpdateCustomAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/custom-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a custom code or image agent in a specified workspace and executes the update asynchronously. Omitted top-level fields retain their original configurations. You can call GetCustomAgent to query the update result.
//
// Description:
//
// Updates the deployment artifact or runtime configuration of a custom code or image agent to release new code, images, or configurations.
//
// ### Update rules
//
// Updates a custom code or image agent in a specified workspace. You must provide at least one top-level field to update. Omitted top-level fields retain their original configurations.
//
// The update rules for each configuration are as follows.
//
// | Configuration | Update method and precautions |
//
// | --- | --- |
//
// | `name` | Omit to retain the original name. The value cannot be `null` or an empty string. |
//
// | `description` | Omit to retain the original description. Pass `null` to clear the description. |
//
// | `artifact` | When submitting, specify `port` and the required fields for the selected source. To retain a custom startup command, also specify `entrypoint`. `code` and `container` are mutually exclusive. |
//
// | `runtime` | When submitting, specify both `compute` and `sessionPolicy`. To retain the elastic policy, submit the complete `hpa`. |
//
// | `environment` | The environment variable list is replaced as a whole. You must include all variables that need to be retained. |
//
// | `network` | When submitting, omitted access items are treated as allowed. To retain access restrictions, specify both `accessInternet` and `accessVpc`. |
//
// | `healthCheck` | When submitting, you must specify `port`. Other unspecified configurations use default values. Pass `null` to remove the health check. |
//
// | `logging` | When submitting `logging.sls`, specify non-empty `project` and `logstore` at the same time. |
//
// | `protocol`, `ossMounts`, `agenticFsMounts` | When submitting, the corresponding list is replaced as a whole. Pass `[]` to clear the corresponding list. `agenticFsMounts` does not accept `null`. |
//
// ### Deployment artifacts
//
// Three deployment methods are supported: code packages, ACR images, and custom images. When updating or switching the deployment method, provide the configuration according to the target method.
//
// | Deployment method | Configuration requirements |
//
// | --- | --- |
//
// | ACR image | When updating the image tag, also provide `artifact.port`, as well as `type=acr`, `instanceId`, `namespace`, `repo`, and `tag` in `artifact.container`. To retain a custom startup command, also provide `entrypoint`. |
//
// | Custom image | Set `artifact.container.type=custom`, provide the complete `customImage`, and provide `clientToken` in the query parameters. When switching to this method, also specify `artifact.port`. Do not include ACR instance, namespace, repository, and tag fields in `container`. |
//
// | Code package | Provide `artifact.port`, as well as `sourceUri`, `languageRuntime`, and `entrypoint` in `artifact.code`. Do not commit `container` at the same time. |
//
// **Custom image**
//
// - `customImage` must include the complete image address `image` and the independent pull network `pullNetwork`. The image address does not include a protocol or user information, and supports addresses with tags or digests.
//
// - For public network pulling, set `pullNetwork.type=public` and do not include VPC fields. For VPC pulling, set `type=vpc` and provide `vpcId`, `vSwitchId`, and `securityGroupId` at the same time. The image pull network and the agent runtime network are configured separately.
//
// - When repository authentication is required, provide non-empty `username` and `password` at the same time. For anonymous access, omit both. The response does not return the password.
//
// - Omitting the entire `customImage` retains the saved image source. Commit the complete object when updating the custom image source. You do not need to resubmit this object when only adjusting the runtime configuration.
//
// - Custom images and code packages do not support enabling `useImageCache`. The value in the final configuration must be `false`. When this field is omitted for a custom image, it is treated as `false`. When switching to a code package, if the original configuration is `true`, you must explicitly change it to `false`. When updating other sources, omitting this field retains the original value.
//
// **Code package**
//
// - Supports ZIP or TAR.GZ. The compressed package size cannot exceed `500000000` bytes. Dependencies are packaged by the invoker. The platform does not execute online builds or install dependencies.
//
// - Query the active `languageRuntime` values by invoking [ListAgentCodeRuntimes](https://api.aliyun.com/document/AgentCore/2026-08-04/ListAgentCodeRuntimes).
//
// - `sourceUri` supports `oss://` addresses or HTTPS object addresses of standard OSS public domain name, including signed URLs.
//
// - For local code packages, first invoke [GetAgentCodePackageUploadUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageUploadUrl), upload the file according to the response information, and then commit the returned `sourceUri`.
//
// - The new source is verified for read permissions, objects, and size before acceptance, and is transferred to a fixed copy on the platform. If the transfer failed, the current agent configuration is not saved. Common updates can reuse the saved fixed code package address.
//
// ### Runtime configuration and elasticity
//
// Omitting the entire `runtime` retains the original runtime configuration. When submitting `runtime`, you must specify both `compute` and `sessionPolicy`.
//
// If you submit `runtime` but omit `hpa`, or omit some fields in `hpa`, the unspecified items use the following default values. **To retain the existing elastic policy, submit the complete `hpa`.**
//
// | Field | Default value | Description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether elastic scaling is enabled. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances, which cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. |
//
// The upper and lower limits of the number of instances are configuration values. Query the current number of instances through `sandboxPhaseCounts`.
//
// When `sessionPolicy.type=ISOLATED_HEADER_FIELD`, you must specify a valid HTTP request header name `headerName`. Do not specify `headerName` for other session policies.
//
// ### Network and mounts
//
// Omitting the entire `network` retains the original configuration. When submitting this object, omitted access items are treated as allowed on the agent side, and actual access is restricted by the workspace network configuration.
//
// To retain the original restrictions, submit both `accessInternet.enabled` and `accessVpc.enabled` at the same time. Do not submit only one of them.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. When updating either list, you must include all mount items that need to be retained in the list. The two lists are updated separately. You can use `agenticFsMounts=[]` to clear AgenticFS mounts, but you cannot use `null` to clear them.
//
// In this operation, ACR images, custom images, and code packages all use the same user mount target rules: `mountPath` is the absolute path in the container, such as `/data/train`, and must not conflict with platform reserved directories. Mount targets must not be duplicated or form parent-child directory overwrites.
//
// ### Environment variables
//
// Omitting the entire `environment` retains the original configuration. When submitting `environment.variables`, the original list is replaced as a whole with the current list. You must include all variables that need to be retained.
//
// The following methods all clear the environment variables configured by the user:
//
// - `environment.variables=[]`.
//
// - `environment=null`.
//
// - Submit `environment` but omit `variables`, such as `environment={}`.
//
// Variable names must be non-empty and unique in the list. Variable values must be strings, and empty strings are allowed.
//
// ### Health checks and logs
//
// Omitting the entire `healthCheck` retains the original configuration. Pass `null` to remove the health check. When submitting this object, you must specify `port`. Other unspecified items use the following default values.
//
// | Field | Default value |
//
// | --- | --- |
//
// | `path` | `/healthz` |
//
// | `intervalSeconds` | `10` seconds |
//
// | `timeoutSeconds` | `2` seconds |
//
// | `unhealthyThreshold` | `3` times |
//
// The health check port ranges from 1 to 65535. Health Check Interval, timeout period, and failure threshold must be positive integers, and the timeout period must be less than Health Check Interval. Do not submit `healthyThreshold`.
//
// When configuring log delivery, `logging.sls.project` and `logging.sls.logstore` must be provided at the same time and cannot be empty.
//
// ### Asynchronous results and retries
//
// `success=true` indicates that the update request has been accepted. Call [GetCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetCustomAgent) to query the configuration processing result.
//
// | Field | Purpose |
//
// | --- | --- |
//
// | `latestSpecVersion` | The latest expected configuration version. |
//
// | `effectiveSpecVersion` | The effective configuration version. |
//
// | `latestVersionStatus` | The processing status of the latest expected configuration version. |
//
// | `latestExecutionError` | The reason why the latest configuration is waiting for a retry or fails to be executed, and whether to automatically retry. |
//
// | `status` | The overall status of the agent. |
//
// | `sandboxPhaseCounts` | A snapshot of the number of instances counted by phase. |
//
// When `latestVersionStatus` is `pending`, `processing`, or `waiting_retry`, you can continue to query later. `succeeded` indicates that the configuration is processed successfully. When the value is `failed`, troubleshoot the configuration based on `latestExecutionError`.
//
// When explicitly submitting `customImage` or explicitly setting `artifact.container.type=custom`, you must provide `clientToken` in the query parameters. You can also optionally pass this parameter when updating ACR images or code packages to reuse the acceptance result of the same submission. The identifier is 1 to 128 characters in length and can contain letters, digits, periods, underscores, colons, and hyphens.
//
// When retrying the same submission, reuse the original `clientToken`. The service returns the existing acceptance result based on the tenant, workspace, and this identifier. Use a new identifier when modifying the submission content.
//
// @param request - UpdateCustomAgentRequest
//
// @return UpdateCustomAgentResponse
func (client *Client) UpdateCustomAgent(workspaceId *string, agentId *string, request *UpdateCustomAgentRequest) (_result *UpdateCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCustomAgentResponse{}
	_body, _err := client.UpdateCustomAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the model source, instructions, tools, and skills of a managed agent, and processes the configuration asynchronously.
//
// Description:
//
// Updates the model, instruction, tool, or skill configurations saved on the platform side for a managed agent to adjust its capability configurations after integration.
//
// ### Update rules
//
// Provide at least one updatable field. Omitted top-level fields retain their original configurations. You can set description=null to clear the description. The tools and skills fields are replaced as a whole upon submission. You must include the items to be retained, or pass [] to clear the corresponding explicit lists.
//
// ### Switch model sources
//
// - To switch to PLATFORM, provide modelSource=PLATFORM and a valid model configuration, including the model connection ID and model name.
//
// - To switch to RUNTIME, change the Settings to modelSource=RUNTIME and do not commit the platform model. The original platform model configuration is removed when the source is switched.
//
// When a template is used, configurations belonging to the template such as instructions are read from the template. A successful update indicates that the request is accepted. Invoke [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent) to query the configuration version and external accessed state. The clientToken parameter does not provide idempotence guarantees at this time.
//
// @param tmpReq - UpdateExternalAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExternalAgentResponse
func (client *Client) UpdateExternalAgentWithOptions(workspaceId *string, agentId *string, tmpReq *UpdateExternalAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExternalAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateExternalAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExternalAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/external-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExternalAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the model source, instructions, tools, and skills of a managed agent, and processes the configuration asynchronously.
//
// Description:
//
// Updates the model, instruction, tool, or skill configurations saved on the platform side for a managed agent to adjust its capability configurations after integration.
//
// ### Update rules
//
// Provide at least one updatable field. Omitted top-level fields retain their original configurations. You can set description=null to clear the description. The tools and skills fields are replaced as a whole upon submission. You must include the items to be retained, or pass [] to clear the corresponding explicit lists.
//
// ### Switch model sources
//
// - To switch to PLATFORM, provide modelSource=PLATFORM and a valid model configuration, including the model connection ID and model name.
//
// - To switch to RUNTIME, change the Settings to modelSource=RUNTIME and do not commit the platform model. The original platform model configuration is removed when the source is switched.
//
// When a template is used, configurations belonging to the template such as instructions are read from the template. A successful update indicates that the request is accepted. Invoke [GetExternalAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetExternalAgent) to query the configuration version and external accessed state. The clientToken parameter does not provide idempotence guarantees at this time.
//
// @param request - UpdateExternalAgentRequest
//
// @return UpdateExternalAgentResponse
func (client *Client) UpdateExternalAgent(workspaceId *string, agentId *string, request *UpdateExternalAgentRequest) (_result *UpdateExternalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExternalAgentResponse{}
	_body, _err := client.UpdateExternalAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the logon configuration of a bound identity provider in a specified workspace. DingTalk and Feishu support the logon switch, incremental synchronization switch, application configuration, and synchronization scope. CustomSSO supports SAML configuration and the logon switch. Updates take effect asynchronously. Call GetIdentityProvider to query the status.
//
// Description:
//
// Modifies the logon or directory configuration of a bound enterprise identity provider to update enterprise application credentials, adjust synchronized departments, or configure custom single sign-on.
//
// ### Select the configuration to modify
//
// | Identity provider | Updatable content |
//
// | --- | --- |
//
// | DingTalk (DingTalk), Feishu (Feishu) | Enterprise application credentials metadata, logon switch loginEnabled, directory change synchronization switch syncEnabled, and synchronization scope syncScope. You must submit at least one item. |
//
// | Custom identity provider (CustomSSO) | SAML protocol protocol, logon configuration saml, and logon switch loginEnabled. SCIM-based user synchronization is configured separately by using [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM). |
//
// ### Update DingTalk or Feishu
//
// metadata is updated by field: submit only the fields that need to be modified. Omitted fields and existing keys remain unchanged. Provided values must be non-empty strings. When you enable directory change synchronization, the configuration must include encryptKey and verificationToken.
//
// To adjust synchronized departments:
//
// **Step 1:*	- Call [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments) to query and select departments.
//
// **Step 2:*	- Save syncScope by using this operation. If you omit this field or pass null, the current scope is retained.
//
// **Step 3:*	- Call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to wait for the configuration processing to complete, and then call [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) to synchronize users based on the new scope.
//
// Saving the scope does not automatically trigger a full synchronization. syncEnabled controls subsequent directory change events. The initial synchronization and manual synchronization are executed separately.
//
// ### Configure custom single sign-on
//
// CustomSSO uses the SAML protocol to allow users to log on through the enterprise identity system. First, call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the access information of AgentCore and configure it in the enterprise identity system. Then, fill in the enterprise-side information in saml.
//
// Choose one of the following two methods:
//
// | Method | Required configuration |
//
// | --- | --- |
//
// | Import metadata | metadataXml: a single IdP metadata XML text exported from the enterprise identity provider, with a maximum size of 1 MiB. Enter the XML content instead of a download URL or file path. |
//
// | Manual entry | Provide entityId, loginUrl, bindingType, and x509Certificates at the same time. Use an HTTPS logon URL and a valid signing certificate. Set bindingType to HTTP-Redirect or HTTP-POST. You can provide up to two certificates. |
//
// The entityId cannot be changed after it is saved. If the XML contains both binding methods, HTTP-POST and its corresponding URL take precedence.
//
// After saving the configuration, set loginEnabled=true to enable logon. You can also enable it when submitting a valid SAML configuration. Setting it to false disables logon and retains the existing configuration.
//
// ### Confirm that the configuration takes effect
//
// After the request is accepted, UPDATING is usually returned. Call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to query the progress. READY indicates that the configuration processing is complete. If it fails, check the returned status and configuration.
//
// If the request times out, query the current configuration and processing status first, and then decide whether to retry. Query the SCIM processing result by calling [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM).
//
// @param tmpReq - UpdateIdentityProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithOptions(workspaceId *string, identityProviderType *string, tmpReq *UpdateIdentityProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProvider"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/identity-providers/" + dara.PercentEncode(dara.StringValue(identityProviderType))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the logon configuration of a bound identity provider in a specified workspace. DingTalk and Feishu support the logon switch, incremental synchronization switch, application configuration, and synchronization scope. CustomSSO supports SAML configuration and the logon switch. Updates take effect asynchronously. Call GetIdentityProvider to query the status.
//
// Description:
//
// Modifies the logon or directory configuration of a bound enterprise identity provider to update enterprise application credentials, adjust synchronized departments, or configure custom single sign-on.
//
// ### Select the configuration to modify
//
// | Identity provider | Updatable content |
//
// | --- | --- |
//
// | DingTalk (DingTalk), Feishu (Feishu) | Enterprise application credentials metadata, logon switch loginEnabled, directory change synchronization switch syncEnabled, and synchronization scope syncScope. You must submit at least one item. |
//
// | Custom identity provider (CustomSSO) | SAML protocol protocol, logon configuration saml, and logon switch loginEnabled. SCIM-based user synchronization is configured separately by using [UpdateIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateIdentityProviderSCIM). |
//
// ### Update DingTalk or Feishu
//
// metadata is updated by field: submit only the fields that need to be modified. Omitted fields and existing keys remain unchanged. Provided values must be non-empty strings. When you enable directory change synchronization, the configuration must include encryptKey and verificationToken.
//
// To adjust synchronized departments:
//
// **Step 1:*	- Call [ListIdentityProviderDepartments](https://api.aliyun.com/document/AgentCore/2026-08-04/ListIdentityProviderDepartments) to query and select departments.
//
// **Step 2:*	- Save syncScope by using this operation. If you omit this field or pass null, the current scope is retained.
//
// **Step 3:*	- Call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to wait for the configuration processing to complete, and then call [CreateIdentityProviderSyncJob](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateIdentityProviderSyncJob) to synchronize users based on the new scope.
//
// Saving the scope does not automatically trigger a full synchronization. syncEnabled controls subsequent directory change events. The initial synchronization and manual synchronization are executed separately.
//
// ### Configure custom single sign-on
//
// CustomSSO uses the SAML protocol to allow users to log on through the enterprise identity system. First, call [GetIdentityProviderServiceProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderServiceProvider) to obtain the access information of AgentCore and configure it in the enterprise identity system. Then, fill in the enterprise-side information in saml.
//
// Choose one of the following two methods:
//
// | Method | Required configuration |
//
// | --- | --- |
//
// | Import metadata | metadataXml: a single IdP metadata XML text exported from the enterprise identity provider, with a maximum size of 1 MiB. Enter the XML content instead of a download URL or file path. |
//
// | Manual entry | Provide entityId, loginUrl, bindingType, and x509Certificates at the same time. Use an HTTPS logon URL and a valid signing certificate. Set bindingType to HTTP-Redirect or HTTP-POST. You can provide up to two certificates. |
//
// The entityId cannot be changed after it is saved. If the XML contains both binding methods, HTTP-POST and its corresponding URL take precedence.
//
// After saving the configuration, set loginEnabled=true to enable logon. You can also enable it when submitting a valid SAML configuration. Setting it to false disables logon and retains the existing configuration.
//
// ### Confirm that the configuration takes effect
//
// After the request is accepted, UPDATING is usually returned. Call [GetIdentityProvider](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProvider) to query the progress. READY indicates that the configuration processing is complete. If it fails, check the returned status and configuration.
//
// If the request times out, query the current configuration and processing status first, and then decide whether to retry. Query the SCIM processing result by calling [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM).
//
// @param request - UpdateIdentityProviderRequest
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProvider(workspaceId *string, identityProviderType *string, request *UpdateIdentityProviderRequest) (_result *UpdateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.UpdateIdentityProviderWithOptions(workspaceId, identityProviderType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables SCIM user synchronization independently. When enabled, an initial full import is executed asynchronously, and the result can be queried by calling GetIdentityProviderSCIM. The SAML logon configuration is retained.
//
// Description:
//
// Enables or disables SCIM user synchronization for AgentCore. When enabled, all existing users in the shared user pool are fully imported into the workspace. Users in the user pool are provisioned by the enterprise identity system through SCIM.
//
// ### Preparations
//
// **Step 1:*	- Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to enable the receiving endpoint and obtain the receiving URL and a valid token.
//
// **Step 2:*	- Complete the SCIM configuration, connection test, and user assignment in the enterprise identity system.
//
// **Step 3:*	- Call this operation and set syncEnabled to true.
//
// ```http
//
// PUT /workspaces/{workspaceId}/scim-provisioning
//
// {"syncEnabled":true}
//
// ```
//
// When the enabling task is executed, the system checks whether the receiving endpoint is enabled and whether valid credentials exist. If no users are provisioned in the enterprise identity system, the initial import may succeed but contain no user records.
//
// ### Impact of disabling
//
// Setting syncEnabled to false disables user synchronization and the upstream SCIM receiving endpoint, but retains the imported users, existing credentials, and SAML logon configuration.
//
// To re-enable the feature later, call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to prepare the receiving endpoint first, and then enable synchronization.
//
// ### Verify the operation result
//
// After this operation accepts the request, it executes the enabling or disabling task in the background. Call [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM) to query the progress:
//
// - syncEnabled indicates whether synchronization is expected to be enabled or disabled.
//
// - operationStatus indicates the execution result of the most recent operation.
//
// - The status field in the query operation comprehensively indicates the current stage of SCIM, such as pending configuration, processing, available, or disabled.
//
// If the operation fails, the saved syncEnabled value is not automatically rolled back. Verify the configuration based on the operation result before deciding on subsequent actions.
//
// @param tmpReq - UpdateIdentityProviderSCIMRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderSCIMResponse
func (client *Client) UpdateIdentityProviderSCIMWithOptions(workspaceId *string, tmpReq *UpdateIdentityProviderSCIMRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderSCIMResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIdentityProviderSCIMShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProviderSCIM"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/scim-provisioning"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityProviderSCIMResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables SCIM user synchronization independently. When enabled, an initial full import is executed asynchronously, and the result can be queried by calling GetIdentityProviderSCIM. The SAML logon configuration is retained.
//
// Description:
//
// Enables or disables SCIM user synchronization for AgentCore. When enabled, all existing users in the shared user pool are fully imported into the workspace. Users in the user pool are provisioned by the enterprise identity system through SCIM.
//
// ### Preparations
//
// **Step 1:*	- Call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to enable the receiving endpoint and obtain the receiving URL and a valid token.
//
// **Step 2:*	- Complete the SCIM configuration, connection test, and user assignment in the enterprise identity system.
//
// **Step 3:*	- Call this operation and set syncEnabled to true.
//
// ```http
//
// PUT /workspaces/{workspaceId}/scim-provisioning
//
// {"syncEnabled":true}
//
// ```
//
// When the enabling task is executed, the system checks whether the receiving endpoint is enabled and whether valid credentials exist. If no users are provisioned in the enterprise identity system, the initial import may succeed but contain no user records.
//
// ### Impact of disabling
//
// Setting syncEnabled to false disables user synchronization and the upstream SCIM receiving endpoint, but retains the imported users, existing credentials, and SAML logon configuration.
//
// To re-enable the feature later, call [PrepareIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/PrepareIdentityProviderSCIM) to prepare the receiving endpoint first, and then enable synchronization.
//
// ### Verify the operation result
//
// After this operation accepts the request, it executes the enabling or disabling task in the background. Call [GetIdentityProviderSCIM](https://api.aliyun.com/document/AgentCore/2026-08-04/GetIdentityProviderSCIM) to query the progress:
//
// - syncEnabled indicates whether synchronization is expected to be enabled or disabled.
//
// - operationStatus indicates the execution result of the most recent operation.
//
// - The status field in the query operation comprehensively indicates the current stage of SCIM, such as pending configuration, processing, available, or disabled.
//
// If the operation fails, the saved syncEnabled value is not automatically rolled back. Verify the configuration based on the operation result before deciding on subsequent actions.
//
// @param request - UpdateIdentityProviderSCIMRequest
//
// @return UpdateIdentityProviderSCIMResponse
func (client *Client) UpdateIdentityProviderSCIM(workspaceId *string, request *UpdateIdentityProviderSCIMRequest) (_result *UpdateIdentityProviderSCIMResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIdentityProviderSCIMResponse{}
	_body, _err := client.UpdateIdentityProviderSCIMWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the model, instructions, skills, and runtime configuration of a managed Harness agent and executes the update asynchronously. Omitted top-level fields retain their original configurations.
//
// Description:
//
// Updates the model, instructions, skills, or runtime configuration of a managed Harness agent to adjust the capabilities and runtime behavior of an existing agent.
//
// ### Update rules
//
// Provide at least one updatable field. Omitted top-level fields retain their original configurations. You can set `description=null` to clear the description.
//
// | Configuration | Update method |
//
// | --- | --- |
//
// | `harness` | The type cannot be changed. When updating the qodercli account binding, keep the original `type` and submit the corresponding `configuration`. |
//
// | `model` | Provide a model configuration that matches the current Harness. For qwenpaw, specify the model connection ID and model name. For qodercli, use the available model identifier of the connector. |
//
// | `runtime` | Specify both `compute` and `sessionPolicy` when submitting. Omitting the entire object retains the original runtime configuration. |
//
// | `tools`, `skills`, `promptVariables`, `subAgents` | The corresponding lists are entirely replaced upon submission. You must include the items to be retained. Pass `[]` to clear the corresponding explicit lists. |
//
// | `environment` | The variable list and credential reference list are entirely replaced respectively. Submit the items to be retained together. |
//
// | `network` | Specify both `accessInternet.enabled` and `accessVpc.enabled` when submitting. |
//
// | `ossMounts`, `agenticFsMounts` | Entirely replaced respectively. Pass `[]` to clear the corresponding lists. `agenticFsMounts` does not accept `null`. |
//
// ### Harnesses, templates, and skills
//
// | Harness | Model and connector configuration |
//
// | --- | --- |
//
// | `qwenpaw` | This type is used by default if `harness` is omitted during creation. Set `model.modelConnectionId` to the workspace model connection ID, and set `model.modelName` to the name of the model created under this connection. |
//
// | `qodercli` | The qodercli connector must be enabled in the workspace. Set `model.modelConnectionId` to the model identifier returned by [ListConnectorModels](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectorModels). `model.modelName` can be omitted. |
//
// `harness` is an object, such as `{"type":"qwenpaw"}`. For qodercli, set `harness.configuration.connectorServiceAccountKey` to the account key record ID returned by [ListConnectors](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectors), not the original service account key. This field can be omitted if there is only one key record, but must be specified if there are multiple records. `connectorServiceAccountName` is the account name. Do not specify these two binding fields for qwenpaw.
//
// `skills` uses the skill names in the workspace AI Registry. If `sourceType=REFERENCE`, versions are referenced by `versionSelector`: use `LABEL` with a tag name, or `VERSION` with a specific version number. If the selector is omitted, `LABEL/latest` is used. If `sourceType=STATIC`, static files bundled with the package are used, and no version selector is provided.
//
// For `template.aiRegistry`, specify the agent template name and an optional version. If the version is omitted during creation, the published version is resolved and fixed. When a template is used, `instruction`, `promptVariables`, and `subAgents` are derived from the template, and the corresponding fields in the current request do not take effect. Explicit skills and template skills are recorded separately. You can query the target and effective versions through `configuredSkills` and `skills`.
//
// ### Elasticity and instances
//
// If `runtime` is committed but `hpa` or some of its fields are omitted, the following default values are used for unspecified items. To retain the existing elastic policy, commit the complete `hpa`.
//
// | Field | Default value | Value and description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether elastic scaling is enabled. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. Valid values: 0 to 100. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If you only modify instructions, models, or skills, you can omit the entire `runtime` to keep the original runtime configuration unchanged. The minimum and maximum numbers of instances are the configuration border. You can query the actual number through `sandboxPhaseCounts`.
//
// ### Environment and mounts
//
// Omitting the entire `environment` retains the original configuration. Submitting `environment={}` or `environment=null` clears the environment configuration. Environment variable values must be strings, and variable names must be non-empty and unique.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. The mount target `mountPath` uses an absolute path in the container, such as `/data/train`. It cannot conflict with platform reserved directories, and cannot be duplicated or form parent-child directory overwrites.
//
// The `path` for OSS mounts must be a non-empty directory prefix in the bucket. For AgenticFS mounts, `server` is the AccessPoint domain name, without a protocol, port, or path. `path` is an existing non-empty relative directory in the AccessPoint. You cannot mount the root directory of the AccessPoint. If `readOnly` is omitted, the default value is `false`.
//
// ### Query update results
//
// A successful response indicates that the update request has been accepted. Call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query `latestSpecVersion`, `latestVersionStatus`, `status`, and `sandboxPhaseCounts` to confirm the configuration processing and instance status. `clientToken` does not provide idempotence guarantees at this time.
//
// @param tmpReq - UpdateManagedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateManagedAgentResponse
func (client *Client) UpdateManagedAgentWithOptions(workspaceId *string, agentId *string, tmpReq *UpdateManagedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateManagedAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateManagedAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateManagedAgent"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/managed-agents/" + dara.PercentEncode(dara.StringValue(agentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateManagedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the model, instructions, skills, and runtime configuration of a managed Harness agent and executes the update asynchronously. Omitted top-level fields retain their original configurations.
//
// Description:
//
// Updates the model, instructions, skills, or runtime configuration of a managed Harness agent to adjust the capabilities and runtime behavior of an existing agent.
//
// ### Update rules
//
// Provide at least one updatable field. Omitted top-level fields retain their original configurations. You can set `description=null` to clear the description.
//
// | Configuration | Update method |
//
// | --- | --- |
//
// | `harness` | The type cannot be changed. When updating the qodercli account binding, keep the original `type` and submit the corresponding `configuration`. |
//
// | `model` | Provide a model configuration that matches the current Harness. For qwenpaw, specify the model connection ID and model name. For qodercli, use the available model identifier of the connector. |
//
// | `runtime` | Specify both `compute` and `sessionPolicy` when submitting. Omitting the entire object retains the original runtime configuration. |
//
// | `tools`, `skills`, `promptVariables`, `subAgents` | The corresponding lists are entirely replaced upon submission. You must include the items to be retained. Pass `[]` to clear the corresponding explicit lists. |
//
// | `environment` | The variable list and credential reference list are entirely replaced respectively. Submit the items to be retained together. |
//
// | `network` | Specify both `accessInternet.enabled` and `accessVpc.enabled` when submitting. |
//
// | `ossMounts`, `agenticFsMounts` | Entirely replaced respectively. Pass `[]` to clear the corresponding lists. `agenticFsMounts` does not accept `null`. |
//
// ### Harnesses, templates, and skills
//
// | Harness | Model and connector configuration |
//
// | --- | --- |
//
// | `qwenpaw` | This type is used by default if `harness` is omitted during creation. Set `model.modelConnectionId` to the workspace model connection ID, and set `model.modelName` to the name of the model created under this connection. |
//
// | `qodercli` | The qodercli connector must be enabled in the workspace. Set `model.modelConnectionId` to the model identifier returned by [ListConnectorModels](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectorModels). `model.modelName` can be omitted. |
//
// `harness` is an object, such as `{"type":"qwenpaw"}`. For qodercli, set `harness.configuration.connectorServiceAccountKey` to the account key record ID returned by [ListConnectors](https://api.aliyun.com/document/AgentCore/2026-08-04/ListConnectors), not the original service account key. This field can be omitted if there is only one key record, but must be specified if there are multiple records. `connectorServiceAccountName` is the account name. Do not specify these two binding fields for qwenpaw.
//
// `skills` uses the skill names in the workspace AI Registry. If `sourceType=REFERENCE`, versions are referenced by `versionSelector`: use `LABEL` with a tag name, or `VERSION` with a specific version number. If the selector is omitted, `LABEL/latest` is used. If `sourceType=STATIC`, static files bundled with the package are used, and no version selector is provided.
//
// For `template.aiRegistry`, specify the agent template name and an optional version. If the version is omitted during creation, the published version is resolved and fixed. When a template is used, `instruction`, `promptVariables`, and `subAgents` are derived from the template, and the corresponding fields in the current request do not take effect. Explicit skills and template skills are recorded separately. You can query the target and effective versions through `configuredSkills` and `skills`.
//
// ### Elasticity and instances
//
// If `runtime` is committed but `hpa` or some of its fields are omitted, the following default values are used for unspecified items. To retain the existing elastic policy, commit the complete `hpa`.
//
// | Field | Default value | Value and description |
//
// | --- | --- | --- |
//
// | `enabled` | `false` | Specifies whether elastic scaling is enabled. |
//
// | `minSandboxCount` | `1` | The minimum number of instances. Valid values: 0 to 100. |
//
// | `maxSandboxCount` | `10` | The maximum number of instances. Valid values: 0 to 100. The value cannot be less than the minimum number. |
//
// | `maxConcurrentSessionsPerSandbox` | `20` | The maximum number of concurrent sessions per instance. Valid values: 1 to 20. |
//
// | `sessionTtlSeconds` | `600` | The maximum idle time of a session, in seconds. Valid values: 60 to 86400. |
//
// If you only modify instructions, models, or skills, you can omit the entire `runtime` to keep the original runtime configuration unchanged. The minimum and maximum numbers of instances are the configuration border. You can query the actual number through `sandboxPhaseCounts`.
//
// ### Environment and mounts
//
// Omitting the entire `environment` retains the original configuration. Submitting `environment={}` or `environment=null` clears the environment configuration. Environment variable values must be strings, and variable names must be non-empty and unique.
//
// The total number of `ossMounts` and `agenticFsMounts` cannot exceed 10. The mount target `mountPath` uses an absolute path in the container, such as `/data/train`. It cannot conflict with platform reserved directories, and cannot be duplicated or form parent-child directory overwrites.
//
// The `path` for OSS mounts must be a non-empty directory prefix in the bucket. For AgenticFS mounts, `server` is the AccessPoint domain name, without a protocol, port, or path. `path` is an existing non-empty relative directory in the AccessPoint. You cannot mount the root directory of the AccessPoint. If `readOnly` is omitted, the default value is `false`.
//
// ### Query update results
//
// A successful response indicates that the update request has been accepted. Call [GetManagedAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/GetManagedAgent) to query `latestSpecVersion`, `latestVersionStatus`, `status`, and `sandboxPhaseCounts` to confirm the configuration processing and instance status. `clientToken` does not provide idempotence guarantees at this time.
//
// @param request - UpdateManagedAgentRequest
//
// @return UpdateManagedAgentResponse
func (client *Client) UpdateManagedAgent(workspaceId *string, agentId *string, request *UpdateManagedAgentRequest) (_result *UpdateManagedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateManagedAgentResponse{}
	_body, _err := client.UpdateManagedAgentWithOptions(workspaceId, agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update process is asynchronous.
//
// Description:
//
// Updates the address, description, authentication, or Swagger configuration of an MCP service. The service type and protocol cannot be modified after creation.
//
// The update process is executed asynchronously. After submission, call [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) to query the processing result. The `clientToken` parameter does not currently provide idempotency guarantees. If the request result is not confirmed, query the current configuration and status first.
//
// @param tmpReq - UpdateMcpRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpResponse
func (client *Client) UpdateMcpWithOptions(workspaceId *string, mcpServerId *string, tmpReq *UpdateMcpRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMcpShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcp"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update process is asynchronous.
//
// Description:
//
// Updates the address, description, authentication, or Swagger configuration of an MCP service. The service type and protocol cannot be modified after creation.
//
// The update process is executed asynchronously. After submission, call [GetMcp](https://api.aliyun.com/document/AgentCore/2026-08-04/GetMcp) to query the processing result. The `clientToken` parameter does not currently provide idempotency guarantees. If the request result is not confirmed, query the current configuration and status first.
//
// @param request - UpdateMcpRequest
//
// @return UpdateMcpResponse
func (client *Client) UpdateMcp(workspaceId *string, mcpServerId *string, request *UpdateMcpRequest) (_result *UpdateMcpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMcpResponse{}
	_body, _err := client.UpdateMcpWithOptions(workspaceId, mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates MCP parameters based on a template.
//
// Description:
//
// Updates parameters based on the marketplace template bound to the MCP service to adjust the configuration of the templated service.
//
// Updates the parameters exposed by the schema using the same template version bound during MCP creation, without upgrading the template version.
//
// @param tmpReq - UpdateMcpTemplateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpTemplateConfigResponse
func (client *Client) UpdateMcpTemplateConfigWithOptions(workspaceId *string, mcpServerId *string, tmpReq *UpdateMcpTemplateConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMcpTemplateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMcpTemplateConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["templateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpTemplateConfig"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId)) + "/template-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpTemplateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates MCP parameters based on a template.
//
// Description:
//
// Updates parameters based on the marketplace template bound to the MCP service to adjust the configuration of the templated service.
//
// Updates the parameters exposed by the schema using the same template version bound during MCP creation, without upgrading the template version.
//
// @param request - UpdateMcpTemplateConfigRequest
//
// @return UpdateMcpTemplateConfigResponse
func (client *Client) UpdateMcpTemplateConfig(workspaceId *string, mcpServerId *string, request *UpdateMcpTemplateConfigRequest) (_result *UpdateMcpTemplateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMcpTemplateConfigResponse{}
	_body, _err := client.UpdateMcpTemplateConfigWithOptions(workspaceId, mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the long-term memory text or metadata based on the memory ID.
//
// Description:
//
// Updates the long-term memory text or metadata based on the memory ID.
//
// @param tmpReq - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(workspaceId *string, memoryStoreName *string, memoryId *string, tmpReq *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMemoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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
// Updates the long-term memory text or metadata based on the memory ID.
//
// Description:
//
// Updates the long-term memory text or metadata based on the memory ID.
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(workspaceId *string, memoryStoreName *string, memoryId *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(workspaceId, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a memory store.
//
// Description:
//
// Updates the mutable configurations of a long-term memory store to adjust the settings of an existing memory store.
//
// @param tmpReq - UpdateMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryStoreResponse
func (client *Client) UpdateMemoryStoreWithOptions(workspaceId *string, memoryStoreName *string, tmpReq *UpdateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMemoryStoreShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemoryStore"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memorystores/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a memory store.
//
// Description:
//
// Updates the mutable configurations of a long-term memory store to adjust the settings of an existing memory store.
//
// @param request - UpdateMemoryStoreRequest
//
// @return UpdateMemoryStoreResponse
func (client *Client) UpdateMemoryStore(workspaceId *string, memoryStoreName *string, request *UpdateMemoryStoreRequest) (_result *UpdateMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryStoreResponse{}
	_body, _err := client.UpdateMemoryStoreWithOptions(workspaceId, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the description, context token limit, maximum output token count, or capability configuration of a specified model.
//
// Description:
//
// Updates the description, context length, maximum output length, or capability configuration of a model.
//
// ### Update rules
//
// - Provide at least one non-null field among description, contextSize, maxTokens, and capabilities.
//
// - Omitted fields or fields passed as null retain their original values.
//
// - The capabilities object is replaced as a whole. Unsubmitted capabilities within the object are treated as false.
//
// ### Impact on agents
//
// After contextSize, maxTokens, or capabilities actually change, the platform asynchronously refreshes the managed Harness agents that reference this model, as well as the managed agents that use platform models (modelSource=PLATFORM). Modifying only the description or repeatedly submitting the same configuration does not trigger a refresh.
//
// A successful response indicates that the model configuration has been saved. Verify the configuration by calling [GetModel](https://api.aliyun.com/document/AgentCore/2026-08-04/GetModel), and then check whether the configuration takes effect by calling the corresponding agent details API operation.
//
// @param tmpReq - UpdateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelResponse
func (client *Client) UpdateModelWithOptions(workspaceId *string, modelId *string, tmpReq *UpdateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModel"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description, context token limit, maximum output token count, or capability configuration of a specified model.
//
// Description:
//
// Updates the description, context length, maximum output length, or capability configuration of a model.
//
// ### Update rules
//
// - Provide at least one non-null field among description, contextSize, maxTokens, and capabilities.
//
// - Omitted fields or fields passed as null retain their original values.
//
// - The capabilities object is replaced as a whole. Unsubmitted capabilities within the object are treated as false.
//
// ### Impact on agents
//
// After contextSize, maxTokens, or capabilities actually change, the platform asynchronously refreshes the managed Harness agents that reference this model, as well as the managed agents that use platform models (modelSource=PLATFORM). Modifying only the description or repeatedly submitting the same configuration does not trigger a refresh.
//
// A successful response indicates that the model configuration has been saved. Verify the configuration by calling [GetModel](https://api.aliyun.com/document/AgentCore/2026-08-04/GetModel), and then check whether the configuration takes effect by calling the corresponding agent details API operation.
//
// @param request - UpdateModelRequest
//
// @return UpdateModelResponse
func (client *Client) UpdateModel(workspaceId *string, modelId *string, request *UpdateModelRequest) (_result *UpdateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelResponse{}
	_body, _err := client.UpdateModelWithOptions(workspaceId, modelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the mutable configurations of a specified model connection and resubmits the publishing task. The invocation protocol cannot be modified after creation.
//
// Description:
//
// Updates the mutable configurations of a specified model connection and resubmits the publishing task. The invocation protocol cannot be modified after creation.
//
// @param tmpReq - UpdateModelConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelConnectionResponse
func (client *Client) UpdateModelConnectionWithOptions(workspaceId *string, connectionId *string, tmpReq *UpdateModelConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateModelConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelConnection"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/model-connections/" + dara.PercentEncode(dara.StringValue(connectionId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the mutable configurations of a specified model connection and resubmits the publishing task. The invocation protocol cannot be modified after creation.
//
// Description:
//
// Updates the mutable configurations of a specified model connection and resubmits the publishing task. The invocation protocol cannot be modified after creation.
//
// @param request - UpdateModelConnectionRequest
//
// @return UpdateModelConnectionResponse
func (client *Client) UpdateModelConnection(workspaceId *string, connectionId *string, request *UpdateModelConnectionRequest) (_result *UpdateModelConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelConnectionResponse{}
	_body, _err := client.UpdateModelConnectionWithOptions(workspaceId, connectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the business tags of a specified skill. Other attributes that are not specified remain unchanged.
//
// Description:
//
// Updates the business tags of a specified skill. Other attributes that are not specified remain unchanged.
//
// @param tmpReq - UpdateSkillBizTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillBizTagsResponse
func (client *Client) UpdateSkillBizTagsWithOptions(workspaceId *string, skillName *string, tmpReq *UpdateSkillBizTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSkillBizTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSkillBizTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillBizTags"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/biz-tags"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillBizTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the business tags of a specified skill. Other attributes that are not specified remain unchanged.
//
// Description:
//
// Updates the business tags of a specified skill. Other attributes that are not specified remain unchanged.
//
// @param request - UpdateSkillBizTagsRequest
//
// @return UpdateSkillBizTagsResponse
func (client *Client) UpdateSkillBizTags(workspaceId *string, skillName *string, request *UpdateSkillBizTagsRequest) (_result *UpdateSkillBizTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSkillBizTagsResponse{}
	_body, _err := client.UpdateSkillBizTagsWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the version labels and their mappings for a specified skill.
//
// Description:
//
// Updates the version labels and their mappings for a specified skill.
//
// @param tmpReq - UpdateSkillLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillLabelsResponse
func (client *Client) UpdateSkillLabelsWithOptions(workspaceId *string, skillName *string, tmpReq *UpdateSkillLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSkillLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSkillLabelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillLabels"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/labels"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version labels and their mappings for a specified skill.
//
// Description:
//
// Updates the version labels and their mappings for a specified skill.
//
// @param request - UpdateSkillLabelsRequest
//
// @return UpdateSkillLabelsResponse
func (client *Client) UpdateSkillLabels(workspaceId *string, skillName *string, request *UpdateSkillLabelsRequest) (_result *UpdateSkillLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSkillLabelsResponse{}
	_body, _err := client.UpdateSkillLabelsWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the visibility scope of a specified skill.
//
// Description:
//
// Updates the visibility scope of a specified skill.
//
// @param tmpReq - UpdateSkillScopeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillScopeResponse
func (client *Client) UpdateSkillScopeWithOptions(workspaceId *string, skillName *string, tmpReq *UpdateSkillScopeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSkillScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSkillScopeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillScope"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skills/" + dara.PercentEncode(dara.StringValue(skillName)) + "/scope"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the visibility scope of a specified skill.
//
// Description:
//
// Updates the visibility scope of a specified skill.
//
// @param request - UpdateSkillScopeRequest
//
// @return UpdateSkillScopeResponse
func (client *Client) UpdateSkillScope(workspaceId *string, skillName *string, request *UpdateSkillScopeRequest) (_result *UpdateSkillScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSkillScopeResponse{}
	_body, _err := client.UpdateSkillScopeWithOptions(workspaceId, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the description and members of a specified team. When users or agents are provided, the corresponding member lists are replaced based on full overwrite semantics. Member lists that are not provided remain unchanged. The team name cannot be modified after creation.
//
// Description:
//
// Updates the description and members of a specified team.
//
// @param tmpReq - UpdateTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTeamResponse
func (client *Client) UpdateTeamWithOptions(workspaceId *string, teamId *string, tmpReq *UpdateTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTeamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTeam"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/teams/" + dara.PercentEncode(dara.StringValue(teamId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description and members of a specified team. When users or agents are provided, the corresponding member lists are replaced based on full overwrite semantics. Member lists that are not provided remain unchanged. The team name cannot be modified after creation.
//
// Description:
//
// Updates the description and members of a specified team.
//
// @param request - UpdateTeamRequest
//
// @return UpdateTeamResponse
func (client *Client) UpdateTeam(workspaceId *string, teamId *string, request *UpdateTeamRequest) (_result *UpdateTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTeamResponse{}
	_body, _err := client.UpdateTeamWithOptions(workspaceId, teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the display name, email address, or note of a specified user. At least one of displayName, email, and note must be specified. Unspecified properties remain unchanged. The username cannot be modified after creation.
//
// Description:
//
// Updates the display name, email address, or note of a specified user.
//
// @param tmpReq - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(workspaceId *string, agentCoreUserId *string, tmpReq *UpdateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/users/" + dara.PercentEncode(dara.StringValue(agentCoreUserId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the display name, email address, or note of a specified user. At least one of displayName, email, and note must be specified. Unspecified properties remain unchanged. The username cannot be modified after creation.
//
// Description:
//
// Updates the display name, email address, or note of a specified user.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(workspaceId *string, agentCoreUserId *string, request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(workspaceId, agentCoreUserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized state can be updated. Status, TenantId, and RegionId are maintained by the server and cannot be modified through this operation.
//
// Description:
//
// Modifies the name or network configuration of a workspace to adjust workspace information and the private network access environment for agents.
//
// The workspace must be in the `Initialized` state. To enable VPC networking, set `Enabled` in the network configuration to `true`, and provide both `VpcId` and at least one `VSwitchIds`.
//
// @param tmpReq - UpdateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(workspaceId *string, tmpReq *UpdateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspace"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized state can be updated. Status, TenantId, and RegionId are maintained by the server and cannot be modified through this operation.
//
// Description:
//
// Modifies the name or network configuration of a workspace to adjust workspace information and the private network access environment for agents.
//
// The workspace must be in the `Initialized` state. To enable VPC networking, set `Enabled` in the network configuration to `true`, and provide both `VpcId` and at least one `VSwitchIds`.
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(workspaceId *string, request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses the package, and imports the AgentSpec into the current workspace.
//
// Description:
//
// Retrieves an uploaded agent template ZIP package from OSS, parses the package, and imports the agent template into the current workspace.
//
// @param tmpReq - UploadAgentSpecViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAgentSpecViaOssResponse
func (client *Client) UploadAgentSpecViaOssWithOptions(workspaceId *string, tmpReq *UploadAgentSpecViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadAgentSpecViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadAgentSpecViaOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadAgentSpecViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-spec-actions/upload-via-oss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadAgentSpecViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses the package, and imports the AgentSpec into the current workspace.
//
// Description:
//
// Retrieves an uploaded agent template ZIP package from OSS, parses the package, and imports the agent template into the current workspace.
//
// @param request - UploadAgentSpecViaOssRequest
//
// @return UploadAgentSpecViaOssResponse
func (client *Client) UploadAgentSpecViaOss(workspaceId *string, request *UploadAgentSpecViaOssRequest) (_result *UploadAgentSpecViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadAgentSpecViaOssResponse{}
	_body, _err := client.UploadAgentSpecViaOssWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a skill into the current workspace by retrieving and parsing an uploaded skill ZIP package from OSS.
//
// Description:
//
// Imports a skill into the current workspace by retrieving and parsing an uploaded skill ZIP package from OSS.
//
// @param tmpReq - UploadSkillViaOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSkillViaOssResponse
func (client *Client) UploadSkillViaOssWithOptions(workspaceId *string, tmpReq *UploadSkillViaOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadSkillViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadSkillViaOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadSkillViaOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/skill-actions/upload-via-oss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadSkillViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a skill into the current workspace by retrieving and parsing an uploaded skill ZIP package from OSS.
//
// Description:
//
// Imports a skill into the current workspace by retrieving and parsing an uploaded skill ZIP package from OSS.
//
// @param request - UploadSkillViaOssRequest
//
// @return UploadSkillViaOssResponse
func (client *Client) UploadSkillViaOss(workspaceId *string, request *UploadSkillViaOssRequest) (_result *UploadSkillViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadSkillViaOssResponse{}
	_body, _err := client.UploadSkillViaOssWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates the read conditions of the workspace role for a specified OSS code package object and returns the source URI and package information.
//
// Description:
//
// Checks the read conditions and package information of a specified OSS object before submitting the code package for deployment, and obtains the source URI for subsequent deployment.
//
// ### Request and validation scope
//
// In the JSON request body, specify `ossBucketName` and the complete `ossObjectName`. Do not specify the HTTPS download URL. The object extension must be ZIP or TAR.GZ, and the compressed size must be between 1 and 500,000,000 bytes.
//
// This operation uses the workspace role to check the object metadata and read permissions, and returns the `sourceUri`, file name, size, and format.
//
// ### Subsequent deployment
//
// Passing the validation indicates that the object check is successful, but does not mean that the archive content is complete or ready to start. Submit the returned `sourceUri` to [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent). The system will recheck and save a fixed copy upon submission.
//
// This operation is an optional precheck and is not a required step for creation or update.
//
// @param tmpReq - ValidateAgentCodePackageFromOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateAgentCodePackageFromOssResponse
func (client *Client) ValidateAgentCodePackageFromOssWithOptions(workspaceId *string, tmpReq *ValidateAgentCodePackageFromOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateAgentCodePackageFromOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ValidateAgentCodePackageFromOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateAgentCodePackageFromOss"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-code-packages/validate-from-oss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateAgentCodePackageFromOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the read conditions of the workspace role for a specified OSS code package object and returns the source URI and package information.
//
// Description:
//
// Checks the read conditions and package information of a specified OSS object before submitting the code package for deployment, and obtains the source URI for subsequent deployment.
//
// ### Request and validation scope
//
// In the JSON request body, specify `ossBucketName` and the complete `ossObjectName`. Do not specify the HTTPS download URL. The object extension must be ZIP or TAR.GZ, and the compressed size must be between 1 and 500,000,000 bytes.
//
// This operation uses the workspace role to check the object metadata and read permissions, and returns the `sourceUri`, file name, size, and format.
//
// ### Subsequent deployment
//
// Passing the validation indicates that the object check is successful, but does not mean that the archive content is complete or ready to start. Submit the returned `sourceUri` to [CreateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/CreateCustomAgent) or [UpdateCustomAgent](https://api.aliyun.com/document/AgentCore/2026-08-04/UpdateCustomAgent). The system will recheck and save a fixed copy upon submission.
//
// This operation is an optional precheck and is not a required step for creation or update.
//
// @param request - ValidateAgentCodePackageFromOssRequest
//
// @return ValidateAgentCodePackageFromOssResponse
func (client *Client) ValidateAgentCodePackageFromOss(workspaceId *string, request *ValidateAgentCodePackageFromOssRequest) (_result *ValidateAgentCodePackageFromOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateAgentCodePackageFromOssResponse{}
	_body, _err := client.ValidateAgentCodePackageFromOssWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding status of the authorization policy for the code package source bucket to the current workspace resource role.
//
// Description:
//
// Queries whether the read authorization for the code package source bucket is bound to the workspace resource role to determine whether supplementary authorization is required.
//
// ### Query results
//
// `ossBucketName` is passed as a query parameter. `authorizationStatus=AUTHORIZED` indicates that the policy is bound, and `UNAUTHORIZED` indicates that it is not bound.
//
// A bound policy does not mean that a specific code package object definitely exists or is readable. To check the object, call [ValidateAgentCodePackageFromOss](https://api.aliyun.com/document/AgentCore/2026-08-04/ValidateAgentCodePackageFromOss). To obtain the authorization link, call [GetAgentCodePackageOssRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageOssRamAuthorizeUrl).
//
// @param request - VerifyAgentCodePackageOssRamAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAgentCodePackageOssRamAuthorizationResponse
func (client *Client) VerifyAgentCodePackageOssRamAuthorizationWithOptions(workspaceId *string, request *VerifyAgentCodePackageOssRamAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyAgentCodePackageOssRamAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssBucketName) {
		query["ossBucketName"] = request.OssBucketName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAgentCodePackageOssRamAuthorization"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agent-code-packages/oss/authorize/verify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAgentCodePackageOssRamAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding status of the authorization policy for the code package source bucket to the current workspace resource role.
//
// Description:
//
// Queries whether the read authorization for the code package source bucket is bound to the workspace resource role to determine whether supplementary authorization is required.
//
// ### Query results
//
// `ossBucketName` is passed as a query parameter. `authorizationStatus=AUTHORIZED` indicates that the policy is bound, and `UNAUTHORIZED` indicates that it is not bound.
//
// A bound policy does not mean that a specific code package object definitely exists or is readable. To check the object, call [ValidateAgentCodePackageFromOss](https://api.aliyun.com/document/AgentCore/2026-08-04/ValidateAgentCodePackageFromOss). To obtain the authorization link, call [GetAgentCodePackageOssRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetAgentCodePackageOssRamAuthorizeUrl).
//
// @param request - VerifyAgentCodePackageOssRamAuthorizationRequest
//
// @return VerifyAgentCodePackageOssRamAuthorizationResponse
func (client *Client) VerifyAgentCodePackageOssRamAuthorization(workspaceId *string, request *VerifyAgentCodePackageOssRamAuthorizationRequest) (_result *VerifyAgentCodePackageOssRamAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyAgentCodePackageOssRamAuthorizationResponse{}
	_body, _err := client.VerifyAgentCodePackageOssRamAuthorizationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether one or more Qoder service account keys are available at a specified site. This operation does not save the keys or modify the connector status.
//
// Description:
//
// Verifies Qoder service account keys for a specified site to check whether the credentials are available before enabling or updating a connector. You can submit multiple keys at the same time. The response returns a list of invalid service account keys.
//
// @param tmpReq - VerifyConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyConnectorResponse
func (client *Client) VerifyConnectorWithOptions(workspaceId *string, connectorName *string, tmpReq *VerifyConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &VerifyConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyConnector"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName)) + "/actions/verify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether one or more Qoder service account keys are available at a specified site. This operation does not save the keys or modify the connector status.
//
// Description:
//
// Verifies Qoder service account keys for a specified site to check whether the credentials are available before enabling or updating a connector. You can submit multiple keys at the same time. The response returns a list of invalid service account keys.
//
// @param request - VerifyConnectorRequest
//
// @return VerifyConnectorResponse
func (client *Client) VerifyConnector(workspaceId *string, connectorName *string, request *VerifyConnectorRequest) (_result *VerifyConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyConnectorResponse{}
	_body, _err := client.VerifyConnectorWithOptions(workspaceId, connectorName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorization status of a workspace resource role for a specified ACR Enterprise Edition repository.
//
// Description:
//
// Queries the authorization status of a workspace resource role for a specified ACR Enterprise Edition image repository to determine whether to proceed with image preparation or grant additional authorizations.
//
// ### Request parameters
//
// `acrInstanceId`, `namespace`, and `repository` use query parameters and must be consistent with the targets used when obtaining the authorization URL.
//
// ### Authorization and deployment
//
// `authorizationStatus=AUTHORIZED` indicates that the target authorization for the selected role is satisfied, and `UNAUTHORIZED` indicates that it is not yet satisfied. If authorization is required, call [GetWorkspaceAcrRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAcrRamAuthorizeUrl) and confirm on the RAM page.
//
// The authorization status is checked separately from whether the specific image tag exists and whether the pull network is reachable.
//
// @param request - VerifyWorkspaceAcrRamAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyWorkspaceAcrRamAuthorizationResponse
func (client *Client) VerifyWorkspaceAcrRamAuthorizationWithOptions(workspaceId *string, request *VerifyWorkspaceAcrRamAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyWorkspaceAcrRamAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		query["acrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Repository) {
		query["repository"] = request.Repository
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyWorkspaceAcrRamAuthorization"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/acr/authorize/verify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyWorkspaceAcrRamAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization status of a workspace resource role for a specified ACR Enterprise Edition repository.
//
// Description:
//
// Queries the authorization status of a workspace resource role for a specified ACR Enterprise Edition image repository to determine whether to proceed with image preparation or grant additional authorizations.
//
// ### Request parameters
//
// `acrInstanceId`, `namespace`, and `repository` use query parameters and must be consistent with the targets used when obtaining the authorization URL.
//
// ### Authorization and deployment
//
// `authorizationStatus=AUTHORIZED` indicates that the target authorization for the selected role is satisfied, and `UNAUTHORIZED` indicates that it is not yet satisfied. If authorization is required, call [GetWorkspaceAcrRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAcrRamAuthorizeUrl) and confirm on the RAM page.
//
// The authorization status is checked separately from whether the specific image tag exists and whether the pull network is reachable.
//
// @param request - VerifyWorkspaceAcrRamAuthorizationRequest
//
// @return VerifyWorkspaceAcrRamAuthorizationResponse
func (client *Client) VerifyWorkspaceAcrRamAuthorization(workspaceId *string, request *VerifyWorkspaceAcrRamAuthorizationRequest) (_result *VerifyWorkspaceAcrRamAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyWorkspaceAcrRamAuthorizationResponse{}
	_body, _err := client.VerifyWorkspaceAcrRamAuthorizationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the workspace resource role has been granted access to the specified NAS (AgenticFS) access point for subsequent Agent mount configurations.
//
// Description:
//
// Checks whether the workspace has been granted access to the specified AgenticFS access point for subsequent Agent file mount configurations.
//
// [AgenticFS](https://www.alibabacloud.com/help/en/nas/product-overview/agenticfs-overview) is a storage type of File Storage NAS designed for AI agents. Access points provide mount entries for file spaces.
//
// ### Query and what to do next
//
// For the `server` query parameter, specify the access point domain name without the protocol, port, or path. Keep it consistent with the value used when obtaining the authorization link.
//
// | `authorizationStatus` | What to do next |
//
// | --- | --- |
//
// | `AUTHORIZED` | Authorization is complete. You can continue to configure `agenticFsMounts` for the agent. |
//
// | `UNAUTHORIZED` | Obtain the link through [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl), complete the authorization, and then query again. |
//
// After authorization is complete, you must still commit the mount configuration in the agent and check the configuration processing result through the agent details. The mount folder must already exist, the access point must be located in the VPC associated with the workspace, and VPC access must be enabled for the agent.
//
// @param request - VerifyWorkspaceAgenticFsMountRamAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyWorkspaceAgenticFsMountRamAuthorizationResponse
func (client *Client) VerifyWorkspaceAgenticFsMountRamAuthorizationWithOptions(workspaceId *string, request *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Server) {
		query["server"] = request.Server
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyWorkspaceAgenticFsMountRamAuthorization"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/agentic-fs/authorize/verify"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyWorkspaceAgenticFsMountRamAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the workspace resource role has been granted access to the specified NAS (AgenticFS) access point for subsequent Agent mount configurations.
//
// Description:
//
// Checks whether the workspace has been granted access to the specified AgenticFS access point for subsequent Agent file mount configurations.
//
// [AgenticFS](https://www.alibabacloud.com/help/en/nas/product-overview/agenticfs-overview) is a storage type of File Storage NAS designed for AI agents. Access points provide mount entries for file spaces.
//
// ### Query and what to do next
//
// For the `server` query parameter, specify the access point domain name without the protocol, port, or path. Keep it consistent with the value used when obtaining the authorization link.
//
// | `authorizationStatus` | What to do next |
//
// | --- | --- |
//
// | `AUTHORIZED` | Authorization is complete. You can continue to configure `agenticFsMounts` for the agent. |
//
// | `UNAUTHORIZED` | Obtain the link through [GetWorkspaceAgenticFsMountRamAuthorizeUrl](https://api.aliyun.com/document/AgentCore/2026-08-04/GetWorkspaceAgenticFsMountRamAuthorizeUrl), complete the authorization, and then query again. |
//
// After authorization is complete, you must still commit the mount configuration in the agent and check the configuration processing result through the agent details. The mount folder must already exist, the access point must be located in the VPC associated with the workspace, and VPC access must be enabled for the agent.
//
// @param request - VerifyWorkspaceAgenticFsMountRamAuthorizationRequest
//
// @return VerifyWorkspaceAgenticFsMountRamAuthorizationResponse
func (client *Client) VerifyWorkspaceAgenticFsMountRamAuthorization(workspaceId *string, request *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) (_result *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyWorkspaceAgenticFsMountRamAuthorizationResponse{}
	_body, _err := client.VerifyWorkspaceAgenticFsMountRamAuthorizationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the OSS mount authorization for a workspace, or checks and updates the authorization status of the default OSS storage for the workspace.
//
// Description:
//
// Verifies the OSS mount authorization for a workspace, or checks and updates the authorization status of the default OSS storage for the workspace.
//
// ### Two query scopes
//
// | Query parameter `bucketName` | Behavior |
//
// | --- | --- |
//
// | Bucket name provided | Checks the mount authorization of the bucket for the workspace resource role without modifying the default OSS binding. |
//
// | Omitted | Checks the authorization of the default private OSS storage for the workspace and updates its authorization status. |
//
// ### Return statuses
//
// `AUTHORIZED` indicates that authorization is granted, `UNAUTHORIZED` indicates that authorization is not granted, and `NOT_REQUIRED` indicates that authorization is not required. The authorization status does not indicate whether a specific mount directory exists or an instance has completed the mount.
//
// If the goal is to read from a code package source, use [VerifyAgentCodePackageOssRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyAgentCodePackageOssRamAuthorization). Do not mix code package verification with runtime directory mounting.
//
// @param request - VerifyWorkspaceOssMountRamAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyWorkspaceOssMountRamAuthorizationResponse
func (client *Client) VerifyWorkspaceOssMountRamAuthorizationWithOptions(workspaceId *string, request *VerifyWorkspaceOssMountRamAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyWorkspaceOssMountRamAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["bucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyWorkspaceOssMountRamAuthorization"),
		Version:     dara.String("2026-08-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/oss-mount/authorize/verify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyWorkspaceOssMountRamAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the OSS mount authorization for a workspace, or checks and updates the authorization status of the default OSS storage for the workspace.
//
// Description:
//
// Verifies the OSS mount authorization for a workspace, or checks and updates the authorization status of the default OSS storage for the workspace.
//
// ### Two query scopes
//
// | Query parameter `bucketName` | Behavior |
//
// | --- | --- |
//
// | Bucket name provided | Checks the mount authorization of the bucket for the workspace resource role without modifying the default OSS binding. |
//
// | Omitted | Checks the authorization of the default private OSS storage for the workspace and updates its authorization status. |
//
// ### Return statuses
//
// `AUTHORIZED` indicates that authorization is granted, `UNAUTHORIZED` indicates that authorization is not granted, and `NOT_REQUIRED` indicates that authorization is not required. The authorization status does not indicate whether a specific mount directory exists or an instance has completed the mount.
//
// If the goal is to read from a code package source, use [VerifyAgentCodePackageOssRamAuthorization](https://api.aliyun.com/document/AgentCore/2026-08-04/VerifyAgentCodePackageOssRamAuthorization). Do not mix code package verification with runtime directory mounting.
//
// @param request - VerifyWorkspaceOssMountRamAuthorizationRequest
//
// @return VerifyWorkspaceOssMountRamAuthorizationResponse
func (client *Client) VerifyWorkspaceOssMountRamAuthorization(workspaceId *string, request *VerifyWorkspaceOssMountRamAuthorizationRequest) (_result *VerifyWorkspaceOssMountRamAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyWorkspaceOssMountRamAuthorizationResponse{}
	_body, _err := client.VerifyWorkspaceOssMountRamAuthorizationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
