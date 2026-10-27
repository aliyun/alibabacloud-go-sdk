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
// Writes messages or text to a memory store. By default, sync is set to true, which means the operation synchronously waits for extraction and returns only memories (an empty array is allowed). If sync is set to false, the operation returns only a taskId, which you can query by calling GetMemoryTask.
//
// Description:
//
// Writes messages or text to a memory store. By default, sync is set to true, which means the operation synchronously waits for extraction and returns only memories (an empty array is allowed). If sync is set to false, the operation returns only a taskId, which you can query by calling GetMemoryTask.
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
// Writes messages or text to a memory store. By default, sync is set to true, which means the operation synchronously waits for extraction and returns only memories (an empty array is allowed). If sync is set to false, the operation returns only a taskId, which you can query by calling GetMemoryTask.
//
// Description:
//
// Writes messages or text to a memory store. By default, sync is set to true, which means the operation synchronously waits for extraction and returns only memories (an empty array is allowed). If sync is set to false, the operation returns only a taskId, which you can query by calling GetMemoryTask.
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
// Uploads skill ZIP packages in batches through OSS and returns the processing result for each skill.
//
// Description:
//
// ## Request description
//
// Uploads skill ZIP packages in batches through OSS and returns the processing result for each skill.
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
// Uploads skill ZIP packages in batches through OSS and returns the processing result for each skill.
//
// Description:
//
// ## Request description
//
// Uploads skill ZIP packages in batches through OSS and returns the processing result for each skill.
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
// Converts an MCP to free-edit mode.
//
// Description:
//
// Disables template usage constraints. After the conversion, the MCP retains its source and tags, but no longer appears on the usage page.
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
// Converts an MCP to free-edit mode.
//
// Description:
//
// Disables template usage constraints. After the conversion, the MCP retains its source and tags, but no longer appears on the usage page.
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
// Creates an IM channel for a specified agent and binds it to a publicly accessible ServiceEndpoint.
//
// Description:
//
// Creates an IM channel for a specified agent and binds it to a publicly accessible ServiceEndpoint.
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
// Creates an IM channel for a specified agent and binds it to a publicly accessible ServiceEndpoint.
//
// Description:
//
// Creates an IM channel for a specified agent and binds it to a publicly accessible ServiceEndpoint.
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
// ## Request description
//
// Creates an AgentSpec in a specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
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
// ## Request description
//
// Creates an AgentSpec in a specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
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
// ## Request description
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist and must not have a draft currently being edited.
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
// ## Request description
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist and must not have a draft currently being edited.
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
// Creates a credential in a specified workspace for authentication when an agent accesses external services. Currently, only the apiKey type is supported. The credential content is passed as a JSON string through credentialMetadata and can only be queried in a desensitized format after it is written.
//
// Description:
//
// Creates a credential in a workspace for authentication of services such as Connector.
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
// Creates a credential in a specified workspace for authentication when an agent accesses external services. Currently, only the apiKey type is supported. The credential content is passed as a JSON string through credentialMetadata and can only be queried in a desensitized format after it is written.
//
// Description:
//
// Creates a credential in a workspace for authentication of services such as Connector.
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
// Creates a custom agent in a specified workspace. Container image and code package deployments are supported. The creation process is asynchronous. You can call GetCustomAgent to query the deployment result.
//
// Description:
//
// Before you call this operation, prepare a writable workspace, a RAM role ARN, a container image or code package, and runtime configurations. You must specify name, artifact, roleArn, and runtime. The runtime parameter must include both compute and sessionPolicy.
//
// In container mode, use artifact.container and specify artifact.port. The image must have the required access permissions and be accessible over the network allowed by the workspace.
//
// In code package pattern, specify artifact.port, and the sourceUri, languageRuntime, and entrypoint of artifact.code. Do not specify container at the same time. Dependencies must be packaged by the user. ZIP and TAR.GZ formats are supported. The compressed package size cannot exceed 500,000,000 bytes. You can call ListAgentCodeRuntimes to query the runtime environment identity. The sourceUri parameter supports oss:// addresses or HTTPS object addresses of standard OSS public domain names, including signed URLs. For local files, call GetAgentCodePackageUploadUrl to obtain an upload URL, upload the file directly, and then commit the returned sourceUri.
//
// When you submit a new code package source, the backend rechecks the read permissions, object, and size. The configuration is saved only after the transfer succeeds. If the transfer fails, the configuration is not saved. Saved fixed packages can be reused. Set the client timeout to 210 seconds for code package requests. Archive integrity and secure decompression are checked at startup. A successful acceptance does not mean that the agent is ready to run.
//
// The creation process is asynchronous. A successful response indicates that the request is accepted. You can use the returned agentId to call GetCustomAgent and check the latestVersionStatus, latestExecutionError, and status corresponding to latestSpecVersion to confirm the deployment result. After the deployment is ready, you can call CreateServiceEndpoint to configure the access endpoint.
//
// The clientToken parameter does not provide idempotency guarantees. Do not rely on this parameter to prevent duplicate requests. The healthyThreshold parameter for health checks does not take effect. Omit this parameter.
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
// Creates a custom agent in a specified workspace. Container image and code package deployments are supported. The creation process is asynchronous. You can call GetCustomAgent to query the deployment result.
//
// Description:
//
// Before you call this operation, prepare a writable workspace, a RAM role ARN, a container image or code package, and runtime configurations. You must specify name, artifact, roleArn, and runtime. The runtime parameter must include both compute and sessionPolicy.
//
// In container mode, use artifact.container and specify artifact.port. The image must have the required access permissions and be accessible over the network allowed by the workspace.
//
// In code package pattern, specify artifact.port, and the sourceUri, languageRuntime, and entrypoint of artifact.code. Do not specify container at the same time. Dependencies must be packaged by the user. ZIP and TAR.GZ formats are supported. The compressed package size cannot exceed 500,000,000 bytes. You can call ListAgentCodeRuntimes to query the runtime environment identity. The sourceUri parameter supports oss:// addresses or HTTPS object addresses of standard OSS public domain names, including signed URLs. For local files, call GetAgentCodePackageUploadUrl to obtain an upload URL, upload the file directly, and then commit the returned sourceUri.
//
// When you submit a new code package source, the backend rechecks the read permissions, object, and size. The configuration is saved only after the transfer succeeds. If the transfer fails, the configuration is not saved. Saved fixed packages can be reused. Set the client timeout to 210 seconds for code package requests. Archive integrity and secure decompression are checked at startup. A successful acceptance does not mean that the agent is ready to run.
//
// The creation process is asynchronous. A successful response indicates that the request is accepted. You can use the returned agentId to call GetCustomAgent and check the latestVersionStatus, latestExecutionError, and status corresponding to latestSpecVersion to confirm the deployment result. After the deployment is ready, you can call CreateServiceEndpoint to configure the access endpoint.
//
// The clientToken parameter does not provide idempotency guarantees. Do not rely on this parameter to prevent duplicate requests. The healthyThreshold parameter for health checks does not take effect. Omit this parameter.
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
// Creates an external agent in a specified workspace.
//
// Description:
//
// Creates an external agent in a specified workspace.
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
// Creates an external agent in a specified workspace.
//
// Description:
//
// Creates an external agent in a specified workspace.
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
// Creates the Bootstrap token and CMS configuration required to onboard a specified external agent.
//
// Description:
//
// Creates the Bootstrap token and CMS configuration required to onboard a specified external agent.
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
// Creates the Bootstrap token and CMS configuration required to onboard a specified external agent.
//
// Description:
//
// Creates the Bootstrap token and CMS configuration required to onboard a specified external agent.
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
// Binds an external identity provider to a specified workspace for single sign-on and organization member synchronization of workspace users. A maximum of one external identity provider can be bound to each workspace. The binding is an asynchronous operation. After the API returns, you can query the status by calling GetIdentityProvider to track the progress.
//
// Description:
//
// ## Features and prerequisites
//
// A workspace can be bound to only one external identity provider. The identityProviderType parameter supports DingTalk, Feishu, and CustomSSO. CustomSSO only initializes the shared user pool. Do not pass non-empty metadata, and do not pass true for loginEnabled or syncEnabled. By default, both switches are turned off. Protocols and SAML configurations are saved by calling UpdateIdentityProvider. SCIM is configured through a separate API. Saving SAML first is not required. After the request is accepted and CONFIGURED is returned, call GetIdentityProvider and wait for the user pool initialization to complete. READY only indicates that the resource is ready, not that SAML is configured. If the same CustomSSO binding is in the CONFIGURED or READY state, the existing binding may be returned. Not all states can be treated as repeatedly creatable.
//
// For DingTalk, appKey, appSecret, and corpId are required in metadata. For Feishu, appId and appSecret are required. The vendor login switch is enabled by default. The synchronization switch is disabled by default. When synchronization is enabled, encryptKey and verificationToken are also required. The vendor syncOnBind and syncScope parameters retain the existing department-based synchronization definitions. Do not use these vendor pull parameters to replace SCIM configurations during CustomSSO initialization.
//
// ## Requests and examples
//
// POST /workspaces/{workspaceId}/identity-providers
//
// Call conventions: Call the API in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/identity-providers
//
// {"identityProviderType": "CustomSSO"}
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": false,
//
//	  "status": "CONFIGURED"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call failure to an upstream service may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 409 IdentityProviderAlreadyBound: An external binding already exists. Query and reuse the correct binding first. Do not automatically unbind the entire binding. Asynchronous initialization failure does not equal HTTP request failure. You must continue to query the status.
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
// Binds an external identity provider to a specified workspace for single sign-on and organization member synchronization of workspace users. A maximum of one external identity provider can be bound to each workspace. The binding is an asynchronous operation. After the API returns, you can query the status by calling GetIdentityProvider to track the progress.
//
// Description:
//
// ## Features and prerequisites
//
// A workspace can be bound to only one external identity provider. The identityProviderType parameter supports DingTalk, Feishu, and CustomSSO. CustomSSO only initializes the shared user pool. Do not pass non-empty metadata, and do not pass true for loginEnabled or syncEnabled. By default, both switches are turned off. Protocols and SAML configurations are saved by calling UpdateIdentityProvider. SCIM is configured through a separate API. Saving SAML first is not required. After the request is accepted and CONFIGURED is returned, call GetIdentityProvider and wait for the user pool initialization to complete. READY only indicates that the resource is ready, not that SAML is configured. If the same CustomSSO binding is in the CONFIGURED or READY state, the existing binding may be returned. Not all states can be treated as repeatedly creatable.
//
// For DingTalk, appKey, appSecret, and corpId are required in metadata. For Feishu, appId and appSecret are required. The vendor login switch is enabled by default. The synchronization switch is disabled by default. When synchronization is enabled, encryptKey and verificationToken are also required. The vendor syncOnBind and syncScope parameters retain the existing department-based synchronization definitions. Do not use these vendor pull parameters to replace SCIM configurations during CustomSSO initialization.
//
// ## Requests and examples
//
// POST /workspaces/{workspaceId}/identity-providers
//
// Call conventions: Call the API in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/identity-providers
//
// {"identityProviderType": "CustomSSO"}
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": false,
//
//	  "status": "CONFIGURED"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call failure to an upstream service may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 409 IdentityProviderAlreadyBound: An external binding already exists. Query and reuse the correct binding first. Do not automatically unbind the entire binding. Asynchronous initialization failure does not equal HTTP request failure. You must continue to query the status.
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
// Issues a SCIM credential explicitly and returns the plaintext token only for the current call. A maximum of two credentials are allowed in the same pool. This operation does not automatically revoke old credentials or enable AC user synchronization.
//
// Description:
//
// ## Features and prerequisites
//
// Issues a new credential explicitly. This requires an available shared pool and the AgentIdentity SCIM reception to be Enabled. This operation does not automatically enable reception. If it is not enabled, call prepare first. Returns the configuration, credential summary, and the current token. The limit of a maximum of two credentials is enforced by AgentIdentity. An existing available token does not prevent users from explicitly creating a second one. Old credentials are not automatically revoked, and no extended rotation process is provided.
//
// The token is returned only once and cannot be retrieved again in subsequent list and query operations. Do not automatically retry the issuance. If the response is lost, check the list first to verify whether it has been created. Creating a credential does not enable AC syncEnabled, start a full user import, or switch SAML logon. Deleting an old token should be performed explicitly after the user confirms the actual usage of the external provisioning system.
//
// ## Request and examples
//
// POST /workspaces/{workspaceId}/scim-provisioning/credentials
//
// Call conventions: Call the operation in the region of the target workspace, using an AK/STS identity that has permissions for the identity source resources of the corresponding workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameter content of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log the key. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/scim-provisioning/credentials
//
// No business request body.
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "credential": {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "expireTime": "<Upstream expiration time>"
//
//	  },
//
//	  "token": "<Newly issued token for the current request>"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that there is no binding of the matching type. Synchronous calls to the upstream may fail and return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 LimitExceeded.ScimCredential: A maximum of two credentials are allowed. Delete one before you can continue to create. 400 InvalidParameter: Reception is not enabled or the shared pool is unavailable. Do not retry quota errors as system failures.
//
// ## Call relationships and usage boundaries
//
// After creation, securely save the token that is returned only once and configure it in the external provisioning system. Existing tokens remain valid. Only DeleteIdentityProviderSCIMCredential revokes the specified credential. Do not automatically retry when the response is lost. Call ListIdentityProviderSCIMCredentials to check first. This operation does not provide automatic rotation or scaling.
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
// Issues a SCIM credential explicitly and returns the plaintext token only for the current call. A maximum of two credentials are allowed in the same pool. This operation does not automatically revoke old credentials or enable AC user synchronization.
//
// Description:
//
// ## Features and prerequisites
//
// Issues a new credential explicitly. This requires an available shared pool and the AgentIdentity SCIM reception to be Enabled. This operation does not automatically enable reception. If it is not enabled, call prepare first. Returns the configuration, credential summary, and the current token. The limit of a maximum of two credentials is enforced by AgentIdentity. An existing available token does not prevent users from explicitly creating a second one. Old credentials are not automatically revoked, and no extended rotation process is provided.
//
// The token is returned only once and cannot be retrieved again in subsequent list and query operations. Do not automatically retry the issuance. If the response is lost, check the list first to verify whether it has been created. Creating a credential does not enable AC syncEnabled, start a full user import, or switch SAML logon. Deleting an old token should be performed explicitly after the user confirms the actual usage of the external provisioning system.
//
// ## Request and examples
//
// POST /workspaces/{workspaceId}/scim-provisioning/credentials
//
// Call conventions: Call the operation in the region of the target workspace, using an AK/STS identity that has permissions for the identity source resources of the corresponding workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameter content of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log the key. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/scim-provisioning/credentials
//
// No business request body.
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "credential": {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "expireTime": "<Upstream expiration time>"
//
//	  },
//
//	  "token": "<Newly issued token for the current request>"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that there is no binding of the matching type. Synchronous calls to the upstream may fail and return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 LimitExceeded.ScimCredential: A maximum of two credentials are allowed. Delete one before you can continue to create. 400 InvalidParameter: Reception is not enabled or the shared pool is unavailable. Do not retry quota errors as system failures.
//
// ## Call relationships and usage boundaries
//
// After creation, securely save the token that is returned only once and configure it in the external provisioning system. Existing tokens remain valid. Only DeleteIdentityProviderSCIMCredential revokes the specified credential. Do not automatically retry when the response is lost. Call ListIdentityProviderSCIMCredentials to check first. This operation does not provide automatic rotation or scaling.
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
// Initiates a manual synchronization of an identity provider.
//
// Description:
//
// Initiates a one-time user synchronization based on the saved scope without saving the configuration or requiring syncEnabled=true. Prerequisites: the identity provider is in the READY or SYNC_FAILED state, a user pool exists, and no other pending identity provider operations are in progress. The returned taskId only indicates that the request is accepted. Call GetIdentityProviderSyncJob to poll the status until it reaches SUCCESS, FAILED, or CANCELED. If the task fails, check the failureReason. The workspace user quota still applies, and previously synchronized users are not automatically deleted if the scope is reduced.
//
// Authorization information: RAM action agentcore:CreateIdentityProviderSyncJob, access level Write, resource ARN: acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is required for authorization.
//
// Idempotency and retry: This operation does not provide a ClientToken parameter for conditional idempotency. Repeated submissions do not guarantee the same task is returned. If an operation is already in progress, an operation conflict is returned. Submitting again after the task ends creates a new synchronization task. If a submission times out, query the identity provider status first to avoid blindly resubmitting the request.
//
// Calling this API is free of charge, and operation records support ActionTrail auditing.
//
// Initial selective synchronization process: CreateIdentityProvider (syncOnBind=false, it is recommended to keep syncEnabled=false initially) → GetIdentityProvider and wait for READY → ListIdentityProviderDepartments to select departments → UpdateIdentityProvider to save syncScope and wait for READY → CreateIdentityProviderSyncJob → GetIdentityProviderSyncJob to query the terminal state. If incremental synchronization is required, configure callbacks and event subscriptions, and then update syncEnabled=true.
//
// Response example:
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
//	"data": {
//
//	  "taskId": "12345678-1234-4234-8234-123456789abc",
//
//	  "status": "PENDING"
//
//	}
//
// }
//
// ```
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
// Initiates a manual synchronization of an identity provider.
//
// Description:
//
// Initiates a one-time user synchronization based on the saved scope without saving the configuration or requiring syncEnabled=true. Prerequisites: the identity provider is in the READY or SYNC_FAILED state, a user pool exists, and no other pending identity provider operations are in progress. The returned taskId only indicates that the request is accepted. Call GetIdentityProviderSyncJob to poll the status until it reaches SUCCESS, FAILED, or CANCELED. If the task fails, check the failureReason. The workspace user quota still applies, and previously synchronized users are not automatically deleted if the scope is reduced.
//
// Authorization information: RAM action agentcore:CreateIdentityProviderSyncJob, access level Write, resource ARN: acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is required for authorization.
//
// Idempotency and retry: This operation does not provide a ClientToken parameter for conditional idempotency. Repeated submissions do not guarantee the same task is returned. If an operation is already in progress, an operation conflict is returned. Submitting again after the task ends creates a new synchronization task. If a submission times out, query the identity provider status first to avoid blindly resubmitting the request.
//
// Calling this API is free of charge, and operation records support ActionTrail auditing.
//
// Initial selective synchronization process: CreateIdentityProvider (syncOnBind=false, it is recommended to keep syncEnabled=false initially) → GetIdentityProvider and wait for READY → ListIdentityProviderDepartments to select departments → UpdateIdentityProvider to save syncScope and wait for READY → CreateIdentityProviderSyncJob → GetIdentityProviderSyncJob to query the terminal state. If incremental synchronization is required, configure callbacks and event subscriptions, and then update syncEnabled=true.
//
// Response example:
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
//	"data": {
//
//	  "taskId": "12345678-1234-4234-8234-123456789abc",
//
//	  "status": "PENDING"
//
//	}
//
// }
//
// ```
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
// Creates an image cache.
//
// Description:
//
// Creates or reuses an image cache in the current workspace based on the image source and tag. Both public and private images uniformly reuse the workspace\\"s source image pull authorization and resource roles. The resource manager parses the tag during the transfer node and performs image replication based on the parsed summary. The digest is empty before parsing is complete. For regular tags, the saved summary is reused after parsing. For the latest tag, it is parsed again when registered after reaching the desired state, and duplicate registrations in progress reuse the node. The logical imageId remains stable. accessInternet and accessVpc only determine the source reading and transfer for the current request and do not modify the agent network. An HTTP 200 response returns an acceptance snapshot of PENDING, READY, or FAILED, which does not indicate that the creation is complete. ClientToken does not guarantee to deduplicate requests or recover the original result. If the result is unconfirmed, query the query cache status first before deciding whether to recommit.
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
// Creates an image cache.
//
// Description:
//
// Creates or reuses an image cache in the current workspace based on the image source and tag. Both public and private images uniformly reuse the workspace\\"s source image pull authorization and resource roles. The resource manager parses the tag during the transfer node and performs image replication based on the parsed summary. The digest is empty before parsing is complete. For regular tags, the saved summary is reused after parsing. For the latest tag, it is parsed again when registered after reaching the desired state, and duplicate registrations in progress reuse the node. The logical imageId remains stable. accessInternet and accessVpc only determine the source reading and transfer for the current request and do not modify the agent network. An HTTP 200 response returns an acceptance snapshot of PENDING, READY, or FAILED, which does not indicate that the creation is complete. ClientToken does not guarantee to deduplicate requests or recover the original result. If the result is unconfirmed, query the query cache status first before deciding whether to recommit.
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
// Creates a managed agent in a specified workspace.
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
// Creates a managed agent in a specified workspace.
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
// Creates an MCP service in a specified workspace. The creation is an asynchronous process. You can check whether the service is ready by using the returned status or by calling a query operation.
//
// Description:
//
// ## Operation description
//
// Creates an MCP service in a specified workspace. The creation is an asynchronous process. You can check whether the service is ready by using the returned status or by calling a query operation.
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
// Creates an MCP service in a specified workspace. The creation is an asynchronous process. You can check whether the service is ready by using the returned status or by calling a query operation.
//
// Description:
//
// ## Operation description
//
// Creates an MCP service in a specified workspace. The creation is an asynchronous process. You can check whether the service is ready by using the returned status or by calling a query operation.
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
// Creates a long-term memory store in a specified workspace.
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
// Creates a long-term memory store in a specified workspace.
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
// Creates a model configuration for model creation under a model connection in a specified workspace.
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
// Creates a model configuration for model creation under a model connection in a specified workspace.
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
// Creates a model connection in a specified workspace and configures the upstream model service endpoint, invoke protocol, and access credentials to support model creation.
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
// Creates a model connection in a specified workspace and configures the upstream model service endpoint, invoke protocol, and access credentials to support model creation.
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
// Creates a Skill in the specified workspace and generates a draft version that can be further edited. You can derive a draft from an existing version or specify a target version and commit message.
//
// Description:
//
// ## Operation description
//
// Creates a Skill in the specified workspace and generates a draft version that can be further edited. You can derive a draft from an existing version or specify a target version and commit message.
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
// Creates a Skill in the specified workspace and generates a draft version that can be further edited. You can derive a draft from an existing version or specify a target version and commit message.
//
// Description:
//
// ## Operation description
//
// Creates a Skill in the specified workspace and generates a draft version that can be further edited. You can derive a draft from an existing version or specify a target version and commit message.
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
// Creates a user in a specified workspace. The username must be unique within the workspace, can contain only lowercase letters, digits, and hyphens, and cannot use reserved names such as manager, admin, or names starting with worker-. If the password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
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
// Creates a user in a specified workspace. The username must be unique within the workspace, can contain only lowercase letters, digits, and hyphens, and cannot use reserved names such as manager, admin, or names starting with worker-. If the password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
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
// Creates an AgentCore workspace control plane record. The server completes the state transition from Initializing to Initialized within the same transaction.
//
// Description:
//
// ## Request description\\nCreates an AgentCore workspace control plane record. The server completes the state transition from `Initializing` to `Initialized` within the same transaction. The network configuration uses `Enabled` to specify whether to enable the VPC network. If enabled, you must provide `VpcId` and at least one `VSwitchIds`.\\n.
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
// Creates an AgentCore workspace control plane record. The server completes the state transition from Initializing to Initialized within the same transaction.
//
// Description:
//
// ## Request description\\nCreates an AgentCore workspace control plane record. The server completes the state transition from `Initializing` to `Initialized` within the same transaction. The network configuration uses `Enabled` to specify whether to enable the VPC network. If enabled, you must provide `VpcId` and at least one `VSwitchIds`.\\n.
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
// Calls a specified model through a published model connection to verify whether the model call chain is available.
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
// Calls a specified model through a published model connection to verify whether the model call chain is available.
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
// Deletes an IM channel of a specified agent.
//
// Description:
//
// Deletes an IM channel of a specified agent.
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
// Deletes an IM channel of a specified agent.
//
// Description:
//
// Deletes an IM channel of a specified agent.
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
// Deletes the specified AgentSpec and all its versions and metadata. The deletion is irreversible.
//
// Description:
//
// ## Operation description
//
// Deletes the specified AgentSpec and all its versions and metadata. The deletion is irreversible.
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
// Deletes the specified AgentSpec and all its versions and metadata. The deletion is irreversible.
//
// Description:
//
// ## Operation description
//
// Deletes the specified AgentSpec and all its versions and metadata. The deletion is irreversible.
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
// Deletes the draft version currently being edited for a specified AgentSpec and clears the draft version pointer.
//
// Description:
//
// ## Request description
//
// Deletes the draft version currently being edited for a specified AgentSpec and clears the draft version pointer.
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
// Deletes the draft version currently being edited for a specified AgentSpec and clears the draft version pointer.
//
// Description:
//
// ## Request description
//
// Deletes the draft version currently being edited for a specified AgentSpec and clears the draft version pointer.
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
// Deletes a credential in a specified workspace and deletes the ciphertext hosted in the credential service. After deletion, agents bound to the credential can no longer obtain the 
//
// Description:
//
// Deletes an access credential in a specified workspace. A credential cannot be deleted if it is still bound to an MCP service.
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
// Deletes an access credential in a specified workspace. A credential cannot be deleted if it is still bound to an MCP service.
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
// Deletes a custom agent from a specified workspace. A successful request indicates that the deletion is accepted. The runtime resources are subsequently cleaned up in the background.
//
// Description:
//
// Before calling this operation, ensure that the workspace is writable. If the agent is still associated with a team, you must dissociate it from the team first.
//
// Deletion is an asynchronous process. After the request is accepted, the associated invocation endpoints become unavailable, and the runtime resources are subsequently cleaned up. The status in the response is typically Deleting. If the agent is already in the deletion process, a subsequent request returns the current summary.
//
// You can call GetCustomAgent to check the deletion progress. After the deletion is complete, a detail query returns a resource not found error.
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
// Deletes a custom agent from a specified workspace. A successful request indicates that the deletion is accepted. The runtime resources are subsequently cleaned up in the background.
//
// Description:
//
// Before calling this operation, ensure that the workspace is writable. If the agent is still associated with a team, you must dissociate it from the team first.
//
// Deletion is an asynchronous process. After the request is accepted, the associated invocation endpoints become unavailable, and the runtime resources are subsequently cleaned up. The status in the response is typically Deleting. If the agent is already in the deletion process, a subsequent request returns the current summary.
//
// You can call GetCustomAgent to check the deletion progress. After the deletion is complete, a detail query returns a resource not found error.
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
// Deletes a specified external agent.
//
// Description:
//
// Deletes a specified external agent.
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
// Deletes a specified external agent.
//
// Description:
//
// Deletes a specified external agent.
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
// Disconnects the identity source integration of a workspace. For CustomSSO, this operation clears the shared user pool and related SAML and SCIM resources. To disable only logon or synchronization, use the corresponding switch API operations.
//
// Description:
//
// ## Features and prerequisites
//
// This operation completely disconnects the integration, rather than simply disabling logon. Disconnecting CustomSSO clears the AC users imported from this identity source, the remote shared user pool and related resources, and the logon publishing configurations and bindings. Therefore, both SAML and SCIM are affected. However, this operation does not delete the enterprise applications created by users on external platforms such as Azure. Before disconnecting, the system checks for references between users and teams. If references exist, the disconnection is rejected.
//
// After the request is accepted, the status returned is DISCONNECTING. You can query the status by calling GetIdentityProvider. The cleanup is complete only when the binding no longer exists (IdentityProviderNotFound). The binding should also no longer appear in the list. DISCONNECT_FAILED indicates that partial cleanup may have occurred, and the state cannot be restored to a completely unconfigured state. You can retry based on the returned status. If the binding does not exist in the first place, the API operation returns a placeholder result indicating no binding without creating new resources. An ongoing disconnection task may return the current binding. Do not use this API operation as a substitute for setting loginEnabled=false, syncEnabled=false, or deleting a single token.
//
// ## Requests and examples
//
// DELETE /workspaces/{workspaceId}/identity-providers/{identityProviderType}
//
// Calling conventions: Call the API operation in the region of the target workspace, using an AK/STS identity that has the corresponding workspace identity source resource permissions. The ws-example, URLs, certificates, and tokens in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log secrets. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// DELETE /workspaces/ws-example/identity-providers/CustomSSO
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "DISCONNECTING"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to upstream services that fail may return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 409 IdentityProvider.Disconnect.Conflict: Users of this identity source still belong to a team. Remove the references first. 409 IdentityProviderOperationConflict: The current status does not allow disconnection. 404 IdentityProviderNotFound: An external binding exists but the type does not match. Asynchronous failures may result in partial cleanup. You cannot claim completion based solely on the first response.
//
// ## Calling relationships and usage boundaries
//
// If you only need to disable a specific capability: To disable SAML logon, use UpdateIdentityProvider and set loginEnabled=false. To disable SCIM, use UpdateIdentityProviderSCIM and set syncEnabled=false. To revoke a single token, use DeleteIdentityProviderSCIMCredential. Call this API operation only when you need to completely disconnect the integration.
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
// Disconnects the identity source integration of a workspace. For CustomSSO, this operation clears the shared user pool and related SAML and SCIM resources. To disable only logon or synchronization, use the corresponding switch API operations.
//
// Description:
//
// ## Features and prerequisites
//
// This operation completely disconnects the integration, rather than simply disabling logon. Disconnecting CustomSSO clears the AC users imported from this identity source, the remote shared user pool and related resources, and the logon publishing configurations and bindings. Therefore, both SAML and SCIM are affected. However, this operation does not delete the enterprise applications created by users on external platforms such as Azure. Before disconnecting, the system checks for references between users and teams. If references exist, the disconnection is rejected.
//
// After the request is accepted, the status returned is DISCONNECTING. You can query the status by calling GetIdentityProvider. The cleanup is complete only when the binding no longer exists (IdentityProviderNotFound). The binding should also no longer appear in the list. DISCONNECT_FAILED indicates that partial cleanup may have occurred, and the state cannot be restored to a completely unconfigured state. You can retry based on the returned status. If the binding does not exist in the first place, the API operation returns a placeholder result indicating no binding without creating new resources. An ongoing disconnection task may return the current binding. Do not use this API operation as a substitute for setting loginEnabled=false, syncEnabled=false, or deleting a single token.
//
// ## Requests and examples
//
// DELETE /workspaces/{workspaceId}/identity-providers/{identityProviderType}
//
// Calling conventions: Call the API operation in the region of the target workspace, using an AK/STS identity that has the corresponding workspace identity source resource permissions. The ws-example, URLs, certificates, and tokens in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log secrets. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// DELETE /workspaces/ws-example/identity-providers/CustomSSO
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "DISCONNECTING"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to upstream services that fail may return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 409 IdentityProvider.Disconnect.Conflict: Users of this identity source still belong to a team. Remove the references first. 409 IdentityProviderOperationConflict: The current status does not allow disconnection. 404 IdentityProviderNotFound: An external binding exists but the type does not match. Asynchronous failures may result in partial cleanup. You cannot claim completion based solely on the first response.
//
// ## Calling relationships and usage boundaries
//
// If you only need to disable a specific capability: To disable SAML logon, use UpdateIdentityProvider and set loginEnabled=false. To disable SCIM, use UpdateIdentityProviderSCIM and set syncEnabled=false. To revoke a single token, use DeleteIdentityProviderSCIMCredential. Call this API operation only when you need to completely disconnect the integration.
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
// 删除 SCIM 凭据
//
// Description:
//
// ## 功能与前提\n撤销指定凭据，credentialId 必填且非空，从凭据列表取得。要求共享池可用，但不要求 SAML 已配置或 SCIM 接收为 Enabled。删除后使用该 Token 的外部预配请求会失效；只影响选定凭据，不自动关闭接收、不自动关闭 AC syncEnabled，也不影响另一枚凭据或 SAML。\n即使删除最后一枚凭据，也不自动回写同步开关或重发新 Token；用户应按需要先停用外部预配或更换其使用的 Token。成功返回公共响应，无业务 data 内容；刷新列表确认。对未知或已删除 ID 的结果由上游决定，不承诺统一 404 或无条件幂等成功。\n\n## 请求与示例\nDELETE /workspaces/{workspaceId}/scim-provisioning/credentials/{credentialId}\n调用约定：按目标 Workspace 的地域调用，使用有相应 Workspace 身份源资源权限的 AK/STS 身份。下文示例中的 ws-example、URL、证书和 Token 均为占位值，需替换为真实值。请求体示例表示公开 API 的 body 参数内容，不是后端直连接口额外的 body 包装。返回示例只展示业务部分，外层还包括 requestId、success、code、message、httpStatusCode；排障请保留 requestId，不记录密钥。业务任务异步不意味着 HTTP 请求长连接等待完成。\n```text\nDELETE /workspaces/ws-example/scim-provisioning/credentials/credential-example\n```\n\n## 返回示例\n```json\n{\n  "success": true,\n  "httpStatusCode": 200\n}\n```\n\n## 错误与处理\n错误处理：400 InvalidParameter 表示参数或当前配置前提不满足；404 IdentityProviderNotFound 表示没有匹配类型的绑定。同步调用上游失败可能返回 503 IdentityProviderProvisioningUnavailable；保留请求 ID 排查，不把读取失败解释为未配置。实际鉴权和 Workspace 不存在错误沿用公共错误处理。\n400 InvalidParameter：credentialId 为空或共享池不可用。404 IdentityProviderNotFound 指绑定不存在，不代表所有未知凭据 ID 都返回该码。其他删除错误按上游结果处理，失败时刷新列表核对，不自动创建替代 Token。\n\n## 调用关系与使用边界\n使用 ListIdentityProviderSCIMCredentials 取得 credentialId 后删除指定凭据。若外部预配系统仍使用该 Token，后续预配可能鉴权失败，因此应先确认实际使用情况。删除后不会自动创建新凭据、关闭 AC 同步开关或关闭 SAML 登录。
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
// 删除 SCIM 凭据
//
// Description:
//
// ## 功能与前提\n撤销指定凭据，credentialId 必填且非空，从凭据列表取得。要求共享池可用，但不要求 SAML 已配置或 SCIM 接收为 Enabled。删除后使用该 Token 的外部预配请求会失效；只影响选定凭据，不自动关闭接收、不自动关闭 AC syncEnabled，也不影响另一枚凭据或 SAML。\n即使删除最后一枚凭据，也不自动回写同步开关或重发新 Token；用户应按需要先停用外部预配或更换其使用的 Token。成功返回公共响应，无业务 data 内容；刷新列表确认。对未知或已删除 ID 的结果由上游决定，不承诺统一 404 或无条件幂等成功。\n\n## 请求与示例\nDELETE /workspaces/{workspaceId}/scim-provisioning/credentials/{credentialId}\n调用约定：按目标 Workspace 的地域调用，使用有相应 Workspace 身份源资源权限的 AK/STS 身份。下文示例中的 ws-example、URL、证书和 Token 均为占位值，需替换为真实值。请求体示例表示公开 API 的 body 参数内容，不是后端直连接口额外的 body 包装。返回示例只展示业务部分，外层还包括 requestId、success、code、message、httpStatusCode；排障请保留 requestId，不记录密钥。业务任务异步不意味着 HTTP 请求长连接等待完成。\n```text\nDELETE /workspaces/ws-example/scim-provisioning/credentials/credential-example\n```\n\n## 返回示例\n```json\n{\n  "success": true,\n  "httpStatusCode": 200\n}\n```\n\n## 错误与处理\n错误处理：400 InvalidParameter 表示参数或当前配置前提不满足；404 IdentityProviderNotFound 表示没有匹配类型的绑定。同步调用上游失败可能返回 503 IdentityProviderProvisioningUnavailable；保留请求 ID 排查，不把读取失败解释为未配置。实际鉴权和 Workspace 不存在错误沿用公共错误处理。\n400 InvalidParameter：credentialId 为空或共享池不可用。404 IdentityProviderNotFound 指绑定不存在，不代表所有未知凭据 ID 都返回该码。其他删除错误按上游结果处理，失败时刷新列表核对，不自动创建替代 Token。\n\n## 调用关系与使用边界\n使用 ListIdentityProviderSCIMCredentials 取得 credentialId 后删除指定凭据。若外部预配系统仍使用该 Token，后续预配可能鉴权失败，因此应先确认实际使用情况。删除后不会自动创建新凭据、关闭 AC 同步开关或关闭 SAML 登录。
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
// Deletes an image cache.
//
// Description:
//
// Deletes the logical image cache of the current workspace. The deletion requires that no enabled agent configurations reference the cache and no active instances are using it. Concurrency is rechecked upon submission. Caches in the PENDING or FAILED state can still be deleted, and the backend handles the cleanup. An HTTP 200 response returns an imageId that matches the request. If deletionPending is true, the current workspace is still undergoing deletion, and a Get request returns DELETING until ImageCacheNotFound is explicitly returned. If false, the deletion is complete or the current workspace no longer exists. Usage by other workspaces does not block the deletion for the current workspace. The physical content is reclaimed only after all workspaces in the same region have removed their associations and no active instances are using it.
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
// Deletes an image cache.
//
// Description:
//
// Deletes the logical image cache of the current workspace. The deletion requires that no enabled agent configurations reference the cache and no active instances are using it. Concurrency is rechecked upon submission. Caches in the PENDING or FAILED state can still be deleted, and the backend handles the cleanup. An HTTP 200 response returns an imageId that matches the request. If deletionPending is true, the current workspace is still undergoing deletion, and a Get request returns DELETING until ImageCacheNotFound is explicitly returned. If false, the deletion is complete or the current workspace no longer exists. Usage by other workspaces does not block the deletion for the current workspace. The physical content is reclaimed only after all workspaces in the same region have removed their associations and no active instances are using it.
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
// Deletes a specified managed agent.
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
// Deletes a specified managed agent.
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
// Deletes a specified MCP server. The deletion is an asynchronous process, and the MCP server is no longer returned after the process is complete.
//
// Description:
//
// ## Request description
//
// Deletes a specified MCP server. The deletion is an asynchronous process, and the MCP server is no longer returned after the process is complete.
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
// Deletes a specified MCP server. The deletion is an asynchronous process, and the MCP server is no longer returned after the process is complete.
//
// Description:
//
// ## Request description
//
// Deletes a specified MCP server. The deletion is an asynchronous process, and the MCP server is no longer returned after the process is complete.
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
// Deletes a long-term memory based on the memory ID.
//
// Description:
//
// Deletes a long-term memory based on the memory ID.
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
// Deletes a long-term memory based on the memory ID.
//
// Description:
//
// Deletes a long-term memory based on the memory ID.
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
// Asynchronously deletes a MemoryStore and all long-term memory within it.
//
// Description:
//
// Asynchronously deletes a MemoryStore and all long-term memory within it.
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
// Asynchronously deletes a MemoryStore and all long-term memory within it.
//
// Description:
//
// Asynchronously deletes a MemoryStore and all long-term memory within it.
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
// Deletes a model from a specified workspace. Models that are currently in use cannot be deleted.
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
// Deletes a model from a specified workspace. Models that are currently in use cannot be deleted.
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
// Deletes a skill and its related version data from a specified workspace. This operation is irreversible.
//
// Description:
//
// ## Operation description
//
// Deletes a skill and its related version data from a specified workspace. This operation is irreversible.
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
// Deletes a skill and its related version data from a specified workspace. This operation is irreversible.
//
// Description:
//
// ## Operation description
//
// Deletes a skill and its related version data from a specified workspace. This operation is irreversible.
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
// Deletes the draft version currently being edited for a specified Skill.
//
// Description:
//
// ## Operation description
//
// Deletes the draft version currently being edited for a specified Skill.
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
// Deletes the draft version currently being edited for a specified Skill.
//
// Description:
//
// ## Operation description
//
// Deletes the draft version currently being edited for a specified Skill.
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
// Deletes a team from a specified workspace. Deleting a team does not delete the users or agents within it. Only the membership associations are removed.
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
// Deletes a team from a specified workspace. Deleting a team does not delete the users or agents within it. Only the membership associations are removed.
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
// Deletes a user from a specified workspace. A user cannot be deleted while the user is still a member of any team. Remove the user from all teams before deleting the user.
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
// Deletes a user from a specified workspace. A user cannot be deleted while the user is still a member of any team. Remove the user from all teams before deleting the user.
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
// Deletes the control plane record of a specified workspace. The server completes the state transition from Deleting to Deleted within the same transaction. The server processes repeated deletion requests for workspaces in the Deleting or Deleted state idempotently.
//
// Description:
//
// ## Request description\\nDeletes the control plane record of a specified workspace. The server completes the state transition from `Deleting` to `Deleted` within the same transaction. The server processes repeated deletion requests for workspaces in the `Deleting` or `Deleted` state idempotently.\\n.
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
// Deletes the control plane record of a specified workspace. The server completes the state transition from Deleting to Deleted within the same transaction. The server processes repeated deletion requests for workspaces in the Deleting or Deleted state idempotently.
//
// Description:
//
// ## Request description\\nDeletes the control plane record of a specified workspace. The server completes the state transition from `Deleting` to `Deleted` within the same transaction. The server processes repeated deletion requests for workspaces in the `Deleting` or `Deleted` state idempotently.\\n.
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
// Disables a QoderCLI Connector in a specified workspace. The Connector can be disabled only when no Agent is using it. After the Connector is disabled, the Connector and its Service Account Key configuration are removed.
//
// Description:
//
// Disables a Connector in a specified workspace.
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
// Disables a QoderCLI Connector in a specified workspace. The Connector can be disabled only when no Agent is using it. After the Connector is disabled, the Connector and its Service Account Key configuration are removed.
//
// Description:
//
// Disables a Connector in a specified workspace.
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
// Retrieves the presigned OSS download URL of a specified AgentSpec to download the AgentSpec ZIP package.
//
// Description:
//
// ## Request description
//
// Retrieves the presigned OSS download URL of a specified AgentSpec to download the AgentSpec ZIP package.
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
// Retrieves the presigned OSS download URL of a specified AgentSpec to download the AgentSpec ZIP package.
//
// Description:
//
// ## Request description
//
// Retrieves the presigned OSS download URL of a specified AgentSpec to download the AgentSpec ZIP package.
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
// Retrieves a pre-signed OSS download URL for a specified Skill version, which is used to download the corresponding Skill ZIP package.
//
// Description:
//
// ## Operation description
//
// Retrieves a pre-signed OSS download URL for a specified Skill version, which is used to download the corresponding Skill ZIP package.
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
// Retrieves a pre-signed OSS download URL for a specified Skill version, which is used to download the corresponding Skill ZIP package.
//
// Description:
//
// ## Operation description
//
// Retrieves a pre-signed OSS download URL for a specified Skill version, which is used to download the corresponding Skill ZIP package.
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
// Enables a QoderCLI Connector in a specified workspace. The request must include the site, API key, and at least one service account key. The service validates all account keys before saving the configuration.
//
// Description:
//
// Enables a Connector in a specified workspace. Credential verification is required before the Connector can be enabled.
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
// Enables a QoderCLI Connector in a specified workspace. The request must include the site, API key, and at least one service account key. The service validates all account keys before saving the configuration.
//
// Description:
//
// Enables a Connector in a specified workspace. Credential verification is required before the Connector can be enabled.
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
// Skips the regular review process and forcibly publishes a specified Skill version.
//
// Description:
//
// ## Operation description
//
// Skips the regular review process and forcibly publishes a specified Skill version.
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
// Skips the regular review process and forcibly publishes a specified Skill version.
//
// Description:
//
// ## Operation description
//
// Skips the regular review process and forcibly publishes a specified Skill version.
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
// Retrieves the download URL for the saved code package of an agent.
//
// Description:
//
// Generates a time-limited, read-only download URL for the code package referenced by the currently saved agent configuration. This applies to code packages uploaded locally or imported from Object Storage Service (OSS). The operation uses workspaceId and agentId for locating and does not accept arbitrary bucket, object key, or historical version parameters. The current configuration must reference an available, fixed code package, and does not require waiting for a running instance to be ready. An error is returned if the code package does not exist, has changed, or is unreadable. The link is valid for up to 10 minutes and does not exceed the underlying credential expiration time. Call the operation again after it expires. The resourceVersionId indicates the saved configuration version corresponding to this download. The sample download link is redacted. Use the downloadUrl in the response.
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
// Retrieves the download URL for the saved code package of an agent.
//
// Description:
//
// Generates a time-limited, read-only download URL for the code package referenced by the currently saved agent configuration. This applies to code packages uploaded locally or imported from Object Storage Service (OSS). The operation uses workspaceId and agentId for locating and does not accept arbitrary bucket, object key, or historical version parameters. The current configuration must reference an available, fixed code package, and does not require waiting for a running instance to be ready. An error is returned if the code package does not exist, has changed, or is unreadable. The link is valid for up to 10 minutes and does not exceed the underlying credential expiration time. Call the operation again after it expires. The resourceVersionId indicates the saved configuration version corresponding to this download. The sample download link is redacted. Use the downloadUrl in the response.
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
// Retrieves the OSS RAM authorization URL for the agent code package.
//
// Description:
//
// Generates a RAM authorization URL for the resource role of the current workspace for a specified source bucket. The client cannot specify or override the role. After you open the authorizeUrl and confirm the authorization, call VerifyAgentCodePackageOssRamAuthorization to check the policy binding, and call ValidateAgentCodePackageFromOss to check specific objects.
//
// This authorization reuses the OSSObjectRestoreCustom template, which includes permissions to read, write, and delete objects, and is not a read-only authorization. Retrieving the URL itself does not complete the authorization, transfer the code package, or modify the default OSS binding of the workspace. Existing authorizations for the same role and bucket can be reused. The example URL is provided to demonstrate the format. Use the actual authorization URL returned by the API operation.
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
// Retrieves the OSS RAM authorization URL for the agent code package.
//
// Description:
//
// Generates a RAM authorization URL for the resource role of the current workspace for a specified source bucket. The client cannot specify or override the role. After you open the authorizeUrl and confirm the authorization, call VerifyAgentCodePackageOssRamAuthorization to check the policy binding, and call ValidateAgentCodePackageFromOss to check specific objects.
//
// This authorization reuses the OSSObjectRestoreCustom template, which includes permissions to read, write, and delete objects, and is not a read-only authorization. Retrieving the URL itself does not complete the authorization, transfer the code package, or modify the default OSS binding of the workspace. Existing authorizations for the same role and bucket can be reused. The example URL is provided to demonstrate the format. Use the actual authorization URL returned by the API operation.
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
// Retrieves the upload URL for an agent code package.
//
// Description:
//
// Retrieves a time-limited upload URL for a code package without uploading the file content. The fileName must be a ZIP or TAR.GZ filename without directories. The format must match the file extension. The fileSize is in bytes, with valid values ranging from 1 to 500,000,000.
//
// Use the uploadMethod, uploadUrl, and headers from the response to upload the file directly. After the upload is complete, pass the sourceUri to the artifact.code.sourceUri parameter of CreateCustomAgent or UpdateCustomAgent. The actual object and size are verified and a fixed copy is saved during creation or update. The upload URL is valid for up to 10 minutes and does not exceed the expiration time of the underlying temporary credentials. Retrieve a new URL after it expires. The example URL is desensitized and cannot be used for actual uploads. Use the original URL and request headers from this response.
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
// Retrieves the upload URL for an agent code package.
//
// Description:
//
// Retrieves a time-limited upload URL for a code package without uploading the file content. The fileName must be a ZIP or TAR.GZ filename without directories. The format must match the file extension. The fileSize is in bytes, with valid values ranging from 1 to 500,000,000.
//
// Use the uploadMethod, uploadUrl, and headers from the response to upload the file directly. After the upload is complete, pass the sourceUri to the artifact.code.sourceUri parameter of CreateCustomAgent or UpdateCustomAgent. The actual object and size are verified and a fixed copy is saved during creation or update. The upload URL is valid for up to 10 minutes and does not exceed the expiration time of the underlying temporary credentials. Retrieve a new URL after it expires. The example URL is desensitized and cannot be used for actual uploads. Use the original URL and request headers from this response.
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
// ## Request description
//
// Queries the governance details of a specified AgentSpec, including basic information, governance pointers, and summaries of all versions.
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
// ## Request description
//
// Queries the governance details of a specified AgentSpec, including basic information, governance pointers, and summaries of all versions.
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
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API to complete the import.
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
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS presigned upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, you can call the AgentSpec OSS upload API to complete the import.
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
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing.
//
// Description:
//
// ## Operation description
//
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing.
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
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing.
//
// Description:
//
// ## Operation description
//
// Queries the full content of the latest version of a specified AgentSpec for editing or viewing.
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
// Queries the complete content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
//
// Description:
//
// ## Request description
//
// Queries the complete content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
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
// Queries the complete content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
//
// Description:
//
// ## Request description
//
// Queries the complete content of a specified AgentSpec version, including the manifest content, resource files, and metadata.
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
// Queries the details of a specified credential and returns the list of agents bound to the credential. The credential content is returned in a masked format.
//
// Description:
//
// Queries the details of a single credential. Sensitive fields are not returned.
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
// Queries the details of a single credential. Sensitive fields are not returned.
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
// Queries the details of a custom agent in a specified workspace, including deployment configurations, the execute status of the latest configuration version, and the running instance status statistics.
//
// Description:
//
// After creating or updating a custom agent, you can poll this operation to query the execution result of the latest configuration version. latestVersionStatus corresponds to latestSpecVersion. latestExecutionError indicates the cause of the pending or failed status for this version. If retryable is true, you can query again later.
//
// The returned configuration is the latest submitted configuration and may not have taken effect. Verify the deployment and running status by checking latestVersionStatus, the overall agent status, and the running instance statistics sandboxPhaseCounts.
//
// If the queried custom agent does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
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
// Queries the details of a custom agent in a specified workspace, including deployment configurations, the execute status of the latest configuration version, and the running instance status statistics.
//
// Description:
//
// After creating or updating a custom agent, you can poll this operation to query the execution result of the latest configuration version. latestVersionStatus corresponds to latestSpecVersion. latestExecutionError indicates the cause of the pending or failed status for this version. If retryable is true, you can query again later.
//
// The returned configuration is the latest submitted configuration and may not have taken effect. Verify the deployment and running status by checking latestVersionStatus, the overall agent status, and the running instance statistics sandboxPhaseCounts.
//
// If the queried custom agent does not exist, has been deleted, or does not belong to the specified workspace, a resource not found error is returned.
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
// Queries the details of a specified external agent.
//
// Description:
//
// Queries the details of a specified external agent.
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
// Queries the details of a specified external agent.
//
// Description:
//
// Queries the details of a specified external agent.
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
// Queries the available access network types for a specified external agent.
//
// Description:
//
// Queries the available access network types for a specified external agent.
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
// Queries the available access network types for a specified external agent.
//
// Description:
//
// Queries the available access network types for a specified external agent.
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
// Queries the binding and expected configuration of a specified identity provider. When CustomSSO is configured with SAML, this operation reads back the actual upstream SSO configuration. The SCIM status is queried through a separate API operation.
//
// Description:
//
// ## Features and prerequisites
//
// Reads existing bindings by type. CustomSSO returns the saved protocol, standardized saml, expected loginEnabled, syncEnabled, and shared binding status. If SAML is not configured, the saml and configuration fields may not be returned. When saml is configured, the system additionally queries AgentIdentity, and configuration represents the actual SSO configuration. The saml field is the data source for the edit form. The status switch for configuration can be Enabled or Disabled. If the upstream query fails, an error is reported. Do not treat this as an unconfigured state.
//
// Only this operation reads back the actual SAML configuration. List and write operations do not guarantee that configuration is returned. A user pool in the READY state does not indicate that SAML is configured. UPDATING and UPDATE_FAILED describe logon configuration tasks, not SCIM tasks. The vendor callback URL for CustomSSO is empty. Use GetIdentityProviderSCIM.operationStatus to query SCIM operation results.
//
// ## Request and examples
//
// GET /workspaces/{workspaceId}/identity-providers/{identityProviderType}
//
// Calling conventions: Call the API operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example shows the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business data. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log keys. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/identity-providers/CustomSSO
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "READY",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": false
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call to the upstream service may fail and return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret a read failure as an unconfigured state. Actual authentication and workspace not found errors follow common error handling.
//
// 404 IdentityProviderNotFound: No binding exists, or the input type is inconsistent with the current binding. If SAML is already configured, an upstream readback failure does not degrade to a successful empty configuration.
//
// ## Call relationships and usage boundaries
//
// SAML configuration completion and logon enabling are independent operations. protocol=SAML and the presence of saml indicate that the logon configuration has been submitted, and loginEnabled indicates the expected logon switch. Check the effective status based on the actual SSO status in configuration. SCIM uses an independent query API operation. Do not use status=READY of this API operation as the basis for successful SCIM synchronization. Currently, OIDC, JIT configuration, or JIT status fields are not provided.
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
// Queries the binding and expected configuration of a specified identity provider. When CustomSSO is configured with SAML, this operation reads back the actual upstream SSO configuration. The SCIM status is queried through a separate API operation.
//
// Description:
//
// ## Features and prerequisites
//
// Reads existing bindings by type. CustomSSO returns the saved protocol, standardized saml, expected loginEnabled, syncEnabled, and shared binding status. If SAML is not configured, the saml and configuration fields may not be returned. When saml is configured, the system additionally queries AgentIdentity, and configuration represents the actual SSO configuration. The saml field is the data source for the edit form. The status switch for configuration can be Enabled or Disabled. If the upstream query fails, an error is reported. Do not treat this as an unconfigured state.
//
// Only this operation reads back the actual SAML configuration. List and write operations do not guarantee that configuration is returned. A user pool in the READY state does not indicate that SAML is configured. UPDATING and UPDATE_FAILED describe logon configuration tasks, not SCIM tasks. The vendor callback URL for CustomSSO is empty. Use GetIdentityProviderSCIM.operationStatus to query SCIM operation results.
//
// ## Request and examples
//
// GET /workspaces/{workspaceId}/identity-providers/{identityProviderType}
//
// Calling conventions: Call the API operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example shows the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business data. The outer layer also includes requestId, success, code, message, and httpStatusCode. Keep the requestId for troubleshooting and do not log keys. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/identity-providers/CustomSSO
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "READY",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": false
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call to the upstream service may fail and return 503 IdentityProviderProvisioningUnavailable. Keep the request ID for troubleshooting and do not interpret a read failure as an unconfigured state. Actual authentication and workspace not found errors follow common error handling.
//
// 404 IdentityProviderNotFound: No binding exists, or the input type is inconsistent with the current binding. If SAML is already configured, an upstream readback failure does not degrade to a successful empty configuration.
//
// ## Call relationships and usage boundaries
//
// SAML configuration completion and logon enabling are independent operations. protocol=SAML and the presence of saml indicate that the logon configuration has been submitted, and loginEnabled indicates the expected logon switch. Check the effective status based on the actual SSO status in configuration. SCIM uses an independent query API operation. Do not use status=READY of this API operation as the basis for successful SCIM synchronization. Currently, OIDC, JIT configuration, or JIT status fields are not provided.
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
// Queries the actual SCIM receiving configuration, the expected AC synchronization switch, and the status of the most recent independent SCIM task. Prior SAML configuration is not required.
//
// Description:
//
// ## Features and prerequisites
//
// The shared CustomSSO user pool must be created and must not be in the overall dissociation process or have failed to be dissociated. Prior SAML configuration is not required. SAML updates in progress or failed SAML updates do not block the query. The configuration.status and baseUrl parameters are obtained from AgentIdentity. The syncEnabled parameter is the expected AC synchronization switch. The operationStatus parameter indicates the most recent independent SCIM switch and the first full import task when the switch is enabled. The presence of a URL does not mean that external provisioning is complete, nor does it mean that AC synchronization is enabled.
//
// operationStatus: PENDING (waiting), RUNNING (in progress), PAUSED (paused), SUCCESS (succeeded), FAILED (failed), CANCELED (canceled). If there are no historical tasks, it returns null or is not returned. You can poll when the status is PENDING or RUNNING. Stop polling when a terminal state is reached. You can manually refresh when the status is PAUSED. This operation does not provide taskId or failure details, and does not indicate credential issuance, external provisioning connection tests, or the health of continuous incremental synchronization. After it is disabled, the URL, credentials, and historical task records may still be retained.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/scim-provisioning
//
// Calling conventions: Call this operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/scim-provisioning
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "syncEnabled": true,
//
//	  "operationStatus": "SUCCESS"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that there is no binding of the matching type. Synchronous calls to upstream services that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow common error handling.
//
// 400 InvalidParameter: The shared user pool is unavailable, being dissociated, or failed to be dissociated. If there is no binding, 404 IdentityProviderNotFound is returned. Upstream read failures must retain the error and cannot be mapped to Disabled, unconfigured, or synchronization success.
//
// ## Call relationships and usage boundaries
//
// SCIM configuration sequence: After preparing the shared pool, call PrepareIdentityProviderSCIM and configure the returned baseUrl and newly issued token to the provisioning system of the external IdP. After completing the external connection test, call UpdateIdentityProviderSCIM to enable synchronization, and then call this operation to query operationStatus. SUCCESS only indicates that the most recent SCIM task is completed, and does not represent the health of continuous incremental synchronization or the success of the external provisioning connection test.
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
// Queries the actual SCIM receiving configuration, the expected AC synchronization switch, and the status of the most recent independent SCIM task. Prior SAML configuration is not required.
//
// Description:
//
// ## Features and prerequisites
//
// The shared CustomSSO user pool must be created and must not be in the overall dissociation process or have failed to be dissociated. Prior SAML configuration is not required. SAML updates in progress or failed SAML updates do not block the query. The configuration.status and baseUrl parameters are obtained from AgentIdentity. The syncEnabled parameter is the expected AC synchronization switch. The operationStatus parameter indicates the most recent independent SCIM switch and the first full import task when the switch is enabled. The presence of a URL does not mean that external provisioning is complete, nor does it mean that AC synchronization is enabled.
//
// operationStatus: PENDING (waiting), RUNNING (in progress), PAUSED (paused), SUCCESS (succeeded), FAILED (failed), CANCELED (canceled). If there are no historical tasks, it returns null or is not returned. You can poll when the status is PENDING or RUNNING. Stop polling when a terminal state is reached. You can manually refresh when the status is PAUSED. This operation does not provide taskId or failure details, and does not indicate credential issuance, external provisioning connection tests, or the health of continuous incremental synchronization. After it is disabled, the URL, credentials, and historical task records may still be retained.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/scim-provisioning
//
// Calling conventions: Call this operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/scim-provisioning
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "syncEnabled": true,
//
//	  "operationStatus": "SUCCESS"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that there is no binding of the matching type. Synchronous calls to upstream services that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow common error handling.
//
// 400 InvalidParameter: The shared user pool is unavailable, being dissociated, or failed to be dissociated. If there is no binding, 404 IdentityProviderNotFound is returned. Upstream read failures must retain the error and cannot be mapped to Disabled, unconfigured, or synchronization success.
//
// ## Call relationships and usage boundaries
//
// SCIM configuration sequence: After preparing the shared pool, call PrepareIdentityProviderSCIM and configure the returned baseUrl and newly issued token to the provisioning system of the external IdP. After completing the external connection test, call UpdateIdentityProviderSCIM to enable synchronization, and then call this operation to query operationStatus. SUCCESS only indicates that the most recent SCIM task is completed, and does not represent the health of continuous incremental synchronization or the success of the external provisioning connection test.
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
// 查询 SAML 服务提供商信息
//
// Description:
//
// ## 功能与前提\n要求已有可用 CustomSSO 共享用户池，identityProviderType=CustomSSO，protocol 必填且仅支持精确值 SAML；不要求先上传 IdP Metadata。返回我方 SP 的 entityId、acsUrl、metadataXml，供用户配置外部 IdP。它们不是外部 IdP 的登录 URL 或 IdP Entity ID，不能将 SP XML 作为 IdP XML 提交 UpdateIdentityProvider。\n本接口只读取，不启用登录、不开启 SCIM、不创建凭据、不执行登录连接测试。共享池初始化中、解绑中或解绑失败时不可读取。\n\n## 请求与示例\nGET /workspaces/{workspaceId}/identity-providers/{identityProviderType}/service-provider\n调用约定：按目标 Workspace 的地域调用，使用有相应 Workspace 身份源资源权限的 AK/STS 身份。下文示例中的 ws-example、URL、证书和 Token 均为占位值，需替换为真实值。请求体示例表示公开 API 的 body 参数内容，不是后端直连接口额外的 body 包装。返回示例只展示业务部分，外层还包括 requestId、success、code、message、httpStatusCode；排障请保留 requestId，不记录密钥。业务任务异步不意味着 HTTP 请求长连接等待完成。\n```text\nGET /workspaces/ws-example/identity-providers/CustomSSO/service-provider?protocol=SAML\n```\n\n## 返回示例\n```json\n{\n  "data": {\n    "protocol": "SAML",\n    "entityId": "<我方 SP Entity ID>",\n    "acsUrl": "https://example.com/saml/acs",\n    "metadataXml": "<我方 SP Metadata XML>"\n  }\n}\n```\n\n## 错误与处理\n错误处理：400 InvalidParameter 表示参数或当前配置前提不满足；404 IdentityProviderNotFound 表示没有匹配类型的绑定。同步调用上游失败可能返回 503 IdentityProviderProvisioningUnavailable；保留请求 ID 排查，不把读取失败解释为未配置。实际鉴权和 Workspace 不存在错误沿用公共错误处理。\n400 InvalidParameter：identityProviderType 或 protocol 不支持。404 IdentityProviderNotFound：尚无 CustomSSO 绑定。409 IdentityProviderOperationConflict：用户池未准备好、解绑中或解绑失败。\n\n## 调用关系与使用边界\nSAML 配置顺序：CreateIdentityProvider 初始化 CustomSSO 共享池，等待 GetIdentityProvider 返回 READY；调用本接口取得 SP 信息并配置到外部 IdP；再将外部 IdP Metadata XML 或完整手填配置提交 UpdateIdentityProvider。保存后按需单独启用 loginEnabled。SP 信息查询不会完成外部 IdP 配置或验证登录。
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
// 查询 SAML 服务提供商信息
//
// Description:
//
// ## 功能与前提\n要求已有可用 CustomSSO 共享用户池，identityProviderType=CustomSSO，protocol 必填且仅支持精确值 SAML；不要求先上传 IdP Metadata。返回我方 SP 的 entityId、acsUrl、metadataXml，供用户配置外部 IdP。它们不是外部 IdP 的登录 URL 或 IdP Entity ID，不能将 SP XML 作为 IdP XML 提交 UpdateIdentityProvider。\n本接口只读取，不启用登录、不开启 SCIM、不创建凭据、不执行登录连接测试。共享池初始化中、解绑中或解绑失败时不可读取。\n\n## 请求与示例\nGET /workspaces/{workspaceId}/identity-providers/{identityProviderType}/service-provider\n调用约定：按目标 Workspace 的地域调用，使用有相应 Workspace 身份源资源权限的 AK/STS 身份。下文示例中的 ws-example、URL、证书和 Token 均为占位值，需替换为真实值。请求体示例表示公开 API 的 body 参数内容，不是后端直连接口额外的 body 包装。返回示例只展示业务部分，外层还包括 requestId、success、code、message、httpStatusCode；排障请保留 requestId，不记录密钥。业务任务异步不意味着 HTTP 请求长连接等待完成。\n```text\nGET /workspaces/ws-example/identity-providers/CustomSSO/service-provider?protocol=SAML\n```\n\n## 返回示例\n```json\n{\n  "data": {\n    "protocol": "SAML",\n    "entityId": "<我方 SP Entity ID>",\n    "acsUrl": "https://example.com/saml/acs",\n    "metadataXml": "<我方 SP Metadata XML>"\n  }\n}\n```\n\n## 错误与处理\n错误处理：400 InvalidParameter 表示参数或当前配置前提不满足；404 IdentityProviderNotFound 表示没有匹配类型的绑定。同步调用上游失败可能返回 503 IdentityProviderProvisioningUnavailable；保留请求 ID 排查，不把读取失败解释为未配置。实际鉴权和 Workspace 不存在错误沿用公共错误处理。\n400 InvalidParameter：identityProviderType 或 protocol 不支持。404 IdentityProviderNotFound：尚无 CustomSSO 绑定。409 IdentityProviderOperationConflict：用户池未准备好、解绑中或解绑失败。\n\n## 调用关系与使用边界\nSAML 配置顺序：CreateIdentityProvider 初始化 CustomSSO 共享池，等待 GetIdentityProvider 返回 READY；调用本接口取得 SP 信息并配置到外部 IdP；再将外部 IdP Metadata XML 或完整手填配置提交 UpdateIdentityProvider。保存后按需单独启用 loginEnabled。SP 信息查询不会完成外部 IdP 配置或验证登录。
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
// Queries the status, current step, number of processed users, and failure cause of a manual synchronization task by using the local taskId of AgentCore. The taskId must be obtained from CreateIdentityProviderSyncJob in the same workspace and identity provider. It is not the synchronizationJobId of AgentIdentity. Poll the task until the status is SUCCESS, FAILED, or CANCELED. The PENDING, RUNNING, and PAUSED statuses do not indicate completion. The syncedUserCount includes updated existing users and does not equal the number of newly added users or the progress percentage.
//
// Authorization information: The RAM operation is agentcore:GetIdentityProviderSyncJob, the access level is Read, and the resource ARN is acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is a required authorization resource.
//
// Idempotence and retry: This operation is read-only and can be repeatedly queried.
//
// Calling this API operation is free of charge, and operation records support ActionTrail auditing.
//
// Response example:
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
//	"data": {
//
//	  "taskId": "12345678-1234-4234-8234-123456789abc",
//
//	  "status": "SUCCESS",
//
//	  "currentStep": "SYNC",
//
//	  "synchronizationJobId": "syncjob_0123456789abcdef0123456789abcdef",
//
//	  "syncedUserCount": 12
//
//	}
//
// }
//
// ```
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
// Queries the status, current step, number of processed users, and failure cause of a manual synchronization task by using the local taskId of AgentCore. The taskId must be obtained from CreateIdentityProviderSyncJob in the same workspace and identity provider. It is not the synchronizationJobId of AgentIdentity. Poll the task until the status is SUCCESS, FAILED, or CANCELED. The PENDING, RUNNING, and PAUSED statuses do not indicate completion. The syncedUserCount includes updated existing users and does not equal the number of newly added users or the progress percentage.
//
// Authorization information: The RAM operation is agentcore:GetIdentityProviderSyncJob, the access level is Read, and the resource ARN is acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is a required authorization resource.
//
// Idempotence and retry: This operation is read-only and can be repeatedly queried.
//
// Calling this API operation is free of charge, and operation records support ActionTrail auditing.
//
// Response example:
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
//	"data": {
//
//	  "taskId": "12345678-1234-4234-8234-123456789abc",
//
//	  "status": "SUCCESS",
//
//	  "currentStep": "SYNC",
//
//	  "synchronizationJobId": "syncjob_0123456789abcdef0123456789abcdef",
//
//	  "syncedUserCount": 12
//
//	}
//
// }
//
// ```
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
// Queries the saved synchronization scope, the validity of selected departments, and the estimated user count downstream of the identity provider. This operation is not used for real-time estimation of unsaved selections. AllAuthorized indicates the address book scope authorized for the application, which is not equivalent to the entire enterprise organization. An empty selectedDepartments array does not prove that the organization has no sub-departments. The estimation can be used only when estimateStatus is Ready and a count is returned. A missing count does not indicate 0.
//
// Authorization information: RAM operation agentcore:GetIdentityProviderSyncScope, access level Read, resource ARN: acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is a required authorization resource.
//
// Idempotence and retry: This operation is read-only and can be queried repeatedly.
//
// Calling this API operation is free of charge. Operation records support ActionTrail auditing.
//
// Sample response:
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
//	"data": {
//
//	  "syncScope": {
//
//	    "mode": "SelectedDepartments",
//
//	    "departmentIds": [
//
//	      "10001"
//
//	    ]
//
//	  },
//
//	  "selectedDepartments": [
//
//	    {
//
//	      "id": "10001",
//
//	      "name": "R&D Department",
//
//	      "exists": true
//
//	    }
//
//	  ],
//
//	  "estimateStatus": "Ready",
//
//	  "expandedDepartmentCount": 2,
//
//	  "estimatedUserCount": 12
//
//	}
//
// }
//
// ```
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
// Queries the saved synchronization scope, the validity of selected departments, and the estimated user count downstream of the identity provider. This operation is not used for real-time estimation of unsaved selections. AllAuthorized indicates the address book scope authorized for the application, which is not equivalent to the entire enterprise organization. An empty selectedDepartments array does not prove that the organization has no sub-departments. The estimation can be used only when estimateStatus is Ready and a count is returned. A missing count does not indicate 0.
//
// Authorization information: RAM operation agentcore:GetIdentityProviderSyncScope, access level Read, resource ARN: acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource is a required authorization resource.
//
// Idempotence and retry: This operation is read-only and can be queried repeatedly.
//
// Calling this API operation is free of charge. Operation records support ActionTrail auditing.
//
// Sample response:
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
//	"data": {
//
//	  "syncScope": {
//
//	    "mode": "SelectedDepartments",
//
//	    "departmentIds": [
//
//	      "10001"
//
//	    ]
//
//	  },
//
//	  "selectedDepartments": [
//
//	    {
//
//	      "id": "10001",
//
//	      "name": "R&D Department",
//
//	      "exists": true
//
//	    }
//
//	  ],
//
//	  "estimateStatus": "Ready",
//
//	  "expandedDepartmentCount": 2,
//
//	  "estimatedUserCount": 12
//
//	}
//
// }
//
// ```
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
// Queries the details of an image cache.
//
// Description:
//
// Queries the details of the logical image cache in the current workspace. Returns the complete source, current digest, actual size, creation time, status, and secure error messages. AgentCount counts configuration references with the switch enabled. CanDelete verifies both configuration references and instance usage. If the switch is disabled, the count is 0, but deletion is still prohibited if old instances are not released. If the resource does not belong to the current workspace, resources with the same content from other workspaces are not returned.
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
// Queries the details of an image cache.
//
// Description:
//
// Queries the details of the logical image cache in the current workspace. Returns the complete source, current digest, actual size, creation time, status, and secure error messages. AgentCount counts configuration references with the switch enabled. CanDelete verifies both configuration references and instance usage. If the switch is disabled, the count is 0, but deletion is still prohibited if old instances are not released. If the resource does not belong to the current workspace, resources with the same content from other workspaces are not returned.
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
// Queries the details of a specified managed agent.
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
// Queries the details of a specified managed agent.
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
// Queries the details of a specified MCP server, including the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// ## Operation description
//
// Queries the details of a specified MCP server, including the address, type, status, authentication configuration, and protocol.
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
// Queries the details of a specified MCP server, including the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// ## Operation description
//
// Queries the details of a specified MCP server, including the address, type, status, authentication configuration, and protocol.
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
// Queries the details of an MCP marketplace template.
//
// Description:
//
// Returns the current template version and installation form schema.
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
// Returns the current template version and installation form schema.
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
// Queries the public information and current status of a MemoryStore in a specified workspace.
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
// Queries the public information and current status of a MemoryStore in a specified workspace.
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
// Queries the task ID returned by the asynchronous AddMemories operation.
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
// Queries the task ID returned by the asynchronous AddMemories operation.
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
// Queries the details of a specified service endpoint, including target routing, access URLs, authentication configuration, and current status.
//
// Description:
//
// ## Operation description\\nQueries the details of a specified service endpoint. A service endpoint (ServiceEndpoint) provides a stable access URL for a specific agent version (AgentVersion) or workspace collaboration component. The response includes target routing, access URL list, authentication configuration, and current lifecycle status.\\n.
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
// Queries the details of a specified service endpoint, including target routing, access URLs, authentication configuration, and current status.
//
// Description:
//
// ## Operation description\\nQueries the details of a specified service endpoint. A service endpoint (ServiceEndpoint) provides a stable access URL for a specific agent version (AgentVersion) or workspace collaboration component. The response includes target routing, access URL list, authentication configuration, and current lifecycle status.\\n.
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
// Queries the currently active API key and its invocation method for a service endpoint in a specified workspace.
//
// Description:
//
// Queries the currently active API key for a specified service endpoint. The call succeeds only when API key authentication is enabled for the service endpoint and the gateway consumer and credentials are ready. The service reads the API key from the gateway in real time, and AgentCore does not persist the plaintext. Keep the returned API key secure and avoid logging it or exposing it in public configurations.
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
// Queries the currently active API key and its invocation method for a service endpoint in a specified workspace.
//
// Description:
//
// Queries the currently active API key for a specified service endpoint. The call succeeds only when API key authentication is enabled for the service endpoint and the gateway consumer and credentials are ready. The service reads the API key from the gateway in real time, and AgentCore does not persist the plaintext. Keep the returned API key secure and avoid logging it or exposing it in public configurations.
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
// Queries the basic information, version status, labels, visibility scope, and version list of a specified Skill.
//
// Description:
//
// ## Operation description
//
// Queries the basic information, version status, labels, visibility scope, and version list of a specified Skill.
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
// Queries the basic information, version status, labels, visibility scope, and version list of a specified Skill.
//
// Description:
//
// ## Operation description
//
// Queries the basic information, version status, labels, visibility scope, and version list of a specified Skill.
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
// Retrieves the OSS presigned upload URL and object name required to import a Skill ZIP package. After the upload is complete, you can call the Skill OSS upload API to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS presigned upload URL and object name required to import a Skill ZIP package. After the upload is complete, you can call the Skill OSS upload API to complete the import.
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
// Retrieves the OSS presigned upload URL and object name required to import a Skill ZIP package. After the upload is complete, you can call the Skill OSS upload API to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS presigned upload URL and object name required to import a Skill ZIP package. After the upload is complete, you can call the Skill OSS upload API to complete the import.
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
// Queries the complete details of a specified Skill version, including version metadata, Skill content, and associated resources.
//
// Description:
//
// ## Operation description
//
// Queries the complete details of a specified Skill version, including version metadata, Skill content, and associated resources.
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
// Queries the complete details of a specified Skill version, including version metadata, Skill content, and associated resources.
//
// Description:
//
// ## Operation description
//
// Queries the complete details of a specified Skill version, including version metadata, Skill content, and associated resources.
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
// Queries the details of a specified team. The response includes the complete attributes and team roles of each user member and agent member in the team.
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
// Queries the details of a specified team. The response includes the complete attributes and team roles of each user member and agent member in the team.
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
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policies.
//
// Description:
//
// ## Request description\\nQueries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policies.\\n.
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
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policies.
//
// Description:
//
// ## Request description\\nQueries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policies.\\n.
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
// Queries the ACR RAM authorization URL of a workspace.
//
// Description:
//
// ACR and user OSS mounts share the workspace role. If no existing role is found, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is reused and permissions are appended without overwriting existing OSS policies. Fixed historical shared roles continue to be reused and are not recreated or renamed due to new naming templates. The acrInstanceId, namespace, and repository parameters are required. The region and account are derived from the workspace context. Returns the role resolved by the backend and the authorization URL for the target repository. This operation does not create ACR or agent execution versions, nor does it deliver runtime configurations. After authorization is complete, you can call Verify before creating or updating an agent, without the need to call an independent Ensure operation.
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
// Queries the ACR RAM authorization URL of a workspace.
//
// Description:
//
// ACR and user OSS mounts share the workspace role. If no existing role is found, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is reused and permissions are appended without overwriting existing OSS policies. Fixed historical shared roles continue to be reused and are not recreated or renamed due to new naming templates. The acrInstanceId, namespace, and repository parameters are required. The region and account are derived from the workspace context. Returns the role resolved by the backend and the authorization URL for the target repository. This operation does not create ACR or agent execution versions, nor does it deliver runtime configurations. After authorization is complete, you can call Verify before creating or updating an agent, without the need to call an independent Ensure operation.
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
// Queries the RAM authorization URL for mounting AgenticFS in a workspace.
//
// Description:
//
// Determines the account, region, and shared role based on the workspace, and concatenates the resource ARN based on the provided fileSystemId and accessPointId. The server parameter is used to generate the policy name. This operation does not query NAS or check the resource status, and returns an authorization URL that contains only the target NAS policy. After completing RAM authorization, you must call the verification API.
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
// Queries the RAM authorization URL for mounting AgenticFS in a workspace.
//
// Description:
//
// Determines the account, region, and shared role based on the workspace, and concatenates the resource ARN based on the provided fileSystemId and accessPointId. The server parameter is used to generate the policy name. This operation does not query NAS or check the resource status, and returns an authorization URL that contains only the target NAS policy. After completing RAM authorization, you must call the verification API.
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
// Queries the plug-in status of a specified workspace.
//
// Description:
//
// ## Operation description\\nQueries the plug-in status of a specified workspace. Returns whether the plug-in is enabled, its lifecycle status, and the currently effective configuration. Currently, two types of plug-ins are supported: collaboration and agentloop. If a plug-in is not installed, its status is DISABLED.\\n.
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
// Queries the plug-in status of a specified workspace.
//
// Description:
//
// ## Operation description\\nQueries the plug-in status of a specified workspace. Returns whether the plug-in is enabled, its lifecycle status, and the currently effective configuration. Currently, two types of plug-ins are supported: collaboration and agentloop. If a plug-in is not installed, its status is DISABLED.\\n.
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
// Validates the input based on the specified template version and creates an MCP in the workspace.
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
// Validates the input based on the specified template version and creates an MCP in the workspace.
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
// Installs a plugin for a specified AgentCore workspace. Currently, only the collaboration plugin is supported. The installation process runs asynchronously.
//
// Description:
//
// ## Request description\\nInstalls a plugin for a specified AgentCore workspace. Currently, only the `collaboration` plugin is supported. Plugin configurations are passed through the `Config` parameter, and different plugins can define their own configuration structures. The `collaboration` plugin uses `Config.NetworkConfiguration` to specify the VPC and public network access policy. The installation procedure executes asynchronously. If you repeatedly invoke this operation for a plugin with the same name that is being installed or already installed, the operation returns the current status based on idempotent semantics if the configurations are identical. If the configurations are different, the operation returns an operation conflict.\\n
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
// Installs a plugin for a specified AgentCore workspace. Currently, only the collaboration plugin is supported. The installation process runs asynchronously.
//
// Description:
//
// ## Request description\\nInstalls a plugin for a specified AgentCore workspace. Currently, only the `collaboration` plugin is supported. Plugin configurations are passed through the `Config` parameter, and different plugins can define their own configuration structures. The `collaboration` plugin uses `Config.NetworkConfiguration` to specify the VPC and public network access policy. The installation procedure executes asynchronously. If you repeatedly invoke this operation for a plugin with the same name that is being installed or already installed, the operation returns the current status based on idempotent semantics if the configurations are identical. If the configurations are different, the operation returns an operation conflict.\\n
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
// Queries the runtime environments supported for code deployment.
//
// Description:
//
// Queries the directories in the current workspace that can be used to create code package agents or switch runtime environments. Runtime environments are dynamically provided based on the current version configuration, and internal image addresses are not returned.
//
// If maxResults is omitted or set to 0, the default value is 10. If the value is greater than 100, it is processed as 100. Omit nextToken in the first request. For subsequent requests, use the token returned in the previous response as is, and keep the workspace unchanged. If a directory change invalidates the token, query from the first page again. An empty nextToken indicates that there are no more pages. Existing agents continue to use the saved images for restarts, scaling, and regular updates.
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
// Queries the runtime environments supported for code deployment.
//
// Description:
//
// Queries the directories in the current workspace that can be used to create code package agents or switch runtime environments. Runtime environments are dynamically provided based on the current version configuration, and internal image addresses are not returned.
//
// If maxResults is omitted or set to 0, the default value is 10. If the value is greater than 100, it is processed as 100. Omit nextToken in the first request. For subsequent requests, use the token returned in the previous response as is, and keep the workspace unchanged. If a directory change invalidates the token, query from the first page again. An empty nextToken indicates that there are no more pages. Existing agents continue to use the saved images for restarts, scaling, and regular updates.
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
// Queries the list of IM channels for a specified agent.
//
// Description:
//
// Queries the list of IM channels for a specified agent.
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
// Queries the list of IM channels for a specified agent.
//
// Description:
//
// Queries the list of IM channels for a specified agent.
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
// Queries AgentSpec resources in a specified workspace by using paging, and supports name-based search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// ## Operation description
//
// Queries AgentSpec resources in a specified workspace by using paging, and supports name-based search, sorting, and filtering by owner, visibility scope, and business labels.
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
// Queries AgentSpec resources in a specified workspace by using paging, and supports name-based search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// ## Operation description
//
// Queries AgentSpec resources in a specified workspace by using paging, and supports name-based search, sorting, and filtering by owner, visibility scope, and business labels.
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
// Queries the teams to which specified agents belong. Specify the agent IDs by using agentIds to retrieve the membership information of each agent in the teams, including the team ID, team name, and the role of the agent in the team.
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
// Queries the teams to which specified agents belong. Specify the agent IDs by using agentIds to retrieve the membership information of each agent in the teams, including the team ID, team name, and the role of the agent in the team.
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
// Queries the available official models and enterprise models for an enabled QoderCLI Connector. The service retrieves models in real time by using the saved API key and organization ID, and deduplicates the results by model ID.
//
// Description:
//
// Queries the list of available models for a specified Connector. Pagination is supported.
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
// Queries the available official models and enterprise models for an enabled QoderCLI Connector. The service retrieves models in real time by using the saved API key and organization ID, and deduplicates the results by model ID.
//
// Description:
//
// Queries the list of available models for a specified Connector. Pagination is supported.
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
// Queries the list of connectors supported by a specified workspace, including their status, number of bound agents, and configurations. The current version returns the QoderCLI connector.
//
// Description:
//
// Queries the list of connectors in a specified workspace.
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
// Queries the list of connectors supported by a specified workspace, including their status, number of bound agents, and configurations. The current version returns the QoderCLI connector.
//
// Description:
//
// Queries the list of connectors in a specified workspace.
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
// Performs a paged query for credentials in a specified workspace. Filters by type using credentialType, performs fuzzy filtering by credential name using nameLike, configures paging by specifying the maximum number of records per page using maxResults, and retrieves the next page using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Performs a paged query for the credential list in a workspace, and supports filtering by type and name during paging.
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
// Performs a paged query for credentials in a specified workspace. Filters by type using credentialType, performs fuzzy filtering by credential name using nameLike, configures paging by specifying the maximum number of records per page using maxResults, and retrieves the next page using nextToken. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Performs a paged query for the credential list in a workspace, and supports filtering by type and name during paging.
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
// Performs a paged query for paging through the custom agent summary in a specified workspace. For the complete configuration and the latest execute results, call GetCustomAgent.
//
// Description:
//
// The default value of maxResults is 20 and the maximum value is 100. If this parameter is omitted or set to 0, the default value is used.
//
// Omit nextToken in the first request. For subsequent requests, pass the nextToken returned in the previous response as is and keep the same workspace. The nextToken is an opaque token. Do not parse or construct it. An empty token in the response indicates that there are no more pages.
//
// This operation returns the summaries of the current page and does not return the complete deployment configurations or the execution errors of the latest version. To obtain this information, call GetCustomAgent by using the agentId in items.
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
// Performs a paged query for paging through the custom agent summary in a specified workspace. For the complete configuration and the latest execute results, call GetCustomAgent.
//
// Description:
//
// The default value of maxResults is 20 and the maximum value is 100. If this parameter is omitted or set to 0, the default value is used.
//
// Omit nextToken in the first request. For subsequent requests, pass the nextToken returned in the previous response as is and keep the same workspace. The nextToken is an opaque token. Do not parse or construct it. An empty token in the response indicates that there are no more pages.
//
// This operation returns the summaries of the current page and does not return the complete deployment configurations or the execution errors of the latest version. To obtain this information, call GetCustomAgent by using the agentId in items.
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
// Queries the list of external agents in a specified workspace.
//
// Description:
//
// Queries the list of external agents in a specified workspace.
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
// Queries the list of external agents in a specified workspace.
//
// Description:
//
// Queries the list of external agents in a specified workspace.
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
// Queries the departments of an identity provider by page.
//
// Description:
//
// Queries DingTalk or Lark departments by hierarchy and page. Prerequisites include an existing identity provider binding, a user pool, valid application credentials, and address book read authorization. If parentDepartmentId is not specified, the root hierarchy is queried. When expanding a node, use the ID of the node to query its direct child departments. The hasChildren field is used to determine expandability and does not indicate that child nodes have been returned. The nextToken field is only used to continue querying the next page under the current parent department. The query only reads the directory, does not save the synchronization scope, and does not import users. After a new department is added to a DingTalk or Lark organization, re-query the corresponding parent department because the old list does not represent the latest directory.
//
// Authorization information: The RAM operation is agentcore:ListIdentityProviderDepartments, the access level is List, and the resource ARN is acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource requires authorization.
//
// Idempotence and retry: This operation is read-only and can be invoked repeatedly.
//
// It is free of charge to invoke this API operation, and operation records support ActionTrail auditing.
//
// Sample response:
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
//	"items": [
//
//	  {
//
//	    "id": "10001",
//
//	    "name": "R&D Department",
//
//	    "parentId": "1",
//
//	    "memberCount": 12,
//
//	    "hasChildren": true
//
//	  }
//
//	]
//
// }
//
// ```
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
// Queries the departments of an identity provider by page.
//
// Description:
//
// Queries DingTalk or Lark departments by hierarchy and page. Prerequisites include an existing identity provider binding, a user pool, valid application credentials, and address book read authorization. If parentDepartmentId is not specified, the root hierarchy is queried. When expanding a node, use the ID of the node to query its direct child departments. The hasChildren field is used to determine expandability and does not indicate that child nodes have been returned. The nextToken field is only used to continue querying the next page under the current parent department. The query only reads the directory, does not save the synchronization scope, and does not import users. After a new department is added to a DingTalk or Lark organization, re-query the corresponding parent department because the old list does not represent the latest directory.
//
// Authorization information: The RAM operation is agentcore:ListIdentityProviderDepartments, the access level is List, and the resource ARN is acs:agentcore:{regionId}:{accountId}:workspace/{workspaceId}/identityProvider/{identityProviderType}. This resource requires authorization.
//
// Idempotence and retry: This operation is read-only and can be invoked repeatedly.
//
// It is free of charge to invoke this API operation, and operation records support ActionTrail auditing.
//
// Sample response:
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
//	"items": [
//
//	  {
//
//	    "id": "10001",
//
//	    "name": "R&D Department",
//
//	    "parentId": "1",
//
//	    "memberCount": 12,
//
//	    "hasChildren": true
//
//	  }
//
//	]
//
// }
//
// ```
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
// Queries the complete summary list of up to two SCIM credentials in the shared user pool, without returning plaintext tokens or supporting pagination.
//
// Description:
//
// ## Features and prerequisites
//
// Retrieves the current credential summary list from AgentIdentity and returns a data array. Each item contains credentialId, status, createTime, and expireTime. If no credentials exist, an empty array is returned. This operation does not support maxResults, nextToken, or pagination. Do not call this operation with fake pagination fields. The list never returns plaintext tokens, nor does it enable the receiving or issuing of new credentials.
//
// Receiving is not required to be enabled, but the shared pool must be available. A credential is considered available if its status is Enabled and its expireTime is later than the current time. The maximum limit of two credentials is determined by AgentIdentity. Client-side counting is for reference only. The list may contain unavailable or expired credentials. Do not use this list to claim that available tokens exist.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/scim-provisioning/credentials
//
// Call conventions: Call this operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity source resources of the workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business data. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/scim-provisioning/credentials
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": [
//
//	  {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "createTime": "<Upstream creation time>",
//
//	    "expireTime": "<Upstream expiration time>"
//
//	  }
//
//	]
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to the upstream service that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling process.
//
// 400 InvalidParameter: The shared user pool is unavailable. 404 IdentityProviderNotFound: No corresponding binding exists. Upstream query failures cannot be replaced with an empty list, which may mislead users into repeatedly creating credentials.
//
// ## Call relationships and usage boundaries
//
// Used for configuration recovery and credential management: Existing plaintext tokens cannot be queried and retrieved again. The list provides only credential summaries. To obtain a new token, call CreateIdentityProviderSCIMCredential. To revoke a credential, submit the credentialId from the list to DeleteIdentityProviderSCIMCredential. A maximum of two credentials are allowed in the same pool. This operation returns the complete array and does not support or fabricate pagination parameters.
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
// Queries the complete summary list of up to two SCIM credentials in the shared user pool, without returning plaintext tokens or supporting pagination.
//
// Description:
//
// ## Features and prerequisites
//
// Retrieves the current credential summary list from AgentIdentity and returns a data array. Each item contains credentialId, status, createTime, and expireTime. If no credentials exist, an empty array is returned. This operation does not support maxResults, nextToken, or pagination. Do not call this operation with fake pagination fields. The list never returns plaintext tokens, nor does it enable the receiving or issuing of new credentials.
//
// Receiving is not required to be enabled, but the shared pool must be available. A credential is considered available if its status is Enabled and its expireTime is later than the current time. The maximum limit of two credentials is determined by AgentIdentity. Client-side counting is for reference only. The list may contain unavailable or expired credentials. Do not use this list to claim that available tokens exist.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/scim-provisioning/credentials
//
// Call conventions: Call this operation in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity source resources of the workspace. The ws-example, URL, certificate, and token in the following examples are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business data. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/scim-provisioning/credentials
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": [
//
//	  {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "createTime": "<Upstream creation time>",
//
//	    "expireTime": "<Upstream expiration time>"
//
//	  }
//
//	]
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to the upstream service that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling process.
//
// 400 InvalidParameter: The shared user pool is unavailable. 404 IdentityProviderNotFound: No corresponding binding exists. Upstream query failures cannot be replaced with an empty list, which may mislead users into repeatedly creating credentials.
//
// ## Call relationships and usage boundaries
//
// Used for configuration recovery and credential management: Existing plaintext tokens cannot be queried and retrieved again. The list provides only credential summaries. To obtain a new token, call CreateIdentityProviderSCIMCredential. To revoke a credential, submit the credentialId from the list to DeleteIdentityProviderSCIMCredential. A maximum of two credentials are allowed in the same pool. This operation returns the complete array and does not support or fabricate pagination parameters.
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
// ## Features and prerequisites
//
// Returns the local binding summary without querying the actual SAML and SCIM configurations of AgentIdentity. If no bindings exist, items is an empty array. A workspace currently supports only one external binding. The default value of maxResults is 10, and the valid values are 1 to 100. Currently, nextToken must be omitted or left empty. A non-empty value returns InvalidParameter, and multi-page results cannot be fabricated.
//
// The status of CustomSSO indicates the task status of shared resources and logon configurations. You must combine protocol and saml to determine whether SAML is configured. SCIM queries are independently performed by using GetIdentityProviderSCIM. The list does not issue tokens and does not return the vendor appSecret, encryptKey, or verificationToken.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/identity-providers
//
// Call conventions: Call the API in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request example shows the body parameters of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/identity-providers?maxResults=10
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"items": [
//
//	  {
//
//	    "workspaceId": "ws-example",
//
//	    "identityProviderType": "CustomSSO",
//
//	    "status": "READY",
//
//	    "loginEnabled": false,
//
//	    "syncEnabled": false
//
//	  }
//
//	],
//
//	"maxResults": 10,
//
//	"totalCount": 1
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call failure to the upstream may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 InvalidParameter: maxResults is out of bounds or nextToken is non-empty, which is currently not supported. When a query fails, an error is displayed instead of being cleared as unconfigured.
//
// ## Call relationships and usage boundaries
//
// This API is used to discover current bindings. The specific SAML configuration of CustomSSO is queried by using GetIdentityProvider. The SP address is queried by using GetIdentityProviderServiceProvider. The actual SCIM configuration and independent operation results are queried by using GetIdentityProviderSCIM.
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
// ## Features and prerequisites
//
// Returns the local binding summary without querying the actual SAML and SCIM configurations of AgentIdentity. If no bindings exist, items is an empty array. A workspace currently supports only one external binding. The default value of maxResults is 10, and the valid values are 1 to 100. Currently, nextToken must be omitted or left empty. A non-empty value returns InvalidParameter, and multi-page results cannot be fabricated.
//
// The status of CustomSSO indicates the task status of shared resources and logon configurations. You must combine protocol and saml to determine whether SAML is configured. SCIM queries are independently performed by using GetIdentityProviderSCIM. The list does not issue tokens and does not return the vendor appSecret, encryptKey, or verificationToken.
//
// ## Requests and examples
//
// GET /workspaces/{workspaceId}/identity-providers
//
// Call conventions: Call the API in the region of the target workspace by using an AccessKey pair or STS token that has permissions on the identity provider resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request example shows the body parameters of the public API, not the extra body wrapper of the backend direct connection API. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// GET /workspaces/ws-example/identity-providers?maxResults=10
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"items": [
//
//	  {
//
//	    "workspaceId": "ws-example",
//
//	    "identityProviderType": "CustomSSO",
//
//	    "status": "READY",
//
//	    "loginEnabled": false,
//
//	    "syncEnabled": false
//
//	  }
//
//	],
//
//	"maxResults": 10,
//
//	"totalCount": 1
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call failure to the upstream may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 InvalidParameter: maxResults is out of bounds or nextToken is non-empty, which is currently not supported. When a query fails, an error is displayed instead of being cleared as unconfigured.
//
// ## Call relationships and usage boundaries
//
// This API is used to discover current bindings. The specific SAML configuration of CustomSSO is queried by using GetIdentityProvider. The SP address is queried by using GetIdentityProviderServiceProvider. The actual SCIM configuration and independent operation results are queried by using GetIdentityProviderSCIM.
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
// Queries image caches.
//
// Description:
//
// Queries the logical image caches within the current account, region, and workspace. WorkspaceId is required, and caches are not visible across workspaces. All conditions filter the full set using the AND operator before pagination is calculated. Results are sorted in descending order of logical resource creation time and ascending order of ImageId. The cursor is bound to the account, region, workspace, all filters, and the sort order. Returns the complete cache objects, the AgentCount based on the configuration switch, and the actual deletion eligibility. Query failures do not return a disguised empty list, and the nextToken on the last page is an empty string.
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
// Queries image caches.
//
// Description:
//
// Queries the logical image caches within the current account, region, and workspace. WorkspaceId is required, and caches are not visible across workspaces. All conditions filter the full set using the AND operator before pagination is calculated. Results are sorted in descending order of logical resource creation time and ascending order of ImageId. The cursor is bound to the account, region, workspace, all filters, and the sort order. Returns the complete cache objects, the AgentCount based on the configuration switch, and the actual deletion eligibility. Query failures do not return a disguised empty list, and the nextToken on the last page is an empty string.
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
// Queries the list of managed agents in a specified workspace.
//
// Description:
//
// Queries the list of managed agents in a specified workspace by using paging. Returns summary information for each agent, including the identity, name, status, template, and specifications.
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
// Queries the list of managed agents in a specified workspace.
//
// Description:
//
// Queries the list of managed agents in a specified workspace by using paging. Returns summary information for each agent, including the identity, name, status, template, and specifications.
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
// Queries all online official MCP templates. You can filter results by keyword, usage tag, and MCP type.
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
// Queries all online official MCP templates. You can filter results by keyword, usage tag, and MCP type.
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
// Queries the list of tools exposed by a specified MCP service and their input/output schemas.
//
// Description:
//
// ## Operation description
//
// Queries the list of tools exposed by a specified MCP service and their input/output schemas.
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
// Queries the list of tools exposed by a specified MCP service and their input/output schemas.
//
// Description:
//
// ## Operation description
//
// Queries the list of tools exposed by a specified MCP service and their input/output schemas.
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
// Queries MCP services and their configurations and statuses in a specified workspace by using paging.
//
// Description:
//
// ## Operation description
//
// Queries MCP services and their configurations and statuses in a specified workspace by using paging.
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
// Queries MCP services and their configurations and statuses in a specified workspace by using paging.
//
// Description:
//
// ## Operation description
//
// Queries MCP services and their configurations and statuses in a specified workspace by using paging.
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
// Performs a paged query on long-term memory records. The results can be filtered independently by user ID, agent ID, and session ID. Paging is supported.
//
// Description:
//
// Performs a paged query on long-term memory records. The results can be filtered independently by user ID, agent ID, and session ID. Paging is supported.
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
// Performs a paged query on long-term memory records. The results can be filtered independently by user ID, agent ID, and session ID. Paging is supported.
//
// Description:
//
// Performs a paged query on long-term memory records. The results can be filtered independently by user ID, agent ID, and session ID. Paging is supported.
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
// Performs a paging query for raw messages by exact scope. SessionId is required, and at least one of UserId and AgentId must be provided.
//
// Description:
//
// Performs a paging query for raw messages by exact scope. SessionId is required, and at least one of UserId and AgentId must be provided.
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
// Performs a paging query for raw messages by exact scope. SessionId is required, and at least one of UserId and AgentId must be provided.
//
// Description:
//
// Performs a paging query for raw messages by exact scope. SessionId is required, and at least one of UserId and AgentId must be provided.
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
// Performs a paging query for sessions in a MemoryStore. You can filter the sessions independently by user ID and agent ID.
//
// Description:
//
// Performs a paging query for sessions in a MemoryStore. You can filter the sessions independently by user ID and agent ID.
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
// Performs a paging query for sessions in a MemoryStore. You can filter the sessions independently by user ID and agent ID.
//
// Description:
//
// Performs a paging query for sessions in a MemoryStore. You can filter the sessions independently by user ID and agent ID.
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
// Performs a paged query for MemoryStores in a specified workspace with paging.
//
// Description:
//
// Performs a paged query for MemoryStores in a specified workspace with paging.
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
// Performs a paged query for MemoryStores in a specified workspace with paging.
//
// Description:
//
// Performs a paged query for MemoryStores in a specified workspace with paging.
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
// Queries model connections in a specified workspace with paging. Supports filtering by name, provider type, and invoke protocol.
//
// Description:
//
// Queries model connections in a specified AgentCore workspace. Supports filtering by name through `Name` and selecting exact match or fuzzy match through `SearchType`. Also supports filtering by model provider type and invoke protocol, and supports paging.
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
// Queries model connections in a specified workspace with paging. Supports filtering by name, provider type, and invoke protocol.
//
// Description:
//
// Queries model connections in a specified AgentCore workspace. Supports filtering by name through `Name` and selecting exact match or fuzzy match through `SearchType`. Also supports filtering by model provider type and invoke protocol, and supports paging.
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
// Queries models in a specified workspace by using paging. Supports filtering by model connection and model name.
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
// Queries models in a specified workspace by using paging. Supports filtering by model connection and model name.
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
// Queries the models and their capabilities in the built-in model catalog of AgentCore for a specified provider.
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
// Queries the models and their capabilities in the built-in model catalog of AgentCore for a specified provider.
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
// Queries the list of active sessions in the sandbox of a specified managed agent by paging, and returns the session source and external channel type.
//
// Description:
//
// Queries the list of active sessions in the sandbox of a specified managed agent.
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
// Queries the list of active sessions in the sandbox of a specified managed agent by paging, and returns the session source and external channel type.
//
// Description:
//
// Queries the list of active sessions in the sandbox of a specified managed agent.
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
// Queries the list of sandboxes for a managed agent in a specified workspace. You can filter results by sandbox ID and active session ID fragments.
//
// Description:
//
// Queries the list of sandboxes for a specified managed agent. The searchText parameter performs a fuzzy match on sandbox ID fragments, and the sessionId parameter performs a fuzzy match on active session ID fragments. Both parameters are case-insensitive and can be used together with AND logic.
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
// Queries the list of sandboxes for a managed agent in a specified workspace. You can filter results by sandbox ID and active session ID fragments.
//
// Description:
//
// Queries the list of sandboxes for a specified managed agent. The searchText parameter performs a fuzzy match on sandbox ID fragments, and the sessionId parameter performs a fuzzy match on active session ID fragments. Both parameters are case-insensitive and can be used together with AND logic.
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
// Queries service endpoints in a specified workspace with paging and supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// ## Request description\\nQueries service endpoints in a specified workspace with paging. You can filter the results by targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status. Use maxResults to specify the maximum number of records per page, and use nextToken to retrieve the next page. If maxResults is not specified, the server returns 20 records by default.\\n
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
// Queries service endpoints in a specified workspace with paging and supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// ## Request description\\nQueries service endpoints in a specified workspace with paging. You can filter the results by targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status. Use maxResults to specify the maximum number of records per page, and use nextToken to retrieve the next page. If maxResults is not specified, the server returns 20 records by default.\\n
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
// Queries skill reference relationships.
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
// Queries skill reference relationships.
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
// Performs a paged query of Skills in a specified workspace and returns basic Skill information, version status, and paging information.
//
// Description:
//
// ## Operation description
//
// Performs a paged query of Skills in a specified workspace and returns basic Skill information, version status, and paging information.
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
// Performs a paged query of Skills in a specified workspace and returns basic Skill information, version status, and paging information.
//
// Description:
//
// ## Operation description
//
// Performs a paged query of Skills in a specified workspace and returns basic Skill information, version status, and paging information.
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
// Queries teams in a specified workspace by using paging. You can use nameLike to filter teams by fuzzy matching the team name, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. The member information in the list includes only the member identifier, name, and team role.
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
// Queries teams in a specified workspace by using paging. You can use nameLike to filter teams by fuzzy matching the team name, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. The member information in the list includes only the member identifier, name, and team role.
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
// Performs a paged query for users in a specified workspace. You can use nameLike to filter users by fuzzy match on the username, maxResults to specify the maximum number of records per page for paging, and nextToken to retrieve the next page. If maxResults is not specified, 10 records are returned by default.
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
// Performs a paged query for users in a specified workspace. You can use nameLike to filter users by fuzzy match on the username, maxResults to specify the maximum number of records per page for paging, and nextToken to retrieve the next page. If maxResults is not specified, 10 records are returned by default.
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
// Performs a paged query on the workspaces under the current tenant. By default, the list does not return soft-deleted records with the Deleted status. The server performs stable sorting on the results by creation order to facilitate paging.
//
// Description:
//
// ## Request description\\nPerforms a paged query on the workspaces under the current tenant. By default, the list does not return soft-deleted records with the Deleted status. The server performs stable sorting on the results by creation order to facilitate paging. You can use nextToken to obtain the next page, skip to skip a specified number of workspaces, maxResults to specify the maximum number of records per page, and nameLike to perform fuzzy filtering by workspace name. If maxResults is not specified or is set to 0, the server returns 20 records by default.\\n
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
// Performs a paged query on the workspaces under the current tenant. By default, the list does not return soft-deleted records with the Deleted status. The server performs stable sorting on the results by creation order to facilitate paging.
//
// Description:
//
// ## Request description\\nPerforms a paged query on the workspaces under the current tenant. By default, the list does not return soft-deleted records with the Deleted status. The server performs stable sorting on the results by creation order to facilitate paging. You can use nextToken to obtain the next page, skip to skip a specified number of workspaces, maxResults to specify the maximum number of records per page, and nameLike to perform fuzzy filtering by workspace name. If maxResults is not specified or is set to 0, the server returns 20 records by default.\\n
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
// Takes the online version of a specified Skill offline so that it is no longer used as the online version.
//
// Description:
//
// ## Operation description
//
// Takes the online version of a specified Skill offline so that it is no longer used as the online version.
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
// Takes the online version of a specified Skill offline so that it is no longer used as the online version.
//
// Description:
//
// ## Operation description
//
// Takes the online version of a specified Skill offline so that it is no longer used as the online version.
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
// Sets a specified Skill version as the online version.
//
// Description:
//
// ## Operation description
//
// Sets a specified Skill version as the online version.
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
// Sets a specified Skill version as the online version.
//
// Description:
//
// ## Operation description
//
// Sets a specified Skill version as the online version.
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
// Parses and checks one or more Skill ZIP packages uploaded to OSS, and returns the name, version, and conflict check results.
//
// Description:
//
// ## Operation description
//
// Parses and checks one or more Skill ZIP packages uploaded to OSS, and returns the name, version, and conflict check results.
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
// Parses and checks one or more Skill ZIP packages uploaded to OSS, and returns the name, version, and conflict check results.
//
// Description:
//
// ## Operation description
//
// Parses and checks one or more Skill ZIP packages uploaded to OSS, and returns the name, version, and conflict check results.
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
// Prepares the SCIM receiving configuration, and returns the provisioning URL and credential summary. Issues a token if no available credentials exist and returns the plaintext only in the current response. Does not automatically enable user synchronization.
//
// Description:
//
// ## Features and prerequisites
//
// A CustomSSO shared user pool must be available. This operation does not depend on SAML metadata and is not blocked by SAML updating or update failure states. It synchronously calls AgentIdentity to enable SCIM receiving, and reads the configuration and credential list. If a credential in the Enabled state that has not expired already exists, its summary is returned and the token is set to null. If no available credentials exist, a new credential is created and the token is returned only in the current response. Existing invalid or expired credentials are not automatically deleted and may still trigger the upstream limit of up to two credentials.
//
// The prepare operation has the side effects of enabling receiving and potentially issuing credentials, but it does not set the AC syncEnabled parameter, initiate the first import, or indicate that the external provisioning connection test has passed. After configuring the baseUrl and the current token in the external provisioning system, call UpdateIdentityProviderSCIM to confirm and enable synchronization. Existing tokens cannot be retrieved again. If the response is lost, query the credential list first. Tokens are not automatically reissued. Tokens must not be logged or persisted in frontend caches.
//
// ## Requests and examples
//
// POST /workspaces/{workspaceId}/scim-provisioning/prepare
//
// Calling conventions: Call this operation in the region of the target workspace by using an AccessKey ID/STS identity that has permissions on the identity source resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not log secrets. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/scim-provisioning/prepare
//
// No business request body.
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "credential": {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "expireTime": "<Expiration time returned by the upstream service>"
//
//	  },
//
//	  "token": "<Token issued only in the current response; null if an available credential already exists>"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to the upstream service that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace not found errors follow the common error handling process.
//
// 400 LimitExceeded.ScimCredential: The maximum limit of two credentials has been reached. You must explicitly delete one credential and try again. Credentials are not automatically deleted, rotated, or expanded. 400 InvalidParameter: The shared pool is unavailable. This call may have enabled upstream receiving and then reported an error due to issuance failure. The error does not mean that the operation has been automatically rolled back.
//
// ## Call relationships and usage boundaries
//
// The shared pool is created by using CreateIdentityProvider with identityProviderType set to CustomSSO. SAML configuration is not required. If an existing credential is reused and token is null, use the token that the user has already saved. Do not configure a null value in the external system. If the token is lost, call ListIdentityProviderSCIMCredentials to check first, and then explicitly call CreateIdentityProviderSCIMCredential to issue a new token. When the limit of two credentials is reached, the user must confirm and delete the unused credential.
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
// Prepares the SCIM receiving configuration, and returns the provisioning URL and credential summary. Issues a token if no available credentials exist and returns the plaintext only in the current response. Does not automatically enable user synchronization.
//
// Description:
//
// ## Features and prerequisites
//
// A CustomSSO shared user pool must be available. This operation does not depend on SAML metadata and is not blocked by SAML updating or update failure states. It synchronously calls AgentIdentity to enable SCIM receiving, and reads the configuration and credential list. If a credential in the Enabled state that has not expired already exists, its summary is returned and the token is set to null. If no available credentials exist, a new credential is created and the token is returned only in the current response. Existing invalid or expired credentials are not automatically deleted and may still trigger the upstream limit of up to two credentials.
//
// The prepare operation has the side effects of enabling receiving and potentially issuing credentials, but it does not set the AC syncEnabled parameter, initiate the first import, or indicate that the external provisioning connection test has passed. After configuring the baseUrl and the current token in the external provisioning system, call UpdateIdentityProviderSCIM to confirm and enable synchronization. Existing tokens cannot be retrieved again. If the response is lost, query the credential list first. Tokens are not automatically reissued. Tokens must not be logged or persisted in frontend caches.
//
// ## Requests and examples
//
// POST /workspaces/{workspaceId}/scim-provisioning/prepare
//
// Calling conventions: Call this operation in the region of the target workspace by using an AccessKey ID/STS identity that has permissions on the identity source resources of the workspace. In the following examples, ws-example, URLs, certificates, and tokens are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API operation, not the extra body wrapper of the backend direct connection API operation. The response example shows only the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not log secrets. Asynchronous business tasks do not mean that the HTTP request uses a persistent connection to wait for completion.
//
// ```text
//
// POST /workspaces/ws-example/scim-provisioning/prepare
//
// No business request body.
//
// ```
//
// ## Response examples
//
// ```json
//
// {
//
//	"data": {
//
//	  "configuration": {
//
//	    "status": "Enabled",
//
//	    "baseUrl": "https://example.com/scim/v2"
//
//	  },
//
//	  "credential": {
//
//	    "credentialId": "credential-example",
//
//	    "status": "Enabled",
//
//	    "expireTime": "<Expiration time returned by the upstream service>"
//
//	  },
//
//	  "token": "<Token issued only in the current response; null if an available credential already exists>"
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. Synchronous calls to the upstream service that fail may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret read failures as unconfigured. Actual authentication and workspace not found errors follow the common error handling process.
//
// 400 LimitExceeded.ScimCredential: The maximum limit of two credentials has been reached. You must explicitly delete one credential and try again. Credentials are not automatically deleted, rotated, or expanded. 400 InvalidParameter: The shared pool is unavailable. This call may have enabled upstream receiving and then reported an error due to issuance failure. The error does not mean that the operation has been automatically rolled back.
//
// ## Call relationships and usage boundaries
//
// The shared pool is created by using CreateIdentityProvider with identityProviderType set to CustomSSO. SAML configuration is not required. If an existing credential is reused and token is null, use the token that the user has already saved. Do not configure a null value in the external system. If the token is lost, call ListIdentityProviderSCIMCredentials to check first, and then explicitly call CreateIdentityProviderSCIMCredential to issue a new token. When the limit of two credentials is reached, the user must confirm and delete the unused credential.
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
// Publishes a specified Skill version to transition it to the published state.
//
// Description:
//
// ## Operation description
//
// Publishes a specified Skill version to transition it to the published state.
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
// Publishes a specified Skill version to transition it to the published state.
//
// Description:
//
// ## Operation description
//
// Publishes a specified Skill version to transition it to the published state.
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
// Converts a specified Skill version back to an editable draft so that you can continue to modify it.
//
// Description:
//
// ## Request description
//
// Converts a specified Skill version back to an editable draft so that you can continue to modify it.
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
// Converts a specified Skill version back to an editable draft so that you can continue to modify it.
//
// Description:
//
// ## Request description
//
// Converts a specified Skill version back to an editable draft so that you can continue to modify it.
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
// Resets the logon password of a specified user. You can specify the user by using agentCoreUserId or username, and you must specify at least one of them. This operation only supports resetting passwords for users who use local password authentication in the workspace. If you do not specify the password parameter, the server generates a random password and returns it in the response.
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
// Resets the logon password of a specified user. You can specify the user by using agentCoreUserId or username, and you must specify at least one of them. This operation only supports resetting passwords for users who use local password authentication in the workspace. If you do not specify the password parameter, the server generates a random password and returns it in the response.
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
// Retries the image cache.
//
// Description:
//
// Retries the image cache for the original ImageId in the current workspace. WorkspaceId is required. The business request body is not required and can be omitted or provided as an empty object. The source authorization and transfer network reuse the trusted context of the original registration. If a digest exists, the retry uses the original digest. If no digest exists, the original tag is resolved again. A new request for a FAILED status can start a new production round. PENDING and READY statuses are reused. DELETING status is rejected, and deleted resources are not recreated. PUBLIC and PRIVATE statuses reuse the pull authorization of the original workspace. Returns a complete acceptance snapshot and retains the original ImageId and CreatedAt. ClientToken does not guarantee request deduplication. If the result is not confirmed, query the real-time status first before deciding whether to resubmit.
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
// Retries the image cache.
//
// Description:
//
// Retries the image cache for the original ImageId in the current workspace. WorkspaceId is required. The business request body is not required and can be omitted or provided as an empty object. The source authorization and transfer network reuse the trusted context of the original registration. If a digest exists, the retry uses the original digest. If no digest exists, the original tag is resolved again. A new request for a FAILED status can start a new production round. PENDING and READY statuses are reused. DELETING status is rejected, and deleted resources are not recreated. PUBLIC and PRIVATE statuses reuse the pull authorization of the original workspace. Returns a complete acceptance snapshot and retains the original ImageId and CreatedAt. ClientToken does not guarantee request deduplication. If the result is not confirmed, query the real-time status first before deciding whether to resubmit.
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
// Retrieves long-term memory from a MemoryStore based on semantics.
//
// Description:
//
// Retrieves long-term memory from a MemoryStore based on semantics.
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
// Retrieves long-term memory from a MemoryStore based on semantics.
//
// Description:
//
// Retrieves long-term memory from a MemoryStore based on semantics.
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
// ## Request description
//
// Submits a specified draft version of an AgentSpec for review. If no release pipeline is configured, the version is directly published online.
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
// ## Request description
//
// Submits a specified draft version of an AgentSpec for review. If no release pipeline is configured, the version is directly published online.
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
// ## Operation description
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
// ## Operation description
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
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel. At least one updatable field must be provided.
//
// Description:
//
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel. At least one updatable field must be provided.
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
// Updates the ServiceEndpoint, enablement status, or channel behavior configuration bound to an IM channel. At least one updatable field must be provided.
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
// Fully updates the channel credentials of a specified IM channel. Secrets are not returned in the response.
//
// Description:
//
// Fully updates the channel credentials of a specified IM channel. Secrets are not returned in the response.
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
// Fully updates the channel credentials of a specified IM channel. Secrets are not returned in the response.
//
// Description:
//
// Fully updates the channel credentials of a specified IM channel. Secrets are not returned in the response.
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
// Updates the business tags, version labels, and visibility scope of a specified AgentSpec. Fields that are not provided remain unchanged.
//
// Description:
//
// ## Operation description
//
// Updates the business tags, version labels, and visibility scope of a specified AgentSpec. Fields that are not provided remain unchanged.
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
// Updates the business tags, version labels, and visibility scope of a specified AgentSpec. Fields that are not provided remain unchanged.
//
// Description:
//
// ## Operation description
//
// Updates the business tags, version labels, and visibility scope of a specified AgentSpec. Fields that are not provided remain unchanged.
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
// Updates the API Key or Service Account Key collection for an enabled QoderCLI Connector. When serviceAccountKeys is submitted, the service treats it as the complete updated key collection and performs additions, renames, rotations, and deletions accordingly.
//
// Description:
//
// Updates the sensitive configuration of a specified Connector and aligns Service Account Keys by ID.
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
// Updates the API Key or Service Account Key collection for an enabled QoderCLI Connector. When serviceAccountKeys is submitted, the service treats it as the complete updated key collection and performs additions, renames, rotations, and deletions accordingly.
//
// Description:
//
// Updates the sensitive configuration of a specified Connector and aligns Service Account Keys by ID.
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
// Updates the content or description of a specified credential. You must specify at least one of credentialMetadata and description. Unspecified attributes remain unchanged. The credential name and credential type cannot be modified after creation.
//
// Description:
//
// Updates the metadata or application scope of a specified credential.
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
// Updates the content or description of a specified credential. You must specify at least one of credentialMetadata and description. Unspecified attributes remain unchanged. The credential name and credential type cannot be modified after creation.
//
// Description:
//
// Updates the metadata or application scope of a specified credential.
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
// Updates the configuration of a specified custom agent in a workspace and executes the update asynchronously. Omitted top-level fields retain their original configurations. You can call the GetCustomAgent operation to query the update result.
//
// Description:
//
// Provide at least one top-level field to update. Omitted top-level fields retain their original configurations. When providing a configuration object, specify its required subfields and include any optional configurations that you want to retain. Arrays are replaced as a whole. Pass an empty array [] to clear a list. The agenticFsMounts field does not accept null. Pass null for the description field to clear the description, and pass null for the healthCheck field to remove the health check.
//
// When updating an image, you must specify artifact.port and the type, instanceId, namespace, repo, and tag fields of artifact.container. To retain a custom startup command, also provide the entrypoint field.
//
// In code package pattern, specify artifact.port and the sourceUri, languageRuntime, and entrypoint fields of artifact.code. Do not specify the container field at the same time. Dependencies must be packaged by the user. ZIP and TAR.GZ formats are supported, and the compressed package size cannot exceed 500,000,000 bytes. Query the runtime environment identity by calling the ListAgentCodeRuntimes operation. The sourceUri field supports oss:// addresses or HTTPS object addresses of standard OSS public domain names, including signed URLs. For local files, first call the GetAgentCodePackageUploadUrl operation to obtain an upload URL, upload the file directly, and then commit the returned sourceUri.
//
// When submitting a new code package source, the backend rechecks read permissions, objects, and sizes. The configuration is saved only after the transfer succeeds. If the transfer fails, the current configuration is not saved. Saved fixed packages can be reused. Set the client timeout to 210 seconds for code package requests. Archive integrity and secure decompression are checked at startup. A successful acceptance does not indicate that the agent is ready to run.
//
// When providing the runtime field, the compute and sessionPolicy fields must be provided at the same time. If the hpa field or its subfields are omitted, unspecified elastic configurations use default values. To retain the original elastic configurations, provide them together. When providing the network field, omitted access items are treated as allowed, but actual access is still restricted by the workspace network configurations. To retain the original restrictions, provide both accessInternet and accessVpc. When providing the healthCheck field, the port field is required, and other unspecified items use default values. When providing the logging.sls field, both project and logstore must be specified.
//
// The update is an asynchronous process. A successful response indicates that the request is accepted. Call the GetCustomAgent operation to query latestSpecVersion, latestVersionStatus, and latestExecutionError, and check the status field to confirm the update result.
//
// The clientToken field currently does not provide idempotency guarantees. Do not rely on this parameter to prevent duplicate requests. The healthyThreshold field for health checks currently does not take effect. Omit this field.
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
// Updates the configuration of a specified custom agent in a workspace and executes the update asynchronously. Omitted top-level fields retain their original configurations. You can call the GetCustomAgent operation to query the update result.
//
// Description:
//
// Provide at least one top-level field to update. Omitted top-level fields retain their original configurations. When providing a configuration object, specify its required subfields and include any optional configurations that you want to retain. Arrays are replaced as a whole. Pass an empty array [] to clear a list. The agenticFsMounts field does not accept null. Pass null for the description field to clear the description, and pass null for the healthCheck field to remove the health check.
//
// When updating an image, you must specify artifact.port and the type, instanceId, namespace, repo, and tag fields of artifact.container. To retain a custom startup command, also provide the entrypoint field.
//
// In code package pattern, specify artifact.port and the sourceUri, languageRuntime, and entrypoint fields of artifact.code. Do not specify the container field at the same time. Dependencies must be packaged by the user. ZIP and TAR.GZ formats are supported, and the compressed package size cannot exceed 500,000,000 bytes. Query the runtime environment identity by calling the ListAgentCodeRuntimes operation. The sourceUri field supports oss:// addresses or HTTPS object addresses of standard OSS public domain names, including signed URLs. For local files, first call the GetAgentCodePackageUploadUrl operation to obtain an upload URL, upload the file directly, and then commit the returned sourceUri.
//
// When submitting a new code package source, the backend rechecks read permissions, objects, and sizes. The configuration is saved only after the transfer succeeds. If the transfer fails, the current configuration is not saved. Saved fixed packages can be reused. Set the client timeout to 210 seconds for code package requests. Archive integrity and secure decompression are checked at startup. A successful acceptance does not indicate that the agent is ready to run.
//
// When providing the runtime field, the compute and sessionPolicy fields must be provided at the same time. If the hpa field or its subfields are omitted, unspecified elastic configurations use default values. To retain the original elastic configurations, provide them together. When providing the network field, omitted access items are treated as allowed, but actual access is still restricted by the workspace network configurations. To retain the original restrictions, provide both accessInternet and accessVpc. When providing the healthCheck field, the port field is required, and other unspecified items use default values. When providing the logging.sls field, both project and logstore must be specified.
//
// The update is an asynchronous process. A successful response indicates that the request is accepted. Call the GetCustomAgent operation to query latestSpecVersion, latestVersionStatus, and latestExecutionError, and check the status field to confirm the update result.
//
// The clientToken field currently does not provide idempotency guarantees. Do not rely on this parameter to prevent duplicate requests. The healthyThreshold field for health checks currently does not take effect. Omit this field.
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
// Updates the configuration of a specified external agent.
//
// Description:
//
// Updates the configuration of a specified external agent.
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
// Updates the configuration of a specified external agent.
//
// Description:
//
// Updates the configuration of a specified external agent.
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
// Modifies the logon configuration of a bound identity provider in a specified workspace. DingTalk and Lark support logon switches, incremental synchronization switches, application configurations, and synchronization scopes. CustomSSO supports SAML configurations and logon switches. Actual updates take effect asynchronously. Call GetIdentityProvider to query the status.
//
// Description:
//
// ### Prerequisites and update methods
//
// The workspace must be bound to an identity provider of the corresponding type. For DingTalk and Lark, you must submit at least one of loginEnabled, syncEnabled, metadata, and syncScope. protocol and saml are not accepted. For CustomSSO, use protocol, saml, and loginEnabled. metadata, syncEnabled, and syncScope are not accepted. The SCIM synchronization switch is managed separately through UpdateIdentityProviderSCIM.
//
// ### DingTalk and Lark
//
// The metadata is merged by field. Submit only the fields that need to be updated. Omitted fields and existing keys remain unchanged. Explicitly submitting null, empty strings, or blank strings will be rejected. The server validates the complete configuration after merging. When incremental synchronization is enabled, the merged configuration must include encryptKey and verificationToken.
//
// syncEnabled controls subsequent address book change events and does not control the initial synchronization or manual synchronization. syncScope applies only to DingTalk and Lark. If it is omitted or set to null, the current scope is retained. Saving the scope does not automatically trigger a full synchronization. To perform manual synchronization, call CreateIdentityProviderSyncJob after saving and waiting for the identity provider to be ready, and then call GetIdentityProviderSyncJob to query the task result.
//
// ### Custom SAML identity provider
//
// CustomSSO currently supports only SAML. The saml parameter supports two mutually exclusive methods: providing metadataXml, or providing the four manual configuration fields: entityId, loginUrl, bindingType, and x509Certificates.
//
// metadataXml is a single IdP EntityDescriptor XML text of up to 1 MiB, not a URL or file path. DTDs and external entities are prohibited. For manual configuration, entityId must be non-empty and up to 2,048 characters, and cannot be changed after saving. loginUrl must be an absolute HTTPS URL without user credentials and fragments. bindingType must be HTTP-Redirect or HTTP-POST. Certificates must be valid X.509 signing certificates, with a maximum of 10 certificates.
//
// jitProvisionEnabled controls the automatic creation of users during the initial logon, and jitUpdateEnabled controls whether user information is updated during logon. Omitting these switches during an update retains the existing values. Modifying only the JIT switches also requires an existing valid SAML configuration. Saving the SAML configuration does not automatically enable logon. A valid configuration must exist before enabling logon. Setting loginEnabled to false only triggers a shutdown of logon and does not delete the saved SAML or SCIM configurations.
//
// ### Asynchronous results and retries
//
// When actual changes exist, the expected configuration and the UPDATING status of the identity provider are returned. The CONFIGURED status may also be returned when resource configurations are restored. Call GetIdentityProvider to wait for the READY status. If the operation fails, check for statuses such as UPDATE_FAILED. If the configuration is identical and no restoration is required, the current binding may be returned directly without generating a new task. A successful HTTP request does not indicate that the background task or upstream configuration has taken effect.
//
// Operation conflicts occur when in-flight operations exist. If a submission times out, query the configuration and status before deciding whether to retry. The status parameter describes the status of the identity provider resource and logon configuration. The SCIM operation results for CustomSSO must be queried through GetIdentityProviderSCIM and cannot be inferred from this operation.
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
// Modifies the logon configuration of a bound identity provider in a specified workspace. DingTalk and Lark support logon switches, incremental synchronization switches, application configurations, and synchronization scopes. CustomSSO supports SAML configurations and logon switches. Actual updates take effect asynchronously. Call GetIdentityProvider to query the status.
//
// Description:
//
// ### Prerequisites and update methods
//
// The workspace must be bound to an identity provider of the corresponding type. For DingTalk and Lark, you must submit at least one of loginEnabled, syncEnabled, metadata, and syncScope. protocol and saml are not accepted. For CustomSSO, use protocol, saml, and loginEnabled. metadata, syncEnabled, and syncScope are not accepted. The SCIM synchronization switch is managed separately through UpdateIdentityProviderSCIM.
//
// ### DingTalk and Lark
//
// The metadata is merged by field. Submit only the fields that need to be updated. Omitted fields and existing keys remain unchanged. Explicitly submitting null, empty strings, or blank strings will be rejected. The server validates the complete configuration after merging. When incremental synchronization is enabled, the merged configuration must include encryptKey and verificationToken.
//
// syncEnabled controls subsequent address book change events and does not control the initial synchronization or manual synchronization. syncScope applies only to DingTalk and Lark. If it is omitted or set to null, the current scope is retained. Saving the scope does not automatically trigger a full synchronization. To perform manual synchronization, call CreateIdentityProviderSyncJob after saving and waiting for the identity provider to be ready, and then call GetIdentityProviderSyncJob to query the task result.
//
// ### Custom SAML identity provider
//
// CustomSSO currently supports only SAML. The saml parameter supports two mutually exclusive methods: providing metadataXml, or providing the four manual configuration fields: entityId, loginUrl, bindingType, and x509Certificates.
//
// metadataXml is a single IdP EntityDescriptor XML text of up to 1 MiB, not a URL or file path. DTDs and external entities are prohibited. For manual configuration, entityId must be non-empty and up to 2,048 characters, and cannot be changed after saving. loginUrl must be an absolute HTTPS URL without user credentials and fragments. bindingType must be HTTP-Redirect or HTTP-POST. Certificates must be valid X.509 signing certificates, with a maximum of 10 certificates.
//
// jitProvisionEnabled controls the automatic creation of users during the initial logon, and jitUpdateEnabled controls whether user information is updated during logon. Omitting these switches during an update retains the existing values. Modifying only the JIT switches also requires an existing valid SAML configuration. Saving the SAML configuration does not automatically enable logon. A valid configuration must exist before enabling logon. Setting loginEnabled to false only triggers a shutdown of logon and does not delete the saved SAML or SCIM configurations.
//
// ### Asynchronous results and retries
//
// When actual changes exist, the expected configuration and the UPDATING status of the identity provider are returned. The CONFIGURED status may also be returned when resource configurations are restored. Call GetIdentityProvider to wait for the READY status. If the operation fails, check for statuses such as UPDATE_FAILED. If the configuration is identical and no restoration is required, the current binding may be returned directly without generating a new task. A successful HTTP request does not indicate that the background task or upstream configuration has taken effect.
//
// Operation conflicts occur when in-flight operations exist. If a submission times out, query the configuration and status before deciding whether to retry. The status parameter describes the status of the identity provider resource and logon configuration. The SCIM operation results for CustomSSO must be queried through GetIdentityProviderSCIM and cannot be inferred from this operation.
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
// Enables or disables SCIM user synchronization independently. When enabled, the initial full import is executed asynchronously. You can query the result by calling GetIdentityProviderSCIM. The SAML logon configuration is retained.
//
// Description:
//
// ## Features and prerequisites
//
// The syncEnabled parameter is required. If you set this parameter to true, a task is submitted to independently enable SCIM and perform the initial full user import. During execution, the AgentIdentity receiving status must be Enabled, and an enabled and unexpired credential must exist. Therefore, you must first call the prepare operation and complete the URL and token configuration in the external system. If you set this parameter to false, a task is submitted to disable synchronization and upstream receiving. Existing users, credentials, and SAML logon configurations are retained. If you re-enable synchronization later and receiving is disabled, you must call the prepare operation again.
//
// Request acceptance only saves the expected switch status in AC and submits the task. The returned data is still the shared binding and does not contain the SCIM taskId. The status is not overwritten by this task and may be READY or UPDATE_FAILED from a previous logon task. Neither can be used to determine the success or failure of SCIM. Query GetIdentityProviderSCIM.operationStatus to wait for the result. Prerequisite checks for enabling are executed within the task. If the prerequisites are not met, you may receive a successful request response first, and then operationStatus=FAILED. A failure does not automatically roll back the expected switch status.
//
// SCIM does not depend on SAML configuration and does not affect the SAML status or switch. A successful task and a successful full import of empty users do not prove that continuous incremental synchronization is healthy. Whether users exist depends on whether the external system has actually pre-provisioned them.
//
// ## Requests and examples
//
// PUT /workspaces/{workspaceId}/scim-provisioning
//
// Call conventions: Call the API in the region of the target workspace by using an AK/STS identity that has permissions on the identity provider resources of the corresponding workspace. In the following examples, ws-example, URL, certificate, and Token are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API, not the additional body wrapper of the backend direct connection API. The response example only shows the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// PUT /workspaces/ws-example/scim-provisioning
//
// {"syncEnabled": true}
//
// To disable: {"syncEnabled": false}
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "READY",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": true
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call to the upstream system that fails may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 InvalidParameter: syncEnabled is missing. 404 IdentityProviderNotFound: No CustomSSO binding exists. 409 IdentityProviderOperationConflict: The user pool is not ready, is being unbound, failed to unbind, or has a concurrent status conflict. Upstream not being enabled or having no valid credentials may manifest as an asynchronous FAILED after acceptance. You cannot only handle HTTP errors.
//
// ## Call relationships and usage boundaries
//
// The call example for enabling is body={"syncEnabled":true}, and for disabling is body={"syncEnabled":false}. This API returns a snapshot of the shared binding, and its status is not the status of this SCIM task. After acceptance, use GetIdentityProviderSCIM to query operationStatus. Pre-provisioning tests or user assignments for external IdPs must still be completed in the external system. AC does not configure external IdPs for users.
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
// Enables or disables SCIM user synchronization independently. When enabled, the initial full import is executed asynchronously. You can query the result by calling GetIdentityProviderSCIM. The SAML logon configuration is retained.
//
// Description:
//
// ## Features and prerequisites
//
// The syncEnabled parameter is required. If you set this parameter to true, a task is submitted to independently enable SCIM and perform the initial full user import. During execution, the AgentIdentity receiving status must be Enabled, and an enabled and unexpired credential must exist. Therefore, you must first call the prepare operation and complete the URL and token configuration in the external system. If you set this parameter to false, a task is submitted to disable synchronization and upstream receiving. Existing users, credentials, and SAML logon configurations are retained. If you re-enable synchronization later and receiving is disabled, you must call the prepare operation again.
//
// Request acceptance only saves the expected switch status in AC and submits the task. The returned data is still the shared binding and does not contain the SCIM taskId. The status is not overwritten by this task and may be READY or UPDATE_FAILED from a previous logon task. Neither can be used to determine the success or failure of SCIM. Query GetIdentityProviderSCIM.operationStatus to wait for the result. Prerequisite checks for enabling are executed within the task. If the prerequisites are not met, you may receive a successful request response first, and then operationStatus=FAILED. A failure does not automatically roll back the expected switch status.
//
// SCIM does not depend on SAML configuration and does not affect the SAML status or switch. A successful task and a successful full import of empty users do not prove that continuous incremental synchronization is healthy. Whether users exist depends on whether the external system has actually pre-provisioned them.
//
// ## Requests and examples
//
// PUT /workspaces/{workspaceId}/scim-provisioning
//
// Call conventions: Call the API in the region of the target workspace by using an AK/STS identity that has permissions on the identity provider resources of the corresponding workspace. In the following examples, ws-example, URL, certificate, and Token are placeholder values and must be replaced with actual values. The request body example represents the body parameters of the public API, not the additional body wrapper of the backend direct connection API. The response example only shows the business part. The outer layer also includes requestId, success, code, message, and httpStatusCode. Retain the requestId for troubleshooting and do not record keys. Asynchronous business tasks do not mean that the HTTP request uses a long connection to wait for completion.
//
// ```text
//
// PUT /workspaces/ws-example/scim-provisioning
//
// {"syncEnabled": true}
//
// To disable: {"syncEnabled": false}
//
// ```
//
// ## Response example
//
// ```json
//
// {
//
//	"data": {
//
//	  "workspaceId": "ws-example",
//
//	  "identityProviderType": "CustomSSO",
//
//	  "status": "READY",
//
//	  "loginEnabled": false,
//
//	  "syncEnabled": true
//
//	}
//
// }
//
// ```
//
// ## Errors and handling
//
// Error handling: 400 InvalidParameter indicates that the parameters or current configuration prerequisites are not met. 404 IdentityProviderNotFound indicates that no binding of the matching type exists. A synchronous call to the upstream system that fails may return 503 IdentityProviderProvisioningUnavailable. Retain the request ID for troubleshooting and do not interpret a read failure as unconfigured. Actual authentication and workspace non-existence errors follow the common error handling.
//
// 400 InvalidParameter: syncEnabled is missing. 404 IdentityProviderNotFound: No CustomSSO binding exists. 409 IdentityProviderOperationConflict: The user pool is not ready, is being unbound, failed to unbind, or has a concurrent status conflict. Upstream not being enabled or having no valid credentials may manifest as an asynchronous FAILED after acceptance. You cannot only handle HTTP errors.
//
// ## Call relationships and usage boundaries
//
// The call example for enabling is body={"syncEnabled":true}, and for disabling is body={"syncEnabled":false}. This API returns a snapshot of the shared binding, and its status is not the status of this SCIM task. After acceptance, use GetIdentityProviderSCIM to query operationStatus. Pre-provisioning tests or user assignments for external IdPs must still be completed in the external system. AC does not configure external IdPs for users.
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
// Updates the configuration of a specified managed agent.
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
// Updates the configuration of a specified managed agent.
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
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update is an asynchronous process.
//
// Description:
//
// ## Operation description
//
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update is an asynchronous process.
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
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update is an asynchronous process.
//
// Description:
//
// ## Operation description
//
// Updates the address, description, authentication, or Swagger configuration of a specified MCP service. The MCP type and protocol cannot be modified after creation. The update is an asynchronous process.
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
// Updates MCP parameters by template.
//
// Description:
//
// Updates the schema-exposed parameters by using the same template version that was bound when the MCP was created, without upgrading the template version.
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
// Updates MCP parameters by template.
//
// Description:
//
// Updates the schema-exposed parameters by using the same template version that was bound when the MCP was created, without upgrading the template version.
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
// Updates the text or metadata of a long-term memory based on the memory ID.
//
// Description:
//
// Updates the text or metadata of a long-term memory based on the memory ID.
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
// Updates the text or metadata of a long-term memory based on the memory ID.
//
// Description:
//
// Updates the text or metadata of a long-term memory based on the memory ID.
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
// Updates a memory store.
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
// Updates a memory store.
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
// Updates the description, context token limit, maximum output token count, or capability configurations of a specified model.
//
// Description:
//
// This operation supports updating description, contextSize, maxTokens, and capabilities. You must specify at least one non-null parameter. Parameters that are not specified or are set to null retain their original values. The capabilities parameter is replaced as a whole object. Capability fields that are not specified in the object are treated as false.
//
// Modifying only the description does not refresh the model configurations of associated agents. When contextSize, maxTokens, or capabilities actually change, the system asynchronously refreshes managed agents that reference the model within the same workspace, as well as external agents whose model source is PLATFORM. External agents whose model source is RUNTIME are not affected. Repeatedly submitting the same configuration does not trigger a new model configuration refresh.
//
// A successful response indicates that the model configuration has been saved, but does not indicate that associated agents have completed the configuration refresh or that the runtime has loaded the new configuration. You can call GetModel to query the saved model configuration.
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
// Updates the description, context token limit, maximum output token count, or capability configurations of a specified model.
//
// Description:
//
// This operation supports updating description, contextSize, maxTokens, and capabilities. You must specify at least one non-null parameter. Parameters that are not specified or are set to null retain their original values. The capabilities parameter is replaced as a whole object. Capability fields that are not specified in the object are treated as false.
//
// Modifying only the description does not refresh the model configurations of associated agents. When contextSize, maxTokens, or capabilities actually change, the system asynchronously refreshes managed agents that reference the model within the same workspace, as well as external agents whose model source is PLATFORM. External agents whose model source is RUNTIME are not affected. Repeatedly submitting the same configuration does not trigger a new model configuration refresh.
//
// A successful response indicates that the model configuration has been saved, but does not indicate that associated agents have completed the configuration refresh or that the runtime has loaded the new configuration. You can call GetModel to query the saved model configuration.
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
// Updates the mutable configurations of a specified model connection and resubmits the publish task. The invocation protocol cannot be modified after creation.
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
// Updates the mutable configurations of a specified model connection and resubmits the publish task. The invocation protocol cannot be modified after creation.
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
// Updates the business tags of a specified skill. Other attributes not provided in the request remain unchanged.
//
// Description:
//
// ## Request description
//
// Updates the business tags of a specified skill. Other attributes not provided in the request remain unchanged.
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
// Updates the business tags of a specified skill. Other attributes not provided in the request remain unchanged.
//
// Description:
//
// ## Request description
//
// Updates the business tags of a specified skill. Other attributes not provided in the request remain unchanged.
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
// Updates the version labels and their mappings for a specified Skill.
//
// Description:
//
// ## Operation description
//
// Updates the version labels and their mappings for a specified Skill.
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
// Updates the version labels and their mappings for a specified Skill.
//
// Description:
//
// ## Operation description
//
// Updates the version labels and their mappings for a specified Skill.
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
// ## Request description
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
// ## Request description
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
// Updates the description and members of a specified team. When users or agents are passed, the corresponding member list is replaced based on full overwrite semantics. Member lists that are not passed remain unchanged. The team name cannot be modified after creation.
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
// Updates the description and members of a specified team. When users or agents are passed, the corresponding member list is replaced based on full overwrite semantics. Member lists that are not passed remain unchanged. The team name cannot be modified after creation.
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
// Updates the display name, email address, or note of a specified user. You must specify at least one of the displayName, email, and note parameters. Unspecified attributes remain unchanged. The username cannot be modified after creation.
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
// Updates the display name, email address, or note of a specified user. You must specify at least one of the displayName, email, and note parameters. Unspecified attributes remain unchanged. The username cannot be modified after creation.
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
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized state can be updated. Status, TenantId, and RegionId are maintained by the server and cannot be modified through this API operation.
//
// Description:
//
// ## Request description\\nUpdates the name or network configuration of a workspace. Only workspaces in the `Initialized` state can be updated. `Status`, `TenantId`, and `RegionId` are maintained by the server and cannot be modified through this API operation. The network configuration uses `Enabled` to specify whether to enable the VPC network. If enabled, you must provide `VpcId` and at least one `VSwitchIds`.\\n.
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
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized state can be updated. Status, TenantId, and RegionId are maintained by the server and cannot be modified through this API operation.
//
// Description:
//
// ## Request description\\nUpdates the name or network configuration of a workspace. Only workspaces in the `Initialized` state can be updated. `Status`, `TenantId`, and `RegionId` are maintained by the server and cannot be modified through this API operation. The network configuration uses `Enabled` to specify whether to enable the VPC network. If enabled, you must provide `VpcId` and at least one `VSwitchIds`.\\n.
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
// Imports an AgentSpec into the current workspace by retrieving and parsing an uploaded AgentSpec ZIP package from OSS.
//
// Description:
//
// ## Request description
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses the package, and imports the AgentSpec into the current workspace.
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
// Imports an AgentSpec into the current workspace by retrieving and parsing an uploaded AgentSpec ZIP package from OSS.
//
// Description:
//
// ## Request description
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses the package, and imports the AgentSpec into the current workspace.
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
// Retrieves an uploaded Skill ZIP package from OSS, parses it, and imports the Skill into the current workspace.
//
// Description:
//
// ## Operation description
//
// Retrieves an uploaded Skill ZIP package from OSS, parses it, and imports the Skill into the current workspace.
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
// Retrieves an uploaded Skill ZIP package from OSS, parses it, and imports the Skill into the current workspace.
//
// Description:
//
// ## Operation description
//
// Retrieves an uploaded Skill ZIP package from OSS, parses it, and imports the Skill into the current workspace.
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
// Validates the OSS code package source for an agent.
//
// Description:
//
// Uses the current workspace role to check whether the customer\\"s OSS object exists and is readable, and checks the file name extension and compressed size. Returns the sourceUri, file name, size, and format read in this request. The bucket and object key are passed separately, and HTTPS download URLs are not accepted.
//
// The validation uses object metadata and read probing. It does not perform full downloads, transfers, or decompression, nor does it generate a SHA-256 hash. Passing the validation does not mean that the archive content is complete or can be started. When you submit a CreateCustomAgent or UpdateCustomAgent request, the system re-validates and transfers the package. You do not need to invoke this operation in advance as a prerequisite for submission.
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
// Validates the OSS code package source for an agent.
//
// Description:
//
// Uses the current workspace role to check whether the customer\\"s OSS object exists and is readable, and checks the file name extension and compressed size. Returns the sourceUri, file name, size, and format read in this request. The bucket and object key are passed separately, and HTTPS download URLs are not accepted.
//
// The validation uses object metadata and read probing. It does not perform full downloads, transfers, or decompression, nor does it generate a SHA-256 hash. Passing the validation does not mean that the archive content is complete or can be started. When you submit a CreateCustomAgent or UpdateCustomAgent request, the system re-validates and transfers the package. You do not need to invoke this operation in advance as a prerequisite for submission.
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
// Checks the OSS RAM authorization status of the agent code package.
//
// Description:
//
// Queries the binding status of the authorization policy for a specified source bucket to the current workspace role. AUTHORIZED indicates that the policy is bound. UNAUTHORIZED indicates that the policy is not bound. A bound policy does not indicate that a specific code package object exists or is readable.
//
// To check a specific object, call ValidateAgentCodePackageFromOss. This operation does not copy code packages or modify the default OSS binding of the workspace.
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
// Checks the OSS RAM authorization status of the agent code package.
//
// Description:
//
// Queries the binding status of the authorization policy for a specified source bucket to the current workspace role. AUTHORIZED indicates that the policy is bound. UNAUTHORIZED indicates that the policy is not bound. A bound policy does not indicate that a specific code package object exists or is readable.
//
// To check a specific object, call ValidateAgentCodePackageFromOss. This operation does not copy code packages or modify the default OSS binding of the workspace.
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
// Validates whether one or more Qoder Service Account Keys are available at a specified site. This operation does not save keys or modify the Connector status.
//
// Description:
//
// Validates whether the credentials of a specified Connector are valid and returns a list of invalid Service Account Keys.
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
// Validates whether one or more Qoder Service Account Keys are available at a specified site. This operation does not save keys or modify the Connector status.
//
// Description:
//
// Validates whether the credentials of a specified Connector are valid and returns a list of invalid Service Account Keys.
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
// Verifies the ACR RAM authorization for a workspace.
//
// Description:
//
// ACR and user OSS mounts share the workspace role. If no existing role is found, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is retained and only new permissions are appended without overwriting existing OSS policies. Fixed historical shared roles continue to be reused and are not recreated or renamed due to new naming templates. This operation performs a read-only check to verify whether the custom policies corresponding to the target instance, namespace, and repository are attached to the shared role, and checks the workspace prerequisites. It does not create execution versions or trigger configuration delivery. AUTHORIZED only indicates that the custom policy for the target name is attached, without parsing the policy content or simulating the final permissions. ELIGIBLE does not mean that RuntimeManager capacity or template admission is complete, nor does it require ACR configuration to be delivered. Creating or updating an agent re-verifies permissions and processes ACR dependencies on the backend. Configuration waiting, failures, and retries are reported through agent details.
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
// Verifies the ACR RAM authorization for a workspace.
//
// Description:
//
// ACR and user OSS mounts share the workspace role. If no existing role is found, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is retained and only new permissions are appended without overwriting existing OSS policies. Fixed historical shared roles continue to be reused and are not recreated or renamed due to new naming templates. This operation performs a read-only check to verify whether the custom policies corresponding to the target instance, namespace, and repository are attached to the shared role, and checks the workspace prerequisites. It does not create execution versions or trigger configuration delivery. AUTHORIZED only indicates that the custom policy for the target name is attached, without parsing the policy content or simulating the final permissions. ELIGIBLE does not mean that RuntimeManager capacity or template admission is complete, nor does it require ACR configuration to be delivered. Creating or updating an agent re-verifies permissions and processes ACR dependencies on the backend. Configuration waiting, failures, and retries are reported through agent details.
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
// Queries the RAM authorization status of AgenticFS mounting for a workspace.
//
// Description:
//
// Determines the shared role and target access point based on the workspace and server, checks whether the role has a mounted policy with the expected name and a type of Custom, and returns AUTHORIZED or UNAUTHORIZED. Consistent with OSS/ACR, this operation does not verify the policy body, role trust, or actual mount read/write permissions. If the upstream query fails, an error is returned.
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
// Queries the RAM authorization status of AgenticFS mounting for a workspace.
//
// Description:
//
// Determines the shared role and target access point based on the workspace and server, checks whether the role has a mounted policy with the expected name and a type of Custom, and returns AUTHORIZED or UNAUTHORIZED. Consistent with OSS/ACR, this operation does not verify the policy body, role trust, or actual mount read/write permissions. If the upstream query fails, an error is returned.
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
// Verifies the RAM authorization for a workspace to access an OSS bucket and returns the current authorization status. When bucketName is specified, verifies the authorization for an agent-mounted bucket. When bucketName is not specified, verifies and updates the authorization status of the workspace\\"s private OSS storage.
//
// Description:
//
// Queries whether the workspace OSS mount role is bound to the custom RAM policy corresponding to the target bucket, and returns AUTHORIZED or UNAUTHORIZED. When bucketName is not specified, the existing user-managed OSS binding of the workspace is used and the authorization status is saved. When bucketName is specified, only the authorization status of the specified bucket is queried without modifying the workspace OSS binding. This operation does not verify OSS data plane access permissions or resume workspace initialization tasks. Alibaba Cloud Container Registry (ACR) and user OSS mounts share the workspace role. If no existing role is available, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is reused and permissions are only appended without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not recreated or renamed due to new naming templates.
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
// Verifies the RAM authorization for a workspace to access an OSS bucket and returns the current authorization status. When bucketName is specified, verifies the authorization for an agent-mounted bucket. When bucketName is not specified, verifies and updates the authorization status of the workspace\\"s private OSS storage.
//
// Description:
//
// Queries whether the workspace OSS mount role is bound to the custom RAM policy corresponding to the target bucket, and returns AUTHORIZED or UNAUTHORIZED. When bucketName is not specified, the existing user-managed OSS binding of the workspace is used and the authorization status is saved. When bucketName is specified, only the authorization status of the specified bucket is queried without modifying the workspace OSS binding. This operation does not verify OSS data plane access permissions or resume workspace initialization tasks. Alibaba Cloud Container Registry (ACR) and user OSS mounts share the workspace role. If no existing role is available, AliyunAgentCoreWorkspace-{workspaceId} is used. If a valid legacy OSS mount role exists, it is reused and permissions are only appended without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not recreated or renamed due to new naming templates.
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
