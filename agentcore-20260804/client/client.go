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
// Deletes models in a specified workspace in batches. If any model is in use, the entire batch request fails.
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
// Deletes models in a specified workspace in batches. If any model is in use, the entire batch request fails.
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
// Uploads Skill ZIP packages in bulk through OSS and returns the processing result of each Skill.
//
// Description:
//
// ## Operation description
//
// Uploads Skill ZIP packages in bulk through OSS and returns the processing result of each Skill.
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
// Uploads Skill ZIP packages in bulk through OSS and returns the processing result of each Skill.
//
// Description:
//
// ## Operation description
//
// Uploads Skill ZIP packages in bulk through OSS and returns the processing result of each Skill.
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
// Converts an MCP to free editing.
//
// Description:
//
// Disables template usage constraints. After the conversion, the MCP retains its source and tags but no longer appears on the usage page.
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
// Converts an MCP to free editing.
//
// Description:
//
// Disables template usage constraints. After the conversion, the MCP retains its source and tags but no longer appears on the usage page.
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
// Creates an IM channel for a specified agent and binds a publicly accessible ServiceEndpoint.
//
// Description:
//
// Creates an IM channel for a specified agent and binds a publicly accessible ServiceEndpoint.
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
// Creates an IM channel for a specified agent and binds a publicly accessible ServiceEndpoint.
//
// Description:
//
// Creates an IM channel for a specified agent and binds a publicly accessible ServiceEndpoint.
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
// Creates an AgentSpec in the specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
//
// Description:
//
// ## Operation description
//
// Creates an AgentSpec in the specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
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
// Creates an AgentSpec in the specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
//
// Description:
//
// ## Operation description
//
// Creates an AgentSpec in the specified workspace and generates the first draft version. Returns a resource conflict error if an AgentSpec with the same name already exists.
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
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist, and there must be no draft currently being edited.
//
// Description:
//
// ## Request description
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist, and there must be no draft currently being edited.
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
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist, and there must be no draft currently being edited.
//
// Description:
//
// ## Request description
//
// Creates a new draft version for an existing AgentSpec. The AgentSpec must exist, and there must be no draft currently being edited.
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
// Creates a credential in a specified workspace for authentication when an agent accesses external services. Currently, only the apiKey type is supported. The credential content is passed in as a JSON string through credentialMetadata and can only be queried in masked form after being written.
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
// Creates a credential in a specified workspace for authentication when an agent accesses external services. Currently, only the apiKey type is supported. The credential content is passed in as a JSON string through credentialMetadata and can only be queried in masked form after being written.
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
// Creates a bootstrap token and CMS configuration required for connecting a specified external agent.
//
// Description:
//
// Creates a bootstrap token and CMS configuration required for connecting a specified external agent.
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
// Creates a bootstrap token and CMS configuration required for connecting a specified external agent.
//
// Description:
//
// Creates a bootstrap token and CMS configuration required for connecting a specified external agent.
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
// Binds an external identity provider to a specified workspace for single sign-on and organization member synchronization. Each workspace can be bound to at most one external identity provider. The binding is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track the progress.
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
// Binds an external identity provider to a specified workspace for single sign-on and organization member synchronization. Each workspace can be bound to at most one external identity provider. The binding is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track the progress.
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
// Creates a model configuration under a specified model connection in a workspace.
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
// Creates a model configuration under a specified model connection in a workspace.
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
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials.
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
// Creates a model connection in a specified workspace and configures the upstream model service address, invoke protocol, and access credentials.
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
// Creates a team in a specified workspace and sets user members and agent members at the same time. The user members must include exactly one member with the ADMIN role. Agent members can only have the LEADER or WORKER role.
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
// Creates a team in a specified workspace and sets user members and agent members at the same time. The user members must include exactly one member with the ADMIN role. Agent members can only have the LEADER or WORKER role.
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
// Creates a user in a specified workspace. The username must be unique within the workspace and can contain only lowercase letters, digits, and hyphens. Reserved names such as manager, admin, or names starting with worker- cannot be used. If password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
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
// Creates a user in a specified workspace. The username must be unique within the workspace and can contain only lowercase letters, digits, and hyphens. Reserved names such as manager, admin, or names starting with worker- cannot be used. If password is not specified, the server generates an initial password and returns it in the initialPassword field of the response.
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
// ## Operation description\\nCreates an AgentCore workspace control plane record. The server completes the state transition from `Initializing` to `Initialized` within the same transaction. Use `Enabled` in the network configuration to specify whether to enable VPC networking. If VPC networking is enabled, you must provide a `VpcId` and at least one entry in `VSwitchIds`.\\n.
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
// ## Operation description\\nCreates an AgentCore workspace control plane record. The server completes the state transition from `Initializing` to `Initialized` within the same transaction. Use `Enabled` in the network configuration to specify whether to enable VPC networking. If VPC networking is enabled, you must provide a `VpcId` and at least one entry in `VSwitchIds`.\\n.
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
// Deletes a specified AgentSpec along with all its versions and metadata. This operation is irreversible.
//
// Description:
//
// ## Request description
//
// Deletes a specified AgentSpec along with all its versions and metadata. This operation is irreversible.
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
// Deletes a specified AgentSpec along with all its versions and metadata. This operation is irreversible.
//
// Description:
//
// ## Request description
//
// Deletes a specified AgentSpec along with all its versions and metadata. This operation is irreversible.
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
// Deletes a credential from a specified workspace and removes the ciphertext hosted in the credential service. After deletion, agents that are bound to this credential can no longer retrieve the 
//
// Description:
//
// Deletes an access credential from a specified workspace. A credential cannot be deleted while it is still bound to an MCP service.
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
// Deletes a credential from a specified workspace and removes the ciphertext hosted in the credential service. After deletion, agents that are bound to this credential can no longer retrieve the credential content.
//
// Description:
//
// Deletes an access credential from a specified workspace. A credential cannot be deleted while it is still bound to an MCP service.
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
// Unbinds the external identity provider from a specified workspace and cleans up users synchronized by that identity provider. The unbinding is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track the progress.
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
// Unbinds the external identity provider from a specified workspace and cleans up users synchronized by that identity provider. The unbinding is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track the progress.
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
// Deletes a specified MCP service. The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
//
// Description:
//
// ## Operation description
//
// Deletes a specified MCP service. The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
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
// ## Operation description
//
// Deletes a specified MCP service. The deletion is an asynchronous process. After the deletion is complete, the MCP service is no longer returned.
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
// ## Request description
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
// ## Request description
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
// ## Request description
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
// ## Request description
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
// Deletes the control plane record of a specified workspace. The server completes the state transition from Deleting to Deleted within the same transaction. When you repeatedly delete a workspace that is in the Deleting or Deleted state, the server handles the request with idempotence semantics.
//
// Description:
//
// ## Request description\\nDeletes the control plane record of a specified workspace. The server completes the state transition from `Deleting` to `Deleted` within the same transaction. When you repeatedly delete a workspace that is in the `Deleting` or `Deleted` state, the server handles the request with idempotence semantics.\\n.
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
// Deletes the control plane record of a specified workspace. The server completes the state transition from Deleting to Deleted within the same transaction. When you repeatedly delete a workspace that is in the Deleting or Deleted state, the server handles the request with idempotence semantics.
//
// Description:
//
// ## Request description\\nDeletes the control plane record of a specified workspace. The server completes the state transition from `Deleting` to `Deleted` within the same transaction. When you repeatedly delete a workspace that is in the `Deleting` or `Deleted` state, the server handles the request with idempotence semantics.\\n.
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
// Retrieves a pre-signed OSS download URL for a specified AgentSpec, which is used to download the AgentSpec ZIP package.
//
// Description:
//
// ## Operation description
//
// Retrieves a pre-signed OSS download URL for a specified AgentSpec, which is used to download the AgentSpec ZIP package.
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
// Retrieves a pre-signed OSS download URL for a specified AgentSpec, which is used to download the AgentSpec ZIP package.
//
// Description:
//
// ## Operation description
//
// Retrieves a pre-signed OSS download URL for a specified AgentSpec, which is used to download the AgentSpec ZIP package.
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
// ## Request description
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
// ## Request description
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
// Skips the regular review process and forcibly publishes the specified Skill version.
//
// Description:
//
// ## Request description
//
// Skips the regular review process and forcibly publishes the specified Skill version.
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
// Skips the regular review process and forcibly publishes the specified Skill version.
//
// Description:
//
// ## Request description
//
// Skips the regular review process and forcibly publishes the specified Skill version.
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
// ## Operation description
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
// ## Operation description
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
// Retrieves the OSS pre-signed upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, call the AgentSpec OSS upload operation to complete the import.
//
// Description:
//
// ## Operation description
//
// Retrieves the OSS pre-signed upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, call the AgentSpec OSS upload operation to complete the import.
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
// Retrieves the OSS pre-signed upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, call the AgentSpec OSS upload operation to complete the import.
//
// Description:
//
// ## Operation description
//
// Retrieves the OSS pre-signed upload URL and object name required to import an AgentSpec ZIP package. After the upload is complete, call the AgentSpec OSS upload operation to complete the import.
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
// Queries the complete content of a specified AgentSpec version, including manifest content, resource files, and metadata.
//
// Description:
//
// ## Operation description
//
// Queries the complete content of a specified AgentSpec version, including manifest content, resource files, and metadata.
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
// Queries the complete content of a specified AgentSpec version, including manifest content, resource files, and metadata.
//
// Description:
//
// ## Operation description
//
// Queries the complete content of a specified AgentSpec version, including manifest content, resource files, and metadata.
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
// Queries the details of a specified credential and returns the list of agents bound to the credential. The credential content is returned in masked form.
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
// Queries the details of a specified credential and returns the list of agents bound to the credential. The credential content is returned in masked form.
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
// Queries the available network types for a specified external agent.
//
// Description:
//
// Queries the available network types for a specified external agent.
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
// Queries the available network types for a specified external agent.
//
// Description:
//
// Queries the available network types for a specified external agent.
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
// Queries the binding details of an external identity provider for a specified workspace, including the binding status, application configuration, and callback URLs that need to be configured on the identity provider side. Application secret configurations are not returned.
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
// Queries the binding details of an external identity provider for a specified workspace, including the binding status, application configuration, and callback URLs that need to be configured on the identity provider side. Application secret configurations are not returned.
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
// Queries the details of a specified MCP service, including the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// ## Operation description
//
// Queries the details of a specified MCP service, including the address, type, status, authentication configuration, and protocol.
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
// Queries the details of a specified MCP service, including the address, type, status, authentication configuration, and protocol.
//
// Description:
//
// ## Operation description
//
// Queries the details of a specified MCP service, including the address, type, status, authentication configuration, and protocol.
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
// Queries the currently active API Key and its invocation method for a specified service endpoint within a workspace.
//
// Description:
//
// Queries the currently active API Key for a specified service endpoint. The call succeeds only when the service endpoint has API_KEY authentication enabled and the gateway consumer and credentials are ready. The service reads the API Key from the gateway in real time. AgentCore does not persist the plaintext. Keep the returned API Key secure and avoid logging it or exposing it in public configurations.
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
// Queries the currently active API Key and its invocation method for a specified service endpoint within a workspace.
//
// Description:
//
// Queries the currently active API Key for a specified service endpoint. The call succeeds only when the service endpoint has API_KEY authentication enabled and the gateway consumer and credentials are ready. The service reads the API Key from the gateway in real time. AgentCore does not persist the plaintext. Keep the returned API Key secure and avoid logging it or exposing it in public configurations.
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
// Retrieves the OSS pre-signed upload URL and object name required for importing a Skill ZIP package. After the upload is complete, call the Skill OSS upload operation to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS pre-signed upload URL and object name required for importing a Skill ZIP package. After the upload is complete, call the Skill OSS upload operation to complete the import.
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
// Retrieves the OSS pre-signed upload URL and object name required for importing a Skill ZIP package. After the upload is complete, call the Skill OSS upload operation to complete the import.
//
// Description:
//
// ## Request description
//
// Retrieves the OSS pre-signed upload URL and object name required for importing a Skill ZIP package. After the upload is complete, call the Skill OSS upload operation to complete the import.
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
// Queries the details of a specified user in a specified workspace. Returns an error if the user does not exist.
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
// Queries the details of a specified user in a specified workspace. Returns an error if the user does not exist.
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
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policy.
//
// Description:
//
// ## Operation description\\nQueries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policy.\\n.
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
// Queries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policy.
//
// Description:
//
// ## Operation description\\nQueries the details of a workspace by workspace ID, including the lifecycle status, CMS Workspace, AIRegistry Namespace, and current network policy.\\n.
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
// Queries the ACR RAM authorization URL for a workspace.
//
// Description:
//
// ACR and user OSS mounts share the same workspace role. If no existing role is found, the system uses AliyunAgentCoreWorkspace-{workspaceId}. If a valid legacy OSS mount role exists, the system reuses it and only appends permissions without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not rebuilt or renamed because of new naming templates. The acrInstanceId, namespace, and repository parameters are required. The region and account are derived from the workspace context. The response contains the backend-resolved role and the authorization URL for the target repository. This operation does not create ACR or Agent execution versions or deliver runtime configurations. After authorization is complete, you can call Verify, and then create or update an Agent without calling a separate Ensure operation.
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
// Queries the ACR RAM authorization URL for a workspace.
//
// Description:
//
// ACR and user OSS mounts share the same workspace role. If no existing role is found, the system uses AliyunAgentCoreWorkspace-{workspaceId}. If a valid legacy OSS mount role exists, the system reuses it and only appends permissions without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not rebuilt or renamed because of new naming templates. The acrInstanceId, namespace, and repository parameters are required. The region and account are derived from the workspace context. The response contains the backend-resolved role and the authorization URL for the target repository. This operation does not create ACR or Agent execution versions or deliver runtime configurations. After authorization is complete, you can call Verify, and then create or update an Agent without calling a separate Ensure operation.
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
// Validates input based on the specified template version and creates an MCP in the workspace.
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
// Validates input based on the specified template version and creates an MCP in the workspace.
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
// Installs a plugin for a specified AgentCore workspace. Currently supports the collaboration plugin. The installation process is executed asynchronously.
//
// Description:
//
// ## Operation description\\nInstalls a plugin for a specified AgentCore workspace. Currently supports the `collaboration` plugin. Plugin configuration is passed through the `Config` parameter, and different plugins can define their own configuration structures. The `collaboration` plugin uses `Config.NetworkConfiguration` to specify VPC and public network access policies. The installation process is executed asynchronously. When you repeatedly call this operation for a plugin with the same name that is being installed or already installed, the operation returns the current status with idempotent semantics if the configuration is the same. If the configuration is different, the operation returns an operation conflict error.\\n.
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
// Installs a plugin for a specified AgentCore workspace. Currently supports the collaboration plugin. The installation process is executed asynchronously.
//
// Description:
//
// ## Operation description\\nInstalls a plugin for a specified AgentCore workspace. Currently supports the `collaboration` plugin. Plugin configuration is passed through the `Config` parameter, and different plugins can define their own configuration structures. The `collaboration` plugin uses `Config.NetworkConfiguration` to specify VPC and public network access policies. The installation process is executed asynchronously. When you repeatedly call this operation for a plugin with the same name that is being installed or already installed, the operation returns the current status with idempotent semantics if the configuration is the same. If the configuration is different, the operation returns an operation conflict error.\\n.
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
// Queries AgentSpec resources in a specified workspace by using paging, supporting name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// ## Operation description
//
// Queries AgentSpec resources in a specified workspace by using paging, supporting name search, sorting, and filtering by owner, visibility scope, and business labels.
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
// Queries AgentSpec resources in a specified workspace by using paging, supporting name search, sorting, and filtering by owner, visibility scope, and business labels.
//
// Description:
//
// ## Operation description
//
// Queries AgentSpec resources in a specified workspace by using paging, supporting name search, sorting, and filtering by owner, visibility scope, and business labels.
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
// Queries the list of teams to which specified agents belong. Specify the agent IDs by using agentIds. The response includes the membership information between each agent and its teams, such as the team ID, team name, and the role that the agent assumes in the team.
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
// Queries the list of teams to which specified agents belong. Specify the agent IDs by using agentIds. The response includes the membership information between each agent and its teams, such as the team ID, team name, and the role that the agent assumes in the team.
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
// Queries credentials in a specified workspace with paging. Use credentialType to filter by type, nameLike to perform a fuzzy match on credential names, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Queries credentials in a workspace with paging. Supports filtering by type and name.
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
// Queries credentials in a specified workspace with paging. Use credentialType to filter by type, nameLike to perform a fuzzy match on credential names, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default.
//
// Description:
//
// Queries credentials in a workspace with paging. Supports filtering by type and name.
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
// Queries the external identity provider bound to a specified workspace. Each workspace can be bound to at most one external identity provider, so the response returns at most one record. Application secret configurations are not returned.
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
// Queries the external identity provider bound to a specified workspace. Each workspace can be bound to at most one external identity provider, so the response returns at most one record. Application secret configurations are not returned.
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
// 查询预定义模型供应商目录
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
// 查询预定义模型供应商目录
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
// Queries the models and their capability information for a specified provider in the AgentCore built-in model catalog.
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
// Queries the models and their capability information for a specified provider in the AgentCore built-in model catalog.
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
// Queries service endpoints in a specified workspace by using paging. Supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// ## Operation description\\nQueries service endpoints in a specified workspace by using paging. Filter results by targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status. Use maxResults to specify the maximum number of records per page and nextToken to retrieve the next page. If maxResults is not specified, the server returns 20 records by default.\\n
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
// Queries service endpoints in a specified workspace by using paging. Supports filtering by target type, agent, collaboration component, and status.
//
// Description:
//
// ## Operation description\\nQueries service endpoints in a specified workspace by using paging. Filter results by targetType, agentId, agentVersion, resourceBindingId, collaborationComponent, and status. Use maxResults to specify the maximum number of records per page and nextToken to retrieve the next page. If maxResults is not specified, the server returns 20 records by default.\\n
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
// Performs a paged query of Skills in a specified workspace, and returns basic Skill information, version status, and paging details.
//
// Description:
//
// ## Operation description
//
// Performs a paged query of Skills in a specified workspace, and returns basic Skill information, version status, and paging details.
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
// Performs a paged query of Skills in a specified workspace, and returns basic Skill information, version status, and paging details.
//
// Description:
//
// ## Operation description
//
// Performs a paged query of Skills in a specified workspace, and returns basic Skill information, version status, and paging details.
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
// Performs a paged query of teams in a specified workspace. Use nameLike to filter by team name with fuzzy match, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. Member information in the list includes only the member identity, name, and team role.
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
// Performs a paged query of teams in a specified workspace. Use nameLike to filter by team name with fuzzy match, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default. Member information in the list includes only the member identity, name, and team role.
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
// Queries users in a specified workspace with paging. Use nameLike to filter by username with fuzzy match, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default.
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
// Queries users in a specified workspace with paging. Use nameLike to filter by username with fuzzy match, maxResults to specify the maximum number of records per page, and nextToken to retrieve the next page. If maxResults is not specified, the server returns 10 records by default.
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
// Queries workspaces under the current tenant by paging. The list does not return soft-deleted records with a status of Deleted by default. Results are stably sorted by creation order on the server side.
//
// Description:
//
// ## Operation description\\nQueries workspaces under the current tenant by paging. The list does not return soft-deleted records with a status of `Deleted` by default. Results are stably sorted by creation order on the server side. Use `nextToken` to retrieve the next page, `skip` to skip a specified number of workspaces, `maxResults` to specify the maximum number of records per paging request, and `nameLike` to filter workspaces by name using fuzzy match. If `maxResults` is not specified or is set to 0, the server returns 20 records by default.\\n
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
// Queries workspaces under the current tenant by paging. The list does not return soft-deleted records with a status of Deleted by default. Results are stably sorted by creation order on the server side.
//
// Description:
//
// ## Operation description\\nQueries workspaces under the current tenant by paging. The list does not return soft-deleted records with a status of `Deleted` by default. Results are stably sorted by creation order on the server side. Use `nextToken` to retrieve the next page, `skip` to skip a specified number of workspaces, `maxResults` to specify the maximum number of records per paging request, and `nameLike` to filter workspaces by name using fuzzy match. If `maxResults` is not specified or is set to 0, the server returns 20 records by default.\\n
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
// ## Request description
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
// ## Request description
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
// ## Request description
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
// ## Request description
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
// Publishes a specified Skill version to change its state to published.
//
// Description:
//
// ## Operation description
//
// Publishes a specified Skill version to change its state to published.
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
// Publishes a specified Skill version to change its state to published.
//
// Description:
//
// ## Operation description
//
// Publishes a specified Skill version to change its state to published.
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
// Converts a specified Skill version back into an editable draft for further modifications.
//
// Description:
//
// ## Operation description
//
// Converts a specified Skill version back into an editable draft for further modifications.
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
// Converts a specified Skill version back into an editable draft for further modifications.
//
// Description:
//
// ## Operation description
//
// Converts a specified Skill version back into an editable draft for further modifications.
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
// Resets the logon password of a specified user. Specify the user by agentCoreUserId or username. At least one of the two parameters must be specified. Only users who use workspace local password authentication can be reset. If password is not specified, the server generates a random password and returns it in the response.
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
// Resets the logon password of a specified user. Specify the user by agentCoreUserId or username. At least one of the two parameters must be specified. Only users who use workspace local password authentication can be reset. If password is not specified, the server generates a random password and returns it in the response.
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
// Submits a specified AgentSpec draft version for review. If no release pipeline is configured, the version is published directly to online status.
//
// Description:
//
// ## Operation description
//
// Submits a specified AgentSpec draft version for review. If no release pipeline is configured, the version is published directly to online status.
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
// Submits a specified AgentSpec draft version for review. If no release pipeline is configured, the version is published directly to online status.
//
// Description:
//
// ## Operation description
//
// Submits a specified AgentSpec draft version for review. If no release pipeline is configured, the version is published directly to online status.
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
// Updates the ServiceEndpoint binding, enabled/disabled status, or channel behavior configuration of an IM channel. At least one updatable field must be provided.
//
// Description:
//
// Updates the ServiceEndpoint binding, enabled/disabled status, or channel behavior configuration of an IM channel. At least one updatable field must be provided.
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
// Updates the ServiceEndpoint binding, enabled/disabled status, or channel behavior configuration of an IM channel. At least one updatable field must be provided.
//
// Description:
//
// Updates the ServiceEndpoint binding, enabled/disabled status, or channel behavior configuration of an IM channel. At least one updatable field must be provided.
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
// Performs a full update of the channel credential for a specified IM channel of an agent. Secrets are not returned in the response.
//
// Description:
//
// Performs a full update of the channel credential for a specified IM channel of an agent. Secrets are not returned in the response.
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
// Performs a full update of the channel credential for a specified IM channel of an agent. Secrets are not returned in the response.
//
// Description:
//
// Performs a full update of the channel credential for a specified IM channel of an agent. Secrets are not returned in the response.
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
// Updates the content or description of a specified credential. At least one of credentialMetadata and description must be specified. Unspecified properties remain unchanged. The credential name and credential type cannot be modified after creation.
//
// Description:
//
// Updates the metadata or resource scope of a specified credential.
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
// Updates the content or description of a specified credential. At least one of credentialMetadata and description must be specified. Unspecified properties remain unchanged. The credential name and credential type cannot be modified after creation.
//
// Description:
//
// Updates the metadata or resource scope of a specified credential.
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
// Updates the login toggle, member synchronization toggle, or application configuration of a specified external identity provider in a workspace. Unspecified properties remain unchanged. The update is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track progress.
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
// Updates the login toggle, member synchronization toggle, or application configuration of a specified external identity provider in a workspace. Unspecified properties remain unchanged. The update is an asynchronous operation. After the API returns, you can call GetIdentityProvider to query the status and track progress.
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
// Updates the schema-exposed parameters by using the same template version that was bound when the MCP was created. This operation does not upgrade the template version.
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
// Updates the schema-exposed parameters by using the same template version that was bound when the MCP was created. This operation does not upgrade the template version.
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
// Updates the description, context token limit, maximum output token count, or capability configuration of a specified model.
//
// Description:
//
// This operation supports updating description, contextSize, maxTokens, and capabilities. At least one non-null parameter must be provided. Parameters that are not provided or set to null retain their original values. The capabilities object is replaced as a whole. Capability fields not included in the object are treated as false.
//
// Modifying only description does not refresh the model configuration of associated Agents. When contextSize, maxTokens, or capabilities actually change, the system asynchronously refreshes managed Agents that reference the model within the same workspace, as well as external Agents whose model source is PLATFORM. External Agents whose model source is RUNTIME are not affected. Submitting the same configuration repeatedly does not trigger a new model configuration refresh.
//
// A successful response indicates that the model configuration has been saved. It does not indicate that associated Agents have completed the configuration refresh or that the runtime has loaded the new configuration. Call GetModel to query the saved model configuration.
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
// This operation supports updating description, contextSize, maxTokens, and capabilities. At least one non-null parameter must be provided. Parameters that are not provided or set to null retain their original values. The capabilities object is replaced as a whole. Capability fields not included in the object are treated as false.
//
// Modifying only description does not refresh the model configuration of associated Agents. When contextSize, maxTokens, or capabilities actually change, the system asynchronously refreshes managed Agents that reference the model within the same workspace, as well as external Agents whose model source is PLATFORM. External Agents whose model source is RUNTIME are not affected. Submitting the same configuration repeatedly does not trigger a new model configuration refresh.
//
// A successful response indicates that the model configuration has been saved. It does not indicate that associated Agents have completed the configuration refresh or that the runtime has loaded the new configuration. Call GetModel to query the saved model configuration.
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
// Updates the mutable configuration of a specified model connection and resubmits the publish task. The protocol cannot be modified after the model connection is created.
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
// Updates the mutable configuration of a specified model connection and resubmits the publish task. The protocol cannot be modified after the model connection is created.
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
// Updates the business tags of a specified Skill. Other attributes that are not included in the request remain unchanged.
//
// Description:
//
// ## Operation description
//
// Updates the business tags of a specified Skill. Other attributes that are not included in the request remain unchanged.
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
// Updates the business tags of a specified Skill. Other attributes that are not included in the request remain unchanged.
//
// Description:
//
// ## Operation description
//
// Updates the business tags of a specified Skill. Other attributes that are not included in the request remain unchanged.
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
// ## Request description
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
// ## Request description
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
// Updates the description and members of a specified team. When users or agents are passed in, the corresponding member list is replaced using full overwrite semantics. Member lists that are not passed in remain unchanged. The team name cannot be modified after creation.
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
// Updates the description and members of a specified team. When users or agents are passed in, the corresponding member list is replaced using full overwrite semantics. Member lists that are not passed in remain unchanged. The team name cannot be modified after creation.
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
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized status can be updated. The Status, TenantId, and RegionId fields are maintained by the server and cannot be modified through this operation.
//
// Description:
//
// ## Operation description\\nUpdates the name or network configuration of a workspace. Only workspaces in the `Initialized` status can be updated. The `Status`, `TenantId`, and `RegionId` fields are maintained by the server and cannot be modified through this operation. The network configuration uses `Enabled` to specify whether to enable VPC networking. When enabled, you must also provide a `VpcId` and at least one entry in `VSwitchIds`.\\n.
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
// Updates the name or network configuration of a workspace. Only workspaces in the Initialized status can be updated. The Status, TenantId, and RegionId fields are maintained by the server and cannot be modified through this operation.
//
// Description:
//
// ## Operation description\\nUpdates the name or network configuration of a workspace. Only workspaces in the `Initialized` status can be updated. The `Status`, `TenantId`, and `RegionId` fields are maintained by the server and cannot be modified through this operation. The network configuration uses `Enabled` to specify whether to enable VPC networking. When enabled, you must also provide a `VpcId` and at least one entry in `VSwitchIds`.\\n.
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
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses it, and imports the AgentSpec into the current workspace.
//
// Description:
//
// ## Operation description
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses it, and imports the AgentSpec into the current workspace.
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
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses it, and imports the AgentSpec into the current workspace.
//
// Description:
//
// ## Operation description
//
// Retrieves an uploaded AgentSpec ZIP package from OSS, parses it, and imports the AgentSpec into the current workspace.
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
// Verifies the ACR RAM authorization of a workspace.
//
// Description:
//
// ACR and user OSS mounts share the same workspace role. When no existing role is found, the system uses AliyunAgentCoreWorkspace-{workspaceId}. When a valid legacy OSS mount role exists, the system reuses it and only appends permissions without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not rebuilt or renamed because of new naming templates. This operation performs a read-only check to verify whether the custom policies corresponding to the target instance, namespace, and repository are attached to the shared role, and checks the workspace prerequisite status. It does not create execution versions or trigger configuration delivery. AUTHORIZED only indicates that the custom policy for the target name is attached. It does not parse policy content or simulate final permissions. ELIGIBLE does not indicate that RuntimeManager capacity or template admission is complete, nor does it require that ACR configurations have been delivered. Creating or updating an Agent re-verifies permissions and processes ACR dependencies on the backend. Configuration waiting, failures, and retries are reported through Agent details.
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
// Verifies the ACR RAM authorization of a workspace.
//
// Description:
//
// ACR and user OSS mounts share the same workspace role. When no existing role is found, the system uses AliyunAgentCoreWorkspace-{workspaceId}. When a valid legacy OSS mount role exists, the system reuses it and only appends permissions without overwriting existing OSS policies. Previously pinned shared roles continue to be reused and are not rebuilt or renamed because of new naming templates. This operation performs a read-only check to verify whether the custom policies corresponding to the target instance, namespace, and repository are attached to the shared role, and checks the workspace prerequisite status. It does not create execution versions or trigger configuration delivery. AUTHORIZED only indicates that the custom policy for the target name is attached. It does not parse policy content or simulate final permissions. ELIGIBLE does not indicate that RuntimeManager capacity or template admission is complete, nor does it require that ACR configurations have been delivered. Creating or updating an Agent re-verifies permissions and processes ACR dependencies on the backend. Configuration waiting, failures, and retries are reported through Agent details.
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
