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
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": dara.String("appflow.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("appflow"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a flow.
//
// Description:
//
// Creates a flow or a flow version.
//
// @param request - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.LaunchStatus) {
		query["LaunchStatus"] = request.LaunchStatus
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow.
//
// Description:
//
// Creates a flow or a flow version.
//
// @param request - CreateFlowRequest
//
// @return CreateFlowResponse
func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - CreateUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserAuthConfigResponse
func (client *Client) CreateUserAuthConfigWithOptions(request *CreateUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfig) {
		query["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthConfigName) {
		query["AuthConfigName"] = request.AuthConfigName
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - CreateUserAuthConfigRequest
//
// @return CreateUserAuthConfigResponse
func (client *Client) CreateUserAuthConfig(request *CreateUserAuthConfigRequest) (_result *CreateUserAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserAuthConfigResponse{}
	_body, _err := client.CreateUserAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a connection flow.
//
// @param request - DeleteFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection flow.
//
// @param request - DeleteFlowRequest
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - DeleteUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserAuthConfigResponse
func (client *Client) DeleteUserAuthConfigWithOptions(request *DeleteUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - DeleteUserAuthConfigRequest
//
// @return DeleteUserAuthConfigResponse
func (client *Client) DeleteUserAuthConfig(request *DeleteUserAuthConfigRequest) (_result *DeleteUserAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserAuthConfigResponse{}
	_body, _err := client.DeleteUserAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a flow.
//
// @param request - DisableFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableFlowResponse
func (client *Client) DisableFlowWithOptions(request *DisableFlowRequest, runtime *dara.RuntimeOptions) (_result *DisableFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a flow.
//
// @param request - DisableFlowRequest
//
// @return DisableFlowResponse
func (client *Client) DisableFlow(request *DisableFlowRequest) (_result *DisableFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableFlowResponse{}
	_body, _err := client.DisableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a flow.
//
// @param request - EnableFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFlowResponse
func (client *Client) EnableFlowWithOptions(request *EnableFlowRequest, runtime *dara.RuntimeOptions) (_result *EnableFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a flow.
//
// @param request - EnableFlowRequest
//
// @return EnableFlowResponse
func (client *Client) EnableFlow(request *EnableFlowRequest) (_result *EnableFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableFlowResponse{}
	_body, _err := client.EnableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a logon session token.
//
// @param request - GenerateUserSessionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionTokenWithOptions(request *GenerateUserSessionTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateUserSessionTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.ExpireSecond) {
		query["ExpireSecond"] = request.ExpireSecond
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.IntegrateId) {
		query["IntegrateId"] = request.IntegrateId
	}

	if !dara.IsNil(request.UserAvatar) {
		query["UserAvatar"] = request.UserAvatar
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUserSessionToken"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a logon session token.
//
// @param request - GenerateUserSessionTokenRequest
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionToken(request *GenerateUserSessionTokenRequest) (_result *GenerateUserSessionTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.GenerateUserSessionTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a flow.
//
// @param request - GetFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowResponse
func (client *Client) GetFlowWithOptions(request *GetFlowRequest, runtime *dara.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a flow.
//
// @param request - GetFlowRequest
//
// @return GetFlowResponse
func (client *Client) GetFlow(request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a user authentication credential.
//
// Description:
//
// This operation gets the details of a specified credential.
//
// @param request - GetUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAuthConfigResponse
func (client *Client) GetUserAuthConfigWithOptions(request *GetUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *GetUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a user authentication credential.
//
// Description:
//
// This operation gets the details of a specified credential.
//
// @param request - GetUserAuthConfigRequest
//
// @return GetUserAuthConfigResponse
func (client *Client) GetUserAuthConfig(request *GetUserAuthConfigRequest) (_result *GetUserAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserAuthConfigResponse{}
	_body, _err := client.GetUserAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes a connector action.
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithSSE(tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions, _yield chan *InvokeActionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeActionWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// Invokes a connector action.
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithOptions(tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) (_result *InvokeActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes a connector action.
//
// @param request - InvokeActionRequest
//
// @return InvokeActionResponse
func (client *Client) InvokeAction(request *InvokeActionRequest) (_result *InvokeActionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvokeActionResponse{}
	_body, _err := client.InvokeActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Launches a flow.
//
// @param request - LaunchFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchFlowResponse
func (client *Client) LaunchFlowWithOptions(request *LaunchFlowRequest, runtime *dara.RuntimeOptions) (_result *LaunchFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LaunchFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LaunchFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Launches a flow.
//
// @param request - LaunchFlowRequest
//
// @return LaunchFlowResponse
func (client *Client) LaunchFlow(request *LaunchFlowRequest) (_result *LaunchFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LaunchFlowResponse{}
	_body, _err := client.LaunchFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of connector flows.
//
// Description:
//
// Creates a connector flow or a connector flow version.
//
// @param request - ListFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowsResponse
func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *dara.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlows"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of connector flows.
//
// Description:
//
// Creates a connector flow or a connector flow version.
//
// @param request - ListFlowsRequest
//
// @return ListFlowsResponse
func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists user authentication credentials.
//
// Description:
//
// This operation retrieves user auth configs that match specified filters.
//
// @param request - ListUserAuthConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAuthConfigsResponse
func (client *Client) ListUserAuthConfigsWithOptions(request *ListUserAuthConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListUserAuthConfigsResponse, _err error) {
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

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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
		Action:      dara.String("ListUserAuthConfigs"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAuthConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists user authentication credentials.
//
// Description:
//
// This operation retrieves user auth configs that match specified filters.
//
// @param request - ListUserAuthConfigsRequest
//
// @return ListUserAuthConfigsResponse
func (client *Client) ListUserAuthConfigs(request *ListUserAuthConfigsRequest) (_result *ListUserAuthConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserAuthConfigsResponse{}
	_body, _err := client.ListUserAuthConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a connection flow.
//
// @param request - UpdateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlowWithOptions(request *UpdateFlowRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a connection flow.
//
// @param request - UpdateFlowRequest
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlow(request *UpdateFlowRequest) (_result *UpdateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlowResponse{}
	_body, _err := client.UpdateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a user authentication credential.
//
// Description:
//
// Updates the configuration of a specific user authentication credential.
//
// @param request - UpdateUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserAuthConfigResponse
func (client *Client) UpdateUserAuthConfigWithOptions(request *UpdateUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfig) {
		query["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.AuthConfigName) {
		query["AuthConfigName"] = request.AuthConfigName
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a user authentication credential.
//
// Description:
//
// Updates the configuration of a specific user authentication credential.
//
// @param request - UpdateUserAuthConfigRequest
//
// @return UpdateUserAuthConfigResponse
func (client *Client) UpdateUserAuthConfig(request *UpdateUserAuthConfigRequest) (_result *UpdateUserAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserAuthConfigResponse{}
	_body, _err := client.UpdateUserAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Withdraws a connection flow.
//
// @param request - WithdrawFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawFlowResponse
func (client *Client) WithdrawFlowWithOptions(request *WithdrawFlowRequest, runtime *dara.RuntimeOptions) (_result *WithdrawFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WithdrawFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Withdraws a connection flow.
//
// @param request - WithdrawFlowRequest
//
// @return WithdrawFlowResponse
func (client *Client) WithdrawFlow(request *WithdrawFlowRequest) (_result *WithdrawFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WithdrawFlowResponse{}
	_body, _err := client.WithdrawFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) invokeActionWithSSE_opYieldFunc(_yield chan *InvokeActionResponse, _yieldErr chan error, tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
